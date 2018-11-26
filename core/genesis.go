// Copyright 2018 The cpchain authors
// Copyright 2014 The go-ethereum Authors

package core

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"bitbucket.org/cpchain/chain/commons/log"
	"bitbucket.org/cpchain/chain/configs"
	"bitbucket.org/cpchain/chain/core/rawdb"
	"bitbucket.org/cpchain/chain/core/state"
	"bitbucket.org/cpchain/chain/database"
	"bitbucket.org/cpchain/chain/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/rlp"
)

//go:generate gencodec -type Genesis -formats json,toml -field-override genesisSpecMarshaling -out gen_genesis.go
//go:generate gencodec -type GenesisAccount -formats json,toml -field-override genesisAccountMarshaling -out gen_genesis_account.go

var (
	errGenesisNoConfig   = errors.New("genesis has no chain configuration")
	errGenesisNoExist    = errors.New("genesis block does not exist")
	errGenesisCfgNoExist = errors.New("genesis block configuration does not exist")
)

// Genesis specifies the header fields, state of a genesis block. It also defines hard
// fork switch-over blocks through the chain configuration.
type Genesis struct {
	Config     *configs.ChainConfig `json:"config"     toml:"config"`
	Nonce      uint64               `json:"nonce"      toml:"nonce"`
	Timestamp  uint64               `json:"timestamp"  toml:"timestamp"`
	ExtraData  []byte               `json:"extraData"  toml:"extraData"`
	GasLimit   uint64               `json:"gasLimit"   toml:"gasLimit"   gencodec:"required"`
	Difficulty *big.Int             `json:"difficulty" toml:"difficulty" gencodec:"required"`
	Mixhash    common.Hash          `json:"mixHash"    toml:"mixHash"`
	Coinbase   common.Address       `json:"coinbase"   toml:"coinbase"`
	Alloc      GenesisAlloc         `json:"alloc"      toml:"alloc"      gencodec:"required"`

	// These fields are used for consensus tests. Please don't use them
	// in actual genesis blocks.
	Number     uint64         `json:"number"     toml:"number"`
	GasUsed    uint64         `json:"gasUsed"    toml:"gasUsed"`
	ParentHash common.Hash    `json:"parentHash" toml:"parentHash"`
	Dpor       types.DporSnap `json:"dpor"       toml:"dpor"`
}

// GenesisAlloc specifies the initial state that is part of the genesis block.
type GenesisAlloc map[common.Address]GenesisAccount

func (ga *GenesisAlloc) UnmarshalJSON(data []byte) error {
	m := make(map[common.UnprefixedAddress]GenesisAccount)
	if err := json.Unmarshal(data, &m); err != nil {
		return err
	}
	*ga = make(GenesisAlloc)
	for addr, a := range m {
		(*ga)[common.Address(addr)] = a
	}
	return nil
}

// GenesisAccount is an account in the state of the genesis block.
type GenesisAccount struct {
	Code       []byte                      `json:"code,omitempty"      toml:"code,omitempty"`
	Storage    map[common.Hash]common.Hash `json:"storage,omitempty"   toml:"storage,omitempty"`
	Balance    *big.Int                    `json:"balance"             toml:"balance"             gencodec:"required"`
	Nonce      uint64                      `json:"nonce,omitempty"     toml:"nonce,omitempty"`
	PrivateKey []byte                      `json:"secretKey,omitempty" toml:"secretKey,omitempty"` // for tests
}

// field type overrides for gencodec
type genesisSpecMarshaling struct {
	Nonce      math.HexOrDecimal64
	Timestamp  math.HexOrDecimal64
	ExtraData  hexutil.Bytes
	GasLimit   math.HexOrDecimal64
	GasUsed    math.HexOrDecimal64
	Number     math.HexOrDecimal64
	Difficulty *math.HexOrDecimal256
	Mixhash    marshalHash
	Alloc      map[common.UnprefixedAddress]GenesisAccount
}

type genesisAccountMarshaling struct {
	Code       hexutil.Bytes
	Balance    *math.HexOrDecimal256
	Nonce      math.HexOrDecimal64
	Storage    map[marshalHash]marshalHash
	PrivateKey hexutil.Bytes
}

// marshalHash represents a 256 bit byte array, but allows less than 256 bits when
// unmarshaling from hex.
type marshalHash common.Hash

