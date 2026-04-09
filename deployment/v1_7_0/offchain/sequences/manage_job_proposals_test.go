package sequences_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	csav1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/csa"
	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain"
	ccvmocks "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/offchain/shared"
)

func newMJPBundle(t *testing.T) cldf_ops.Bundle {
	t.Helper()
	lggr := logger.Test(t)
	return cldf_ops.NewBundle(
		func() context.Context { return context.Background() },
		lggr,
		cldf_ops.NewMemoryReporter(),
	)
}

// offchainAdapter adapts a shared.JDClient (the subset used by CCV operations) to
// the broader offchain.Client interface required by deployment.Environment.
// Methods beyond shared.JDClient panic if called, surfacing unexpected interactions.
type offchainAdapter struct{ shared.JDClient }

func (a *offchainAdapter) GetJob(ctx context.Context, in *jobv1.GetJobRequest, opts ...grpc.CallOption) (*jobv1.GetJobResponse, error) {
	panic("GetJob not expected in test")
}
func (a *offchainAdapter) GetProposal(ctx context.Context, in *jobv1.GetProposalRequest, opts ...grpc.CallOption) (*jobv1.GetProposalResponse, error) {
	panic("GetProposal not expected in test")
}
func (a *offchainAdapter) ListJobs(ctx context.Context, in *jobv1.ListJobsRequest, opts ...grpc.CallOption) (*jobv1.ListJobsResponse, error) {
	panic("ListJobs not expected in test")
}
func (a *offchainAdapter) BatchProposeJob(ctx context.Context, in *jobv1.BatchProposeJobRequest, opts ...grpc.CallOption) (*jobv1.BatchProposeJobResponse, error) {
	panic("BatchProposeJob not expected in test")
}
func (a *offchainAdapter) DeleteJob(ctx context.Context, in *jobv1.DeleteJobRequest, opts ...grpc.CallOption) (*jobv1.DeleteJobResponse, error) {
	panic("DeleteJob not expected in test")
}
func (a *offchainAdapter) UpdateJob(ctx context.Context, in *jobv1.UpdateJobRequest, opts ...grpc.CallOption) (*jobv1.UpdateJobResponse, error) {
	panic("UpdateJob not expected in test")
}
func (a *offchainAdapter) DisableNode(ctx context.Context, in *nodev1.DisableNodeRequest, opts ...grpc.CallOption) (*nodev1.DisableNodeResponse, error) {
	panic("DisableNode not expected in test")
}
func (a *offchainAdapter) EnableNode(ctx context.Context, in *nodev1.EnableNodeRequest, opts ...grpc.CallOption) (*nodev1.EnableNodeResponse, error) {
	panic("EnableNode not expected in test")
}
func (a *offchainAdapter) GetNode(ctx context.Context, in *nodev1.GetNodeRequest, opts ...grpc.CallOption) (*nodev1.GetNodeResponse, error) {
	panic("GetNode not expected in test")
}
func (a *offchainAdapter) RegisterNode(ctx context.Context, in *nodev1.RegisterNodeRequest, opts ...grpc.CallOption) (*nodev1.RegisterNodeResponse, error) {
	panic("RegisterNode not expected in test")
}
func (a *offchainAdapter) UpdateNode(ctx context.Context, in *nodev1.UpdateNodeRequest, opts ...grpc.CallOption) (*nodev1.UpdateNodeResponse, error) {
	panic("UpdateNode not expected in test")
}
func (a *offchainAdapter) GetKeypair(ctx context.Context, in *csav1.GetKeypairRequest, opts ...grpc.CallOption) (*csav1.GetKeypairResponse, error) {
	panic("GetKeypair not expected in test")
}
func (a *offchainAdapter) ListKeypairs(ctx context.Context, in *csav1.ListKeypairsRequest, opts ...grpc.CallOption) (*csav1.ListKeypairsResponse, error) {
	panic("ListKeypairs not expected in test")
}

func newMJPEnv(t *testing.T, client shared.JDClient, ds datastore.DataStore) deployment.Environment {
	t.Helper()
	lggr := logger.Test(t)
	env := deployment.Environment{
		Logger:    lggr,
		DataStore: ds,
		NodeIDs:   []string{"node-1"},
	}
	if client != nil {
		env.Offchain = &offchainAdapter{client}
	}
	return env
}

func emptyDS() datastore.DataStore {
	return datastore.NewMemoryDataStore().Seal()
}

