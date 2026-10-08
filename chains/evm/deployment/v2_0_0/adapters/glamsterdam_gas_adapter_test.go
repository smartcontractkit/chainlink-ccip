package adapters

import (
	"testing"

	"github.com/stretchr/testify/require"

	v2_0_0_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
)

// TestAdapterRegistration verifies the EVM adapter is correctly registered in the registry.
func TestAdapterRegistration(t *testing.T) {
	registry := v2_0_0_adapters.GetGasUpdateAdapterRegistry()

	// Adapter should be registered for EVM family
	adapter := registry.GetGasUpdateAdapter(chain_selectors.FamilyEVM)
	require.NotNil(t, adapter)

	// Verify it's the correct type
	_, ok := adapter.(*GlamsterdamGasAdapter)
	require.True(t, ok, "registered adapter should be GlamsterdamGasAdapter")
}

// TestAdapterInterface verifies the adapter implements the interface
func TestAdapterInterface(t *testing.T) {
	adapter := &GlamsterdamGasAdapter{}

	// Verify it satisfies the interface
	var _ v2_0_0_adapters.GasUpdateAdapter = adapter
}
