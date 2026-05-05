package fees

import (
	"fmt"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// UpdateFeeQuoterDests creates a changeset that updates FeeQuoter destination chain configs
// with upsert semantics. For each destination, it reads the current on-chain config (or uses
// defaults if none exists), applies the user's override function, and writes the result.
func UpdateFeeQuoterDests() cldf.ChangeSetV2[UpdateFeeQuoterDestsInput] {
	feeRegistry := GetRegistry()
	mcmsRegistry := changesets.GetRegistry()
	return cldf.CreateChangeSet(makeFQDestsApply(feeRegistry, mcmsRegistry), makeFQDestsVerify())
}

func makeFQDestsVerify() func(cldf.Environment, UpdateFeeQuoterDestsInput) error {
	return func(_ cldf.Environment, cfg UpdateFeeQuoterDestsInput) error {
		seenSrc := utils.NewSet[uint64]()
		for i, src := range cfg.Args {
			if exists := seenSrc.Add(src.Selector); exists {
				return fmt.Errorf("duplicate src chain selector at args[%d]: %d", i, src.Selector)
			}

			seenDst := utils.NewSet[uint64]()
			for j, dst := range src.Settings {
				if src.Selector == dst.Selector {
					return fmt.Errorf("src and dst chain selectors cannot be the same at args[%d].settings[%d]: %d", i, j, src.Selector)
				}
				if exists := seenDst.Add(dst.Selector); exists {
					return fmt.Errorf("duplicate dst chain selector at args[%d].settings[%d] (src=%d): %d", i, j, src.Selector, dst.Selector)
				}
			}
		}

		return nil
	}
}

func makeFQDestsApply(feeRegistry *FeeAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateFeeQuoterDestsInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg UpdateFeeQuoterDestsInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		type DestsGroup struct {
			settings map[uint64]lanes.FeeQuoterDestChainConfig
			fqRefDS  datastore.AddressRef
			adapter  FeeAdapter
		}

		for _, src := range cfg.Args {
			srcFamily, err := chain_selectors.GetSelectorFamily(src.Selector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", src.Selector, err)
			}
			srcResolver, ok := feeRegistry.GetFeeResolver(srcFamily)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("no fee resolver found for chain family %s (src selector %d)", srcFamily, src.Selector)
			}

			// NOTE: we could have a pair A (src --> dst1) & a pair B (src --> dst2) where pair A has
			// an FeeQ with version v1.6.0 and pair B has an FeeQ with version v2.0.0. In these cases
			// we need to execute the fee update for pair A using the v1.6 adapter and the fee update
			// for pair B using the v2.0 adapter as the logic differs between versions. The map below
			// will be used to group updates by AddressRefKey so that we can execute them correctly.
			feeGroups := map[datastore.AddressRefKey]DestsGroup{}
			for _, dst := range src.Settings {
				// Version inference part 1: we use the router contract to infer the currently configured on ramp
				onRampRef, err := srcResolver.GetOnRampRef(e, src.Selector, dst.Selector)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get OnRamp address ref from Router for src %d and dst %d: %w", src.Selector, dst.Selector, err)
				}
				onRampAdp, ok := feeRegistry.GetFeeAdapter(srcFamily, onRampRef.Version)
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no fee adapter found for chain family %s and version %s", srcFamily, onRampRef.Version.String())
				}

				// Version inference part 2: we use the on ramp to get the currently configure fee contract (e.g. EVM2EVMOnRamp for v1.5.x and FeeQuoter for v1.6.x and v2.0.x)
				feeQuoterRef, err := onRampAdp.GetFeeContractRef(e, onRampRef, src.Selector, dst.Selector)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get fee contract ref for src %d and dst %d: %w", src.Selector, dst.Selector, err)
				}
				feeQuoterAdp, ok := feeRegistry.GetFeeAdapter(srcFamily, feeQuoterRef.Version)
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no fee adapter found for chain family %s and version %s", srcFamily, feeQuoterRef.Version.String())
				}

				// Version inference part 3: the fee quoter adapter is used to configure the dest chain config
				resolved, err := resolveDestChainConfig(feeQuoterAdp, e, feeQuoterRef, src.Selector, dst.Selector, dst.Override)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve dest chain config for src %d, dst %d: %w", src.Selector, dst.Selector, err)
				}

				// Operations are grouped by fee contract to ensure the correct bindings are used
				if _, exists := feeGroups[feeQuoterRef.Key()]; !exists {
					feeGroups[feeQuoterRef.Key()] = DestsGroup{
						settings: map[uint64]lanes.FeeQuoterDestChainConfig{},
						fqRefDS:  feeQuoterRef,
						adapter:  feeQuoterAdp,
					}
				}

				// Assign the settings for this dst to the appropriate group
				feeGroups[feeQuoterRef.Key()].settings[dst.Selector] = resolved
			}

			for _, group := range feeGroups {
				if len(group.settings) == 0 {
					continue
				}
				report, err := cldf_ops.ExecuteSequence(
					e.OperationsBundle,
					group.adapter.ApplyDestChainConfigUpdates(e, group.fqRefDS),
					e.BlockChains,
					ApplyDestChainConfigSequenceInput{
						Selector: src.Selector,
						Settings: group.settings,
					},
				)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to apply FQ dest chain config updates for src %d: %w", src.Selector, err)
				}
				batchOps = append(batchOps, report.Output.BatchOps...)
				reports = append(reports, report.ExecutionReports...)
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithBatchOps(batchOps).
			WithReports(reports).
			Build(cfg.MCMS)
	}
}

// resolveDestChainConfig reads on-chain state (or defaults) and applies the user's override.
func resolveDestChainConfig(adapter FeeAdapter, e cldf.Environment, fq datastore.AddressRef, src, dst uint64, override *lanes.FeeQuoterDestChainConfigOverride) (lanes.FeeQuoterDestChainConfig, error) {
	onchain, err := adapter.GetOnchainDestChainConfig(e, fq, src, dst)
	if err != nil {
		return lanes.FeeQuoterDestChainConfig{}, fmt.Errorf("failed to read on-chain dest chain config for src %d, dst %d: %w", src, dst, err)
	}

	var base lanes.FeeQuoterDestChainConfig
	if onchain.IsEnabled {
		base = onchain
	} else {
		base = adapter.GetDefaultDestChainConfig(src, dst)
	}

	if override != nil {
		(*override)(&base)
	}
	return base, nil
}
