package tokens

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// UnresolvedTokenTransferFeeArgs allows for partial specification of token transfer fee configurations.
type UnresolvedTokenTransferFeeArgs struct {
	DefaultFinalityTransferFeeBps utils.Optional[uint16] `json:"defaultFinalityTransferFeeBps" yaml:"defaultFinalityTransferFeeBps"`
	CustomFinalityTransferFeeBps  utils.Optional[uint16] `json:"customFinalityTransferFeeBps" yaml:"customFinalityTransferFeeBps"`
	DefaultFinalityFeeUSDCents    utils.Optional[uint32] `json:"defaultFinalityFeeUSDCents" yaml:"defaultFinalityFeeUSDCents"`
	CustomFinalityFeeUSDCents     utils.Optional[uint32] `json:"customFinalityFeeUSDCents" yaml:"customFinalityFeeUSDCents"`
	DestBytesOverhead             utils.Optional[uint32] `json:"destBytesOverhead" yaml:"destBytesOverhead"`
	DestGasOverhead               utils.Optional[uint32] `json:"destGasOverhead" yaml:"destGasOverhead"`
	IsEnabled                     utils.Optional[bool]   `json:"isEnabled" yaml:"isEnabled"`
}

// Infer fills in any unset fields in the unresolved configuration using the provided fallback values.
func (cfg UnresolvedTokenTransferFeeArgs) Infer(fallbacks TokenTransferFeeConfig) *TokenTransferFeeConfig {
	return &TokenTransferFeeConfig{
		DefaultFinalityTransferFeeBps: cfg.DefaultFinalityTransferFeeBps.Infer(fallbacks.DefaultFinalityTransferFeeBps),
		CustomFinalityTransferFeeBps:  cfg.CustomFinalityTransferFeeBps.Infer(fallbacks.CustomFinalityTransferFeeBps),
		DefaultFinalityFeeUSDCents:    cfg.DefaultFinalityFeeUSDCents.Infer(fallbacks.DefaultFinalityFeeUSDCents),
		CustomFinalityFeeUSDCents:     cfg.CustomFinalityFeeUSDCents.Infer(fallbacks.CustomFinalityFeeUSDCents),
		DestBytesOverhead:             cfg.DestBytesOverhead.Infer(fallbacks.DestBytesOverhead),
		DestGasOverhead:               cfg.DestGasOverhead.Infer(fallbacks.DestGasOverhead),
		IsEnabled:                     cfg.IsEnabled.Infer(fallbacks.IsEnabled),
	}
}

// TokenTransferFeeConfig defines the standardized configuration for token transfer fees for all chain families.
type TokenTransferFeeForDst struct {
	Settings UnresolvedTokenTransferFeeArgs `json:"settings" yaml:"settings"`
	Selector uint64                         `json:"selector" yaml:"selector"`
	IsReset  bool                           `json:"isReset" yaml:"isReset"`
}

// TokenTransferFeeForPool organizes token transfer fee configurations by pool address, then by destination
// chain selector. This allows the user to set multiple destination chain configurations for the same token
// pool address without repeating the pool address for each one.
type TokenTransferFeeForPool struct {
	MinBlockConfirmations utils.Optional[uint16]   `json:"minBlockConfirmations" yaml:"minBlockConfirmations"`
	Destinations          []TokenTransferFeeForDst `json:"destinations" yaml:"destinations"`
	PoolAddress           string                   `json:"poolAddress" yaml:"poolAddress"`
}

// TokenTransferFeeForSrc organizes token transfer fee configurations by source chain selector, then by pool
// address, then by destination chain selector. This allows the user to set multiple pool configurations for
// a source chain without repeating the source chain selector for each one.
type TokenTransferFeeForSrc struct {
	TokenPools []TokenTransferFeeForPool `json:"tokenPools" yaml:"tokenPools"`
	Selector   uint64                    `json:"selector" yaml:"selector"`
}

// SetTokenTransferFeeInput defines the input for the SetTokenTransferFee change set, which allows users to set
// token transfer fee configurations for multiple source chains, with multiple pools, and multiple destinations
// for each pool.
type SetTokenTransferFeeInput struct {
	Version *semver.Version          `json:"version" yaml:"version"`
	Args    []TokenTransferFeeForSrc `json:"args" yaml:"args"`
	MCMS    mcms.Input               `json:"mcms" yaml:"mcms"`
}

