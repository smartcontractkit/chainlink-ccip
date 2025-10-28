package lanes

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type LaneAdapter interface {
	// high level API
	ConfigureLaneLegAsSource() *cldf_ops.Sequence[UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	ConfigureLaneLegAsDest() *cldf_ops.Sequence[UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains]

	// helpers to expose lower level functionality if needed
	// needed for populating values in chain specific configs
	GetOnRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetOffRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetRouterAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetFQAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
}

type laneAdapterID string

// LaneAdapterRegistry maintains a registry of LaneAdapters.
type LaneAdapterRegistry struct {
	mu sync.Mutex
	m  map[laneAdapterID]LaneAdapter
}

// NewLaneAdapterRegistry creates a fresh registry.  It is kept unexported
// because callers should obtain the singleton via GetLaneAdapterRegistry().
func newLaneAdapterRegistry() *LaneAdapterRegistry {
	return &LaneAdapterRegistry{
		m: make(map[laneAdapterID]LaneAdapter),
	}
}

// RegisterLaneAdapter registers a new adapter; panics if the key already exists.
func (r *LaneAdapterRegistry) RegisterLaneAdapter(chainFamily string, version *semver.Version, adapter LaneAdapter) {
	id := newLaneAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.m[id]; exists {
		panic(fmt.Errorf("LaneAdapter '%s %s' already registered", chainFamily, version))
	}
	r.m[id] = adapter
}

// GetLaneAdapter looks up an adapter; the second return value tells you if it was found.
func (r *LaneAdapterRegistry) GetLaneAdapter(chainFamily string, version *semver.Version) (LaneAdapter, bool) {
	id := newLaneAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	adapter, ok := r.m[id]
	return adapter, ok
}

var (
	singletonRegistry *LaneAdapterRegistry
	once              sync.Once
)

// GetLaneAdapterRegistry returns the global singleton instance.
// The first call creates the registry; subsequent calls return the same pointer.
func GetLaneAdapterRegistry() *LaneAdapterRegistry {
	once.Do(func() {
		singletonRegistry = newLaneAdapterRegistry()
	})
	return singletonRegistry
}

func newLaneAdapterID(chainFamily string, version *semver.Version) laneAdapterID {
	return laneAdapterID(fmt.Sprintf("%s-%s", chainFamily, version.String()))
}
