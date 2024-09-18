package merkleroot

import (
	"context"
	"encoding/hex"
	"fmt"
	"sort"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
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
	if err := w.verifyQuery(ctx, prevOutcome, q); err != nil {
		return Observation{}, fmt.Errorf("verify query: %w", err)
	}

	tStart := time.Now()
	observation, nextState := w.getObservation(ctx, q, prevOutcome)
	w.lggr.Infow("Sending MerkleRootObs",
		"observation", observation, "nextState", nextState, "observationDuration", time.Since(tStart))
	return observation, nil
}

// verifyQuery verifies the query based to the following rules.
// 1. If RMN is enabled, RMN signatures are required in the BuildingReport state but not expected in other states.
// 2. If RMN signatures are provided, they are verified against the current RMN node config.
func (w *Processor) verifyQuery(ctx context.Context, prevOutcome Outcome, q Query) error {
	if !w.cfg.RMNEnabled {
		return nil
	}

	nextState := prevOutcome.NextState()

	// If we are in the BuildingReport state, and we are not retrying RMN signatures in the next round, we expect RMN
	// signatures to be provided by the leader.
	if nextState == BuildingReport && !q.RetryRMNSignatures && q.RMNSignatures == nil {
		return fmt.Errorf("RMN signatures are required in the BuildingReport state but not provided by leader")
	}

	// If we are not in the BuildingReport state, we do not expect RMN signatures to be provided.
	if nextState != BuildingReport && q.RMNSignatures != nil {
		return fmt.Errorf("RMN signatures are provided but not expected in the %d state", nextState)
	}

	ch, exists := chainsel.ChainBySelector(uint64(w.cfg.DestChain))
	if !exists {
		return fmt.Errorf("failed to get chain by selector %d", w.cfg.DestChain)
	}

	offRampAddress, err := w.ccipReader.GetContractAddress(consts.ContractNameOffRamp, w.cfg.DestChain)
	if err != nil {
		return fmt.Errorf("failed to get offramp contract address: %w", err)
	}

	sigs, err := rmn.NewECDSASigsFromPB(q.RMNSignatures.Signatures)
	if err != nil {
		return fmt.Errorf("failed to convert signatures from protobuf: %w", err)
	}

	signerAddresses := make([]cciptypes.Bytes, 0, len(sigs))
	for _, rmnNode := range w.rmnConfig.Home.RmnNodes {
		signerAddresses = append(signerAddresses, rmnNode.SignReportsAddress)
	}

	laneUpdates, err := rmn.NewLaneUpdatesFromPB(q.RMNSignatures.LaneUpdates)
	if err != nil {
		return fmt.Errorf("failed to convert lane updates from protobuf: %w", err)
	}

	rmnReport := cciptypes.RMNReport{
		ReportVersion:               w.rmnConfig.Home.RmnReportVersion,
		DestChainID:                 cciptypes.NewBigIntFromInt64(int64(ch.EvmChainID)),
		DestChainSelector:           cciptypes.ChainSelector(ch.Selector),
		RmnRemoteContractAddress:    w.rmnConfig.Remote.ContractAddress,
		OfframpAddress:              offRampAddress,
		RmnHomeContractConfigDigest: w.rmnConfig.Home.ConfigDigest,
		LaneUpdates:                 laneUpdates,
	}

	if err := w.rmnCrypto.VerifyReportSignatures(ctx, sigs, rmnReport, signerAddresses); err != nil {
		return fmt.Errorf("failed to verify RMN signatures: %w", err)
	}
	return nil
}

func (w *Processor) getObservation(ctx context.Context, q Query, previousOutcome Outcome) (Observation, State) {
	nextState := previousOutcome.NextState()
	switch nextState {
	case SelectingRangesForReport:
		offRampNextSeqNums := w.observer.ObserveOffRampNextSeqNums(ctx)
		onRampLatestSeqNums := w.observer.ObserveLatestOnRampSeqNums(ctx, w.cfg.DestChain)

		return Observation{
			OnRampMaxSeqNums:   onRampLatestSeqNums,
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
	// ObserveOffRampNextSeqNums observes the next OffRamp sequence numbers for each source chain
	ObserveOffRampNextSeqNums(ctx context.Context) []plugintypes.SeqNumChain

	// ObserveLatestOnRampSeqNums observes the latest OnRamp sequence numbers for each configured source chain.
	ObserveLatestOnRampSeqNums(ctx context.Context, destChain cciptypes.ChainSelector) []plugintypes.SeqNumChain

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

// ObserveLatestOnRampSeqNums observes the latest onRamp sequence numbers for each configured source chain.
func (o ObserverImpl) ObserveLatestOnRampSeqNums(
	ctx context.Context, destChain cciptypes.ChainSelector) []plugintypes.SeqNumChain {

	allSourceChains, err := o.chainSupport.KnownSourceChainsSlice()
	if err != nil {
		o.lggr.Warnw("call to KnownSourceChainsSlice failed", "err", err)
		return nil
	}

	supportedChains, err := o.chainSupport.SupportedChains(o.nodeID)
	if err != nil {
		o.lggr.Warnw("call to KnownSourceChainsSlice failed", "err", err)
		return nil
	}

	sourceChains := mapset.NewSet(allSourceChains...).Intersect(supportedChains).ToSlice()
	sort.Slice(sourceChains, func(i, j int) bool { return sourceChains[i] < sourceChains[j] })

	latestOnRampSeqNums := make([]plugintypes.SeqNumChain, len(sourceChains))
	eg := &errgroup.Group{}

	for i, sourceChain := range sourceChains {
		i, sourceChain := i, sourceChain
		eg.Go(func() error {
			nextOnRampSeqNum, err := o.ccipReader.GetExpectedNextSequenceNumber(ctx, sourceChain, destChain)
			if err != nil {
				return fmt.Errorf("failed to get expected next sequence number for source chain %d: %w", sourceChain, err)
			}
			if nextOnRampSeqNum == 0 {
				return fmt.Errorf("expected next sequence number for source chain %d is 0", sourceChain)
			}

			latestOnRampSeqNums[i] = plugintypes.SeqNumChain{
				ChainSel: sourceChain,
				SeqNum:   nextOnRampSeqNum - 1, // Latest is the next one minus one.
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		o.lggr.Warnw("call to GetExpectedNextSequenceNumber failed", "err", err)
		return nil
	}

	return latestOnRampSeqNums
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
