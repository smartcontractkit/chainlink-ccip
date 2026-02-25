package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/ccv"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/propose_jobs"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/revoke_jobs"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
)

type NOPContext struct {
	Modes      map[shared.NOPAlias]shared.NOPMode
	TargetNOPs []shared.NOPAlias
	AllNOPs    []shared.NOPAlias
}

type ManageJobProposalsInput struct {
	JobSpecs      shared.NOPJobSpecs
	AffectedScope shared.JobScope
	Labels        map[string]string
	NOPs          NOPContext
}

type ManageJobProposalsOutput struct {
	Jobs        []shared.JobInfo
	RevokedJobs []shared.JobInfo
	DataStore   datastore.MutableDataStore
}

type ManageJobProposalsDeps struct {
	Env deployment.Environment
}

var ManageJobProposals = operations.NewSequence(
	"manage-job-proposals",
	semver.MustParse("1.0.0"),
	"Manages job proposals by proposing new jobs and revoking orphaned ones via JD",
	func(b operations.Bundle, deps ManageJobProposalsDeps, input ManageJobProposalsInput) (ManageJobProposalsOutput, error) {
		e := deps.Env

		ds := datastore.NewMemoryDataStore()
		if e.DataStore != nil {
			if err := ds.Merge(e.DataStore); err != nil {
				return ManageJobProposalsOutput{}, fmt.Errorf("failed to merge datastore: %w", err)
			}
		}

		existingJobs, err := ccv.GetAllJobs(ds.Seal())
		if err != nil {
			e.Logger.Warnw("Failed to load existing jobs, will propose all jobs", "error", err)
			existingJobs = nil
		}

		jobs := buildJobsFromJobSpecs(input.JobSpecs, input.NOPs.Modes)

		changedJobs := filterChangedJobs(jobs, existingJobs)
		clModeSpecs := extractCLModeSpecs(changedJobs)

		if len(changedJobs) < len(jobs) {
			e.Logger.Infow("Skipping unchanged jobs", "total", len(jobs), "changed", len(changedJobs), "skipped", len(jobs)-len(changedJobs))
		}

		if len(clModeSpecs) > 0 {
			if e.Offchain == nil {
				return ManageJobProposalsOutput{}, fmt.Errorf("CL mode NOPs require JD but e.Offchain is nil")
			}

			proposeReport, err := operations.ExecuteOperation(
				b,
				propose_jobs.ProposeJobs,
				propose_jobs.ProposeJobsDeps{
					JDClient: e.Offchain,
					Logger:   e.Logger,
					NodeIDs:  e.NodeIDs,
				},
				propose_jobs.ProposeJobsInput{
					JobSpecs: clModeSpecs,
					Labels:   input.Labels,
				},
			)
			if err != nil {
				return ManageJobProposalsOutput{}, fmt.Errorf("failed to propose CL mode jobs: %w", err)
			}

			updateJobsWithJDResponse(jobs, proposeReport.Output.ProposedJobs)
		}

		if err := ccv.SaveJobs(ds, jobs); err != nil {
			return ManageJobProposalsOutput{}, fmt.Errorf("failed to save jobs: %w", err)
		}

		expectedJobsByNOP := extractExpectedJobsByNOP(jobs)

		orphanedJobs, err := ccv.CollectOrphanedJobs(
			ds.Seal(),
			input.AffectedScope,
			expectedJobsByNOP,
			shared.NOPAliasSliceToSet(input.NOPs.TargetNOPs),
			shared.NOPAliasSliceToSet(input.NOPs.AllNOPs),
		)
		if err != nil {
			e.Logger.Warnw("Failed to collect orphaned jobs", "error", err)
			orphanedJobs = nil
		}

		revokedJobs := make([]shared.JobInfo, 0)
		if len(orphanedJobs) > 0 {
			clOrphanedJobs := filterCLModeJobs(orphanedJobs)
			if len(clOrphanedJobs) > 0 && e.Offchain != nil {
				revokeReport, revokeErr := operations.ExecuteOperation(
					b,
					revoke_jobs.RevokeJobs,
					revoke_jobs.RevokeJobsDeps{
						JDClient: e.Offchain,
						Logger:   e.Logger,
						NodeIDs:  e.NodeIDs,
					},
					revoke_jobs.RevokeJobsInput{
						Jobs: clOrphanedJobs,
					},
				)
				if revokeErr != nil {
					return ManageJobProposalsOutput{
						Jobs:      jobs,
						DataStore: ds,
					}, fmt.Errorf("failed to revoke orphaned CL mode jobs: %w", revokeErr)
				}
				e.Logger.Infow("Revoked orphaned jobs", "count", len(revokeReport.Output.RevokedJobs))
			}

			if cleanupErr := ccv.CleanupOrphanedJobs(ds, orphanedJobs); cleanupErr != nil {
				return ManageJobProposalsOutput{
					Jobs:      jobs,
					DataStore: ds,
				}, fmt.Errorf("failed to cleanup orphaned jobs: %w", cleanupErr)
			}
			revokedJobs = orphanedJobs
		}

		return ManageJobProposalsOutput{
			Jobs:        jobs,
			RevokedJobs: revokedJobs,
			DataStore:   ds,
		}, nil
	},
)

