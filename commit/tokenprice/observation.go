package tokenprice

import (
	"context"
	"sort"
	"sync"
	"time"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	pkgreader "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func (p *processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	query Query,
) (Observation, error) {
	lggr := logutil.WithContextValues(ctx, p.lggr)

	fChain := p.observeFChain(lggr)
	if len(fChain) == 0 {
		return Observation{}, nil
	}

	feedTokenPrices := p.obs.observeFeedTokenPrices(ctx, lggr)
	feeQuoterUpdates := p.obs.observeFeeQuoterTokenUpdates(ctx, lggr)
	now := time.Now().UTC()
	lggr.Infow(
		"observed token prices",
		"feedPrices", feedTokenPrices,
		"feeQuoterUpdates", feeQuoterUpdates,
		"timestampNow", now,
	)

	obs := Observation{
		FeedTokenPrices:       feedTokenPrices,
		FeeQuoterTokenUpdates: feeQuoterUpdates,
		FChain:                fChain,
		Timestamp:             now,
	}
	return obs, nil
}

func (p *processor) observeFChain(lggr logger.Logger) map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		lggr.Errorw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}

type observer interface {
	observeFeedTokenPrices(ctx context.Context, lggr logger.Logger) cciptypes.TokenPriceMap
	observeFeeQuoterTokenUpdates(
		ctx context.Context,
		lggr logger.Logger) map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig
	close()
}

type baseObserver struct {
	oracleID         commontypes.OracleID
	tokenPriceReader pkgreader.PriceReader
	chainSupport     plugincommon.ChainSupport
	offChainCfg      pluginconfig.CommitOffchainConfig
	destChain        cciptypes.ChainSelector
}

func newBaseObserver(
	tokenPriceReader pkgreader.PriceReader,
	destChain cciptypes.ChainSelector,
	oracleID commontypes.OracleID,
	chainSupport plugincommon.ChainSupport,
	offchainCfg pluginconfig.CommitOffchainConfig,
) *baseObserver {
	return &baseObserver{
		oracleID:         oracleID,
		tokenPriceReader: tokenPriceReader,
		chainSupport:     chainSupport,
		destChain:        destChain,
		offChainCfg:      offchainCfg,
	}
}

func (b *baseObserver) observeFeedTokenPrices(ctx context.Context, lggr logger.Logger) cciptypes.TokenPriceMap {
	if b.tokenPriceReader == nil {
		lggr.Debugw("no token price reader available")
		return cciptypes.TokenPriceMap{}
	}

	supportedChains, err := b.chainSupport.SupportedChains(b.oracleID)
	if err != nil {
		lggr.Warnw("call to SupportedChains failed", "err", err)
		return cciptypes.TokenPriceMap{}
	}

	if !supportedChains.Contains(b.offChainCfg.PriceFeedChainSelector) {
		lggr.Debugf("oracle does not support feed chain %d", b.offChainCfg.PriceFeedChainSelector)
		return cciptypes.TokenPriceMap{}
	}

	tokensToQuery := maps.Keys(b.offChainCfg.TokenInfo)
	lggr.Infow("observing feed token prices", "tokens", tokensToQuery)
	tokenPrices, err := b.tokenPriceReader.GetFeedPricesUSD(ctx, tokensToQuery)
	if err != nil {
		lggr.Errorw("call to GetFeedPricesUSD failed",
			"err", err)
		return cciptypes.TokenPriceMap{}
	}

	return tokenPrices
}

func (b *baseObserver) observeFeeQuoterTokenUpdates(
	ctx context.Context,
	lggr logger.Logger) map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig {
	if b.tokenPriceReader == nil {
		lggr.Debugw("no token price reader available")
		return map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{}
	}

	supportsDestChain, err := b.chainSupport.SupportsDestChain(b.oracleID)
	if err != nil {
		lggr.Warnw("call to SupportsDestChain failed", "err", err)
		return map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{}
	}
	if !supportsDestChain {
		lggr.Debugw("oracle does not support fee quoter observation")
		return map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{}
	}

	tokensToQuery := maps.Keys(b.offChainCfg.TokenInfo)
	// sort tokens to query to ensure deterministic order
	sort.Slice(tokensToQuery, func(i, j int) bool { return tokensToQuery[i] < tokensToQuery[j] })
	lggr.Infow("observing fee quoter token updates")
	priceUpdates, err := b.tokenPriceReader.GetFeeQuoterTokenUpdates(ctx, tokensToQuery, b.destChain)
	if err != nil {
		lggr.Errorw("call to GetFeeQuoterTokenUpdates failed", "err", err)
		return map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig{}
	}

	tokenUpdates := make(map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig)

	for token, update := range priceUpdates {
		tokenUpdates[token] = cciptypes.TimestampedBig{
			Value:     update.Value,
			Timestamp: update.Timestamp,
		}
	}

	return tokenUpdates
}

