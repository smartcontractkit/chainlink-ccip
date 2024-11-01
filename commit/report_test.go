package commit

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
)

func TestPluginReports(t *testing.T) {
	testCases := []struct {
		name          string
		outc          Outcome
		expErr        bool
		expReports    []ccipocr3.CommitPluginReport
		expReportInfo ReportInfo
	}{
		{
			name: "wrong outcome type gives an empty report but no error",
			outc: Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportIntervalsSelected,
				},
			},
			expErr: false,
		},
		{
			name: "correct outcome type but empty data",
			outc: Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportGenerated,
				},
			},
			expErr: false,
		},
		{
			name: "token prices reported but no merkle roots so report is empty",
			outc: Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportTransmissionFailed,
				},
				TokenPriceOutcome: tokenprice.Outcome{
					TokenPrices: []ccipocr3.TokenPrice{
						{TokenID: "a", Price: ccipocr3.NewBigIntFromInt64(123)},
					},
				},
				ChainFeeOutcome: chainfee.Outcome{
					GasPrices: []ccipocr3.GasPriceChain{
						{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
					},
				},
			},
			expErr: false,
		},
		{
			name: "token prices reported but no merkle roots so report is empty",
			outc: Outcome{
				MerkleRootOutcome: merkleroot.Outcome{
					OutcomeType: merkleroot.ReportGenerated,
					RootsToReport: []ccipocr3.MerkleRootChain{
						{
							ChainSel:      3,
							OnRampAddress: []byte{1, 2, 3},
							SeqNumsRange:  ccipocr3.NewSeqNumRange(10, 20),
							MerkleRoot:    ccipocr3.Bytes32{1, 2, 3, 4, 5, 6},
						},
					},
					RMNRemoteCfg: rmntypes.RemoteConfig{F: 123},
				},
				TokenPriceOutcome: tokenprice.Outcome{
					TokenPrices: []ccipocr3.TokenPrice{
						{TokenID: "a", Price: ccipocr3.NewBigIntFromInt64(123)},
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
					MerkleRoots: []ccipocr3.MerkleRootChain{
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
					RMNSignatures: nil,
				},
			},
			expReportInfo: ReportInfo{RemoteF: 123},
		},
	}

	ctx := tests.Context(t)
	lggr := logger.Test(t)
	reportCodec := mocks.NewCommitPluginJSONReportCodec()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			p := Plugin{
				lggr:        lggr,
				reportCodec: reportCodec,
			}

			outcomeBytes, err := tc.outc.Encode()

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
	p := Plugin{lggr: lggr}
	_, err := p.Reports(tests.Context(t), 0, []byte("invalid json"))
	require.Error(t, err)
}
