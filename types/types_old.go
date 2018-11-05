// Copyright 2014 The go-ethereum Authors
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

package types

import (
	"errors"
	"io"
	"math/big"
	"sync/atomic"

	"time"

	"bytes"
	"fmt"

	"bitbucket.org/cpchain/chain/commons/log"
	"bitbucket.org/cpchain/chain/crypto"
	"bitbucket.org/cpchain/chain/crypto/sha3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rlp"
)

// This file contains old Transaction, txdata, Block, extblock, for unit test updating their old serialization string.

//go:generate gencodec -type txdataold -field-override txdataMarshalingOld -out gen_tx_json_old.go
//go:generate gencodec -type HeaderOld -field-override headerOldMarshaling -out gen_header_json_old.go

var (
	ErrInvalidSigOld = errors.New("invalid transaction v, r, s values")
)

// example is the example code for demonstrating how to use to update serialization string.
func example(oldHex string) {
	data := common.Hex2Bytes(oldHex)
	txold := TransactionOld{}
	rlp.Decode(bytes.NewReader(data), &txold)

	tx := txold.ToNewTx()
	bb, _ := rlp.EncodeToBytes(tx)
	fmt.Println(bb)
}

func ConvertOldBlockHexToNewBlockHex(oldBlockHex string) string {
	var oldblock BlockOld
	err := rlp.DecodeBytes(common.Hex2Bytes(oldBlockHex), &oldblock)
	if err != nil {
		log.Warn("Error when converting", err)
		return ""
	}

	nb := oldblock.ToNewBlock()
	b, _ := rlp.EncodeToBytes(nb)
	return common.Bytes2Hex(b)
}

func ConvertOldTxHexToNewTxHex(oldTxHex string) string {
	var oldtx TransactionOld
	rlp.DecodeBytes(common.Hex2Bytes(oldTxHex), &oldtx)
	nt := oldtx.ToNewTx()
	b, _ := rlp.EncodeToBytes(nt)
	return common.Bytes2Hex(b)
}

// Header represents a block header in the Ethereum blockchain.
type HeaderOld struct {
	ParentHash  common.Hash    `json:"parentHash"       gencodec:"required"`
	UncleHash   common.Hash    `json:"sha3Uncles"       gencodec:"required"`
	Coinbase    common.Address `json:"miner"            gencodec:"required"`
	Root        common.Hash    `json:"stateRoot"        gencodec:"required"`
	TxHash      common.Hash    `json:"transactionsRoot" gencodec:"required"`
	ReceiptHash common.Hash    `json:"receiptsRoot"     gencodec:"required"`
	Bloom       Bloom          `json:"logsBloom"        gencodec:"required"`
	Difficulty  *big.Int       `json:"difficulty"       gencodec:"required"`
	Number      *big.Int       `json:"number"           gencodec:"required"`
	GasLimit    uint64         `json:"gasLimit"         gencodec:"required"`
	GasUsed     uint64         `json:"gasUsed"          gencodec:"required"`
	Time        *big.Int       `json:"timestamp"        gencodec:"required"`
	Extra       []byte         `json:"extraData"        gencodec:"required"`
	Extra2      []byte         `json:"extraData2"       gencodec:"required"`
	MixDigest   common.Hash    `json:"mixHash"          gencodec:"required"`
	Nonce       BlockNonce     `json:"nonce"            gencodec:"required"`
}

func (h *HeaderOld) ToNewType() *Header {
	return &Header{
		ParentHash:   h.ParentHash,
		Coinbase:     h.Coinbase,
		StateRoot:    h.Root,
		TxsRoot:      h.TxHash,
		ReceiptsRoot: h.ReceiptHash,
		LogsBloom:    h.Bloom,
		Difficulty:   h.Difficulty,
		Number:       h.Number,
		GasLimit:     h.GasLimit,
		GasUsed:      h.GasUsed,
		Time:         h.Time,
		Extra:        h.Extra,
		Extra2:       h.Extra2,
		MixHash:      h.MixDigest,
		Nonce:        h.Nonce,
	}
}

func (header *HeaderOld) Hash() (hash common.Hash) {
	hasher := sha3.NewKeccak256()

	rlp.Encode(hasher, []interface{}{
		header.ParentHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra,
		header.MixDigest,
		header.Nonce,
	})
	hasher.Sum(hash[:0])
	return hash
}

// field type overrides for gencodec
type headerOldMarshaling struct {
	Difficulty *hexutil.Big
	Number     *hexutil.Big
	GasLimit   hexutil.Uint64
	GasUsed    hexutil.Uint64
	Time       *hexutil.Big
	Extra      hexutil.Bytes
	Extra2     hexutil.Bytes
	Hash       common.Hash `json:"hash"` // adds call to Hash() in MarshalJSON
}

