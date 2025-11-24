package adapters

import (
	"fmt"
	"sync"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// CommitteeVerifierDestChainConfig configures the CommitteeVerifier for a remote chain.
type CommitteeVerifierDestChainConfig struct {
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
}

// ExecutorDestChainConfig configures the Executor for a remote chain.
type ExecutorDestChainConfig struct {
	// The fee charged by the executor to process messages to this chain.
	USDCentsFee uint16
	// Whether this destination chain is enabled.
	Enabled bool
}

// FeeQuoterDestChainConfig configures the FeeQuoter for a remote chain.
type FeeQuoterDestChainConfig struct {
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
	ChainFamilySelector [4]byte
	// Default token fee charged per token transfer.
	DefaultTokenFeeUSDCents uint16
	// Default gas charged to execute a token transfer on the destination chain.
	DefaultTokenDestGasOverhead uint32
	// Default gas limit for a tx.
	DefaultTxGasLimit uint32
	// Flat network fee to charge for messages, multiples of 0.01 USD.
	NetworkFeeUSDCents uint16
	// Percent multiplier for payments in LINK token.
	LinkFeeMultiplierPercent uint8
}

// RemoteChainConfig defines the configuration for a remote chain.
type RemoteChainConfig[RemoteContract any, LocalContract any] struct {
	// Whether to allow traffic FROM this remote chain.
	AllowTrafficFrom bool
	// The OnRamp address on the remote chain.
	OnRamp RemoteContract
	// The OffRamp address on the remote chain.
	OffRamp RemoteContract
	// The addresses of CCVs that will be applied to messages FROM this remote chain if no receiver is specified.
	DefaultInboundCCVs []LocalContract
	// Addresses of any CCVs that must always be used for messages FROM this remote chain.
	LaneMandatedInboundCCVs []LocalContract
	// Addresses of CCVs that will be used for messages TO this remote chain if none are specified.
	DefaultOutboundCCVs []LocalContract
	// Addresses of CCVs that will always be applied to messages TO this remote chain.
	LaneMandatedOutboundCCVs []LocalContract
	// The Executor address that will be used for messages TO this remote chain if none is specified.
	DefaultExecutor LocalContract
	// CommitteeVerifierDestChainConfig configures the CommitteeVerifier for this remote chain
	CommitteeVerifierDestChainConfig CommitteeVerifierDestChainConfig
	// FeeQuoterDestChainConfig configures the FeeQuoter for this remote chain
	FeeQuoterDestChainConfig FeeQuoterDestChainConfig
	// ExecutorDestChainConfig configures the Executor for this remote chain
	ExecutorDestChainConfig ExecutorDestChainConfig
	// Length of addresses on the destination chain, in bytes.
	AddressBytesLength uint8
	// Execution gas cost, excluding pool/CCV/receiver gas.
	BaseExecutionGasCost uint32
}

type CommitteeVerifier[Contract any] struct {
	// Resolver is the contract responsible for directing traffic to the correct CommitteeVerifier implementation.
	Resolver Contract
	// Implementation is the actual CommitteeVerifier contract.
	Implementation Contract
}

// ConfigureChainForLanesInput is the input for the ConfigureChainForLanes sequence.
type ConfigureChainForLanesInput struct {
	// The selector of the chain being configured.
	ChainSelector uint64
	// The Router address on the chain being configured.
	// We assume that all connections defined will use the same router, either test or production.
	Router string
	// The OnRamp address on the chain being configured.
	// Similarly, we assume that all connections will use the same OnRamp.
	OnRamp string
	// The CommitteeVerifier addresses on the chain being configured.
	// There can be multiple committee verifiers on a chain, each controlled by a different entity.
	CommitteeVerifiers []CommitteeVerifier[string]
	// The FeeQuoter address on the chain being configured.
	FeeQuoter string
	// The OffRamp address on the chain being configured
	OffRamp string
	// The configuration for each remote chain that we want to connect to.
	RemoteChains map[uint64]RemoteChainConfig[[]byte, string]
}

// ChainFamily is a configurable chain family.
type ChainFamily interface {
	// ConfigureChainForLanes performs all configuration required for a chain of this family to send messages to other chains.
	// The sequence should target a single chain.
	ConfigureChainForLanes() *cldf_ops.Sequence[ConfigureChainForLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	// AddressRefToBytes returns the byte representation of an address for this chain family.
	AddressRefToBytes(ref datastore.AddressRef) ([]byte, error)
}

// ChainFamilyRegistry maintains a registry of chain families.
type ChainFamilyRegistry struct {
	mu sync.Mutex
	m  map[string]ChainFamily
}

// NewChainFamilyRegistry creates a new chain family registry.
func NewChainFamilyRegistry() *ChainFamilyRegistry {
	return &ChainFamilyRegistry{
		m: make(map[string]ChainFamily),
	}
}

// RegisterChainFamily allows chain families to register their changeset logic.
func (r *ChainFamilyRegistry) RegisterChainFamily(chainFamily string, adapter ChainFamily) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.m[chainFamily]; exists {
		panic(fmt.Errorf("ChainFamily '%s' already registered", chainFamily))
	}
	r.m[chainFamily] = adapter
}

// GetChainFamily retrieves a registered adapter for the given chain family.
// The boolean return value indicates whether an adapter was found.
func (r *ChainFamilyRegistry) GetChainFamily(chainFamily string) (ChainFamily, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	adapter, ok := r.m[chainFamily]
	return adapter, ok
}
