package fees

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
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
	GetTokenTransferFeeConfigDefaults(src uint64, dst uint64) TokenTransferFeeArgs
	SetTokenTransferFeeConfig(ds datastore.DataStore, src uint64) *cldf_ops.Sequence[SetTokenTransferFeeConfigSequenceInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	GetTokenTransferFeeConfig(e cldf.Environment, src uint64, dst uint64, token string) (TokenTransferFeeArgs, error)
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

// RegisterFeeAdapter registers a new adapter; panics if the key already exists.
func (r *FeeAdapterRegistry) RegisterFeeAdapter(chainFamily string, version *semver.Version, adapter FeeAdapter) {
	id := newFeeAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.m[id]; exists {
		panic(fmt.Errorf("FeeAdapter '%s %s' already registered", chainFamily, version))
	}
	r.m[id] = adapter
}

// GetFeeAdapter looks up an adapter; the second return value tells you if it was found.
func (r *FeeAdapterRegistry) GetFeeAdapter(chainFamily string, version *semver.Version) (FeeAdapter, bool) {
	id := newFeeAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	adapter, ok := r.m[id]
	return adapter, ok
}

// GetFeeAdapterRegistry returns the global singleton instance.
func GetFeeAdapterRegistry() *FeeAdapterRegistry {
	once.Do(func() {
		singletonRegistry = newFeeAdapterRegistry()
	})
	return singletonRegistry
}
