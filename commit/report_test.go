package commit

import (
	"fmt"
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	reader2 "github.com/smartcontractkit/chainlink-ccip/pkg/reader"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	libocrtypes "github.com/smartcontractkit/libocr/ragep2p/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/internal/builder"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	ccipocr3mock "github.com/smartcontractkit/chainlink-ccip/mocks/chainlink_common/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/plugincommon"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
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

	ctx := t.Context()
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
	_, err := p.Reports(t.Context(), 0, []byte("invalid json"))
	require.Error(t, err)
}

func Test_IsStaleReportMerkleRoots(t *testing.T) {
	sourceChainConfig := map[ccipocr3.ChainSelector]reader2.StaticSourceChainConfig{
		10: {IsRMNVerificationDisabled: false, IsEnabled: true},
		20: {IsRMNVerificationDisabled: false, IsEnabled: true},
		30: {IsRMNVerificationDisabled: true, IsEnabled: true},
	}

	testCases := []struct {
		name                 string
		onRampNextSeqNum     []plugintypes.SeqNumChain
		offRampExpNextSeqNum map[ccipocr3.ChainSelector]ccipocr3.SeqNum
		readerErr            error
		expErr               bool
	}{
		{
			name: "happy path",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: map[ccipocr3.ChainSelector]ccipocr3.SeqNum{10: 100, 20: 200},
			expErr:               false,
		},
		{
			name: "one root is stale",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			// <- 200 is already on chain
			offRampExpNextSeqNum: map[ccipocr3.ChainSelector]ccipocr3.SeqNum{10: 100, 20: 201},
			expErr:               true,
		},
		{
			name: "one root has gap",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 101), // <- onchain 99 but we submit 101 instead of 100
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: map[ccipocr3.ChainSelector]ccipocr3.SeqNum{10: 100, 20: 200},
			expErr:               true,
		},
		{
			name: "reader returned wrong number of seq nums, should be ok",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: map[ccipocr3.ChainSelector]ccipocr3.SeqNum{10: 100, 20: 200, 30: 300},
			expErr:               false,
		},
		{
			name: "reader error",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
			},
			offRampExpNextSeqNum: map[ccipocr3.ChainSelector]ccipocr3.SeqNum{10: 100, 20: 200},
			readerErr:            fmt.Errorf("reader error"),
			expErr:               true,
		},
		{
			name: "happy path one root unblessed",
			onRampNextSeqNum: []plugintypes.SeqNumChain{
				plugintypes.NewSeqNumChain(10, 100),
				plugintypes.NewSeqNumChain(20, 200),
				plugintypes.NewSeqNumChain(30, 300),
			},
			offRampExpNextSeqNum: map[ccipocr3.ChainSelector]ccipocr3.SeqNum{10: 100, 20: 200, 30: 300},
			expErr:               false,
		},
	}

	ctx := t.Context()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := readermock.NewMockCCIPReader(t)
			rep := ccipocr3.CommitPluginReport{}
			chains := make([]ccipocr3.ChainSelector, 0, len(tc.onRampNextSeqNum))
			for _, snc := range tc.onRampNextSeqNum {
				if sourceChainConfig[snc.ChainSel].IsRMNVerificationDisabled {
					rep.UnblessedMerkleRoots = append(rep.UnblessedMerkleRoots, ccipocr3.MerkleRootChain{
						ChainSel:     snc.ChainSel,
						SeqNumsRange: ccipocr3.NewSeqNumRange(snc.SeqNum, snc.SeqNum+10),
					})
				} else {
					rep.BlessedMerkleRoots = append(rep.BlessedMerkleRoots, ccipocr3.MerkleRootChain{
						ChainSel:     snc.ChainSel,
						SeqNumsRange: ccipocr3.NewSeqNumRange(snc.SeqNum, snc.SeqNum+10),
					})
				}
				chains = append(chains, snc.ChainSel)
			}
			reader.EXPECT().NextSeqNum(ctx, chains).Return(tc.offRampExpNextSeqNum, tc.readerErr)

			reader.EXPECT().GetOffRampSourceChainsConfig(ctx, chains).Return(sourceChainConfig, nil).Maybe()
			// on chain seq num is 2, larger than the `1` we send to isStaleReport
			reader.EXPECT().GetLatestPriceSeqNr(ctx).Return(2, nil)

			p := Plugin{
				lggr:       logger.Test(t),
				ccipReader: reader,
			}
			report := ccipocr3.CommitPluginReport{
				BlessedMerkleRoots:   rep.BlessedMerkleRoots,
				UnblessedMerkleRoots: rep.UnblessedMerkleRoots,
			}
			err := p.isStaleReport(ctx, 1, report)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
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

	ctx := t.Context()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			reader := readermock.NewMockCCIPReader(t)
			reader.EXPECT().GetLatestPriceSeqNr(mock.Anything).Return(tc.onChainSeqNum, nil)

			reader.EXPECT().NextSeqNum(ctx, mock.Anything).Return(
				map[ccipocr3.ChainSelector]ccipocr3.SeqNum{}, nil).Maybe()

			reader.EXPECT().GetOffRampSourceChainsConfig(ctx, mock.Anything).Return(
				map[ccipocr3.ChainSelector]reader2.StaticSourceChainConfig{}, nil).Maybe()

			p := Plugin{
				lggr:       logger.Test(t),
				ccipReader: reader,
			}
			report := ccipocr3.CommitPluginReport{
				BlessedMerkleRoots: make([]ccipocr3.MerkleRootChain, tc.lenMerkleRoots),
			}
			err := p.isStaleReport(ctx, tc.reportSeqNum, report)
			if tc.shouldBeStale {
				require.Error(t, err, "Expected report to be err but it was not")
			} else {
				require.NoError(t, err, "Expected report to be err but it was not")
			}
		})
	}
}

func Test_encodeReports(t *testing.T) {
	ctx := t.Context()
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