// jobID builds a consistent JobID for use in tests.
func jobID(nop shared.NOPAlias, qualifier string) shared.JobID {
	return shared.JobID(string(nop) + "-" + qualifier)
}

// TestManageJobProposals_StandaloneNOPs verifies that when all NOPs are in
// standalone mode (no CL mode), the JD client is never called and the jobs are
// persisted to the output DataStore.
func TestManageJobProposals_StandaloneNOPs(t *testing.T) {
	nop := shared.NOPAlias("nop-1")
	jid := jobID(nop, "executor-pool1")

	input := sequences.ManageJobProposalsInput{
		JobSpecs: shared.NOPJobSpecs{
			nop: {jid: "type = 'standalone'\n"},
		},
		NOPs: sequences.NOPContext{
			Modes:      map[shared.NOPAlias]shared.NOPMode{nop: shared.NOPModeStandalone},
			TargetNOPs: []shared.NOPAlias{nop},
			AllNOPs:    []shared.NOPAlias{nop},
		},
	}
	deps := sequences.ManageJobProposalsDeps{
		Env: newMJPEnv(t, nil, emptyDS()), // nil Offchain is fine for standalone
	}

	report, err := cldf_ops.ExecuteSequence(newMJPBundle(t), sequences.ManageJobProposals, deps, input)
	require.NoError(t, err)
	assert.Len(t, report.Output.Jobs, 1)
	assert.Equal(t, shared.NOPModeStandalone, report.Output.Jobs[0].Mode)

	savedJobs, err := offchain.GetAllJobs(report.Output.DataStore.Seal())
	require.NoError(t, err)
	assert.Len(t, savedJobs[nop], 1)
}

// TestManageJobProposals_CLModeWithNilOffchain verifies that CL-mode NOPs
// require an Offchain client and return an error when one is absent.
func TestManageJobProposals_CLModeWithNilOffchain(t *testing.T) {
	nop := shared.NOPAlias("nop-1")
	jid := jobID(nop, "executor-pool1")

	input := sequences.ManageJobProposalsInput{
		JobSpecs: shared.NOPJobSpecs{
			nop: {jid: "type = 'ccvexecutor'\n"},
		},
		NOPs: sequences.NOPContext{
			Modes:      map[shared.NOPAlias]shared.NOPMode{nop: shared.NOPModeCL},
			TargetNOPs: []shared.NOPAlias{nop},
			AllNOPs:    []shared.NOPAlias{nop},
		},
	}
	deps := sequences.ManageJobProposalsDeps{
		Env: newMJPEnv(t, nil, emptyDS()),
	}

	_, err := cldf_ops.ExecuteSequence(newMJPBundle(t), sequences.ManageJobProposals, deps, input)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "e.Offchain is nil")
}

// TestManageJobProposals_CLModeUnchangedSpecSkipsJD verifies that a CL-mode
// job whose spec has not changed since it was last proposed is not re-proposed,
// i.e. JD ProposeJob is not called.
func TestManageJobProposals_CLModeUnchangedSpecSkipsJD(t *testing.T) {
	nop := shared.NOPAlias("nop-1")
	jid := jobID(nop, "executor-pool1")
	spec := "type = 'ccvexecutor'\n"

	// Pre-populate the datastore with the existing job that already has the same spec.
	mutableDS := datastore.NewMemoryDataStore()
	err := offchain.SaveJobs(mutableDS, []shared.JobInfo{
		{
			JobID:         jid,
			ExternalJobID: jid.ToExternalJobID(),
			NOPAlias:      nop,
			Mode:          shared.NOPModeCL,
			Spec:          spec,
			Proposals: map[string]shared.ProposalRevision{
				"prop-1": {
					ProposalID: "prop-1",
					Revision:   1,
					Status:     shared.JobProposalStatusApproved,
					Spec:       spec, // same spec → should be filtered out as unchanged
				},
			},
		},
	})
	require.NoError(t, err)

	mockClient := ccvmocks.NewMockJDClient(t)
	// ProposeJob must NOT be called — no expectations set.
	// mockClient will fail the test if any unexpected call is made.

	input := sequences.ManageJobProposalsInput{
		JobSpecs: shared.NOPJobSpecs{
			nop: {jid: spec},
		},
		NOPs: sequences.NOPContext{
			Modes:      map[shared.NOPAlias]shared.NOPMode{nop: shared.NOPModeCL},
			TargetNOPs: []shared.NOPAlias{nop},
			AllNOPs:    []shared.NOPAlias{nop},
		},
	}
	deps := sequences.ManageJobProposalsDeps{
		Env: newMJPEnv(t, mockClient, mutableDS.Seal()),
	}

	_, err = cldf_ops.ExecuteSequence(newMJPBundle(t), sequences.ManageJobProposals, deps, input)
	require.NoError(t, err)
}

