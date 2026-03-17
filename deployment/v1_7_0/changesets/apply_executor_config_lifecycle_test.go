package changesets_test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/shared"
)

func newExecutorTestEnvWithDS(t *testing.T, selectors []uint64, ds datastore.DataStore) deployment.Environment {
	t.Helper()
	lggr := logger.Test(t)
	return deployment.Environment{
		Name:        "test",
		BlockChains: newExecutorTestBlockChains(selectors),
		DataStore:   ds,
		Logger:      lggr,
		OperationsBundle: operations.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			operations.NewMemoryReporter(),
		),
	}
}

func TestApplyExecutorConfig_GeneratesValidJobSpec(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1, sel2},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOffRamp1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xProxy1"},
			sel2: {OffRampAddress: "0xOffRamp2", RmnAddress: "0xRmn2", ExecutorProxyAddress: "0xProxy2"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newTopologyWithChainConfigs(
		[]string{"nop1", "nop2"},
		"pool1",
		shared.NOPModeStandalone,
		map[string]offchain.ChainExecutorPoolConfig{
			sel1Str: {NOPAliases: []string{"nop1", "nop2"}, ExecutionInterval: 15 * time.Second},
			sel2Str: {NOPAliases: []string{"nop1", "nop2"}, ExecutionInterval: 15 * time.Second},
		},
	)
	topo.IndexerAddress = []string{"http://indexer:8100", "http://indexer:8101"}
	env := newTestExecutorEnv(t, []uint64{sel1, sel2})

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	job, err := offchain.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-pool1-executor"))
	require.NoError(t, err)
	spec := job.Spec

	assert.Contains(t, spec, `schemaVersion = 1`)
	assert.Contains(t, spec, `type = "ccvexecutor"`)
	assert.Contains(t, spec, `executorConfig = """`)
	assert.Contains(t, spec, `executor_id = "nop1"`)
	assert.Contains(t, spec, `indexer_address = ["http://indexer:8100", "http://indexer:8101"]`)
	assert.Contains(t, spec, `[chain_configuration]`)
	assert.Contains(t, spec, sel1Str)
	assert.Contains(t, spec, sel2Str)
}

func TestApplyExecutorConfig_PreservesExistingJobSpecs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOffRamp1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xProxy1"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		JobID:    "existing-verifier",
		NOPAlias: "existing-nop",
		Spec:     "existing-verifier-job-spec",
	}))

	topo := newMinimalTopology([]string{"nop1"}, "pool1", shared.NOPModeStandalone)
	env := newExecutorTestEnvWithDS(t, []uint64{sel1}, ds.Seal())

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	sealed := output.DataStore.Seal()
	_, err = offchain.GetJob(sealed, shared.NOPAlias("nop1"), shared.JobID("nop1-pool1-executor"))
	require.NoError(t, err)

	verifierJob, err := offchain.GetJob(sealed, shared.NOPAlias("existing-nop"), shared.JobID("existing-verifier"))
	require.NoError(t, err)
	assert.Equal(t, "existing-verifier-job-spec", verifierJob.Spec)
}

func TestApplyExecutorConfig_RemovesOrphanedJobSpecs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOffRamp1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xProxy1"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		JobID:    "removed-nop-pool1-executor",
		NOPAlias: "removed-nop",
		Spec:     "orphaned-job-spec",
		Mode:     shared.NOPModeStandalone,
	}))
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		JobID:    "nop1-custom-executor",
		NOPAlias: "nop1",
		Spec:     "other-qualifier-job-spec",
		Mode:     shared.NOPModeStandalone,
	}))

	topo := newMinimalTopology([]string{"nop1"}, "pool1", shared.NOPModeStandalone)
	env := newExecutorTestEnvWithDS(t, []uint64{sel1}, ds.Seal())

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:           topo,
		ExecutorQualifier:  "pool1",
		RevokeOrphanedJobs: true,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	sealed := output.DataStore.Seal()
	_, err = offchain.GetJob(sealed, shared.NOPAlias("removed-nop"), shared.JobID("removed-nop-pool1-executor"))
	require.Error(t, err)

	otherJob, err := offchain.GetJob(sealed, shared.NOPAlias("nop1"), shared.JobID("nop1-custom-executor"))
	require.NoError(t, err)
	assert.Equal(t, "other-qualifier-job-spec", otherJob.Spec)
}

