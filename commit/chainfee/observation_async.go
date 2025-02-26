package chainfee

import (
	"context"
	"sort"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	ccipreader "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type observer interface {
	getChainsFeeComponents(ctx context.Context, lggr logger.Logger) map[cciptypes.ChainSelector]types.ChainFeeComponents
	getNativeTokenPrices(ctx context.Context, lggr logger.Logger) map[cciptypes.ChainSelector]cciptypes.BigInt
	getChainFeePriceUpdates(ctx context.Context, lggr logger.Logger) map[cciptypes.ChainSelector]Update
	close()
}

type baseObserver struct {
	cs         plugincommon.ChainSupport
	oracleID   commontypes.OracleID
	destChain  cciptypes.ChainSelector
	ccipReader ccipreader.CCIPReader
}

func newBaseObserver(
	ccipReader ccipreader.CCIPReader,
	destChain cciptypes.ChainSelector,
	oracleID commontypes.OracleID,
	cs plugincommon.ChainSupport,
) *baseObserver {
	return &baseObserver{
		cs:         cs,
		oracleID:   oracleID,
		destChain:  destChain,
		ccipReader: ccipReader,
	}
}

func (o *baseObserver) getChainsFeeComponents(
	ctx context.Context,
	lggr logger.Logger,
) map[cciptypes.ChainSelector]types.ChainFeeComponents {
	supportedChains, err := o.getSupportedChains(lggr, o.cs, o.oracleID, o.destChain)
	if err != nil {
		lggr.Errorw("failed to get supported chains unable to get chains fee components", "err", err)
		return map[cciptypes.ChainSelector]types.ChainFeeComponents{}
	}
	return o.ccipReader.GetChainsFeeComponents(ctx, supportedChains)
}

func (o *baseObserver) getNativeTokenPrices(
	ctx context.Context,
	lggr logger.Logger,
) map[cciptypes.ChainSelector]cciptypes.BigInt {
	supportedChains, err := o.getSupportedChains(lggr, o.cs, o.oracleID, o.destChain)
	if err != nil {
		lggr.Errorw("failed to get supported chains unable to get native token prices", "err", err)
		return map[cciptypes.ChainSelector]cciptypes.BigInt{}
	}
	return o.ccipReader.GetWrappedNativeTokenPriceUSD(ctx, supportedChains)
}

func (o *baseObserver) getChainFeePriceUpdates(
	ctx context.Context,
	lggr logger.Logger,
) map[cciptypes.ChainSelector]Update {
	supportedChains, err := o.getSupportedChains(lggr, o.cs, o.oracleID, o.destChain)
	if err != nil {
		lggr.Errorw("failed to get supported chains unable to get chain fee price updates", "err", err)
		return map[cciptypes.ChainSelector]Update{}
	}
	return feeUpdatesFromTimestampedBig(o.ccipReader.GetChainFeePriceUpdate(ctx, supportedChains))
}

func (o *baseObserver) close() {
}

func (o *baseObserver) getSupportedChains(
	lggr logger.Logger,
	cs plugincommon.ChainSupport,
	oracleID commontypes.OracleID,
	destChain cciptypes.ChainSelector,
) ([]cciptypes.ChainSelector, error) {
	supportedChains, err := cs.SupportedChains(oracleID)
	if err != nil {
		return nil, err
	}

	supportedChains.Remove(destChain)
	if supportedChains.Cardinality() == 0 {
		lggr.Info("no supported chains other than dest chain to observe")
		return nil, nil
	}

	supportedChainsSlice := supportedChains.ToSlice()
	sort.Slice(supportedChainsSlice, func(i, j int) bool { return supportedChainsSlice[i] < supportedChainsSlice[j] })

	return supportedChainsSlice, nil
}

type asyncObserver struct {
	baseObserver         *baseObserver
	lggr                 logger.Logger
	cancelFunc           func()
	mu                   *sync.RWMutex
	chainsFeeComponents  map[cciptypes.ChainSelector]types.ChainFeeComponents
	nativeTokenPrices    map[cciptypes.ChainSelector]cciptypes.BigInt
	chainFeePriceUpdates map[cciptypes.ChainSelector]Update
}

