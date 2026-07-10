package cctp

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/erc20"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_2/operations/hybrid_lock_release_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_usdc_token_pool"
	hybrid_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_2/hybrid_lock_release_usdc_token_pool"
	elb_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/erc20_lock_box"
	siloed_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v2_0_0/adapters"
)

var MigrateHybridLockReleaseLiquidity = cldf_ops.NewSequence(
	"migrate-hybrid-lock-release-liquidity",
	semver.MustParse("2.0.0"),
	"Migrates absolute amounts of liquidity from HybridLockReleaseUSDCTokenPool into per-chain Siloed lockboxes",
	migrateHybridLockReleaseLiquidity,
)

func migrateHybridLockReleaseLiquidity(
	b cldf_ops.Bundle,
	deps adapters.MigrateHybridLockReleaseLiquidityDeps,
	input adapters.MigrateHybridLockReleaseLiquidityInput,
) (sequences.OnChainOutput, error) {
	chain, ok := deps.BlockChains.EVMChains()[input.ChainSelector]
	if !ok {
		return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
	}

	if chain.Selector != chain_selectors.ETHEREUM_MAINNET.Selector && chain.Selector != chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector {
		return sequences.OnChainOutput{}, fmt.Errorf("liquidity migration is only supported on home chains")
	}

	hybridPoolAddr, err := parseHexAddress("HybridLockReleaseUSDCTokenPool", input.HybridLockReleaseTokenPool)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}
	siloedPoolAddr, err := parseHexAddress("SiloedUSDCTokenPool", input.SiloedUSDCTokenPool)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}
	tokenAddr, err := parseHexAddress("USDC", input.USDCToken)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}
	timelockAddr, err := parseHexAddress("MCMSTimelock", input.MCMSTimelockAddress)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	// Sorting the selectors before iterating gives a stable, reproducible order (e.g. by chain selector)
	// and keeps the generated proposal and batch structure consistent for a given input.
	selectors := sortedSelectors(input.WithdrawAmounts)

	// Load lockbox mappings from the siloed pool.
	lockBoxFromSiloedPool, err := fetchLockBoxesFromSiloedPool(b, chain, siloedPoolAddr)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	// Validate all selectors: configured for lock-release, have lockboxes, amounts don't exceed locked.
	lockBoxes := make(map[uint64]common.Address, len(selectors))
	for _, sel := range selectors {
		shouldUseReport, err := evmops.ExecuteRead(
			b, chain, hybridPoolAddr,
			evmops.BindAs[hybrid_bindings.HybridLockReleaseUSDCTokenPoolInterface](hybrid_bindings.NewHybridLockReleaseUSDCTokenPool),
			hybrid_lock_release_usdc_token_pool.NewReadShouldUseLockRelease,
			sel,
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check lock-release mechanism for chain %d: %w", sel, err)
		}
		if !shouldUseReport.Output {
			return sequences.OnChainOutput{}, fmt.Errorf("hybrid pool not configured for lock-release on chain %d", sel)
		}

		lockBoxAddr, ok := lockBoxFromSiloedPool[sel]
		if !ok || lockBoxAddr == (common.Address{}) {
			return sequences.OnChainOutput{}, fmt.Errorf("lockbox not configured for chain %d", sel)
		}
		lockBoxes[sel] = lockBoxAddr

		lockedReport, err := evmops.ExecuteRead(
			b, chain, hybridPoolAddr,
			evmops.BindAs[hybrid_bindings.HybridLockReleaseUSDCTokenPoolInterface](hybrid_bindings.NewHybridLockReleaseUSDCTokenPool),
			hybrid_lock_release_usdc_token_pool.NewReadGetLockedTokensForChain,
			sel,
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get locked tokens for chain %d: %w", sel, err)
		}
		withdrawAmount := new(big.Int).SetUint64(input.WithdrawAmounts[sel])
		locked := lockedReport.Output
		if locked == nil || locked.Cmp(withdrawAmount) < 0 {
			lockedStr := "0"
			if locked != nil {
				lockedStr = locked.String()
			}
			return sequences.OnChainOutput{}, fmt.Errorf(
				"validation failed for chain %d: withdraw amount %s exceeds locked tokens %s (hybridPool=%s)",
				sel, withdrawAmount.String(), lockedStr, hybridPoolAddr.Hex(),
			)
		}
	}

	// Migration runs in three phases, all batched into a single MCMS proposal:
	// 1. Authorize callers: add siloed pool, timelock, and LP to each lockbox so deposits succeed.
	// 2. Migrate liquidity: withdraw from hybrid pool, approve lockbox, deposit into each lockbox.
	// 3. Transfer ownership: propose lockbox ownership to LPs (each LP must call acceptOwnership separately).
	ctx := &migratePhaseCtx{
		b:              b,
		chain:          chain,
		input:          input,
		hybridPoolAddr: hybridPoolAddr,
		siloedPoolAddr: siloedPoolAddr,
		tokenAddr:      tokenAddr,
		timelockAddr:   timelockAddr,
		selectors:      selectors,
		lockBoxes:      lockBoxes,
	}

	writes, err := authorizeLockboxCallers(ctx)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}
	liquidityWrites, err := migrateLiquidityToLockboxes(ctx)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}
	writes = append(writes, liquidityWrites...)
	ownershipWrites, err := transferLockboxOwnershipToLPs(ctx)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}
	writes = append(writes, ownershipWrites...)

	batchOps := make([]mcms_types.BatchOperation, 0)
	if len(writes) > 0 {
		batchOp, err := contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
		}
		batchOps = append(batchOps, batchOp)
	}

	return sequences.OnChainOutput{
		BatchOps: batchOps,
	}, nil
}

