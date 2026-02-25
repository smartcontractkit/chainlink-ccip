package changesets_test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	chainsel "github.com/smartcontractkit/chain-selectors"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/ccv"
	deploymocks "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/topology"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/testutils"
)

func TestApplyVerifierConfig_Validation(t *testing.T) {
	tests := []struct {
		name        string
		cfg         changesets.ApplyVerifierConfigCfg
		expectedErr string
	}{
		{
			name: "missing topology",
			cfg: changesets.ApplyVerifierConfigCfg{
				Topology:                 nil,
				CommitteeQualifier:       testCommittee,
				DefaultExecutorQualifier: testDefaultQualifier,
			},
			expectedErr: "topology is required",
		},
		{
			name: "missing default executor qualifier",
			cfg: changesets.ApplyVerifierConfigCfg{
				Topology:                 newTestTopology(),
				CommitteeQualifier:       testCommittee,
				DefaultExecutorQualifier: "",
			},
			expectedErr: "default executor qualifier is required",
		},
		{
			name: "missing committee qualifier",
			cfg: changesets.ApplyVerifierConfigCfg{
				Topology:                 newTestTopology(),
				CommitteeQualifier:       "",
				DefaultExecutorQualifier: testDefaultQualifier,
			},
			expectedErr: "committee qualifier is required",
		},
		{
			name: "committee not found in topology",
			cfg: changesets.ApplyVerifierConfigCfg{
				Topology:                 newTestTopology(),
				CommitteeQualifier:       "non-existent-committee",
				DefaultExecutorQualifier: testDefaultQualifier,
			},
			expectedErr: "committee \"non-existent-committee\" not found in topology",
		},
		{
			name: "NOP not found in topology",
			cfg: changesets.ApplyVerifierConfigCfg{
				Topology:                 newTestTopology(),
				CommitteeQualifier:       testCommittee,
				DefaultExecutorQualifier: testDefaultQualifier,
				TargetNOPs:               []shared.NOPAlias{"non-existent-nop"},
			},
			expectedErr: "NOP alias \"non-existent-nop\" not found in topology",
		},
		{
			name: "chain selector not in environment",
			cfg: changesets.ApplyVerifierConfigCfg{
				Topology:                 newTestTopology(),
				CommitteeQualifier:       testCommittee,
				DefaultExecutorQualifier: testDefaultQualifier,
				ChainSelectors:           []uint64{999999999},
			},
			expectedErr: "selector 999999999 is not available in environment",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			changeset := changesets.ApplyVerifierConfig()
			env := newTestEnvironment(t, defaultSelectors)

			err := changeset.VerifyPreconditions(env, tt.cfg)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}

func TestApplyVerifierConfig_ValidatesAggregators(t *testing.T) {
	changeset := changesets.ApplyVerifierConfig()
	env := newTestEnvironment(t, defaultSelectors)

	topology := newTestTopology(
		WithCommittee(testCommittee, topology.CommitteeConfig{
			Qualifier:   testCommittee,
			Aggregators: []topology.AggregatorConfig{},
			ChainConfigs: map[string]topology.ChainCommitteeConfig{
				strconv.FormatUint(chainsel.TEST_90000001.Selector, 10): {
					NOPAliases: []string{"nop-1", "nop-2"},
					Threshold:  2,
				},
			},
		}),
	)

	err := changeset.VerifyPreconditions(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topology,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "at least one aggregator is required")
}

func TestApplyVerifierConfig_ValidatesChainNotInCommittee(t *testing.T) {
	changeset := changesets.ApplyVerifierConfig()

	selectors := []uint64{
		chainsel.TEST_90000001.Selector,
		chainsel.TEST_90000002.Selector,
		chainsel.TEST_90000003.Selector,
	}
	env := newTestEnvironment(t, selectors)

	err := changeset.VerifyPreconditions(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 newTestTopology(),
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           []uint64{chainsel.TEST_90000003.Selector},
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "not configured in committee")
}

func TestApplyVerifierConfig_PyroscopeNotAllowedInProduction(t *testing.T) {
	changeset := changesets.ApplyVerifierConfig()

	env := newTestEnvironment(t, defaultSelectors)
	env.Name = "mainnet"

	topology := newTestTopology(WithPyroscopeURL("http://pyroscope:4040"))

	err := changeset.VerifyPreconditions(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topology,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
	})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "pyroscope URL is not supported for production environments")
}

