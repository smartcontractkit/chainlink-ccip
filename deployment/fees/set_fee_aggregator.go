package fees

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// FeeAggregatorForChain represents a fee aggregator address update for a single chain.
// When Contracts is empty the adapter uses its default contract(s).
// When Contracts is provided, the adapter updates only the specified contracts,
// using the full ref (Type + Version + Qualifier) for unambiguous datastore resolution.
type FeeAggregatorForChain struct {
	ChainSelector uint64                `json:"chainSelector" yaml:"chainSelector"`
	FeeAggregator string                `json:"feeAggregator" yaml:"feeAggregator"`
	Contracts     []datastore.AddressRef `json:"contracts,omitempty" yaml:"contracts,omitempty"`
}

// SetFeeAggregatorInput is the top-level input for the SetFeeAggregator changeset.
type SetFeeAggregatorInput struct {
	Version *semver.Version         `json:"version" yaml:"version"`
	Args    []FeeAggregatorForChain `json:"args" yaml:"args"`
	MCMS    mcms.Input              `json:"mcms" yaml:"mcms"`
}

// SetFeeAggregator creates a changeset that sets the fee aggregator address on one or more chains.
func SetFeeAggregator() cldf.ChangeSetV2[SetFeeAggregatorInput] {
	registry := GetFeeAggregatorRegistry()
	mcmsRegistry := changesets.GetRegistry()
	return cldf.CreateChangeSet(makeSetFeeAggregatorApply(registry, mcmsRegistry), makeSetFeeAggregatorVerify(registry))
}

func makeSetFeeAggregatorVerify(registry *FeeAggregatorAdapterRegistry) func(cldf.Environment, SetFeeAggregatorInput) error {
	return func(_ cldf.Environment, cfg SetFeeAggregatorInput) error {
		if cfg.Version == nil {
			return fmt.Errorf("version is required")
		}
		if len(cfg.Args) == 0 {
			return fmt.Errorf("at least one chain must be specified")
		}

		seen := map[uint64]bool{}
		for i, arg := range cfg.Args {
			if seen[arg.ChainSelector] {
				return fmt.Errorf("duplicate chain selector at args[%d]: %d", i, arg.ChainSelector)
			}
			seen[arg.ChainSelector] = true

			if arg.FeeAggregator == "" {
				return fmt.Errorf("fee aggregator address is required at args[%d] (chain=%d)", i, arg.ChainSelector)
			}

			family, err := chain_selectors.GetSelectorFamily(arg.ChainSelector)
			if err != nil {
				return fmt.Errorf("failed to get chain family for selector %d: %w", arg.ChainSelector, err)
			}

			if _, exists := registry.GetFeeAggregatorAdapter(family, cfg.Version); !exists {
				return fmt.Errorf("no fee aggregator adapter found for chain family %s and version %s", family, cfg.Version.String())
			}
		}

		return nil
	}
}

func makeSetFeeAggregatorApply(registry *FeeAggregatorAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, SetFeeAggregatorInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg SetFeeAggregatorInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for _, arg := range cfg.Args {
			family, err := chain_selectors.GetSelectorFamily(arg.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", arg.ChainSelector, err)
			}

			adapter, exists := registry.GetFeeAggregatorAdapter(family, cfg.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no fee aggregator adapter found for chain family %s and version %s", family, cfg.Version.String())
			}

			report, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle,
				adapter.SetFeeAggregator(e),
				e.BlockChains,
				FeeAggregatorForChain{
					ChainSelector: arg.ChainSelector,
					FeeAggregator: arg.FeeAggregator,
					Contracts:     arg.Contracts,
				},
			)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to set fee aggregator for chain %d: %w", arg.ChainSelector, err)
			}

			batchOps = append(batchOps, report.Output.BatchOps...)
			reports = append(reports, report.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithBatchOps(batchOps).
			WithReports(reports).
			Build(cfg.MCMS)
	}
}
