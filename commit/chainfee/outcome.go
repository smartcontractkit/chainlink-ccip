package chainfee

import (
	"fmt"
	"sort"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/mathslib"

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
	if len(aggObs.Timestamps) < fDestChain {
		return Observation{},
			fmt.Errorf("not enough observations for timestamps to reach consensus, have %d, need %d",
				len(aggObs.Timestamps), fDestChain)
	}
	timestamp := consensus.Median(aggObs.Timestamps, consensus.TimestampComparator)

	chainFeeUpdatesConsensus := consensus.GetConsensusMapAggregator(
		p.lggr,
		"ChainFeeUpdates",
		aggObs.ChainFeeUpdates,
		consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(fDestChain)),
		plugincommon.ChainFeeUpdateAggregator,
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
		FChain:            fChains,
		FeeComponents:     feeComponents,
		NativeTokenPrices: nativeTokenPrices,
		ChainFeeUpdates:   chainFeeUpdatesConsensus,
		Timestamp:         timestamp,
	}

	return consensusObs, nil
}

func aggregateObservations(aos []plugincommon.AttributedObservation[Observation]) AggregateObservation {
	aggObs := AggregateObservation{
		FeeComponents:     make(map[cciptypes.ChainSelector][]types.ChainFeeComponents),
		NativeTokenPrices: make(map[cciptypes.ChainSelector][]cciptypes.BigInt),
		FChain:            make(map[cciptypes.ChainSelector][]int),
		ChainFeeUpdates:   make(map[cciptypes.ChainSelector][]plugincommon.ChainFeeUpdate),
		Timestamps:        []time.Time{},
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

		for chainSel, feeUpdate := range obs.ChainFeeUpdates {
			aggObs.ChainFeeUpdates[chainSel] = append(aggObs.ChainFeeUpdates[chainSel], feeUpdate)
		}

		// Timestamps
		aggObs.Timestamps = append(aggObs.Timestamps, obs.Timestamp)
	}

	return aggObs
}

// getGasPricesToUpdate checks which chain fees need to be updated based on the observed chain fee prices and
// the fee quoter updates.
// A chain fee is selected for update if it meets one of 2 conditions:
// 1. If time passed since the last update is greater than the stale threshold.
// 2. If deviation between the fee quoter and feed exceeds the chain's configured threshold.
func (p *processor) getGasPricesToUpdate(
	currentChainUSDFees map[cciptypes.ChainSelector]plugincommon.ChainFeeUSDPrices,
	latestUpdates map[cciptypes.ChainSelector]plugincommon.ChainFeeUpdate,
	obsTimestamp time.Time,
) []cciptypes.GasPriceChain {
	var gasPrices []cciptypes.GasPriceChain
	feeInfo := p.cfg.FeeInfo

	for chain, currentChainFee := range currentChainUSDFees {
		packedFee := cciptypes.NewBigInt(currentChainFee.ToPackedFee())
		lastUpdate, exists := latestUpdates[chain]
		nextUpdateTime := lastUpdate.Timestamp.Add(p.cfg.RemoteGasPriceBatchWriteFrequency.Duration())
		// If the chain is not in the fee quoter updates or is stale, then we should update it
		if !exists || obsTimestamp.After(nextUpdateTime) {
			gasPrices = append(gasPrices, cciptypes.GasPriceChain{
				ChainSel: chain,
				GasPrice: packedFee,
			})
			continue
		}

		if feeInfo == nil {
			continue
		}
		ci, ok := feeInfo[chain]
		if !ok {
			p.lggr.Warnf("could not find fee info for chain %d", chain)
			continue
		}

		// Checks if executionFee or dataAvFee deviates from the last update
		if mathslib.Deviates(
			currentChainFee.ExecutionFeePriceUSD,
			lastUpdate.ChainFee.ExecutionFeePriceUSD,
			ci.ExecDeviationPPB.Int64()) ||
			mathslib.Deviates(
				currentChainFee.DataAvFeePriceUSD,
				lastUpdate.ChainFee.DataAvFeePriceUSD,
				ci.DataAvDeviationPPB.Int64()) {
			gasPrices = append(gasPrices, cciptypes.GasPriceChain{
				ChainSel: chain,
				GasPrice: packedFee,
			})
		}
	}

	// sort chainFeeUSDPrices based on chainSel
	sort.Slice(gasPrices, func(i, j int) bool {
		return gasPrices[i].ChainSel < gasPrices[j].ChainSel
	})
	return gasPrices
}
