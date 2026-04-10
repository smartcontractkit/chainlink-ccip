package offchain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/shared"
)

func TestCollectOrphanedJobs_TargetNOPsPreservesOutOfScopeNOPs(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	nop1 := shared.NOPAlias("nop1")
	nop2 := shared.NOPAlias("nop2")
	nop3 := shared.NOPAlias("nop3")
	scope := shared.ExecutorJobScope{ExecutorQualifier: "pool1"}

	jobNop1 := shared.JobInfo{
		JobID:    shared.JobID("nop1-pool1-executor"),
		NOPAlias: nop1,
		Proposals: map[string]shared.ProposalRevision{
			"p1": {Status: shared.JobProposalStatusApproved},
		},
	}
	jobNop2 := shared.JobInfo{
		JobID:    shared.JobID("nop2-pool1-executor"),
		NOPAlias: nop2,
		Proposals: map[string]shared.ProposalRevision{
			"p2": {Status: shared.JobProposalStatusApproved},
		},
	}
	jobNop3 := shared.JobInfo{
		JobID:    shared.JobID("nop3-pool1-executor"),
		NOPAlias: nop3,
		Proposals: map[string]shared.ProposalRevision{
			"p3": {Status: shared.JobProposalStatusApproved},
		},
	}
	require.NoError(t, SaveJobs(ds, []shared.JobInfo{jobNop1, jobNop2, jobNop3}))

	scopedNOPs := map[shared.NOPAlias]bool{nop1: true}
	expectedByNOP := map[shared.NOPAlias]map[shared.JobID]bool{
		nop1: {jobNop1.JobID: true},
	}
	environmentNOPs := map[shared.NOPAlias]bool{nop1: true, nop2: true, nop3: true}

	orphaned, err := CollectOrphanedJobs(ds.Seal(), scope, expectedByNOP, scopedNOPs, environmentNOPs)
	require.NoError(t, err)
	assert.Empty(t, orphaned, "nop1's job is expected so not orphaned; nop2/nop3 are out of TargetNOPs so must not appear in orphaned list")
}

func TestCollectOrphanedJobs_ScopeIsolatesCrossTypeJobs(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	nop := shared.NOPAlias("nop1")
	verifierScope := shared.VerifierJobScope{CommitteeQualifier: "c1"}

	verifierJob := shared.JobInfo{
		JobID:    shared.JobID("nop1-agg-c1-verifier"),
		NOPAlias: nop,
		Proposals: map[string]shared.ProposalRevision{
			"p1": {Status: shared.JobProposalStatusApproved},
		},
	}
	executorJob := shared.JobInfo{
		JobID:    shared.JobID("nop1-pool1-executor"),
		NOPAlias: nop,
		Proposals: map[string]shared.ProposalRevision{
			"p2": {Status: shared.JobProposalStatusApproved},
		},
	}
	require.NoError(t, SaveJobs(ds, []shared.JobInfo{verifierJob, executorJob}))

	expectedByNOP := map[shared.NOPAlias]map[shared.JobID]bool{
		nop: {verifierJob.JobID: true},
	}

	orphaned, err := CollectOrphanedJobs(ds.Seal(), verifierScope, expectedByNOP, nil, map[shared.NOPAlias]bool{nop: true})
	require.NoError(t, err)
	assert.Empty(t, orphaned, "verifier job is expected; executor job must not be in scope so must not appear in orphaned list")
}

func TestCollectOrphanedJobs_EmptyTargetNOPsConsidersAllNOPs(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	nop1 := shared.NOPAlias("nop1")
	nop2 := shared.NOPAlias("nop2")
	scope := shared.ExecutorJobScope{ExecutorQualifier: "pool1"}

	job1 := shared.JobInfo{
		JobID:    shared.JobID("nop1-pool1-executor"),
		NOPAlias: nop1,
		Proposals: map[string]shared.ProposalRevision{
			"p1": {Status: shared.JobProposalStatusApproved},
		},
	}
	job2 := shared.JobInfo{
		JobID:    shared.JobID("nop2-pool1-executor"),
		NOPAlias: nop2,
		Proposals: map[string]shared.ProposalRevision{
			"p2": {Status: shared.JobProposalStatusApproved},
		},
	}
	require.NoError(t, SaveJobs(ds, []shared.JobInfo{job1, job2}))

	expectedByNOP := map[shared.NOPAlias]map[shared.JobID]bool{
		nop1: {job1.JobID: true},
	}
	var scopedNOPs map[shared.NOPAlias]bool
	environmentNOPs := map[shared.NOPAlias]bool{nop1: true, nop2: true}

	orphaned, err := CollectOrphanedJobs(ds.Seal(), scope, expectedByNOP, scopedNOPs, environmentNOPs)
	require.NoError(t, err)
	assert.Len(t, orphaned, 1, "with empty TargetNOPs (scopedNOPs=nil) all NOPs are considered; nop2's job is not expected so must be orphaned")
	assert.Equal(t, nop2, orphaned[0].NOPAlias)
	assert.Equal(t, job2.JobID, orphaned[0].JobID)
}