func TestApplyExecutorConfig_PreservesOrphanedJobSpecsWhenRevokeOrphanedJobsFalse(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOffRamp1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xProxy1"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		JobID:    "removed-nop-pool1-executor",
		NOPAlias: "removed-nop",
		Spec:     "orphaned-job-spec",
		Mode:     shared.NOPModeStandalone,
	}))

	topo := newMinimalTopology([]string{"nop1"}, "pool1", shared.NOPModeStandalone)
	env := newExecutorTestEnvWithDS(t, []uint64{sel1}, ds.Seal())

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	sealed := output.DataStore.Seal()
	orphanedJob, err := offchain.GetJob(sealed, shared.NOPAlias("removed-nop"), shared.JobID("removed-nop-pool1-executor"))
	require.NoError(t, err)
	assert.Equal(t, "orphaned-job-spec", orphanedJob.Spec)
}

func TestApplyExecutorConfig_TargetNOPsScopingPreservesUntargetedJobs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOffRamp1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xProxy1"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		JobID:    "nop1-pool1-executor",
		NOPAlias: "nop1",
		Spec:     "nop1-job-spec",
	}))
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		JobID:    "nop2-pool1-executor",
		NOPAlias: "nop2",
		Spec:     "nop2-job-spec",
	}))

	topo := newMinimalTopology([]string{"nop1", "nop2"}, "pool1", shared.NOPModeStandalone)
	env := newExecutorTestEnvWithDS(t, []uint64{sel1}, ds.Seal())

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
		TargetNOPs:        []shared.NOPAlias{"nop1"},
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	sealed := output.DataStore.Seal()
	nop1Job, err := offchain.GetJob(sealed, shared.NOPAlias("nop1"), shared.JobID("nop1-pool1-executor"))
	require.NoError(t, err)
	assert.NotEqual(t, "nop1-job-spec", nop1Job.Spec)
	assert.Contains(t, nop1Job.Spec, "nop2")

	nop2Job, err := offchain.GetJob(sealed, shared.NOPAlias("nop2"), shared.JobID("nop2-pool1-executor"))
	require.NoError(t, err)
	assert.Equal(t, "nop2-job-spec", nop2Job.Spec)
}

func TestApplyExecutorConfig_OnlyIncludesDeployedChains(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	sel3 := chainsel.TEST_90000003.Selector
	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)
	sel3Str := fmt.Sprintf("%d", sel3)

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1, sel2},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOffRamp1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xProxy1"},
			sel2: {OffRampAddress: "0xOffRamp2", RmnAddress: "0xRmn2", ExecutorProxyAddress: "0xProxy2"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newTopologyWithChainConfigs(
		[]string{"nop1"},
		"pool1",
		shared.NOPModeStandalone,
		map[string]offchain.ChainExecutorPoolConfig{
			sel1Str: {NOPAliases: []string{"nop1"}, ExecutionInterval: 15 * time.Second},
			sel2Str: {NOPAliases: []string{"nop1"}, ExecutionInterval: 15 * time.Second},
			sel3Str: {NOPAliases: []string{"nop1"}, ExecutionInterval: 15 * time.Second},
		},
	)
	env := newTestExecutorEnv(t, []uint64{sel1, sel2, sel3})

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	job, err := offchain.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-pool1-executor"))
	require.NoError(t, err)

	assert.Contains(t, job.Spec, sel1Str)
	assert.Contains(t, job.Spec, sel2Str)
	assert.NotContains(t, job.Spec, sel3Str)
}

func TestApplyExecutorConfig_PoolMembershipIncludedInJobSpec(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOffRamp1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xProxy1"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	nops := []offchain.NOPConfig{
		{Alias: "nop1", Name: "nop1-name", Mode: shared.NOPModeStandalone},
		{Alias: "nop2", Name: "nop2-name", Mode: shared.NOPModeStandalone},
		{Alias: "nop3", Name: "nop3-name", Mode: shared.NOPModeStandalone},
	}
	sel1Str := fmt.Sprintf("%d", sel1)
	topo := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &offchain.NOPTopology{
			NOPs:       nops,
			Committees: map[string]offchain.CommitteeConfig{},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{
			"pool1": {
				ChainConfigs: map[string]offchain.ChainExecutorPoolConfig{
					sel1Str: {NOPAliases: []string{"nop1", "nop2", "nop3"}, ExecutionInterval: 15 * time.Second},
				},
			},
		},
	}

	env := newTestExecutorEnv(t, []uint64{sel1})

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	job, err := offchain.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-pool1-executor"))
	require.NoError(t, err)

	assert.Contains(t, job.Spec, "nop1")
	assert.Contains(t, job.Spec, "nop2")
	assert.Contains(t, job.Spec, "nop3")
}

