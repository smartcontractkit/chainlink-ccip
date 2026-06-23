package merkleroot

import (
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

// Processor is the processor responsible for
// reading next messages and building merkle roots for them.
type Processor struct {
	oracleID        commontypes.OracleID
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID
	offchainCfg     pluginconfig.CommitOffchainConfig
	destChain       cciptypes.ChainSelector
	// Don't use this logger directly but rather through logutil\.WithContextValues where possible
	lggr            logger.Logger
	observer        Observer
	ccipReader      readerpkg.CCIPReader
	reportingCfg    ocr3types.ReportingPluginConfig
	chainSupport    plugincommon.ChainSupport
	metricsReporter MetricsReporter
	addressCodec    cciptypes.AddressCodec
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
	metricsReporter MetricsReporter,
	addressCodec cciptypes.AddressCodec,
) plugincommon.PluginProcessor[Query, Observation, Outcome] {
	var observer Observer
	baseObserver := newObserverImpl(
		lggr,
		homeChain,
		oracleID,
		chainSupport,
		ccipReader,
		msgHasher,
	)
	if !offchainCfg.MerkleRootAsyncObserverDisabled {
		observer = newAsyncObserver(
			lggr,
			baseObserver,
			offchainCfg.MerkleRootAsyncObserverSyncFreq,
			offchainCfg.MerkleRootAsyncObserverSyncTimeout,
		)
	} else {
		observer = baseObserver
	}

	p := &Processor{
		oracleID:        oracleID,
		oracleIDToP2pID: oracleIDToP2pID,
		offchainCfg:     offchainCfg,
		destChain:       destChain,
		lggr:            lggr,
		observer:        observer,
		ccipReader:      ccipReader,
		reportingCfg:    reportingCfg,
		chainSupport:    chainSupport,
		metricsReporter: metricsReporter,
		addressCodec:    addressCodec,
	}
	return plugincommon.NewTrackedProcessor(lggr, p, processorLabel, metricsReporter)
}

var _ plugincommon.PluginProcessor[Query, Observation, Outcome] = &Processor{}

func (p *Processor) Close() error {
	return p.observer.Close()
}
