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

type ObservedProcessor[Query any, Observation plugintypes.Trackable, Outcome plugintypes.Trackable] struct {
	PluginProcessor[Query, Observation, Outcome]
	lggr          logger.Logger
	processorName string
	reporter      MetricsReporter
}

func NewObservedProcessor[Query any, Observation plugintypes.Trackable, Outcome plugintypes.Trackable](
	lggr logger.Logger,
	origin PluginProcessor[Query, Observation, Outcome],
	processorName string,
	reporter MetricsReporter,
) *ObservedProcessor[Query, Observation, Outcome] {
	return &ObservedProcessor[Query, Observation, Outcome]{
		PluginProcessor: origin,
		lggr:            lggr,
		processorName:   processorName,
		reporter:        reporter,
	}
}

func (p *ObservedProcessor[Query, Observation, Outcome]) Query(ctx context.Context, prev Outcome) (Query, error) {
	return withObservedQuery[Query](p, "query", func() (Query, error) {
		return p.PluginProcessor.Query(ctx, prev)
	})
}

func (p *ObservedProcessor[Query, Observation, Outcome]) Observation(ctx context.Context, prev Outcome, query Query) (Observation, error) {
	obs, err := withObservedQuery[Observation](p, "observation", func() (Observation, error) {
		return p.PluginProcessor.Observation(ctx, prev, query)
	})
	if err != nil {
		p.reporter.TrackProcessorObservation(p.processorName, obs)
	}
	return obs, err
}

func (p *ObservedProcessor[Query, Observation, Outcome]) Outcome(ctx context.Context, prev Outcome, query Query, aos []AttributedObservation[Observation]) (Outcome, error) {
	out, err := withObservedQuery[Outcome](p, "outcome", func() (Outcome, error) {
		return p.PluginProcessor.Outcome(ctx, prev, query, aos)
	})
	if err != nil {
		p.reporter.TrackProcessorOutcome(p.processorName, out)
	}
	return out, err
}

func withObservedQuery[T any, Query any, Observation plugintypes.Trackable, Outcome plugintypes.Trackable](
	p *ObservedProcessor[Query, Observation, Outcome],
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
