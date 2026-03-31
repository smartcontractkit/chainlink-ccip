package fees

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
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
	return cldf.CreateChangeSet(makeFQDestsApply(feeRegistry, mcmsRegistry), makeFQDestsVerify(feeRegistry))
}

func makeFQDestsVerify(feeRegistry *FeeAdapterRegistry) func(cldf.Environment, UpdateFeeQuoterDestsInput) error {
	return func(_ cldf.Environment, cfg UpdateFeeQuoterDestsInput) error {
		if cfg.Version == nil {
			return fmt.Errorf("version is required")
		}

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

			srcFamily, err := chain_selectors.GetSelectorFamily(src.Selector)
			if err != nil {
				return fmt.Errorf("failed to get chain family for selector %d: %w", src.Selector, err)
			}

			if _, exists := feeRegistry.GetFeeAdapter(srcFamily, cfg.Version); !exists {
				return fmt.Errorf("no fee adapter found for chain family %s and version %s", srcFamily, cfg.Version.String())
			}
		}

		return nil
	}
}

// resolveDestChainConfig reads on-chain state (or defaults) and applies the user's override.
func resolveDestChainConfig(adapter FeeAdapter, e cldf.Environment, src, dst uint64, override *lanes.FeeQuoterDestChainConfigOverride) (lanes.FeeQuoterDestChainConfig, error) {
	onchain, err := adapter.GetOnchainDestChainConfig(e, src, dst)
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

type fqDestsGroup struct {
	adapter  FeeAdapter
	settings map[uint64]lanes.FeeQuoterDestChainConfig
}

func makeFQDestsApply(feeRegistry *FeeAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateFeeQuoterDestsInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg UpdateFeeQuoterDestsInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for _, src := range cfg.Args {
			srcFamily, err := chain_selectors.GetSelectorFamily(src.Selector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", src.Selector, err)
			}

			adapter, exists := feeRegistry.GetFeeAdapter(srcFamily, cfg.Version)
			if !exists {
				return cldf.ChangesetOutput{}, fmt.Errorf("no fee adapter found for chain family %s and version %s", srcFamily, cfg.Version.String())
			}

			versionGroups := map[string]fqDestsGroup{}

			for _, dst := range src.Settings {
				fqRef, err := adapter.GetFeeContractRef(e, src.Selector, dst.Selector)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get FQ contract ref for src %d, dst %d: %w", src.Selector, dst.Selector, err)
				}

				v := fqRef.Version
				lookupVersion := semver.MustParse(fmt.Sprintf("%d.%d.0", v.Major(), v.Minor()))

				updater, exists := feeRegistry.GetFeeAdapter(srcFamily, lookupVersion)
				if !exists {
					return cldf.ChangesetOutput{}, fmt.Errorf("no fee adapter found for chain family %s and version %s (detected from contract at %s)", srcFamily, lookupVersion.String(), fqRef.Address)
				}

				resolved, err := resolveDestChainConfig(updater, e, src.Selector, dst.Selector, dst.Override)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve dest chain config for src %d, dst %d: %w", src.Selector, dst.Selector, err)
				}

				versionKey := lookupVersion.String()
				if _, exists := versionGroups[versionKey]; !exists {
					versionGroups[versionKey] = fqDestsGroup{
						adapter:  updater,
						settings: map[uint64]lanes.FeeQuoterDestChainConfig{},
					}
				}

				versionGroups[versionKey].settings[dst.Selector] = resolved
			}

			for _, group := range versionGroups {
				if len(group.settings) == 0 {
					continue
				}

				report, err := cldf_ops.ExecuteSequence(
					e.OperationsBundle,
					group.adapter.ApplyDestChainConfigUpdates(e),
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
