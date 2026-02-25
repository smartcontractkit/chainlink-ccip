package sync_job_proposals

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/ccv"
	deploymocks "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
)

func newTestBundle(t *testing.T) operations.Bundle {
	t.Helper()
	lggr, err := logger.New()
	require.NoError(t, err)
	return operations.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		operations.NewMemoryReporter(),
	)
}

func newTestEnvironment(t *testing.T) deployment.Environment {
	t.Helper()
	lggr, err := logger.New()
	require.NoError(t, err)
	return deployment.Environment{
		GetContext: func() context.Context { return context.Background() },
		Logger:     lggr,
	}
}

func newTestEnvironmentWithDataStore(t *testing.T, ds datastore.DataStore, nodeIDs []string) deployment.Environment {
	t.Helper()
	lggr, err := logger.New()
	require.NoError(t, err)
	return deployment.Environment{
		GetContext: func() context.Context { return context.Background() },
		Logger:     lggr,
		DataStore:  ds,
		NodeIDs:    nodeIDs,
	}
}

func TestSyncJobProposals_NoJDClient_ReturnsError(t *testing.T) {
	bundle := newTestBundle(t)
	env := newTestEnvironment(t)

	_, err := operations.ExecuteOperation(bundle, SyncJobProposals, SyncJobProposalsDeps{
		Env: env,
	}, SyncJobProposalsInput{})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "offchain client not available")
}

func TestSyncJobProposals_EmptyJobs_ReturnsEmpty(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	bundle := newTestBundle(t)
	env := newTestEnvironment(t)

	report, err := operations.ExecuteOperation(bundle, SyncJobProposals, SyncJobProposalsDeps{
		Env:      env,
		JDClient: mockClient,
	}, SyncJobProposalsInput{})

	require.NoError(t, err)
	assert.Empty(t, report.Output.StatusChanges)
	assert.Empty(t, report.Output.SpecDrifts)
	assert.Empty(t, report.Output.OrphanedJobs)
	assert.Empty(t, report.Output.Errors)
}

func TestSyncJobProposals_StatusChange_UpdatesDatastore(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	localJob := shared.JobInfo{
		Spec:          "type = 'ccv'",
		ExternalJobID: "ext-1",
		JDJobID:       "jd-job-1",
		NodeID:        "node-1",
		JobID:         "nop-1-default-executor",
		NOPAlias:      "nop-1",
		Mode:          shared.NOPModeCL,
		Proposals: map[string]shared.ProposalRevision{
			"jd-proposal-1": {
				ProposalID: "jd-proposal-1",
				Revision:   1,
				Status:     shared.JobProposalStatusPending,
				Spec:       "type = 'ccv'",
			},
		},
	}
	err := ccv.SaveJob(ds, localJob)
	require.NoError(t, err)

	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListProposals(mock.Anything, mock.Anything).Return(
		&jobv1.ListProposalsResponse{
			Proposals: []*jobv1.Proposal{
				{
					Id:       "jd-proposal-1",
					JobId:    "jd-job-1",
					Revision: 1,
					Status:   jobv1.ProposalStatus_PROPOSAL_STATUS_APPROVED,
					Spec:     "type = 'ccv'",
				},
			},
		}, nil,
	)

	bundle := newTestBundle(t)
	env := newTestEnvironmentWithDataStore(t, ds.Seal(), []string{"node-1"})

	report, err := operations.ExecuteOperation(bundle, SyncJobProposals, SyncJobProposalsDeps{
		Env:      env,
		JDClient: mockClient,
	}, SyncJobProposalsInput{})

	require.NoError(t, err)
	require.Len(t, report.Output.StatusChanges, 1)
	assert.Equal(t, shared.NOPAlias("nop-1"), report.Output.StatusChanges[0].NOPAlias)
	assert.Equal(t, shared.JobProposalStatusPending, report.Output.StatusChanges[0].OldStatus)
	assert.Equal(t, shared.JobProposalStatusApproved, report.Output.StatusChanges[0].NewStatus)
	assert.Empty(t, report.Output.SpecDrifts)
	assert.Empty(t, report.Output.OrphanedJobs)
	assert.Empty(t, report.Output.Errors)

	updatedJob, err := ccv.GetJob(report.Output.DataStore.Seal(), "nop-1", "nop-1-default-executor")
	require.NoError(t, err)
	assert.Equal(t, shared.JobProposalStatusApproved, updatedJob.LatestStatus())
	assert.Equal(t, "jd-proposal-1", updatedJob.ActiveProposalID)
}

