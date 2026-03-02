package tokens

import (
	"bytes"
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	burn_mint_token_pool_latest "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/burn_mint_token_pool"
	tp_bindings "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/token_pool"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
	v17seq "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
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

		localDecimalsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetTokenDecimals, chain, evm_contract.FunctionInput[any]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenPoolAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token decimals: %w", err)
		}

		defaultFinalityOutboundRateLimiterConfig, defaultFinalityInboundRateLimiterConfig := tokens.GenerateTPRLConfigs(
			input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig,
			input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig,
			localDecimalsReport.Output,
			input.RemoteChainConfig.RemoteDecimals,
			chain.Family(),
			token_pool.Version,
		)

		customFinalityOutboundRateLimiterConfig, customFinalityInboundRateLimiterConfig := tokens.GenerateTPRLConfigs(
			input.RemoteChainConfig.CustomFinalityOutboundRateLimiterConfig,
			input.RemoteChainConfig.CustomFinalityInboundRateLimiterConfig,
			localDecimalsReport.Output,
			input.RemoteChainConfig.RemoteDecimals,
			chain.Family(),
			token_pool.Version,
		)

		// Update token transfer fee configuration for the remote chain.
		tokenTransferFeeConfigUpdates, err := makeTokenTransferFeeConfigUpdates(b, chain, input, input.RemoteChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to make token transfer fee config updates: %w", err)
		}
		if len(tokenTransferFeeConfigUpdates) > 0 {
			applyTokenTransferFeeConfigUpdatesReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyTokenTransferFeeConfigUpdates, chain, evm_contract.FunctionInput[token_pool.TokenTransferFeeConfigArgs]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenPoolAddress,
				Args: token_pool.TokenTransferFeeConfigArgs{
					TokenTransferFeeConfigUpdates: tokenTransferFeeConfigUpdates,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply token transfer fee config updates: %w", err)
			}
			writes = append(writes, applyTokenTransferFeeConfigUpdatesReport.Output)
		}

		// Set CCVs for the remote chain (idempotent: only apply when on-chain differs from desired)
		if input.AdvancedPoolHooks != (common.Address{}) {
			ccvArg, needCCVUpdate, err := makeCCVUpdates(b, chain, input.ChainSelector, input.AdvancedPoolHooks, input.RemoteChainSelector,
				input.RemoteChainConfig.InboundCCVs,
				input.RemoteChainConfig.OutboundCCVs,
				input.RemoteChainConfig.InboundCCVsToAddAboveThreshold,
				input.RemoteChainConfig.OutboundCCVsToAddAboveThreshold,
			)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("make CCV updates: %w", err)
			}
			if needCCVUpdate && ccvArg != nil {
				setCCVsReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.ApplyCCVConfigUpdates, chain, evm_contract.FunctionInput[[]advanced_pool_hooks.CCVConfigArg]{
					ChainSelector: input.ChainSelector,
					Address:       input.AdvancedPoolHooks,
					Args:          []advanced_pool_hooks.CCVConfigArg{*ccvArg},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to set CCVs: %w", err)
				}
				writes = append(writes, setCCVsReport.Output)
			}
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
				defaultFinalityRateLimitersReport, err := maybeUpdateRateLimiters(
					b,
					chain,
					input.ChainSelector,
					input.TokenPoolAddress,
					input.RemoteChainSelector,
					false,
					customFinalityInboundRateLimiterConfig,
					customFinalityOutboundRateLimiterConfig,
					defaultFinalityInboundRateLimiterConfig,
					defaultFinalityOutboundRateLimiterConfig,
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to maybe update rate limiters: %w", err)
				}
				if defaultFinalityRateLimitersReport != nil {
					writes = append(writes, *defaultFinalityRateLimitersReport)
				}

				// Check and update custom finality rate limiters
				customFinalityRateLimitersReport, err := maybeUpdateRateLimiters(
					b,
					chain,
					input.ChainSelector,
					input.TokenPoolAddress,
					input.RemoteChainSelector,
					true,
					customFinalityInboundRateLimiterConfig,
					customFinalityOutboundRateLimiterConfig,
					defaultFinalityInboundRateLimiterConfig,
					defaultFinalityOutboundRateLimiterConfig,
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to maybe update rate limiters: %w", err)
				}
				if customFinalityRateLimitersReport != nil {
					writes = append(writes, *customFinalityRateLimitersReport)
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
						OutboundRateLimiterConfig: defaultFinalityOutboundRateLimiterConfig,
						InboundRateLimiterConfig:  defaultFinalityInboundRateLimiterConfig,
					},
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply chain updates: %w", err)
		}
		writes = append(writes, applyChainUpdatesReport.Output)

		// Check and update custom finality rate limiters
		customFinalityRateLimitersReport, err := maybeUpdateRateLimiters(
			b,
			chain,
			input.ChainSelector,
			input.TokenPoolAddress,
			input.RemoteChainSelector,
			true,
			customFinalityInboundRateLimiterConfig,
			customFinalityOutboundRateLimiterConfig,
			defaultFinalityInboundRateLimiterConfig,
			defaultFinalityOutboundRateLimiterConfig,
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to maybe update rate limiters: %w", err)
		}
		if customFinalityRateLimitersReport != nil {
			writes = append(writes, *customFinalityRateLimitersReport)
		}

		batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	},
)

