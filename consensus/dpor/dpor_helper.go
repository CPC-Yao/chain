package dpor

import (
	"bytes"
	"math/big"
	"time"

	"bitbucket.org/cpchain/chain/accounts"
	"bitbucket.org/cpchain/chain/commons/log"
	"bitbucket.org/cpchain/chain/consensus"
	"bitbucket.org/cpchain/chain/types"
	"github.com/ethereum/go-ethereum/common"
)

type dporHelper interface {
	dporUtil
	verifyHeader(d *Dpor, chain consensus.ChainReader, header *types.Header, parents []*types.Header, refHeader *types.Header) error
	verifyCascadingFields(d *Dpor, chain consensus.ChainReader, header *types.Header, parents []*types.Header, refHeader *types.Header) error
	snapshot(d *Dpor, chain consensus.ChainReader, number uint64, hash common.Hash, parents []*types.Header) (*DporSnapshot, error)
	verifySeal(d *Dpor, chain consensus.ChainReader, header *types.Header, parents []*types.Header, refHeader *types.Header) error
	signHeader(d *Dpor, chain consensus.ChainReader, header *types.Header, state consensus.State) error
}

type defaultDporHelper struct {
	dporUtil
}

// verifyHeader checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (dh *defaultDporHelper) verifyHeader(c *Dpor, chain consensus.ChainReader, header *types.Header, parents []*types.Header, refHeader *types.Header) error {
	if header.Number == nil {
		return errUnknownBlock
	}

	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time.Cmp(big.NewInt(time.Now().Unix())) > 0 {
		return consensus.ErrFutureBlock
	}

	switch c.fake {
	case DoNothingFakeMode:
		// do nothing
	case FakeMode:
		return nil
	}

	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}

	// Check if extraData is valid
	signersBytes := len(header.Extra) - extraVanity - extraSeal
	if signersBytes%common.AddressLength != 0 {
		return errInvalidSigners
	}

	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixHash != (common.Hash{}) {
		return errInvalidMixHash
	}

	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if header.Difficulty == nil || (header.Difficulty.Cmp(diffInTurn) != 0 && header.Difficulty.Cmp(diffNoTurn) != 0) {
			return errInvalidDifficulty
		}
	}

	// All basic checks passed, verify cascading fields
	return c.dh.verifyCascadingFields(c, chain, header, parents, refHeader)
}

// verifyCascadingFields verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (dh *defaultDporHelper) verifyCascadingFields(dpor *Dpor, chain consensus.ChainReader, header *types.Header, parents []*types.Header, refHeader *types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}

	// Ensure that the block's timestamp isn't too close to it's parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}

	// Ensure that the block's parent is valid
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}

	// Ensure that the block's timestamp is valid
	if parent.Time.Uint64()+dpor.config.Period > header.Time.Uint64() {
		return ErrInvalidTimestamp
	}

	// Retrieve the Snapshot needed to verify this header and cache it
	snap, err := dh.snapshot(dpor, chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}

	// Check signers bytes in extraData
	signers := make([]byte, dpor.config.Epoch*common.AddressLength)
	for round, signer := range snap.SignersOf(number) {
		copy(signers[round*common.AddressLength:(round+1)*common.AddressLength], signer[:])
	}
	extraSuffix := len(header.Extra) - extraSeal
	if !bytes.Equal(header.Extra[extraVanity:extraSuffix], signers) {
		if NormalMode == dpor.fake {
			return errInvalidSigners
		}
	}

	// All basic checks passed, verify the seal and return
	return dh.verifySeal(dpor, chain, header, parents, refHeader)
}

