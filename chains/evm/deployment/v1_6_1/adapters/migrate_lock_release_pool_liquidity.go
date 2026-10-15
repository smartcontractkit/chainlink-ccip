package adapters

import (
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	erc20_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	lrtp_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/lock_release_token_pool"
	siloed_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/siloed_lock_release_token_pool"
	siloed_sink_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/siloed_lock_release_token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// MigrateLockReleasePoolLiquiditySequence returns the SOURCE half of a lock-release liquidity
// migration for v1.6.1 pools (plain and siloed), also serving v1.6.2 which registers this adapter. It
// extracts liquidity from the old pool to the timelock with the v1.6.1 bindings and records the moved
// amounts in input.Resolved for the v2.0.0 sink.
func (t *TokenAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokensapi.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.6.1-adapter:migrate-lock-release-pool-liquidity-source",
		semver.MustParse("1.6.1"),
		"Extract lock-release liquidity from a v1.6.1 pool (plain or siloed) to the timelock",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.MigrateLockReleasePoolLiquidityInput) (sequences.OnChainOutput, error) {
			evmChain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
			}
			oldPool := common.HexToAddress(input.OldPoolAddress)
			newPool := common.HexToAddress(input.NewPoolAddress)
			timelock := common.HexToAddress(input.TimelockAddress)

			oldTV, _, err := t.ValidateLockReleaseMigrationPoolPair(b, evmChain, oldPool, newPool)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			// The adapter may have been resolved from a datastore version; only drive pools whose
			// on-chain version actually shares the v1.6.1 bindings.
			if !oldTV.Version.Equal(cciputils.Version_1_6_1) && !oldTV.Version.Equal(cciputils.Version_1_6_2) {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"old pool %s is %s %s on-chain; the v1.6.1 liquidity migration source only supports v1.6.1 and v1.6.2 pools",
					oldPool, oldTV.Type, oldTV.Version)
			}
			if string(oldTV.Type) == cciputils.SiloedLockReleaseTokenPool.String() {
				return t.migrateSiloedLockReleaseSource(b, evmChain, input, oldPool, newPool, timelock)
			}
			if input.IsExactSiloMode() {
				return sequences.OnChainOutput{}, fmt.Errorf("SiloExactAmounts/UnsiloedExactAmount are only supported for siloed pools")
			}
			return t.migratePlainLockReleaseSource(b, evmChain, input, oldPool, newPool, timelock)
		})
}

func (t *TokenAdapter) migratePlainLockReleaseSource(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	input tokensapi.MigrateLockReleasePoolLiquidityInput,
	oldPool, newPool, timelock common.Address,
) (sequences.OnChainOutput, error) {
	chainSel := input.ChainSelector

	tokenAddr, err := t.ResolveMigrationToken(b, evmChain, oldPool, newPool)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}

	balanceReport, err := cldf_ops.ExecuteOperation(b, erc20_ops.BalanceOf, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       tokenAddr,
		Args:          oldPool,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get balance of old pool %s: %w", oldPool, err)
	}
	balance := balanceReport.Output

	amount, err := input.ComputeAmount(balance)
	if err != nil {
		return sequences.OnChainOutput{}, err
	}
	if amount.Sign() == 0 {
		return sequences.OnChainOutput{}, fmt.Errorf("computed migration amount is zero")
	}
	if amount.Cmp(balance) > 0 {
		return sequences.OnChainOutput{}, fmt.Errorf("migration amount %s exceeds old pool balance %s", amount, balance)
	}

	rebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops.GetRebalancer, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       oldPool,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get rebalancer from old pool %s: %w", oldPool, err)
	}

	setRebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPool,
		Args:          timelock,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to set rebalancer on old pool %s: %w", oldPool, err)
	}

	withdrawReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops.WithdrawLiquidity, evmChain, evm_contract.FunctionInput[*big.Int]{
		ChainSelector: chainSel,
		Address:       oldPool,
		Args:          amount,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw liquidity from old pool %s: %w", oldPool, err)
	}

	restoreRebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPool,
		Args:          rebalancerReport.Output,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to restore rebalancer on old pool %s: %w", oldPool, err)
	}

	batchOp, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{
		setRebalancerReport.Output,
		withdrawReport.Output,
		restoreRebalancerReport.Output,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
	}

	if input.Resolved != nil {
		input.Resolved.Amount = amount
	}
	return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
}

