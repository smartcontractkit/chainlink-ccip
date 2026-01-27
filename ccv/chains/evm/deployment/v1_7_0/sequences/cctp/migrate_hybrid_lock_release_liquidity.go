package cctp

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/hybrid_lock_release_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/siloed_usdc_token_pool"
	contract_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type MigrateHybridLockReleaseLiquidityInput struct {
	ChainSelector              uint64
	HybridLockReleaseTokenPool string
	SiloedUSDCTokenPool        string
	USDCToken                  string
	LockReleaseChainSelectors  []uint64
	// LiquidityWithdrawPercent is the percent of locked liquidity to migrate (1-100).
	LiquidityWithdrawPercent uint8
}

type MigrateHybridLockReleaseLiquidityOutput struct {
	Addresses []datastore.AddressRef
	BatchOps  []mcms_types.BatchOperation
	LockBoxes map[uint64]string
}

var MigrateHybridLockReleaseLiquidity = cldf_ops.NewSequence(
	"migrate-hybrid-lock-release-liquidity",
	semver.MustParse("1.7.0"),
	"Migrates a share of liquidity from HybridLockReleaseUSDCTokenPool into per-chain Siloed lockboxes",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input MigrateHybridLockReleaseLiquidityInput) (output MigrateHybridLockReleaseLiquidityOutput, err error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}
		if len(input.LockReleaseChainSelectors) == 0 {
			return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("lock release chain selectors must be provided")
		}
		if chain.Selector != chain_selectors.ETHEREUM_MAINNET.Selector && chain.Selector != chain_selectors.ETHEREUM_TESTNET_SEPOLIA.Selector {
			return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("liquidity migration is only supported on home chains")
		}
		if input.LiquidityWithdrawPercent == 0 || input.LiquidityWithdrawPercent > 100 {
			return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("liquidity withdraw percent must be between 1 and 100")
		}

		hybridPoolAddr, err := parseHexAddress("HybridLockReleaseUSDCTokenPool", input.HybridLockReleaseTokenPool)
		if err != nil {
			return MigrateHybridLockReleaseLiquidityOutput{}, err
		}
		siloedPoolAddr, err := parseHexAddress("SiloedUSDCTokenPool", input.SiloedUSDCTokenPool)
		if err != nil {
			return MigrateHybridLockReleaseLiquidityOutput{}, err
		}
		tokenAddr, err := parseHexAddress("USDC", input.USDCToken)
		if err != nil {
			return MigrateHybridLockReleaseLiquidityOutput{}, err
		}

		addresses := make([]datastore.AddressRef, 0)
		writes := make([]contract_utils.WriteOutput, 0)
		lockBoxes := make(map[uint64]string)

		// Load lockbox mappings from the siloed pool to validate inputs.
		lockBoxFromSiloedPool, err := fetchLockBoxesFromSiloedPool(b, chain, input.ChainSelector, siloedPoolAddr)
		if err != nil {
			return MigrateHybridLockReleaseLiquidityOutput{}, err
		}
		lockReleaseSelectors := input.LockReleaseChainSelectors
		// Validate selectors are unique and configured for lock-release in the hybrid pool.
		seenSelectors := make(map[uint64]struct{}, len(lockReleaseSelectors))
		for _, sel := range lockReleaseSelectors {
			if _, exists := seenSelectors[sel]; exists {
				return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("duplicate lock release chain selector %d", sel)
			}
			seenSelectors[sel] = struct{}{}
			// Validate that the hybrid pool is configured for lock-release on this chain.
			shouldUseReport, err := cldf_ops.ExecuteOperation(b, hybrid_lock_release_usdc_token_pool.ShouldUseLockRelease, chain, contract_utils.FunctionInput[uint64]{
				ChainSelector: input.ChainSelector,
				Address:       hybridPoolAddr,
				Args:          sel,
			})
			if err != nil {
				return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("failed to check lock-release mechanism for chain %d: %w", sel, err)
			}
			if !shouldUseReport.Output {
				return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("hybrid pool not configured for lock-release on chain %d", sel)
			}
		}
		// Ensure each selector has a configured lockbox in the siloed pool.
		for _, sel := range lockReleaseSelectors {
			if lockBoxAddr, ok := lockBoxFromSiloedPool[sel]; ok && lockBoxAddr != (common.Address{}) {
				lockBoxes[sel] = lockBoxAddr.Hex()
				continue
			}
			return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("lockbox not configured for chain %d", sel)
		}

		// Make sure the siloed pool is authorized on each lockbox before deposits.
		for sel, lockBox := range lockBoxes {
			lockBoxAddr := common.HexToAddress(lockBox)
			callersReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.GetAllAuthorizedCallers, chain, contract_utils.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       lockBoxAddr,
			})
			if err != nil {
				return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("failed to get authorized callers for lockbox %s (chain %d): %w", lockBox, sel, err)
			}
			if containsAddress(callersReport.Output, siloedPoolAddr) {
				continue
			}
			authReport, err := cldf_ops.ExecuteOperation(b, erc20_lock_box.ApplyAuthorizedCallerUpdates, chain, contract_utils.FunctionInput[erc20_lock_box.AuthorizedCallerArgs]{
				ChainSelector: input.ChainSelector,
				Address:       lockBoxAddr,
				Args: erc20_lock_box.AuthorizedCallerArgs{
					AddedCallers: []common.Address{siloedPoolAddr},
				},
			})
			if err != nil {
				return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("failed to authorize siloed pool on lockbox %s (chain %d): %w", lockBox, sel, err)
			}
			writes = append(writes, authReport.Output)
		}

		// For each lock-release chain, move the requested share of liquidity into the lockbox.
		for _, sel := range lockReleaseSelectors {
			lockBoxAddr, ok := lockBoxes[sel]
			if !ok {
				return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("lockbox address missing for chain %d", sel)
			}
			lockedReport, err := cldf_ops.ExecuteOperation(b, hybrid_lock_release_usdc_token_pool.GetLockedTokensForChain, chain, contract_utils.FunctionInput[uint64]{
				ChainSelector: input.ChainSelector,
				Address:       hybridPoolAddr,
				Args:          sel,
			})
			if err != nil {
				return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("failed to get locked tokens for chain %d: %w", sel, err)
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
				return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("failed to withdraw liquidity for chain %d: %w", sel, err)
			}
			writes = append(writes, withdrawReport.Output)

			approveReport, err := cldf_ops.ExecuteOperation(b, erc20.Approve, chain, contract_utils.FunctionInput[erc20.ApproveArgs]{
				ChainSelector: input.ChainSelector,
				Address:       tokenAddr,
				Args: erc20.ApproveArgs{
					Spender: common.HexToAddress(lockBoxAddr),
					Amount:  withdrawAmount,
				},
			})
			if err != nil {
				return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("failed to approve lockbox for chain %d: %w", sel, err)
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
				return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("failed to deposit into lockbox for chain %d: %w", sel, err)
			}
			writes = append(writes, depositReport.Output)
		}

		// Batch all writes into a single atomic MCMS operation.
		batchOps := make([]mcms_types.BatchOperation, 0)
		if len(writes) > 0 {
			batchOp, err := contract_utils.NewBatchOperationFromWrites(writes)
			if err != nil {
				return MigrateHybridLockReleaseLiquidityOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
			}
			batchOps = append(batchOps, batchOp)
		}

		return MigrateHybridLockReleaseLiquidityOutput{
			Addresses: addresses,
			BatchOps:  batchOps,
			LockBoxes: lockBoxes,
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