func (h *marshalHash) UnmarshalText(text []byte) error {
	text = bytes.TrimPrefix(text, []byte("0x"))
	if len(text) > 64 {
		return fmt.Errorf("too many hex characters in storage key/value %q", text)
	}
	offset := len(h) - len(text)/2 // pad on the left
	if _, err := hex.Decode(h[offset:], text); err != nil {
		fmt.Println(err)
		return fmt.Errorf("invalid hex storage key/value %q", text)
	}
	return nil
}

func (h marshalHash) MarshalText() ([]byte, error) {
	return hexutil.Bytes(h[:]).MarshalText()
}

// GenesisMismatchError is raised when trying to overwrite an existing
// genesis block with an incompatible one.
type GenesisMismatchError struct {
	Stored, New common.Hash
}

func (e *GenesisMismatchError) Error() string {
	return fmt.Sprintf("database already contains an incompatible genesis block (have %x, new %x)", e.Stored[:8], e.New[:8])
}

// SetupGenesisBlock writes or updates the genesis block in db.
// The block that will be used is:
//
//                          genesis == nil       genesis != nil
//                       +------------------------------------------
//     db has no genesis |  main-net default  |  genesis
//     db has genesis    |  from DB           |  genesis (must match)
//
// The stored chain configuration will be updated if it is compatible (i.e. does not
// specify a fork block below the local head block). In case of a conflict, the
// error is a *params.ConfigCompatError and the new, unwritten config is returned.
//
// The returned chain configuration is never nil.
func SetupGenesisBlock(db database.Database, genesis *Genesis) (*configs.ChainConfig, common.Hash, error) {
	if genesis != nil && genesis.Config == nil {
		return configs.DefaultChainConfig, common.Hash{}, errGenesisNoConfig
	}

	stored := rawdb.ReadCanonicalHash(db, 0)
	if (stored == common.Hash{}) {
		// Just commit the new block if there is no stored genesis block.
		if genesis == nil {
			log.Info("Writing default main-net genesis block")
			genesis = DefaultGenesisBlock()
		} else {
			log.Info("Writing custom genesis block")
		}
		block, err := genesis.Commit(db)
		return genesis.Config, block.Hash(), err
	} else {
		// Get the existing chain configuration.
		storedCfg := rawdb.ReadChainConfig(db, stored)
		newCfg := genesis.configOrDefault(stored)
		var finalCfg *configs.ChainConfig
		if genesis != nil {
			// Check whether the genesis block is already written.
			hash := genesis.ToBlock(nil).Hash()
			if hash != stored {
				return genesis.Config, hash, &GenesisMismatchError{stored, hash}
			}
			finalCfg = updateChainConfig(storedCfg, newCfg, db, stored)
		} else {
			// Special case: don't change the existing config of a non-mainnet chain if no new
			// config is supplied. These chains would get AllProtocolChanges (and a compat error)
			// if we just continued here.
			if stored != MainnetGenesisHash {
				return storedCfg, stored, nil
			} else {
				finalCfg = updateChainConfig(storedCfg, newCfg, db, stored)
			}
		}
		return finalCfg, stored, nil
	}
}

func updateChainConfig(storedcfg *configs.ChainConfig, newcfg *configs.ChainConfig, db database.Database, stored common.Hash) *configs.ChainConfig {
	if storedcfg == nil {
		log.Warn("Found genesis block without chain config")
	}
	// NOTICE: in future we may need to check the compatibility of old and new configurations.
	rawdb.WriteChainConfig(db, stored, newcfg)
	return newcfg
}

// OpenGenesisBlock opens genesis block and returns its chain configuration and hash.
// Return errors when genesis block not exist or genesis block configuration not exist.
func OpenGenesisBlock(db database.Database) (*configs.ChainConfig, common.Hash, error) {
	// the hash of the stored block
	stored := rawdb.ReadCanonicalHash(db, 0)
	if (stored == common.Hash{}) {
		return nil, common.Hash{}, errGenesisNoExist
	}
	storedcfg := rawdb.ReadChainConfig(db, stored)
	if storedcfg != nil {
		return storedcfg, stored, nil
	} else {
		return nil, stored, errGenesisCfgNoExist
	}
}

func (g *Genesis) configOrDefault(ghash common.Hash) *configs.ChainConfig {
	switch {
	case g != nil:
		return g.Config
	case ghash == MainnetGenesisHash:
		return configs.MainnetChainConfig
	default:
		return configs.DefaultChainConfig
	}
}