// TestManageJobProposals_CLModeChangedSpecProposesViaJD verifies that a CL-mode
// job with a changed spec IS re-proposed through JD.
func TestManageJobProposals_CLModeChangedSpecProposesViaJD(t *testing.T) {
	nop := shared.NOPAlias("nop-1")
	jid := jobID(nop, "executor-pool1")
	oldSpec := "type = 'ccvexecutor'\nversion = '1'\n"
	newSpec := "type = 'ccvexecutor'\nversion = '2'\n"

	mutableDS := datastore.NewMemoryDataStore()
	err := offchain.SaveJobs(mutableDS, []shared.JobInfo{
		{
			JobID:         jid,
			ExternalJobID: jid.ToExternalJobID(),
			NOPAlias:      nop,
			Mode:          shared.NOPModeCL,
			Spec:          oldSpec,
			Proposals: map[string]shared.ProposalRevision{
				"prop-1": {ProposalID: "prop-1", Revision: 1, Spec: oldSpec},
			},
		},
	})
	require.NoError(t, err)

	mockClient := ccvmocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{Nodes: []*nodev1.Node{{Id: "node-1", Name: string(nop)}}}, nil,
	)
	mockClient.EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		&jobv1.ProposeJobResponse{
			Proposal: &jobv1.Proposal{Id: "prop-2", JobId: "jd-job-1", Spec: newSpec,
				Status: jobv1.ProposalStatus_PROPOSAL_STATUS_PENDING, Revision: 2},
		}, nil,
	)

	input := sequences.ManageJobProposalsInput{
		JobSpecs: shared.NOPJobSpecs{
			nop: {jid: newSpec},
		},
		NOPs: sequences.NOPContext{
			Modes:      map[shared.NOPAlias]shared.NOPMode{nop: shared.NOPModeCL},
			TargetNOPs: []shared.NOPAlias{nop},
			AllNOPs:    []shared.NOPAlias{nop},
		},
	}
	deps := sequences.ManageJobProposalsDeps{
		Env: newMJPEnv(t, mockClient, mutableDS.Seal()),
	}

	report, err := cldf_ops.ExecuteSequence(newMJPBundle(t), sequences.ManageJobProposals, deps, input)
	require.NoError(t, err)
	require.Len(t, report.Output.Jobs, 1)
	assert.Equal(t, newSpec, report.Output.Jobs[0].Spec)
}

// TestManageJobProposals_RevokeOrphanedJobsFalseSkipsRevoke verifies that
// setting RevokeOrphanedJobs=false never triggers a JD revoke call, even when
// there are stale jobs in the datastore.
func TestManageJobProposals_RevokeOrphanedJobsFalseSkipsRevoke(t *testing.T) {
	nop := shared.NOPAlias("nop-1")
	orphanJID := jobID(nop, "old-pool")

	mutableDS := datastore.NewMemoryDataStore()
	err := offchain.SaveJobs(mutableDS, []shared.JobInfo{
		{
			JobID:    orphanJID,
			NOPAlias: nop,
			Mode:     shared.NOPModeCL,
			Spec:     "type = 'ccvexecutor'\n",
		},
	})
	require.NoError(t, err)

	mockClient := ccvmocks.NewMockJDClient(t)
	// No RevokeJob expectation; any unexpected call will fail the test.

	input := sequences.ManageJobProposalsInput{
		JobSpecs:           shared.NOPJobSpecs{}, // no new jobs for nop-1
		RevokeOrphanedJobs: false,
		NOPs: sequences.NOPContext{
			Modes:      map[shared.NOPAlias]shared.NOPMode{nop: shared.NOPModeCL},
			TargetNOPs: []shared.NOPAlias{nop},
			AllNOPs:    []shared.NOPAlias{nop},
		},
	}
	deps := sequences.ManageJobProposalsDeps{
		Env: newMJPEnv(t, mockClient, mutableDS.Seal()),
	}

	_, err = cldf_ops.ExecuteSequence(newMJPBundle(t), sequences.ManageJobProposals, deps, input)
	require.NoError(t, err)
}

