package commit

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"sort"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	ocrtypes "github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	sel "github.com/smartcontractkit/chain-selectors"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/chainconfig"
	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/internal/builder"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/metrics"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	reader_mock "github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/consts"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	reader2 "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

const (
	destChain         = ccipocr3.ChainSelector(1)
	arbAddr           = ccipocr3.UnknownEncodedAddress("0xa100000000000000000000000000000000000000")
	arbAggregatorAddr = ccipocr3.UnknownEncodedAddress("0xa2000000000000000000000000000000000000000")

	ethAddr           = ccipocr3.UnknownEncodedAddress("0xe100000000000000000000000000000000000000")
	ethAggregatorAddr = ccipocr3.UnknownEncodedAddress("0xe200000000000000000000000000000000000000")
)

var (
	sourceEvmChain1 = ccipocr3.ChainSelector(sel.ETHEREUM_TESTNET_SEPOLIA.Selector)
	sourceSolChain  = ccipocr3.ChainSelector(sel.SOLANA_DEVNET.Selector)

	oracleIDs = []commontypes.OracleID{1, 2, 3}
	peerIDs   = []libocrtypes.PeerID{{1}, {2}, {3}}

	arbPrice = new(big.Int).Mul(big.NewInt(5), big.NewInt(1e18))
	ethPrice = new(big.Int).Mul(big.NewInt(7), big.NewInt(1e18))

	// a map to ease working with tests
	tokenPriceMap = ccipocr3.TokenPriceMap{
		arbAddr: ccipocr3.NewBigInt(arbPrice),
		ethAddr: ccipocr3.NewBigInt(ethPrice),
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

	sourceChainConfigs = map[ccipocr3.ChainSelector]reader2.StaticSourceChainConfig{
		sourceEvmChain1: {IsEnabled: true, IsRMNVerificationDisabled: true},
		sourceSolChain:  {IsEnabled: true, IsRMNVerificationDisabled: true},
	}
)

var ocrTypCodec = ocrtypecodec.DefaultCommitCodec

func TestPlugin_E2E_AllNodesAgree_MerkleRoots(t *testing.T) {
	params := defaultNodeParams(t)
	nodes := make([]ocr3types.ReportingPlugin[[]byte], len(oracleIDs))

	outcomeIntervalsSelected := committypes.Outcome{
		MerkleRootOutcome: merkleroot.Outcome{
			OutcomeType: merkleroot.ReportIntervalsSelected,
			RangesSelectedForReport: []plugintypes.ChainRange{
				{ChainSel: sourceEvmChain1, SeqNumRange: ccipocr3.SeqNumRange{10, 10}},
				{ChainSel: sourceSolChain, SeqNumRange: ccipocr3.SeqNumRange{20, 20}},
			},
			OffRampNextSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: sourceEvmChain1, SeqNum: 10},
				{ChainSel: sourceSolChain, SeqNum: 20},
			},
			RMNRemoteCfg: params.rmnReportCfg,
		},
	}

	outcomeReportGenerated := committypes.Outcome{
		MerkleRootOutcome: merkleroot.Outcome{
			OutcomeType: merkleroot.ReportGenerated,
			RootsToReport: []ccipocr3.MerkleRootChain{
				{
					ChainSel:      sourceEvmChain1,
					OnRampAddress: ccipocr3.UnknownAddress{1},
					SeqNumsRange:  ccipocr3.SeqNumRange{0xa, 0xa},
					MerkleRoot:    merkleRoot1,
				},
			},
			OffRampNextSeqNums: []plugintypes.SeqNumChain{
				{ChainSel: sourceEvmChain1, SeqNum: 10},
				{ChainSel: sourceSolChain, SeqNum: 20},
			},
			RMNRemoteCfg: params.rmnReportCfg,
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
					UnblessedMerkleRoots: []ccipocr3.MerkleRootChain{
						{
							ChainSel:      sourceEvmChain1,
							SeqNumsRange:  ccipocr3.NewSeqNumRange(0xa, 0xa),
							OnRampAddress: ccipocr3.UnknownAddress{1},
							MerkleRoot:    merkleRoot1,
						},
					},
					BlessedMerkleRoots: make([]ccipocr3.MerkleRootChain, 0),
					PriceUpdates:       ccipocr3.PriceUpdates{},
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
						{ChainSel: sourceEvmChain1, SeqNum: 10},
						{ChainSel: sourceSolChain, SeqNum: 20},
					},
					RootsToReport: []ccipocr3.MerkleRootChain{
						{
							ChainSel:      sourceEvmChain1,
							SeqNumsRange:  ccipocr3.NewSeqNumRange(0xa, 0xa),
							OnRampAddress: ccipocr3.UnknownAddress{1},
							MerkleRoot:    merkleRoot1,
						},
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
			offRampNextSeqNumDefaultOverrideKeys: []ccipocr3.ChainSelector{sourceEvmChain1, sourceSolChain},
			offRampNextSeqNumDefaultOverrideValues: map[ccipocr3.ChainSelector]ccipocr3.SeqNum{
				sourceEvmChain1: 11,
				sourceSolChain:  20,
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
				prepareCcipReaderMock(n.ccipReader, false, tc.enableDiscovery, true)

				if len(tc.offRampNextSeqNumDefaultOverrideKeys) > 0 {
					require.Equal(t, len(tc.offRampNextSeqNumDefaultOverrideKeys), len(tc.offRampNextSeqNumDefaultOverrideValues))
					n.ccipReader.EXPECT().NextSeqNum(mock.Anything, tc.offRampNextSeqNumDefaultOverrideKeys).Unset()
					n.ccipReader.EXPECT().
						NextSeqNum(mock.Anything, tc.offRampNextSeqNumDefaultOverrideKeys).
						Return(tc.offRampNextSeqNumDefaultOverrideValues, nil).
						Maybe()
				}

				preparePriceReaderMock(n.priceReader)
			}

			encodedPrevOutcome, err := ocrTypCodec.EncodeOutcome(tc.prevOutcome)
			assert.NoError(t, err)
			runner := testhelpers.NewOCR3Runner(nodes, oracleIDs, encodedPrevOutcome)
			res, err := runner.RunRound(params.ctx)
			assert.NoError(t, err)

			decodedOutcome, err := ocrTypCodec.DecodeOutcome(res.Outcome)
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

	nodes := make([]ocr3types.ReportingPlugin[[]byte], len(oracleIDs))

	merkleOutcome := reportEmptyMerkleRootOutcome()

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
					GetFeedPricesUSD(mock.Anything, mock.MatchedBy(func(tokens []ccipocr3.UnknownEncodedAddress) bool {
						expectedTokens := mapset.NewSet(arbAddr, ethAddr)
						actualTokens := mapset.NewSet(tokens...)
						return expectedTokens.Equal(actualTokens)
					})).
					Return(ccipocr3.TokenPriceMap{
						arbAddr: ccipocr3.NewBigInt(arbPrice),
						ethAddr: ccipocr3.NewBigInt(ethPrice),
					}, nil).Maybe()

				m.EXPECT().
					GetFeeQuoterTokenUpdates(mock.Anything, mock.Anything, mock.Anything).
					Return(
						map[ccipocr3.UnknownEncodedAddress]ccipocr3.TimestampedBig{}, nil,
					).
					Maybe()
			},
			expOutcome: committypes.Outcome{
				MerkleRootOutcome: merkleOutcome,
				TokenPriceOutcome: tokenprice.Outcome{
					TokenPrices: tokenPriceMap,
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
					GetFeedPricesUSD(mock.Anything, mock.MatchedBy(func(tokens []ccipocr3.UnknownEncodedAddress) bool {
						expectedTokens := mapset.NewSet(arbAddr, ethAddr)
						actualTokens := mapset.NewSet(tokens...)
						return expectedTokens.Equal(actualTokens)
					})).
					Return(ccipocr3.TokenPriceMap{
						arbAddr: ccipocr3.NewBigInt(arbPrice),
						ethAddr: ccipocr3.NewBigInt(ethPrice),
					}, nil).
					Maybe()

				// Arb is fresh, will not be updated
				m.EXPECT().
					GetFeeQuoterTokenUpdates(mock.Anything, mock.Anything, mock.Anything).
					Return(
						map[ccipocr3.UnknownEncodedAddress]ccipocr3.TimestampedBig{
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
					GetFeedPricesUSD(mock.Anything, mock.MatchedBy(func(tokens []ccipocr3.UnknownEncodedAddress) bool {
						expectedTokens := mapset.NewSet(arbAddr, ethAddr)
						actualTokens := mapset.NewSet(tokens...)
						return expectedTokens.Equal(actualTokens)
					})).
					Return(ccipocr3.TokenPriceMap{
						arbAddr: ccipocr3.NewBigInt(arbPrice),
						ethAddr: ccipocr3.NewBigInt(ethPrice),
					}, nil).Maybe()

				m.EXPECT().
					GetFeeQuoterTokenUpdates(mock.Anything, mock.Anything, mock.Anything).
					Return(
						map[ccipocr3.UnknownEncodedAddress]ccipocr3.TimestampedBig{
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
					TokenPrices: ccipocr3.TokenPriceMap{
						ethAddr: ccipocr3.NewBigInt(ethPrice),
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

				prepareCcipReaderMock(n.ccipReader, true, false, true)
				tc.mockPriceReader(n.priceReader)
			}

			encodedPrevOutcome, err := ocrTypCodec.EncodeOutcome(tc.prevOutcome)
			assert.NoError(t, err)
			runner := testhelpers.NewOCR3Runner(nodes, oracleIDs, encodedPrevOutcome)
			res, err := runner.RunRound(params.ctx)
			assert.NoError(t, err)

			decodedOutcome, err := ocrTypCodec.DecodeOutcome(res.Outcome)
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
	merkleOutcome := reportEmptyMerkleRootOutcome()
	nodes := make([]ocr3types.ReportingPlugin[[]byte], len(oracleIDs))

	// evm selector
	evmFeeComponents, evmNativePrice, evmPackedGasPrice := newRandomFees(sourceEvmChain1)
	expectedChain1FeeOutcome := chainfee.Outcome{
		GasPrices: []ccipocr3.GasPriceChain{
			{
				GasPrice: evmPackedGasPrice,
				ChainSel: sourceEvmChain1,
			},
		},
	}

	// sol selector
	solFeeComponents, solNativePrice, solPackedPrice := newRandomFees(sourceSolChain)

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
							GasPrice: evmPackedGasPrice,
							ChainSel: sourceEvmChain1,
						},
						{
							GasPrice: solPackedPrice,
							ChainSel: sourceSolChain,
						},
					},
				},
				MainOutcome: committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 10},
			},
			expTransmittedReportLen: 1,
			mockCCIPReader: func(m *readerpkg_mock.MockCCIPReader) {
				m.EXPECT().
					GetChainsFeeComponents(mock.Anything, mock.Anything).
					Return(
						map[ccipocr3.ChainSelector]types.ChainFeeComponents{
							sourceEvmChain1: evmFeeComponents,
							sourceSolChain:  solFeeComponents,
						})

				m.EXPECT().
					GetWrappedNativeTokenPriceUSD(mock.Anything, mock.Anything).
					Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						sourceEvmChain1: evmNativePrice,
						sourceSolChain:  solNativePrice,
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
					GetChainsFeeComponents(mock.Anything, mock.Anything).
					Return(
						map[ccipocr3.ChainSelector]types.ChainFeeComponents{
							sourceEvmChain1: evmFeeComponents,
						})

				m.EXPECT().
					GetWrappedNativeTokenPriceUSD(mock.Anything, mock.Anything).
					Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						sourceEvmChain1: evmNativePrice,
						sourceSolChain:  solNativePrice,
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
				MerkleRootOutcome: merkleOutcome,
				ChainFeeOutcome:   expectedChain1FeeOutcome,
				MainOutcome:       committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 10},
			},
			expTransmittedReportLen: 1,
			mockCCIPReader: func(m *readerpkg_mock.MockCCIPReader) {
				m.EXPECT().
					GetChainsFeeComponents(mock.Anything, mock.Anything).
					Return(
						map[ccipocr3.ChainSelector]types.ChainFeeComponents{
							sourceEvmChain1: evmFeeComponents,
						})

				m.EXPECT().
					GetWrappedNativeTokenPriceUSD(mock.Anything, mock.Anything).
					Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						sourceEvmChain1: evmNativePrice,
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
				MerkleRootOutcome: reportEmptyMerkleRootOutcome(),
				ChainFeeOutcome: chainfee.Outcome{
					GasPrices: []ccipocr3.GasPriceChain{
						{
							GasPrice: solPackedPrice,
							ChainSel: sourceSolChain,
						},
					},
				},
				MainOutcome: committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 10},
			},
			expTransmittedReportLen: 1,
			mockCCIPReader: func(m *readerpkg_mock.MockCCIPReader) {
				m.EXPECT().
					GetChainsFeeComponents(mock.Anything, mock.Anything).
					Return(
						map[ccipocr3.ChainSelector]types.ChainFeeComponents{
							sourceSolChain: solFeeComponents,
						})

				m.EXPECT().
					GetWrappedNativeTokenPriceUSD(mock.Anything, mock.Anything).
					Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						sourceSolChain: solNativePrice,
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
				MerkleRootOutcome: reportEmptyMerkleRootOutcome(),
				ChainFeeOutcome: chainfee.Outcome{
					GasPrices: []ccipocr3.GasPriceChain{
						{
							GasPrice: solPackedPrice,
							ChainSel: sourceSolChain,
						},
					},
				},
				MainOutcome: committypes.MainOutcome{InflightPriceOcrSequenceNumber: 1, RemainingPriceChecks: 10},
			},
			expTransmittedReportLen: 1,
			mockCCIPReader: func(m *readerpkg_mock.MockCCIPReader) {
				m.EXPECT().
					GetChainsFeeComponents(mock.Anything, mock.Anything).
					Return(
						map[ccipocr3.ChainSelector]types.ChainFeeComponents{
							sourceSolChain: solFeeComponents,
						})
				m.EXPECT().
					GetWrappedNativeTokenPriceUSD(mock.Anything, mock.Anything).
					Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{
						sourceSolChain: solNativePrice,
					})

				m.EXPECT().GetChainFeePriceUpdate(mock.Anything, mock.Anything).Unset()
				elapsed, err := time.ParseDuration("3h")
				require.NoError(t, err)
				t := time.Now().UTC().Add(-elapsed)
				m.EXPECT().
					GetChainFeePriceUpdate(mock.Anything, mock.Anything).
					Return(map[ccipocr3.ChainSelector]ccipocr3.TimestampedBig{
						sourceSolChain: {
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

				prepareCcipReaderMock(n.ccipReader, true, false, false)

				preparePriceReaderMock(n.priceReader)

				tc.mockCCIPReader(n.ccipReader)
			}

			encodedPrevOutcome, err := ocrTypCodec.EncodeOutcome(tc.prevOutcome)
			assert.NoError(t, err)
			runner := testhelpers.NewOCR3Runner(nodes, oracleIDs, encodedPrevOutcome)
			res, err := runner.RunRound(params.ctx)
			require.NoError(t, err)

			decodedOutcome, err := ocrTypCodec.DecodeOutcome(res.Outcome)
			require.NoError(t, err)
			require.Equal(t, normalizeOutcome(tc.expOutcome), normalizeOutcome(decodedOutcome))

			require.Len(t, res.Transmitted, tc.expTransmittedReportLen)
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
	ccipReader *readerpkg_mock.MockCCIPReader,
	mockEmptySeqNrs bool,
	enableDiscovery bool,
	mockChainFee bool,
) {
	if mockChainFee {
		ccipReader.EXPECT().
			GetChainsFeeComponents(mock.Anything, mock.Anything).
			Return(map[ccipocr3.ChainSelector]types.ChainFeeComponents{}).Maybe()
		ccipReader.EXPECT().
			GetWrappedNativeTokenPriceUSD(mock.Anything, mock.Anything).
			Return(map[ccipocr3.ChainSelector]ccipocr3.BigInt{}).Maybe()
	}
	ccipReader.EXPECT().
		GetLatestPriceSeqNr(mock.Anything).
		Return(0, nil).Maybe()
	ccipReader.EXPECT().
		GetChainFeePriceUpdate(mock.Anything, mock.Anything).
		Return(map[ccipocr3.ChainSelector]ccipocr3.TimestampedBig{}).Maybe()
	ccipReader.EXPECT().
		GetContractAddress(mock.Anything, mock.Anything).
		Return(ccipocr3.Bytes{1}, nil).Maybe()
	ccipReader.EXPECT().GetRmnCurseInfo(mock.Anything).
		Return(reader2.CurseInfo{}, nil).Maybe()
	ccipReader.EXPECT().GetOffRampSourceChainsConfig(mock.Anything, mock.Anything).
		Return(sourceChainConfigs, nil).Maybe()

	if mockEmptySeqNrs {
		ccipReader.EXPECT().NextSeqNum(mock.Anything, mock.Anything).Unset()
		ccipReader.EXPECT().NextSeqNum(mock.Anything, mock.Anything).
			Return(map[ccipocr3.ChainSelector]ccipocr3.SeqNum{}, nil).
			Maybe()
	}

	if enableDiscovery {
		ccipReader.EXPECT().DiscoverContracts(mock.Anything, mock.Anything).Return(nil, nil)
		ccipReader.EXPECT().Sync(mock.Anything, mock.Anything).Return(nil)
	}
}

func preparePriceReaderMock(priceReader *readerpkg_mock.MockPriceReader) {
	priceReader.EXPECT().
		GetFeeQuoterTokenUpdates(mock.Anything, mock.Anything, mock.Anything).
		Return(
			nil, nil,
		).
		Maybe()

	priceReader.EXPECT().
		GetFeedPricesUSD(mock.Anything, mock.Anything).
		Return(nil, nil).Maybe()
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
	rmnReportCfg      ccipocr3.RemoteConfig
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

	rmnHomeReader.EXPECT().GetRMNEnabledSourceChains(mock.Anything).Return(map[ccipocr3.ChainSelector]bool{
		sourceEvmChain1: false,
		sourceSolChain:  false,
	}, nil).Maybe()

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
	if params.enableDiscovery {
		homeChainReader.EXPECT().GetAllChainConfigs().Return(params.chainCfg, nil)
	}
	homeChainReader.EXPECT().GetOCRConfigs(mock.Anything, params.donID, consts.PluginTypeCommit).
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
					OnRamp:         ccipocr3.UnknownAddress{1},
				},
			})
		}

		ccipReader.EXPECT().MsgsBetweenSeqNums(mock.Anything, sourceChain,
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
		ccipReader.EXPECT().NextSeqNum(mock.Anything, chainsWithNewMsgs).Return(seqNumsOfChainsWithNewMsgs, nil).Maybe()
	}
	ccipReader.EXPECT().NextSeqNum(mock.Anything, sourceChains).Return(offRampNextSeqNums, nil).Maybe()

	for _, ch := range sourceChains {
		ccipReader.EXPECT().LatestMsgSeqNum(
			mock.Anything, ch).Return(params.offRampNextSeqNum[ch], nil).Maybe()
	}

	ccipReader.EXPECT().
		GetRMNRemoteConfig(mock.Anything).
		Return(params.rmnReportCfg, nil).Maybe()

	ccipReader.EXPECT().
		GetOffRampConfigDigest(mock.Anything, consts.PluginTypeCommit).
		Return(params.reportingCfg.ConfigDigest, nil).Maybe()

	cfg := pluginconfig.CommitOffchainConfig{}
	err := cfg.ApplyDefaultsAndValidate()
	require.NoError(params.t, err)
	reportBuilder, err := builder.NewReportBuilder(cfg.RMNEnabled, cfg.MaxMerkleRootsPerReport, cfg.MaxPricesPerReport)
	require.NoError(params.t, err)

	mockAddrCodec := internal.NewMockAddressCodecHex(params.t)
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
		mockAddrCodec,
		reportBuilder,
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
		destChain:       {FChain: 1, SupportedNodes: peerIDsMap, Config: chainconfig.ChainConfig{}},
		sourceEvmChain1: {FChain: 1, SupportedNodes: peerIDsMap, Config: chainconfig.ChainConfig{}},
		sourceSolChain:  {FChain: 1, SupportedNodes: peerIDsMap, Config: chainconfig.ChainConfig{}},
	}

	offRampNextSeqNum := map[ccipocr3.ChainSelector]ccipocr3.SeqNum{
		sourceEvmChain1: 10,
		sourceSolChain:  20,
	}

	onRampLastSeqNum := map[ccipocr3.ChainSelector]ccipocr3.SeqNum{
		sourceEvmChain1: 10, // one new msg -> 10
		sourceSolChain:  19, // no new msg, still on 19
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
		PriceFeedChainSelector:          sourceEvmChain1,
		InflightPriceCheckRetries:       10,
		MerkleRootAsyncObserverDisabled: true, // we want to keep it disabled since this test is deterministic
		ChainFeeAsyncObserverDisabled:   true,
		TokenPriceAsyncObserverDisabled: true,
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

func reportEmptyMerkleRootOutcome() merkleroot.Outcome {
	return merkleroot.Outcome{OutcomeType: merkleroot.ReportEmpty}
}

func newRandomFees(selector ccipocr3.ChainSelector) (components types.ChainFeeComponents,
	nativePrice ccipocr3.BigInt,
	usdPrices ccipocr3.BigInt) {
	execFee := big.NewInt(rand.RandomInt64())
	dataAvFee := big.NewInt(rand.RandomInt64())
	nativePriceI := big.NewInt(rand.RandomInt64())
	usdPricesF := chainfee.FeeComponentsToPackedFee(chainfee.ComponentsUSDPrices{
		ExecutionFeePriceUSD: internal.MustCalculateUsdPerUnitGas(selector, execFee, nativePriceI),
		DataAvFeePriceUSD:    internal.MustCalculateUsdPerUnitGas(selector, dataAvFee, nativePriceI),
	})

	return types.ChainFeeComponents{ExecutionFee: execFee, DataAvailabilityFee: dataAvFee},
		ccipocr3.NewBigInt(nativePriceI),
		ccipocr3.NewBigInt(usdPricesF)
}
