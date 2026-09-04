package metrics

import (
	"context"
	"fmt"
	"maps"
	"slices"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"

	"github.com/smartcontractkit/chainlink-common/pkg/beholder"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/metricsutil"
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
	promLooppCCIPProviderSupported = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_exec_loopp_ccip_provider_supported",
		Help: "Tracks whether LOOPP CCIP provider is supported for each chain family (1 = supported, 0 = not supported)",
	}, []string{"chain_family"})
	promExecConfigDigestMismatch = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "ccip_exec_config_digest_mismatch",
		Help: "Reports whether the home chain config digest differs from the offramp config digest (1 = mismatch, 0 = match)",
	}, []string{"chain_family", "chain_id"})
)

type PromReporter struct {
	lggr          logger.Logger
	bhClient      beholder.Client
	chainSelector cciptypes.ChainSelector
	chainFamily   string
	chainID       string

	// Prometheus reporters
	latencyHistogram          *prometheus.HistogramVec
	execErrors                *prometheus.CounterVec
	outputDetailsCounter      *prometheus.CounterVec
	sequenceNumbers           *prometheus.GaugeVec
	processorLatencyHistogram *prometheus.HistogramVec
	processorErrors           *prometheus.CounterVec
	latestRoundID             *prometheus.GaugeVec
	looppProviderSupported    *prometheus.GaugeVec
	// Beholder reporters
	bhProcessorLatencyHistogram metric.Int64Histogram
	bhLatencyHistogram          metric.Int64Histogram
	bhExecErrors                metric.Int64Counter
	bhOutputDetailsCounter      metric.Int64Counter
	bhSequenceNumbers           metric.Int64Gauge
	beholderProcessorErrors     metric.Int64Counter
	bhExecLatestRound           metric.Int64Gauge
	bhLooppProviderSupported    metric.Int64Gauge

	configDigestMismatch   *prometheus.GaugeVec
	bhConfigDigestMismatch metric.Int64Gauge

	// Beholder-only metrics (no prometheus equivalent): mainnet NOP nodes aren't scraped by our
	// prometheus, only ingested via beholder, so anything meant to give us visibility into NOP-run
	// nodes should be defined here rather than as a promauto metric. Same pattern as
	// commit/metrics/prom.go; design rationale for every metric lives in docs/metrics/execute-metrics.md.
	bhPluginHeartbeat                metric.Int64Counter
	bhCurrentState                   metric.Int64Gauge
	bhPhaseErrors                    metric.Int64Counter
	bhOldestPendingCommitAge         metric.Int64Gauge
	bhLastExecutedSeqNum             metric.Int64Gauge
	bhConsensusDropped               metric.Int64Counter
	bhMessageConsensusConflicts      metric.Int64Counter
	bhReportValidationRejected       metric.Int64Counter
	bhPendingReports                 metric.Int64Gauge
	bhMessagesSkipped                metric.Int64Counter
	bhReportBuildErrors              metric.Int64Counter
	bhMessageReadErrors              metric.Int64Counter
	bhNonceReadErrors                metric.Int64Counter
	bhCommitReportCacheRefreshErrors metric.Int64Counter
	bhCommitReportCacheRefreshAge    metric.Int64Gauge
}

