package chainfee

import (
	"context"
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
)

type processor struct {
	destChain                        cciptypes.ChainSelector
	lggr                             logger.Logger
	homeChain                        reader.HomeChain
	chainSupport                     plugincommon.ChainSupport
	ccipReader                       readerpkg.CCIPReader
	ChainFeePriceBatchWriteFrequency commonconfig.Duration
	bigF                             int
}

// nolint: revive
func NewProcessor(
	lggr logger.Logger,
	destChain cciptypes.ChainSelector,
	homeChain reader.HomeChain,
	chainSupport plugincommon.ChainSupport,
	ccipReader readerpkg.CCIPReader,
	chainFeePriceBatchWriteFrequency commonconfig.Duration,
	bigF int,
) *processor {
	return &processor{
		lggr:                             lggr,
		destChain:                        destChain,
		homeChain:                        homeChain,
		chainSupport:                     chainSupport,
		ccipReader:                       ccipReader,
		ChainFeePriceBatchWriteFrequency: chainFeePriceBatchWriteFrequency,
		bigF:                             bigF,
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
	feeComponents := p.ccipReader.GetAvailableChainsFeeComponents(ctx)
	nativeTokenPrices := p.ccipReader.GetWrappedNativeTokenPriceUSD(ctx, maps.Keys(feeComponents))
	return Observation{
		FChain:           p.ObserveFChain(),
		FeeComponents:    feeComponents,
		NativeTokenPrice: nativeTokenPrices,
		Timestamp:        time.Now().UTC(),
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
	if !consensusObs.ShouldUpdate || len(consensusObs.FeeComponents) == 0 {
		return Outcome{}, nil
	}
	gasPrices := make([]cciptypes.GasPriceChain, 0, len(consensusObs.FeeComponents))
	for chain, feeComp := range consensusObs.FeeComponents {
		dataAvailabilityPrice := cciptypes.NewBigIntFromInt64(1).
			Mul(consensusObs.NativeTokenPrices[chain].Int,
				feeComp.DataAvailabilityFee)
		execPrice := cciptypes.NewBigIntFromInt64(1).
			Mul(consensusObs.NativeTokenPrices[chain].Int,
				feeComp.ExecutionFee)

		// Bitwise operation here like:
		// gasPrice << 112 | nativeTokenPrice * executionFee
		// nolint:lll
		// https://github.com/smartcontractkit/chainlink/blob/60e8b1181dd74b66903cf5b9a8427557b85357ec/contracts/src/v0.8/ccip/FeeQuoter.sol#L498
		price := dataAvailabilityPrice.Lsh(dataAvailabilityPrice, 112)
		combinedPrice := price.Or(price, execPrice)

		gasPrice := cciptypes.GasPriceChain{
			ChainSel: chain,
			GasPrice: cciptypes.NewBigInt(combinedPrice),
		}
		gasPrices = append(gasPrices, gasPrice)
	}

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