func TestApplyExecutorConfig_PerChainPoolMembershipAndIntervalInJobSpec(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1, sel2},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOffRamp1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xProxy1"},
			sel2: {OffRampAddress: "0xOffRamp2", RmnAddress: "0xRmn2", ExecutorProxyAddress: "0xProxy2"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	nops := []offchain.NOPConfig{
		{Alias: "nop1", Name: "nop1-name", Mode: shared.NOPModeStandalone},
		{Alias: "nop2", Name: "nop2-name", Mode: shared.NOPModeStandalone},
		{Alias: "nop3", Name: "nop3-name", Mode: shared.NOPModeStandalone},
	}
	topo := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &offchain.NOPTopology{
			NOPs:       nops,
			Committees: map[string]offchain.CommitteeConfig{},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{
			"pool1": {
				ChainConfigs: map[string]offchain.ChainExecutorPoolConfig{
					sel1Str: {
						NOPAliases:        []string{"nop1", "nop2"},
						ExecutionInterval: 10 * time.Second,
					},
					sel2Str: {
						NOPAliases:        []string{"nop2", "nop3"},
						ExecutionInterval: 20 * time.Second,
					},
				},
			},
		},
	}

	env := newTestExecutorEnv(t, []uint64{sel1, sel2})

	cs := changesets.ApplyExecutorConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	sealed := output.DataStore.Seal()

	nop1Job, err := offchain.GetJob(sealed, shared.NOPAlias("nop1"), shared.JobID("nop1-pool1-executor"))
	require.NoError(t, err)
	assert.Contains(t, nop1Job.Spec, sel1Str)
	assert.Contains(t, nop1Job.Spec, `executor_pool = ["nop1", "nop2"]`)
	assert.Contains(t, nop1Job.Spec, `execution_interval = "10s"`)

	nop2Job, err := offchain.GetJob(sealed, shared.NOPAlias("nop2"), shared.JobID("nop2-pool1-executor"))
	require.NoError(t, err)
	assert.Contains(t, nop2Job.Spec, sel1Str)
	assert.Contains(t, nop2Job.Spec, sel2Str)
	assert.Contains(t, nop2Job.Spec, `executor_pool = ["nop1", "nop2"]`)
	assert.Contains(t, nop2Job.Spec, `executor_pool = ["nop2", "nop3"]`)
	assert.Contains(t, nop2Job.Spec, `execution_interval = "10s"`)
	assert.Contains(t, nop2Job.Spec, `execution_interval = "20s"`)

	nop3Job, err := offchain.GetJob(sealed, shared.NOPAlias("nop3"), shared.JobID("nop3-pool1-executor"))
	require.NoError(t, err)
	assert.Contains(t, nop3Job.Spec, sel2Str)
	assert.Contains(t, nop3Job.Spec, `executor_pool = ["nop2", "nop3"]`)
	assert.Contains(t, nop3Job.Spec, `execution_interval = "20s"`)
}

func TestApplyExecutorConfig_FailsWhenPoolChainHasNoDeployedContracts(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1, sel2},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOffRamp1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xProxy1"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	nops := []offchain.NOPConfig{
		{Alias: "nop1", Name: "nop1-name", Mode: shared.NOPModeStandalone},
		{Alias: "nop2", Name: "nop2-name", Mode: shared.NOPModeStandalone},
	}
	topo := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &offchain.NOPTopology{
			NOPs:       nops,
			Committees: map[string]offchain.CommitteeConfig{},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{
			"pool1": {
				ChainConfigs: map[string]offchain.ChainExecutorPoolConfig{
					sel1Str: {NOPAliases: []string{"nop1", "nop2"}, ExecutionInterval: 15 * time.Second},
					sel2Str: {NOPAliases: []string{"nop1", "nop2"}, ExecutionInterval: 15 * time.Second},
				},
			},
		},
	}

	env := newTestExecutorEnv(t, []uint64{sel1})

	cs := changesets.ApplyExecutorConfig(registry)
	_, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to build config")
}

