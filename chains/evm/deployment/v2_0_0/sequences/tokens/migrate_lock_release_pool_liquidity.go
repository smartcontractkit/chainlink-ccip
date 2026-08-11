package tokens

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	erc20_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	tar_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	lrtp_ops_v161 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/lock_release_token_pool"
	siloed_ops_v161 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/siloed_lock_release_token_pool"
	lrtp_ops_v170 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/lock_release_token_pool"
	siloed_lrtp_ops_v170 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_lock_release_token_pool"
	token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
)

var MigrateLockReleasePoolLiquidity = cldf_ops.NewSequence(
	"migrate-lock-release-pool-liquidity",
	semver.MustParse("2.0.0"),
	"Migrates liquidity from a legacy LockReleaseTokenPool (v1.5.1/v1.6.1) to a v2.0 lockbox-based pool",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input tokens.MigrateLockReleasePoolLiquidityInput) (sequences.OnChainOutput, error) {
		evmChain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		if err := validateMigrationInput(input); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid migration input: %w", err)
		}

		oldPoolAddr := common.HexToAddress(input.OldPoolAddress)
		newPoolAddr := common.HexToAddress(input.NewPoolAddress)
		timelockAddr := common.HexToAddress(input.TimelockAddress)

		tvReport, err := cldf_ops.ExecuteOperation(b, type_and_version.GetTypeAndVersion, evmChain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       oldPoolAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get typeAndVersion from old pool %s: %w", oldPoolAddr, err)
		}
		oldPoolType := string(tvReport.Output.Type)

		tokenReport, err := cldf_ops.ExecuteOperation(b, token_pool_ops.GetToken, evmChain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       newPoolAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from new pool %s: %w", newPoolAddr, err)
		}
		tokenAddr := tokenReport.Output

		// Only the generic siloed lock-release pool is handled here. A substring match on "Siloed"
		// would also catch SiloedUSDCTokenPool, which migrates through the CCTP hybrid path instead.
		isSiloed := oldPoolType == utils.SiloedLockReleaseTokenPool.String()

		if isSiloed {
			if input.Amount != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("siloed pools only support BasisPoints, not exact Amount")
			}
			return migrateSiloedPool(b, evmChain, input, oldPoolAddr, newPoolAddr, tokenAddr, timelockAddr)
		}
		return migrateUnsiloedPool(b, evmChain, input, oldPoolAddr, newPoolAddr, tokenAddr, timelockAddr)
	},
)

func validateMigrationInput(input tokens.MigrateLockReleasePoolLiquidityInput) error {
	if input.Amount != nil && input.BasisPoints != nil {
		return fmt.Errorf("Amount and BasisPoints are mutually exclusive")
	}
	if input.Amount == nil && input.BasisPoints == nil {
		return fmt.Errorf("one of Amount or BasisPoints must be provided")
	}
	if input.BasisPoints != nil {
		bp := *input.BasisPoints
		if bp == 0 || bp > 10000 {
			return fmt.Errorf("BasisPoints must be between 1 and 10000, got %d", bp)
		}
	}
	if input.Amount != nil && input.Amount.Sign() <= 0 {
		return fmt.Errorf("Amount must be positive")
	}
	if input.OldPoolAddress == "" || input.NewPoolAddress == "" {
		return fmt.Errorf("OldPoolAddress and NewPoolAddress must be provided")
	}
	if input.TimelockAddress == "" {
		return fmt.Errorf("TimelockAddress must be provided")
	}
	return nil
}

func computeAmount(balance *big.Int, input tokens.MigrateLockReleasePoolLiquidityInput) *big.Int {
	if input.Amount != nil {
		return new(big.Int).Set(input.Amount)
	}
	bp := *input.BasisPoints
	if bp == 10000 {
		return new(big.Int).Set(balance)
	}
	amount := new(big.Int).Mul(balance, big.NewInt(int64(bp)))
	return amount.Div(amount, big.NewInt(10000))
}

