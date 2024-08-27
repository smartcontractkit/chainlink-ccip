package commit

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sort"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	reader_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func TestPlugin_E2E(t *testing.T) {
	destChain := ccipocr3.ChainSelector(1)
	sourceChain1 := ccipocr3.ChainSelector(2)
	sourceChain2 := ccipocr3.ChainSelector(3)
	require.Len(t, mapset.NewSet(destChain, sourceChain1, sourceChain2).ToSlice(), 3)

	ctx := context.Background()
	lggr := logger.Test(t)

	oracleIDs := []commontypes.OracleID{1, 2, 3}
	peerIDs := []libocrtypes.PeerID{{1}, {2}, {3}}
	require.Equal(t, len(oracleIDs), len(peerIDs))

	oracleIDToPeerID := make(map[commontypes.OracleID]libocrtypes.PeerID)
	for i := range oracleIDs {
		oracleIDToPeerID[oracleIDs[i]] = peerIDs[i]
	}

	peerIDsMap := mapset.NewSet(peerIDs...)
	homeChainConfig := map[ccipocr3.ChainSelector]reader.ChainConfig{
		destChain:    {FChain: 1, SupportedNodes: peerIDsMap, Config: chainconfig.ChainConfig{FinalityDepth: 1}},
		sourceChain1: {FChain: 1, SupportedNodes: peerIDsMap, Config: chainconfig.ChainConfig{FinalityDepth: 1}},
		sourceChain2: {FChain: 1, SupportedNodes: peerIDsMap, Config: chainconfig.ChainConfig{FinalityDepth: 1}},
	}

	offRampNextSeqNum := map[ccipocr3.ChainSelector]ccipocr3.SeqNum{
		sourceChain1: 10,
		sourceChain2: 20,
	}

	onRampLastSeqNum := map[ccipocr3.ChainSelector]ccipocr3.SeqNum{
		sourceChain1: 10, // one new msg -> 10
		sourceChain2: 19, // no new msg, still on 19
	}

	cfg := pluginconfig.CommitPluginConfig{
		DestChain:                          destChain,
		NewMsgScanBatchSize:                100,
		MaxReportTransmissionCheckAttempts: 0,
		SyncTimeout:                        10 * time.Second,
		SyncFrequency:                      time.Hour,
	}

	nodes := make([]ocr3types.ReportingPlugin[[]byte], len(oracleIDs))
	reportingCfg := ocr3types.ReportingPluginConfig{F: 1}
	for i := range oracleIDs {
		n := setupNode(ctx, t, lggr, oracleIDs[i], reportingCfg, oracleIDToPeerID,
			cfg, homeChainConfig, offRampNextSeqNum, onRampLastSeqNum)
		nodes[i] = n.node
	}

	testCases := []struct {
		name        string
		prevOutcome Outcome
		expOutcome  Outcome
	}{
		{
			name:        "empty previous outcome, should select ranges for report",
			prevOutcome: Outcome{},
			expOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: sourceChain1, SeqNumRange: ccipocr3.SeqNumRange{10, 10}},
					{ChainSel: sourceChain2, SeqNumRange: ccipocr3.SeqNumRange{20, 20}},
				},
			},
		},
		{
			name: "selected ranges for report in previous outcome",
			prevOutcome: Outcome{
				OutcomeType: ReportIntervalsSelected,
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: sourceChain1, SeqNumRange: ccipocr3.SeqNumRange{10, 10}},
					{ChainSel: sourceChain2, SeqNumRange: ccipocr3.SeqNumRange{20, 20}},
				},
			},
			expOutcome: Outcome{
				OutcomeType: ReportGenerated,
				RootsToReport: []ccipocr3.MerkleRootChain{
					{
						ChainSel:     sourceChain1,
						SeqNumsRange: ccipocr3.SeqNumRange{0xa, 0xa},
						MerkleRoot: ccipocr3.Bytes32{0x4a, 0x44, 0xdc, 0x15, 0x36, 0x42, 0x4, 0xa8, 0xf, 0xe8, 0xe,
							0x90, 0x39, 0x45, 0x5c, 0xc1, 0x60, 0x82, 0x81, 0x82, 0xf, 0xe2, 0xb2, 0x4f, 0x1e, 0x52,
							0x33, 0xad, 0xe6, 0xaf, 0x1d, 0xd5},
					},
				},
				TokenPrices: make([]ccipocr3.TokenPrice, 0),
				GasPrices:   make([]ccipocr3.GasPriceChain, 0),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			encodedPrevOutcome, err := tc.prevOutcome.Encode()
			assert.NoError(t, err)
			runner := testhelpers.NewOCR3Runner(nodes, oracleIDs, encodedPrevOutcome)
			res, err := runner.RunRound(ctx)
			assert.NoError(t, err)

			decodedOutcome, err := DecodeOutcome(res.Outcome)
			assert.NoError(t, err)
			assert.Equal(t, tc.expOutcome, decodedOutcome)
		})
	}
}

