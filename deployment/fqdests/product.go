package fqdests

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// FQDestsAdapter defines the interface for chain-family and version-specific
// FeeQuoter destination chain config operations.
type FQDestsAdapter interface {
	// ApplyDestChainConfigUpdates returns a sequence that applies destination chain config
	// updates to the FeeQuoter contract for the given chain family and version.
	ApplyDestChainConfigUpdates(e cldf.Environment) *operations.Sequence[
		ApplyDestChainConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains,
	]

	// GetOnchainDestChainConfig reads the current on-chain FeeQuoter destination config
	// for a given (source, destination) pair.
	GetOnchainDestChainConfig(e cldf.Environment, src, dst uint64) (lanes.FeeQuoterDestChainConfig, error)

	// GetDefaultDestChainConfig returns adapter-provided default destination chain config
	// for a given (source, destination) pair.
	GetDefaultDestChainConfig(src, dst uint64) lanes.FeeQuoterDestChainConfig

	// GetFeeContractRef resolves the FeeQuoter contract address reference (including version)
	// for the given source chain. The dst parameter may be unused by some implementations.
	GetFeeContractRef(e cldf.Environment, src, dst uint64) (datastore.AddressRef, error)
}

type fqDestsAdapterID string

// FQDestsAdapterRegistry maintains a registry of FQDestsAdapters keyed by chain family and version.
type FQDestsAdapterRegistry struct {
	mu sync.Mutex
	m  map[fqDestsAdapterID]FQDestsAdapter
}

func newFQDestsAdapterID(chainFamily string, version *semver.Version) fqDestsAdapterID {
	return fqDestsAdapterID(fmt.Sprintf("%s-%s", chainFamily, version.String()))
}

func newFQDestsAdapterRegistry() *FQDestsAdapterRegistry {
	return &FQDestsAdapterRegistry{
		m: make(map[fqDestsAdapterID]FQDestsAdapter),
	}
}

// RegisterFQDestsAdapter registers a new adapter; panics if the key already exists.
func (r *FQDestsAdapterRegistry) RegisterFQDestsAdapter(chainFamily string, version *semver.Version, adapter FQDestsAdapter) {
	id := newFQDestsAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.m[id]; exists {
		panic(fmt.Errorf("FQDestsAdapter '%s %s' already registered", chainFamily, version))
	}
	r.m[id] = adapter
}

// GetFQDestsAdapter looks up an adapter; the second return value tells you if it was found.
func (r *FQDestsAdapterRegistry) GetFQDestsAdapter(chainFamily string, version *semver.Version) (FQDestsAdapter, bool) {
	id := newFQDestsAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	adapter, ok := r.m[id]
	return adapter, ok
}

var (
	singletonRegistry *FQDestsAdapterRegistry
	once              sync.Once
)

// GetRegistry returns the global singleton instance.
func GetRegistry() *FQDestsAdapterRegistry {
	once.Do(func() {
		singletonRegistry = newFQDestsAdapterRegistry()
	})
	return singletonRegistry
}
