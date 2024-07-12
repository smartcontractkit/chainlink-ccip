package commitrmnocb

import (
	"fmt"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

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
	return &Plugin{
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
}

// TODO: doc
// SelectingRangesForReport doesn't depend on the previous outcome, explain how this is resilient (to being unable
// to parse previous outcome)
func (p *Plugin) decodeOutcome(outcome ocr3types.Outcome) (CommitPluginOutcome, CommitPluginState) {
	if outcome == nil || len(outcome) == 0 {
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
