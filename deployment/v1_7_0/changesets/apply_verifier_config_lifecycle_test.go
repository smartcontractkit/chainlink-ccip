package changesets_test

import (
	"context"
	"fmt"
	"strings"
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

func newVerifierTestEnvWithDS(t *testing.T, selectors []uint64, ds datastore.DataStore) deployment.Environment {
	t.Helper()
	lggr := logger.Test(t)
	chains := make(map[uint64]cldf_chain.BlockChain, len(selectors))
	for _, sel := range selectors {
		chains[sel] = cldfevm.Chain{Selector: sel}
	}
	return deployment.Environment{
		Name:        "test",
		BlockChains: cldf_chain.NewBlockChains(chains),
		DataStore:   ds,
		Logger:      lggr,
		OperationsBundle: operations.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			operations.NewMemoryReporter(),
		),
	}
}

func newVerifierTopologyWithAggregators(
	nopAliases []string,
	committeeQualifier string,
	selectors []uint64,
	mode shared.NOPMode,
	aggregators []offchain.AggregatorConfig,
	chainConfigs map[string]offchain.ChainCommitteeConfig,
) *offchain.EnvironmentTopology {
	if chainConfigs == nil {
		chainConfigs = make(map[string]offchain.ChainCommitteeConfig, len(selectors))
		for _, sel := range selectors {
			chainConfigs[fmt.Sprintf("%d", sel)] = offchain.ChainCommitteeConfig{
				NOPAliases: nopAliases,
				Threshold:  1,
			}
		}
	}
	if len(aggregators) == 0 {
		aggregators = []offchain.AggregatorConfig{
			{Name: "agg-1", Address: "ws://agg:9090"},
		}
	}
	nops := make([]offchain.NOPConfig, len(nopAliases))
	for i, alias := range nopAliases {
		nops[i] = offchain.NOPConfig{
			Alias:                 alias,
			Name:                  alias + "-name",
			Mode:                  mode,
			SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xabc123"},
		}
	}
	return &offchain.EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &offchain.NOPTopology{
			NOPs: nops,
			Committees: map[string]offchain.CommitteeConfig{
				committeeQualifier: {
					Qualifier:    committeeQualifier,
					ChainConfigs: chainConfigs,
					Aggregators:  aggregators,
				},
			},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{},
	}
}

func TestApplyVerifierConfig_GeneratesValidJobSpec(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, shared.NOPModeStandalone)
	env := newVerifierTestEnv(t, []uint64{sel1})

	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	job, err := offchain.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)
	spec := job.Spec

	assert.Contains(t, spec, `schemaVersion = 1`)
	assert.Contains(t, spec, `type = "ccvcommitteeverifier"`)
	assert.Contains(t, spec, `committeeVerifierConfig = """`)
	assert.Contains(t, spec, fmt.Sprintf("%d", sel1))
	assert.Contains(t, spec, "0xCommitteeVerifier")
	assert.Contains(t, spec, "0xOnRamp")
	assert.Contains(t, spec, "0xExecutorProxy")
	assert.Contains(t, spec, "0xRMNRemote")
	assert.Contains(t, spec, "0xabc123")
	assert.Contains(t, spec, "ws://agg:9090")
}

func TestApplyVerifierConfig_PreservesExistingJobSpecs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		Spec:     "existing-executor-job-spec",
		JobID:    shared.JobID("existing-nop-pool1-executor"),
		NOPAlias: shared.NOPAlias("existing-nop"),
	}))
	env := newVerifierTestEnvWithDS(t, []uint64{sel1}, ds.Seal())

	topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, shared.NOPModeStandalone)
	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	sealed := output.DataStore.Seal()
	_, err = offchain.GetJob(sealed, shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)

	executorJob, err := offchain.GetJob(sealed, shared.NOPAlias("existing-nop"), shared.JobID("existing-nop-pool1-executor"))
	require.NoError(t, err)
	assert.Equal(t, "existing-executor-job-spec", executorJob.Spec)
}