func TestApplyExecutorConfig_UpdatesJobsWhenChainRemovedFromTopology(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1, sel2},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOffRamp1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xProxy1"},
			sel2: {OffRampAddress: "0xOffRamp2", RmnAddress: "0xRmn2", ExecutorProxyAddress: "0xProxy2"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newTopologyWithChainConfigs(
		[]string{"nop1", "nop2"},
		"pool1",
		shared.NOPModeStandalone,
		map[string]offchain.ChainExecutorPoolConfig{
			sel1Str: {NOPAliases: []string{"nop1", "nop2"}, ExecutionInterval: 15 * time.Second},
			sel2Str: {NOPAliases: []string{"nop1", "nop2"}, ExecutionInterval: 15 * time.Second},
		},
	)
	env := newTestExecutorEnv(t, []uint64{sel1, sel2})

	cs := changesets.ApplyExecutorConfig(registry)
	firstOutput, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	firstJob, err := offchain.GetJob(firstOutput.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-pool1-executor"))
	require.NoError(t, err)
	assert.True(t, strings.Contains(firstJob.Spec, sel1Str) && strings.Contains(firstJob.Spec, sel2Str))

	mock.deployedChains["pool1"] = []uint64{sel1}

	topoOneChain := newTopologyWithChainConfigs(
		[]string{"nop1", "nop2"},
		"pool1",
		shared.NOPModeStandalone,
		map[string]offchain.ChainExecutorPoolConfig{
			sel1Str: {NOPAliases: []string{"nop1", "nop2"}, ExecutionInterval: 15 * time.Second},
		},
	)
	env.DataStore = firstOutput.DataStore.Seal()
	secondOutput, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          topoOneChain,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	secondJob, err := offchain.GetJob(secondOutput.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-pool1-executor"))
	require.NoError(t, err)
	assert.Contains(t, secondJob.Spec, sel1Str)
	assert.NotContains(t, secondJob.Spec, sel2Str)
}

func TestApplyExecutorConfig_AllNOPsUpdatedWhenMemberAddedToPool(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel1Str := fmt.Sprintf("%d", sel1)

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOffRamp1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xProxy1"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	initialTopo := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{Alias: "nop1", Name: "nop1-name", Mode: shared.NOPModeStandalone},
				{Alias: "nop2", Name: "nop2-name", Mode: shared.NOPModeStandalone},
			},
			Committees: map[string]offchain.CommitteeConfig{},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{
			"pool1": {ChainConfigs: map[string]offchain.ChainExecutorPoolConfig{
				sel1Str: {NOPAliases: []string{"nop1", "nop2"}, ExecutionInterval: 15 * time.Second},
			}},
		},
	}

	env := newTestExecutorEnv(t, []uint64{sel1})

	cs := changesets.ApplyExecutorConfig(registry)
	firstOutput, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          initialTopo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	firstSealed := firstOutput.DataStore.Seal()
	nop1Job, err := offchain.GetJob(firstSealed, shared.NOPAlias("nop1"), shared.JobID("nop1-pool1-executor"))
	require.NoError(t, err)
	nop2Job, err := offchain.GetJob(firstSealed, shared.NOPAlias("nop2"), shared.JobID("nop2-pool1-executor"))
	require.NoError(t, err)

	assert.Contains(t, nop1Job.Spec, "nop1")
	assert.Contains(t, nop1Job.Spec, "nop2")
	assert.NotContains(t, nop1Job.Spec, "nop3")
	assert.Contains(t, nop2Job.Spec, "nop1")
	assert.Contains(t, nop2Job.Spec, "nop2")
	assert.NotContains(t, nop2Job.Spec, "nop3")

	expandedTopo := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{Alias: "nop1", Name: "nop1-name", Mode: shared.NOPModeStandalone},
				{Alias: "nop2", Name: "nop2-name", Mode: shared.NOPModeStandalone},
				{Alias: "nop3", Name: "nop3-name", Mode: shared.NOPModeStandalone},
			},
			Committees: map[string]offchain.CommitteeConfig{},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{
			"pool1": {ChainConfigs: map[string]offchain.ChainExecutorPoolConfig{
				sel1Str: {NOPAliases: []string{"nop1", "nop2", "nop3"}, ExecutionInterval: 15 * time.Second},
			}},
		},
	}

	env.DataStore = firstSealed
	secondOutput, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          expandedTopo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	secondSealed := secondOutput.DataStore.Seal()
	nop1Updated, err := offchain.GetJob(secondSealed, shared.NOPAlias("nop1"), shared.JobID("nop1-pool1-executor"))
	require.NoError(t, err)
	nop2Updated, err := offchain.GetJob(secondSealed, shared.NOPAlias("nop2"), shared.JobID("nop2-pool1-executor"))
	require.NoError(t, err)
	nop3Job, err := offchain.GetJob(secondSealed, shared.NOPAlias("nop3"), shared.JobID("nop3-pool1-executor"))
	require.NoError(t, err)

	assert.Contains(t, nop1Updated.Spec, "nop3")
	assert.Contains(t, nop2Updated.Spec, "nop3")
	assert.Contains(t, nop3Job.Spec, "nop1")
	assert.Contains(t, nop3Job.Spec, "nop2")
}