func buildJobsFromJobSpecs(jobSpecs shared.NOPJobSpecs, nopModes map[shared.NOPAlias]shared.NOPMode) []shared.JobInfo {
	totalJobs := 0
	for _, jobSpecsByID := range jobSpecs {
		totalJobs += len(jobSpecsByID)
	}
	jobs := make([]shared.JobInfo, 0, totalJobs)

	for nopAlias, jobSpecsByID := range jobSpecs {
		mode := nopModes[nopAlias]
		if mode == "" {
			mode = shared.NOPModeCL
		}

		for jobID, spec := range jobSpecsByID {
			job := shared.JobInfo{
				Spec:          spec,
				ExternalJobID: jobID.ToExternalJobID(),
				JobID:         jobID,
				NOPAlias:      nopAlias,
				Mode:          mode,
				Proposals:     make(map[string]shared.ProposalRevision),
			}
			jobs = append(jobs, job)
		}
	}

	return jobs
}

func extractCLModeSpecs(jobs []shared.JobInfo) []propose_jobs.JobSpecInput {
	clModeSpecs := make([]propose_jobs.JobSpecInput, 0)
	for _, job := range jobs {
		if job.Mode == shared.NOPModeCL {
			clModeSpecs = append(clModeSpecs, propose_jobs.JobSpecInput{
				NOPAlias:      job.NOPAlias,
				InternalJobID: job.JobID,
				Spec:          job.Spec,
			})
		}
	}
	return clModeSpecs
}

func updateJobsWithJDResponse(jobs []shared.JobInfo, proposedJobs []propose_jobs.ProposedJob) {
	proposedJobMap := make(map[shared.JobID]propose_jobs.ProposedJob, len(proposedJobs))
	for _, pj := range proposedJobs {
		proposedJobMap[pj.InternalJobID] = pj
	}

	for i := range jobs {
		if jobs[i].Mode != shared.NOPModeCL {
			continue
		}
		if pj, ok := proposedJobMap[jobs[i].JobID]; ok {
			jobs[i].JDJobID = pj.JobID
			jobs[i].NodeID = pj.NodeID
			jobs[i].AddProposal(shared.ProposalRevision{
				ProposalID: pj.ProposalID,
				Revision:   pj.Revision,
				Status:     shared.JobProposalStatusPending,
				Spec:       pj.Spec,
			})
		}
	}
}

func extractExpectedJobsByNOP(jobs []shared.JobInfo) map[shared.NOPAlias]map[shared.JobID]bool {
	result := make(map[shared.NOPAlias]map[shared.JobID]bool)
	for _, job := range jobs {
		if result[job.NOPAlias] == nil {
			result[job.NOPAlias] = make(map[shared.JobID]bool)
		}
		result[job.NOPAlias][job.JobID] = true
	}
	return result
}

func filterCLModeJobs(jobs []shared.JobInfo) []shared.JobInfo {
	result := make([]shared.JobInfo, 0)
	for _, j := range jobs {
		if j.Mode == shared.NOPModeCL {
			result = append(result, j)
		}
	}
	return result
}

func filterChangedJobs(newJobs []shared.JobInfo, existingJobs map[shared.NOPAlias]map[shared.JobID]shared.JobInfo) []shared.JobInfo {
	if existingJobs == nil {
		return newJobs
	}

	changed := make([]shared.JobInfo, 0, len(newJobs))
	for _, job := range newJobs {
		nopJobs, nopExists := existingJobs[job.NOPAlias]
		if !nopExists {
			changed = append(changed, job)
			continue
		}

		existing, jobExists := nopJobs[job.JobID]
		if !jobExists {
			changed = append(changed, job)
			continue
		}

		latestProposal := existing.LatestProposal()
		if latestProposal == nil || latestProposal.Spec != job.Spec {
			changed = append(changed, job)
		}
	}
	return changed
}
