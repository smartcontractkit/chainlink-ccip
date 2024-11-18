package merkleroot

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"sort"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/sync/errgroup"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/smartcontractkit/libocr/quorumhelper"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"
)

func (w *Processor) ObservationQuorum(
	_ context.Context, _ ocr3types.OutcomeContext, _ types.Query, aos []types.AttributedObservation,
) (bool, error) {
	// Across all chains we require at least 2F+1 observations.
	return quorumhelper.ObservationCountReachesObservationQuorum(
		quorumhelper.QuorumTwoFPlusOne, w.reportingCfg.N, w.reportingCfg.F, aos), nil
}

func (w *Processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	q Query,
) (Observation, error) {
	if err := w.initializeRMNController(ctx, prevOutcome); err != nil {
		return Observation{}, fmt.Errorf("initialize RMN controller: %w", err)
	}

	if err := w.verifyQuery(ctx, prevOutcome, q); err != nil {
		return Observation{}, fmt.Errorf("verify query: %w", err)
	}

	tStart := time.Now()
	observation, nextState := w.getObservation(ctx, q, prevOutcome)
	w.lggr.Infow("Sending MerkleRootObs",
		"observation", observation, "nextState", nextState, "observationDuration", time.Since(tStart))
	return observation, nil
}

// initializeRMNController initializes the RMN controller iff:
// 1. RMN is enabled.
// 2. RMN controller is not already initialized with the same cfg digest.
// 3. RMN remote config is available from previous outcome.
func (w *Processor) initializeRMNController(ctx context.Context, prevOutcome Outcome) error {
	if !w.offchainCfg.RMNEnabled {
		return nil
	}

	if prevOutcome.RMNRemoteCfg.IsEmpty() {
		w.lggr.Debug("RMN remote config is empty, skipping RMN controller initialization in this round")
		return nil
	}

	if prevOutcome.RMNRemoteCfg.ConfigDigest == w.rmnControllerCfgDigest {
		w.lggr.Debugw("RMN controller already initialized with the same config digest",
			"configDigest", w.rmnControllerCfgDigest)
		return nil
	}

	w.lggr.Infow("Initializing RMN controller", "rmnRemoteCfg", prevOutcome.RMNRemoteCfg)

	rmnNodesInfo, err := w.rmnHomeReader.GetRMNNodesInfo(prevOutcome.RMNRemoteCfg.ConfigDigest)
	if err != nil {
		return fmt.Errorf("failed to get RMN nodes info: %w", err)
	}

	oraclePeerIDs := make([]ragep2ptypes.PeerID, 0, len(w.oracleIDToP2pID))
	for _, p2pID := range w.oracleIDToP2pID {
		w.lggr.Infow("Adding oracle node to peerIDs", "p2pID", p2pID.String())
		oraclePeerIDs = append(oraclePeerIDs, p2pID)
	}

	if err := w.rmnController.InitConnection(
		ctx,
		cciptypes.Bytes32(w.reportingCfg.ConfigDigest),
		prevOutcome.RMNRemoteCfg.ConfigDigest,
		oraclePeerIDs,
		rmnNodesInfo,
	); err != nil {
		return fmt.Errorf("failed to init connection to RMN: %w", err)
	}

	w.rmnControllerCfgDigest = prevOutcome.RMNRemoteCfg.ConfigDigest

	return nil
}

