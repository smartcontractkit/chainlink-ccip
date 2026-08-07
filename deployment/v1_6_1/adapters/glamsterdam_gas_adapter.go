package adapters

import (
	"sync"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// GasUpdateAdapter provides the interface for implementing Glamsterdam gas config updates
// for a specific chain family.
type GasUpdateAdapter interface {
	// HasLaneToTarget reports whether srcChainSelector has an enabled lane pointed at
	// targetChainSelector, per that family's dest-chain-config-equivalent contract.
	// Returns an error only for hard failures (RPC error, stale address); "no lane" is a
	// normal false, not an error.
	HasLaneToTarget(b cldf_ops.Bundle, chains chain.BlockChains, ds datastore.DataStore, srcChainSelector, targetChainSelector uint64) (bool, error)

	// ReadDestGasFields reads the current on-chain values for every field this version's
	// mapping table defines (FeeQuoter.DestChainConfig.DestGasOverhead, DefaultTokenDestGasOverhead, ...)
	// for one lane, keyed by field name matching the FieldSpec.Name values in fields.go.
	ReadDestGasFields(b cldf_ops.Bundle, chains chain.BlockChains, ds datastore.DataStore, srcChainSelector, targetChainSelector uint64) (map[string]uint32, error)

	// WriteDestGasFields applies resolved values back on-chain, returning one MCMS batch op per
	// contract this lane's writes touch. resolved is keyed the same way as ReadDestGasFields's return.
	WriteDestGasFields(b cldf_ops.Bundle, chains chain.BlockChains, ds datastore.DataStore, srcChainSelector, targetChainSelector uint64, resolved map[string]uint32) ([]mcms_types.BatchOperation, error)

	// ReadImmutableSanityFields reads fields that exist on-chain but have no setter (e.g. OffRamp
	// GasForCallExactCheck) purely to flag an unexpected deployment. Returns field name -> value;
	// caller compares against the expected baseline and only logs, never writes.
	ReadImmutableSanityFields(b cldf_ops.Bundle, chains chain.BlockChains, ds datastore.DataStore, srcChainSelector uint64) (map[string]uint32, error)

	// DiscoverCandidateTokens returns every token this chain's token registry knows about, as raw
	// on-chain address bytes in that chain's native encoding. Used to build the per-token
	// TokenTransferFeeConfig candidate list.
	DiscoverCandidateTokens(b cldf_ops.Bundle, chains chain.BlockChains, ds datastore.DataStore, srcChainSelector uint64) ([][]byte, error)

	// ReadTokenGasField / WriteTokenGasField mirror ReadDestGasFields/WriteDestGasFields but for
	// the one per-(chain,token) field each version's table defines. token is raw address bytes.
	ReadTokenGasField(b cldf_ops.Bundle, chains chain.BlockChains, ds datastore.DataStore, srcChainSelector, targetChainSelector uint64, token []byte) (uint32, bool, error) // bool = field is configured at all
	WriteTokenGasField(b cldf_ops.Bundle, chains chain.BlockChains, ds datastore.DataStore, srcChainSelector, targetChainSelector uint64, token []byte, value uint32) (mcms_types.BatchOperation, error)
}

// GasUpdateAdapterRegistry maintains a registry of GasUpdateAdapter implementations, one per chain family.
type GasUpdateAdapterRegistry struct {
	mu       sync.Mutex
	adapters map[string]GasUpdateAdapter
}

var (
	registryOnce sync.Once
	registry     *GasUpdateAdapterRegistry
)

// GetGasUpdateAdapterRegistry returns the singleton GasUpdateAdapterRegistry.
func GetGasUpdateAdapterRegistry() *GasUpdateAdapterRegistry {
	registryOnce.Do(func() {
		registry = &GasUpdateAdapterRegistry{
			adapters: make(map[string]GasUpdateAdapter),
		}
	})
	return registry
}

// RegisterGasUpdateAdapter registers a GasUpdateAdapter for the given chain family.
// Panics if the family is already registered.
func (r *GasUpdateAdapterRegistry) RegisterGasUpdateAdapter(family string, adapter GasUpdateAdapter) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.adapters[family]; exists {
		panic("GasUpdateAdapter already registered for family: " + family)
	}
	r.adapters[family] = adapter
}

// GetGasUpdateAdapter returns the registered GasUpdateAdapter for the given chain family,
// or nil if no adapter is registered.
func (r *GasUpdateAdapterRegistry) GetGasUpdateAdapter(family string) GasUpdateAdapter {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.adapters[family]
}