func TestApplyVerifierConfig_MultipleAggregators(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	aggregators := []offchain.AggregatorConfig{
		{Name: "agg-primary", Address: "ws://agg1:9090"},
		{Name: "agg-secondary", Address: "ws://agg2:9090"},
		{Name: "agg-tertiary", Address: "ws://agg3:9090"},
	}
	topo := newVerifierTopologyWithAggregators([]string{"nop1"}, "c1", []uint64{sel1}, shared.NOPModeStandalone, aggregators, nil)
	env := newVerifierTestEnv(t, []uint64{sel1})

	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	sealed := output.DataStore.Seal()
	job1, err := offchain.GetJob(sealed, shared.NOPAlias("nop1"), shared.JobID("nop1-agg-primary-c1-verifier"))
	require.NoError(t, err)
	assert.Contains(t, job1.Spec, "ws://agg1:9090")

	job2, err := offchain.GetJob(sealed, shared.NOPAlias("nop1"), shared.JobID("nop1-agg-secondary-c1-verifier"))
	require.NoError(t, err)
	assert.Contains(t, job2.Spec, "ws://agg2:9090")

	job3, err := offchain.GetJob(sealed, shared.NOPAlias("nop1"), shared.JobID("nop1-agg-tertiary-c1-verifier"))
	require.NoError(t, err)
	assert.Contains(t, job3.Spec, "ws://agg3:9090")
}

func TestApplyVerifierConfig_RemovesOrphanedJobSpecs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		Spec:     "orphaned-job-spec",
		JobID:    shared.JobID("removed-nop-agg-1-c1-verifier"),
		NOPAlias: shared.NOPAlias("removed-nop"),
		Mode:     shared.NOPModeStandalone,
	}))
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		Spec:     "other-committee-job-spec",
		JobID:    shared.JobID("nop1-agg-1-other-committee-verifier"),
		NOPAlias: shared.NOPAlias("nop1"),
		Mode:     shared.NOPModeStandalone,
	}))
	env := newVerifierTestEnvWithDS(t, []uint64{sel1}, ds.Seal())

	topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, shared.NOPModeStandalone)
	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
		RevokeOrphanedJobs:       true,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	sealed := output.DataStore.Seal()
	_, err = offchain.GetJob(sealed, shared.NOPAlias("removed-nop"), shared.JobID("removed-nop-agg-1-c1-verifier"))
	require.Error(t, err)

	otherJob, err := offchain.GetJob(sealed, shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-other-committee-verifier"))
	require.NoError(t, err)
	assert.Equal(t, "other-committee-job-spec", otherJob.Spec)
}

func TestApplyVerifierConfig_PreservesOrphanedJobSpecsWhenRevokeOrphanedJobsFalse(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		Spec:     "orphaned-job-spec",
		JobID:    shared.JobID("removed-nop-agg-1-c1-verifier"),
		NOPAlias: shared.NOPAlias("removed-nop"),
		Mode:     shared.NOPModeStandalone,
	}))
	env := newVerifierTestEnvWithDS(t, []uint64{sel1}, ds.Seal())

	topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, shared.NOPModeStandalone)
	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	sealed := output.DataStore.Seal()
	orphanedJob, err := offchain.GetJob(sealed, shared.NOPAlias("removed-nop"), shared.JobID("removed-nop-agg-1-c1-verifier"))
	require.NoError(t, err)
	assert.Equal(t, "orphaned-job-spec", orphanedJob.Spec)
}

func TestApplyVerifierConfig_TargetNOPsScopingPreservesUntargetedJobs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	ds := datastore.NewMemoryDataStore()
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		Spec:     "nop1-job-spec",
		JobID:    shared.JobID("nop1-agg-1-c1-verifier"),
		NOPAlias: shared.NOPAlias("nop1"),
	}))
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		Spec:     "nop2-job-spec",
		JobID:    shared.JobID("nop2-agg-1-c1-verifier"),
		NOPAlias: shared.NOPAlias("nop2"),
	}))
	env := newVerifierTestEnvWithDS(t, []uint64{sel1}, ds.Seal())

	topo := newVerifierTopology([]string{"nop1", "nop2"}, "c1", []uint64{sel1}, shared.NOPModeStandalone)
	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
		TargetNOPs:               []shared.NOPAlias{"nop1"},
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	sealed := output.DataStore.Seal()
	nop1Job, err := offchain.GetJob(sealed, shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)
	assert.NotEqual(t, "nop1-job-spec", nop1Job.Spec)

	nop2Job, err := offchain.GetJob(sealed, shared.NOPAlias("nop2"), shared.JobID("nop2-agg-1-c1-verifier"))
	require.NoError(t, err)
	assert.Equal(t, "nop2-job-spec", nop2Job.Spec)
}