// verifyQuery verifies the query based to the following rules.
// 1. If RMN is enabled, RMN signatures are required in the BuildingReport state but not expected in other states.
// 2. If RMN signatures are provided, they are verified against the current RMN node config.
func (w *Processor) verifyQuery(ctx context.Context, prevOutcome Outcome, q Query) error {
	if !w.offchainCfg.RMNEnabled {
		return nil
	}

	nextState := prevOutcome.NextState()

	skipVerification, err := shouldSkipRMNVerification(nextState, q, prevOutcome)
	if skipVerification {
		return nil
	}
	if err != nil {
		return err
	}

	ch, exists := chainsel.ChainBySelector(uint64(w.destChain))
	if !exists {
		return fmt.Errorf("failed to get chain by selector %d", w.destChain)
	}

	offRampAddress, err := w.ccipReader.GetContractAddress(consts.ContractNameOffRamp, w.destChain)
	if err != nil {
		return fmt.Errorf("failed to get offramp contract address: %w", err)
	}

	sigs, err := rmn.NewECDSASigsFromPB(q.RMNSignatures.Signatures)
	if err != nil {
		return fmt.Errorf("failed to convert signatures from protobuf: %w", err)
	}

	rmnRemoteCfg := prevOutcome.RMNRemoteCfg
	if rmnRemoteCfg.IsEmpty() {
		return fmt.Errorf("RMN remote config is not provided in the previous outcome")
	}

	signerAddresses := make([]cciptypes.UnknownAddress, 0, len(sigs))
	for _, rmnNode := range rmnRemoteCfg.Signers {
		signerAddresses = append(signerAddresses, rmnNode.OnchainPublicKey)
	}

	laneUpdates, err := rmn.NewLaneUpdatesFromPB(q.RMNSignatures.LaneUpdates)
	if err != nil {
		return fmt.Errorf("failed to convert lane updates from protobuf: %w", err)
	}

	rmnReport := cciptypes.RMNReport{
		ReportVersionDigest:         rmnRemoteCfg.RmnReportVersion,
		DestChainID:                 cciptypes.NewBigIntFromInt64(int64(ch.EvmChainID)),
		DestChainSelector:           cciptypes.ChainSelector(ch.Selector),
		RmnRemoteContractAddress:    rmnRemoteCfg.ContractAddress,
		OfframpAddress:              offRampAddress,
		RmnHomeContractConfigDigest: rmnRemoteCfg.ConfigDigest,
		LaneUpdates:                 laneUpdates,
	}

	if err := w.rmnCrypto.VerifyReportSignatures(ctx, sigs, rmnReport, signerAddresses); err != nil {
		return fmt.Errorf("failed to verify RMN signatures: %w", err)
	}
	return nil
}

// shouldSkipRMNVerification checks whether RMN verification should be skipped based on the current state and query.
func shouldSkipRMNVerification(nextState State, q Query, prevOutcome Outcome) (bool, error) {
	// Skip verification if RMN signatures are not expected in the current state.
	if nextState != BuildingReport && q.RMNSignatures == nil {
		return true, nil
	}

	// Skip verification if we are retrying RMN signatures in the next round.
	if nextState == BuildingReport && q.RetryRMNSignatures {
		return true, nil
	}

	// If in the BuildingReport state and RMN signatures are required but not provided, return an error.
	if nextState == BuildingReport && !q.RetryRMNSignatures && q.RMNSignatures == nil {
		return false, fmt.Errorf("RMN signatures are required in the BuildingReport state but not provided by leader")
	}

	// If in the BuildingReport state but RMN remote config is not available, return an error.
	if nextState == BuildingReport && prevOutcome.RMNRemoteCfg.IsEmpty() {
		return false, fmt.Errorf("RMN report config is not provided in the previous outcome")
	}

	// If RMN signatures are unexpectedly provided in a non-BuildingReport state, return an error.
	if nextState != BuildingReport && q.RMNSignatures != nil {
		return false, fmt.Errorf("RMN signatures are provided but not expected in the %d state", nextState)
	}

	// Proceed with RMN verification.
	return false, nil
}

