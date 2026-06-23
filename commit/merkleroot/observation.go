package merkleroot

import (
	"context"
	"errors"
	"fmt"
	"slices"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"golang.org/x/sync/errgroup"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/smartcontractkit/libocr/quorumhelper"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/merklemulti"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccip/consts"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/rpctimeout"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

const skippedLanesLogFrequency = 30 * time.Minute

var lastSkippedLanesLog atomic.Pointer[time.Time]

// isLiveOffRampSourceLane reports whether the offramp has an enabled inbound lane from sourceChain.
func isLiveOffRampSourceLane(cfg readerpkg.StaticSourceChainConfig, exists bool) bool {
	return exists && cfg.IsEnabled && len(cfg.OnRamp) > 0
}

// offRampLaneClassification buckets source chains by their offramp lane status.
type offRampLaneClassification struct {
	// live lanes are enabled, have an onRamp, and have RMN verification disabled (queryable).
	live []cciptypes.ChainSelector
	// skippedNotALane have no offramp config or no onRamp configured.
	skippedNotALane []cciptypes.ChainSelector
	// skippedDisabled have an offramp config but are disabled.
	skippedDisabled []cciptypes.ChainSelector
	// rmnMisconfigured are live lanes that unexpectedly have RMN verification enabled.
	rmnMisconfigured []cciptypes.ChainSelector
}

// classifyOffRampSourceLanes buckets the supported source chains based on their offramp source chain
// config. Only "live" lanes should be queried for onRamp seq nums, the other buckets are skipped and
// exist for logging.
func classifyOffRampSourceLanes(
	supported []cciptypes.ChainSelector,
	cfgs map[cciptypes.ChainSelector]readerpkg.StaticSourceChainConfig,
) offRampLaneClassification {
	var c offRampLaneClassification
	for _, sourceChain := range supported {
		cfg, ok := cfgs[sourceChain]
		switch {
		case !ok:
			c.skippedNotALane = append(c.skippedNotALane, sourceChain)
		case !cfg.IsEnabled:
			c.skippedDisabled = append(c.skippedDisabled, sourceChain)
		case len(cfg.OnRamp) == 0:
			c.skippedNotALane = append(c.skippedNotALane, sourceChain)
		case !cfg.IsRMNVerificationDisabled:
			c.rmnMisconfigured = append(c.rmnMisconfigured, sourceChain)
		default:
			c.live = append(c.live, sourceChain)
		}
	}
	return c
}

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
func (p *Processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	_ Query,
) (Observation, error) {
	lggr := logutil.WithContextValues(ctx, p.lggr)

	tStart := time.Now()
	observation, nextState, err := p.getObservation(ctx, prevOutcome)
	if err != nil {
		return Observation{}, fmt.Errorf("get observation: %w", err)
	}
	lggr.Infow(SendingObservation,
		"observation", observation, "nextState", nextState, "observationDuration", time.Since(tStart))
	return observation, nil
}

func (p *Processor) getObservation(
	ctx context.Context, previousOutcome Outcome) (Observation, processorState, error) {
	nextState := previousOutcome.nextState()
	switch nextState {
	case selectingRangesForReport:
		return Observation{
			OnRampMaxSeqNums:   p.observer.ObserveLatestOnRampSeqNums(ctx),
			OffRampNextSeqNums: p.observer.ObserveOffRampNextSeqNums(ctx),
			FChain:             p.observer.ObserveFChain(ctx),
		}, nextState, nil
	case buildingReport:
		return Observation{
			MerkleRoots: p.observer.ObserveMerkleRoots(ctx, previousOutcome.RangesSelectedForReport),
			FChain:      p.observer.ObserveFChain(ctx),
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

	// Helpful log for operators to know which source chains are being ignored due to being cursed on the destination.
	if len(curseInfo.CursedSourceChains) > 0 {
		lggr.Infow(
			"ignoring some cursed source chains, won't read their messages",
			"cursedSourceChainsToIgnore", curseInfo.CursedSourceChains,
			"sourceChainsToRead", sourceChains,
		)
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

	supportedChains, err := o.chainSupport.SupportedChains(o.oracleID)
	if err != nil {
		lggr.Warnw("call to SupportedChains failed", "err", err, "oracleID", o.oracleID)
		return nil
	}

	supportedSourceChains := mapset.NewSet(allSourceChains...).
		Intersect(supportedChains).ToSlice()

	slices.Sort(supportedSourceChains)

	configCtx, configCancel := rpctimeout.Context(ctx)
	defer configCancel()
	sourceChainsCfg, err := o.ccipReader.GetOffRampSourceChainsConfig(configCtx, supportedSourceChains)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			lggr.Warnw("call to GetOffRampSourceChainsConfig timed out", "err", err)
		} else {
			lggr.Warnw("call to GetOffRampSourceChainsConfig failed", "err", err)
		}
		return nil
	}

	classification := classifyOffRampSourceLanes(supportedSourceChains, sourceChainsCfg)

	if len(classification.skippedNotALane) > 0 || len(classification.skippedDisabled) > 0 {
		logutil.LogWhenExceedFrequency(&lastSkippedLanesLog, skippedLanesLogFrequency, func() {
			lggr.Debugw("skipping onRamp seq num observations for non-live lanes",
				"noConfigOrOnRamp", classification.skippedNotALane,
				"disabled", classification.skippedDisabled,
			)
		})
	}
	if len(classification.rmnMisconfigured) > 0 {
		lggr.Warnw("rmn enablement is misconfigured on these lanes, skipping observations",
			"sources", classification.rmnMisconfigured,
		)
	}

	mu := &sync.Mutex{}
	latestOnRampSeqNums := make([]plugintypes.SeqNumChain, 0, len(classification.live))

	wg := &sync.WaitGroup{}
	for _, sourceChain := range classification.live {
		wg.Go(func() {
			chainCtx, chainCancel := rpctimeout.Context(ctx)
			defer chainCancel()
			latestOnRampSeqNum, err := o.ccipReader.LatestMsgSeqNum(chainCtx, sourceChain)
			if err != nil {
				if isNoBindingsError(err) {
					// when a source chain is disabled there will not be a binding for the onRamp contract
					// we don't want to log this as an error.
					lggr.Debugw("no bindings for source chain, ignore if chain is disabled", "sourceChain", sourceChain)
				} else if errors.Is(err, context.DeadlineExceeded) {
					lggr.Warnw("timed out getting latest msg seq num for source chain", "sourceChain", sourceChain)
				} else {
					lggr.Errorf("failed to get latest msg seq num for source chain %d: %s", sourceChain, err)
				}
				return
			}

			mu.Lock()
			latestOnRampSeqNums = append(
				latestOnRampSeqNums,
				plugintypes.NewSeqNumChain(sourceChain, latestOnRampSeqNum),
			)
			mu.Unlock()
		})
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
			wg.Go(func() {
				chainCtx, chainCancel := rpctimeout.Context(ctx)
				defer chainCancel()
				msgs, err := o.ccipReader.MsgsBetweenSeqNums(chainCtx, chainRange.ChainSel, chainRange.SeqNumRange)
				if err != nil {
					if errors.Is(err, context.DeadlineExceeded) {
						lggr.Warnw("timed out getting messages for source chain",
							"sourceChain", chainRange.ChainSel,
							"range", chainRange.SeqNumRange,
						)
					} else {
						lggr.Warnw("call to MsgsBetweenSeqNums failed", "err", err)
					}
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
			})
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

	supportsDestChain, err := o.chainSupport.SupportsDestChain(o.oracleID)
	if err != nil {
		lggr.Errorw("call to SupportsDestChain failed", "err", err)
		return cciptypes.RemoteConfig{}
	}

	if !supportsDestChain {
		lggr.Debugw("cannot observe RMN remote config since destination chain is not supported")
		return cciptypes.RemoteConfig{}
	}

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

// isNoBindingsError checks if the error is a no bindings error. We check both for the sentinel error
// and for the error message containing "no bindings" to cover different implementations since not all
// chain accessors use contract reader anymore.
// TODO: consider adding chain-agnostic ChainAccessor error types to cl-common.
func isNoBindingsError(err error) bool {
	if err == nil {
		return false
	}
	return errors.Is(err, contractreader.ErrNoBindings) || strings.Contains(err.Error(), "no bindings")
}
