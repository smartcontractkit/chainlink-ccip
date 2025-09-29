package metrics

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

var (
	PromExecOutputCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_output_sizes",
			Help: "This metric tracks the number of different items in the exec plugin",
		},
		[]string{"chainFamily", "chainID", "method", "state", "type"},
	)
	PromExecLatencyHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ccip_exec_latency",
			Help: "This metric tracks the client-observed latency of a single exec plugin method",
			Buckets: []float64{
				float64(50 * time.Millisecond),
				float64(100 * time.Millisecond),
				float64(200 * time.Millisecond),
				float64(500 * time.Millisecond),
				float64(700 * time.Millisecond),
				float64(time.Second),
				float64(2 * time.Second),
				float64(5 * time.Second),
				float64(7 * time.Second),
				float64(10 * time.Second),
				float64(20 * time.Second),
			},
		},
		[]string{"chainFamily", "chainID", "method", "state"},
	)
	PromExecErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_errors",
			Help: "This metric tracks the number of errors in the exec plugin",
		},
		[]string{"chainFamily", "chainID", "method", "state"},
	)
	PromExecProcessorLatencyHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ccip_exec_processor_latency",
			Help: "This metric tracks the client-observed latency of a single processor method",
			Buckets: []float64{
				float64(50 * time.Millisecond),
				float64(100 * time.Millisecond),
				float64(200 * time.Millisecond),
				float64(500 * time.Millisecond),
				float64(700 * time.Millisecond),
				float64(time.Second),
				float64(2 * time.Second),
				float64(5 * time.Second),
				float64(7 * time.Second),
				float64(10 * time.Second),
				float64(20 * time.Second),
			},
		},
		[]string{"chainFamily", "chainID", "processor", "method"},
	)
	PromExecProcessorErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_processor_errors",
			Help: "This metric tracks the number of errors in the exec plugin processor",
		},
		[]string{"chainFamily", "chainID", "processor", "method"},
	)
	PromSequenceNumbers = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ccip_exec_max_sequence_number",
			Help: "This metric tracks the max sequence number observed by the commit processor",
		},
		[]string{"chainFamily", "chainID", "sourceChainFamily", "sourceChain",
			"method", "source_network_name", "dest_network_name"},
	)
	PromExecLatestRoundID = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_exec_latest_round_id",
		Help: "The latest round ID observed by the exec plugin",
	}, []string{"source_network_name", "dest_network_name", "plugin"})
)

type PromReporter struct {
	lggr        logger.Logger
	bhClient    beholder.Client
	chainFamily string
	chainID     string

	// Prometheus reporters
	latencyHistogram          *prometheus.HistogramVec
	execErrors                *prometheus.CounterVec
	outputDetailsCounter      *prometheus.CounterVec
	sequenceNumbers           *prometheus.GaugeVec
	processorLatencyHistogram *prometheus.HistogramVec
	processorErrors           *prometheus.CounterVec
	latestRoundID             *prometheus.GaugeVec
	// Beholder reporters
	bhProcessorLatencyHistogram metric.Int64Histogram
	bhLatencyHistogram          metric.Int64Histogram
	bhExecErrors                metric.Int64Counter
	bhOutputDetailsCounter      metric.Int64Counter
	bhSequenceNumbers           metric.Int64Gauge
	beholderProcessorErrors     metric.Int64Counter
	bhExecLatestRound           metric.Int64Gauge
}

