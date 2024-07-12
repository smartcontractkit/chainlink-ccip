package consts

// This package contains ChainReader and ChainWriter related constants.

// Contract Names
const (
	ContractNameOffRamp       = "OffRamp"
	ContractNameOnRamp        = "OnRamp"
	ContractNamePriceRegistry = "PriceRegistry"
)

// Method Names
const (
	// Offramp methods
	MethodNameGetSourceChainConfig         = "GetSourceChainConfig"
	MethodNameOfframpGetDynamicConfig      = "OfframpGetDynamicConfig"
	MethodNameOfframpGetStaticConfig       = "OfframpGetStaticConfig"
	MethodNameGetLatestPriceSequenceNumber = "GetLatestPriceSequenceNumber"
	MethodNameIsBlessed                    = "IsBlessed"
	MethodNameGetMerkleRoot                = "GetMerkleRoot"
	MethodNameGetExecutionState            = "GetExecutionState"

	// Onramp methods
	MethodNameGetDestChainConfig            = "GetDestChainConfig"
	MethodNameOnrampGetDynamicConfig        = "OnrampGetDynamicConfig"
	MethodNameOnrampGetStaticConfig         = "OnrampGetStaticConfig"
	MethodNameGetExpectedNextSequenceNumber = "GetExpectedNextSequenceNumber"
	MethodNameGetPremiumMultiplierWeiPerEth = "GetPremiumMultiplierWeiPerEth"
	MethodNameGetTokenTransferFeeConfig     = "GetTokenTransferFeeConfig"

	/*
		// On EVM:
		function commit(
			bytes32[3] calldata reportContext,
			    bytes calldata report,
			    bytes32[] calldata rs,
			    bytes32[] calldata ss,
			    bytes32 rawVs // signatures
			  ) external
	*/
	MethodCommit = "Commit"

	// On EVM:
	// function execute(bytes32[3] calldata reportContext, bytes calldata report) external
	MethodExecute = "Execute"
)

// Event Names
const (
	EventNameCCIPSendRequested     = "CCIPSendRequested"
	EventNameExecutionStateChanged = "ExecutionStateChanged"
	EventNameCommitReportAccepted  = "CommitReportAccepted"
)

// Event Attributes
const (
	EventAttributeSequenceNumber = "SequenceNumber"
)
