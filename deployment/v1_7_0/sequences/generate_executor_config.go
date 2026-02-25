package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	executorconfig "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/executor_config"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
)

type GenerateExecutorConfigInput struct {
	// ExecutorQualifier is the qualifier of the executor that is configured as part of this operation.
	ExecutorQualifier string
	// ChainSelectors is the list of chain selectors to consider. Defaults to all chain selectors in the environment.
	ChainSelectors []uint64
	// TargetNOPs limits which NOPs will have their job specs updated. Defaults to all NOPs in the executor pool when empty.
	TargetNOPs []shared.NOPAlias
	// ExecutorPool is the executor pool configuration containing pool membership and execution parameters.
	ExecutorPool executorconfig.ExecutorPoolInput
	// IndexerAddress is the address of the indexer service used by executors.
	IndexerAddress []string
	// PyroscopeURL is the URL of the Pyroscope server for profiling (optional).
	PyroscopeURL string
	// Monitoring is the monitoring configuration containing beholder settings.
	Monitoring shared.MonitoringInput
	// NOPModes maps NOP alias to job proposal mode (cl or standalone).
	NOPModes map[shared.NOPAlias]shared.NOPMode
}

type GenerateExecutorConfigOutput struct {
	JobSpecs      shared.NOPJobSpecs
	AffectedScope shared.ExecutorJobScope
}

type GenerateExecutorConfigDeps struct {
	Env deployment.Environment
}

var GenerateExecutorConfig = operations.NewSequence(
	"generate-executor-config",
	semver.MustParse("1.0.0"),
	"Generates executor job specs from datastore contract addresses and explicit input",
	func(b operations.Bundle, deps GenerateExecutorConfigDeps, input GenerateExecutorConfigInput) (GenerateExecutorConfigOutput, error) {
		buildResult, err := operations.ExecuteOperation(b, executorconfig.BuildConfig, executorconfig.BuildConfigDeps{
			Env: deps.Env,
		}, executorconfig.BuildConfigInput{
			ExecutorQualifier: input.ExecutorQualifier,
			ChainSelectors:    input.ChainSelectors,
		})
		if err != nil {
			return GenerateExecutorConfigOutput{}, fmt.Errorf("failed to build executor config: %w", err)
		}

		jobSpecsResult, err := operations.ExecuteOperation(b, executorconfig.BuildJobSpecs, struct{}{}, executorconfig.BuildJobSpecsInput{
			GeneratedConfig:   buildResult.Output.Config,
			ExecutorQualifier: input.ExecutorQualifier,
			TargetNOPs:        input.TargetNOPs,
			ExecutorPool:      input.ExecutorPool,
			IndexerAddress:    input.IndexerAddress,
			PyroscopeURL:      input.PyroscopeURL,
			Monitoring:        input.Monitoring,
		})
		if err != nil {
			return GenerateExecutorConfigOutput{}, fmt.Errorf("failed to build executor job specs: %w", err)
		}

		return GenerateExecutorConfigOutput{
			JobSpecs:      jobSpecsResult.Output.JobSpecs,
			AffectedScope: jobSpecsResult.Output.AffectedScope,
		}, nil
	},
)
