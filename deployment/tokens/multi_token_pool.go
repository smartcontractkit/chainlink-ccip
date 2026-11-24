package tokens

import (
	"sync"

	"github.com/Masterminds/semver/v3"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// DeployMultiTokenPoolInput is the input for deploying one multi-token pool stack on one chain.
// The stack consists of the connection, fee, and security policy stores, a PoolConfigManager,
// and the pool itself. Token registration and lane configuration are separate workflows.
type DeployMultiTokenPoolInput struct {
	ChainSelector uint64          `json:"chainSelector" yaml:"chainSelector"`
	PoolType      string          `json:"poolType" yaml:"poolType"`
	PoolVersion   *semver.Version `json:"poolVersion" yaml:"poolVersion"`
	// PoolQualifier identifies the stack in the datastore. It is required because an MTTP is
	// not tied to a single token whose address could otherwise be used as the qualifier.
	PoolQualifier string `json:"poolQualifier" yaml:"poolQualifier"`
	// Admin is the initial PoolConfigManager default admin, in the chain family's native
	// string format. When empty, the changeset uses the selected MCMS timelock when available;
	// otherwise the chain adapter may use its deployer.
	Admin string `json:"admin,omitempty" yaml:"admin,omitempty"`
	// ExistingDataStore is populated by the changeset and is not user input.
	ExistingDataStore datastore.DataStore `json:"-" yaml:"-"`
}

// MultiTokenPoolAdapter implements multi-token pool operations for a chain family and pool version.
// It is separate from TokenAdapter because an MTTP is deployed before, and independently of, any
// particular token it will serve.
type MultiTokenPoolAdapter interface {
	// ValidateDeployMultiTokenPool rejects input the chain family's sequence cannot deploy, such
	// as an unsupported pool type or version, a malformed admin address, an inconsistent existing
	// stack, or a TokenAdminRegistry that differs from the ramps'. It may read chain state but
	// must not send transactions. The changeset calls it for every chain before deploying to any.
	ValidateDeployMultiTokenPool(e cldf.Environment, input DeployMultiTokenPoolInput) error
	// DeployMultiTokenPool returns a sequence that deploys or reuses the stack identified by
	// PoolType, PoolVersion, and PoolQualifier. Implementations must resolve well-known addresses
	// from ExistingDataStore, use native address strings in inputs, and emit only missing writes.
	DeployMultiTokenPool() *cldf_ops.Sequence[DeployMultiTokenPoolInput, sequences.OnChainOutput, cldf_chain.BlockChains]
}

// MultiTokenPoolAdapterRegistry maintains MultiTokenPoolAdapters by chain family and pool version.
type MultiTokenPoolAdapterRegistry struct {
	mu       sync.Mutex
	adapters map[tokenAdapterID]MultiTokenPoolAdapter
	versions map[string][]*semver.Version
}

// NewMultiTokenPoolAdapterRegistry creates an empty registry. Tests and consumers that need
// isolated registrations can use this instead of the process-wide singleton.
func NewMultiTokenPoolAdapterRegistry() *MultiTokenPoolAdapterRegistry {
	return &MultiTokenPoolAdapterRegistry{
		adapters: make(map[tokenAdapterID]MultiTokenPoolAdapter),
		versions: make(map[string][]*semver.Version),
	}
}

// RegisterMultiTokenPoolAdapter registers an adapter. The first registration for a family and
// version wins, matching TokenAdapterRegistry's registration semantics.
func (r *MultiTokenPoolAdapterRegistry) RegisterMultiTokenPoolAdapter(chainFamily string, version *semver.Version, adapter MultiTokenPoolAdapter) {
	id := newTokenAdapterID(chainFamily, version)
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.adapters[id]; exists {
		return
	}
	r.adapters[id] = adapter
	r.versions[chainFamily] = append(r.versions[chainFamily], version)
}

// GetMultiTokenPoolAdapter retrieves the adapter registered for the chain family and pool version.
// Versions on the same supported release line resolve using the shared version-resolution rules.
func (r *MultiTokenPoolAdapterRegistry) GetMultiTokenPoolAdapter(chainFamily string, version *semver.Version) (MultiTokenPoolAdapter, bool) {
	if version == nil {
		return nil, false
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	registered, ok := utils.ResolveRegisteredVersion(r.versions[chainFamily], version)
	if !ok {
		return nil, false
	}
	adapter, ok := r.adapters[newTokenAdapterID(chainFamily, registered)]
	return adapter, ok
}

var (
	multiTokenPoolRegistry     *MultiTokenPoolAdapterRegistry
	multiTokenPoolRegistryOnce sync.Once
)

// GetMultiTokenPoolAdapterRegistry returns the process-wide multi-token pool adapter registry.
func GetMultiTokenPoolAdapterRegistry() *MultiTokenPoolAdapterRegistry {
	multiTokenPoolRegistryOnce.Do(func() {
		multiTokenPoolRegistry = NewMultiTokenPoolAdapterRegistry()
	})
	return multiTokenPoolRegistry
}
