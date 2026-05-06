package authorizedcallers

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

var (
	singletonAuthorizedCallersRegistry *AuthorizedCallersRegistry
	authorizedCallersRegistryOnce      sync.Once
)

// AuthorizedCallersAdapter is the chain-agnostic interface for managing the
// authorized callers on any contract that inherits AuthorizedCallers.sol.
type AuthorizedCallersAdapter interface {
	// Initialize resolves and caches the target contract address for the given input.
	// Must be called before GetAllAuthorizedCallers or ApplyAuthorizedCallerUpdates.
	Initialize(e cldf.Environment, in ApplyInput) error

	// GetAllAuthorizedCallers returns the current set of authorized callers on the
	// target contract. Adapters should execute the generated read op against live chain
	// RPC without replaying cached ExecuteOperation reports from the environment bundle.
	GetAllAuthorizedCallers(e cldf.Environment, selector uint64, contractType cldf.ContractType, version *semver.Version) ([]Caller, error)

	// ApplyAuthorizedCallerUpdates returns the sequence that calls
	// applyAuthorizedCallerUpdates on the target contract.
	ApplyAuthorizedCallerUpdates() *cldf_ops.Sequence[ApplyInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

// AuthorizedCallersRegistry holds one AuthorizedCallersAdapter per
// (chain family, ContractType, version) triple. It is the counterpart of
// fastcurse.CurseRegistry for authorized-callers management.
type AuthorizedCallersRegistry struct {
	mu       sync.Mutex
	adapters map[string]AuthorizedCallersAdapter
}

func newAuthorizedCallersRegistry() *AuthorizedCallersRegistry {
	return &AuthorizedCallersRegistry{
		adapters: make(map[string]AuthorizedCallersAdapter),
	}
}

// GetAuthorizedCallersRegistry returns the process-global singleton registry.
// EVM adapters register themselves via init() in their respective adapters packages
// (blank-imported by changeset callers) — the same pattern as fastcurse.GetCurseRegistry().
func GetAuthorizedCallersRegistry() *AuthorizedCallersRegistry {
	authorizedCallersRegistryOnce.Do(func() {
		singletonAuthorizedCallersRegistry = newAuthorizedCallersRegistry()
	})
	return singletonAuthorizedCallersRegistry
}

// RegisterAdapter registers an AuthorizedCallersAdapter for the given
// (chain family, ContractType, version) triple. Silently skips duplicate registrations.
// Skips (no-op) when version is nil or contractType is empty so adapterKey does not mis-register.
func (r *AuthorizedCallersRegistry) RegisterAdapter(
	family string,
	contractType cldf.ContractType,
	version *semver.Version,
	a AuthorizedCallersAdapter,
) {
	if version == nil || contractType == "" {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	key := adapterKey(family, contractType, version)
	if _, exists := r.adapters[key]; !exists {
		r.adapters[key] = a
	}
}

// GetAdapter retrieves the adapter for the given triple. Returns false if none is registered.
func (r *AuthorizedCallersRegistry) GetAdapter(
	family string,
	contractType cldf.ContractType,
	version *semver.Version,
) (AuthorizedCallersAdapter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	a, ok := r.adapters[adapterKey(family, contractType, version)]
	return a, ok
}

// adapterKey builds the registry lookup key. It extends utils.NewRegistererID
// (which only encodes family+version) by also including ContractType, because
// multiple AuthorizedCallers-inheriting contracts may exist for the same
// chain family and version (e.g. RMN, FeeQuoter, AdvancedPoolHooks).
// A nil version yields an empty semver segment (avoid panics); callers should
// validate inputs instead of relying on that encoding.
func adapterKey(family string, contractType cldf.ContractType, version *semver.Version) string {
	verStr := ""
	if version != nil {
		verStr = version.String()
	}
	return fmt.Sprintf("%s-%s-%s", family, contractType, verStr)
}