// "external" block encoding. used for eth protocol, etc.
type extblockold struct {
	Header *HeaderOld
	Txs    []*Transaction
	Uncles []*HeaderOld
}

// Block represents an entire block in the Ethereum blockchain.
type BlockOld struct {
	header       *HeaderOld
	uncles       []*HeaderOld
	transactions Transactions

	// caches
	hash atomic.Value
	size atomic.Value

	// Td is used by package core to store the total difficulty
	// of the chain up to and including the block.
	td *big.Int

	// These fields are used by package eth to track
	// inter-peer block relay.
	ReceivedAt   time.Time
	ReceivedFrom interface{}
}

func (b *BlockOld) ToNewBlock() *Block {
	return &Block{
		header:       b.header.ToNewType(),
		transactions: b.transactions,
		hash:         b.hash,
		size:         b.size,
		td:           b.td,

		ReceivedAt:   b.ReceivedAt,
		ReceivedFrom: b.ReceivedFrom,
	}
}

// DecodeRLP decodes the Ethereum
func (b *BlockOld) DecodeRLP(s *rlp.Stream) error {
	var eb extblockold
	_, size, _ := s.Kind()
	if err := s.Decode(&eb); err != nil {
		return err
	}
	b.header, b.uncles, b.transactions = eb.Header, eb.Uncles, eb.Txs
	b.size.Store(common.StorageSize(rlp.ListSize(size)))
	return nil
}

// EncodeRLP serializes b into the Ethereum RLP block format.
func (b *BlockOld) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, extblockold{
		Header: b.header,
		Txs:    b.transactions,
		Uncles: b.uncles,
	})
}

// TransactionOld is created for compatibility, usually used by unittest.
type TransactionOld struct {
	data txdataold
	// caches
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}

type txdataold struct {
	AccountNonce uint64          `json:"nonce"    gencodec:"required"`
	Price        *big.Int        `json:"gasPrice" gencodec:"required"`
	GasLimit     uint64          `json:"gas"      gencodec:"required"`
	Recipient    *common.Address `json:"to"       rlp:"nil"` // nil means contract creation
	Amount       *big.Int        `json:"value"    gencodec:"required"`
	Payload      []byte          `json:"input"    gencodec:"required"`

	// Signature values
	V *big.Int `json:"v" gencodec:"required"`
	R *big.Int `json:"r" gencodec:"required"`
	S *big.Int `json:"s" gencodec:"required"`

	// This is only used when marshaling to JSON.
	Hash *common.Hash `json:"hash" rlp:"-"`

	// IsPrivate indicates if the transaction is private.
	IsPrivate bool `json:"private" gencodec:"required"`
}

