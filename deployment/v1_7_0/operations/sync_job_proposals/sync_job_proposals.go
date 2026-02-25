package sync_job_proposals

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/ccv"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
)

type SyncJobProposalsInput struct {
	NOPAliases []shared.NOPAlias
}

type SyncJobProposalsOutput struct {
	StatusChanges []ProposalStatusChange
	SpecDrifts    []ProposalSpecDrift
	OrphanedJobs  []OrphanedJob
	Errors        []SyncError
	DataStore     datastore.MutableDataStore
}

type OrphanedJob struct {
	NOPAlias shared.NOPAlias
	JobID    shared.JobID
	JDJobID  string
	Reason   string
}

type ProposalStatusChange struct {
	NOPAlias  shared.NOPAlias
	JobID     shared.JobID
	OldStatus shared.JobProposalStatus
	NewStatus shared.JobProposalStatus
}

type ProposalSpecDrift struct {
	NOPAlias  shared.NOPAlias
	JobID     shared.JobID
	LocalSpec string
	JDSpec    string
}

type SyncError struct {
	ProposalID string
	NOPAlias   shared.NOPAlias
	JobID      shared.JobID
	Error      string
}

type SyncJobProposalsDeps struct {
	Env      deployment.Environment
	JDClient shared.JDClient // Optional: if nil, uses Env.Offchain
}