func NewPromReporter(
	lggr logger.Logger, selector cciptypes.ChainSelector, bhClient beholder.Client) (*PromReporter, error,
) {
	chainFamily, chainID, ok := libs.GetChainInfoFromSelector(selector)
	if !ok {
		return nil, fmt.Errorf("chainFamily and chainID not found for selector %d", selector)
	}

	instruments, err := registerBeholderExecInstruments(bhClient.Meter)
	if err != nil {
		return nil, err
	}

	return &PromReporter{
		lggr:          lggr,
		chainSelector: selector,
		chainFamily:   chainFamily,
		bhClient:      bhClient,
		chainID:       chainID,

		latencyHistogram:          PromExecLatencyHistogram,
		execErrors:                PromExecErrors,
		outputDetailsCounter:      PromExecOutputCounter,
		sequenceNumbers:           PromSequenceNumbers,
		processorLatencyHistogram: PromExecProcessorLatencyHistogram,
		processorErrors:           PromExecProcessorErrors,
		latestRoundID:             PromExecLatestRoundID,
		looppProviderSupported:    promLooppCCIPProviderSupported,
		configDigestMismatch:      promExecConfigDigestMismatch,

		bhLatencyHistogram:          instruments.latencyHistogram,
		bhProcessorLatencyHistogram: instruments.processorLatencyHistogram,
		bhExecErrors:                instruments.execErrors,
		bhOutputDetailsCounter:      instruments.outputDetailsCounter,
		bhSequenceNumbers:           instruments.sequenceNumbers,
		beholderProcessorErrors:     instruments.processorErrors,
		bhExecLatestRound:           instruments.latestRoundID,
		bhLooppProviderSupported:    instruments.looppProviderSupported,
		bhConfigDigestMismatch:      instruments.configDigestMismatch,

		bhPluginHeartbeat:                instruments.pluginHeartbeat,
		bhCurrentState:                   instruments.currentState,
		bhPhaseErrors:                    instruments.phaseErrors,
		bhOldestPendingCommitAge:         instruments.oldestPendingCommitAge,
		bhLastExecutedSeqNum:             instruments.lastExecutedSeqNum,
		bhConsensusDropped:               instruments.consensusDropped,
		bhMessageConsensusConflicts:      instruments.messageConsensusConflicts,
		bhReportValidationRejected:       instruments.reportValidationRejected,
		bhPendingReports:                 instruments.pendingReports,
		bhMessagesSkipped:                instruments.messagesSkipped,
		bhReportBuildErrors:              instruments.reportBuildErrors,
		bhMessageReadErrors:              instruments.messageReadErrors,
		bhNonceReadErrors:                instruments.nonceReadErrors,
		bhCommitReportCacheRefreshErrors: instruments.commitReportCacheRefreshErrors,
		bhCommitReportCacheRefreshAge:    instruments.commitReportCacheRefreshAge,
	}, nil
}

// beholderExecInstruments holds every beholder instrument the execute reporter records to.
// They are registered together in registerBeholderExecInstruments so that NewPromReporter
// stays under the lint complexity limit and instrument names live in one table.
type beholderExecInstruments struct {
	latencyHistogram          metric.Int64Histogram
	processorLatencyHistogram metric.Int64Histogram

	execErrors                     metric.Int64Counter
	outputDetailsCounter           metric.Int64Counter
	processorErrors                metric.Int64Counter
	pluginHeartbeat                metric.Int64Counter
	phaseErrors                    metric.Int64Counter
	consensusDropped               metric.Int64Counter
	messageConsensusConflicts      metric.Int64Counter
	reportValidationRejected       metric.Int64Counter
	messagesSkipped                metric.Int64Counter
	reportBuildErrors              metric.Int64Counter
	messageReadErrors              metric.Int64Counter
	nonceReadErrors                metric.Int64Counter
	commitReportCacheRefreshErrors metric.Int64Counter

	sequenceNumbers             metric.Int64Gauge
	latestRoundID               metric.Int64Gauge
	looppProviderSupported      metric.Int64Gauge
	configDigestMismatch        metric.Int64Gauge
	currentState                metric.Int64Gauge
	oldestPendingCommitAge      metric.Int64Gauge
	lastExecutedSeqNum          metric.Int64Gauge
	pendingReports              metric.Int64Gauge
	commitReportCacheRefreshAge metric.Int64Gauge
}

// registerBeholderInstruments registers instruments of one kind by name, failing fast on
// the first registration error.
func registerBeholderInstruments[T any](
	register func(name string) (T, error), names ...string,
) ([]T, error) {
	instruments := make([]T, 0, len(names))
	for _, name := range names {
		inst, err := register(name)
		if err != nil {
			return nil, fmt.Errorf("failed to register %s: %w", name, err)
		}
		instruments = append(instruments, inst)
	}
	return instruments, nil
}

