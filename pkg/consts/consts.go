package consts

// This package contains ChainReader and ChainWriter related constants.

// Contract Names
const (
	ContractNameOffRamp = "OffRamp"
	ContractNameOnRamp  = "OnRamp"
)

// Function Names
const (
	MethodNameGetSourceChainConfig = "GetSourceChainConfig"

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
	MethodCommit = "commit"

	// On EVM:
	// function execute(bytes32[3] calldata reportContext, bytes calldata report) external
	MethodExecute = "execute"
)

// Event Names
const (
	EventNameCCIPSendRequested = "CCIPSendRequested"
)

// Event Attributes
const (
	EventAttributeSequenceNumber = "SequenceNumber"
)
