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
	promLooppCCIPProviderSupported = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_commit_loopp_ccip_provider_supported",
		Help: "Tracks whether LOOPP CCIP provider is supported for each chain family (1 = supported, 0 = not supported)",
	}, []string{"chain_family"})
	promCommitConfigDigestMismatch = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_commit_config_digest_mismatch",
		Help: "Reports whether the home chain config digest differs from the offramp config digest (1 = mismatch, 0 = match)",
	}, []string{"chain_family", "chain_id"})
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
	looppProviderSupported    *prometheus.GaugeVec
	// Beholder components
	bhProcessorLatencyHistogram metric.Int64Histogram
	bhProcessorOutputCounter    metric.Int64Counter
	bhProcessorErrors           metric.Int64Counter
	bhSequenceNumbers           metric.Int64Gauge
	bhCommitLatestRound         metric.Int64Gauge
	bhLooppProviderSupported    metric.Int64Gauge

	configDigestMismatch   *prometheus.GaugeVec
	bhConfigDigestMismatch metric.Int64Gauge

	// Beholder-only metrics (no prometheus equivalent): mainnet NOP nodes aren't scraped by our
	// prometheus, only ingested via beholder, so anything meant to give us visibility into NOP-run
	// nodes should be defined here rather than as a promauto metric.
	bhPluginHeartbeat            metric.Int64Counter
	bhReportValidationRejected   metric.Int64Counter
	bhOnRampMaxSeqNum            metric.Int64Gauge
	bhOffRampNextSeqNum          metric.Int64Gauge
	bhPendingMessages            metric.Int64Gauge
	bhRmnCurseActive             metric.Int64Gauge
	bhSourceChainCursed          metric.Int64Gauge
	bhMerkleRootObservationErrs  metric.Int64Counter
	bhFChainReadErrors           metric.Int64Counter
	bhConsensusObservationFailed metric.Int64Counter
	bhReportTransmissionGaveUp   metric.Int64Counter
	bhReportTransmissionAttempts metric.Int64Histogram
	bhRangeTruncated             metric.Int64Counter
	bhOffRampLaneStatus          metric.Int64Gauge
	bhSeqNumInvariantViolation   metric.Int64Counter
	bhOffRampConsensusInsuff     metric.Int64Counter
	bhConsensusDropped           metric.Int64Counter
}

