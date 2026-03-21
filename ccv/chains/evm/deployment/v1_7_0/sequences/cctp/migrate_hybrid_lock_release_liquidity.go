package cctp

import (
	"fmt"
	"math/big"
	"slices"
	"sort"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/siloed_usdc_token_pool"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_2/operations/hybrid_lock_release_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
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
	lockBoxFromSiloedPool, err := fetchLockBoxesFromSiloedPool(b, chain, input.ChainSelector, siloedPoolAddr)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	// Validate all selectors: configured for lock-release, have lockboxes, amounts don't exceed locked.
	lockBoxes := make(map[uint64]common.Address, len(selectors))
	for _, sel := range selectors {
		shouldUseReport, err := cldf_ops.ExecuteOperation(b, hybrid_lock_release_usdc_token_pool.ShouldUseLockRelease, chain, contract_utils.FunctionInput[uint64]{
			ChainSelector: input.ChainSelector,
			Address:       hybridPoolAddr,
			Args:          sel,
		})
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

		lockedReport, err := cldf_ops.ExecuteOperation(b, hybrid_lock_release_usdc_token_pool.GetLockedTokensForChain, chain, contract_utils.FunctionInput[uint64]{
			ChainSelector: input.ChainSelector,
			Address:       hybridPoolAddr,
			Args:          sel,
		})
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
		batchOp, err := contract_utils.NewBatchOperationFromWrites(writes)
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
func authorizeLockboxCallers(ctx *migratePhaseCtx) ([]contract_utils.WriteOutput, error) {
	if ctx == nil {
		return nil, fmt.Errorf("migratePhaseCtx is nil")
	}
	writes := make([]contract_utils.WriteOutput, 0)
	for _, sel := range ctx.selectors {
		lockBoxAddr := ctx.lockBoxes[sel]

		lpReport, err := cldf_ops.ExecuteOperation(ctx.b, hybrid_lock_release_usdc_token_pool.GetLiquidityProvider, ctx.chain, contract_utils.FunctionInput[uint64]{
			ChainSelector: ctx.input.ChainSelector,
			Address:       ctx.hybridPoolAddr,
			Args:          sel,
		})
		if err != nil {
			return nil, fmt.Errorf("authorizeLockboxCallers: chain %d lockbox %s: get liquidity provider: %w", sel, lockBoxAddr.Hex(), err)
		}
		lp := lpReport.Output

		callersReport, err := cldf_ops.ExecuteOperation(ctx.b, erc20_lock_box.GetAllAuthorizedCallers, ctx.chain, contract_utils.FunctionInput[struct{}]{
			ChainSelector: ctx.input.ChainSelector,
			Address:       lockBoxAddr,
			Args:          struct{}{},
		})
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
			authReport, err := cldf_ops.ExecuteOperation(ctx.b, erc20_lock_box.ApplyAuthorizedCallerUpdates, ctx.chain, contract_utils.FunctionInput[erc20_lock_box.AuthorizedCallerArgs]{
				ChainSelector: ctx.input.ChainSelector,
				Address:       lockBoxAddr,
				Args: erc20_lock_box.AuthorizedCallerArgs{
					AddedCallers: callersToAdd,
				},
			})
			if err != nil {
				return nil, fmt.Errorf("authorizeLockboxCallers: chain %d lockbox %s: apply authorized caller updates: %w", sel, lockBoxAddr.Hex(), err)
			}
			writes = append(writes, authReport.Output)
		}
	}

	return writes, nil

}

// migrateLiquidityToLockboxes withdraws from the hybrid pool, approves each lockbox, and deposits into it.
func migrateLiquidityToLockboxes(ctx *migratePhaseCtx) ([]contract_utils.WriteOutput, error) {
	if ctx == nil {
		return nil, fmt.Errorf("migratePhaseCtx is nil")
	}
	writes := make([]contract_utils.WriteOutput, 0)
	for _, sel := range ctx.selectors {
		lockBoxAddr := ctx.lockBoxes[sel]
		withdrawAmount := new(big.Int).SetUint64(ctx.input.WithdrawAmounts[sel])

		withdrawReport, err := cldf_ops.ExecuteOperation(ctx.b, hybrid_lock_release_usdc_token_pool.WithdrawLiquidity, ctx.chain, contract_utils.FunctionInput[hybrid_lock_release_usdc_token_pool.WithdrawLiquidityArgs]{
			ChainSelector: ctx.input.ChainSelector,
			Address:       ctx.hybridPoolAddr,
			Args: hybrid_lock_release_usdc_token_pool.WithdrawLiquidityArgs{
				RemoteChainSelector: sel,
				Amount:              withdrawAmount,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("migrateLiquidityToLockboxes: chain %d lockbox %s amount %s: withdraw: %w", sel, lockBoxAddr.Hex(), withdrawAmount.String(), err)
		}
		writes = append(writes, withdrawReport.Output)

		approveReport, err := cldf_ops.ExecuteOperation(ctx.b, erc20.ApproveProposalOnly, ctx.chain, contract_utils.FunctionInput[erc20.ApproveArgs]{
			ChainSelector: ctx.input.ChainSelector,
			Address:       ctx.tokenAddr,
			Args: erc20.ApproveArgs{
				Spender: lockBoxAddr,
				Amount:  withdrawAmount,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("migrateLiquidityToLockboxes: chain %d lockbox %s: approve USDC: %w", sel, lockBoxAddr.Hex(), err)
		}
		writes = append(writes, approveReport.Output)

		depositReport, err := cldf_ops.ExecuteOperation(ctx.b, erc20_lock_box.Deposit, ctx.chain, contract_utils.FunctionInput[erc20_lock_box.DepositArgs]{
			ChainSelector: ctx.input.ChainSelector,
			Address:       lockBoxAddr,
			Args: erc20_lock_box.DepositArgs{
				Token:               ctx.tokenAddr,
				RemoteChainSelector: sel,
				Amount:              withdrawAmount,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("migrateLiquidityToLockboxes: chain %d lockbox %s amount %s: deposit: %w", sel, lockBoxAddr.Hex(), withdrawAmount.String(), err)
		}
		writes = append(writes, depositReport.Output)
	}
	return writes, nil
}

// transferLockboxOwnershipToLPs proposes ownership transfer to the liquidity provider for each lockbox.
// Runs last so all deposits complete before ownership changes. LP must call acceptOwnership() afterward.
func transferLockboxOwnershipToLPs(ctx *migratePhaseCtx) ([]contract_utils.WriteOutput, error) {
	if ctx == nil {
		return nil, fmt.Errorf("migratePhaseCtx is nil")
	}
	writes := make([]contract_utils.WriteOutput, 0)
	for _, sel := range ctx.selectors {
		lockBoxAddr := ctx.lockBoxes[sel]

		lpReport, err := cldf_ops.ExecuteOperation(ctx.b, hybrid_lock_release_usdc_token_pool.GetLiquidityProvider, ctx.chain, contract_utils.FunctionInput[uint64]{
			ChainSelector: ctx.input.ChainSelector,
			Address:       ctx.hybridPoolAddr,
			Args:          sel,
		})
		if err != nil {
			return nil, fmt.Errorf("transferLockboxOwnershipToLPs: chain %d lockbox %s: get liquidity provider: %w", sel, lockBoxAddr.Hex(), err)
		}
		lp := lpReport.Output

		if lp == (common.Address{}) || lp == ctx.timelockAddr {
			continue
		}

		ownerReport, err := cldf_ops.ExecuteOperation(ctx.b, erc20_lock_box.Owner, ctx.chain, contract_utils.FunctionInput[struct{}]{
			ChainSelector: ctx.input.ChainSelector,
			Address:       lockBoxAddr,
			Args:          struct{}{},
		})
		if err != nil {
			return nil, fmt.Errorf("transferLockboxOwnershipToLPs: chain %d lockbox %s: get owner: %w", sel, lockBoxAddr.Hex(), err)
		}
		if ownerReport.Output == lp {
			continue
		}

		transferReport, err := cldf_ops.ExecuteOperation(ctx.b, erc20_lock_box.TransferOwnership, ctx.chain, contract_utils.FunctionInput[common.Address]{
			ChainSelector: ctx.input.ChainSelector,
			Address:       lockBoxAddr,
			Args:          lp,
		})
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
	sort.Slice(sels, func(i, j int) bool { return sels[i] < sels[j] })
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

func fetchLockBoxesFromSiloedPool(b cldf_ops.Bundle, chain evm.Chain, chainSelector uint64, poolAddress common.Address) (map[uint64]common.Address, error) {
	lockBoxReport, err := cldf_ops.ExecuteOperation(b, siloed_usdc_token_pool.GetAllLockBoxConfigs, chain, contract_utils.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       poolAddress,
		Args:          struct{}{},
	})
	if err != nil {
		return nil, fmt.Errorf("fetchLockBoxesFromSiloedPool: siloedPool=%s chainSelector=%d: %w", poolAddress.Hex(), chainSelector, err)
	}
	configs := lockBoxReport.Output
	if configs == nil {
		configs = []siloed_usdc_token_pool.LockBoxConfig{}
	}

	lockBoxes := make(map[uint64]common.Address, len(configs))
	for _, cfg := range lockBoxReport.Output {
		lockBoxes[cfg.RemoteChainSelector] = cfg.LockBox
	}
	return lockBoxes, nil
}
