package changesets

import (
	"fmt"
	"slices"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/ccv"
	aggconfig "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/operations/aggregator_config"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/sequences"
)

// GenerateAggregatorConfigCfg is an alias for the changeset input type.
type GenerateAggregatorConfigCfg = aggconfig.BuildConfigInput

// GenerateAggregatorConfig creates a changeset that generates the aggregator configuration
// by scanning on-chain CommitteeVerifier contracts.
func GenerateAggregatorConfig() deployment.ChangeSetV2[aggconfig.BuildConfigInput] {
	validate := func(e deployment.Environment, cfg aggconfig.BuildConfigInput) error {
		if cfg.ServiceIdentifier == "" {
			return fmt.Errorf("service identifier is required")
		}
		if cfg.CommitteeQualifier == "" {
			return fmt.Errorf("committee qualifier is required")
		}
		envSelectors := e.BlockChains.ListChainSelectors()
		for _, s := range cfg.ChainSelectors {
			if !slices.Contains(envSelectors, s) {
				return fmt.Errorf("selector %d is not available in environment", s)
			}
		}
		return nil
	}

	apply := func(e deployment.Environment, cfg aggconfig.BuildConfigInput) (deployment.ChangesetOutput, error) {
		input := cfg
		if len(input.ChainSelectors) == 0 {
			input.ChainSelectors = e.BlockChains.ListChainSelectors()
		}
		deps := sequences.GenerateAggregatorConfigDeps{
			Env: e,
		}

		report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.GenerateAggregatorConfig, deps, input)
		if err != nil {
			return deployment.ChangesetOutput{
				Reports: report.ExecutionReports,
			}, fmt.Errorf("failed to generate aggregator config: %w", err)
		}

		outputDS := datastore.NewMemoryDataStore()
		if e.DataStore != nil {
			if err := outputDS.Merge(e.DataStore); err != nil {
				return deployment.ChangesetOutput{
					Reports: report.ExecutionReports,
				}, fmt.Errorf("failed to merge existing datastore: %w", err)
			}
		}

		aggCfg := report.Output.Committee.ToModelCommittee()
		if err := ccv.SaveAggregatorConfig(outputDS, report.Output.ServiceIdentifier, aggCfg); err != nil {
			return deployment.ChangesetOutput{
				Reports: report.ExecutionReports,
			}, fmt.Errorf("failed to save aggregator config: %w", err)
		}

		return deployment.ChangesetOutput{
			Reports:   report.ExecutionReports,
			DataStore: outputDS,
		}, nil
	}

	return deployment.CreateChangeSet(apply, validate)
}
