package fees

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var (
	feeAggregatorSingletonRegistry *FeeAggregatorAdapterRegistry
	feeAggregatorOnce              sync.Once
)

type feeAggregatorAdapterID string

// FeeAggregatorAdapter defines the interface for setting and reading the fee aggregator address.
// Each chain family + version must provide an implementation.
//
// On EVM 1.6, the fee aggregator is stored in the OnRamp's DynamicConfig.
// On EVM 2.0, the fee aggregator is stored on the Proxy contract.
// On Solana, the fee aggregator is stored on the Router.
type FeeAggregatorAdapter interface {
	SetFeeAggregator(e cldf.Environment) *cldf_ops.Sequence[FeeAggregatorForChain, sequences.OnChainOutput, cldf_chain.BlockChains]
	GetFeeAggregator(e cldf.Environment, chainSelector uint64) (string, error)
}

// FeeAggregatorAdapterRegistry maintains a registry of FeeAggregatorAdapters for different chain families and versions.
type FeeAggregatorAdapterRegistry struct {
	mu sync.Mutex
	m  map[feeAggregatorAdapterID]FeeAggregatorAdapter
}

func newFeeAggregatorAdapterID(chainFamily string, version *semver.Version) feeAggregatorAdapterID {
	return feeAggregatorAdapterID(fmt.Sprintf("%s-%s", chainFamily, version.String()))
}

func newFeeAggregatorAdapterRegistry() *FeeAggregatorAdapterRegistry {
	return &FeeAggregatorAdapterRegistry{
		m: make(map[feeAggregatorAdapterID]FeeAggregatorAdapter),
	}
}

// RegisterFeeAggregatorAdapter registers a new adapter.
func (r *FeeAggregatorAdapterRegistry) RegisterFeeAggregatorAdapter(chainFamily string, version *semver.Version, adapter FeeAggregatorAdapter) {
	id := newFeeAggregatorAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.m[id]; !exists {
		r.m[id] = adapter
	}
}

// GetFeeAggregatorAdapter looks up an adapter; the second return value tells you if it was found.
func (r *FeeAggregatorAdapterRegistry) GetFeeAggregatorAdapter(chainFamily string, version *semver.Version) (FeeAggregatorAdapter, bool) {
	id := newFeeAggregatorAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	adapter, ok := r.m[id]
	return adapter, ok
}

// GetFeeAggregatorRegistry returns the global singleton instance.
func GetFeeAggregatorRegistry() *FeeAggregatorAdapterRegistry {
	feeAggregatorOnce.Do(func() {
		feeAggregatorSingletonRegistry = newFeeAggregatorAdapterRegistry()
	})
	return feeAggregatorSingletonRegistry
}
