// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"errors"

	"crypto/rsa"

	"bitbucket.org/cpchain/chain/consensus"
	"bitbucket.org/cpchain/chain/consensus/misc"
	"bitbucket.org/cpchain/chain/core/state"
	"bitbucket.org/cpchain/chain/core/types"
	"bitbucket.org/cpchain/chain/core/vm"
	"bitbucket.org/cpchain/chain/crypto"
	"bitbucket.org/cpchain/chain/ethdb"
	"bitbucket.org/cpchain/chain/params"
	"bitbucket.org/cpchain/chain/private"
	"github.com/ethereum/go-ethereum/common"
)

var (
	RemoteDBAbsenceError = errors.New("RemoteDB is not set, no capacibility of processing private transaction.")
	NoPermissionError    = errors.New("The node doesn't have the permission/responsibility to process the private tx.")
	RSAKeyAbsenceError   = errors.New("RSA private key is not set, no capacibility of processing private transaction.")
)

// StateProcessor is a basic Processor, which takes care of transitioning
// state from one point to another.
//
// StateProcessor implements Processor.
type StateProcessor struct {
	config *params.ChainConfig // Chain configuration options
	bc     *BlockChain         // Canonical block chain
	engine consensus.Engine    // Consensus engine used for block rewards
}

// NewStateProcessor initialises a new StateProcessor.
func NewStateProcessor(config *params.ChainConfig, bc *BlockChain, engine consensus.Engine) *StateProcessor {
	return &StateProcessor{
		config: config,
		bc:     bc,
		engine: engine,
	}
}

// Process processes the state changes according to the Ethereum rules by running
// the transaction messages using the pubStateDB and applying any rewards to both
// the processor (coinbase) and any included uncles.
//
// Process returns the public receipts, private receipts(if have) and logs accumulated during the process and
// returns the amount of gas that was used in the process. If any of the
// transactions failed to execute due to insufficient gas it will return an error.
func (p *StateProcessor) Process(block *types.Block, statedb *state.StateDB, statePrivDB *state.StateDB,
	remoteDB ethdb.RemoteDatabase, cfg vm.Config, rsaPrivKey *rsa.PrivateKey) (types.Receipts, types.Receipts, []*types.Log, uint64, error) {
	var (
		pubReceipts  types.Receipts
		privReceipts types.Receipts
		usedGas      = new(uint64)
		header       = block.Header()
		allLogs      []*types.Log
		gp           = new(GasPool).AddGas(block.GasLimit())
	)
	// Mutate the the block and state according to any hard-fork specs
	if p.config.DAOForkSupport && p.config.DAOForkBlock != nil && p.config.DAOForkBlock.Cmp(block.Number()) == 0 {
		misc.ApplyDAOHardFork(statedb)
	}
	// Iterate over and process the individual transactions
	for i, tx := range block.Transactions() {
		statedb.Prepare(tx.Hash(), block.Hash(), i)
		statePrivDB.Prepare(tx.Hash(), block.Hash(), i)
		pubReceipt, privReceipt, _, err := ApplyTransaction(p.config, p.bc, nil, gp, statedb, statePrivDB, remoteDB, header, tx,
			usedGas, cfg, rsaPrivKey)
		if err != nil {
			return nil, nil, nil, 0, err
		}
		pubReceipts = append(pubReceipts, pubReceipt)
		if privReceipt != nil {
			privReceipts = append(privReceipts, privReceipt)
		}
		allLogs = append(allLogs, pubReceipt.Logs...) // not include private receipt's logs.
		// TODO: if need to add private receipt's logs to allLogs variable.
	}
	// Finalize the block, applying any consensus engine specific extras (e.g. block rewards)
	p.engine.Finalize(p.bc, header, statedb, block.Transactions(), block.Uncles(), pubReceipts)

	// TODO: if return private logs separately or merge them together as a whole logs collection?
	return pubReceipts, privReceipts, allLogs, *usedGas, nil
}

