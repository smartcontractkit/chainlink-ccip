package commit

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"sort"
	"testing"
	"time"

	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/metrics"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/mathslib"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"

	"golang.org/x/exp/maps"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	reader_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	reader2 "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

const (
	destChain         = ccipocr3.ChainSelector(1)
	sourceChain1      = ccipocr3.ChainSelector(2)
	sourceChain2      = ccipocr3.ChainSelector(3)
	arbAddr           = ccipocr3.UnknownEncodedAddress("0xa100000000000000000000000000000000000000")
	arbAggregatorAddr = ccipocr3.UnknownEncodedAddress("0xa2000000000000000000000000000000000000000")

	ethAddr           = ccipocr3.UnknownEncodedAddress("0xe100000000000000000000000000000000000000")
	ethAggregatorAddr = ccipocr3.UnknownEncodedAddress("0xe200000000000000000000000000000000000000")
)

var (
	oracleIDs = []commontypes.OracleID{1, 2, 3}
	peerIDs   = []libocrtypes.PeerID{{1}, {2}, {3}}

	arbPrice = new(big.Int).Mul(big.NewInt(5), big.NewInt(1e18))
	ethPrice = new(big.Int).Mul(big.NewInt(7), big.NewInt(1e18))

	// a map to ease working with tests
	tokenPriceMap = map[ccipocr3.UnknownEncodedAddress]ccipocr3.TokenPrice{
		arbAddr: {
			TokenID: arbAddr,
			Price:   ccipocr3.NewBigInt(arbPrice),
		},
		ethAddr: {
			TokenID: ethAddr,
			Price:   ccipocr3.NewBigInt(ethPrice),
		},
	}

	decimals18 = uint8(18)

	arbInfo = pluginconfig.TokenInfo{
		AggregatorAddress: arbAggregatorAddr,
		DeviationPPB:      ccipocr3.NewBigInt(big.NewInt(1e5)),
		Decimals:          decimals18,
	}
	ethInfo = pluginconfig.TokenInfo{
		AggregatorAddress: ethAggregatorAddr,
		DeviationPPB:      ccipocr3.NewBigInt(big.NewInt(1e5)),
		Decimals:          decimals18,
	}
)

