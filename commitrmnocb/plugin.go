package commitrmnocb

import (
	"context"
	"fmt"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type Plugin struct {
	nodeID            commontypes.OracleID
	oracleIDToP2pID   map[commontypes.OracleID]libocrtypes.PeerID
	cfg               pluginconfig.CommitPluginConfig
	ccipReader        reader.CCIP
	readerSyncer      *plugincommon.BackgroundReaderSyncer
	tokenPricesReader reader.TokenPrices
	reportCodec       cciptypes.CommitPluginCodec
	msgHasher         cciptypes.MessageHasher
	lggr              logger.Logger
	homeChain         reader.HomeChain
	reportingCfg      ocr3types.ReportingPluginConfig
}

func NewPlugin(
	_ context.Context,
	nodeID commontypes.OracleID,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	cfg pluginconfig.CommitPluginConfig,
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
		syncTimeout(cfg.SyncTimeout),
		syncFrequency(cfg.SyncFrequency),
	)
	if err := readerSyncer.Start(context.Background()); err != nil {
		lggr.Errorw("error starting background reader syncer", "err", err)
	}

	return &Plugin{
		nodeID:            nodeID,
		oracleIDToP2pID:   oracleIDToP2pID,
		lggr:              lggr,
		cfg:               cfg,
		tokenPricesReader: tokenPricesReader,
		ccipReader:        ccipReader,
		homeChain:         homeChain,
		readerSyncer:      readerSyncer,
		reportCodec:       reportCodec,
		msgHasher:         msgHasher,
		reportingCfg:      reportingCfg,
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

func (p *Plugin) decodeOutcome(outcome ocr3types.Outcome) (Outcome, CommitPluginState) {
	if len(outcome) == 0 {
		return Outcome{}, SelectingRangesForReport
	}

	decodedOutcome, err := DecodeCommitPluginOutcome(outcome)
	if err != nil {
		p.lggr.Errorw("Failed to decode Outcome", "outcome", outcome, "err", err)
		return Outcome{}, SelectingRangesForReport
	}

	return decodedOutcome, decodedOutcome.NextState()
}

func (p *Plugin) knownSourceChainsSlice() []cciptypes.ChainSelector {
	knownSourceChains, err := p.homeChain.GetKnownCCIPChains()
	if err != nil {
		p.lggr.Errorw("error getting known chains", "err", err)
		return nil
	}
	knownSourceChainsSlice := knownSourceChains.ToSlice()
	sort.Slice(
		knownSourceChainsSlice,
		func(i, j int) bool { return knownSourceChainsSlice[i] < knownSourceChainsSlice[j] },
	)
	return slicelib.Filter(knownSourceChainsSlice, func(ch cciptypes.ChainSelector) bool { return ch != p.cfg.DestChain })
}

// Returns the set of chains that the given Oracle is configured to access
func (p *Plugin) supportedChains(oracleID commontypes.OracleID) (mapset.Set[cciptypes.ChainSelector], error) {
	p2pID, exists := p.oracleIDToP2pID[oracleID]
	if !exists {
		return nil, fmt.Errorf("oracle ID %d not found in oracleIDToP2pID", p.nodeID)
	}
	supportedChains, err := p.homeChain.GetSupportedChainsForPeer(p2pID)
	if err != nil {
		p.lggr.Warnw("error getting supported chains", err)
		return mapset.NewSet[cciptypes.ChainSelector](), fmt.Errorf("error getting supported chains: %w", err)
	}

	return supportedChains, nil
}

// supportsDestChain returns true if the given oracle supports the dest chain, returns false otherwise
func (p *Plugin) supportsDestChain(oracle commontypes.OracleID) (bool, error) {
	destChainConfig, err := p.homeChain.GetChainConfig(p.cfg.DestChain)
	if err != nil {
		return false, fmt.Errorf("get chain config: %w", err)
	}
	return destChainConfig.SupportedNodes.Contains(p.oracleIDToP2pID[oracle]), nil
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
