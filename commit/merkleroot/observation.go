package merkleroot

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"

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
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"

	ragep2ptypes "github.com/smartcontractkit/libocr/ragep2p/types"
)

var ErrSignaturesNotProvidedByLeader = errors.New("rmn signatures were not provided by the leader, " +
	"in most cases this indicates that the RMN nodes did not include any chain in their response")

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

const SendingObservation = "sending merkle root processor observation"

// Observation makes external calls to observe information according to the current processor state.
// According to the state it either observes sequence numbers, root hashes, RMN remote config, etc...
func (p *Processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	q Query,
) (Observation, error) {
	lggr := logutil.WithContextValues(ctx, p.lggr)

	if err := p.prepareRMNController(ctx, lggr, prevOutcome); err != nil {
		return Observation{}, fmt.Errorf("initialize RMN controller: %w", err)
	}

	if err := p.verifyQuery(ctx, prevOutcome, q); err != nil {
		if errors.Is(err, ErrSignaturesNotProvidedByLeader) {
			lggr.Warnw("RMN signatures not available, returning only fChain", "err", err)
			return Observation{
				// We observe fChain to avoid errors in the outcome phase.
				FChain: p.observer.ObserveFChain(ctx),
			}, nil
		}
		return Observation{}, fmt.Errorf("verify query: %w", err)
	}

	tStart := time.Now()
	observation, nextState, err := p.getObservation(ctx, lggr, q, prevOutcome)
	if err != nil {
		return Observation{}, fmt.Errorf("get observation: %w", err)
	}
	lggr.Infow(SendingObservation,
		"observation", observation, "nextState", nextState, "observationDuration", time.Since(tStart))
	return observation, nil
}

