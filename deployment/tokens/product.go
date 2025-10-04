package tokens

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type tokenAdapterID string

func newTokenAdapterID(chainFamily string, version *semver.Version) tokenAdapterID {
	return tokenAdapterID(fmt.Sprintf("%s-%s", chainFamily, version.String()))
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

// RemoteChainConfig specifies configuration for a remote chain on a token pool.
type RemoteChainConfig[R any, CCV any] struct {
	// The token on the remote chain.
	RemoteToken R
	// The token pool on the remote chain.
	RemotePool R
	// InboundRateLimiterConfig specifies the desired rate limiter configuration for inbound traffic.
	InboundRateLimiterConfig RateLimiterConfig
	// OutboundRateLimiterConfig specifies the desired rate limiter configuration for outbound traffic.
	OutboundRateLimiterConfig RateLimiterConfig
	// OutboundCCVs specifies the verifiers to apply to outbound traffic.
	OutboundCCVs []CCV
	// InboundCCVs specifies the verifiers to apply to inbound traffic.
	InboundCCVs []CCV
}

// ConfigureTokenForTransfersInput is the input for the ConfigureTokenForTransfers sequence.
type ConfigureTokenForTransfersInput struct {
	// ChainSelector is the chain selector for the chain being configured.
	ChainSelector uint64
	// TokenPoolAddress is the address of the token pool to be configured.
	TokenPoolAddress string
	// RemoteChains specifies the remote chains to configure on the token pool.
	RemoteChains map[uint64]RemoteChainConfig[[]byte, string]
	// ExternalAdmin is specified when we want to propose an admin that we don't control.
	ExternalAdmin string
	// RegistryAddress is the address of the contract on which the token pool must be registered.
	RegistryAddress string
}

type TokenAdapter interface {
	ConfigureTokenForTransfersSequence() *cldf_ops.Sequence[ConfigureTokenForTransfersInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	ConvertRefToBytes(ref datastore.AddressRef) ([]byte, error)
}

var registeredTokenAdapters = make(map[tokenAdapterID]TokenAdapter)

// RegisterTokenAdapter allows chains to register their changeset logic.
// Configuration logic not only differs by chain family, but also by version.
// For example, 1.7.0 token pools require CCV configuration, while earlier versions do not.
// 1.5.0 pools require remote pool addresses to be set, while earlier versions do not.
// Thus each version of a token pool on a chain family should have its own adapter implementation.
func RegisterTokenAdapter(chainFamily string, version *semver.Version, adapter TokenAdapter) {
	id := newTokenAdapterID(chainFamily, version)
	if _, exists := registeredTokenAdapters[id]; exists {
		panic(fmt.Sprintf("TokenAdapter '%s' already registered", chainFamily))
	}
	registeredTokenAdapters[id] = adapter
}
