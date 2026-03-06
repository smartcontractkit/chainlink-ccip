package adapters

import (
	"context"
	"fmt"
	"sync"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

type CommitteeState struct {
	Qualifier        string
	ChainSelector    uint64
	Address          string
	SignatureConfigs []SignatureConfig
}

type SignatureConfig struct {
	SourceChainSelector uint64
	Signers             []string
	Threshold           uint8
}

// AggregatorConfigAdapter provides chain-family-specific logic to discover committee state
// and resolve verifier addresses for aggregator offchain config.
type AggregatorConfigAdapter interface {
	// ScanCommitteeStates returns committee states for the given chain from the deployment env.
	// Each state includes qualifier, chain selector, committee address, and signature configs (signers, threshold per source chain).
	ScanCommitteeStates(ctx context.Context, env deployment.Environment, chainSelector uint64) ([]*CommitteeState, error)
	// ResolveVerifierAddress returns the verifier contract address for the given chain and qualifier using the datastore.
	ResolveVerifierAddress(ds datastore.DataStore, chainSelector uint64, qualifier string) (string, error)
}

type AggregatorConfigRegistry struct {
	mu       sync.Mutex
	adapters map[string]AggregatorConfigAdapter
}

var (
	singletonOffchainConfigRegistry *AggregatorConfigRegistry
	aggregatorConfigRegistryOnce    sync.Once
)

func newOffchainConfigRegistry() *AggregatorConfigRegistry {
	return &AggregatorConfigRegistry{
		adapters: make(map[string]AggregatorConfigAdapter),
	}
}

func GetAggregatorConfigRegistry() *AggregatorConfigRegistry {
	aggregatorConfigRegistryOnce.Do(func() {
		singletonOffchainConfigRegistry = newOffchainConfigRegistry()
	})
	return singletonOffchainConfigRegistry
}

func (r *AggregatorConfigRegistry) Register(family string, a AggregatorConfigAdapter) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.adapters == nil {
		r.adapters = make(map[string]AggregatorConfigAdapter)
	}
	if _, exists := r.adapters[family]; !exists {
		r.adapters[family] = a
	}
}

func (r *AggregatorConfigRegistry) Get(family string) (AggregatorConfigAdapter, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.adapters == nil {
		return nil, false
	}
	a, ok := r.adapters[family]
	return a, ok
}

func (r *AggregatorConfigRegistry) GetByChain(chainSelector uint64) (AggregatorConfigAdapter, error) {
	family, err := chainsel.GetSelectorFamily(chainSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to get chain family for selector %d: %w", chainSelector, err)
	}
	adapter, ok := r.Get(family)
	if !ok {
		return nil, fmt.Errorf("no offchain config adapter registered for chain family %q", family)
	}
	return adapter, nil
}
