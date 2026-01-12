package deploy

import (
	"sync"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

type TransferOwnershipAdapterRegistry struct {
	mu       sync.Mutex
	adapters map[string]TransferOwnershipAdapter
}

func newTransferOwnershipRegistry() *TransferOwnershipAdapterRegistry {
	return &TransferOwnershipAdapterRegistry{
		mu:       sync.Mutex{},
		adapters: make(map[string]TransferOwnershipAdapter),
	}
}

func GetTransferOwnershipRegistry() *TransferOwnershipAdapterRegistry {
	chainAdapterOnce.Do(func() {
		singletonAdapterRegistry = newTransferOwnershipRegistry()
	})
	return singletonAdapterRegistry
}

func (r *TransferOwnershipAdapterRegistry) RegisterAdapter(chainFamily string, version *semver.Version, adapter TransferOwnershipAdapter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := utils.NewRegistererID(chainFamily, version)
	if _, exists := r.adapters[id]; !exists {
		r.adapters[id] = adapter
	}
}

func (r *TransferOwnershipAdapterRegistry) GetAdapterByChainSelector(chainSelector uint64, version *semver.Version) (TransferOwnershipAdapter, error) {
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
