package commit

import (
	"context"
	"encoding/json"
	rand2 "math/rand"
	"testing"
	"time"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func TestPluginReports(t *testing.T) {
	testCases := []struct {
		name          string
		outc          committypes.Outcome
		expErr        bool
		expReports    []ccipocr3.CommitPluginReport
		expReportInfo ccipocr3.CommitReportInfo
	}{
		{
			name: "wrong outcome type gives an empty report but no error",
			outc: committypes.Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportIntervalsSelected,
				},
			},
			expErr: false,
		},
		{
			name: "correct outcome type but empty data",
			outc: committypes.Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportGenerated,
				},
			},
			expErr: false,
		},
		{
			name: "token prices reported without merkle root is still transmitted",
			outc: committypes.Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportTransmissionFailed,
				},
				TokenPriceOutcome: tokenprice.Outcome{
					TokenPrices: ccipocr3.TokenPriceMap{
						"a": ccipocr3.NewBigIntFromInt64(123),
					},
				},
				ChainFeeOutcome: chainfee.Outcome{
					GasPrices: []ccipocr3.GasPriceChain{
						{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
					},
				},
			},
			expReports: []ccipocr3.CommitPluginReport{
				{
					PriceUpdates: ccipocr3.PriceUpdates{
						TokenPriceUpdates: []ccipocr3.TokenPrice{
							{TokenID: "a", Price: ccipocr3.NewBigIntFromInt64(123)},
						},
						GasPriceUpdates: []ccipocr3.GasPriceChain{
							{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
						},
					},
					RMNSignatures:        nil,
					UnblessedMerkleRoots: make([]ccipocr3.MerkleRootChain, 0),
					BlessedMerkleRoots:   make([]ccipocr3.MerkleRootChain, 0),
				},
			},
			expReportInfo: ccipocr3.CommitReportInfo{
				PriceUpdates: ccipocr3.PriceUpdates{
					TokenPriceUpdates: []ccipocr3.TokenPrice{
						{TokenID: "a", Price: ccipocr3.NewBigIntFromInt64(123)},
					},
					GasPriceUpdates: []ccipocr3.GasPriceChain{
						{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
					},
				},
			},
			expErr: false,
		},
		{
			name: "only chain fee reported without merkle root is still transmitted",
			outc: committypes.Outcome{
				ChainFeeOutcome: chainfee.Outcome{
					GasPrices: []ccipocr3.GasPriceChain{
						{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
					},
				},
			},
			expReports: []ccipocr3.CommitPluginReport{
				{
					PriceUpdates: ccipocr3.PriceUpdates{
						GasPriceUpdates: []ccipocr3.GasPriceChain{
							{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
						},
					},
					UnblessedMerkleRoots: make([]ccipocr3.MerkleRootChain, 0),
					BlessedMerkleRoots:   make([]ccipocr3.MerkleRootChain, 0),
					RMNSignatures:        nil,
				},
			},
			expReportInfo: ccipocr3.CommitReportInfo{
				PriceUpdates: ccipocr3.PriceUpdates{
					GasPriceUpdates: []ccipocr3.GasPriceChain{
						{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
					},
				},
			},
			expErr: false,
		},
		{
			name: "token prices reported but no merkle roots so report is not empty",
			outc: committypes.Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportGenerated,
					RootsToReport: []ccipocr3.MerkleRootChain{
						{
							ChainSel:      3,
							OnRampAddress: []byte{1, 2, 3},
							SeqNumsRange:  ccipocr3.NewSeqNumRange(10, 20),
							MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6},
						},
						{
							ChainSel:      2,
							OnRampAddress: []byte{1, 2, 3},
							SeqNumsRange:  ccipocr3.NewSeqNumRange(110, 210),
							MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6, 7},
						},
					},
					RMNRemoteCfg:     rmntypes.RemoteConfig{FSign: 123},
					RMNEnabledChains: map[ccipocr3.ChainSelector]bool{3: true, 2: false},
				},
				TokenPriceOutcome: tokenprice.Outcome{
					TokenPrices: ccipocr3.TokenPriceMap{
						"a": ccipocr3.NewBigIntFromInt64(123),
					},
				},
				ChainFeeOutcome: chainfee.Outcome{
					GasPrices: []ccipocr3.GasPriceChain{
						{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
					},
				},
			},
			expReports: []ccipocr3.CommitPluginReport{
				{
					BlessedMerkleRoots: []ccipocr3.MerkleRootChain{
						{
							ChainSel:      3,
							OnRampAddress: []byte{1, 2, 3},
							SeqNumsRange:  ccipocr3.NewSeqNumRange(10, 20),
							MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6},
						},
					},
					UnblessedMerkleRoots: []ccipocr3.MerkleRootChain{
						{
							ChainSel:      2,
							OnRampAddress: []byte{1, 2, 3},
							SeqNumsRange:  ccipocr3.NewSeqNumRange(110, 210),
							MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6, 7},
						},
					},
					PriceUpdates: ccipocr3.PriceUpdates{
						TokenPriceUpdates: []ccipocr3.TokenPrice{
							{TokenID: "a", Price: ccipocr3.NewBigIntFromInt64(123)},
						},
						GasPriceUpdates: []ccipocr3.GasPriceChain{
							{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
						},
					},
					RMNSignatures: nil,
				},
			},
			expReportInfo: ccipocr3.CommitReportInfo{
				RemoteF: 123,
				MerkleRoots: []ccipocr3.MerkleRootChain{
					{
						ChainSel:      2,
						OnRampAddress: []byte{1, 2, 3},
						SeqNumsRange:  ccipocr3.NewSeqNumRange(110, 210),
						MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6, 7},
					},
					{
						ChainSel:      3,
						OnRampAddress: []byte{1, 2, 3},
						SeqNumsRange:  ccipocr3.NewSeqNumRange(10, 20),
						MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6},
					},
				},
				PriceUpdates: ccipocr3.PriceUpdates{
					TokenPriceUpdates: []ccipocr3.TokenPrice{
						{TokenID: "a", Price: ccipocr3.NewBigIntFromInt64(123)},
					},
					GasPriceUpdates: []ccipocr3.GasPriceChain{
						{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
					},
				},
			},
		},
	}

	ctx := tests.Context(t)
	lggr := logger.Test(t)
	reportCodec := mocks.NewCommitPluginJSONReportCodec()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ocrTypeCodec := ocrtypecodec.DefaultCommitCodec

			cs := plugincommon.NewMockChainSupport(t)
			p := Plugin{
				lggr:            lggr,
				reportCodec:     reportCodec,
				ocrTypeCodec:    ocrTypeCodec,
				oracleIDToP2PID: map[commontypes.OracleID]libocrtypes.PeerID{1: {1}},
				chainSupport:    cs,
				reportBuilder:   buildStandardReport,
			}
			cs.EXPECT().SupportsDestChain(commontypes.OracleID(1)).Return(true, nil).Maybe()

			outcomeBytes, err := ocrTypeCodec.EncodeOutcome(tc.outc)
			require.NoError(t, err)

			reports, err := p.Reports(ctx, 0, outcomeBytes)
			if tc.expErr {
				require.Error(t, err)
				return
			}

			require.NoError(t, err)
			require.Len(t, reports, len(tc.expReports))
			for i, report := range reports {
				expEncodedReport, err := reportCodec.Encode(ctx, tc.expReports[i])
				require.NoError(t, err)
				require.Equal(t, expEncodedReport, []byte(report.ReportWithInfo.Report))

				expReportInfoBytes, err := tc.expReportInfo.Encode()
				require.NoError(t, err)
				require.Equal(t, expReportInfoBytes, report.ReportWithInfo.Info)
			}
		})
	}
}

