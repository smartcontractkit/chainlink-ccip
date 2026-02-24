package tokens

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type tokenAdapterID string

// TokenAdapter defines the interface that each chain family + token pool version combo must implement to support cross-chain token configuration.
type TokenAdapter interface {
	// ConfigureTokenForTransfersSequence returns a sequence that configures a token pool for cross-chain transfers.
	// The sequence should target a single chain, performing anything required on that chain to enable the token for CCIP transfers.
	// This assumes that the token and token pool are already deployed and registered on-chain.
	ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	// AddressRefToBytes converts an AddressRef to a byte slice representing the address.
	// Each chain family has their own way of serializing addresses from strings and needs to specify this logic.
	AddressRefToBytes(ref datastore.AddressRef) ([]byte, error)
	// DeriveTokenAddress derives the token address (in bytes) from the given token pool reference.
	// For example, if this address is stored on the pool, this method should fetch it.
	DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) ([]byte, error)
	// DeriveTokenDecimals derives the token decimals from the given token pool reference.
	DeriveTokenDecimals(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef, token []byte) (uint8, error)
	// For some chains, the token pool address is not the deployed address and must be derived from the token reference.
	// This method performs that derivation.
	DeriveTokenPoolCounterpart(e deployment.Environment, chainSelector uint64, tokenPool []byte, token []byte) ([]byte, error)
	// ManualRegistration manually registers a customer token with the token admin registry.
	// This is usally done as they no longer have mint authority over the token.
	ManualRegistration() *cldf_ops.Sequence[ManualRegistrationSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	// SetTokenPoolRateLimits returns a sequence that sets rate limits on a token pool.
	SetTokenPoolRateLimits() *cldf_ops.Sequence[TPRLRemotes, sequences.OnChainOutput, cldf_chain.BlockChains]
	DeployToken() *cldf_ops.Sequence[DeployTokenInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	DeployTokenVerify(e deployment.Environment, in any) error
	DeployTokenPoolForToken() *cldf_ops.Sequence[DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	UpdateAuthorities() *cldf_ops.Sequence[UpdateAuthoritiesInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

// RateLimiterConfig specifies configuration for a rate limiter on a token pool.
type RateLimiterConfig struct {
	// IsEnabled specifies whether the rate limiter should be enabled.
	IsEnabled bool
	// Capacity is the maximum number of tokens that can be in a rate limiter bucket.
	Capacity *big.Int
	// Rate is the rate at which the rate limiter bucket refills, in tokens per second.
	Rate *big.Int
}

// TokenTransferFeeConfig specifies configuration for a token transfer fee on a token pool.
type TokenTransferFeeConfig struct {
	// DestGasOverhead is the gas overhead for the token transfer.
	DestGasOverhead uint32
	// DestBytesOverhead is the bytes overhead for the token transfer.
	DestBytesOverhead uint32
	// DefaultFinalityFeeUSDCents is the flat fee for a default finality transfer.
	DefaultFinalityFeeUSDCents uint32
	// CustomFinalityFeeUSDCents is the flat fee for a custom finality transfer.
	CustomFinalityFeeUSDCents uint32
	// DefaultFinalityTransferFeeBps is the bps fee for a default finality transfer.
	DefaultFinalityTransferFeeBps uint16
	// CustomFinalityTransferFeeBps is the bps fee for a custom finality transfer.
	CustomFinalityTransferFeeBps uint16
	// IsEnabled is whether the token transfer fee config is enabled.
	IsEnabled bool
}

// RateLimiterConfigFloatInput is the user-friendly version of RateLimiterConfig that accepts
// float inputs for capacity and rate, which are then converted to big.Int internally after scaling by token decimals.
type RateLimiterConfigFloatInput struct {
	// IsEnabled specifies whether the rate limiter should be enabled.
	IsEnabled bool
	// Capacity is the maximum number of tokens that can be in a rate limiter bucket.
	Capacity float64
	// Rate is the rate at which the rate limiter bucket refills, in tokens per second.
	Rate float64
}

// RemoteChainConfig specifies configuration for a remote chain on a token pool.
type RemoteChainConfig[R any, CCV any] struct {
	// The token on the remote chain.
	// If not provided, the token will be derived from the pool reference.
	RemoteToken R
	// The token pool on the remote chain.
	RemotePool R
	// DefaultFinalityInboundRateLimiterConfig specifies the desired rate limiter configuration for default-finality inbound traffic.
	// DO NOT SET THIS VALUE WHEN PASSING IN INPUTS.
	// This value is derived from the configuration specified for outbound traffic to the remote chain, as the same limits should apply in both directions.
	DefaultFinalityInboundRateLimiterConfig RateLimiterConfigFloatInput
	// DefaultFinalityOutboundRateLimiterConfig specifies the desired rate limiter configuration for default-finality outbound traffic.
	DefaultFinalityOutboundRateLimiterConfig RateLimiterConfigFloatInput
	// CustomFinalityInboundRateLimiterConfig specifies the desired rate limiter configuration for custom-finality inbound traffic.
	// DO NOT SET THIS VALUE WHEN PASSING IN INPUTS.
	// This value is derived from the configuration specified for outbound traffic to the remote chain, as the same limits should apply in both directions.
	CustomFinalityInboundRateLimiterConfig RateLimiterConfigFloatInput
	// CustomFinalityOutboundRateLimiterConfig specifies the desired rate limiter configuration for custom-finality outbound traffic.
	CustomFinalityOutboundRateLimiterConfig RateLimiterConfigFloatInput
	// Decimals of the token on the remote chain.
	RemoteDecimals uint8
	// OutboundCCVs specifies the verifiers to apply to outbound traffic.
	OutboundCCVs []CCV
	// InboundCCVs specifies the verifiers to apply to inbound traffic.
	InboundCCVs []CCV
	// OutboundCCVsToAddAboveThreshold specifies the verifiers to apply to outbound traffic above the threshold.
	OutboundCCVsToAddAboveThreshold []CCV
	// InboundCCVsToAddAboveThreshold specifies the verifiers to apply to inbound traffic above the threshold.
	InboundCCVsToAddAboveThreshold []CCV
	// TokenTransferFeeConfig specifies the desired token transfer fee configuration for this remote chain.
	TokenTransferFeeConfig TokenTransferFeeConfig
}

// ConfigureTokenForTransfersInput is the input for the ConfigureTokenForTransfers sequence.
type ConfigureTokenForTransfersInput struct {
	// ChainSelector is the chain selector for the chain being configured.
	ChainSelector uint64
	// TokenAddress is the address of the token being registered and configured.
	TokenAddress string
	// TokenPoolAddress is the address of the token pool to be configured.
	TokenPoolAddress string
	// RegistryTokenPoolAddress overrides the pool address to register in the token admin registry.
	// If empty, TokenPoolAddress is used.
	RegistryTokenPoolAddress string
	// RemoteChains specifies the remote chains to configure on the token pool.
	RemoteChains map[uint64]RemoteChainConfig[[]byte, string]
	// ExternalAdmin is specified when we want to propose an admin that we don't control.
	ExternalAdmin string
	// RegistryAddress is the address of the contract on which the token pool must be registered.
	RegistryAddress string
	// MinFinalityValue is the minimum finality value required by the token pool.
	// This can be interpreted as # of block confirmations, an ID, or otherwise.
	// Interpretation is left to each chain family.
	MinFinalityValue uint16
	// Below are not provided by the user and populated programmatically.
	// ExistingDataStore is the datastore containing existing deployment data.
	ExistingDataStore datastore.DataStore
	// PoolType specifies the type of the token pool. Needed for Solana token pools.
	PoolType string
	// TokenAddress is the address of the token being configured.
	TokenRef datastore.AddressRef
}

// TokenAdapterRegistry maintains a registry of TokenAdapters.
type TokenAdapterRegistry struct {
	mu sync.Mutex
	m  map[tokenAdapterID]TokenAdapter
}

func newTokenAdapterRegistry() *TokenAdapterRegistry {
	return &TokenAdapterRegistry{
		m: make(map[tokenAdapterID]TokenAdapter),
	}
}

// RegisterTokenAdapter allows chains to register their changeset logic.
// Configuration logic not only differs by chain family, but also by version.
// For example, 1.7.0 token pools require CCV configuration, while earlier versions do not.
// 1.5.0 pools require remote pool addresses to be set, while earlier versions do not.
// Thus each version of a token pool on a chain family should have its own adapter implementation.
func (r *TokenAdapterRegistry) RegisterTokenAdapter(chainFamily string, version *semver.Version, adapter TokenAdapter) {
	id := newTokenAdapterID(chainFamily, version)
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.m[id]; exists {
		panic(fmt.Errorf("TokenAdapter '%s %s' already registered", chainFamily, version))
	}
	r.m[id] = adapter
}

// GetTokenAdapter retrieves a registered TokenAdapter for the given chain family and version.
// The boolean return value indicates whether an adapter was found.
func (r *TokenAdapterRegistry) GetTokenAdapter(chainFamily string, version *semver.Version) (TokenAdapter, bool) {
	id := newTokenAdapterID(chainFamily, version)
	r.mu.Lock()
	defer r.mu.Unlock()
	adapter, ok := r.m[id]
	return adapter, ok
}

var (
	singletonRegistry *TokenAdapterRegistry
	once              sync.Once
)

// GetTokenAdapterRegistry returns the global singleton instance.
// The first call creates the registry; subsequent calls return the same pointer.
func GetTokenAdapterRegistry() *TokenAdapterRegistry {
	once.Do(func() {
		singletonRegistry = newTokenAdapterRegistry()
	})
	return singletonRegistry
}

func newTokenAdapterID(chainFamily string, version *semver.Version) tokenAdapterID {
	return tokenAdapterID(fmt.Sprintf("%s-%s", chainFamily, version.String()))
}