func TestApplyVerifierConfig_OnlyIncludesCommitteeChains(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector
	sel3 := chainsel.TEST_90000003.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCV1",
				OnRampAddress:            "0xOnRamp1",
				ExecutorProxyAddress:     "0xExec1",
				RMNRemoteAddress:         "0xRMN1",
			},
			sel2: {
				CommitteeVerifierAddress: "0xCV2",
				OnRampAddress:            "0xOnRamp2",
				ExecutorProxyAddress:     "0xExec2",
				RMNRemoteAddress:         "0xRMN2",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	chainConfigs := map[string]offchain.ChainCommitteeConfig{
		fmt.Sprintf("%d", sel1): {NOPAliases: []string{"nop1"}, Threshold: 1},
		fmt.Sprintf("%d", sel2): {NOPAliases: []string{"nop1"}, Threshold: 1},
	}
	topo := newVerifierTopologyWithAggregators([]string{"nop1"}, "c1", nil, shared.NOPModeStandalone, nil, chainConfigs)
	env := newVerifierTestEnv(t, []uint64{sel1, sel2, sel3})

	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	job, err := offchain.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)

	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)
	sel3Str := fmt.Sprintf("%d", sel3)
	assert.Contains(t, job.Spec, sel1Str)
	assert.Contains(t, job.Spec, sel2Str)
	assert.NotContains(t, job.Spec, sel3Str)
}

func TestApplyVerifierConfig_FiltersChainsPerNOPMembership(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCV1",
				OnRampAddress:            "0xOnRamp1",
				ExecutorProxyAddress:     "0xExec1",
				RMNRemoteAddress:         "0xRMN1",
			},
			sel2: {
				CommitteeVerifierAddress: "0xCV2",
				OnRampAddress:            "0xOnRamp2",
				ExecutorProxyAddress:     "0xExec2",
				RMNRemoteAddress:         "0xRMN2",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	chainConfigs := map[string]offchain.ChainCommitteeConfig{
		fmt.Sprintf("%d", sel1): {NOPAliases: []string{"nop1"}, Threshold: 1},
		fmt.Sprintf("%d", sel2): {NOPAliases: []string{"nop2"}, Threshold: 1},
	}
	topo := newVerifierTopologyWithAggregators([]string{"nop1", "nop2"}, "c1", nil, shared.NOPModeStandalone, nil, chainConfigs)
	env := newVerifierTestEnv(t, []uint64{sel1, sel2})

	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	sealed := output.DataStore.Seal()
	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)

	nop1Job, err := offchain.GetJob(sealed, shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)
	assert.Contains(t, nop1Job.Spec, sel1Str)
	assert.NotContains(t, nop1Job.Spec, sel2Str)

	nop2Job, err := offchain.GetJob(sealed, shared.NOPAlias("nop2"), shared.JobID("nop2-agg-1-c1-verifier"))
	require.NoError(t, err)
	assert.NotContains(t, nop2Job.Spec, sel1Str)
	assert.Contains(t, nop2Job.Spec, sel2Str)
}

func TestApplyVerifierConfig_SkipsNOPsNotInAnyChainConfig(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	chainConfigs := map[string]offchain.ChainCommitteeConfig{
		fmt.Sprintf("%d", sel1): {NOPAliases: []string{"nop1"}, Threshold: 1},
	}
	nops := []offchain.NOPConfig{
		{Alias: "nop1", Name: "nop1", Mode: shared.NOPModeStandalone, SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xabc123"}},
		{Alias: "nop-not-in-committee", Name: "nop-not-in-committee", Mode: shared.NOPModeStandalone, SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xdef456"}},
	}
	topo := &offchain.EnvironmentTopology{
		IndexerAddress: []string{"http://indexer:8080"},
		NOPTopology: &offchain.NOPTopology{
			NOPs: nops,
			Committees: map[string]offchain.CommitteeConfig{
				"c1": {
					Qualifier:    "c1",
					ChainConfigs: chainConfigs,
					Aggregators:  []offchain.AggregatorConfig{{Name: "agg-1", Address: "ws://agg:9090"}},
				},
			},
		},
		ExecutorPools: map[string]offchain.ExecutorPoolConfig{},
	}
	env := newVerifierTestEnv(t, []uint64{sel1})

	cs := changesets.ApplyVerifierConfig(registry)
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
		TargetNOPs:               []shared.NOPAlias{"nop1", "nop-not-in-committee"},
	})
	require.NoError(t, err)

	sealed := output.DataStore.Seal()
	_, err = offchain.GetJob(sealed, shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)

	_, err = offchain.GetJob(sealed, shared.NOPAlias("nop-not-in-committee"), shared.JobID("nop-not-in-committee-agg-1-c1-verifier"))
	require.Error(t, err)
}

