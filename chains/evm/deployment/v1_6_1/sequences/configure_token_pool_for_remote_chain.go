package sequences

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
)

// ConfigureTokenPoolForRemoteChainInput is the input for the ConfigureTokenPoolForRemoteChain sequence.
type ConfigureTokenPoolForRemoteChainInput struct {
	// ChainSelector is the chain selector for the chain being configured.
	ChainSelector uint64
	// TokenPoolAddress is the address of the token pool.
	TokenPoolAddress common.Address
	// RemoteChainSelector is the selector of the remote chain to configure.
	RemoteChainSelector uint64
	// RemoteChainConfig is the configuration for the remote chain.
	RemoteChainConfig tokens.RemoteChainConfig[[]byte, string]
}

func (c ConfigureTokenPoolForRemoteChainInput) Validate(chain evm.Chain) error {
	if c.ChainSelector != chain.Selector {
		return fmt.Errorf("chain selector %d does not match chain %s", c.ChainSelector, chain)
	}
	return nil
}

var ConfigureTokenPoolForRemoteChain = cldf_ops.NewSequence(
	"configure-token-pool-for-remote-chain",
	semver.MustParse("1.7.0"),
	"Configures a token pool on an EVM chain for transfers with other chains",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}
		writes := make([]evm_contract.WriteOutput, 0)

		// Get remote chains that are currently supported by the token pools
		supportedChainsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetSupportedChains, chain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenPoolAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get supported chains: %w", err)
		}

		// If the chain is supported
		// 1. Check remote token, remove and re-add remote config if requested remote token is different
		// 2. Check existing rate limiters and update if necessary
		// 3. Check existing remote pools and add requested remote pool if it does not exist
		removes := make([]uint64, 0, 1) // Cap == 1 because we may need to remove the chain if the remote token is different
		if slices.Contains(supportedChainsReport.Output, input.RemoteChainSelector) {
			// Check existing remote token
			getRemoteTokenReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRemoteToken, chain, evm_contract.FunctionInput[uint64]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenPoolAddress,
				Args:          input.RemoteChainSelector,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote token: %w", err)
			}
			if !bytes.Equal(getRemoteTokenReport.Output, input.RemoteChainConfig.RemoteToken) {
				removes = append(removes, input.RemoteChainSelector)
			}

			// Only proceed further if we do NOT need to remove and re-add the chain
			if len(removes) == 0 {
				// Check and update rate limiters
				rateLimitersReport, err := maybeUpdateRateLimiters(b, chain, input)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to maybe update rate limiters: %w", err)
				}
				// Only append if there's an actual write (maybeUpdateRateLimiters returns empty WriteOutput when no update needed)
				if rateLimitersReport.ChainSelector != 0 {
					writes = append(writes, rateLimitersReport)
				}

				// Check existing remote pools
				getRemotePoolsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRemotePools, chain, evm_contract.FunctionInput[uint64]{
					ChainSelector: input.ChainSelector,
					Address:       input.TokenPoolAddress,
					Args:          input.RemoteChainSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote pools: %w", err)
				}
				if !slices.ContainsFunc(getRemotePoolsReport.Output, func(addr []byte) bool {
					return bytes.Equal(addr, input.RemoteChainConfig.RemotePool)
				}) {
					// Add the requested remote pool
					addRemotePoolsReport, err := cldf_ops.ExecuteOperation(b, token_pool.AddRemotePool, chain, evm_contract.FunctionInput[token_pool.AddRemotePoolArgs]{
						ChainSelector: input.ChainSelector,
						Address:       input.TokenPoolAddress,
						Args: token_pool.AddRemotePoolArgs{
							RemoteChainSelector: input.RemoteChainSelector,
							RemotePoolAddress:   common.LeftPadBytes(input.RemoteChainConfig.RemotePool, 32),
						},
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to add remote pool: %w", err)
					}
					writes = append(writes, addRemotePoolsReport.Output)
				}

				// Return early as no further action is required
				batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
				}

				return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
			}
		}

		// If the chain is not supported, apply the config for the remote chain
		applyChainUpdatesReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyChainUpdates, chain, evm_contract.FunctionInput[token_pool.ApplyChainUpdatesArgs]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenPoolAddress,
			Args: token_pool.ApplyChainUpdatesArgs{
				RemoteChainSelectorsToRemove: removes,
				ChainsToAdd: []token_pool.ChainUpdate{
					{
						RemoteChainSelector: input.RemoteChainSelector,
						RemotePoolAddresses: [][]byte{
							common.LeftPadBytes(input.RemoteChainConfig.RemotePool, 32),
						},
						RemoteTokenAddress: common.LeftPadBytes(input.RemoteChainConfig.RemoteToken, 32),
						OutboundRateLimiterConfig: token_pool.Config{
							IsEnabled: input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.IsEnabled,
							Capacity:  input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.Capacity,
							Rate:      input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.Rate,
						},
						InboundRateLimiterConfig: token_pool.Config{
							IsEnabled: input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.IsEnabled,
							Capacity:  input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.Capacity,
							Rate:      input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.Rate,
						},
					},
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply chain updates: %w", err)
		}
		writes = append(writes, applyChainUpdatesReport.Output)

		batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	},
)

