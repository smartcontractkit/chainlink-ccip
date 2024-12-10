package merkleroot

import (
	"context"
	"errors"
	"fmt"
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

// ObservationQuorum requires "across all chains" at least 2F+1 observations.
func (p *Processor) ObservationQuorum(
	_ context.Context, _ ocr3types.OutcomeContext, _ types.Query, aos []types.AttributedObservation,
) (bool, error) {
	return quorumhelper.ObservationCountReachesObservationQuorum(
		quorumhelper.QuorumTwoFPlusOne,
		p.reportingCfg.N,
		p.reportingCfg.F,
		aos,
	), nil
}

// Observation makes external calls to observe information according to the current processor state.
// According to the state it either observes sequence numbers, root hashes, RMN remote config, etc...
func (p *Processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	q Query,
) (Observation, error) {
	if err := p.prepareRMNController(ctx, prevOutcome); err != nil {
		return Observation{}, fmt.Errorf("initialize RMN controller: %w", err)
	}

	if err := p.verifyQuery(ctx, prevOutcome, q); err != nil {
		return Observation{}, fmt.Errorf("verify query: %w", err)
	}

	tStart := time.Now()
	observation, nextState, err := p.getObservation(ctx, q, prevOutcome)
	if err != nil {
		return Observation{}, fmt.Errorf("get observation: %w", err)
	}

	p.lggr.Infow("sending merkle root processor observation",
		"observation", observation, "nextState", nextState, "observationDuration", time.Since(tStart))
	return observation, nil
}

// prepareRMNController initializes the RMN controller iff:
// 1. RMN is enabled.
// 2. RMN controller is not already initialized with the same cfg digest.
// 3. RMN remote config is available from previous outcome.
func (p *Processor) prepareRMNController(ctx context.Context, prevOutcome Outcome) error {
	if !p.offchainCfg.RMNEnabled {
		return nil
	}

	if prevOutcome.RMNRemoteCfg.IsEmpty() {
		p.lggr.Debug("RMN remote config is empty, skipping RMN controller initialization in this round")
		return nil
	}

	if prevOutcome.RMNRemoteCfg.ConfigDigest == p.rmnControllerCfgDigest {
		p.lggr.Debugw("RMN controller already initialized with the same config digest",
			"configDigest", p.rmnControllerCfgDigest)
		return nil
	}

	p.lggr.Infow("Initializing RMN controller", "rmnRemoteCfg", prevOutcome.RMNRemoteCfg)

	rmnNodesInfo, err := p.rmnHomeReader.GetRMNNodesInfo(prevOutcome.RMNRemoteCfg.ConfigDigest)
	if err != nil {
		return fmt.Errorf("failed to get RMN nodes info: %w", err)
	}

	oraclePeerIDs := make([]ragep2ptypes.PeerID, 0, len(p.oracleIDToP2pID))
	for _, p2pID := range p.oracleIDToP2pID {
		p.lggr.Infow("Adding oracle node to peerIDs", "p2pID", p2pID.String())
		oraclePeerIDs = append(oraclePeerIDs, p2pID)
	}

	if err := p.rmnController.InitConnection(
		ctx,
		cciptypes.Bytes32(p.reportingCfg.ConfigDigest),
		prevOutcome.RMNRemoteCfg.ConfigDigest,
		oraclePeerIDs,
		rmnNodesInfo,
	); err != nil {
		return fmt.Errorf("failed to init connection to RMN: %w", err)
	}

	p.rmnControllerCfgDigest = prevOutcome.RMNRemoteCfg.ConfigDigest

	return nil
}

// verifyQuery verifies the query based on the following rules.
// 1. If RMN is enabled, RMN signatures are required in the BuildingReport state but not expected in other states.
// 2. If RMN signatures are provided, they are verified against the current RMN node configuration.
func (p *Processor) verifyQuery(ctx context.Context, prevOutcome Outcome, q Query) error {
	if !p.offchainCfg.RMNEnabled {
		return nil
	}

	nextState := prevOutcome.nextState()

	skipVerification, err := shouldSkipRMNVerification(nextState, q, prevOutcome)
	if skipVerification {
		return nil
	}
	if err != nil {
		return err
	}

	ch, exists := chainsel.ChainBySelector(uint64(p.destChain))
	if !exists {
		return fmt.Errorf("get chain by selector %d", p.destChain)
	}

	offRampAddress, err := p.ccipReader.GetContractAddress(consts.ContractNameOffRamp, p.destChain)
	if err != nil {
		return fmt.Errorf("get offramp contract address: %w", err)
	}

	sigs, err := rmn.NewECDSASigsFromPB(q.RMNSignatures.Signatures)
	if err != nil {
		return fmt.Errorf("parse protobuf signatures %v: %w", q.RMNSignatures.Signatures, err)
	}

	rmnRemoteCfg := prevOutcome.RMNRemoteCfg
	if rmnRemoteCfg.IsEmpty() {
		return fmt.Errorf("RMN remote configuration was not provided by the previous outcome")
	}

	signerAddresses := make([]cciptypes.UnknownAddress, 0, len(sigs))
	for _, rmnNode := range rmnRemoteCfg.Signers {
		signerAddresses = append(signerAddresses, rmnNode.OnchainPublicKey)
	}

	laneUpdates, err := rmn.NewLaneUpdatesFromPB(q.RMNSignatures.LaneUpdates)
	if err != nil {
		return fmt.Errorf("parse protobuf lane updates %v: %w", q.RMNSignatures.LaneUpdates, err)
	}

	rmnReport := cciptypes.NewRMNReport(
		rmnRemoteCfg.RmnReportVersion,
		cciptypes.NewBigIntFromInt64(int64(ch.EvmChainID)),
		cciptypes.ChainSelector(ch.Selector),
		rmnRemoteCfg.ContractAddress,
		offRampAddress,
		rmnRemoteCfg.ConfigDigest,
		laneUpdates,
	)

	if err := p.rmnCrypto.VerifyReportSignatures(ctx, sigs, rmnReport, signerAddresses); err != nil {
		return fmt.Errorf("failed to verify RMN signatures: %w", err)
	}
	return nil
}

// shouldSkipRMNVerification checks whether RMN verification should be skipped based on the current state and query.
func shouldSkipRMNVerification(nextState processorState, q Query, prevOutcome Outcome) (bool, error) {
	// Skip verification if RMN signatures are not expected in the current state.
	if nextState != buildingReport && q.RMNSignatures == nil {
		return true, nil
	}

	// Skip verification if we are retrying RMN signatures in the next round.
	if nextState == buildingReport && q.RetryRMNSignatures {
		if q.RMNSignatures != nil {
			return false, fmt.Errorf("RMN signatures are provided but not expected if retrying is set to true")
		}
		return true, nil
	}

	// If in the BuildingReport state and RMN signatures are required but not provided, return an error.
	if nextState == buildingReport && !q.RetryRMNSignatures && q.RMNSignatures == nil {
		return false, fmt.Errorf("RMN signatures are required in the BuildingReport state but not provided by leader")
	}

	// If in the BuildingReport state but RMN remote config is not available, return an error.
	if nextState == buildingReport && prevOutcome.RMNRemoteCfg.IsEmpty() {
		return false, fmt.Errorf("RMN report config is not provided from the previous outcome")
	}

	// If RMN signatures are unexpectedly provided in a non-BuildingReport state, return an error.
	if nextState != buildingReport && q.RMNSignatures != nil {
		return false, fmt.Errorf("RMN signatures are provided but not expected in the %d state", nextState)
	}

	// Proceed with RMN verification.
	return false, nil
}

func (p *Processor) getObservation(
	ctx context.Context, q Query, previousOutcome Outcome) (Observation, processorState, error) {
	nextState := previousOutcome.nextState()
	switch nextState {
	case selectingRangesForReport:
		offRampNextSeqNums := p.observer.ObserveOffRampNextSeqNums(ctx)
		onRampLatestSeqNums := p.observer.ObserveLatestOnRampSeqNums(ctx, p.destChain)
		rmnRemoteCfg := p.observer.ObserveRMNRemoteCfg(ctx, p.destChain)

		return Observation{
			OnRampMaxSeqNums:   onRampLatestSeqNums,
			OffRampNextSeqNums: offRampNextSeqNums,
			RMNRemoteConfig:    rmnRemoteCfg,
			FChain:             p.observer.ObserveFChain(),
		}, nextState, nil
	case buildingReport:
		if q.RetryRMNSignatures {
			// RMN signature computation failed, we only want to retry getting the RMN signatures in the next round.
			// So there's nothing to observe, i.e. we don't want to build the report yet.
			return Observation{}, nextState, nil
		}
		return Observation{
			MerkleRoots: p.observer.ObserveMerkleRoots(ctx, previousOutcome.RangesSelectedForReport),
			FChain:      p.observer.ObserveFChain(),
		}, nextState, nil
	case waitingForReportTransmission:
		return Observation{
			OffRampNextSeqNums: p.observer.ObserveOffRampNextSeqNums(ctx),
			FChain:             p.observer.ObserveFChain(),
		}, nextState, nil
	default:
		return Observation{},
			nextState,
			fmt.Errorf("unexpected nextState=%d with prevOutcome=%d", nextState, previousOutcome.OutcomeType)
	}
}

// Observer is an interface for observing data from the offRamp, onRamp, RMN remote config, etc...
type Observer interface {
	// ObserveOffRampNextSeqNums observes the next OffRamp sequence numbers for each source chain.
	// If the destination chain is cursed it returns nil or
	// if some source chain is cursed, it's skipped from the results.
	// NOTE: Make sure that caller supports the destination chain.
	ObserveOffRampNextSeqNums(ctx context.Context) []plugintypes.SeqNumChain

	// ObserveLatestOnRampSeqNums observes the latest OnRamp sequence numbers for each configured source chain.
	// NOTE: Make sure that caller supports the provided destination chain.
	ObserveLatestOnRampSeqNums(ctx context.Context, destChain cciptypes.ChainSelector) []plugintypes.SeqNumChain

	// ObserveMerkleRoots computes and returns the merkle roots for the provided sequence number ranges.
	// NOTE: Make sure that caller supports the provided chains.
	ObserveMerkleRoots(ctx context.Context, ranges []plugintypes.ChainRange) []cciptypes.MerkleRootChain

	// ObserveRMNRemoteCfg observes the RMN remote config for the given destination chain.
	// Check implementation specific details to learn if external calls are made, if values are cached, etc...
	// NOTE: Make sure that caller supports the provided destination chain.
	ObserveRMNRemoteCfg(ctx context.Context, dstChain cciptypes.ChainSelector) rmntypes.RemoteConfig

	// ObserveFChain observes the FChain for each supported chain. Check implementation specific details to learn
	// if external calls are made, if values are cached, etc...
	// NOTE: You can assume that every oracle can call this method, since data are fetched from home chain.
	ObserveFChain() map[cciptypes.ChainSelector]int
}

type observerImpl struct {
	lggr         logger.Logger
	homeChain    reader.HomeChain
	oracleID     commontypes.OracleID
	chainSupport plugincommon.ChainSupport
	ccipReader   readerpkg.CCIPReader
	msgHasher    cciptypes.MessageHasher
}

func newObserverImpl(
	lggr logger.Logger,
	homeChain reader.HomeChain,
	oracleID commontypes.OracleID,
	chainSupport plugincommon.ChainSupport,
	ccipReader readerpkg.CCIPReader,
	msgHasher cciptypes.MessageHasher,
) observerImpl {
	return observerImpl{
		lggr:         lggr,
		homeChain:    homeChain,
		oracleID:     oracleID,
		chainSupport: chainSupport,
		ccipReader:   ccipReader,
		msgHasher:    msgHasher,
	}
}

// ObserveOffRampNextSeqNums observes the next sequence numbers for each source chain from the OffRamp.
// If the destination chain is cursed it returns nil.
// If some source chain is cursed, it is not included in the results.
func (o observerImpl) ObserveOffRampNextSeqNums(ctx context.Context) []plugintypes.SeqNumChain {
	supportsDestChain, err := o.chainSupport.SupportsDestChain(o.oracleID)
	if err != nil {
		o.lggr.Warnw("call to SupportsDestChain failed", "err", err)
		return nil
	}

	if !supportsDestChain {
		o.lggr.Debugw("cannot observe off ramp seq nums since destination chain is not supported")
		return nil
	}

	allSourceChains, err := o.chainSupport.KnownSourceChainsSlice()
	if err != nil {
		o.lggr.Warnw("call to KnownSourceChainsSlice failed", "err", err)
		return nil
	}

	curseInfo, err := o.ccipReader.GetRmnCurseInfo(ctx, o.chainSupport.DestChain(), allSourceChains)
	if err != nil {
		o.lggr.Errorw("nothing to observe: rmn read error",
			"err", err,
			"curseInfo", curseInfo,
			"sourceChains", allSourceChains,
		)
		return nil
	}
	if curseInfo.GlobalCurse || curseInfo.CursedDestination {
		o.lggr.Warnw("nothing to observe: rmn curse", "curseInfo", curseInfo)
		return nil
	}

	sourceChains := curseInfo.NonCursedSourceChains(allSourceChains)
	if len(sourceChains) == 0 {
		o.lggr.Warnw(
			"nothing to observe from the offRamp, no active source chains exist",
			"curseInfo", curseInfo)
		return nil
	}

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
func (o observerImpl) ObserveLatestOnRampSeqNums(
	ctx context.Context, destChain cciptypes.ChainSelector) []plugintypes.SeqNumChain {

	allSourceChains, err := o.chainSupport.KnownSourceChainsSlice()
	if err != nil {
		o.lggr.Warnw("call to KnownSourceChainsSlice failed", "err", err)
		return nil
	}

	supportedChains, err := o.chainSupport.SupportedChains(o.oracleID)
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
func (o observerImpl) ObserveMerkleRoots(
	ctx context.Context,
	ranges []plugintypes.ChainRange,
) []cciptypes.MerkleRootChain {

	supportedChains, err := o.chainSupport.SupportedChains(o.oracleID)
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

				if uint64(len(msgs)) != uint64(chainRange.SeqNumRange.End()-chainRange.SeqNumRange.Start()+1) {
					o.lggr.Warnw("call to MsgsBetweenSeqNums returned unexpected number of messages chain skipped",
						"chain", chainRange.ChainSel,
						"range", chainRange.SeqNumRange,
						"expected", chainRange.SeqNumRange.End()-chainRange.SeqNumRange.Start()+1,
						"actual", len(msgs),
					)
					return
				}

				// If the returned messages do not match the sequence numbers range
				// there is nothing to observe for this chain since messages are missing.
				msgIdx := 0
				for seqNum := chainRange.SeqNumRange.Start(); seqNum <= chainRange.SeqNumRange.End(); seqNum++ {
					msgSeqNum := msgs[msgIdx].Header.SequenceNumber
					if msgSeqNum != seqNum {
						o.lggr.Warnw("message sequence number does not match seqNum range chain skipped",
							"chain", chainRange.ChainSel,
							"seqNum", seqNum,
							"msgSeqNum", msgSeqNum,
							"range", chainRange.SeqNumRange,
						)
						return
					}
					msgIdx++
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

// computeMerkleRoot computes the merkle root of a list of messages.
// Messages should be sorted by sequence number and not have any gaps.
func (o observerImpl) computeMerkleRoot(ctx context.Context, msgs []cciptypes.Message) (cciptypes.Bytes32, error) {
	hashes := make([][32]byte, len(msgs))
	hashesStr := make([]string, len(hashes)) // also keep hashes as strings for logging purposes

	eg := &errgroup.Group{}
	for i, msg := range msgs {
		i := i
		msg := msg
		eg.Go(func() error {
			// Assert there are no sequence number gaps in msgs
			if i > 0 {
				if msg.Header.SequenceNumber != msgs[i-1].Header.SequenceNumber+1 {
					return fmt.Errorf("found non-consecutive sequence numbers when computing merkle root, "+
						"gap between sequence nums %d and %d, messages: %v", msgs[i-1].Header.SequenceNumber,
						msg.Header.SequenceNumber, msgs)
				}
			}

			msgHash, err := o.msgHasher.Hash(ctx, msg)
			if err != nil {
				o.lggr.Warnw("failed to hash message", "msg", msg, "err", err)
				return fmt.Errorf("hash message with id %s: %w", msg.Header.MessageID, err)
			}

			hashes[i] = msgHash
			hashesStr[i] = msgHash.String()
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return [32]byte{}, err
	}

	// TODO: Do not hard code the hash function, it should be derived from the message hasher
	tree, err := merklemulti.NewTree(hashutil.NewKeccak(), hashes)
	if err != nil {
		return [32]byte{}, fmt.Errorf("failed to construct merkle tree from %d leaves: %w", len(hashes), err)
	}

	root := tree.Root()
	o.lggr.Infow("Computed merkle root", "hashes", hashesStr, "root", cciptypes.Bytes32(root).String())
	return root, nil
}

// ObserveRMNRemoteCfg observes the RMN remote config for the given destination chain.
// NOTE: At least two external calls are made.
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

// ObserveFChain observes the FChain for each supported chain.
// NOTE: It does not make any external calls, values are cached.
func (o observerImpl) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := o.homeChain.GetFChain()
	if err != nil {
		// TODO: metrics
		o.lggr.Errorw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}
