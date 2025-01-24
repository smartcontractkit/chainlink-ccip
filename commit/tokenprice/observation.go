package tokenprice

import (
	"context"
	"sort"
	"time"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func (p *processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	query Query,
) (Observation, error) {
	lggr := logutil.WithContextValues(ctx, p.lggr)

	fChain := p.ObserveFChain(lggr)
	if len(fChain) == 0 {
		return Observation{}, nil
	}

	feedTokenPrices := p.ObserveFeedTokenPrices(ctx, lggr)
	feeQuoterUpdates := p.ObserveFeeQuoterTokenUpdates(ctx, lggr)
	now := time.Now().UTC()
	lggr.Infow(
		"observed token prices",
		"feed prices", feedTokenPrices,
		"fee quoter updates", feeQuoterUpdates,
		"timestampNow", now,
	)

	obs := Observation{
		FeedTokenPrices:       feedTokenPrices,
		FeeQuoterTokenUpdates: feeQuoterUpdates,
		FChain:                fChain,
		Timestamp:             now,
	}
	p.metricsReporter.TrackTokenPricesObservation(obs)
	return obs, nil
}

func (p *processor) ObserveFChain(lggr logger.Logger) map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		lggr.Errorw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}

func (p *processor) ObserveFeedTokenPrices(ctx context.Context, lggr logger.Logger) cciptypes.TokenPriceMap {
	if p.tokenPriceReader == nil {
		lggr.Debugw("no token price reader available")
		return cciptypes.TokenPriceMap{}
	}

	supportedChains, err := p.chainSupport.SupportedChains(p.oracleID)
	if err != nil {
		lggr.Warnw("call to SupportedChains failed", "err", err)
		return cciptypes.TokenPriceMap{}
	}

	if !supportedChains.Contains(p.offChainCfg.PriceFeedChainSelector) {
		lggr.Debugf("oracle does not support feed chain %d", p.offChainCfg.PriceFeedChainSelector)
		return cciptypes.TokenPriceMap{}
	}

	tokensToQuery := maps.Keys(p.offChainCfg.TokenInfo)
	// sort tokens to query to ensure deterministic order
	sort.Slice(tokensToQuery, func(i, j int) bool { return tokensToQuery[i] < tokensToQuery[j] })
	lggr.Infow("observing feed token prices", "tokens", tokensToQuery)
	tokenPrices, err := p.tokenPriceReader.GetFeedPricesUSD(ctx, tokensToQuery)
	if err != nil {
		lggr.Errorw("call to GetFeedPricesUSD failed",
			"err", err)
		return cciptypes.TokenPriceMap{}
	}

	return tokenPrices
}

func (p *processor) ObserveFeeQuoterTokenUpdates(
	ctx context.Context,
	lggr logger.Logger,
) map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig {
	if p.tokenPriceReader == nil {
		lggr.Debugw("no token price reader available")
		return map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{}
	}

	supportsDestChain, err := p.chainSupport.SupportsDestChain(p.oracleID)
	if err != nil {
		lggr.Warnw("call to SupportsDestChain failed", "err", err)
		return map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{}
	}
	if !supportsDestChain {
		lggr.Debugw("oracle does not support fee quoter observation")
		return map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{}
	}

	tokensToQuery := maps.Keys(p.offChainCfg.TokenInfo)
	// sort tokens to query to ensure deterministic order
	sort.Slice(tokensToQuery, func(i, j int) bool { return tokensToQuery[i] < tokensToQuery[j] })
	lggr.Infow("observing fee quoter token updates")
	priceUpdates, err := p.tokenPriceReader.GetFeeQuoterTokenUpdates(ctx, tokensToQuery, p.destChain)
	if err != nil {
		lggr.Errorw("call to GetFeeQuoterTokenUpdates failed", "err", err)
		return map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig{}
	}

	tokenUpdates := make(map[cciptypes.UnknownEncodedAddress]plugintypes.TimestampedBig)

	for token, update := range priceUpdates {
		tokenUpdates[token] = plugintypes.TimestampedBig{
			Value:     update.Value,
			Timestamp: update.Timestamp,
		}
	}

	return tokenUpdates
}