var SyncJobProposals = operations.NewOperation(
	"sync-job-proposals",
	semver.MustParse("1.0.0"),
	"Syncs job proposal statuses from JD and detects spec drift",
	func(b operations.Bundle, deps SyncJobProposalsDeps, input SyncJobProposalsInput) (SyncJobProposalsOutput, error) {
		ctx := b.GetContext()
		e := deps.Env

		ds := datastore.NewMemoryDataStore()
		if e.DataStore != nil {
			if err := ds.Merge(e.DataStore); err != nil {
				return SyncJobProposalsOutput{}, fmt.Errorf("failed to merge datastore: %w", err)
			}
		}

		output := SyncJobProposalsOutput{
			StatusChanges: make([]ProposalStatusChange, 0),
			SpecDrifts:    make([]ProposalSpecDrift, 0),
			OrphanedJobs:  make([]OrphanedJob, 0),
			Errors:        make([]SyncError, 0),
			DataStore:     ds,
		}

		jdClient := deps.JDClient
		if jdClient == nil {
			if e.Offchain == nil {
				return output, fmt.Errorf("offchain client not available")
			}
			jdClient = e.Offchain
		}

		allJobs, err := ccv.GetAllJobs(ds.Seal())
		if err != nil {
			e.Logger.Debugw("No jobs found in datastore, nothing to sync", "error", err)
			return output, nil
		}

		nopFilter := shared.NOPAliasSliceToSet(input.NOPAliases)

		allowedNodeIDs := shared.NodeIDsToSet(e.NodeIDs)
		if allowedNodeIDs == nil {
			return SyncJobProposalsOutput{}, fmt.Errorf("NodeIDs must be specified")
		}

		for nopAlias, nopJobs := range allJobs {
			if len(nopFilter) > 0 && !nopFilter[nopAlias] {
				continue
			}
			for jobID, localJob := range nopJobs {
				if localJob.Mode != shared.NOPModeCL {
					continue
				}
				if localJob.JDJobID == "" {
					continue
				}
				if localJob.NodeID == "" {
					return SyncJobProposalsOutput{}, fmt.Errorf("job %s has no NodeID", jobID)
				}
				if !allowedNodeIDs[localJob.NodeID] {
					return SyncJobProposalsOutput{}, fmt.Errorf("job %s has NodeID %s which is not in the allowed node list", jobID, localJob.NodeID)
				}
			}
		}

		for nopAlias, nopJobs := range allJobs {
			if len(nopFilter) > 0 && !nopFilter[nopAlias] {
				continue
			}

			for jobID, localJob := range nopJobs {
				if localJob.Mode != shared.NOPModeCL {
					continue
				}

				if localJob.JDJobID == "" {
					e.Logger.Debugw("Skipping job without JD job ID",
						"nopAlias", nopAlias,
						"jobId", jobID)
					continue
				}

				e.Logger.Debugw("Fetching proposals from JD",
					"jdJobId", localJob.JDJobID,
					"nopAlias", nopAlias,
					"jobId", jobID)

				listResp, err := jdClient.ListProposals(ctx, &jobv1.ListProposalsRequest{
					Filter: &jobv1.ListProposalsRequest_Filter{
						JobIds: []string{localJob.JDJobID},
					},
				})
				if err != nil {
					e.Logger.Warnw("Failed to list proposals from JD",
						"jdJobId", localJob.JDJobID,
						"error", err)
					output.Errors = append(output.Errors, SyncError{
						NOPAlias: nopAlias,
						JobID:    jobID,
						Error:    err.Error(),
					})
					continue
				}

				if listResp == nil || len(listResp.Proposals) == 0 {
					e.Logger.Infow("Job no longer exists in JD, marking as orphaned",
						"jdJobId", localJob.JDJobID,
						"nopAlias", nopAlias,
						"jobId", jobID)

					output.OrphanedJobs = append(output.OrphanedJobs, OrphanedJob{
						NOPAlias: nopAlias,
						JobID:    jobID,
						JDJobID:  localJob.JDJobID,
						Reason:   "no proposals found in JD",
					})

					if err := ccv.DeleteJob(ds, nopAlias, jobID); err != nil {
						e.Logger.Errorw("Failed to delete orphaned job",
							"jdJobId", localJob.JDJobID,
							"error", err)
						output.Errors = append(output.Errors, SyncError{
							NOPAlias: nopAlias,
							JobID:    jobID,
							Error:    fmt.Sprintf("failed to delete orphaned job: %s", err.Error()),
						})
					}
					continue
				}

				oldStatus := localJob.LatestStatus()

				for _, jdProposal := range listResp.Proposals {
					localJob.AddProposal(shared.ProposalRevision{
						ProposalID: jdProposal.Id,
						Revision:   jdProposal.Revision,
						Status:     mapJDStatusToLocal(jdProposal.Status),
						Spec:       jdProposal.Spec,
					})

					if jdProposal.Status == jobv1.ProposalStatus_PROPOSAL_STATUS_APPROVED {
						localJob.SetActiveProposal(jdProposal.Id)
					}
				}

				newStatus := localJob.LatestStatus()
				if newStatus != oldStatus {
					e.Logger.Infow("Status change detected",
						"jdJobId", localJob.JDJobID,
						"nopAlias", nopAlias,
						"jobId", jobID,
						"oldStatus", oldStatus,
						"newStatus", newStatus)

					output.StatusChanges = append(output.StatusChanges, ProposalStatusChange{
						NOPAlias:  nopAlias,
						JobID:     jobID,
						OldStatus: oldStatus,
						NewStatus: newStatus,
					})
				}

				latestProposal := localJob.LatestProposal()
				if latestProposal != nil && latestProposal.Spec != localJob.Spec {
					e.Logger.Warnw("Spec drift detected",
						"jdJobId", localJob.JDJobID,
						"nopAlias", nopAlias,
						"jobId", jobID)

					output.SpecDrifts = append(output.SpecDrifts, ProposalSpecDrift{
						NOPAlias:  nopAlias,
						JobID:     jobID,
						LocalSpec: localJob.Spec,
						JDSpec:    latestProposal.Spec,
					})
				}

				if err := ccv.SaveJob(ds, localJob); err != nil {
					e.Logger.Errorw("Failed to save job",
						"jdJobId", localJob.JDJobID,
						"error", err)
					output.Errors = append(output.Errors, SyncError{
						NOPAlias: nopAlias,
						JobID:    jobID,
						Error:    fmt.Sprintf("failed to save job: %s", err.Error()),
					})
				}
			}
		}

		e.Logger.Infow("Job proposals sync completed",
			"statusChanges", len(output.StatusChanges),
			"specDrifts", len(output.SpecDrifts),
			"orphanedJobs", len(output.OrphanedJobs),
			"errors", len(output.Errors))

		return output, nil
	},
)

func mapJDStatusToLocal(jdStatus jobv1.ProposalStatus) shared.JobProposalStatus {
	switch jdStatus {
	case jobv1.ProposalStatus_PROPOSAL_STATUS_PENDING:
		return shared.JobProposalStatusPending
	case jobv1.ProposalStatus_PROPOSAL_STATUS_PROPOSED:
		return shared.JobProposalStatusPending
	case jobv1.ProposalStatus_PROPOSAL_STATUS_APPROVED:
		return shared.JobProposalStatusApproved
	case jobv1.ProposalStatus_PROPOSAL_STATUS_REJECTED:
		return shared.JobProposalStatusRejected
	case jobv1.ProposalStatus_PROPOSAL_STATUS_REVOKED:
		return shared.JobProposalStatusRevoked
	case jobv1.ProposalStatus_PROPOSAL_STATUS_CANCELLED:
		return shared.JobProposalStatusRevoked
	default:
		return shared.JobProposalStatusPending
	}
}
