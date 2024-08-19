package commitrmnocb

import (
	"context"
	"fmt"
	"time"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type Plugin struct {
	nodeID            commontypes.OracleID
	oracleIDToP2pID   map[commontypes.OracleID]libocrtypes.PeerID
	destChain         cciptypes.ChainSelector
	offchainConfig    pluginconfig.CommitOffchainConfig
	ccipReader        reader.CCIP
	readerSyncer      *plugincommon.BackgroundReaderSyncer
	tokenPricesReader reader.TokenPrices
	reportCodec       cciptypes.CommitPluginCodec
	lggr              logger.Logger
	homeChain         reader.HomeChain
	reportingCfg      ocr3types.ReportingPluginConfig
	chainSupport      ChainSupport
	observer          Observer
}

func NewPlugin(
	_ context.Context,
	nodeID commontypes.OracleID,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	destChain cciptypes.ChainSelector,
	offchainConfig pluginconfig.CommitOffchainConfig,
	ccipReader reader.CCIP,
	tokenPricesReader reader.TokenPrices,
	reportCodec cciptypes.CommitPluginCodec,
	msgHasher cciptypes.MessageHasher,
	lggr logger.Logger,
	homeChain reader.HomeChain,
	reportingCfg ocr3types.ReportingPluginConfig,
) *Plugin {
	readerSyncer := plugincommon.NewBackgroundReaderSyncer(
		lggr,
		ccipReader,
		syncTimeout(offchainConfig.SyncTimeout),
		syncFrequency(offchainConfig.SyncFrequency),
	)
	if err := readerSyncer.Start(context.Background()); err != nil {
		lggr.Errorw("error starting background reader syncer", "err", err)
	}

	chainSupport := CCIPChainSupport{
		lggr:            lggr,
		homeChain:       homeChain,
		oracleIDToP2pID: oracleIDToP2pID,
		nodeID:          nodeID,
		destChain:       destChain,
	}

	observer := ObserverImpl{
		lggr:         lggr,
		homeChain:    homeChain,
		nodeID:       nodeID,
		chainSupport: chainSupport,
		ccipReader:   ccipReader,
		msgHasher:    msgHasher,
	}

	return &Plugin{
		nodeID:            nodeID,
		oracleIDToP2pID:   oracleIDToP2pID,
		lggr:              lggr,
		offchainConfig:    offchainConfig,
		destChain:         destChain,
		tokenPricesReader: tokenPricesReader,
		ccipReader:        ccipReader,
		homeChain:         homeChain,
		readerSyncer:      readerSyncer,
		reportCodec:       reportCodec,
		reportingCfg:      reportingCfg,
		chainSupport:      chainSupport,
		observer:          observer,
	}
}

func (p *Plugin) Close() error {
	timeout := 10 * time.Second
	ctx, cf := context.WithTimeout(context.Background(), timeout)
	defer cf()

	if err := p.readerSyncer.Close(); err != nil {
		p.lggr.Errorw("error closing reader syncer", "err", err)
	}

	if err := p.ccipReader.Close(ctx); err != nil {
		return fmt.Errorf("close ccip reader: %w", err)
	}

	return nil
}

func (p *Plugin) decodeOutcome(outcome ocr3types.Outcome) (Outcome, State) {
	if len(outcome) == 0 {
		return Outcome{}, SelectingRangesForReport
	}

	decodedOutcome, err := DecodeOutcome(outcome)
	if err != nil {
		p.lggr.Errorw("Failed to decode Outcome", "outcome", outcome, "err", err)
		return Outcome{}, SelectingRangesForReport
	}

	return decodedOutcome, decodedOutcome.NextState()
}

func syncFrequency(configuredValue time.Duration) time.Duration {
	if configuredValue.Milliseconds() == 0 {
		return 10 * time.Second
	}
	return configuredValue
}

func syncTimeout(configuredValue time.Duration) time.Duration {
	if configuredValue.Milliseconds() == 0 {
		return 3 * time.Second
	}
	return configuredValue
}

// Interface compatibility checks.
var _ ocr3types.ReportingPlugin[[]byte] = &Plugin{}
