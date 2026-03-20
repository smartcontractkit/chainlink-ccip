package tokens

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	mcms_types "github.com/smartcontractkit/mcms/types"

	erc20_ops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/latest/operations/erc20"
	lockbox_ops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/erc20_lock_box"
	lrtp_ops_v170 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/lock_release_token_pool"
	siloed_lrtp_ops_v170 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/siloed_lock_release_token_pool"
	token_pool_ops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/token_pool"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	type_and_version "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	tar_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	lrtp_ops_v161 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/lock_release_token_pool"
	siloed_ops_v161 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/siloed_lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
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

		isSiloed := strings.Contains(oldPoolType, "Siloed")

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

	ops, err = appendAuthApproveDeposit(b, evmChain, chainSel, lockboxAddr, tokenAddr, timelockAddr, amount, 0, ops)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	ops, err = appendCleanup(b, evmChain, chainSel, lockboxAddr, oldPoolAddr, timelockAddr, originalRebalancer, ops)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

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
		lockboxByChain[config.RemoteChainSelector] = config.LockBox
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
		isSiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops_v161.IsSiloed, evmChain, evm_contract.FunctionInput[uint64]{
			ChainSelector: chainSel,
			Address:       oldPoolAddr,
			Args:          remoteChain,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check if chain %d is siloed on old pool %s: %w", remoteChain, oldPoolAddr, err)
		}

		if isSiloedReport.Output {
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
	usedLockboxes := make(map[common.Address]bool)
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

		ops, err = appendAuthApproveDeposit(b, evmChain, chainSel, lockbox, tokenAddr, timelockAddr, siloAmount, info.chainSelector, ops)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build deposit ops for chain %d: %w", info.chainSelector, err)
		}
		usedLockboxes[lockbox] = true
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

		ops, err = appendAuthApproveDeposit(b, evmChain, chainSel, depositLockbox, tokenAddr, timelockAddr, unsiloedAmount, 0, ops)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to build deposit ops for unsiloed liquidity: %w", err)
		}
		usedLockboxes[depositLockbox] = true
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

	for lb := range usedLockboxes {
		removeAuthReport, err := cldf_ops.ExecuteOperation(b, lockbox_ops.ApplyAuthorizedCallerUpdates, evmChain, evm_contract.FunctionInput[lockbox_ops.AuthorizedCallerArgs]{
			ChainSelector: chainSel,
			Address:       lb,
			Args: lockbox_ops.AuthorizedCallerArgs{
				AddedCallers:   []common.Address{},
				RemovedCallers: []common.Address{timelockAddr},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to remove timelock from lockbox %s authorized callers: %w", lb, err)
		}
		ops = append(ops, removeAuthReport.Output)
	}

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

func appendAuthApproveDeposit(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	chainSel uint64,
	lockboxAddr, tokenAddr, timelockAddr common.Address,
	amount *big.Int,
	remoteChainSelector uint64,
	ops []evm_contract.WriteOutput,
) ([]evm_contract.WriteOutput, error) {
	addAuthReport, err := cldf_ops.ExecuteOperation(b, lockbox_ops.ApplyAuthorizedCallerUpdates, evmChain, evm_contract.FunctionInput[lockbox_ops.AuthorizedCallerArgs]{
		ChainSelector: chainSel,
		Address:       lockboxAddr,
		Args: lockbox_ops.AuthorizedCallerArgs{
			AddedCallers:   []common.Address{timelockAddr},
			RemovedCallers: []common.Address{},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to add timelock as authorized caller on lockbox %s: %w", lockboxAddr, err)
	}
	ops = append(ops, addAuthReport.Output)

	approveReport, err := cldf_ops.ExecuteOperation(b, erc20_ops.Approve, evmChain, evm_contract.FunctionInput[erc20_ops.ApproveArgs]{
		ChainSelector: chainSel,
		Address:       tokenAddr,
		Args: erc20_ops.ApproveArgs{
			Spender: lockboxAddr,
			Amount:  amount,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to approve lockbox %s to spend tokens: %w", lockboxAddr, err)
	}
	ops = append(ops, approveReport.Output)

	depositReport, err := cldf_ops.ExecuteOperation(b, lockbox_ops.Deposit, evmChain, evm_contract.FunctionInput[lockbox_ops.DepositArgs]{
		ChainSelector: chainSel,
		Address:       lockboxAddr,
		Args: lockbox_ops.DepositArgs{
			Token:               tokenAddr,
			RemoteChainSelector: remoteChainSelector,
			Amount:              amount,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to deposit into lockbox %s: %w", lockboxAddr, err)
	}
	ops = append(ops, depositReport.Output)

	return ops, nil
}

func appendCleanup(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	chainSel uint64,
	lockboxAddr, oldPoolAddr, timelockAddr, originalRebalancer common.Address,
	ops []evm_contract.WriteOutput,
) ([]evm_contract.WriteOutput, error) {
	removeAuthReport, err := cldf_ops.ExecuteOperation(b, lockbox_ops.ApplyAuthorizedCallerUpdates, evmChain, evm_contract.FunctionInput[lockbox_ops.AuthorizedCallerArgs]{
		ChainSelector: chainSel,
		Address:       lockboxAddr,
		Args: lockbox_ops.AuthorizedCallerArgs{
			AddedCallers:   []common.Address{},
			RemovedCallers: []common.Address{timelockAddr},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to remove timelock as authorized caller on lockbox %s: %w", lockboxAddr, err)
	}
	ops = append(ops, removeAuthReport.Output)

	restoreRebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops_v161.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPoolAddr,
		Args:          originalRebalancer,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to restore rebalancer on old pool %s: %w", oldPoolAddr, err)
	}
	ops = append(ops, restoreRebalancerReport.Output)

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
