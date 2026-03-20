package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences/cctp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	mcms_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
)

// MigrateHybridLockReleaseLiquidityConfig is the configuration for the MigrateHybridLockReleaseLiquidity changeset.
// This changeset is EVM-only and intended for Ethereum mainnet/Sepolia (home chains) where the hybrid lock-release
// pool and siloed lockboxes live.
type MigrateHybridLockReleaseLiquidityConfig struct {
	// ChainSelector is the home chain where liquidity will be migrated (must be Ethereum mainnet or Sepolia).
	ChainSelector uint64
	// HybridLockReleaseTokenPool is the address of the existing HybridLockReleaseUSDCTokenPool.
	HybridLockReleaseTokenPool string
	// SiloedUSDCTokenPool is the address of the SiloedUSDCTokenPool to migrate liquidity into.
	SiloedUSDCTokenPool string
	// USDCToken is the address of the USDC token contract.
	USDCToken string
	// WithdrawAmounts maps each remote chain selector to the absolute amount of USDC to migrate (raw units, e.g. 6 decimals for USDC).
	// Each chain must be configured for lock-release on the hybrid pool and have a lockbox on the siloed pool.
	WithdrawAmounts map[uint64]uint64
	// MCMS configures the resulting proposal. Required because this migration
	// operates on timelock-owned contracts and all operations must execute
	// atomically within a single MCMS proposal batch.
	MCMS mcms_utils.Input
}

// MigrateHybridLockReleaseLiquidity returns an EVM-only changeset that migrates liquidity from a
// HybridLockReleaseUSDCTokenPool into per-chain siloed lockboxes on the home chain.
func MigrateHybridLockReleaseLiquidity(mcmsRegistry *changesets.MCMSReaderRegistry) cldf_deployment.ChangeSetV2[MigrateHybridLockReleaseLiquidityConfig] {
	return cldf_deployment.CreateChangeSet(
		makeApplyMigrateHybridLockReleaseLiquidity(mcmsRegistry),
		makeVerifyMigrateHybridLockReleaseLiquidity(),
	)
}

func makeVerifyMigrateHybridLockReleaseLiquidity() func(cldf_deployment.Environment, MigrateHybridLockReleaseLiquidityConfig) error {
	return func(e cldf_deployment.Environment, cfg MigrateHybridLockReleaseLiquidityConfig) error {
		if err := cfg.MCMS.Validate(); err != nil {
			return fmt.Errorf("MCMS config is required and must be valid: %w", err)
		}
		if _, err := chain_selectors.GetSelectorFamily(cfg.ChainSelector); err != nil {
			return fmt.Errorf("invalid chain selector %d: %w", cfg.ChainSelector, err)
		}
		if cfg.ChainSelector != chain_selectors.ETHEREUM_MAINNET.Selector && cfg.ChainSelector != chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector {
			return fmt.Errorf("liquidity migration is only supported on Ethereum mainnet or Sepolia, got chain selector %d", cfg.ChainSelector)
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
		if len(cfg.WithdrawAmounts) == 0 {
			return fmt.Errorf("at least one withdraw amount must be provided")
		}
		for sel, amt := range cfg.WithdrawAmounts {
			if _, err := chain_selectors.GetSelectorFamily(sel); err != nil {
				return fmt.Errorf("invalid chain selector %d in withdraw amounts: %w", sel, err)
			}
			if amt == 0 {
				return fmt.Errorf("withdraw amount for chain %d must be greater than zero", sel)
			}
		}
		return nil
	}
}

func makeApplyMigrateHybridLockReleaseLiquidity(
	mcmsRegistry *changesets.MCMSReaderRegistry,
) func(cldf_deployment.Environment, MigrateHybridLockReleaseLiquidityConfig) (cldf_deployment.ChangesetOutput, error) {
	return func(e cldf_deployment.Environment, cfg MigrateHybridLockReleaseLiquidityConfig) (cldf_deployment.ChangesetOutput, error) {
		reader, ok := mcmsRegistry.GetMCMSReader(chain_selectors.FamilyEVM)
		if !ok {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("no MCMS reader registered for chain family 'evm'")
		}
		timelockRef, err := reader.GetTimelockRef(e, cfg.ChainSelector, cfg.MCMS)
		if err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to resolve timelock address on chain %d: %w", cfg.ChainSelector, err)
		}

		deps := adapters.MigrateHybridLockReleaseLiquidityDeps{
			BlockChains: e.BlockChains,
		}
		in := adapters.MigrateHybridLockReleaseLiquidityInput{
			ChainSelector:              cfg.ChainSelector,
			HybridLockReleaseTokenPool: cfg.HybridLockReleaseTokenPool,
			SiloedUSDCTokenPool:        cfg.SiloedUSDCTokenPool,
			USDCToken:                  cfg.USDCToken,
			WithdrawAmounts:            cfg.WithdrawAmounts,
			MCMSTimelockAddress:        timelockRef.Address,
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, cctp.MigrateHybridLockReleaseLiquidity, deps, in)
		if err != nil {
			return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to migrate hybrid lock-release liquidity on chain %d: %w", cfg.ChainSelector, err)
		}

		batchOps := report.Output.BatchOps
		reports := report.ExecutionReports

		newDS := datastore.NewMemoryDataStore()
		for _, addr := range report.Output.Addresses {
			if err := newDS.Addresses().Add(addr); err != nil {
				return cldf_deployment.ChangesetOutput{}, fmt.Errorf("failed to add address %s to datastore: %w", addr.Address, err)
			}
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			WithDataStore(newDS).
			Build(cfg.MCMS)
	}
}
