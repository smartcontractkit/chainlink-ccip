package token_pool

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	tpops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_1/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/token_pool"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

type ConfigureTokenPoolForRemoteChainsInput struct {
	TokenPoolAddress common.Address
	RemoteChains     map[uint64]tokensapi.RemoteChainConfig[[]byte, string]
}

type ConfigureTokenPoolForRemoteChainInput struct {
	TokenPoolAddress    common.Address
	RemoteChainSelector uint64
	RemoteChainConfig   tokensapi.RemoteChainConfig[[]byte, string]
}

// ConfigureTokenPoolForRemoteChains configures a token pool on an EVM chain for cross-
// chain token transfers with other remote chains. It's capable of configuring multiple
// remote chains with a single invocation.
var ConfigureTokenPoolForRemoteChains = cldf_ops.NewSequence(
	"token-pool:configure-token-pool-for-remote-chains",
	tpops.Version,
	"Configure a token on an EVM chain for cross-chain transfers",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainsInput) (sequences.OnChainOutput, error) {
		// NOTE: this sequence will be called repeatedly as part of a larger changeset (e.g.
		// ConfigureTokensForTransfers) so we intentionally use the direct contract bindings
		// over ExecuteOperation to avoid the possibility of reading stale onchain data from
		// the operation reports cache.
		tokenPool, err := token_pool.NewTokenPool(input.TokenPoolAddress, chain.Client)
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
	tpops.Version,
	"Configures a token pool on an EVM chain for transfers with other chains",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput) (sequences.OnChainOutput, error) {
		// Below, we read onchain state directly from the contract binding. We intentionally
		// avoid the use of ExecuteOperation because it could return stale onchain data from
		// the operations reports cache if this sequence is called as part of a broader, and
		// more complex changeset that repeatedly reads and writes to the same config during
		// execution (e.g. ConfigureTokensForTransfers)
		tp, err := token_pool.NewTokenPool(input.TokenPoolAddress, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token pool contract: %w", err)
		}
		sc, err := tp.GetSupportedChains(&bind.CallOpts{Context: b.GetContext()})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get supported chains: %w", err)
		}

		// Token pool remote chain configuration can vary depending on whether the remote
		// pool is or isn't supported. The different cases to consider are recorded below
		// in the code.
		reportWrites := []contract.WriteOutput{}
		remotesToDel := []uint64{}
		if slices.Contains(sc, input.RemoteChainSelector) {
			remoteToken, err := tp.GetRemoteToken(&bind.CallOpts{Context: b.GetContext()}, input.RemoteChainSelector)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote token: %w", err)
			}

			// Token pool remote chain configuration can also vary depending on whether the
			// remote token matches or not - see comment further below for more details.
			if !bytes.Equal(remoteToken, input.RemoteChainConfig.RemoteToken) {
				// If the remote token onchain is different from the one provided as input, then we
				// need to ensure that ApplyChainUpdates removes any existing config for the remote
				// chain before a new one is used.
				remotesToDel = []uint64{input.RemoteChainSelector}
			} else {
				// If the remote token onchain matches the one provided as input, then we won't call
				// ApplyChainUpdates and instead handle the onchain updates via SetRateLimiterConfig
				// and AddRemotePool.
				inputORL := input.RemoteChainConfig.OutboundRateLimiterConfig
				inputIRL := input.RemoteChainConfig.InboundRateLimiterConfig
				remoteTP := input.RemoteChainConfig.RemotePool
				remoteCS := input.RemoteChainSelector

				// Query rate limits and remote pools
				onchainORL, err := tp.GetCurrentOutboundRateLimiterState(&bind.CallOpts{Context: b.GetContext()}, remoteCS)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get outbound rate limiter state: %w", err)
				}
				onchainIRL, err := tp.GetCurrentInboundRateLimiterState(&bind.CallOpts{Context: b.GetContext()}, remoteCS)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get inbound rate limiter state: %w", err)
				}
				remoteTPs, err := tp.GetRemotePools(&bind.CallOpts{Context: b.GetContext()}, remoteCS)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote token pools: %w", err)
				}

				// Check if the provided outbound RL matches the onchain outbound RL
				isOutboundEqual := inputORL.IsEnabled == onchainORL.IsEnabled &&
					inputORL.Capacity.Cmp(onchainORL.Capacity) == 0 &&
					inputORL.Rate.Cmp(onchainORL.Rate) == 0

				// Check if the provided inbound RL matches the onchain inbound RL
				isInboundEqual := inputIRL.IsEnabled == onchainIRL.IsEnabled &&
					inputIRL.Capacity.Cmp(onchainIRL.Capacity) == 0 &&
					inputIRL.Rate.Cmp(onchainIRL.Rate) == 0

				// Check if the provided remote token pool is already registered onchain
				hasRemoteTP := slices.ContainsFunc(remoteTPs, func(rtp []byte) bool {
					return bytes.Equal(rtp, remoteTP)
				})

				// If either rate limiter config is different, then update it
				if !isOutboundEqual || !isInboundEqual {
					report, err := cldf_ops.ExecuteOperation(b, tpops.SetChainRateLimiterConfig, chain, contract.FunctionInput[tpops.SetChainRateLimiterConfigArgs]{
						ChainSelector: chain.Selector,
						Address:       input.TokenPoolAddress,
						Args: tpops.SetChainRateLimiterConfigArgs{
							OutboundRateLimitConfig: token_pool.RateLimiterConfig{IsEnabled: inputORL.IsEnabled, Capacity: inputORL.Capacity, Rate: inputORL.Rate},
							InboundRateLimitConfig:  token_pool.RateLimiterConfig{IsEnabled: inputIRL.IsEnabled, Capacity: inputIRL.Capacity, Rate: inputIRL.Rate},
							RemoteChainSelector:     remoteCS,
						},
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to set rate limiter config: %w", err)
					}
					reportWrites = append(reportWrites, report.Output)
				}

				// If the remote token pool is not registered, then add it
				if !hasRemoteTP {
					report, err := cldf_ops.ExecuteOperation(b, tpops.AddRemotePool, chain, contract.FunctionInput[tpops.AddRemotePoolArgs]{
						ChainSelector: chain.Selector,
						Address:       input.TokenPoolAddress,
						Args: tpops.AddRemotePoolArgs{
							RemoteChainSelector: remoteCS,
							RemotePoolAddress:   remoteTP,
						},
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to add remote token pool: %w", err)
					}
					reportWrites = append(reportWrites, report.Output)
				}
			}
		}

		// Three cases to consider here:
		// --
		//   1. The chain is not supported yet in which case the only thing that's needed is to add
		//      it via ApplyChainUpdates. No removals are necessary, and rate limiters will be set.
		// --
		//   2. The chain is already supported AND the input remote token EQUALS the onchain remote
		//      token. In this case, we need to ensure that any existing remote configs are removed
		//      before adding a new one via ApplyChainUpdates.
		// --
		//   3. The chain is already supported AND the input remote token DIFFERS from the onchain
		//      remote token. In this case, we will never call ApplyChainUpdates. Instead, we will
		//      update on chain state purely using SetRateLimiterConfig and AddRemotePool above.
		//
		if len(reportWrites) == 0 {
			paddedRemoteTokenPoolAddress := common.LeftPadBytes(input.RemoteChainConfig.RemotePool, 32)
			applyChainUpdatesInput := contract.FunctionInput[tpops.ApplyChainUpdatesArgs]{
				ChainSelector: chain.Selector,
				Address:       input.TokenPoolAddress,
				Args: tpops.ApplyChainUpdatesArgs{
					RemoteChainSelectorsToRemove: remotesToDel,
					ChainsToAdd: []token_pool.TokenPoolChainUpdate{
						{
							RemotePoolAddresses: [][]byte{paddedRemoteTokenPoolAddress},
							RemoteChainSelector: input.RemoteChainSelector,
							RemoteTokenAddress:  input.RemoteChainConfig.RemoteToken,
							OutboundRateLimiterConfig: token_pool.RateLimiterConfig{
								IsEnabled: input.RemoteChainConfig.OutboundRateLimiterConfig.IsEnabled,
								Capacity:  input.RemoteChainConfig.OutboundRateLimiterConfig.Capacity,
								Rate:      input.RemoteChainConfig.OutboundRateLimiterConfig.Rate,
							},
							InboundRateLimiterConfig: token_pool.RateLimiterConfig{
								IsEnabled: input.RemoteChainConfig.InboundRateLimiterConfig.IsEnabled,
								Capacity:  input.RemoteChainConfig.InboundRateLimiterConfig.Capacity,
								Rate:      input.RemoteChainConfig.InboundRateLimiterConfig.Rate,
							},
						},
					},
				},
			}

			report, err := cldf_ops.ExecuteOperation(b, tpops.ApplyChainUpdates, chain, applyChainUpdatesInput)
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