// TODO: add new parameter 'isPrivate'.
func NewTransactionOld(nonce uint64, to common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *TransactionOld {
	return newTransactionOld(nonce, &to, amount, gasLimit, gasPrice, data, false)
}

// TODO: add new parameter 'isPrivate'.
func NewContractCreationOld(nonce uint64, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte) *TransactionOld {
	return newTransactionOld(nonce, nil, amount, gasLimit, gasPrice, data, false)
}

func newTransactionOld(nonce uint64, to *common.Address, amount *big.Int, gasLimit uint64, gasPrice *big.Int, data []byte, isPrivate bool) *TransactionOld {
	if len(data) > 0 {
		data = common.CopyBytes(data)
	}
	d := txdataold{
		AccountNonce: nonce,
		Recipient:    to,
		Payload:      data,
		Amount:       new(big.Int),
		GasLimit:     gasLimit,
		Price:        new(big.Int),
		V:            new(big.Int),
		R:            new(big.Int),
		S:            new(big.Int),
		IsPrivate:    isPrivate,
	}
	if amount != nil {
		d.Amount.Set(amount)
	}
	if gasPrice != nil {
		d.Price.Set(gasPrice)
	}

	return &TransactionOld{data: d}
}

type txdataMarshalingOld struct {
	AccountNonce hexutil.Uint64
	Price        *hexutil.Big
	GasLimit     hexutil.Uint64
	Amount       *hexutil.Big
	Payload      hexutil.Bytes
	V            *hexutil.Big
	R            *hexutil.Big
	S            *hexutil.Big
}

// ChainId returns which chain id this transaction was signed for (if at all)
func (tx *TransactionOld) ChainId() *big.Int {
	return deriveChainId(tx.data.V)
}

// Protected returns whether the transaction is protected from replay protection.
func (tx *TransactionOld) Protected() bool {
	return isProtectedV(tx.data.V)
}

// EncodeRLP implements rlp.Encoder
func (tx *TransactionOld) EncodeRLP(w io.Writer) error {
	return rlp.Encode(w, &tx.data)
}

// DecodeRLP implements rlp.Decoder
func (tx *TransactionOld) DecodeRLP(s *rlp.Stream) error {
	_, size, _ := s.Kind()
	err := s.Decode(&tx.data)
	if err == nil {
		tx.size.Store(common.StorageSize(rlp.ListSize(size)))
	}

	return err
}

// MarshalJSON encodes the web3 RPC transaction format.
func (tx *TransactionOld) MarshalJSON() ([]byte, error) {
	hash := tx.Hash()
	data := tx.data
	data.Hash = &hash
	return data.MarshalJSON()
}

// UnmarshalJSON decodes the web3 RPC transaction format.
func (tx *TransactionOld) UnmarshalJSON(input []byte) error {
	var dec txdataold
	if err := dec.UnmarshalJSON(input); err != nil {
		return err
	}
	var V byte
	if isProtectedV(dec.V) {
		chainID := deriveChainId(dec.V).Uint64()
		V = byte(dec.V.Uint64() - 35 - 2*chainID)
	} else {
		V = byte(dec.V.Uint64() - 27)
	}
	if !crypto.ValidateSignatureValues(V, dec.R, dec.S, false) {
		return ErrInvalidSig
	}
	*tx = TransactionOld{data: dec}
	return nil
}

func (tx *TransactionOld) Data() []byte       { return common.CopyBytes(tx.data.Payload) }
func (tx *TransactionOld) Gas() uint64        { return tx.data.GasLimit }
func (tx *TransactionOld) GasPrice() *big.Int { return new(big.Int).Set(tx.data.Price) }
func (tx *TransactionOld) Value() *big.Int    { return new(big.Int).Set(tx.data.Amount) }
func (tx *TransactionOld) Nonce() uint64      { return tx.data.AccountNonce }
func (tx *TransactionOld) CheckNonce() bool   { return true }

// To returns the recipient address of the transaction.
// It returns nil if the transaction is a contract creation.
func (tx *TransactionOld) To() *common.Address {
	if tx.data.Recipient == nil {
		return nil
	}
	to := *tx.data.Recipient
	return &to
}

// Hash hashes the RLP encoding of tx.
// It uniquely identifies the transaction.
func (tx *TransactionOld) Hash() common.Hash {
	if hash := tx.hash.Load(); hash != nil {
		return hash.(common.Hash)
	}
	v := rlpHash(tx)
	tx.hash.Store(v)
	return v
}

// Size returns the true RLP encoded storage size of the transaction, either by
// encoding and returning it, or returning a previsouly cached value.
func (tx *TransactionOld) Size() common.StorageSize {
	if size := tx.size.Load(); size != nil {
		return size.(common.StorageSize)
	}
	c := writeCounter(0)
	rlp.Encode(&c, &tx.data)
	tx.size.Store(common.StorageSize(c))
	return common.StorageSize(c)
}

func (tx *TransactionOld) ToNewTx() *Transaction {
	types := uint64(0)
	if tx.IsPrivate() {
		types |= TxTypePrivate
	}

	d := txdata{
		Types:        types,
		AccountNonce: tx.data.AccountNonce,
		Recipient:    tx.data.Recipient,
		Payload:      tx.data.Payload,
		Amount:       tx.data.Amount,
		GasLimit:     tx.data.GasLimit,
		Price:        tx.data.Price,
		V:            tx.data.V,
		R:            tx.data.R,
		S:            tx.data.S,
	}

	return &Transaction{data: d}
}

// Cost returns amount + gasprice * gaslimit.
func (tx *TransactionOld) Cost() *big.Int {
	total := new(big.Int).Mul(tx.data.Price, new(big.Int).SetUint64(tx.data.GasLimit))
	total.Add(total, tx.data.Amount)
	return total
}

func (tx *TransactionOld) RawSignatureValues() (*big.Int, *big.Int, *big.Int) {
	return tx.data.V, tx.data.R, tx.data.S
}

// IsPrivate checks if the tx is private.
func (tx *TransactionOld) IsPrivate() bool {
	return tx.data.IsPrivate
}

// SetPrivate sets the tx as private.
func (tx *TransactionOld) SetPrivate(isPrivate bool) {
	tx.data.IsPrivate = isPrivate
}

// Transactions is a Transaction slice type for basic sorting.
type TransactionsOld []*TransactionOld

// Len returns the length of s.
func (s TransactionsOld) Len() int { return len(s) }

// Swap swaps the i'th and the j'th element in s.
func (s TransactionsOld) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// GetRlp implements Rlpable and returns the i'th element of s in rlp.
func (s TransactionsOld) GetRlp(i int) []byte {
	enc, _ := rlp.EncodeToBytes(s[i])
	return enc
}