package adapters

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	erc20_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	lrtp_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/lock_release_token_pool"
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
// migration for v1.5.0 pools, plain and *AndProxy. It extracts liquidity from the effective holder
// to the timelock and records the moved amount in input.Resolved for the v2.0.0 sink.
//
// A plain pool always holds its own liquidity. A LockReleaseTokenPoolAndProxy routes every
// lock/release through its previous pool when getPreviousPool() is set, so the liquidity may sit in
// either contract:
//   - only the proxy holds a balance: it is extracted with the v1.5.0 bindings (rebalancer take-over,
//     withdraw, restore), whether or not a previous pool is set;
//   - only the previous pool holds a balance: the migration is delegated to the source of the
//     previous pool's on-chain version (v1.4.0 or v1.2.0), which runs its own checks against it;
//   - both hold a balance (e.g. the previous pool was set after the proxy had locked tokens, or
//     someone transferred tokens to the proxy): the amount inputs are ambiguous, so the migration
//     fails. Naming the proxy keeps failing until one balance is zero; to recover, name the
//     previous pool directly, migrate its entire balance, execute that proposal, then rerun naming
//     the proxy. A previous pool that is itself an *AndProxy also fails, before any balance
//     decides the holder.
//
// getPreviousPool() is only called on the *AndProxy type; it does not exist on a plain pool.
func (t *TokenAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokensapi.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.5.0-adapter:migrate-lock-release-pool-liquidity-source",
		semver.MustParse("1.5.0"),
		"Extract lock-release liquidity from a v1.5.0 pool, or the previous pool behind an *AndProxy, to the timelock",
		func(b cldf_ops.Bundle, chains cldf_chain.BlockChains, input tokensapi.MigrateLockReleasePoolLiquidityInput) (sequences.OnChainOutput, error) {
			evmChain, ok := chains.EVMChains()[input.ChainSelector]
			if !ok {
				return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
			}
			oldPool := common.HexToAddress(input.OldPoolAddress)
			newPool := common.HexToAddress(input.NewPoolAddress)

			oldTV, _, err := t.ValidateLockReleaseMigrationPoolPair(b, evmChain, oldPool, newPool)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			// The adapter may have been resolved from a datastore version; only drive pools whose
			// on-chain version actually is v1.5.0 (plain or *AndProxy).
			if !oldTV.Version.Equal(cciputils.Version_1_5_0) {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"old pool %s is %s %s on-chain; the v1.5.0 liquidity migration source only supports v1.5.0 pools",
					oldPool, oldTV.Type, oldTV.Version)
			}
			if input.IsExactSiloMode() {
				return sequences.OnChainOutput{}, fmt.Errorf("SiloExactAmounts/UnsiloedExactAmount are only supported for siloed pools")
			}

			tokenAddr, err := t.ResolveMigrationToken(b, evmChain, oldPool, newPool)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			oldPoolBalance, err := tokenBalance(b, evmChain, tokenAddr, oldPool)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get balance of old pool %s: %w", oldPool, err)
			}

			if oldTV.Type.String() != cciputils.LockReleaseTokenPoolAndProxy.String() {
				return t.migrateV150Holder(b, evmChain, input, oldPool, oldPoolBalance)
			}

			previous, err := previousPool(evmChain, oldPool, nil)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to read getPreviousPool from proxy %s: %w", oldPool, err)
			}
			if previous == (common.Address{}) {
				return t.migrateV150Holder(b, evmChain, input, oldPool, oldPoolBalance)
			}

			previousVersion, previousType, err := previousTypeAndVersion(b, evmChain, previous)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}
			if previousType == cciputils.LockReleaseTokenPoolAndProxy.String() || previousType == cciputils.BurnMintTokenPoolAndProxy.String() {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"previous pool %s of proxy %s is itself a %s; nested proxies are not supported",
					previous, oldPool, previousType)
			}
			previousBalance, err := tokenBalance(b, evmChain, tokenAddr, previous)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get balance of previous pool %s of proxy %s: %w", previous, oldPool, err)
			}

			switch {
			case oldPoolBalance.Sign() > 0 && previousBalance.Sign() > 0:
				// Routing is all-or-nothing on getPreviousPool(), so at most one holder is expected.
				// With both funded, an Amount or BasisPoints input has no single balance to apply to.
				// Naming the proxy always lands here while both are funded, so the error spells out the
				// order that clears it.
				return sequences.OnChainOutput{}, fmt.Errorf(
					"liquidity is split between proxy %s (%s) and its previous pool %s (%s); first migrate the previous pool's entire balance by naming %s as the old pool and execute that proposal, then rerun naming the proxy",
					oldPool, oldPoolBalance, previous, previousBalance, previous)
			case oldPoolBalance.Sign() > 0:
				return t.migrateV150Holder(b, evmChain, input, oldPool, oldPoolBalance)
			case previousBalance.Sign() > 0:
				return delegateToPreviousPool(b, chains, input, oldPool, previous, previousVersion, previousType)
			default:
				return sequences.OnChainOutput{}, fmt.Errorf(
					"neither proxy %s nor its previous pool %s holds any liquidity to migrate", oldPool, previous)
			}
		},
	)
}