func TestPluginReports_InvalidOutcome(t *testing.T) {
	lggr := logger.Test(t)
	p := Plugin{
		lggr:         lggr,
		ocrTypeCodec: ocrtypecodec.DefaultCommitCodec,
	}
	_, err := p.Reports(tests.Context(t), 0, []byte("invalid json"))
	require.Error(t, err)
}

func Test_Plugin_isStaleReport(t *testing.T) {
	testCases := []struct {
		name           string
		onChainSeqNum  uint64
		reportSeqNum   uint64
		lenMerkleRoots int
		shouldBeStale  bool
	}{
		{
			name:           "report is not stale when merkle roots exist no matter the seq nums",
			onChainSeqNum:  rand2.Uint64(),
			reportSeqNum:   rand2.Uint64(),
			lenMerkleRoots: 1,
			shouldBeStale:  false,
		},
		{
			name:           "report is stale when onChainSeqNum is equal to report seq num",
			onChainSeqNum:  33,
			reportSeqNum:   33,
			lenMerkleRoots: 0,
			shouldBeStale:  true,
		},
		{
			name:           "report is stale when onChainSeqNum is greater than report seq num",
			onChainSeqNum:  34,
			reportSeqNum:   33,
			lenMerkleRoots: 0,
			shouldBeStale:  true,
		},
		{
			name:           "report is not stale when onChainSeqNum is less than report seq num",
			onChainSeqNum:  32,
			reportSeqNum:   33,
			lenMerkleRoots: 0,
			shouldBeStale:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := Plugin{
				lggr: logger.Test(t),
			}
			report := ccipocr3.CommitPluginReport{
				BlessedMerkleRoots: make([]ccipocr3.MerkleRootChain, tc.lenMerkleRoots),
			}
			stale := p.isStaleReport(p.lggr, tc.reportSeqNum, tc.onChainSeqNum, report)
			require.Equal(t, tc.shouldBeStale, stale)
		})
	}
}

