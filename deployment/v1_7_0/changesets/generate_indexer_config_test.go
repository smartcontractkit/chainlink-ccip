package changesets_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
)

var _ adapters.IndexerConfigAdapter = (*mockIndexerConfigAdapter)(nil)

type mockIndexerConfigAdapter struct {
	addressesByChainAndKind map[uint64]map[adapters.VerifierKind]map[string][]string
	resolveErr              error
}

func (m *mockIndexerConfigAdapter) ResolveVerifierAddresses(
	_ datastore.DataStore, chainSelector uint64, qualifier string, kind adapters.VerifierKind,
) ([]string, error) {
	if m.resolveErr != nil {
		return nil, m.resolveErr
	}
	byKind, ok := m.addressesByChainAndKind[chainSelector]
	if !ok {
		return nil, nil
	}
	byQualifier, ok := byKind[kind]
	if !ok {
		return nil, nil
	}
	return byQualifier[qualifier], nil
}

func newIndexerConfigRegistry() *adapters.IndexerConfigRegistry {
	return &adapters.IndexerConfigRegistry{}
}

func TestGenerateIndexerConfig_Validation(t *testing.T) {
	registry := newIndexerConfigRegistry()
	cs := changesets.GenerateIndexerConfig(registry)

	tests := []struct {
		name        string
		input       changesets.GenerateIndexerConfigInput
		errContains string
	}{
		{
			name: "missing service identifier returns error",
			input: changesets.GenerateIndexerConfigInput{
				CommitteeVerifierNameToQualifier: map[string]string{"v1": "default"},
			},
			errContains: "service identifier is required",
		},
		{
			name: "no verifier mappings returns error",
			input: changesets.GenerateIndexerConfigInput{
				ServiceIdentifier: "indexer",
			},
			errContains: "at least one verifier name to qualifier mapping is required",
		},
	}

	env := deployment.Environment{}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := cs.VerifyPreconditions(env, tc.input)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tc.errContains)
		})
	}
}

