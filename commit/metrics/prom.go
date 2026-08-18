package metrics

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
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

	// Beholder equivalents of the above.
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

// used for labels values that could not be resolved.
const unknownLabelValue = "unknown"

func (p *PromReporter) trackLatestRoundID(
	latestRoundID uint32, sourceChainSelector cciptypes.ChainSelector, onramp string, method string,
) {
	// Could happen due to out of date chain-selectors lib.
	sourceFamily, sourceChainID, ok := libs.GetChainInfoFromSelector(sourceChainSelector)
	if !ok {
		// graceful fallback - we could even alert on such a thing.
		sourceFamily = unknownLabelValue
		sourceChainID = unknownLabelValue
	}
	sourceName, err := libs.GetNameFromIDAndFamily(sourceChainID, sourceFamily)
	if err != nil {
		sourceName = unknownLabelValue
	}
	destName, err := libs.GetNameFromIDAndFamily(p.chainID, p.chainFamily)
	if err != nil {
		destName = unknownLabelValue
	}

	p.commitLatestRound.WithLabelValues(sourceName, destName, onramp, method).Set(float64(latestRoundID))

	attrs := append(p.sourceChainAttrs(sourceChainSelector), p.destChainAttrs()...)
	attrs = append(attrs, attribute.String("onrampAddress", onramp), attribute.String("method", method))
	p.bhCommitLatestRound.Record(context.Background(), int64(latestRoundID), metric.WithAttributes(attrs...))
}

func (p *PromReporter) trackMaxSequenceNumber(
	sourceChainSelector cciptypes.ChainSelector,
	maxSeqNr cciptypes.SeqNum,
	method string,
) {
	if maxSeqNr == 0 {
		return
	}

	// Could happen due to out of date chain-selectors lib.
	sourceFamily, sourceChainID, ok := libs.GetChainInfoFromSelector(sourceChainSelector)
	if !ok {
		// graceful fallback - we could even alert on such a thing.
		sourceFamily = unknownLabelValue
		sourceChainID = unknownLabelValue
	}
	sourceName, err := libs.GetNameFromIDAndFamily(sourceChainID, sourceFamily)
	if err != nil {
		sourceName = unknownLabelValue
	}
	destName, err := libs.GetNameFromIDAndFamily(p.chainID, p.chainFamily)
	if err != nil {
		destName = unknownLabelValue
	}

	p.sequenceNumbers.
		WithLabelValues(p.chainFamily, p.chainID, sourceFamily, sourceChainID, method, sourceName, destName).
		Set(float64(maxSeqNr))

	attrs := append(p.sourceChainAttrs(sourceChainSelector), p.destChainAttrs()...)
	attrs = append(attrs, attribute.String("method", method))
	p.bhSequenceNumbers.Record(context.Background(), int64(maxSeqNr), metric.WithAttributes(attrs...))

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
	attrs := append(p.destChainAttrs(), attribute.String("processor", processor), attribute.String("method", method))
	p.bhProcessorLatencyHistogram.Record(context.Background(), int64(latency), metric.WithAttributes(attrs...))
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
		attrs := append(
			p.destChainAttrs(),
			attribute.String("processor", processor),
			attribute.String("method", method),
			attribute.String("type", key),
		)
		p.bhProcessorOutputCounter.Add(context.Background(), int64(val), metric.WithAttributes(attrs...))
	}
}

func (p *PromReporter) TrackLooppProviderSupported(looppCCIPProviderSupported map[string]bool) {
	for chainFamily, supported := range looppCCIPProviderSupported {
		value := float64(0)
		if supported {
			value = 1
		}
		p.looppProviderSupported.WithLabelValues(chainFamily).Set(value)
		attrs := append(p.destChainAttrs(), attribute.String("loopChainFamily", chainFamily))
		p.bhLooppProviderSupported.Record(context.Background(), int64(value), metric.WithAttributes(attrs...))
	}
}

func (p *PromReporter) TrackConfigDigestMismatch(mismatch bool) {
	var value float64
	if mismatch {
		value = 1
	}
	p.configDigestMismatch.WithLabelValues(p.chainFamily, p.chainID).Set(value)
	p.bhConfigDigestMismatch.Record(
		context.Background(),
		int64(value),
		metric.WithAttributes(p.destChainAttrs()...),
	)
}

