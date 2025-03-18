package builder

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-ccip/commit/chainfee"
	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/commit/tokenprice"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

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
		reportBuilder   ReportBuilderFunc
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
