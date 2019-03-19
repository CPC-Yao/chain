package backend

import (
	"hash/fnv"

	"bitbucket.org/cpchain/chain/commons/log"
	"bitbucket.org/cpchain/chain/consensus"
	"bitbucket.org/cpchain/chain/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/p2p"
	lru "github.com/hashicorp/golang-lru"
)

// handleLBFTMsg handles given msg with lbft (lightweighted bft) mode
func (vh *Handler) handleLBFTMsg(msg p2p.Msg, p *RemoteSigner) error {
	if vh.lbft == nil {
		vh.lbft = NewLBFT(vh)
	}

	return vh.lbft.Handle(msg, p)
}

func logMsgReceived(number uint64, hash common.Hash, msgCode MsgCode, p *RemoteSigner) {
	log.Debug("received msg", "number", number, "hash", hash.Hex(), "msg code", msgCode.String(), "remote peer", func() string {
		if p != nil {
			return p.Coinbase().Hex()
		}
		return "nil peer"
	}())
}

func (vh *Handler) handleProposerConnectionMsg(version int, p *p2p.Peer, rw p2p.MsgReadWriter, msg p2p.Msg) (string, error) {
	switch msg.Code {
	case NewSignerMsg:

		log.Debug("received new signer msg from", "peer.RemoteAddress", p.RemoteAddr().String())

		var signerStatus SignerStatusData
		address, err := ReadSignerStatus(msg, &signerStatus)
		if err != nil {
			return common.Address{}.Hex(), err
		}

		var (
			currectNumber = vh.dpor.GetCurrentBlock().NumberU64()
			term          = vh.dpor.TermOf(currectNumber)
			futureTerm    = vh.dpor.FutureTermOf(currectNumber)
		)

		// if current or future proposer, add to local peer set
		if vh.dialer.isCurrentOrFutureProposer(address, term, futureTerm) {
			vh.dialer.addRemoteProposer(version, p, rw, address)
		}

		log.Debug("added the signer as a proposer", "address", address.Hex(), "peer.RemoteAddress", p.RemoteAddr().String())
		return address.Hex(), nil

	default:
		log.Warn("unknown msg code", "msg", msg.Code)
	}
	return common.Address{}.Hex(), nil
}

