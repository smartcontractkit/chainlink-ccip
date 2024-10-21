package merkleroot

import (
	"errors"
	"fmt"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
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
	observer := ObserverImpl{
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
		lggr:            lggr,
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

	errs := make([]error, 0)

	// close rmn controller
	if p.rmnController != nil {
		if err := p.rmnController.Close(); err != nil {
			errs = append(errs, fmt.Errorf("close RMN controller: %w", err))
			p.lggr.Errorw("Failed to close RMN controller", "err", err)
		}
	}

	// close rmn home reader
	if p.rmnHomeReader != nil {
		if err := p.rmnHomeReader.Close(); err != nil {
			errs = append(errs, fmt.Errorf("close RMNHome reader: %w", err))
			p.lggr.Errorw("Failed to close RMNHome reader", "err", err)
		}
	}

	return errors.Join(errs...)
}
