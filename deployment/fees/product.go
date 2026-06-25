package fees

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	singletonRegistry *FeeAdapterRegistry
	once              sync.Once
)

// feeResolverID is a unique identifier for a fee resolver based on chain family.
type feeResolverID string

// feeAdapterID is a unique identifier for a fee adapter based on chain family and version.
type feeAdapterID string

// FeeResolver defines the interface for fee resolvers that can infer the appropriate fee adapter version based on the chain family.
type FeeResolver interface {
	GetOnRampRef(b cldf_ops.Bundle, chains cldf_chain.BlockChains, ds datastore.DataStore, src uint64, dst uint64) (datastore.AddressRef, error)
}

// FeeAdapter defines the interface for fee adapters.
type FeeAdapter interface {
	GetFeeContractRef(b cldf_ops.Bundle, chains cldf_chain.BlockChains, ds datastore.DataStore, onRamp datastore.AddressRef, src uint64, dst uint64) (datastore.AddressRef, error)

	SetTokenTransferFee(ds datastore.DataStore, fq datastore.AddressRef) *cldf_ops.Sequence[SetTokenTransferFeeSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	GetOnchainTokenTransferFeeConfig(b cldf_ops.Bundle, chains cldf_chain.BlockChains, fq datastore.AddressRef, src uint64, dst uint64, token string) (TokenTransferFeeArgs, error)
	GetDefaultTokenTransferFeeConfig(src uint64, dst uint64) TokenTransferFeeArgs

	ApplyDestChainConfigUpdates(ds datastore.DataStore, fq datastore.AddressRef) *cldf_ops.Sequence[ApplyDestChainConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	GetOnchainDestChainConfig(b cldf_ops.Bundle, chains cldf_chain.BlockChains, fq datastore.AddressRef, src uint64, dst uint64) (lanes.FeeQuoterDestChainConfig, error)
	GetDefaultDestChainConfig(src, dst uint64) lanes.FeeQuoterDestChainConfig
}

// FeeAdapterRegistry maintains a registry of FeeAdapters for different chain families and versions.
type FeeAdapterRegistry struct {
	muResolvers sync.Mutex
	muAdapters  sync.Mutex
	resolvers   map[feeResolverID]FeeResolver
	adapters    map[feeAdapterID]FeeAdapter
}

// newFeeResolverID constructs a unique identifier for a fee resolver based on chain family.
func newFeeResolverID(chainFamily string) feeResolverID {
	return feeResolverID(chainFamily)
}

// newFeeAdapterID constructs a unique identifier for a fee adapter based on chain family and version.
func newFeeAdapterID(chainFamily string, version *semver.Version) feeAdapterID {
	return feeAdapterID(fmt.Sprintf("%s-%s", chainFamily, version.String()))
}

// newFeeAdapterRegistry creates a fresh registry.
func newFeeAdapterRegistry() *FeeAdapterRegistry {
	return &FeeAdapterRegistry{
		resolvers: make(map[feeResolverID]FeeResolver),
		adapters:  make(map[feeAdapterID]FeeAdapter),
	}
}

// RegisterFeeResolver registers a new resolver.
func (r *FeeAdapterRegistry) RegisterFeeResolver(chainFamily string, resolver FeeResolver) {
	id := newFeeResolverID(chainFamily)

	r.muResolvers.Lock()
	defer r.muResolvers.Unlock()

	if _, exists := r.resolvers[id]; !exists {
		r.resolvers[id] = resolver
	}
}

// RegisterFeeAdapter registers a new adapter.
func (r *FeeAdapterRegistry) RegisterFeeAdapter(chainFamily string, version *semver.Version, adapter FeeAdapter) {
	id := newFeeAdapterID(chainFamily, version)

	r.muAdapters.Lock()
	defer r.muAdapters.Unlock()

	if _, exists := r.adapters[id]; !exists {
		r.adapters[id] = adapter
	}
}

// GetFeeResolver looks up a fee resolver; the second return value tells you if it was found.
func (r *FeeAdapterRegistry) GetFeeResolver(chainFamily string) (FeeResolver, bool) {
	id := newFeeResolverID(chainFamily)

	r.muResolvers.Lock()
	defer r.muResolvers.Unlock()

	adapter, ok := r.resolvers[id]
	return adapter, ok
}

// GetFeeAdapter looks up an adapter; the second return value tells you if it was found.
func (r *FeeAdapterRegistry) GetFeeAdapter(chainFamily string, version *semver.Version) (FeeAdapter, bool) {
	id := newFeeAdapterID(chainFamily, version)

	r.muAdapters.Lock()
	defer r.muAdapters.Unlock()

	adapter, ok := r.adapters[id]
	return adapter, ok
}

// GetRegistry returns the global singleton instance.
func GetRegistry() *FeeAdapterRegistry {
	once.Do(func() {
		singletonRegistry = newFeeAdapterRegistry()
	})
	return singletonRegistry
}
