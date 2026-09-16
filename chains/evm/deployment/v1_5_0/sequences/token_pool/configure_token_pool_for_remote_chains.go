package token_pool

import (
	"bytes"
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	evmutils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	bmtpap "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_token_pool_and_proxy"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type ConfigureTokenPoolForRemoteChainsInput struct {
	TokenPoolAddress common.Address
	TokenPoolVersion *semver.Version
	RemoteChains     map[uint64]tokensapi.RemoteChainConfig[[]byte, string]
}

type ConfigureTokenPoolForRemoteChainInput struct {
	TokenPoolAddress    common.Address
	RemoteChainSelector uint64
	RemoteChainConfig   tokensapi.RemoteChainConfig[[]byte, string]
}

func (c ConfigureTokenPoolForRemoteChainInput) Validate() error {
	return evmutils.Wrap(
		c.RemoteChainConfig.Validate(),
		fmt.Sprintf("invalid remote chain config for remote chain selector %d", c.RemoteChainSelector),
	)
}

// disabledRateLimit is the zero-value rate limiter config. v1.5.0's applyChainUpdates validates
// each entry with `_validateTokenBucketConfig(cfg, mustBeDisabled: !allowed)`, so a removal entry
// (allowed=false) that carries a non-zero or enabled config reverts with RateLimitMustBeDisabled.
func disabledRateLimit() bmtpap.Config {
	return bmtpap.Config{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
}

// ConfigureTokenPoolForRemoteChains configures a v1.5.0 token pool on an EVM chain for cross-
// chain token transfers with other remote chains. It's capable of configuring multiple
// remote chains with a single invocation.
//
// This is the v1.5.0 analogue of the v1.5.1 sequence of the same name. It cannot share that
// implementation because the v1.5.0 pool ABI differs in four load-bearing ways:
//
//   - there is no getTokenDecimals(); local decimals are read from the token's ERC20 decimals()
//   - getRemotePool(uint64) returns a SINGLE remote pool, not the v1.5.1 getRemotePools() list
//   - there is no addRemotePool/removeRemotePool; setRemotePool REPLACES the single entry
//   - applyChainUpdates takes one ChainUpdate[] (with a singular remotePoolAddress) and expresses
//     removals as allowed=false entries, rather than v1.5.1's (selectorsToRemove, chainsToAdd)
var ConfigureTokenPoolForRemoteChains = cldf_ops.NewSequence(
	"token-pool:configure-token-pool-for-remote-chains",
	bmtpap.Version,
	"Configure a v1.5.0 token pool on an EVM chain for cross-chain transfers",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainsInput) (sequences.OnChainOutput, error) {
		// NOTE: this sequence will be called repeatedly as part of a larger changeset (e.g.
		// ConfigureTokensForTransfers) so we intentionally use the direct contract bindings
		// over ExecuteOperation to avoid the possibility of reading stale onchain data from
		// the operation reports cache.
		tokenPool, err := bmtpap.NewBurnMintTokenPoolAndProxyContract(input.TokenPoolAddress, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token pool contract: %w", err)
		}

		tokenAddress, err := tokenPool.GetToken(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token from token pool: %w", err)
		}

		isSupported, err := tokenPool.IsSupportedToken(&bind.CallOpts{Context: b.GetContext()}, tokenAddress)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check if token is supported: %w", err)
		}
		if !isSupported {
			return sequences.OnChainOutput{}, fmt.Errorf("token %s is not supported by token pool %s", tokenAddress.Hex(), input.TokenPoolAddress)
		}

		batchOps := make([]mcms_types.BatchOperation, 0)
		for remoteChainSelector, remoteChainConfig := range input.RemoteChains {
			report, err := cldf_ops.ExecuteSequence(b,
				ConfigureTokenPoolForRemoteChain,
				chain,
				ConfigureTokenPoolForRemoteChainInput{
					TokenPoolAddress:    tokenPool.Address(),
					RemoteChainSelector: remoteChainSelector,
					RemoteChainConfig:   remoteChainConfig,
				},
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for remote chain %d: %w", remoteChainSelector, err)
			}

			batchOps = append(batchOps, report.Output.BatchOps...)
		}

		return sequences.OnChainOutput{BatchOps: batchOps}, nil
	})

