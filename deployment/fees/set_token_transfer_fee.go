package fees

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type TokenTransferFee struct {
	FeeArgs UnresolvedTokenTransferFeeArgs `json:"feeArgs" yaml:"feeArgs"`
	Address string                         `json:"address" yaml:"address"`
	IsReset bool                           `json:"isReset" yaml:"isReset"`
}

type TokenTransferFeeForDst struct {
	Settings []TokenTransferFee `json:"settings" yaml:"settings"`
	Selector uint64             `json:"selector" yaml:"selector"`
}

type TokenTransferFeeForSrc struct {
	Settings []TokenTransferFeeForDst `json:"settings" yaml:"settings"`
	Selector uint64                   `json:"selector" yaml:"selector"`
}

type SetTokenTransferFeeInput struct {
	// Version represents the chain adapter version to use. Historically this was
	// equal to the version of the currently configured OnRamp, but we have since
	// modified this adapter such that it is now capable of inferring the correct
	// contract version purely from onchain data. Thus, specifying this field has
	// no effect, but we have left it here for backwards compatibility - or if we
	// want to implement a later feature that allows the user to override version
	// inference.
	Version *semver.Version          `json:"version" yaml:"version"`
	Args    []TokenTransferFeeForSrc `json:"args" yaml:"args"`
	MCMS    mcms.Input               `json:"mcms" yaml:"mcms"`
}

func SetTokenTransferFee() cldf.ChangeSetV2[SetTokenTransferFeeInput] {
	feeRegistry := GetRegistry()
	mcmsRegistry := changesets.GetRegistry()
	return cldf.CreateChangeSet(makeApply(feeRegistry, mcmsRegistry), makeVerify(feeRegistry, mcmsRegistry))
}

func makeVerify(_ *FeeAdapterRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, SetTokenTransferFeeInput) error {
	return func(_ cldf.Environment, cfg SetTokenTransferFeeInput) error {
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

				seenAddresses := utils.NewSet[string]()
				for k, entry := range dst.Settings {
					trimmed := strings.TrimSpace(entry.Address)
					if trimmed == "" {
						return fmt.Errorf("empty token address at args[%d].settings[%d].settings[%d] (src=%d,dst=%d)", i, j, k, src.Selector, dst.Selector)
					}
					if exists := seenAddresses.Add(trimmed); exists {
						return fmt.Errorf("duplicate token address at args[%d].settings[%d].settings[%d] (src=%d,dst=%d): %q", i, j, k, src.Selector, dst.Selector, trimmed)
					}
				}
			}
		}

		return nil
	}
}

