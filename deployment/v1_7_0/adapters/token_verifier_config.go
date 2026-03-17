package adapters

import (
	"fmt"
	"sync"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type TokenVerifierChainAddresses struct {
	OnRampAddress                  string
	RMNRemoteAddress               string
	CCTPVerifierAddress            string
	CCTPVerifierResolverAddress    string
	LombardVerifierResolverAddress string
}

type TokenVerifierConfigAdapter interface {
	ResolveTokenVerifierAddresses(
		ds datastore.DataStore,
		chainSelector uint64,
		cctpQualifier string,
		lombardQualifier string,
	) (*TokenVerifierChainAddresses, error)
}

type TokenVerifierConfigRegistry struct {
	mu       sync.Mutex
	adapters map[string]TokenVerifierConfigAdapter
}

var (
	singletonTokenVerifierConfigRegistry *TokenVerifierConfigRegistry
	tokenVerifierConfigRegistryOnce      sync.Once
)

func NewTokenVerifierConfigRegistry() *TokenVerifierConfigRegistry {
	return &TokenVerifierConfigRegistry{
		adapters: make(map[string]TokenVerifierConfigAdapter),
	}
}

func GetTokenVerifierConfigRegistry() *TokenVerifierConfigRegistry {
	tokenVerifierConfigRegistryOnce.Do(func() {
		singletonTokenVerifierConfigRegistry = NewTokenVerifierConfigRegistry()
	})
	return singletonTokenVerifierConfigRegistry
}

func (r *TokenVerifierConfigRegistry) Register(family string, a TokenVerifierConfigAdapter) {
	if a == nil {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.adapters[family]; !exists {
		r.adapters[family] = a
	}
}

func (r *TokenVerifierConfigRegistry) Get(family string) (TokenVerifierConfigAdapter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.adapters[family]
	return a, ok
}

func (r *TokenVerifierConfigRegistry) GetByChain(chainSelector uint64) (TokenVerifierConfigAdapter, error) {
	family, err := chainsel.GetSelectorFamily(chainSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain family for selector %d: %w", chainSelector, err)
	}
	adapter, ok := r.Get(family)
	if !ok {
		return nil, fmt.Errorf("no token verifier config adapter registered for chain family %q", family)
	}
	return adapter, nil
}