// ConfigureTokenPoolForRemoteChain is a helper sequence that performs the logic for
// configuring a token pool for a SINGLE remote chain. The sequence allows the upper
// level ConfigureTokenPoolForRemoteChains sequence to handle multiple remote chains
var ConfigureTokenPoolForRemoteChain = cldf_ops.NewSequence(
	"token-pool:configure-token-pool-for-remote-chain",
	bmtpap.Version,
	"Configures a v1.5.0 token pool on an EVM chain for transfers with other chains",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput) (sequences.OnChainOutput, error) {
		if err := input.Validate(); err != nil {
			return sequences.OnChainOutput{}, err
		}

		// Below, we read onchain state directly from the contract binding. We intentionally
		// avoid the use of ExecuteOperation because it could return stale onchain data from
		// the operations reports cache if this sequence is called as part of a broader, and
		// more complex changeset that repeatedly reads and writes to the same config during
		// execution (e.g. ConfigureTokensForTransfers)
		tp, err := bmtpap.NewBurnMintTokenPoolAndProxyContract(input.TokenPoolAddress, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token pool contract: %w", err)
		}
		sc, err := tp.GetSupportedChains(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get supported chains: %w", err)
		}
		chainSupported := slices.Contains(sc, input.RemoteChainSelector)

		// Read the currently configured remote token up front: whether it is changing decides both
		// how the rate limits are resolved below and whether ApplyChainUpdates has to remove the
		// existing lane config first.
		var remoteTokenChanged bool
		if chainSupported {
			onchainRemoteToken, err := tp.GetRemoteToken(&bind.CallOpts{Context: b.GetContext()}, input.RemoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote token: %w", err)
			}
			remoteTokenChanged = !bytes.Equal(onchainRemoteToken, input.RemoteChainConfig.RemoteToken)
		}

		// v1.5.0 pools have no getTokenDecimals(); read the token's own ERC20 decimals() instead.
		// Both the token address and its decimals are immutable, so ExecuteOperation (and its
		// report cache) is safe here, unlike the mutable reads above.
		localTokenAddr, err := tp.GetToken(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token from token pool: %w", err)
		}
		decimalsReport, err := cldf_ops.ExecuteOperation(b, erc20.GetDecimals, chain, contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       localTokenAddr,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get decimals for token %s: %w", localTokenAddr.Hex(), err)
		}
		localDecimals := decimalsReport.Output

		// A pool's type and version is immutable so we can safely use ExecuteOperation here
		// without worrying about stale data from the cache.
		tvReport, err := cldf_ops.ExecuteOperation(b, type_and_version.GetTypeAndVersion, chain, contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       input.TokenPoolAddress,
			Args:          struct{}{},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get type and version of token pool: %w", err)
		}

		// Get outbound and inbound rate limits from the input
		outboundRL, outboundOk := input.RemoteChainConfig.GetOutboundRateLimitBuckets().DefaultBucket()
		inboundRL, inboundOk := input.RemoteChainConfig.GetInboundRateLimitBuckets().DefaultBucket()

		// Resolve the outbound and inbound rate limits
		var inputORL, inputIRL tokensapi.RateLimiterConfig
		switch {
		case outboundOk && inboundOk:
			// If the user explicitly provided both the outbound and inbound rate limits, then
			// we use them.
			inputORL, inputIRL = tokensapi.GenerateTPRLConfigs(
				outboundRL.RateLimit,
				inboundRL.RateLimit,
				localDecimals,
				input.RemoteChainConfig.RemoteDecimals,
				chain.Family(),
				tvReport.Output.Version,
				tvReport.Output.Type.String(),
			)

		case !outboundOk && !inboundOk:
			if chainSupported {
				// Idempotent behavior: if we're re-calling this sequence and no rate limits are
				// specified, then we re-use whatever is currently onchain to avoid accidentally
				// overwriting existing onchain config
				onchainOutboundBucket, err := tp.GetCurrentOutboundRateLimiterState(&bind.CallOpts{Context: b.GetContext()}, input.RemoteChainSelector)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get outbound rate limiter state for remote chain %d: %w", input.RemoteChainSelector, err)
				}
				onchainInboundBucket, err := tp.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: b.GetContext()}, input.RemoteChainSelector)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get inbound rate limiter state for remote chain %d: %w", input.RemoteChainSelector, err)
				}

				// Carrying the existing buckets forward is only safe while the lane keeps pointing at
				// the same remote token. See tokensapi.ErrInboundRateLimitNotPortable for why, and
				// DoesPoolUseLocalDecimals for which pools are affected (v1.5.0 always is, but the
				// predicate is asked rather than assumed so this stays correct if that changes).
				if remoteTokenChanged && onchainInboundBucket.IsEnabled &&
					!tokensapi.DoesPoolUseLocalDecimals(chain.Family(), tvReport.Output.Version, tvReport.Output.Type.String()) {
					return sequences.OnChainOutput{}, tokensapi.ErrInboundRateLimitNotPortable(
						input.RemoteChainSelector, onchainInboundBucket.Capacity, onchainInboundBucket.Rate,
					)
				}

				inputORL = tokensapi.RateLimiterConfig{
					IsEnabled: onchainOutboundBucket.IsEnabled,
					Capacity:  onchainOutboundBucket.Capacity,
					Rate:      onchainOutboundBucket.Rate,
				}
				inputIRL = tokensapi.RateLimiterConfig{
					IsEnabled: onchainInboundBucket.IsEnabled,
					Capacity:  onchainInboundBucket.Capacity,
					Rate:      onchainInboundBucket.Rate,
				}
			} else {
				// If this is a fresh configuration for a remote chain (i.e. the remote chain selector
				// is not currently supported onchain), and no rate limits are specified in the input,
				// then we default to disabled rate limiters.
				inputORL = tokensapi.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
				inputIRL = tokensapi.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
			}

		default:
			return sequences.OnChainOutput{}, fmt.Errorf(
				"default outbound and inbound rate limits must both be specified together or both omitted for remote chain %d",
				input.RemoteChainSelector,
			)
		}

		// NOTE: v1.5.0 pools share v1.5.1's rate limiter validation: an enabled bucket with both
		// rate and capacity at zero is rejected (InvalidRateLimitRate), and a disabled bucket with
		// a non-zero rate or capacity is rejected (DisabledNonZeroRateLimit).
		if inputORL.IsEnabled && inputORL.Capacity.Cmp(big.NewInt(0)) == 0 && inputORL.Rate.Cmp(big.NewInt(0)) == 0 {
			return sequences.OnChainOutput{}, fmt.Errorf("outbound rate limiter config is enabled but rate and capacity are both zero")
		}
		if inputIRL.IsEnabled && inputIRL.Capacity.Cmp(big.NewInt(0)) == 0 && inputIRL.Rate.Cmp(big.NewInt(0)) == 0 {
			return sequences.OnChainOutput{}, fmt.Errorf("inbound rate limiter config is enabled but rate and capacity are both zero")
		}

		// Token pool remote chain configuration can vary depending on whether the remote
		// pool is or isn't supported. The different cases to consider are recorded below
		// in the code.
		reportWrites := []contract.WriteOutput{}
		removeRemoteFirst := false
		if chainSupported {
			// Token pool remote chain configuration can also vary depending on whether the
			// remote token matches or not - see comment further below for more details.
			if remoteTokenChanged {
				// If the remote token onchain is different from the one provided as input, then we
				// need to ensure that ApplyChainUpdates removes any existing config for the remote
				// chain before a new one is added. v1.5.0 expresses that as an allowed=false entry
				// preceding the allowed=true entry in the same ChainUpdate array.
				removeRemoteFirst = true
			} else {
				// If the remote token onchain matches the one provided as input, then we won't call
				// ApplyChainUpdates and instead handle the onchain updates via
				// SetChainRateLimiterConfig and SetRemotePool.
				// Remote pool addresses in CCIP messages are ABI-encoded (32-byte left-padded).
				// Using left-padded addresses here ensures the stored value matches what
				// the protocol sends, preventing "invalid source pool" errors on delivery.
				remoteTP := common.LeftPadBytes(input.RemoteChainConfig.RemotePool, 32)
				remoteCS := input.RemoteChainSelector

				// Query rate limits and the (single) remote pool
				onchainORL, err := tp.GetCurrentOutboundRateLimiterState(&bind.CallOpts{Context: b.GetContext()}, remoteCS)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get outbound rate limiter state: %w", err)
				}
				onchainIRL, err := tp.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: b.GetContext()}, remoteCS)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get inbound rate limiter state: %w", err)
				}
				onchainRemoteTP, err := tp.GetRemotePool(&bind.CallOpts{Context: b.GetContext()}, remoteCS)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote token pool: %w", err)
				}

				// Check if the provided outbound RL matches the onchain outbound RL
				isOutboundEqual := inputORL.IsEnabled == onchainORL.IsEnabled &&
					inputORL.Capacity.Cmp(onchainORL.Capacity) == 0 &&
					inputORL.Rate.Cmp(onchainORL.Rate) == 0

				// Check if the provided inbound RL matches the onchain inbound RL
				isInboundEqual := inputIRL.IsEnabled == onchainIRL.IsEnabled &&
					inputIRL.Capacity.Cmp(onchainIRL.Capacity) == 0 &&
					inputIRL.Rate.Cmp(onchainIRL.Rate) == 0

				// Exact (not normalized) comparison, matching v1.5.1: if only a 20-byte entry
				// exists from a prior run, this returns false and we rewrite the correct
				// 32-byte value.
				hasRemoteTP := bytes.Equal(onchainRemoteTP, remoteTP)

				// If either rate limiter config is different, then update it
				if !isOutboundEqual || !isInboundEqual {
					report, err := cldf_ops.ExecuteOperation(b, bmtpap.SetChainRateLimiterConfig, chain, contract.FunctionInput[bmtpap.SetChainRateLimiterConfigArgs]{
						ChainSelector: chain.Selector,
						Address:       input.TokenPoolAddress,
						Args: bmtpap.SetChainRateLimiterConfigArgs{
							OutboundConfig:      bmtpap.Config{IsEnabled: inputORL.IsEnabled, Capacity: inputORL.Capacity, Rate: inputORL.Rate},
							InboundConfig:       bmtpap.Config{IsEnabled: inputIRL.IsEnabled, Capacity: inputIRL.Capacity, Rate: inputIRL.Rate},
							RemoteChainSelector: remoteCS,
						},
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limiter config: %w", err)
					}
					reportWrites = append(reportWrites, report.Output)
				}

				// If the remote pool differs, REPLACE it.
				//
				// ⚠️ This is a genuine behavioural difference from v1.5.1 and newer pools, which
				// can hold several remote pools per lane and therefore append the new pool
				// (addRemotePool) while the old one keeps accepting in-flight messages. A v1.5.0
				// pool stores exactly one remote pool per lane, so setRemotePool is a hard
				// cutover: messages already in flight from the previous remote pool will be
				// rejected with InvalidSourcePoolAddress on arrival. Drain the lane before
				// retargeting a v1.5.0 pool.
				if !hasRemoteTP {
					b.Logger.Warnf(
						"v1.5.0 pool %s on chain %d stores a single remote pool per lane: replacing the remote pool for chain %d (%s -> %s). "+
							"This is a cutover, not an append - in-flight messages from the previous remote pool will be rejected.",
						input.TokenPoolAddress.Hex(), chain.Selector, remoteCS,
						common.Bytes2Hex(onchainRemoteTP), common.Bytes2Hex(remoteTP),
					)
					report, err := cldf_ops.ExecuteOperation(b, bmtpap.SetRemotePool, chain, contract.FunctionInput[bmtpap.SetRemotePoolArgs]{
						ChainSelector: chain.Selector,
						Address:       input.TokenPoolAddress,
						Args: bmtpap.SetRemotePoolArgs{
							RemoteChainSelector: remoteCS,
							RemotePoolAddress:   remoteTP,
						},
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to set remote token pool: %w", err)
					}
					reportWrites = append(reportWrites, report.Output)
				}

				// The chain is already supported with a matching remote token. If
				// reportWrites is still empty here, rate limiters and the pool address are
				// all already correct — nothing left to do.
				if len(reportWrites) == 0 {
					return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{}}, nil
				}
			}
		}

		// Three cases to consider here:
		// --
		//   1. The chain is not supported yet in which case the only thing that's needed is to add
		//      it via ApplyChainUpdates. No removals are necessary, and rate limiters will be set.
		// --
		//   2. The chain is already supported AND the input remote token DIFFERS from the onchain
		//      remote token. In this case we prepend an allowed=false entry for the same selector
		//      so the existing config is removed before the new one is added in the same call
		//      (v1.5.0 applyChainUpdates processes the array in order and reverts with
		//      ChainAlreadyExists if an existing selector is added without being removed first).
		// --
		//   3. The chain is already supported AND the input remote token EQUALS the onchain remote
		//      token. In this case, we will never call ApplyChainUpdates. Instead, we handle
		//      onchain updates via SetChainRateLimiterConfig and SetRemotePool above, returning
		//      early if the chain is already fully configured.
		//
		if len(reportWrites) == 0 {
			paddedRemoteTokenPoolAddress := common.LeftPadBytes(input.RemoteChainConfig.RemotePool, 32)
			chainUpdates := make([]bmtpap.ChainUpdate, 0, 2)
			if removeRemoteFirst {
				// Removal entries must carry disabled, zeroed rate limiter configs or the
				// contract reverts with RateLimitMustBeDisabled.
				chainUpdates = append(chainUpdates, bmtpap.ChainUpdate{
					RemoteChainSelector:       input.RemoteChainSelector,
					Allowed:                   false,
					RemotePoolAddress:         []byte{},
					RemoteTokenAddress:        []byte{},
					OutboundRateLimiterConfig: disabledRateLimit(),
					InboundRateLimiterConfig:  disabledRateLimit(),
				})
			}
			chainUpdates = append(chainUpdates, bmtpap.ChainUpdate{
				RemoteChainSelector: input.RemoteChainSelector,
				Allowed:             true,
				RemotePoolAddress:   paddedRemoteTokenPoolAddress,
				RemoteTokenAddress:  input.RemoteChainConfig.RemoteToken,
				OutboundRateLimiterConfig: bmtpap.Config{
					IsEnabled: inputORL.IsEnabled,
					Capacity:  inputORL.Capacity,
					Rate:      inputORL.Rate,
				},
				InboundRateLimiterConfig: bmtpap.Config{
					IsEnabled: inputIRL.IsEnabled,
					Capacity:  inputIRL.Capacity,
					Rate:      inputIRL.Rate,
				},
			})

			report, err := cldf_ops.ExecuteOperation(b, bmtpap.ApplyChainUpdates, chain, contract.FunctionInput[[]bmtpap.ChainUpdate]{
				ChainSelector: chain.Selector,
				Address:       input.TokenPoolAddress,
				Args:          chainUpdates,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply chain updates: %w", err)
			}

			reportWrites = append(reportWrites, report.Output)
		}

		batchOp, err := contract.NewBatchOperationFromWrites(reportWrites)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation: %w", err)
		}

		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	})
