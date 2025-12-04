package tokenpools

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type TokenPoolAdapter interface {
	// TODO: Should ds datastore.DataStore be passed in?
	ManualRegistration() *cldf_ops.Sequence[ManualRegistrationInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

type TokenPoolAdapterID string

// TokenPoolAdapterRegistry maintains a registry of TokenPoolAdapters.
type TokenPoolAdapterRegistry struct {
	mu sync.Mutex
	m  map[TokenPoolAdapterID]TokenPoolAdapter
}

// NewTokenPoolAdapterRegistry creates a fresh registry.  It is kept unexported
// because callers should obtain the singleton via GetTokenPoolAdapterRegistry().
func newTokenPoolAdapterRegistry() *TokenPoolAdapterRegistry {
	return &TokenPoolAdapterRegistry{
		m: make(map[TokenPoolAdapterID]TokenPoolAdapter),
	}
}

// RegisterTokenPoolAdapter registers a new adapter; panics if the key already exists.
func (r *TokenPoolAdapterRegistry) RegisterTokenPoolAdapter(chainFamily string, version *semver.Version, adapter TokenPoolAdapter) {
	id := newTokenPoolAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.m[id]; exists {
		panic(fmt.Errorf("TokenPoolAdapter '%s %s' already registered", chainFamily, version))
	}
	r.m[id] = adapter
}

// GetTokenPoolAdapter looks up an adapter; the second return value tells you if it was found.
func (r *TokenPoolAdapterRegistry) GetTokenPoolAdapter(chainFamily string, version *semver.Version) (TokenPoolAdapter, bool) {
	id := newTokenPoolAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	adapter, ok := r.m[id]
	return adapter, ok
}

var (
	singletonRegistry *TokenPoolAdapterRegistry
	once              sync.Once
)

// GetTokenPoolAdapterRegistry returns the global singleton instance.
// The first call creates the registry; subsequent calls return the same pointer.
func GetTokenPoolAdapterRegistry() *TokenPoolAdapterRegistry {
	once.Do(func() {
		singletonRegistry = newTokenPoolAdapterRegistry()
	})
	return singletonRegistry
}

func newTokenPoolAdapterID(chainFamily string, version *semver.Version) TokenPoolAdapterID {
	return TokenPoolAdapterID(fmt.Sprintf("%s-%s", chainFamily, version.String()))
}