func TestApplyExecutorConfig_AllNOPsUpdatedWhenMemberRemovedFromPool(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel1Str := fmt.Sprintf("%d", sel1)

	mock := &mockExecutorConfigAdapter{
		deployedChains: map[string][]uint64{
			"pool1": {sel1},
		},
		chainConfigs: map[uint64]adapters.ExecutorChainConfig{
			sel1: {OffRampAddress: "0xOffRamp1", RmnAddress: "0xRmn1", ExecutorProxyAddress: "0xProxy1"},
		},
	}

	registry := adapters.NewExecutorConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	initialTopo := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{Alias: "nop1", Name: "nop1-name", Mode: shared.NOPModeStandalone},
				{Alias: "nop2", Name: "nop2-name", Mode: shared.NOPModeStandalone},
				{Alias: "nop3", Name: "nop3-name", Mode: shared.NOPModeStandalone},
			},
			Committees: map[string]offchain.CommitteeConfig{},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{
			"pool1": {ChainConfigs: map[string]offchain.ChainExecutorPoolConfig{
				sel1Str: {NOPAliases: []string{"nop1", "nop2", "nop3"}, ExecutionInterval: 15 * time.Second},
			}},
		},
	}

	env := newTestExecutorEnv(t, []uint64{sel1})

	cs := changesets.ApplyExecutorConfig(registry)
	firstOutput, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:          initialTopo,
		ExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	firstSealed := firstOutput.DataStore.Seal()
	for _, nop := range []shared.NOPAlias{"nop1", "nop2", "nop3"} {
		job, err := offchain.GetJob(firstSealed, nop, shared.JobID(string(nop)+"-pool1-executor"))
		require.NoError(t, err)
		assert.Contains(t, job.Spec, "nop1")
		assert.Contains(t, job.Spec, "nop2")
		assert.Contains(t, job.Spec, "nop3")
	}

	reducedTopo := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &offchain.NOPTopology{
			NOPs: []offchain.NOPConfig{
				{Alias: "nop1", Name: "nop1-name", Mode: shared.NOPModeStandalone},
				{Alias: "nop2", Name: "nop2-name", Mode: shared.NOPModeStandalone},
				{Alias: "nop3", Name: "nop3-name", Mode: shared.NOPModeStandalone},
			},
			Committees: map[string]offchain.CommitteeConfig{},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{
			"pool1": {ChainConfigs: map[string]offchain.ChainExecutorPoolConfig{
				sel1Str: {NOPAliases: []string{"nop1", "nop2"}, ExecutionInterval: 15 * time.Second},
			}},
		},
	}

	env.DataStore = firstSealed
	secondOutput, err := cs.Apply(env, changesets.ApplyExecutorConfigInput{
		Topology:           reducedTopo,
		ExecutorQualifier:  "pool1",
		RevokeOrphanedJobs: true,
	})
	require.NoError(t, err)

	secondSealed := secondOutput.DataStore.Seal()
	nop1Updated, err := offchain.GetJob(secondSealed, shared.NOPAlias("nop1"), shared.JobID("nop1-pool1-executor"))
	require.NoError(t, err)
	nop2Updated, err := offchain.GetJob(secondSealed, shared.NOPAlias("nop2"), shared.JobID("nop2-pool1-executor"))
	require.NoError(t, err)

	assert.Contains(t, nop1Updated.Spec, "nop1")
	assert.Contains(t, nop1Updated.Spec, "nop2")
	assert.NotContains(t, nop1Updated.Spec, "nop3")
	assert.Contains(t, nop2Updated.Spec, "nop1")
	assert.Contains(t, nop2Updated.Spec, "nop2")
	assert.NotContains(t, nop2Updated.Spec, "nop3")

	_, err = offchain.GetJob(secondSealed, shared.NOPAlias("nop3"), shared.JobID("nop3-pool1-executor"))
	require.Error(t, err)
}