func TestApplyVerifierConfig_SkipsUnchangedJobSpecs(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCommitteeVerifier",
				OnRampAddress:            "0xOnRamp",
				ExecutorProxyAddress:     "0xExecutorProxy",
				RMNRemoteAddress:         "0xRMNRemote",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1}, shared.NOPModeStandalone)
	env := newVerifierTestEnv(t, []uint64{sel1})

	cs := changesets.ApplyVerifierConfig(registry)
	firstOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	firstJob, err := offchain.GetJob(firstOutput.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)
	firstSpec := firstJob.Spec

	env.DataStore = firstOutput.DataStore.Seal()
	secondOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	secondJob, err := offchain.GetJob(secondOutput.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)
	assert.Equal(t, firstSpec, secondJob.Spec)
}

func TestApplyVerifierConfig_UpdatesJobsWhenChainRemoved(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCV1",
				OnRampAddress:            "0xOnRamp1",
				ExecutorProxyAddress:     "0xExec1",
				RMNRemoteAddress:         "0xRMN1",
			},
			sel2: {
				CommitteeVerifierAddress: "0xCV2",
				OnRampAddress:            "0xOnRamp2",
				ExecutorProxyAddress:     "0xExec2",
				RMNRemoteAddress:         "0xRMN2",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topo := newVerifierTopology([]string{"nop1"}, "c1", []uint64{sel1, sel2}, shared.NOPModeStandalone)
	env := newVerifierTestEnv(t, []uint64{sel1, sel2})

	cs := changesets.ApplyVerifierConfig(registry)
	firstOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topo,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	firstJob, err := offchain.GetJob(firstOutput.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)
	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)
	assert.True(t, strings.Contains(firstJob.Spec, sel1Str) && strings.Contains(firstJob.Spec, sel2Str))

	chainConfigsOneChain := map[string]offchain.ChainCommitteeConfig{
		fmt.Sprintf("%d", sel1): {NOPAliases: []string{"nop1"}, Threshold: 1},
	}
	topoOneChain := newVerifierTopologyWithAggregators([]string{"nop1"}, "c1", nil, shared.NOPModeStandalone, nil, chainConfigsOneChain)

	env.DataStore = firstOutput.DataStore.Seal()
	secondOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topoOneChain,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	secondJob, err := offchain.GetJob(secondOutput.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)
	assert.Contains(t, secondJob.Spec, sel1Str)
	assert.NotContains(t, secondJob.Spec, sel2Str)
}