// TestManageJobProposals_ModeChangeStandaloneToCLExpectJDProposal verifies that
// when a NOP transitions from standalone to CL mode with the same spec,
// the job IS detected as changed and proposed to JD.
func TestManageJobProposals_ModeChangeStandaloneToCLExpectJDProposal(t *testing.T) {
	nop := shared.NOPAlias("nop-1")
	jid := jobID(nop, "executor-pool1")
	spec := "type = 'ccvexecutor'\n"

	mutableDS := datastore.NewMemoryDataStore()
	err := offchain.SaveJobs(mutableDS, []shared.JobInfo{
		{
			JobID:         jid,
			ExternalJobID: jid.ToExternalJobID(),
			NOPAlias:      nop,
			Mode:          shared.NOPModeStandalone,
			Spec:          spec,
			Proposals: map[string]shared.ProposalRevision{
				"prop-1": {
					ProposalID: "prop-1",
					Revision:   1,
					Status:     shared.JobProposalStatusApproved,
					Spec:       spec,
				},
			},
		},
	})
	require.NoError(t, err)

	mockClient := ccvmocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{Nodes: []*nodev1.Node{{Id: "node-1", Name: string(nop)}}}, nil,
	)
	mockClient.EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		&jobv1.ProposeJobResponse{
			Proposal: &jobv1.Proposal{Id: "prop-2", JobId: "jd-job-1", Spec: spec,
				Status: jobv1.ProposalStatus_PROPOSAL_STATUS_PENDING, Revision: 2},
		}, nil,
	)

	input := sequences.ManageJobProposalsInput{
		JobSpecs: shared.NOPJobSpecs{
			nop: {jid: spec},
		},
		NOPs: sequences.NOPContext{
			Modes:      map[shared.NOPAlias]shared.NOPMode{nop: shared.NOPModeCL},
			TargetNOPs: []shared.NOPAlias{nop},
			AllNOPs:    []shared.NOPAlias{nop},
		},
	}
	deps := sequences.ManageJobProposalsDeps{
		Env: newMJPEnv(t, mockClient, mutableDS.Seal()),
	}

	report, err := cldf_ops.ExecuteSequence(newMJPBundle(t), sequences.ManageJobProposals, deps, input)
	require.NoError(t, err)
	require.Len(t, report.Output.Jobs, 1)
	assert.Equal(t, shared.NOPModeCL, report.Output.Jobs[0].Mode)
	assert.Equal(t, "jd-job-1", report.Output.Jobs[0].JDJobID)
}

// TestManageJobProposals_ModeChangeCLToStandaloneExpectModeUpdatedInDataStore verifies that
// when a NOP transitions from CL to standalone mode with the same spec,
// the job is detected as changed and saved with the new mode.
func TestManageJobProposals_ModeChangeCLToStandaloneExpectModeUpdatedInDataStore(t *testing.T) {
	nop := shared.NOPAlias("nop-1")
	jid := jobID(nop, "executor-pool1")
	spec := "type = 'ccvexecutor'\n"

	mutableDS := datastore.NewMemoryDataStore()
	err := offchain.SaveJobs(mutableDS, []shared.JobInfo{
		{
			JobID:         jid,
			ExternalJobID: jid.ToExternalJobID(),
			NOPAlias:      nop,
			Mode:          shared.NOPModeCL,
			JDJobID:       "jd-job-1",
			NodeID:        "node-1",
			Spec:          spec,
			Proposals: map[string]shared.ProposalRevision{
				"prop-1": {
					ProposalID: "prop-1",
					Revision:   1,
					Status:     shared.JobProposalStatusApproved,
					Spec:       spec,
				},
			},
		},
	})
	require.NoError(t, err)

	input := sequences.ManageJobProposalsInput{
		JobSpecs: shared.NOPJobSpecs{
			nop: {jid: spec},
		},
		NOPs: sequences.NOPContext{
			Modes:      map[shared.NOPAlias]shared.NOPMode{nop: shared.NOPModeStandalone},
			TargetNOPs: []shared.NOPAlias{nop},
			AllNOPs:    []shared.NOPAlias{nop},
		},
	}
	deps := sequences.ManageJobProposalsDeps{
		Env: newMJPEnv(t, nil, mutableDS.Seal()),
	}

	report, err := cldf_ops.ExecuteSequence(newMJPBundle(t), sequences.ManageJobProposals, deps, input)
	require.NoError(t, err)
	require.Len(t, report.Output.Jobs, 1)
	assert.Equal(t, shared.NOPModeStandalone, report.Output.Jobs[0].Mode)

	savedJobs, err := offchain.GetAllJobs(report.Output.DataStore.Seal())
	require.NoError(t, err)
	savedJob := savedJobs[nop][jid]
	assert.Equal(t, shared.NOPModeStandalone, savedJob.Mode)
}