// prepareRMNController initializes the RMN controller iff:
// 1. RMN is enabled.
// 2. RMN controller is not already initialized with the same cfg digest.
// 3. RMN remote config is available from previous outcome.
func (p *Processor) prepareRMNController(ctx context.Context, lggr logger.Logger, prevOutcome Outcome) error {
	if !p.offchainCfg.RMNEnabled {
		return nil
	}

	if prevOutcome.RMNRemoteCfg.IsEmpty() {
		lggr.Debug("RMN remote config is empty, skipping RMN controller initialization in this round")
		return nil
	}

	if prevOutcome.RMNRemoteCfg.ConfigDigest == p.rmnControllerCfgDigest {
		lggr.Debugw("RMN controller already initialized with the same config digest",
			"configDigest", p.rmnControllerCfgDigest)
		return nil
	}

	lggr.Infow("Initializing RMN controller", "rmnRemoteCfg", prevOutcome.RMNRemoteCfg)

	rmnNodesInfo, err := p.rmnHomeReader.GetRMNNodesInfo(prevOutcome.RMNRemoteCfg.ConfigDigest)
	if err != nil {
		return fmt.Errorf("failed to get RMN nodes info: %w", err)
	}

	oraclePeerIDs := make([]ragep2ptypes.PeerID, 0, len(p.oracleIDToP2pID))
	for _, p2pID := range p.oracleIDToP2pID {
		lggr.Infow("Adding oracle node to peerIDs", "p2pID", p2pID.String())
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
	emptySigs := !q.ContainsRmnSignatures()

	switch nextState {
	case buildingReport:
		if q.RetryRMNSignatures {
			if emptySigs {
				return true, nil
			}
			return false, fmt.Errorf("RMN signatures are provided but not expected if retrying is set to true")
		}

		if prevOutcome.RMNRemoteCfg.IsEmpty() {
			return false, fmt.Errorf("RMN report config is not provided from the previous outcome")
		}

		// we don't want to check for empty sigs since at this point we don't know which chains are RMN-disabled.
		// if signatures are missing for specific chains they will be caught in the outcome phase.

		return false, nil
	default:
		if emptySigs {
			return true, nil // Sigs not expected
		}
		// If RMN signatures are unexpectedly provided in a non-BuildingReport state, return an error.
		return false, fmt.Errorf("RMN signatures are provided but not expected in the %d state", nextState)
	}
}

func (p *Processor) getObservation(
	ctx context.Context, lggr logger.Logger, q Query, previousOutcome Outcome) (Observation, processorState, error) {
	nextState := previousOutcome.nextState()
	switch nextState {
	case selectingRangesForReport:
		return Observation{
			OnRampMaxSeqNums:   p.observer.ObserveLatestOnRampSeqNums(ctx),
			OffRampNextSeqNums: p.observer.ObserveOffRampNextSeqNums(ctx),
			RMNRemoteConfig:    p.observer.ObserveRMNRemoteCfg(ctx),
			FChain:             p.observer.ObserveFChain(ctx),
		}, nextState, nil
	case buildingReport:
		if q.RetryRMNSignatures {
			// RMN signature computation failed, we only want to retry getting the RMN signatures in the next round.
			// So there's nothing to observe except for fChain, i.e. we don't want to build the report yet.
			return Observation{
				// We observe fChain to avoid errors in the outcome phase.
				// We check q.RetryRMNSignatures there and return the appropriate state and outcome
				// in order to retry.
				FChain: p.observer.ObserveFChain(ctx),
			}, nextState, nil
		}

		rmnEnabledChains := make(map[cciptypes.ChainSelector]bool)

		if p.offchainCfg.RMNEnabled {
			var err error
			rmnEnabledChains, err = p.rmnHomeReader.GetRMNEnabledSourceChains(previousOutcome.RMNRemoteCfg.ConfigDigest)
			if err != nil {
				return Observation{}, nextState, fmt.Errorf("failed to get RMN enabled source chains for %s: %w",
					previousOutcome.RMNRemoteCfg.ConfigDigest.String(), err)
			}
			lggr.Debugw("fetched RMN-enabled chains from rmnHome", "rmnEnabledChains", rmnEnabledChains)
		}

		return Observation{
			MerkleRoots:      p.observer.ObserveMerkleRoots(ctx, previousOutcome.RangesSelectedForReport),
			FChain:           p.observer.ObserveFChain(ctx),
			RMNEnabledChains: rmnEnabledChains,
		}, nextState, nil
	case waitingForReportTransmission:
		return Observation{
			OffRampNextSeqNums: p.observer.ObserveOffRampNextSeqNums(ctx),
			FChain:             p.observer.ObserveFChain(ctx),
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
	// NOTE: Make sure that caller supports the destination chain.
	ObserveLatestOnRampSeqNums(ctx context.Context) []plugintypes.SeqNumChain

	// ObserveMerkleRoots computes and returns the merkle roots for the provided sequence number ranges.
	// NOTE: Make sure that caller supports the provided chains.
	ObserveMerkleRoots(ctx context.Context, ranges []plugintypes.ChainRange) []cciptypes.MerkleRootChain

	// ObserveRMNRemoteCfg observes the RMN remote config from the configured destination chain.
	// Check implementation specific details to learn if external calls are made, if values are cached, etc...
	// NOTE: Make sure that caller supports the destination chain.
	ObserveRMNRemoteCfg(ctx context.Context) cciptypes.RemoteConfig

	// ObserveFChain observes the FChain for each supported chain. Check implementation specific details to learn
	// if external calls are made, if values are cached, etc...
	// NOTE: You can assume that every oracle can call this method, since data are fetched from home chain.
	ObserveFChain(ctx context.Context) map[cciptypes.ChainSelector]int

	// Close closes the observer and releases any resources.
	Close() error
}

// asyncObserver is an Observer implementation that fetches the data asynchronously.
type asyncObserver struct {
	lggr                logger.Logger
	syncObserver        observerImpl
	cancelFunc          func()
	mu                  *sync.RWMutex
	offRampNextSeqNums  []plugintypes.SeqNumChain
	onRampLatestSeqNums []plugintypes.SeqNumChain
}

// newAsyncObserver creates a new asyncObserver.
// It fetches the data from the base observer asynchronously and caches the results.
// It fetches the data every tickDur and uses a timeout of syncTimeout to kill long RPC calls.
func newAsyncObserver(lggr logger.Logger, observer observerImpl, tickDur, syncTimeout time.Duration) *asyncObserver {
	ctx, cf := context.WithCancel(context.Background())

	o := &asyncObserver{
		lggr:                lggr,
		syncObserver:        observer,
		cancelFunc:          cf,
		mu:                  &sync.RWMutex{},
		offRampNextSeqNums:  make([]plugintypes.SeqNumChain, 0),
		onRampLatestSeqNums: make([]plugintypes.SeqNumChain, 0),
	}

	ticker := time.NewTicker(tickDur)
	lggr.Debugw("async observer started", "tickDur", tickDur, "syncTimeout", syncTimeout)
	o.start(ctx, ticker.C, syncTimeout)

	o.cancelFunc = func() {
		cf()
		ticker.Stop()
	}

	return o
}

func (o *asyncObserver) start(ctx context.Context, ticker <-chan time.Time, syncTimeout time.Duration) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker:
				o.sync(ctx, syncTimeout)
			}
		}
	}()
}

