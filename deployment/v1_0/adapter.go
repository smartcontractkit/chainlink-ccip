package v1_0

import (
	"sync"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

type ChainAdapterRegistry struct {
	mu       sync.Mutex
	adapters map[string]ChainAdapter
}

func newChainAdapterRegistry() *ChainAdapterRegistry {
	return &ChainAdapterRegistry{
		mu:       sync.Mutex{},
		adapters: make(map[string]ChainAdapter),
	}
}

func GetChainAdapterRegistry() *ChainAdapterRegistry {
	chainAdapterOnce.Do(func() {
		singletonAdapterRegistry = newChainAdapterRegistry()
	})
	return singletonAdapterRegistry
}

func (r *ChainAdapterRegistry) RegisterAdapter(chainFamily string, version *semver.Version, adapter ChainAdapter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(chainFamily, version)
	if _, exists := r.adapters[id]; exists {
		panic("ChainAdapter already registered for " + id)
	}
	r.adapters[id] = adapter
}

func (r *ChainAdapterRegistry) GetAdapterByChainSelector(chainSelector uint64, version *semver.Version) (ChainAdapter, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	chainFamily, err := chain_selectors.GetSelectorFamily(chainSelector)
	if err != nil {
		return nil, err
	}
	id := utils.NewRegistererID(chainFamily, version)
	adapter, exists := r.adapters[id]
	if !exists {
		return nil, utils.ErrNoAdapterRegistered(chainFamily, version)
	}
	return adapter, nil
}
