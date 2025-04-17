package tokenprice

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/mathslib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// getConsensusObservation Combine the list of observations into a single consensus observation
func (p *processor) getConsensusObservation(
	lggr logger.Logger,
	aos []plugincommon.AttributedObservation[Observation],
) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)

	// consensus on the fChain map uses the role DON F value
	// because all nodes can observe the home chain.
	donThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(p.fRoleDON))
	fChains := consensus.GetConsensusMap(lggr, "fChain", aggObs.FChain, donThresh)

	fDestChain, exists := fChains[p.destChain]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for fDestChain, destChain: %d", p.destChain)
	}

	if consensus.LtTwoFPlusOne(fDestChain, len(aggObs.Timestamps)) {
		return ConsensusObservation{},
			fmt.Errorf("not enough observations for timestamps to reach consensus, have %d, need %d",
				len(aggObs.Timestamps), consensus.TwoFPlus1(fDestChain))
	}
	timestamp := consensus.Median(aggObs.Timestamps, consensus.TimestampComparator)

	fFeedChain, exists := fChains[p.offChainCfg.PriceFeedChainSelector]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for f for FeedChain: %d", p.offChainCfg.PriceFeedChainSelector)
	}

	feedPricesConsensus := consensus.GetConsensusMapAggregator(
		lggr,
		"FeedTokenPrices",
		aggObs.FeedTokenPrices,
		consensus.MakeConstantThreshold[cciptypes.UnknownEncodedAddress](consensus.TwoFPlus1(fFeedChain)),
		func(vals []cciptypes.TokenPrice) cciptypes.TokenPrice {
			return consensus.Median(vals, consensus.TokenPriceComparator)
		},
	)

	feeQuoterUpdatesConsensus := consensus.GetConsensusMapAggregator(
		lggr,
		"FeeQuoterUpdates",
		aggObs.FeeQuoterTokenUpdates,
		consensus.MakeConstantThreshold[cciptypes.UnknownEncodedAddress](consensus.TwoFPlus1(fDestChain)),
		// each key will have one object with the median for timestamps as timestamp value
		// and the median prices as price value
		consensus.TimestampedBigAggregator,
	)

	consensusObs := ConsensusObservation{
		FChain:                fChains,
		Timestamp:             timestamp,
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
	lggr logger.Logger,
	obs ConsensusObservation,
) cciptypes.TokenPriceMap {
	tokenPrices := make(cciptypes.TokenPriceMap)
	cfg := p.offChainCfg
	tokenInfo := cfg.TokenInfo

	for token, feedPrice := range obs.FeedTokenPrices {
		lastUpdate, exists := obs.FeeQuoterTokenUpdates[token]
		lggr := logger.With(lggr,
			"token", token,
			"feedPrice", feedPrice,
			"lastUpdate", lastUpdate,
			"consensusTimestamp", obs.Timestamp,
		)
		if !exists {
			lggr.Infow("token price update needed: no previous update exists")
			tokenPrices[token] = cciptypes.NewBigInt(feedPrice.Price.Int)
			continue
		}

		ti, ok := tokenInfo[token]
		if !ok {
			lggr.Warnf("could not find token info for token %s", token)
			continue
		}

		nextUpdateTime := lastUpdate.Timestamp.Add(cfg.TokenPriceBatchWriteFrequency.Duration())
		priceDeviates := mathslib.Deviates(feedPrice.Price.Int, lastUpdate.Value.Int, ti.DeviationPPB.Int64())
		heartbeatPassed := obs.Timestamp.After(nextUpdateTime)

		if heartbeatPassed {
			lggr.Infow("token price update needed: heartbeat time passed",
				"nextUpdateTime", nextUpdateTime,
				"heartbeatInterval", cfg.TokenPriceBatchWriteFrequency)
			tokenPrices[token] = cciptypes.NewBigInt(feedPrice.Price.Int)
		} else if priceDeviates {
			lggr.Infow("token price update needed: deviation threshold exceeded",
				"deviationPPB", ti.DeviationPPB)
			tokenPrices[token] = cciptypes.NewBigInt(feedPrice.Price.Int)
		} else {
			lggr.Debugw("token price update not needed: within deviation threshold",
				"deviationPPB", ti.DeviationPPB)
		}
	}

	return tokenPrices
}

// aggregateObservations takes a list of observations and produces an AggregateObservation
func aggregateObservations(aos []plugincommon.AttributedObservation[Observation]) AggregateObservation {
	aggObs := AggregateObservation{
		FeedTokenPrices:       make(map[cciptypes.UnknownEncodedAddress][]cciptypes.TokenPrice),
		FeeQuoterTokenUpdates: make(map[cciptypes.UnknownEncodedAddress][]cciptypes.TimestampedBig),
		FChain:                make(map[cciptypes.ChainSelector][]int),
		Timestamps:            []time.Time{},
	}

	for _, ao := range aos {
		obs := ao.Observation
		// FeedTokenPrices
		for tokenID, price := range obs.FeedTokenPrices {
			aggObs.FeedTokenPrices[tokenID] = append(
				aggObs.FeedTokenPrices[tokenID],
				cciptypes.NewTokenPrice(tokenID, price.Int),
			)
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
