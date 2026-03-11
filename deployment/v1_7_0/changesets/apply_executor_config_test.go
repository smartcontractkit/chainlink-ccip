package changesets_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldfevm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/shared"
)

var _ adapters.ExecutorConfigAdapter = (*mockExecutorConfigAdapter)(nil)

type mockExecutorConfigAdapter struct {
	deployedChains map[string][]uint64
	chainConfigs   map[uint64]adapters.ExecutorChainConfig
	buildErr       error
}

func (m *mockExecutorConfigAdapter) GetDeployedChains(_ datastore.DataStore, qualifier string) []uint64 {
	return m.deployedChains[qualifier]
}

func (m *mockExecutorConfigAdapter) BuildChainConfig(_ datastore.DataStore, chainSelector uint64, _ string) (adapters.ExecutorChainConfig, error) {
	if m.buildErr != nil {
		return adapters.ExecutorChainConfig{}, m.buildErr
	}
	cfg, ok := m.chainConfigs[chainSelector]
	if !ok {
		return adapters.ExecutorChainConfig{}, fmt.Errorf("no config for chain %d", chainSelector)
	}
	return cfg, nil
}

func newMinimalTopology(nopAliases []string, executorQualifier string, mode shared.NOPMode) *offchain.EnvironmentTopology {
	nops := make([]offchain.NOPConfig, len(nopAliases))
	for i, alias := range nopAliases {
		nops[i] = offchain.NOPConfig{Alias: alias, Name: alias + "-name", Mode: mode}
	}
	return &offchain.EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &offchain.NOPTopology{
			NOPs:       nops,
			Committees: map[string]offchain.CommitteeConfig{},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{
			executorQualifier: {
				NOPAliases: nopAliases,
			},
		},
	}
}

func newTestExecutorEnv(t *testing.T, selectors []uint64) deployment.Environment {
	t.Helper()
	lggr := logger.Test(t)
	return deployment.Environment{
		Name:        "test",
		BlockChains: newExecutorTestBlockChains(selectors),
		DataStore:   datastore.NewMemoryDataStore().Seal(),
		Logger:      lggr,
		OperationsBundle: operations.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			operations.NewMemoryReporter(),
		),
	}
}

func newExecutorTestBlockChains(selectors []uint64) cldf_chain.BlockChains {
	chains := make(map[uint64]cldf_chain.BlockChain, len(selectors))
	for _, sel := range selectors {
		chains[sel] = cldfevm.Chain{Selector: sel}
	}
	return cldf_chain.NewBlockChains(chains)
}