// ToBlock creates the genesis block and writes state of a genesis specification
// to the given database (or discards it if nil).
func (g *Genesis) ToBlock(db database.Database) *types.Block {
	if db == nil {
		db = database.NewMemDatabase()
	}
	statedb, _ := state.New(common.Hash{}, state.NewDatabase(db))
	for addr, account := range g.Alloc {
		statedb.AddBalance(addr, account.Balance)
		statedb.SetCode(addr, account.Code)
		statedb.SetNonce(addr, account.Nonce)
		for key, value := range account.Storage {
			statedb.SetState(addr, key, value)
		}
	}
	root := statedb.IntermediateRoot(false)
	head := &types.Header{
		Number:     new(big.Int).SetUint64(g.Number),
		Nonce:      types.EncodeNonce(g.Nonce),
		Time:       new(big.Int).SetUint64(g.Timestamp),
		ParentHash: g.ParentHash,
		Extra:      g.ExtraData,
		GasLimit:   g.GasLimit,
		GasUsed:    g.GasUsed,
		Difficulty: g.Difficulty,
		MixHash:    g.Mixhash,
		Coinbase:   g.Coinbase,
		StateRoot:  root,
		Dpor:       g.Dpor,
	}
	if g.GasLimit == 0 {
		head.GasLimit = configs.GenesisGasLimit
	}
	if g.Difficulty == nil {
		head.Difficulty = configs.GenesisDifficulty
	}
	if _, err := statedb.Commit(false); err != nil {
		log.Error("Error in genesis", "error", err)
	}
	if err := statedb.Database().TrieDB().Commit(root, true); err != nil {
		log.Error("Error in genesis", "error", err)
	}
	return types.NewBlock(head, nil, nil)
}

// Commit writes the block and state of a genesis specification to the database.
// The block is committed as the canonical head block.
func (g *Genesis) Commit(db database.Database) (*types.Block, error) {
	block := g.ToBlock(db)
	if block.Number().Sign() != 0 {
		return nil, fmt.Errorf("can't commit genesis block with number > 0")
	}
	rawdb.WriteTd(db, block.Hash(), block.NumberU64(), g.Difficulty)
	rawdb.WriteBlock(db, block)
	rawdb.WriteReceipts(db, block.Hash(), block.NumberU64(), nil)
	rawdb.WriteCanonicalHash(db, block.Hash(), block.NumberU64())
	rawdb.WriteHeadBlockHash(db, block.Hash())
	rawdb.WriteHeadHeaderHash(db, block.Hash())

	config := g.Config
	if config == nil {
		config = configs.DefaultChainConfig
	}
	rawdb.WriteChainConfig(db, block.Hash(), config)
	return block, nil
}

// MustCommit writes the genesis block and state to db, panicking on error.
// The block is committed as the canonical head block.
func (g *Genesis) MustCommit(db database.Database) *types.Block {
	block, err := g.Commit(db)
	if err != nil {
		panic(err)
	}
	return block
}

// GenesisBlockForTesting creates and writes a block in which addr has the given wei balance.
func GenesisBlockForTesting(db database.Database, addr common.Address, balance *big.Int) *types.Block {
	g := DefaultGenesisBlock()
	g.Alloc = GenesisAlloc{addr: {Balance: balance}}
	return g.MustCommit(db)
}

// Genesis hashes to enforce below configs on.
var MainnetGenesisHash = common.HexToHash("0xb0d9b86767138ce7b1e8fd898642d6ec2aca73f7393be0deb496f95acfebc531")

