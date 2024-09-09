package tokenprice

import (
	"context"
	"fmt"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/shared"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"golang.org/x/exp/maps"
)

type Processor struct {
	oracleID         commontypes.OracleID
	lggr             logger.Logger
	cfg              pluginconfig.CommitPluginConfig
	chainSupport     plugincommon.ChainSupport
	tokenPriceReader reader.PriceReader
	homeChain        reader.HomeChain
}

func NewProcessor(
	oracleID commontypes.OracleID,
	lggr logger.Logger,
	cfg pluginconfig.CommitPluginConfig,
	chainSupport plugincommon.ChainSupport,
	tokenPriceReader reader.PriceReader,
	homeChain reader.HomeChain,
) *Processor {
	return &Processor{
		oracleID:         oracleID,
		lggr:             lggr,
		cfg:              cfg,
		chainSupport:     chainSupport,
		tokenPriceReader: tokenPriceReader,
		homeChain:        homeChain,
	}
}

func (p *Processor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	return Query{}, nil
}

func (p *Processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	query Query,
) (Observation, error) {

	fDestChain, err := p.ObserveFDestChain()
	if err != nil {
		return Observation{}, err
	}

	return Observation{
		FeedTokenPrices:       p.ObserveFeedTokenPrices(ctx),
		FeeQuoterTokenUpdates: p.ObserveFeeQuoterTokenUpdates(ctx),
		FDestChain:            *fDestChain,
		Timestamp:             time.Now().UTC(),
	}, nil
}

func (p *Processor) Outcome(
	prevOutcome Outcome,
	query Query,
	aos []shared.AttributedObservation[Observation],
) (Outcome, error) {
	return Outcome{}, nil
}

func (p *Processor) ValidateObservation(
	prevOutcome Outcome,
	query Query,
	ao shared.AttributedObservation[Observation],
) error {
	//TODO: Validate token prices
	return nil
}

func (p *Processor) ObserveFDestChain() (*int, error) {
	fChain, err := p.homeChain.GetFChain()
	if err != nil {
		// TODO: metrics
		p.lggr.Warnw("call to GetFChain failed", "err", err)
		return nil, fmt.Errorf("failed to get FChain")
	}

	fDestChain, ok := fChain[p.cfg.DestChain]
	if !ok {
		return nil, fmt.Errorf("fChain does not have dest chain")
	}

	return &fDestChain, nil
}

func (p *Processor) ObserveFeedTokenPrices(ctx context.Context) []cciptypes.TokenPrice {
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

	// Continue if we couldn't fetch all prices
	if len(tokenPrices) != len(tokensToQuery) {
		p.lggr.Errorw("token prices length mismatch", "got", len(tokenPrices), "want", len(tokensToQuery))
	}

	tokenPricesUSD := make([]cciptypes.TokenPrice, 0, len(tokenPrices))
	for i, token := range tokensToQuery {
		tokenPricesUSD = append(tokenPricesUSD, cciptypes.NewTokenPrice(token, tokenPrices[i]))
	}

	return tokenPricesUSD
}

func (p *Processor) ObserveFeeQuoterTokenUpdates(ctx context.Context) map[types.Account]shared.TimestampedBig {
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

func validateObservedTokenPrices(tokenPrices []cciptypes.TokenPrice) error {
	tokensWithPrice := mapset.NewSet[types.Account]()
	for _, t := range tokenPrices {
		if tokensWithPrice.Contains(t.TokenID) {
			return fmt.Errorf("duplicate token price for token: %s", t.TokenID)
		}
		tokensWithPrice.Add(t.TokenID)

		if t.Price.IsEmpty() {
			return fmt.Errorf("token price must not be empty")
		}
	}
	return nil
}

var _ shared.PluginProcessor[Query, Observation, Outcome] = &Processor{}
