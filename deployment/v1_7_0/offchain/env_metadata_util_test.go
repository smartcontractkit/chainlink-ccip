package offchain

import (
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
