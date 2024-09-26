package chainfee

import (
	"context"
	"fmt"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/mathslib"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
)

type processor struct {
	destChain  cciptypes.ChainSelector
	lggr       logger.Logger
	homeChain  reader.HomeChain
	ccipReader readerpkg.CCIPReader
	cfg        pluginconfig.CommitOffchainConfig
	fRoleDON   int
}

func NewProcessor(
	lggr logger.Logger,
	destChain cciptypes.ChainSelector,
	homeChain reader.HomeChain,
	ccipReader readerpkg.CCIPReader,
	offChainConfig pluginconfig.CommitOffchainConfig,
	fRoleDON int,
) plugincommon.PluginProcessor[Query, Observation, Outcome] {
	return &processor{
		lggr:       lggr,
		destChain:  destChain,
		homeChain:  homeChain,
		ccipReader: ccipReader,
		fRoleDON:   fRoleDON,
		cfg:        offChainConfig,
	}
}

func (p *processor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	return Query{}, nil
}

func (p *processor) Observation(
	ctx context.Context,
	prevOutcome Outcome,
	query Query,
) (Observation, error) {
	// Get the fee components for all available chains that we can read from
	feeComponents := p.ccipReader.GetAvailableChainsFeeComponents(ctx)
	// Get the native token prices for all available chains that we can read from
	nativeTokenPrices := p.ccipReader.GetWrappedNativeTokenPriceUSD(ctx, maps.Keys(feeComponents))
	// Get the latest chain fee price updates for the source chains
	timestampedPriceUpdates := p.ccipReader.GetChainFeePriceUpdate(ctx, maps.Keys(feeComponents))
	// Convert the timestamped price updates to a map of chain fee updates
	chainFeeUpdates := FeeUpdatesFromTimestampedBig(timestampedPriceUpdates)

	fChain := p.ObserveFChain()

	p.lggr.Infow("observed fee components",
		"feeComponents", feeComponents,
		"nativeTokenPrices", nativeTokenPrices,
		"chainFeeUpdates", chainFeeUpdates,
		"fChain", fChain,
	)

	return Observation{
		FChain:            fChain,
		FeeComponents:     feeComponents,
		NativeTokenPrices: nativeTokenPrices,
		ChainFeeUpdates:   chainFeeUpdates,
		TimestampNow:      time.Now().UTC(),
	}, nil
}

func (p *processor) Outcome(
	prevOutcome Outcome,
	query Query,
	aos []plugincommon.AttributedObservation[Observation],
) (Outcome, error) {

	consensusObs, err := p.getConsensusObservation(aos)
	if err != nil {
		return Outcome{}, fmt.Errorf("failed to get consensus observation: %w", err)
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
		usdPerFeeToken := consensusObs.NativeTokenPrices[chain].Int

		// Example with Wei as the lowest denominator and Eth as the Fee token
		// usdPerEthToken = Xe18USD18
		// Price per Wei = Xe18USD18/1e18 = XUSD18
		// 1 gas = 1 wei = XUSD18
		// execFee = 30 Gwei = 30e9 wei = 30e9 * XUSD18
		chainFeeUsd := ComponentsUSDPrices{
			ExecutionFeePriceUSD: mathslib.CalculateUsdPerUnitGas(feeComp.ExecutionFee, usdPerFeeToken),
			DataAvFeePriceUSD:    mathslib.CalculateUsdPerUnitGas(feeComp.DataAvailabilityFee, usdPerFeeToken),
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

func (p *processor) ValidateObservation(
	prevOutcome Outcome,
	query Query,
	ao plugincommon.AttributedObservation[Observation],
) error {
	// TODO: Validate token prices
	//  Validate FChain
	//  Validate no nil token prices
	return nil
}

func validateObservedGasPrices(gasPrices []cciptypes.GasPriceChain) error {
	// Duplicate gas prices must not appear for the same chain and must not be empty.
	gasPriceChains := mapset.NewSet[cciptypes.ChainSelector]()
	for _, g := range gasPrices {
		if gasPriceChains.Contains(g.ChainSel) {
			return fmt.Errorf("duplicate gas price for chain %d", g.ChainSel)
		}
		gasPriceChains.Add(g.ChainSel)
		if g.GasPrice.IsEmpty() {
			return fmt.Errorf("gas price must not be empty")
		}
	}

	return nil
}

var _ plugincommon.PluginProcessor[Query, Observation, Outcome] = &processor{}