func SetTokenTransferFee() deployment.ChangeSetV2[SetTokenTransferFeeInput] {
	return deployment.CreateChangeSet(setTokenTransferFeeApply(), setTokenTransferFeeVerify())
}

func setTokenTransferFeeVerify() func(deployment.Environment, SetTokenTransferFeeInput) error {
	return func(_ deployment.Environment, cfg SetTokenTransferFeeInput) error {
		seenSrc := utils.NewSet[uint64]()
		for i, src := range cfg.Args {
			if exists := seenSrc.Add(src.Selector); exists {
				return fmt.Errorf("duplicate src chain selector at args[%d]: %d", i, src.Selector)
			}

			seenPools := utils.NewSet[string]()
			for j, pool := range src.TokenPools {
				trimmed := strings.TrimSpace(pool.PoolAddress)
				if trimmed == "" {
					return fmt.Errorf("empty pool address at args[%d].settings[%d] (src=%d)", i, j, src.Selector)
				}
				if exists := seenPools.Add(pool.PoolAddress); exists {
					return fmt.Errorf("duplicate pool address at args[%d].settings[%d] (src=%d): %s", i, j, src.Selector, pool.PoolAddress)
				}

				updateSet := utils.NewSet[string]()
				resetsSet := utils.NewSet[string]()
				seenDests := utils.NewSet[uint64]()
				for k, dst := range pool.Destinations {
					if exists := seenDests.Add(dst.Selector); exists {
						return fmt.Errorf("duplicate dst chain selector at args[%d].settings[%d].settings[%d] (src=%d): %d", i, j, k, src.Selector, dst.Selector)
					}
					if src.Selector == dst.Selector {
						return fmt.Errorf("src and dst chain selectors cannot be the same at args[%d].settings[%d].settings[%d] (src=%d)", i, j, k, src.Selector)
					}
					if addr := pool.PoolAddress; dst.IsReset {
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

func setTokenTransferFeeApply() func(deployment.Environment, SetTokenTransferFeeInput) (deployment.ChangesetOutput, error) {
	poolRegistry := GetTokenAdapterRegistry()
	mcmsRegistry := changesets.GetRegistry()

	return func(e deployment.Environment, cfg SetTokenTransferFeeInput) (deployment.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)
		for _, src := range cfg.Args {
			srcChainFam, err := chainsel.GetSelectorFamily(src.Selector)
			if err != nil {
				return deployment.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", src.Selector, err)
			}
			poolAdapter, exists := poolRegistry.GetTokenAdapter(srcChainFam, cfg.Version)
			if !exists {
				return deployment.ChangesetOutput{}, fmt.Errorf("no fee adapter found for chain family %s and version %s", srcChainFam, cfg.Version.String())
			}
			feesAdapter, ok := poolAdapter.(TokenFeeAdapter)
			if !ok {
				return deployment.ChangesetOutput{}, fmt.Errorf("adapter for chain family %s and version %s does not implement TokenFeeAdapter", srcChainFam, cfg.Version.String())
			}

			feeConfigSettings := map[string]map[uint64]*TokenTransferFeeConfig{}
			minBlocksSettings := map[string]uint16{}
			for _, pool := range src.TokenPools {
				feeConfigSettings[pool.PoolAddress] = map[uint64]*TokenTransferFeeConfig{}
				for _, dst := range pool.Destinations {
					if args, err := inferTokenTransferFeeArgs(feesAdapter, e, pool.PoolAddress, src.Selector, dst.Selector, dst); err != nil {
						return deployment.ChangesetOutput{}, fmt.Errorf("failed to infer token transfer fee args for src %d, dst %d, and pool %s: %w", src.Selector, dst.Selector, pool.PoolAddress, err)
					} else {
						feeConfigSettings[pool.PoolAddress][dst.Selector] = args
					}
				}
				if pool.MinBlockConfirmations.Valid {
					minBlocksSettings[pool.PoolAddress] = pool.MinBlockConfirmations.Value
				}
			}

			if len(minBlocksSettings) > 0 {
				minBlocksReport, err := cldf_ops.ExecuteSequence(
					e.OperationsBundle,
					feesAdapter.SetMinBlockConfirmations(&e),
					e.BlockChains,
					SetMinBlockConfirmationsSequenceInput{
						Selector: src.Selector,
						Settings: minBlocksSettings,
					},
				)
				if err != nil {
					return deployment.ChangesetOutput{}, fmt.Errorf("failed to execute SetMinBlockConfirmations operation for src %d: %w", src.Selector, err)
				}
				batchOps = append(batchOps, minBlocksReport.Output.BatchOps...)
				reports = append(reports, minBlocksReport.ExecutionReports...)
			}

			if len(feeConfigSettings) > 0 {
				feeConfigsReport, err := cldf_ops.ExecuteSequence(
					e.OperationsBundle,
					feesAdapter.SetTokenTransferFee(&e),
					e.BlockChains,
					SetTokenTransferFeeSequenceInput{
						Selector: src.Selector,
						Settings: feeConfigSettings,
					},
				)
				if err != nil {
					return deployment.ChangesetOutput{}, fmt.Errorf("failed to execute SetTokenTransferFee operation for src %d: %w", src.Selector, err)
				}
				batchOps = append(batchOps, feeConfigsReport.Output.BatchOps...)
				reports = append(reports, feeConfigsReport.ExecutionReports...)
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithBatchOps(batchOps).
			WithReports(reports).
			Build(cfg.MCMS)
	}
}

func inferTokenTransferFeeArgs(adapter TokenFeeAdapter, e deployment.Environment, poolAddress string, src uint64, dst uint64, cfg TokenTransferFeeForDst) (*TokenTransferFeeConfig, error) {
	if cfg.IsReset {
		e.Logger.Infof("Reset requested for token transfer fee config for src %d, dst %d, and pool %s; skipping inference", src, dst, poolAddress)
		return nil, nil
	}

	e.Logger.Infof("Inferring token transfer fee config for src %d, dst %d, and pool %s", src, dst, poolAddress)
	onchainCfg, err := adapter.GetOnchainTokenTransferFeeConfig(e, poolAddress, src, dst)
	if err != nil {
		return nil, fmt.Errorf("failed to get on-chain token transfer fee config for src %d, dst %d, and pool %s: %w", src, dst, poolAddress, err)
	}

	var fallbacks TokenTransferFeeConfig
	if onchainCfg.IsEnabled {
		fallbacks = onchainCfg
		e.Logger.Infof("Token transfer fee config for src %d, dst %d, and pool %s is already set on-chain; using on-chain values as defaults: %+v", src, dst, poolAddress, fallbacks)
	} else {
		fallbacks = adapter.GetDefaultTokenTransferFeeConfig(src, dst)
		e.Logger.Infof("Token transfer fee config for src %d, dst %d, and pool %s is not set on-chain; using adapter defaults: %+v", src, dst, poolAddress, fallbacks)
	}

	return cfg.Settings.Infer(fallbacks), nil
}

func GetDefaultChainAgnosticTokenTransferFeeConfig(src uint64, dst uint64, overrides ...func(*TokenTransferFeeConfig)) TokenTransferFeeConfig {
	minFeeUSDCents := uint32(25)

	// NOTE: we validate that src != dst so only one of these if statements will execute
	if src == chainsel.ETHEREUM_MAINNET.Selector {
		minFeeUSDCents = 50
	}
	if dst == chainsel.ETHEREUM_MAINNET.Selector {
		minFeeUSDCents = 150
	}

	cfg := TokenTransferFeeConfig{
		DefaultFinalityTransferFeeBps: 0,
		CustomFinalityTransferFeeBps:  0,
		DefaultFinalityFeeUSDCents:    minFeeUSDCents,
		CustomFinalityFeeUSDCents:     minFeeUSDCents,
		DestBytesOverhead:             32,
		DestGasOverhead:               90_000,
		IsEnabled:                     true,
	}

	for _, override := range overrides {
		override(&cfg)
	}

	return cfg
}
