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

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
)

type processor struct {
	lggr                          logger.Logger
	homeChain                     reader.HomeChain
	chainSupport                  plugincommon.ChainSupport
	ccipReader                    readerpkg.CCIPReader
	TokenPriceBatchWriteFrequency commonconfig.Duration
	bigF                          int
}

func NewProcessor(
	lggr logger.Logger,
	homeChain reader.HomeChain,
	chainSupport plugincommon.ChainSupport,
	ccipReader readerpkg.CCIPReader,
	TokenPriceBatchWriteFrequency commonconfig.Duration,
	bigF int,
) *processor {
	return &processor{
		lggr:                          lggr,
		homeChain:                     homeChain,
		chainSupport:                  chainSupport,
		ccipReader:                    ccipReader,
		TokenPriceBatchWriteFrequency: TokenPriceBatchWriteFrequency,
		bigF:                          bigF,
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
	return Observation{
		FChain:        p.ObserveFChain(),
		FeeComponents: p.ObserveFeeComponents(ctx),
		Timestamp:     time.Now().UTC(),
	}, nil
}

func (p *processor) Outcome(
	prevOutcome Outcome,
	query Query,
	aos []plugincommon.AttributedObservation[Observation],
) (Outcome, error) {
	return Outcome{}, nil
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
