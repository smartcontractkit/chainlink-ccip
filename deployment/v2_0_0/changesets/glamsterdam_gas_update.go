package changesets

import (
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	v2_0_0_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
	glamsterdam_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/glamsterdam"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// GlamsterdamGasUpdateV200Cfg is the config for the v2.0.0 Glamsterdam gas update changeset.
type GlamsterdamGasUpdateV200Cfg struct {
	// TargetChainSelector is the Glamsterdam target chain
	TargetChainSelector uint64
	// SkipChainSelectors are chains to skip
	SkipChainSelectors []uint64
}

// UpdateGasConfigForGlamsterdamV200 returns the v2.0.0 Glamsterdam gas update changeset.
func UpdateGasConfigForGlamsterdamV200(registry *changesets.MCMSReaderRegistry) deployment.ChangeSetV2[changesets.WithMCMS[GlamsterdamGasUpdateV200Cfg]] {
	validate := func(e deployment.Environment, cfg changesets.WithMCMS[GlamsterdamGasUpdateV200Cfg]) error {
		if cfg.Cfg.TargetChainSelector == 0 {
			return fmt.Errorf("TargetChainSelector must be set")
		}
		return nil
	}

	apply := func(e deployment.Environment, cfg changesets.WithMCMS[GlamsterdamGasUpdateV200Cfg]) (deployment.ChangesetOutput, error) {
		if err := validate(e, cfg); err != nil {
			return deployment.ChangesetOutput{}, err
		}

		// Collect all batch ops and report
		var allBatchOps []mcms_types.BatchOperation
		report := glamsterdam_utils.NewReport()

		// Build skip set for quick lookup
		skipSet := make(map[uint64]bool)
		for _, sel := range cfg.Cfg.SkipChainSelectors {
			skipSet[sel] = true
		}

		// Get all candidate chain selectors (all chains except target and skip list)
		candidateChains := []uint64{}
		for sel := range e.BlockChains.EVMChains() {
			if sel != cfg.Cfg.TargetChainSelector && !skipSet[sel] {
				candidateChains = append(candidateChains, sel)
			}
		}

		// Add skip list entries to report
		for sel := range skipSet {
			report.AddLine(fmt.Sprintf("chain %d: skipped (explicit SkipChainSelectors entry)", sel))
		}

		// Get the EVM adapter from the registry
		adapterRegistry := v2_0_0_adapters.GetGasUpdateAdapterRegistry()
		adapter := adapterRegistry.GetGasUpdateAdapter(chain_selectors.FamilyEVM)
		if adapter == nil {
			// Gracefully skip if no adapter is registered for EVM (shouldn't happen in practice)
			report.AddLine(fmt.Sprintf("EVM family: no gas update adapter registered, skipped"))
			return changesets.NewOutputBuilder(e, registry).
				WithBatchOps(allBatchOps).
				Build(cfg.MCMS)
		}

		// Run the orchestration sequence for EVM chains
		seqOutput, err := v2_0_0_adapters.GlamsterdamGasUpdateSequence(
			e.OperationsBundle,
			e.BlockChains,
			e.DataStore,
			v2_0_0_adapters.GlamsterdamGasUpdateSequenceInput{
				Adapter:                 adapter,
				TargetChainSelector:     cfg.Cfg.TargetChainSelector,
				CandidateChainSelectors: candidateChains,
				Report:                  report,
			},
		)
		if err != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to execute glamsterdam gas update sequence: %w", err)
		}

		// Merge batch ops
		allBatchOps = append(allBatchOps, seqOutput.BatchOps...)

		// Build the output with MCMS proposal
		return changesets.NewOutputBuilder(e, registry).
			WithBatchOps(allBatchOps).
			Build(cfg.MCMS)
	}

	return deployment.CreateChangeSet(apply, validate)
}
