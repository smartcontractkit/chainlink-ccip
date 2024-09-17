package commit

import (
	"context"
	"crypto/sha256"
	"fmt"
	"sort"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocr2types "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	reader_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/shared"

	"github.com/stretchr/testify/mock"
)

const (
	destChain    = ccipocr3.ChainSelector(1)
	sourceChain1 = ccipocr3.ChainSelector(2)
	sourceChain2 = ccipocr3.ChainSelector(3)
)

func TestPlugin_E2E_AllNodesAgree(t *testing.T) {
	ctx := tests.Context(t)
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
		destChain:    {FChain: 1, SupportedNodes: peerIDsMap, Config: chainconfig.ChainConfig{}},
		sourceChain1: {FChain: 1, SupportedNodes: peerIDsMap, Config: chainconfig.ChainConfig{}},
		sourceChain2: {FChain: 1, SupportedNodes: peerIDsMap, Config: chainconfig.ChainConfig{}},
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
		MaxReportTransmissionCheckAttempts: 2,
		SyncTimeout:                        10 * time.Second,
		SyncFrequency:                      time.Hour,
	}

	nodes := make([]ocr3types.ReportingPlugin[[]byte], len(oracleIDs))

	reportingCfg := ocr3types.ReportingPluginConfig{F: 1}

	outcomeIntervalsSelected := Outcome{
		MerkleRootOutcome: merkleroot.Outcome{
			OutcomeType: merkleroot.ReportIntervalsSelected,
			RangesSelectedForReport: []plugintypes.ChainRange{
				{ChainSel: sourceChain1, SeqNumRange: ccipocr3.SeqNumRange{10, 10}},
				{ChainSel: sourceChain2, SeqNumRange: ccipocr3.SeqNumRange{20, 20}},
			},
			OffRampNextSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: sourceChain1, SeqNum: 10},
				{ChainSel: sourceChain2, SeqNum: 20},
			},
		},
	}

	outcomeReportGenerated := Outcome{
		MerkleRootOutcome: merkleroot.Outcome{
			OutcomeType: merkleroot.ReportGenerated,
			RootsToReport: []ccipocr3.MerkleRootChain{
				{
					ChainSel:     sourceChain1,
					SeqNumsRange: ccipocr3.SeqNumRange{0xa, 0xa},
					MerkleRoot:   merkleRoot1,
				},
			},
			OffRampNextSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: sourceChain1, SeqNum: 10},
				{ChainSel: sourceChain2, SeqNum: 20},
			},
			RMNReportSignatures: []ccipocr3.RMNECDSASignature{},
		},
	}

	outcomeReportGeneratedOneInflightCheck := outcomeReportGenerated
	outcomeReportGeneratedOneInflightCheck.MerkleRootOutcome.ReportTransmissionCheckAttempts = 1

	testCases := []struct {
		name                  string
		prevOutcome           Outcome
		expOutcome            Outcome
		expTransmittedReports []ccipocr3.CommitPluginReport

		offRampNextSeqNumDefaultOverrideKeys   []ccipocr3.ChainSelector
		offRampNextSeqNumDefaultOverrideValues []ccipocr3.SeqNum
	}{
		{
			name:        "empty previous outcome, should select ranges for report",
			prevOutcome: Outcome{},
			expOutcome:  outcomeIntervalsSelected,
		},
		{
			name:        "selected ranges for report in previous outcome",
			prevOutcome: outcomeIntervalsSelected,
			expOutcome:  outcomeReportGenerated,
			expTransmittedReports: []ccipocr3.CommitPluginReport{
				{
					MerkleRoots: []ccipocr3.MerkleRootChain{
						{
							ChainSel:     sourceChain1,
							SeqNumsRange: ccipocr3.NewSeqNumRange(0xa, 0xa),
							MerkleRoot:   merkleRoot1,
						},
					},
					PriceUpdates:  ccipocr3.PriceUpdates{},
					RMNSignatures: []ccipocr3.RMNECDSASignature{},
				},
			},
		},
		{
			name:        "report generated in previous outcome, still inflight",
			prevOutcome: outcomeReportGenerated,
			expOutcome: Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType:                     merkleroot.ReportInFlight,
					ReportTransmissionCheckAttempts: 1,
					OffRampNextSeqNums: []plugintypes.SeqNumChain{
						{ChainSel: sourceChain1, SeqNum: 10},
						{ChainSel: sourceChain2, SeqNum: 20},
					},
				},
			},
		},
		{
			name:        "report generated in previous outcome, still inflight, reached all inflight check attempts",
			prevOutcome: outcomeReportGeneratedOneInflightCheck,
			expOutcome: Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportTransmissionFailed,
				},
			},
		},
		{
			name:                                   "report generated in previous outcome, transmitted with success",
			prevOutcome:                            outcomeReportGenerated,
			offRampNextSeqNumDefaultOverrideKeys:   []ccipocr3.ChainSelector{sourceChain1, sourceChain2},
			offRampNextSeqNumDefaultOverrideValues: []ccipocr3.SeqNum{11, 20},
			expOutcome: Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportTransmitted,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var reportCodec ccipocr3.CommitPluginCodec
			for i := range oracleIDs {
				n := setupNode(ctx, t, lggr, oracleIDs[i], reportingCfg, oracleIDToPeerID,
					cfg, homeChainConfig, offRampNextSeqNum, onRampLastSeqNum)
				nodes[i] = n.node
				if i == 0 {
					reportCodec = n.reportCodec
				}

				if len(tc.offRampNextSeqNumDefaultOverrideKeys) > 0 {
					assert.Equal(t, len(tc.offRampNextSeqNumDefaultOverrideKeys), len(tc.offRampNextSeqNumDefaultOverrideValues))
					n.ccipReader.EXPECT().NextSeqNum(ctx, tc.offRampNextSeqNumDefaultOverrideKeys).Unset()
					n.ccipReader.EXPECT().
						NextSeqNum(ctx, tc.offRampNextSeqNumDefaultOverrideKeys).
						Return(tc.offRampNextSeqNumDefaultOverrideValues, nil).
						Maybe()
				}
				n.priceReader.EXPECT().
					GetFeeQuoterTokenUpdates(ctx, mock.Anything).
					Return(
						map[ocr2types.Account]shared.TimestampedBig{}, nil,
					).
					Maybe()
			}

			encodedPrevOutcome, err := tc.prevOutcome.Encode()
			assert.NoError(t, err)
			runner := testhelpers.NewOCR3Runner(nodes, oracleIDs, encodedPrevOutcome)
			res, err := runner.RunRound(ctx)
			assert.NoError(t, err)

			decodedOutcome, err := DecodeOutcome(res.Outcome)
			assert.NoError(t, err)
			assert.Equal(t, tc.expOutcome, decodedOutcome)

			assert.Len(t, res.Transmitted, len(tc.expTransmittedReports))
			for i := range res.Transmitted {
				decoded, err := reportCodec.Decode(ctx, res.Transmitted[i].Report)
				assert.NoError(t, err)
				assert.Equal(t, tc.expTransmittedReports[i], decoded)
			}
		})
	}
}

