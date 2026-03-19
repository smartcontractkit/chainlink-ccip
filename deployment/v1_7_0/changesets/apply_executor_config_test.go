package changesets_test

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/BurntSushi/toml"
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
	dummyChain := strconv.FormatUint(chainsel.TEST_90000001.Selector, 10)
	return &offchain.EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &offchain.NOPTopology{
			NOPs:       nops,
			Committees: map[string]offchain.CommitteeConfig{},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{
			executorQualifier: {
				ChainConfigs: map[string]offchain.ChainExecutorPoolConfig{
					dummyChain: {NOPAliases: nopAliases, ExecutionInterval: 15_000_000_000},
				},
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
					ExecutorPools: map[string]offchain.ExecutorPoolConfig{"pool1": {
						ChainConfigs: map[string]offchain.ChainExecutorPoolConfig{
							strconv.FormatUint(sel1, 10): {NOPAliases: []string{"nop1"}, ExecutionInterval: 15_000_000_000},
						},
					}},
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
	require.NotNil(t, output.DataStore)

	sealed := output.DataStore.Seal()
	sel1Str := fmt.Sprintf("%d", sel1)
	for _, nop := range []string{"nop1", "nop2"} {
		jobID := shared.NewExecutorJobID(shared.NOPAlias(nop), shared.ExecutorJobScope{ExecutorQualifier: "pool1"})
		job, err := offchain.GetJob(sealed, shared.NOPAlias(nop), jobID.ToJobID())
		require.NoError(t, err, "job should exist for %s", nop)
		assert.Contains(t, job.Spec, `type = "ccvexecutor"`)
		assert.Contains(t, job.Spec, sel1Str)
	}
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
	require.NotNil(t, output.DataStore)

	sealed := output.DataStore.Seal()
	scope := shared.ExecutorJobScope{ExecutorQualifier: "pool1"}

	nop1JobID := shared.NewExecutorJobID(shared.NOPAlias("nop1"), scope)
	_, err = offchain.GetJob(sealed, shared.NOPAlias("nop1"), nop1JobID.ToJobID())
	require.NoError(t, err, "targeted nop1 should have a job")

	nop2JobID := shared.NewExecutorJobID(shared.NOPAlias("nop2"), scope)
	_, err = offchain.GetJob(sealed, shared.NOPAlias("nop2"), nop2JobID.ToJobID())
	require.Error(t, err, "untargeted nop2 should not have a job")

	nop3JobID := shared.NewExecutorJobID(shared.NOPAlias("nop3"), scope)
	_, err = offchain.GetJob(sealed, shared.NOPAlias("nop3"), nop3JobID.ToJobID())
	require.Error(t, err, "untargeted nop3 should not have a job")
}

func newTopologyWithChainConfigs(
	nopAliases []string,
	qualifier string,
	mode shared.NOPMode,
	chainConfigs map[string]offchain.ChainExecutorPoolConfig,
) *offchain.EnvironmentTopology {
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
			qualifier: {
				ChainConfigs: chainConfigs,
			},
		},
	}
}

func extractExecutorChainSelectors(t *testing.T, ds datastore.DataStore, nopAlias shared.NOPAlias, qualifier string) []string {
	t.Helper()
	jobID := shared.NewExecutorJobID(nopAlias, shared.ExecutorJobScope{ExecutorQualifier: qualifier})
	job, err := offchain.GetJob(ds, nopAlias, jobID.ToJobID())
	require.NoError(t, err)

	var cfg offchain.ExecutorConfiguration
	_, err = toml.Decode(extractExecutorConfig(t, job.Spec), &cfg)
	require.NoError(t, err)

	chains := make([]string, 0, len(cfg.ChainConfiguration))
	for sel := range cfg.ChainConfiguration {
		chains = append(chains, sel)
	}
	return chains
}

func extractExecutorConfig(t *testing.T, jobSpec string) string {
	t.Helper()
	const marker = "executorConfig = \"\"\"\n"
	start := len(jobSpec)
	for i := 0; i < len(jobSpec)-len(marker); i++ {
		if jobSpec[i:i+len(marker)] == marker {
			start = i + len(marker)
			break
		}
	}
	end := len(jobSpec)
	for i := start; i < len(jobSpec)-3; i++ {
		if jobSpec[i:i+3] == "\"\"\"" {
			end = i
			break
		}
	}
	return jobSpec[start:end]
}

