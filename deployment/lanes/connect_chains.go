package lanes

import (
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
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
			chainA, chainB := &lane.ChainA, &lane.ChainB
			chainAFamily, err := chain_selectors.GetSelectorFamily(chainA.Selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			chainAAdapter, exists := laneRegistry.GetLaneAdapter(chainAFamily, lane.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", chainAFamily)
			}
			chainBFamily, err := chain_selectors.GetSelectorFamily(chainB.Selector)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			chainBAdapter, exists := laneRegistry.GetLaneAdapter(chainBFamily, lane.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", chainBFamily)
			}
			err = populateAddresses(e.DataStore, chainA, chainAAdapter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error fetching address for src chain %d: %w", chainA.Selector, err)
			}
			err = populateAddresses(e.DataStore, chainB, chainBAdapter)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error fetching address for dest chain %d: %w", chainB.Selector, err)
			}
			type lanePair struct {
				src         *ChainDefinition
				dest        *ChainDefinition
				srcAdapter  LaneAdapter
				destAdapter LaneAdapter
			}
			for _, pair := range []lanePair{
				{src: chainA, dest: chainB, srcAdapter: chainAAdapter, destAdapter: chainBAdapter},
				{src: chainB, dest: chainA, srcAdapter: chainBAdapter, destAdapter: chainAAdapter},
			} {
				configureLaneReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, pair.srcAdapter.ConfigureLaneLegAsSource(), e.BlockChains, UpdateLanesInput{
					Source:       pair.src,
					Dest:         pair.dest,
					IsDisabled:   lane.IsDisabled,
					TestRouter:   lane.TestRouter,
					ExtraConfigs: lane.ExtraConfigs,
				})
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to lane leg as source with selector %d: %w", pair.src.Selector, err)
				}
				batchOps = append(batchOps, configureLaneReport.Output.BatchOps...)
				reports = append(reports, configureLaneReport.ExecutionReports...)

				configureLaneReport, err = cldf_ops.ExecuteSequence(e.OperationsBundle, pair.destAdapter.ConfigureLaneLegAsDest(), e.BlockChains, UpdateLanesInput{
					Source:       pair.src,
					Dest:         pair.dest,
					IsDisabled:   lane.IsDisabled,
					TestRouter:   lane.TestRouter,
					ExtraConfigs: lane.ExtraConfigs,
				})
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure lane leg as dest with selector %d: %w", pair.dest.Selector, err)
				}
				batchOps = append(batchOps, configureLaneReport.Output.BatchOps...)
				reports = append(reports, configureLaneReport.ExecutionReports...)
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}

func populateAddresses(ds datastore.DataStore, chainDef *ChainDefinition, adapter LaneAdapter) error {
	var err error
	chainDef.OnRamp, err = adapter.GetOnRampAddress(ds, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching onramp address for chain %d: %w", chainDef.Selector, err)
	}
	chainDef.OffRamp, err = adapter.GetOffRampAddress(ds, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching offramp address for chain %d: %w", chainDef.Selector, err)
	}
	chainDef.FeeQuoter, err = adapter.GetFQAddress(ds, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching fee quoter address for chain %d: %w", chainDef.Selector, err)
	}
	chainDef.Router, err = adapter.GetRouterAddress(ds, chainDef.Selector)
	if err != nil {
		return fmt.Errorf("error fetching router address for chain %d: %w", chainDef.Selector, err)
	}
	return nil
}
