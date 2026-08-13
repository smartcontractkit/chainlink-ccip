package adapters

import (
	"fmt"
	"sync"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// DeployTestVerifierChainInput is the input for deploying test verifier infrastructure on one chain.
// It deploys a TESTVTR drip token, a matching BurnMintTokenPool wired to the TestRouter,
// a VerifierTestHelper, and a VersionedVerifierResolver.
type DeployTestVerifierChainInput struct {
	// ChainSelector is the selector for the chain being deployed to.
	ChainSelector uint64
	// TokenName is the name of the drip token (e.g. "TESTVTR").
	TokenName string
	// TokenSymbol is the symbol of the drip token (e.g. "TESTVTR").
	TokenSymbol string
	// TokenDecimals is the number of decimals for the token.
	TokenDecimals uint8
	// PreMintAccounts maps addresses to initial mint amounts (as decimal strings).
	PreMintAccounts map[string]string
	// AllowedSenders are addresses allowlisted on the test verifier.
	AllowedSenders []string
	// StorageLocations are passed to the VerifierTestHelper constructor. The verifier's
	// getFee delegates to RMN, which needs valid storage locations to compute fees.
	StorageLocations []string
}

// DeployTestVerifierChainDeps are the dependencies for the DeployTestVerifierChain sequence.
type DeployTestVerifierChainDeps struct {
	// BlockChains are the chains in the environment.
	BlockChains cldf_chain.BlockChains
	// DataStore defines all addresses in the environment.
	DataStore datastore.DataStore
}

// RemoteTestVerifierChainConfig configures a test-verifier-enabled chain for a remote counterpart.
type RemoteTestVerifierChainConfig struct {
	// FeeUSDCents is the flat fee for verification on the remote chain.
	FeeUSDCents uint16
	// GasForVerification is the gas required to verify on the remote chain.
	GasForVerification uint32
	// PayloadSizeBytes is the payload size for verification on the remote chain.
	PayloadSizeBytes uint16
}

// ConfigureTestVerifierForLanesInput is the input for the ConfigureTestVerifierChainForLanes sequence.
type ConfigureTestVerifierForLanesInput struct {
	// ChainSelector is the selector for the chain being configured.
	ChainSelector uint64
	// AllowedSenders are addresses allowlisted on the test verifier for each remote.
	AllowedSenders []string
	// RemoteChains is the set of remote chains to configure lanes for.
	RemoteChains map[uint64]RemoteTestVerifierChainConfig
}

// ConfigureTestVerifierForLanesDeps are the dependencies for the ConfigureTestVerifierChainForLanes sequence.
type ConfigureTestVerifierForLanesDeps struct {
	// BlockChains are the chains in the environment.
	BlockChains cldf_chain.BlockChains
	// DataStore defines all addresses in the environment.
	DataStore datastore.DataStore
	// RemoteChains are the remote chain adapters for resolving remote addresses.
	RemoteChains map[uint64]RemoteTestVerifierChain
}

// RemoteTestVerifierChain resolves test verifier infrastructure addresses on a remote chain.
type RemoteTestVerifierChain interface {
	// TokenAddress returns the TESTVTR token address on the remote chain in bytes.
	TokenAddress(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error)
	// PoolAddress returns the TESTVTR pool address on the remote chain in bytes.
	PoolAddress(d datastore.DataStore, b cldf_chain.BlockChains, chainSelector uint64) ([]byte, error)
}

// TestVerifierChainAdapter is the interface for chains that support test verifier deployment.
type TestVerifierChainAdapter interface {
	RemoteTestVerifierChain
	// DeployTestVerifierChain deploys the TESTVTR token, pool, verifier, and resolver on one chain.
	DeployTestVerifierChain() *cldf_ops.Sequence[DeployTestVerifierChainInput, sequences.OnChainOutput, DeployTestVerifierChainDeps]
	// ConfigureTestVerifierChainForLanes configures the test verifier and pool for lanes.
	ConfigureTestVerifierChainForLanes() *cldf_ops.Sequence[ConfigureTestVerifierForLanesInput, sequences.OnChainOutput, ConfigureTestVerifierForLanesDeps]
}

// TestVerifierChainRegistry maintains a registry of TestVerifierChainAdapters.
type TestVerifierChainRegistry struct {
	mu sync.Mutex
	m  map[string]TestVerifierChainAdapter
}

// NewTestVerifierChainRegistry creates a new TestVerifierChainRegistry.
func NewTestVerifierChainRegistry() *TestVerifierChainRegistry {
	return &TestVerifierChainRegistry{
		m: make(map[string]TestVerifierChainAdapter),
	}
}

var (
	testVerifierChainRegistrySingleton     *TestVerifierChainRegistry
	testVerifierChainRegistrySingletonOnce sync.Once
)

// GetTestVerifierChainRegistry returns the global singleton instance.
func GetTestVerifierChainRegistry() *TestVerifierChainRegistry {
	testVerifierChainRegistrySingletonOnce.Do(func() {
		testVerifierChainRegistrySingleton = NewTestVerifierChainRegistry()
	})
	return testVerifierChainRegistrySingleton
}

// Register registers a TestVerifierChainAdapter for a chain family.
func (r *TestVerifierChainRegistry) Register(chainFamily string, adapter TestVerifierChainAdapter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.m[chainFamily]; !exists {
		r.m[chainFamily] = adapter
	}
}

// Get retrieves a registered TestVerifierChainAdapter for the given chain family.
func (r *TestVerifierChainRegistry) Get(chainFamily string) (TestVerifierChainAdapter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	adapter, ok := r.m[chainFamily]
	return adapter, ok
}

// String returns a string representation of the registry contents (for debugging).
func (r *TestVerifierChainRegistry) String() string {
	r.mu.Lock()
	defer r.mu.Unlock()
	families := make([]string, 0, len(r.m))
	for family := range r.m {
		families = append(families, family)
	}
	return fmt.Sprintf("TestVerifierChainRegistry(families: %v)", families)
}