func TestApplyVerifierConfig_GeneratesValidJobSpec(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)
	env.DataStore = ds.Seal()

	sel1Str := strconv.FormatUint(defaultSelectors[0], 10)
	sel2Str := strconv.FormatUint(defaultSelectors[1], 10)

	topology := newTestTopology(
		WithPyroscopeURL("http://pyroscope:4040"),
		WithMonitoring(defaultMonitoringConfig()),
	)

	cs := changesets.ApplyVerifierConfig()
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topology,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	job, err := ccv.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop-1"), shared.JobID("nop-1-instance-1-test-committee-verifier"))
	require.NoError(t, err)
	jobSpec := job.Spec

	assert.Contains(t, jobSpec, `schemaVersion = 1`)
	assert.Contains(t, jobSpec, `type = "ccvcommitteeverifier"`)
	assert.Contains(t, jobSpec, `committeeVerifierConfig = """`)
	assert.Contains(t, jobSpec, `verifier_id = "instance-1-test-committee-verifier"`)
	assert.Contains(t, jobSpec, `aggregator_address = "aggregator-1:443"`)
	assert.Contains(t, jobSpec, `signer_address = "0xABCDEF1234567890ABCDEF1234567890ABCDEF12"`)
	assert.Contains(t, jobSpec, `pyroscope_url = "http://pyroscope:4040"`)
	assert.Contains(t, jobSpec, `[committee_verifier_addresses]`)
	assert.Contains(t, jobSpec, sel1Str)
	assert.Contains(t, jobSpec, sel2Str)
	assert.Contains(t, jobSpec, `[on_ramp_addresses]`)
	assert.Contains(t, jobSpec, `[default_executor_on_ramp_addresses]`)
	assert.Contains(t, jobSpec, `[rmn_remote_addresses]`)
	assert.Contains(t, jobSpec, `[monitoring]`)
	assert.Contains(t, jobSpec, `Enabled = true`)
	assert.Contains(t, jobSpec, `Type = "beholder"`)
}

func TestApplyVerifierConfig_PreservesExistingJobSpecs(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)

	existingJobSpec := "existing-executor-job-spec"
	err := ccv.SaveJob(ds, shared.JobInfo{
		Spec:     existingJobSpec,
		JobID:    "existing-executor",
		NOPAlias: "existing-nop",
	})
	require.NoError(t, err)

	env.DataStore = ds.Seal()

	cs := changesets.ApplyVerifierConfig()
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 newTestTopology(),
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	outputSealed := output.DataStore.Seal()

	_, err = ccv.GetJob(outputSealed, shared.NOPAlias("nop-1"), shared.JobID("nop-1-instance-1-test-committee-verifier"))
	require.NoError(t, err, "new verifier job spec should be present")

	retrievedJob, err := ccv.GetJob(outputSealed, shared.NOPAlias("existing-nop"), shared.JobID("existing-executor"))
	require.NoError(t, err, "existing executor job spec should be preserved")
	assert.Equal(t, existingJobSpec, retrievedJob.Spec, "executor job spec should be unchanged")
}

