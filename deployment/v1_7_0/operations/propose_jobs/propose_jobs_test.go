package propose_jobs

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

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

func newTestLogger(t *testing.T) logger.Logger {
	t.Helper()
	lggr, err := logger.New()
	require.NoError(t, err)
	return lggr
}

func TestProposeJobs_EmptyJobSpecs_ReturnsEmptySuccess(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	bundle := newTestBundle(t)

	input := ProposeJobsInput{
		JobSpecs: []JobSpecInput{},
	}

	report, err := operations.ExecuteOperation(bundle, ProposeJobs, ProposeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)
	require.NoError(t, err)
	assert.Empty(t, report.Output.ProposedJobs)
}

func TestProposeJobs_SingleJob_Success(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop-1"}}}, nil,
	)
	mockClient.EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		&jobv1.ProposeJobResponse{
			Proposal: &jobv1.Proposal{
				Id:     "proposal-1",
				JobId:  "jd-job-1",
				Spec:   "type = 'executor'",
				Status: jobv1.ProposalStatus_PROPOSAL_STATUS_PENDING,
			},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := ProposeJobsInput{
		JobSpecs: []JobSpecInput{
			{NOPAlias: "nop-1", InternalJobID: "nop-1-default-executor", Spec: "type = 'executor'"},
		},
	}

	report, err := operations.ExecuteOperation(bundle, ProposeJobs, ProposeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)
	require.NoError(t, err)
	require.Len(t, report.Output.ProposedJobs, 1)
	assert.Equal(t, shared.NOPAlias("nop-1"), report.Output.ProposedJobs[0].NOPAlias)
	assert.Equal(t, shared.JobID("nop-1-default-executor"), report.Output.ProposedJobs[0].InternalJobID)
	assert.Equal(t, "proposal-1", report.Output.ProposedJobs[0].ProposalID)
}

func TestProposeJobs_MultipleJobsSameNOP_AllProposed(t *testing.T) {
	proposalCount := 0
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop-1"}}}, nil,
	)
	mockClient.EXPECT().ProposeJob(mock.Anything, mock.Anything).RunAndReturn(
		func(ctx context.Context, req *jobv1.ProposeJobRequest, opts ...grpc.CallOption) (*jobv1.ProposeJobResponse, error) {
			proposalCount++
			return &jobv1.ProposeJobResponse{
				Proposal: &jobv1.Proposal{
					Id:     fmt.Sprintf("proposal-%d", proposalCount),
					JobId:  fmt.Sprintf("jd-job-%d", proposalCount),
					Spec:   req.Spec,
					Status: jobv1.ProposalStatus_PROPOSAL_STATUS_PENDING,
				},
			}, nil
		},
	).Times(2)

	bundle := newTestBundle(t)

	input := ProposeJobsInput{
		JobSpecs: []JobSpecInput{
			{NOPAlias: "nop-1", InternalJobID: "nop-1-default-executor", Spec: "type = 'executor'"},
			{NOPAlias: "nop-1", InternalJobID: "aggregator-1-default-verifier", Spec: "type = 'verifier'"},
		},
	}

	report, err := operations.ExecuteOperation(bundle, ProposeJobs, ProposeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)
	require.NoError(t, err)
	require.Len(t, report.Output.ProposedJobs, 2)

	assert.Equal(t, shared.JobID("nop-1-default-executor"), report.Output.ProposedJobs[0].InternalJobID)
	assert.Equal(t, "proposal-1", report.Output.ProposedJobs[0].ProposalID)
	assert.Equal(t, "type = 'executor'", report.Output.ProposedJobs[0].Spec)

	assert.Equal(t, shared.JobID("aggregator-1-default-verifier"), report.Output.ProposedJobs[1].InternalJobID)
	assert.Equal(t, "proposal-2", report.Output.ProposedJobs[1].ProposalID)
	assert.Equal(t, "type = 'verifier'", report.Output.ProposedJobs[1].Spec)
}

func TestProposeJobs_MultipleNOPs_AllProposed(t *testing.T) {
	proposalCount := 0
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{Nodes: []*nodev1.Node{
			{Id: "node-1", Name: "nop-1"},
			{Id: "node-2", Name: "nop-2"},
		}}, nil,
	)
	mockClient.EXPECT().ProposeJob(mock.Anything, mock.Anything).RunAndReturn(
		func(ctx context.Context, req *jobv1.ProposeJobRequest, opts ...grpc.CallOption) (*jobv1.ProposeJobResponse, error) {
			proposalCount++
			return &jobv1.ProposeJobResponse{
				Proposal: &jobv1.Proposal{
					Id:     fmt.Sprintf("proposal-%d", proposalCount),
					JobId:  fmt.Sprintf("jd-job-%d", proposalCount),
					Spec:   req.Spec,
					Status: jobv1.ProposalStatus_PROPOSAL_STATUS_PENDING,
				},
			}, nil
		},
	).Times(2)

	bundle := newTestBundle(t)

	input := ProposeJobsInput{
		JobSpecs: []JobSpecInput{
			{NOPAlias: "nop-1", InternalJobID: "nop-1-default-executor", Spec: "type = 'executor' nop = 1"},
			{NOPAlias: "nop-2", InternalJobID: "nop-2-default-executor", Spec: "type = 'executor' nop = 2"},
		},
	}

	report, err := operations.ExecuteOperation(bundle, ProposeJobs, ProposeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1", "node-2"},
	}, input)
	require.NoError(t, err)
	require.Len(t, report.Output.ProposedJobs, 2)

	nopAliases := make([]shared.NOPAlias, 0, 2)
	for _, pj := range report.Output.ProposedJobs {
		nopAliases = append(nopAliases, pj.NOPAlias)
	}
	assert.Contains(t, nopAliases, shared.NOPAlias("nop-1"))
	assert.Contains(t, nopAliases, shared.NOPAlias("nop-2"))
}