func TestManageJobProposals_CLToStandaloneTransition_WithRevokeFlag_RevokesOldCLJobAndClearsMetadata(t *testing.T) {
	nop := shared.NOPAlias("nop-1")
	jid := jobID(nop, "executor-pool1")
	spec := "type = 'ccvexecutor'\n"

	mutableDS := datastore.NewMemoryDataStore()
	err := offchain.SaveJobs(mutableDS, []shared.JobInfo{
		{
			JobID:            jid,
			ExternalJobID:    jid.ToExternalJobID(),
			NOPAlias:         nop,
			Mode:             shared.NOPModeCL,
			JDJobID:          "jd-job-1",
			NodeID:           "node-1",
			ActiveProposalID: "prop-1",
			Spec:             spec,
			Proposals: map[string]shared.ProposalRevision{
				"prop-1": {
					ProposalID: "prop-1",
					Revision:   1,
					Status:     shared.JobProposalStatusApproved,
					Spec:       spec,
				},
			},
		},
	})
	require.NoError(t, err)

	mockClient := ccvmocks.NewMockJDClient(t)
	mockClient.EXPECT().RevokeJob(mock.Anything, mock.Anything).Return(
		&jobv1.RevokeJobResponse{}, nil,
	)

	input := sequences.ManageJobProposalsInput{
		JobSpecs: shared.NOPJobSpecs{
			nop: {jid: spec},
		},
		RevokeOrphanedJobs: true,
		AffectedScope:      shared.ExecutorJobScope{ExecutorQualifier: "pool1"},
		NOPs: sequences.NOPContext{
			Modes:      map[shared.NOPAlias]shared.NOPMode{nop: shared.NOPModeStandalone},
			TargetNOPs: []shared.NOPAlias{nop},
			AllNOPs:    []shared.NOPAlias{nop},
		},
	}
	deps := sequences.ManageJobProposalsDeps{
		Env: newMJPEnv(t, mockClient, mutableDS.Seal()),
	}

	report, err := cldf_ops.ExecuteSequence(newMJPBundle(t), sequences.ManageJobProposals, deps, input)
	require.NoError(t, err)
	require.Len(t, report.Output.Jobs, 1)

	savedJobs, err := offchain.GetAllJobs(report.Output.DataStore.Seal())
	require.NoError(t, err)
	savedJob := savedJobs[nop][jid]
	assert.Equal(t, shared.NOPModeStandalone, savedJob.Mode)
	assert.Empty(t, savedJob.JDJobID, "JD metadata must be cleared after CL-to-standalone transition")
	assert.Empty(t, savedJob.NodeID, "JD metadata must be cleared after CL-to-standalone transition")
	assert.Empty(t, savedJob.ActiveProposalID, "JD metadata must be cleared after CL-to-standalone transition")
	assert.Nil(t, savedJob.Proposals, "JD metadata must be cleared after CL-to-standalone transition")
}