func TestPlugin_E2E_AllNodesAgree_MerkleRoots(t *testing.T) {
	params := defaultNodeParams(t)
	nodes := make([]ocr3types.ReportingPlugin[[]byte], len(oracleIDs))

	outcomeIntervalsSelected := committypes.Outcome{
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
			RMNRemoteCfg: params.rmnReportCfg,
		},
	}

	outcomeReportGenerated := committypes.Outcome{
		MerkleRootOutcome: merkleroot.Outcome{
			OutcomeType: merkleroot.ReportGenerated,
			RootsToReport: []ccipocr3.MerkleRootChain{
				{
					ChainSel:      sourceChain1,
					OnRampAddress: ccipocr3.UnknownAddress{},
					SeqNumsRange:  ccipocr3.SeqNumRange{0xa, 0xa},
					MerkleRoot:    merkleRoot1,
				},
			},
			OffRampNextSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: sourceChain1, SeqNum: 10},
				{ChainSel: sourceChain2, SeqNum: 20},
			},
			RMNReportSignatures: []ccipocr3.RMNECDSASignature{},
			RMNRemoteCfg:        params.rmnReportCfg,
		},
	}

	outcomeReportGeneratedOneInflightCheck := outcomeReportGenerated
	outcomeReportGeneratedOneInflightCheck.MerkleRootOutcome.ReportTransmissionCheckAttempts = 1

	testCases := []struct {
		name                  string
		prevOutcome           committypes.Outcome
		expOutcome            committypes.Outcome
		expTransmittedReports []ccipocr3.CommitPluginReport

		offRampNextSeqNumDefaultOverrideKeys   []ccipocr3.ChainSelector
		offRampNextSeqNumDefaultOverrideValues map[ccipocr3.ChainSelector]ccipocr3.SeqNum

		enableDiscovery bool
	}{
		{
			name:        "empty previous outcome, should select ranges for report",
			prevOutcome: committypes.Outcome{},
			expOutcome:  outcomeIntervalsSelected,
		},
		{
			name:            "discovery enabled, should discover contracts",
			prevOutcome:     committypes.Outcome{},
			expOutcome:      committypes.Outcome{},
			enableDiscovery: true,
		},
		{
			name:        "selected ranges for report in previous outcome",
			prevOutcome: outcomeIntervalsSelected,
			expOutcome:  outcomeReportGenerated,
			expTransmittedReports: []ccipocr3.CommitPluginReport{
				{
					MerkleRoots: []ccipocr3.MerkleRootChain{
						{
							ChainSel:      sourceChain1,
							SeqNumsRange:  ccipocr3.NewSeqNumRange(0xa, 0xa),
							OnRampAddress: ccipocr3.UnknownAddress{},
							MerkleRoot:    merkleRoot1,
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
			expOutcome: committypes.Outcome{
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
			expOutcome: committypes.Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportTransmissionFailed,
				},
			},
		},
		{
			name:                                 "report generated in previous outcome, transmitted with success",
			prevOutcome:                          outcomeReportGenerated,
			offRampNextSeqNumDefaultOverrideKeys: []ccipocr3.ChainSelector{sourceChain1, sourceChain2},
			offRampNextSeqNumDefaultOverrideValues: map[ccipocr3.ChainSelector]ccipocr3.SeqNum{
				sourceChain1: 11,
				sourceChain2: 20,
			},
			expOutcome: committypes.Outcome{
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
				paramsCp := params
				paramsCp.enableDiscovery = tc.enableDiscovery
				paramsCp.reportingCfg.OracleID = oracleIDs[i]
				n := setupNode(paramsCp)
				nodes[i] = n.node
				if i == 0 {
					reportCodec = n.reportCodec
				}
				prepareCcipReaderMock(params.ctx, n.ccipReader, false, tc.enableDiscovery, true)

				if len(tc.offRampNextSeqNumDefaultOverrideKeys) > 0 {
					require.Equal(t, len(tc.offRampNextSeqNumDefaultOverrideKeys), len(tc.offRampNextSeqNumDefaultOverrideValues))
					n.ccipReader.EXPECT().NextSeqNum(params.ctx, tc.offRampNextSeqNumDefaultOverrideKeys).Unset()
					n.ccipReader.EXPECT().
						NextSeqNum(params.ctx, tc.offRampNextSeqNumDefaultOverrideKeys).
						Return(tc.offRampNextSeqNumDefaultOverrideValues, nil).
						Maybe()
				}

				preparePriceReaderMock(params.ctx, n.priceReader)
			}

			encodedPrevOutcome, err := tc.prevOutcome.Encode()
			assert.NoError(t, err)
			runner := testhelpers.NewOCR3Runner(nodes, oracleIDs, encodedPrevOutcome)
			res, err := runner.RunRound(params.ctx)
			assert.NoError(t, err)

			decodedOutcome, err := committypes.DecodeOutcome(res.Outcome)
			assert.NoError(t, err)
			assert.Equal(t, normalizeOutcome(tc.expOutcome), normalizeOutcome(decodedOutcome))

			assert.Len(t, res.Transmitted, len(tc.expTransmittedReports))
			for i := range res.Transmitted {
				decoded, err := reportCodec.Decode(params.ctx, res.Transmitted[i].Report)
				assert.NoError(t, err)
				assert.Equal(t, tc.expTransmittedReports[i], decoded)
			}
		})
	}
}

func TestPlugin_E2E_AllNodesAgree_TokenPrices(t *testing.T) {
	params := defaultNodeParams(t)

	tokensToQuery := maps.Keys(params.offchainCfg.TokenInfo)
	sort.Slice(tokensToQuery, func(i, j int) bool { return tokensToQuery[i] < tokensToQuery[j] })
	orderedTokenPrices := make([]ccipocr3.TokenPrice, 0, len(tokensToQuery))
	for _, token := range tokensToQuery {
		orderedTokenPrices = append(orderedTokenPrices, tokenPriceMap[token])
	}

	nodes := make([]ocr3types.ReportingPlugin[[]byte], len(oracleIDs))

	merkleOutcome := baseMerkleOutcome(params.rmnReportCfg)

	testCases := []struct {
		name                  string
		prevOutcome           committypes.Outcome
		expOutcome            committypes.Outcome
		mockPriceReader       func(*readerpkg_mock.MockPriceReader)
		expTransmittedReports []ccipocr3.CommitPluginReport
		enableDiscovery       bool
	}{
		{
			name:        "empty fee_quoter token updates, should select all token prices for update",
			prevOutcome: committypes.Outcome{},
			mockPriceReader: func(m *readerpkg_mock.MockPriceReader) {
				m.EXPECT().
					// tokens need to be ordered, plugin checks all tokens from commit offchain config
					GetFeedPricesUSD(params.ctx, []ccipocr3.UnknownEncodedAddress{arbAddr, ethAddr}).
					Return([]*big.Int{arbPrice, ethPrice}, nil).
					Maybe()

				m.EXPECT().
					GetFeeQuoterTokenUpdates(params.ctx, mock.Anything, mock.Anything).
					Return(
						map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig{}, nil,
					).
					Maybe()
			},
			expOutcome: committypes.Outcome{
				MerkleRootOutcome: merkleOutcome,
				TokenPriceOutcome: tokenprice.Outcome{
					TokenPrices: orderedTokenPrices,
				},
				MainOutcome: committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 10},
			},
			expTransmittedReports: []ccipocr3.CommitPluginReport{
				{
					PriceUpdates: ccipocr3.PriceUpdates{
						TokenPriceUpdates: []ccipocr3.TokenPrice{
							{
								TokenID: arbAddr,
								Price:   ccipocr3.NewBigInt(arbPrice),
							},
							{
								TokenID: ethAddr,
								Price:   ccipocr3.NewBigInt(ethPrice),
							},
						},
					},
				},
			},
		},
		{
			name: "prices already inflight, no prices to report",
			prevOutcome: committypes.Outcome{
				MainOutcome: committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 4},
			},
			mockPriceReader: func(m *readerpkg_mock.MockPriceReader) {},
			expOutcome: committypes.Outcome{
				MerkleRootOutcome: merkleOutcome,
				TokenPriceOutcome: tokenprice.Outcome{},
				MainOutcome:       committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 3},
			},
			expTransmittedReports: []ccipocr3.CommitPluginReport{},
		},
		{
			name:        "fresh tokens don't need new updates",
			prevOutcome: committypes.Outcome{},
			mockPriceReader: func(m *readerpkg_mock.MockPriceReader) {
				m.EXPECT().
					// tokens need to be ordered, plugin checks all tokens from commit offchain config
					GetFeedPricesUSD(params.ctx, []ccipocr3.UnknownEncodedAddress{arbAddr, ethAddr}).
					Return([]*big.Int{arbPrice, ethPrice}, nil).
					Maybe()

				// Arb is fresh, will not be updated
				m.EXPECT().
					GetFeeQuoterTokenUpdates(params.ctx, mock.Anything, mock.Anything).
					Return(
						map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig{
							arbAddr: {Value: ccipocr3.NewBigInt(arbPrice), Timestamp: time.Now()},
							ethAddr: {Value: ccipocr3.NewBigInt(ethPrice), Timestamp: time.Now()},
						}, nil,
					).
					Maybe()
			},
			expOutcome: committypes.Outcome{
				MerkleRootOutcome: merkleOutcome,
				TokenPriceOutcome: tokenprice.Outcome{},
			},
		},
		{
			name:        "stale tokens need new updates",
			prevOutcome: committypes.Outcome{},
			mockPriceReader: func(m *readerpkg_mock.MockPriceReader) {
				m.EXPECT().
					// tokens need to be ordered, plugin checks all tokens from commit offchain config
					GetFeedPricesUSD(params.ctx, []ccipocr3.UnknownEncodedAddress{arbAddr, ethAddr}).
					Return([]*big.Int{arbPrice, ethPrice}, nil).
					Maybe()

				m.EXPECT().
					GetFeeQuoterTokenUpdates(params.ctx, mock.Anything, mock.Anything).
					Return(
						map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig{
							// Arb is fresh, will not be updated
							arbAddr: {Value: ccipocr3.NewBigInt(arbPrice), Timestamp: time.Now()},
							// Eth is stale, should update
							ethAddr: {
								Value: ccipocr3.NewBigInt(ethPrice),
								Timestamp: time.Now().
									Add(-params.offchainCfg.TokenPriceBatchWriteFrequency.Duration() * 2),
							},
						}, nil,
					).
					Maybe()
			},
			expOutcome: committypes.Outcome{
				MerkleRootOutcome: merkleOutcome,
				TokenPriceOutcome: tokenprice.Outcome{
					TokenPrices: []ccipocr3.TokenPrice{
						tokenPriceMap[ethAddr],
					},
				},
				MainOutcome: committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 10},
			},
			expTransmittedReports: []ccipocr3.CommitPluginReport{
				{
					PriceUpdates: ccipocr3.PriceUpdates{
						TokenPriceUpdates: []ccipocr3.TokenPrice{
							{
								TokenID: ethAddr,
								Price:   ccipocr3.NewBigInt(ethPrice),
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var reportCodec ccipocr3.CommitPluginCodec
			for i := range oracleIDs {
				paramsCp := params
				paramsCp.reportingCfg.OracleID = oracleIDs[i]
				n := setupNode(paramsCp)
				nodes[i] = n.node
				if i == 0 {
					reportCodec = n.reportCodec
				}

				prepareCcipReaderMock(params.ctx, n.ccipReader, true, false, true)
				tc.mockPriceReader(n.priceReader)
			}

			encodedPrevOutcome, err := tc.prevOutcome.Encode()
			assert.NoError(t, err)
			runner := testhelpers.NewOCR3Runner(nodes, oracleIDs, encodedPrevOutcome)
			res, err := runner.RunRound(params.ctx)
			assert.NoError(t, err)

			decodedOutcome, err := committypes.DecodeOutcome(res.Outcome)
			assert.NoError(t, err)
			assert.Equal(t, normalizeOutcome(tc.expOutcome), normalizeOutcome(decodedOutcome))

			assert.Len(t, res.Transmitted, len(tc.expTransmittedReports))
			for i := range res.Transmitted {
				decoded, err := reportCodec.Decode(params.ctx, res.Transmitted[i].Report)
				assert.NoError(t, err)
				assert.Equal(t, tc.expTransmittedReports[i], decoded)
			}
		})
	}
}

func TestPlugin_E2E_AllNodesAgree_ChainFee(t *testing.T) {
	params := defaultNodeParams(t)
	merkleOutcome := baseMerkleOutcome(params.rmnReportCfg)
	nodes := make([]ocr3types.ReportingPlugin[[]byte], len(oracleIDs))

	newFeeComponents, newNativePrice, packedGasPrice := newRandomFees()
	expectedChain1FeeOutcome := chainfee.Outcome{
		GasPrices: []ccipocr3.GasPriceChain{
			{
				GasPrice: packedGasPrice,
				ChainSel: sourceChain1,
			},
		},
	}

	newFeeComponents2, newNativePrice2, newPackedGasPrice2 := newRandomFees()

	testCases := []struct {
		name                    string
		prevOutcome             committypes.Outcome
		expOutcome              committypes.Outcome
		expTransmittedReportLen int

		mockCCIPReader func(*readerpkg_mock.MockCCIPReader)
	}{
		{
			name:        "fee components should be updated",
			prevOutcome: committypes.Outcome{},
			expOutcome: committypes.Outcome{
				MerkleRootOutcome: merkleOutcome,
				ChainFeeOutcome: chainfee.Outcome{
					GasPrices: []ccipocr3.GasPriceChain{
						{
							GasPrice: packedGasPrice,
							ChainSel: sourceChain1,
						},
						{
							GasPrice: packedGasPrice,
							ChainSel: sourceChain2,
						},
					},
				},
				MainOutcome: committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 10},
			},
			expTransmittedReportLen: 1,
			mockCCIPReader: func(m *readerpkg_mock.MockCCIPReader) {
				m.EXPECT().
					GetChainsFeeComponents(params.ctx, mock.Anything).
					Return(
						map[ccipocr3.ChainSelector]types.ChainFeeComponents{
							sourceChain1: newFeeComponents,
							sourceChain2: newFeeComponents,
						})

				m.EXPECT().
					GetWrappedNativeTokenPriceUSD(params.ctx, mock.Anything).
					Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						sourceChain1: newNativePrice,
						sourceChain2: newNativePrice,
					})
			},
		},
		{
			name:        "fee components should be updated when there's a subset of chains",
			prevOutcome: committypes.Outcome{},
			expOutcome: committypes.Outcome{
				MerkleRootOutcome: merkleOutcome,
				ChainFeeOutcome:   expectedChain1FeeOutcome,
				MainOutcome:       committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 10},
			},
			expTransmittedReportLen: 1,
			mockCCIPReader: func(m *readerpkg_mock.MockCCIPReader) {
				m.EXPECT().
					GetChainsFeeComponents(params.ctx, mock.Anything).
					Return(
						map[ccipocr3.ChainSelector]types.ChainFeeComponents{
							sourceChain1: newFeeComponents,
						})

				m.EXPECT().
					GetWrappedNativeTokenPriceUSD(params.ctx, mock.Anything).
					Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						sourceChain1: newNativePrice,
						sourceChain2: newNativePrice,
					})
			},
		},
		{
			name: "fee components should not be updated when there's a subset of chains but we wait for prices",
			prevOutcome: committypes.Outcome{
				MainOutcome: committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 4},
			},
			expOutcome: committypes.Outcome{
				MerkleRootOutcome: merkleOutcome,
				MainOutcome:       committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 3},
			},
			expTransmittedReportLen: 0,
			mockCCIPReader: func(m *readerpkg_mock.MockCCIPReader) {
				m.EXPECT().GetLatestPriceSeqNr(mock.Anything).Unset()
				m.EXPECT().GetLatestPriceSeqNr(mock.Anything).Return(0, nil).Maybe()
			},
		},
		{
			name: "fee components should not be updated within deviation",
			prevOutcome: committypes.Outcome{
				MerkleRootOutcome: merkleOutcome,
				ChainFeeOutcome:   expectedChain1FeeOutcome,
			},
			expOutcome: committypes.Outcome{
				MerkleRootOutcome: noReportMerkleOutcome(params.rmnReportCfg),
				ChainFeeOutcome:   expectedChain1FeeOutcome,
				MainOutcome:       committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 10},
			},
			expTransmittedReportLen: 1,
			mockCCIPReader: func(m *readerpkg_mock.MockCCIPReader) {
				m.EXPECT().
					GetChainsFeeComponents(params.ctx, mock.Anything).
					Return(
						map[ccipocr3.ChainSelector]types.ChainFeeComponents{
							sourceChain1: newFeeComponents,
						})

				m.EXPECT().
					GetWrappedNativeTokenPriceUSD(params.ctx, mock.Anything).
					Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						sourceChain1: newNativePrice,
					})
			},
		},
		{
			name: "fresh fees (timestamped) should not be updated, even outside of deviation",
			prevOutcome: committypes.Outcome{
				MerkleRootOutcome: merkleOutcome,
				ChainFeeOutcome:   expectedChain1FeeOutcome,
			},
			expOutcome: committypes.Outcome{
				MerkleRootOutcome: noReportMerkleOutcome(params.rmnReportCfg),
				ChainFeeOutcome: chainfee.Outcome{
					GasPrices: []ccipocr3.GasPriceChain{
						{
							GasPrice: newPackedGasPrice2,
							ChainSel: sourceChain1,
						},
					},
				},
				MainOutcome: committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 10},
			},
			expTransmittedReportLen: 1,
			mockCCIPReader: func(m *readerpkg_mock.MockCCIPReader) {
				m.EXPECT().
					GetChainsFeeComponents(params.ctx, mock.Anything).
					Return(
						map[ccipocr3.ChainSelector]types.ChainFeeComponents{
							sourceChain1: newFeeComponents2,
						})

				m.EXPECT().
					GetWrappedNativeTokenPriceUSD(params.ctx, mock.Anything).
					Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						sourceChain1: newNativePrice2,
					})
			},
		},
		{
			name: "stale fees should be updated",
			prevOutcome: committypes.Outcome{
				MerkleRootOutcome: merkleOutcome,
				ChainFeeOutcome:   expectedChain1FeeOutcome,
			},
			expOutcome: committypes.Outcome{
				MerkleRootOutcome: noReportMerkleOutcome(params.rmnReportCfg),
				ChainFeeOutcome: chainfee.Outcome{
					GasPrices: []ccipocr3.GasPriceChain{
						{
							GasPrice: newPackedGasPrice2,
							ChainSel: sourceChain1,
						},
					},
				},
				MainOutcome: committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 10},
			},
			expTransmittedReportLen: 1,
			mockCCIPReader: func(m *readerpkg_mock.MockCCIPReader) {
				m.EXPECT().
					GetChainsFeeComponents(params.ctx, mock.Anything).
					Return(
						map[ccipocr3.ChainSelector]types.ChainFeeComponents{
							sourceChain1: newFeeComponents2,
						})
				m.EXPECT().
					GetWrappedNativeTokenPriceUSD(params.ctx, mock.Anything).
					Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						sourceChain1: newNativePrice2,
					})

				m.EXPECT().GetChainFeePriceUpdate(params.ctx, mock.Anything).Unset()
				elapsed, err := time.ParseDuration("3h")
				require.NoError(t, err)
				t := time.Now().UTC().Add(-elapsed)
				m.EXPECT().
					GetChainFeePriceUpdate(params.ctx, mock.Anything).
					Return(map[ccipocr3.ChainSelector]plugintypes.TimestampedBig{
						sourceChain1: {
							Timestamp: t,
							Value:     expectedChain1FeeOutcome.GasPrices[0].GasPrice,
						},
					})
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for i := range oracleIDs {
				paramsCp := params
				paramsCp.reportingCfg.OracleID = oracleIDs[i]
				n := setupNode(paramsCp)
				nodes[i] = n.node

				prepareCcipReaderMock(params.ctx, n.ccipReader, true, false, false)

				preparePriceReaderMock(params.ctx, n.priceReader)

				tc.mockCCIPReader(n.ccipReader)
			}

			encodedPrevOutcome, err := tc.prevOutcome.Encode()
			assert.NoError(t, err)
			runner := testhelpers.NewOCR3Runner(nodes, oracleIDs, encodedPrevOutcome)
			res, err := runner.RunRound(params.ctx)
			assert.NoError(t, err)

			decodedOutcome, err := committypes.DecodeOutcome(res.Outcome)
			assert.NoError(t, err)
			assert.Equal(t, normalizeOutcome(tc.expOutcome), normalizeOutcome(decodedOutcome))

			assert.Len(t, res.Transmitted, tc.expTransmittedReportLen)
		})
	}
}

