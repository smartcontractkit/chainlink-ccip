package revoke_jobs

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"

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

func newJobWithProposal(jobID shared.JobID, jdJobID string, nopAlias shared.NOPAlias, nodeID string) shared.JobInfo {
	return shared.JobInfo{
		JobID:    jobID,
		JDJobID:  jdJobID,
		NOPAlias: nopAlias,
		NodeID:   nodeID,
		Mode:     shared.NOPModeCL,
	}
}

func TestRevokeJobs_EmptyJobs_ReturnsEmpty(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	bundle := newTestBundle(t)

	input := RevokeJobsInput{
		Jobs: []shared.JobInfo{},
	}

	report, err := operations.ExecuteOperation(bundle, RevokeJobs, RevokeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
	}, input)
	require.NoError(t, err)
	assert.Empty(t, report.Output.RevokedJobs)
	assert.Empty(t, report.Output.Errors)
}

func TestRevokeJobs_Success_RevokesLatestProposal(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().RevokeJob(mock.Anything, mock.Anything).Return(
		&jobv1.RevokeJobResponse{}, nil,
	)

	bundle := newTestBundle(t)

	input := RevokeJobsInput{
		Jobs: []shared.JobInfo{
			newJobWithProposal("nop-1-default-executor", "jd-proposal-123", "nop-1", "node-1"),
		},
	}

	report, err := operations.ExecuteOperation(bundle, RevokeJobs, RevokeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)
	require.NoError(t, err)
	require.Len(t, report.Output.RevokedJobs, 1)
	assert.Equal(t, shared.JobID("nop-1-default-executor"), report.Output.RevokedJobs[0].JobID)
	assert.Equal(t, shared.NOPAlias("nop-1"), report.Output.RevokedJobs[0].NOPAlias)
	assert.Empty(t, report.Output.Errors)
}

func TestRevokeJobs_MultipleJobs_AllRevoked(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().RevokeJob(mock.Anything, mock.Anything).Return(
		&jobv1.RevokeJobResponse{}, nil,
	).Times(3)

	bundle := newTestBundle(t)

	input := RevokeJobsInput{
		Jobs: []shared.JobInfo{
			newJobWithProposal("job-1", "proposal-1", "nop-1", "node-1"),
			newJobWithProposal("job-2", "proposal-2", "nop-2", "node-2"),
			newJobWithProposal("job-3", "proposal-3", "nop-1", "node-1"),
		},
	}

	report, err := operations.ExecuteOperation(bundle, RevokeJobs, RevokeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1", "node-2"},
	}, input)
	require.NoError(t, err)
	require.Len(t, report.Output.RevokedJobs, 3)
	assert.Empty(t, report.Output.Errors)
}

func TestRevokeJobs_NoJDJobID_SkipsAndReportsError(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().RevokeJob(mock.Anything, mock.Anything).Return(
		&jobv1.RevokeJobResponse{}, nil,
	).Once()

	bundle := newTestBundle(t)

	input := RevokeJobsInput{
		Jobs: []shared.JobInfo{
			{JobID: "job-1", NOPAlias: "nop-1", NodeID: "node-1"},
			newJobWithProposal("job-2", "jd-job-2", "nop-2", "node-2"),
		},
	}

	report, err := operations.ExecuteOperation(bundle, RevokeJobs, RevokeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1", "node-2"},
	}, input)
	require.NoError(t, err)

	require.Len(t, report.Output.RevokedJobs, 1)
	assert.Equal(t, shared.JobID("job-2"), report.Output.RevokedJobs[0].JobID)

	require.Len(t, report.Output.Errors, 1)
	assert.Equal(t, shared.JobID("job-1"), report.Output.Errors[0].JobID)
	assert.Contains(t, report.Output.Errors[0].Error, "no JD job ID")
}

func TestRevokeJobs_JDError_ReportsErrorAndContinues(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().RevokeJob(mock.Anything, mock.Anything).Return(
		nil, fmt.Errorf("JD unavailable"),
	).Times(2)

	bundle := newTestBundle(t)

	input := RevokeJobsInput{
		Jobs: []shared.JobInfo{
			newJobWithProposal("job-1", "proposal-1", "nop-1", "node-1"),
			newJobWithProposal("job-2", "proposal-2", "nop-2", "node-2"),
		},
	}

	report, err := operations.ExecuteOperation(bundle, RevokeJobs, RevokeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1", "node-2"},
	}, input)
	require.NoError(t, err)

	assert.Empty(t, report.Output.RevokedJobs)
	require.Len(t, report.Output.Errors, 2)
	assert.Contains(t, report.Output.Errors[0].Error, "JD unavailable")
	assert.Contains(t, report.Output.Errors[1].Error, "JD unavailable")
}

func TestRevokeJobs_EmptyNodeIDsDeps_ReturnsError(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	bundle := newTestBundle(t)

	input := RevokeJobsInput{
		Jobs: []shared.JobInfo{
			newJobWithProposal("job-1", "proposal-1", "nop-1", "node-1"),
		},
	}

	_, err := operations.ExecuteOperation(bundle, RevokeJobs, RevokeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  nil,
	}, input)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "NodeIDs must be specified")
}

func TestRevokeJobs_NodeNotInAllowedList_ReturnsError(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	bundle := newTestBundle(t)

	input := RevokeJobsInput{
		Jobs: []shared.JobInfo{
			newJobWithProposal("job-1", "proposal-1", "nop-1", "node-1"),
			newJobWithProposal("job-2", "proposal-2", "nop-2", "node-not-allowed"),
		},
	}

	_, err := operations.ExecuteOperation(bundle, RevokeJobs, RevokeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1", "node-2"},
	}, input)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "node-not-allowed")
	assert.Contains(t, err.Error(), "not in the allowed node list")
}

func TestRevokeJobs_EmptyNodeID_ReturnsError(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	bundle := newTestBundle(t)

	input := RevokeJobsInput{
		Jobs: []shared.JobInfo{
			newJobWithProposal("job-1", "proposal-1", "nop-1", ""),
		},
	}

	_, err := operations.ExecuteOperation(bundle, RevokeJobs, RevokeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1"},
	}, input)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "has no NodeID")
}

func TestRevokeJobs_WithNodeIDs_ValidatesBeforeRevoke(t *testing.T) {
	mockClient := deploymocks.NewMockJDClient(t)
	mockClient.EXPECT().RevokeJob(mock.Anything, mock.Anything).Return(
		&jobv1.RevokeJobResponse{}, nil,
	).Times(2)

	bundle := newTestBundle(t)

	input := RevokeJobsInput{
		Jobs: []shared.JobInfo{
			newJobWithProposal("job-1", "proposal-1", "nop-1", "node-1"),
			newJobWithProposal("job-2", "proposal-2", "nop-2", "node-2"),
		},
	}

	report, err := operations.ExecuteOperation(bundle, RevokeJobs, RevokeJobsDeps{
		JDClient: mockClient,
		Logger:   newTestLogger(t),
		NodeIDs:  []string{"node-1", "node-2"},
	}, input)
	require.NoError(t, err)
	require.Len(t, report.Output.RevokedJobs, 2)
	assert.Empty(t, report.Output.Errors)
}
