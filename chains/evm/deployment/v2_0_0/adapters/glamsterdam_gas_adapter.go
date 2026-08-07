package adapters

import (
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// GlamsterdamGasAdapter implements v2_0_0_adapters.GasUpdateAdapter for EVM chains.
type GlamsterdamGasAdapter struct{}

// HasLaneToTarget checks if a lane exists from srcChainSelector to targetChainSelector
// by reading the OnRamp's dest chain config.
func (a *GlamsterdamGasAdapter) HasLaneToTarget(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector, targetChainSelector uint64,
) (bool, error) {
	// Stub: v2.0.0 implementation would read from OnRamp, FeeQuoter, or verifier contracts
	// For now, return true to allow the adapter to be registered
	return true, nil
}

// ReadDestGasFields reads the current gas field values from all contracts.
func (a *GlamsterdamGasAdapter) ReadDestGasFields(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector, targetChainSelector uint64,
) (map[string]uint32, error) {
	// Stub: v2.0.0 would read from OnRamp, FeeQuoter, CommitteeVerifier, LombardVerifier, CCTPVerifier
	return map[string]uint32{}, nil
}

// WriteDestGasFields writes the resolved gas field values to all contracts via MCMS.
func (a *GlamsterdamGasAdapter) WriteDestGasFields(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector, targetChainSelector uint64,
	resolved map[string]uint32,
) ([]mcms_types.BatchOperation, error) {
	// Stub: v2.0.0 would call ApplyOnRampDestChainConfigUpdates, ApplyFeeQuoterDestChainConfigUpdates,
	// ApplyCommitteeVerifierRemoteChainConfigUpdates, ApplyLombardVerifierRemoteChainConfigUpdates,
	// ApplyCCTPVerifierRemoteChainConfigUpdates
	return []mcms_types.BatchOperation{}, nil
}

// ReadImmutableSanityFields reads OffRamp's immutable fields for sanity checking.
func (a *GlamsterdamGasAdapter) ReadImmutableSanityFields(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector uint64,
) (map[string]uint32, error) {
	// Stub implementation
	return map[string]uint32{}, nil
}

// DiscoverCandidateTokens returns all tokens known to the token registry on the chain.
func (a *GlamsterdamGasAdapter) DiscoverCandidateTokens(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector uint64,
) ([][]byte, error) {
	// Stub implementation
	return [][]byte{}, nil
}

// ReadTokenGasField reads a token's gas field from TokenPool.
func (a *GlamsterdamGasAdapter) ReadTokenGasField(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector, targetChainSelector uint64,
	token []byte,
) (uint32, bool, error) {
	// Stub implementation
	return 0, false, nil
}

// WriteTokenGasField writes a token's gas field to TokenPool via MCMS.
func (a *GlamsterdamGasAdapter) WriteTokenGasField(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	ds datastore.DataStore,
	srcChainSelector, targetChainSelector uint64,
	token []byte,
	value uint32,
) (mcms_types.BatchOperation, error) {
	// Stub implementation
	return mcms_types.BatchOperation{}, nil
}