func TestSyncJobProposals_SpecDrift_DetectedAndReported(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	localJob := shared.JobInfo{
		Spec:          "type = 'ccv'\nname = 'local'",
		ExternalJobID: "ext-1",
		JDJobID:       "jd-job-1",
		NodeID:        "node-1",
		JobID:         "nop-1-default-executor",
		NOPAlias:      "nop-1",
		Mode:          shared.NOPModeCL,
		Proposals: map[string]shared.ProposalRevision{
			"jd-proposal-1": {
				ProposalID: "jd-proposal-1",
				Revision:   1,
				Status:     shared.JobProposalStatusPending,
				Spec:       "type = 'ccv'\nname = 'local'",
			},
		},
	}
	err := ccv.SaveJob(ds, localJob)
	require.NoError(t, err)

	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListProposals(mock.Anything, mock.Anything).Return(
		&jobv1.ListProposalsResponse{
			Proposals: []*jobv1.Proposal{
				{
					Id:       "jd-proposal-1",
					JobId:    "jd-job-1",
					Revision: 1,
					Status:   jobv1.ProposalStatus_PROPOSAL_STATUS_PENDING,
					Spec:     "type = 'ccv'\nname = 'different'",
				},
			},
		}, nil,
	)

	bundle := newTestBundle(t)
	env := newTestEnvironmentWithDataStore(t, ds.Seal(), []string{"node-1"})

	report, err := operations.ExecuteOperation(bundle, SyncJobProposals, SyncJobProposalsDeps{
		Env:      env,
		JDClient: mockClient,
	}, SyncJobProposalsInput{})

	require.NoError(t, err)
	assert.Empty(t, report.Output.StatusChanges)
	require.Len(t, report.Output.SpecDrifts, 1)
	assert.Equal(t, shared.NOPAlias("nop-1"), report.Output.SpecDrifts[0].NOPAlias)
	assert.Equal(t, "type = 'ccv'\nname = 'local'", report.Output.SpecDrifts[0].LocalSpec)
	assert.Equal(t, "type = 'ccv'\nname = 'different'", report.Output.SpecDrifts[0].JDSpec)
	assert.Empty(t, report.Output.OrphanedJobs)
	assert.Empty(t, report.Output.Errors)
}

func TestSyncJobProposals_JobNotFoundInJD_MarkedAsOrphanedAndDeleted(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	localJob := shared.JobInfo{
		Spec:          "type = 'ccv'",
		ExternalJobID: "ext-1",
		JDJobID:       "jd-job-deleted",
		NodeID:        "node-1",
		JobID:         "nop-1-default-executor",
		NOPAlias:      "nop-1",
		Mode:          shared.NOPModeCL,
		Proposals: map[string]shared.ProposalRevision{
			"old-proposal": {
				ProposalID: "old-proposal",
				Revision:   1,
				Status:     shared.JobProposalStatusPending,
			},
		},
	}
	err := ccv.SaveJob(ds, localJob)
	require.NoError(t, err)

	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListProposals(mock.Anything, mock.Anything).Return(
		&jobv1.ListProposalsResponse{}, nil,
	)

	bundle := newTestBundle(t)
	env := newTestEnvironmentWithDataStore(t, ds.Seal(), []string{"node-1"})

	report, err := operations.ExecuteOperation(bundle, SyncJobProposals, SyncJobProposalsDeps{
		Env:      env,
		JDClient: mockClient,
	}, SyncJobProposalsInput{})

	require.NoError(t, err)
	assert.Empty(t, report.Output.StatusChanges)
	assert.Empty(t, report.Output.SpecDrifts)
	require.Len(t, report.Output.OrphanedJobs, 1)
	assert.Equal(t, shared.NOPAlias("nop-1"), report.Output.OrphanedJobs[0].NOPAlias)
	assert.Equal(t, shared.JobID("nop-1-default-executor"), report.Output.OrphanedJobs[0].JobID)
	assert.Equal(t, "jd-job-deleted", report.Output.OrphanedJobs[0].JDJobID)
	assert.Empty(t, report.Output.Errors)

	_, err = ccv.GetJob(report.Output.DataStore.Seal(), "nop-1", "nop-1-default-executor")
	require.Error(t, err)
}

