package chainfee

import (
	"context"
	"fmt"
	"math/big"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"golang.org/x/exp/maps"

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

// nolint: revive
func NewProcessor(
	lggr logger.Logger,
	destChain cciptypes.ChainSelector,
	homeChain reader.HomeChain,
	ccipReader readerpkg.CCIPReader,
	offChainConfig pluginconfig.CommitOffchainConfig,
	fRoleDON int,
) *processor {
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
	chainFeePriceUpdates := p.ccipReader.GetChainFeePriceUpdate(ctx, maps.Keys(feeComponents))

	return Observation{
		FChain:            p.ObserveFChain(),
		FeeComponents:     feeComponents,
		NativeTokenPrices: nativeTokenPrices,
		ChainFeeUpdates:   chainFeePriceUpdates,
		Timestamp:         time.Now().UTC(),
	}, nil
}

func (p *processor) Outcome(
	prevOutcome Outcome,
	query Query,
	aos []plugincommon.AttributedObservation[Observation],
) (Outcome, error) {

	consensusObs, err := p.getConsensusObservation(aos)
	if err != nil {
		return Outcome{}, err
	}
	// No need to update yet
	if len(consensusObs.FeeComponents) == 0 {
		return Outcome{}, nil
	}

	// Stop early if earliest updated timestamp is still fresh
	//earliestUpdateTime := consensus.EarliestTimestamp(maps.Values(consensusObs.ChainFeeUpdates))
	//nextUpdateTime := earliestUpdateTime.Add(p.ChainFeePriceBatchWriteFrequency.Duration())
	//if nextUpdateTime.After(consensusObs.Timestamp) {
	//	return Outcome{}, nil
	//}

	chainFeeUSDPrices := make(map[cciptypes.ChainSelector]plugincommon.ChainFeeUSDPrices)
	for chain, feeComp := range consensusObs.FeeComponents {
		// We need to report a packed GasPrice
		// The packed GasPrice is a 224-bit integer with the following format:
		// (dataAvFeePriceUSD) << 112 | (executionFeePriceUSD)
		// nolint:lll
		// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L498
		nativeTokenPriceUSD := consensusObs.NativeTokenPrices[chain].Int

		// Calculate the price in USD for the data availability and execution fees.
		// Raw fee components are in native token units
		chainFeeUsd := plugincommon.ChainFeeUSDPrices{
			ExecutionFeePriceUSD: new(big.Int).Mul(nativeTokenPriceUSD, feeComp.ExecutionFee),
			DataAvFeePriceUSD:    new(big.Int).Mul(nativeTokenPriceUSD, feeComp.DataAvailabilityFee),
		}

		chainFeeUSDPrices[chain] = chainFeeUsd
	}

	return Outcome{
		GasPrices: p.getGasPricesToUpdate(
			chainFeeUSDPrices,
			consensusObs.ChainFeeUpdates,
			consensusObs.Timestamp,
		),
	}, nil
}

func (p *processor) ValidateObservation(
	prevOutcome Outcome,
	query Query,
	ao plugincommon.AttributedObservation[Observation],
) error {
	// TODO: Validate token prices
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