// maybeUpdateRateLimiters checks and updates the rate limiters for a given remote chain if they do not match the
// desired config. Returns nil if no update is needed.
func maybeUpdateRateLimiters(
	b cldf_ops.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	tokenPoolAddress common.Address,
	remoteChainSelector uint64,
	customBlockConfirmation bool,
	customFinalityInboundConfig tokens.RateLimiterConfig,
	customFinalityOutboundConfig tokens.RateLimiterConfig,
	defaultFinalityInboundConfig tokens.RateLimiterConfig,
	defaultFinalityOutboundConfig tokens.RateLimiterConfig,
) (*evm_contract.WriteOutput, error) {
	desiredInboundRateLimiterConfig := defaultFinalityInboundConfig
	desiredOutboundRateLimiterConfig := defaultFinalityOutboundConfig
	if customBlockConfirmation {
		desiredInboundRateLimiterConfig = customFinalityInboundConfig
		desiredOutboundRateLimiterConfig = customFinalityOutboundConfig
	}

	// Check existing rate limiters
	rateLimiterStateReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetCurrentRateLimiterState, chain, evm_contract.FunctionInput[token_pool.GetCurrentRateLimiterStateArgs]{
		ChainSelector: chainSelector,
		Address:       tokenPoolAddress,
		Args: token_pool.GetCurrentRateLimiterStateArgs{
			RemoteChainSelector:     remoteChainSelector,
			CustomBlockConfirmation: customBlockConfirmation,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get rate limiter state: %w", err)
	}
	currentStates := rateLimiterStateReport.Output

	// Update the rate limiters if they do not match the desired config.
	// We only allow updates to the rate limiters, we do not allow disabling them.
	// This is to reduce the risk of accidentally disabling the rate limiters.
	// Disabling traffic on a token is allowed, as in this case IsEnabled would be true with rate = capacity = 0.
	if (!rateLimiterConfigsEqual(currentStates.InboundRateLimiterState, desiredInboundRateLimiterConfig) && desiredInboundRateLimiterConfig.IsEnabled) ||
		(!rateLimiterConfigsEqual(currentStates.OutboundRateLimiterState, desiredOutboundRateLimiterConfig) && desiredOutboundRateLimiterConfig.IsEnabled) {
		setInboundRateLimiterReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetRateLimitConfig, chain, evm_contract.FunctionInput[[]token_pool.SetRateLimitConfigArg]{
			ChainSelector: chainSelector,
			Address:       tokenPoolAddress,
			Args: []token_pool.SetRateLimitConfigArg{
				{
					RemoteChainSelector:       remoteChainSelector,
					CustomBlockConfirmation:   customBlockConfirmation,
					InboundRateLimiterConfig:  desiredInboundRateLimiterConfig,
					OutboundRateLimiterConfig: desiredOutboundRateLimiterConfig,
				},
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to set rate limiters config: %w", err)
		}
		return &setInboundRateLimiterReport.Output, nil
	}

	return nil, nil
}

// rateLimiterConfigsEqual returns true if the current rate limiter config on-chain matches the desired config.
func rateLimiterConfigsEqual(current tp_bindings.RateLimiterTokenBucket, desired tokens.RateLimiterConfig) bool {
	return current.IsEnabled == desired.IsEnabled &&
		current.Capacity.Cmp(desired.Capacity) == 0 &&
		current.Rate.Cmp(desired.Rate) == 0
}

func makeTokenTransferFeeConfigUpdates(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput, remoteChainSelector uint64) ([]token_pool.TokenTransferFeeConfigUpdate, error) {
	desiredTokenTransferFeeConfig := input.RemoteChainConfig.TokenTransferFeeConfig
	if !desiredTokenTransferFeeConfig.IsEnabled {
		return nil, nil
	}

	report, err := cldf_ops.ExecuteOperation(b, token_pool.GetTokenTransferFeeConfig, chain, evm_contract.FunctionInput[uint64]{
		ChainSelector: input.ChainSelector,
		Address:       input.TokenPoolAddress,
		Args:          remoteChainSelector,
	})
	if err != nil {
		// Print the token pool address and type and version
		boundTP, _ := burn_mint_token_pool_latest.NewBurnMintTokenPool(input.TokenPoolAddress, chain.Client)
		typeAndVersion, _ := boundTP.TypeAndVersion(nil)
		fmt.Println(typeAndVersion)
		fmt.Println(input.TokenPoolAddress)
		return nil, fmt.Errorf("failed to get token transfer fee config: %w", err)
	}

	currentTokenTransferFeeConfig := report.Output

	// Fall back to on-chain values if inputted values are empty
	if desiredTokenTransferFeeConfig.DestGasOverhead == 0 {
		desiredTokenTransferFeeConfig.DestGasOverhead = currentTokenTransferFeeConfig.DestGasOverhead
	}
	if desiredTokenTransferFeeConfig.DestBytesOverhead == 0 {
		desiredTokenTransferFeeConfig.DestBytesOverhead = currentTokenTransferFeeConfig.DestBytesOverhead
	}
	if desiredTokenTransferFeeConfig.DefaultFinalityFeeUSDCents == 0 {
		desiredTokenTransferFeeConfig.DefaultFinalityFeeUSDCents = currentTokenTransferFeeConfig.DefaultBlockConfirmationsFeeUSDCents
	}
	if desiredTokenTransferFeeConfig.CustomFinalityFeeUSDCents == 0 {
		desiredTokenTransferFeeConfig.CustomFinalityFeeUSDCents = currentTokenTransferFeeConfig.CustomBlockConfirmationsFeeUSDCents
	}
	if desiredTokenTransferFeeConfig.DefaultFinalityTransferFeeBps == 0 {
		desiredTokenTransferFeeConfig.DefaultFinalityTransferFeeBps = currentTokenTransferFeeConfig.DefaultBlockConfirmationsTransferFeeBps
	}
	if desiredTokenTransferFeeConfig.CustomFinalityTransferFeeBps == 0 {
		desiredTokenTransferFeeConfig.CustomFinalityTransferFeeBps = currentTokenTransferFeeConfig.CustomBlockConfirmationsTransferFeeBps
	}

	updates := make([]token_pool.TokenTransferFeeConfigUpdate, 0)

	if desiredTokenTransferFeeConfig.DestGasOverhead != currentTokenTransferFeeConfig.DestGasOverhead ||
		desiredTokenTransferFeeConfig.DestBytesOverhead != currentTokenTransferFeeConfig.DestBytesOverhead ||
		desiredTokenTransferFeeConfig.DefaultFinalityFeeUSDCents != currentTokenTransferFeeConfig.DefaultBlockConfirmationsFeeUSDCents ||
		desiredTokenTransferFeeConfig.CustomFinalityFeeUSDCents != currentTokenTransferFeeConfig.CustomBlockConfirmationsFeeUSDCents ||
		desiredTokenTransferFeeConfig.DefaultFinalityTransferFeeBps != currentTokenTransferFeeConfig.DefaultBlockConfirmationsTransferFeeBps ||
		desiredTokenTransferFeeConfig.CustomFinalityTransferFeeBps != currentTokenTransferFeeConfig.CustomBlockConfirmationsTransferFeeBps {
		updates = append(updates, token_pool.TokenTransferFeeConfigUpdate{
			DestChainSelector:                      remoteChainSelector,
			DestGasOverhead:                        desiredTokenTransferFeeConfig.DestGasOverhead,
			DestBytesOverhead:                      desiredTokenTransferFeeConfig.DestBytesOverhead,
			DefaultBlockConfirmationFeeUSDCents:    desiredTokenTransferFeeConfig.DefaultFinalityFeeUSDCents,
			CustomBlockConfirmationFeeUSDCents:     desiredTokenTransferFeeConfig.CustomFinalityFeeUSDCents,
			DefaultBlockConfirmationTransferFeeBps: desiredTokenTransferFeeConfig.DefaultFinalityTransferFeeBps,
			CustomBlockConfirmationTransferFeeBps:  desiredTokenTransferFeeConfig.CustomFinalityTransferFeeBps,
			IsEnabled:                              true,
		})
	}

	return updates, nil
}

// makeCCVUpdates returns the CCV config update to apply for the remote chain, or (nil, false, nil) if on-chain
// state already matches desired. Uses GetRequiredCCVs(amount=0) for standard CCVs and GetRequiredCCVs(amount>threshold)
// to derive threshold-only CCVs; if an input list is empty, the on-chain value is used for comparison.
func makeCCVUpdates(b cldf_ops.Bundle, chain evm.Chain, chainSelector uint64, advancedPoolHooksAddr common.Address, remoteChainSelector uint64, inboundCCVs, outboundCCVs, inboundCCVsToAddAboveThreshold, outboundCCVsToAddAboveThreshold []string) (arg *advanced_pool_hooks.CCVConfigArg, needUpdate bool, err error) {
	// MessageDirection: Outbound = 0, Inbound = 1 (IPoolV2.MessageDirection)
	const directionInbound uint8 = 1
	const directionOutbound uint8 = 0

	thresholdReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.GetThresholdAmount, chain, evm_contract.FunctionInput[any]{
		ChainSelector: chainSelector,
		Address:       advancedPoolHooksAddr,
	})
	if err != nil {
		return nil, false, fmt.Errorf("get threshold amount: %w", err)
	}
	threshold := thresholdReport.Output
	amountAboveThreshold := new(big.Int).Add(threshold, big.NewInt(1))

	getRequiredCCVs := func(amount *big.Int, direction uint8) ([]common.Address, error) {
		report, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.GetRequiredCCVs, chain, evm_contract.FunctionInput[advanced_pool_hooks.GetRequiredCCVsArgs]{
			ChainSelector: chainSelector,
			Address:       advancedPoolHooksAddr,
			Args: advanced_pool_hooks.GetRequiredCCVsArgs{
				RemoteChainSelector: remoteChainSelector,
				Amount:              amount,
				Direction:           direction,
			},
		})
		if err != nil {
			return nil, err
		}
		return report.Output, nil
	}

	stdInbound, err := getRequiredCCVs(big.NewInt(0), directionInbound)
	if err != nil {
		return nil, false, fmt.Errorf("get required CCVs inbound amount=0: %w", err)
	}
	stdOutbound, err := getRequiredCCVs(big.NewInt(0), directionOutbound)
	if err != nil {
		return nil, false, fmt.Errorf("get required CCVs outbound amount=0: %w", err)
	}
	aboveInbound, err := getRequiredCCVs(amountAboveThreshold, directionInbound)
	if err != nil {
		return nil, false, fmt.Errorf("get required CCVs inbound above threshold: %w", err)
	}
	aboveOutbound, err := getRequiredCCVs(amountAboveThreshold, directionOutbound)
	if err != nil {
		return nil, false, fmt.Errorf("get required CCVs outbound above threshold: %w", err)
	}
	onChainThresholdInbound := v17seq.AddressesNotIn(aboveInbound, stdInbound)
	onChainThresholdOutbound := v17seq.AddressesNotIn(aboveOutbound, stdOutbound)

	parse := func(ss []string) []common.Address {
		out := make([]common.Address, len(ss))
		for i, s := range ss {
			out[i] = common.HexToAddress(s)
		}
		return out
	}
	desiredInbound := parse(inboundCCVs)
	if len(desiredInbound) == 0 {
		desiredInbound = stdInbound
	}
	desiredOutbound := parse(outboundCCVs)
	if len(desiredOutbound) == 0 {
		desiredOutbound = stdOutbound
	}
	desiredThresholdInbound := parse(inboundCCVsToAddAboveThreshold)
	if len(desiredThresholdInbound) == 0 {
		desiredThresholdInbound = onChainThresholdInbound
	}
	desiredThresholdOutbound := parse(outboundCCVsToAddAboveThreshold)
	if len(desiredThresholdOutbound) == 0 {
		desiredThresholdOutbound = onChainThresholdOutbound
	}

	addrEq := func(a, b common.Address) bool { return a == b }
	if v17seq.UnorderedSliceEqual(desiredInbound, stdInbound, addrEq) &&
		v17seq.UnorderedSliceEqual(desiredOutbound, stdOutbound, addrEq) &&
		v17seq.UnorderedSliceEqual(desiredThresholdInbound, onChainThresholdInbound, addrEq) &&
		v17seq.UnorderedSliceEqual(desiredThresholdOutbound, onChainThresholdOutbound, addrEq) {
		return nil, false, nil
	}
	return &advanced_pool_hooks.CCVConfigArg{
		RemoteChainSelector:   remoteChainSelector,
		OutboundCCVs:          desiredOutbound,
		ThresholdOutboundCCVs: desiredThresholdOutbound,
		InboundCCVs:           desiredInbound,
		ThresholdInboundCCVs:  desiredThresholdInbound,
	}, true, nil
}