// delegateToPreviousPool runs the liquidity migration source registered for the previous pool's
// on-chain version against the previous pool. Only OldPoolAddress changes: Resolved is a pointer,
// so the amount the delegated source resolves reaches the sink unchanged.
//
// The adapter is resolved from the global registry, the one the MigrateLockReleasePoolLiquidity
// changeset resolves this source from.
func delegateToPreviousPool(
	b cldf_ops.Bundle,
	chains cldf_chain.BlockChains,
	input tokensapi.MigrateLockReleasePoolLiquidityInput,
	proxy, previous common.Address,
	previousVersion *semver.Version,
	previousType string,
) (sequences.OnChainOutput, error) {
	adapter, _, err := tokensapi.ResolveAdapter(tokensapi.GetTokenAdapterRegistry(), input.ChainSelector, previousVersion)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf(
			"previous pool %s (%s %s) of proxy %s does not support liquidity migration yet: %w",
			previous, previousType, previousVersion, proxy, err)
	}
	seq := adapter.MigrateLockReleasePoolLiquiditySequence()
	if seq == nil {
		return sequences.OnChainOutput{}, fmt.Errorf(
			"previous pool %s (%s %s) of proxy %s does not support liquidity migration yet",
			previous, previousType, previousVersion, proxy)
	}

	previousInput := input
	previousInput.OldPoolAddress = previous.Hex()
	report, err := cldf_ops.ExecuteSequence(b, seq, chains, previousInput)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("previous pool %s of proxy %s: %w", previous, proxy, err)
	}
	return report.Output, nil
}

// migrateV150Holder extracts liquidity held by the v1.5.0 pool at holder itself: set the timelock as
// rebalancer, withdraw the amount to it, and restore the original rebalancer. withdrawLiquidity is
// rebalancer-gated, so it uses the proposal-only op and is always batched; the take-over and restore
// are skipped when the timelock already is the rebalancer.
func (t *TokenAdapter) migrateV150Holder(
	b cldf_ops.Bundle,
	evmChain evm.Chain,
	input tokensapi.MigrateLockReleasePoolLiquidityInput,
	holder common.Address,
	balance *big.Int,
) (sequences.OnChainOutput, error) {
	chainSel := input.ChainSelector
	timelock := common.HexToAddress(input.TimelockAddress)

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

	if err := t.RequireTimelockOwnsOldPool(b, evmChain, holder, timelock); err != nil {
		return sequences.OnChainOutput{}, err
	}

	rebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops.GetRebalancer, evmChain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSel,
		Address:       holder,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to get rebalancer from old pool %s: %w", holder, err)
	}

	// Skip the take-over and the restore when the timelock already is the rebalancer: both
	// would be no-op writes in the proposal.
	originalRebalancer := rebalancerReport.Output
	handOver := originalRebalancer != timelock

	var writes []evm_contract.WriteOutput
	if handOver {
		setRebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: chainSel,
			Address:       holder,
			Args:          timelock,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set rebalancer on old pool %s: %w", holder, err)
		}
		writes = append(writes, setRebalancerReport.Output)
	}

	withdrawReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops.WithdrawLiquidityProposalOnly, evmChain, evm_contract.FunctionInput[*big.Int]{
		ChainSelector: chainSel,
		Address:       holder,
		Args:          amount,
	})
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to withdraw liquidity from old pool %s: %w", holder, err)
	}
	writes = append(writes, withdrawReport.Output)

	if handOver {
		restoreRebalancerReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops.SetRebalancer, evmChain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: chainSel,
			Address:       holder,
			Args:          originalRebalancer,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to restore rebalancer on old pool %s: %w", holder, err)
		}
		writes = append(writes, restoreRebalancerReport.Output)
	}

	batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
	if err != nil {
		return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
	}

	if input.Resolved != nil {
		input.Resolved.Amount = amount
	}
	return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
}

// tokenBalance reads token.balanceOf(holder).
func tokenBalance(b cldf_ops.Bundle, evmChain evm.Chain, token, holder common.Address) (*big.Int, error) {
	report, err := cldf_ops.ExecuteOperation(b, erc20_ops.BalanceOf, evmChain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: evmChain.Selector,
		Address:       token,
		Args:          holder,
	})
	if err != nil {
		return nil, err
	}
	return report.Output, nil
}
