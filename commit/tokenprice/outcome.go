package tokenprice

import (
	"fmt"
	"sort"
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/mathslib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// getConsensusObservation Combine the list of observations into a single consensus observation
func (p *processor) getConsensusObservation(
	aos []plugincommon.AttributedObservation[Observation],
) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)

	fMin := mathslib.RepeatedF(func() int { return p.bigF }, maps.Keys(aggObs.FChain))

	// consensus on the fChain map uses the role DON F value
	// because all nodes can observe the home chain.
	fChains := plugincommon.GetConsensusMap(p.lggr, "fChain", aggObs.FChain, fMin)

	fDestChain, exists := fChains[p.cfg.DestChain]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for fDestChain, destChain: %d", p.cfg.DestChain)
	}

	fFeedChain, exists := fChains[p.cfg.OffchainConfig.PriceFeedChainSelector]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for f for FeedChain: %d", p.cfg.OffchainConfig.PriceFeedChainSelector)
	}

	feedPricesConsensus := plugincommon.GetConsensusMapAggregator(
		p.lggr,
		"FeedTokenPrices",
		aggObs.FeedTokenPrices,
		mathslib.RepeatedF(
			func() int { return mathslib.TwoFPlus1(fFeedChain) },
			maps.Keys(aggObs.FeedTokenPrices),
		),
		func(vals []cciptypes.TokenPrice) cciptypes.TokenPrice {
			return plugincommon.Median(vals, plugincommon.TokenPriceComparator)
		},
	)

	feeQuoterUpdatesConsensus := plugincommon.GetConsensusMapAggregator(
		p.lggr,
		"FeeQuoterUpdates",
		aggObs.FeeQuoterTokenUpdates,
		mathslib.RepeatedF(
			func() int { return mathslib.TwoFPlus1(fDestChain) },
			maps.Keys(aggObs.FeeQuoterTokenUpdates),
		),
		plugincommon.TimestampedBigAggregator,
	)

	consensusObs := ConsensusObservation{
		FChain:                fChains,
		Timestamp:             plugincommon.Median(aggObs.Timestamps, plugincommon.TimestampComparator),
		FeedTokenPrices:       feedPricesConsensus,
		FeeQuoterTokenUpdates: feeQuoterUpdatesConsensus,
	}

	return consensusObs, nil
}

// selectTokensForUpdate checks which tokens need to be updated based on the observed token prices and
// the fee quoter updates
// a token is selected for update if it meets one of 2 conditions:
// 1. if time passed since the last update is greater than the stale threshold
// 2. if deviation between the fee quoter and feed exceeds token's configured threshold
func (p *processor) selectTokensForUpdate(
	obs ConsensusObservation,
) []cciptypes.TokenPrice {
	var tokenPrices []cciptypes.TokenPrice
	cfg := p.cfg.OffchainConfig
	tokenInfo := cfg.TokenInfo

	for token, feedPrice := range obs.FeedTokenPrices {
		lastUpdate, exists := obs.FeeQuoterTokenUpdates[token]
		if !exists {
			// if the token is not in the fee quoter updates, then we should update it
			tokenPrices = append(tokenPrices, cciptypes.TokenPrice{
				TokenID: token,
				Price:   cciptypes.NewBigInt(feedPrice.Price.Int),
			})
			continue
		}

		ti, ok := tokenInfo[token]
		if !ok {
			p.lggr.Warnf("could not find token info for token %s", token)
			continue
		}

		nextUpdateTime := lastUpdate.Timestamp.Add(cfg.TokenPriceBatchWriteFrequency.Duration())
		shouldUpdate :=
			obs.Timestamp.After(nextUpdateTime) ||
				mathslib.Deviates(feedPrice.Price.Int, lastUpdate.Value.Int, ti.DeviationPPB.Int64())
		if shouldUpdate {
			tokenPrices = append(tokenPrices, cciptypes.TokenPrice{
				TokenID: token,
				Price:   cciptypes.NewBigInt(feedPrice.Price.Int),
			})
		}
	}

	// sort the token prices by tokenID
	sort.Slice(tokenPrices, func(i, j int) bool {
		return tokenPrices[i].TokenID < tokenPrices[j].TokenID
	})
	return tokenPrices
}

// aggregateObservations takes a list of observations and produces an AggregateObservation
func aggregateObservations(aos []plugincommon.AttributedObservation[Observation]) AggregateObservation {
	aggObs := AggregateObservation{
		FeedTokenPrices:       make(map[types.Account][]cciptypes.TokenPrice),
		FeeQuoterTokenUpdates: make(map[types.Account][]plugintypes.TimestampedBig),
		FChain:                make(map[cciptypes.ChainSelector][]int),
		Timestamps:            []time.Time{},
	}

	for _, ao := range aos {
		obs := ao.Observation
		// FeedTokenPrices
		for _, tokenPrice := range obs.FeedTokenPrices {
			aggObs.FeedTokenPrices[tokenPrice.TokenID] = append(aggObs.FeedTokenPrices[tokenPrice.TokenID], tokenPrice)
		}

		// FeeQuoterTokenUpdates
		for account, update := range obs.FeeQuoterTokenUpdates {
			aggObs.FeeQuoterTokenUpdates[account] = append(aggObs.FeeQuoterTokenUpdates[account], update)
		}

		// FChain
		for chainSel, f := range obs.FChain {
			aggObs.FChain[chainSel] = append(aggObs.FChain[chainSel], f)
		}

		// Timestamps
		aggObs.Timestamps = append(aggObs.Timestamps, obs.Timestamp)
	}

	return aggObs
}