func TestApplyVerifierConfig_MultipleAggregators(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)
	env.DataStore = ds.Seal()

	sel1Str := strconv.FormatUint(defaultSelectors[0], 10)
	sel2Str := strconv.FormatUint(defaultSelectors[1], 10)

	topology := newTestTopology(
		WithCommittee(testCommittee, topology.CommitteeConfig{
			Qualifier:       testCommittee,
			VerifierVersion: semver.MustParse("1.7.0"),
			Aggregators: []topology.AggregatorConfig{
				{Name: "agg-primary", Address: "aggregator-primary:443", InsecureAggregatorConnection: true},
				{Name: "agg-secondary", Address: "aggregator-secondary:443", InsecureAggregatorConnection: true},
				{Name: "agg-tertiary", Address: "aggregator-tertiary:443", InsecureAggregatorConnection: true},
			},
			ChainConfigs: map[string]topology.ChainCommitteeConfig{
				sel1Str: {NOPAliases: []string{"nop-1"}, Threshold: 1},
				sel2Str: {NOPAliases: []string{"nop-1"}, Threshold: 1},
			},
		}),
	)

	cs := changesets.ApplyVerifierConfig()
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topology,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	outputSealed := output.DataStore.Seal()

	job1, err := ccv.GetJob(outputSealed, shared.NOPAlias("nop-1"), shared.JobID("nop-1-agg-primary-test-committee-verifier"))
	require.NoError(t, err)
	assert.Contains(t, job1.Spec, `aggregator_address = "aggregator-primary:443"`)

	job2, err := ccv.GetJob(outputSealed, shared.NOPAlias("nop-1"), shared.JobID("nop-1-agg-secondary-test-committee-verifier"))
	require.NoError(t, err)
	assert.Contains(t, job2.Spec, `aggregator_address = "aggregator-secondary:443"`)

	job3, err := ccv.GetJob(outputSealed, shared.NOPAlias("nop-1"), shared.JobID("nop-1-agg-tertiary-test-committee-verifier"))
	require.NoError(t, err)
	assert.Contains(t, job3.Spec, `aggregator_address = "aggregator-tertiary:443"`)
}

func TestApplyVerifierConfig_RemovesOrphanedJobSpecs(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)

	err := ccv.SaveJob(ds, shared.JobInfo{
		Spec:     "orphaned-job-spec",
		JobID:    "removed-nop-instance-1-test-committee-verifier",
		NOPAlias: "removed-nop",
		Mode:     shared.NOPModeStandalone,
	})
	require.NoError(t, err)

	err = ccv.SaveJob(ds, shared.JobInfo{
		Spec:     "other-committee-job-spec",
		JobID:    "nop-1-instance-1-other-committee-verifier",
		NOPAlias: "nop-1",
		Mode:     shared.NOPModeStandalone,
	})
	require.NoError(t, err)

	env.DataStore = ds.Seal()

	cs := changesets.ApplyVerifierConfig()
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 newTestTopology(),
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	outputSealed := output.DataStore.Seal()

	_, err = ccv.GetJob(outputSealed, shared.NOPAlias("removed-nop"), shared.JobID("removed-nop-instance-1-test-committee-verifier"))
	require.Error(t, err, "orphaned job spec should be removed")

	otherCommitteeJob, err := ccv.GetJob(outputSealed, shared.NOPAlias("nop-1"), shared.JobID("nop-1-instance-1-other-committee-verifier"))
	require.NoError(t, err, "job spec for other committee should be preserved")
	assert.Equal(t, "other-committee-job-spec", otherCommitteeJob.Spec)
}