func makeApply(feeRegistry *FeeAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, SetTokenTransferFeeInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg SetTokenTransferFeeInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		type FeeGroup struct {
			settings map[uint64]map[string]*TokenTransferFeeArgs
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
			feeGroups := map[datastore.AddressRefKey]FeeGroup{}
			for _, dst := range src.Settings {
				// Version inference part 1: we use the router contract to infer the currently configured on ramp
				onRampRef, err := srcResolver.GetOnRampRef(e.OperationsBundle, e.BlockChains, e.DataStore, src.Selector, dst.Selector)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get OnRamp address ref from Router for src %d and dst %d: %w", src.Selector, dst.Selector, err)
				}
				onRampAdp, ok := feeRegistry.GetFeeAdapter(srcFamily, onRampRef.Version)
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no fee adapter found for chain family %s and version %s", srcFamily, onRampRef.Version.String())
				}

				// Version inference part 2: we use the on ramp to get the currently configure fee contract (e.g. EVM2EVMOnRamp for v1.5.x and FeeQuoter for v1.6.x and v2.0.x)
				feeQuoterRef, err := onRampAdp.GetFeeContractRef(e.OperationsBundle, e.BlockChains, e.DataStore, onRampRef, src.Selector, dst.Selector)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get fee contract ref for src %d and dst %d: %w", src.Selector, dst.Selector, err)
				}
				feeQuoterAdp, ok := feeRegistry.GetFeeAdapter(srcFamily, feeQuoterRef.Version)
				if !ok {
					return cldf.ChangesetOutput{}, fmt.Errorf("no fee adapter found for chain family %s and version %s", srcFamily, feeQuoterRef.Version.String())
				}

				// Version inference part 3: the fee quoter adapter is used to configure the fees
				dstSettings := map[string]*TokenTransferFeeArgs{}
				for _, feeCfg := range dst.Settings {
					args, shouldApply, err := inferTokenTransferFeeArgs(feeQuoterAdp, e, feeQuoterRef, src.Selector, dst.Selector, feeCfg)
					if err != nil {
						return cldf.ChangesetOutput{}, fmt.Errorf("failed to infer token transfer fee args for token %s: %w", feeCfg.Address, err)
					}
					if !shouldApply {
						continue
					}
					dstSettings[feeCfg.Address] = args
				}
				if len(dstSettings) == 0 {
					continue
				}

				// Operations are grouped by fee contract to ensure the correct bindings are used
				if _, exists := feeGroups[feeQuoterRef.Key()]; !exists {
					feeGroups[feeQuoterRef.Key()] = FeeGroup{
						settings: map[uint64]map[string]*TokenTransferFeeArgs{},
						fqRefDS:  feeQuoterRef,
						adapter:  feeQuoterAdp,
					}
				}

				// Assign the settings for this dst to the appropriate group
				feeGroups[feeQuoterRef.Key()].settings[dst.Selector] = dstSettings
			}

			for _, group := range feeGroups {
				if len(group.settings) == 0 {
					continue
				}
				report, err := cldf_ops.ExecuteSequence(
					e.OperationsBundle,
					group.adapter.SetTokenTransferFee(e.DataStore, group.fqRefDS),
					e.BlockChains,
					SetTokenTransferFeeSequenceInput{
						Selector: src.Selector,
						Settings: group.settings,
					},
				)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to set token transfer fee config for selector %d: %w", src.Selector, err)
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

func inferTokenTransferFeeArgs(adapter FeeAdapter, e cldf.Environment, fq datastore.AddressRef, src uint64, dst uint64, cfg TokenTransferFee) (*TokenTransferFeeArgs, bool, error) {
	e.Logger.Infof("Inferring token transfer fee config for src %d, dst %d, and token %s", src, dst, cfg.Address)
	onchainCfg, err := adapter.GetOnchainTokenTransferFeeConfig(e.OperationsBundle, e.BlockChains, fq, src, dst, cfg.Address)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get on-chain token transfer fee config for src %d, dst %d, and token %s: %w", src, dst, cfg.Address, err)
	}

	if cfg.IsReset {
		if !onchainCfg.IsEnabled {
			e.Logger.Infof("Token transfer fee config for src %d, dst %d, and token %s is already disabled on-chain; skipping reset", src, dst, cfg.Address)
			return nil, false, nil
		}

		e.Logger.Infof("Reset requested for token transfer fee config for src %d, dst %d, and token %s", src, dst, cfg.Address)
		return nil, true, nil
	}

	var fallbacks TokenTransferFeeArgs
	if onchainCfg.IsEnabled {
		fallbacks = onchainCfg
		e.Logger.Infof("Token transfer fee config for src %d, dst %d, and token %s is already set on-chain; using on-chain values as defaults: %+v", src, dst, cfg.Address, fallbacks)
	} else {
		fallbacks = adapter.GetDefaultTokenTransferFeeConfig(src, dst)
		e.Logger.Infof("Token transfer fee config for src %d, dst %d, and token %s is not set on-chain; using adapter defaults: %+v", src, dst, cfg.Address, fallbacks)
	}

	resolved := cfg.FeeArgs.Resolve(fallbacks)
	if *resolved == onchainCfg {
		e.Logger.Infof("Token transfer fee config for src %d, dst %d, and token %s already matches on-chain config; skipping update", src, dst, cfg.Address)
		return nil, false, nil
	}

	return resolved, true, nil
}