// ApplyTransaction attempts to apply a transaction to the given state database
// and uses the input parameters for its environment. It returns the receipt
// for the transaction, gas used and an error if the transaction failed,
// indicating the block was invalid.
func ApplyTransaction(config *params.ChainConfig, bc ChainContext, author *common.Address, gp *GasPool, pubStateDb *state.StateDB,
	privateStateDb *state.StateDB, remoteDB ethdb.RemoteDatabase, header *types.Header, tx *types.Transaction, usedGas *uint64,
	cfg vm.Config, rsaPrivKey *rsa.PrivateKey) (*types.Receipt, *types.Receipt, uint64, error) {
	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return nil, nil, 0, err
	}

	// For private tx, its payload is a replacement which cannot be executed as normal tx payload, thus set it to be empty to skip execution.
	// This around of execution generates stuff stored in public blockchain.
	if tx.IsPrivate() {
		msg.SetData([]byte{})
	}

	// Create a new context to be used in the EVM environment
	context := NewEVMContext(msg, header, bc, author)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(context, pubStateDb, config, cfg)
	// Apply the transaction to the current state (included in the env)
	_, gas, failed, err := ApplyMessage(vmenv, msg, gp)
	if err != nil {
		return nil, nil, 0, err
	}
	// Update the state with pending changes
	var root []byte
	// TODO: investigate whether root is empty and everything seems good when IsByzantium returns false.
	if config.IsByzantium(header.Number) {
		pubStateDb.Finalise(true)
	} else {
		root = pubStateDb.IntermediateRoot(config.IsEIP158(header.Number)).Bytes()
	}
	*usedGas += gas

	// Create a new pubReceipt for the transaction, storing the intermediate root and gas used by the tx
	// based on the eip phase, we're passing wether the root touch-delete accounts.
	pubReceipt := types.NewReceipt(root, failed, *usedGas)
	pubReceipt.TxHash = tx.Hash()
	pubReceipt.GasUsed = gas
	// if the transaction created a contract, store the creation address in the pubReceipt.
	if msg.To() == nil {
		pubReceipt.ContractAddress = crypto.CreateAddress(vmenv.Context.Origin, tx.Nonce())
	}
	// Set the pubReceipt logs and create a bloom for filtering
	pubReceipt.Logs = pubStateDb.GetLogs(tx.Hash())
	pubReceipt.Bloom = types.CreateBloom(types.Receipts{pubReceipt})

	var privReceipt *types.Receipt
	// For private tx, it should process its real private tx payload in participant's node.
	if tx.IsPrivate() {
		privReceipt, _ = tryApplyPrivateTx(config, bc, author, gp, privateStateDb, remoteDB, header, tx, cfg, rsaPrivKey)
	}

	return pubReceipt, privReceipt, gas, err
}

// applyPrivateTx attempts to apply a private transaction to the given state database
func tryApplyPrivateTx(config *params.ChainConfig, bc ChainContext, author *common.Address, gp *GasPool, privateStateDb *state.StateDB,
	remoteDB ethdb.RemoteDatabase, header *types.Header, tx *types.Transaction, cfg vm.Config, rsaPrivKey *rsa.PrivateKey) (*types.Receipt, error) {
	msg, err := tx.AsMessage(types.MakeSigner(config, header.Number))
	if err != nil {
		return nil, err
	}

	if remoteDB == nil {
		return nil, RemoteDBAbsenceError
	}

	if rsaPrivKey == nil {
		return nil, RSAKeyAbsenceError
	}

	pub := rsaPrivKey.PublicKey
	if err != nil {
		return nil, err
	}

	payload, hasPermission, _ := private.RetrieveAndDecryptPayload(tx.Data(), tx.Nonce(), remoteDB, &pub, rsaPrivKey)
	if !hasPermission {
		return nil, NoPermissionError
	}

	// Replace with the real payload decrypted from remote database.
	msg.SetData(payload)
	msg.GasPrice().SetUint64(0)
	privateStateDb.SetNonce(msg.From(), msg.Nonce())

	// Create a new context to be used in the EVM environment
	context := NewEVMContext(msg, header, bc, author)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(context, privateStateDb, config, cfg)
	// Apply the transaction to the current state (included in the env)
	_, _, failed, err := ApplyMessage(vmenv, msg, gp)
	if err != nil {
		return nil, err
	}

	root := privateStateDb.IntermediateRoot(true).Bytes()

	// Create a new receipt for the transaction, storing the intermediate root and gas used by the tx
	// based on the eip phase, we're passing wether the root touch-delete accounts.
	receipt := types.NewReceipt(root, failed, 0)
	receipt.TxHash = tx.Hash()
	receipt.GasUsed = 0 // for private tx, consume no gas.
	// if the transaction created a contract, store the creation address in the receipt.
	if msg.To() == nil {
		receipt.ContractAddress = crypto.CreateAddress(vmenv.Context.Origin, tx.Nonce())
	}
	// Set the receipt logs and create a bloom for filtering
	receipt.Logs = privateStateDb.GetLogs(tx.Hash())
	receipt.Bloom = types.CreateBloom(types.Receipts{receipt})

	return receipt, err
}
