package changesets_test

import (
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/ccv"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	deploymocks "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/topology"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/testutils"
)

func TestApplyExecutorConfig_Validation(t *testing.T) {
	tests := []struct {
		name        string
		setupEnv    func(t *testing.T) changesets.ApplyExecutorConfigCfg
		expectedErr string
	}{
		{
			name: "missing topology",
			setupEnv: func(t *testing.T) changesets.ApplyExecutorConfigCfg {
				return changesets.ApplyExecutorConfigCfg{
					Topology:          nil,
					ExecutorQualifier: testDefaultQualifier,
				}
			},
			expectedErr: "topology is required",
		},
		{
			name: "missing executor qualifier",
			setupEnv: func(t *testing.T) changesets.ApplyExecutorConfigCfg {
				return changesets.ApplyExecutorConfigCfg{
					Topology:          newTestTopology(),
					ExecutorQualifier: "",
				}
			},
			expectedErr: "executor qualifier is required",
		},
		{
			name: "missing indexer address",
			setupEnv: func(t *testing.T) changesets.ApplyExecutorConfigCfg {
				topology := newTestTopology()
				topology.IndexerAddress = []string{}
				return changesets.ApplyExecutorConfigCfg{
					Topology:          topology,
					ExecutorQualifier: testDefaultQualifier,
				}
			},
			expectedErr: "indexer address is required",
		},
		{
			name: "executor pool not found",
			setupEnv: func(t *testing.T) changesets.ApplyExecutorConfigCfg {
				return changesets.ApplyExecutorConfigCfg{
					Topology:          newTestTopology(),
					ExecutorQualifier: "non-existent-pool",
				}
			},
			expectedErr: "executor pool \"non-existent-pool\" not found in topology",
		},
		{
			name: "target NOP not in pool",
			setupEnv: func(t *testing.T) changesets.ApplyExecutorConfigCfg {
				return changesets.ApplyExecutorConfigCfg{
					Topology:          newTestTopology(),
					ExecutorQualifier: testDefaultQualifier,
					TargetNOPs:        []shared.NOPAlias{"unknown-nop"},
				}
			},
			expectedErr: "not found in executor pool",
		},
		{
			name: "chain selector not in environment",
			setupEnv: func(t *testing.T) changesets.ApplyExecutorConfigCfg {
				return changesets.ApplyExecutorConfigCfg{
					Topology:          newTestTopology(),
					ExecutorQualifier: testDefaultQualifier,
					ChainSelectors:    []uint64{999999999},
				}
			},
			expectedErr: "selector 999999999 is not available in environment",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			changeset := changesets.ApplyExecutorConfig()
			env := newTestEnvironment(t, defaultSelectors)
			cfg := tt.setupEnv(t)

			err := changeset.VerifyPreconditions(env, cfg)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}

func TestApplyExecutorConfig_ValidatesExecutorPoolNOPs(t *testing.T) {
	changeset := changesets.ApplyExecutorConfig()
	env := newTestEnvironment(t, defaultSelectors)

	topology := newTestTopology(
		WithExecutorPool(testDefaultQualifier, topology.ExecutorPoolConfig{
			NOPAliases: []string{},
		}),
	)

	err := changeset.VerifyPreconditions(env, changesets.ApplyExecutorConfigCfg{
		Topology:          topology,
		ExecutorQualifier: testDefaultQualifier,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "executor pool \"default\" has no NOPs")
}

func TestApplyExecutorConfig_PyroscopeNotAllowedInProduction(t *testing.T) {
	changeset := changesets.ApplyExecutorConfig()

	env := newTestEnvironment(t, defaultSelectors)
	env.Name = "mainnet"

	topology := newTestTopology(WithPyroscopeURL("http://pyroscope:4040"))

	err := changeset.VerifyPreconditions(env, changesets.ApplyExecutorConfigCfg{
		Topology:          topology,
		ExecutorQualifier: testDefaultQualifier,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "pyroscope URL is not supported for production environments")
}

func TestApplyExecutorConfig_GeneratesValidJobSpec(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupExecutorDatastore(t, ds, defaultSelectors, testDefaultQualifier)
	env.DataStore = ds.Seal()

	sel1Str := strconv.FormatUint(defaultSelectors[0], 10)
	sel2Str := strconv.FormatUint(defaultSelectors[1], 10)

	topology := newTestTopology(
		WithPyroscopeURL("http://pyroscope:4040"),
		WithMonitoring(defaultMonitoringConfig()),
		WithIndexerAddress([]string{"http://indexer:8100", "http://indexer:8101"}),
	)

	cs := changesets.ApplyExecutorConfig()
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigCfg{
		Topology:          topology,
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    defaultSelectors,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	job, err := ccv.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop-1"), shared.JobID("nop-1-default-executor"))
	require.NoError(t, err)
	jobSpec := job.Spec

	assert.Contains(t, jobSpec, `schemaVersion = 1`)
	assert.Contains(t, jobSpec, `type = "ccvexecutor"`)
	assert.Contains(t, jobSpec, `executorConfig = """`)
	assert.Contains(t, jobSpec, `executor_id = "nop-1"`)
	assert.Contains(t, jobSpec, `indexer_address = ["http://indexer:8100", "http://indexer:8101"]`)
	assert.Contains(t, jobSpec, `pyroscope_url = "http://pyroscope:4040"`)
	assert.Contains(t, jobSpec, `[chain_configuration]`)
	assert.Contains(t, jobSpec, sel1Str)
	assert.Contains(t, jobSpec, sel2Str)
	assert.Contains(t, jobSpec, `[Monitoring]`)
	assert.Contains(t, jobSpec, `Enabled = true`)
	assert.Contains(t, jobSpec, `Type = "beholder"`)
}

func TestApplyExecutorConfig_PreservesExistingJobSpecs(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupExecutorDatastore(t, ds, defaultSelectors, testDefaultQualifier)

	existingJobSpec := "existing-verifier-job-spec"
	err := ccv.SaveJob(ds, shared.JobInfo{
		Spec:     existingJobSpec,
		JobID:    "existing-verifier",
		NOPAlias: "existing-nop",
	})
	require.NoError(t, err)

	env.DataStore = ds.Seal()

	cs := changesets.ApplyExecutorConfig()
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigCfg{
		Topology:          newTestTopology(),
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    defaultSelectors,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	outputSealed := output.DataStore.Seal()

	_, err = ccv.GetJob(outputSealed, shared.NOPAlias("nop-1"), shared.JobID("nop-1-default-executor"))
	require.NoError(t, err, "new executor job spec should be present")

	retrievedJob, err := ccv.GetJob(outputSealed, shared.NOPAlias("existing-nop"), shared.JobID("existing-verifier"))
	require.NoError(t, err, "existing verifier job spec should be preserved")
	assert.Equal(t, existingJobSpec, retrievedJob.Spec, "verifier job spec should be unchanged")
}

func TestApplyExecutorConfig_RemovesOrphanedJobSpecs(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupExecutorDatastore(t, ds, defaultSelectors, testDefaultQualifier)

	err := ccv.SaveJob(ds, shared.JobInfo{
		Spec:     "orphaned-job-spec",
		JobID:    "removed-nop-default-executor",
		NOPAlias: "removed-nop",
		Mode:     shared.NOPModeStandalone,
	})
	require.NoError(t, err)

	err = ccv.SaveJob(ds, shared.JobInfo{
		Spec:     "other-qualifier-job-spec",
		JobID:    "nop-1-custom-executor",
		NOPAlias: "nop-1",
		Mode:     shared.NOPModeStandalone,
	})
	require.NoError(t, err)

	env.DataStore = ds.Seal()

	cs := changesets.ApplyExecutorConfig()
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigCfg{
		Topology:          newTestTopology(),
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    defaultSelectors,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	outputSealed := output.DataStore.Seal()

	_, err = ccv.GetJob(outputSealed, shared.NOPAlias("removed-nop"), shared.JobID("removed-nop-default-executor"))
	require.Error(t, err, "orphaned job spec should be removed")

	otherQualifierJob, err := ccv.GetJob(outputSealed, shared.NOPAlias("nop-1"), shared.JobID("nop-1-custom-executor"))
	require.NoError(t, err, "job spec for other qualifier should be preserved")
	assert.Equal(t, "other-qualifier-job-spec", otherQualifierJob.Spec)
}

func TestApplyExecutorConfig_TargetNOPsScoping(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupExecutorDatastore(t, ds, defaultSelectors, testDefaultQualifier)

	err := ccv.SaveJob(ds, shared.JobInfo{
		Spec:     "nop-1-job-spec",
		JobID:    "nop-1-default-executor",
		NOPAlias: "nop-1",
	})
	require.NoError(t, err)
	err = ccv.SaveJob(ds, shared.JobInfo{
		Spec:     "nop-2-job-spec",
		JobID:    "nop-2-default-executor",
		NOPAlias: "nop-2",
	})
	require.NoError(t, err)

	env.DataStore = ds.Seal()

	cs := changesets.ApplyExecutorConfig()
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigCfg{
		Topology:          newTestTopology(),
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    defaultSelectors,
		TargetNOPs:        []shared.NOPAlias{"nop-1"},
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	outputSealed := output.DataStore.Seal()

	job, err := ccv.GetJob(outputSealed, shared.NOPAlias("nop-1"), shared.JobID("nop-1-default-executor"))
	require.NoError(t, err, "nop-1 executor job spec should exist")
	assert.NotEqual(t, "nop-1-job-spec", job.Spec, "nop-1 job spec should be updated")

	nop2Job, err := ccv.GetJob(outputSealed, shared.NOPAlias("nop-2"), shared.JobID("nop-2-default-executor"))
	require.NoError(t, err, "nop-2 executor job spec should be preserved when not in scope")
	assert.Equal(t, "nop-2-job-spec", nop2Job.Spec, "nop-2 job spec should be unchanged")
}

func TestApplyExecutorConfig_UsesDeployedChainsWhenEmptyChainSelectors(t *testing.T) {
	selectors := []uint64{
		chainsel.TEST_90000001.Selector,
		chainsel.TEST_90000002.Selector,
		chainsel.TEST_90000003.Selector,
	}
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, selectors)
	setupExecutorDatastore(t, ds, selectors[:2], testDefaultQualifier)
	env.DataStore = ds.Seal()

	sel1Str := strconv.FormatUint(selectors[0], 10)
	sel2Str := strconv.FormatUint(selectors[1], 10)
	sel3Str := strconv.FormatUint(selectors[2], 10)

	cs := changesets.ApplyExecutorConfig()
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigCfg{
		Topology:          newTestTopology(),
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    nil,
	})
	require.NoError(t, err)

	job, err := ccv.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop-1"), shared.JobID("nop-1-default-executor"))
	require.NoError(t, err)

	assert.Contains(t, job.Spec, sel1Str, "should include chain where executor is deployed")
	assert.Contains(t, job.Spec, sel2Str, "should include chain where executor is deployed")
	assert.NotContains(t, job.Spec, sel3Str, "should NOT include chain where executor is NOT deployed")
}

func TestApplyExecutorConfig_PoolMembershipIncludedInJobSpec(t *testing.T) {
	selectors := []uint64{chainsel.TEST_90000001.Selector}
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, selectors)
	setupExecutorDatastore(t, ds, selectors, testDefaultQualifier)
	env.DataStore = ds.Seal()

	topology := newTestTopology(
		WithNOPs([]topology.NOPConfig{
			{Alias: "nop-1", Name: "nop-1", Mode: shared.NOPModeStandalone},
			{Alias: "nop-2", Name: "nop-2", Mode: shared.NOPModeStandalone},
			{Alias: "nop-3", Name: "nop-3", Mode: shared.NOPModeStandalone},
		}),
		WithExecutorPool(testDefaultQualifier, topology.ExecutorPoolConfig{
			NOPAliases:        []string{"nop-1", "nop-2", "nop-3"},
			ExecutionInterval: 15 * time.Second,
		}),
	)

	cs := changesets.ApplyExecutorConfig()
	output, err := cs.Apply(env, changesets.ApplyExecutorConfigCfg{
		Topology:          topology,
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    selectors,
	})
	require.NoError(t, err)

	job, err := ccv.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop-1"), shared.JobID("nop-1-default-executor"))
	require.NoError(t, err)

	assert.Contains(t, job.Spec, "nop-1", "job spec should include nop-1 in pool")
	assert.Contains(t, job.Spec, "nop-2", "job spec should include nop-2 in pool")
	assert.Contains(t, job.Spec, "nop-3", "job spec should include nop-3 in pool")
}

func TestApplyExecutorConfig_UpdatesJobsWhenChainRemoved(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupExecutorDatastore(t, ds, defaultSelectors, testDefaultQualifier)
	env.DataStore = ds.Seal()

	topology := newTestTopology()

	cs := changesets.ApplyExecutorConfig()
	firstOutput, err := cs.Apply(env, changesets.ApplyExecutorConfigCfg{
		Topology:          topology,
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    defaultSelectors,
	})
	require.NoError(t, err)

	firstJob, err := ccv.GetJob(firstOutput.DataStore.Seal(), shared.NOPAlias("nop-1"), shared.JobID("nop-1-default-executor"))
	require.NoError(t, err)

	sel1Str := strconv.FormatUint(defaultSelectors[0], 10)
	sel2Str := strconv.FormatUint(defaultSelectors[1], 10)
	assert.True(t, strings.Contains(firstJob.Spec, sel1Str) && strings.Contains(firstJob.Spec, sel2Str),
		"first job should contain both chains")

	env.DataStore = firstOutput.DataStore.Seal()
	secondOutput, err := cs.Apply(env, changesets.ApplyExecutorConfigCfg{
		Topology:          topology,
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    []uint64{defaultSelectors[0]},
	})
	require.NoError(t, err)

	secondJob, err := ccv.GetJob(secondOutput.DataStore.Seal(), shared.NOPAlias("nop-1"), shared.JobID("nop-1-default-executor"))
	require.NoError(t, err)

	assert.Contains(t, secondJob.Spec, sel1Str, "updated job should contain chain 1")
	assert.NotContains(t, secondJob.Spec, sel2Str, "updated job should NOT contain removed chain 2")
}

func TestApplyExecutorConfig_AllNOPsUpdatedWhenMemberAddedToPool(t *testing.T) {
	selectors := []uint64{chainsel.TEST_90000001.Selector}
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, selectors)
	setupExecutorDatastore(t, ds, selectors, testDefaultQualifier)
	env.DataStore = ds.Seal()

	initialTopology := newTestTopology(
		WithNOPs([]topology.NOPConfig{
			{Alias: "nop-1", Name: "nop-1", Mode: shared.NOPModeStandalone},
			{Alias: "nop-2", Name: "nop-2", Mode: shared.NOPModeStandalone},
		}),
		WithExecutorPool(testDefaultQualifier, topology.ExecutorPoolConfig{
			NOPAliases:        []string{"nop-1", "nop-2"},
			ExecutionInterval: 15 * time.Second,
		}),
	)

	cs := changesets.ApplyExecutorConfig()
	firstOutput, err := cs.Apply(env, changesets.ApplyExecutorConfigCfg{
		Topology:          initialTopology,
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    selectors,
	})
	require.NoError(t, err)

	firstSealedDS := firstOutput.DataStore.Seal()
	nop1Job, err := ccv.GetJob(firstSealedDS, shared.NOPAlias("nop-1"), shared.JobID("nop-1-default-executor"))
	require.NoError(t, err)
	nop2Job, err := ccv.GetJob(firstSealedDS, shared.NOPAlias("nop-2"), shared.JobID("nop-2-default-executor"))
	require.NoError(t, err)

	assert.Contains(t, nop1Job.Spec, "nop-1")
	assert.Contains(t, nop1Job.Spec, "nop-2")
	assert.NotContains(t, nop1Job.Spec, "nop-3")
	assert.Contains(t, nop2Job.Spec, "nop-1")
	assert.Contains(t, nop2Job.Spec, "nop-2")
	assert.NotContains(t, nop2Job.Spec, "nop-3")

	expandedTopology := newTestTopology(
		WithNOPs([]topology.NOPConfig{
			{Alias: "nop-1", Name: "nop-1", Mode: shared.NOPModeStandalone},
			{Alias: "nop-2", Name: "nop-2", Mode: shared.NOPModeStandalone},
			{Alias: "nop-3", Name: "nop-3", Mode: shared.NOPModeStandalone},
		}),
		WithExecutorPool(testDefaultQualifier, topology.ExecutorPoolConfig{
			NOPAliases:        []string{"nop-1", "nop-2", "nop-3"},
			ExecutionInterval: 15 * time.Second,
		}),
	)

	env.DataStore = firstSealedDS
	secondOutput, err := cs.Apply(env, changesets.ApplyExecutorConfigCfg{
		Topology:          expandedTopology,
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    selectors,
	})
	require.NoError(t, err)

	secondSealedDS := secondOutput.DataStore.Seal()
	nop1JobUpdated, err := ccv.GetJob(secondSealedDS, shared.NOPAlias("nop-1"), shared.JobID("nop-1-default-executor"))
	require.NoError(t, err)
	nop2JobUpdated, err := ccv.GetJob(secondSealedDS, shared.NOPAlias("nop-2"), shared.JobID("nop-2-default-executor"))
	require.NoError(t, err)
	nop3Job, err := ccv.GetJob(secondSealedDS, shared.NOPAlias("nop-3"), shared.JobID("nop-3-default-executor"))
	require.NoError(t, err)

	assert.Contains(t, nop1JobUpdated.Spec, "nop-3", "nop-1 job should include new member nop-3")
	assert.Contains(t, nop2JobUpdated.Spec, "nop-3", "nop-2 job should include new member nop-3")
	assert.Contains(t, nop3Job.Spec, "nop-1", "new nop-3 job should include existing member nop-1")
	assert.Contains(t, nop3Job.Spec, "nop-2", "new nop-3 job should include existing member nop-2")
}

func TestApplyExecutorConfig_AllNOPsUpdatedWhenMemberRemovedFromPool(t *testing.T) {
	selectors := []uint64{chainsel.TEST_90000001.Selector}
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, selectors)
	setupExecutorDatastore(t, ds, selectors, testDefaultQualifier)
	env.DataStore = ds.Seal()

	initialTopology := newTestTopology(
		WithNOPs([]topology.NOPConfig{
			{Alias: "nop-1", Name: "nop-1", Mode: shared.NOPModeStandalone},
			{Alias: "nop-2", Name: "nop-2", Mode: shared.NOPModeStandalone},
			{Alias: "nop-3", Name: "nop-3", Mode: shared.NOPModeStandalone},
		}),
		WithExecutorPool(testDefaultQualifier, topology.ExecutorPoolConfig{
			NOPAliases:        []string{"nop-1", "nop-2", "nop-3"},
			ExecutionInterval: 15 * time.Second,
		}),
	)

	cs := changesets.ApplyExecutorConfig()
	firstOutput, err := cs.Apply(env, changesets.ApplyExecutorConfigCfg{
		Topology:          initialTopology,
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    selectors,
	})
	require.NoError(t, err)

	firstSealedDS := firstOutput.DataStore.Seal()
	for _, nop := range []shared.NOPAlias{"nop-1", "nop-2", "nop-3"} {
		job, err := ccv.GetJob(firstSealedDS, nop, shared.JobID(string(nop)+"-default-executor"))
		require.NoError(t, err)
		assert.Contains(t, job.Spec, "nop-1")
		assert.Contains(t, job.Spec, "nop-2")
		assert.Contains(t, job.Spec, "nop-3")
	}

	reducedTopology := newTestTopology(
		WithNOPs([]topology.NOPConfig{
			{Alias: "nop-1", Name: "nop-1", Mode: shared.NOPModeStandalone},
			{Alias: "nop-2", Name: "nop-2", Mode: shared.NOPModeStandalone},
			{Alias: "nop-3", Name: "nop-3", Mode: shared.NOPModeStandalone},
		}),
		WithExecutorPool(testDefaultQualifier, topology.ExecutorPoolConfig{
			NOPAliases:        []string{"nop-1", "nop-2"},
			ExecutionInterval: 15 * time.Second,
		}),
	)

	env.DataStore = firstSealedDS
	secondOutput, err := cs.Apply(env, changesets.ApplyExecutorConfigCfg{
		Topology:          reducedTopology,
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    selectors,
	})
	require.NoError(t, err)

	secondSealedDS := secondOutput.DataStore.Seal()
	nop1JobUpdated, err := ccv.GetJob(secondSealedDS, shared.NOPAlias("nop-1"), shared.JobID("nop-1-default-executor"))
	require.NoError(t, err)
	nop2JobUpdated, err := ccv.GetJob(secondSealedDS, shared.NOPAlias("nop-2"), shared.JobID("nop-2-default-executor"))
	require.NoError(t, err)

	assert.Contains(t, nop1JobUpdated.Spec, "nop-1")
	assert.Contains(t, nop1JobUpdated.Spec, "nop-2")
	assert.NotContains(t, nop1JobUpdated.Spec, "nop-3", "nop-1 job should not include removed member nop-3")

	assert.Contains(t, nop2JobUpdated.Spec, "nop-1")
	assert.Contains(t, nop2JobUpdated.Spec, "nop-2")
	assert.NotContains(t, nop2JobUpdated.Spec, "nop-3", "nop-2 job should not include removed member nop-3")

	_, err = ccv.GetJob(secondSealedDS, shared.NOPAlias("nop-3"), shared.JobID("nop-3-default-executor"))
	require.Error(t, err, "nop-3 job should be removed when NOP is removed from pool")
}

func TestApplyExecutorConfig_FailsWhenNOPMissingChainSupport(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupExecutorDatastore(t, ds, defaultSelectors, testDefaultQualifier)
	env.DataStore = ds.Seal()

	mockJD := deploymocks.NewMockJDClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop-1"},
				{Id: "node-2", Name: "nop-2"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "90000001"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0x123"},
					},
				},
				{
					NodeId: "node-2",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "90000001"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0x456"},
					},
				},
				{
					NodeId: "node-2",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "90000002"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0x456"},
					},
				},
			},
		}, nil,
	)

	deps := changesets.ExecutorApplyDeps{
		Env:      env,
		JDClient: mockJD,
		NodeIDs:  []string{"node-1", "node-2"},
	}

	cfg := changesets.ApplyExecutorConfigCfg{
		Topology:          newTestTopology(),
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    defaultSelectors,
	}

	_, err := changesets.ApplyExecutorConfigWithDeps(deps, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "chain support validation failed")
	assert.Contains(t, err.Error(), "nop-1")
}

