package tokenprice

import (
	"fmt"
	"sort"
	"time"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/mathslib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

// getConsensusObservation Combine the list of observations into a single consensus observation
func (p *processor) getConsensusObservation(
	aos []plugincommon.AttributedObservation[Observation],
) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)

	// consensus on the fChain map uses the role DON f value
	// because all nodes can observe the home chain.
	donThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(p.fRoleDON))
	fChains := consensus.GetConsensusMap(p.lggr, "fChain", aggObs.FChain, donThresh)

	fDestChain, exists := fChains[p.destChain]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for fDestChain, destChain: %d", p.destChain)
	}

	fFeedChain, exists := fChains[p.offChainCfg.PriceFeedChainSelector]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for f for FeedChain: %d", p.offChainCfg.PriceFeedChainSelector)
	}

	feedPricesConsensus := consensus.GetConsensusMapAggregator(
		p.lggr,
		"FeedTokenPrices",
		aggObs.FeedTokenPrices,
		consensus.MakeConstantThreshold[types.Account](consensus.TwoFPlus1(fFeedChain)),
		func(vals []cciptypes.TokenPrice) cciptypes.TokenPrice {
			return consensus.Median(vals, consensus.TokenPriceComparator)
		},
	)

	feeQuoterUpdatesConsensus := consensus.GetConsensusMapAggregator(
		p.lggr,
		"FeeQuoterUpdates",
		aggObs.FeeQuoterTokenUpdates,
		consensus.MakeConstantThreshold[types.Account](consensus.TwoFPlus1(fDestChain)),
		// each key will have one object with the median for timestamps as timestamp value
		// and the median prices as price value
		consensus.TimestampedBigAggregator,
	)

	consensusObs := ConsensusObservation{
		FChain:                fChains,
		Timestamp:             consensus.Median(aggObs.Timestamps, consensus.TimestampComparator),
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
	cfg := p.offChainCfg
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
