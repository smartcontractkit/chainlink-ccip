package tokenprice

import (
	"context"
	"fmt"
	"sort"

	"github.com/smartcontractkit/chainlink-ccip/shared"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"golang.org/x/exp/maps"
)

func (p *processor) ObserveFDestChain() (int, error) {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		// TODO: metrics
		p.lggr.Warnw("call to GetFChain failed", "err", err)
		return 0, fmt.Errorf("failed to get FChain")
	}

	fDestChain, ok := fChain[p.cfg.DestChain]
	if !ok {
		return 0, fmt.Errorf("fChain does not have dest chain")
	}

	return fDestChain, nil
}

func (p *processor) ObserveFeedTokenPrices(ctx context.Context) []cciptypes.TokenPrice {
	supportedChains, err := p.chainSupport.SupportedChains(p.oracleID)
	if err != nil {
		p.lggr.Warnw("call to SupportedChains failed", "err", err)
		return []cciptypes.TokenPrice{}
	}

	if !supportedChains.Contains(cciptypes.ChainSelector(p.cfg.OffchainConfig.TokenPriceChainSelector)) {
		p.lggr.Debugw("oracle does not support token price observation", "oracleID", p.oracleID)
		return []cciptypes.TokenPrice{}

	}

	if p.tokenPriceReader == nil {
		p.lggr.Errorw("no token price reader available")
		return []cciptypes.TokenPrice{}
	}

	tokensToQuery := maps.Keys(p.cfg.OffchainConfig.TokenInfo)
	//sort tokens to query to ensure deterministic order
	sort.Slice(tokensToQuery, func(i, j int) bool { return tokensToQuery[i] < tokensToQuery[j] })
	p.lggr.Infow("observing feed token prices")
	tokenPrices, err := p.tokenPriceReader.GetTokenFeedPricesUSD(ctx, tokensToQuery)
	if err != nil {
		p.lggr.Errorw("call to GetTokenFeedPricesUSD failed", "err", err)
		return []cciptypes.TokenPrice{}
	}

	// If we couldn't fetch all prices log and return only the ones we could fetch
	if len(tokenPrices) != len(tokensToQuery) {
		p.lggr.Errorw("token prices length mismatch", "got", len(tokenPrices), "want", len(tokensToQuery))
	}

	tokenPricesUSD := make([]cciptypes.TokenPrice, 0, len(tokenPrices))
	for i, token := range tokensToQuery {
		tokenPricesUSD = append(tokenPricesUSD, cciptypes.NewTokenPrice(token, tokenPrices[i]))
	}

	return tokenPricesUSD
}

func (p *processor) ObserveFeeQuoterTokenUpdates(ctx context.Context) map[types.Account]shared.TimestampedBig {
	supportsDestChain, err := p.chainSupport.SupportsDestChain(p.oracleID)
	if err != nil {
		p.lggr.Warnw("call to SupportsDestChain failed", "err", err)
		return map[types.Account]shared.TimestampedBig{}
	}
	if !supportsDestChain {
		p.lggr.Debugw("oracle does not support price registry observation", "oracleID", p.oracleID)
		return map[types.Account]shared.TimestampedBig{}
	}

	if p.tokenPriceReader == nil {
		p.lggr.Errorw("no token price reader available")
		return map[types.Account]shared.TimestampedBig{}
	}

	tokensToQuery := maps.Keys(p.cfg.OffchainConfig.TokenInfo)
	//sort tokens to query to ensure deterministic order
	sort.Slice(tokensToQuery, func(i, j int) bool { return tokensToQuery[i] < tokensToQuery[j] })
	p.lggr.Infow("observing price registry token updates")
	priceUpdates, err := p.tokenPriceReader.GetFeeQuoterTokenUpdates(ctx, tokensToQuery)
	if err != nil {
		p.lggr.Errorw("call to GetFeeQuoterTokenUpdates failed", "err", err)
		return map[types.Account]shared.TimestampedBig{}
	}

	tokenUpdates := make(map[types.Account]shared.TimestampedBig)

	for token, update := range priceUpdates {
		tokenUpdates[token] = shared.TimestampedBig{
			Value:     update.Value,
			Timestamp: update.Timestamp,
		}
	}

	return tokenUpdates
}
