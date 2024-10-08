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
)

// Processor is the processor responsible for
// reading next messages and building merkle roots for them,
// It's setup to use RMN to query which messages to include in the merkle root and ensures
// the newly built merkle roots are the same as RMN roots.
type Processor struct {
	oracleID      commontypes.OracleID
	offchainCfg   pluginconfig.CommitOffchainConfig
	destChain     cciptypes.ChainSelector
	lggr          logger.Logger
	observer      Observer
	ccipReader    readerpkg.CCIPReader
	reportingCfg  ocr3types.ReportingPluginConfig
	chainSupport  plugincommon.ChainSupport
	rmnClient     rmn.Controller
	rmnCrypto     cciptypes.RMNCrypto
	rmnHomeReader reader.RMNHome
}

// NewProcessor creates a new Processor
func NewProcessor(
	oracleID commontypes.OracleID,
	lggr logger.Logger,
	offchainCfg pluginconfig.CommitOffchainConfig,
	destChain cciptypes.ChainSelector,
	homeChain reader.HomeChain,
	ccipReader readerpkg.CCIPReader,
	msgHasher cciptypes.MessageHasher,
	reportingCfg ocr3types.ReportingPluginConfig,
	chainSupport plugincommon.ChainSupport,
	rmnClient rmn.Controller,
	rmnCrypto cciptypes.RMNCrypto,
	rmnHomeReader reader.RMNHome,
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
		oracleID:      oracleID,
		offchainCfg:   offchainCfg,
		destChain:     destChain,
		lggr:          lggr,
		observer:      observer,
		ccipReader:    ccipReader,
		reportingCfg:  reportingCfg,
		chainSupport:  chainSupport,
		rmnClient:     rmnClient,
		rmnCrypto:     rmnCrypto,
		rmnHomeReader: rmnHomeReader,
	}
}

var _ plugincommon.PluginProcessor[Query, Observation, Outcome] = &Processor{}