func (vh *Handler) handleLBFT2Msg(msg p2p.Msg, p *RemoteSigner) error {

	var (
		input         = &BlockOrHeader{}
		msgCode       = NoMsgCode
		currentNumber = vh.dpor.GetCurrentBlock().NumberU64()
	)

	switch msg.Code {
	case PreprepareBlockMsg:
		// recover the block from msg
		block, err := RecoverBlockFromMsg(msg, p)
		if err != nil {
			return err
		}

		// prepare input and msg code for the fsm
		input = &BlockOrHeader{
			block: block,
		}
		msgCode = PreprepareMsgCode

	case PrepareHeaderMsg:
		// recover the header from msg
		header, err := RecoverHeaderFromMsg(msg, p)
		if err != nil {
			return err
		}

		// prepare input and msg code for the fsm
		input = &BlockOrHeader{
			header: header,
		}
		msgCode = PrepareMsgCode

	case CommitHeaderMsg:
		// recover the header from msg
		header, err := RecoverHeaderFromMsg(msg, p)
		if err != nil {
			return err
		}

		// prepare input and msg code for the fsm
		input = &BlockOrHeader{
			header: header,
		}
		msgCode = CommitMsgCode

	case ValidateBlockMsg:
		// recover the block from msg
		block, err := RecoverBlockFromMsg(msg, p)
		if err != nil {
			return err
		}

		// prepare input and msg code for the fsm
		input = &BlockOrHeader{
			block: block,
		}
		msgCode = ValidateMsgCode

	case PreprepareImpeachBlockMsg:
		// recover the block from msg
		block, err := RecoverBlockFromMsg(msg, p)
		if err != nil {
			return err
		}

		// prepare input and msg code for the fsm
		input = &BlockOrHeader{
			block: block,
		}
		msgCode = ImpeachPreprepareMsgCode

	case PrepareImpeachHeaderMsg:
		// recover the header from msg
		header, err := RecoverHeaderFromMsg(msg, p)
		if err != nil {
			return err
		}

		// prepare input and msg code for the fsm
		input = &BlockOrHeader{
			header: header,
		}
		msgCode = ImpeachPrepareMsgCode

	case CommitImpeachHeaderMsg:
		// recover the header from msg
		header, err := RecoverHeaderFromMsg(msg, p)
		if err != nil {
			return err
		}

		// prepare input and msg code for the fsm
		input = &BlockOrHeader{
			header: header,
		}
		msgCode = ImpeachCommitMsgCode

	case ValidateImpeachBlockMsg:
		// recover the block from msg
		block, err := RecoverBlockFromMsg(msg, p)
		if err != nil {
			return err
		}

		// prepare input and msg code for the fsm
		input = &BlockOrHeader{
			block: block,
		}
		msgCode = ImpeachValidateMsgCode

	default:
		log.Warn("unknown msg code", "msg", msg.Code)
	}

	// log output received msg
	logMsgReceived(input.Number(), input.Hash(), msgCode, p)

	// if number is larger than local current number, sync from remote peer
	if input.Number() > currentNumber+1 && p != nil {
		go vh.dpor.SyncFrom(p.Peer)
		log.Debug("I am slow, syncing with peer", "peer", p.address)
	}

	// if number is equal or less than current number, drop the msg
	if input.Number() < currentNumber {
		log.Debug("received outdated msg, discarding...")
		return nil
	}

	// rebroadcast the msg
	// go vh.reBroadcast(input, msgCode, msg)

	// this is just for debug
	switch msgCode {
	// if received a impeach validate msg, log out some debug infos
	case ImpeachValidateMsgCode:

		log.Debug("-----------------------------")
		log.Debug("now received an ImpeachValidateMsg", "number", input.Number(), "hash", input.Hash().Hex())

		correctProposer, err := vh.dpor.ProposerOf(input.Number())
		if err != nil {
			log.Debug("err when get proposer of number", "err", err, "number", input.Number())
		}

		correctProposerPeer, exist := vh.dialer.getProposer(correctProposer.Hex())
		if !exist || correctProposerPeer == nil {
			log.Debug("proposer for the block is not in local proposer peer set")
			log.Debug("for this block number, the correct proposer should be", "addr", correctProposer.Hex())
		} else {
			log.Debug("for this block number, the correct proposer should be", "addr", correctProposer.Hex(), "ip:port", correctProposerPeer.Peer.RemoteAddr())
		}

		log.Debug("-----------------------------")
	}

	// if the msg is PreprepareImpeachBlockMsg, or msg code is ImpeachPreprepareMsgCode, the sender must be nil(self)
	switch msgCode {
	case ImpeachPreprepareMsgCode:
		if p != nil {
			// invalid impeach preprepare msg sender!
			return nil
		}
	}

	// call fsm
	output, action, msgCode, err := vh.fsm.FSM(input, msgCode)
	switch err {
	case nil:
	case consensus.ErrUnknownAncestor:
		log.Debug("added block to unknown ancestor cache", "number", input.Number(), "hash", input.Hash().Hex())

		vh.unknownAncestorBlocks.AddBlock(input.block)

	default:
		log.Error("received an error when run fsm", "err", err)

		return nil
	}

	// handle fsm result
	switch output {
	case nil:
		// nil output, do nothing

	default:
		switch action {
		case BroadcastMsgAction:

			switch msgCode {
			case PrepareMsgCode:
				go vh.BroadcastPrepareHeader(output[0].header)

			case CommitMsgCode:
				go vh.BroadcastCommitHeader(output[0].header)

			case PrepareAndCommitMsgCode:
				go vh.BroadcastPrepareHeader(output[0].header)
				go vh.BroadcastCommitHeader(output[1].header)

			case ValidateMsgCode:
				go vh.BroadcastValidateBlock(output[0].block)

			case ImpeachPrepareMsgCode:
				go vh.BroadcastPrepareImpeachHeader(output[0].header)

			case ImpeachCommitMsgCode:
				go vh.BroadcastCommitImpeachHeader(output[0].header)

			case ImpeachPrepareAndCommitMsgCode:
				go vh.BroadcastPrepareImpeachHeader(output[0].header)
				go vh.BroadcastCommitImpeachHeader(output[1].header)

			case ImpeachValidateMsgCode:
				go vh.BroadcastValidateImpeachBlock(output[0].block)

			// unknown msg code
			default:
				log.Debug("unknown msg code for fsm output", "msgCode", msgCode)
			}

		case BroadcastAndInsertBlockAction:
			switch msgCode {
			case ValidateMsgCode:
				go vh.dpor.InsertChain(output[0].block)
				go vh.dpor.BroadcastBlock(output[0].block, true)

			case ImpeachValidateMsgCode:
				go vh.dpor.InsertChain(output[0].block)
				go vh.dpor.BroadcastBlock(output[0].block, true)

			default:
				log.Debug("unknown msg code for fsm output", "msgCode", msgCode)
			}

		// other actions
		default:
			log.Debug("unknown action code for fsm output", "action", action)
		}

	}

	return nil
}

