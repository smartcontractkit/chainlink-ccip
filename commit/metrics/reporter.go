package metrics

import (
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/committypes"
	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
)

// Reporter is a simple interface used for tracking observations and outcomes of the commit plugin.
// Default implementation is based on the Prometheus metrics, but it can be extended to support other metrics systems.
// It allows you to track observation/outcome on the processor level as well as on the individual plugin level.
// That gives us more flexibility and granularity in tracking the performance of the commit plugin.
// Processors have a dedicated sub-interfaces covering only the relevant methods for reporting, please see:
// - merkleroot.MetricsReporter
// - CommitPluginReporter
// This split is required to define the reporting logic in one place but inject only relevant dependencies to
// plugins/processors. Also, it solves the problem of cyclic dependencies between the plugins/processors.
type Reporter interface {
	TrackObservation(obs committypes.Observation, round uint64)
	TrackOutcome(outcome committypes.Outcome, round uint64)

	TrackRmnReport(latency float64, success bool)
	TrackRmnRequest(method string, latency float64, nodeID uint64, err string)

	TrackProcessorLatency(processor string, method plugincommon.MethodType, latency time.Duration, err error)
	TrackProcessorOutput(processor string, method plugincommon.MethodType, obs plugintypes.Trackable)

	TrackConfigDigestMismatch(mismatch bool)

	// TrackPluginHeartbeat is called unconditionally near the top of Observation()/Outcome(), before
	// anything downstream can fail, so it keeps ticking even if the rest of the round is wedged. Absence
	// of this metric (not a bad value) is the signal: it means the plugin isn't scheduling rounds at all.
	TrackPluginHeartbeat(phase string)

	// TrackReportValidationRejected is called for every rejection reason inside validateReport, from both
	// ShouldAcceptAttestedReport (phase="should_accept") and ShouldTransmitAcceptedReport
	// (phase="should_transmit"). This is what lets an on-call engineer tell "nobody ever tried to transmit
	// because every oracle's local validation rejected it" apart from "a tx was submitted and reverted".
	TrackReportValidationRejected(phase string, reason string)

	// The methods below mirror merkleroot.MetricsReporter (see that type for docs on each). They're
	// duplicated here rather than embedded, matching the existing convention for TrackRmnReport etc.,
	// because Reporter is what's passed down to merkleroot.NewProcessor and Go requires the static
	// interface type doing the passing to already declare every method the narrower interface needs.
	TrackOnRampMaxSeqNum(sourceChain cciptypes.ChainSelector, seqNum cciptypes.SeqNum)
	TrackOffRampNextSeqNum(sourceChain cciptypes.ChainSelector, seqNum cciptypes.SeqNum)
	TrackPendingMessages(sourceChain cciptypes.ChainSelector, pending uint64)
	TrackRmnCurseActive(curseType string, active bool)
	TrackSourceChainCursed(sourceChain cciptypes.ChainSelector, cursed bool)
	TrackObservationError(sourceChain cciptypes.ChainSelector, reason string)
	TrackFChainReadError()
	TrackConsensusObservationFailed()
	TrackReportTransmissionGaveUp(sourceChain cciptypes.ChainSelector)
	TrackReportTransmissionAttempts(attempts uint, success bool)
	TrackRangeTruncated(sourceChain cciptypes.ChainSelector)
	TrackOffRampLaneStatus(sourceChain cciptypes.ChainSelector, status string, active bool)
	TrackSeqNumInvariantViolation(sourceChain cciptypes.ChainSelector, violationType string)
	TrackOffRampConsensusInsufficient(sourceChain cciptypes.ChainSelector)
}

type CommitPluginReporter interface {
	TrackObservation(obs committypes.Observation, round uint64)
	TrackOutcome(outcome committypes.Outcome, round uint64)
	TrackConfigDigestMismatch(mismatch bool)
	TrackPluginHeartbeat(phase string)
	TrackReportValidationRejected(phase string, reason string)
}

type Noop struct{}

func (n *Noop) TrackObservation(committypes.Observation, uint64) {}

func (n *Noop) TrackOutcome(committypes.Outcome, uint64) {}

func (n *Noop) TrackRmnReport(float64, bool) {}

func (n *Noop) TrackRmnRequest(string, float64, uint64, string) {}

func (n *Noop) TrackProcessorLatency(string, plugincommon.MethodType, time.Duration, error) {}

func (n *Noop) TrackProcessorOutput(string, plugincommon.MethodType, plugintypes.Trackable) {}

func (n *Noop) TrackConfigDigestMismatch(bool) {}

func (n *Noop) TrackPluginHeartbeat(string) {}

func (n *Noop) TrackReportValidationRejected(string, string) {}

func (n *Noop) TrackOnRampMaxSeqNum(cciptypes.ChainSelector, cciptypes.SeqNum) {}

func (n *Noop) TrackOffRampNextSeqNum(cciptypes.ChainSelector, cciptypes.SeqNum) {}

func (n *Noop) TrackPendingMessages(cciptypes.ChainSelector, uint64) {}

func (n *Noop) TrackRmnCurseActive(string, bool) {}

func (n *Noop) TrackSourceChainCursed(cciptypes.ChainSelector, bool) {}

func (n *Noop) TrackObservationError(cciptypes.ChainSelector, string) {}

func (n *Noop) TrackFChainReadError() {}

func (n *Noop) TrackConsensusObservationFailed() {}

func (n *Noop) TrackReportTransmissionGaveUp(cciptypes.ChainSelector) {}

func (n *Noop) TrackReportTransmissionAttempts(uint, bool) {}

func (n *Noop) TrackRangeTruncated(cciptypes.ChainSelector) {}

func (n *Noop) TrackOffRampLaneStatus(cciptypes.ChainSelector, string, bool) {}

func (n *Noop) TrackSeqNumInvariantViolation(cciptypes.ChainSelector, string) {}

func (n *Noop) TrackOffRampConsensusInsufficient(cciptypes.ChainSelector) {}

var _ Reporter = &PromReporter{}
var _ CommitPluginReporter = &PromReporter{}
var _ merkleroot.MetricsReporter = &PromReporter{}