// normalizeOutcome converts empty slices to nil or nil slices to empty where needed.
func normalizeOutcome(o committypes.Outcome) committypes.Outcome {
	if len(o.MerkleRootOutcome.RMNRemoteCfg.ContractAddress) == 0 {
		// Normalize to `nil` if it's an empty slice
		o.MerkleRootOutcome.RMNRemoteCfg.ContractAddress = nil
	}
	return o
}

func prepareCcipReaderMock(
	ctx context.Context,
	ccipReader *readerpkg_mock.MockCCIPReader,
	mockEmptySeqNrs bool,
	enableDiscovery bool,
	mockChainFee bool) {

	if mockChainFee {
		ccipReader.EXPECT().
			GetChainsFeeComponents(ctx, mock.Anything).
			Return(map[ccipocr3.ChainSelector]types.ChainFeeComponents{}).Maybe()
		ccipReader.EXPECT().
			GetWrappedNativeTokenPriceUSD(ctx, mock.Anything).
			Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{}).Maybe()
	}
	ccipReader.EXPECT().
		GetLatestPriceSeqNr(ctx).
		Return(0, nil).Maybe()
	ccipReader.EXPECT().
		GetChainFeePriceUpdate(ctx, mock.Anything).
		Return(map[ccipocr3.ChainSelector]plugintypes.TimestampedBig{}).Maybe()
	ccipReader.EXPECT().
		GetContractAddress(mock.Anything, mock.Anything).
		Return(ccipocr3.Bytes{}, nil).Maybe()
	ccipReader.EXPECT().GetRmnCurseInfo(mock.Anything, mock.Anything, mock.Anything).
		Return(&reader2.CurseInfo{}, nil).Maybe()

	if mockEmptySeqNrs {
		ccipReader.EXPECT().NextSeqNum(ctx, mock.Anything).Unset()
		ccipReader.EXPECT().NextSeqNum(ctx, mock.Anything).Return(map[ccipocr3.ChainSelector]ccipocr3.SeqNum{}, nil).
			Maybe()
	}

	if enableDiscovery {
		ccipReader.EXPECT().DiscoverContracts(mock.Anything).Return(nil, nil)
		ccipReader.EXPECT().Sync(mock.Anything, mock.Anything).Return(nil)
	}
}