func TestSaveAggregatorConfig_ReducedChainSet_RemovesStaleChains(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	svc := "agg-1"

	fullCfg := &Committee{
		QuorumConfigs: map[string]*QuorumConfig{
			"chain1": {SourceVerifierAddress: "0x1", Threshold: 1},
			"chain2": {SourceVerifierAddress: "0x2", Threshold: 2},
		},
	}
	require.NoError(t, SaveAggregatorConfig(ds, svc, fullCfg))

	reducedCfg := &Committee{
		QuorumConfigs: map[string]*QuorumConfig{
			"chain1": {SourceVerifierAddress: "0x1", Threshold: 1},
		},
	}
	require.NoError(t, SaveAggregatorConfig(ds, svc, reducedCfg))

	loaded, err := GetAggregatorConfig(ds.Seal(), svc)
	require.NoError(t, err)
	assert.Len(t, loaded.QuorumConfigs, 1)
	assert.Contains(t, loaded.QuorumConfigs, "chain1")
	assert.NotContains(t, loaded.QuorumConfigs, "chain2")
}

func TestDeleteJob_PreservesUnknownOffchainConfigKeys(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	require.NoError(t, SaveJob(ds, shared.JobInfo{
		JobID:    shared.JobID("nop1-pool1-executor"),
		NOPAlias: shared.NOPAlias("nop1"),
		Spec:     "spec-1",
	}))

	rawMeta := map[string]any{
		"offchainConfigs": map[string]any{
			"nopJobs": map[string]any{
				"nop1": map[string]any{
					"nop1-pool1-executor": map[string]any{
						"jobId":    "nop1-pool1-executor",
						"nopAlias": "nop1",
						"spec":     "spec-1",
					},
				},
			},
			"customExtension": map[string]any{
				"key": "value",
			},
		},
	}
	require.NoError(t, ds.EnvMetadata().Set(datastore.EnvMetadata{Metadata: rawMeta}))

	require.NoError(t, DeleteJob(ds, shared.NOPAlias("nop1"), shared.JobID("nop1-pool1-executor")))

	envMeta, err := ds.Seal().EnvMetadata().Get()
	require.NoError(t, err)

	data, err := json.Marshal(envMeta.Metadata)
	require.NoError(t, err)
	var result map[string]any
	require.NoError(t, json.Unmarshal(data, &result))

	oc, ok := result["offchainConfigs"].(map[string]any)
	require.True(t, ok)
	assert.Contains(t, oc, "customExtension")
}

func TestSaveTokenVerifierConfig_ReducedChainSet_RemovesStaleEntries(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	svc := "tv-1"

	fullCfg := &TokenVerifierGeneratedConfig{
		OnRampAddresses: map[string]string{
			"chain1": "0x1",
			"chain2": "0x2",
		},
	}
	require.NoError(t, SaveTokenVerifierConfig(ds, svc, fullCfg))

	reducedCfg := &TokenVerifierGeneratedConfig{
		OnRampAddresses: map[string]string{
			"chain1": "0x1",
		},
	}
	require.NoError(t, SaveTokenVerifierConfig(ds, svc, reducedCfg))

	loaded, err := GetTokenVerifierConfig(ds.Seal(), svc)
	require.NoError(t, err)
	assert.Len(t, loaded.OnRampAddresses, 1)
	assert.Contains(t, loaded.OnRampAddresses, "chain1")
	assert.NotContains(t, loaded.OnRampAddresses, "chain2")
}

func TestCollectOrphanedJobs_IncludesAlreadyRevokedJobsForCleanup(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	nop := shared.NOPAlias("nop1")
	scope := shared.ExecutorJobScope{ExecutorQualifier: "pool1"}

	revokedJob := shared.JobInfo{
		JobID:    shared.JobID("nop1-pool1-executor"),
		NOPAlias: nop,
		Proposals: map[string]shared.ProposalRevision{
			"p1": {Status: shared.JobProposalStatusRevoked},
		},
	}
	require.NoError(t, SaveJobs(ds, []shared.JobInfo{revokedJob}))

	expectedByNOP := map[shared.NOPAlias]map[shared.JobID]bool{}
	environmentNOPs := map[shared.NOPAlias]bool{nop: true}

	orphaned, err := CollectOrphanedJobs(ds.Seal(), scope, expectedByNOP, nil, environmentNOPs)
	require.NoError(t, err)
	assert.Len(t, orphaned, 1)
	assert.Equal(t, revokedJob.JobID, orphaned[0].JobID)
}