func TestApplyExecutorConfig_PassesWhenNOPSupportsExtraChains(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupExecutorDatastore(t, ds, defaultSelectors, testDefaultQualifier)
	env.DataStore = ds.Seal()

	mockJD := deploymocks.NewMockJDClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop-1"},
				{Id: "node-2", Name: "nop-2"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "90000001"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0x123"},
					},
				},
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "90000002"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0x123"},
					},
				},
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "90000003"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0x123"},
					},
				},
				{
					NodeId: "node-2",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "90000001"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0x456"},
					},
				},
				{
					NodeId: "node-2",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "90000002"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0x456"},
					},
				},
			},
		}, nil,
	)

	deps := changesets.ExecutorApplyDeps{
		Env:      env,
		JDClient: mockJD,
		NodeIDs:  []string{"node-1", "node-2"},
	}

	cfg := changesets.ApplyExecutorConfigCfg{
		Topology:          newTestTopology(),
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    defaultSelectors,
	}

	output, err := changesets.ApplyExecutorConfigWithDeps(deps, cfg)
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)
}

func TestApplyExecutorConfig_PassesWhenNonTargetNOPMissingChainSupport(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupExecutorDatastore(t, ds, defaultSelectors, testDefaultQualifier)
	env.DataStore = ds.Seal()

	mockJD := deploymocks.NewMockJDClient(t)
	mockJD.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{
			Nodes: []*nodev1.Node{
				{Id: "node-1", Name: "nop-1"},
			},
		}, nil,
	)
	mockJD.EXPECT().ListNodeChainConfigs(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodeChainConfigsResponse{
			ChainConfigs: []*nodev1.ChainConfig{
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "90000001"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0x123"},
					},
				},
				{
					NodeId: "node-1",
					Chain:  &nodev1.Chain{Type: nodev1.ChainType_CHAIN_TYPE_EVM, Id: "90000002"},
					Ocr2Config: &nodev1.OCR2Config{
						OcrKeyBundle: &nodev1.OCR2Config_OCRKeyBundle{OnchainSigningAddress: "0x123"},
					},
				},
			},
		}, nil,
	)

	deps := changesets.ExecutorApplyDeps{
		Env:      env,
		JDClient: mockJD,
		NodeIDs:  []string{"node-1", "node-2"},
	}

	cfg := changesets.ApplyExecutorConfigCfg{
		Topology:          newTestTopology(),
		ExecutorQualifier: testDefaultQualifier,
		ChainSelectors:    defaultSelectors,
		TargetNOPs:        []shared.NOPAlias{"nop-1"},
	}

	output, err := changesets.ApplyExecutorConfigWithDeps(deps, cfg)
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)
}
