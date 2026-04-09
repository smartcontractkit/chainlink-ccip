package adapters_test

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
)

type mockChainFamily struct{}

func (m *mockChainFamily) ConfigureChainForLanes() *cldf_ops.Sequence[adapters.ConfigureChainForLanesInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"mock-chain-family",
		semver.MustParse("1.0.0"),
		"mock",
		func(_ cldf_ops.Bundle, _ cldf_chain.BlockChains, _ adapters.ConfigureChainForLanesInput) (sequences.OnChainOutput, error) {
			return sequences.OnChainOutput{}, nil
		},
	)
}

func (m *mockChainFamily) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return []byte(ref.Address), nil
}

func (m *mockChainFamily) GetOnRampAddress(_ datastore.DataStore, _ uint64) ([]byte, error) {
	return nil, nil
}

func (m *mockChainFamily) GetOffRampAddress(_ datastore.DataStore, _ uint64) ([]byte, error) {
	return nil, nil
}

func (m *mockChainFamily) GetFQAddress(_ datastore.DataStore, _ uint64) ([]byte, error) {
	return nil, nil
}

func (m *mockChainFamily) GetRouterAddress(_ datastore.DataStore, _ uint64) ([]byte, error) {
	return nil, nil
}

func (m *mockChainFamily) GetTestRouter(_ datastore.DataStore, _ uint64) ([]byte, error) {
	return nil, nil
}

func (m *mockChainFamily) ResolveExecutor(_ datastore.DataStore, _ uint64, _ string) (string, error) {
	return "", nil
}

func TestChainFamilyRegistry_RegisterAndGet(t *testing.T) {
	registry := adapters.NewChainFamilyRegistry()
	adapter := &mockChainFamily{}

	registry.RegisterChainFamily(chainsel.FamilyEVM, adapter)

	got, ok := registry.GetChainFamily(chainsel.FamilyEVM)
	require.True(t, ok)
	assert.Same(t, adapter, got)
}

func TestChainFamilyRegistry_DuplicateRegisterDoesNotPanic(t *testing.T) {
	registry := adapters.NewChainFamilyRegistry()
	adapter1 := &mockChainFamily{}

	registry.RegisterChainFamily(chainsel.FamilyEVM, adapter1)
	require.NotPanics(t, func() {
		registry.RegisterChainFamily(chainsel.FamilyEVM, adapter1)
	})

	got, ok := registry.GetChainFamily(chainsel.FamilyEVM)
	require.True(t, ok)
	assert.Same(t, adapter1, got)
}