func preparePriceReaderMock(ctx context.Context, priceReader *readerpkg_mock.MockPriceReader) {
	priceReader.EXPECT().
		GetFeeQuoterTokenUpdates(ctx, mock.Anything, mock.Anything).
		Return(
			map[ccipocr3.UnknownEncodedAddress]plugintypes.TimestampedBig{}, nil,
		).
		Maybe()

	priceReader.EXPECT().
		GetFeedPricesUSD(ctx, mock.Anything).
		Return([]*big.Int{}, nil).
		Maybe()
}

type nodeSetup struct {
	node        *Plugin
	ccipReader  *readerpkg_mock.MockCCIPReader
	priceReader *readerpkg_mock.MockPriceReader
	reportCodec *mocks.CommitPluginJSONReportCodec
	msgHasher   *mocks.MessageHasher
}

// Define a struct to hold the parameters
type SetupNodeParams struct {
	ctx               context.Context
	t                 *testing.T
	lggr              logger.Logger
	donID             plugintypes.DonID
	reportingCfg      ocr3types.ReportingPluginConfig
	oracleIDToP2pID   map[commontypes.OracleID]libocrtypes.PeerID
	offchainCfg       pluginconfig.CommitOffchainConfig
	chainCfg          map[ccipocr3.ChainSelector]reader.ChainConfig
	offRampNextSeqNum map[ccipocr3.ChainSelector]ccipocr3.SeqNum
	onRampLastSeqNum  map[ccipocr3.ChainSelector]ccipocr3.SeqNum
	rmnReportCfg      rmntypes.RemoteConfig
	enableDiscovery   bool
}

