package reader_test

import (
	"errors"
	"math/big"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/internal"
	mock_reader "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func Test_GetChainsFeeComponents(t *testing.T) {
	t.Cleanup(func() { reader.PromChainFeeGauge.Reset() })

	ctx := tests.Context(t)
	chain1 := "2337"
	selector1 := cciptypes.ChainSelector(12922642891491394802)
	chain2 := "3337"
	selector2 := cciptypes.ChainSelector(4793464827907405086)

	fees := map[cciptypes.ChainSelector]types.ChainFeeComponents{
		selector1: {
			ExecutionFee:        big.NewInt(10),
			DataAvailabilityFee: big.NewInt(15),
		},
		selector2: {
			ExecutionFee:        nil,
			DataAvailabilityFee: big.NewInt(2),
		},
	}

	origin := mock_reader.NewMockCCIPReader(t)
	r := reader.NewObservedCCIPReader(origin, logger.Test(t), selector1)

	origin.EXPECT().
		GetChainsFeeComponents(mock.Anything, mock.Anything).
		Return(fees)

	r.GetChainsFeeComponents(ctx, []cciptypes.ChainSelector{})

	require.Equal(
		t,
		10.0,
		testutil.ToFloat64(reader.PromChainFeeGauge.WithLabelValues(chain1, "execCost")),
	)
	require.Equal(
		t,
		15.0,
		testutil.ToFloat64(reader.PromChainFeeGauge.WithLabelValues(chain1, "daCost")),
	)
	require.Equal(
		t,
		0.0,
		testutil.ToFloat64(reader.PromChainFeeGauge.WithLabelValues(chain2, "execCost")),
	)
	require.Equal(
		t,
		2.0,
		testutil.ToFloat64(reader.PromChainFeeGauge.WithLabelValues(chain2, "daCost")),
	)
}

func Test_GetDestChainFeeComponents(t *testing.T) {
	ctx := tests.Context(t)
	chainID := "2337"
	chainSelector := cciptypes.ChainSelector(12922642891491394802)

	tt := []struct {
		name          string
		feeComponents types.ChainFeeComponents
		err           error
		expExec       float64
		expDa         float64
	}{
		{
			name: "fee components are reported",
			feeComponents: types.ChainFeeComponents{
				ExecutionFee:        big.NewInt(10),
				DataAvailabilityFee: big.NewInt(15),
			},
			expExec: 10,
			expDa:   15,
		},
		{
			name: "single missing components is ignored",
			feeComponents: types.ChainFeeComponents{
				ExecutionFee:        big.NewInt(20),
				DataAvailabilityFee: nil,
			},
			expExec: 20,
			expDa:   0,
		},
		{
			name: "missing components are ignored",
			feeComponents: types.ChainFeeComponents{
				ExecutionFee:        nil,
				DataAvailabilityFee: nil,
			},
			expExec: 0,
			expDa:   0,
		},
		{
			name:    "error doesn't report",
			err:     errors.New("something went wrong"),
			expExec: 0,
			expDa:   0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			t.Cleanup(func() { reader.PromChainFeeGauge.Reset() })

			origin := mock_reader.NewMockCCIPReader(t)
			r := reader.NewObservedCCIPReader(origin, logger.Test(t), chainSelector)

			origin.EXPECT().
				GetDestChainFeeComponents(ctx).
				Return(tc.feeComponents, tc.err)

			_, err := r.GetDestChainFeeComponents(ctx)
			if tc.err != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			require.Equal(
				t,
				tc.expExec,
				testutil.ToFloat64(reader.PromChainFeeGauge.WithLabelValues(chainID, "execCost")),
			)
			require.Equal(
				t,
				tc.expDa,
				testutil.ToFloat64(reader.PromChainFeeGauge.WithLabelValues(chainID, "daCost")),
			)
		})
	}
}

func Test_CommitReportsGTETimestamp(t *testing.T) {
	ctx := tests.Context(t)
	chainID := "2337"
	chainSelector := cciptypes.ChainSelector(12922642891491394802)

	t.Cleanup(cleanupMetrics())

	tt := []struct {
		name          string
		result        []plugintypes.CommitPluginReportWithMeta
		err           error
		expectedCount int
	}{
		{
			name:          "nil reports",
			result:        nil,
			err:           nil,
			expectedCount: 0,
		},
		{
			name:          "empty reports",
			result:        []plugintypes.CommitPluginReportWithMeta{},
			err:           nil,
			expectedCount: 0,
		},
		{
			name: "reports with some items",
			result: []plugintypes.CommitPluginReportWithMeta{
				{
					Timestamp: time.Now(),
				},
				{
					Timestamp: time.Now(),
				},
			},
			err:           nil,
			expectedCount: 2,
		},
		{
			name:          "error case",
			result:        nil,
			err:           errors.New("some error occurred"),
			expectedCount: 0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			cleanupMetrics()()

			origin := mock_reader.NewMockCCIPReader(t)
			r := reader.NewObservedCCIPReader(origin, logger.Test(t), chainSelector)

			origin.EXPECT().
				CommitReportsGTETimestamp(ctx, mock.Anything, mock.Anything, mock.Anything).
				Return(tc.result, tc.err)

			result, err := r.CommitReportsGTETimestamp(ctx, time.Now(), primitives.Unconfirmed, 1)
			require.Equal(t, tc.err, err)
			require.Equal(t, tc.result, result)
			require.Equal(
				t,
				float64(tc.expectedCount),
				testutil.ToFloat64(reader.PromDataSetSizeGauge.WithLabelValues(chainID, "CommitReportsGTETimestamp")),
			)

			count := internal.CounterFromHistogramByLabels(t, reader.PromQueryHistogram, chainID, "CommitReportsGTETimestamp")
			require.Equal(t, 1, count)
		})
	}

	//
	//origin.EXPECT().
	//	CommitReportsGTETimestamp(ctx, mock.Anything, mock.Anything, 2).
	//	Return([]plugintypes.CommitPluginReportWithMeta{}, nil)
	//
	//origin.EXPECT().
	//	CommitReportsGTETimestamp(ctx, mock.Anything, mock.Anything, 3).
	//	Return([]plugintypes.CommitPluginReportWithMeta{}, nil)
	//
	//origin.EXPECT().
	//	CommitReportsGTETimestamp(ctx, mock.Anything, mock.Anything, 4).
	//	Return(nil, fmt.Errorf("error"))

	//
	//reports2, err := r.CommitReportsGTETimestamp(ctx, time.Now(), primitives.Unconfirmed, 2)
}

func cleanupMetrics() func() {
	return func() {
		reader.PromChainFeeGauge.Reset()
		reader.PromDataSetSizeGauge.Reset()
		reader.PromQueryHistogram.Reset()
	}
}
