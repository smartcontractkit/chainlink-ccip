package lanes

import (
	"fmt"
	"math/big"
	"sync"

	"github.com/Masterminds/semver/v3"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

type TestRouterProvider interface {
	GetTestRouter(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
}

type LaneAdapter interface {
	// high level API
	ConfigureLaneLegAsSource() *cldf_ops.Sequence[UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	ConfigureLaneLegAsDest() *cldf_ops.Sequence[UpdateLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	DisableRemoteChain() *cldf_ops.Sequence[DisableRemoteChainInput, sequences.OnChainOutput, cldf_chain.BlockChains]
	// helpers to expose lower level functionality if needed
	// needed for populating values in chain specific configs
	GetOnRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetOffRampAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetFQAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetRouterAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	GetFeeQuoterDestChainConfig() FeeQuoterDestChainConfig
	// GasPrice defines the USD price (18 decimals) per unit gas for this chain as a destination.
	GetDefaultGasPrice() *big.Int
}

// FeeQuoterVersionProvider is an optional interface that LaneAdapters can implement
// to report the FeeQuoter contract version for a chain (e.g. 1.6.x vs 2.0.x).
// When set, it is used to choose the correct FeeQuoter operations in update_lanes.
type FeeQuoterVersionProvider interface {
	GetFQVersion(ds datastore.DataStore, address []byte, chainSelector uint64, chains cldf_chain.BlockChains) (*semver.Version, error)
}

// DynamicFeeQuoter is an optional interface that LaneAdapters can implement to dynamically retrieve the FeeQuoter.
type DynamicFeeQuoter interface {
	GetFQAddressDynamic(ds datastore.DataStore, chainSelector uint64, chains cldf_chain.BlockChains) ([]byte, error)
}

// TokenPriceProvider is an optional interface that LaneAdapters can implement
// to provide default fee token prices for a chain.
// This is primarily used by EVM chains; other chains can skip this.
type TokenPriceProvider interface {
	// GetDefaultTokenPrices returns default fee token prices for a chain.
	// Returns a map of contract type to USD price (18 decimals).
	// The caller is responsible for resolving contract types to addresses.
	GetDefaultTokenPrices() map[datastore.ContractType]*big.Int
}

// ChainMetadataProvider is an optional interface that LaneAdapters can implement
// to expose the on-chain 4-byte chain family selector (e.g. 0x2812d52c for EVM).
// When implemented, the selector is automatically registered in the global
// family-selector registry so that GetSelectorHex can resolve it without a
// hardcoded switch.
type ChainMetadataProvider interface {
	GetChainFamilySelector() [4]byte
}

type laneAdapterID string

// LaneAdapterRegistry maintains a registry of LaneAdapters.
type LaneAdapterRegistry struct {
	mu sync.Mutex
	m  map[laneAdapterID]LaneAdapter
	// versions are the registered versions per chain family.
	versions map[string][]*semver.Version
}

// NewLaneAdapterRegistry creates a fresh registry.  It is kept unexported
// because callers should obtain the singleton via GetLaneAdapterRegistry().
func newLaneAdapterRegistry() *LaneAdapterRegistry {
	return &LaneAdapterRegistry{
		m:        make(map[laneAdapterID]LaneAdapter),
		versions: make(map[string][]*semver.Version),
	}
}

// RegisterLaneAdapter registers a new adapter.
// If the adapter implements ChainMetadataProvider, its chain family selector is
// also registered in the global family-selector registry for GetSelectorHex.
func (r *LaneAdapterRegistry) RegisterLaneAdapter(chainFamily string, version *semver.Version, adapter LaneAdapter) {
	id := newLaneAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.m[id]; !exists {
		r.m[id] = adapter
		r.versions[chainFamily] = append(r.versions[chainFamily], version)
	}

	if mp, ok := adapter.(ChainMetadataProvider); ok {
		utils.RegisterChainFamilySelector(chainFamily, mp.GetChainFamilySelector())
	}
}

// GetLaneAdapter looks up an adapter; the second return value tells you if it was found.
func (r *LaneAdapterRegistry) GetLaneAdapter(chainFamily string, version *semver.Version) (LaneAdapter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Exact registration first, then the same major from 2.0.0 up (a 2.0.1 ref uses the 2.0.0 adapter).
	registered, ok := utils.ResolveRegisteredVersion(r.versions[chainFamily], version)
	if !ok {
		return nil, false
	}
	adapter, ok := r.m[newLaneAdapterID(chainFamily, registered)]
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