// ReceiveImpeachPendingBlock receives a block to add to pending block channel
func (vh *Handler) ReceiveImpeachPendingBlock(block *types.Block) error {
	select {
	case vh.pendingImpeachBlockCh <- block:
		err := vh.knownBlocks.AddBlock(block)
		if err != nil {
			return err
		}

		return nil
	}
}

type msgID struct {
	blockID BlockIdentifier
	msgCode MsgCode
	msgHash common.Hash
}

func newMsgID(number uint64, hash common.Hash, msgCode MsgCode, msg p2p.Msg) msgID {

	var payload []byte
	msgHash := fnv.New32a()
	msg.Payload.Read(payload)
	// msgHash.Write([]byte(msg.String()))
	msgHash.Write(payload)

	return msgID{
		blockID: BlockIdentifier{number: number, hash: hash},
		msgCode: msgCode,
		msgHash: common.BytesToHash(msgHash.Sum(nil)),
	}
}

type broadcastRecord struct {
	record *lru.ARCCache
}

func newBroadcastRecord() *broadcastRecord {
	record, _ := lru.NewARC(1000)
	return &broadcastRecord{
		record: record,
	}
}

func (br *broadcastRecord) markAsBroadcasted(number uint64, hash common.Hash, msgCode MsgCode, msg p2p.Msg) {
	msgID := newMsgID(number, hash, msgCode, msg)
	br.record.Add(msgID, true)
}

func (br *broadcastRecord) ifBroadcasted(number uint64, hash common.Hash, msgCode MsgCode, msg p2p.Msg) bool {
	msgID := newMsgID(number, hash, msgCode, msg)
	broadcasted, exists := br.record.Get(msgID)
	return exists && broadcasted.(bool) == true
}

type impeachmentRecord struct {
	record *lru.ARCCache
}

func newImpeachmentRecord() *impeachmentRecord {
	record, _ := lru.NewARC(1000)
	return &impeachmentRecord{
		record: record,
	}
}

func (ir *impeachmentRecord) markAsImpeached(number uint64, hash common.Hash) {
	bi := NewBlockIdentifier(number, hash)
	ir.record.Add(bi, true)
}

func (ir *impeachmentRecord) ifImpeached(number uint64, hash common.Hash) bool {
	bi := NewBlockIdentifier(number, hash)
	impeached, exists := ir.record.Get(bi)
	return exists && impeached.(bool) == true
}

func (vh *Handler) reBroadcast(input *BlockOrHeader, msgCode MsgCode, msg p2p.Msg) {
	if !vh.broadcastRecord.ifBroadcasted(input.Number(), input.Hash(), msgCode, msg) {
		switch msgCode {
		case PreprepareMsgCode:
			vh.BroadcastPreprepareBlock(input.block)
		case PrepareMsgCode:
			vh.BroadcastPrepareHeader(input.header)
		case CommitMsgCode:
			vh.BroadcastCommitHeader(input.header)
		case ValidateMsgCode:
			vh.BroadcastValidateBlock(input.block)
		case ImpeachPreprepareMsgCode:
			vh.BroadcastPreprepareImpeachBlock(input.block)
		case ImpeachPrepareMsgCode:
			vh.BroadcastPrepareImpeachHeader(input.header)
		case ImpeachCommitMsgCode:
			vh.BroadcastCommitImpeachHeader(input.header)
		case ImpeachValidateMsgCode:
			vh.BroadcastValidateImpeachBlock(input.block)
		default:
		}
		vh.broadcastRecord.markAsBroadcasted(input.Number(), input.Hash(), msgCode, msg)
	}
}