type nodeSetup struct {
	node        *Plugin
	ccipReader  *readerpkg_mock.MockCCIPReader
	priceReader *reader_mock.MockPriceReader
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
	offRampNextSeqNum map[ccipocr3.ChainSelector]ccipocr3.SeqNum,
	onRampLastSeqNum map[ccipocr3.ChainSelector]ccipocr3.SeqNum,
) nodeSetup {
	ccipReader := readerpkg_mock.NewMockCCIPReader(t)
	tokenPricesReader := reader_mock.NewMockPriceReader(t)
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

	homeChainReader.EXPECT().GetFChain().Return(fChain, nil)

	for peerID, supportedChains := range supportedChainsForPeer {
		homeChainReader.EXPECT().GetSupportedChainsForPeer(peerID).Return(supportedChains, nil).Maybe()
	}

	knownCCIPChains := mapset.NewSet[ccipocr3.ChainSelector]()

	for chainSel, cfg := range chainCfg {
		homeChainReader.EXPECT().GetChainConfig(chainSel).Return(cfg, nil).Maybe()
		knownCCIPChains.Add(chainSel)
	}
	homeChainReader.EXPECT().GetKnownCCIPChains().Return(knownCCIPChains, nil).Maybe()

	sourceChains := make([]ccipocr3.ChainSelector, 0)
	for chainSel := range offRampNextSeqNum {
		sourceChains = append(sourceChains, chainSel)
	}
	sort.Slice(sourceChains, func(i, j int) bool { return sourceChains[i] < sourceChains[j] })

	offRampNextSeqNums := make([]ccipocr3.SeqNum, 0)
	chainsWithNewMsgs := make([]ccipocr3.ChainSelector, 0)
	for _, sourceChain := range sourceChains {
		offRampNextSeqNum, ok := offRampNextSeqNum[sourceChain]
		assert.True(t, ok)
		offRampNextSeqNums = append(offRampNextSeqNums, offRampNextSeqNum)

		newMsgs := make([]ccipocr3.Message, 0)
		numNewMsgs := (onRampLastSeqNum[sourceChain] - offRampNextSeqNum) + 1
		for i := uint64(0); i < uint64(numNewMsgs); i++ {
			messageID := sha256.Sum256([]byte(fmt.Sprintf("%d", uint64(offRampNextSeqNum)+i)))
			newMsgs = append(newMsgs, ccipocr3.Message{
				Header: ccipocr3.RampMessageHeader{
					MessageID:      messageID,
					SequenceNumber: offRampNextSeqNum + ccipocr3.SeqNum(i),
				},
			})
		}

		ccipReader.EXPECT().MsgsBetweenSeqNums(ctx, sourceChain,
			ccipocr3.NewSeqNumRange(offRampNextSeqNum, offRampNextSeqNum)).
			Return(newMsgs, nil).Maybe()

		if len(newMsgs) > 0 {
			chainsWithNewMsgs = append(chainsWithNewMsgs, sourceChain)
		}
	}

	seqNumsOfChainsWithNewMsgs := make([]ccipocr3.SeqNum, 0)
	for _, chainSel := range chainsWithNewMsgs {
		seqNumsOfChainsWithNewMsgs = append(seqNumsOfChainsWithNewMsgs, offRampNextSeqNum[chainSel])
	}
	if len(chainsWithNewMsgs) > 0 {
		ccipReader.EXPECT().NextSeqNum(ctx, chainsWithNewMsgs).Return(seqNumsOfChainsWithNewMsgs, nil).Maybe()
	}
	ccipReader.EXPECT().NextSeqNum(ctx, sourceChains).Return(offRampNextSeqNums, nil).Maybe()

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
		rmntypes.RMNConfig{},
	)

	return nodeSetup{
		node:        p,
		ccipReader:  ccipReader,
		priceReader: tokenPricesReader,
		reportCodec: reportCodec,
		msgHasher:   msgHasher,
	}
}

// merkleRoot1 is the markle root that the test generates, the merkle root generation logic is not supposed to be
// tested in this context, so we just assume it's correct.
var merkleRoot1 = ccipocr3.Bytes32{0x4a, 0x44, 0xdc, 0x15, 0x36, 0x42, 0x4, 0xa8, 0xf, 0xe8, 0xe,
	0x90, 0x39, 0x45, 0x5c, 0xc1, 0x60, 0x82, 0x81, 0x82, 0xf, 0xe2, 0xb2, 0x4f, 0x1e, 0x52,
	0x33, 0xad, 0xe6, 0xaf, 0x1d, 0xd5}
