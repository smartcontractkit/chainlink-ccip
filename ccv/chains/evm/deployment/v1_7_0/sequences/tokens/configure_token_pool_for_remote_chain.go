package tokens

import (
	"bytes"
	"fmt"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
	tp_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"

	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// ConfigureTokenPoolForRemoteChainInput is the input for the ConfigureTokenPoolForRemoteChain sequence.
type ConfigureTokenPoolForRemoteChainInput struct {
	// ChainSelector is the chain selector for the chain being configured.
	ChainSelector uint64
	// TokenPoolAddress is the address of the token pool.
	TokenPoolAddress common.Address
	// AdvancedPoolHooks is the address of the AdvancedPoolHooks contract.
	AdvancedPoolHooks common.Address
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

		// Update token transfer fee configuration for the remote chain
		if input.RemoteChainConfig.TokenTransferFeeConfig.IsEnabled {
			applyTokenTransferFeeConfigUpdatesReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyTokenTransferFeeConfigUpdates, chain, evm_contract.FunctionInput[token_pool.TokenTransferFeeConfigArgs]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenPoolAddress,
				Args: token_pool.TokenTransferFeeConfigArgs{
					TokenTransferFeeConfigUpdates: []token_pool.TokenTransferFeeConfigUpdate{
						{
							DestChainSelector:                      input.RemoteChainSelector,
							DestGasOverhead:                        input.RemoteChainConfig.TokenTransferFeeConfig.DestGasOverhead,
							DestBytesOverhead:                      input.RemoteChainConfig.TokenTransferFeeConfig.DestBytesOverhead,
							DefaultBlockConfirmationFeeUSDCents:    input.RemoteChainConfig.TokenTransferFeeConfig.DefaultFinalityFeeUSDCents,
							CustomBlockConfirmationFeeUSDCents:     input.RemoteChainConfig.TokenTransferFeeConfig.CustomFinalityFeeUSDCents,
							DefaultBlockConfirmationTransferFeeBps: input.RemoteChainConfig.TokenTransferFeeConfig.DefaultFinalityTransferFeeBps,
							CustomBlockConfirmationTransferFeeBps:  input.RemoteChainConfig.TokenTransferFeeConfig.CustomFinalityTransferFeeBps,
							IsEnabled:                              input.RemoteChainConfig.TokenTransferFeeConfig.IsEnabled,
						},
					},
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply token transfer fee config updates: %w", err)
			}
			writes = append(writes, applyTokenTransferFeeConfigUpdatesReport.Output)
		}

		// Set the requested CCVs
		inboundCCVs := make([]common.Address, len(input.RemoteChainConfig.InboundCCVs))
		for i, addr := range input.RemoteChainConfig.InboundCCVs {
			inboundCCVs[i] = common.HexToAddress(addr)
		}
		outboundCCVs := make([]common.Address, len(input.RemoteChainConfig.OutboundCCVs))
		for i, addr := range input.RemoteChainConfig.OutboundCCVs {
			outboundCCVs[i] = common.HexToAddress(addr)
		}
		outboundCCVsToAddAboveThreshold := make([]common.Address, len(input.RemoteChainConfig.OutboundCCVsToAddAboveThreshold))
		for i, addr := range input.RemoteChainConfig.OutboundCCVsToAddAboveThreshold {
			outboundCCVsToAddAboveThreshold[i] = common.HexToAddress(addr)
		}
		inboundCCVsToAddAboveThreshold := make([]common.Address, len(input.RemoteChainConfig.InboundCCVsToAddAboveThreshold))
		for i, addr := range input.RemoteChainConfig.InboundCCVsToAddAboveThreshold {
			inboundCCVsToAddAboveThreshold[i] = common.HexToAddress(addr)
		}
		if len(inboundCCVs) > 0 || len(outboundCCVs) > 0 || len(outboundCCVsToAddAboveThreshold) > 0 || len(inboundCCVsToAddAboveThreshold) > 0 {
			setCCVsReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.ApplyCCVConfigUpdates, chain, evm_contract.FunctionInput[[]advanced_pool_hooks.CCVConfigArg]{
				ChainSelector: input.ChainSelector,
				Address:       input.AdvancedPoolHooks,
				Args: []advanced_pool_hooks.CCVConfigArg{
					{
						RemoteChainSelector:             input.RemoteChainSelector,
						OutboundCCVs:                    outboundCCVs,
						OutboundCCVsToAddAboveThreshold: outboundCCVsToAddAboveThreshold,
						InboundCCVs:                     inboundCCVs,
						InboundCCVsToAddAboveThreshold:  inboundCCVsToAddAboveThreshold,
					},
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set CCVs: %w", err)
			}
			writes = append(writes, setCCVsReport.Output)
		}

		// Get remote chains that are currently supported by the token pools
		supportedChainsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetSupportedChains, chain, evm_contract.FunctionInput[any]{
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
				// Check and update default finality rate limiters
				defaultFinalityRateLimitersReport, err := maybeUpdateRateLimiters(b, chain, input, false)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to maybe update rate limiters: %w", err)
				}
				writes = append(writes, defaultFinalityRateLimitersReport)

				// Check and update custom finality rate limiters
				customFinalityRateLimitersReport, err := maybeUpdateRateLimiters(b, chain, input, true)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to maybe update rate limiters: %w", err)
				}
				writes = append(writes, customFinalityRateLimitersReport)

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
					addRemotePoolsReport, err := cldf_ops.ExecuteOperation(b, token_pool.AddRemotePool, chain, evm_contract.FunctionInput[token_pool.RemotePoolArgs]{
						ChainSelector: input.ChainSelector,
						Address:       input.TokenPoolAddress,
						Args: token_pool.RemotePoolArgs{
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
						RemoteTokenAddress:        common.LeftPadBytes(input.RemoteChainConfig.RemoteToken, 32),
						OutboundRateLimiterConfig: input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig,
						InboundRateLimiterConfig:  input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig,
					},
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply chain updates: %w", err)
		}
		writes = append(writes, applyChainUpdatesReport.Output)
		// Check and update custom finality rate limiters
		customFinalityRateLimitersReport, err := maybeUpdateRateLimiters(b, chain, input, true)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to maybe update rate limiters: %w", err)
		}
		writes = append(writes, customFinalityRateLimitersReport)

		batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	},
)

