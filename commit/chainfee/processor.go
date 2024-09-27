package chainfee

import (
	"context"

	"github.com/smartcontractkit/libocr/commontypes"

	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
)

type processor struct {
	oracleID     commontypes.OracleID
	destChain    cciptypes.ChainSelector
	lggr         logger.Logger
	homeChain    reader.HomeChain
	ccipReader   readerpkg.CCIPReader
	cfg          pluginconfig.CommitOffchainConfig
	chainSupport plugincommon.ChainSupport
	fRoleDON     int
}

func NewProcessor(
	oracleID commontypes.OracleID,
	lggr logger.Logger,
	destChain cciptypes.ChainSelector,
	homeChain reader.HomeChain,
	ccipReader readerpkg.CCIPReader,
	offChainConfig pluginconfig.CommitOffchainConfig,
	chainSupport plugincommon.ChainSupport,
	fRoleDON int,
) plugincommon.PluginProcessor[Query, Observation, Outcome] {
	return &processor{
		oracleID:     oracleID,
		lggr:         lggr,
		destChain:    destChain,
		homeChain:    homeChain,
		ccipReader:   ccipReader,
		fRoleDON:     fRoleDON,
		chainSupport: chainSupport,
		cfg:          offChainConfig,
	}
}

func (p *processor) Query(ctx context.Context, prevOutcome Outcome) (Query, error) {
	return Query{}, nil
}

var _ plugincommon.PluginProcessor[Query, Observation, Outcome] = &processor{}
