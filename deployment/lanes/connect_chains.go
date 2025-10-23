package lanes

import (
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// ConfigureTokensForTransfers returns a changeset that configures tokens on multiple chains for transfers with other chains.
func ConnectChains(tokenRegistry *LaneAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[ConnectChainsConfig] {
	return cldf.CreateChangeSet(makeApply(tokenRegistry, mcmsRegistry), makeVerify(tokenRegistry, mcmsRegistry))
}

func makeVerify(_ *LaneAdapterRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, ConnectChainsConfig) error {
	return func(e cldf.Environment, cfg ConnectChainsConfig) error {
		// TODO: implement
		return nil
	}
}

func makeApply(laneRegistry *LaneAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ConnectChainsConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ConnectChainsConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for _, lane := range cfg.Lanes {
			src, dest := &lane.Source, &lane.Dest
			srcFamily, err := chain_selectors.GetSelectorFamily(src.Selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			srcAdapter, exists := laneRegistry.GetLaneAdapter(srcFamily, lane.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", srcFamily)
			}
			destFamily, err := chain_selectors.GetSelectorFamily(dest.Selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			destAdapter, exists := laneRegistry.GetLaneAdapter(destFamily, lane.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", destFamily)
			}
			err = populateAddresses(&e, src, srcAdapter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error fetching address for src chain %d: %w", src.Selector, err)
			}
			err = populateAddresses(&e, dest, destAdapter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error fetching address for dest chain %d: %w", dest.Selector, err)
			}
			type lanePair struct {
				chainA  *ChainDefinition
				chainB  *ChainDefinition
				adapter LaneAdapter
			}
			configureLaneReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, srcAdapter.adapter.ConfigureLaneLegAsSource(), e.BlockChains, UpdateLanesInput{
				Source:       src,
				Dest:         dest,
				IsDisabled:   lane.IsDisabled,
				TestRouter:   lane.TestRouter,
				ExtraConfigs: lane.ExtraConfigs,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure token pool on chain with selector %d: %w", src.Selector, err)
			}
			batchOps = append(batchOps, configureLaneReport.Output.BatchOps...)
			reports = append(reports, configureLaneReport.ExecutionReports...)

			configureLaneReport, err = cldf_ops.ExecuteSequence(e.OperationsBundle, destAdapter.adapter.ConfigureLaneLegAsDest(), e.BlockChains, UpdateLanesInput{
				Source:       src,
				Dest:         dest,
				IsDisabled:   lane.IsDisabled,
				TestRouter:   lane.TestRouter,
				ExtraConfigs: lane.ExtraConfigs,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure lane leg as on chain with selector %d: %w", dest.Selector, err)
			}
			batchOps = append(batchOps, configureLaneReport.Output.BatchOps...)
			reports = append(reports, configureLaneReport.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}

func populateAddresses(e *cldf.Environment, chainDef *ChainDefinition, adapter LaneAdapter) error {
	var err error
	chainDef.OnRamp, err = adapter.GetOnRampAddress(e, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching onramp address for chain %d: %w", chainDef.Selector, err)
	}
	chainDef.OffRamp, err = adapter.GetOffRampAddress(e, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching offramp address for chain %d: %w", chainDef.Selector, err)
	}
	chainDef.FeeQuoter, err = adapter.GetFQAddress(e, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching fee quoter address for chain %d: %w", chainDef.Selector, err)
	}
	chainDef.Router, err = adapter.GetRouterAddress(e, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching router address for chain %d: %w", chainDef.Selector, err)
	}
	return nil
}
