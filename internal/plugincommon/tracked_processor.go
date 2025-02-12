package plugincommon

import (
	"context"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

type MetricsReporter interface {
	TrackProcessorObservation(processor string, obs plugintypes.Trackable)
	TrackProcessorOutcome(processor string, out plugintypes.Trackable)
	TrackProcessorLatency(processor string, method string, latency time.Duration, err error)
}

// TrackedProcessor wraps a PluginProcessor and tracks
// * latencies of most of the perf critical methods (Query, Observation, Outcome)
// * observations and outcomes (and their stats) of the processor
// * errors in the tracked methods
type TrackedProcessor[Query any, Observation plugintypes.Trackable, Outcome plugintypes.Trackable] struct {
	PluginProcessor[Query, Observation, Outcome]
	lggr          logger.Logger
	processorName string
	reporter      MetricsReporter
}

func NewTrackedProcessor[Query any, Observation plugintypes.Trackable, Outcome plugintypes.Trackable](
	lggr logger.Logger,
	origin PluginProcessor[Query, Observation, Outcome],
	processorName string,
	reporter MetricsReporter,
) *TrackedProcessor[Query, Observation, Outcome] {
	return &TrackedProcessor[Query, Observation, Outcome]{
		PluginProcessor: origin,
		lggr:            lggr,
		processorName:   processorName,
		reporter:        reporter,
	}
}

func (p *TrackedProcessor[Query, Observation, Outcome]) Query(ctx context.Context, prev Outcome) (Query, error) {
	return withTrackedMethod[Query](p, "query", func() (Query, error) {
		return p.PluginProcessor.Query(ctx, prev)
	})
}

func (p *TrackedProcessor[Query, Observation, Outcome]) Observation(
	ctx context.Context,
	prev Outcome,
	query Query,
) (Observation, error) {
	obs, err := withTrackedMethod[Observation](p, "observation", func() (Observation, error) {
		return p.PluginProcessor.Observation(ctx, prev, query)
	})
	if err != nil {
		p.reporter.TrackProcessorObservation(p.processorName, obs)
	}
	return obs, err
}

func (p *TrackedProcessor[Query, Observation, Outcome]) Outcome(
	ctx context.Context,
	prev Outcome,
	query Query,
	aos []AttributedObservation[Observation],
) (Outcome, error) {
	out, err := withTrackedMethod[Outcome](p, "outcome", func() (Outcome, error) {
		return p.PluginProcessor.Outcome(ctx, prev, query, aos)
	})
	if err != nil {
		p.reporter.TrackProcessorOutcome(p.processorName, out)
	}
	return out, err
}

func withTrackedMethod[T any, Query any, Observation plugintypes.Trackable, Outcome plugintypes.Trackable](
	p *TrackedProcessor[Query, Observation, Outcome],
	method string,
	f func() (T, error),
) (T, error) {
	queryStarted := time.Now()
	resp, err := f()

	latency := time.Since(queryStarted)
	p.reporter.TrackProcessorLatency(p.processorName, method, latency, err)
	p.lggr.Debugw("tracking processor latency",
		"processor", p.processorName,
		"method", method,
		"latency", latency,
	)

	return resp, err
}

type NoopReporter struct{}

func (n NoopReporter) TrackProcessorLatency(string, string, time.Duration, error) {}

func (n NoopReporter) TrackProcessorObservation(string, plugintypes.Trackable) {}

func (n NoopReporter) TrackProcessorOutcome(string, plugintypes.Trackable) {}
