package tokenprice

import (
	"context"
	"sort"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"golang.org/x/exp/maps"
)

func (p *processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	query Query,
) (Observation, error) {

	fChain := p.ObserveFChain()
	if len(fChain) == 0 {
		return Observation{}, nil
	}

	feedTokenPrices := p.ObserveFeedTokenPrices(ctx)
	feeQuoterUpdates := p.ObserveFeeQuoterTokenUpdates(ctx)
	ts := time.Now().UTC()
	p.lggr.Infow(
		"observed token prices",
		"feed prices", feedTokenPrices,
		"fee quoter updates", feeQuoterUpdates,
		"timestamp", ts,
	)

	return Observation{
		FeedTokenPrices:       feedTokenPrices,
		FeeQuoterTokenUpdates: feeQuoterUpdates,
		FChain:                fChain,
		Timestamp:             ts,
	}, nil
}

func (p *processor) ObserveFChain() map[cciptypes.ChainSelector]int {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		p.lggr.Errorw("call to GetFChain failed", "err", err)
		return map[cciptypes.ChainSelector]int{}
	}
	return fChain
}

func (p *processor) ObserveFeedTokenPrices(ctx context.Context) []cciptypes.TokenPrice {
	if p.tokenPriceReader == nil {
		p.lggr.Debugw("no token price reader available")
		return []cciptypes.TokenPrice{}
	}

	supportedChains, err := p.chainSupport.SupportedChains(p.oracleID)
	if err != nil {
		p.lggr.Warnw("call to SupportedChains failed", "err", err)
		return []cciptypes.TokenPrice{}
	}

	if !supportedChains.Contains(p.offChainCfg.PriceFeedChainSelector) {
		p.lggr.Debugf("oracle does not support feed chain %d", p.offChainCfg.PriceFeedChainSelector)
		return []cciptypes.TokenPrice{}

	}

	tokensToQuery := maps.Keys(p.offChainCfg.TokenInfo)
	// sort tokens to query to ensure deterministic order
	sort.Slice(tokensToQuery, func(i, j int) bool { return tokensToQuery[i] < tokensToQuery[j] })
	p.lggr.Infow("observing feed token prices", "tokens", tokensToQuery)
	tokenPrices, err := p.tokenPriceReader.GetFeedPricesUSD(ctx, tokensToQuery)
	if err != nil {
		p.lggr.Errorw("call to GetFeedPricesUSD failed",
			"err", err)
		return []cciptypes.TokenPrice{}
	}

	// If we couldn't fetch all prices log and return only the ones we could fetch
	if len(tokenPrices) != len(tokensToQuery) {
		p.lggr.Errorw("token prices length mismatch", "got", tokenPrices, "want", tokensToQuery)
		return []cciptypes.TokenPrice{}
	}

	tokenPricesUSD := make([]cciptypes.TokenPrice, 0, len(tokenPrices))
	for i, token := range tokensToQuery {
		tokenPricesUSD = append(tokenPricesUSD, cciptypes.NewTokenPrice(token, tokenPrices[i]))
	}

	return tokenPricesUSD
}

func (p *processor) ObserveFeeQuoterTokenUpdates(ctx context.Context) map[types.Account]plugintypes.TimestampedBig {
	if p.tokenPriceReader == nil {
		p.lggr.Debugw("no token price reader available")
		return map[types.Account]plugintypes.TimestampedBig{}
	}

	supportsDestChain, err := p.chainSupport.SupportsDestChain(p.oracleID)
	if err != nil {
		p.lggr.Warnw("call to SupportsDestChain failed", "err", err)
		return map[types.Account]plugintypes.TimestampedBig{}
	}
	if !supportsDestChain {
		p.lggr.Debugw("oracle does not support fee quoter observation")
		return map[types.Account]plugintypes.TimestampedBig{}
	}

	tokensToQuery := maps.Keys(p.offChainCfg.TokenInfo)
	// sort tokens to query to ensure deterministic order
	sort.Slice(tokensToQuery, func(i, j int) bool { return tokensToQuery[i] < tokensToQuery[j] })
	p.lggr.Infow("observing fee quoter token updates")
	priceUpdates, err := p.tokenPriceReader.GetFeeQuoterTokenUpdates(ctx, tokensToQuery, p.destChain)
	if err != nil {
		p.lggr.Errorw("call to GetFeeQuoterTokenUpdates failed", "err", err)
		return map[types.Account]plugintypes.TimestampedBig{}
	}

	tokenUpdates := make(map[types.Account]plugintypes.TimestampedBig)

	for token, update := range priceUpdates {
		tokenUpdates[token] = plugintypes.TimestampedBig{
			Value:     update.Value,
			Timestamp: update.Timestamp,
		}
	}

	return tokenUpdates
}
