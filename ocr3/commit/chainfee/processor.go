package chainfee

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/commontypes"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	plugincommon2 "github.com/smartcontractkit/chainlink-ccip/ocr3/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/ocr3/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/pluginconfig"
)

type processor struct {
	oracleID  commontypes.OracleID
	destChain cciptypes.ChainSelector
	// Don't use this logger directly but rather through logutil\.WithContextValues where possible
	lggr            logger.Logger
	homeChain       reader.HomeChain
	ccipReader      readerpkg.CCIPReader
	cfg             pluginconfig.CommitOffchainConfig
	chainSupport    plugincommon2.ChainSupport
	metricsReporter plugincommon2.MetricsReporter
	fRoleDON        int
	obs             observer
}

func NewProcessor(
	lggr logger.Logger,
	oracleID commontypes.OracleID,
	destChain cciptypes.ChainSelector,
	homeChain reader.HomeChain,
	ccipReader readerpkg.CCIPReader,
	offChainConfig pluginconfig.CommitOffchainConfig,
	chainSupport plugincommon2.ChainSupport,
	fRoleDON int,
	metricsReporter plugincommon2.MetricsReporter,
) plugincommon2.PluginProcessor[Query, Observation, Outcome] {
	var obs observer
	baseObs := newBaseObserver(
		ccipReader,
		destChain,
		oracleID,
		chainSupport,
	)
	if !offChainConfig.ChainFeeAsyncObserverDisabled {
		obs = newAsyncObserver(
			lggr,
			baseObs,
			offChainConfig.ChainFeeAsyncObserverSyncFreq,
			offChainConfig.ChainFeeAsyncObserverSyncTimeout,
		)
	} else {
		obs = baseObs
	}

	p := &processor{
		lggr:            lggr,
		oracleID:        oracleID,
		destChain:       destChain,
		homeChain:       homeChain,
		ccipReader:      ccipReader,
		fRoleDON:        fRoleDON,
		chainSupport:    chainSupport,
		cfg:             offChainConfig,
		metricsReporter: metricsReporter,
		obs:             obs,
	}
	return plugincommon2.NewTrackedProcessor(lggr, p, processorLabel, metricsReporter)
}

func (p *processor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	return Query{}, nil
}

var _ plugincommon2.PluginProcessor[Query, Observation, Outcome] = &processor{}

func (p *processor) Close() error {
	p.obs.close()
	return nil
}
