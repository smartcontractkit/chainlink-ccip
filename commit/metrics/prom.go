package metrics

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
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
		[]string{"chainFamily", "chainID", "processor", "method", "type"},
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
		[]string{"chainFamily", "chainID", "processor", "method"},
	)
	promProcessorErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ccip_commit_processor_errors",
			Help: "This metric tracks the number of errors in the commit processor observation",
		},
		[]string{"chainFamily", "chainID", "processor", "method"},
	)
	promSequenceNumbers = promauto.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ccip_commit_max_sequence_number",
			Help: "This metric tracks the max sequence number observed by the commit processor",
		},
		[]string{"chainFamily", "chainID", "sourceChainFamily", "sourceChain",
			"method", "source_network_name", "dest_network_name"},
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
	promCommitLatestRoundID = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_commit_latest_round_id",
		Help: "The latest round ID observed by the commit plugin",
	}, []string{"source_network_name", "dest_network_name", "contract_address", "plugin"})
)

type PromReporter struct {
	lggr        logger.Logger
	bhClient    beholder.Client
	chainFamily string
	chainID     string
	// Prometheus components
	merkleProcessorRmnReportHistogram *prometheus.HistogramVec
	rmnControllerRmnRequestHistogram  *prometheus.HistogramVec

	processorLatencyHistogram *prometheus.HistogramVec
	processorOutputCounter    *prometheus.CounterVec
	processorErrors           *prometheus.CounterVec
	commitLatestRound         *prometheus.GaugeVec
	sequenceNumbers           *prometheus.GaugeVec
	// Beholder components
	bhProcessorLatencyHistogram metric.Int64Histogram
	bhProcessorOutputCounter    metric.Int64Counter
	bhProcessorErrors           metric.Int64Counter
	bhSequenceNumbers           metric.Int64Gauge
	bhCommitLatestRound         metric.Int64Gauge
}

func NewPromReporter(
	lggr logger.Logger, selector cciptypes.ChainSelector, bhClient beholder.Client) (*PromReporter, error,
) {
	chainFamily, chainID, ok := libs.GetChainInfoFromSelector(selector)
	if !ok {
		return nil, fmt.Errorf("chainFamily and chainID not found for selector %d", selector)
	}
	processorLatencyHistogram, err := bhClient.Meter.Int64Histogram("ccip_commit_processor_latency")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_processor_latency histogram: %w", err)
	}
	processorOutputCounter, err := bhClient.Meter.Int64Counter("ccip_unexpired_commit_roots")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_unexpired_commit_roots gauge: %w", err)
	}
	processorErrors, err := bhClient.Meter.Int64Counter("ccip_commit_processor_errors")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_processor_errors gauge: %w", err)
	}
	sequenceNumbers, err := bhClient.Meter.Int64Gauge("ccip_commit_max_sequence_number")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_max_sequence_number gauge: %w", err)
	}
	commitLatestRoundID, err := bhClient.Meter.Int64Gauge("ccip_commit_latest_round_id")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_latest_round_id gauge: %w", err)
	}

	return &PromReporter{
		lggr:        lggr,
		bhClient:    bhClient,
		chainFamily: chainFamily,
		chainID:     chainID,

		merkleProcessorRmnReportHistogram: promMerkleProcessorRmnReportLatency,
		rmnControllerRmnRequestHistogram:  promRmnControllerRmnRequestLatency,

		sequenceNumbers:   promSequenceNumbers,
		commitLatestRound: promCommitLatestRoundID,

		processorLatencyHistogram: promProcessorLatencyHistogram,
		processorOutputCounter:    promProcessorOutputCounter,
		processorErrors:           promProcessorErrors,

		bhProcessorLatencyHistogram: processorLatencyHistogram,
		bhProcessorOutputCounter:    processorOutputCounter,
		bhProcessorErrors:           processorErrors,
		bhSequenceNumbers:           sequenceNumbers,
		bhCommitLatestRound:         commitLatestRoundID,
	}, nil
}

func (p *PromReporter) TrackObservation(obs committypes.Observation, round uint64) {
	for _, root := range obs.MerkleRootObs.MerkleRoots {
		sourceChainSelector := root.ChainSel
		maxSeqNr := root.SeqNumsRange.End()
		onramp := root.OnRampAddress.String()
		p.trackMaxSequenceNumber(sourceChainSelector, maxSeqNr, plugincommon.ObservationMethod)
		p.trackLatestRoundID(uint32(round), sourceChainSelector, onramp, plugincommon.ObservationMethod)
	}
}

func (p *PromReporter) TrackOutcome(outcome committypes.Outcome, round uint64) {
	for _, root := range outcome.MerkleRootOutcome.RootsToReport {
		sourceChainSelector := root.ChainSel
		maxSeqNr := root.SeqNumsRange.End()
		onramp := root.OnRampAddress.String()
		p.trackMaxSequenceNumber(sourceChainSelector, maxSeqNr, plugincommon.OutcomeMethod)
		p.trackLatestRoundID(uint32(round), sourceChainSelector, onramp, plugincommon.OutcomeMethod)
	}
}

func (p *PromReporter) trackLatestRoundID(
	latestRoundID uint32, sourceChainSelector cciptypes.ChainSelector, onramp string, method string,
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

	p.commitLatestRound.WithLabelValues(sourceName, destName, onramp, method).Set(float64(latestRoundID))
	p.bhCommitLatestRound.Record(context.Background(), int64(latestRoundID), metric.WithAttributes(
		attribute.String("source_network_name", sourceName),
		attribute.String("dest_network_name", destName),
		attribute.String("contract_address", onramp),
		attribute.String("plugin", method),
	))
}

func (p *PromReporter) trackMaxSequenceNumber(
	sourceChainSelector cciptypes.ChainSelector,
	maxSeqNr cciptypes.SeqNum,
	method string,
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
		"exec latest max seq num",
		"method", method,
		"sourceChain", sourceChainID,
		"sourceChainFamily", sourceFamily,
		"destChain", p.chainID,
		"destChainFamily", p.chainFamily,
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
	processor string,
	method plugincommon.MethodType,
	obs plugintypes.Trackable,
) {
	for key, val := range obs.Stats() {
		p.processorOutputCounter.
			WithLabelValues(p.chainFamily, p.chainID, processor, method, key).
			Add(float64(val))
		p.bhProcessorOutputCounter.Add(context.Background(), int64(val), metric.WithAttributes(
			attribute.String("chainFamily", p.chainFamily),
			attribute.String("chainID", p.chainID),
			attribute.String("processor", processor),
			attribute.String("method", method),
			attribute.String("type", key),
		))
	}
}
