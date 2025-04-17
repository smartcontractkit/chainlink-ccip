package commit

import (
	unsaferand "math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/internal/builder"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	ccipocr3mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/types/ccipocr3"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
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
					UnblessedMerkleRoots: nil,
					BlessedMerkleRoots:   nil,
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
					UnblessedMerkleRoots: nil,
					BlessedMerkleRoots:   nil,
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
					RMNRemoteCfg:     ccipocr3.RemoteConfig{FSign: 123},
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

	cfg := pluginconfig.CommitOffchainConfig{}
	err := cfg.ApplyDefaultsAndValidate()
	require.NoError(t, err)
	reportBuilder, err := builder.NewReportBuilder(cfg.RMNEnabled, cfg.MaxMerkleRootsPerReport, cfg.MaxPricesPerReport)
	require.NoError(t, err)

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
				reportBuilder:   reportBuilder,
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
				// Decoded reports should be equal (check first for a helpful error message)
				decodedReport, err := reportCodec.Decode(ctx, report.ReportWithInfo.Report)
				require.NoError(t, err)
				require.Equal(t, tc.expReports[i], decodedReport)

				// Encoded bytes should be equal
				encodedExpectedReport, err := reportCodec.Encode(ctx, tc.expReports[i])
				require.NoError(t, err)
				var reportBytes []byte = report.ReportWithInfo.Report
				require.Equal(t, encodedExpectedReport, reportBytes)

				// Decoded report info should be equal.
				decodedReportInfo, err := ccipocr3.DecodeCommitReportInfo(report.ReportWithInfo.Info)
				require.NoError(t, err)
				require.Equal(t, tc.expReportInfo, decodedReportInfo)

				// Encoded report info bytes should be equal.
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
			onChainSeqNum:  unsaferand.Uint64(),
			reportSeqNum:   unsaferand.Uint64(),
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

func Test_encodeReports(t *testing.T) {
	ctx := tests.Context(t)
	lggr := logger.Test(t)
	var transmissionSchedule *ocr3types.TransmissionSchedule

	var emptyReport builder.Report

	nonEmptyReport := builder.Report{
		Report: ccipocr3.CommitPluginReport{
			PriceUpdates: ccipocr3.PriceUpdates{
				GasPriceUpdates: []ccipocr3.GasPriceChain{
					{
						GasPrice: ccipocr3.NewBigIntFromInt64(3),
						ChainSel: 123,
					},
				},
			},
		},
	}

	type args struct {
		reports     []builder.Report
		reportCodec func() ccipocr3.CommitPluginCodec
	}
	testcases := []struct {
		name       string
		args       args
		numReports int
		wantErr    assert.ErrorAssertionFunc
	}{
		{
			name: "Happy path:one report with no errors",
			args: args{
				reports: []builder.Report{nonEmptyReport},
				reportCodec: func() ccipocr3.CommitPluginCodec {
					codec := ccipocr3mock.NewMockCommitPluginCodec(t)
					codec.EXPECT().Encode(mock.Anything, mock.Anything).Return([]byte{0x1}, nil)
					return codec
				},
			},
			numReports: 1,
			wantErr:    assert.NoError,
		},
		{
			name: "Happy path:multiple report with no errors",
			args: args{
				reports: []builder.Report{nonEmptyReport, nonEmptyReport},
				reportCodec: func() ccipocr3.CommitPluginCodec {
					codec := ccipocr3mock.NewMockCommitPluginCodec(t)
					codec.EXPECT().Encode(mock.Anything, mock.Anything).Return([]byte{0x1}, nil)
					return codec
				},
			},
			numReports: 2,
			wantErr:    assert.NoError,
		},
		{
			name: "Empty report error",
			args: args{
				reports: []builder.Report{emptyReport},
				reportCodec: func() ccipocr3.CommitPluginCodec {
					return ccipocr3mock.NewMockCommitPluginCodec(t)
				},
			},
			numReports: 0,
			wantErr:    assert.Error,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := encodeReports(ctx, lggr, tt.args.reports, transmissionSchedule, tt.args.reportCodec())
			if !tt.wantErr(t, err, "encodeReports(...)") {
				return
			}
			assert.Equal(t, tt.numReports, len(got))
		})
	}
}