func (o *asyncObserver) sync(ctx context.Context, syncTimeout time.Duration) {
	o.lggr.Debugw("async observer is syncing", "syncTimeout", syncTimeout)
	ctxSync, cf := context.WithTimeout(ctx, syncTimeout)
	defer cf()

	syncOps := []struct {
		id string
		op func(context.Context)
	}{
		{
			id: "offRampNextSeqNums",
			op: func(ctx context.Context) {
				offRampNext := o.syncObserver.ObserveOffRampNextSeqNums(ctxSync)
				o.mu.Lock()
				o.offRampNextSeqNums = offRampNext
				o.mu.Unlock()
			},
		},
		{
			id: "onRampLatestSeqNums",
			op: func(ctx context.Context) {
				onRampLast := o.syncObserver.ObserveLatestOnRampSeqNums(ctxSync)
				o.mu.Lock()
				o.onRampLatestSeqNums = onRampLast
				o.mu.Unlock()
			},
		},
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(syncOps))
	for _, op := range syncOps {
		go o.applySyncOp(ctxSync, o.lggr, op.id, wg, op.op)
	}
	wg.Wait()
}

// applySyncOp applies the given operation synchronously.
func (o *asyncObserver) applySyncOp(
	ctx context.Context, lggr logger.Logger, id string, wg *sync.WaitGroup, op func(ctx context.Context)) {
	defer wg.Done()
	tStart := time.Now()
	o.lggr.Debugw("async observer applying sync operation", "id", id)
	op(ctx)
	lggr.Debugw("async observer has applied the sync operation",
		"id", id, "duration", time.Since(tStart))
}

// ObserveOffRampNextSeqNums observes the next sequence numbers for each source chain from the OffRamp.
// Values are fetched from observers state which are fetched async.
func (o *asyncObserver) ObserveOffRampNextSeqNums(_ context.Context) []plugintypes.SeqNumChain {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.offRampNextSeqNums
}

// ObserveLatestOnRampSeqNums observes the latest onRamp sequence numbers for each configured source chain.
// Values are fetched from observers state which are fetched async.
func (o *asyncObserver) ObserveLatestOnRampSeqNums(_ context.Context) []plugintypes.SeqNumChain {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.onRampLatestSeqNums
}

// ObserveMerkleRoots observes the merkle roots for the given sequence number ranges.
// It directly calls the base observer since this values cannot be known in advance.
func (o *asyncObserver) ObserveMerkleRoots(
	ctx context.Context, ranges []plugintypes.ChainRange) []cciptypes.MerkleRootChain {
	return o.syncObserver.ObserveMerkleRoots(ctx, ranges)
}

