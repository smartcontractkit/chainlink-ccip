package fees

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
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

type SetTokenTransferFeeConfig struct {
	Version *semver.Version          `json:"version" yaml:"version"`
	Args    []TokenTransferFeeForSrc `json:"args" yaml:"args"`
	MCMS    mcms.Input               `json:"mcms" yaml:"mcms"`
}

func SetTokenTransferFee(feeRegistry *FeeAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[SetTokenTransferFeeConfig] {
	return cldf.CreateChangeSet(makeApply(feeRegistry, mcmsRegistry), makeVerify(feeRegistry, mcmsRegistry))
}

func makeVerify(_ *FeeAdapterRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, SetTokenTransferFeeConfig) error {
	return func(_ cldf.Environment, cfg SetTokenTransferFeeConfig) error {
		seenSrc := utils.NewSet[uint64]()
		for i, src := range cfg.Args {
			// Disallow duplicate src selectors
			if exists := seenSrc.Add(src.Selector); exists {
				return fmt.Errorf("duplicate src chain selector at args[%d]: %d", i, src.Selector)
			}

			seenDst := utils.NewSet[uint64]()
			for j, dst := range src.Settings {
				// Disallow cyclic src/dst selectors
				if src.Selector == dst.Selector {
					return fmt.Errorf("src and dst chain selectors cannot be the same at args[%d].settings[%d]: %d", i, j, src.Selector)
				}

				// Disallow duplicate dst selectors within the same src
				if exists := seenDst.Add(dst.Selector); exists {
					return fmt.Errorf("duplicate dst chain selector at args[%d].settings[%d] (src=%d): %d", i, j, src.Selector, dst.Selector)
				}

				// Duplicate tracking
				updateSet := utils.NewSet[string]() // Track addresses for IsReset == false
				resetsSet := utils.NewSet[string]() // Track addresses for IsReset == true
				for k, entry := range dst.Settings {
					// Disallow empty token address
					trimmed := strings.TrimSpace(entry.Address)
					if trimmed == "" {
						return fmt.Errorf("empty token address at args[%d].settings[%d].settings[%d] (src=%d,dst=%d)", i, j, k, src.Selector, dst.Selector)
					}

					// Track by update vs reset to catch overlaps and duplicates
					addr := entry.Address
					if entry.IsReset {
						if updateSet.Has(addr) {
							return fmt.Errorf("the same address cannot be referenced in both updates and resets (src=%d,dst=%d,addr=%q)", src.Selector, dst.Selector, addr)
						}
						if exists := resetsSet.Add(addr); exists {
							return fmt.Errorf("duplicate reset for address at (src=%d,dst=%d) args[%d].settings[%d].settings[%d]: %q", src.Selector, dst.Selector, i, j, k, addr)
						}
					} else {
						if resetsSet.Has(addr) {
							return fmt.Errorf("the same address cannot be referenced in both updates and resets (src=%d,dst=%d,addr=%q)", src.Selector, dst.Selector, addr)
						}
						if exists := updateSet.Add(addr); exists {
							return fmt.Errorf("duplicate update for address at (src=%d,dst=%d) args[%d].settings[%d].settings[%d]: %q", src.Selector, dst.Selector, i, j, k, addr)
						}
					}
				}
			}
		}

		return nil
	}
}

func makeApply(feeRegistry *FeeAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, SetTokenTransferFeeConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg SetTokenTransferFeeConfig) (cldf.ChangesetOutput, error) {
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

			settings := map[uint64]map[string]*TokenTransferFeeArgs{}
			for _, dst := range src.Settings {
				settings[dst.Selector] = map[string]*TokenTransferFeeArgs{}
				for _, feeCfg := range dst.Settings {
					if args, err := inferTokenTransferFeeArgs(adapter, e, src.Selector, dst.Selector, feeCfg); err != nil {
						return cldf.ChangesetOutput{}, fmt.Errorf("failed to infer token transfer fee args for address %s: %w", feeCfg.Address, err)
					} else {
						settings[dst.Selector][feeCfg.Address] = args
					}
				}
			}

			report, err := cldf_ops.ExecuteSequence(
				e.OperationsBundle,
				adapter.SetTokenTransferFeeConfig(e.DataStore, src.Selector),
				e.BlockChains,
				SetTokenTransferFeeConfigSequenceInput{
					Selector: src.Selector,
					Settings: settings,
				},
			)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to set token transfer fee config for selector %d: %w", src.Selector, err)
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

func inferTokenTransferFeeArgs(adapter FeeAdapter, e cldf.Environment, src uint64, dst uint64, cfg TokenTransferFee) (*TokenTransferFeeArgs, error) {
	if cfg.IsReset {
		e.Logger.Infof("Reset requested for token transfer fee config for src %d, dst %d, and address %s; skipping inference", src, dst, cfg.Address)
		return nil, nil
	}

	e.Logger.Infof("Inferring token transfer fee config for src %d, dst %d, and address %s", src, dst, cfg.Address)
	onchainCfg, err := adapter.GetTokenTransferFeeConfig(e, src, dst, cfg.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to get on-chain token transfer fee config for src %d, dst %d, and address %s: %w", src, dst, cfg.Address, err)
	}

	var fallbacks TokenTransferFeeArgs
	if onchainCfg.IsEnabled {
		fallbacks = onchainCfg
		e.Logger.Infof("Token transfer fee config for src %d, dst %d, and address %s is already set on-chain; using on-chain values as defaults: %+v", src, dst, cfg.Address, fallbacks)
	} else {
		fallbacks = adapter.GetTokenTransferFeeConfigDefaults(src, dst)
		e.Logger.Infof("Token transfer fee config for src %d, dst %d, and address %s is not set on-chain; using adapter defaults: %+v", src, dst, cfg.Address, fallbacks)
	}

	return cfg.FeeArgs.Infer(fallbacks), nil
}