func newAsyncObserver(
	lggr logger.Logger,
	baseObserver *baseObserver,
	tickDur, syncTimeout time.Duration,
) *asyncObserver {
	ctx, cf := context.WithCancel(context.Background())

	obs := &asyncObserver{
		baseObserver: baseObserver,
		lggr:         logutil.WithComponent(lggr, "chainfeeAsyncObserver"),
		cancelFunc:   nil,
		mu:           &sync.RWMutex{},
	}

	ticker := time.NewTicker(tickDur)
	lggr.Debugw("async chainfee observer started", "tickDur", tickDur, "syncTimeout", syncTimeout)
	obs.start(ctx, ticker.C, syncTimeout)

	obs.cancelFunc = func() {
		cf()
		ticker.Stop()
	}

	return obs
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
	o.lggr.Debugw("async chainfee observer is syncing", "syncTimeout", syncTimeout)
	ctxSync, cf := context.WithTimeout(ctx, syncTimeout)
	defer cf()

	syncOps := []struct {
		id string
		op func(context.Context)
	}{
		{
			id: "chainsFeeComponents",
			op: func(ctx context.Context) {
				chainsFeeComponents := o.baseObserver.getChainsFeeComponents(ctx, o.lggr)
				o.mu.Lock()
				o.chainsFeeComponents = chainsFeeComponents
				o.mu.Unlock()
			},
		},
		{
			id: "nativeTokenPrices",
			op: func(ctx context.Context) {
				nativeTokenPrices := o.baseObserver.getNativeTokenPrices(ctx, o.lggr)
				o.mu.Lock()
				o.nativeTokenPrices = nativeTokenPrices
				o.mu.Unlock()
			},
		},
		{
			id: "chainFeePriceUpdates",
			op: func(ctx context.Context) {
				chainFeePriceUpdates := o.baseObserver.getChainFeePriceUpdates(ctx, o.lggr)
				o.mu.Lock()
				o.chainFeePriceUpdates = chainFeePriceUpdates
				o.mu.Unlock()
			},
		},
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(syncOps))
	for _, op := range syncOps {
		go o.applySyncOp(ctxSync, op.id, wg, op.op)
	}
	wg.Wait()
}

// applySyncOp applies the given operation synchronously.
func (o *asyncObserver) applySyncOp(
	ctx context.Context, id string, wg *sync.WaitGroup, op func(ctx context.Context)) {
	defer wg.Done()
	tStart := time.Now()
	o.lggr.Debugw("async observer applying sync operation", "id", id)
	op(ctx)
	o.lggr.Debugw("async observer has applied the sync operation",
		"id", id, "duration", time.Since(tStart))
}

func (o *asyncObserver) getChainsFeeComponents(
	_ context.Context,
	lggr logger.Logger,
) map[cciptypes.ChainSelector]types.ChainFeeComponents {
	o.mu.RLock()
	defer o.mu.RUnlock()
	lggr.Debugw("getChainsFeeComponents returning cached value", "numFeeComponents", len(o.chainsFeeComponents))
	return o.chainsFeeComponents
}

func (o *asyncObserver) getNativeTokenPrices(
	_ context.Context,
	lggr logger.Logger,
) map[cciptypes.ChainSelector]cciptypes.BigInt {
	o.mu.RLock()
	defer o.mu.RUnlock()
	lggr.Debugw("getNativeTokenPrices returning cached value", "numPrices", len(o.nativeTokenPrices))
	return o.nativeTokenPrices
}

func (o *asyncObserver) getChainFeePriceUpdates(
	_ context.Context,
	lggr logger.Logger,
) map[cciptypes.ChainSelector]Update {
	o.mu.RLock()
	defer o.mu.RUnlock()
	lggr.Debugw("getChainFeePriceUpdates returning cached value", "numUpdates", len(o.chainFeePriceUpdates))
	return o.chainFeePriceUpdates
}

func (o *asyncObserver) close() {
	if o.cancelFunc != nil {
		o.cancelFunc()
		o.cancelFunc = nil
	}
}
