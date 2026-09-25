package adapters

import (
	"testing"

	"github.com/stretchr/testify/require"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	v1_6_1_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_6_1/adapters"
)

// TestAdapterRegistration verifies the EVM adapter is correctly registered in the registry.
func TestAdapterRegistration(t *testing.T) {
	registry := v1_6_1_adapters.GetGasUpdateAdapterRegistry()

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
	var _ v1_6_1_adapters.GasUpdateAdapter = adapter
}
