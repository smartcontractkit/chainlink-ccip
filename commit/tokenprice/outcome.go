package tokenprice

import (
	"fmt"
	"sort"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/shared"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
)

// getConsensusObservation Combine the list of observations into a single consensus observation
func (p *processor) getConsensusObservation(
	aos []shared.AttributedObservation[Observation],
) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)

	if len(aggObs.FDestChain) < p.bigF {
		return ConsensusObservation{},
			fmt.Errorf("not enough observations for fDestChain")
	}

	fDestChain := shared.Median(aggObs.FDestChain, func(a, b int) bool {
		return a < b
	})

	//twoFPlus1 := fDestChain*2 + 1

	consensusObs := ConsensusObservation{
		FDestChain: fDestChain,
		Timestamp:  shared.MedianTimestamp(aggObs.Timestamps),
	}

	return consensusObs, nil
}

//// feedPricesConsensus returns the median of the feed token prices for each token given all observed prices
//func feedPricesConsensus(
//	lggr logger.Logger,
//	feedTokenPrices map[types.Account][]cciptypes.BigInt,
//	fDestChain int,
//) map[types.Account]cciptypes.BigInt {
//	tokenPrices := make(map[types.Account]cciptypes.BigInt)
//	for token, prices := range feedTokenPrices {
//		if len(prices) < 2*fDestChain+1 {
//			lggr.Warnf("could not reach consensus on feed token prices for token %s ", token)
//			continue
//		}
//		tokenPrices[token] = shared.MedianBigInt(prices)
//	}
//	return tokenPrices
//}
//
//// registryPricesConsensus returns the median of the price registry token prices for each
//// token given all observed updates
//func registryPricesConsensus(
//	lggr logger.Logger,
//	priceRegistryPrices map[types.Account][]cciptypes.BigInt,
//	fDestChain int,
//) map[types.Account]cciptypes.BigInt {
//	tokenPrices := make(map[types.Account]cciptypes.BigInt)
//	for token, prices := range priceRegistryPrices {
//		if len(prices) < 2*fDestChain+1 {
//			lggr.Warnf("could not reach consensus on fee quoter token updates for token %s ", token)
//			continue
//		}
//		medianPrice := shared.MedianBigInt(prices)
//		tokenPrices[token] = cciptypes.NewBigInt(medianPrice.Int)
//	}
//
//	return tokenPrices
//}

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
			p.lggr.Warnf("could not find fee quoter update for token %s", token)
			continue
		}

		ti, ok := tokenInfo[token]
		if !ok {
			p.lggr.Warnf("could not find token info for token %s", token)
			continue
		}

		nextUpdateTime := lastUpdate.Timestamp.Add(cfg.TokenPriceBatchWriteFrequency.Duration())
		if obs.Timestamp.After(nextUpdateTime) {
			tokenPrices = append(tokenPrices, cciptypes.TokenPrice{
				TokenID: token,
				Price:   cciptypes.NewBigInt(feedPrice.Price.Int),
			})
		} else if shared.Deviates(feedPrice.Price.Int, lastUpdate.Value.Int, ti.DeviationPPB.Int64()) {
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
func aggregateObservations(aos []shared.AttributedObservation[Observation]) AggregateObservation {
	aggObs := AggregateObservation{
		FeedTokenPrices:       make(map[types.Account][]cciptypes.TokenPrice),
		FeeQuoterTokenUpdates: make(map[types.Account][]shared.TimestampedBig),
		FDestChain:            []int{},
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

		// FDestChain
		aggObs.FDestChain = append(aggObs.FDestChain, obs.FDestChain)

		// Timestamps
		aggObs.Timestamps = append(aggObs.Timestamps, obs.Timestamp)
	}

	return aggObs
}
