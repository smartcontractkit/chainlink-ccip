package lanes

import (
	"fmt"
	"time"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

const (
	DefaultValidUntil = 72 * time.Hour
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

// type ConnectChainsBidirectional struct{}

// func (cs ConnectChainsBidirectional) VerifyPreconditions(env cldf.Environment, cfg ConnectChainsConfig) error {
// 	// TODO: implement this
// 	return nil
// }

// func (cs ConnectChainsBidirectional) Apply(e cldf.Environment, cfg ConnectChainsConfig) (cldf.ChangesetOutput, error) {
// 	finalOutput := cldf.ChangesetOutput{}
// 	for i, lane := range cfg.Lanes {
// 		src, dest := lane.Source, lane.Dest
// 		srcFamily, err := chain_selectors.GetSelectorFamily(src.Selector)
// 		if err != nil {
// 			return cldf.ChangesetOutput{}, err
// 		}
// 		if _, exists := registeredChainAdapters[srcFamily]; !exists {
// 			return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", srcFamily)
// 		}
// 		destFamily, err := chain_selectors.GetSelectorFamily(dest.Selector)
// 		if err != nil {
// 			return cldf.ChangesetOutput{}, err
// 		}
// 		if _, exists := registeredChainAdapters[destFamily]; !exists {
// 			return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", destFamily)
// 		}
// 		srcOnRamp, err := registeredChainAdapters[srcFamily].GetOnRampAddress(e, src.Selector)
// 		if err != nil {
// 			return cldf.ChangesetOutput{}, fmt.Errorf("error fetching onramp address for src chain %d: %w", src.Selector, err)
// 		}
// 		src.OnRamp = srcOnRamp
// 		// coalesce src -> dest
// 		output, err := registeredChainAdapters[srcFamily].ConfigureLaneAsSourceAndDest(e, UpdateLanesInput{
// 			Source:       src,
// 			Dest:         dest,
// 			IsDisabled:   lane.IsDisabled,
// 			TestRouter:   lane.TestRouter,
// 			ExtraConfigs: lane.ExtraConfigs,
// 			MCMS:         cfg.MCMS,
// 		})
// 		if err != nil {
// 			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
// 			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to apply changeset at index %d: %w", i, err)
// 		}
// 		err = MergeChangesetOutput(e, &finalOutput, output)
// 		if err != nil {
// 			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
// 			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to merge output of changeset at index %d: %w", i, err)
// 		}
// 		// coalesce dest -> src
// 		output, err = registeredChainAdapters[destFamily].ConfigureLaneAsSourceAndDest(e, UpdateLanesInput{
// 			Source:       dest,
// 			Dest:         src,
// 			IsDisabled:   lane.IsDisabled,
// 			TestRouter:   lane.TestRouter,
// 			ExtraConfigs: lane.ExtraConfigs,
// 			MCMS:         cfg.MCMS,
// 		})
// 		if err != nil {
// 			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
// 			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to apply changeset at index %d: %w", i, err)
// 		}
// 		err = MergeChangesetOutput(e, &finalOutput, output)
// 		if err != nil {
// 			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
// 			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to merge output of changeset at index %d: %w", i, err)
// 		}
// 	}
// 	// Aggregate all Timelock proposals into 1 proposal
// 	proposal, err := AggregateProposals(
// 		e,
// 		finalOutput.MCMSTimelockProposals,
// 		"connect chains bidirectionally",
// 		cfg.MCMS,
// 	)
// 	if err != nil {
// 		return finalOutput, fmt.Errorf("failed to aggregate proposals: %w", err)
// 	}

// 	// If no proposal was created, we return the final output without a proposal
// 	if proposal == nil {
// 		return finalOutput, nil
// 	}

// 	// Reset proposals to only include the aggregated proposal
// 	finalOutput.MCMSTimelockProposals = []mcmslib.TimelockProposal{*proposal}
// 	return finalOutput, nil
// }

func makeApply(laneRegistry *LaneAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ConnectChainsConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg ConnectChainsConfig) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for _, lane := range cfg.Lanes {
			src, dest := lane.Source, lane.Dest
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
			srcOnRamp, err := srcAdapter.GetOnRampAddress(e, src.Selector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("error fetching onramp address for src chain %d: %w", src.Selector, err)
			}
			src.OnRamp = srcOnRamp
			type lanePair struct {
				chainA  ChainDefinition
				chainB  ChainDefinition
				adapter LaneAdapter
			}
			for _, pair := range []lanePair{
				{chainA: src, chainB: dest, adapter: srcAdapter},
				{chainA: dest, chainB: src, adapter: destAdapter},
			} {
				configureLaneReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, pair.adapter.ConfigureLaneLegAsSource(), e.BlockChains, UpdateLanesInput{
					Source:       pair.chainA,
					Dest:         pair.chainB,
					IsDisabled:   lane.IsDisabled,
					TestRouter:   lane.TestRouter,
					ExtraConfigs: lane.ExtraConfigs,
				})
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to configure token pool on chain with selector %d: %w", src.Selector, err)
				}
				batchOps = append(batchOps, configureLaneReport.Output.BatchOps...)
				reports = append(reports, configureLaneReport.ExecutionReports...)

				configureLaneReport, err = cldf_ops.ExecuteSequence(e.OperationsBundle, pair.adapter.ConfigureLaneLegAsDest(), e.BlockChains, UpdateLanesInput{
					Source:       pair.chainA,
					Dest:         pair.chainB,
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
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
