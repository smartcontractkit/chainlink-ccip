package adapters

import (
	"fmt"
	"sync"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type ExecutorChainConfig struct {
	OffRampAddress       string
	RmnAddress           string
	ExecutorProxyAddress string
}

type ExecutorConfigAdapter interface {
	GetDeployedChains(ds datastore.DataStore, qualifier string) []uint64
	BuildChainConfig(ds datastore.DataStore, chainSelector uint64, qualifier string) (ExecutorChainConfig, error)
}

type ExecutorConfigRegistry struct {
	mu       sync.Mutex
	adapters map[string]ExecutorConfigAdapter
}

var (
	singletonExecutorConfigRegistry *ExecutorConfigRegistry
	executorConfigRegistryOnce      sync.Once
)

func NewExecutorConfigRegistry() *ExecutorConfigRegistry {
	return &ExecutorConfigRegistry{
		adapters: make(map[string]ExecutorConfigAdapter),
	}
}

func GetExecutorConfigRegistry() *ExecutorConfigRegistry {
	executorConfigRegistryOnce.Do(func() {
		singletonExecutorConfigRegistry = NewExecutorConfigRegistry()
	})
	return singletonExecutorConfigRegistry
}

func (r *ExecutorConfigRegistry) Register(family string, a ExecutorConfigAdapter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.adapters[family]; !exists {
		r.adapters[family] = a
	}
}

func (r *ExecutorConfigRegistry) Get(family string) (ExecutorConfigAdapter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.adapters[family]
	return a, ok
}

func (r *ExecutorConfigRegistry) GetByChain(chainSelector uint64) (ExecutorConfigAdapter, error) {
	family, err := chainsel.GetSelectorFamily(chainSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain family for selector %d: %w", chainSelector, err)
	}
	adapter, ok := r.Get(family)
	if !ok {
		return nil, fmt.Errorf("no executor config adapter registered for chain family %q", family)
	}
	return adapter, nil
}

// AllDeployedChains collects deployed chains across all registered adapters.
func (r *ExecutorConfigRegistry) AllDeployedChains(ds datastore.DataStore, qualifier string) []uint64 {
	r.mu.Lock()
	defer r.mu.Unlock()
	var chains []uint64
	for _, adapter := range r.adapters {
		chains = append(chains, adapter.GetDeployedChains(ds, qualifier)...)
	}
	return chains
}