func (t *TokenAdapter) migrateSiloedLockReleaseSource(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	input tokensapi.MigrateLockReleasePoolLiquidityInput,
	oldPool, newPool, timelock common.Address,
) (sequences.OnChainOutput, error) {
	chainSel := input.ChainSelector
	exactMode := input.IsExactSiloMode()

	if _, err := t.ResolveMigrationToken(b, evmChain, oldPool, newPool); err != nil {
		return sequences.OnChainOutput{}, err
	}

	chainsReport, err := cldf_ops.ExecuteOperation(b, siloed_ops.GetSupportedChains, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       oldPool,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get supported chains from old siloed pool %s: %w", oldPool, err)
	}
	supportedChains := chainsReport.Output

	// Resolve which of the old pool's chains are siloed once; used by every check below and by the
	// rebalancer handover.
	var siloedChains []uint64
	for _, remoteChain := range supportedChains {
		isSiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops.IsSiloed, evmChain, evm_contract.FunctionInput[uint64]{
			ChainSelector: chainSel,
			Address:       oldPool,
			Args:          remoteChain,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check if chain %d is siloed on old pool %s: %w", remoteChain, oldPool, err)
		}
		if isSiloedReport.Output {
			siloedChains = append(siloedChains, remoteChain)
		}
	}

	// In exact mode, every siloed chain must be explicit - a chain silently falling back to zero
	// migration is more likely to be an oversight than intent, so surface it before any write is
	// emitted rather than migrating a partial set of silos.
	if exactMode {
		// A ChainSelector that isn't actually a siloed chain of the old pool would otherwise be
		// silently ignored, leaving no trace in the execution logs that the entry was never applied -
		// fail loudly instead so a typo'd or stale selector surfaces immediately.
		var unknownSilos []uint64
		for _, sa := range input.SiloExactAmounts {
			if !slices.Contains(siloedChains, sa.ChainSelector) {
				unknownSilos = append(unknownSilos, sa.ChainSelector)
			}
		}
		if len(unknownSilos) > 0 {
			slices.Sort(unknownSilos)
			return sequences.OnChainOutput{}, fmt.Errorf(
				"SiloExactAmounts references chain selectors %v that are not siloed chains on old pool %s",
				unknownSilos, oldPool,
			)
		}

		var missingSilos []uint64
		for _, remoteChain := range siloedChains {
			if _, ok := input.SiloExactAmount(remoteChain); !ok {
				missingSilos = append(missingSilos, remoteChain)
			}
		}
		if len(missingSilos) > 0 {
			slices.Sort(missingSilos)
			return sequences.OnChainOutput{}, fmt.Errorf(
				"SiloExactAmounts is missing entries for siloed chains %v of old pool %s; exact mode requires every siloed chain to be explicit",
				missingSilos, oldPool,
			)
		}
	}

	unsiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops.GetUnsiloedLiquidity, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       oldPool,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get unsiloed liquidity from old pool %s: %w", oldPool, err)
	}

	var unsiloedAmount *big.Int
	if exactMode {
		if unsiloedReport.Output.Sign() > 0 && input.UnsiloedExactAmount == nil {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"old pool %s holds %s unsiloed liquidity to migrate but UnsiloedExactAmount was not set; "+
					"exact mode requires an explicit amount for the shared bucket",
				oldPool, unsiloedReport.Output,
			)
		}
		if input.UnsiloedExactAmount != nil {
			unsiloedAmount = new(big.Int).Set(input.UnsiloedExactAmount)
		} else {
			unsiloedAmount = big.NewInt(0)
		}
		if unsiloedAmount.Cmp(unsiloedReport.Output) > 0 {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"UnsiloedExactAmount %s exceeds old pool %s unsiloed balance %s",
				unsiloedAmount, oldPool, unsiloedReport.Output,
			)
		}
	} else {
		unsiloedAmount, err = input.ComputeAmount(unsiloedReport.Output)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
	}

	// Resolve every siloed chain's migration amount and validate it against the on-chain balance.
	siloAmounts := make([]tokensapi.SiloExactAmount, 0, len(siloedChains))
	for _, remoteChain := range siloedChains {
		availableReport, err := cldf_ops.ExecuteOperation(b, siloed_ops.GetAvailableTokens, evmChain, evm_contract.FunctionInput[uint64]{
			ChainSelector: chainSel,
			Address:       oldPool,
			Args:          remoteChain,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get available tokens for chain %d: %w", remoteChain, err)
		}
		siloBalance := availableReport.Output

		var siloAmount *big.Int
		if exactMode {
			amt, _ := input.SiloExactAmount(remoteChain) // presence guaranteed by the missing-silo check above
			siloAmount = new(big.Int).Set(amt)
		} else {
			siloAmount, err = input.ComputeAmount(siloBalance)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
		}
		if siloAmount.Cmp(siloBalance) > 0 {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"migration amount %s for chain %d exceeds silo balance %s", siloAmount, remoteChain, siloBalance,
			)
		}
		siloAmounts = append(siloAmounts, tokensapi.SiloExactAmount{ChainSelector: remoteChain, Amount: siloAmount})
	}

	// Check the new pool can receive what is about to be extracted BEFORE any write is emitted. This
	// reads the v2.0.0 sink pool: the framework executes a write immediately when the deployer is an
	// allowed caller, so the check cannot be left to the sink alone.
	if err := checkSiloedSinkCoverage(b, evmChain, input, oldPool, newPool, siloedChains, unsiloedAmount); err != nil {
		return sequences.OnChainOutput{}, err
	}

	var ops []evm_contract.WriteOutput

	rebalancerReport, err := cldf_ops.ExecuteOperation(b, siloed_ops.GetRebalancer, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       oldPool,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get unsiloed rebalancer from old pool %s: %w", oldPool, err)
	}
	originalUnsiloedRebalancer := rebalancerReport.Output

	setRebalancerReport, err := cldf_ops.ExecuteOperation(b, siloed_ops.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPool,
		Args:          timelock,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to set unsiloed rebalancer on old pool %s: %w", oldPool, err)
	}
	ops = append(ops, setRebalancerReport.Output)

	originalSiloRebalancers := make(map[uint64]common.Address, len(siloedChains))
	for _, remoteChain := range siloedChains {
		chainRebalancerReport, err := cldf_ops.ExecuteOperation(b, siloed_ops.GetChainRebalancer, evmChain, evm_contract.FunctionInput[uint64]{
			ChainSelector: chainSel,
			Address:       oldPool,
			Args:          remoteChain,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get chain rebalancer for chain %d: %w", remoteChain, err)
		}
		originalSiloRebalancers[remoteChain] = chainRebalancerReport.Output

		setSiloReport, err := cldf_ops.ExecuteOperation(b, siloed_ops.SetSiloRebalancer, evmChain, evm_contract.FunctionInput[siloed_ops.SetSiloRebalancerArgs]{
			ChainSelector: chainSel,
			Address:       oldPool,
			Args: siloed_ops.SetSiloRebalancerArgs{
				RemoteChainSelector: remoteChain,
				NewRebalancer:       timelock,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set silo rebalancer for chain %d: %w", remoteChain, err)
		}
		ops = append(ops, setSiloReport.Output)
	}

	for _, sa := range siloAmounts {
		if sa.Amount.Sign() == 0 {
			continue
		}
		withdrawReport, err := cldf_ops.ExecuteOperation(b, siloed_ops.WithdrawSiloedLiquidity, evmChain, evm_contract.FunctionInput[siloed_ops.WithdrawSiloedLiquidityArgs]{
			ChainSelector: chainSel,
			Address:       oldPool,
			Args: siloed_ops.WithdrawSiloedLiquidityArgs{
				RemoteChainSelector: sa.ChainSelector,
				Amount:              sa.Amount,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw siloed liquidity for chain %d: %w", sa.ChainSelector, err)
		}
		ops = append(ops, withdrawReport.Output)
	}

	if unsiloedAmount.Sign() > 0 {
		withdrawUnsiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops.WithdrawLiquidity, evmChain, evm_contract.FunctionInput[*big.Int]{
			ChainSelector: chainSel,
			Address:       oldPool,
			Args:          unsiloedAmount,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw unsiloed liquidity: %w", err)
		}
		ops = append(ops, withdrawUnsiloedReport.Output)
	}

	for _, remoteChain := range siloedChains {
		restoreReport, err := cldf_ops.ExecuteOperation(b, siloed_ops.SetSiloRebalancer, evmChain, evm_contract.FunctionInput[siloed_ops.SetSiloRebalancerArgs]{
			ChainSelector: chainSel,
			Address:       oldPool,
			Args: siloed_ops.SetSiloRebalancerArgs{
				RemoteChainSelector: remoteChain,
				NewRebalancer:       originalSiloRebalancers[remoteChain],
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to restore silo rebalancer for chain %d: %w", remoteChain, err)
		}
		ops = append(ops, restoreReport.Output)
	}

	restoreUnsiloedReport, err := cldf_ops.ExecuteOperation(b, siloed_ops.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSel,
		Address:       oldPool,
		Args:          originalUnsiloedRebalancer,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to restore unsiloed rebalancer: %w", err)
	}
	ops = append(ops, restoreUnsiloedReport.Output)

	batchOp, err := evm_contract.NewBatchOperationFromWrites(ops)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
	}

	if input.Resolved != nil {
		input.Resolved.SiloAmounts = siloAmounts
		input.Resolved.UnsiloedAmount = unsiloedAmount
	}
	return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
}

// checkSiloedSinkCoverage fails if the new siloed pool has no lockbox for a siloed chain being
// migrated, or if unsiloed liquidity is being migrated without a valid destination lockbox.
func checkSiloedSinkCoverage(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	input tokensapi.MigrateLockReleasePoolLiquidityInput,
	oldPool, newPool common.Address,
	siloedChains []uint64,
	unsiloedAmount *big.Int,
) error {
	lockboxConfigsReport, err := cldf_ops.ExecuteOperation(b, siloed_sink_ops.GetAllLockBoxConfigs, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: input.ChainSelector,
		Address:       newPool,
	})
	if err != nil {
		return fmt.Errorf("failed to get lockbox configs from new pool %s: %w", newPool, err)
	}

	lockboxByChain := make(map[uint64]common.Address)
	configuredLockBoxes := make(map[common.Address]bool)
	for _, config := range lockboxConfigsReport.Output {
		if config.LockBox == (common.Address{}) {
			continue
		}
		lockboxByChain[config.RemoteChainSelector] = config.LockBox
		configuredLockBoxes[config.LockBox] = true
	}

	var chainsWithoutLockBox []uint64
	for _, remoteChain := range siloedChains {
		if _, ok := lockboxByChain[remoteChain]; !ok {
			chainsWithoutLockBox = append(chainsWithoutLockBox, remoteChain)
		}
	}
	if len(chainsWithoutLockBox) > 0 {
		slices.Sort(chainsWithoutLockBox)
		return fmt.Errorf(
			"new siloed pool %s has no lockbox configured for siloed chains %v of old pool %s; the new pool's lockBoxGroups must cover every siloed chain being migrated",
			newPool, chainsWithoutLockBox, oldPool,
		)
	}

	if input.UnsiloedLockBoxAddress != "" {
		lockBox := common.HexToAddress(input.UnsiloedLockBoxAddress)
		if lockBox == (common.Address{}) {
			return fmt.Errorf("UnsiloedLockBoxAddress %q is not a valid address", input.UnsiloedLockBoxAddress)
		}
		if !configuredLockBoxes[lockBox] {
			return fmt.Errorf(
				"UnsiloedLockBoxAddress %s is not one of the lockboxes configured on new pool %s",
				lockBox, newPool,
			)
		}
	} else if unsiloedAmount.Sign() > 0 {
		return fmt.Errorf(
			"old pool %s holds %s unsiloed liquidity to migrate but UnsiloedLockBoxAddress was not set; "+
				"the shared balance backs the pool's non-siloed chains and its destination cannot be inferred - "+
				"set it to the lockbox serving those chains on new pool %s",
			oldPool, unsiloedAmount, newPool,
		)
	}
	return nil
}
