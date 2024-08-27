package commit

import (
	"context"
	"sort"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	reader_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func TestPlugin_E2E(t *testing.T) {
	destChain := ccipocr3.ChainSelector(123)
	sourceChain1 := ccipocr3.ChainSelector(1)
	sourceChain2 := ccipocr3.ChainSelector(2)

	ctx := context.Background()
	lggr := logger.Test(t)

	nodeIDs := []commontypes.OracleID{1, 2, 3}
	peerIDs := []libocrtypes.PeerID{{1}, {2}, {3}}

	oracleIDToP2pID := make(map[commontypes.OracleID]libocrtypes.PeerID)
	for i := range nodeIDs {
		oracleIDToP2pID[nodeIDs[i]] = peerIDs[i]
	}

	chainConfig := map[ccipocr3.ChainSelector]reader.ChainConfig{
		destChain: {
			FChain:         1,
			SupportedNodes: mapset.NewSet(peerIDs...),
			Config:         chainconfig.ChainConfig{FinalityDepth: 1},
		},
		sourceChain1: {
			FChain:         1,
			SupportedNodes: mapset.NewSet(peerIDs...),
			Config:         chainconfig.ChainConfig{FinalityDepth: 1},
		},
		sourceChain2: {
			FChain:         1,
			SupportedNodes: mapset.NewSet(peerIDs...),
			Config:         chainconfig.ChainConfig{FinalityDepth: 1},
		},
	}

	nextSeqNum := map[ccipocr3.ChainSelector]uint64{
		sourceChain1: 10,
		sourceChain2: 20,
	}

	cfg := pluginconfig.CommitPluginConfig{
		DestChain:                          destChain,
		NewMsgScanBatchSize:                0,
		MaxReportTransmissionCheckAttempts: 0,
		SyncTimeout:                        0,
		SyncFrequency:                      0,
		OffchainConfig:                     pluginconfig.CommitOffchainConfig{},
	}

	reportingCfg := ocr3types.ReportingPluginConfig{
		ConfigDigest:                            ocr2types.ConfigDigest{},
		OracleID:                                0,
		N:                                       0,
		F:                                       0,
		OnchainConfig:                           nil,
		OffchainConfig:                          nil,
		EstimatedRoundInterval:                  0,
		MaxDurationQuery:                        0,
		MaxDurationObservation:                  0,
		MaxDurationShouldAcceptAttestedReport:   0,
		MaxDurationShouldTransmitAcceptedReport: 0,
	}

	n0 := setupNode(ctx, t, lggr, nodeIDs[0], reportingCfg, oracleIDToP2pID, cfg, chainConfig, nextSeqNum)
	n1 := setupNode(ctx, t, lggr, nodeIDs[1], reportingCfg, oracleIDToP2pID, cfg, chainConfig, nextSeqNum)
	n2 := setupNode(ctx, t, lggr, nodeIDs[2], reportingCfg, oracleIDToP2pID, cfg, chainConfig, nextSeqNum)

	nodes := []ocr3types.ReportingPlugin[[]byte]{
		n0.node,
		n1.node,
		n2.node,
	}

	initialOutcome := ocr3types.Outcome{}

	runner := testhelpers.NewOCR3Runner(nodes, nodeIDs, initialOutcome)

	res, err := runner.RunRound(ctx)
	assert.NoError(t, err)
	t.Logf("%#v", res)
}

type nodeSetup struct {
	node        *Plugin
	ccipReader  *reader_mock.MockCCIP
	priceReader *mocks.TokenPricesReader
	reportCodec *mocks.CommitPluginJSONReportCodec
	msgHasher   *mocks.MessageHasher
}

func setupNode(
	ctx context.Context,
	t *testing.T,
	lggr logger.Logger,
	nodeID commontypes.OracleID,
	reportingCfg ocr3types.ReportingPluginConfig,
	oracleIDToP2pID map[commontypes.OracleID]libocrtypes.PeerID,
	pluginCfg pluginconfig.CommitPluginConfig,
	chainCfg map[ccipocr3.ChainSelector]reader.ChainConfig,
	nextSeqNum map[ccipocr3.ChainSelector]uint64,
) nodeSetup {
	ccipReader := reader_mock.NewMockCCIP(t)
	tokenPricesReader := mocks.NewTokenPricesReader()
	reportCodec := mocks.NewCommitPluginJSONReportCodec()
	msgHasher := mocks.NewMessageHasher()
	homeChainReader := reader_mock.NewMockHomeChain(t)

	fChain := map[ccipocr3.ChainSelector]int{}
	supportedChainsForPeer := make(map[libocrtypes.PeerID]mapset.Set[ccipocr3.ChainSelector])
	for chainSel, cfg := range chainCfg {
		fChain[chainSel] = cfg.FChain

		for _, peerID := range cfg.SupportedNodes.ToSlice() {
			if _, ok := supportedChainsForPeer[peerID]; !ok {
				supportedChainsForPeer[peerID] = mapset.NewSet[ccipocr3.ChainSelector]()
			}
			supportedChainsForPeer[peerID].Add(chainSel)
		}
	}

	homeChainReader.On("GetFChain").Return(fChain, nil)

	for peerID, supportedChains := range supportedChainsForPeer {
		homeChainReader.On("GetSupportedChainsForPeer", peerID).Return(supportedChains, nil).Maybe()
	}

	knownCCIPChains := mapset.NewSet[ccipocr3.ChainSelector]()

	for chainSel, cfg := range chainCfg {
		homeChainReader.On("GetChainConfig", chainSel).Return(cfg, nil).Maybe()
		knownCCIPChains.Add(chainSel)
	}
	homeChainReader.On("GetKnownCCIPChains").Return(knownCCIPChains, nil)

	sourceChains := make([]ccipocr3.ChainSelector, 0)
	seqNums := make([]ccipocr3.SeqNum, 0)
	for chainSel, ns := range nextSeqNum {
		sourceChains = append(sourceChains, chainSel)
		seqNums = append(seqNums, ccipocr3.SeqNum(ns))
	}

	ccipReader.On("NextSeqNum", ctx, mock.Anything).Run(func(args mock.Arguments) {
		providedChains := args[1].([]ccipocr3.ChainSelector)
		sort.Slice(providedChains, func(i, j int) bool { return providedChains[i] < providedChains[j] })
		sort.Slice(sourceChains, func(i, j int) bool { return sourceChains[i] < sourceChains[j] })
		assert.Equal(t, sourceChains, providedChains)
	}).Return(seqNums, nil)

	p := NewPlugin(
		ctx,
		nodeID,
		oracleIDToP2pID,
		pluginCfg,
		ccipReader,
		tokenPricesReader,
		reportCodec,
		msgHasher,
		lggr,
		homeChainReader,
		reportingCfg,
	)

	return nodeSetup{
		node:        p,
		ccipReader:  ccipReader,
		priceReader: tokenPricesReader,
		reportCodec: reportCodec,
		msgHasher:   msgHasher,
	}
}