func TestApplyVerifierConfig_TargetNOPsScoping(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)

	err := ccv.SaveJob(ds, shared.JobInfo{
		Spec:     "nop-1-job-spec",
		JobID:    "nop-1-instance-1-test-committee-verifier",
		NOPAlias: "nop-1",
	})
	require.NoError(t, err)
	err = ccv.SaveJob(ds, shared.JobInfo{
		Spec:     "nop-2-job-spec",
		JobID:    "nop-2-instance-1-test-committee-verifier",
		NOPAlias: "nop-2",
	})
	require.NoError(t, err)

	env.DataStore = ds.Seal()

	cs := changesets.ApplyVerifierConfig()
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 newTestTopology(),
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
		TargetNOPs:               []shared.NOPAlias{"nop-1"},
	})
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)

	outputSealed := output.DataStore.Seal()

	job, err := ccv.GetJob(outputSealed, shared.NOPAlias("nop-1"), shared.JobID("nop-1-instance-1-test-committee-verifier"))
	require.NoError(t, err, "nop-1 verifier job spec should exist")
	assert.NotEqual(t, "nop-1-job-spec", job.Spec, "nop-1 job spec should be updated")

	nop2Job, err := ccv.GetJob(outputSealed, shared.NOPAlias("nop-2"), shared.JobID("nop-2-instance-1-test-committee-verifier"))
	require.NoError(t, err, "nop-2 verifier job spec should be preserved when not in scope")
	assert.Equal(t, "nop-2-job-spec", nop2Job.Spec, "nop-2 job spec should be unchanged")
}

func TestApplyVerifierConfig_UsesCommitteeChainsWhenEmptyChainSelectors(t *testing.T) {
	selectors := []uint64{
		chainsel.TEST_90000001.Selector,
		chainsel.TEST_90000002.Selector,
		chainsel.TEST_90000003.Selector,
	}
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, selectors)
	setupVerifierDatastore(t, ds, selectors[:2], testCommittee, testDefaultQualifier)
	env.DataStore = ds.Seal()

	sel1Str := strconv.FormatUint(selectors[0], 10)
	sel2Str := strconv.FormatUint(selectors[1], 10)
	sel3Str := strconv.FormatUint(selectors[2], 10)

	cs := changesets.ApplyVerifierConfig()
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 newTestTopology(),
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           nil,
	})
	require.NoError(t, err)

	job, err := ccv.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop-1"), shared.JobID("nop-1-instance-1-test-committee-verifier"))
	require.NoError(t, err)

	assert.Contains(t, job.Spec, sel1Str, "should include chain from committee config")
	assert.Contains(t, job.Spec, sel2Str, "should include chain from committee config")
	assert.NotContains(t, job.Spec, sel3Str, "should NOT include chain not in committee config")
}

func TestApplyVerifierConfig_FiltersChainsPerNOPMembership(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)
	env.DataStore = ds.Seal()

	sel1Str := strconv.FormatUint(defaultSelectors[0], 10)
	sel2Str := strconv.FormatUint(defaultSelectors[1], 10)

	topology := newTestTopology(
		WithCommittee(testCommittee, topology.CommitteeConfig{
			Qualifier:       testCommittee,
			VerifierVersion: semver.MustParse("1.7.0"),
			Aggregators: []topology.AggregatorConfig{
				{Name: "instance-1", Address: "aggregator-1:443", InsecureAggregatorConnection: true},
			},
			ChainConfigs: map[string]topology.ChainCommitteeConfig{
				sel1Str: {NOPAliases: []string{"nop-1"}, Threshold: 1},
				sel2Str: {NOPAliases: []string{"nop-2"}, Threshold: 1},
			},
		}),
	)

	cs := changesets.ApplyVerifierConfig()
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topology,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	})
	require.NoError(t, err)

	nop1Job, err := ccv.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop-1"), shared.JobID("nop-1-instance-1-test-committee-verifier"))
	require.NoError(t, err)

	assert.Contains(t, nop1Job.Spec, sel1Str, "nop-1 should have chain 1 where they are member")
	assert.NotContains(t, nop1Job.Spec, sel2Str, "nop-1 should NOT have chain 2 where they are NOT member")

	nop2Job, err := ccv.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop-2"), shared.JobID("nop-2-instance-1-test-committee-verifier"))
	require.NoError(t, err)

	assert.NotContains(t, nop2Job.Spec, sel1Str, "nop-2 should NOT have chain 1 where they are NOT member")
	assert.Contains(t, nop2Job.Spec, sel2Str, "nop-2 should have chain 2 where they are member")
}

