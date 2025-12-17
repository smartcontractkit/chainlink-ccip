package chainfee

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/commontypes"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/asynclib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

const (
	opGetChainsFeeComponents  = "getChainsFeeComponents"
	opGetNativeTokenPrices    = "getNativeTokenPrices"
	opGetChainFeePriceUpdates = "getChainFeePriceUpdates"
	opObserveFChain           = "observeFChain"

	// pool sizes are set to 1 to avoid blocking the runner if one operation blocks forever.
	poolSizeGetChainsFeeComponents  = 1
	poolSizeGetNativeTokenPrices    = 1
	poolSizeGetChainFeePriceUpdates = 1
	poolSizeObserveFChain           = 1
)

var (
	poolSizePerOp = map[string]int{
		opGetChainsFeeComponents:  poolSizeGetChainsFeeComponents,
		opGetNativeTokenPrices:    poolSizeGetNativeTokenPrices,
		opGetChainFeePriceUpdates: poolSizeGetChainFeePriceUpdates,
		opObserveFChain:           poolSizeObserveFChain,
	}
)

type processor struct {
	oracleID  commontypes.OracleID
	destChain cciptypes.ChainSelector
	// Don't use this logger directly but rather through logutil\.WithContextValues where possible
	lggr            logger.Logger
	homeChain       reader.HomeChain
	ccipReader      readerpkg.CCIPReader
	cfg             pluginconfig.CommitOffchainConfig
	chainSupport    plugincommon.ChainSupport
	metricsReporter plugincommon.MetricsReporter
	fRoleDON        int
	obs             observer

	// the same runner is used for operations across rounds just in case
	// we run into a situation where a particular operation blocks forever
	// and we keep enqueuing it, which could cause a goroutine leak.
	runner *asynclib.AsyncOpsRunner
}

func NewProcessor(
	lggr logger.Logger,
	oracleID commontypes.OracleID,
	destChain cciptypes.ChainSelector,
	homeChain reader.HomeChain,
	ccipReader readerpkg.CCIPReader,
	offChainConfig pluginconfig.CommitOffchainConfig,
	chainSupport plugincommon.ChainSupport,
	fRoleDON int,
	metricsReporter plugincommon.MetricsReporter,
) plugincommon.PluginProcessor[Query, Observation, Outcome] {
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

	runner, err := asynclib.NewAsyncOpsRunner(poolSizePerOp)
	if err != nil {
		lggr.Errorw("failed to create async ops runner", "err", err)
		return nil
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
		runner:          runner,
	}
	return plugincommon.NewTrackedProcessor(lggr, p, processorLabel, metricsReporter)
}

func (p *processor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	return Query{}, nil
}

var _ plugincommon.PluginProcessor[Query, Observation, Outcome] = &processor{}

func (p *processor) Close() error {
	p.obs.close()
	return nil
}
