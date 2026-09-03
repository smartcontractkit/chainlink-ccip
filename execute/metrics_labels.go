package execute

// Label-value constants for the execute plugin's health metrics
// (docs/metrics/execute-metrics.md). Grouped here so the enums are
// discoverable and greppable instead of scattered as string literals.
const (
	// Phase-error reasons for ccip_exec_phase_errors{phase, reason}. The phase
	// label reuses plugincommon.ObservationMethod/plugincommon.OutcomeMethod.
	phaseErrGetMessages              = "get_messages"
	phaseErrGetFilter                = "get_filter"
	phaseErrFilter                   = "filter"
	phaseErrDiscovery                = "discovery"
	phaseErrCurseRead                = "curse_read"
	phaseErrNonceRead                = "nonce_read"
	phaseErrConfigDigestCheck        = "config_digest_check"
	phaseErrCommitReportCacheRefresh = "commit_report_cache_refresh"
	phaseErrHashesValidation         = "hashes_validation"
	phaseErrTokenDataValidation      = "token_data_validation"

	// objectName values for ccip_exec_consensus_dropped{objectName, ...}.
	consensusObjectFChain           = "fChain"
	consensusObjectMerkleRoot       = "merkle_root"
	consensusObjectExecutedMessages = "executed_messages"
	consensusObjectOnRampAddress    = "onramp_address"

	// Inline drop reasons for ccip_exec_consensus_dropped{..., reason}. Reasons
	// from the shared consensus package arrive via ConsensusDropReason.String().
	consensusReasonInsufficientAgreement = "insufficient_agreement"
	consensusReasonSplit                 = "split"
	consensusReasonDecodeError           = "decode_error"

	// Message-skip reasons for ccip_exec_messages_skipped{reason} raised at the
	// observation stage. Values mirror report.AlreadyInflight/report.AlreadyExecuted;
	// the report-package statuses are not imported here to avoid clashing with the
	// report loop variable in observation.go.
	skipReasonAlreadyInflight = "already_inflight"
	skipReasonAlreadyExecuted = "already_executed"

	// kind values for ccip_exec_message_consensus_conflicts{source_network_name, kind}.
	conflictKindNone         = "none"
	conflictKindMultiMessage = "multi_message"
	conflictKindOverTwo      = "over_two"
	conflictKindHash         = "hash"
)
