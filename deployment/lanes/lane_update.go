package lanes

import (
	"math/big"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type ChainDefinition struct {
	// Selector is the chain selector of this chain.
	// This is provided by the user
	Selector uint64
	// GasPrice defines the USD price (18 decimals) per unit gas for this chain as a destination.
	// This is provided by the user
	// 1.6 only
	GasPrice *big.Int
	// TokenPrices define the USD price (18 decimals) per 1e18 of the smallest token denomination for various tokens on this chain.
	// This is provided by the user
	// 1.6 only
	TokenPrices map[string]*big.Int
	// FeeQuoterDestChainConfigOverrides is a functional option that mutates a
	// FeeQuoterDestChainConfig in place. Pass one or more overrides to selectively change default values.
	FeeQuoterDestChainConfigOverrides *FeeQuoterDestChainConfigOverride
	// RMNVerificationEnabled is true if we want the RMN to bless messages FROM this chain.
	// This is provided by the user
	// 1.6 only
	RMNVerificationEnabled bool
	// AllowListEnabled is true if we want an allowlist to restrict who on this chain can send outbound messages.
	// This is provided by the user
	AllowListEnabled bool
	// AllowList is the list of addresses on this chain that are allowed to send outbound messages.
	// Addresses must be in this chain's native format (e.g. hex for EVM, base58 for Solana).
	// This is provided by the user
	AllowList []string
	// CommitteeVerifiers holds fully resolved committee verifier configuration.
	// For v2.0 lanes, either this or CommitteeVerifierInputs must be populated.
	// Mutually exclusive with CommitteeVerifierInputs.
	CommitteeVerifiers []CommitteeVerifierConfig[datastore.AddressRef]
	// CommitteeVerifierInputs holds raw committee verifier configuration that
	// requires resolution (contract lookup + signing key fetch) during apply.
	// When set, ConnectChainsConfig.CommitteePopulator must also be provided.
	// Mutually exclusive with CommitteeVerifiers.
	CommitteeVerifierInputs []CommitteeVerifierInput
	// The addresses of CCVs that will be applied to messages FROM this chain if no receiver is specified.
	DefaultInboundCCVs []datastore.AddressRef
	// Addresses of any CCVs that must always be used for messages FROM this chain.
	LaneMandatedInboundCCVs []datastore.AddressRef
	// Addresses of CCVs that will be used for messages TO this chain if none are specified.
	DefaultOutboundCCVs []datastore.AddressRef
	// Addresses of CCVs that will always be applied to messages TO this chain.
	LaneMandatedOutboundCCVs []datastore.AddressRef
	// The Executor address that will be used for messages TO this chain if none is specified.
	DefaultExecutor datastore.AddressRef
	// ExecutorDestChainConfig configures the Executor for this chain
	ExecutorDestChainConfig ExecutorDestChainConfig
	// Length of addresses on the destination chain, in bytes.
	AddressBytesLength uint8
	// Execution gas cost, excluding pool/CCV/receiver gas.
	BaseExecutionGasCost uint32
	// Whether token receiver is allowed on the destination chain.
	TokenReceiverAllowed *bool
	// Message network fee in USD cents.
	MessageNetworkFeeUSDCents uint16
	// Token network fee in USD cents.
	TokenNetworkFeeUSDCents uint16
	// CantonLaneConfig holds Canton-specific configuration for lane setup.
	CantonLaneConfig *CantonLaneConfig
	// OnRamp is the address of the OnRamp contract(s) on this chain.
	// This is populated programmatically
	OnRamp []byte
	// OffRamp is the address of the OffRamp contract on this chain.
	// This is populated programmatically
	OffRamp []byte
	// Router is the address of the Router contract on this chain.
	// This is populated programmatically
	Router []byte
	// FeeQuoter is the address of the FeeQuoter contract on this chain.
	// This is populated programmatically
	FeeQuoter []byte
	// FeeQuoterDestChainConfig is the configuration that should be applied to this chain's FeeQuoter for it to be a destination in the lane.
	// This is populated programmatically and is based on the chain family, with possible overrides from the user.
	FeeQuoterDestChainConfig FeeQuoterDestChainConfig
	// FeeQuoterVersion is the contract version of the FeeQuoter (e.g. 1.6.0 or 2.0.0).
	FeeQuoterVersion *semver.Version
}

// CantonLaneConfig holds Canton-specific configuration for lane setup.
type CantonLaneConfig struct {
	// GlobalConfig is the Canton global config control contract address.
	GlobalConfig datastore.AddressRef
}

type FeeQuoterDestChainConfig struct {
	// If DestChainConfig already exists for this chain, whether to override it with the provided config.
	// If false, the existing config will be preserved and result in a noop.
	OverrideExistingConfig bool
	// Whether this destination chain is enabled.
	IsEnabled bool
	// Maximum data payload size in bytes.
	MaxDataBytes uint32
	// Maximum gas limit.
	MaxPerMsgGasLimit uint32
	// Gas charged on top of the gasLimit to cover destination chain costs.
	DestGasOverhead uint32
	// Default dest-chain gas charged for each byte of `data` payload.
	DestGasPerPayloadByteBase uint8
	// Selector that identifies the destination chain's family. Used to determine the correct validations to perform for the dest chain.
	ChainFamilySelector uint32
	// Default token fee charged per token transfer.
	DefaultTokenFeeUSDCents uint16
	// Default gas charged to execute a token transfer on the destination chain.
	DefaultTokenDestGasOverhead uint32
	// Default gas limit for a tx.
	DefaultTxGasLimit uint32
	// Flat network fee to charge for messages, multiples of 0.01 USD.
	NetworkFeeUSDCents uint16
	// V1Params holds fields specific to CCIP v1.6 FeeQuoter contracts. Populate when configuring v1.6 lanes.
	V1Params *FeeQuoterV1Params
	// V2Params holds fields specific to CCIP v2.0 FeeQuoter contracts. Populate when configuring v2.0/CCV lanes.
	V2Params *FeeQuoterV2Params
}

// FeeQuoterV1Params contains fields used only by CCIP v1.6 FeeQuoter contracts.
type FeeQuoterV1Params struct {
	MaxNumberOfTokensPerMsg           uint16
	DestGasPerPayloadByteHigh         uint8
	DestGasPerPayloadByteThreshold    uint16
	DestDataAvailabilityOverheadGas   uint32
	DestGasPerDataAvailabilityByte    uint16
	DestDataAvailabilityMultiplierBps uint16
	EnforceOutOfOrder                 bool
	GasMultiplierWeiPerEth            uint64
	GasPriceStalenessThreshold        uint32
}

// FeeQuoterV2Params contains fields used only by CCIP v2.0 (CCV) FeeQuoter contracts.
type FeeQuoterV2Params struct {
	// Percent multiplier for payments in LINK token.
	LinkFeeMultiplierPercent uint8
	// USD per unit gas for the destination chain.
	USDPerUnitGas *big.Int
}

// CommitteeVerifierSignatureQuorumConfig specifies the quorum required for any given message.
type CommitteeVerifierSignatureQuorumConfig struct {
	// Signers specifies valid signer addresses.
	Signers []string
	// Threshold specifies the number of signatures required for the message to be verified.
	Threshold uint8
}

// CommitteeVerifierRemoteChainConfig configures the CommitteeVerifier for a remote chain.
type CommitteeVerifierRemoteChainConfig struct {
	// Whether to allow traffic TO the remote chain.
	AllowlistEnabled bool
	// Addresses that are allowed to send messages TO the remote chain.
	AddedAllowlistedSenders []string
	// Addresses that are no longer allowed to send messages TO the remote chain.
	RemovedAllowlistedSenders []string
	// The fee in USD cents charged for verification on the remote chain.
	FeeUSDCents uint16
	// The gas required to execute the verification call on the destination chain (used for billing).
	GasForVerification uint32
	// The size of the CCV specific payload in bytes (used for billing).
	PayloadSizeBytes uint32
	// SignatureConfig specifies the signature configuration for the remote chain.
	SignatureConfig CommitteeVerifierSignatureQuorumConfig
}

// CommitteeVerifierConfig configures a CommitteeVerifier contract.
type CommitteeVerifierConfig[C any] struct {
	// CommitteeVerifier is a set of addresses comprising the committee verifier system.
	CommitteeVerifier []C
	// RemoteChains specifies the configuration for each remote chain supported by the committee verifier.
	RemoteChains map[uint64]CommitteeVerifierRemoteChainConfig
}

// ExecutorDestChainConfig configures the Executor for a remote chain.
type ExecutorDestChainConfig struct {
	// The fee charged by the executor to process messages to this chain.
	USDCentsFee uint16
	// Whether this destination chain is enabled.
	Enabled bool
}

type ConnectChainsConfig struct {
	Lanes []LaneConfig
	MCMS  mcms.Input
	// CommitteePopulator populates CommitteeVerifierInputs into fully configured
	// CommitteeVerifierConfig values during apply. Required when any ChainDefinition
	// in Lanes has CommitteeVerifierInputs set. Nil for 1.6-only usage.
	CommitteePopulator CommitteeConfigPopulator
}
type LaneConfig struct {
	ChainA       ChainDefinition
	ChainB       ChainDefinition
	Version      *semver.Version
	IsDisabled   bool
	TestRouter   bool
	ExtraConfigs ExtraConfigs
}

type ExtraConfigs struct {
	OnRampVersion []byte
}

type UpdateLanesInput struct {
	Source       *ChainDefinition
	Dest         *ChainDefinition
	IsDisabled   bool
	TestRouter   bool
	ExtraConfigs ExtraConfigs
}

// FeeQuoterDestChainConfigOverride is a functional option that mutates a
// FeeQuoterDestChainConfig in place. Pass one or more overrides to selectively change default values.
type FeeQuoterDestChainConfigOverride func(*FeeQuoterDestChainConfig)

// CommitteeConfigPopulator populates raw CommitteeVerifierInput values into
// fully configured CommitteeVerifierConfig values. Implementations encapsulate
// contract registry lookup, signing key fetch, and topology mapping.
type CommitteeConfigPopulator interface {
	PopulateCommitteeConfig(
		e cldf.Environment,
		chainSelector uint64,
		inputs []CommitteeVerifierInput,
	) ([]CommitteeVerifierConfig[datastore.AddressRef], error)
}

// CommitteeVerifierInput describes raw committee verifier configuration before
// resolution. The resolver transforms these into CommitteeVerifierConfig values
// by looking up contracts and computing signing quorums.
type CommitteeVerifierInput struct {
	CommitteeQualifier string
	RemoteChains       map[uint64]CommitteeVerifierRemoteChainInput
}

// CommitteeVerifierRemoteChainInput is the user-provided portion of remote
// chain config. The resolver adds SignatureConfig (signers + threshold) during
// resolution to produce CommitteeVerifierRemoteChainConfig.
type CommitteeVerifierRemoteChainInput struct {
	AllowlistEnabled          bool
	AddedAllowlistedSenders   []string
	RemovedAllowlistedSenders []string
	FeeUSDCents               uint16
	GasForVerification        uint32
	PayloadSizeBytes          uint32
}