//nolint:gocyclo // todo
func setupNode(params SetupNodeParams) nodeSetup {
	ccipReader := readerpkg_mock.NewMockCCIPReader(params.t)
	tokenPricesReader := readerpkg_mock.NewMockPriceReader(params.t)
	reportCodec := mocks.NewCommitPluginJSONReportCodec()
	msgHasher := mocks.NewMessageHasher()
	homeChainReader := reader_mock.NewMockHomeChain(params.t)
	rmnHomeReader := readerpkg_mock.NewMockRMNHome(params.t)

	fChain := map[ccipocr3.ChainSelector]int{}
	supportedChainsForPeer := make(map[libocrtypes.PeerID]mapset.Set[ccipocr3.ChainSelector])
	for chainSel, cfg := range params.chainCfg {
		fChain[chainSel] = cfg.FChain

		for _, peerID := range cfg.SupportedNodes.ToSlice() {
			if _, ok := supportedChainsForPeer[peerID]; !ok {
				supportedChainsForPeer[peerID] = mapset.NewSet[ccipocr3.ChainSelector]()
			}
			supportedChainsForPeer[peerID].Add(chainSel)
		}
	}

	homeChainReader.EXPECT().GetFChain().Return(fChain, nil)
	homeChainReader.EXPECT().
		GetOCRConfigs(mock.Anything, params.donID, consts.PluginTypeCommit).
		Return(reader.ActiveAndCandidate{
			ActiveConfig: reader.OCR3ConfigWithMeta{
				ConfigDigest: params.reportingCfg.ConfigDigest,
			},
		}, nil).Maybe()

	for peerID, supportedChains := range supportedChainsForPeer {
		homeChainReader.EXPECT().GetSupportedChainsForPeer(peerID).Return(supportedChains, nil).Maybe()
	}

	knownCCIPChains := mapset.NewSet[ccipocr3.ChainSelector]()

	for chainSel, cfg := range params.chainCfg {
		homeChainReader.EXPECT().GetChainConfig(chainSel).Return(cfg, nil).Maybe()
		knownCCIPChains.Add(chainSel)
	}
	homeChainReader.EXPECT().GetKnownCCIPChains().Return(knownCCIPChains, nil).Maybe()

	sourceChains := make([]ccipocr3.ChainSelector, 0)
	for chainSel := range params.offRampNextSeqNum {
		sourceChains = append(sourceChains, chainSel)
	}
	sort.Slice(sourceChains, func(i, j int) bool { return sourceChains[i] < sourceChains[j] })

	offRampNextSeqNums := make(map[ccipocr3.ChainSelector]ccipocr3.SeqNum, 0)
	chainsWithNewMsgs := make([]ccipocr3.ChainSelector, 0)
	for _, sourceChain := range sourceChains {
		offRampNextSeqNum, ok := params.offRampNextSeqNum[sourceChain]
		assert.True(params.t, ok)
		offRampNextSeqNums[sourceChain] = offRampNextSeqNum

		newMsgs := make([]ccipocr3.Message, 0)
		numNewMsgs := (params.onRampLastSeqNum[sourceChain] - offRampNextSeqNum) + 1
		for i := uint64(0); i < uint64(numNewMsgs); i++ {
			messageID := sha256.Sum256([]byte(fmt.Sprintf("%d", uint64(offRampNextSeqNum)+i)))
			newMsgs = append(newMsgs, ccipocr3.Message{
				Header: ccipocr3.RampMessageHeader{
					MessageID:      messageID,
					SequenceNumber: offRampNextSeqNum + ccipocr3.SeqNum(i),
				},
			})
		}

		ccipReader.EXPECT().MsgsBetweenSeqNums(params.ctx, sourceChain,
			ccipocr3.NewSeqNumRange(offRampNextSeqNum, offRampNextSeqNum)).
			Return(newMsgs, nil).Maybe()

		if len(newMsgs) > 0 {
			chainsWithNewMsgs = append(chainsWithNewMsgs, sourceChain)
		}
	}

	seqNumsOfChainsWithNewMsgs := map[ccipocr3.ChainSelector]ccipocr3.SeqNum{}
	for _, chainSel := range chainsWithNewMsgs {
		seqNumsOfChainsWithNewMsgs[chainSel] = params.offRampNextSeqNum[chainSel]
	}
	if len(chainsWithNewMsgs) > 0 {
		ccipReader.EXPECT().NextSeqNum(params.ctx, chainsWithNewMsgs).Return(seqNumsOfChainsWithNewMsgs, nil).Maybe()
	}
	ccipReader.EXPECT().NextSeqNum(params.ctx, sourceChains).Return(offRampNextSeqNums, nil).Maybe()

	for _, ch := range sourceChains {
		ccipReader.EXPECT().GetExpectedNextSequenceNumber(
			params.ctx, ch, destChain).Return(params.offRampNextSeqNum[ch]+1, nil).Maybe()
	}

	ccipReader.EXPECT().
		GetRMNRemoteConfig(params.ctx, mock.Anything).
		Return(params.rmnReportCfg, nil).Maybe()

	ccipReader.EXPECT().
		GetOffRampConfigDigest(params.ctx, consts.PluginTypeCommit).
		Return(params.reportingCfg.ConfigDigest, nil).Maybe()

	p := NewPlugin(
		params.donID,
		params.oracleIDToP2pID,
		params.offchainCfg,
		destChain,
		ccipReader,
		tokenPricesReader,
		reportCodec,
		msgHasher,
		params.lggr,
		homeChainReader,
		rmnHomeReader,
		nil,
		nil,
		params.reportingCfg,
		&metrics.Noop{},
	)

	if !params.enableDiscovery {
		p.discoveryProcessor = nil
	}
	return nodeSetup{
		node:        p,
		ccipReader:  ccipReader,
		priceReader: tokenPricesReader,
		reportCodec: reportCodec,
		msgHasher:   msgHasher,
	}
}