func TestProposeJobs_NodeNotFound_ReturnsError(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{Nodes: []*nodev1.Node{{Id: "node-1", Name: "existing-nop"}}}, nil,
	)

	bundle := newTestBundle(t)

	input := ProposeJobsInput{
		JobSpecs: []JobSpecInput{
			{NOPAlias: "non-existent-nop", InternalJobID: "job-1", Spec: "type = 'ccv'"},
		},
	}

	_, err := operations.ExecuteOperation(bundle, ProposeJobs, ProposeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "node not found")
}

func TestProposeJobs_ProposeError_ReturnsError(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{Nodes: []*nodev1.Node{{Id: "node-1", Name: "nop-1"}}}, nil,
	)
	mockClient.EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		nil, fmt.Errorf("JD unavailable"),
	)

	bundle := newTestBundle(t)

	input := ProposeJobsInput{
		JobSpecs: []JobSpecInput{
			{NOPAlias: "nop-1", InternalJobID: "job-1", Spec: "type = 'ccv'"},
		},
	}

	_, err := operations.ExecuteOperation(bundle, ProposeJobs, ProposeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to propose job")
}

func TestProposeJobs_CaseInsensitiveNodeLookup(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().ListNodes(mock.Anything, mock.Anything).Return(
		&nodev1.ListNodesResponse{Nodes: []*nodev1.Node{{Id: "node-1", Name: "NOP-1"}}}, nil,
	)
	mockClient.EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		&jobv1.ProposeJobResponse{
			Proposal: &jobv1.Proposal{
				Id:     "proposal-1",
				JobId:  "jd-job-1",
				Spec:   "type = 'ccv'",
				Status: jobv1.ProposalStatus_PROPOSAL_STATUS_PENDING,
			},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := ProposeJobsInput{
		JobSpecs: []JobSpecInput{
			{NOPAlias: "nop-1", InternalJobID: "job-1", Spec: "type = 'ccv'"},
		},
	}

	report, err := operations.ExecuteOperation(bundle, ProposeJobs, ProposeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)
	require.NoError(t, err)
	require.Len(t, report.Output.ProposedJobs, 1)
}

func TestProposeJobs_EmptyNodeIDs_ReturnsError(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)

	bundle := newTestBundle(t)

	input := ProposeJobsInput{
		JobSpecs: []JobSpecInput{
			{NOPAlias: "nop-1", InternalJobID: "job-1", Spec: "type = 'ccv'"},
		},
	}

	_, err := operations.ExecuteOperation(bundle, ProposeJobs, ProposeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  nil,
	}, input)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "nodeIDs must be specified")
}

func TestProposeJobs_ListNodesCalledWithNodeIDsFilter(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)

	expectedNodeIDs := []string{"node-abc", "node-xyz"}

	mockClient.EXPECT().ListNodes(mock.Anything, mock.MatchedBy(func(req *nodev1.ListNodesRequest) bool {
		if req.Filter == nil {
			return false
		}
		if len(req.Filter.Ids) != len(expectedNodeIDs) {
			return false
		}
		for i, id := range req.Filter.Ids {
			if id != expectedNodeIDs[i] {
				return false
			}
		}
		return true
	})).Return(
		&nodev1.ListNodesResponse{Nodes: []*nodev1.Node{
			{Id: "node-abc", Name: "nop-1"},
			{Id: "node-xyz", Name: "nop-2"},
		}}, nil,
	)
	mockClient.EXPECT().ProposeJob(mock.Anything, mock.Anything).Return(
		&jobv1.ProposeJobResponse{
			Proposal: &jobv1.Proposal{
				Id:     "proposal-1",
				JobId:  "jd-job-1",
				Spec:   "type = 'ccv'",
				Status: jobv1.ProposalStatus_PROPOSAL_STATUS_PENDING,
			},
		}, nil,
	)

	bundle := newTestBundle(t)

	input := ProposeJobsInput{
		JobSpecs: []JobSpecInput{
			{NOPAlias: "nop-1", InternalJobID: "job-1", Spec: "type = 'ccv'"},
		},
	}

	report, err := operations.ExecuteOperation(bundle, ProposeJobs, ProposeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  expectedNodeIDs,
	}, input)
	require.NoError(t, err)
	require.Len(t, report.Output.ProposedJobs, 1)
}