func NewPromReporter(
	lggr logger.Logger, selector cciptypes.ChainSelector, bhClient beholder.Client) (*PromReporter, error,
) {
	chainFamily, chainID, ok := libs.GetChainInfoFromSelector(selector)
	if !ok {
		return nil, fmt.Errorf("chainFamily and chainID not found for selector %d", selector)
	}

	latencyHistogram, err := bhClient.Meter.Int64Histogram("ccip_exec_latency")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_exec_latency histogram: %w", err)
	}
	processorLatencyHistogram, err := bhClient.Meter.Int64Histogram("ccip_exec_processor_latency")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_exec_processor_latency histogram: %w", err)
	}
	execErrors, err := bhClient.Meter.Int64Counter("ccip_exec_errors")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_exec_errors counter: %w", err)
	}
	outputDetailsCounter, err := bhClient.Meter.Int64Counter("ccip_exec_output_sizes")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_exec_output_sizes counter: %w", err)
	}
	sequenceNumbers, err := bhClient.Meter.Int64Gauge("ccip_exec_max_sequence_number")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_exec_max_sequence_number gauge: %w", err)
	}
	processorErrors, err := bhClient.Meter.Int64Counter("ccip_exec_processor_errors")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_exec_processor_errors counter: %w", err)
	}
	execLatestRoundID, err := bhClient.Meter.Int64Gauge("ccip_exec_latest_round_id")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_exec_latest_round_id gauge: %w", err)
	}

	return &PromReporter{
		lggr:        lggr,
		chainFamily: chainFamily,
		bhClient:    bhClient,
		chainID:     chainID,

		latencyHistogram:          PromExecLatencyHistogram,
		execErrors:                PromExecErrors,
		outputDetailsCounter:      PromExecOutputCounter,
		sequenceNumbers:           PromSequenceNumbers,
		processorLatencyHistogram: PromExecProcessorLatencyHistogram,
		processorErrors:           PromExecProcessorErrors,
		latestRoundID:             PromExecLatestRoundID,

		bhLatencyHistogram:          latencyHistogram,
		bhProcessorLatencyHistogram: processorLatencyHistogram,
		bhExecErrors:                execErrors,
		bhOutputDetailsCounter:      outputDetailsCounter,
		bhSequenceNumbers:           sequenceNumbers,
		beholderProcessorErrors:     processorErrors,
		bhExecLatestRound:           execLatestRoundID,
	}, nil
}

func (p *PromReporter) TrackObservation(obs exectypes.Observation, state exectypes.PluginState, round uint64) {
	p.trackOutputStats(obs, state, plugincommon.ObservationMethod)

	for sourceChainSelector, cr := range obs.Messages {
		maxSeqNr := pickHighestSeqNr(maps.Keys(cr))
		p.trackMaxSequenceNumber(sourceChainSelector, maxSeqNr, plugincommon.ObservationMethod)
		p.trackLatestRoundID(round, sourceChainSelector, plugincommon.OutcomeMethod)
	}
}

func (p *PromReporter) TrackOutcome(outcome exectypes.Outcome, state exectypes.PluginState, round uint64) {
	p.trackOutputStats(&outcome, state, plugincommon.OutcomeMethod)

	for _, cr := range outcome.CommitReports {
		sourceChainSelector := cr.SourceChain
		maxSeqNr := pickHighestSeqNrInMessages(cr.Messages)
		p.trackMaxSequenceNumber(sourceChainSelector, maxSeqNr, plugincommon.OutcomeMethod)
		p.trackLatestRoundID(round, sourceChainSelector, plugincommon.OutcomeMethod)
	}
}

func (p *PromReporter) TrackLatency(
	state exectypes.PluginState,
	method plugincommon.MethodType,
	latency time.Duration,
	err error,
) {
	if err != nil {
		p.execErrors.
			WithLabelValues(p.chainFamily, p.chainID, method, string(state)).
			Inc()
		p.bhExecErrors.Add(context.Background(), 1, metric.WithAttributes(
			attribute.String("chainFamily", p.chainFamily),
			attribute.String("chainID", p.chainID),
			attribute.String("method", method),
			attribute.String("state", string(state)),
		))
		return
	}
	p.latencyHistogram.WithLabelValues(p.chainFamily, p.chainID, method, string(state)).
		Observe(float64(latency))
	p.bhLatencyHistogram.Record(context.Background(), int64(latency), metric.WithAttributes(
		attribute.String("chainFamily", p.chainFamily),
		attribute.String("chainID", p.chainID),
		attribute.String("method", method),
		attribute.String("state", string(state))))
}

func (p *PromReporter) TrackProcessorLatency(
	processor string,
	method plugincommon.MethodType,
	latency time.Duration,
	err error,
) {
	if err != nil {
		p.processorErrors.
			WithLabelValues(p.chainFamily, p.chainID, processor, method).
			Inc()
		return
	}

	p.processorLatencyHistogram.
		WithLabelValues(p.chainFamily, p.chainID, processor, method).
		Observe(float64(latency))
	p.bhProcessorLatencyHistogram.Record(context.Background(), int64(latency), metric.WithAttributes(
		attribute.String("chainFamily", p.chainFamily),
		attribute.String("chainID", p.chainID),
		attribute.String("processor", processor),
		attribute.String("method", method),
	))
}

func (p *PromReporter) TrackProcessorOutput(
	string, plugincommon.MethodType, plugintypes.Trackable,
) {
	// noop
}

