package adapters

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	erc20_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	lrtp_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/lock_release_token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// MigrateLockReleasePoolLiquiditySequence returns the SOURCE half of a lock-release liquidity
// migration for v1.5.1 pools. It extracts liquidity from the old pool to the timelock with the v1.5.1
// bindings and records the moved amount in input.Resolved for the v2.0.0 sink.
func (t *TokenAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokensapi.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.5.1-adapter:migrate-lock-release-pool-liquidity-source",
		semver.MustParse("1.5.1"),
		"Extract lock-release liquidity from a v1.5.1 pool to the timelock",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.MigrateLockReleasePoolLiquidityInput) (sequences.OnChainOutput, error) {
			evmChain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
			}
			chainSel := input.ChainSelector
			oldPool := common.HexToAddress(input.OldPoolAddress)
			newPool := common.HexToAddress(input.NewPoolAddress)
			timelock := common.HexToAddress(input.TimelockAddress)

			oldTV, _, err := t.ValidateLockReleaseMigrationPoolPair(b, evmChain, oldPool, newPool)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			// The adapter may have been resolved from a datastore version; only drive pools whose
			// on-chain version actually is v1.5.1.
			if !oldTV.Version.Equal(cciputils.Version_1_5_1) {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"old pool %s is %s %s on-chain; the v1.5.1 liquidity migration source only supports v1.5.1 pools",
					oldPool, oldTV.Type, oldTV.Version)
			}
			if input.IsExactSiloMode() {
				return sequences.OnChainOutput{}, fmt.Errorf("SiloExactAmounts/UnsiloedExactAmount are only supported for siloed pools")
			}

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

			if err := t.RequireTimelockOwnsOldPool(b, evmChain, oldPool, timelock); err != nil {
				return sequences.OnChainOutput{}, err
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

			withdrawReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops.WithdrawLiquidityProposalOnly, evmChain, evm_contract.FunctionInput[*big.Int]{
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
		},
	)
}