// DefaultGenesisBlock returns the cpchain main net genesis block.
func DefaultGenesisBlock() *Genesis {
	return &Genesis{
		Config:    configs.MainnetChainConfig,
		Timestamp: 1492009146,
		ExtraData: hexutil.MustDecode("0x0000000000000000000000000000000000000000000000000000000000000000"),
		GasLimit:  4700000,
		// GasLimit:   1000000000,
		Difficulty: big.NewInt(1),
		Alloc: map[common.Address]GenesisAccount{
			common.HexToAddress("0xe94b7b6c5a0e526a4d97f9768ad6097bde25c62a"): {Balance: big.NewInt(math.MaxInt64)},
			common.HexToAddress("0xc05302acebd0730e3a18a058d7d1cb1204c4a092"): {Balance: big.NewInt(math.MaxInt64)},
			common.HexToAddress("0xef3dd127de235f15ffb4fc0d71469d1339df6465"): {Balance: big.NewInt(math.MaxInt64)},

			common.HexToAddress("0x3a18598184ef84198db90c28fdfdfdf56544f747"): {Balance: big.NewInt(math.MaxInt64)},
			common.HexToAddress("0x6e31e5b68a98dcd17264bd1ba547d0b3e874da1e"): {Balance: big.NewInt(math.MaxInt64)},

			common.HexToAddress("0x22a672eab2b1a3ff3ed91563205a56ca5a560e08"): {Balance: big.NewInt(math.MaxInt64)},
			common.HexToAddress("0x7b2f052a372951d02798853e39ee56c895109992"): {Balance: big.NewInt(math.MaxInt64)},
			common.HexToAddress("0x2f0176cc3a8617b6ddea6a501028fa4c6fc25ca1"): {Balance: big.NewInt(math.MaxInt64)},
			common.HexToAddress("0xe4d51117832e84f1d082e9fc12439b771a57e7b2"): {Balance: big.NewInt(math.MaxInt64)},
			common.HexToAddress("0x32bd7c33bb5060a85f361caf20c0bda9075c5d51"): {Balance: big.NewInt(math.MaxInt64)},

			common.HexToAddress("0x0000000000000000000000000000000000000000"): {Balance: big.NewInt(0x00000000000000000)},
			common.HexToAddress("0x0000000000000000000000000000000000000001"): {Balance: big.NewInt(0x00000000000000000)},
			common.HexToAddress("0x0000000000000000000000000000000000000002"): {Balance: big.NewInt(0x00000000000000000)},
			common.HexToAddress("0x00000000000000000000000000000000000000ff"): {Balance: big.NewInt(0x00000000000000000)},
		},
		Dpor: types.DporSnap{
			Proposers: []common.Address{
				common.HexToAddress("0xc05302acebd0730e3a18a058d7d1cb1204c4a092"),
				common.HexToAddress("0xe94b7b6c5a0e526a4d97f9768ad6097bde25c62a"),
				common.HexToAddress("0xef3dd127de235f15ffb4fc0d71469d1339df6465"),
				common.HexToAddress("0x6e31e5b68a98dcd17264bd1ba547d0b3e874da1e"),
			},
			Seal:       types.DporSignature{},
			Sigs:       make([]types.DporSignature, 4),
			Validators: []common.Address{},
		},
	}
}

// DeveloperGenesisBlock returns the 'cpchain --dev' genesis block. Note, this must
// be seeded with the
func DeveloperGenesisBlock(period uint64, faucet common.Address) *Genesis {
	// Override the default period to the user requested one
	config := *configs.DefaultChainConfig
	config.Dpor.Period = period

	// Assemble and return the genesis with the precompiles and faucet pre-funded
	return &Genesis{
		Config:     &config,
		Timestamp:  1492009146,
		ExtraData:  make([]byte, 32),
		GasLimit:   4700000,
		Difficulty: big.NewInt(1),
		Alloc: map[common.Address]GenesisAccount{
			common.BytesToAddress([]byte{1}): {Balance: big.NewInt(1)}, // ECRecover
			common.BytesToAddress([]byte{2}): {Balance: big.NewInt(1)}, // SHA256
			common.BytesToAddress([]byte{3}): {Balance: big.NewInt(1)}, // RIPEMD
			common.BytesToAddress([]byte{4}): {Balance: big.NewInt(1)}, // Identity
			common.BytesToAddress([]byte{5}): {Balance: big.NewInt(1)}, // ModExp
			common.BytesToAddress([]byte{6}): {Balance: big.NewInt(1)}, // ECAdd
			common.BytesToAddress([]byte{7}): {Balance: big.NewInt(1)}, // ECScalarMul
			common.BytesToAddress([]byte{8}): {Balance: big.NewInt(1)}, // ECPairing
			faucet: {Balance: new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), 256), big.NewInt(9))},
		},
		Dpor: types.DporSnap{
			Proposers: []common.Address{
				faucet,
			},
			Seal:       types.DporSignature{},
			Sigs:       make([]types.DporSignature, 4),
			Validators: []common.Address{},
		},
	}
}

func decodePrealloc(data string) GenesisAlloc {
	var p []struct{ Addr, Balance *big.Int }
	if err := rlp.NewStream(strings.NewReader(data), 0).Decode(&p); err != nil {
		panic(err)
	}
	ga := make(GenesisAlloc, len(p))
	for _, account := range p {
		ga[common.BigToAddress(account.Addr)] = GenesisAccount{Balance: account.Balance}
	}
	return ga
}