// migratePhaseCtx holds shared context for the migration phases.
type migratePhaseCtx struct {
	b              cldf_ops.Bundle
	chain          evm.Chain
	input          adapters.MigrateHybridLockReleaseLiquidityInput
	hybridPoolAddr common.Address
	siloedPoolAddr common.Address
	tokenAddr      common.Address
	timelockAddr   common.Address
	selectors      []uint64
	lockBoxes      map[uint64]common.Address
}

// authorizeLockboxCallers adds siloed pool, timelock, and LP as authorized callers on each lockbox.
// This must run before the MCMS batch so the timelock can deposit and the LP has access post-migration.
func authorizeLockboxCallers(ctx *migratePhaseCtx) ([]contract.WriteOutput, error) {
	if ctx == nil {
		return nil, fmt.Errorf("migratePhaseCtx is nil")
	}
	writes := make([]contract.WriteOutput, 0)
	for _, sel := range ctx.selectors {
		lockBoxAddr := ctx.lockBoxes[sel]

		lpReport, err := evmops.ExecuteRead(
			ctx.b, ctx.chain, ctx.hybridPoolAddr,
			evmops.BindAs[hybrid_bindings.HybridLockReleaseUSDCTokenPoolInterface](hybrid_bindings.NewHybridLockReleaseUSDCTokenPool),
			hybrid_lock_release_usdc_token_pool.NewReadGetLiquidityProvider,
			sel,
		)
		if err != nil {
			return nil, fmt.Errorf("authorizeLockboxCallers: chain %d lockbox %s: get liquidity provider: %w", sel, lockBoxAddr.Hex(), err)
		}
		lp := lpReport.Output

		callersReport, err := evmops.ExecuteRead(
			ctx.b, ctx.chain, lockBoxAddr,
			evmops.BindAs[elb_bindings.ERC20LockBoxInterface](elb_bindings.NewERC20LockBox),
			erc20_lock_box.NewReadGetAllAuthorizedCallers,
			struct{}{},
		)
		if err != nil {
			return nil, fmt.Errorf("authorizeLockboxCallers: chain %d lockbox %s: get authorized callers: %w", sel, lockBoxAddr.Hex(), err)
		}
		existingCallers := callersReport.Output
		if existingCallers == nil {
			existingCallers = []common.Address{}
		}

		callersToAdd := make([]common.Address, 0, 3)
		if !slices.Contains(existingCallers, ctx.siloedPoolAddr) {
			callersToAdd = append(callersToAdd, ctx.siloedPoolAddr)
		}
		if !slices.Contains(existingCallers, ctx.timelockAddr) {
			callersToAdd = append(callersToAdd, ctx.timelockAddr)
		}
		if lp != (common.Address{}) && lp != ctx.timelockAddr && !slices.Contains(existingCallers, lp) {
			callersToAdd = append(callersToAdd, lp)
		}

		if len(callersToAdd) > 0 {
			authReport, err := evmops.ExecuteWrite(
				ctx.b, ctx.chain, lockBoxAddr,
				evmops.BindAs[elb_bindings.ERC20LockBoxInterface](elb_bindings.NewERC20LockBox),
				erc20_lock_box.NewWriteApplyAuthorizedCallerUpdates,
				elb_bindings.AuthorizedCallersAuthorizedCallerArgs{
					AddedCallers:   callersToAdd,
					RemovedCallers: []common.Address{},
				},
			)
			if err != nil {
				return nil, fmt.Errorf("authorizeLockboxCallers: chain %d lockbox %s: apply authorized caller updates: %w", sel, lockBoxAddr.Hex(), err)
			}
			writes = append(writes, authReport.Output)
		}
	}

	return writes, nil

}

