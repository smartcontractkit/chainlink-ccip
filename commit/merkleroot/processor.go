package merkleroot

import (
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/shared"
)

// Processor is the processor responsible for
// reading next messages and building merkle roots for them,
// It's setup to use RMN to query which messages to include in the merkle root and ensures
// the newly built merkle roots are the same as RMN roots.
type Processor struct {
	oracleID     commontypes.OracleID
	cfg          pluginconfig.CommitPluginConfig
	lggr         logger.Logger
	observer     Observer
	ccipReader   readerpkg.CCIPReader
	reportingCfg ocr3types.ReportingPluginConfig
	chainSupport plugincommon.ChainSupport
	rmnClient    rmn.Client
	rmnNodeInfo  []rmn.RMNNodeInfo
}

// NewProcessor creates a new Processor
func NewProcessor(
	oracleID commontypes.OracleID,
	lggr logger.Logger,
	cfg pluginconfig.CommitPluginConfig,
	homeChain reader.HomeChain,
	ccipReader readerpkg.CCIPReader,
	msgHasher cciptypes.MessageHasher,
	reportingCfg ocr3types.ReportingPluginConfig,
	chainSupport plugincommon.ChainSupport,
	rmnNodeInfo []rmn.RMNNodeInfo,
) *Processor {
	observer := ObserverImpl{
		lggr,
		homeChain,
		oracleID,
		chainSupport,
		ccipReader,
		msgHasher,
	}
	return &Processor{
		oracleID:     oracleID,
		cfg:          cfg,
		lggr:         lggr,
		ccipReader:   ccipReader,
		observer:     observer,
		reportingCfg: reportingCfg,
		chainSupport: chainSupport,
		rmnNodeInfo:  rmnNodeInfo,
	}
}

var _ shared.PluginProcessor[Query, Observation, Outcome] = &Processor{}
