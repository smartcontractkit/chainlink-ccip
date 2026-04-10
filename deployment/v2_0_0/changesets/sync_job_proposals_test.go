package changesets_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_offchain "github.com/smartcontractkit/chainlink-deployments-framework/offchain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	csav1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/csa"
	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"
	nodev1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/node"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/offchain/shared"
)

// stubOffchainClient satisfies offchain.Client for tests that only verify
// a non-nil client is present; no methods are ever called on it.
type stubOffchainClient struct {
	jobv1.JobServiceClient
	nodev1.NodeServiceClient
	csav1.CSAServiceClient
}

var _ cldf_offchain.Client = (*stubOffchainClient)(nil)

// syncMockOffchain embeds stubOffchainClient and overrides ListProposals so
// that tests that DO invoke ListProposals don't hit the nil-interface panic.
type syncMockOffchain struct {
	stubOffchainClient
	listProposalsFn func(context.Context, *jobv1.ListProposalsRequest, ...grpc.CallOption) (*jobv1.ListProposalsResponse, error)
}

func (m *syncMockOffchain) ListProposals(ctx context.Context, in *jobv1.ListProposalsRequest, opts ...grpc.CallOption) (*jobv1.ListProposalsResponse, error) {
	return m.listProposalsFn(ctx, in, opts...)
}

func newSyncEnv(t *testing.T, client cldf_offchain.Client) deployment.Environment {
	t.Helper()
	lggr := logger.Test(t)
	return deployment.Environment{
		Logger:   lggr,
		Offchain: client,
		OperationsBundle: cldf_ops.NewBundle(
			func() context.Context { return context.Background() },
			lggr,
			cldf_ops.NewMemoryReporter(),
		),
	}
}

func TestSyncJobProposals_Validate_NilOffchainReturnsError(t *testing.T) {
	env := newSyncEnv(t, nil)
	cs := changesets.SyncJobProposals()

	err := cs.VerifyPreconditions(env, changesets.SyncJobProposalsInput{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "offchain client (JD) is required")
}

func TestSyncJobProposals_Validate_NonNilOffchainPasses(t *testing.T) {
	env := newSyncEnv(t, &stubOffchainClient{})
	cs := changesets.SyncJobProposals()

	err := cs.VerifyPreconditions(env, changesets.SyncJobProposalsInput{})
	require.NoError(t, err)
}

func TestSyncJobProposals_Apply_EmptyInputSucceeds(t *testing.T) {
	// No jobs in the datastore and no JD responses to mock; the operation
	// should return cleanly with an empty output.
	env := newSyncEnv(t, &stubOffchainClient{})
	cs := changesets.SyncJobProposals()

	out, err := cs.Apply(env, changesets.SyncJobProposalsInput{
		NOPAliases: []shared.NOPAlias{},
	})
	require.NoError(t, err)
	// DataStore may be nil or empty; just confirm no error path was hit.
	_ = out
}

// TestSyncJobProposals_Apply_SpecDriftLogged verifies that when JD returns a
// proposal with a spec that differs from the locally stored spec, the Apply
// function completes without error (spec drift is logged, not fatal).
func TestSyncJobProposals_Apply_SpecDriftLogged(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		JobID:    shared.JobID("nop1-agg-1-c1-verifier"),
		JDJobID:  "jd-job-id-1",
		NOPAlias: shared.NOPAlias("nop1"),
		NodeID:   "node-1",
		Mode:     shared.NOPModeCL,
		Spec:     "local-spec-v1",
	}))

	mockClient := &syncMockOffchain{
		listProposalsFn: func(_ context.Context, _ *jobv1.ListProposalsRequest, _ ...grpc.CallOption) (*jobv1.ListProposalsResponse, error) {
			return &jobv1.ListProposalsResponse{
				Proposals: []*jobv1.Proposal{
					{
						Id:     "proposal-1",
						Spec:   "jd-spec-v2", // differs from local spec
						Status: jobv1.ProposalStatus_PROPOSAL_STATUS_PROPOSED,
					},
				},
			}, nil
		},
	}

	env := newSyncEnv(t, mockClient)
	env.NodeIDs = []string{"node-1"}
	env.DataStore = ds.Seal()

	cs := changesets.SyncJobProposals()
	out, err := cs.Apply(env, changesets.SyncJobProposalsInput{})
	require.NoError(t, err)
	require.NotNil(t, out)
}

// TestSyncJobProposals_Apply_OperationError exercises the err != nil branch in
// the Apply body: the operation returns an error when NodeIDs is nil but the
// datastore contains a CL-mode job with a JD job ID.
func TestSyncJobProposals_Apply_OperationError(t *testing.T) {
	ds := datastore.NewMemoryDataStore()
	require.NoError(t, offchain.SaveJob(ds, shared.JobInfo{
		JobID:    shared.JobID("nop1-agg-1-c1-verifier"),
		JDJobID:  "jd-job-id-1",
		NOPAlias: shared.NOPAlias("nop1"),
		NodeID:   "node-1",
		Mode:     shared.NOPModeCL,
		Spec:     "some-spec",
	}))

	env := newSyncEnv(t, &syncMockOffchain{
		listProposalsFn: func(_ context.Context, _ *jobv1.ListProposalsRequest, _ ...grpc.CallOption) (*jobv1.ListProposalsResponse, error) {
			return &jobv1.ListProposalsResponse{}, nil
		},
	})
	// NodeIDs is deliberately left nil → operation errors with "NodeIDs must be specified".
	env.DataStore = ds.Seal()

	cs := changesets.SyncJobProposals()
	_, err := cs.Apply(env, changesets.SyncJobProposalsInput{})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "failed to sync job proposals")
}
