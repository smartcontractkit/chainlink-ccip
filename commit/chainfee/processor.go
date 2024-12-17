package chainfee

import (
	"context"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type processor struct {
	oracleID        commontypes.OracleID
	destChain       cciptypes.ChainSelector
	lggr            logger.Logger
	homeChain       reader.HomeChain
	ccipReader      readerpkg.CCIPReader
	cfg             pluginconfig.CommitOffchainConfig
	chainSupport    plugincommon.ChainSupport
	metricsReporter MetricsReporter
	fRoleDON        int
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
	metricsReporter MetricsReporter,
) plugincommon.PluginProcessor[Query, Observation, Outcome] {
	return &processor{
		lggr:            lggr,
		oracleID:        oracleID,
		destChain:       destChain,
		homeChain:       homeChain,
		ccipReader:      ccipReader,
		fRoleDON:        fRoleDON,
		chainSupport:    chainSupport,
		cfg:             offChainConfig,
		metricsReporter: metricsReporter,
	}
}

func (p *processor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	return Query{}, nil
}

var _ plugincommon.PluginProcessor[Query, Observation, Outcome] = &processor{}

func (p *processor) Close() error {
	return nil
}
