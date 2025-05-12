package chainfee

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/mathslib"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/consensus"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func (p *processor) Outcome(
	ctx context.Context,
	_ Outcome,
	_ Query,
	aos []plugincommon.AttributedObservation[Observation],
) (Outcome, error) {
	lggr := logutil.WithContextValues(ctx, p.lggr)

	consensusObs, err := p.getConsensusObservation(lggr, aos)
	if err != nil {
		return Outcome{}, fmt.Errorf("get consensus observation: %w", err)
	}

	// No need to update yet
	if len(consensusObs.FeeComponents) == 0 {
		lggr.Warn("no consensus on fee components, nothing to update",
			"consensusObs", consensusObs)
		return Outcome{}, nil
	}

	chainFeeUSDPrices := make(map[cciptypes.ChainSelector]ComponentsUSDPrices)
	// We need to report a packed GasPrice
	// The packed GasPrice is a 224-bit integer with the following format:
	// (dataAvFeePriceUSD) << 112 | (executionFeePriceUSD)
	//
	// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L498
	// In next loop we calculate the price in USD for the data availability and execution fees.
	// And getGasPricesToUpdate will select and calculate the **packed** gas price to update based.
	//
	//nolint:lll
	for chain, feeComp := range consensusObs.FeeComponents {
		// The price, in USD with 18 decimals, per 1e18 of the smallest token denomination.
		// 1 USDC = 1.00 USD per full token, each full token is 1e6 units -> 1 * 1e18 * 1e18 / 1e6 = 1e30
		// 1 ETH = 2,000 USD per full token, each full token is 1e18 units -> 2000 * 1e18 * 1e18 / 1e18 = 2_000e18
		// 1 LINK = 5.00 USD per full token, each full token is 1e18 units -> 5 * 1e18 * 1e18 / 1e18 = 5e18
		usdPerFeeToken, ok := consensusObs.NativeTokenPrices[chain]
		if !ok {
			lggr.Warnw("missing native token price for chain, chain fee will not be updated",
				"chain", chain,
			)
			continue
		}
		lggr.Debugw("USD per fee token", "chain", chain, "usdPerFeeToken", usdPerFeeToken)

		// Example with Wei as the lowest denominator and Eth as the Fee token
		// usdPerEthToken = Xe18USD18
		// Price per Wei = Xe18USD18/1e18 = XUSD18
		// 1 gas = 1 wei = XUSD18
		// execFee = 30 Gwei = 30e9 wei = 30e9 * XUSD18
		execFee, err := mathslib.CalculateUsdPerUnitGas(chain, feeComp.ExecutionFee, usdPerFeeToken.Int)
		if err != nil {
			lggr.Errorw("error calculating USD per unit gas", "chain", chain, "err", err)
			continue
		}
		daFee, err := mathslib.CalculateUsdPerUnitGas(chain, feeComp.DataAvailabilityFee, usdPerFeeToken.Int)
		if err != nil {
			lggr.Errorw("error calculating USD per unit gas", "chain", chain, "err", err)
			continue
		}
		chainFeeUsd := ComponentsUSDPrices{
			ExecutionFeePriceUSD: execFee,
			DataAvFeePriceUSD:    daFee,
		}

		chainFeeUSDPrices[chain] = chainFeeUsd
	}

	gasPrices := p.getGasPricesToUpdate(
		lggr,
		chainFeeUSDPrices,
		consensusObs.ChainFeeUpdates,
		consensusObs.TimestampNow,
	)

	// sort chainFeeUSDPrices based on chainSel
	sort.Slice(gasPrices, func(i, j int) bool {
		return gasPrices[i].ChainSel < gasPrices[j].ChainSel
	})

	lggr.Infow("Gas Prices Outcome",
		"gasPrices", gasPrices,
		"consensusTimestamp", consensusObs.TimestampNow,
	)
	out := Outcome{GasPrices: gasPrices}
	return out, nil
}