func TestApplyVerifierConfig_SkipsNOPsNotInAnyChainConfig(t *testing.T) {
	selectors := []uint64{chainsel.TEST_90000001.Selector}
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, selectors)
	setupVerifierDatastore(t, ds, selectors, testCommittee, testDefaultQualifier)
	env.DataStore = ds.Seal()

	sel1Str := strconv.FormatUint(selectors[0], 10)

	topology := newTestTopology(
		WithNOPs([]topology.NOPConfig{
			{
				Alias:                 "nop-1",
				Name:                  "nop-1",
				SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0xABCDEF1234567890ABCDEF1234567890ABCDEF12"},
				Mode:                  shared.NOPModeStandalone,
			},
			{
				Alias:                 "nop-not-in-committee",
				Name:                  "nop-not-in-committee",
				SignerAddressByFamily: map[string]string{chainsel.FamilyEVM: "0x1234567890ABCDEF1234567890ABCDEF12345678"},
				Mode:                  shared.NOPModeStandalone,
			},
		}),
		WithCommittee(testCommittee, topology.CommitteeConfig{
			Qualifier:       testCommittee,
			VerifierVersion: semver.MustParse("1.7.0"),
			Aggregators: []topology.AggregatorConfig{
				{Name: "instance-1", Address: "aggregator-1:443", InsecureAggregatorConnection: true},
			},
			ChainConfigs: map[string]topology.ChainCommitteeConfig{
				sel1Str: {NOPAliases: []string{"nop-1"}, Threshold: 1},
			},
		}),
		WithExecutorPool(testDefaultQualifier, topology.ExecutorPoolConfig{
			NOPAliases: []string{"nop-1", "nop-not-in-committee"},
		}),
	)

	cs := changesets.ApplyVerifierConfig()
	output, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topology,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           selectors,
		TargetNOPs:               []shared.NOPAlias{"nop-1", "nop-not-in-committee"},
	})
	require.NoError(t, err)

	_, err = ccv.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop-1"), shared.JobID("nop-1-instance-1-test-committee-verifier"))
	require.NoError(t, err, "nop-1 should have a job since they are in the committee")

	_, err = ccv.GetJob(output.DataStore.Seal(), shared.NOPAlias("nop-not-in-committee"), shared.JobID("nop-not-in-committee-instance-1-test-committee-verifier"))
	require.Error(t, err, "nop-not-in-committee should NOT have a job since they are not in any chain config")
}

func TestApplyVerifierConfig_SkipsUnchangedJobSpecs(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)
	env.DataStore = ds.Seal()

	topology := newTestTopology()

	cs := changesets.ApplyVerifierConfig()
	firstOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topology,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	})
	require.NoError(t, err)

	firstJob, err := ccv.GetJob(firstOutput.DataStore.Seal(), "nop-1", "nop-1-instance-1-test-committee-verifier")
	require.NoError(t, err)
	firstSpec := firstJob.Spec

	env.DataStore = firstOutput.DataStore.Seal()
	secondOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topology,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	})
	require.NoError(t, err)

	secondJob, err := ccv.GetJob(secondOutput.DataStore.Seal(), "nop-1", "nop-1-instance-1-test-committee-verifier")
	require.NoError(t, err)

	assert.Equal(t, firstSpec, secondJob.Spec, "job spec should remain unchanged on re-run")
}

