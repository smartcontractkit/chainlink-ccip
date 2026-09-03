package report

import (
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Metrics is a narrow interface for report-building observability, following the same pattern
// as commit/merkleroot's local MetricsReporter: the plugin's full metrics reporter satisfies
// it, tests pass NoopMetrics, and this package stays decoupled from execute/metrics.
// Design rationale for every metric lives in docs/metrics/execute-metrics.md.
type Metrics interface {
	// TrackMessageSkipped counts a message excluded from a report with a coarse skip
	// reason (the messageStatus values below).
	TrackMessageSkipped(sourceChain cciptypes.ChainSelector, reason string)

	// TrackReportBuildError counts a commit report that died in the builder funnel
	// (malformed commit data, merkle reconstruction failures, token-data length mismatch,
	// empty-after-checks).
	TrackReportBuildError(sourceChain cciptypes.ChainSelector, reason string)

	// TrackReportRejected counts a report that was built but rejected by verifyReport
	// (skipped_nonce, size_limit, gas_limit).
	TrackReportRejected(sourceChain cciptypes.ChainSelector, reason string)
}

type NoopMetrics struct{}

func (NoopMetrics) TrackMessageSkipped(cciptypes.ChainSelector, string) {}

func (NoopMetrics) TrackReportBuildError(cciptypes.ChainSelector, string) {}

func (NoopMetrics) TrackReportRejected(cciptypes.ChainSelector, string) {}

var _ Metrics = NoopMetrics{}

// Rejection reasons for TrackReportRejected. Named constants so the label values are a
// fixed enum.
const (
	rejectReasonSkippedNonce = "skipped_nonce"
	rejectReasonSizeLimit    = "size_limit"
	rejectReasonGasLimit     = "gas_limit"
)

// Build-error reasons for TrackReportBuildError. Named constants so the label values are a
// fixed enum.
const (
	buildErrorTokenDataLengthMismatch = "token_data_length_mismatch"
	buildErrorMerkleTreeConstruction  = "merkle_tree_construction"
	buildErrorMerkleRootMismatch      = "merkle_root_mismatch"
	buildErrorProveFailed             = "prove_failed"
	buildErrorEncodeFailed            = "encode_failed"
	buildErrorEmptyReport             = "empty_report"
)