func (w *Processor) getObservation(ctx context.Context, q Query, previousOutcome Outcome) (Observation, State) {
	nextState := previousOutcome.NextState()
	switch nextState {
	case SelectingRangesForReport:
		cursedChains := w.observer.ObserveCursedChains(ctx, w.destChain)
		if slices.Contains(cursedChains, w.destChain) {
			w.lggr.Warnw("destination chain is cursed, nothing to observe", "destChain", w.destChain)
			return Observation{}, SelectingRangesForReport
		}
		if len(cursedChains) > 0 {
			w.lggr.Warnw("some chains are cursed, ranges are not reported: %v", cursedChains)
		}
		offRampNextSeqNums := w.observer.ObserveOffRampNextSeqNums(ctx, cursedChains)
		onRampLatestSeqNums := w.observer.ObserveLatestOnRampSeqNums(ctx, w.destChain, cursedChains)
		rmnRemoteCfg := w.observer.ObserveRMNRemoteCfg(ctx, w.destChain)

		return Observation{
			OnRampMaxSeqNums:   onRampLatestSeqNums,
			OffRampNextSeqNums: offRampNextSeqNums,
			RMNRemoteConfig:    rmnRemoteCfg,
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
			OffRampNextSeqNums: w.observer.ObserveOffRampNextSeqNums(ctx, []cciptypes.ChainSelector{}),
			FChain:             w.observer.ObserveFChain(),
		}, nextState
	default:
		w.lggr.Errorw("Unexpected state", "state", nextState)
		return Observation{}, nextState
	}
}

type Observer interface {
	// ObserveOffRampNextSeqNums observes the next OffRamp sequence numbers for each source chain excluding cursed.
	ObserveOffRampNextSeqNums(ctx context.Context, cursedChains []cciptypes.ChainSelector) []plugintypes.SeqNumChain

	// ObserveLatestOnRampSeqNums observes the latest OnRamp sequence numbers for
	// each configured source chain excluding cursed.
	ObserveLatestOnRampSeqNums(
		ctx context.Context,
		destChain cciptypes.ChainSelector,
		cursedChains []cciptypes.ChainSelector,
	) []plugintypes.SeqNumChain

	// ObserveMerkleRoots computes the merkle roots for the given sequence number ranges
	ObserveMerkleRoots(ctx context.Context, ranges []plugintypes.ChainRange) []cciptypes.MerkleRootChain

	// ObserveRMNRemoteCfg observes the RMN remote config for the given destination chain
	ObserveRMNRemoteCfg(ctx context.Context, dstChain cciptypes.ChainSelector) rmntypes.RemoteConfig

	// ObserveCursedChains observes the chains that are cursed.
	// The results are sorted in ascending order and guaranteed to not contain duplicates.
	ObserveCursedChains(ctx context.Context, destChain cciptypes.ChainSelector) []cciptypes.ChainSelector

	ObserveFChain() map[cciptypes.ChainSelector]int
}

type observerImpl struct {
	lggr         logger.Logger
	homeChain    reader.HomeChain
	nodeID       commontypes.OracleID
	chainSupport plugincommon.ChainSupport
	ccipReader   readerpkg.CCIPReader
	msgHasher    cciptypes.MessageHasher
}

// ObserveOffRampNextSeqNums observes the next sequence numbers for each source chain from the OffRamp
func (o observerImpl) ObserveOffRampNextSeqNums(
	ctx context.Context, cursedChains []cciptypes.ChainSelector) []plugintypes.SeqNumChain {
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

	// Exclude the cursed chains from the source chains.
	chains := mapset.NewSet(sourceChains...).Difference(mapset.NewSet(cursedChains...)).ToSlice()

	sort.Slice(chains, func(i, j int) bool { return chains[i] < chains[j] })
	offRampNextSeqNums, err := o.ccipReader.NextSeqNum(ctx, chains)
	if err != nil {
		o.lggr.Warnw("call to NextSeqNum failed", "err", err)
		return nil
	}

	if len(offRampNextSeqNums) != len(chains) {
		o.lggr.Errorf("call to NextSeqNum returned unexpected number of seq nums, got %d, expected %d",
			len(offRampNextSeqNums), len(chains))
		return nil
	}

	result := make([]plugintypes.SeqNumChain, len(chains))
	for i := range chains {
		result[i] = plugintypes.SeqNumChain{ChainSel: chains[i], SeqNum: offRampNextSeqNums[i]}
	}

	return result
}