// Returns default values for setting up a Node
// Note:
// oracleID will be set to 0
func defaultNodeParams(t *testing.T) SetupNodeParams {
	ctx := tests.Context(t)
	lggr := logger.Test(t)

	donID := uint32(1)
	rb := rand.RandomBytes32()
	digest := ocrtypes.ConfigDigest(rb[:])

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

	rmnRemoteCfg := testhelpers.CreateRMNRemoteCfg()

	writeFrequency := *commonconfig.MustNewDuration(1 * time.Minute)
	cfg := pluginconfig.CommitOffchainConfig{
		NewMsgScanBatchSize:                100,
		MaxReportTransmissionCheckAttempts: 2,
		TokenPriceBatchWriteFrequency:      writeFrequency,
		TokenInfo: map[ccipocr3.UnknownEncodedAddress]pluginconfig.TokenInfo{
			arbAddr: arbInfo,
			ethAddr: ethInfo,
		},
		PriceFeedChainSelector:    sourceChain1,
		InflightPriceCheckRetries: 10,
	}

	reportingCfg := ocr3types.ReportingPluginConfig{F: 1, ConfigDigest: digest}

	params := SetupNodeParams{
		ctx:               ctx,
		t:                 t,
		lggr:              lggr,
		donID:             donID,
		reportingCfg:      reportingCfg,
		oracleIDToP2pID:   oracleIDToPeerID,
		offchainCfg:       cfg,
		chainCfg:          homeChainConfig,
		offRampNextSeqNum: offRampNextSeqNum,
		onRampLastSeqNum:  onRampLastSeqNum,
		rmnReportCfg:      rmnRemoteCfg,
		enableDiscovery:   false,
	}

	return params
}

