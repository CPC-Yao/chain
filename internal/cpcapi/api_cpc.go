// Copyright 2018 The cpchain authors
// This file is part of the cpchain library.
//
// The cpchain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The cpchain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the cpchain library. If not, see <http://www.gnu.org/licenses/>.

package cpcapi

import (
	"bitbucket.org/cpchain/chain/accounts"
	pb "bitbucket.org/cpchain/chain/api/grpc/v1/common"
	"bitbucket.org/cpchain/chain/api/grpc/v1/cpc"
	"bitbucket.org/cpchain/chain/api/rpc"
	"bitbucket.org/cpchain/chain/core/rawdb"
	"bitbucket.org/cpchain/chain/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// ChainStateReader provides an API to access Ethereum related information.
// It offers only methods that operate on public data that is freely available to anyone.
type ChainStateReader struct {
	b Backend
}

// NewChainStateReader creates a new cpc protocol API.
func NewChainStateReader(b Backend) *ChainStateReader {
	return &ChainStateReader{b: b}
}

// IsPublic if public default
func (c *ChainStateReader) IsPublic() bool {
	return true
}

// Namespace namespace
func (c *ChainStateReader) Namespace() string {
	return "cpc"
}

// RegisterServer register api to grpc
func (c *ChainStateReader) RegisterServer(s *grpc.Server) {
	cpc.RegisterChainStateReaderServer(s, c)
}

// RegisterJsonRpc register api to restfull json
func (c *ChainStateReader) RegisterJsonRpc(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	cpc.RegisterChainStateReaderHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

// GasPrice returns a suggestion for a gas price.
func (c *ChainStateReader) GasPrice(ctx context.Context, req *empty.Empty) (*pb.GasPrice, error) {
	price, err := c.b.SuggestPrice(ctx)
	return &pb.GasPrice{GasPrice: price.Int64()}, err
}

// ProtocolVersion returns the current Ethereum protocol version this node supports
func (c *ChainStateReader) ProtocolVersion(ctx context.Context, req *empty.Empty) (*wrappers.UInt32Value, error) {
	return &wrappers.UInt32Value{Value: uint32(c.b.ProtocolVersion())}, nil
}

// Syncing returns false in case the node is currently not syncing with the network. It can be up to date or has not
// yet received the latest block headers from its pears. In case it is synchronizing:
// - startingBlock: block number this node started to synchronise from
// - currentBlock:  block number this node is currently importing
// - highestBlock:  block number of the highest block header this node has received from peers
// - pulledStates:  number of state entries processed until now
// - knownStates:   number of known state entries that still need to be pulled
func (c *ChainStateReader) Syncing(ctx context.Context, req *empty.Empty) (*pb.SyncingInfo, error) {
	progress := c.b.Downloader().Progress()

	// Return not syncing if the synchronisation already completed
	if progress.CurrentBlock >= progress.HighestBlock {
		return &pb.SyncingInfo{IsSyncing: false}, nil
	}
	// Otherwise gather the block sync stats
	return &pb.SyncingInfo{
		StartBlock:   progress.StartingBlock,
		CurrentBlock: progress.CurrentBlock,
		HighestBlock: progress.HighestBlock,
		PulledStates: progress.PulledStates,
		KnownStates:  progress.KnownStates,
	}, nil
}

// ChainReader provides an API to access the Ethereum blockchain.
// It offers only methods that operate on public data that is freely available to anyone.
type ChainReader struct {
	b Backend
}

// NewChainReader creates a new cpc blockchain API.
func NewChainReader(b Backend) *ChainReader {
	return &ChainReader{b}
}

// GetBlockCount returns the block number of the chain head.
func (c *ChainReader) GetBlockCount(ctx context.Context, req *empty.Empty) (*pb.BlockNumber, error) {
	header, _ := c.b.HeaderByNumber(context.Background(), rpc.LatestBlockNumber) // latest header should always be available
	return &pb.BlockNumber{BlockNumber: header.Number.Int64()}, nil
}

// IsPublic if public default
func (c *ChainReader) IsPublic() bool {
	return true
}

// Namespace namespace
func (c *ChainReader) Namespace() string {
	return "cpc"
}

// RegisterServer register api to grpc
func (c *ChainReader) RegisterServer(s *grpc.Server) {
	cpc.RegisterChainReaderServer(s, c)
}

// RegisterJsonRpc register api to restfull json
func (c *ChainReader) RegisterJsonRpc(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	cpc.RegisterChainReaderHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

// GetBalance returns the amount of wei for the given address in the state of the
// given block number. The rpc.LatestBlockNumber and rpc.PendingBlockNumber meta
// block numbers are also allowed.
func (s *ChainReader) GetBalance(ctx context.Context, req *cpc.ChainReaderRequest) (*pb.Balance, error) {
	state, _, err := s.b.StateAndHeaderByNumber(ctx, rpc.BlockNumber(req.BlockNumber), req.IsFull)
	if state == nil || err != nil {
		return nil, err
	}
	balance := (*hexutil.Big)(state.GetBalance(common.HexToAddress(req.Address)))
	return &pb.Balance{Balance: balance.String()}, state.Error()
}

// newGRPCTransaction returns a transaction that will serialize to the RPC
// representation, with the given location metadata set (if available).
func newGRPCTransaction(tx *types.Transaction, blockHash common.Hash, blockNumber uint64, index uint64) *pb.RpcTransaction {
	var signer types.Signer = types.FrontierSigner{}
	if tx.Protected() {
		signer = types.NewPrivTxSupportEIP155Signer(tx.ChainId())
	}
	from, _ := types.Sender(signer, tx)
	v, r, s := tx.RawSignatureValues()

	result := &pb.RpcTransaction{
		From:     from.Hex(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice().Uint64(),
		Hash:     tx.Hash().Hex(),
		Nonce:    tx.Nonce(),
		To:       &wrappers.StringValue{Value: tx.To().Hex()},
		Input:    tx.Data(),
		Value:    tx.Value().String(),
		V:        v.String(),
		R:        r.String(),
		S:        s.String(),
	}

	if blockHash != (common.Hash{}) {
		result.BlockHash = blockHash.Hex()
		result.BlockNumber = blockNumber
		result.TransactionIndex = index
	}
	return result
}

// newRPCPendingTransaction returns a pending transaction that will serialize to the RPC representation
func newGRPCPendingTransaction(tx *types.Transaction) *pb.RpcTransaction {
	return newGRPCTransaction(tx, common.Hash{}, 0, 0)
}

// newGRPCTransactionFromBlockIndex returns a transaction that will serialize to the RPC representation.
func newGRPCTransactionFromBlockIndex(b *types.Block, index uint64) *pb.RpcTransaction {
	txs := b.Transactions()
	if index >= uint64(len(txs)) {
		return nil
	}
	return newGRPCTransaction(txs[index], b.Hash(), b.NumberU64(), index)
}

// newGRPCTransactionFromBlockHash returns a transaction that will serialize to the RPC representation.
func newGRPCTransactionFromBlockHash(b *types.Block, hash common.Hash) *pb.RpcTransaction {
	for idx, tx := range b.Transactions() {
		if tx.Hash() == hash {
			return newGRPCTransactionFromBlockIndex(b, uint64(idx))
		}
	}
	return nil
}

// GRPCMarshalBlock converts the given block to the RPC output which depends on fullTx. If inclTx is true transactions are
// returned. When fullTx is true the returned block contains full transaction details, otherwise it will only contain
// transaction hashes.
func GRPCMarshalBlock(b *types.Block, inclTx bool, fullTx bool) (*pb.Block, error) {
	head := b.Header() // copies the header once
	block := &pb.Block{
		Number:           head.Number.Uint64(),
		Hash:             b.Hash().Hex(),
		ParentHash:       head.ParentHash.Hex(),
		Nonce:            head.Nonce.Uint64(),
		MixHash:          head.MixHash.Hex(),
		StateRoot:        head.StateRoot.Hex(),
		Miner:            head.Coinbase.Hex(),
		Difficulty:       head.Difficulty.Uint64(),
		Size:             uint64(head.Size()),
		GasLimit:         head.GasLimit,
		GasUsed:          head.GasUsed,
		Timestamp:        head.Time.String(),
		TransactionsRoot: head.TxsRoot.Hex(),
		ReceiptsRoot:     head.ReceiptsRoot.Hex(),
		LogsBloom:        head.LogsBloom.Bytes(),
		ExtraData:        head.Extra,
	}

	if inclTx {
		formatTx := func(tx *types.Transaction) (*pb.RpcTransaction, error) {
			return &pb.RpcTransaction{TxHash: tx.Hash().Hex()}, nil
		}
		if fullTx {
			formatTx = func(tx *types.Transaction) (*pb.RpcTransaction, error) {
				return newGRPCTransactionFromBlockHash(b, tx.Hash()), nil
			}
		}
		txs := b.Transactions()
		block.Transactions = make([]*pb.RpcTransaction, len(txs))
		var err error
		for i, tx := range txs {
			if block.Transactions[i], err = formatTx(tx); err != nil {
				return nil, err
			}
		}
	}

	return block, nil
}

// rpcOutputBlock uses the generalized output filler, then adds the total difficulty field, which requires
// a `PublicBlockchainAPI`.
func (s *ChainReader) grpcOutputBlock(b *types.Block, inclTx bool, fullTx bool) (*pb.Block, error) {
	block, err := GRPCMarshalBlock(b, inclTx, fullTx)
	if err != nil {
		return nil, err
	}
	return block, err
}

// GetBlockByNumber returns the requested block. When blockNr is -1 the chain head is returned. When fullTx is true all
// transactions in the block are returned in full detail, otherwise only the transaction hash is returned.
func (s *ChainReader) GetBlockByNumber(ctx context.Context, req *cpc.ChainReaderRequest) (*pb.Block, error) {
	block, err := s.b.BlockByNumber(ctx, rpc.BlockNumber(req.BlockNumber))
	if block != nil {
		response, err := s.grpcOutputBlock(block, true, req.IsFull)
		if err == nil && rpc.BlockNumber(req.BlockNumber) == rpc.PendingBlockNumber {
			// Pending blocks need to nil out a few fields
			response.Hash = ""
			response.Nonce = 0
			response.Miner = ""
		}
		return response, err
	}
	return &pb.Block{}, err
}

// GetBlockByHash returns the requested block. When fullTx is true all transactions in the block are returned in full
// detail, otherwise only the transaction hash is returned.
func (s *ChainReader) GetBlockByHash(ctx context.Context, req *cpc.ChainReaderRequest) (*pb.Block, error) {
	block, err := s.b.GetBlock(ctx, common.HexToHash(req.BlockHash))
	if block != nil {
		return s.grpcOutputBlock(block, true, req.IsFull)
	}
	return &pb.Block{}, err
}

// GetCode returns the code stored at the given address in the state for the given block number.
func (s *ChainReader) GetCode(ctx context.Context, req *cpc.ChainReaderRequest) (*pb.Code, error) {
	state, _, err := s.b.StateAndHeaderByNumber(ctx, rpc.BlockNumber(req.BlockNumber), false)
	if state == nil || err != nil {
		return &pb.Code{}, err
	}
	code := state.GetCode(common.HexToAddress(req.Address))

	return &pb.Code{Code: code}, state.Error()
}

// TransactionPoolReader exposes methods for the RPC interface
type TransactionPoolReader struct {
	b         Backend
	nonceLock *AddrLocker
}

// NewTransactionPoolReader a new RPC service with methods specific for the transaction pool.
func NewTransactionPoolReader(b Backend, nonceLock *AddrLocker) *TransactionPoolReader {
	return &TransactionPoolReader{b, nonceLock}
}

// IsPublic if public default
func (t *TransactionPoolReader) IsPublic() bool {
	return true
}

// Namespace namespace
func (t *TransactionPoolReader) Namespace() string {
	return "cpc"
}

// RegisterServer register api to grpc
func (t *TransactionPoolReader) RegisterServer(s *grpc.Server) {
	cpc.RegisterTransactionPoolReaderServer(s, t)
}

// RegisterJsonRpc register api to restfull json
func (t *TransactionPoolReader) RegisterJsonRpc(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	cpc.RegisterTransactionPoolReaderHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

// GetBlockTransactionCountByNumber returns the number of transactions in the block with the given block number.
func (t *TransactionPoolReader) GetTransactionCountByBlockNumber(ctx context.Context, blockNr *pb.BlockNumber) (*cpc.TransactionCount, error) {
	if block, _ := t.b.BlockByNumber(ctx, rpc.BlockNumber(blockNr.BlockNumber)); block != nil {
		return &cpc.TransactionCount{TransactionCount: uint64(len(block.Transactions()))}, nil
	}
	return &cpc.TransactionCount{}, nil
}

// GetBlockTransactionCountByHash returns the number of transactions in the block with the given hash.
func (t *TransactionPoolReader) GetTransactionCountByBlockHash(ctx context.Context, blockHash *pb.BlockHash) (*cpc.TransactionCount, error) {
	if block, _ := t.b.GetBlock(ctx, common.HexToHash(blockHash.BlockHash)); block != nil {
		return &cpc.TransactionCount{TransactionCount: uint64(len(block.Transactions()))}, nil
	}
	return &cpc.TransactionCount{}, nil
}

// GetTransactionByBlockNumberAndIndex returns the transaction for the given block number and index.
func (t *TransactionPoolReader) GetTransactionByBlockNumberAndIndex(ctx context.Context, req *cpc.TransactionPoolReaderRequest) (*pb.RpcTransaction, error) {
	if block, _ := t.b.BlockByNumber(ctx, rpc.BlockNumber(req.BlockNumber)); block != nil {
		return newGRPCTransactionFromBlockIndex(block, req.Index), nil
	}
	return &pb.RpcTransaction{}, nil
}

// GetTransactionByBlockHashAndIndex returns the transaction for the given block hash and index.
func (t *TransactionPoolReader) GetTransactionByBlockHashAndIndex(ctx context.Context, req *cpc.TransactionPoolReaderRequest) (*pb.RpcTransaction, error) {
	if block, _ := t.b.GetBlock(ctx, common.HexToHash(req.BlockHash)); block != nil {
		return newGRPCTransactionFromBlockIndex(block, req.Index), nil
	}
	return &pb.RpcTransaction{}, nil
}

// GetRawTransactionByBlockNumberAndIndex returns the bytes of the transaction for the given block number and index.
func (t *TransactionPoolReader) GetRawTransactionByBlockNumberAndIndex(ctx context.Context, req *cpc.TransactionPoolReaderRequest) (*cpc.RawTransaction, error) {
	var out cpc.RawTransaction
	if block, _ := t.b.BlockByNumber(ctx, rpc.BlockNumber(req.BlockNumber)); block != nil {
		raw := newRPCRawTransactionFromBlockIndex(block, req.Index)
		out.RawTransaction = raw
	}
	return &out, nil
}

// GetRawTransactionByBlockHashAndIndex returns the bytes of the transaction for the given block hash and index.
func (t *TransactionPoolReader) GetRawTransactionByBlockHashAndIndex(ctx context.Context, req *cpc.TransactionPoolReaderRequest) (*cpc.RawTransaction, error) {
	var out cpc.RawTransaction
	if block, _ := t.b.GetBlock(ctx, common.HexToHash(req.BlockHash)); block != nil {
		raw := newRPCRawTransactionFromBlockIndex(block, req.Index)
		out.RawTransaction = raw
	}
	return &out, nil
}

// GetTransactionCount returns the number of transactions the given address has sent for the given block number
func (t *TransactionPoolReader) GetTransactionCount(ctx context.Context, req *cpc.TransactionPoolReaderRequest) (*cpc.TransactionCount, error) {
	state, _, err := t.b.StateAndHeaderByNumber(ctx, rpc.BlockNumber(req.BlockNumber), false)
	if state == nil || err != nil {
		return nil, err
	}
	nonce := state.GetNonce(common.HexToAddress(req.Address))
	return &cpc.TransactionCount{TransactionCount: nonce}, nil
}

// GetTransactionByHash returns the transaction for the given hash
func (t *TransactionPoolReader) GetTransactionByHash(ctx context.Context, txHash *cpc.TransactionHash) (*pb.RpcTransaction, error) {
	// Try to return an already finalized transaction
	if tx, blockHash, blockNumber, index := rawdb.ReadTransaction(t.b.ChainDb(), common.HexToHash(txHash.TransactionHash)); tx != nil {
		return newGRPCTransaction(tx, blockHash, blockNumber, index), nil
	}
	// No finalized transaction, try to retrieve it from the pool
	if tx := t.b.GetPoolTransaction(common.HexToHash(txHash.TransactionHash)); tx != nil {
		return newGRPCPendingTransaction(tx), nil
	}
	// Transaction unknown, return as such
	return &pb.RpcTransaction{}, nil
}

// GetRawTransactionByHash returns the bytes of the transaction for the given hash.
func (t *TransactionPoolReader) GetRawTransactionByHash(ctx context.Context, hash common.Hash) (hexutil.Bytes, error) {
	var tx *types.Transaction

	// Retrieve a finalized transaction, or a pooled otherwise
	if tx, _, _, _ = rawdb.ReadTransaction(t.b.ChainDb(), hash); tx == nil {
		if tx = t.b.GetPoolTransaction(hash); tx == nil {
			// Transaction not found anywhere, abort
			return nil, nil
		}
	}
	// Serialize to RLP and return
	return rlp.EncodeToBytes(tx)
}

// GetTransactionReceipt returns the transaction receipt for the given transaction hash.
func (t *TransactionPoolReader) GetTransactionReceipt(ctx context.Context, txHash *cpc.TransactionHash) (*pb.Receipt, error) {
	tx, blockHash, blockNumber, index := rawdb.ReadTransaction(t.b.ChainDb(), common.HexToHash(txHash.TransactionHash))
	if tx == nil {
		return &pb.Receipt{}, nil
	}

	var receipt *types.Receipt
	if tx.IsPrivate() {
		receipt, _ = t.b.GetPrivateReceipt(ctx, tx.Hash())
		if receipt == nil {
			return nil, nil
		}
	} else {
		receipts, err := t.b.GetReceipts(ctx, blockHash)
		if err != nil {
			return nil, err
		}
		if len(receipts) <= int(index) {
			return nil, nil
		}
		receipt = receipts[index]
	}

	var signer types.Signer = types.FrontierSigner{}
	if tx.Protected() {
		signer = types.NewPrivTxSupportEIP155Signer(tx.ChainId())
	}
	from, _ := types.Sender(signer, tx)

	result := &pb.Receipt{
		BlockHash:        blockHash.Hex(),
		BlockNumber:      blockNumber,
		TransactionHash:  txHash.TransactionHash,
		TransactionIndex: index,
		From:             from.Hex(),
		To:               tx.To().Hex(),
		GasUsed:          receipt.GasUsed,
		Logs:             make([]*pb.Log, len(receipt.Logs)),
		LogsBloom:        receipt.Bloom.Bytes(),
	}

	for _, log := range receipt.Logs {
		l := &pb.Log{
			Removed:     log.Removed,
			BlockNumber: log.BlockNumber,
			TxIndex:     uint64(log.TxIndex),
			BlockHash:   log.BlockHash.Hex(),
			Index:       uint32(log.Index),
			Address:     log.Address.Hex(),
			TxHash:      log.TxHash.Hex(),
			Data:        log.Data,
			Topics:      make([]string, len(log.Topics)),
		}
		for _, t := range log.Topics {
			l.Topics = append(l.Topics, t.Hex())
		}
		result.Logs = append(result.Logs, l)
	}

	// Assign receipt status or post state.
	if len(receipt.PostState) > 0 {
		result.Root = receipt.PostState
	} else {
		result.Status = receipt.Status
	}
	if receipt.Logs == nil {
		result.Logs = []*pb.Log{}
	}
	// If the ContractAddress is 20 0x0 bytes, assume it is not a contract creation
	if receipt.ContractAddress != (common.Address{}) {
		result.ContractAddress = receipt.ContractAddress.Hex()
	}
	return result, nil
}

// AccountReader provides an API to access accounts managed by this node.
// It offers only methods that can retrieve accounts.
type AccountReader struct {
	am *accounts.Manager
}

// NewAccountReader creates a new AccountReader.
func NewAccountReader(am *accounts.Manager) *AccountReader {
	return &AccountReader{am: am}
}

// IsPublic if public default
func (c *AccountReader) IsPublic() bool {
	return true
}

// Namespace namespace
func (c *AccountReader) Namespace() string {
	return "cpc"
}

// RegisterServer register api to grpc
func (c *AccountReader) RegisterServer(s *grpc.Server) {
	cpc.RegisterAccountReaderServer(s, c)
}

// RegisterJsonRpc register api to restfull json
func (c *AccountReader) RegisterJsonRpc(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
	cpc.RegisterAccountReaderHandlerFromEndpoint(ctx, mux, endpoint, opts)
}

// Accounts returns the collection of accounts this node manages
func (s *AccountReader) Accounts(ctx context.Context, req *empty.Empty) (*pb.Addresses, error) {
	addresses := &pb.Addresses{Addresses: make([]string, 0, 0)}
	for _, wallet := range s.am.Wallets() {
		for _, account := range wallet.Accounts() {
			addresses.Addresses = append(addresses.Addresses, account.Address.Hex())
		}
	}
	return addresses, nil
}