// ObserveRMNRemoteCfg observes the RMN Remote Config by directly calling the base observer since this value is cached.
func (o *asyncObserver) ObserveRMNRemoteCfg(ctx context.Context) cciptypes.RemoteConfig {
	return o.syncObserver.ObserveRMNRemoteCfg(ctx)
}

// ObserveFChain observes the FChain by directly calling the base observer since this value is cached.
func (o *asyncObserver) ObserveFChain(ctx context.Context) map[cciptypes.ChainSelector]int {
	return o.syncObserver.ObserveFChain(ctx)
}

// Close closes the observer and releases any resources.
func (o *asyncObserver) Close() error {
	if o.cancelFunc != nil {
		o.cancelFunc()
		o.cancelFunc = nil
	}
	return nil
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
	lggr := logutil.WithContextValues(ctx, o.lggr)

	supportsDestChain, err := o.chainSupport.SupportsDestChain(o.oracleID)
	if err != nil {
		lggr.Warnw("call to SupportsDestChain failed", "err", err)
		return nil
	}

	if !supportsDestChain {
		lggr.Debugw("cannot observe off ramp seq nums since destination chain is not supported")
		return nil
	}

	allSourceChains, err := o.chainSupport.KnownSourceChainsSlice()
	if err != nil {
		lggr.Warnw("call to KnownSourceChainsSlice failed", "err", err)
		return nil
	}

	curseInfo, err := o.ccipReader.GetRmnCurseInfo(ctx)
	if err != nil {
		lggr.Errorw("nothing to observe: rmn read error",
			"err", err,
			"curseInfo", curseInfo,
			"sourceChains", allSourceChains,
		)
		return nil
	}
	if curseInfo.GlobalCurse || curseInfo.CursedDestination {
		lggr.Warnw("nothing to observe: rmn curse", "curseInfo", curseInfo)
		return nil
	}

	sourceChains := curseInfo.NonCursedSourceChains(allSourceChains)
	if len(sourceChains) == 0 {
		lggr.Warnw(
			"nothing to observe from the offRamp, no active source chains exist",
			"curseInfo", curseInfo)
		return nil
	}

	offRampNextSeqNums, err := o.ccipReader.NextSeqNum(ctx, sourceChains)
	if err != nil {
		lggr.Warnw("call to NextSeqNum failed", "err", err)
		return nil
	}

	result := make([]plugintypes.SeqNumChain, 0, len(sourceChains))
	for chainSelector, seqNum := range offRampNextSeqNums {
		result = append(result, plugintypes.NewSeqNumChain(chainSelector, seqNum))
	}

	sort.Slice(result, func(i, j int) bool { return result[i].ChainSel < result[j].ChainSel })
	return result
}

// ObserveLatestOnRampSeqNums observes the latest onRamp sequence numbers for each configured source chain.
func (o observerImpl) ObserveLatestOnRampSeqNums(ctx context.Context) []plugintypes.SeqNumChain {
	lggr := logutil.WithContextValues(ctx, o.lggr)

	allSourceChains, err := o.chainSupport.KnownSourceChainsSlice()
	if err != nil {
		lggr.Warnw("call to KnownSourceChainsSlice failed", "err", err)
		return nil
	}

	sourceChainsConfig, err := o.ccipReader.GetOffRampSourceChainsConfig(ctx, allSourceChains)
	if err != nil {
		lggr.Errorw("get offRamp source chains config failed", "err", err)
		return nil
	}

	mu := &sync.Mutex{}
	latestOnRampSeqNums := make([]plugintypes.SeqNumChain, 0, len(sourceChainsConfig))

	wg := &sync.WaitGroup{}
	for sourceChain, cfg := range sourceChainsConfig {
		if !cfg.IsEnabled {
			lggr.Debugw("ObserveLatestOnRampSeqNums source chain is disabled, skipping", "chain", sourceChain)
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			latestOnRampSeqNum, err := o.ccipReader.LatestMsgSeqNum(ctx, sourceChain)
			if err != nil {
				lggr.Errorf("failed to get latest msg seq num for source chain %d: %s", sourceChain, err)
				return
			}

			mu.Lock()
			latestOnRampSeqNums = append(
				latestOnRampSeqNums,
				plugintypes.NewSeqNumChain(sourceChain, latestOnRampSeqNum),
			)
			mu.Unlock()
		}()
	}
	wg.Wait()

	sort.Slice(latestOnRampSeqNums, func(i, j int) bool {
		return latestOnRampSeqNums[i].ChainSel < latestOnRampSeqNums[j].ChainSel
	})
	lggr.Debugw("fetched latestOnRampSeqNums", "seqNums", latestOnRampSeqNums)
	return latestOnRampSeqNums
}

