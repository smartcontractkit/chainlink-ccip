package revoke_jobs

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
)

type RevokeJobsInput struct {
	Jobs []shared.JobInfo
}

type RevokeJobsOutput struct {
	RevokedJobs []RevokedJob
	Errors      []RevokeError
}

type RevokedJob struct {
	JobID    shared.JobID
	NOPAlias shared.NOPAlias
}

type RevokeError struct {
	JobID    shared.JobID
	NOPAlias shared.NOPAlias
	Error    string
}

type RevokeJobsDeps struct {
	JDClient shared.JDClient
	Logger   logger.Logger
	NodeIDs  []string
}

var RevokeJobs = operations.NewOperation(
	"revoke-jobs",
	semver.MustParse("1.0.0"),
	"Revokes job proposals from nodes via the job distributor",
	func(b operations.Bundle, deps RevokeJobsDeps, input RevokeJobsInput) (RevokeJobsOutput, error) {
		ctx := b.GetContext()
		lggr := deps.Logger

		output := RevokeJobsOutput{
			RevokedJobs: make([]RevokedJob, 0),
			Errors:      make([]RevokeError, 0),
		}

		if len(input.Jobs) == 0 {
			return output, nil
		}

		allowedNodeIDs := shared.NodeIDsToSet(deps.NodeIDs)
		if allowedNodeIDs == nil {
			return output, fmt.Errorf("NodeIDs must be specified")
		}

		for _, job := range input.Jobs {
			if job.NodeID == "" {
				return output, fmt.Errorf("job %s has no NodeID - cannot validate node ownership", job.JobID)
			}
			if !allowedNodeIDs[job.NodeID] {
				return output, fmt.Errorf("job %s has NodeID %s which is not in the allowed node list", job.JobID, job.NodeID)
			}
		}

		for _, job := range input.Jobs {
			if job.JDJobID == "" {
				lggr.Warnw("Skipping job with no JD job ID",
					"jobId", job.JobID,
					"nopAlias", job.NOPAlias)
				output.Errors = append(output.Errors, RevokeError{
					JobID:    job.JobID,
					NOPAlias: job.NOPAlias,
					Error:    "no JD job ID",
				})
				continue
			}

			_, err := deps.JDClient.RevokeJob(ctx, &jobv1.RevokeJobRequest{
				IdOneof: &jobv1.RevokeJobRequest_Id{Id: job.JDJobID},
			})
			if err != nil {
				lggr.Errorw("Failed to revoke job",
					"jobId", job.JobID,
					"error", err)
				output.Errors = append(output.Errors, RevokeError{
					JobID:    job.JobID,
					NOPAlias: job.NOPAlias,
					Error:    err.Error(),
				})
				continue
			}

			lggr.Infow("Job proposal revoked successfully",
				"jobId", job.JobID,
				"nopAlias", job.NOPAlias)

			output.RevokedJobs = append(output.RevokedJobs, RevokedJob{
				JobID:    job.JobID,
				NOPAlias: job.NOPAlias,
			})
		}

		return output, nil
	},
)
