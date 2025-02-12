package metrics

import (
	"fmt"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	sel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

var (
	RequestLatencyBucketsMilliseconds = []float64{
		5, 10, 25, 50, 100, 250, 500, 1000, 2500, 5000, 10000,
	}
	promProcessorObservationCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_processor_observation_sizes",
			Help: "This metric tracks the number of different items in the commit processor",
		},
		[]string{"chainID", "processor", "type"},
	)
	promProcessorOutcomeCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_processor_outcome_sizes",
			Help: "This metric tracks the number of different items in the commit processor",
		},
		[]string{"chainID", "processor", "type"},
	)
	promProcessorLatencyHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ccip_commit_processor_latency_ms",
			Help:    "This metric tracks the client-observed latency of a single processor method",
			Buckets: RequestLatencyBucketsMilliseconds,
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
	promMerkleProcessorRmnReportLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ccip_commit_merkle_processor_rmn_report_latency_ms",
			Help:    "This metric tracks the client-observed latency of building an full RMN report with signatures",
			Buckets: RequestLatencyBucketsMilliseconds,
		},
		[]string{"chainID", "success"},
	)
	promRmnControllerRmnRequestLatency = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "ccip_commit_rmn_controller_rmn_request_latency_ms",
			Help:    "This metric tracks the client-observed latency of a single RMN request",
			Buckets: RequestLatencyBucketsMilliseconds,
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
	processorObservationCounter       *prometheus.CounterVec
	processorOutcomeCounter           *prometheus.CounterVec
	processorErrors                   *prometheus.CounterVec
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

		processorLatencyHistogram:   promProcessorLatencyHistogram,
		processorObservationCounter: promProcessorObservationCounter,
		processorOutcomeCounter:     promProcessorOutcomeCounter,
		processorErrors:             promProcessorErrors,
	}, nil
}

func (p *PromReporter) TrackObservation(obs committypes.Observation) {
	for _, root := range obs.MerkleRootObs.MerkleRoots {
		sourceChainSelector := root.ChainSel
		maxSeqNr := root.SeqNumsRange.End()

		// TODO Implement me in next PR!
		fmt.Println(sourceChainSelector, maxSeqNr)
	}
}

func (p *PromReporter) TrackOutcome(outcome committypes.Outcome) {
	for _, root := range outcome.MerkleRootOutcome.RootsToReport {
		sourceChainSelector := root.ChainSel
		maxSeqNr := root.SeqNumsRange.End()

		// TODO Implement me in next PR!
		fmt.Println(sourceChainSelector, maxSeqNr)
	}
}

func (p *PromReporter) TrackRmnReport(latency float64, success bool) {
	successStr := strconv.FormatBool(success)
	p.merkleProcessorRmnReportHistogram.WithLabelValues(p.chainID, successStr).Observe(latency)
}

func (p *PromReporter) TrackRmnRequest(method string, latency float64, nodeID uint64, err string) {
	nodeIDStr := strconv.FormatUint(nodeID, 10)
	p.rmnControllerRmnRequestHistogram.WithLabelValues(method, nodeIDStr, err).Observe(latency)
}

func (p *PromReporter) TrackProcessorLatency(processor string, method string, latency time.Duration, err error) {
	if err != nil {
		p.processorErrors.
			WithLabelValues(p.chainID, processor, method).
			Inc()
		return
	}

	p.processorLatencyHistogram.
		WithLabelValues(p.chainID, processor, method).
		Observe(float64(latency.Milliseconds()))
}

func (p *PromReporter) TrackProcessorObservation(processor string, obs plugintypes.Trackable) {
	for key, val := range obs.Stats() {
		p.processorObservationCounter.
			WithLabelValues(p.chainID, processor, key).
			Add(float64(val))
	}
}

func (p *PromReporter) TrackProcessorOutcome(processor string, out plugintypes.Trackable) {
	for key, val := range out.Stats() {
		p.processorOutcomeCounter.
			WithLabelValues(p.chainID, processor, key).
			Add(float64(val))
	}
}
