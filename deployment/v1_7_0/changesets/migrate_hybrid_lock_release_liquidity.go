package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
)

// MigrateHybridLockReleaseLiquidityConfig is the configuration for the MigrateHybridLockReleaseLiquidity changeset.
type MigrateHybridLockReleaseLiquidityConfig struct {
	// ChainSelector is the home chain where liquidity will be migrated.
	ChainSelector uint64
	// HybridLockReleaseTokenPool is the address of the existing HybridLockReleaseUSDCTokenPool.
	HybridLockReleaseTokenPool string
	// SiloedUSDCTokenPool is the address of the SiloedUSDCTokenPool to migrate liquidity into.
	SiloedUSDCTokenPool string
	// USDCToken is the address of the USDC token contract.
	USDCToken string
	// LockReleaseChainSelectors specifies which remote chains' locked liquidity to migrate.
	LockReleaseChainSelectors []uint64
	// LiquidityWithdrawPercent is the percent of locked liquidity to migrate (1-100).
	LiquidityWithdrawPercent uint8
	// USDCType specifies the type of the USDC on the chain.
	USDCType adapters.USDCType
	// MCMS configures the resulting proposal. Required because this migration
	// operates on timelock-owned contracts and all operations must execute
	// atomically within a single MCMS proposal batch.
	MCMS mcms.Input
}

// MigrateHybridLockReleaseLiquidity returns a changeset that migrates liquidity from a
// HybridLockReleaseUSDCTokenPool into per-chain siloed lockboxes.
func MigrateHybridLockReleaseLiquidity(cctpChainRegistry *adapters.CCTPChainRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[MigrateHybridLockReleaseLiquidityConfig] {
	return cldf.CreateChangeSet(
		makeApplyMigrateHybridLockReleaseLiquidity(cctpChainRegistry, mcmsRegistry),
		makeVerifyMigrateHybridLockReleaseLiquidity(),
	)
}

func makeVerifyMigrateHybridLockReleaseLiquidity() func(cldf.Environment, MigrateHybridLockReleaseLiquidityConfig) error {
	return func(e cldf.Environment, cfg MigrateHybridLockReleaseLiquidityConfig) error {
		if err := cfg.MCMS.Validate(); err != nil {
			return fmt.Errorf("MCMS config is required and must be valid: %w", err)
		}
		if _, err := chain_selectors.GetSelectorFamily(cfg.ChainSelector); err != nil {
			return fmt.Errorf("invalid chain selector %d: %w", cfg.ChainSelector, err)
		}
		if !cfg.USDCType.IsValid() {
			return fmt.Errorf("invalid USDC type: %s", cfg.USDCType)
		}
		if !common.IsHexAddress(cfg.HybridLockReleaseTokenPool) {
			return fmt.Errorf("invalid HybridLockReleaseTokenPool address for chain %d", cfg.ChainSelector)
		}
		if !common.IsHexAddress(cfg.SiloedUSDCTokenPool) {
			return fmt.Errorf("invalid SiloedUSDCTokenPool address for chain %d", cfg.ChainSelector)
		}
		if !common.IsHexAddress(cfg.USDCToken) {
			return fmt.Errorf("invalid USDCToken address for chain %d", cfg.ChainSelector)
		}
		if len(cfg.LockReleaseChainSelectors) == 0 {
			return fmt.Errorf("at least one lock release chain selector must be provided")
		}
		if cfg.LiquidityWithdrawPercent == 0 || cfg.LiquidityWithdrawPercent > 100 {
			return fmt.Errorf("liquidity withdraw percent must be between 1 and 100")
		}
		for _, sel := range cfg.LockReleaseChainSelectors {
			if _, err := chain_selectors.GetSelectorFamily(sel); err != nil {
				return fmt.Errorf("invalid lock release chain selector %d: %w", sel, err)
			}
		}
		return nil
	}
}

func makeApplyMigrateHybridLockReleaseLiquidity(
	cctpChainRegistry *adapters.CCTPChainRegistry,
	mcmsRegistry *changesets.MCMSReaderRegistry,
) func(cldf.Environment, MigrateHybridLockReleaseLiquidityConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg MigrateHybridLockReleaseLiquidityConfig) (cldf.ChangesetOutput, error) {
		family, err := chain_selectors.GetSelectorFamily(cfg.ChainSelector)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for chain selector %d: %w", cfg.ChainSelector, err)
		}
		adapter, ok := cctpChainRegistry.GetCCTPChain(family, cfg.USDCType)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no CCTP adapter registered for chain family '%s' and type '%s'", family, cfg.USDCType)
		}

		reader, ok := mcmsRegistry.GetMCMSReader(family)
		if !ok {
			return cldf.ChangesetOutput{}, fmt.Errorf("no MCMS reader registered for chain family '%s'", family)
		}
		timelockRef, err := reader.GetTimelockRef(e, cfg.ChainSelector, cfg.MCMS)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to resolve timelock address on chain %d: %w", cfg.ChainSelector, err)
		}

		deps := adapters.MigrateHybridLockReleaseLiquidityDeps{
			BlockChains: e.BlockChains,
		}
		in := adapters.MigrateHybridLockReleaseLiquidityInput{
			ChainSelector:              cfg.ChainSelector,
			HybridLockReleaseTokenPool: cfg.HybridLockReleaseTokenPool,
			SiloedUSDCTokenPool:        cfg.SiloedUSDCTokenPool,
			USDCToken:                  cfg.USDCToken,
			LockReleaseChainSelectors:  cfg.LockReleaseChainSelectors,
			LiquidityWithdrawPercent:   cfg.LiquidityWithdrawPercent,
			MCMSTimelockAddress:        timelockRef.Address,
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, adapter.MigrateHybridLockReleaseLiquidity(), deps, in)
		if err != nil {
			return cldf.ChangesetOutput{}, fmt.Errorf("failed to migrate hybrid lock-release liquidity on chain %d: %w", cfg.ChainSelector, err)
		}

		batchOps := report.Output.BatchOps
		reports := report.ExecutionReports

		newDS := datastore.NewMemoryDataStore()
		for _, addr := range report.Output.Addresses {
			if err := newDS.Addresses().Add(addr); err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to add address %s to datastore: %w", addr.Address, err)
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			WithDataStore(newDS).
			Build(cfg.MCMS)
	}
}
