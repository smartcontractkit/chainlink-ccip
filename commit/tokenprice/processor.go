package tokenprice

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logger"
	pkgreader "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type processor struct {
	oracleID         commontypes.OracleID
	lggr             logger.Logger
	offChainCfg      pluginconfig.CommitOffchainConfig
	destChain        cciptypes.ChainSelector
	chainSupport     plugincommon.ChainSupport
	tokenPriceReader pkgreader.PriceReader
	homeChain        reader.HomeChain
	fRoleDON         int
}

func NewProcessor(
	oracleID commontypes.OracleID,
	lggr logger.Logger,
	offChainCfg pluginconfig.CommitOffchainConfig,
	destChain cciptypes.ChainSelector,
	chainSupport plugincommon.ChainSupport,
	tokenPriceReader pkgreader.PriceReader,
	homeChain reader.HomeChain,
	fRoleDON int,
) plugincommon.PluginProcessor[Query, Observation, Outcome] {
	return &processor{
		oracleID:         oracleID,
		lggr:             logger.NewProcessorLogWrapper(lggr, "TokenPrice"),
		offChainCfg:      offChainCfg,
		destChain:        destChain,
		chainSupport:     chainSupport,
		tokenPriceReader: tokenPriceReader,
		homeChain:        homeChain,
		fRoleDON:         fRoleDON,
	}
}

func (p *processor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	return Query{}, nil
}

func (p *processor) Outcome(
	ctx context.Context,
	_ Outcome,
	_ Query,
	aos []plugincommon.AttributedObservation[Observation],
) (Outcome, error) {
	p.lggr.Infow("processing token price outcome")
	// If set to zero, no prices will be reported (i.e keystone feeds would be active).
	if p.offChainCfg.TokenPriceBatchWriteFrequency.Duration() == 0 {
		p.lggr.Debugw("TokenPriceBatchWriteFrequency is set to zero, no prices will be reported")
		return Outcome{}, nil
	}

	consensusObservation, err := p.getConsensusObservation(aos)
	if err != nil {
		return Outcome{}, fmt.Errorf("get consensus observation: %w", err)
	}

	tokenPriceOutcome := p.selectTokensForUpdate(consensusObservation)
	p.lggr.Infow(
		"outcome token prices",
		"tokenPrices", tokenPriceOutcome,
	)
	return Outcome{
		TokenPrices: tokenPriceOutcome,
	}, nil
}

func (p *processor) Close() error {
	return nil
}

var _ plugincommon.PluginProcessor[Query, Observation, Outcome] = &processor{}
