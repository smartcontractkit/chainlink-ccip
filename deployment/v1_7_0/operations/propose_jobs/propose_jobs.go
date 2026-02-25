package propose_jobs

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	jobv1 "github.com/smartcontractkit/chainlink-protos/job-distributor/v1/job"
	"github.com/smartcontractkit/chainlink-protos/job-distributor/v1/shared/ptypes"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
)

type JobSpecInput struct {
	NOPAlias      shared.NOPAlias
	InternalJobID shared.JobID
	Spec          string
}

type ProposeJobsInput struct {
	JobSpecs []JobSpecInput
	Labels   map[string]string
}

type ProposeJobsOutput struct {
	ProposedJobs []ProposedJob
}

type ProposedJob struct {
	NodeID        string
	NodeName      string
	ProposalID    string
	JobID         string
	InternalJobID shared.JobID
	Spec          string
	NOPAlias      shared.NOPAlias
	Revision      int64
}

type ProposeJobsDeps struct {
	JDClient shared.JDClient
	Logger   logger.Logger
	NodeIDs  []string
}

var ProposeJobs = operations.NewOperation(
	"propose-jobs",
	semver.MustParse("1.0.0"),
	"Proposes jobs to nodes via the job distributor",
	func(b operations.Bundle, deps ProposeJobsDeps, input ProposeJobsInput) (ProposeJobsOutput, error) {
		ctx := b.GetContext()
		lggr := deps.Logger

		output := ProposeJobsOutput{
			ProposedJobs: make([]ProposedJob, 0, len(input.JobSpecs)),
		}

		if len(input.JobSpecs) == 0 {
			return output, nil
		}

		lookup, err := shared.FetchNodeLookup(ctx, deps.JDClient, deps.NodeIDs)
		if err != nil {
			return output, err
		}

		for _, jobSpecInput := range input.JobSpecs {
			nopAlias := string(jobSpecInput.NOPAlias)
			node, ok := lookup.FindByName(nopAlias)
			if !ok {
				lggr.Warnw("Node not found for job spec",
					"nopAlias", nopAlias,
					"internalJobId", jobSpecInput.InternalJobID)
				return output, fmt.Errorf("node not found for job spec: %s", nopAlias)
			}

			var jobLabels []*ptypes.Label
			for k, v := range input.Labels {
				jobLabels = append(jobLabels, &ptypes.Label{
					Key:   k,
					Value: &v,
				})
			}

			lggr.Debugw("Proposing job",
				"nodeId", node.Id,
				"nodeName", node.Name,
				"nopAlias", nopAlias,
				"internalJobId", jobSpecInput.InternalJobID)

			resp, err := deps.JDClient.ProposeJob(ctx, &jobv1.ProposeJobRequest{
				NodeId: node.Id,
				Spec:   jobSpecInput.Spec,
				Labels: jobLabels,
			})
			if err != nil {
				lggr.Errorw("Failed to propose job",
					"nodeId", node.Id,
					"nodeName", node.Name,
					"internalJobId", jobSpecInput.InternalJobID,
					"error", err)
				return output, fmt.Errorf("failed to propose job %s for node %s: %w", jobSpecInput.InternalJobID, node.Name, err)
			}

			lggr.Infow("Job proposed successfully",
				"nodeId", node.Id,
				"nodeName", node.Name,
				"proposalId", resp.Proposal.Id,
				"jdJobId", resp.Proposal.JobId,
				"internalJobId", jobSpecInput.InternalJobID)

			output.ProposedJobs = append(output.ProposedJobs, ProposedJob{
				NodeID:        node.Id,
				NodeName:      node.Name,
				ProposalID:    resp.Proposal.Id,
				JobID:         resp.Proposal.JobId,
				InternalJobID: jobSpecInput.InternalJobID,
				Spec:          resp.Proposal.Spec,
				NOPAlias:      jobSpecInput.NOPAlias,
				Revision:      resp.Proposal.Revision,
			})
		}

		return output, nil
	},
)
