package adapters

import (
	"fmt"
	"math/big"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	erc20_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	lrtp_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/lock_release_token_pool"
	tp "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/token_pool"
	rmn_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	cciputils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// MigrateLockReleasePoolLiquiditySequence returns the SOURCE half of a lock-release liquidity
// migration for v1.2.0 pools. It extracts liquidity from the old pool to the timelock and records the
// moved amount in input.Resolved for the v2.0.0 sink.
//
// A v1.2.0 pool has no rebalancer and no admin withdrawal: removeLiquidity only returns a provider's
// own credit, and tokens locked through lockOrBurn (e.g. forwarded by the fronting v1.5.0 proxy) are
// credited to no provider. The only function that moves them is the onlyOffRamp releaseOrMint, so
// this source rescues them through it: register the timelock as an offRamp with a disabled rate
// limiter, release the amount to the timelock, then deregister it. When the timelock already is an
// offRamp it is left registered and only the release is emitted.
//
// It runs directly when the old pool named is the v1.2.0 pool, or delegated from the v1.5.0 source
// when the old pool named is the LockReleaseTokenPoolAndProxy in front of it; either way every check
// and write targets the v1.2.0 pool.
//
// Fails fast, before emitting anything, when:
//   - the pool can accept liquidity: releaseOrMint does not decrement provider credits, so it would
//     release provider funds and strand their claims. canAcceptLiquidity is immutable, so false
//     proves no provider credit was ever added;
//   - the pool's ARM is cursed: releaseOrMint is whenHealthy and would revert;
//   - the timelock does not own the pool: applyRampUpdates is onlyOwner, and all three writes must
//     land in the timelock batch in order;
//   - the timelock already is an offRamp whose enabled rate limiter cannot cover the amount.
func (t *TokenAdapter) MigrateLockReleasePoolLiquiditySequence() *cldf_ops.Sequence[tokensapi.MigrateLockReleasePoolLiquidityInput, sequences.OnChainOutput, cldf_chain.BlockChains] {
	return cldf_ops.NewSequence(
		"evm-v1.2.0-adapter:migrate-lock-release-pool-liquidity-source",
		semver.MustParse("1.2.0"),
		"Rescue lock-release liquidity from a v1.2.0 pool to the timelock via releaseOrMint",
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
			// on-chain version actually is v1.2.0.
			if !oldTV.Version.Equal(cciputils.Version_1_2_0) {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"old pool %s is %s %s on-chain; the v1.2.0 liquidity migration source only supports v1.2.0 pools",
					oldPool, oldTV.Type, oldTV.Version)
			}
			if input.IsExactSiloMode() {
				return sequences.OnChainOutput{}, fmt.Errorf("SiloExactAmounts/UnsiloedExactAmount are only supported for siloed pools")
			}

			tokenAddr, err := t.ResolveMigrationToken(b, evmChain, oldPool, newPool)
			if err != nil {
				return sequences.OnChainOutput{}, err
			}

			canAcceptReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops.CanAcceptLiquidity, evmChain, evm_contract.FunctionInput[struct{}]{
				ChainSelector: chainSel,
				Address:       oldPool,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to read canAcceptLiquidity from old pool %s: %w", oldPool, err)
			}
			if canAcceptReport.Output {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"v1.2.0 pool %s accepts liquidity, so provider credits may exist; releaseOrMint would release provider funds without decrementing their balances",
					oldPool)
			}

			armProxyReport, err := cldf_ops.ExecuteOperation(b, tp.GetArmProxy, evmChain, evm_contract.FunctionInput[struct{}]{
				ChainSelector: chainSel,
				Address:       oldPool,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get ARM proxy of old pool %s: %w", oldPool, err)
			}
			// isCursed() has the same selector on every ARM/RMN generation, and the ARM proxy
			// forwards it, so the latest RMN read serves whichever ARM the pool points at.
			cursedReport, err := cldf_ops.ExecuteOperation(b, rmn_ops.IsCursed0, evmChain, evm_contract.FunctionInput[struct{}]{
				ChainSelector: chainSel,
				Address:       armProxyReport.Output,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to read curse status from ARM proxy %s of old pool %s: %w", armProxyReport.Output, oldPool, err)
			}
			if cursedReport.Output {
				return sequences.OnChainOutput{}, fmt.Errorf(
					"ARM proxy %s of old pool %s reports a curse; releaseOrMint is whenHealthy and would revert",
					armProxyReport.Output, oldPool)
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

			isOffRampReport, err := cldf_ops.ExecuteOperation(b, tp.IsOffRamp, evmChain, evm_contract.FunctionInput[common.Address]{
				ChainSelector: chainSel,
				Address:       oldPool,
				Args:          timelock,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to check whether the timelock is an offRamp of old pool %s: %w", oldPool, err)
			}
			// Registering an existing offRamp reverts (RampAlreadyExists), so an already registered
			// timelock is left alone: no registration and no deregistration. Its limiter then applies
			// to the release; tokens only refill until execution, so the current level is a safe bound.
			register := !isOffRampReport.Output
			if !register {
				bucketReport, err := cldf_ops.ExecuteOperation(b, tp.CurrentOffRampRateLimiterState, evmChain, evm_contract.FunctionInput[common.Address]{
					ChainSelector: chainSel,
					Address:       oldPool,
					Args:          timelock,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to read the timelock offRamp rate limiter of old pool %s: %w", oldPool, err)
				}
				if bucket := bucketReport.Output; bucket.IsEnabled && bucket.Tokens.Cmp(amount) < 0 {
					return sequences.OnChainOutput{}, fmt.Errorf(
						"timelock %s already is an offRamp of old pool %s with an enabled rate limiter holding %s tokens, less than the migration amount %s",
						timelock, oldPool, bucket.Tokens, amount)
				}
			}

			var writes []evm_contract.WriteOutput
			if register {
				// isEnabled=false: the release does not consume any limiter capacity.
				addRampReport, err := cldf_ops.ExecuteOperation(b, tp.ApplyRampUpdates, evmChain, evm_contract.FunctionInput[tp.ApplyRampUpdatesArgs]{
					ChainSelector: chainSel,
					Address:       oldPool,
					Args: tp.ApplyRampUpdatesArgs{
						OffRamps: []tp.RampUpdate{{
							Ramp:              timelock,
							Allowed:           true,
							RateLimiterConfig: tp.Config{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
						}},
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to register the timelock as an offRamp of old pool %s: %w", oldPool, err)
				}
				writes = append(writes, addRampReport.Output)
			}

			releaseReport, err := cldf_ops.ExecuteOperation(b, lrtp_ops.ReleaseOrMintProposalOnly, evmChain, evm_contract.FunctionInput[lrtp_ops.ReleaseOrMintArgs]{
				ChainSelector: chainSel,
				Address:       oldPool,
				Args: lrtp_ops.ReleaseOrMintArgs{
					Receiver: timelock,
					Amount:   amount,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to release liquidity from old pool %s: %w", oldPool, err)
			}
			writes = append(writes, releaseReport.Output)

			if register {
				removeRampReport, err := cldf_ops.ExecuteOperation(b, tp.ApplyRampUpdates, evmChain, evm_contract.FunctionInput[tp.ApplyRampUpdatesArgs]{
					ChainSelector: chainSel,
					Address:       oldPool,
					Args: tp.ApplyRampUpdatesArgs{
						OffRamps: []tp.RampUpdate{{
							Ramp:              timelock,
							Allowed:           false,
							RateLimiterConfig: tp.Config{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
						}},
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to deregister the timelock offRamp of old pool %s: %w", oldPool, err)
				}
				writes = append(writes, removeRampReport.Output)
			}

			batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
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