func (p *PromReporter) trackLatestRoundID(
	latestRoundID uint64, sourceChainSelector cciptypes.ChainSelector, method plugincommon.MethodType,
) {
	sourceFamily, sourceChainID, ok := libs.GetChainInfoFromSelector(sourceChainSelector)
	if !ok {
		p.lggr.Errorw("failed to get chain ID from selector", "selector", sourceChainSelector)
		return
	}
	sourceName, err := libs.GetNameFromIDAndFamily(sourceChainID, sourceFamily)
	if err != nil {
		p.lggr.Errorw("failed to get chain name from ID and family", "chainID",
			sourceChainID, "family", sourceFamily, "err", err)
	}
	destName, err := libs.GetNameFromIDAndFamily(p.chainID, p.chainFamily)
	if err != nil {
		p.lggr.Errorw("failed to get chain name from ID and family", "chainID",
			p.chainID, "family", p.chainFamily, "err", err)
	}
	p.latestRoundID.WithLabelValues(sourceName, destName, method).Set(float64(latestRoundID))
	p.bhExecLatestRound.Record(context.Background(), int64(latestRoundID), metric.WithAttributes(
		attribute.String("source_network_name", sourceName),
		attribute.String("dest_network_name", destName),
		attribute.String("plugin", method),
	))
}

func (p *PromReporter) trackMaxSequenceNumber(
	sourceChainSelector cciptypes.ChainSelector,
	maxSeqNr int,
	method plugincommon.MethodType,
) {
	if maxSeqNr == 0 {
		return
	}

	sourceFamily, sourceChainID, ok := libs.GetChainInfoFromSelector(sourceChainSelector)
	if !ok {
		p.lggr.Errorw("failed to get chain ID from selector", "selector", sourceChainSelector)
		return
	}
	sourceName, err := libs.GetNameFromIDAndFamily(sourceChainID, sourceFamily)
	if err != nil {
		p.lggr.Errorw("failed to get chain name from ID and family", "chainID",
			sourceChainID, "family", sourceFamily, "err", err)
	}
	destName, err := libs.GetNameFromIDAndFamily(p.chainID, p.chainFamily)
	if err != nil {
		p.lggr.Errorw("failed to get chain name from ID and family", "chainID",
			p.chainID, "family", p.chainFamily, "err", err)
	}

	p.sequenceNumbers.
		WithLabelValues(p.chainFamily, p.chainID, sourceFamily, sourceChainID, method, sourceName, destName).
		Set(float64(maxSeqNr))
	p.bhSequenceNumbers.Record(context.Background(), int64(maxSeqNr), metric.WithAttributes(
		attribute.String("chainFamily", p.chainFamily),
		attribute.String("chainID", p.chainID),
		attribute.String("sourceChainFamily", sourceFamily),
		attribute.String("sourceChainID", sourceChainID),
		attribute.String("method", method),
		attribute.String("source_network_name", sourceName),
		attribute.String("dest_network_name", destName),
	))

	p.lggr.Debugw(
		"commit latest max seq num",
		"method", method,
		"sourceChain", sourceChainID,
		"sourceChainFamily", sourceFamily,
		"destChain", p.chainID,
		"destChainFamily", p.chainFamily,
		"maxSeqNr", maxSeqNr,
	)
}

func (p *PromReporter) trackOutputStats(
	output plugintypes.Trackable,
	state exectypes.PluginState,
	method plugincommon.MethodType,
) {
	stringState := string(state)
	for key, val := range output.Stats() {
		p.outputDetailsCounter.
			WithLabelValues(p.chainFamily, p.chainID, method, stringState, key).
			Add(float64(val))
		p.bhOutputDetailsCounter.Add(context.Background(), int64(val), metric.WithAttributes(
			attribute.String("chainFamily", p.chainFamily),
			attribute.String("chainID", p.chainID),
			attribute.String("method", method),
			attribute.String("state", stringState),
			attribute.String("type", key),
		))
	}
}

func pickHighestSeqNrInMessages(messages []cciptypes.Message) int {
	seqNrs := make([]cciptypes.SeqNum, len(messages))
	for i, m := range messages {
		seqNrs[i] = m.Header.SequenceNumber
	}
	return pickHighestSeqNr(seqNrs)
}

func pickHighestSeqNr(seqNrs []cciptypes.SeqNum) int {
	seqNr := cciptypes.SeqNum(0)
	for _, s := range seqNrs {
		if s > seqNr {
			seqNr = s
		}
	}
	return int(seqNr)
}
