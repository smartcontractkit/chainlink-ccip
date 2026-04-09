package adapters

import (
	"fmt"
	"sync"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type CommitteeVerifierContractAdapter interface {
	ResolveCommitteeVerifierContracts(
		ds datastore.DataStore,
		chainSelector uint64,
		qualifier string,
	) ([]datastore.AddressRef, error)
}

type CommitteeVerifierContractRegistry struct {
	mu       sync.Mutex
	adapters map[string]CommitteeVerifierContractAdapter
}

var (
	singletonCommitteeVerifierContractRegistry *CommitteeVerifierContractRegistry
	committeeVerifierContractRegistryOnce      sync.Once
)

func NewCommitteeVerifierContractRegistry() *CommitteeVerifierContractRegistry {
	return &CommitteeVerifierContractRegistry{
		adapters: make(map[string]CommitteeVerifierContractAdapter),
	}
}

func GetCommitteeVerifierContractRegistry() *CommitteeVerifierContractRegistry {
	committeeVerifierContractRegistryOnce.Do(func() {
		singletonCommitteeVerifierContractRegistry = NewCommitteeVerifierContractRegistry()
	})
	return singletonCommitteeVerifierContractRegistry
}

func (r *CommitteeVerifierContractRegistry) Register(family string, a CommitteeVerifierContractAdapter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.adapters[family]; !exists {
		r.adapters[family] = a
	}
}

func (r *CommitteeVerifierContractRegistry) Get(family string) (CommitteeVerifierContractAdapter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.adapters[family]
	return a, ok
}

func (r *CommitteeVerifierContractRegistry) GetByChain(chainSelector uint64) (CommitteeVerifierContractAdapter, error) {
	family, err := chainsel.GetSelectorFamily(chainSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain family for selector %d: %w", chainSelector, err)
	}
	adapter, ok := r.Get(family)
	if !ok {
		return nil, fmt.Errorf("no committee verifier contract adapter registered for chain family %q", family)
	}
	return adapter, nil
}