func (b *baseObserver) close() {}

// asyncObserver wraps baseObserver and periodically syncs the tokenPriceMap and tokenUpdates.
// It is used to avoid blocking the processor when querying the tokenPriceReader.
type asyncObserver struct {
	lggr       logger.Logger
	base       *baseObserver
	cancelFunc func()
	mu         sync.RWMutex

	// cached values, only ever read thru mutex.
	tokenPriceMap cciptypes.TokenPriceMap
	tokenUpdates  map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig
}

func newAsyncObserver(
	lggr logger.Logger,
	base *baseObserver,
	tickDuration, syncTimeout time.Duration,
) *asyncObserver {
	ctx, cancel := context.WithCancel(context.Background())

	obs := &asyncObserver{
		lggr: logutil.WithComponent(lggr, "tokenpriceAsyncObserver"),
		base: base,
		mu:   sync.RWMutex{},
	}

	ticker := time.NewTicker(tickDuration)
	lggr.Debugw("async tokenprice observer started", "tickDuration", tickDuration, "syncTimeout", syncTimeout)
	obs.start(ctx, ticker.C, syncTimeout)

	obs.cancelFunc = func() {
		cancel()
		ticker.Stop()
	}

	return obs
}

func (a *asyncObserver) start(ctx context.Context, tickerC <-chan time.Time, syncTimeout time.Duration) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-tickerC:
				a.sync(ctx, syncTimeout)
			}
		}
	}()
}

func (a *asyncObserver) sync(ctx context.Context, syncTimeout time.Duration) {
	a.lggr.Debugw("async tokenprice observer is syncing")
	ctxSync, cancel := context.WithTimeout(ctx, syncTimeout)
	defer cancel()

	syncOps := []struct {
		id string
		op func(context.Context)
	}{
		{
			id: "feedTokenPrices",
			op: func(ctx context.Context) {
				tokenPriceMap := a.base.observeFeedTokenPrices(ctx, a.lggr)
				a.mu.Lock()
				a.tokenPriceMap = tokenPriceMap
				a.mu.Unlock()
			},
		},
		{
			id: "feeQuoterTokenUpdates",
			op: func(ctx context.Context) {
				tokenUpdates := a.base.observeFeeQuoterTokenUpdates(ctx, a.lggr)
				a.mu.Lock()
				a.tokenUpdates = tokenUpdates
				a.mu.Unlock()
			},
		},
	}

	wg := &sync.WaitGroup{}
	wg.Add(len(syncOps))
	for _, op := range syncOps {
		go a.applySyncOp(ctxSync, op.id, wg, op.op)
	}
	wg.Wait()
}

// applySyncOp applies the given operation synchronously.
func (a *asyncObserver) applySyncOp(
	ctx context.Context, id string, wg *sync.WaitGroup, op func(ctx context.Context)) {
	defer wg.Done()
	tStart := time.Now()
	a.lggr.Debugw("async observer applying sync operation", "id", id)
	op(ctx)
	a.lggr.Debugw("async observer has applied the sync operation",
		"id", id, "duration", time.Since(tStart))
}

// observeFeeQuoterTokenUpdates implements observer by returning the cached tokenUpdates.
func (a *asyncObserver) observeFeeQuoterTokenUpdates(
	ctx context.Context,
	lggr logger.Logger,
) map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedBig {
	a.mu.RLock()
	defer a.mu.RUnlock()
	lggr.Debugw("observeFeeQuoterTokenUpdates returning cached value", "numUpdates", len(a.tokenUpdates))
	return a.tokenUpdates
}

// observeFeedTokenPrices implements observer by returning the cached tokenPriceMap.
func (a *asyncObserver) observeFeedTokenPrices(
	ctx context.Context,
	lggr logger.Logger,
) cciptypes.TokenPriceMap {
	a.mu.RLock()
	defer a.mu.RUnlock()
	lggr.Debugw("observeFeedTokenPrices returning cached value", "numPrices", len(a.tokenPriceMap))
	return a.tokenPriceMap
}

func (a *asyncObserver) close() {
	a.cancelFunc()
}

var _ observer = &asyncObserver{}
var _ observer = &baseObserver{}