func TestManageJobProposals_CLToStandaloneTransition_WithoutRevokeFlag_ClearsMetadataButSkipsJDRevoke(t *testing.T) {
	nop := shared.NOPAlias("nop-1")
	jid := jobID(nop, "executor-pool1")
	spec := "type = 'ccvexecutor'\n"

	mutableDS := datastore.NewMemoryDataStore()
	err := offchain.SaveJobs(mutableDS, []shared.JobInfo{
		{
			JobID:            jid,
			ExternalJobID:    jid.ToExternalJobID(),
			NOPAlias:         nop,
			Mode:             shared.NOPModeCL,
			JDJobID:          "jd-job-1",
			NodeID:           "node-1",
			ActiveProposalID: "prop-1",
			Spec:             spec,
			Proposals: map[string]shared.ProposalRevision{
				"prop-1": {
					ProposalID: "prop-1",
					Revision:   1,
					Status:     shared.JobProposalStatusApproved,
					Spec:       spec,
				},
			},
		},
	})
	require.NoError(t, err)

	mockClient := ccvmocks.NewMockJDClient(t)
	// RevokeJob must NOT be called — mockClient will fail if any unexpected call is made.

	input := sequences.ManageJobProposalsInput{
		JobSpecs: shared.NOPJobSpecs{
			nop: {jid: spec},
		},
		RevokeOrphanedJobs: false,
		AffectedScope:      shared.ExecutorJobScope{ExecutorQualifier: "pool1"},
		NOPs: sequences.NOPContext{
			Modes:      map[shared.NOPAlias]shared.NOPMode{nop: shared.NOPModeStandalone},
			TargetNOPs: []shared.NOPAlias{nop},
			AllNOPs:    []shared.NOPAlias{nop},
		},
	}
	deps := sequences.ManageJobProposalsDeps{
		Env: newMJPEnv(t, mockClient, mutableDS.Seal()),
	}

	report, err := cldf_ops.ExecuteSequence(newMJPBundle(t), sequences.ManageJobProposals, deps, input)
	require.NoError(t, err)
	require.Len(t, report.Output.Jobs, 1)

	savedJobs, err := offchain.GetAllJobs(report.Output.DataStore.Seal())
	require.NoError(t, err)
	savedJob := savedJobs[nop][jid]
	assert.Equal(t, shared.NOPModeStandalone, savedJob.Mode)
	assert.Empty(t, savedJob.JDJobID, "JD metadata must be cleared even without revoke flag")
	assert.Empty(t, savedJob.NodeID, "JD metadata must be cleared even without revoke flag")
	assert.Empty(t, savedJob.ActiveProposalID, "JD metadata must be cleared even without revoke flag")
	assert.Nil(t, savedJob.Proposals, "JD metadata must be cleared even without revoke flag")
}

func TestManageJobProposals_RevokedOrphanCleanedFromDatastoreWithoutJDCall(t *testing.T) {
	nop := shared.NOPAlias("nop-1")
	orphanJID := shared.NewExecutorJobID(nop, shared.ExecutorJobScope{ExecutorQualifier: "old-pool"}).ToJobID()

	mutableDS := datastore.NewMemoryDataStore()
	err := offchain.SaveJobs(mutableDS, []shared.JobInfo{
		{
			JobID:    orphanJID,
			JDJobID:  "jd-job-orphan",
			NodeID:   "node-1",
			NOPAlias: nop,
			Mode:     shared.NOPModeCL,
			Spec:     "type = 'ccvexecutor'\n",
			Proposals: map[string]shared.ProposalRevision{
				"p1": {ProposalID: "p1", Status: shared.JobProposalStatusRevoked},
			},
		},
	})
	require.NoError(t, err)

	mockClient := ccvmocks.NewMockJDClient(t)
	// RevokeJob must NOT be called for an already-revoked job

	input := sequences.ManageJobProposalsInput{
		JobSpecs:           shared.NOPJobSpecs{},
		RevokeOrphanedJobs: true,
		AffectedScope:      shared.ExecutorJobScope{ExecutorQualifier: "old-pool"},
		NOPs: sequences.NOPContext{
			Modes:      map[shared.NOPAlias]shared.NOPMode{nop: shared.NOPModeCL},
			TargetNOPs: []shared.NOPAlias{nop},
			AllNOPs:    []shared.NOPAlias{nop},
		},
	}
	deps := sequences.ManageJobProposalsDeps{
		Env: newMJPEnv(t, mockClient, mutableDS.Seal()),
	}

	report, err := cldf_ops.ExecuteSequence(newMJPBundle(t), sequences.ManageJobProposals, deps, input)
	require.NoError(t, err)
	assert.Len(t, report.Output.RevokedJobs, 1, "revoked orphan should be in cleanup list")

	allJobs, err := offchain.GetAllJobs(report.Output.DataStore.Seal())
	require.NoError(t, err)
	assert.Empty(t, allJobs[nop], "orphaned job should be deleted from datastore")
}