//nolint:gocyclo
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
	looppProviderSupported, err := bhClient.Meter.Int64Gauge("ccip_commit_loopp_ccip_provider_supported")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_loopp_ccip_provider_supported gauge: %w", err)
	}
	configDigestMismatch, err := bhClient.Meter.Int64Gauge("ccip_commit_config_digest_mismatch")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_config_digest_mismatch gauge: %w", err)
	}

	pluginHeartbeat, err := bhClient.Meter.Int64Counter("ccip_commit_plugin_heartbeat")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_plugin_heartbeat counter: %w", err)
	}
	reportValidationRejected, err := bhClient.Meter.Int64Counter("ccip_commit_report_validation_rejected")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_report_validation_rejected counter: %w", err)
	}
	onRampMaxSeqNum, err := bhClient.Meter.Int64Gauge("ccip_commit_onramp_max_seq_num")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_onramp_max_seq_num gauge: %w", err)
	}
	offRampNextSeqNum, err := bhClient.Meter.Int64Gauge("ccip_commit_offramp_next_seq_num")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_offramp_next_seq_num gauge: %w", err)
	}
	pendingMessages, err := bhClient.Meter.Int64Gauge("ccip_commit_pending_messages")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_pending_messages gauge: %w", err)
	}
	rmnCurseActive, err := bhClient.Meter.Int64Gauge("ccip_commit_rmn_curse_active")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_rmn_curse_active gauge: %w", err)
	}
	sourceChainCursed, err := bhClient.Meter.Int64Gauge("ccip_commit_source_chain_cursed")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_source_chain_cursed gauge: %w", err)
	}
	merkleRootObservationErrs, err := bhClient.Meter.Int64Counter("ccip_commit_merkleroot_observation_errors")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_merkleroot_observation_errors counter: %w", err)
	}
	fChainReadErrors, err := bhClient.Meter.Int64Counter("ccip_commit_fchain_read_errors")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_fchain_read_errors counter: %w", err)
	}
	consensusObservationFailed, err := bhClient.Meter.Int64Counter("ccip_commit_consensus_observation_failed")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_consensus_observation_failed counter: %w", err)
	}
	reportTransmissionGaveUp, err := bhClient.Meter.Int64Counter("ccip_commit_report_transmission_gave_up")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_report_transmission_gave_up counter: %w", err)
	}
	reportTransmissionAttempts, err := bhClient.Meter.Int64Histogram("ccip_commit_report_transmission_attempts")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_report_transmission_attempts histogram: %w", err)
	}
	rangeTruncated, err := bhClient.Meter.Int64Counter("ccip_commit_range_truncated")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_range_truncated counter: %w", err)
	}
	offRampLaneStatus, err := bhClient.Meter.Int64Gauge("ccip_commit_offramp_lane_status")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_offramp_lane_status gauge: %w", err)
	}
	seqNumInvariantViolation, err := bhClient.Meter.Int64Counter("ccip_commit_seqnum_invariant_violation")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_seqnum_invariant_violation counter: %w", err)
	}
	offRampConsensusInsuff, err := bhClient.Meter.Int64Counter("ccip_commit_offramp_consensus_insufficient")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_offramp_consensus_insufficient counter: %w", err)
	}
	consensusDropped, err := bhClient.Meter.Int64Counter("ccip_commit_consensus_dropped")
	if err != nil {
		return nil, fmt.Errorf("failed to register ccip_commit_consensus_dropped counter: %w", err)
	}

	return &PromReporter{
		lggr:        lggr,
		bhClient:    bhClient,
		chainFamily: chainFamily,
		chainID:     chainID,

		merkleProcessorRmnReportHistogram: promMerkleProcessorRmnReportLatency,
		rmnControllerRmnRequestHistogram:  promRmnControllerRmnRequestLatency,

		sequenceNumbers:        promSequenceNumbers,
		commitLatestRound:      promCommitLatestRoundID,
		looppProviderSupported: promLooppCCIPProviderSupported,
		configDigestMismatch:   promCommitConfigDigestMismatch,

		processorLatencyHistogram: promProcessorLatencyHistogram,
		processorOutputCounter:    promProcessorOutputCounter,
		processorErrors:           promProcessorErrors,

		bhProcessorLatencyHistogram: processorLatencyHistogram,
		bhProcessorOutputCounter:    processorOutputCounter,
		bhProcessorErrors:           processorErrors,
		bhSequenceNumbers:           sequenceNumbers,
		bhCommitLatestRound:         commitLatestRoundID,
		bhLooppProviderSupported:    looppProviderSupported,
		bhConfigDigestMismatch:      configDigestMismatch,

		bhPluginHeartbeat:            pluginHeartbeat,
		bhReportValidationRejected:   reportValidationRejected,
		bhOnRampMaxSeqNum:            onRampMaxSeqNum,
		bhOffRampNextSeqNum:          offRampNextSeqNum,
		bhPendingMessages:            pendingMessages,
		bhRmnCurseActive:             rmnCurseActive,
		bhSourceChainCursed:          sourceChainCursed,
		bhMerkleRootObservationErrs:  merkleRootObservationErrs,
		bhFChainReadErrors:           fChainReadErrors,
		bhConsensusObservationFailed: consensusObservationFailed,
		bhReportTransmissionGaveUp:   reportTransmissionGaveUp,
		bhReportTransmissionAttempts: reportTransmissionAttempts,
		bhRangeTruncated:             rangeTruncated,
		bhOffRampLaneStatus:          offRampLaneStatus,
		bhSeqNumInvariantViolation:   seqNumInvariantViolation,
		bhOffRampConsensusInsuff:     offRampConsensusInsuff,
		bhConsensusDropped:           consensusDropped,
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

func (p *PromReporter) TrackLooppProviderSupported(looppCCIPProviderSupported map[string]bool) {
	for chainFamily, supported := range looppCCIPProviderSupported {
		value := float64(0)
		if supported {
			value = 1
		}
		p.looppProviderSupported.WithLabelValues(chainFamily).Set(value)
		p.bhLooppProviderSupported.Record(context.Background(), int64(value), metric.WithAttributes(
			attribute.String("chain_family", chainFamily),
		))
	}
}

func (p *PromReporter) TrackConfigDigestMismatch(mismatch bool) {
	var value float64
	if mismatch {
		value = 1
	}
	p.configDigestMismatch.WithLabelValues(p.chainFamily, p.chainID).Set(value)
	p.bhConfigDigestMismatch.Record(context.Background(), int64(value), metric.WithAttributes(
		attribute.String("chain_family", p.chainFamily),
		attribute.String("chain_id", p.chainID),
	))
}

func (p *PromReporter) TrackPluginHeartbeat(phase string) {
	p.bhPluginHeartbeat.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("chainFamily", p.chainFamily),
		attribute.String("chainID", p.chainID),
		attribute.String("phase", phase),
	))
}

func (p *PromReporter) TrackReportValidationRejected(phase string, reason string) {
	p.bhReportValidationRejected.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("chainFamily", p.chainFamily),
		attribute.String("chainID", p.chainID),
		attribute.String("phase", phase),
		attribute.String("reason", reason),
	))
}

// sourceChainAttrs resolves a source chain selector into the same family/ID/network-name attributes
// used elsewhere in this reporter (see trackMaxSequenceNumber), so dashboards can join on network name
// consistently across metrics.
func (p *PromReporter) sourceChainAttrs(sourceChainSelector cciptypes.ChainSelector) []attribute.KeyValue {
	sourceFamily, sourceChainID, ok := libs.GetChainInfoFromSelector(sourceChainSelector)
	if !ok {
		p.lggr.Errorw("failed to get chain ID from selector", "selector", sourceChainSelector)
		return []attribute.KeyValue{
			attribute.String("sourceChainSelector", strconv.FormatUint(uint64(sourceChainSelector), 10)),
		}
	}
	sourceName, err := libs.GetNameFromIDAndFamily(sourceChainID, sourceFamily)
	if err != nil {
		p.lggr.Errorw("failed to get chain name from ID and family", "chainID",
			sourceChainID, "family", sourceFamily, "err", err)
	}
	return []attribute.KeyValue{
		attribute.String("sourceChainFamily", sourceFamily),
		attribute.String("sourceChainID", sourceChainID),
		attribute.String("source_network_name", sourceName),
	}
}

