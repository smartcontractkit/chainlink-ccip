package tokens

import (
	"fmt"
	"math/big"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// LockReleasePoolMigration specifies a single pool migration to perform.
type LockReleasePoolMigration struct {
	// ChainSelector identifies the chain on which both the old and new pools live.
	ChainSelector uint64
	// OldPoolRef is a reference to the legacy lock-release pool to migrate from. Its version selects
	// the adapter whose source extracts the liquidity; supported today: LockReleaseTokenPool v1.5.1,
	// v1.6.1 and v1.6.2, and SiloedLockReleaseTokenPool v1.6.1 and v1.6.2. Other versions fail
	// with an explicit "unsupported" error. Required because in step-2 migrations, the TAR already
	// points to the new pool.
	OldPoolRef datastore.AddressRef
	// NewPoolRef is a reference to the new v2.0 LockReleaseTokenPool (with lockbox) to migrate to.
	NewPoolRef datastore.AddressRef
	// LiquidityMigrationAmount specifies how much liquidity to migrate, either as a raw token amount
	// (RAW) or as a percentage of the old pool's balance (BPS). Mutually exclusive with
	// SiloExactAmounts/UnsiloedExactAmount.
	LiquidityMigrationAmount *LockReleasePoolLiquidityMigrationAmount
	// SiloExactAmounts specifies exact per-silo migration amounts, keyed by remote chain selector.
	// Mutually exclusive with LiquidityMigrationAmount. See MigrateLockReleasePoolLiquidityInput.SiloExactAmounts.
	SiloExactAmounts []SiloExactAmount
	// UnsiloedExactAmount specifies the exact amount to migrate from the unsiloed (shared) balance.
	// Mutually exclusive with LiquidityMigrationAmount. See MigrateLockReleasePoolLiquidityInput.UnsiloedExactAmount.
	UnsiloedExactAmount *big.Int
	// UnsiloedLockBoxRef references the lockbox that receives the old pool's unsiloed (shared)
	// balance. Required when migrating a siloed pool that holds unsiloed liquidity; ignored for
	// non-siloed pools. See MigrateLockReleasePoolLiquidityInput.UnsiloedLockBoxAddress for why
	// this is specified rather than inferred.
	UnsiloedLockBoxRef *datastore.AddressRef
	// UsePlainTransfer, when true, transfers tokens directly to the lockbox via ERC20.transfer instead of
	// using the lockbox's deposit() function, bypassing the Deposit event emission. This is a break-glass
	// option for the standalone changeset; the standard deposit() path is the default.
	UsePlainTransfer bool
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
	return func(_ cldf.Environment, cfg MigrateLockReleasePoolLiquidityConfig) error {
		if len(cfg.Migrations) == 0 {
			return fmt.Errorf("at least one migration is required")
		}
		for i, migration := range cfg.Migrations {
			exactMode := len(migration.SiloExactAmounts) > 0 || migration.UnsiloedExactAmount != nil
			legacyMode := migration.LiquidityMigrationAmount != nil

			if exactMode && legacyMode {
				return fmt.Errorf("migration[%d]: SiloExactAmounts/UnsiloedExactAmount are mutually exclusive with LiquidityMigrationAmount", i)
			}
			if !exactMode && !legacyMode {
				return fmt.Errorf("migration[%d]: one of LiquidityMigrationAmount or SiloExactAmounts/UnsiloedExactAmount must be provided", i)
			}
			if migration.LiquidityMigrationAmount != nil {
				if err := migration.LiquidityMigrationAmount.Validate(); err != nil {
					return fmt.Errorf("migration[%d]: invalid LiquidityMigrationAmount: %w", i, err)
				}
			}
			if exactMode {
				if migration.UnsiloedExactAmount != nil && len(migration.SiloExactAmounts) == 0 {
					return fmt.Errorf("migration[%d]: UnsiloedExactAmount requires SiloExactAmounts to also be set; exact mode cannot migrate the unsiloed bucket alone", i)
				}
				seen := make(map[uint64]bool, len(migration.SiloExactAmounts))
				for j, sa := range migration.SiloExactAmounts {
					if seen[sa.ChainSelector] {
						return fmt.Errorf("migration[%d]: duplicate ChainSelector %d in SiloExactAmounts", i, sa.ChainSelector)
					}
					seen[sa.ChainSelector] = true
					if sa.Amount == nil || sa.Amount.Sign() < 0 {
						return fmt.Errorf("migration[%d]: SiloExactAmounts[%d].Amount must be positive", i, j)
					}
				}
				if migration.UnsiloedExactAmount != nil && migration.UnsiloedExactAmount.Sign() <= 0 {
					return fmt.Errorf("migration[%d]: UnsiloedExactAmount must be positive", i)
				}
			}
			if (migration.RegistryRef == nil) != (migration.TokenRef == nil) {
				return fmt.Errorf("migration[%d]: RegistryRef and TokenRef must both be set or both be omitted", i)
			}
		}
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

			oldPoolRef, err := ResolveTokenPoolRef(e, tokenRegistry, migration.ChainSelector, migration.OldPoolRef)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to resolve old pool ref: %w", i, err)
			}
			newPoolRef, err := ResolveTokenPoolRef(e, tokenRegistry, migration.ChainSelector, migration.NewPoolRef)
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
				regRef, err := deploy.TryNormalizeAddressRef(migration.ChainSelector, *migration.RegistryRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to normalize registry ref address: %w", i, err)
				}
				registryRef, err := datastore_utils.FindAndFormatRef(e.DataStore, regRef, migration.ChainSelector, datastore_utils.FullRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to resolve registry ref: %w", i, err)
				}
				tokenRef, err := ResolveTokenRef(e, tokenRegistry, migration.ChainSelector, *migration.TokenRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to resolve token ref: %w", i, err)
				}
				setPoolConfig = &MigrationSetPoolConfig{
					RegistryAddress: registryRef.Address,
					TokenAddress:    tokenRef.Address,
				}
			}

			var unsiloedLockBoxAddress string
			if migration.UnsiloedLockBoxRef != nil {
				lockBoxRef, err := deploy.TryNormalizeAddressRef(migration.ChainSelector, *migration.UnsiloedLockBoxRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to normalize unsiloed lockbox ref address: %w", i, err)
				}

				resolved, err := datastore_utils.FindAndFormatRef(e.DataStore, lockBoxRef, migration.ChainSelector, datastore_utils.FullRef)
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to resolve unsiloed lockbox ref: %w", i, err)
				}

				unsiloedLockBoxAddress = resolved.Address
			}

			var basisPoints *uint16
			var rawAmount *big.Int
			if migration.LiquidityMigrationAmount != nil {
				basisPoints, err = migration.LiquidityMigrationAmount.BasisPoints()
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to parse basis points: %w", i, err)
				}
				rawAmount, err = migration.LiquidityMigrationAmount.RawAmount()
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to parse raw amount: %w", i, err)
				}
			}

			migrationOutput, migrationReports, err := runLockReleaseLiquidityMigration(e, tokenRegistry, migration.ChainSelector, MigrateLockReleasePoolLiquidityInput{
				ChainSelector:          migration.ChainSelector,
				OldPoolAddress:         oldPoolRef.Address,
				NewPoolAddress:         newPoolRef.Address,
				TimelockAddress:        timelockRef.Address,
				Amount:                 rawAmount,
				BasisPoints:            basisPoints,
				SiloExactAmounts:       migration.SiloExactAmounts,
				UnsiloedExactAmount:    migration.UnsiloedExactAmount,
				UnsiloedLockBoxAddress: unsiloedLockBoxAddress,
				UsePlainTransfer:       migration.UsePlainTransfer,
				SetPoolConfig:          setPoolConfig,
			})
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("migration[%d]: failed to execute liquidity migration: %w", i, err)
			}

			batchOps = append(batchOps, migrationOutput.BatchOps...)
			reports = append(reports, migrationReports...)
		}

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(cfg.MCMS)
	}
}