// maybeUpdateRateLimiters checks and updates the rate limiters for a given remote chain if they do not match the desired config.
func maybeUpdateRateLimiters(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput, customBlockConfirmation bool) (output evm_contract.WriteOutput, err error) {
	var desiredInboundRateLimiterConfig tokens.RateLimiterConfig
	var desiredOutboundRateLimiterConfig tokens.RateLimiterConfig
	if customBlockConfirmation {
		desiredInboundRateLimiterConfig = input.RemoteChainConfig.CustomFinalityInboundRateLimiterConfig
		desiredOutboundRateLimiterConfig = input.RemoteChainConfig.CustomFinalityOutboundRateLimiterConfig
	} else {
		desiredInboundRateLimiterConfig = input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig
		desiredOutboundRateLimiterConfig = input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig
	}

	// Check existing rate limiters
	rateLimiterStateReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetCurrentRateLimiterState, chain, evm_contract.FunctionInput[token_pool.GetCurrentRateLimiterStateArgs]{
		ChainSelector: input.ChainSelector,
		Address:       input.TokenPoolAddress,
		Args: token_pool.GetCurrentRateLimiterStateArgs{
			RemoteChainSelector:     input.RemoteChainSelector,
			CustomBlockConfirmation: customBlockConfirmation,
		},
	})
	if err != nil {
		return evm_contract.WriteOutput{}, fmt.Errorf("failed to get rate limiter state: %w", err)
	}
	currentStates := rateLimiterStateReport.Output

	// Update the rate limiters if they do not match the desired config
	if !rateLimiterConfigsEqual(currentStates.InboundRateLimiterState, desiredInboundRateLimiterConfig) ||
		!rateLimiterConfigsEqual(currentStates.OutboundRateLimiterState, desiredOutboundRateLimiterConfig) {
		setInboundRateLimiterReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetRateLimitConfig, chain, evm_contract.FunctionInput[[]token_pool.SetRateLimitConfigArg]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenPoolAddress,
			Args: []token_pool.SetRateLimitConfigArg{
				{
					RemoteChainSelector:       input.RemoteChainSelector,
					CustomBlockConfirmation:   customBlockConfirmation,
					InboundRateLimiterConfig:  desiredInboundRateLimiterConfig,
					OutboundRateLimiterConfig: desiredOutboundRateLimiterConfig,
				},
			},
		})
		if err != nil {
			return evm_contract.WriteOutput{}, fmt.Errorf("failed to set inbound rate limiter config: %w", err)
		}
		return setInboundRateLimiterReport.Output, nil
	}

	return evm_contract.WriteOutput{}, nil
}

// rateLimiterConfigsEqual returns true if the current rate limiter config on-chain matches the desired config.
func rateLimiterConfigsEqual(current tp_bindings.RateLimiterTokenBucket, desired tokens.RateLimiterConfig) bool {
	return current.IsEnabled == desired.IsEnabled &&
		current.Capacity == desired.Capacity &&
		current.Rate == desired.Rate
}