func TestMultiReportBuilders(t *testing.T) {
	outcome := committypes.Outcome{
		MerkleRootOutcome: merkleroot.Outcome{
			OutcomeType: merkleroot.ReportGenerated,
			RootsToReport: []ccipocr3.MerkleRootChain{
				{
					ChainSel:      2,
					OnRampAddress: []byte{1, 2, 3},
					SeqNumsRange:  ccipocr3.NewSeqNumRange(10, 20),
					MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6},
				},
				{ // this one is blessed.
					ChainSel:      3,
					OnRampAddress: []byte{1, 2, 3},
					SeqNumsRange:  ccipocr3.NewSeqNumRange(110, 210),
					MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6, 7},
				},
				{
					ChainSel:      4,
					OnRampAddress: []byte{1, 2, 3},
					SeqNumsRange:  ccipocr3.NewSeqNumRange(110, 210),
					MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6, 7},
				},
				{
					ChainSel:      5,
					OnRampAddress: []byte{1, 2, 3},
					SeqNumsRange:  ccipocr3.NewSeqNumRange(110, 210),
					MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6, 7},
				},
			},
			RMNRemoteCfg:     rmntypes.RemoteConfig{FSign: 123},
			RMNEnabledChains: map[ccipocr3.ChainSelector]bool{3: true, 2: false},
		},
		TokenPriceOutcome: tokenprice.Outcome{
			TokenPrices: ccipocr3.TokenPriceMap{
				"a": ccipocr3.NewBigIntFromInt64(123),
			},
		},
		ChainFeeOutcome: chainfee.Outcome{
			GasPrices: []ccipocr3.GasPriceChain{
				{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
			},
		},
	}

	testcases := []struct {
		name            string
		reportBuilder   reportBuilderFunc
		maxRoots        uint64
		expectedReports int
		checkReport     func(t *testing.T, i int, report ccipocr3.CommitPluginReport)
	}{
		{
			name:            "standard report builder",
			reportBuilder:   buildStandardReport,
			maxRoots:        1,
			expectedReports: 1,
			checkReport: func(t *testing.T, i int, report ccipocr3.CommitPluginReport) {
				// only one report, it contains all price updates.
				priceUpdates := ccipocr3.PriceUpdates{
					TokenPriceUpdates: outcome.TokenPriceOutcome.TokenPrices.ToSortedSlice(),
					GasPriceUpdates:   outcome.ChainFeeOutcome.GasPrices,
				}
				require.Equal(t, report.PriceUpdates, priceUpdates)

				roots := append(report.BlessedMerkleRoots, report.UnblessedMerkleRoots...)
				require.Len(t, roots, 4)
			},
		},
		{
			name:            "multi report builder",
			reportBuilder:   buildMultipleReports,
			maxRoots:        1,
			expectedReports: 4,
			checkReport: func(t *testing.T, i int, report ccipocr3.CommitPluginReport) {
				// only the first report contains price updates.
				if i > 0 {
					require.Equal(t, report.PriceUpdates, ccipocr3.PriceUpdates{})
				} else {
					priceUpdates := ccipocr3.PriceUpdates{
						TokenPriceUpdates: outcome.TokenPriceOutcome.TokenPrices.ToSortedSlice(),
						GasPriceUpdates:   outcome.ChainFeeOutcome.GasPrices,
					}
					require.Equal(t, report.PriceUpdates, priceUpdates)
				}

				roots := append(report.BlessedMerkleRoots, report.UnblessedMerkleRoots...)
				require.Len(t, roots, 1)
			},
		},
		{
			name:            "truncated report builder",
			reportBuilder:   buildTruncatedReport,
			maxRoots:        1,
			expectedReports: 1,
			checkReport: func(t *testing.T, i int, report ccipocr3.CommitPluginReport) {
				// single report contains all price updates.
				priceUpdates := ccipocr3.PriceUpdates{
					TokenPriceUpdates: outcome.TokenPriceOutcome.TokenPrices.ToSortedSlice(),
					GasPriceUpdates:   outcome.ChainFeeOutcome.GasPrices,
				}
				require.Equal(t, report.PriceUpdates, priceUpdates)
				roots := append(report.BlessedMerkleRoots, report.UnblessedMerkleRoots...)
				require.Len(t, roots, 1)
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tests.Context(t)
			lggr := logger.Test(t)
			reportCodec := mocks.NewCommitPluginJSONReportCodec()
			ts := &ocr3types.TransmissionSchedule{
				Transmitters:       nil,
				TransmissionDelays: nil,
			}

			reports, err := tc.reportBuilder(ctx, lggr, reportCodec, ts, outcome, tc.maxRoots)
			require.NoError(t, err)
			require.Len(t, reports, tc.expectedReports)

			for i, report := range reports {
				r, err := reportCodec.Decode(ctx, report.ReportWithInfo.Report)
				require.NoError(t, err)
				tc.checkReport(t, i, r)
			}
		})
	}
}

