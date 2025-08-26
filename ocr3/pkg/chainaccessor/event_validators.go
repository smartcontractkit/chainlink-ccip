package chainaccessor

import (
	"fmt"
	"math/big"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func ValidateSendRequestedEvent(
	ev *SendRequestedEvent, source, dest cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) error {
	if ev == nil {
		return fmt.Errorf("send requested event is nil")
	}

	if ev.Message.Header.DestChainSelector != dest {
		return fmt.Errorf("msg dest chain is not the expected queried one")
	}
	if ev.DestChainSelector != dest {
		return fmt.Errorf("dest chain is not the expected queried one")
	}

	if ev.Message.Header.SourceChainSelector != source {
		return fmt.Errorf("source chain is not the expected queried one")
	}

	if ev.SequenceNumber != ev.Message.Header.SequenceNumber {
		return fmt.Errorf("event sequence number does not match the message sequence number %d != %d",
			ev.SequenceNumber, ev.Message.Header.SequenceNumber)
	}

	if ev.SequenceNumber < seqNumRange.Start() || ev.SequenceNumber > seqNumRange.End() {
		return fmt.Errorf("send requested event sequence number is not in the expected range")
	}

	if ev.Message.Header.MessageID.IsEmpty() {
		return fmt.Errorf("message ID is zero")
	}

	if len(ev.Message.Receiver) == 0 {
		return fmt.Errorf("empty receiver address: %s", ev.Message.Receiver.String())
	}

	if ev.Message.Sender.IsZeroOrEmpty() {
		return fmt.Errorf("invalid sender address: %s", ev.Message.Sender.String())
	}

	if ev.Message.FeeTokenAmount.IsEmpty() {
		return fmt.Errorf("fee token amount is zero")
	}

	if ev.Message.FeeToken.IsZeroOrEmpty() {
		return fmt.Errorf("invalid fee token: %s", ev.Message.FeeToken.String())
	}

	return nil
}

func validateCommitReportAcceptedEvent(
	seq types.Sequence, gteTimestamp time.Time,
) (*CommitReportAcceptedEvent, error) {
	ev, is := (seq.Data).(*CommitReportAcceptedEvent)
	if !is {
		return nil, fmt.Errorf("unexpected type %T while expecting a commit report", seq)
	}

	if ev == nil {
		return nil, fmt.Errorf("commit report accepted event is nil")
	}

	if seq.Timestamp < uint64(gteTimestamp.Unix()) {
		return nil, fmt.Errorf("commit report accepted event timestamp is less than the minimum timestamp %v<%v",
			seq.Timestamp, gteTimestamp.Unix())
	}

	if err := validateMerkleRoots(append(ev.BlessedMerkleRoots, ev.UnblessedMerkleRoots...)); err != nil {
		return nil, fmt.Errorf("merkle roots: %w", err)
	}

	for _, tpus := range ev.PriceUpdates.TokenPriceUpdates {
		if tpus.SourceToken.IsZeroOrEmpty() {
			return nil, fmt.Errorf("invalid source token address: %s", tpus.SourceToken.String())
		}
		if tpus.UsdPerToken == nil || tpus.UsdPerToken.Cmp(big.NewInt(0)) <= 0 {
			return nil, fmt.Errorf("nil or non-positive usd per token")
		}
	}

	for _, gpus := range ev.PriceUpdates.GasPriceUpdates {
		if gpus.UsdPerUnitGas == nil || gpus.UsdPerUnitGas.Cmp(big.NewInt(0)) < 0 {
			return nil, fmt.Errorf("nil or negative usd per unit gas: %s", gpus.UsdPerUnitGas.String())
		}
	}

	return ev, nil
}

func validateMerkleRoots(merkleRoots []MerkleRoot) error {
	seenRoots := mapset.NewSet[cciptypes.Bytes32]()

	for _, mr := range merkleRoots {
		if seenRoots.Contains(mr.MerkleRoot) {
			return fmt.Errorf("duplicate merkle root: %s", mr.MerkleRoot.String())
		}
		seenRoots.Add(mr.MerkleRoot)

		if mr.SourceChainSelector == 0 {
			return fmt.Errorf("source chain is zero")
		}
		if mr.MinSeqNr == 0 {
			return fmt.Errorf("minSeqNr is zero")
		}
		if mr.MaxSeqNr == 0 {
			return fmt.Errorf("maxSeqNr is zero")
		}
		if mr.MinSeqNr > mr.MaxSeqNr {
			return fmt.Errorf("minSeqNr is greater than maxSeqNr")
		}
		if mr.MerkleRoot.IsEmpty() {
			return fmt.Errorf("empty merkle root")
		}
		if mr.OnRampAddress.IsZeroOrEmpty() {
			return fmt.Errorf("invalid onramp address: %s", mr.OnRampAddress.String())
		}
	}

	return nil
}

func validateExecutionStateChangedEvent(
	ev *ExecutionStateChangedEvent, rangesByChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange) error {
	if ev == nil {
		return fmt.Errorf("execution state changed event is nil")
	}

	if _, ok := rangesByChain[ev.SourceChainSelector]; !ok {
		return fmt.Errorf("source chain of messages was not queries")
	}

	if !ev.SequenceNumber.IsWithinRanges(rangesByChain[ev.SourceChainSelector]) {
		return fmt.Errorf("execution state changed event sequence number is not in the expected range")
	}

	if ev.MessageHash.IsEmpty() {
		return fmt.Errorf("nil message hash")
	}

	if ev.MessageID.IsEmpty() {
		return fmt.Errorf("message ID is zero")
	}

	if ev.State == 0 {
		return fmt.Errorf("state is zero")
	}

	return nil
}
