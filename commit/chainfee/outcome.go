package chainfee

import (
	"fmt"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
)

func (p *processor) getConsensusObservation(
	aos []plugincommon.AttributedObservation[Observation],
) (Observation, error) {
	aggObs := aggregateObservations(aos)

	// consensus on the fChain map uses the role DON F value
	// because all nodes can observe the home chain.
	donThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(p.fRoleDON))
	fChains := consensus.GetConsensusMap(p.lggr, "fChain", aggObs.FChain, donThresh)

	fDestChain, exists := fChains[p.destChain]
	if !exists {
		return Observation{},
			fmt.Errorf("no consensus value for fDestChain, destChain: %d", p.destChain)
	}

	timestamp := consensus.Median(aggObs.Timestamps, consensus.TimestampComparator)
	chainFeeUpdatesConsensus := consensus.GetConsensusMapAggregator(
		p.lggr,
		"ChainFeeLatestUpdates",
		aggObs.ChainFeeLatestUpdates,
		consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(fDestChain)),
		consensus.TimestampMedianAggregator,
	)

	twoFPlus1 := consensus.MakeMultiThreshold(fChains, consensus.TwoFPlus1)

	feeComponents := consensus.GetConsensusMapAggregator(
		p.lggr,
		"FeeComponents",
		aggObs.FeeComponents,
		twoFPlus1,
		// Aggregator function
		func(vals []types.ChainFeeComponents) types.ChainFeeComponents {
			executionFees := make([]cciptypes.BigInt, len(vals))
			dataAvailabilityFees := make([]cciptypes.BigInt, len(vals))
			for i, feeComp := range vals {
				executionFees[i] = cciptypes.NewBigInt(feeComp.ExecutionFee)
				dataAvailabilityFees[i] = cciptypes.NewBigInt(feeComp.DataAvailabilityFee)
			}
			return types.ChainFeeComponents{
				ExecutionFee:        consensus.Median(executionFees, consensus.BigIntComparator).Int,
				DataAvailabilityFee: consensus.Median(dataAvailabilityFees, consensus.BigIntComparator).Int,
			}
		},
	)

	nativeTokenPrices := consensus.GetConsensusMapAggregator(
		p.lggr,
		"NativeTokenPrices",
		aggObs.NativeTokenPrices,
		twoFPlus1,
		// Aggregator function
		func(vals []cciptypes.BigInt) cciptypes.BigInt {
			return consensus.Median(vals, consensus.BigIntComparator)
		},
	)

	consensusObs := Observation{
		FChain:                fChains,
		FeeComponents:         feeComponents,
		NativeTokenPrices:     nativeTokenPrices,
		ChainFeeLatestUpdates: chainFeeUpdatesConsensus,
		Timestamp:             timestamp,
	}

	return consensusObs, nil
}

func aggregateObservations(aos []plugincommon.AttributedObservation[Observation]) AggregateObservation {
	aggObs := AggregateObservation{
		FeeComponents:         make(map[cciptypes.ChainSelector][]types.ChainFeeComponents),
		NativeTokenPrices:     make(map[cciptypes.ChainSelector][]cciptypes.BigInt),
		FChain:                make(map[cciptypes.ChainSelector][]int),
		ChainFeeLatestUpdates: make(map[cciptypes.ChainSelector][]time.Time),
		Timestamps:            []time.Time{},
	}

	for _, ao := range aos {
		obs := ao.Observation

		// FeeComponents
		for chainSel, feeComp := range obs.FeeComponents {
			aggObs.FeeComponents[chainSel] = append(aggObs.FeeComponents[chainSel], feeComp)
		}

		// NativeTokenPrices
		for chainSel, tokenPrice := range obs.NativeTokenPrices {
			aggObs.NativeTokenPrices[chainSel] = append(aggObs.NativeTokenPrices[chainSel], tokenPrice)
		}

		// FChain
		for chainSel, f := range obs.FChain {
			aggObs.FChain[chainSel] = append(aggObs.FChain[chainSel], f)
		}

		for chainSel, feeUpdate := range obs.ChainFeeLatestUpdates {
			aggObs.ChainFeeLatestUpdates[chainSel] = append(aggObs.ChainFeeLatestUpdates[chainSel], feeUpdate)
		}

		// Timestamps
		aggObs.Timestamps = append(aggObs.Timestamps, obs.Timestamp)
	}

	return aggObs
}