// ObserveLatestOnRampSeqNums observes the latest onRamp sequence numbers for each configured source chain.
func (o observerImpl) ObserveLatestOnRampSeqNums(
	ctx context.Context, destChain cciptypes.ChainSelector, cursedChains []cciptypes.ChainSelector,
) []plugintypes.SeqNumChain {
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

	sourceChains := mapset.NewSet(allSourceChains...).
		Intersect(supportedChains).
		Difference(mapset.NewSet(cursedChains...)).ToSlice()

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
func (o observerImpl) ObserveMerkleRoots(
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
		chainRange := chainRange
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

				onRampAddress, err := o.ccipReader.GetContractAddress(consts.ContractNameOnRamp, chainRange.ChainSel)
				if err != nil {
					o.lggr.Warnw(
						fmt.Sprintf("getting onramp contract address failed for selector %d", chainRange.ChainSel),
						"err", err,
						"chainSelector", chainRange.ChainSel,
					)
					return
				}

				merkleRoot := cciptypes.MerkleRootChain{
					ChainSel:      chainRange.ChainSel,
					SeqNumsRange:  chainRange.SeqNumRange,
					OnRampAddress: onRampAddress,
					MerkleRoot:    root,
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
func (o observerImpl) computeMerkleRoot(ctx context.Context, msgs []cciptypes.Message) (cciptypes.Bytes32, error) {
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
			o.lggr.Warnw("failed to hash message", "msg", msg, "err", err)
			return cciptypes.Bytes32{}, fmt.Errorf("hash message with id %s: %w", msg.Header.MessageID, err)
		}

		hashes = append(hashes, msgHash)
	}

	// TODO: Do not hard code the hash function, it should be derived from the message hasher
	tree, err := merklemulti.NewTree(hashutil.NewKeccak(), hashes)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to construct merkle tree from %d leaves: %w", len(hashes), err)
	}

	hashesStr := make([]string, len(hashes))
	for i, h := range hashes {
		hashesStr[i] = cciptypes.Bytes32(h).String()
	}
	root := tree.Root()
	o.lggr.Infow("Computed merkle root", "hashes", hashesStr, "root", cciptypes.Bytes32(root).String())
	return root, nil
}

func (o observerImpl) ObserveRMNRemoteCfg(
	ctx context.Context,
	dstChain cciptypes.ChainSelector) rmntypes.RemoteConfig {
	rmnRemoteCfg, err := o.ccipReader.GetRMNRemoteConfig(ctx, dstChain)
	if err != nil {
		if errors.Is(err, readerpkg.ErrContractReaderNotFound) {
			// destination chain not supported
			return rmntypes.RemoteConfig{}
		}
		// legitimate error
		o.lggr.Errorw("call to GetRMNRemoteConfig failed", "err", err)
		return rmntypes.RemoteConfig{}
	}
	return rmnRemoteCfg
}

// ObserveCursedChains observes the cursed chains for the current node.
// The results are sorted in ascending order and guaranteed to not contain duplicates.
func (o observerImpl) ObserveCursedChains(
	ctx context.Context, destChain cciptypes.ChainSelector) []cciptypes.ChainSelector {
	sourceChains, err := o.chainSupport.KnownSourceChainsSlice()
	if err != nil {
		o.lggr.Warnw("call to KnownSourceChainsSlice failed", "err", err)
		return nil
	}

	allChains := append(sourceChains, destChain)

	wg := sync.WaitGroup{}
	cursedChains := mapset.NewSet[cciptypes.ChainSelector]() // thread-safe by default

	for _, chain := range allChains {
		chain := chain
		wg.Add(1)
		go func() {
			defer wg.Done()

			isCursed, err := o.ccipReader.IsRMNRemoteCursed(ctx, chain)
			if err != nil {
				o.lggr.Errorw("call to IsChainCursed failed", "err", err)
				return
			}
			if isCursed {
				cursedChains.Add(chain)
			}
		}()
	}

	wg.Wait()

	cursedChainsSlice := cursedChains.ToSlice()
	sort.Slice(cursedChainsSlice, func(i, j int) bool { return cursedChainsSlice[i] < cursedChainsSlice[j] })
	return cursedChainsSlice
}

func (o observerImpl) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := o.homeChain.GetFChain()
	if err != nil {
		// TODO: metrics
		o.lggr.Warnw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}

// interface compliance check
var _ Observer = observerImpl{}
