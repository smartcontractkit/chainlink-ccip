package consts

// This package contains ChainReader and ChainWriter related constants.

// Contract Names
const (
	ContractNameOffRamp                = "OffRamp"
	ContractNameOnRamp                 = "OnRamp"
	ContractNameFeeQuoter              = "FeeQuoter"
	ContractNameCapabilitiesRegistry   = "CapabilitiesRegistry"
	ContractNameCCIPConfig             = "CCIPHome"
	ContractNamePriceAggregator        = "AggregatorV3Interface"
	ContractNameNonceManager           = "NonceManager"
	ContractNameRMNHome                = "RMNHome"
	ContractNameRMNRemote              = "RMNRemote"
	ContractNameRMNProxy               = "RMNProxy"
	ContractNameRouter                 = "Router"
	ContractNameCCTPMessageTransmitter = "MessageTransmitter"
	ContractNameUSDCTokenPool          = "USDCTokenPool" //nolint:gosec // Solana USDC Token Pool contract name
)

func AllContractNames() []string {
	return []string{
		ContractNameOffRamp,
		ContractNameOnRamp,
		ContractNameFeeQuoter,
		ContractNameCapabilitiesRegistry,
		ContractNameCCIPConfig,
		ContractNamePriceAggregator,
		ContractNameNonceManager,
		ContractNameRMNHome,
		ContractNameRMNRemote,
		ContractNameRMNProxy,
		ContractNameRouter,
		ContractNameCCTPMessageTransmitter,
	}
}

// Method Names
// TODO: these should be better organized, maybe separate packages.
const (
	// Router methods
	MethodNameRouterGetWrappedNative = "GetWrappedNative"

	// OffRamp methods
	MethodNameGetSourceChainConfig            = "GetSourceChainConfig"
	MethodNameOffRampGetAllSourceChainConfigs = "OffRampGetAllSourceChainConfigs"
	MethodNameOffRampGetDynamicConfig         = "OffRampGetDynamicConfig"
	MethodNameOffRampGetStaticConfig          = "OffRampGetStaticConfig"
	MethodNameOffRampLatestConfigDetails      = "OffRampLatestConfigDetails"
	MethodNameGetLatestPriceSequenceNumber    = "GetLatestPriceSequenceNumber"
	MethodNameGetMerkleRoot                   = "GetMerkleRoot"
	MethodNameGetExecutionState               = "GetExecutionState"

	// OnRamp methods
	MethodNameOnRampGetDynamicConfig        = "OnRampGetDynamicConfig"
	MethodNameOnRampGetStaticConfig         = "OnRampGetStaticConfig"
	MethodNameOnRampGetDestChainConfig      = "OnRampGetDestChainConfig"
	MethodNameGetExpectedNextSequenceNumber = "GetExpectedNextSequenceNumber"

	// FeeQuoter view/pure methods
	MethodNameFeeQuoterGetTokenPrices       = "GetTokenPrices"
	MethodNameFeeQuoterGetTokenPrice        = "GetTokenPrice"
	MethodNameGetFeePriceUpdate             = "GetDestinationChainGasPrice"
	MethodNameFeeQuoterGetStaticConfig      = "GetStaticConfig"
	MethodNameGetDestChainConfig            = "GetDestChainConfig"
	MethodNameGetPremiumMultiplierWeiPerEth = "GetPremiumMultiplierWeiPerEth"
	MethodNameGetTokenTransferFeeConfig     = "GetTokenTransferFeeConfig"
	MethodNameProcessMessageArgs            = "ProcessMessageArgs"
	MethodNameGetValidatedTokenPrice        = "GetValidatedTokenPrice"
	MethodNameGetFeeTokens                  = "GetFeeTokens"

	// Aggregator methods
	MethodNameGetLatestRoundData = "latestRoundData"
	MethodNameGetDecimals        = "decimals"

	// NonceManager methods
	MethodNameGetInboundNonce  = "GetInboundNonce"
	MethodNameGetOutboundNonce = "GetOutboundNonce"

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
	MethodCommit          = "Commit"
	MethodCommitPriceOnly = "CommitPriceOnly"

	// On EVM:
	// function execute(bytes32[3] calldata reportContext, bytes calldata report) external
	MethodExecute = "Execute"

	// Capability registry methods.
	// Used by the home chain reader.
	MethodNameGetCapability = "GetCapability"

	// CCIPHome.sol methods.
	// Used by the home chain reader.
	// TODO: change them to getConfig, getAllConfigs
	MethodNameGetAllChainConfigs = "GetAllChainConfigs"
	MethodNameGetOCRConfig       = "GetOCRConfig"

	// RMNHome.sol methods
	// Used by the rmn home reader.
	MethodNameGetAllConfigs = "GetAllConfigs"

	// RMNRemote.sol methods
	// Used by the rmn remote reader.
	MethodNameGetVersionedConfig    = "GetVersionedConfig"
	MethodNameGetReportDigestHeader = "GetReportDigestHeader"
	MethodNameGetCursedSubjects     = "GetCursedSubjects"

	// RMNProxy.sol methods
	MethodNameGetARM = "GetARM"
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
	EventAttributeDestChain      = "DestChain"
	EventAttributeState          = "State"

	// Required for Solana CCTP
	EventAttributeCCTPNonce    = "CctpNonce"
	EventAttributeSourceDomain = "SourceDomain"
)

// Dedicated filters
const (
	CCTPMessageSentValue = "CCTPMessageSentValue"
)

// Mirrors of Internal.sol's OCRPluginType
const (
	PluginTypeCommit  uint8 = 0
	PluginTypeExecute uint8 = 1
)

type CtxKey string

const InvalidateCacheKey = CtxKey("invalidate-cache")