func TestApplyVerifierConfig_UpdatesJobsWhenChainRemoved(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)
	env.DataStore = ds.Seal()

	envTopology := newTestTopology()

	cs := changesets.ApplyVerifierConfig()
	firstOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 envTopology,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	})
	require.NoError(t, err)

	firstJob, err := ccv.GetJob(firstOutput.DataStore.Seal(), "nop-1", "nop-1-instance-1-test-committee-verifier")
	require.NoError(t, err)

	sel1Str := strconv.FormatUint(defaultSelectors[0], 10)
	sel2Str := strconv.FormatUint(defaultSelectors[1], 10)
	assert.True(t, strings.Contains(firstJob.Spec, sel1Str) && strings.Contains(firstJob.Spec, sel2Str),
		"first job should contain both chains")

	topologyWithOneChain := newTestTopology(
		WithCommittee(testCommittee, topology.CommitteeConfig{
			Qualifier:       testCommittee,
			VerifierVersion: semver.MustParse("1.7.0"),
			Aggregators: []topology.AggregatorConfig{
				{Name: "instance-1", Address: "aggregator-1:443", InsecureAggregatorConnection: true},
			},
			ChainConfigs: map[string]topology.ChainCommitteeConfig{
				sel1Str: {NOPAliases: []string{"nop-1", "nop-2"}, Threshold: 2},
			},
		}),
	)

	env.DataStore = firstOutput.DataStore.Seal()
	secondOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topologyWithOneChain,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           []uint64{defaultSelectors[0]},
	})
	require.NoError(t, err)

	secondJob, err := ccv.GetJob(secondOutput.DataStore.Seal(), "nop-1", "nop-1-instance-1-test-committee-verifier")
	require.NoError(t, err)

	assert.Contains(t, secondJob.Spec, sel1Str, "updated job should contain chain 1")
	assert.NotContains(t, secondJob.Spec, sel2Str, "updated job should NOT contain removed chain 2")
}

func TestApplyVerifierConfig_RemovesJobWhenNOPRemovedFromCommittee(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)
	env.DataStore = ds.Seal()

	sel1Str := strconv.FormatUint(defaultSelectors[0], 10)
	sel2Str := strconv.FormatUint(defaultSelectors[1], 10)

	topologyWithBothNOPs := newTestTopology()

	cs := changesets.ApplyVerifierConfig()
	firstOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topologyWithBothNOPs,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	})
	require.NoError(t, err)

	_, err = ccv.GetJob(firstOutput.DataStore.Seal(), "nop-1", "nop-1-instance-1-test-committee-verifier")
	require.NoError(t, err, "nop-1 should have verifier job initially")
	_, err = ccv.GetJob(firstOutput.DataStore.Seal(), "nop-2", "nop-2-instance-1-test-committee-verifier")
	require.NoError(t, err, "nop-2 should have verifier job initially")

	topologyWithNOP1Only := newTestTopology(
		WithCommittee(testCommittee, topology.CommitteeConfig{
			Qualifier:       testCommittee,
			VerifierVersion: semver.MustParse("1.7.0"),
			Aggregators: []topology.AggregatorConfig{
				{Name: testAggregatorName, Address: testAggregatorAddress, InsecureAggregatorConnection: true},
			},
			ChainConfigs: map[string]topology.ChainCommitteeConfig{
				sel1Str: {NOPAliases: []string{"nop-1"}, Threshold: 1},
				sel2Str: {NOPAliases: []string{"nop-1"}, Threshold: 1},
			},
		}),
	)

	env.DataStore = firstOutput.DataStore.Seal()
	secondOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topologyWithNOP1Only,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	})
	require.NoError(t, err)

	_, err = ccv.GetJob(secondOutput.DataStore.Seal(), "nop-1", "nop-1-instance-1-test-committee-verifier")
	require.NoError(t, err, "nop-1 should still have verifier job")

	_, err = ccv.GetJob(secondOutput.DataStore.Seal(), "nop-2", "nop-2-instance-1-test-committee-verifier")
	require.Error(t, err, "nop-2 verifier job should be removed after being removed from committee")
}