func TestApplyVerifierConfig_RemovesJobWhenNOPRemovedFromCommittee(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCV1",
				OnRampAddress:            "0xOnRamp1",
				ExecutorProxyAddress:     "0xExec1",
				RMNRemoteAddress:         "0xRMN1",
			},
			sel2: {
				CommitteeVerifierAddress: "0xCV2",
				OnRampAddress:            "0xOnRamp2",
				ExecutorProxyAddress:     "0xExec2",
				RMNRemoteAddress:         "0xRMN2",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	topoBoth := newVerifierTopology([]string{"nop1", "nop2"}, "c1", []uint64{sel1, sel2}, shared.NOPModeStandalone)
	env := newVerifierTestEnv(t, []uint64{sel1, sel2})

	cs := changesets.ApplyVerifierConfig(registry)
	firstOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topoBoth,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	_, err = offchain.GetJob(firstOutput.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)
	_, err = offchain.GetJob(firstOutput.DataStore.Seal(), shared.NOPAlias("nop2"), shared.JobID("nop2-agg-1-c1-verifier"))
	require.NoError(t, err)

	chainConfigsNOP1Only := map[string]offchain.ChainCommitteeConfig{
		fmt.Sprintf("%d", sel1): {NOPAliases: []string{"nop1"}, Threshold: 1},
		fmt.Sprintf("%d", sel2): {NOPAliases: []string{"nop1"}, Threshold: 1},
	}
	topoNOP1Only := newVerifierTopologyWithAggregators([]string{"nop1", "nop2"}, "c1", nil, shared.NOPModeStandalone, nil, chainConfigsNOP1Only)

	env.DataStore = firstOutput.DataStore.Seal()
	secondOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topoNOP1Only,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
		RevokeOrphanedJobs:       true,
	})
	require.NoError(t, err)

	_, err = offchain.GetJob(secondOutput.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)

	_, err = offchain.GetJob(secondOutput.DataStore.Seal(), shared.NOPAlias("nop2"), shared.JobID("nop2-agg-1-c1-verifier"))
	require.Error(t, err)
}

func TestApplyVerifierConfig_CreatesJobWhenNOPAddedToCommittee(t *testing.T) {
	sel1 := chainsel.TEST_90000001.Selector
	sel2 := chainsel.TEST_90000002.Selector

	mock := &mockVerifierJobConfigAdapter{
		chainConfigs: map[uint64]*adapters.VerifierContractAddresses{
			sel1: {
				CommitteeVerifierAddress: "0xCV1",
				OnRampAddress:            "0xOnRamp1",
				ExecutorProxyAddress:     "0xExec1",
				RMNRemoteAddress:         "0xRMN1",
			},
			sel2: {
				CommitteeVerifierAddress: "0xCV2",
				OnRampAddress:            "0xOnRamp2",
				ExecutorProxyAddress:     "0xExec2",
				RMNRemoteAddress:         "0xRMN2",
			},
		},
	}
	registry := adapters.NewVerifierConfigRegistry()
	registry.Register(chainsel.FamilyEVM, mock)

	chainConfigsNOP1Only := map[string]offchain.ChainCommitteeConfig{
		fmt.Sprintf("%d", sel1): {NOPAliases: []string{"nop1"}, Threshold: 1},
		fmt.Sprintf("%d", sel2): {NOPAliases: []string{"nop1"}, Threshold: 1},
	}
	topoNOP1Only := newVerifierTopologyWithAggregators([]string{"nop1", "nop2"}, "c1", nil, shared.NOPModeStandalone, nil, chainConfigsNOP1Only)
	env := newVerifierTestEnv(t, []uint64{sel1, sel2})

	cs := changesets.ApplyVerifierConfig(registry)
	firstOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topoNOP1Only,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	_, err = offchain.GetJob(firstOutput.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)
	_, err = offchain.GetJob(firstOutput.DataStore.Seal(), shared.NOPAlias("nop2"), shared.JobID("nop2-agg-1-c1-verifier"))
	require.Error(t, err)

	topoBoth := newVerifierTopology([]string{"nop1", "nop2"}, "c1", []uint64{sel1, sel2}, shared.NOPModeStandalone)

	env.DataStore = firstOutput.DataStore.Seal()
	secondOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigInput{
		Topology:                 topoBoth,
		CommitteeQualifier:       "c1",
		DefaultExecutorQualifier: "pool1",
	})
	require.NoError(t, err)

	_, err = offchain.GetJob(secondOutput.DataStore.Seal(), shared.NOPAlias("nop1"), shared.JobID("nop1-agg-1-c1-verifier"))
	require.NoError(t, err)

	nop2Job, err := offchain.GetJob(secondOutput.DataStore.Seal(), shared.NOPAlias("nop2"), shared.JobID("nop2-agg-1-c1-verifier"))
	require.NoError(t, err)
	sel1Str := fmt.Sprintf("%d", sel1)
	sel2Str := fmt.Sprintf("%d", sel2)
	assert.Contains(t, nop2Job.Spec, sel1Str)
	assert.Contains(t, nop2Job.Spec, sel2Str)
}
