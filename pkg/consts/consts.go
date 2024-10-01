package consts

// This package contains ChainReader and ChainWriter related constants.

// Contract Names
const (
	ContractNameOffRamp                = "OffRamp"
	ContractNameOnRamp                 = "OnRamp"
	ContractNameFeeQuoter              = "FeeQuoter"
	ContractNameCapabilitiesRegistry   = "CapabilitiesRegistry"
	ContractNameCCIPConfig             = "CCIPConfig"
	ContractNamePriceAggregator        = "AggregatorV3Interface"
	ContractNameNonceManager           = "NonceManager"
	ContractNameRMNHome                = "RMNHome"
	ContractNameRMNRemote              = "RMNRemote"
	ContractNameRouter                 = "Router"
	ContractNameCCTPMessageTransmitter = "MessageTransmitter"
)

// Method Names
// TODO: these should be better organized, maybe separate packages.
const (
	// Router methods
	MethodNameRouterGetWrappedNative = "GetWrappedNative"

	// OffRamp methods
	MethodNameGetSourceChainConfig         = "GetSourceChainConfig"
	MethodNameOffRampGetDynamicConfig      = "OffRampGetDynamicConfig"
	MethodNameOffRampGetStaticConfig       = "OffRampGetStaticConfig"
	MethodNameOffRampGetDestChainConfig    = "OffRampGetDestChainConfig"
	MethodNameGetLatestPriceSequenceNumber = "GetLatestPriceSequenceNumber"
	MethodNameIsBlessed                    = "IsBlessed"
	MethodNameGetMerkleRoot                = "GetMerkleRoot"
	MethodNameGetExecutionState            = "GetExecutionState"

	// OnRamp methods
	MethodNameOnRampGetDynamicConfig        = "OnRampGetDynamicConfig"
	MethodNameOnRampGetStaticConfig         = "OnRampGetStaticConfig"
	MethodNameGetExpectedNextSequenceNumber = "GetExpectedNextSequenceNumber"

	// FeeQuoter view/pure methods
	MethodNameFeeQuoterGetTokenPrices       = "GetTokenPrices"
	MethodNameGetFeePriceUpdate             = "GetDestinationChainGasPrice"
	MethodNameFeeQuoterGetStaticConfig      = "GetStaticConfig"
	MethodNameGetDestChainConfig            = "GetDestChainConfig"
	MethodNameGetPremiumMultiplierWeiPerEth = "GetPremiumMultiplierWeiPerEth"
	MethodNameGetTokenTransferFeeConfig     = "GetTokenTransferFeeConfig"
	MethodNameProcessMessageArgs            = "ProcessMessageArgs"
	MethodNameProcessPoolReturnData         = "ProcessPoolReturnData"
	MethodNameGetValidatedTokenPrice        = "GetValidatedTokenPrice"
	MethodNameGetFeeTokens                  = "GetFeeTokens"

	// Aggregator methods
	MethodNameGetLatestRoundData = "latestRoundData"
	MethodNameGetDecimals        = "decimals"

	// NonceManager methods
	MethodNameGetInboundNonce  = "GetInboundNonce"
	MethodNameGetOutboundNonce = "GetOutboundNonce"

	// Deprecated: TODO: remove after chainlink is updated.
	MethodNameOfframpGetDynamicConfig = "OfframpGetDynamicConfig"
	// Deprecated: TODO: remove after chainlink is updated.
	MethodNameOfframpGetStaticConfig = "OfframpGetStaticConfig"
	// Deprecated: TODO: remove after chainlink is updated.
	MethodNameOnrampGetDynamicConfig = "OnrampGetDynamicConfig"
	// Deprecated: TODO: remove after chainlink is updated.
	MethodNameOnrampGetStaticConfig = "OnrampGetStaticConfig"

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

	// Capability registry methods.
	// Used by the home chain reader.
	MethodNameGetCapability = "GetCapability"

	// CCIPConfig.sol methods.
	// Used by the home chain reader.
	MethodNameGetAllChainConfigs = "GetAllChainConfigs"
	MethodNameGetOCRConfig       = "GetOCRConfig"

	// RMNHome.sol methods
	// Used by the rmn home reader.
	MethodNameGetAllConfigs = "GetAllConfigs"

	// RMNRemote.sol methods
	// Used by the rmn remote reader.
	MethodNameGetVersionedConfig = "GetVersionedConfig"
)

// Event Names
const (
	EventNameCCIPMessageSent       = "CCIPMessageSent"
	EventNameExecutionStateChanged = "ExecutionStateChanged"
	EventNameCommitReportAccepted  = "CommitReportAccepted"
	EventNameCCTPMessageSent       = "MessageSent"
)

// Event Attributes
const (
	EventAttributeSequenceNumber = "SequenceNumber"
	EventAttributeSourceChain    = "SourceChain"
	EventAttributeDestChain      = "destChain"
)

// Mirrors of Internal.sol's OCRPluginType
const (
	PluginTypeCommit  uint8 = 0
	PluginTypeExecute uint8 = 1
)