func (p *processor) getConsensusObservation(
	lggr logger.Logger, aos []plugincommon.AttributedObservation[Observation],
) (Observation, error) {
	aggObs := aggregateObservations(aos)

	// consensus on the fChain map uses the role DON F value
	// because all nodes can observe the home chain.
	donThresh := consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(p.fRoleDON))
	fChains := consensus.GetConsensusMap(lggr, "fChain", aggObs.FChain, donThresh)

	fDestChain, exists := fChains[p.destChain]
	if !exists {
		return Observation{},
			fmt.Errorf("no consensus value for fDestChain, destChain: %d", p.destChain)
	}

	if consensus.LtTwoFPlusOne(fDestChain, len(aggObs.Timestamps)) {
		return Observation{},
			fmt.Errorf("not enough observations for timestamps to reach consensus, have %d, need %d",
				len(aggObs.Timestamps), consensus.TwoFPlus1(fDestChain))
	}
	timestamp := consensus.Median(aggObs.Timestamps, consensus.TimestampComparator)

	chainFeeUpdatesConsensus := consensus.GetConsensusMapAggregator(
		lggr,
		"ChainFeeUpdates",
		aggObs.ChainFeeUpdates,
		consensus.MakeConstantThreshold[cciptypes.ChainSelector](consensus.TwoFPlus1(fDestChain)),
		chainFeeUpdateAggregator,
	)

	twoFChainPlus1 := consensus.MakeMultiThreshold(fChains, consensus.TwoFPlus1)

	feeComponents := consensus.GetConsensusMapAggregator(
		lggr,
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
		lggr,
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
	lggr logger.Logger,
	currentChainUSDFees map[cciptypes.ChainSelector]ComponentsUSDPrices,
	latestUpdates map[cciptypes.ChainSelector]Update,
	consensusTimestamp time.Time,
) []cciptypes.GasPriceChain {
	var gasPrices []cciptypes.GasPriceChain

	for chain, currentChainFee := range currentChainUSDFees {
		chainCfg, err := p.homeChain.GetChainConfig(chain)
		if err != nil {
			lggr.Errorw("error getting chain config", "chain", chain, "err", err)
			continue
		}

		feeConfig := chainCfg.Config
		packedFee := cciptypes.NewBigInt(FeeComponentsToPackedFee(currentChainFee))
		lastUpdate, exists := latestUpdates[chain]
		lggr := logger.With(lggr,
			"chain", chain,
			"consensusTimestamp", consensusTimestamp,
			"currentChainFee", currentChainFee,
			"packedFee", packedFee,
			"lastUpdate", lastUpdate)
		// If the chain is not in the fee quoter updates or is stale, then we should update it
		if !exists {
			lggr.Infow("chain fee update needed: no previous update exists")
			gasPrices = append(gasPrices, cciptypes.GasPriceChain{
				ChainSel: chain,
				GasPrice: packedFee,
			})
			continue
		}

		nextUpdateTime := lastUpdate.Timestamp.Add(p.cfg.RemoteGasPriceBatchWriteFrequency.Duration())
		if consensusTimestamp.After(nextUpdateTime) {
			lggr.Infow("chain fee update needed: heartbeat time passed",
				"nextUpdateTime", nextUpdateTime,
				"consensusTimestamp", consensusTimestamp,
				"heartbeatInterval", p.cfg.RemoteGasPriceBatchWriteFrequency)
			gasPrices = append(gasPrices, cciptypes.GasPriceChain{
				ChainSel: chain,
				GasPrice: packedFee,
			})
			continue
		}

		if feeConfig.ChainFeeDeviationDisabled {
			lggr.Debugw("chain fee deviation disabled")
			continue
		}

		// Validating later as chain can be updated even if the config is invalid when write frequency is reached
		if feeConfig.Validate() != nil {
			lggr.Errorw("invalid fee config for chain", "err", err)
			continue
		}

		executionFeeDeviates := mathslib.Deviates(
			currentChainFee.ExecutionFeePriceUSD,
			lastUpdate.ChainFee.ExecutionFeePriceUSD,
			feeConfig.GasPriceDeviationPPB.Int64(),
		)

		dataAvFeeDeviates := mathslib.Deviates(
			currentChainFee.DataAvFeePriceUSD,
			lastUpdate.ChainFee.DataAvFeePriceUSD,
			feeConfig.DAGasPriceDeviationPPB.Int64(),
		)

		if executionFeeDeviates || dataAvFeeDeviates {
			lggr.Infow(
				"chain fee update needed: deviation threshold exceeded for either execution or data availability fee",
				"executionFeeDeviates", executionFeeDeviates,
				"dataAvFeeDeviates", dataAvFeeDeviates,
				"executionFeeDeviationPPB", feeConfig.GasPriceDeviationPPB,
				"dataAvFeeDeviationPPB", feeConfig.DAGasPriceDeviationPPB)
			gasPrices = append(gasPrices, cciptypes.GasPriceChain{
				ChainSel: chain,
				GasPrice: packedFee,
			})
		} else {
			lggr.Debugw("chain fee update not needed: within deviation thresholds",
				"executionFeeDeviationPPB", feeConfig.GasPriceDeviationPPB,
				"dataAvFeeDeviationPPB", feeConfig.DAGasPriceDeviationPPB)
		}
	}

	return gasPrices
}

// chainFeeUpdateAggregator aggregates a slice of ChainFeeUpdates into a single Update
// by taking the median of each price component and the timestamps
func chainFeeUpdateAggregator(updates []Update) Update {
	execFeeUSDs := make([]*big.Int, len(updates))
	dataAvFeeUSDs := make([]*big.Int, len(updates))
	timestamps := make([]time.Time, len(updates))
	for i := range updates {
		execFeeUSDs[i] = updates[i].ChainFee.ExecutionFeePriceUSD
		dataAvFeeUSDs[i] = updates[i].ChainFee.DataAvFeePriceUSD
		timestamps[i] = updates[i].Timestamp
	}
	medianExecFeeUSD := consensus.Median(execFeeUSDs, func(a, b *big.Int) bool {
		return a.Cmp(b) == -1
	})
	medianDataAvFeeUSD := consensus.Median(dataAvFeeUSDs, func(a, b *big.Int) bool {
		return a.Cmp(b) == -1
	})

	return Update{
		ChainFee: ComponentsUSDPrices{
			ExecutionFeePriceUSD: medianExecFeeUSD,
			DataAvFeePriceUSD:    medianDataAvFeeUSD,
		},
		Timestamp: consensus.TimestampsMedian(timestamps),
	}
}

// FeeComponentsToPackedFee is a Bitwise operation:
// (dataAvFeeUSD << 112) | executionFeeUSD
//
// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L498
//
//nolint:lll
func FeeComponentsToPackedFee(c ComponentsUSDPrices) *big.Int {
	daShifted := new(big.Int).Lsh(c.DataAvFeePriceUSD, 112)
	return new(big.Int).Or(daShifted, c.ExecutionFeePriceUSD)
}