func migrateUnsiloedPool(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	input tokens.MigrateLockReleasePoolLiquidityInput,
	oldPoolAddr, newPoolAddr, tokenAddr, timelockAddr common.Address,
) (sequences.OnChainOutput, error) {
	chainSel := input.ChainSelector
	var ops []evm_contract.WriteOutput

	lockboxReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops_v170.GetLockBox, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       newPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get lockbox from new pool %s: %w", newPoolAddr, err)
	}
	lockboxAddr := lockboxReport.Output

	balanceReport, err := cldf_ops.ExecuteOperation(b, erc20_ops.BalanceOf, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       tokenAddr,
		Args:          oldPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get balance of old pool %s: %w", oldPoolAddr, err)
	}
	balance := balanceReport.Output

	amount := computeAmount(balance, input)
	if amount.Sign() == 0 {
		return sequences.OnChainOutput{}, fmt.Errorf("computed migration amount is zero")
	}
	if amount.Cmp(balance) > 0 {
		return sequences.OnChainOutput{}, fmt.Errorf("migration amount %s exceeds old pool balance %s", amount, balance)
	}

	rebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops_v161.GetRebalancer, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get rebalancer from old pool %s: %w", oldPoolAddr, err)
	}
	originalRebalancer := rebalancerReport.Output

	ops, err = appendSetRebalancerAndWithdraw(b, evmChain, chainSel, oldPoolAddr, timelockAddr, amount, ops)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	transferReport, err := cldf_ops.ExecuteOperation(b, erc20_ops.Transfer, evmChain, evm_contract.FunctionInput[erc20_ops.TransferArgs]{
		ChainSelector: chainSel,
		Address:       tokenAddr,
		Args: erc20_ops.TransferArgs{
			Receiver: lockboxAddr,
			Amount:   amount,
		},
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer tokens to lockbox %s: %w", lockboxAddr, err)
	}
	ops = append(ops, transferReport.Output)

	restoreRebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops_v161.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
		Args:          originalRebalancer,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to restore rebalancer on old pool %s: %w", oldPoolAddr, err)
	}
	ops = append(ops, restoreRebalancerReport.Output)

	if input.SetPoolConfig != nil {
		ops, err = appendSetPool(b, evmChain, chainSel, input.SetPoolConfig, newPoolAddr, ops)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
	}

	batchOp, err := evm_contract.NewBatchOperationFromWrites(ops)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
	}

	return sequences.OnChainOutput{
		BatchOps: []mcms_types.BatchOperation{batchOp},
	}, nil
}