func TestApplyExecutorConfig_PerChainNOPsIncludesAllChains(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	sel3 := chainsel.TEST_90000003.Selector
	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)
	sel3Str := fmt.Sprintf("%d", sel3)

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1, sel2, sel3},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOff1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xExec1"},
			sel2: {OffRampAddress: "0xOff2", RmnAddress: "0xRmn2", ExecutorProxyAddress: "0xExec2"},
			sel3: {OffRampAddress: "0xOff3", RmnAddress: "0xRmn3", ExecutorProxyAddress: "0xExec3"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newTopologyWithChainConfigs(
		[]string{"exec1", "exec2"},
		"pool1",
		shared.NOPModeStandalone,
		map[string]offchain.ChainExecutorPoolConfig{
			sel1Str: {NOPAliases: []string{"exec1"}, ExecutionInterval: 15_000_000_000},
			sel2Str: {NOPAliases: []string{"exec1", "exec2"}, ExecutionInterval: 15_000_000_000},
			sel3Str: {NOPAliases: []string{"exec2"}, ExecutionInterval: 15_000_000_000},
		},
	)

	env := newTestExecutorEnv(t, []uint64{sel1, sel2, sel3})
	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	exec1Chains := extractExecutorChainSelectors(t, output.DataStore.Seal(), "exec1", "pool1")
	assert.ElementsMatch(t, []string{sel1Str, sel2Str, sel3Str}, exec1Chains,
		"exec1 should have all chains when listed on all chain configs")

	exec2Chains := extractExecutorChainSelectors(t, output.DataStore.Seal(), "exec2", "pool1")
	assert.ElementsMatch(t, []string{sel2Str, sel3Str, sel1Str}, exec2Chains,
		"exec2 should have all chains when listed on all chain configs")
}

func TestApplyExecutorConfig_AllChainsNOPsGetAllChains(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1, sel2},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOff1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xExec1"},
			sel2: {OffRampAddress: "0xOff2", RmnAddress: "0xRmn2", ExecutorProxyAddress: "0xExec2"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newTopologyWithChainConfigs(
		[]string{"exec1", "exec2"},
		"pool1",
		shared.NOPModeStandalone,
		map[string]offchain.ChainExecutorPoolConfig{
			sel1Str: {NOPAliases: []string{"exec1", "exec2"}, ExecutionInterval: 15_000_000_000},
			sel2Str: {NOPAliases: []string{"exec1", "exec2"}, ExecutionInterval: 15_000_000_000},
		},
	)
	env := newTestExecutorEnv(t, []uint64{sel1, sel2})

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	exec1Chains := extractExecutorChainSelectors(t, output.DataStore.Seal(), "exec1", "pool1")
	assert.ElementsMatch(t, []string{sel1Str, sel2Str}, exec1Chains,
		"exec1 should have all chains when listed on all chain configs")

	exec2Chains := extractExecutorChainSelectors(t, output.DataStore.Seal(), "exec2", "pool1")
	assert.ElementsMatch(t, []string{sel1Str, sel2Str}, exec2Chains,
		"exec2 should have all chains when listed on all chain configs")
}

func TestApplyExecutorConfig_ExecutorPoolFieldMatchesChainPool(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1, sel2},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOff1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xExec1"},
			sel2: {OffRampAddress: "0xOff2", RmnAddress: "0xRmn2", ExecutorProxyAddress: "0xExec2"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newTopologyWithChainConfigs(
		[]string{"exec1", "exec2"},
		"pool1",
		shared.NOPModeStandalone,
		map[string]offchain.ChainExecutorPoolConfig{
			sel1Str: {NOPAliases: []string{"exec1", "exec2"}, ExecutionInterval: 15_000_000_000},
			sel2Str: {NOPAliases: []string{"exec2"}, ExecutionInterval: 15_000_000_000},
		},
	)

	env := newTestExecutorEnv(t, []uint64{sel1, sel2})
	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	jobID := shared.NewExecutorJobID("exec2", shared.ExecutorJobScope{ExecutorQualifier: "pool1"})
	job, err := offchain.GetJob(output.DataStore.Seal(), "exec2", jobID.ToJobID())
	require.NoError(t, err)

	var cfg offchain.ExecutorConfiguration
	_, err = toml.Decode(extractExecutorConfig(t, job.Spec), &cfg)
	require.NoError(t, err)

	chain1Cfg, ok := cfg.ChainConfiguration[sel1Str]
	require.True(t, ok, "exec2 should have chain %s", sel1Str)
	assert.ElementsMatch(t, []string{"exec1", "exec2"}, chain1Cfg.ExecutorPool)

	chain2Cfg, ok := cfg.ChainConfiguration[sel2Str]
	require.True(t, ok, "exec2 should have chain %s", sel2Str)
	assert.ElementsMatch(t, []string{"exec2"}, chain2Cfg.ExecutorPool)
}
