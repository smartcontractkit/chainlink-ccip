package reader_test

import (
	"errors"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/testhelpers/rand"
	mock_reader "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
)

func Test_GetChainsFeeComponents(t *testing.T) {
	t.Cleanup(func() { reader.PromChainFeeGauge.Reset() })

	ctx := t.Context()
	chain1 := "2337"
	selector1 := cciptypes.ChainSelector(12922642891491394802)
	chain2 := "5eykt4UsFv8P8NJdTREpY1vzqKqZKvdpKuc147dw2N9d"
	selector2 := cciptypes.ChainSelector(124615329519749607)

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
		testutil.ToFloat64(reader.PromChainFeeGauge.WithLabelValues("evm", chain1, "execCost")),
	)
	require.Equal(
		t,
		15.0,
		testutil.ToFloat64(reader.PromChainFeeGauge.WithLabelValues("evm", chain1, "daCost")),
	)
	require.Equal(
		t,
		0.0,
		testutil.ToFloat64(reader.PromChainFeeGauge.WithLabelValues("solana", chain2, "execCost")),
	)
	require.Equal(
		t,
		2.0,
		testutil.ToFloat64(reader.PromChainFeeGauge.WithLabelValues("solana", chain2, "daCost")),
	)
}

func Test_GetDestChainFeeComponents(t *testing.T) {
	ctx := t.Context()
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
				testutil.ToFloat64(reader.PromChainFeeGauge.WithLabelValues("evm", chainID, "execCost")),
			)
			require.Equal(
				t,
				tc.expDa,
				testutil.ToFloat64(reader.PromChainFeeGauge.WithLabelValues("evm", chainID, "daCost")),
			)
		})
	}
}

func Test_CommitReportsGTETimestamp(t *testing.T) {
	ctx := t.Context()
	chainID := "2337"
	chainSelector := cciptypes.ChainSelector(12922642891491394802)

	t.Cleanup(cleanupMetrics())

	tt := []struct {
		name          string
		result        []cciptypes.CommitPluginReportWithMeta
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
			result:        []cciptypes.CommitPluginReportWithMeta{},
			err:           nil,
			expectedCount: 0,
		},
		{
			name: "reports with some items",
			result: []cciptypes.CommitPluginReportWithMeta{
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
				testutil.ToFloat64(reader.PromDataSetSizeGauge.WithLabelValues("evm", chainID, "CommitReportsGTETimestamp")),
			)

			count := internal.CounterFromHistogramByLabels(
				t, reader.PromQueryHistogram, "evm", chainID, "CommitReportsGTETimestamp",
			)
			require.Equal(t, 1, count)
		})
	}
}

func Test_LatestMsgSeqNum(t *testing.T) {
	ctx := t.Context()
	chainID := "1"
	chainSelector := cciptypes.ChainSelector(4741433654826277614)

	t.Cleanup(cleanupMetrics())

	origin := mock_reader.NewMockCCIPReader(t)
	r := reader.NewObservedCCIPReader(origin, logger.Test(t), chainSelector)

	seqNr := cciptypes.SeqNum(1234)
	success := cciptypes.ChainSelector(rand.RandomInt64())
	failure := cciptypes.ChainSelector(rand.RandomInt64())

	origin.EXPECT().
		LatestMsgSeqNum(ctx, success).
		Return(seqNr, nil)

	origin.EXPECT().
		LatestMsgSeqNum(ctx, failure).
		Return(seqNr, fmt.Errorf("erro"))

	result, err := r.LatestMsgSeqNum(ctx, success)
	require.NoError(t, err)
	require.Equal(t, seqNr, result)

	require.Equal(
		t,
		1,
		internal.CounterFromHistogramByLabels(t, reader.PromQueryHistogram, "aptos", chainID, "LatestMsgSeqNum"),
	)
	require.Equal(
		t,
		0.0,
		testutil.ToFloat64(reader.PromDataSetSizeGauge.WithLabelValues("aptos", chainID, "CommitReportsGTETimestamp")),
	)

	_, err = r.LatestMsgSeqNum(ctx, failure)
	require.Error(t, err)
	require.Equal(
		t,
		2,
		internal.CounterFromHistogramByLabels(t, reader.PromQueryHistogram, "aptos", chainID, "LatestMsgSeqNum"),
	)
	require.Equal(
		t,
		0.0,
		testutil.ToFloat64(reader.PromDataSetSizeGauge.WithLabelValues("aptos", chainID, "CommitReportsGTETimestamp")),
	)
}

func cleanupMetrics() func() {
	return func() {
		reader.PromChainFeeGauge.Reset()
		reader.PromDataSetSizeGauge.Reset()
		reader.PromQueryHistogram.Reset()
	}
}
