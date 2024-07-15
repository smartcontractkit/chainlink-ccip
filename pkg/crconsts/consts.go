package crconsts

// This package contains chainReader related constants.

// Contract Names
const (
	ContractNameOffRamp = "OffRamp"
	ContractNameOnRamp  = "OnRamp"
)

// Function Names
const (
	FunctionNameGetSourceChainConfig = "GetSourceChainConfig"
)

// Event Names
const (
	EventNameCCIPSendRequested     = "CCIPSendRequested"
	EventNameCommitReportAccepted  = "CommitReportAccepted"
	EventNameExecutionStateChanged = "ExecutionStateChanged"
)

// Event Attributes
const (
	EventAttributeSequenceNumber = "SequenceNumber"
	EventAttributeSourceChain    = "SourceChain"
	EventAttributeDestChain      = "DestChain"
)