// maybeUpdateRateLimiters checks and updates the rate limiters for a given remote chain if they do not match the desired config.
func maybeUpdateRateLimiters(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput) (output evm_contract.WriteOutput, err error) {
	inboundRateLimiterStateReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetCurrentInboundRateLimiterState, chain, evm_contract.FunctionInput[uint64]{
		ChainSelector: input.ChainSelector,
		Address:       input.TokenPoolAddress,
		Args:          input.RemoteChainSelector,
	})
	if err != nil {
		return evm_contract.WriteOutput{}, fmt.Errorf("failed to get inbound rate limiter state: %w", err)
	}
	currentInboundRateLimiterState := inboundRateLimiterStateReport.Output

	outboundRateLimiterStateReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetCurrentOutboundRateLimiterState, chain, evm_contract.FunctionInput[uint64]{
		ChainSelector: input.ChainSelector,
		Address:       input.TokenPoolAddress,
		Args:          input.RemoteChainSelector,
	})
	if err != nil {
		return evm_contract.WriteOutput{}, fmt.Errorf("failed to get outbound rate limiter state: %w", err)
	}
	currentOutboundRateLimiterState := outboundRateLimiterStateReport.Output

	// Update the rate limiters if they do not match the desired config
	if !rateLimiterConfigsEqual(currentInboundRateLimiterState, input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig) ||
		!rateLimiterConfigsEqual(currentOutboundRateLimiterState, input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig) {
		setInboundRateLimiterReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetChainRateLimiterConfig, chain, evm_contract.FunctionInput[token_pool.SetChainRateLimiterConfigArgs]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenPoolAddress,
			Args: token_pool.SetChainRateLimiterConfigArgs{
				RemoteChainSelector: input.RemoteChainSelector,
				InboundConfig: token_pool.Config{
					IsEnabled: input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.IsEnabled,
					Capacity:  input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.Capacity,
					Rate:      input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.Rate,
				},
				OutboundConfig: token_pool.Config{
					IsEnabled: input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.IsEnabled,
					Capacity:  input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.Capacity,
					Rate:      input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.Rate,
				},
			},
		})
		if err != nil {
			return evm_contract.WriteOutput{}, fmt.Errorf("failed to set rate limiters config: %w", err)
		}
		return setInboundRateLimiterReport.Output, nil
	}

	return evm_contract.WriteOutput{}, nil
}

// rateLimiterConfigsEqual returns true if the current rate limiter config on-chain matches the desired config.
func rateLimiterConfigsEqual(current token_pool.TokenBucket, desired tokens.RateLimiterConfig) bool {
	return current.IsEnabled == desired.IsEnabled &&
		current.Capacity.Cmp(desired.Capacity) == 0 &&
		current.Rate.Cmp(desired.Rate) == 0
}
