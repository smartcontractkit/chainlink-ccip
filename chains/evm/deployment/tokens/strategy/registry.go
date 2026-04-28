package strategy

import (
	"sync"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// Registry holds the per-token-type strategies for each chain family.
// EVM is the only family populated today; other families would gain
// their own interface and map alongside.
type Registry struct {
	mu  sync.Mutex
	evm map[deployment.ContractType]EVMTokenStrategy
}

func newRegistry() *Registry {
	return &Registry{evm: make(map[deployment.ContractType]EVMTokenStrategy)}
}

// RegisterEVM registers a strategy for an EVM token contract type.
// First registration wins; subsequent registrations for the same
// ContractType are no-ops, matching the semantics of
// tokens.RegisterTokenAdapter.
func (r *Registry) RegisterEVM(s EVMTokenStrategy) {
	if s == nil {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.evm[s.ContractType()]; !exists {
		r.evm[s.ContractType()] = s
	}
}

// GetEVM returns the registered strategy for an EVM token contract type.
// The boolean is false when no strategy is registered.
func (r *Registry) GetEVM(ct deployment.ContractType) (EVMTokenStrategy, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	s, ok := r.evm[ct]
	return s, ok
}

// CapabilitiesEVM returns the Capabilities for an EVM token contract
// type, or the zero value if no strategy is registered. The zero value
// preserves the historical "unknown type implies all-false" predicate
// behavior at the call sites.
func (r *Registry) CapabilitiesEVM(ct deployment.ContractType) Capabilities {
	if s, ok := r.GetEVM(ct); ok {
		return s.Capabilities()
	}
	return Capabilities{}
}

var (
	singleton *Registry
	once      sync.Once
)

// GetRegistry returns the global singleton registry. The first call
// constructs the registry; subsequent calls return the same pointer.
func GetRegistry() *Registry {
	once.Do(func() { singleton = newRegistry() })
	return singleton
}