func (p *PromReporter) TrackOnRampMaxSeqNum(sourceChain cciptypes.ChainSelector, seqNum cciptypes.SeqNum) {
	p.bhOnRampMaxSeqNum.Record(context.Background(), int64(seqNum),
		metric.WithAttributes(p.sourceChainAttrs(sourceChain)...))
}

func (p *PromReporter) TrackOffRampNextSeqNum(sourceChain cciptypes.ChainSelector, seqNum cciptypes.SeqNum) {
	p.bhOffRampNextSeqNum.Record(context.Background(), int64(seqNum),
		metric.WithAttributes(p.sourceChainAttrs(sourceChain)...))
}

func (p *PromReporter) TrackPendingMessages(sourceChain cciptypes.ChainSelector, pending uint64) {
	p.bhPendingMessages.Record(context.Background(), int64(pending),
		metric.WithAttributes(p.sourceChainAttrs(sourceChain)...))
}

func (p *PromReporter) TrackRmnCurseActive(curseType string, active bool) {
	var value int64
	if active {
		value = 1
	}
	p.bhRmnCurseActive.Record(context.Background(), value, metric.WithAttributes(
		attribute.String("chain_family", p.chainFamily),
		attribute.String("chain_id", p.chainID),
		attribute.String("curse_type", curseType),
	))
}

func (p *PromReporter) TrackSourceChainCursed(sourceChain cciptypes.ChainSelector, cursed bool) {
	var value int64
	if cursed {
		value = 1
	}
	p.bhSourceChainCursed.Record(context.Background(), value,
		metric.WithAttributes(p.sourceChainAttrs(sourceChain)...))
}

func (p *PromReporter) TrackObservationError(sourceChain cciptypes.ChainSelector, reason string) {
	attrs := append(p.sourceChainAttrs(sourceChain), attribute.String("reason", reason))
	p.bhMerkleRootObservationErrs.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackFChainReadError() {
	p.bhFChainReadErrors.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("chainFamily", p.chainFamily),
		attribute.String("chainID", p.chainID),
	))
}

func (p *PromReporter) TrackConsensusObservationFailed() {
	p.bhConsensusObservationFailed.Add(context.Background(), 1, metric.WithAttributes(
		attribute.String("chain_family", p.chainFamily),
		attribute.String("chain_id", p.chainID),
	))
}

func (p *PromReporter) TrackReportTransmissionGaveUp(sourceChain cciptypes.ChainSelector) {
	p.bhReportTransmissionGaveUp.Add(context.Background(), 1,
		metric.WithAttributes(p.sourceChainAttrs(sourceChain)...))
}

func (p *PromReporter) TrackReportTransmissionAttempts(attempts uint, success bool) {
	p.bhReportTransmissionAttempts.Record(context.Background(), int64(attempts), metric.WithAttributes(
		attribute.String("chainFamily", p.chainFamily),
		attribute.String("chainID", p.chainID),
		attribute.Bool("success", success),
	))
}

func (p *PromReporter) TrackRangeTruncated(sourceChain cciptypes.ChainSelector) {
	p.bhRangeTruncated.Add(context.Background(), 1, metric.WithAttributes(p.sourceChainAttrs(sourceChain)...))
}

func (p *PromReporter) TrackOffRampLaneStatus(sourceChain cciptypes.ChainSelector, status string, active bool) {
	var value int64
	if active {
		value = 1
	}
	attrs := append(p.sourceChainAttrs(sourceChain), attribute.String("status", status))
	p.bhOffRampLaneStatus.Record(context.Background(), value, metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackSeqNumInvariantViolation(sourceChain cciptypes.ChainSelector, violationType string) {
	attrs := append(p.sourceChainAttrs(sourceChain), attribute.String("type", violationType))
	p.bhSeqNumInvariantViolation.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackOffRampConsensusInsufficient(sourceChain cciptypes.ChainSelector) {
	p.bhOffRampConsensusInsuff.Add(context.Background(), 1, metric.WithAttributes(p.sourceChainAttrs(sourceChain)...))
}

func (p *PromReporter) TrackConsensusDropped(
	objectName string, key string, reason string, sourceChain cciptypes.ChainSelector,
) {
	attrs := []attribute.KeyValue{
		attribute.String("chainFamily", p.chainFamily),
		attribute.String("chainID", p.chainID),
		attribute.String("objectName", objectName),
		attribute.String("key", key),
		attribute.String("reason", reason),
	}
	if sourceChain != 0 {
		attrs = append(attrs, p.sourceChainAttrs(sourceChain)...)
	}
	p.bhConsensusDropped.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}