func TestApplyVerifierConfig_CreatesJobWhenNOPAddedToCommittee(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)
	env.DataStore = ds.Seal()

	sel1Str := strconv.FormatUint(defaultSelectors[0], 10)
	sel2Str := strconv.FormatUint(defaultSelectors[1], 10)

	topologyWithNOP1Only := newTestTopology(
		WithCommittee(testCommittee, topology.CommitteeConfig{
			Qualifier:       testCommittee,
			VerifierVersion: semver.MustParse("1.7.0"),
			Aggregators: []topology.AggregatorConfig{
				{Name: testAggregatorName, Address: testAggregatorAddress, InsecureAggregatorConnection: true},
			},
			ChainConfigs: map[string]topology.ChainCommitteeConfig{
				sel1Str: {NOPAliases: []string{"nop-1"}, Threshold: 1},
				sel2Str: {NOPAliases: []string{"nop-1"}, Threshold: 1},
			},
		}),
	)

	cs := changesets.ApplyVerifierConfig()
	firstOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topologyWithNOP1Only,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	})
	require.NoError(t, err)

	_, err = ccv.GetJob(firstOutput.DataStore.Seal(), "nop-1", "nop-1-instance-1-test-committee-verifier")
	require.NoError(t, err, "nop-1 should have verifier job")
	_, err = ccv.GetJob(firstOutput.DataStore.Seal(), "nop-2", "nop-2-instance-1-test-committee-verifier")
	require.Error(t, err, "nop-2 should NOT have verifier job initially")

	topologyWithBothNOPs := newTestTopology()

	env.DataStore = firstOutput.DataStore.Seal()
	secondOutput, err := cs.Apply(env, changesets.ApplyVerifierConfigCfg{
		Topology:                 topologyWithBothNOPs,
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	})
	require.NoError(t, err)

	_, err = ccv.GetJob(secondOutput.DataStore.Seal(), "nop-1", "nop-1-instance-1-test-committee-verifier")
	require.NoError(t, err, "nop-1 should still have verifier job")

	nop2Job, err := ccv.GetJob(secondOutput.DataStore.Seal(), "nop-2", "nop-2-instance-1-test-committee-verifier")
	require.NoError(t, err, "nop-2 should now have verifier job after being added to committee")
	assert.Contains(t, nop2Job.Spec, sel1Str, "nop-2 job should include chain 1")
	assert.Contains(t, nop2Job.Spec, sel2Str, "nop-2 job should include chain 2")
}

func TestApplyVerifierConfig_FailsWhenNOPMissingChainSupport(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)
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

	deps := changesets.VerifierApplyDeps{
		Env:      env,
		JDClient: mockJD,
		NodeIDs:  []string{"node-1", "node-2"},
	}

	cfg := changesets.ApplyVerifierConfigCfg{
		Topology:                 newTestTopology(),
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	}

	_, err := changesets.ApplyVerifierConfigWithDeps(deps, cfg)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "chain support validation failed")
	assert.Contains(t, err.Error(), "nop-1")
}

func TestApplyVerifierConfig_PassesWhenNOPSupportsExtraChains(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)
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

	deps := changesets.VerifierApplyDeps{
		Env:      env,
		JDClient: mockJD,
		NodeIDs:  []string{"node-1", "node-2"},
	}

	cfg := changesets.ApplyVerifierConfigCfg{
		Topology:                 newTestTopology(),
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
	}

	output, err := changesets.ApplyVerifierConfigWithDeps(deps, cfg)
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)
}

func TestApplyVerifierConfig_PassesWhenNonTargetNOPMissingChainSupport(t *testing.T) {
	env, ds := testutils.NewSimulatedEVMEnvironmentWithDataStore(t, defaultSelectors)
	setupVerifierDatastore(t, ds, defaultSelectors, testCommittee, testDefaultQualifier)
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

	deps := changesets.VerifierApplyDeps{
		Env:      env,
		JDClient: mockJD,
		NodeIDs:  []string{"node-1", "node-2"},
	}

	cfg := changesets.ApplyVerifierConfigCfg{
		Topology:                 newTestTopology(),
		CommitteeQualifier:       testCommittee,
		DefaultExecutorQualifier: testDefaultQualifier,
		ChainSelectors:           defaultSelectors,
		TargetNOPs:               []shared.NOPAlias{"nop-1"},
	}

	output, err := changesets.ApplyVerifierConfigWithDeps(deps, cfg)
	require.NoError(t, err)
	require.NotNil(t, output.DataStore)
}
