package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"golang.org/x/exp/maps"

	sel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	PromExecOutputCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_output_sizes",
			Help: "This metric tracks the number of different items in the exec plugin",
		},
		[]string{"chainID", "method", "state", "type"},
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
		[]string{"chainID", "method", "state"},
	)
	PromExecErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_errors",
			Help: "This metric tracks the number of errors in the exec plugin",
		},
		[]string{"chainID", "method", "state"},
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
		[]string{"chainID", "processor", "method"},
	)
	PromExecProcessorErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_exec_processor_errors",
			Help: "This metric tracks the number of errors in the exec plugin processor",
		},
		[]string{"chainID", "processor", "method"},
	)
	PromSequenceNumbers = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ccip_exec_max_sequence_number",
			Help: "This metric tracks the max sequence number observed by the commit processor",
		},
		[]string{"chainID", "sourceChain", "method"},
	)
)

type PromReporter struct {
	lggr    logger.Logger
	chainID string

	// Prometheus reporters
	latencyHistogram          *prometheus.HistogramVec
	execErrors                *prometheus.CounterVec
	outputDetailsCounter      *prometheus.CounterVec
	sequenceNumbers           *prometheus.GaugeVec
	processorLatencyHistogram *prometheus.HistogramVec
	processorErrors           *prometheus.CounterVec
}

func NewPromReporter(lggr logger.Logger, selector cciptypes.ChainSelector) (*PromReporter, error) {
	chainID, err := sel.GetChainIDFromSelector(uint64(selector))
	if err != nil {
		return nil, err
	}

	return &PromReporter{
		lggr:    lggr,
		chainID: chainID,

		latencyHistogram:          PromExecLatencyHistogram,
		execErrors:                PromExecErrors,
		outputDetailsCounter:      PromExecOutputCounter,
		sequenceNumbers:           PromSequenceNumbers,
		processorLatencyHistogram: PromExecProcessorLatencyHistogram,
		processorErrors:           PromExecProcessorErrors,
	}, nil
}

func (p *PromReporter) TrackObservation(obs exectypes.Observation, state exectypes.PluginState) {
	p.trackOutputStats(obs, state, plugincommon.ObservationMethod)

	for sourceChainSelector, cr := range obs.Messages {
		maxSeqNr := pickHighestSeqNr(maps.Keys(cr))
		p.trackMaxSequenceNumber(sourceChainSelector, maxSeqNr, plugincommon.ObservationMethod)
	}
}

func (p *PromReporter) TrackOutcome(outcome exectypes.Outcome, state exectypes.PluginState) {
	p.trackOutputStats(&outcome, state, plugincommon.OutcomeMethod)

	for _, cr := range outcome.CommitReports {
		sourceChainSelector := cr.SourceChain
		maxSeqNr := pickHighestSeqNrInMessages(cr.Messages)
		p.trackMaxSequenceNumber(sourceChainSelector, maxSeqNr, plugincommon.OutcomeMethod)
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
			WithLabelValues(p.chainID, method, string(state)).
			Inc()
		return
	}

	p.latencyHistogram.
		WithLabelValues(p.chainID, method, string(state)).
		Observe(float64(latency))
}

func (p *PromReporter) TrackProcessorLatency(
	processor string,
	method plugincommon.MethodType,
	latency time.Duration,
	err error,
) {
	if err != nil {
		p.processorErrors.
			WithLabelValues(p.chainID, processor, method).
			Inc()
		return
	}

	p.processorLatencyHistogram.
		WithLabelValues(p.chainID, processor, method).
		Observe(float64(latency))
}

func (p *PromReporter) TrackProcessorOutput(
	string, plugincommon.MethodType, plugintypes.Trackable,
) {
	// noop
}

func (p *PromReporter) trackMaxSequenceNumber(
	sourceChainSelector cciptypes.ChainSelector,
	maxSeqNr int,
	method plugincommon.MethodType,
) {
	if maxSeqNr == 0 {
		return
	}

	sourceChain, err := sel.GetChainIDFromSelector(uint64(sourceChainSelector))
	if err != nil {
		p.lggr.Errorw("failed to get chain ID from selector", "err", err)
		return
	}

	p.sequenceNumbers.
		WithLabelValues(p.chainID, sourceChain, method).
		Set(float64(maxSeqNr))

	p.lggr.Debugw(
		"commit latest max seq num",
		"method", method,
		"sourceChain", sourceChain,
		"destChain", p.chainID,
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
			WithLabelValues(p.chainID, method, stringState, key).
			Add(float64(val))
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