type nodeSetup struct {
	node        *Plugin
	ccipReader  *reader_mock.MockCCIP
	priceReader *reader_mock.MockTokenPrices
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
	nextSeqNum map[ccipocr3.ChainSelector]ccipocr3.SeqNum,
	onRampLastSeqNum map[ccipocr3.ChainSelector]ccipocr3.SeqNum,
) nodeSetup {
	ccipReader := reader_mock.NewMockCCIP(t)
	tokenPricesReader := reader_mock.NewMockTokenPrices(t)
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
	homeChainReader.On("GetKnownCCIPChains").Return(knownCCIPChains, nil).Maybe()

	sourceChains := make([]ccipocr3.ChainSelector, 0)
	seqNums := make([]ccipocr3.SeqNum, 0)
	chainsWithNewMsgs := make([]ccipocr3.ChainSelector, 0)
	for chainSel, offRampNextSeqNum := range nextSeqNum {
		sourceChains = append(sourceChains, chainSel)
		seqNums = append(seqNums, offRampNextSeqNum)

		newMsgs := make([]ccipocr3.Message, 0)
		numNewMsgs := (onRampLastSeqNum[chainSel] - offRampNextSeqNum) + 1
		for i := uint64(0); i < uint64(numNewMsgs); i++ {
			messageID := sha256.Sum256([]byte(fmt.Sprintf("%d", uint64(offRampNextSeqNum)+i)))
			newMsgs = append(newMsgs, ccipocr3.Message{
				Header: ccipocr3.RampMessageHeader{
					MessageID:      messageID,
					SequenceNumber: offRampNextSeqNum + ccipocr3.SeqNum(i),
				},
			})
		}

		ccipReader.On("MsgsBetweenSeqNums", ctx, chainSel,
			ccipocr3.NewSeqNumRange(offRampNextSeqNum, offRampNextSeqNum)).
			Return(newMsgs, nil).Maybe()

		if len(newMsgs) > 0 {
			chainsWithNewMsgs = append(chainsWithNewMsgs, chainSel)
		}
	}

	sort.Slice(chainsWithNewMsgs, func(i, j int) bool { return chainsWithNewMsgs[i] < chainsWithNewMsgs[j] })
	seqNumsOfChainsWithNewMsgs := make([]ccipocr3.SeqNum, 0)
	for _, chainSel := range chainsWithNewMsgs {
		seqNumsOfChainsWithNewMsgs = append(seqNumsOfChainsWithNewMsgs, nextSeqNum[chainSel])
	}
	if len(chainsWithNewMsgs) > 0 {
		ccipReader.On("NextSeqNum", ctx, chainsWithNewMsgs).Return(seqNumsOfChainsWithNewMsgs, nil).Maybe()
	}
	sort.Slice(sourceChains, func(i, j int) bool { return sourceChains[i] < sourceChains[j] })
	ccipReader.On("NextSeqNum", ctx, sourceChains).Return(seqNums, nil).Maybe()

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
