package merkleroot

import (
	"context"
	"encoding/hex"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func (w *Processor) ObservationQuorum(_ ocr3types.OutcomeContext, _ types.Query) (ocr3types.Quorum, error) {
	// Across all chains we require at least 2F+1 observations.
	return ocr3types.QuorumTwoFPlusOne, nil
}

func (w *Processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	q Query,
) (Observation, error) {
	tStart := time.Now()
	observation, nextState := w.getObservation(ctx, q, prevOutcome)
	w.lggr.Infow("Sending MerkleRootObs",
		"observation", observation, "nextState", nextState, "observationDuration", time.Since(tStart))
	return observation, nil
}

func (w *Processor) getObservation(ctx context.Context, q Query, previousOutcome Outcome) (Observation, State) {
	nextState := previousOutcome.NextState()
	switch nextState {
	case SelectingRangesForReport:
		offRampNextSeqNums := w.observer.ObserveOffRampNextSeqNums(ctx)
		return Observation{
			// TODO: observe OnRamp max seq nums. The use of offRampNextSeqNums here effectively disables batching,
			// e.g. the ranges selected for each chain will be [x, x] (e.g. [46, 46]), which means reports will only
			// contain one message per chain. Querying the OnRamp contract requires changes to reader.CCIPReader,
			// which will need to be done in a future change.
			OnRampMaxSeqNums:   offRampNextSeqNums,
			OffRampNextSeqNums: offRampNextSeqNums,
			FChain:             w.observer.ObserveFChain(),
		}, nextState
	case BuildingReport:
		if q.RetryRMNSignatures {
			// RMN signature computation failed, we only want to retry getting the RMN signatures in the next round.
			// So there's nothing to observe, i.e. we don't want to build the report yet.
			return Observation{}, nextState
		}
		return Observation{
			MerkleRoots: w.observer.ObserveMerkleRoots(ctx, previousOutcome.RangesSelectedForReport),
			FChain:      w.observer.ObserveFChain(),
		}, nextState
	case WaitingForReportTransmission:
		return Observation{
			OffRampNextSeqNums: w.observer.ObserveOffRampNextSeqNums(ctx),
			FChain:             w.observer.ObserveFChain(),
		}, nextState
	default:
		w.lggr.Errorw("Unexpected state", "state", nextState)
		return Observation{}, nextState
	}
}

type Observer interface {
	// ObserveOffRampNextSeqNums observes the next sequence numbers for each source chain from the OffRamp
	ObserveOffRampNextSeqNums(ctx context.Context) []plugintypes.SeqNumChain

	// ObserveMerkleRoots computes the merkle roots for the given sequence number ranges
	ObserveMerkleRoots(ctx context.Context, ranges []plugintypes.ChainRange) []cciptypes.MerkleRootChain

	ObserveFChain() map[cciptypes.ChainSelector]int
}

type ObserverImpl struct {
	lggr         logger.Logger
	homeChain    reader.HomeChain
	nodeID       commontypes.OracleID
	chainSupport plugincommon.ChainSupport
	ccipReader   readerpkg.CCIPReader
	msgHasher    cciptypes.MessageHasher
}

// ObserveOffRampNextSeqNums observes the next sequence numbers for each source chain from the OffRamp
func (o ObserverImpl) ObserveOffRampNextSeqNums(ctx context.Context) []plugintypes.SeqNumChain {
	supportsDestChain, err := o.chainSupport.SupportsDestChain(o.nodeID)
	if err != nil {
		o.lggr.Warnw("call to SupportsDestChain failed", "err", err)
		return nil
	}

	if !supportsDestChain {
		return nil
	}

	sourceChains, err := o.chainSupport.KnownSourceChainsSlice()
	if err != nil {
		o.lggr.Warnw("call to KnownSourceChainsSlice failed", "err", err)
		return nil
	}

	sort.Slice(sourceChains, func(i, j int) bool { return sourceChains[i] < sourceChains[j] })
	offRampNextSeqNums, err := o.ccipReader.NextSeqNum(ctx, sourceChains)
	if err != nil {
		o.lggr.Warnw("call to NextSeqNum failed", "err", err)
		return nil
	}

	if len(offRampNextSeqNums) != len(sourceChains) {
		o.lggr.Errorf("call to NextSeqNum returned unexpected number of seq nums, got %d, expected %d",
			len(offRampNextSeqNums), len(sourceChains))
		return nil
	}

	result := make([]plugintypes.SeqNumChain, len(sourceChains))
	for i := range sourceChains {
		result[i] = plugintypes.SeqNumChain{ChainSel: sourceChains[i], SeqNum: offRampNextSeqNums[i]}
	}

	return result
}

// ObserveMerkleRoots computes the merkle roots for the given sequence number ranges
func (o ObserverImpl) ObserveMerkleRoots(
	ctx context.Context,
	ranges []plugintypes.ChainRange,
) []cciptypes.MerkleRootChain {

	supportedChains, err := o.chainSupport.SupportedChains(o.nodeID)
	if err != nil {
		o.lggr.Warnw("call to supportedChains failed", "err", err)
		return nil
	}

	var roots []cciptypes.MerkleRootChain
	rootsMu := &sync.Mutex{}
	wg := sync.WaitGroup{}
	for _, chainRange := range ranges {
		if supportedChains.Contains(chainRange.ChainSel) {
			wg.Add(1)
			go func() {
				defer wg.Done()
				msgs, err := o.ccipReader.MsgsBetweenSeqNums(ctx, chainRange.ChainSel, chainRange.SeqNumRange)
				if err != nil {
					o.lggr.Warnw("call to MsgsBetweenSeqNums failed", "err", err)
					return
				}

				root, err := o.computeMerkleRoot(ctx, msgs)
				if err != nil {
					o.lggr.Warnw("call to computeMerkleRoot failed", "err", err)
					return
				}

				merkleRoot := cciptypes.MerkleRootChain{
					ChainSel:     chainRange.ChainSel,
					SeqNumsRange: chainRange.SeqNumRange,
					MerkleRoot:   root,
				}

				rootsMu.Lock()
				roots = append(roots, merkleRoot)
				rootsMu.Unlock()
			}()
		}
	}
	wg.Wait()

	return roots
}

// computeMerkleRoot computes the merkle root of a list of messages
func (o ObserverImpl) computeMerkleRoot(ctx context.Context, msgs []cciptypes.Message) (cciptypes.Bytes32, error) {
	var hashes [][32]byte
	sort.Slice(msgs, func(i, j int) bool { return msgs[i].Header.SequenceNumber < msgs[j].Header.SequenceNumber })

	for i, msg := range msgs {
		// Assert there are no sequence number gaps in msgs
		if i > 0 {
			if msg.Header.SequenceNumber != msgs[i-1].Header.SequenceNumber+1 {
				return [32]byte{}, fmt.Errorf("found non-consecutive sequence numbers when computing merkle root, "+
					"gap between sequence nums %d and %d, messages: %v", msgs[i-1].Header.SequenceNumber,
					msg.Header.SequenceNumber, msgs)
			}
		}

		msgHash, err := o.msgHasher.Hash(ctx, msg)
		if err != nil {
			msgID := hex.EncodeToString(msg.Header.MessageID[:])
			o.lggr.Warnw("failed to hash message", "msg", msg, "msg_id", msgID, "err", err)
			return cciptypes.Bytes32{}, fmt.Errorf("failed to hash message with id %s: %w", msgID, err)
		}

		hashes = append(hashes, msgHash)
	}

	// TODO: Do not hard code the hash function, it should be derived from the message hasher
	tree, err := merklemulti.NewTree(hashutil.NewKeccak(), hashes)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to construct merkle tree from %d leaves: %w", len(hashes), err)
	}

	root := tree.Root()
	o.lggr.Infow("computeMerkleRoot: Computed merkle root", "root", hex.EncodeToString(root[:]))

	return root, nil
}

func (o ObserverImpl) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := o.homeChain.GetFChain()
	if err != nil {
		// TODO: metrics
		o.lggr.Warnw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}