func TestSyncJobProposals_NOPFilter_OnlySyncsFilteredNOPs(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	job1 := shared.JobInfo{
		Spec:     "type = 'ccv'",
		JDJobID:  "jd-job-1",
		NodeID:   "node-1",
		JobID:    "nop-1-default-executor",
		NOPAlias: "nop-1",
		Mode:     shared.NOPModeCL,
		Proposals: map[string]shared.ProposalRevision{
			"prop-1": {ProposalID: "prop-1", Revision: 1, Status: shared.JobProposalStatusPending},
		},
	}
	job2 := shared.JobInfo{
		Spec:     "type = 'ccv'",
		JDJobID:  "jd-job-2",
		NodeID:   "node-2",
		JobID:    "nop-2-default-executor",
		NOPAlias: "nop-2",
		Mode:     shared.NOPModeCL,
		Proposals: map[string]shared.ProposalRevision{
			"prop-2": {ProposalID: "prop-2", Revision: 1, Status: shared.JobProposalStatusPending},
		},
	}
	require.NoError(t, ccv.SaveJob(ds, job1))
	require.NoError(t, ccv.SaveJob(ds, job2))

	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListProposals(mock.Anything, mock.Anything).Return(
		&jobv1.ListProposalsResponse{
			Proposals: []*jobv1.Proposal{
				{Id: "prop-1", JobId: "jd-job-1", Revision: 1, Status: jobv1.ProposalStatus_PROPOSAL_STATUS_APPROVED, Spec: "type = 'ccv'"},
			},
		}, nil,
	).Once()

	bundle := newTestBundle(t)
	env := newTestEnvironmentWithDataStore(t, ds.Seal(), []string{"node-1", "node-2"})

	report, err := operations.ExecuteOperation(bundle, SyncJobProposals, SyncJobProposalsDeps{
		Env:      env,
		JDClient: mockClient,
	}, SyncJobProposalsInput{
		NOPAliases: []shared.NOPAlias{"nop-1"},
	})

	require.NoError(t, err)
	require.Len(t, report.Output.StatusChanges, 1)
	assert.Equal(t, shared.NOPAlias("nop-1"), report.Output.StatusChanges[0].NOPAlias)
}

func TestSyncJobProposals_StandaloneModeJobs_Skipped(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	job := shared.JobInfo{
		Spec:     "type = 'ccv'",
		JDJobID:  "jd-job-1",
		JobID:    "nop-1-default-executor",
		NOPAlias: "nop-1",
		Mode:     shared.NOPModeStandalone,
	}
	require.NoError(t, ccv.SaveJob(ds, job))

	mockClient := deploymocks.NewMockJDClient(t)

	bundle := newTestBundle(t)
	env := newTestEnvironmentWithDataStore(t, ds.Seal(), []string{"node-1"})

	report, err := operations.ExecuteOperation(bundle, SyncJobProposals, SyncJobProposalsDeps{
		Env:      env,
		JDClient: mockClient,
	}, SyncJobProposalsInput{})

	require.NoError(t, err)
	assert.Empty(t, report.Output.StatusChanges)
	assert.Empty(t, report.Output.SpecDrifts)
	assert.Empty(t, report.Output.OrphanedJobs)
	assert.Empty(t, report.Output.Errors)
}

func TestSyncJobProposals_JobWithoutJDJobID_Skipped(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	job := shared.JobInfo{
		Spec:     "type = 'ccv'",
		JDJobID:  "",
		JobID:    "nop-1-default-executor",
		NOPAlias: "nop-1",
		Mode:     shared.NOPModeCL,
	}
	require.NoError(t, ccv.SaveJob(ds, job))

	mockClient := deploymocks.NewMockJDClient(t)

	bundle := newTestBundle(t)
	env := newTestEnvironmentWithDataStore(t, ds.Seal(), []string{"node-1"})

	report, err := operations.ExecuteOperation(bundle, SyncJobProposals, SyncJobProposalsDeps{
		Env:      env,
		JDClient: mockClient,
	}, SyncJobProposalsInput{})

	require.NoError(t, err)
	assert.Empty(t, report.Output.StatusChanges)
	assert.Empty(t, report.Output.SpecDrifts)
	assert.Empty(t, report.Output.OrphanedJobs)
	assert.Empty(t, report.Output.Errors)
}

func TestSyncJobProposals_NewRevision_AddsToProposalHistory(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	localJob := shared.JobInfo{
		Spec:     "type = 'ccv'",
		JDJobID:  "jd-job-1",
		NodeID:   "node-1",
		JobID:    "nop-1-default-executor",
		NOPAlias: "nop-1",
		Mode:     shared.NOPModeCL,
		Proposals: map[string]shared.ProposalRevision{
			"prop-1": {ProposalID: "prop-1", Revision: 1, Status: shared.JobProposalStatusRejected, Spec: "type = 'ccv'"},
		},
	}
	err := ccv.SaveJob(ds, localJob)
	require.NoError(t, err)

	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListProposals(mock.Anything, mock.Anything).Return(
		&jobv1.ListProposalsResponse{
			Proposals: []*jobv1.Proposal{
				{Id: "prop-1", JobId: "jd-job-1", Revision: 1, Status: jobv1.ProposalStatus_PROPOSAL_STATUS_REJECTED, Spec: "type = 'ccv'"},
				{Id: "prop-2", JobId: "jd-job-1", Revision: 2, Status: jobv1.ProposalStatus_PROPOSAL_STATUS_PENDING, Spec: "type = 'ccv' v2"},
			},
		}, nil,
	)

	bundle := newTestBundle(t)
	env := newTestEnvironmentWithDataStore(t, ds.Seal(), []string{"node-1"})

	report, err := operations.ExecuteOperation(bundle, SyncJobProposals, SyncJobProposalsDeps{
		Env:      env,
		JDClient: mockClient,
	}, SyncJobProposalsInput{})

	require.NoError(t, err)
	require.Len(t, report.Output.StatusChanges, 1)
	assert.Equal(t, shared.JobProposalStatusRejected, report.Output.StatusChanges[0].OldStatus)
	assert.Equal(t, shared.JobProposalStatusPending, report.Output.StatusChanges[0].NewStatus)

	updatedJob, err := ccv.GetJob(report.Output.DataStore.Seal(), "nop-1", "nop-1-default-executor")
	require.NoError(t, err)
	assert.Len(t, updatedJob.Proposals, 2)
	assert.Equal(t, int64(2), updatedJob.LatestProposal().Revision)
}

