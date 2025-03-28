package commit

import (
	"fmt"
	unsaferand "math/rand"
	"testing"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	ccipocr3mocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

func Test_buildOneReport(t *testing.T) {
	ctx := t.Context()
	lggr := logger.Test(t)

	transmissionSchedule := &ocr3types.TransmissionSchedule{}

	blessedMerkleRoots := []cciptypes.MerkleRootChain{
		{ChainSel: 1, MerkleRoot: mustMakeBytes("0x0102030405060708090102030405060708090102030405060708090102030405")},
	}
	unblessedMerkleRoots := []cciptypes.MerkleRootChain{
		{ChainSel: 2, MerkleRoot: mustMakeBytes("0x0202030405060708090102030405060708090102030405060708090102030405")},
	}
	rmnSignatures := []cciptypes.RMNECDSASignature{
		{
			R: [32]byte{0x01},
			S: [32]byte{0x02},
		},
		{
			R: [32]byte{0x02},
			S: [32]byte{0x03},
		},
	}
	priceUpdates := cciptypes.PriceUpdates{
		TokenPriceUpdates: []cciptypes.TokenPrice{
			{TokenID: "ETH", Price: cciptypes.NewBigIntFromInt64(1234)},
		},
		GasPriceUpdates: []cciptypes.GasPriceChain{
			{ChainSel: 1, GasPrice: cciptypes.NewBigIntFromInt64(4567)},
		},
	}

	tests := []struct {
		name              string
		merkleOutcomeType merkleroot.OutcomeType
		blessedRoots      []cciptypes.MerkleRootChain
		unblessedRoots    []cciptypes.MerkleRootChain
		rmnSignatures     []cciptypes.RMNECDSASignature
		rmnRemoteFSign    uint64
		priceUpdates      cciptypes.PriceUpdates
		mockCodecFn       func() *ccipocr3mocks.MockCommitPluginCodec
		wantErr           bool
		reportNil         bool
	}{
		{
			name:              "empty merkle root outcome, no prices either",
			merkleOutcomeType: merkleroot.ReportEmpty,
			mockCodecFn: func() *ccipocr3mocks.MockCommitPluginCodec {
				// no calls
				return ccipocr3mocks.NewMockCommitPluginCodec(t)
			},
			wantErr:   false,
			reportNil: true,
		},
		{
			name:              "empty merkle root outcome, with prices",
			merkleOutcomeType: merkleroot.ReportEmpty,
			priceUpdates:      priceUpdates,
			mockCodecFn: func() *ccipocr3mocks.MockCommitPluginCodec {
				m := ccipocr3mocks.NewMockCommitPluginCodec(t)
				m.EXPECT().Encode(mock.Anything, mock.MatchedBy(func(r ccipocr3.CommitPluginReport) bool {
					return len(r.PriceUpdates.TokenPriceUpdates) == 1 && len(r.PriceUpdates.GasPriceUpdates) == 1
				})).Once().Return([]byte("report"), nil)
				return m
			},
			wantErr:   false,
			reportNil: false,
		},
		{
			name:              "merkle outcome with blessed and unblessed roots, no price updates",
			merkleOutcomeType: merkleroot.ReportGenerated,
			blessedRoots:      blessedMerkleRoots,
			unblessedRoots:    unblessedMerkleRoots,
			rmnRemoteFSign:    1,
			rmnSignatures:     rmnSignatures,
			mockCodecFn: func() *ccipocr3mocks.MockCommitPluginCodec {
				m := ccipocr3mocks.NewMockCommitPluginCodec(t)
				m.EXPECT().Encode(mock.Anything, mock.MatchedBy(func(r ccipocr3.CommitPluginReport) bool {
					return len(r.BlessedMerkleRoots) == 1 && len(r.UnblessedMerkleRoots) == 1
				})).Once().Return([]byte("report"), nil)
				return m
			},
			wantErr:   false,
			reportNil: false,
		},
		{
			name:              "merkle outcome with blessed and unblessed roots, with price updates",
			merkleOutcomeType: merkleroot.ReportGenerated,
			blessedRoots:      blessedMerkleRoots,
			unblessedRoots:    unblessedMerkleRoots,
			priceUpdates:      priceUpdates,
			mockCodecFn: func() *ccipocr3mocks.MockCommitPluginCodec {
				m := ccipocr3mocks.NewMockCommitPluginCodec(t)
				m.EXPECT().Encode(mock.Anything, mock.MatchedBy(func(r ccipocr3.CommitPluginReport) bool {
					return len(r.BlessedMerkleRoots) == 1 && len(r.UnblessedMerkleRoots) == 1 &&
						len(r.PriceUpdates.TokenPriceUpdates) == 1 && len(r.PriceUpdates.GasPriceUpdates) == 1
				})).Once().Return([]byte("report"), nil)
				return m
			},
			wantErr:   false,
			reportNil: false,
		},
		{
			name:              "merkle outcome ReportInFlight, no price updates",
			merkleOutcomeType: merkleroot.ReportInFlight,

			// notice that blessed and unblessed roots are still set since they're
			// set in the merkle outcome.
			// However, they wouldn't be included in the report.
			blessedRoots:   blessedMerkleRoots,
			unblessedRoots: unblessedMerkleRoots,
			rmnRemoteFSign: 1,
			rmnSignatures:  rmnSignatures,
			mockCodecFn: func() *ccipocr3mocks.MockCommitPluginCodec {
				// no calls
				return ccipocr3mocks.NewMockCommitPluginCodec(t)
			},
			wantErr:   false,
			reportNil: true,
		},
		{
			name:              "merkle outcome ReportInFlight, with price updates",
			merkleOutcomeType: merkleroot.ReportInFlight,

			// notice that blessed and unblessed roots are still set since they're
			// set in the merkle outcome.
			// However, they wouldn't be included in the report.
			blessedRoots:   blessedMerkleRoots,
			unblessedRoots: unblessedMerkleRoots,
			rmnRemoteFSign: 1,
			rmnSignatures:  rmnSignatures,
			priceUpdates:   priceUpdates,
			mockCodecFn: func() *ccipocr3mocks.MockCommitPluginCodec {
				m := ccipocr3mocks.NewMockCommitPluginCodec(t)
				m.EXPECT().Encode(mock.Anything, mock.MatchedBy(func(r ccipocr3.CommitPluginReport) bool {
					return len(r.PriceUpdates.TokenPriceUpdates) == 1 && len(r.PriceUpdates.GasPriceUpdates) == 1
				})).Once().Return([]byte("report"), nil)
				return m
			},
			wantErr:   false,
			reportNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			report, err := buildOneReport(
				ctx,
				lggr,
				tt.mockCodecFn(),
				transmissionSchedule,
				tt.merkleOutcomeType,
				tt.blessedRoots,
				tt.unblessedRoots,
				tt.rmnSignatures,
				tt.rmnRemoteFSign,
				tt.priceUpdates,
			)

			if tt.wantErr {
				require.Error(t, err)
				require.Nil(t, report)
			} else {
				require.NoError(t, err)

				if tt.reportNil {
					require.Nil(t, report)
				} else {
					require.NotNil(t, report)
				}
			}
		})
	}
}

// mustMakeBytes parses a given string into a byte array, any error causes a panic. Pass in an empty string for a
// random byte array.
func mustMakeBytes(byteStr string) cciptypes.Bytes32 {
	if byteStr == "" {
		var randomBytes cciptypes.Bytes32
		n, err := unsaferand.New(unsaferand.NewSource(0)).Read(randomBytes[:])
		if n != 32 {
			panic(fmt.Sprintf("Unexpected number of bytes read for placeholder id: want 32, got %d", n))
		}
		if err != nil {
			panic(fmt.Sprintf("Error reading random bytes: %v", err))
		}
		return randomBytes
	}
	b, err := cciptypes.NewBytes32FromString(byteStr)
	if err != nil {
		panic(err)
	}
	return b
}
