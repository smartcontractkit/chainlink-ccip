package fees

import (
	"errors"
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
	// Deprecated: ignored for chain families that have a registered
	// FeeContractResolver; the fee contract version is inferred per-lane.
	// Required for families that fall back to FeeAdapter.GetFeeContractRef.
	Version *semver.Version          `json:"version" yaml:"version"`
	Args    []TokenTransferFeeForSrc `json:"args" yaml:"args"`
	MCMS    mcms.Input               `json:"mcms" yaml:"mcms"`
}

func SetTokenTransferFee() cldf.ChangeSetV2[SetTokenTransferFeeInput] {
	feeRegistry := GetRegistry()
	resolverRegistry := GetFeeContractResolverRegistry()
	mcmsRegistry := changesets.GetRegistry()
	return cldf.CreateChangeSet(makeApply(feeRegistry, resolverRegistry, mcmsRegistry), makeVerify(feeRegistry, resolverRegistry, mcmsRegistry))
}

func makeVerify(_ *FeeAdapterRegistry, resolverRegistry *FeeContractResolverRegistry, _ *changesets.MCMSReaderRegistry) func(cldf.Environment, SetTokenTransferFeeInput) error {
	return func(_ cldf.Environment, cfg SetTokenTransferFeeInput) error {
		seenSrc := utils.NewSet[uint64]()
		for i, src := range cfg.Args {
			if exists := seenSrc.Add(src.Selector); exists {
				return fmt.Errorf("duplicate src chain selector at args[%d]: %d", i, src.Selector)
			}

			// Fail fast: if the src family has no registered FeeContractResolver,
			// the apply will need cfg.Version for the legacy fallback path.
			// Surface that requirement at verify rather than partway through apply.
			if cfg.Version == nil {
				family, err := chain_selectors.GetSelectorFamily(src.Selector)
				if err != nil {
					return fmt.Errorf("failed to get chain family for selector %d at args[%d]: %w", src.Selector, i, err)
				}
				if _, ok := resolverRegistry.GetFeeContractResolver(family); !ok {
					return fmt.Errorf("Version is required because chain family %s (src=%d at args[%d]) has no registered FeeContractResolver and the legacy fallback path needs it", family, src.Selector, i)
				}
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

func makeApply(feeRegistry *FeeAdapterRegistry, resolverRegistry *FeeContractResolverRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, SetTokenTransferFeeInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg SetTokenTransferFeeInput) (cldf.ChangesetOutput, error) {
		// Warn per-source family that has a registered resolver, so a mixed
		// input still surfaces the deprecation for migrated families instead
		// of being silenced by the unmigrated one.
		if cfg.Version != nil {
			for _, src := range cfg.Args {
				family, err := chain_selectors.GetSelectorFamily(src.Selector)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", src.Selector, err)
				}
				if _, ok := resolverRegistry.GetFeeContractResolver(family); ok {
					e.Logger.Warnf("SetTokenTransferFeeInput.Version is deprecated for chain family %s and ignored; the fee contract version is inferred per-lane from Router.GetOnRamp() (got %s for src %d)", family, cfg.Version.String(), src.Selector)
				}
			}
		}

		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		type FeeGroup struct {
			adapter  FeeAdapter
			settings map[uint64]map[string]*TokenTransferFeeArgs
		}

		for _, src := range cfg.Args {
			srcFamily, err := chain_selectors.GetSelectorFamily(src.Selector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", src.Selector, err)
			}

			resolver, hasResolver := resolverRegistry.GetFeeContractResolver(srcFamily)

			// legacyAdapter is selected lazily — either eagerly when no resolver
			// is registered for the family, or on demand when the resolver
			// returns ErrNoLiveLane (lane not yet wired in the Router) and we
			// need to fall back to the user-supplied Version path.
			var legacyAdapter FeeAdapter
			ensureLegacyAdapter := func() error {
				if legacyAdapter != nil {
					return nil
				}
				if cfg.Version == nil {
					return fmt.Errorf("legacy fee-contract-ref discovery for chain family %s requires Version on the input", srcFamily)
				}
				a, ok := feeRegistry.GetFeeAdapter(srcFamily, cfg.Version)
				if !ok {
					return fmt.Errorf("no FeeAdapter found for chain family %s and version %s", srcFamily, cfg.Version.String())
				}
				legacyAdapter = a
				return nil
			}
			if !hasResolver {
				if err := ensureLegacyAdapter(); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("no FeeContractResolver registered for chain family %s: %w", srcFamily, err)
				}
			}

			// Build version-grouped settings: version -> settings map
			versionGroups := map[string]FeeGroup{}

			for _, dst := range src.Settings {

				var feeContractRef datastore.AddressRef
				if hasResolver {
					feeContractRef, err = resolver.ResolveFeeContractRef(e, src.Selector, dst.Selector)
					if errors.Is(err, ErrNoLiveLane) {
						// Lane not yet wired in the Router. Fall back to the
						// version-supplied legacy path for this (src, dst) only.
						if legacyErr := ensureLegacyAdapter(); legacyErr != nil {
							return cldf.ChangesetOutput{}, fmt.Errorf("Router has no live lane for src %d dst %d and the legacy fallback is unavailable: %w (resolver error: %v)", src.Selector, dst.Selector, legacyErr, err)
						}
						e.Logger.Infof("EVMFeeContractResolver: no live lane for src=%d dst=%d; falling back to FeeAdapter v%s", src.Selector, dst.Selector, cfg.Version.String())
						feeContractRef, err = legacyAdapter.GetFeeContractRef(e, src.Selector, dst.Selector)
					}
				} else {
					feeContractRef, err = legacyAdapter.GetFeeContractRef(e, src.Selector, dst.Selector)
				}
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get fee contract ref for src %d and dst %d: %w", src.Selector, dst.Selector, err)
				}
				if feeContractRef.Version == nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("fee contract ref for src %d and dst %d has nil Version; cannot select a fee adapter", src.Selector, dst.Selector)
				}

				// Only feeContractRef.Version is consumed; the address itself is not
				// threaded through SetTokenTransferFeeSequenceInput. See the
				// EVMFeeContractResolver doc comment (R7) for the datastore-vs-chain
				// invariant this depends on.
				lookupVersion := utils.StripPatchVersion(feeContractRef.Version)

				updater, exists := feeRegistry.GetFeeAdapter(srcFamily, lookupVersion)
				if !exists {
					return cldf.ChangesetOutput{}, fmt.Errorf("no fee adapter found for chain family %s and version %s (resolved from feeContractRef.Version=%s)", srcFamily, lookupVersion.String(), feeContractRef.Version.String())
				}

				dstSettings := map[string]*TokenTransferFeeArgs{}
				for _, feeCfg := range dst.Settings {
					args, shouldApply, err := inferTokenTransferFeeArgs(updater, e, src.Selector, dst.Selector, feeCfg)
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

				versionKey := lookupVersion.String()
				if _, exists := versionGroups[versionKey]; !exists {
					versionGroups[versionKey] = FeeGroup{
						adapter:  updater,
						settings: map[uint64]map[string]*TokenTransferFeeArgs{},
					}
				}

				// Process settings for this dst with its version's adapter
				versionGroups[versionKey].settings[dst.Selector] = dstSettings
			}

			// Execute updates grouped by adapter version
			for _, group := range versionGroups {
				report, err := cldf_ops.ExecuteSequence(
					e.OperationsBundle,
					group.adapter.SetTokenTransferFee(e),
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

func inferTokenTransferFeeArgs(adapter FeeAdapter, e cldf.Environment, src uint64, dst uint64, cfg TokenTransferFee) (*TokenTransferFeeArgs, bool, error) {
	e.Logger.Infof("Inferring token transfer fee config for src %d, dst %d, and token %s", src, dst, cfg.Address)
	onchainCfg, err := adapter.GetOnchainTokenTransferFeeConfig(e, src, dst, cfg.Address)
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
