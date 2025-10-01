package tokens

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// RemoteChainConfig specifies configuration for a remote chain on a token pool.
type RemoteChainConfig struct {
	// The address of the token on the remote chain in bytes.
	RemoteTokenAddress []byte
	// The address of the token pool on the remote chain in bytes.
	RemotePoolAddress []byte
	// InboundRateLimiterConfig specifies the desired rate limiter configuration for inbound traffic.
	InboundRateLimiterConfig token_pool.RateLimiterConfig
	// OutboundRateLimiterConfig specifies the desired rate limiter configuration for outbound traffic.
	OutboundRateLimiterConfig token_pool.RateLimiterConfig
	// OutboundCCVs specifies the contract addresses of verifiers to apply to outbound traffic.
	OutboundCCVs []common.Address
	// InboundCCVs specifies the contract addresses of verifiers to apply to inbound traffic.
	InboundCCVs []common.Address
}

// ConfigureTokenPoolForRemoteChainInput is the input for the ConfigureTokenPoolForRemoteChain sequence.
type ConfigureTokenPoolForRemoteChainInput struct {
	// ChainSelector is the chain selector for the chain being configured.
	ChainSelector uint64
	// TokenPoolAddress is the address of the token pool.
	TokenPoolAddress common.Address
	// RemoteChainSelector is the selector of the remote chain to configure.
	RemoteChainSelector uint64
	// RemoteChainConfig is the configuration for the remote chain.
	RemoteChainConfig RemoteChainConfig
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
	func(b operations.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}
		writes := make([]contract.WriteOutput, 0)

		// Set the requested CCVs
		setCCVsReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyCCVConfigUpdates, chain, contract.FunctionInput[[]token_pool.CCVConfigArg]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenPoolAddress,
			Args: []token_pool.CCVConfigArg{
				{
					RemoteChainSelector: input.RemoteChainSelector,
					OutboundCCVs:        input.RemoteChainConfig.InboundCCVs,
					InboundCCVs:         input.RemoteChainConfig.OutboundCCVs,
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to set CCVs: %w", err)
		}
		writes = append(writes, setCCVsReport.Output)

		// Get remote chains that are currently supported by the token pools
		supportedChainsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetSupportedChains, chain, contract.FunctionInput[any]{
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
		removes := make([]uint64, 1) // Cap == 1 because we may need to remove the chain if the remote token is different
		if slices.Contains(supportedChainsReport.Output, input.RemoteChainSelector) {
			// Check existing remote token
			getRemoteTokenReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRemoteToken, chain, contract.FunctionInput[uint64]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenPoolAddress,
				Args:          input.RemoteChainSelector,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote token: %w", err)
			}
			if !bytes.Equal(getRemoteTokenReport.Output, input.RemoteChainConfig.RemoteTokenAddress) {
				removes = append(removes, input.RemoteChainSelector)
			}

			// Only proceed further if we do NOT need to remove and re-add the chain
			if len(removes) == 0 {
				// Check existing rate limiters
				inboundRateLimiterReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetCurrentInboundRateLimiterState, chain, contract.FunctionInput[uint64]{
					ChainSelector: input.ChainSelector,
					Address:       input.TokenPoolAddress,
					Args:          input.RemoteChainSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get inbound rate limiter state: %w", err)
				}
				outboundRateLimiterReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetCurrentOutboundRateLimiterState, chain, contract.FunctionInput[uint64]{
					ChainSelector: input.ChainSelector,
					Address:       input.TokenPoolAddress,
					Args:          input.RemoteChainSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get outbound rate limiter state: %w", err)
				}
				// Update the rate limiters if they do not match the desired config
				if !rateLimiterConfigsEqual(inboundRateLimiterReport.Output, input.RemoteChainConfig.InboundRateLimiterConfig) || !rateLimiterConfigsEqual(outboundRateLimiterReport.Output, input.RemoteChainConfig.OutboundRateLimiterConfig) {
					setInboundRateLimiterReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetChainRateLimiterConfigs, chain, contract.FunctionInput[[]token_pool.SetChainRateLimiterConfigArg]{
						ChainSelector: input.ChainSelector,
						Address:       input.TokenPoolAddress,
						Args: []token_pool.SetChainRateLimiterConfigArg{
							{
								RemoteChainSelector:       input.RemoteChainSelector,
								InboundRateLimiterConfig:  input.RemoteChainConfig.InboundRateLimiterConfig,
								OutboundRateLimiterConfig: input.RemoteChainConfig.OutboundRateLimiterConfig,
							},
						},
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to set inbound rate limiter config: %w", err)
					}
					writes = append(writes, setInboundRateLimiterReport.Output)
				}

				// Check existing remote pools
				getRemotePoolsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRemotePools, chain, contract.FunctionInput[uint64]{
					ChainSelector: input.ChainSelector,
					Address:       input.TokenPoolAddress,
					Args:          input.RemoteChainSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote pools: %w", err)
				}
				if !slices.ContainsFunc(getRemotePoolsReport.Output, func(addr []byte) bool {
					return bytes.Equal(addr, input.RemoteChainConfig.RemotePoolAddress)
				}) {
					// Add the requested remote pool
					addRemotePoolsReport, err := cldf_ops.ExecuteOperation(b, token_pool.AddRemotePool, chain, contract.FunctionInput[token_pool.RemotePoolArgs]{
						ChainSelector: input.ChainSelector,
						Address:       input.TokenPoolAddress,
						Args: token_pool.RemotePoolArgs{
							RemoteChainSelector: input.RemoteChainSelector,
							RemotePoolAddress:   input.RemoteChainConfig.RemotePoolAddress,
						},
					})
					if err != nil {
						return sequences.OnChainOutput{}, fmt.Errorf("failed to add remote pool: %w", err)
					}
					writes = append(writes, addRemotePoolsReport.Output)
				}

				// Return early as no further action is required
				return sequences.OnChainOutput{Writes: writes}, nil
			}
		}

		// If the chain is not supported, apply the config for the remote chain
		applyChainUpdatesReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyChainUpdates, chain, contract.FunctionInput[token_pool.ApplyChainUpdatesArgs]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenPoolAddress,
			Args: token_pool.ApplyChainUpdatesArgs{
				RemoteChainSelectorsToRemove: []uint64{},
				ChainsToAdd: []token_pool.ChainUpdate{
					{
						RemoteChainSelector: input.RemoteChainSelector,
						RemotePoolAddresses: [][]byte{
							input.RemoteChainConfig.RemotePoolAddress,
						},
						RemoteTokenAddress:        input.RemoteChainConfig.RemoteTokenAddress,
						OutboundRateLimiterConfig: input.RemoteChainConfig.InboundRateLimiterConfig,
						InboundRateLimiterConfig:  input.RemoteChainConfig.OutboundRateLimiterConfig,
					},
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply chain updates: %w", err)
		}
		writes = append(writes, applyChainUpdatesReport.Output)

		return sequences.OnChainOutput{Writes: writes}, nil
	},
)

// rateLimiterConfigsEqual returns true if the current rate limiter config on-chain matches the desired config.
func rateLimiterConfigsEqual(current token_pool.RateLimiterBucket, desired token_pool.RateLimiterConfig) bool {
	return current.IsEnabled == desired.IsEnabled &&
		current.Capacity == desired.Capacity &&
		current.Rate == desired.Rate
}
