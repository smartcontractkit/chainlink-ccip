package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/shared"
	verifierconfig "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/verifier_config"
)

type GenerateVerifierConfigInput struct {
	// DefaultExecutorQualifier is the qualifier of the executor considered as the default executor. The executor proxy address is used for the verifier.
	DefaultExecutorQualifier string
	// ChainSelectors is the list of chain selectors to consider. Defaults to all chain selectors in the environment.
	ChainSelectors []uint64
	// TargetNOPs limits which NOPs will have their job specs updated. Defaults to all NOPs in the committee when empty.
	TargetNOPs []shared.NOPAlias
	// EnvironmentNOPs is the list of NOP configurations containing signing addresses for each NOP.
	EnvironmentNOPs []verifierconfig.NOPInput
	// Committee contains the committee configuration including aggregators and membership.
	Committee verifierconfig.CommitteeInput
	// PyroscopeURL is the URL of the Pyroscope server for profiling (optional).
	PyroscopeURL string
	// Monitoring is the monitoring configuration containing beholder settings.
	Monitoring shared.MonitoringInput
	// DisableFinalityCheckers is a list of chain selectors (as strings) for which
	// the finality violation checker should be disabled.
	DisableFinalityCheckers []string
}

type GenerateVerifierConfigOutput struct {
	JobSpecs      shared.NOPJobSpecs
	AffectedScope shared.VerifierJobScope
}

type GenerateVerifierConfigDeps struct {
	Env deployment.Environment
}

var GenerateVerifierConfig = operations.NewSequence(
	"generate-verifier-config",
	semver.MustParse("1.0.0"),
	"Generates verifier job specs from datastore contract addresses and explicit input",
	func(b operations.Bundle, deps GenerateVerifierConfigDeps, input GenerateVerifierConfigInput) (GenerateVerifierConfigOutput, error) {
		buildResult, err := operations.ExecuteOperation(b, verifierconfig.BuildConfig, verifierconfig.BuildConfigDeps{
			Env: deps.Env,
		}, verifierconfig.BuildConfigInput{
			CommitteeQualifier: input.Committee.Qualifier,
			ExecutorQualifier:  input.DefaultExecutorQualifier,
			ChainSelectors:     input.ChainSelectors,
		})
		if err != nil {
			return GenerateVerifierConfigOutput{}, fmt.Errorf("failed to build verifier config: %w", err)
		}

		jobSpecsResult, err := operations.ExecuteOperation(b, verifierconfig.BuildJobSpecs, struct{}{}, verifierconfig.BuildJobSpecsInput{
			GeneratedConfig:         buildResult.Output.Config,
			TargetNOPs:              input.TargetNOPs,
			EnvironmentNOPs:         input.EnvironmentNOPs,
			Committee:               input.Committee,
			PyroscopeURL:            input.PyroscopeURL,
			Monitoring:              input.Monitoring,
			DisableFinalityCheckers: input.DisableFinalityCheckers,
		})
		if err != nil {
			return GenerateVerifierConfigOutput{}, fmt.Errorf("failed to build verifier job specs: %w", err)
		}

		return GenerateVerifierConfigOutput{
			JobSpecs:      jobSpecsResult.Output.JobSpecs,
			AffectedScope: jobSpecsResult.Output.AffectedScope,
		}, nil
	},
)
