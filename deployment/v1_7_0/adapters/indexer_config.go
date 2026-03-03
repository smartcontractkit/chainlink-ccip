package adapters

import (
	"fmt"
	"sync"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type VerifierKind string

const (
	CommitteeVerifierKind VerifierKind = "committee"
	CCTPVerifierKind      VerifierKind = "cctp"
	LombardVerifierKind   VerifierKind = "lombard"
)

type IndexerConfigAdapter interface {
	ResolveVerifierAddresses(ds datastore.DataStore, chainSelector uint64, qualifier string, kind VerifierKind) ([]string, error)
}

type IndexerConfigRegistry struct {
	mu       sync.Mutex
	adapters map[string]IndexerConfigAdapter
}

var (
	singletonIndexerConfigRegistry *IndexerConfigRegistry
	indexerConfigRegistryOnce      sync.Once
)

func NewIndexerConfigRegistry() *IndexerConfigRegistry {
	return &IndexerConfigRegistry{
		adapters: make(map[string]IndexerConfigAdapter),
	}
}

func GetIndexerConfigRegistry() *IndexerConfigRegistry {
	indexerConfigRegistryOnce.Do(func() {
		singletonIndexerConfigRegistry = NewIndexerConfigRegistry()
	})
	return singletonIndexerConfigRegistry
}

func (r *IndexerConfigRegistry) Register(family string, a IndexerConfigAdapter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.adapters[family]; !exists {
		r.adapters[family] = a
	}
}

func (r *IndexerConfigRegistry) Get(family string) (IndexerConfigAdapter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.adapters[family]
	return a, ok
}

func (r *IndexerConfigRegistry) GetByChain(chainSelector uint64) (IndexerConfigAdapter, error) {
	family, err := chainsel.GetSelectorFamily(chainSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain family for selector %d: %w", chainSelector, err)
	}
	adapter, ok := r.Get(family)
	if !ok {
		return nil, fmt.Errorf("no indexer config adapter registered for chain family %q", family)
	}
	return adapter, nil
}