func registerBeholderExecInstruments(meter metric.Meter) (*beholderExecInstruments, error) {
	counters, err := registerBeholderInstruments(
		func(name string) (metric.Int64Counter, error) { return meter.Int64Counter(name) },
		"ccip_exec_errors",
		"ccip_exec_output_sizes",
		"ccip_exec_processor_errors",
		"ccip_exec_plugin_heartbeat",
		"ccip_exec_phase_errors",
		"ccip_exec_consensus_dropped",
		"ccip_exec_message_consensus_conflicts",
		"ccip_exec_report_validation_rejected",
		"ccip_exec_messages_skipped",
		"ccip_exec_report_build_errors",
		"ccip_exec_message_read_errors",
		"ccip_exec_nonce_read_errors",
		"ccip_exec_commit_report_cache_refresh_errors",
	)
	if err != nil {
		return nil, err
	}
	gauges, err := registerBeholderInstruments(
		func(name string) (metric.Int64Gauge, error) { return meter.Int64Gauge(name) },
		"ccip_exec_max_sequence_number",
		"ccip_exec_latest_round_id",
		"ccip_exec_loopp_ccip_provider_supported",
		"ccip_exec_config_digest_mismatch",
		"ccip_exec_current_state",
		"ccip_exec_oldest_pending_commit_age_seconds",
		"ccip_exec_last_executed_seq_num",
		"ccip_exec_pending_reports",
		"ccip_exec_commit_report_cache_last_refresh_age_seconds",
	)
	if err != nil {
		return nil, err
	}
	histograms, err := registerBeholderInstruments(
		func(name string) (metric.Int64Histogram, error) { return meter.Int64Histogram(name) },
		"ccip_exec_latency",
		"ccip_exec_processor_latency",
	)
	if err != nil {
		return nil, err
	}

	ci, gi, hi := 0, 0, 0
	nextCounter := func() metric.Int64Counter {
		inst := counters[ci]
		ci++
		return inst
	}
	nextGauge := func() metric.Int64Gauge {
		inst := gauges[gi]
		gi++
		return inst
	}
	nextHistogram := func() metric.Int64Histogram {
		inst := histograms[hi]
		hi++
		return inst
	}

	return &beholderExecInstruments{
		latencyHistogram:          nextHistogram(),
		processorLatencyHistogram: nextHistogram(),

		execErrors:                     nextCounter(),
		outputDetailsCounter:           nextCounter(),
		processorErrors:                nextCounter(),
		pluginHeartbeat:                nextCounter(),
		phaseErrors:                    nextCounter(),
		consensusDropped:               nextCounter(),
		messageConsensusConflicts:      nextCounter(),
		reportValidationRejected:       nextCounter(),
		messagesSkipped:                nextCounter(),
		reportBuildErrors:              nextCounter(),
		messageReadErrors:              nextCounter(),
		nonceReadErrors:                nextCounter(),
		commitReportCacheRefreshErrors: nextCounter(),

		sequenceNumbers:             nextGauge(),
		latestRoundID:               nextGauge(),
		looppProviderSupported:      nextGauge(),
		configDigestMismatch:        nextGauge(),
		currentState:                nextGauge(),
		oldestPendingCommitAge:      nextGauge(),
		lastExecutedSeqNum:          nextGauge(),
		pendingReports:              nextGauge(),
		commitReportCacheRefreshAge: nextGauge(),
	}, nil
}

func (p *PromReporter) TrackObservation(obs exectypes.Observation, state exectypes.PluginState, round uint64) {
	p.trackOutputStats(obs, state, plugincommon.ObservationMethod)

	for sourceChainSelector, cr := range obs.Messages {
		maxSeqNr := pickHighestSeqNr(slices.Collect(maps.Keys(cr)))
		p.trackMaxSequenceNumber(sourceChainSelector, maxSeqNr, plugincommon.ObservationMethod)
		p.trackLatestRoundID(round, sourceChainSelector, plugincommon.OutcomeMethod)
	}
}

