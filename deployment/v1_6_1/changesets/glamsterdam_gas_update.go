package changesets

import (
	"fmt"
	"sort"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	v1_6_1_adapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_6_1/adapters"
	glamsterdam_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/glamsterdam"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// GlamsterdamGasUpdateV16Cfg is the config for the v1.6.1 Glamsterdam gas update changeset.
type GlamsterdamGasUpdateV16Cfg struct {
	// TargetChainSelector is the Glamsterdam target chain
	TargetChainSelector uint64
	// SkipChainSelectors are chains to skip
	SkipChainSelectors []uint64
}

// UpdateGasConfigForGlamsterdamV16 returns the v1.6.1 Glamsterdam gas update changeset.
//
// adapterRegistry is injected rather than looked up from the GetGasUpdateAdapterRegistry global
// so the changeset is testable with mocks (see ConfigureChainsForLanesFromTopology's
// ChainFamilyRegistry for the established convention this follows). Candidate chains are
// discovered across every chain family present in the environment, grouped by family, and
// dispatched to that family's registered adapter — a family with no registered adapter fails the
// changeset rather than silently skipping chains from that family.
func UpdateGasConfigForGlamsterdamV16(mcmsRegistry *changesets.MCMSReaderRegistry, adapterRegistry *v1_6_1_adapters.GasUpdateAdapterRegistry) deployment.ChangeSetV2[changesets.WithMCMS[GlamsterdamGasUpdateV16Cfg]] {
	validate := func(e deployment.Environment, cfg changesets.WithMCMS[GlamsterdamGasUpdateV16Cfg]) error {
		if adapterRegistry == nil {
			return fmt.Errorf("gas update adapter registry is required")
		}
		if cfg.Cfg.TargetChainSelector == 0 {
			return fmt.Errorf("TargetChainSelector must be set")
		}
		return nil
	}

	apply := func(e deployment.Environment, cfg changesets.WithMCMS[GlamsterdamGasUpdateV16Cfg]) (deployment.ChangesetOutput, error) {
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

		// Get all candidate chain selectors across every chain family (all chains except target
		// and skip list). Sorted so that identical inputs always produce operations/descriptions
		// in the same order, regardless of Go's randomized map iteration order.
		candidateChains := []uint64{}
		for _, sel := range e.BlockChains.ListChainSelectors() {
			if sel != cfg.Cfg.TargetChainSelector && !skipSet[sel] {
				candidateChains = append(candidateChains, sel)
			}
		}
		sort.Slice(candidateChains, func(i, j int) bool { return candidateChains[i] < candidateChains[j] })

		// Add skip list entries to report, sorted for the same reason.
		sortedSkips := make([]uint64, 0, len(skipSet))
		for sel := range skipSet {
			sortedSkips = append(sortedSkips, sel)
		}
		sort.Slice(sortedSkips, func(i, j int) bool { return sortedSkips[i] < sortedSkips[j] })
		for _, sel := range sortedSkips {
			report.AddLine(fmt.Sprintf("chain %d: skipped (explicit SkipChainSelectors entry)", sel))
		}

		// Group candidate chains by chain family, so each family's registered adapter only ever
		// sees its own chains. Sorted for deterministic dispatch order.
		byFamily := make(map[string][]uint64)
		for _, sel := range candidateChains {
			family, err := chain_selectors.GetSelectorFamily(sel)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain %d: %w", sel, err)
			}
			byFamily[family] = append(byFamily[family], sel)
		}
		families := make([]string, 0, len(byFamily))
		for family := range byFamily {
			families = append(families, family)
		}
		sort.Strings(families)

		for _, family := range families {
			adapter := adapterRegistry.GetGasUpdateAdapter(family)
			if adapter == nil {
				// A missing adapter means this migration cannot run at all for this chain family —
				// fail loudly rather than silently returning a "successful" no-op output, which would
				// let automation treat an unexecuted update as complete.
				return deployment.ChangesetOutput{}, fmt.Errorf("no gas update adapter registered for chain family %q", family)
			}

			seqOutput, err := v1_6_1_adapters.GlamsterdamGasUpdateSequence(
				e.OperationsBundle,
				e.BlockChains,
				e.DataStore,
				v1_6_1_adapters.GlamsterdamGasUpdateSequenceInput{
					Adapter:                 adapter,
					TargetChainSelector:     cfg.Cfg.TargetChainSelector,
					CandidateChainSelectors: byFamily[family],
					Report:                  report,
				},
			)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to execute glamsterdam gas update sequence for chain family %q: %w", family, err)
			}

			allBatchOps = append(allBatchOps, seqOutput.BatchOps...)
		}

		// Build the output with MCMS proposal
		mcmsInput := cfg.MCMS
		if mcmsInput.Description == "" {
			mcmsInput.Description = report.String()
		} else {
			mcmsInput.Description = mcmsInput.Description + "\n\n" + report.String()
		}
		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithBatchOps(allBatchOps).
			Build(mcmsInput)
	}

	return deployment.CreateChangeSet(apply, validate)
}
