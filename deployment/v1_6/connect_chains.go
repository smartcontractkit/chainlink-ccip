package v1_6

import (
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	mcmslib "github.com/smartcontractkit/mcms"
	"github.com/smartcontractkit/mcms/types"
)

const (
	DefaultValidUntil = 72 * time.Hour
)

type ConnectChainsUnidirectional struct{}

func (cs ConnectChainsUnidirectional) VerifyPreconditions(env cldf.Environment, cfg ConnectChainsConfig) error {
	// TODO: implement this
	return nil
}

func (cs ConnectChainsUnidirectional) Apply(e cldf.Environment, cfg ConnectChainsConfig) (cldf.ChangesetOutput, error) {
	finalOutput := cldf.ChangesetOutput{}
	for i, lane := range cfg.Lanes {
		src, dest := lane.Source, lane.Dest
		srcFamily, err := chain_selectors.GetSelectorFamily(src.Selector)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		if _, exists := registeredChainAdapters[srcFamily]; !exists {
			return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", srcFamily)
		}
		destFamily, err := chain_selectors.GetSelectorFamily(dest.Selector)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		if _, exists := registeredChainAdapters[destFamily]; !exists {
			return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", destFamily)
		}
		srcOnRamp, err := registeredChainAdapters[srcFamily].GetOnRampAddress(e, src.Selector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("error fetching onramp address for src chain %d: %w", src.Selector, err)
		}
		src.OnRamp = srcOnRamp
		// coalesce src
		output, err := registeredChainAdapters[srcFamily].ConfigureLaneLegAsSource(e, UpdateLanesInput{
			Source:       src,
			Dest:         dest,
			IsDisabled:   lane.IsDisabled,
			TestRouter:   lane.TestRouter,
			ExtraConfigs: lane.ExtraConfigs,
			MCMS:         cfg.MCMS,
		})
		if err != nil {
			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to apply changeset at index %d: %w", i, err)
		}
		err = MergeChangesetOutput(e, &finalOutput, output)
		if err != nil {
			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to merge output of changeset at index %d: %w", i, err)
		}
		// coalesce dest
		output, err = registeredChainAdapters[destFamily].ConfigureLaneLegAsDest(e, UpdateLanesInput{
			Source:       src,
			Dest:         dest,
			IsDisabled:   lane.IsDisabled,
			TestRouter:   lane.TestRouter,
			ExtraConfigs: lane.ExtraConfigs,
			MCMS:         cfg.MCMS,
		})
		if err != nil {
			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to apply changeset at index %d: %w", i, err)
		}
		err = MergeChangesetOutput(e, &finalOutput, output)
		if err != nil {
			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to merge output of changeset at index %d: %w", i, err)
		}
	}
	// Aggregate all Timelock proposals into 1 proposal
	proposal, err := AggregateProposals(
		e,
		finalOutput.MCMSTimelockProposals,
		"connect chains unidirectionally",
		cfg.MCMS,
	)
	if err != nil {
		return finalOutput, fmt.Errorf("failed to aggregate proposals: %w", err)
	}

	// If no proposal was created, we return the final output without a proposal
	if proposal == nil {
		return finalOutput, nil
	}

	// Reset proposals to only include the aggregated proposal
	finalOutput.MCMSTimelockProposals = []mcmslib.TimelockProposal{*proposal}
	return finalOutput, nil
}

type ConnectChainsBidirectional struct{}

func (cs ConnectChainsBidirectional) VerifyPreconditions(env cldf.Environment, cfg ConnectChainsConfig) error {
	// TODO: implement this
	return nil
}

