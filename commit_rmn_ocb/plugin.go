package commitrmnocb

import (
	"context"
	"fmt"
	"sync"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type Plugin struct {
	reportingCfg      ocr3types.ReportingPluginConfig
	nodeID            commontypes.OracleID
	oracleIDToP2pID   map[commontypes.OracleID]libocrtypes.PeerID
	log               logger.Logger
	rmn               Rmn
	cfg               CommitPluginConfig
	onChain           OnChain
	tokenPricesReader reader.TokenPrices
	ccipReader        reader.CCIP
	homeChain         reader.HomeChain
	bgSyncCancelFunc  context.CancelFunc
	bgSyncWG          *sync.WaitGroup
}

func NewPlugin(
	reportingCfg ocr3types.ReportingPluginConfig,
	nodeID commontypes.OracleID,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	log logger.Logger,
	rmn Rmn,
	cfg CommitPluginConfig,
	onChain OnChain,
	tokenPricesReader reader.TokenPrices,
	ccipReader reader.CCIP,
	homeChain reader.HomeChain,
) *Plugin {
	p := &Plugin{
		reportingCfg:      reportingCfg,
		nodeID:            nodeID,
		oracleIDToP2pID:   oracleIDToP2pID,
		log:               log,
		rmn:               rmn,
		cfg:               cfg,
		onChain:           onChain,
		tokenPricesReader: tokenPricesReader,
		ccipReader:        ccipReader,
		homeChain:         homeChain,
	}

	bgSyncCtx, bgSyncCf := context.WithCancel(context.Background())
	p.bgSyncCancelFunc = bgSyncCf
	p.bgSyncWG = &sync.WaitGroup{}
	p.bgSyncWG.Add(1)
	commit.BackgroundReaderSync(
		bgSyncCtx,
		p.bgSyncWG,
		log,
		ccipReader,
		syncTimeout(cfg.SyncTimeout),
		time.NewTicker(syncFrequency(p.cfg.SyncFrequency)).C,
	)

	return p
}

func (p *Plugin) Close() error {
	timeout := 10 * time.Second
	ctx, cf := context.WithTimeout(context.Background(), timeout)
	defer cf()

	if err := p.ccipReader.Close(ctx); err != nil {
		return fmt.Errorf("close ccip reader: %w", err)
	}

	p.bgSyncCancelFunc()
	p.bgSyncWG.Wait()
	return nil
}

// TODO: doc
// SelectingRangesForReport doesn't depend on the previous outcome, explain how this is resilient (to being unable
// to parse previous outcome)
func (p *Plugin) decodeOutcome(outcome ocr3types.Outcome) (CommitPluginOutcome, CommitPluginState) {
	if len(outcome) == 0 {
		return CommitPluginOutcome{}, SelectingRangesForReport
	}

	decodedOutcome, err := DecodeCommitPluginOutcome(outcome)
	if err != nil {
		p.log.Errorw("Failed to decode CommitPluginOutcome", "outcome", outcome, "err", err)
		return CommitPluginOutcome{}, SelectingRangesForReport
	}

	return decodedOutcome, decodedOutcome.NextState()
}

// TODO: doc
// is this all source chains? across all nodes?
func (p *Plugin) sourceChains() []cciptypes.ChainSelector {
	return []cciptypes.ChainSelector{}
}

// TODO: doc
func (p *Plugin) supportedChains(oracleID commontypes.OracleID) (mapset.Set[cciptypes.ChainSelector], error) {
	p2pID, exists := p.oracleIDToP2pID[oracleID]
	if !exists {
		return nil, fmt.Errorf("oracle ID %d not found in oracleIDToP2pID", p.nodeID)
	}
	supportedChains, err := p.homeChain.GetSupportedChainsForPeer(p2pID)
	if err != nil {
		p.log.Warnw("error getting supported chains", err)
		return mapset.NewSet[cciptypes.ChainSelector](), fmt.Errorf("error getting supported chains: %w", err)
	}

	return supportedChains, nil
}

// supportsDestChain Returns true if the given oracle supports the dest chain, returns false otherwise
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