func TestGenerateIndexerConfig_BuildsCorrectConfigAcrossChains(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	addr1 := "0x1111111111111111111111111111111111111111"
	addr2 := "0x2222222222222222222222222222222222222222"
	addr3 := "0x3333333333333333333333333333333333333333"

	mock := &mockIndexerConfigAdapter{
		addressesByChainAndKind: map[uint64]map[adapters.VerifierKind]map[string][]string{
			sel1: {
				adapters.CommitteeVerifierKind: {"default": {addr1}},
				adapters.CCTPVerifierKind:      {"cctp-q": {addr2}},
			},
			sel2: {
				adapters.CommitteeVerifierKind: {"default": {addr3}},
			},
		},
	}

	registry := newIndexerConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	env := deployment.Environment{
		DataStore:   ds.Seal(),
		BlockChains: newTestBlockChains([]uint64{sel1, sel2}),
	}

	cs := changesets.GenerateIndexerConfig(registry)
	output, err := cs.Apply(env, changesets.GenerateIndexerConfigInput{
		ServiceIdentifier:                "test-indexer",
		CommitteeVerifierNameToQualifier: map[string]string{"committee": "default"},
		CCTPVerifierNameToQualifier:      map[string]string{"cctp": "cctp-q"},
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	cfg, err := offchain.GetIndexerConfig(output.DataStore.Seal(), "test-indexer")
	require.NoError(t, err)

	verifiersByName := make(map[string][]string)
	for _, v := range cfg.Verifiers {
		verifiersByName[v.Name] = v.IssuerAddresses
	}

	require.Contains(t, verifiersByName, "committee")
	assert.ElementsMatch(t, []string{addr1, addr3}, verifiersByName["committee"])

	require.Contains(t, verifiersByName, "cctp")
	assert.ElementsMatch(t, []string{addr2}, verifiersByName["cctp"])
}

func TestGenerateIndexerConfig_DeduplicatesAddressesAcrossChains(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	sharedAddr := "0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

	mock := &mockIndexerConfigAdapter{
		addressesByChainAndKind: map[uint64]map[adapters.VerifierKind]map[string][]string{
			sel1: {
				adapters.CommitteeVerifierKind: {"default": {sharedAddr}},
			},
			sel2: {
				adapters.CommitteeVerifierKind: {"default": {sharedAddr}},
			},
		},
	}

	registry := newIndexerConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	env := deployment.Environment{
		DataStore:   ds.Seal(),
		BlockChains: newTestBlockChains([]uint64{sel1, sel2}),
	}

	cs := changesets.GenerateIndexerConfig(registry)
	output, err := cs.Apply(env, changesets.GenerateIndexerConfigInput{
		ServiceIdentifier:                "test-indexer",
		CommitteeVerifierNameToQualifier: map[string]string{"committee": "default"},
	})
	require.NoError(t, err)

	cfg, err := offchain.GetIndexerConfig(output.DataStore.Seal(), "test-indexer")
	require.NoError(t, err)

	require.Len(t, cfg.Verifiers, 1)
	assert.Equal(t, "committee", cfg.Verifiers[0].Name)
	assert.Equal(t, []string{sharedAddr}, cfg.Verifiers[0].IssuerAddresses)
}

func TestGenerateIndexerConfig_AdapterErrorPropagates(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockIndexerConfigAdapter{
		addressesByChainAndKind: map[uint64]map[adapters.VerifierKind]map[string][]string{
			sel1: {
				adapters.CommitteeVerifierKind: {"default": {"0x1111111111111111111111111111111111111111"}},
			},
		},
		resolveErr: fmt.Errorf("datastore unavailable"),
	}

	registry := newIndexerConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	env := deployment.Environment{
		DataStore:   ds.Seal(),
		BlockChains: newTestBlockChains([]uint64{sel1}),
	}

	cs := changesets.GenerateIndexerConfig(registry)
	_, err := cs.Apply(env, changesets.GenerateIndexerConfigInput{
		ServiceIdentifier:                "test-indexer",
		CommitteeVerifierNameToQualifier: map[string]string{"committee": "default"},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "datastore unavailable")
}

func TestGenerateIndexerConfig_MissingAdapterReturnsError(t *testing.T) {
	registry := newIndexerConfigRegistry()

	ds := datastore.NewMemoryDataStore()
	env := deployment.Environment{
		DataStore:   ds.Seal(),
		BlockChains: newTestBlockChains([]uint64{chainsel.TEST_90000001.Selector}),
	}

	cs := changesets.GenerateIndexerConfig(registry)
	_, err := cs.Apply(env, changesets.GenerateIndexerConfigInput{
		ServiceIdentifier:                "test-indexer",
		CommitteeVerifierNameToQualifier: map[string]string{"committee": "default"},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "no indexer config adapter registered")
}

func TestGenerateIndexerConfig_IgnoresUndeployedEnvChains(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	addr1 := "0x1111111111111111111111111111111111111111"

	mock := &mockIndexerConfigAdapter{
		addressesByChainAndKind: map[uint64]map[adapters.VerifierKind]map[string][]string{
			sel1: {
				adapters.CommitteeVerifierKind: {"default": {addr1}},
			},
		},
	}

	registry := newIndexerConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	env := deployment.Environment{
		DataStore:   ds.Seal(),
		BlockChains: newTestBlockChains([]uint64{sel1, sel2}),
	}

	cs := changesets.GenerateIndexerConfig(registry)
	output, err := cs.Apply(env, changesets.GenerateIndexerConfigInput{
		ServiceIdentifier:                "test-indexer",
		CommitteeVerifierNameToQualifier: map[string]string{"committee": "default"},
	})
	require.NoError(t, err)

	cfg, err := offchain.GetIndexerConfig(output.DataStore.Seal(), "test-indexer")
	require.NoError(t, err)

	require.Len(t, cfg.Verifiers, 1)
	assert.Equal(t, []string{addr1}, cfg.Verifiers[0].IssuerAddresses)
}

func TestGenerateIndexerConfig_NoDeployedChainsForQualifierReturnsError(t *testing.T) {
	registry := newIndexerConfigRegistry()
	registry.Register(chainsel.FamilyEVM, &mockIndexerConfigAdapter{
		addressesByChainAndKind: map[uint64]map[adapters.VerifierKind]map[string][]string{},
	})

	ds := datastore.NewMemoryDataStore()
	env := deployment.Environment{
		DataStore:   ds.Seal(),
		BlockChains: newTestBlockChains([]uint64{chainsel.TEST_90000001.Selector}),
	}

	cs := changesets.GenerateIndexerConfig(registry)
	_, err := cs.Apply(env, changesets.GenerateIndexerConfigInput{
		ServiceIdentifier:                "test-indexer",
		CommitteeVerifierNameToQualifier: map[string]string{"committee": "default"},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), `failed to resolve addresses for verifier "committee" (qualifier "default"): no deployed committee verifier addresses found for qualifier "default"`)
}
