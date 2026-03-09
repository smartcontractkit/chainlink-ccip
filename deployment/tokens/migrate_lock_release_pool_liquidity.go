package tokens

import (
	"fmt"
	"math/big"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// LockReleasePoolMigration specifies a single pool migration to perform.
type LockReleasePoolMigration struct {
	// ChainSelector identifies the chain on which both the old and new pools live.
	ChainSelector uint64
	// OldPoolRef is a reference to the legacy LockReleaseTokenPool (v1.5.1 or v1.6.1) to migrate from.
	// Required because in step-2 migrations, the TAR already points to the new pool.
	OldPoolRef datastore.AddressRef
	// NewPoolRef is a reference to the new v2.0 LockReleaseTokenPool (with lockbox) to migrate to.
	NewPoolRef datastore.AddressRef
	// Amount specifies an exact token amount to migrate. Mutually exclusive with BasisPoints.
	Amount *big.Int
	// BasisPoints specifies a percentage (1-10000, where 10000 = 100%) of the old pool's balance to migrate.
	// Mutually exclusive with Amount.
	BasisPoints *uint16
	// RegistryRef, if provided, triggers a setPool call on the TokenAdminRegistry after migration.
	RegistryRef *datastore.AddressRef
	// TokenRef, if provided along with RegistryRef, specifies the token address for the setPool call.
	TokenRef *datastore.AddressRef
}

// MigrateLockReleasePoolLiquidityConfig is the configuration for the MigrateLockReleasePoolLiquidity changeset.
type MigrateLockReleasePoolLiquidityConfig struct {
	// Migrations specifies the pool migrations to perform.
	Migrations []LockReleasePoolMigration
	// MCMS configures the resulting proposal.
	MCMS mcms.Input
}

// MigrateLockReleasePoolLiquidity returns a changeset that migrates liquidity from legacy LockReleaseTokenPools
// to v2.0 lockbox-based pools. This is intended for the step-2 drain or standalone migration operations.
func MigrateLockReleasePoolLiquidity(tokenRegistry *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[MigrateLockReleasePoolLiquidityConfig] {
	return cldf.CreateChangeSet(makeMigrationApply(tokenRegistry, mcmsRegistry), makeMigrationVerify())
}

func makeMigrationVerify() func(cldf.Environment, MigrateLockReleasePoolLiquidityConfig) error {
	return func(_ cldf.Environment, _ MigrateLockReleasePoolLiquidityConfig) error {
		return nil
	}
}

func makeMigrationApply(_ *TokenAdapterRegistry, mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, MigrateLockReleasePoolLiquidityConfig) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, cfg MigrateLockReleasePoolLiquidityConfig) (cldf.ChangesetOutput, error) {
		tokenRegistry := GetTokenAdapterRegistry()
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		for i, migration := range cfg.Migrations {
			family, err := chain_selectors.GetSelectorFamily(migration.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to get chain family for selector %d: %w", i, migration.ChainSelector, err)
			}
			adapter, ok := tokenRegistry.GetTokenAdapter(family, migration.NewPoolRef.Version)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: no token adapter registered for chain family '%s' and version '%s'", i, family, migration.NewPoolRef.Version)
			}

			migrationSeq := adapter.MigrateLockReleasePoolLiquiditySequence()
			if migrationSeq == nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: adapter for family '%s' version '%s' does not support liquidity migration", i, family, migration.NewPoolRef.Version)
			}

			oldPoolRef, err := datastore_utils.FindAndFormatRef(e.DataStore, migration.OldPoolRef, migration.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to resolve old pool ref: %w", i, err)
			}
			newPoolRef, err := datastore_utils.FindAndFormatRef(e.DataStore, migration.NewPoolRef, migration.ChainSelector, datastore_utils.FullRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to resolve new pool ref: %w", i, err)
			}

			// Derive the timelock address from the MCMS config.
			mcmsReader, ok := mcmsRegistry.GetMCMSReader(family)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: no MCMS reader registered for chain family '%s'", i, family)
			}
			timelockRef, err := mcmsReader.GetTimelockRef(e, migration.ChainSelector, cfg.MCMS)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to get timelock address from MCMS config: %w", i, err)
			}

			var setPoolConfig *MigrationSetPoolConfig
			if migration.RegistryRef != nil && migration.TokenRef != nil {
				registryRef, err := datastore_utils.FindAndFormatRef(e.DataStore, *migration.RegistryRef, migration.ChainSelector, datastore_utils.FullRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to resolve registry ref: %w", i, err)
				}
				tokenRef, err := datastore_utils.FindAndFormatRef(e.DataStore, *migration.TokenRef, migration.ChainSelector, datastore_utils.FullRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to resolve token ref: %w", i, err)
				}
				setPoolConfig = &MigrationSetPoolConfig{
					RegistryAddress: registryRef.Address,
					TokenAddress:    tokenRef.Address,
				}
			}

			migrationReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, migrationSeq, e.BlockChains, MigrateLockReleasePoolLiquidityInput{
				ChainSelector:   migration.ChainSelector,
				OldPoolAddress:  oldPoolRef.Address,
				NewPoolAddress:  newPoolRef.Address,
				TimelockAddress: timelockRef.Address,
				Amount:          migration.Amount,
				BasisPoints:     migration.BasisPoints,
				SetPoolConfig:   setPoolConfig,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to execute liquidity migration: %w", i, err)
			}

			batchOps = append(batchOps, migrationReport.Output.BatchOps...)
			reports = append(reports, migrationReport.ExecutionReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}
