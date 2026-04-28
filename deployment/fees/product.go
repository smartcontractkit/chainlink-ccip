package fees

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	singletonRegistry *FeeAdapterRegistry
	once              sync.Once
)

// feeAdapterID is a unique identifier for a fee adapter based on chain family and version.
type feeAdapterID string

// FeeAdapter defines the interface for fee adapters.
type FeeAdapter interface {
	SetTokenTransferFee(e cldf.Environment) *cldf_ops.Sequence[SetTokenTransferFeeSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	GetOnchainTokenTransferFeeConfig(e cldf.Environment, src uint64, dst uint64, token string) (TokenTransferFeeArgs, error)
	GetDefaultTokenTransferFeeConfig(src uint64, dst uint64) TokenTransferFeeArgs
	GetFeeContractRef(e cldf.Environment, src uint64, dst uint64) (datastore.AddressRef, error)

	ApplyDestChainConfigUpdates(e cldf.Environment) *cldf_ops.Sequence[ApplyDestChainConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	GetDefaultDestChainConfig(src, dst uint64) lanes.FeeQuoterDestChainConfig
	GetOnchainDestChainConfig(e cldf.Environment, src uint64, dst uint64) (lanes.FeeQuoterDestChainConfig, error)
}

// FeeAdapterRegistry maintains a registry of FeeAdapters for different chain families and versions.
type FeeAdapterRegistry struct {
	mu sync.Mutex
	m  map[feeAdapterID]FeeAdapter
}

// newFeeAdapterID constructs a unique identifier for a fee adapter based on chain family and version.
func newFeeAdapterID(chainFamily string, version *semver.Version) feeAdapterID {
	return feeAdapterID(fmt.Sprintf("%s-%s", chainFamily, version.String()))
}

// newFeeAdapterRegistry creates a fresh registry.
func newFeeAdapterRegistry() *FeeAdapterRegistry {
	return &FeeAdapterRegistry{
		m: make(map[feeAdapterID]FeeAdapter),
	}
}

// RegisterFeeAdapter registers a new adapter.
func (r *FeeAdapterRegistry) RegisterFeeAdapter(chainFamily string, version *semver.Version, adapter FeeAdapter) {
	id := newFeeAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.m[id]; !exists {
		r.m[id] = adapter
	}
}

// GetFeeAdapter looks up an adapter; the second return value tells you if it was found.
func (r *FeeAdapterRegistry) GetFeeAdapter(chainFamily string, version *semver.Version) (FeeAdapter, bool) {
	id := newFeeAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	adapter, ok := r.m[id]
	return adapter, ok
}

// GetRegistry returns the global singleton instance.
func GetRegistry() *FeeAdapterRegistry {
	once.Do(func() {
		singletonRegistry = newFeeAdapterRegistry()
	})
	return singletonRegistry
}

// ErrNoLiveLane is returned (wrapped) by a FeeContractResolver implementation
// when the (src, dst) lane has no live router-level mapping — e.g. for EVM,
// when Router.GetOnRamp(dst) yields the zero address. Callers may use
// errors.Is to detect this and fall back to a non-Router discovery path
// (e.g. the legacy FeeAdapter.GetFeeContractRef using a supplied Version)
// when configuring fees before a lane is wired.
var ErrNoLiveLane = errors.New("no live lane on the source chain's router for the given dst")

// FeeContractResolver discovers, for a given (src, dst) lane, the AddressRef
// of the contract that holds token-transfer fee config — without requiring the
// caller to know which CCIP version that lane is on. Implementations are
// registered per chain family.
type FeeContractResolver interface {
	ResolveFeeContractRef(e cldf.Environment, src uint64, dst uint64) (datastore.AddressRef, error)
}

// FeeContractResolverRegistry maintains a per-chain-family map of
// FeeContractResolvers.
type FeeContractResolverRegistry struct {
	mu sync.Mutex
	m  map[string]FeeContractResolver
}

func newFeeContractResolverRegistry() *FeeContractResolverRegistry {
	return &FeeContractResolverRegistry{
		m: make(map[string]FeeContractResolver),
	}
}

// RegisterFeeContractResolver registers the resolver for a chain family. The
// first registration for a given family wins; subsequent calls are ignored.
func (r *FeeContractResolverRegistry) RegisterFeeContractResolver(chainFamily string, resolver FeeContractResolver) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.m[chainFamily]; !exists {
		r.m[chainFamily] = resolver
	}
}

// GetFeeContractResolver returns the resolver for a chain family.
func (r *FeeContractResolverRegistry) GetFeeContractResolver(chainFamily string) (FeeContractResolver, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	resolver, ok := r.m[chainFamily]
	return resolver, ok
}

var (
	singletonResolverRegistry *FeeContractResolverRegistry
	resolverOnce              sync.Once
)

// GetFeeContractResolverRegistry returns the global singleton instance.
func GetFeeContractResolverRegistry() *FeeContractResolverRegistry {
	resolverOnce.Do(func() {
		singletonResolverRegistry = newFeeContractResolverRegistry()
	})
	return singletonResolverRegistry
}
