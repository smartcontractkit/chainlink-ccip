package consensus

// ConsensusDropReason explains why a key was dropped during consensus aggregation.
// It is returned alongside the consensus result so callers can route the reason to
// their own metrics reporter instead of emitting plugin-specific metrics from this
// shared package.
type ConsensusDropReason int

const (
	// DropReasonNone means the key was not dropped.
	DropReasonNone ConsensusDropReason = iota
	// DropReasonThresholdNotDefined means the caller did not supply a threshold for the key.
	DropReasonThresholdNotDefined
	// DropReasonInsufficientAgreement means no single value reached the required threshold.
	DropReasonInsufficientAgreement
	// DropReasonSplit means two or more distinct values each independently reached the threshold.
	DropReasonSplit
	// DropReasonInvalidThreshold means the configured threshold was zero or negative.
	DropReasonInvalidThreshold
)

func (r ConsensusDropReason) String() string {
	switch r {
	case DropReasonThresholdNotDefined:
		return "threshold_not_defined"
	case DropReasonInsufficientAgreement:
		return "insufficient_agreement"
	case DropReasonSplit:
		return "split"
	case DropReasonInvalidThreshold:
		return "invalid_threshold"
	default:
		return ""
	}
}
