package cctp

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/erc20"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_2/operations/hybrid_lock_release_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var MigrateHybridLockReleaseLiquidity = cldf_ops.NewSequence(
	"migrate-hybrid-lock-release-liquidity",
	semver.MustParse("1.7.0"),
	"Migrates a share of liquidity from HybridLockReleaseUSDCTokenPool into per-chain Siloed lockboxes",
	func(b cldf_ops.Bundle, deps adapters.MigrateHybridLockReleaseLiquidityDeps, input adapters.MigrateHybridLockReleaseLiquidityInput) (output sequences.OnChainOutput, err error) {
		chain, ok := deps.BlockChains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}
		if len(input.LockReleaseChainSelectors) == 0 {
			return sequences.OnChainOutput{}, fmt.Errorf("lock release chain selectors must be provided")
		}
		if chain.Selector != chain_selectors.ETHEREUM_MAINNET.Selector && chain.Selector != chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector {
			return sequences.OnChainOutput{}, fmt.Errorf("liquidity migration is only supported on home chains")
		}
		if input.LiquidityWithdrawPercent == 0 || input.LiquidityWithdrawPercent > 100 {
			return sequences.OnChainOutput{}, fmt.Errorf("liquidity withdraw percent must be between 1 and 100")
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

		writes := make([]contract_utils.WriteOutput, 0)
		lockBoxes := make(map[uint64]string)

		// Load lockbox mappings from the siloed pool to validate inputs.
		lockBoxFromSiloedPool, err := fetchLockBoxesFromSiloedPool(b, chain, input.ChainSelector, siloedPoolAddr)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		lockReleaseSelectors := input.LockReleaseChainSelectors
		// Validate selectors are unique and configured for lock-release in the hybrid pool.
		seenSelectors := make(map[uint64]struct{}, len(lockReleaseSelectors))
		for _, sel := range lockReleaseSelectors {
			if _, exists := seenSelectors[sel]; exists {
				return sequences.OnChainOutput{}, fmt.Errorf("duplicate lock release chain selector %d", sel)
			}
			seenSelectors[sel] = struct{}{}
			// Validate that the hybrid pool is configured for lock-release on this chain.
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
		}
		// Ensure each selector has a configured lockbox in the siloed pool.
		for _, sel := range lockReleaseSelectors {
			if lockBoxAddr, ok := lockBoxFromSiloedPool[sel]; ok && lockBoxAddr != (common.Address{}) {
				lockBoxes[sel] = lockBoxAddr.Hex()
				continue
			}
			return sequences.OnChainOutput{}, fmt.Errorf("lockbox not configured for chain %d", sel)
		}

		// Authorize the siloed pool and the MCMS timelock on each lockbox.
		// The siloed pool needs authorization for normal CCTP lock-release operations.
		// The timelock needs authorization because it will be the msg.sender calling deposit()
		// when the MCMS proposal executes.
		for _, sel := range lockReleaseSelectors {
			lockBox, ok := lockBoxes[sel]
			if !ok {
				continue
			}
			lockBoxAddr := common.HexToAddress(lockBox)
			callersReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.GetAllAuthorizedCallers, chain, contract_utils.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       lockBoxAddr,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get authorized callers for lockbox %s (chain %d): %w", lockBox, sel, err)
			}
			callersToAdd := make([]common.Address, 0, 2)
			if !slices.Contains(callersReport.Output, siloedPoolAddr) {
				callersToAdd = append(callersToAdd, siloedPoolAddr)
			}
			if !slices.Contains(callersReport.Output, timelockAddr) {
				callersToAdd = append(callersToAdd, timelockAddr)
			}
			if len(callersToAdd) == 0 {
				continue
			}
			authReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[erc20_lock_box.AuthorizedCallerArgs]{
				ChainSelector: input.ChainSelector,
				Address:       lockBoxAddr,
				Args: erc20_lock_box.AuthorizedCallerArgs{
					AddedCallers: callersToAdd,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to authorize callers on lockbox %s (chain %d): %w", lockBox, sel, err)
			}
			writes = append(writes, authReport.Output)
		}

		// For each lock-release chain, move the requested share of liquidity into the lockbox.
		for _, sel := range lockReleaseSelectors {
			lockBoxAddr, ok := lockBoxes[sel]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("lockbox address missing for chain %d", sel)
			}
			lockedReport, err := cldf_ops.ExecuteOperation(b, hybrid_lock_release_usdc_token_pool.GetLockedTokensForChain, chain, contract_utils.FunctionInput[uint64]{
				ChainSelector: input.ChainSelector,
				Address:       hybridPoolAddr,
				Args:          sel,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get locked tokens for chain %d: %w", sel, err)
			}
			if lockedReport.Output == nil || lockedReport.Output.Sign() <= 0 {
				continue
			}
			withdrawAmount := new(big.Int).Mul(lockedReport.Output, big.NewInt(int64(input.LiquidityWithdrawPercent)))
			withdrawAmount.Div(withdrawAmount, big.NewInt(100))
			if withdrawAmount.Sign() == 0 {
				continue
			}

			withdrawReport, err := cldf_ops.ExecuteOperation(b, hybrid_lock_release_usdc_token_pool.WithdrawLiquidity, chain, contract_utils.FunctionInput[hybrid_lock_release_usdc_token_pool.WithdrawLiquidityArgs]{
				ChainSelector: input.ChainSelector,
				Address:       hybridPoolAddr,
				Args: hybrid_lock_release_usdc_token_pool.WithdrawLiquidityArgs{
					RemoteChainSelector: sel,
					Amount:              withdrawAmount,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw liquidity for chain %d: %w", sel, err)
			}
			writes = append(writes, withdrawReport.Output)

			approveReport, err := cldf_ops.ExecuteOperation(b, erc20.ApproveProposalOnly, chain, contract_utils.FunctionInput[erc20.ApproveArgs]{
				ChainSelector: input.ChainSelector,
				Address:       tokenAddr,
				Args: erc20.ApproveArgs{
					Spender: common.HexToAddress(lockBoxAddr),
					Amount:  withdrawAmount,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to approve lockbox for chain %d: %w", sel, err)
			}
			writes = append(writes, approveReport.Output)

			depositReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.Deposit, chain, contract_utils.FunctionInput[erc20_lock_box.DepositArgs]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(lockBoxAddr),
				Args: erc20_lock_box.DepositArgs{
					Token:               tokenAddr,
					RemoteChainSelector: sel,
					Amount:              withdrawAmount,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to deposit into lockbox for chain %d: %w", sel, err)
			}
			writes = append(writes, depositReport.Output)
		}

		// Batch all writes into a single atomic MCMS operation.
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
	},
)

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
	lockBoxReport, err := cldf_ops.ExecuteOperation(b, siloed_usdc_token_pool.GetAllLockBoxConfigs, chain, contract_utils.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       poolAddress,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get lockbox configs: %w", err)
	}

	lockBoxes := make(map[uint64]common.Address, len(lockBoxReport.Output))
	for _, cfg := range lockBoxReport.Output {
		lockBoxes[cfg.RemoteChainSelector] = cfg.LockBox
	}
	return lockBoxes, nil
}
