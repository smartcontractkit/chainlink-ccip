package metrics

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	sel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	rmnLatencyBucketsMilliseconds = []float64{
		5, 10, 25, 50, 100, 250, 500, 1000, 2500, 5000, 10000,
	}
	promProcessorOutputCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_processor_output_sizes",
			Help: "This metric tracks the number of different items in the commit processor",
		},
		[]string{"chainID", "processor", "method", "type"},
	)
	promProcessorLatencyHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "ccip_commit_processor_latency",
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
	promProcessorErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_processor_errors",
			Help: "This metric tracks the number of errors in the commit processor observation",
		},
		[]string{"chainID", "processor", "method"},
	)
	promSequenceNumbers = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ccip_commit_max_sequence_number",
			Help: "This metric tracks the max sequence number observed by the commit processor",
		},
		[]string{"chainID", "sourceChain", "method"},
	)
	promMerkleProcessorRmnReportLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ccip_commit_merkle_processor_rmn_report_latency_ms",
			Help:    "This metric tracks the client-observed latency of building an full RMN report with signatures",
			Buckets: rmnLatencyBucketsMilliseconds,
		},
		[]string{"chainID", "success"},
	)
	promRmnControllerRmnRequestLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ccip_commit_rmn_controller_rmn_request_latency_ms",
			Help:    "This metric tracks the client-observed latency of a single RMN request",
			Buckets: rmnLatencyBucketsMilliseconds,
		},
		[]string{"method", "nodeID", "error"},
	)
)

type PromReporter struct {
	lggr    logger.Logger
	chainID string
	// Prometheus components
	merkleProcessorRmnReportHistogram *prometheus.HistogramVec
	rmnControllerRmnRequestHistogram  *prometheus.HistogramVec
	processorLatencyHistogram         *prometheus.HistogramVec
	processorOutputCounter            *prometheus.CounterVec
	processorErrors                   *prometheus.CounterVec
	sequenceNumbers                   *prometheus.GaugeVec
}

func NewPromReporter(lggr logger.Logger, selector cciptypes.ChainSelector) (*PromReporter, error) {
	chainID, err := sel.GetChainIDFromSelector(uint64(selector))
	if err != nil {
		return nil, err
	}

	return &PromReporter{
		lggr:    lggr,
		chainID: chainID,

		merkleProcessorRmnReportHistogram: promMerkleProcessorRmnReportLatency,
		rmnControllerRmnRequestHistogram:  promRmnControllerRmnRequestLatency,

		sequenceNumbers: promSequenceNumbers,

		processorLatencyHistogram: promProcessorLatencyHistogram,
		processorOutputCounter:    promProcessorOutputCounter,
		processorErrors:           promProcessorErrors,
	}, nil
}

func (p *PromReporter) TrackObservation(obs committypes.Observation) {
	for _, root := range obs.MerkleRootObs.MerkleRoots {
		sourceChainSelector := root.ChainSel
		maxSeqNr := root.SeqNumsRange.End()

		p.trackMaxSequenceNumber(sourceChainSelector, maxSeqNr, plugincommon.ObservationMethod)
	}
}

func (p *PromReporter) TrackOutcome(outcome committypes.Outcome) {
	for _, root := range outcome.MerkleRootOutcome.RootsToReport {
		sourceChainSelector := root.ChainSel
		maxSeqNr := root.SeqNumsRange.End()

		p.trackMaxSequenceNumber(sourceChainSelector, maxSeqNr, plugincommon.OutcomeMethod)
	}
}

func (p *PromReporter) trackMaxSequenceNumber(
	sourceChainSelector cciptypes.ChainSelector,
	maxSeqNr cciptypes.SeqNum,
	method string,
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
		"exec latest max seq num",
		"method", method,
		"sourceChain", sourceChain,
		"destChain", p.chainID,
		"maxSeqNr", maxSeqNr,
	)
}

func (p *PromReporter) TrackRmnReport(latency float64, success bool) {
	successStr := strconv.FormatBool(success)
	p.merkleProcessorRmnReportHistogram.WithLabelValues(p.chainID, successStr).Observe(latency)
}

func (p *PromReporter) TrackRmnRequest(method string, latency float64, nodeID uint64, err string) {
	nodeIDStr := strconv.FormatUint(nodeID, 10)
	p.rmnControllerRmnRequestHistogram.WithLabelValues(method, nodeIDStr, err).Observe(latency)
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
	processor string,
	method plugincommon.MethodType,
	obs plugintypes.Trackable,
) {
	for key, val := range obs.Stats() {
		p.processorOutputCounter.
			WithLabelValues(p.chainID, processor, method, key).
			Add(float64(val))
	}
}