// TestManageJobProposals_UnchangedCLJobExpectJDMetadataPreserved verifies that
// when a CL-mode job is re-submitted with the same spec (no change),
// the existing JD metadata (jdJobId, nodeId, proposals) is preserved
// in the output DataStore rather than being wiped.
func TestManageJobProposals_UnchangedCLJobExpectJDMetadataPreserved(t *testing.T) {
	nop := shared.NOPAlias("nop-1")
	jid := jobID(nop, "executor-pool1")
	spec := "type = 'ccvexecutor'\n"

	mutableDS := datastore.NewMemoryDataStore()
	err := offchain.SaveJobs(mutableDS, []shared.JobInfo{
		{
			JobID:         jid,
			ExternalJobID: jid.ToExternalJobID(),
			NOPAlias:      nop,
			Mode:          shared.NOPModeCL,
			JDJobID:       "jd-job-1",
			NodeID:        "node-1",
			Spec:          spec,
			Proposals: map[string]shared.ProposalRevision{
				"prop-1": {
					ProposalID: "prop-1",
					Revision:   1,
					Status:     shared.JobProposalStatusApproved,
					Spec:       spec,
				},
			},
		},
	})
	require.NoError(t, err)

	mockClient := ccvmocks.NewMockJDClient(t)

	input := sequences.ManageJobProposalsInput{
		JobSpecs: shared.NOPJobSpecs{
			nop: {jid: spec},
		},
		NOPs: sequences.NOPContext{
			Modes:      map[shared.NOPAlias]shared.NOPMode{nop: shared.NOPModeCL},
			TargetNOPs: []shared.NOPAlias{nop},
			AllNOPs:    []shared.NOPAlias{nop},
		},
	}
	deps := sequences.ManageJobProposalsDeps{
		Env: newMJPEnv(t, mockClient, mutableDS.Seal()),
	}

	report, err := cldf_ops.ExecuteSequence(newMJPBundle(t), sequences.ManageJobProposals, deps, input)
	require.NoError(t, err)

	savedJobs, err := offchain.GetAllJobs(report.Output.DataStore.Seal())
	require.NoError(t, err)

	savedJob := savedJobs[nop][jid]
	assert.Equal(t, "jd-job-1", savedJob.JDJobID, "jdJobId must be preserved for unchanged jobs")
	assert.Equal(t, "node-1", savedJob.NodeID, "nodeId must be preserved for unchanged jobs")
	require.Len(t, savedJob.Proposals, 1, "proposals must be preserved for unchanged jobs")
	assert.Equal(t, "prop-1", savedJob.Proposals["prop-1"].ProposalID)
}

// TestManageJobProposals_RevokeOrphanedJobsTrueRevokesViaCL verifies that
// CL-mode orphaned jobs are revoked through JD when RevokeOrphanedJobs=true.
func TestManageJobProposals_RevokeOrphanedJobsTrueRevokesViaCL(t *testing.T) {
	nop := shared.NOPAlias("nop-1")
	// ExecutorJobScope.IsJobInScope expects suffix "-<qualifier>-executor".
	orphanJID := shared.NewExecutorJobID(nop, shared.ExecutorJobScope{ExecutorQualifier: "old-pool"}).ToJobID()

	mutableDS := datastore.NewMemoryDataStore()
	err := offchain.SaveJobs(mutableDS, []shared.JobInfo{
		{
			JobID:    orphanJID,
			JDJobID:  "jd-job-orphan",
			NodeID:   "node-1",
			NOPAlias: nop,
			Mode:     shared.NOPModeCL,
			Spec:     "type = 'ccvexecutor'\n",
		},
	})
	require.NoError(t, err)

	mockClient := ccvmocks.NewMockJDClient(t)
	mockClient.EXPECT().RevokeJob(mock.Anything, mock.Anything).Return(
		&jobv1.RevokeJobResponse{}, nil,
	)

	input := sequences.ManageJobProposalsInput{
		JobSpecs:           shared.NOPJobSpecs{}, // no expected jobs for nop-1 → orphan detected
		RevokeOrphanedJobs: true,
		AffectedScope:      shared.ExecutorJobScope{ExecutorQualifier: "old-pool"},
		NOPs: sequences.NOPContext{
			Modes:      map[shared.NOPAlias]shared.NOPMode{nop: shared.NOPModeCL},
			TargetNOPs: []shared.NOPAlias{nop},
			AllNOPs:    []shared.NOPAlias{nop},
		},
	}
	deps := sequences.ManageJobProposalsDeps{
		Env: newMJPEnv(t, mockClient, mutableDS.Seal()),
	}

	report, err := cldf_ops.ExecuteSequence(newMJPBundle(t), sequences.ManageJobProposals, deps, input)
	require.NoError(t, err)
	assert.Len(t, report.Output.RevokedJobs, 1)
	assert.Equal(t, orphanJID, report.Output.RevokedJobs[0].JobID)
}
