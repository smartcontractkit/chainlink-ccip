package merkleroot

import (
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/services"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logger"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

// Processor is the processor responsible for
// reading next messages and building merkle roots for them,
// It's setup to use RMN to query which messages to include in the merkle root and ensures
// the newly built merkle roots are the same as RMN roots.
type Processor struct {
	oracleID               commontypes.OracleID
	oracleIDToP2pID        map[commontypes.OracleID]libocrtypes.PeerID
	offchainCfg            pluginconfig.CommitOffchainConfig
	destChain              cciptypes.ChainSelector
	lggr                   logger.Logger
	observer               Observer
	ccipReader             readerpkg.CCIPReader
	reportingCfg           ocr3types.ReportingPluginConfig
	chainSupport           plugincommon.ChainSupport
	rmnController          rmn.Controller
	rmnControllerCfgDigest cciptypes.Bytes32
	rmnCrypto              cciptypes.RMNCrypto
	rmnHomeReader          readerpkg.RMNHome
}

// NewProcessor creates a new Processor
func NewProcessor(
	oracleID commontypes.OracleID,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	lggr logger.Logger,
	offchainCfg pluginconfig.CommitOffchainConfig,
	destChain cciptypes.ChainSelector,
	homeChain reader.HomeChain,
	ccipReader readerpkg.CCIPReader,
	msgHasher cciptypes.MessageHasher,
	reportingCfg ocr3types.ReportingPluginConfig,
	chainSupport plugincommon.ChainSupport,
	rmnController rmn.Controller,
	rmnCrypto cciptypes.RMNCrypto,
	rmnHomeReader readerpkg.RMNHome,
) *Processor {
	observer := observerImpl{
		lggr,
		homeChain,
		oracleID,
		chainSupport,
		ccipReader,
		msgHasher,
	}
	return &Processor{
		oracleID:        oracleID,
		oracleIDToP2pID: oracleIDToP2pID,
		offchainCfg:     offchainCfg,
		destChain:       destChain,
		lggr:            logger.NewProcessorLogWrapper(lggr, "MerkleRoot"),
		observer:        observer,
		ccipReader:      ccipReader,
		reportingCfg:    reportingCfg,
		chainSupport:    chainSupport,
		rmnController:   rmnController,
		rmnCrypto:       rmnCrypto,
		rmnHomeReader:   rmnHomeReader,
	}
}

var _ plugincommon.PluginProcessor[Query, Observation, Outcome] = &Processor{}

func (p *Processor) Close() error {
	if !p.offchainCfg.RMNEnabled {
		return nil
	}

	return services.CloseAll(p.rmnController, p.rmnHomeReader)
}
