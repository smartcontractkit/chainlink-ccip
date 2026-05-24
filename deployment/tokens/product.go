package tokens

import (
	"errors"
	"fmt"
	"math/big"
	"sync"

	"github.com/Masterminds/semver/v3"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// ArtificialAddressRefLabel is a special label that can be used by adapters that implement
// TokenRefResolver. Adapters can add this label to reconstructed `AddressRefs` to indicate
// that the ref is artificial. It's not required to use this label - if there is no need to
// distinguish between artificial AddressRefs and datastore AddressRefs in the adapter then
// this label can be ignored.
const ArtificialAddressRefLabel = "ArtificialAddressRef"

type tokenAdapterID string

// TokenFeeAdapter is an optional interface that can be implemented by TokenAdapters to support setting token transfer fee configurations.
type TokenFeeAdapter interface {
	SetAllowedFinalityConfig(e *deployment.Environment) *cldf_ops.Sequence[SetAllowedFinalityConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	SetTokenTransferFee(e *deployment.Environment) *cldf_ops.Sequence[SetTokenTransferFeeSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	GetOnchainTokenTransferFeeConfig(e deployment.Environment, poolAddress string, src uint64, dst uint64) (TokenTransferFeeConfig, error)
	GetDefaultTokenTransferFeeConfig(src uint64, dst uint64) TokenTransferFeeConfig
}

// TokenRefResolver is an optional interface that can be implemented by TokenAdapters. It acts as a form of middleware that allows token
// and pool references to be resolved in a particular way before they are passed into the adapter logic. For example, a ref resolver can
// reconstruct refs from onchain data, normalize addresses, apply transformations on the raw input ref, etc.
type TokenRefResolver interface {
	ResolveTokenPoolRef(b cldf_ops.Bundle, chains cldf_chain.BlockChains, ds datastore.DataStore, chainSelector uint64, address string) (datastore.AddressRef, error)
	ResolveTokenRef(b cldf_ops.Bundle, chains cldf_chain.BlockChains, ds datastore.DataStore, chainSelector uint64, address string) (datastore.AddressRef, error)
}

// RateLimitReaderAdapter is an optional interface that exposes on-chain rate limit reads
// for a token pool's lane. The OutboundOnly TPRL path uses it twice: once on the local
// adapter to read the current inbound and pass it through unchanged (when the on-chain
// setter takes both directions atomically), and once on the counterpart adapter to
// validate chain B's existing inbound against chain A's new outbound.
type RateLimitReaderAdapter interface {
	// GetOnchainInboundRateLimit returns the existing on-chain inbound RateLimiterConfig
	// for the given lane (chainSelector, remoteSelector) and FastFinality bucket. Adapters
	// that do not distinguish FastFinality buckets should return an error when called with
	// fastFinality=true. If no bucket has been configured on-chain for the lane, the adapter
	// should return a zero-value RateLimiterConfig (IsEnabled=false, Capacity=0, Rate=0) and
	// no error so the caller can apply its own minimum-threshold checks.
	//
	// tokenRef is the resolved token reference for the pool on chainSelector. Some chain
	// families (Solana) need the token mint to derive the PDA holding the inbound config;
	// chain families keyed by pool address alone (EVM) may ignore it.
	GetOnchainInboundRateLimit(
		e deployment.Environment,
		chainSelector uint64,
		poolRef datastore.AddressRef,
		tokenRef datastore.AddressRef,
		remoteSelector uint64,
		fastFinality bool,
	) (RateLimiterConfig, error)
}

// TokenAdapter defines the interface that each chain family + token pool version combo must implement to support cross-chain token configuration.
type TokenAdapter interface {
	// ConfigureTokenForTransfersSequence returns a sequence that configures a token pool for cross-chain transfers.
	// The sequence should target a single chain, performing anything required on that chain to enable the token for CCIP transfers.
	// This assumes that the token and token pool are already deployed and registered on-chain.
	ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	// AddressRefToBytes converts an AddressRef to a byte slice representing the address.
	// Each chain family has their own way of serializing addresses from strings and needs to specify this logic.
	AddressRefToBytes(ref datastore.AddressRef) ([]byte, error)
	// DeriveTokenAddress derives the token address (as a string) from the given token pool reference.
	// For example, if this address is stored on the pool, this method should fetch it.
	DeriveTokenAddress(e deployment.Environment, chainSelector uint64, poolRef datastore.AddressRef) (string, error)
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
	DeployTokenVerify(e deployment.Environment, in DeployTokenInput) error
	DeployTokenPoolForToken() *cldf_ops.Sequence[DeployTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	UpdateAuthorities() *cldf_ops.Sequence[UpdateAuthoritiesInput, sequences.OnChainOutput, *deployment.Environment]
	// MigrateLockReleasePoolLiquiditySequence returns a sequence that migrates liquidity from a legacy
	// LockReleaseTokenPool (v1.5.1/v1.6.1) to a v2.0 lockbox-based pool. Returns nil if not supported.
	// Used by the standalone MigrateLockReleasePoolLiquidity changeset.
	MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

// MigrateLockReleasePoolLiquidityInput is the input for the liquidity migration sequence.
type MigrateLockReleasePoolLiquidityInput struct {
	ChainSelector  uint64
	OldPoolAddress string
	NewPoolAddress string
	// TimelockAddress is the MCMS timelock address that will execute the migration operations.
	// Required because the timelock must be set as the rebalancer and authorized caller.
	TimelockAddress string
	// Amount specifies an exact token amount to migrate. Mutually exclusive with BasisPoints.
	Amount *big.Int
	// BasisPoints specifies a percentage of the old pool's balance to migrate (1-10000, where 10000 = 100%).
	// Mutually exclusive with Amount. For siloed pools, only BasisPoints is supported.
	BasisPoints *uint16
	// SetPoolConfig, if provided, triggers a setPool call on the TokenAdminRegistry after migration.
	SetPoolConfig *MigrationSetPoolConfig
}

// MigrationSetPoolConfig configures the optional setPool call during migration.
type MigrationSetPoolConfig struct {
	RegistryAddress string
	TokenAddress    string
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

// PartialTokenTransferFeeConfig is a version of TokenTransferFeeConfig where all fields are optional. This
// is used for user input, where the user may only want to specify a subset of the fields and have the rest
// be filled in with defaults or existing on-chain values.
type PartialTokenTransferFeeConfig struct {
	DefaultFinalityTransferFeeBps utils.Optional[uint16] `yaml:"defaultFinalityTransferFeeBps" json:"defaultFinalityTransferFeeBps"`
	CustomFinalityTransferFeeBps  utils.Optional[uint16] `yaml:"customFinalityTransferFeeBps" json:"customFinalityTransferFeeBps"`
	DefaultFinalityFeeUSDCents    utils.Optional[uint32] `yaml:"defaultFinalityFeeUSDCents" json:"defaultFinalityFeeUSDCents"`
	CustomFinalityFeeUSDCents     utils.Optional[uint32] `yaml:"customFinalityFeeUSDCents" json:"customFinalityFeeUSDCents"`
	DestBytesOverhead             utils.Optional[uint32] `yaml:"destBytesOverhead" json:"destBytesOverhead"`
	DestGasOverhead               utils.Optional[uint32] `yaml:"destGasOverhead" json:"destGasOverhead"`
	IsEnabled                     utils.Optional[bool]   `yaml:"isEnabled" json:"isEnabled"`
}

// Populate fills in the fields of the PartialTokenTransferFeeConfig with values from the provided TokenTransferFeeConfig
// and returns a new PartialTokenTransferFeeConfig.
func (cfg PartialTokenTransferFeeConfig) Populate(input TokenTransferFeeConfig) PartialTokenTransferFeeConfig {
	return PartialTokenTransferFeeConfig{
		DefaultFinalityTransferFeeBps: utils.NewOptional(input.DefaultFinalityTransferFeeBps),
		CustomFinalityTransferFeeBps:  utils.NewOptional(input.CustomFinalityTransferFeeBps),
		DefaultFinalityFeeUSDCents:    utils.NewOptional(input.DefaultFinalityFeeUSDCents),
		CustomFinalityFeeUSDCents:     utils.NewOptional(input.CustomFinalityFeeUSDCents),
		DestBytesOverhead:             utils.NewOptional(input.DestBytesOverhead),
		DestGasOverhead:               utils.NewOptional(input.DestGasOverhead),
		IsEnabled:                     utils.NewOptional(input.IsEnabled),
	}
}

// MergeWith fills in the missing fields in the PartialTokenTransferFeeConfig with values from
// the provided fallbacks TokenTransferFeeConfig and returns a complete TokenTransferFeeConfig.
func (cfg PartialTokenTransferFeeConfig) MergeWith(fallbacks TokenTransferFeeConfig) TokenTransferFeeConfig {
	return TokenTransferFeeConfig{
		DefaultFinalityTransferFeeBps: cfg.DefaultFinalityTransferFeeBps.GetOrDefault(fallbacks.DefaultFinalityTransferFeeBps),
		CustomFinalityTransferFeeBps:  cfg.CustomFinalityTransferFeeBps.GetOrDefault(fallbacks.CustomFinalityTransferFeeBps),
		DefaultFinalityFeeUSDCents:    cfg.DefaultFinalityFeeUSDCents.GetOrDefault(fallbacks.DefaultFinalityFeeUSDCents),
		CustomFinalityFeeUSDCents:     cfg.CustomFinalityFeeUSDCents.GetOrDefault(fallbacks.CustomFinalityFeeUSDCents),
		DestBytesOverhead:             cfg.DestBytesOverhead.GetOrDefault(fallbacks.DestBytesOverhead),
		DestGasOverhead:               cfg.DestGasOverhead.GetOrDefault(fallbacks.DestGasOverhead),
		IsEnabled:                     cfg.IsEnabled.GetOrDefault(fallbacks.IsEnabled),
	}
}

// TokenTransferFeeConfig specifies configuration for a token transfer fee on a token pool.
type TokenTransferFeeConfig struct {
	DefaultFinalityTransferFeeBps uint16 `yaml:"defaultFinalityTransferFeeBps" json:"defaultFinalityTransferFeeBps"`
	CustomFinalityTransferFeeBps  uint16 `yaml:"customFinalityTransferFeeBps" json:"customFinalityTransferFeeBps"`
	DefaultFinalityFeeUSDCents    uint32 `yaml:"defaultFinalityFeeUSDCents" json:"defaultFinalityFeeUSDCents"`
	CustomFinalityFeeUSDCents     uint32 `yaml:"customFinalityFeeUSDCents" json:"customFinalityFeeUSDCents"`
	DestBytesOverhead             uint32 `yaml:"destBytesOverhead" json:"destBytesOverhead"`
	DestGasOverhead               uint32 `yaml:"destGasOverhead" json:"destGasOverhead"`
	IsEnabled                     bool   `yaml:"isEnabled" json:"isEnabled"`
}

// RateLimiterConfigFloatInput is the user-friendly version of RateLimiterConfig that accepts
// float inputs for capacity and rate, which are then converted to big.Int internally after scaling by token decimals.
type RateLimiterConfigFloatInput struct {
	// IsEnabled specifies whether the rate limiter should be enabled.
	IsEnabled bool `yaml:"isEnabled" json:"isEnabled"`
	// Capacity is the maximum number of tokens that can be in a rate limiter bucket.
	Capacity float64 `yaml:"capacity" json:"capacity"`
	// Rate is the rate at which the rate limiter bucket refills, in tokens per second.
	Rate float64 `yaml:"rate" json:"rate"`
}

// Validate checks the validity of the RateLimiterConfigFloatInput.
func (rl RateLimiterConfigFloatInput) Validate() error {
	// NOTE: EVM v1.5.1 token pools reject IsEnabled=true,rate=0,capacity=0
	// whereas v1.6+ pools (for most if not all chain families) treat it as
	// a valid config. We intentionally enforce a more lenient check at the
	// top-level so that an adapter can superimpose more strict checks at a
	// lower level if needed.
	if rl.IsEnabled {
		if rl.Rate < 0 || rl.Capacity < 0 {
			return errors.New("rate limiter config cannot have negative capacity or rate")
		}
		if rl.Rate > rl.Capacity {
			return errors.New("rate limiter config has rate greater than capacity")
		}
	} else {
		if rl.Capacity != 0 || rl.Rate != 0 {
			return errors.New("rate limiter config is disabled but capacity or rate is non-zero")
		}
	}

	return nil
}

// RemoteChainConfig specifies configuration for a remote chain on a token pool.
type RemoteChainConfig[R any, CCV any] struct {
	// The token on the remote chain.
	// If not provided, the token will be derived from the pool reference.
	RemoteToken R `yaml:"remoteToken" json:"remoteToken"`
	// The token pool on the remote chain.
	RemotePool R `yaml:"remotePool" json:"remotePool"`
	// InboundRateLimiterConfig specifies the desired rate limiter configuration for inbound traffic.
	// DO NOT SET THIS VALUE WHEN PASSING IN INPUTS.
	// This value is derived from the configuration specified for outbound traffic to the remote chain, as the same limits should apply in both directions.
	InboundRateLimiterConfig *RateLimiterConfigFloatInput `yaml:"inboundRateLimiterConfig,omitempty" json:"inboundRateLimiterConfig,omitempty"`
	// OutboundRateLimiterConfig specifies the desired rate limiter configuration for outbound traffic.
	// This is a backwards compatible alias for the default rate limit bucket. If OutboundRateLimits is
	// defined and it has a FastFinality=false bucket, then that bucket takes precedence as the default
	// rate limit bucket and this config is ignored. Otherwise, this config is used.
	OutboundRateLimiterConfig *RateLimiterConfigFloatInput `yaml:"outboundRateLimiterConfig,omitempty" json:"outboundRateLimiterConfig,omitempty"`
	// InboundRateLimits is populated by ConfigureTokensForTransfers from the counterpart chain's outbound buckets.
	// Do not set in YAML. Follows the same semantics as OutboundRateLimits, but for inbound traffic.
	InboundRateLimits []RateLimitConfig `yaml:"inboundRateLimits" json:"inboundRateLimits"`
	// OutboundRateLimits specifies outbound rate limits per bucket for token pools that distinguish default vs fast-finality.
	// This has higher precedence than OutboundRateLimiterConfig when a FastFinality=false bucket is present. Only v2 adapters
	// support FastFinality=true rate limits; older adapters will ignore the FastFinality field.
	OutboundRateLimits []RateLimitConfig `yaml:"outboundRateLimits" json:"outboundRateLimits"`
	// Decimals of the token on the remote chain.
	RemoteDecimals uint8 `yaml:"remoteDecimals,string" json:"remoteDecimals,string"`
	// OutboundCCVs specifies the verifiers to apply to outbound traffic.
	OutboundCCVs []CCV `yaml:"outboundCCVs" json:"outboundCCVs"`
	// InboundCCVs specifies the verifiers to apply to inbound traffic.
	InboundCCVs []CCV `yaml:"inboundCCVs" json:"inboundCCVs"`
	// OutboundCCVsToAddAboveThreshold specifies the verifiers to apply to outbound traffic above the threshold.
	OutboundCCVsToAddAboveThreshold []CCV `yaml:"outboundCCVsToAddAboveThreshold" json:"outboundCCVsToAddAboveThreshold"`
	// InboundCCVsToAddAboveThreshold specifies the verifiers to apply to inbound traffic above the threshold.
	InboundCCVsToAddAboveThreshold []CCV `yaml:"inboundCCVsToAddAboveThreshold" json:"inboundCCVsToAddAboveThreshold"`
	// TokenTransferFeeConfig specifies the desired token transfer fee configuration for this remote chain.
	TokenTransferFeeConfig *PartialTokenTransferFeeConfig `yaml:"tokenTransferFeeConfig" json:"tokenTransferFeeConfig"`
}

// GetOutboundRateLimitBuckets returns the outbound RL configuration as a RemoteOutbounds struct. The
// methods on the RemoteOutbounds struct should be used to reconcile the rate limit buckets between V2
// and older versions.
func (c RemoteChainConfig[R, CCV]) GetOutboundRateLimitBuckets() RemoteOutbounds {
	return RemoteOutbounds{RateLimit: c.OutboundRateLimiterConfig, Outbounds: c.OutboundRateLimits}
}

// GetInboundRateLimitBuckets returns the inbound RL configuration as a RemoteOutbounds struct. The
// methods on the RemoteOutbounds struct should be used to reconcile the rate limit buckets between
// V2 and older versions.
func (c RemoteChainConfig[R, CCV]) GetInboundRateLimitBuckets() RemoteOutbounds {
	return RemoteOutbounds{RateLimit: c.InboundRateLimiterConfig, Outbounds: c.InboundRateLimits}
}

// Validate checks the validity of the RemoteChainConfig, including its rate limiter configurations.
func (c RemoteChainConfig[R, CCV]) Validate() error {
	if err := c.GetOutboundRateLimitBuckets().Validate(); err != nil {
		return fmt.Errorf("outbound rate limit config for remote chain: %w", err)
	}
	if err := c.GetInboundRateLimitBuckets().Validate(); err != nil {
		return fmt.Errorf("inbound rate limit config for remote chain: %w", err)
	}
	return nil
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
	RegistryAddress       string
	AllowedFinalityConfig finality.Config
	// LiquidityMigrationAmount, if set, specifies an exact token amount to migrate from the old pool
	// to the new pool's lockbox. Mutually exclusive with LiquidityMigrationBasisPoints.
	// The old pool is derived from the TokenAdminRegistry. Only used by EVM adapters.
	LiquidityMigrationAmount *big.Int
	// LiquidityMigrationBasisPoints specifies a percentage of the old pool's balance to migrate (1-10000, where 10000 = 100%).
	// Mutually exclusive with LiquidityMigrationAmount. Only used by EVM adapters.
	LiquidityMigrationBasisPoints *uint16
	// TimelockAddress is the MCMS timelock address, resolved by the changeset from MCMS config.
	// Required when a liquidity migration is triggered.
	TimelockAddress string
	// Below are not provided by the user and populated programmatically.
	// ExistingDataStore is the datastore containing existing deployment data.
	ExistingDataStore datastore.DataStore
	// PoolType specifies the type of the token pool. Needed for Solana token pools.
	PoolType string
	// TokenAddress is the address of the token being configured.
	TokenRef datastore.AddressRef
}

// SetAllowedFinalityConfigSequenceInput defines the input for setting the allowed finality config on a V2 token pool.
type SetAllowedFinalityConfigSequenceInput struct {
	// Settings are provided as a map of pool address to finality config.
	Settings map[string]finality.Config `json:"settings" yaml:"settings"`
	// Selector is the chain selector for the chain on which to set the allowed finality configs.
	Selector uint64 `json:"selector" yaml:"selector"`
}

// SetTokenTransferFeeSequenceInput defines the input for setting token transfer fee configurations in a sequence.
type SetTokenTransferFeeSequenceInput struct {
	// Settings are provided as a map of pool address to a map of dest chain selector to fee config (use nil to disable the fee config for a dest).
	Settings map[string]map[uint64]*TokenTransferFeeConfig `json:"settings" yaml:"settings"`
	// Selector is the chain selector for the chain on which to set the token transfer fee configs.
	Selector uint64 `json:"selector" yaml:"selector"`
}

// TokenAdapterRegistry maintains a registry of TokenAdapters.
type TokenAdapterRegistry struct {
	tokenRefResolverReg map[string]TokenRefResolver
	tokenAdapterReg     map[tokenAdapterID]TokenAdapter
	tokenRefResolverMu  sync.Mutex
	tokenAdapterMu      sync.Mutex
}

func newTokenAdapterRegistry() *TokenAdapterRegistry {
	return &TokenAdapterRegistry{
		tokenRefResolverReg: make(map[string]TokenRefResolver),
		tokenAdapterReg:     make(map[tokenAdapterID]TokenAdapter),
	}
}

// Given a chain family, RegisterTokenRefResolver registers a TokenRefResolver that can resolve token and pool references for that chain family.
func (r *TokenAdapterRegistry) RegisterTokenRefResolver(chainFamily string, resolver TokenRefResolver) {
	r.tokenRefResolverMu.Lock()
	defer r.tokenRefResolverMu.Unlock()
	if _, exists := r.tokenRefResolverReg[chainFamily]; !exists {
		r.tokenRefResolverReg[chainFamily] = resolver
	}
}

// GetTokenRefResolver retrieves a registered TokenRefResolver for the given chain family.
func (r *TokenAdapterRegistry) GetTokenRefResolver(chainFamily string) (TokenRefResolver, bool) {
	r.tokenRefResolverMu.Lock()
	defer r.tokenRefResolverMu.Unlock()
	resolver, ok := r.tokenRefResolverReg[chainFamily]
	return resolver, ok
}

// RegisterTokenAdapter allows chains to register their changeset logic.
// Configuration logic not only differs by chain family, but also by version.
// For example, 2.0.0 token pools require CCV configuration, while earlier versions do not.
// 1.5.0 pools require remote pool addresses to be set, while earlier versions do not.
// Thus each version of a token pool on a chain family should have its own adapter implementation.
func (r *TokenAdapterRegistry) RegisterTokenAdapter(chainFamily string, version *semver.Version, adapter TokenAdapter) {
	id := newTokenAdapterID(chainFamily, version)
	r.tokenAdapterMu.Lock()
	defer r.tokenAdapterMu.Unlock()
	if _, exists := r.tokenAdapterReg[id]; !exists {
		r.tokenAdapterReg[id] = adapter
	}
}

// GetTokenAdapter retrieves a registered TokenAdapter for the given chain family and version.
// The boolean return value indicates whether an adapter was found.
// Returns (nil, false) if version is nil to avoid panics when token config has no version set.
func (r *TokenAdapterRegistry) GetTokenAdapter(chainFamily string, version *semver.Version) (TokenAdapter, bool) {
	if version == nil {
		return nil, false
	}
	id := newTokenAdapterID(chainFamily, version)
	r.tokenAdapterMu.Lock()
	defer r.tokenAdapterMu.Unlock()
	adapter, ok := r.tokenAdapterReg[id]
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