// migrateLiquidityToLockboxes withdraws from the hybrid pool, approves each lockbox, and deposits into it.
func migrateLiquidityToLockboxes(ctx *migratePhaseCtx) ([]contract.WriteOutput, error) {
	if ctx == nil {
		return nil, fmt.Errorf("migratePhaseCtx is nil")
	}
	writes := make([]contract.WriteOutput, 0)
	for _, sel := range ctx.selectors {
		lockBoxAddr := ctx.lockBoxes[sel]
		withdrawAmount := new(big.Int).SetUint64(ctx.input.WithdrawAmounts[sel])

		withdrawReport, err := evmops.ExecuteWrite(
			ctx.b, ctx.chain, ctx.hybridPoolAddr,
			evmops.BindAs[hybrid_bindings.HybridLockReleaseUSDCTokenPoolInterface](hybrid_bindings.NewHybridLockReleaseUSDCTokenPool),
			hybrid_lock_release_usdc_token_pool.NewWriteWithdrawLiquidity,
			hybrid_lock_release_usdc_token_pool.WithdrawLiquidityArgs{
				RemoteChainSelector: sel,
				Amount:              withdrawAmount,
			},
		)
		if err != nil {
			return nil, fmt.Errorf("migrateLiquidityToLockboxes: chain %d lockbox %s amount %s: withdraw: %w", sel, lockBoxAddr.Hex(), withdrawAmount.String(), err)
		}
		writes = append(writes, withdrawReport.Output)

		approveReport, err := evmops.ExecuteWrite(
			ctx.b, ctx.chain, ctx.tokenAddr,
			erc20_bindings.NewERC20,
			erc20.NewWriteApproveProposalOnly,
			erc20.ApproveArgs{
				Spender: lockBoxAddr,
				Value:   withdrawAmount,
			},
		)
		if err != nil {
			return nil, fmt.Errorf("migrateLiquidityToLockboxes: chain %d lockbox %s: approve USDC: %w", sel, lockBoxAddr.Hex(), err)
		}
		writes = append(writes, approveReport.Output)

		depositReport, err := evmops.ExecuteWrite(
			ctx.b, ctx.chain, lockBoxAddr,
			evmops.BindAs[elb_bindings.ERC20LockBoxInterface](elb_bindings.NewERC20LockBox),
			erc20_lock_box.NewWriteDeposit,
			erc20_lock_box.DepositArgs{
				Token:               ctx.tokenAddr,
				RemoteChainSelector: sel,
				Amount:              withdrawAmount,
			},
		)
		if err != nil {
			return nil, fmt.Errorf("migrateLiquidityToLockboxes: chain %d lockbox %s amount %s: deposit: %w", sel, lockBoxAddr.Hex(), withdrawAmount.String(), err)
		}
		writes = append(writes, depositReport.Output)
	}
	return writes, nil
}

