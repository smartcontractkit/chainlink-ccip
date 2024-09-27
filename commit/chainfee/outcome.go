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

func (p *processor) Outcome(
	prevOutcome Outcome,
	query Query,
	aos []plugincommon.AttributedObservation[Observation],
) (Outcome, error) {
	consensusObs, err := p.getConsensusObservation(aos)
	if err != nil {
		return Outcome{}, fmt.Errorf("get consensus observation: %w", err)
	}
	// No need to update yet
	if len(consensusObs.FeeComponents) == 0 {
		p.lggr.Debug("no consensus on fee components, nothing to update",
			"consensusObs", consensusObs)
		return Outcome{}, nil
	}

	// Stop early if earliest updated timestamp is still fresh
	//earliestUpdateTime := consensus.EarliestTimestamp(maps.Values(consensusObs.ChainFeeUpdates))
	//nextUpdateTime := earliestUpdateTime.Add(p.ChainFeePriceBatchWriteFrequency.Duration())
	//if nextUpdateTime.After(consensusObs.TimestampNow) {
	//	return Outcome{}, nil
	//}

	chainFeeUSDPrices := make(map[cciptypes.ChainSelector]ComponentsUSDPrices)
	// We need to report a packed GasPrice
	// The packed GasPrice is a 224-bit integer with the following format:
	// (dataAvFeePriceUSD) << 112 | (executionFeePriceUSD)
	// nolint:lll
	// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L498
	// In next loop we calculate the price in USD for the data availability and execution fees.
	// And getGasPricesToUpdate will select and calculate the **packed** gas price to update based.
	for chain, feeComp := range consensusObs.FeeComponents {
		// The price, in USD with 18 decimals, per 1e18 of the smallest token denomination.
		// 1 USDC = 1.00 USD per full token, each full token is 1e6 units -> 1 * 1e18 * 1e18 / 1e6 = 1e30
		// 1 ETH = 2,000 USD per full token, each full token is 1e18 units -> 2000 * 1e18 * 1e18 / 1e18 = 2_000e18
		// 1 LINK = 5.00 USD per full token, each full token is 1e18 units -> 5 * 1e18 * 1e18 / 1e18 = 5e18
		usdPerFeeToken, ok := consensusObs.NativeTokenPrices[chain]
		if !ok {
			p.lggr.Warnw("missing native token price for chain",
				"chain", chain,
			)
			continue
		}

		// Example with Wei as the lowest denominator and Eth as the Fee token
		// usdPerEthToken = Xe18USD18
		// Price per Wei = Xe18USD18/1e18 = XUSD18
		// 1 gas = 1 wei = XUSD18
		// execFee = 30 Gwei = 30e9 wei = 30e9 * XUSD18
		chainFeeUsd := ComponentsUSDPrices{
			ExecutionFeePriceUSD: mathslib.CalculateUsdPerUnitGas(feeComp.ExecutionFee, usdPerFeeToken.Int),
			DataAvFeePriceUSD:    mathslib.CalculateUsdPerUnitGas(feeComp.DataAvailabilityFee, usdPerFeeToken.Int),
		}

		chainFeeUSDPrices[chain] = chainFeeUsd
	}

	gasPrices := p.getGasPricesToUpdate(
		chainFeeUSDPrices,
		consensusObs.ChainFeeUpdates,
		consensusObs.TimestampNow,
	)

	// sort chainFeeUSDPrices based on chainSel
	sort.Slice(gasPrices, func(i, j int) bool {
		return gasPrices[i].ChainSel < gasPrices[j].ChainSel
	})

	p.lggr.Infow("Gas Prices Outcome",
		"gasPrices", gasPrices,
	)

	return Outcome{
		GasPrices: gasPrices,
	}, nil
}

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
			fmt.Errorf("no consensus value for fDestChain, destChain: %d, fChainObs: %+v, threshold: %d",
				p.destChain, aggObs.FChain, consensus.TwoFPlus1(p.fRoleDON))
	}

	if len(aggObs.Timestamps) < int(consensus.TwoFPlus1(fDestChain)) {
		return Observation{},
			fmt.Errorf("not enough observations for timestamps to reach consensus, have %d, need %d",
				len(aggObs.Timestamps), consensus.TwoFPlus1(fDestChain))
	}
	timestamp := consensus.Median(aggObs.Timestamps, consensus.TimestampComparator)

	chainFeeUpdatesConsensus := consensus.GetConsensusMapAggregator(
		p.lggr,
		"ChainFeeUpdates",
		aggObs.ChainFeeUpdates,
		consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(fDestChain)),
		ChainFeeUpdateAggregator,
	)

	twoFChainPlus1 := consensus.MakeMultiThreshold(fChains, consensus.TwoFPlus1)

	feeComponents := consensus.GetConsensusMapAggregator(
		p.lggr,
		"FeeComponents",
		aggObs.FeeComponents,
		twoFChainPlus1,
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
		twoFChainPlus1,
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
		TimestampNow:      timestamp,
	}

	return consensusObs, nil
}

func aggregateObservations(aos []plugincommon.AttributedObservation[Observation]) AggregateObservation {
	aggObs := AggregateObservation{
		FeeComponents:     make(map[cciptypes.ChainSelector][]types.ChainFeeComponents),
		NativeTokenPrices: make(map[cciptypes.ChainSelector][]cciptypes.BigInt),
		FChain:            make(map[cciptypes.ChainSelector][]int),
		ChainFeeUpdates:   make(map[cciptypes.ChainSelector][]Update),
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
		aggObs.Timestamps = append(aggObs.Timestamps, obs.TimestampNow)
	}

	return aggObs
}

// getGasPricesToUpdate checks which chain fees need to be updated based on the observed chain fee prices and
// the fee quoter updates.
// A chain fee is selected for update if it meets one of 2 conditions:
// 1. If time passed since the last update is greater than the stale threshold.
// 2. If deviation between the fee quoter and latest observed chain fee exceeds the chain's configured threshold.
func (p *processor) getGasPricesToUpdate(
	currentChainUSDFees map[cciptypes.ChainSelector]ComponentsUSDPrices,
	latestUpdates map[cciptypes.ChainSelector]Update,
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

		executionFeeDeviates := mathslib.Deviates(
			currentChainFee.ExecutionFeePriceUSD,
			lastUpdate.ChainFee.ExecutionFeePriceUSD,
			ci.ExecDeviationPPB.Int64(),
		)

		dataAvFeeDeviates := mathslib.Deviates(
			currentChainFee.DataAvFeePriceUSD,
			lastUpdate.ChainFee.DataAvFeePriceUSD,
			ci.DataAvailabilityDeviationPPB.Int64(),
		)

		if executionFeeDeviates || dataAvFeeDeviates {
			gasPrices = append(gasPrices, cciptypes.GasPriceChain{
				ChainSel: chain,
				GasPrice: packedFee,
			})
		}
	}

	return gasPrices
}