// Snapshot retrieves the authorization Snapshot at a given point in time.
func (dh *defaultDporHelper) snapshot(dpor *Dpor, chain consensus.ChainReader, number uint64, hash common.Hash, parents []*types.Header) (*DporSnapshot, error) {
	// Search for a Snapshot in memory or on disk for checkpoints
	var (
		headers []*types.Header
		snap    *DporSnapshot
	)
	for snap == nil {
		// If an in-memory Snapshot was found, use that
		if s, ok := dpor.recents.Get(hash); ok {
			snap = s.(*DporSnapshot)
			break
		}

		// If an on-disk checkpoint Snapshot can be found, use that
		// if number%checkpointInterval == 0 {
		if IsCheckPoint(number, dpor.config.Epoch, dpor.config.View) {
			if s, err := loadSnapshot(dpor.config, dpor.db, hash); err == nil {
				log.Debug("Loaded voting Snapshot from disk", "number", number, "hash", hash)
				snap = s
				break
			}
		}

		// If we're at block zero, make a Snapshot
		if number == 0 {
			// Retrieve genesis block and verify it
			genesis := chain.GetHeaderByNumber(0)
			if err := dpor.dh.verifyHeader(dpor, chain, genesis, nil, nil); err != nil {
				return nil, err
			}

			var signers []common.Address
			if dpor.fake == FakeMode || dpor.fake == DoNothingFakeMode {
				// do nothing when test,empty signers assigned
			} else {
				// Create a snapshot from the genesis block
				signers = make([]common.Address, (len(genesis.Extra)-extraVanity-extraSeal)/common.AddressLength)
				for i := 0; i < len(signers); i++ {
					copy(signers[i][:], genesis.Extra[extraVanity+i*common.AddressLength:])
				}
			}

			snap = newSnapshot(dpor.config, 0, genesis.Hash(), signers)
			if err := snap.store(dpor.db); err != nil {
				return nil, err
			}
			log.Debug("Stored genesis voting Snapshot to disk")
			break
		}

		// No Snapshot for this header, gather the header and move backward
		var header *types.Header
		if len(parents) > 0 {
			// If we have explicit parents, pick from there (enforced)
			header = parents[len(parents)-1]
			if header.Hash() != hash || header.Number.Uint64() != number {
				return nil, consensus.ErrUnknownAncestor
			}
			parents = parents[:len(parents)-1]
		} else {
			// No explicit parents (or no more left), reach out to the database
			header = chain.GetHeader(hash, number)
			if header == nil {
				return nil, consensus.ErrUnknownAncestor
			}
		}
		headers = append(headers, header)
		number, hash = number-1, header.ParentHash
	}

	// Previous Snapshot found, apply any pending headers on top of it
	for i := 0; i < len(headers)/2; i++ {
		headers[i], headers[len(headers)-1-i] = headers[len(headers)-1-i], headers[i]
	}

	dpor.lock.Lock()
	contractCaller := dpor.contractCaller
	dpor.lock.Unlock()

	// Apply headers to the snapshot and updates RPTs
	snap, err := snap.apply(headers, contractCaller)
	if err != nil {
		return nil, err
	}

	// Save to cache
	dpor.recents.Add(snap.Hash, snap)

	// If we've generated a new checkpoint Snapshot, save to disk
	if IsCheckPoint(snap.Number, dpor.config.Epoch, dpor.config.View) && len(headers) > 0 {
		if err = snap.store(dpor.db); err != nil {
			return nil, err
		}
		log.Debug("Stored voting Snapshot to disk", "number", snap.Number, "hash", snap.Hash)
	}
	return snap, err
}

// verifySeal checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (dh *defaultDporHelper) verifySeal(dpor *Dpor, chain consensus.ChainReader, header *types.Header, parents []*types.Header, refHeader *types.Header) error {
	hash := header.Hash()
	number := header.Number.Uint64()

	// Verifying the genesis block is not supported
	if number == 0 {
		return errUnknownBlock
	}

	// Fake Dpor doesn't do seal check
	if dpor.fake == FakeMode || dpor.fake == DoNothingFakeMode {
		time.Sleep(dpor.fakeDelay)
		if dpor.fakeFail == number {
			return errFakerFail
		}
		return nil
	}

	// Resolve the authorization key and check against signers
	leader, signers, err := dpor.dh.ecrecover(header, dpor.signatures)
	if err != nil {
		return err
	}

	// Retrieve the Snapshot needed to verify this header and cache it
	snap, err := dh.snapshot(dpor, chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}

	// Ensure that the difficulty corresponds to the turn-ness of the signer
	inturn, _ := snap.IsLeaderOf(leader, header.Number.Uint64())
	if inturn && header.Difficulty.Cmp(diffInTurn) != 0 {
		return errInvalidDifficulty
	}
	if !inturn && header.Difficulty.Cmp(diffNoTurn) != 0 {
		return errInvalidDifficulty
	}

	// Some debug infos here
	log.Debug("--------dpor.verifySeal start--------")
	log.Debug("hash", "hash", hash.Hex())
	log.Debug("number", "number", number)
	log.Debug("current header", "number", chain.CurrentHeader().Number.Uint64())
	log.Debug("leader", "address", leader.Hex())
	log.Debug("signers recoverd from header: ")
	for _, signer := range signers {
		log.Debug("signer", "address", signer.Hex())
	}
	log.Debug("signers in snapshot: ")
	for _, signer := range snap.SignersOf(number) {
		log.Debug("signer", "address", signer.Hex())
	}

	// Check if the leader is the real leader
	ok, err := snap.IsLeaderOf(leader, number)
	if err != nil {
		return err
	}
	// If leader is a wrong leader, return err
	if !ok {
		return consensus.ErrUnauthorized
	}

	// Check if accept the sigs
	accept, err := dpor.dh.acceptSigs(header, dpor.signatures, snap.SignersOf(number), uint(dpor.config.Epoch))
	if err != nil {
		return err
	}

	// We haven't reached the 2/3 rule
	if !accept {
		return consensus.ErrNotEnoughSigs
	}

	return nil
}