func TestApplyExecutorConfig_Validation(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, &mockExecutorConfigAdapter{})

	tests := []struct {
		name    string
		env     deployment.Environment
		input   changesets.ApplyExecutorConfigInput
		wantErr string
	}{
		{
			name:    "missing topology returns error",
			env:     deployment.Environment{BlockChains: newExecutorTestBlockChains([]uint64{sel1})},
			input:   changesets.ApplyExecutorConfigInput{},
			wantErr: "topology is required",
		},
		{
			name: "nil NOP topology returns error",
			env:  deployment.Environment{BlockChains: newExecutorTestBlockChains([]uint64{sel1})},
			input: changesets.ApplyExecutorConfigInput{
				Topology: &offchain.EnvironmentTopology{
					IndexerAddress: []string{"http://indexer:8080"},
					ExecutorPools:  map[string]offchain.ExecutorPoolConfig{"pool1": {NOPAliases: []string{"nop1"}}},
				},
				ExecutorQualifier: "pool1",
			},
			wantErr: "NOP topology with at least one NOP is required",
		},
		{
			name: "missing executor qualifier returns error",
			env:  deployment.Environment{BlockChains: newExecutorTestBlockChains([]uint64{sel1})},
			input: changesets.ApplyExecutorConfigInput{
				Topology: newMinimalTopology([]string{"nop1"}, "pool1", ""),
			},
			wantErr: "executor qualifier is required",
		},
		{
			name: "unknown pool returns error",
			env:  deployment.Environment{BlockChains: newExecutorTestBlockChains([]uint64{sel1})},
			input: changesets.ApplyExecutorConfigInput{
				Topology:          newMinimalTopology([]string{"nop1"}, "pool1", ""),
				ExecutorQualifier: "nonexistent",
			},
			wantErr: "executor pool \"nonexistent\" not found in topology",
		},
		{
			name: "target NOP not in pool returns error",
			env:  deployment.Environment{BlockChains: newExecutorTestBlockChains([]uint64{sel1})},
			input: changesets.ApplyExecutorConfigInput{
				Topology:          newMinimalTopology([]string{"nop1"}, "pool1", ""),
				ExecutorQualifier: "pool1",
				TargetNOPs:        []shared.NOPAlias{"unknown-nop"},
			},
			wantErr: "NOP alias \"unknown-nop\" not found in executor pool",
		},
		{
			name: "pyroscope URL in production returns error",
			env:  deployment.Environment{Name: "mainnet", BlockChains: newExecutorTestBlockChains([]uint64{sel1})},
			input: changesets.ApplyExecutorConfigInput{
				Topology: func() *offchain.EnvironmentTopology {
					topo := newMinimalTopology([]string{"nop1"}, "pool1", "")
					topo.PyroscopeURL = "http://pyroscope"
					return topo
				}(),
				ExecutorQualifier: "pool1",
			},
			wantErr: "pyroscope URL is not supported for production environments",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cs := changesets.ApplyExecutorConfig(registry)
			err := cs.VerifyPreconditions(tt.env, tt.input)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.wantErr)
		})
	}
}

func TestApplyExecutorConfig_HappyPathBuildsJobSpecs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {
				OffRampAddress:       "0xOffRamp",
				RmnAddress:           "0xRmn",
				ExecutorProxyAddress: "0xExecutorProxy",
			},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newMinimalTopology([]string{"nop1", "nop2"}, "pool1", shared.NOPModeStandalone)
	env := newTestExecutorEnv(t, []uint64{sel1})

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
}

func TestApplyExecutorConfig_NoDeployedChainsReturnsEmpty(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{},
		chainConfigs:   map[uint64]adapters.ExecutorChainConfig{},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newMinimalTopology([]string{"nop1"}, "pool1", shared.NOPModeStandalone)
	env := newTestExecutorEnv(t, []uint64{sel1})

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
}

func TestApplyExecutorConfig_BuildChainConfigErrorPropagates(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		buildErr: assert.AnError,
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newMinimalTopology([]string{"nop1"}, "pool1", shared.NOPModeStandalone)
	env := newTestExecutorEnv(t, []uint64{sel1})

	cs := changesets.ApplyExecutorConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to build config for chain")
}

func TestApplyExecutorConfig_UsesAllDeployedChains(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {
				OffRampAddress:       "0xOffRamp",
				RmnAddress:           "0xRmn",
				ExecutorProxyAddress: "0xExecutorProxy",
			},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newMinimalTopology([]string{"nop1"}, "pool1", shared.NOPModeStandalone)
	env := newTestExecutorEnv(t, []uint64{sel1, sel2})

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
}

func TestApplyExecutorConfig_TargetNOPsFiltersJobSpecs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {
				OffRampAddress:       "0xOffRamp",
				RmnAddress:           "0xRmn",
				ExecutorProxyAddress: "0xExecutorProxy",
			},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newMinimalTopology([]string{"nop1", "nop2", "nop3"}, "pool1", shared.NOPModeStandalone)
	env := newTestExecutorEnv(t, []uint64{sel1})

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
		TargetNOPs:        []shared.NOPAlias{"nop1"},
	})
	require.NoError(t, err)
	assert.NotNil(t, output.DataStore)
}