func (p *PromReporter) TrackPluginHeartbeat(phase string) {
	attrs := append(
		p.destChainAttrs(),
		attribute.String("phase", phase),
	)
	p.bhPluginHeartbeat.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackReportValidationRejected(phase string, reason string) {
	attrs := append(
		p.destChainAttrs(),
		attribute.String("phase", phase),
		attribute.String("reason", reason),
	)
	p.bhReportValidationRejected.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// Label key constants for attributes used in this reporter - to reduce the chance of drift.
const (
	destChainIDLabelKey       = "destChainID"
	destChainFamilyLabelKey   = "destChainFamily"
	destChainNameLabelKey     = "destChainName"
	destChainSelectorLabelKey = "destChainSelector"

	sourceChainIDLabelKey       = "sourceChainID"
	sourceChainFamilyLabelKey   = "sourceChainFamily"
	sourceChainNameLabelKey     = "sourceChainName"
	sourceChainSelectorLabelKey = "sourceChainSelector"
)

// destChainAttrs returns this reporter's own (destination) chain-family/chain-ID attributes. The
// commit plugin runs one instance per destination chain, so every per-lane metric needs this to stay
// unambiguous once the same source chain is reporting into more than one destination chain on the
// same Beholder/Prometheus pipeline -- without it, two plugin instances (e.g. dest=chain-B and
// dest=chain-C) both reading source=chain-A emit identical label sets and collide into one series.
func (p *PromReporter) destChainAttrs() []attribute.KeyValue {
	// Could happen due to out of date chain-selectors lib.
	destChainInfo, err := chain_selectors.GetChainDetailsByChainIDAndFamily(p.chainID, p.chainFamily)
	if err != nil {
		// graceful fallback - we could even alert on such a thing.
		destChainInfo = chain_selectors.ChainDetails{
			ChainSelector: 0,
			ChainName:     unknownLabelValue,
		}
	}

	return []attribute.KeyValue{
		attribute.String(destChainIDLabelKey, p.chainID),
		attribute.String(destChainFamilyLabelKey, p.chainFamily),
		attribute.String(destChainNameLabelKey, destChainInfo.ChainName),
		attribute.String(destChainSelectorLabelKey, fmt.Sprintf("%d", destChainInfo.ChainSelector)),
	}
}

// sourceChainAttrs resolves a source chain selector into the same family/ID/network-name attributes
// used elsewhere in this reporter (see trackMaxSequenceNumber), so dashboards can join on network name
// consistently across metrics. Callers that identify a lane (source, but not dest) must also append
// destChainAttrs() -- see its doc comment for why.
func (p *PromReporter) sourceChainAttrs(sourceChainSelector cciptypes.ChainSelector) []attribute.KeyValue {
	// Could happen due to out of date chain-selectors lib.
	sourceFamily, sourceChainID, ok := libs.GetChainInfoFromSelector(sourceChainSelector)
	if !ok {
		// graceful fallback - we could even alert on such a thing.
		sourceFamily = unknownLabelValue
		sourceChainID = unknownLabelValue
	}

	sourceName, err := libs.GetNameFromIDAndFamily(sourceChainID, sourceFamily)
	if err != nil {
		sourceName = unknownLabelValue
	}

	return []attribute.KeyValue{
		attribute.String(sourceChainFamilyLabelKey, sourceFamily),
		attribute.String(sourceChainIDLabelKey, sourceChainID),
		attribute.String(sourceChainNameLabelKey, sourceName),
		attribute.String(sourceChainSelectorLabelKey, fmt.Sprintf("%d", sourceChainSelector)),
	}
}

func (p *PromReporter) TrackOnRampMaxSeqNum(sourceChain cciptypes.ChainSelector, seqNum cciptypes.SeqNum) {
	attrs := append(p.sourceChainAttrs(sourceChain), p.destChainAttrs()...)
	p.bhOnRampMaxSeqNum.Record(context.Background(), int64(seqNum), metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackOffRampNextSeqNum(sourceChain cciptypes.ChainSelector, seqNum cciptypes.SeqNum) {
	attrs := append(p.sourceChainAttrs(sourceChain), p.destChainAttrs()...)
	p.bhOffRampNextSeqNum.Record(context.Background(), int64(seqNum), metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackPendingMessages(sourceChain cciptypes.ChainSelector, pending uint64) {
	attrs := append(p.sourceChainAttrs(sourceChain), p.destChainAttrs()...)
	p.bhPendingMessages.Record(context.Background(), int64(pending), metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackRmnCurseActive(curseType string, active bool) {
	var value int64
	if active {
		value = 1
	}
	attrs := append(p.destChainAttrs(), attribute.String("curse_type", curseType))
	p.bhRmnCurseActive.Record(context.Background(), value, metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackSourceChainCursed(sourceChain cciptypes.ChainSelector, cursed bool) {
	var value int64
	if cursed {
		value = 1
	}
	attrs := append(p.sourceChainAttrs(sourceChain), p.destChainAttrs()...)
	p.bhSourceChainCursed.Record(context.Background(), value, metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackObservationError(sourceChain cciptypes.ChainSelector, reason string) {
	attrs := append(p.sourceChainAttrs(sourceChain), p.destChainAttrs()...)
	attrs = append(attrs, attribute.String("reason", reason))
	p.bhMerkleRootObservationErrs.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackFChainReadError() {
	p.bhFChainReadErrors.Add(context.Background(), 1, metric.WithAttributes(p.destChainAttrs()...))
}

func (p *PromReporter) TrackConsensusObservationFailed() {
	p.bhConsensusObservationFailed.Add(context.Background(), 1, metric.WithAttributes(p.destChainAttrs()...))
}

// TrackReportTransmissionGaveUp reports one source-chain lane within a report abandoned before
// transmission. sourceChain identifies which lane's seq-num range is still pending on this report --
// a single report can cover multiple source chains, and they can each resolve independently (see
// the pendingSources map in merkleroot/outcome.go); destChainAttrs identifies which destination
// chain's offramp the report itself was headed for.
func (p *PromReporter) TrackReportTransmissionGaveUp(sourceChain cciptypes.ChainSelector) {
	attrs := append(p.sourceChainAttrs(sourceChain), p.destChainAttrs()...)
	p.bhReportTransmissionGaveUp.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackReportTransmissionAttempts(attempts uint, success bool) {
	attrs := append(p.destChainAttrs(), attribute.Bool("success", success))
	p.bhReportTransmissionAttempts.Record(context.Background(), int64(attempts), metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackRangeTruncated(sourceChain cciptypes.ChainSelector) {
	attrs := append(p.sourceChainAttrs(sourceChain), p.destChainAttrs()...)
	p.bhRangeTruncated.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackOffRampLaneStatus(sourceChain cciptypes.ChainSelector, status string, active bool) {
	var value int64
	if active {
		value = 1
	}
	attrs := append(p.sourceChainAttrs(sourceChain), p.destChainAttrs()...)
	attrs = append(attrs, attribute.String("status", status))
	p.bhOffRampLaneStatus.Record(context.Background(), value, metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackSeqNumInvariantViolation(sourceChain cciptypes.ChainSelector, violationType string) {
	attrs := append(p.sourceChainAttrs(sourceChain), p.destChainAttrs()...)
	attrs = append(attrs, attribute.String("type", violationType))
	p.bhSeqNumInvariantViolation.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackOffRampConsensusInsufficient(sourceChain cciptypes.ChainSelector) {
	attrs := append(p.sourceChainAttrs(sourceChain), p.destChainAttrs()...)
	p.bhOffRampConsensusInsuff.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

func (p *PromReporter) TrackConsensusDropped(
	objectName string, key string, reason string, sourceChain cciptypes.ChainSelector,
) {
	attrs := append(p.destChainAttrs(),
		attribute.String("objectName", objectName),
		attribute.String("key", key),
		attribute.String("reason", reason))
	if sourceChain != 0 {
		attrs = append(attrs, p.sourceChainAttrs(sourceChain)...)
	}
	p.bhConsensusDropped.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}
