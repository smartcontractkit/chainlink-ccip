package lanes

import (
	"fmt"
	"sync"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// PingPongAdapter is an optional interface that chain adapters can implement
// to support PingPong demo contract configuration.
// Chains that don't support PingPong (like Solana) should not implement this interface.
type PingPongAdapter interface {
	// GetPingPongDemoAddress returns the PingPongDemo contract address for the given chain selector.
	// Returns an error if the contract is not deployed.
	GetPingPongDemoAddress(ds datastore.DataStore, chainSelector uint64) ([]byte, error)
	// ConfigurePingPong returns a sequence that configures PingPong for a lane between source and dest.
	ConfigurePingPong() *cldf_ops.Sequence[PingPongInput, PingPongOutput, cldf_chain.BlockChains]
}

// PingPongInput contains the input for configuring PingPong between two chains.
type PingPongInput struct {
	SourceSelector     uint64
	DestSelector       uint64
	SourcePingPongAddr []byte
	DestPingPongAddr   []byte
}

// PingPongOutput contains the output from configuring PingPong.
type PingPongOutput struct{}

type pingPongAdapterID string

// PingPongAdapterRegistry maintains a registry of PingPongAdapters.
type PingPongAdapterRegistry struct {
	mu sync.Mutex
	m  map[pingPongAdapterID]PingPongAdapter
}

// newPingPongAdapterRegistry creates a fresh registry.
func newPingPongAdapterRegistry() *PingPongAdapterRegistry {
	return &PingPongAdapterRegistry{
		m: make(map[pingPongAdapterID]PingPongAdapter),
	}
}

// RegisterPingPongAdapter registers a new adapter; panics if the key already exists.
func (r *PingPongAdapterRegistry) RegisterPingPongAdapter(chainFamily string, version *semver.Version, adapter PingPongAdapter) {
	id := newPingPongAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.m[id]; exists {
		panic(fmt.Errorf("PingPongAdapter '%s %s' already registered", chainFamily, version))
	}
	r.m[id] = adapter
}

// GetPingPongAdapter looks up an adapter; the second return value tells you if it was found.
func (r *PingPongAdapterRegistry) GetPingPongAdapter(chainFamily string, version *semver.Version) (PingPongAdapter, bool) {
	id := newPingPongAdapterID(chainFamily, version)

	r.mu.Lock()
	defer r.mu.Unlock()

	adapter, ok := r.m[id]
	return adapter, ok
}

var (
	singletonPingPongRegistry *PingPongAdapterRegistry
	oncePingPong              sync.Once
)

// GetPingPongAdapterRegistry returns the global singleton instance.
func GetPingPongAdapterRegistry() *PingPongAdapterRegistry {
	oncePingPong.Do(func() {
		singletonPingPongRegistry = newPingPongAdapterRegistry()
	})
	return singletonPingPongRegistry
}

func newPingPongAdapterID(chainFamily string, version *semver.Version) pingPongAdapterID {
	return pingPongAdapterID(fmt.Sprintf("%s-%s", chainFamily, version.String()))
}

// ConfigurePingPongForLanes configures PingPong contracts for all lanes that support it.
// Chains without PingPong adapters are silently skipped.
func ConfigurePingPongForLanes(e cldf.Environment, registry *PingPongAdapterRegistry, version *semver.Version, selector uint64, remoteSelectors []uint64) error {
	chainFamily, err := chain_selectors.GetSelectorFamily(selector)
	if err != nil {
		return fmt.Errorf("failed to get chain family for selector %d: %w", selector, err)
	}

	// Get the adapter for this chain - if not found, skip (not all chains support PingPong)
	adapter, exists := registry.GetPingPongAdapter(chainFamily, version)
	if !exists {
		// Chain doesn't support PingPong - silently skip
		return nil
	}

	// Get the PingPongDemo address on this chain - if not deployed, skip
	localPingPongAddr, err := adapter.GetPingPongDemoAddress(e.DataStore, selector)
	if err != nil {
		// PingPong not deployed on this chain - silently skip
		return nil
	}

	for _, remoteSelector := range remoteSelectors {
		remoteFamily, err := chain_selectors.GetSelectorFamily(remoteSelector)
		if err != nil {
			return fmt.Errorf("failed to get chain family for remote selector %d: %w", remoteSelector, err)
		}

		// Get remote adapter - if not found, skip this pair
		remoteAdapter, exists := registry.GetPingPongAdapter(remoteFamily, version)
		if !exists {
			continue
		}

		// Get remote PingPong address - if not deployed, skip this pair
		remotePingPongAddr, err := remoteAdapter.GetPingPongDemoAddress(e.DataStore, remoteSelector)
		if err != nil {
			continue
		}

		// Configure PingPong for this lane
		_, err = cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.ConfigurePingPong(), e.BlockChains, PingPongInput{
			SourceSelector:     selector,
			DestSelector:       remoteSelector,
			SourcePingPongAddr: localPingPongAddr,
			DestPingPongAddr:   remotePingPongAddr,
		})
		if err != nil {
			return fmt.Errorf("failed to configure PingPong for lane %d -> %d: %w", selector, remoteSelector, err)
		}
	}

	return nil
}
