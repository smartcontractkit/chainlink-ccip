package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	aggconfig "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/aggregator_config"
)

// GenerateAggregatorConfigOutput contains the output of the aggregator config generation sequence.
type GenerateAggregatorConfigOutput struct {
	ServiceIdentifier string
	Committee         *aggconfig.Committee
}

// GenerateAggregatorConfigDeps contains the dependencies for the sequence.
type GenerateAggregatorConfigDeps struct {
	Env deployment.Environment
}

// GenerateAggregatorConfig is a sequence that generates the aggregator configuration
// by scanning on-chain CommitteeVerifier contracts.
var GenerateAggregatorConfig = operations.NewSequence(
	"generate-aggregator-config",
	semver.MustParse("1.0.0"),
	"Generates the aggregator committee configuration from on-chain state",
	func(b operations.Bundle, deps GenerateAggregatorConfigDeps, input aggconfig.BuildConfigInput) (GenerateAggregatorConfigOutput, error) {
		// Execute the build config operation
		result, err := operations.ExecuteOperation(b, aggconfig.BuildConfig, aggconfig.BuildConfigDeps{
			Env: deps.Env,
		}, input)
		if err != nil {
			return GenerateAggregatorConfigOutput{}, fmt.Errorf("failed to build aggregator config: %w", err)
		}

		return GenerateAggregatorConfigOutput{
			ServiceIdentifier: result.Output.ServiceIdentifier,
			Committee:         result.Output.Committee,
		}, nil
	},
)