// sighHeader signs the given refHeader if self is in the committee
func (dh *defaultDporHelper) signHeader(dpor *Dpor, chain consensus.ChainReader, header *types.Header, state consensus.State) error {
	hash := header.Hash()
	number := header.Number.Uint64()

	// Retrieve the Snapshot needed to verify this header and cache it
	snap, err := dh.snapshot(dpor, chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}

	// Retrieve signatures of the block in cache
	s, _ := dpor.signatures.Get(hash)

	// Copy all signatures recovered to allSigs.
	allSigs := make([]byte, int(dpor.config.Epoch)*extraSeal)
	for round, signer := range snap.SignersOf(number) {
		if sigHash, ok := s.(*Signatures).GetSig(signer); ok {
			copy(allSigs[round*extraSeal:(round+1)*extraSeal], sigHash)
		}
	}

	// Encode allSigs to header.extra2.
	err = header.EncodeToExtra2(types.Extra2Struct{Type: types.TypeExtra2Signatures, Data: allSigs})
	if err != nil {
		return err
	}

	// Sign the block if self is in the committee
	if snap.IsSignerOf(dpor.signer, number) {

		// NOTE: sign a block only once
		if signedHash, signed := dpor.signedBlocks[header.Number.Uint64()]; signed && signedHash != header.Hash() {
			return errMultiBlocksInOneHeight
		}

		// Sign it!
		sighash, err := dpor.signFn(accounts.Account{Address: dpor.signer}, dpor.dh.sigHash(header).Bytes())
		if err != nil {
			return err
		}

		// Copy signer's signature to the right position in the allSigs
		round, _ := snap.SignerRoundOf(dpor.signer, number)
		copy(allSigs[round*extraSeal:(round+1)*extraSeal], sighash)

		// Encode to header.extra2
		err = header.EncodeToExtra2(types.Extra2Struct{Type: types.TypeExtra2Signatures, Data: allSigs})
		if err != nil {
			return err
		}

		return nil
	}
	return errSignerNotInCommittee
}

// timeToDialCommittee checks if it is time to dial remote signers, and dails them if time is up
func (dh *defaultDporHelper) timeToDialCommittee(dpor *Dpor, chain consensus.ChainReader) bool {

	header := chain.CurrentHeader()
	number := header.Number.Uint64()

	// Retrieve the Snapshot needed to verify this header and cache it
	snap, err := dh.snapshot(dpor, chain, number, header.Hash(), nil)
	if err != nil {
		return false
	}

	// Some debug infos
	log.Debug("my address", "eb", dpor.signer.Hex())
	log.Debug("current block number", "number", number)
	log.Debug("ISCheckPoint", "bool", IsCheckPoint(number, dpor.config.Epoch, dpor.config.View))
	log.Debug("is future signer", "bool", snap.IsFutureSignerOf(dpor.signer, number))
	log.Debug("epoch idx of block number", "block epochIdx", snap.EpochIdxOf(number))

	log.Debug("recent signers: ")
	for i := snap.EpochIdxOf(number); i < snap.EpochIdxOf(number)+5; i++ {
		log.Debug("----------------------")
		log.Debug("signers in snapshot of:", "epoch idx", i)
		for _, s := range snap.RecentSigners[i] {
			log.Debug("signer", "s", s.Hex())
		}
	}

	// If in a checkpoint and self is in the future committee, try to build the committee network
	isCheckpoint := IsCheckPoint(number, dpor.config.Epoch, dpor.config.View)
	isFutureSigner := snap.IsFutureSignerOf(dpor.signer, number)
	ifStartDynamic := number >= dpor.config.MaxInitBlockNumber

	return isCheckpoint && isFutureSigner && ifStartDynamic
}

func (dh *defaultDporHelper) dialCommittee(dpor *Dpor, snap *DporSnapshot, number uint64) error {
	log.Info("In future committee, building the committee network...")

	epochIdx := snap.FutureEpochIdxOf(number)
	signers := snap.FutureSignersOf(number)

	go func(eIdx uint64, committee []common.Address) {
		// Updates handler.signers
		dpor.handler.UpdateSigners(eIdx, committee)
		// Connect all
		dpor.handler.DialAll()
	}(epochIdx, signers)

	return nil
}