func migrateSiloedPool(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	input tokens.MigrateLockReleasePoolLiquidityInput,
	oldPoolAddr, newPoolAddr, tokenAddr, timelockAddr common.Address,
) (sequences.OnChainOutput, error) {
	chainSel := input.ChainSelector
	var ops []evm_contract.WriteOutput

	chainsReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.GetSupportedChains, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get supported chains from old siloed pool %s: %w", oldPoolAddr, err)
	}
	supportedChains := chainsReport.Output

	lockboxConfigsReport, err := cldf_ops.ExecuteOperation(b, siloed_lrtp_ops_v170.GetAllLockBoxConfigs, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       newPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get lockbox configs from new pool %s: %w", newPoolAddr, err)
	}

	lockboxByChain := make(map[uint64]common.Address)
	for _, config := range lockboxConfigsReport.Output {
		if config.LockBox == (common.Address{}) {
			continue
		}
		lockboxByChain[config.RemoteChainSelector] = config.LockBox
	}

	// Resolve which of the old pool's chains are siloed once; used both for the coverage check below
	// and for the rebalancer handover further down.
	isSiloedByChain := make(map[uint64]bool, len(supportedChains))
	for _, remoteChain := range supportedChains {
		isSiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.IsSiloed, evmChain, evm_contract.FunctionInput[uint64]{
			ChainSelector: chainSel,
			Address:       oldPoolAddr,
			Args:          remoteChain,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check if chain %d is siloed on old pool %s: %w", remoteChain, oldPoolAddr, err)
		}
		isSiloedByChain[remoteChain] = isSiloedReport.Output
	}

	// Check lockbox coverage before emitting any writes. Without this the batch is built chain by
	// chain and a gap surfaces partway through - after the rebalancer has already been repointed at
	// the timelock - leaving the operator to work out which silo was missing.
	var chainsWithoutLockBox []uint64
	for _, remoteChain := range supportedChains {
		if !isSiloedByChain[remoteChain] {
			continue
		}
		if _, ok := lockboxByChain[remoteChain]; !ok {
			chainsWithoutLockBox = append(chainsWithoutLockBox, remoteChain)
		}
	}
	if len(chainsWithoutLockBox) > 0 {
		slices.Sort(chainsWithoutLockBox)
		return sequences.OnChainOutput{}, fmt.Errorf(
			"new siloed pool %s has no lockbox configured for siloed chains %v of old pool %s; the new pool's lockBoxGroups must cover every siloed chain being migrated",
			newPoolAddr, chainsWithoutLockBox, oldPoolAddr,
		)
	}

	rebalancerReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.GetRebalancer, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get unsiloed rebalancer from old pool %s: %w", oldPoolAddr, err)
	}
	originalUnsiloedRebalancer := rebalancerReport.Output

	setRebalancerReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
		Args:          timelockAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to set unsiloed rebalancer on old pool %s: %w", oldPoolAddr, err)
	}
	ops = append(ops, setRebalancerReport.Output)

	type chainRebalancerInfo struct {
		chainSelector      uint64
		originalRebalancer common.Address
		isSiloed           bool
	}
	var siloInfos []chainRebalancerInfo

	for _, remoteChain := range supportedChains {
		if isSiloedByChain[remoteChain] {
			chainRebalancerReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.GetChainRebalancer, evmChain, evm_contract.FunctionInput[uint64]{
				ChainSelector: chainSel,
				Address:       oldPoolAddr,
				Args:          remoteChain,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get chain rebalancer for chain %d: %w", remoteChain, err)
			}

			setSiloReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.SetSiloRebalancer, evmChain, evm_contract.FunctionInput[siloed_ops_v161.SetSiloRebalancerArgs]{
				ChainSelector: chainSel,
				Address:       oldPoolAddr,
				Args: siloed_ops_v161.SetSiloRebalancerArgs{
					RemoteChainSelector: remoteChain,
					NewRebalancer:       timelockAddr,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set silo rebalancer for chain %d: %w", remoteChain, err)
			}
			ops = append(ops, setSiloReport.Output)

			siloInfos = append(siloInfos, chainRebalancerInfo{
				chainSelector:      remoteChain,
				originalRebalancer: chainRebalancerReport.Output,
				isSiloed:           true,
			})
		} else {
			siloInfos = append(siloInfos, chainRebalancerInfo{
				chainSelector: remoteChain,
				isSiloed:      false,
			})
		}
	}

	var firstLockbox common.Address
	for _, info := range siloInfos {
		if !info.isSiloed {
			continue
		}

		lockbox, ok := lockboxByChain[info.chainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("no lockbox configured for chain %d on new siloed pool", info.chainSelector)
		}
		if firstLockbox == (common.Address{}) {
			firstLockbox = lockbox
		}

		availableReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.GetAvailableTokens, evmChain, evm_contract.FunctionInput[uint64]{
			ChainSelector: chainSel,
			Address:       oldPoolAddr,
			Args:          info.chainSelector,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get available tokens for chain %d: %w", info.chainSelector, err)
		}

		siloBalance := availableReport.Output
		siloAmount := computeAmount(siloBalance, input)
		if siloAmount.Sign() == 0 {
			continue
		}

		withdrawReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.WithdrawSiloedLiquidity, evmChain, evm_contract.FunctionInput[siloed_ops_v161.WithdrawSiloedLiquidityArgs]{
			ChainSelector: chainSel,
			Address:       oldPoolAddr,
			Args: siloed_ops_v161.WithdrawSiloedLiquidityArgs{
				RemoteChainSelector: info.chainSelector,
				Amount:              siloAmount,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw siloed liquidity for chain %d: %w", info.chainSelector, err)
		}
		ops = append(ops, withdrawReport.Output)

		siloTransferReport, err := cldf_ops.ExecuteOperation(b, erc20_ops.Transfer, evmChain, evm_contract.FunctionInput[erc20_ops.TransferArgs]{
			ChainSelector: chainSel,
			Address:       tokenAddr,
			Args: erc20_ops.TransferArgs{
				Receiver: lockbox,
				Amount:   siloAmount,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer siloed liquidity to lockbox for chain %d: %w", info.chainSelector, err)
		}
		ops = append(ops, siloTransferReport.Output)
	}

	unsiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.GetUnsiloedLiquidity, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get unsiloed liquidity from old pool %s: %w", oldPoolAddr, err)
	}
	unsiloedBalance := unsiloedReport.Output
	unsiloedAmount := computeAmount(unsiloedBalance, input)

	if unsiloedAmount.Sign() > 0 {
		depositLockbox := firstLockbox
		if depositLockbox == (common.Address{}) {
			for _, lb := range lockboxByChain {
				depositLockbox = lb
				break
			}
		}
		if depositLockbox == (common.Address{}) {
			return sequences.OnChainOutput{}, fmt.Errorf("no lockbox available for unsiloed liquidity deposit")
		}

		withdrawUnsiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.WithdrawLiquidity, evmChain, evm_contract.FunctionInput[*big.Int]{
			ChainSelector: chainSel,
			Address:       oldPoolAddr,
			Args:          unsiloedAmount,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw unsiloed liquidity: %w", err)
		}
		ops = append(ops, withdrawUnsiloedReport.Output)

		unsiloedTransferReport, err := cldf_ops.ExecuteOperation(b, erc20_ops.Transfer, evmChain, evm_contract.FunctionInput[erc20_ops.TransferArgs]{
			ChainSelector: chainSel,
			Address:       tokenAddr,
			Args: erc20_ops.TransferArgs{
				Receiver: depositLockbox,
				Amount:   unsiloedAmount,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer unsiloed liquidity to lockbox: %w", err)
		}
		ops = append(ops, unsiloedTransferReport.Output)
	}

	for _, info := range siloInfos {
		if info.isSiloed {
			restoreReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.SetSiloRebalancer, evmChain, evm_contract.FunctionInput[siloed_ops_v161.SetSiloRebalancerArgs]{
				ChainSelector: chainSel,
				Address:       oldPoolAddr,
				Args: siloed_ops_v161.SetSiloRebalancerArgs{
					RemoteChainSelector: info.chainSelector,
					NewRebalancer:       info.originalRebalancer,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to restore silo rebalancer for chain %d: %w", info.chainSelector, err)
			}
			ops = append(ops, restoreReport.Output)
		}
	}

	restoreUnsiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
		Args:          originalUnsiloedRebalancer,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to restore unsiloed rebalancer: %w", err)
	}
	ops = append(ops, restoreUnsiloedReport.Output)

	if input.SetPoolConfig != nil {
		ops, err = appendSetPool(b, evmChain, chainSel, input.SetPoolConfig, newPoolAddr, ops)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
	}

	batchOp, err := evm_contract.NewBatchOperationFromWrites(ops)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
	}

	return sequences.OnChainOutput{
		BatchOps: []mcms_types.BatchOperation{batchOp},
	}, nil
}

func appendSetRebalancerAndWithdraw(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	chainSel uint64,
	oldPoolAddr, timelockAddr common.Address,
	amount *big.Int,
	ops []evm_contract.WriteOutput,
) ([]evm_contract.WriteOutput, error) {
	setRebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops_v161.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
		Args:          timelockAddr,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set rebalancer on old pool %s: %w", oldPoolAddr, err)
	}
	ops = append(ops, setRebalancerReport.Output)

	withdrawReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops_v161.WithdrawLiquidity, evmChain, evm_contract.FunctionInput[*big.Int]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
		Args:          amount,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to withdraw liquidity from old pool %s: %w", oldPoolAddr, err)
	}
	ops = append(ops, withdrawReport.Output)

	return ops, nil
}

func appendSetPool(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	chainSel uint64,
	config *tokens.MigrationSetPoolConfig,
	newPoolAddr common.Address,
	ops []evm_contract.WriteOutput,
) ([]evm_contract.WriteOutput, error) {
	setPoolReport, err := cldf_ops.ExecuteOperation(b, tar_ops.SetPool, evmChain, evm_contract.FunctionInput[tar_ops.SetPoolArgs]{
		ChainSelector: chainSel,
		Address:       common.HexToAddress(config.RegistryAddress),
		Args: tar_ops.SetPoolArgs{
			TokenAddress:     common.HexToAddress(config.TokenAddress),
			TokenPoolAddress: newPoolAddr,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set pool on token admin registry: %w", err)
	}
	ops = append(ops, setPoolReport.Output)

	return ops, nil
}
