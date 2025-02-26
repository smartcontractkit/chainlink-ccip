package tokenprice

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	pkgreader "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type processor struct {
	oracleID commontypes.OracleID
	// Don't use this logger directly but rather through logutil\.WithContextValues where possible
	lggr             logger.Logger
	offChainCfg      pluginconfig.CommitOffchainConfig
	destChain        cciptypes.ChainSelector
	chainSupport     plugincommon.ChainSupport
	tokenPriceReader pkgreader.PriceReader
	homeChain        reader.HomeChain
	metricsReporter  plugincommon.MetricsReporter
	fRoleDON         int
	obs              observer
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
	metricsReporter plugincommon.MetricsReporter,
) plugincommon.PluginProcessor[Query, Observation, Outcome] {
	var obs observer
	baseObs := newBaseObserver(
		tokenPriceReader,
		destChain,
		oracleID,
		chainSupport,
		offChainCfg,
	)
	if !offChainCfg.TokenPriceAsyncObserverDisabled {
		obs = newAsyncObserver(
			lggr,
			baseObs,
			offChainCfg.TokenPriceAsyncObserverSyncFreq.Duration(),
			offChainCfg.TokenPriceAsyncObserverSyncTimeout.Duration(),
		)
	} else {
		obs = baseObs
	}
	p := &processor{
		oracleID:         oracleID,
		lggr:             lggr,
		offChainCfg:      offChainCfg,
		destChain:        destChain,
		chainSupport:     chainSupport,
		tokenPriceReader: tokenPriceReader,
		homeChain:        homeChain,
		fRoleDON:         fRoleDON,
		metricsReporter:  metricsReporter,
		obs:              obs,
	}
	return plugincommon.NewTrackedProcessor(lggr, p, processorsLabel, metricsReporter)
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
	lggr := logutil.WithContextValues(ctx, p.lggr)

	lggr.Infow("processing token price outcome")
	// If set to zero, no prices will be reported (i.e keystone feeds would be active).
	if p.offChainCfg.TokenPriceBatchWriteFrequency.Duration() == 0 {
		lggr.Debugw("TokenPriceBatchWriteFrequency is set to zero, no prices will be reported")
		return Outcome{}, nil
	}

	consensusObservation, err := p.getConsensusObservation(lggr, aos)
	if err != nil {
		return Outcome{}, fmt.Errorf("get consensus observation: %w", err)
	}

	tokenPriceOutcome := p.selectTokensForUpdate(lggr, consensusObservation)
	lggr.Infow(
		"outcome token prices",
		"tokenPrices", tokenPriceOutcome,
	)

	if len(tokenPriceOutcome) == 0 {
		lggr.Debugw("No token prices to report")
		return Outcome{}, nil
	}

	out := Outcome{TokenPrices: tokenPriceOutcome}
	return out, nil
}

func (p *processor) Close() error {
	p.obs.close()
	return nil
}

var _ plugincommon.PluginProcessor[Query, Observation, Outcome] = &processor{}
