// Copyright 2017 The go-ethereum Authors
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

package consensus

import (
	"errors"

	"bitbucket.org/cpchain/chain/core/types"
	"github.com/ethereum/go-ethereum/common"
)

var (
	// ErrUnknownAncestor is returned when validating a block requires an ancestor
	// that is unknown.
	ErrUnknownAncestor = errors.New("unknown ancestor")

	// ErrPrunedAncestor is returned when validating a block requires an ancestor
	// that is known, but the state of which is not available.
	ErrPrunedAncestor = errors.New("pruned ancestor")

	// ErrFutureBlock is returned when a block's timestamp is in the future according
	// to the current node.
	ErrFutureBlock = errors.New("block in the future")

	// ErrInvalidNumber is returned if a block's number doesn't equal it's parent's
	// plus one.
	ErrInvalidNumber = errors.New("invalid block number")

	// ErrNotEnoughSigs is returned if there is not enough signatures for a block.
	ErrNotEnoughSigs = &ErrNotEnoughSigsType{NotEnoughSigsBlockHash: common.Hash{}}

	// ErrUnauthorized is returned if a header is signed by a non-authorized entity.
	ErrUnauthorized = errors.New("unauthorized leader")

	// ErrNewSignedHeader is returned if i sign the block, but not accept the block yet.
	ErrNewSignedHeader = &ErrNewSignedHeaderType{SignedHeader: &types.Header{}}
	// ErrNewSignedHeader = errors.New("new signed header")
)

// ErrNotEnoughSigsType is returned if there is not enough signatures for a block.
type ErrNotEnoughSigsType struct {
	NotEnoughSigsBlockHash common.Hash
}

func (e *ErrNotEnoughSigsType) Error() string {
	return "not enough sigs: block hash: " + e.NotEnoughSigsBlockHash.Hex()
}

// ErrNewSignedHeaderType is type for ErrNewSignedHeader.
type ErrNewSignedHeaderType struct {
	SignedHeader *types.Header
}

func (e *ErrNewSignedHeaderType) Error() string {
	return "signed the header: header hash: " + e.SignedHeader.Hash().Hex()
}
