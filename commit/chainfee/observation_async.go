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
	ccipreader "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type observer interface {
	getChainsFeeComponents() map[cciptypes.ChainSelector]types.ChainFeeComponents
	getNativeTokenPrices() map[cciptypes.ChainSelector]cciptypes.BigInt
	getChainFeePriceUpdates() map[cciptypes.ChainSelector]Update
	close()
}

type asyncObserver struct {
	cs         plugincommon.ChainSupport
	oracleID   commontypes.OracleID
	destChain  cciptypes.ChainSelector
	ccipReader ccipreader.CCIPReader
	lggr       logger.Logger
	cancelFunc func()
	mu         *sync.RWMutex

	chainsFeeComponents  map[cciptypes.ChainSelector]types.ChainFeeComponents
	nativeTokenPrices    map[cciptypes.ChainSelector]cciptypes.BigInt
	chainFeePriceUpdates map[cciptypes.ChainSelector]Update
}

// newAsyncObserver creates a new asyncObserver.
// It fetches the data from the base observer asynchronously and caches the results.
// It fetches the data every tickDur and uses a timeout of syncTimeout to kill long RPC calls.
func newAsyncObserver(
	lggr logger.Logger,
	ccipReader ccipreader.CCIPReader,
	destChain cciptypes.ChainSelector,
	oracleID commontypes.OracleID,
	cs plugincommon.ChainSupport,
	tickDur, syncTimeout time.Duration,
) *asyncObserver {
	ctx, cf := context.WithCancel(context.Background())

	obs := &asyncObserver{
		cs:         cs,
		oracleID:   oracleID,
		destChain:  destChain,
		ccipReader: ccipReader,
		lggr:       lggr,
		cancelFunc: nil,
		mu:         &sync.RWMutex{},
	}

	ticker := time.NewTicker(tickDur)
	lggr.Debugw("async observer started", "tickDur", tickDur, "syncTimeout", syncTimeout)
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
	o.lggr.Debugw("async observer is syncing", "syncTimeout", syncTimeout)
	ctxSync, cf := context.WithTimeout(ctx, syncTimeout)
	defer cf()

	supportedChains, err := o.getSupportedChains()
	if err != nil {
		o.lggr.Errorw("failed to get supported chains unable to sync", "err", err)
		return
	}

	syncOps := []struct {
		id string
		op func(context.Context)
	}{
		{
			id: "chainsFeeComponents",
			op: func(ctx context.Context) {
				chainsFeeComponents := o.ccipReader.GetChainsFeeComponents(ctx, supportedChains)
				o.mu.Lock()
				o.chainsFeeComponents = chainsFeeComponents
				o.mu.Unlock()
			},
		},
		{
			id: "nativeTokenPrices",
			op: func(ctx context.Context) {
				nativeTokenPrices := o.ccipReader.GetWrappedNativeTokenPriceUSD(ctx, supportedChains)
				o.mu.Lock()
				o.nativeTokenPrices = nativeTokenPrices
				o.mu.Unlock()
			},
		},
		{
			id: "chainFeePriceUpdates",
			op: func(ctx context.Context) {
				chainFeePriceUpdates := feeUpdatesFromTimestampedBig(
					o.ccipReader.GetChainFeePriceUpdate(ctx, supportedChains))
				o.mu.Lock()
				o.chainFeePriceUpdates = chainFeePriceUpdates
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

func (o *asyncObserver) getChainsFeeComponents() map[cciptypes.ChainSelector]types.ChainFeeComponents {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.chainsFeeComponents
}

func (o *asyncObserver) getNativeTokenPrices() map[cciptypes.ChainSelector]cciptypes.BigInt {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.nativeTokenPrices
}

func (o *asyncObserver) getChainFeePriceUpdates() map[cciptypes.ChainSelector]Update {
	o.mu.RLock()
	defer o.mu.RUnlock()
	return o.chainFeePriceUpdates
}

func (o *asyncObserver) close() {
	if o.cancelFunc != nil {
		o.cancelFunc()
		o.cancelFunc = nil
	}
}

func (o *asyncObserver) getSupportedChains() ([]cciptypes.ChainSelector, error) {
	supportedChains, err := o.cs.SupportedChains(o.oracleID)
	if err != nil {
		return nil, err
	}

	supportedChains.Remove(o.destChain)
	if supportedChains.Cardinality() == 0 {
		o.lggr.Info("no supported chains other than dest chain to observe")
		return nil, nil
	}

	supportedChainsSlice := supportedChains.ToSlice()
	sort.Slice(supportedChainsSlice, func(i, j int) bool { return supportedChainsSlice[i] < supportedChainsSlice[j] })

	return supportedChainsSlice, nil
}