// transferLockboxOwnershipToLPs proposes ownership transfer to the liquidity provider for each lockbox.
// Runs last so all deposits complete before ownership changes. LP must call acceptOwnership() afterward.
func transferLockboxOwnershipToLPs(ctx *migratePhaseCtx) ([]contract.WriteOutput, error) {
	if ctx == nil {
		return nil, fmt.Errorf("migratePhaseCtx is nil")
	}
	writes := make([]contract.WriteOutput, 0)
	for _, sel := range ctx.selectors {
		lockBoxAddr := ctx.lockBoxes[sel]

		lpReport, err := evmops.ExecuteRead(
			ctx.b, ctx.chain, ctx.hybridPoolAddr,
			evmops.BindAs[hybrid_bindings.HybridLockReleaseUSDCTokenPoolInterface](hybrid_bindings.NewHybridLockReleaseUSDCTokenPool),
			hybrid_lock_release_usdc_token_pool.NewReadGetLiquidityProvider,
			sel,
		)
		if err != nil {
			return nil, fmt.Errorf("transferLockboxOwnershipToLPs: chain %d lockbox %s: get liquidity provider: %w", sel, lockBoxAddr.Hex(), err)
		}
		lp := lpReport.Output

		if lp == (common.Address{}) || lp == ctx.timelockAddr {
			continue
		}

		ownerReport, err := evmops.ExecuteRead(
			ctx.b, ctx.chain, lockBoxAddr,
			evmops.BindAs[elb_bindings.ERC20LockBoxInterface](elb_bindings.NewERC20LockBox),
			erc20_lock_box.NewReadOwner,
			struct{}{},
		)
		if err != nil {
			return nil, fmt.Errorf("transferLockboxOwnershipToLPs: chain %d lockbox %s: get owner: %w", sel, lockBoxAddr.Hex(), err)
		}
		if ownerReport.Output == lp {
			continue
		}

		transferReport, err := evmops.ExecuteWrite(
			ctx.b, ctx.chain, lockBoxAddr,
			evmops.BindAs[elb_bindings.ERC20LockBoxInterface](elb_bindings.NewERC20LockBox),
			erc20_lock_box.NewWriteTransferOwnership,
			lp,
		)
		if err != nil {
			return nil, fmt.Errorf("transferLockboxOwnershipToLPs: chain %d lockbox %s lp %s: transfer ownership: %w", sel, lockBoxAddr.Hex(), lp.Hex(), err)
		}
		writes = append(writes, transferReport.Output)
	}
	return writes, nil
}

func sortedSelectors(m map[uint64]uint64) []uint64 {
	sels := make([]uint64, 0, len(m))
	for sel := range m {
		sels = append(sels, sel)
	}
	slices.Sort(sels)
	return sels
}

func parseHexAddress(name, address string) (common.Address, error) {
	if address == "" {
		return common.Address{}, fmt.Errorf("%s address is required", name)
	}
	if !common.IsHexAddress(address) {
		return common.Address{}, fmt.Errorf("%s address %q is not a valid hex address", name, address)
	}
	parsed := common.HexToAddress(address)
	if parsed == (common.Address{}) {
		return common.Address{}, fmt.Errorf("%s address is zero", name)
	}
	return parsed, nil
}

func fetchLockBoxesFromSiloedPool(b cldf_ops.Bundle, chain evm.Chain, poolAddress common.Address) (map[uint64]common.Address, error) {
	lockBoxReport, err := evmops.ExecuteRead(
		b, chain, poolAddress,
		evmops.BindAs[siloed_bindings.SiloedUSDCTokenPoolInterface](siloed_bindings.NewSiloedUSDCTokenPool),
		siloed_usdc_token_pool.NewReadGetAllLockBoxConfigs,
		struct{}{},
	)
	if err != nil {
		return nil, fmt.Errorf("fetchLockBoxesFromSiloedPool: siloedPool=%s chainSelector=%d: %w", poolAddress.Hex(), chain.Selector, err)
	}
	configs := lockBoxReport.Output
	if configs == nil {
		configs = []siloed_bindings.SiloedLockReleaseTokenPoolLockBoxConfig{}
	}

	lockBoxes := make(map[uint64]common.Address, len(configs))
	for _, cfg := range configs {
		lockBoxes[cfg.RemoteChainSelector] = cfg.LockBox
	}
	return lockBoxes, nil
}
