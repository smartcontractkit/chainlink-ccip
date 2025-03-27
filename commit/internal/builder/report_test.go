package builder

import (
	"fmt"
	unsaferand "math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	ccipocr3mocks "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/types/ccipocr3"
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

func Test_buildOneReport(t *testing.T) {
	ctx := t.Context()
	lggr := logger.Test(t)

	transmissionSchedule := &ocr3types.TransmissionSchedule{}

	blessedMerkleRoots := []ccipocr3.MerkleRootChain{
		{ChainSel: 1, MerkleRoot: mustMakeBytes("0x0102030405060708090102030405060708090102030405060708090102030405")},
	}
	unblessedMerkleRoots := []ccipocr3.MerkleRootChain{
		{ChainSel: 2, MerkleRoot: mustMakeBytes("0x0202030405060708090102030405060708090102030405060708090102030405")},
	}
	rmnSignatures := []ccipocr3.RMNECDSASignature{
		{
			R: [32]byte{0x01},
			S: [32]byte{0x02},
		},
		{
			R: [32]byte{0x02},
			S: [32]byte{0x03},
		},
	}
	priceUpdates := ccipocr3.PriceUpdates{
		TokenPriceUpdates: []ccipocr3.TokenPrice{
			{TokenID: "ETH", Price: ccipocr3.NewBigIntFromInt64(1234)},
		},
		GasPriceUpdates: []ccipocr3.GasPriceChain{
			{ChainSel: 1, GasPrice: ccipocr3.NewBigIntFromInt64(4567)},
		},
	}

	tests := []struct {
		name              string
		merkleOutcomeType merkleroot.OutcomeType
		blessedRoots      []ccipocr3.MerkleRootChain
		unblessedRoots    []ccipocr3.MerkleRootChain
		rmnSignatures     []ccipocr3.RMNECDSASignature
		rmnRemoteFSign    uint64
		priceUpdates      ccipocr3.PriceUpdates
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
func mustMakeBytes(byteStr string) ccipocr3.Bytes32 {
	if byteStr == "" {
		var randomBytes ccipocr3.Bytes32
		n, err := unsaferand.New(unsaferand.NewSource(0)).Read(randomBytes[:])
		if n != 32 {
			panic(fmt.Sprintf("Unexpected number of bytes read for placeholder id: want 32, got %d", n))
		}
		if err != nil {
			panic(fmt.Sprintf("Error reading random bytes: %v", err))
		}
		return randomBytes
	}
	b, err := ccipocr3.NewBytes32FromString(byteStr)
	if err != nil {
		panic(err)
	}
	return b
}