// merkleRoot1 is the markle root that the test generates, the merkle root generation logic is not supposed to be
// tested in this context, so we just assume it's correct.
var merkleRoot1 = ccipocr3.Bytes32{0x4a, 0x44, 0xdc, 0x15, 0x36, 0x42, 0x4, 0xa8, 0xf, 0xe8, 0xe,
	0x90, 0x39, 0x45, 0x5c, 0xc1, 0x60, 0x82, 0x81, 0x82, 0xf, 0xe2, 0xb2, 0x4f, 0x1e, 0x52,
	0x33, 0xad, 0xe6, 0xaf, 0x1d, 0xd5}

func baseMerkleOutcome(r rmntypes.RemoteConfig) merkleroot.Outcome {
	return merkleroot.Outcome{
		OutcomeType:             merkleroot.ReportIntervalsSelected,
		RangesSelectedForReport: []plugintypes.ChainRange{},
		OffRampNextSeqNums:      []plugintypes.SeqNumChain{},
		RMNRemoteCfg:            r,
	}
}

func noReportMerkleOutcome(r rmntypes.RemoteConfig) merkleroot.Outcome {
	return merkleroot.Outcome{
		OutcomeType:             merkleroot.ReportEmpty,
		RangesSelectedForReport: nil,
		RootsToReport:           []ccipocr3.MerkleRootChain{},
		OffRampNextSeqNums:      []plugintypes.SeqNumChain{},
		RMNReportSignatures:     []ccipocr3.RMNECDSASignature{},
		RMNRemoteCfg:            r,
	}
}

func newRandomFees() (types.ChainFeeComponents, ccipocr3.BigInt, ccipocr3.BigInt) {
	execFee := big.NewInt(rand.RandomInt64())
	dataAvFee := big.NewInt(rand.RandomInt64())
	nativePrice := big.NewInt(rand.RandomInt64())
	usdPrices := chainfee.FeeComponentsToPackedFee(chainfee.ComponentsUSDPrices{
		ExecutionFeePriceUSD: mathslib.CalculateUsdPerUnitGas(execFee, nativePrice),
		DataAvFeePriceUSD:    mathslib.CalculateUsdPerUnitGas(dataAvFee, nativePrice),
	})

	return types.ChainFeeComponents{ExecutionFee: execFee, DataAvailabilityFee: dataAvFee},
		ccipocr3.NewBigInt(nativePrice),
		ccipocr3.NewBigInt(usdPrices)
}