// ObserveMerkleRoots computes the merkle roots for the given sequence number ranges
func (o observerImpl) ObserveMerkleRoots(
	ctx context.Context,
	ranges []plugintypes.ChainRange,
) []cciptypes.MerkleRootChain {
	lggr := logutil.WithContextValues(ctx, o.lggr)

	supportedChains, err := o.chainSupport.SupportedChains(o.oracleID)
	if err != nil {
		lggr.Warnw("call to supportedChains failed", "err", err)
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
					lggr.Warnw("call to MsgsBetweenSeqNums failed", "err", err)
					return
				}

				if uint64(len(msgs)) != uint64(chainRange.SeqNumRange.End()-chainRange.SeqNumRange.Start()+1) {
					lggr.Warnw("call to MsgsBetweenSeqNums returned unexpected number of messages, chain skipped",
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
						lggr.Warnw("message sequence number does not match seqNum range, chain skipped",
							"chain", chainRange.ChainSel,
							"seqNum", seqNum,
							"msgSeqNum", msgSeqNum,
							"range", chainRange.SeqNumRange,
						)
						return
					}
					msgIdx++
				}

				root, err := o.computeMerkleRoot(ctx, lggr, msgs)
				if err != nil {
					lggr.Warnw("call to computeMerkleRoot failed", "err", err)
					return
				}

				onRampAddress, err := o.ccipReader.GetContractAddress(consts.ContractNameOnRamp, chainRange.ChainSel)
				if err != nil {
					lggr.Warnw(
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
func (o observerImpl) computeMerkleRoot(
	ctx context.Context,
	lggr logger.Logger,
	msgs []cciptypes.Message,
) (cciptypes.Bytes32, error) {
	hashes := make([][32]byte, len(msgs))
	hashesStr := make([]string, len(hashes)) // also keep hashes as strings for logging purposes

	eg := &errgroup.Group{}
	for i, msg := range msgs {
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
				lggr.Warnw("failed to hash message", "message", msg, "err", err)
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
	lggr.Infow("Computed merkle root", "hashes", hashesStr, "root", cciptypes.Bytes32(root).String())
	return root, nil
}

// ObserveRMNRemoteCfg observes the RMN remote config for the given destination chain.
// NOTE: At least two external calls are made.
func (o observerImpl) ObserveRMNRemoteCfg(ctx context.Context) cciptypes.RemoteConfig {
	lggr := logutil.WithContextValues(ctx, o.lggr)

	rmnRemoteCfg, err := o.ccipReader.GetRMNRemoteConfig(ctx)
	if err != nil {
		if errors.Is(err, readerpkg.ErrContractReaderNotFound) {
			// destination chain not supported
			return cciptypes.RemoteConfig{}
		}
		// legitimate error
		lggr.Errorw("call to GetRMNRemoteConfig failed", "err", err)
		return cciptypes.RemoteConfig{}
	}
	return rmnRemoteCfg
}

// ObserveFChain observes the FChain for each supported chain.
// NOTE: It does not make any external calls, values are cached.
func (o observerImpl) ObserveFChain(ctx context.Context) map[cciptypes.ChainSelector]int {
	lggr := logutil.WithContextValues(ctx, o.lggr)

	fChain, err := o.homeChain.GetFChain()
	if err != nil {
		// TODO: metrics
		lggr.Errorw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}

func (o observerImpl) Close() error {
	return nil
}