func TestMapJDStatusToLocal_AllStatusesMapped(t *testing.T) {
	tests := []struct {
		name     string
		jdStatus jobv1.ProposalStatus
		expected shared.JobProposalStatus
	}{
		{"pending", jobv1.ProposalStatus_PROPOSAL_STATUS_PENDING, shared.JobProposalStatusPending},
		{"proposed", jobv1.ProposalStatus_PROPOSAL_STATUS_PROPOSED, shared.JobProposalStatusPending},
		{"approved", jobv1.ProposalStatus_PROPOSAL_STATUS_APPROVED, shared.JobProposalStatusApproved},
		{"rejected", jobv1.ProposalStatus_PROPOSAL_STATUS_REJECTED, shared.JobProposalStatusRejected},
		{"revoked", jobv1.ProposalStatus_PROPOSAL_STATUS_REVOKED, shared.JobProposalStatusRevoked},
		{"cancelled", jobv1.ProposalStatus_PROPOSAL_STATUS_CANCELLED, shared.JobProposalStatusRevoked},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := mapJDStatusToLocal(tc.jdStatus)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestSyncJobProposals_EmptyNodeIDs_ReturnsError(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	job := shared.JobInfo{
		Spec:     "type = 'ccv'",
		JDJobID:  "jd-job-1",
		NodeID:   "node-1",
		JobID:    "nop-1-default-executor",
		NOPAlias: "nop-1",
		Mode:     shared.NOPModeCL,
	}
	require.NoError(t, ccv.SaveJob(ds, job))

	mockClient := deploymocks.NewMockJDClient(t)

	bundle := newTestBundle(t)
	env := newTestEnvironmentWithDataStore(t, ds.Seal(), nil)

	_, err := operations.ExecuteOperation(bundle, SyncJobProposals, SyncJobProposalsDeps{
		Env:      env,
		JDClient: mockClient,
	}, SyncJobProposalsInput{})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "NodeIDs must be specified")
}

func TestSyncJobProposals_JobNodeIDNotInAllowedList_ReturnsError(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	job := shared.JobInfo{
		Spec:     "type = 'ccv'",
		JDJobID:  "jd-job-1",
		NodeID:   "node-not-allowed",
		JobID:    "nop-1-default-executor",
		NOPAlias: "nop-1",
		Mode:     shared.NOPModeCL,
	}
	require.NoError(t, ccv.SaveJob(ds, job))

	mockClient := deploymocks.NewMockJDClient(t)

	bundle := newTestBundle(t)
	env := newTestEnvironmentWithDataStore(t, ds.Seal(), []string{"node-1", "node-2"})

	_, err := operations.ExecuteOperation(bundle, SyncJobProposals, SyncJobProposalsDeps{
		Env:      env,
		JDClient: mockClient,
	}, SyncJobProposalsInput{})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "not in the allowed node list")
}

func TestSyncJobProposals_JobMissingNodeID_ReturnsError(t *testing.T) {
	ds := datastore.NewMemoryDataStore()

	job := shared.JobInfo{
		Spec:     "type = 'ccv'",
		JDJobID:  "jd-job-1",
		NodeID:   "",
		JobID:    "nop-1-default-executor",
		NOPAlias: "nop-1",
		Mode:     shared.NOPModeCL,
	}
	require.NoError(t, ccv.SaveJob(ds, job))

	mockClient := deploymocks.NewMockJDClient(t)

	bundle := newTestBundle(t)
	env := newTestEnvironmentWithDataStore(t, ds.Seal(), []string{"node-1"})

	_, err := operations.ExecuteOperation(bundle, SyncJobProposals, SyncJobProposalsDeps{
		Env:      env,
		JDClient: mockClient,
	}, SyncJobProposalsInput{})

	require.Error(t, err)
	assert.Contains(t, err.Error(), "has no NodeID")
}