func (p *PromReporter) TrackOutcome(outcome exectypes.Outcome, state exectypes.PluginState, round uint64) {
	p.trackOutputStats(&outcome, state, plugincommon.OutcomeMethod)

	// Per-lane derived gauges: age of the oldest pending commit report and the highest
	// sequence number exec believes is executed. Both are computed from data already carried
	// by the outcome's commit reports; see docs/metrics/execute-metrics.md.
	oldestAgeByChain := make(map[cciptypes.ChainSelector]float64)
	lastExecutedByChain := make(map[cciptypes.ChainSelector]cciptypes.SeqNum)
	for _, cr := range outcome.CommitReports {
		sourceChainSelector := cr.SourceChain
		maxSeqNr := pickHighestSeqNrInMessages(cr.Messages)
		p.trackMaxSequenceNumber(sourceChainSelector, maxSeqNr, plugincommon.OutcomeMethod)
		p.trackLatestRoundID(round, sourceChainSelector, plugincommon.OutcomeMethod)

		if !cr.Timestamp.IsZero() {
			age := time.Since(cr.Timestamp).Seconds()
			if existing, ok := oldestAgeByChain[sourceChainSelector]; !ok || age < existing {
				oldestAgeByChain[sourceChainSelector] = age
			}
		}
		for _, seqNum := range cr.ExecutedMessages {
			if seqNum > lastExecutedByChain[sourceChainSelector] {
				lastExecutedByChain[sourceChainSelector] = seqNum
			}
		}
	}
	for sourceChainSelector, age := range oldestAgeByChain {
		p.TrackOldestPendingCommitAge(sourceChainSelector, age)
	}
	for sourceChainSelector, seqNum := range lastExecutedByChain {
		p.TrackLastExecutedSeqNum(sourceChainSelector, seqNum)
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

// TrackPluginHeartbeat is documented on Reporter.
func (p *PromReporter) TrackPluginHeartbeat(phase string) {
	attrs := append(
		metricsutil.DestChainAttrs(p.chainSelector),
		attribute.String("phase", phase),
	)
	p.bhPluginHeartbeat.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// TrackCurrentState is documented on Reporter.
func (p *PromReporter) TrackCurrentState(state exectypes.PluginState) {
	attrs := append(
		metricsutil.DestChainAttrs(p.chainSelector),
		attribute.String("state", string(state)),
	)
	p.bhCurrentState.Record(context.Background(), 1, metric.WithAttributes(attrs...))
}

// TrackPhaseError is documented on Reporter.
func (p *PromReporter) TrackPhaseError(phase, reason string) {
	attrs := append(
		metricsutil.DestChainAttrs(p.chainSelector),
		attribute.String("phase", phase),
		attribute.String("reason", reason),
	)
	p.bhPhaseErrors.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// TrackOldestPendingCommitAge is documented on Reporter.
func (p *PromReporter) TrackOldestPendingCommitAge(sourceChain cciptypes.ChainSelector, ageSeconds float64) {
	attrs := append(
		metricsutil.SourceChainAttrs(sourceChain),
		metricsutil.DestChainAttrs(p.chainSelector)...)
	p.bhOldestPendingCommitAge.Record(context.Background(), int64(ageSeconds), metric.WithAttributes(attrs...))
}

// TrackLastExecutedSeqNum is documented on Reporter.
func (p *PromReporter) TrackLastExecutedSeqNum(sourceChain cciptypes.ChainSelector, seqNum cciptypes.SeqNum) {
	attrs := append(
		metricsutil.SourceChainAttrs(sourceChain),
		metricsutil.DestChainAttrs(p.chainSelector)...)
	p.bhLastExecutedSeqNum.Record(context.Background(), int64(seqNum), metric.WithAttributes(attrs...))
}

// TrackConsensusDropped is documented on Reporter.
func (p *PromReporter) TrackConsensusDropped(
	objectName string, key string, reason string, sourceChain cciptypes.ChainSelector,
) {
	attrs := append(
		metricsutil.DestChainAttrs(p.chainSelector),
		attribute.String("objectName", objectName),
		attribute.String("key", key),
		attribute.String("reason", reason))
	if sourceChain != 0 {
		attrs = append(attrs, metricsutil.SourceChainAttrs(sourceChain)...)
	}
	p.bhConsensusDropped.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// TrackMessageConsensusConflict is documented on Reporter.
func (p *PromReporter) TrackMessageConsensusConflict(sourceChain cciptypes.ChainSelector, kind string) {
	attrs := append(
		metricsutil.SourceChainAttrs(sourceChain),
		metricsutil.DestChainAttrs(p.chainSelector)...)
	attrs = append(attrs, attribute.String("kind", kind))
	p.bhMessageConsensusConflicts.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// TrackReportValidationRejected is documented on Reporter.
func (p *PromReporter) TrackReportValidationRejected(phase, reason string) {
	attrs := append(
		metricsutil.DestChainAttrs(p.chainSelector),
		attribute.String("phase", phase),
		attribute.String("reason", reason),
	)
	p.bhReportValidationRejected.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// reportPhaseBuild is the phase label value for ccip_exec_report_validation_rejected
// rejections raised in the report builder (verifyReport), as opposed to the
// should_accept/should_transmit OCR phases.
const reportPhaseBuild = "build"

// TrackReportRejected satisfies report.Metrics: build-path rejections (verifyReport's
// skipped_nonce/size_limit/gas_limit) ride the same ccip_exec_report_validation_rejected
// family as the accept/transmit rejections, with phase="build".
func (p *PromReporter) TrackReportRejected(sourceChain cciptypes.ChainSelector, reason string) {
	attrs := append(
		metricsutil.DestChainAttrs(p.chainSelector),
		attribute.String("phase", reportPhaseBuild),
		attribute.String("reason", reason))
	if sourceChain != 0 {
		attrs = append(attrs, metricsutil.SourceChainAttrs(sourceChain)...)
	}
	p.bhReportValidationRejected.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// TrackPendingReports is documented on Reporter.
func (p *PromReporter) TrackPendingReports(sourceChain cciptypes.ChainSelector, pending int) {
	attrs := append(
		metricsutil.SourceChainAttrs(sourceChain),
		metricsutil.DestChainAttrs(p.chainSelector)...)
	p.bhPendingReports.Record(context.Background(), int64(pending), metric.WithAttributes(attrs...))
}

// TrackMessageReadError is documented on Reporter.
func (p *PromReporter) TrackMessageReadError(sourceChain cciptypes.ChainSelector, reason string) {
	attrs := append(
		metricsutil.SourceChainAttrs(sourceChain),
		metricsutil.DestChainAttrs(p.chainSelector)...)
	attrs = append(attrs, attribute.String("reason", reason))
	p.bhMessageReadErrors.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// TrackNonceReadError is documented on Reporter.
func (p *PromReporter) TrackNonceReadError() {
	p.bhNonceReadErrors.Add(
		context.Background(),
		1,
		metric.WithAttributes(metricsutil.DestChainAttrs(p.chainSelector)...))
}

// TrackMessageSkipped is documented on Reporter.
func (p *PromReporter) TrackMessageSkipped(sourceChain cciptypes.ChainSelector, reason string) {
	attrs := append(
		metricsutil.SourceChainAttrs(sourceChain),
		metricsutil.DestChainAttrs(p.chainSelector)...)
	attrs = append(attrs, attribute.String("reason", reason))
	p.bhMessagesSkipped.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// TrackReportBuildError is documented on Reporter.
func (p *PromReporter) TrackReportBuildError(sourceChain cciptypes.ChainSelector, reason string) {
	attrs := append(
		metricsutil.DestChainAttrs(p.chainSelector),
		attribute.String("reason", reason))
	if sourceChain != 0 {
		attrs = append(attrs, metricsutil.SourceChainAttrs(sourceChain)...)
	}
	p.bhReportBuildErrors.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// TrackCommitReportCacheRefreshError is documented on Reporter.
func (p *PromReporter) TrackCommitReportCacheRefreshError() {
	p.bhCommitReportCacheRefreshErrors.Add(
		context.Background(),
		1,
		metric.WithAttributes(metricsutil.DestChainAttrs(p.chainSelector)...))
}

// TrackCommitReportCacheRefreshAge is documented on Reporter.
func (p *PromReporter) TrackCommitReportCacheRefreshAge(ageSeconds float64) {
	p.bhCommitReportCacheRefreshAge.Record(
		context.Background(),
		int64(ageSeconds),
		metric.WithAttributes(metricsutil.DestChainAttrs(p.chainSelector)...))
}
