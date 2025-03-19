package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func TestReportBuilders(t *testing.T) {
	// This outcome is sliced in different ways depending on the config.
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
				"b": ccipocr3.NewBigIntFromInt64(123),
				"c": ccipocr3.NewBigIntFromInt64(123),
				"d": ccipocr3.NewBigIntFromInt64(123),
				"e": ccipocr3.NewBigIntFromInt64(123),
			},
		},
		ChainFeeOutcome: chainfee.Outcome{
			GasPrices: []ccipocr3.GasPriceChain{
				{GasPrice: ccipocr3.NewBigIntFromInt64(1), ChainSel: 123},
				{GasPrice: ccipocr3.NewBigIntFromInt64(2), ChainSel: 123},
				{GasPrice: ccipocr3.NewBigIntFromInt64(3), ChainSel: 123},
				{GasPrice: ccipocr3.NewBigIntFromInt64(4), ChainSel: 123},
				{GasPrice: ccipocr3.NewBigIntFromInt64(5), ChainSel: 123},
			},
		},
	}

	testcases := []struct {
		name            string
		reportBuilder   ReportBuilderFunc
		maxRoots        uint64
		maxPrices       uint64
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
			name:            "multi root report builder",
			reportBuilder:   buildMultipleMerkleRootReports,
			maxRoots:        1,
			expectedReports: 4,
			checkReport: func(t *testing.T, i int, report ccipocr3.CommitPluginReport) {
				// only the first report contains price updates.
				if i > 0 {
					assert.Equal(t, report.PriceUpdates, ccipocr3.PriceUpdates{})
				} else {
					// contains all price updates
					priceUpdates := ccipocr3.PriceUpdates{
						TokenPriceUpdates: outcome.TokenPriceOutcome.TokenPrices.ToSortedSlice(),
						GasPriceUpdates:   outcome.ChainFeeOutcome.GasPrices,
					}
					assert.Equal(t, report.PriceUpdates, priceUpdates)
				}

				roots := append(report.BlessedMerkleRoots, report.UnblessedMerkleRoots...)
				assert.Len(t, roots, 1)
			},
		},
		{
			name:            "multi price report builder",
			reportBuilder:   buildMultiplePriceAndMerkleRootReports,
			maxPrices:       3, // chosen to have one report with both gas and token prices.
			expectedReports: 5,
			checkReport: func(t *testing.T, i int, report ccipocr3.CommitPluginReport) {
				numRoots := len(report.BlessedMerkleRoots) + len(report.UnblessedMerkleRoots)
				numPrices := len(report.PriceUpdates.TokenPriceUpdates) + len(report.PriceUpdates.GasPriceUpdates)

				if i < 3 {
					assert.Equalf(t, 3, numPrices,
						"first 3 price reports should have maxPrices.")
					assert.Equalf(t, 0, numRoots, "There should be no roots in a price report.")
				} else if i < 4 {
					assert.Equalf(t, 1, numPrices,
						"final price report should have the remaining 1 gas price.")
					assert.Equalf(t, 0, numRoots, "There should be no roots in a price report.")
				} else {
					require.Equalf(t, 4, numRoots, "All roots are in the final report.")
					assert.Equalf(t, 0, numPrices, "There should be no prices in a root report.")
				}
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

			cfg := pluginconfig.CommitOffchainConfig{}
			require.NoError(t, cfg.ApplyDefaultsAndValidate())
			cfg.MaxMerkleRootsPerReport = tc.maxRoots
			cfg.MaxPricesPerReport = tc.maxPrices
			cfg.MultipleReportsEnabled = true
			reports, err := tc.reportBuilder(ctx, lggr, reportCodec, ts, outcome, cfg)
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