// runLockReleaseLiquidityMigration composes a lock-release liquidity migration from two adapters:
// the SOURCE, resolved from the old pool's version, extracts liquidity to the timelock and records
// the amounts it moved in input.Resolved; the SINK, always the v2.0.0 adapter, funds the new pool's
// lockbox(es) with exactly those amounts.
//
// The source and sink batches for a chain are merged into a single batch operation so the timelock
// executes the extraction and the funding atomically - otherwise a failed funding batch would leave
// the extracted liquidity stranded in the timelock. Merging only combines emitted batch operations;
// it cannot make writes that already executed directly atomic. Sources therefore must emit their
// extraction writes into the batch rather than executing them directly (the EVM sources require
// the old pool to be owned by the timelock, so every owner-gated write is batched).
func runLockReleaseLiquidityMigration(
	e cldf.Environment,
	reg *TokenAdapterRegistry,
	sel uint64,
	input MigrateLockReleasePoolLiquidityInput,
) (sequences.OnChainOutput, []cldf_ops.Report[any, any], error) {
	if err := input.Validate(); err != nil {
		return sequences.OnChainOutput{}, nil, fmt.Errorf("invalid migration input: %w", err)
	}

	// SOURCE - resolved from the OLD pool's version (never the new pool's).
	oldPoolRef, err := ResolveTokenPoolRef(e, reg, sel, datastore.AddressRef{ChainSelector: sel, Address: input.OldPoolAddress})
	if err != nil {
		return sequences.OnChainOutput{}, nil, fmt.Errorf("failed to resolve old pool ref: %w", err)
	}
	// The source must be a legacy pool. A v2.0.0+ pool would resolve to the sink adapter, which must
	// never be used as a source.
	if oldPoolRef.Version == nil || !oldPoolRef.Version.LessThan(utils.Version_2_0_0) {
		return sequences.OnChainOutput{}, nil, fmt.Errorf(
			"old pool %s (chain %d) is version %v; the migration source must be older than v2.0.0",
			input.OldPoolAddress, sel, oldPoolRef.Version)
	}
	srcAdapter, _, err := ResolveAdapter(reg, sel, oldPoolRef.Version)
	if err != nil {
		return sequences.OnChainOutput{}, nil, fmt.Errorf(
			"old pool %s (chain %d) does not support liquidity migration yet: %w", input.OldPoolAddress, sel, err)
	}
	srcSeq := srcAdapter.MigrateLockReleasePoolLiquiditySequence()
	if srcSeq == nil {
		return sequences.OnChainOutput{}, nil, fmt.Errorf(
			"old pool %s (chain %d) does not support liquidity migration yet", input.OldPoolAddress, sel)
	}

	// The source writes the resolved amounts into input.Resolved.
	input.Resolved = &MigrateLockReleasePoolLiquidityResolved{}
	srcReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, srcSeq, e.BlockChains, input)
	if err != nil {
		return sequences.OnChainOutput{}, nil, fmt.Errorf("source migration failed: %w", err)
	}
	// Resolved is filled in by the source handler, which does not run when the framework replays a
	// previous report for an identical input. Fail rather than fund the sink with nothing.
	if input.Resolved.Amount == nil && len(input.Resolved.SiloAmounts) == 0 && input.Resolved.UnsiloedAmount == nil {
		return sequences.OnChainOutput{}, nil, fmt.Errorf(
			"source migration for old pool %s (chain %d) resolved no amounts; if it was replayed from a previous report, rerun with a fresh operations reporter",
			input.OldPoolAddress, sel)
	}

	// SINK - always the registered v2.0.0 adapter.
	sinkAdapter, _, err := ResolveAdapter(reg, sel, utils.Version_2_0_0)
	if err != nil {
		return sequences.OnChainOutput{}, nil, fmt.Errorf("failed to resolve v2.0.0 sink adapter: %w", err)
	}
	sinkSeq := sinkAdapter.MigrateLockReleasePoolLiquiditySequence()
	if sinkSeq == nil {
		return sequences.OnChainOutput{}, nil, fmt.Errorf("no v2.0.0 sink registered for chain %d", sel)
	}
	sinkInput := input
	sinkInput.Amount = input.Resolved.Amount
	sinkInput.BasisPoints = nil
	sinkInput.SiloExactAmounts = input.Resolved.SiloAmounts
	sinkInput.UnsiloedExactAmount = input.Resolved.UnsiloedAmount
	sinkInput.Resolved = nil
	sinkReport, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sinkSeq, e.BlockChains, sinkInput)
	if err != nil {
		return sequences.OnChainOutput{}, nil, fmt.Errorf("sink migration failed: %w", err)
	}

	return sequences.OnChainOutput{
		BatchOps:  mergeBatchOpsByChain(append(srcReport.Output.BatchOps, sinkReport.Output.BatchOps...)),
		Addresses: append(srcReport.Output.Addresses, sinkReport.Output.Addresses...),
	}, append(srcReport.ExecutionReports, sinkReport.ExecutionReports...), nil
}

// mergeBatchOpsByChain merges batch operations targeting the same chain into one, preserving the
// order of chains and of transactions within each chain. Empty batches (every write was executed
// directly) are dropped.
func mergeBatchOpsByChain(batchOps []mcms_types.BatchOperation) []mcms_types.BatchOperation {
	merged := make([]mcms_types.BatchOperation, 0, len(batchOps))
	indexByChain := make(map[mcms_types.ChainSelector]int, len(batchOps))
	for _, op := range batchOps {
		if len(op.Transactions) == 0 {
			continue
		}
		if i, ok := indexByChain[op.ChainSelector]; ok {
			merged[i].Transactions = append(merged[i].Transactions, op.Transactions...)
			continue
		}
		indexByChain[op.ChainSelector] = len(merged)
		merged = append(merged, mcms_types.BatchOperation{
			ChainSelector: op.ChainSelector,
			Transactions:  append([]mcms_types.Transaction(nil), op.Transactions...),
		})
	}
	return merged
}
