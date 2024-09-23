package chainfee

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/mathslib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"golang.org/x/exp/maps"
)

func (p *processor) getConsensusObservation(
	aos []plugincommon.AttributedObservation[Observation],
) (ConsensusObservation, error) {
	aggObs := aggregateObservations(aos)

	fMin := mathslib.RepeatedF(func() int { return p.bigF }, maps.Keys(aggObs.FChain))
	// consensus on the fChain map uses the role DON F value
	// because all nodes can observe the home chain.
	fChains := plugincommon.GetConsensusMap(p.lggr, "fChain", aggObs.FChain, fMin)

	fDestChain, exists := fChains[p.destChain]
	if !exists {
		return ConsensusObservation{},
			fmt.Errorf("no consensus value for fDestChain, destChain: %d", p.destChain)
	}

	timestamp := plugincommon.Median(aggObs.Timestamps, plugincommon.TimestampComparator)
	chainFeeUpdatesConsensus := plugincommon.GetConsensusMapAggregator(
		p.lggr,
		"ChainFeePriceUpdates",
		aggObs.ChainFeePriceUpdates,
		mathslib.RepeatedF(
			func() int { return mathslib.TwoFPlus1(fDestChain) },
			maps.Keys(aggObs.ChainFeePriceUpdates),
		),
		plugincommon.TimestampedBigAggregator,
	)

	// Stop early if earliest updated timestamp is still fresh
	earliestUpdateTime := plugincommon.EarliestTimestamp(maps.Values(chainFeeUpdatesConsensus), timestamp)
	nextUpdateTime := earliestUpdateTime.Add(p.ChainFeePriceBatchWriteFrequency.Duration())
	if earliestUpdateTime.Before(nextUpdateTime) {
		return ConsensusObservation{
			ShouldUpdate: false,
		}, nil
	}

	feeComponents := plugincommon.GetConsensusMapAggregator(
		p.lggr,
		"FeeComponents",
		aggObs.FeeComponents,
		mathslib.TwoFPlus1Map(fChains),
		func(vals []types.ChainFeeComponents) types.ChainFeeComponents {
			executionFees := make([]cciptypes.BigInt, len(vals))
			dataAvailabilityFees := make([]cciptypes.BigInt, len(vals))
			for i, feeComp := range vals {
				executionFees[i] = cciptypes.NewBigInt(feeComp.ExecutionFee)
				dataAvailabilityFees[i] = cciptypes.NewBigInt(feeComp.DataAvailabilityFee)
			}
			return types.ChainFeeComponents{
				ExecutionFee:        plugincommon.Median(executionFees, plugincommon.BigIntComparator).Int,
				DataAvailabilityFee: plugincommon.Median(dataAvailabilityFees, plugincommon.BigIntComparator).Int,
			}
		},
	)

	nativeTokenPrices := plugincommon.GetConsensusMapAggregator(
		p.lggr,
		"NativeTokenPrices",
		aggObs.NativeTokenPrices,
		mathslib.TwoFPlus1Map(fChains),
		func(vals []cciptypes.BigInt) cciptypes.BigInt {
			return plugincommon.Median(vals, plugincommon.BigIntComparator)
		},
	)

	consensusObs := ConsensusObservation{
		FChain:               fChains,
		FeeComponents:        feeComponents,
		NativeTokenPrices:    nativeTokenPrices,
		ChainFeePriceUpdates: chainFeeUpdatesConsensus,
		Timestamp:            timestamp,
		ShouldUpdate:         true,
	}

	return consensusObs, nil
}

func aggregateObservations(aos []plugincommon.AttributedObservation[Observation]) AggregateObservation {
	aggObs := AggregateObservation{
		FeeComponents:        make(map[cciptypes.ChainSelector][]types.ChainFeeComponents),
		NativeTokenPrices:    make(map[cciptypes.ChainSelector][]cciptypes.BigInt),
		FChain:               make(map[cciptypes.ChainSelector][]int),
		ChainFeePriceUpdates: make(map[cciptypes.ChainSelector][]plugintypes.TimestampedBig),
		Timestamps:           []time.Time{},
	}

	for _, ao := range aos {
		obs := ao.Observation

		// FeeComponents
		for chainSel, feeComp := range obs.FeeComponents {
			aggObs.FeeComponents[chainSel] = append(aggObs.FeeComponents[chainSel], feeComp)
		}

		// NativeTokenPrices
		for chainSel, tokenPrice := range obs.NativeTokenPrice {
			aggObs.NativeTokenPrices[chainSel] = append(aggObs.NativeTokenPrices[chainSel], tokenPrice)
		}

		// FChain
		for chainSel, f := range obs.FChain {
			aggObs.FChain[chainSel] = append(aggObs.FChain[chainSel], f)
		}

		for chainSel, feeUpdate := range obs.ChainFeePriceUpdates {
			aggObs.ChainFeePriceUpdates[chainSel] = append(aggObs.ChainFeePriceUpdates[chainSel], feeUpdate)
		}

		// Timestamps
		aggObs.Timestamps = append(aggObs.Timestamps, obs.Timestamp)
	}

	return aggObs
}
