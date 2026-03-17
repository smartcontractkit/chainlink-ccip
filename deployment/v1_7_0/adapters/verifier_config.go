package adapters

import (
	"fmt"
	"sync"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type VerifierContractAddresses struct {
	CommitteeVerifierAddress string
	OnRampAddress            string
	ExecutorProxyAddress     string
	RMNRemoteAddress         string
}

type VerifierConfigAdapter interface {
	ResolveVerifierContractAddresses(
		ds datastore.DataStore,
		chainSelector uint64,
		committeeQualifier string,
		executorQualifier string,
	) (*VerifierContractAddresses, error)
}

type VerifierConfigRegistry struct {
	mu       sync.Mutex
	adapters map[string]VerifierConfigAdapter
}

var (
	singletonVerifierJobConfigRegistry *VerifierConfigRegistry
	verifierJobConfigRegistryOnce      sync.Once
)

func NewVerifierConfigRegistry() *VerifierConfigRegistry {
	return &VerifierConfigRegistry{
		adapters: make(map[string]VerifierConfigAdapter),
	}
}

func GetVerifierJobConfigRegistry() *VerifierConfigRegistry {
	verifierJobConfigRegistryOnce.Do(func() {
		singletonVerifierJobConfigRegistry = NewVerifierConfigRegistry()
	})
	return singletonVerifierJobConfigRegistry
}

func (r *VerifierConfigRegistry) Register(family string, a VerifierConfigAdapter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.adapters[family]; !exists {
		r.adapters[family] = a
	}
}

func (r *VerifierConfigRegistry) Get(family string) (VerifierConfigAdapter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.adapters[family]
	return a, ok
}

func (r *VerifierConfigRegistry) GetByChain(chainSelector uint64) (VerifierConfigAdapter, error) {
	family, err := chainsel.GetSelectorFamily(chainSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain family for selector %d: %w", chainSelector, err)
	}
	adapter, ok := r.Get(family)
	if !ok {
		return nil, fmt.Errorf("no verifier job config adapter registered for chain family %q", family)
	}
	return adapter, nil
}
