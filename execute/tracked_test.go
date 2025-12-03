package execute

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/metrics"
	"github.com/smartcontractkit/chainlink-ccip/internal"
	ocrtypecodec "github.com/smartcontractkit/chainlink-ccip/pkg/ocrtypecodec/v1"
)

const (
	chainID  = "2"
	selector = cciptypes.ChainSelector(743186221051783445)
)

func Test_LatencyIsTracked(t *testing.T) {
	t.Cleanup(func() {
		metrics.PromExecErrors.Reset()
		metrics.PromExecLatencyHistogram.Reset()
	})

	var b strings.Builder
	bhClient, _ := beholder.NewWriterClient(&b)
	query := types.Query([]byte("query"))
	observation := types.Observation([]byte("observation"))
	outcome := ocr3types.Outcome([]byte("outcome"))
	ctx := t.Context()
	lggr := logger.Test(t)
	origin := FakePlugin{
		query:       query,
		observation: observation,
		outcome:     outcome,
	}
	reporter, err := metrics.NewPromReporter(lggr, selector, *bhClient)
	require.NoError(t, err)
	tracked := NewTrackedPlugin(origin, lggr, reporter, ocrtypecodec.DefaultExecCodec)

	count := 100
	for i := 0; i < count; i++ {
		q, err := tracked.Query(ctx, ocr3types.OutcomeContext{})
		require.Equal(t, query, q)
		require.NoError(t, err)

		obs, err := tracked.Observation(ctx, ocr3types.OutcomeContext{}, types.Query{})
		require.Equal(t, observation, obs)
		require.NoError(t, err)

		out, err := tracked.Outcome(ctx, ocr3types.OutcomeContext{}, types.Query{}, []types.AttributedObservation{})
		require.Equal(t, outcome, out)
		require.NoError(t, err)
	}

	l1 := internal.CounterFromHistogramByLabels(
		t, metrics.PromExecLatencyHistogram, "aptos", chainID, "observation", "GetCommitReports",
	)
	require.Equal(t, count, l1)

	l2 := internal.CounterFromHistogramByLabels(
		t, metrics.PromExecLatencyHistogram, "aptos", chainID, "outcome", "GetCommitReports",
	)
	require.Equal(t, count, l2)

	l3 := internal.CounterFromHistogramByLabels(
		t, metrics.PromExecLatencyHistogram, "aptos", chainID, "query", "GetCommitReports",
	)
	require.Equal(t, count, l3)
}

func Test_ErrorIsTrackedWhenOriginReturns(t *testing.T) {
	t.Cleanup(func() {
		metrics.PromExecErrors.Reset()
		metrics.PromExecLatencyHistogram.Reset()
	})
	var b strings.Builder
	bhClient, _ := beholder.NewWriterClient(&b)
	lggr := logger.Test(t)
	origin := FakePlugin{err: fmt.Errorf("error")}
	reporter, err := metrics.NewPromReporter(lggr, selector, *bhClient)
	require.NoError(t, err)
	tracked := NewTrackedPlugin(origin, lggr, reporter, ocrtypecodec.DefaultExecCodec)

	count := 100
	for i := 0; i < count; i++ {
		_, err = tracked.Outcome(
			t.Context(), ocr3types.OutcomeContext{}, types.Query{}, []types.AttributedObservation{},
		)
		require.Error(t, err)
	}

	l1 := internal.CounterFromHistogramByLabels(
		t, metrics.PromExecLatencyHistogram, "aptos", chainID, "outcome", "GetCommitReports",
	)
	require.Equal(t, 0, l1)

	l2 := testutil.ToFloat64(
		metrics.PromExecErrors.WithLabelValues("aptos", chainID, "outcome", "GetCommitReports"),
	)
	require.Equal(t, float64(count), l2)
}

type FakePlugin struct {
	query       types.Query
	observation types.Observation
	outcome     ocr3types.Outcome
	err         error
}

func (f FakePlugin) Query(ctx context.Context, outctx ocr3types.OutcomeContext) (types.Query, error) {
	return f.query, f.err
}

func (f FakePlugin) Observation(
	ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query,
) (types.Observation, error) {
	return f.observation, f.err
}

func (f FakePlugin) ValidateObservation(
	ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query, ao types.AttributedObservation,
) error {
	panic("implement me")
}

func (f FakePlugin) ObservationQuorum(
	ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (quorumReached bool, err error) {
	panic("implement me")
}

func (f FakePlugin) Outcome(
	ctx context.Context, outctx ocr3types.OutcomeContext, query types.Query, aos []types.AttributedObservation,
) (ocr3types.Outcome, error) {
	return f.outcome, f.err
}

func (f FakePlugin) Reports(
	ctx context.Context, seqNr uint64, outcome ocr3types.Outcome,
) ([]ocr3types.ReportPlus[[]byte], error) {
	panic("implement me")
}

func (f FakePlugin) ShouldAcceptAttestedReport(
	ctx context.Context, seqNr uint64, reportWithInfo ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	panic("implement me")
}

func (f FakePlugin) ShouldTransmitAcceptedReport(
	ctx context.Context, seqNr uint64, reportWithInfo ocr3types.ReportWithInfo[[]byte],
) (bool, error) {
	panic("implement me")
}

func (f FakePlugin) Close() error {
	panic("implement me")
}