func (cs ConnectChainsBidirectional) Apply(e cldf.Environment, cfg ConnectChainsConfig) (cldf.ChangesetOutput, error) {
	finalOutput := cldf.ChangesetOutput{}
	for i, lane := range cfg.Lanes {
		src, dest := lane.Source, lane.Dest
		srcFamily, err := chain_selectors.GetSelectorFamily(src.Selector)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		if _, exists := registeredChainAdapters[srcFamily]; !exists {
			return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", srcFamily)
		}
		destFamily, err := chain_selectors.GetSelectorFamily(dest.Selector)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}
		if _, exists := registeredChainAdapters[destFamily]; !exists {
			return cldf.ChangesetOutput{}, fmt.Errorf("no ChainAdapter registered for chain family '%s'", destFamily)
		}
		srcOnRamp, err := registeredChainAdapters[srcFamily].GetOnRampAddress(e, src.Selector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("error fetching onramp address for src chain %d: %w", src.Selector, err)
		}
		src.OnRamp = srcOnRamp
		// coalesce src -> dest
		output, err := registeredChainAdapters[srcFamily].ConfigureLaneAsSourceAndDest(e, UpdateLanesInput{
			Source:       src,
			Dest:         dest,
			IsDisabled:   lane.IsDisabled,
			TestRouter:   lane.TestRouter,
			ExtraConfigs: lane.ExtraConfigs,
			MCMS:         cfg.MCMS,
		})
		if err != nil {
			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to apply changeset at index %d: %w", i, err)
		}
		err = MergeChangesetOutput(e, &finalOutput, output)
		if err != nil {
			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to merge output of changeset at index %d: %w", i, err)
		}
		// coalesce dest -> src
		output, err = registeredChainAdapters[destFamily].ConfigureLaneAsSourceAndDest(e, UpdateLanesInput{
			Source:       dest,
			Dest:         src,
			IsDisabled:   lane.IsDisabled,
			TestRouter:   lane.TestRouter,
			ExtraConfigs: lane.ExtraConfigs,
			MCMS:         cfg.MCMS,
		})
		if err != nil {
			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to apply changeset at index %d: %w", i, err)
		}
		err = MergeChangesetOutput(e, &finalOutput, output)
		if err != nil {
			finalOutput.Reports = append(finalOutput.Reports, output.Reports...)
			return cldf.ChangesetOutput{Reports: finalOutput.Reports}, fmt.Errorf("failed to merge output of changeset at index %d: %w", i, err)
		}
	}
	// Aggregate all Timelock proposals into 1 proposal
	proposal, err := AggregateProposals(
		e,
		finalOutput.MCMSTimelockProposals,
		"connect chains bidirectionally",
		cfg.MCMS,
	)
	if err != nil {
		return finalOutput, fmt.Errorf("failed to aggregate proposals: %w", err)
	}

	// If no proposal was created, we return the final output without a proposal
	if proposal == nil {
		return finalOutput, nil
	}

	// Reset proposals to only include the aggregated proposal
	finalOutput.MCMSTimelockProposals = []mcmslib.TimelockProposal{*proposal}
	return finalOutput, nil
}

// copied from core
// move all of this to utils post-PoC
func MergeChangesetOutput(e cldf.Environment, dest *cldf.ChangesetOutput, src cldf.ChangesetOutput) error {
	err := cldf.MergeChangesetOutput(e, dest, src)
	if err != nil {
		return fmt.Errorf("failed to merge changeset output: %w", err)
	}

	// The following merges are not included in cldf.MergeChangesetOutput
	// TODO @ccip-tooling: Open PR in chainlink-deployments-framework to include these merges
	// 1. Merge DataStores
	if dest.DataStore == nil {
		dest.DataStore = src.DataStore
	} else if src.DataStore != nil {
		err := dest.DataStore.Merge(src.DataStore.Seal())
		if err != nil {
			return fmt.Errorf("failed to merge data store: %w", err)
		}
	}
	// 2. Merge Reports
	if dest.Reports == nil {
		dest.Reports = src.Reports
	} else if src.Reports != nil {
		dest.Reports = append(dest.Reports, src.Reports...)
	}

	return nil
}

func AggregateProposals(
	e cldf.Environment,
	proposals []mcmslib.TimelockProposal,
	description string,
	mcmsConfig *utils.MCMSInput,
) (*mcmslib.TimelockProposal, error) {
	var batches []types.BatchOperation

	// Add proposals to the aggregate.
	for _, proposal := range proposals {
		batches = append(batches, proposal.Operations...)
	}

	// Return early if there are no operations.
	if len(batches) == 0 {
		return nil, nil
	}

	chains := mapset.NewSet[uint64]()
	for _, op := range batches {
		chains.Add(uint64(op.ChainSelector))
	}
	tlsPerChainID := make(map[types.ChainSelector]string)
	metadataPerChain := make(map[types.ChainSelector]types.ChainMetadata)
	for _, chain := range chains.ToSlice() {
		family, err := chain_selectors.GetSelectorFamily(chain)
		if err != nil {
			return nil, fmt.Errorf("error getting family for chain %d: %w", chain, err)
		}

		tl, err := registeredChainAdapters[family].GetTimelockAddress(e, chain)
		if err != nil {
			return nil, fmt.Errorf("error getting timelock address for chain %d: %w", chain, err)
		}
		tlsPerChainID[types.ChainSelector(chain)] = tl

		metadata, err := registeredChainAdapters[family].GetMCMSMetadata(e, chain, mcmsConfig.TimelockAction)
		if err != nil {
			return nil, fmt.Errorf("error getting mcms metadata for chain %d: %w", chain, err)
		}
		metadataPerChain[types.ChainSelector(chain)] = metadata
	}

	validUntil := mcmsConfig.ValidUntil
	if validUntil <= 0 {
		//nolint:gosec // G115
		validUntil = uint32(time.Now().Unix() + int64(DefaultValidUntil.Seconds()))
	}

	builder := mcmslib.NewTimelockProposalBuilder()
	builder.
		SetVersion("v1").
		SetAction(mcmsConfig.TimelockAction).
		SetValidUntil(uint32(validUntil)).
		SetDescription(description).
		SetDelay(mcmsConfig.TimelockDelay).
		SetOverridePreviousRoot(mcmsConfig.OverridePreviousRoot).
		SetChainMetadata(metadataPerChain).
		SetTimelockAddresses(tlsPerChainID).
		SetOperations(batches)

	build, err := builder.Build()
	if err != nil {
		return nil, err
	}
	return build, nil
}