func BenchmarkReports(b *testing.B) {
	oracleID := commontypes.OracleID(1)
	mockChainSupport := plugincommon.NewMockChainSupport(b)
	mockChainSupport.EXPECT().SupportsDestChain(mock.Anything).Return(true, nil)
	oracleIDToP2PID := map[commontypes.OracleID]libocrtypes.PeerID{0: {0}, oracleID: {1}, 2: {2}, 3: {3}}

	plugin := &Plugin{
		oracleID:        oracleID,
		lggr:            logger.Test(b),
		ocrTypeCodec:    ocrtypecodec.DefaultCommitCodec,
		chainSupport:    mockChainSupport,
		oracleIDToP2PID: oracleIDToP2PID,
		offchainCfg: pluginconfig.CommitOffchainConfig{
			TransmissionDelayMultiplier: 30 * time.Second,
			MaxMerkleRootsPerReport:     0,
		},
		reportBuilder: buildStandardReport,
		reportCodec:   mocks.NewCommitPluginJSONReportCodec(),
	}

	// Mock input data
	ctx := context.Background()
	seqNr := uint64(1)
	outcomeJSON := `{
    "merkleRootOutcome": {
      "outcomeType": 2,
      "rangesSelectedForReport": null,
      "rootsToReport": [
        {
          "chain": 909606746561742100,
          "onRampAddress": "0x00000000000000000000000078c2e4a7ef593bf4759f5401ee22556e42a1b334",
          "seqNumsRange": [
            1,
            256
          ],
          "merkleRoot": "0x1da9da57091881823f3a3791b5ba351bcd3012b670e499d46ef4b069754ff808"
        }
      ],
      "rmnEnabledChains": {},
      "offRampNextSeqNums": [
        {
          "chainSel": 909606746561742100,
          "seqNum": 1
        },
        {
          "chainSel": 5548718428018410000,
          "seqNum": 1
        }
      ],
      "reportTransmissionCheckAttempts": 0,
      "rmnReportSignatures": [],
      "rmnRemoteCfg": {
        "contractAddress": "0xd59800ca659eca2b8d4dc0d2a9f4c929951e6e89",
        "configDigest": "0x000b6accb0239044979ae0df2c58e36ab26d31da43997dcb10c38f47215437e8",
        "signers": [
          {
            "onchainPublicKey": "0x0100000000000000000000000000000000000000",
            "nodeIndex": 0
          }
        ],
        "fSign": 0,
        "configVersion": 1,
        "rmnReportVersion": "0x9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf53"
      }
    },
    "tokenPriceOutcome": {
      "tokenPrices": {
        "0x5050E3E132aE4c03cAD5fF0747fd210544366caF": "9000000000000000000",
        "0x72867881ee6AA5b4fc57adD6AE02CA9dB24Afc07": "500000000000000000000"
      }
    },
    "chainFeeOutcome": {
      "gasPrices": [
        {
          "chainSel": 909606746561742100,
          "gasPrice": "80000000000000"
        }
      ]
    },
    "mainOutcome": {
      "inflightPriceOcrSequenceNumber": 3,
      "remainingPriceChecks": 5
    }
  }`
	var outcome committypes.Outcome
	err := json.Unmarshal([]byte(outcomeJSON), &outcome)
	require.NoError(b, err)
	outcomeBytes, err := plugin.ocrTypeCodec.EncodeOutcome(outcome)
	require.NoError(b, err)

	// Reset the timer to exclude setup time

	for b.Loop() {
		_, err := plugin.Reports(ctx, seqNr, outcomeBytes)
		if err != nil {
			b.Fatalf("Reports function failed: %v", err)
		}
	}
}
