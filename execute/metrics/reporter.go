package metrics

import (
	"time"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Reporter is a simple interface used for tracking observations and outcomes of the execution plugin.
// Default implementation is based on the Prometheus metrics, but it can be extended to support other metrics systems.
// Main goal is to provide a simple way to track the performance of the execution plugin, for instance:
// - understand how efficiently we batch (number of messages, number of token data, number of source chains used etc.)
// - understand how many messages, reports, token data are observed by plugins
type Reporter interface {
	TrackObservation(obs exectypes.Observation, state exectypes.PluginState, round uint64)
	TrackOutcome(outcome exectypes.Outcome, state exectypes.PluginState, round uint64)
	TrackLatency(state exectypes.PluginState, method plugincommon.MethodType, latency time.Duration, err error)
	TrackProcessorOutput(string, plugincommon.MethodType, plugintypes.Trackable)
	TrackProcessorLatency(processor string, method plugincommon.MethodType, latency time.Duration, err error)
	TrackConfigDigestMismatch(mismatch bool)

	// TrackPluginHeartbeat is called unconditionally near the top of Observation()/Outcome(),
	// before any decode, state check, or early return. It is the execute analogue of
	// ccip_commit_plugin_heartbeat: the only signal that distinguishes "process wedged" from
	// "valid but stale" for every other metric. See docs/metrics/execute-metrics.md.
	TrackPluginHeartbeat(phase string)

	// TrackCurrentState reports the state machine state (Initialized/GetCommitReports/
	// GetMessages/Filter) once per Outcome() call. A DON frozen in one non-idle state is
	// wedged; a cycling state with an advancing heartbeat is healthy-idle.
	TrackCurrentState(state exectypes.PluginState)

	// TrackPhaseError counts a phase-level failure that is deliberately swallowed (the
	// `return nil, nil` idiom or a degrade-to-empty continuation) and therefore never trips
	// ccip_exec_errors. Reasons are a fixed coarse enum.
	TrackPhaseError(phase, reason string)

	// TrackOldestPendingCommitAge reports the age (seconds) of the oldest pending commit
	// report for a source chain, computed in TrackOutcome. Absent series = no pending work
	// (idle); growing age with an advancing heartbeat = the lane sees work it cannot finish.
	TrackOldestPendingCommitAge(sourceChain cciptypes.ChainSelector, ageSeconds float64)

	// TrackLastExecutedSeqNum reports the highest sequence number exec believes is executed
	// for a source chain, derived from ExecutedMessages in the outcome's commit data. The
	// execute twin of commit's offramp_next_seq_num gauge.
	TrackLastExecutedSeqNum(sourceChain cciptypes.ChainSelector, seqNum cciptypes.SeqNum)

	// TrackConsensusDropped reports a key dropped during consensus computation, with the
	// reason distinguishing threshold_not_defined (config bug) from insufficient_agreement
	// (routine) from split (integrity signal). Mirror of ccip_commit_consensus_dropped.
	// sourceChain may be 0 when the dropped key is not chain-keyed.
	TrackConsensusDropped(objectName, key, reason string, sourceChain cciptypes.ChainSelector)

	// TrackMessageConsensusConflict counts a message/hash that either failed to reach
	// consensus or reached conflicting consensus and was dropped from execution entirely.
	TrackMessageConsensusConflict(sourceChain cciptypes.ChainSelector, kind string)

	// TrackReportValidationRejected is called for every rejection reason in the report
	// acceptance/transmission path (ShouldAcceptAttestedReport/ShouldTransmitAcceptedReport,
	// phase="should_accept"/"should_transmit") and the build-path rejections
	// (phase="build"). Mirror of ccip_commit_report_validation_rejected.
	TrackReportValidationRejected(phase, reason string)

	// TrackPendingReports reports the per-source-chain count of commit reports still pending
	// execution after selectReports, gauged every round for every chain seen (0 included) so
	// series cannot go stale between rounds.
	TrackPendingReports(sourceChain cciptypes.ChainSelector, pending int)

	// TrackMessageReadError counts a failure to read a lane's messages from the source chain
	// (reason="rpc") or a range conformance failure (reason="range_mismatch") — the canonical
	// false-idle driver: one flaky read starves the lane across all oracles.
	TrackMessageReadError(sourceChain cciptypes.ChainSelector, reason string)

	// TrackNonceReadError counts a destination-chain nonce read failure, which empties the
	// observation's Nonces and cascades into every ordered message skipping with
	// missing_nonces_for_chain.
	TrackNonceReadError()

	// TrackMessageSkipped counts a message excluded from a report with a coarse skip reason
	// (already_executed, pseudo_deleted, token_data_not_ready, missing_nonce,
	// invalid_nonce, missing_nonces_for_chain, already_inflight, deferred).
	TrackMessageSkipped(sourceChain cciptypes.ChainSelector, reason string)

	// TrackReportBuildError counts a commit report that died in the report-builder funnel
	// (malformed commit data, merkle reconstruction failures, token-data length mismatch,
	// empty-after-checks). sourceChain may be 0 when not attributable to a lane.
	TrackReportBuildError(sourceChain cciptypes.ChainSelector, reason string)

	// TrackReportRejected counts a built report rejected by the builder's verifyReport
	// (skipped_nonce, size_limit, gas_limit). Present on Reporter so the PromReporter
	// satisfies the report package's narrow Metrics interface.
	TrackReportRejected(sourceChain cciptypes.ChainSelector, reason string)

	// TrackCommitReportCacheRefreshError counts a failed commit-report cache refresh, which
	// degrades the plugin to a wider query window silently (false idle).
	TrackCommitReportCacheRefreshError()

	// TrackCommitReportCacheRefreshAge reports seconds since the last successful commit-report
	// cache refresh. A stuck cache looks identical to "no new reports" without it.
	TrackCommitReportCacheRefreshAge(ageSeconds float64)
}

type Noop struct{}

func (n Noop) TrackObservation(exectypes.Observation, exectypes.PluginState, uint64) {}

func (n Noop) TrackOutcome(exectypes.Outcome, exectypes.PluginState, uint64) {}

func (n Noop) TrackLatency(exectypes.PluginState, plugincommon.MethodType, time.Duration, error) {}

func (n Noop) TrackProcessorOutput(string, plugincommon.MethodType, plugintypes.Trackable) {}

func (n Noop) TrackProcessorLatency(string, plugincommon.MethodType, time.Duration, error) {}

func (n Noop) TrackConfigDigestMismatch(bool) {}

func (n Noop) TrackPluginHeartbeat(string) {}

func (n Noop) TrackCurrentState(exectypes.PluginState) {}

func (n Noop) TrackPhaseError(string, string) {}

func (n Noop) TrackOldestPendingCommitAge(cciptypes.ChainSelector, float64) {}

func (n Noop) TrackLastExecutedSeqNum(cciptypes.ChainSelector, cciptypes.SeqNum) {}

func (n Noop) TrackConsensusDropped(string, string, string, cciptypes.ChainSelector) {}

func (n Noop) TrackMessageConsensusConflict(cciptypes.ChainSelector, string) {}

func (n Noop) TrackReportValidationRejected(string, string) {}

func (n Noop) TrackPendingReports(cciptypes.ChainSelector, int) {}

func (n Noop) TrackMessageReadError(cciptypes.ChainSelector, string) {}

func (n Noop) TrackNonceReadError() {}

func (n Noop) TrackMessageSkipped(cciptypes.ChainSelector, string) {}

func (n Noop) TrackReportBuildError(cciptypes.ChainSelector, string) {}

func (n Noop) TrackReportRejected(cciptypes.ChainSelector, string) {}

func (n Noop) TrackCommitReportCacheRefreshError() {}

func (n Noop) TrackCommitReportCacheRefreshAge(float64) {}

var _ Reporter = Noop{}
var _ Reporter = &Noop{}
var _ Reporter = &PromReporter{}
