package tokens

import (
	"bytes"
	"fmt"
	"math/big"
	"slices"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	evm_contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/finality"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/operations/token_pool"
	v17seq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_0_0/sequences"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
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
	// RegistryAddress is the TokenAdminRegistry address; used by ConfigureTokenPoolForRemoteChains for upgrade validation.
	RegistryAddress common.Address
	// TokenAddress is the token address; used with RegistryAddress for supported-chains validation on upgrade.
	TokenAddress common.Address
	// RemoteChainAlreadySupported is true when the pool already has this remote chain in its supported list (avoids an on-chain read).
	RemoteChainAlreadySupported bool
}

func (c ConfigureTokenPoolForRemoteChainInput) Validate(chain evm.Chain) error {
	if c.ChainSelector != chain.Selector {
		return fmt.Errorf("chain selector %d does not match chain %s", c.ChainSelector, chain)
	}
	if err := c.RemoteChainConfig.Validate(); err != nil {
		return fmt.Errorf("invalid remote chain config: %w", err)
	}
	return nil
}

var ConfigureTokenPoolForRemoteChain = cldf_ops.NewSequence(
	"configure-token-pool-for-remote-chain",
	semver.MustParse("2.0.0"),
	"Configures a token pool on an EVM chain for transfers with other chains",
	func(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}
		writes := make([]evm_contract.WriteOutput, 0)

		localDecimalsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetTokenDecimals, chain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenPoolAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token decimals: %w", err)
		}

		tvReport, err := cldf_ops.ExecuteOperation(b, type_and_version.GetTypeAndVersion, chain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       input.TokenPoolAddress,
			Args:          struct{}{},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get type and version of token pool: %w", err)
		}

		// Fetch rate limit buckets from input config
		outbounds := input.RemoteChainConfig.GetOutboundRateLimitBuckets()
		inbounds := input.RemoteChainConfig.GetInboundRateLimitBuckets()
		defaultOutboundBucket, defaultOutboundExists := outbounds.DefaultBucket()
		defaultInboundBucket, defaultInboundExists := inbounds.DefaultBucket()
		if defaultOutboundExists != defaultInboundExists {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"default outbound and inbound rate limits must both be specified together in deployment input, provided on MigrationMetadata, or fully omitted from both for remote chain %d",
				input.RemoteChainSelector,
			)
		}

		// Resolve default outbound/inbound TPRL limits using the following precedence rules:
		// (1) if both default buckets are provided in the deployment input → GenerateTPRLConfigs (explicit user input has the highest precedence)
		// (2) else if rate limits are provided via migration metadata → use imported legacy rate limits as-is (top-level changeset will have already handled decimal rebasing)
		// (3) else if rate limits are omitted from input and metadata → use the same limits that are on-chain if the remote is already supported otherwise use disabled limits
		// (4) else error — partial defaults in input or metadata are too ambiguous to safely proceed
		var outboundRateLimiterConfig tokens.RateLimiterConfig
		var inboundRateLimiterConfig tokens.RateLimiterConfig
		imported := input.RemoteChainConfig.MigrationMetadata
		switch {
		case defaultOutboundExists && defaultInboundExists:
			outboundRateLimiterConfig, inboundRateLimiterConfig = tokens.GenerateTPRLConfigs(
				defaultOutboundBucket.RateLimit,
				defaultInboundBucket.RateLimit,
				localDecimalsReport.Output,
				input.RemoteChainConfig.RemoteDecimals,
				chain.Family(),
				tvReport.Output.Version,
				tvReport.Output.Type.String(),
			)

		case imported.LegacyRateLimits != nil:
			outboundRateLimiterConfig = imported.LegacyRateLimits.Outbound
			inboundRateLimiterConfig = imported.LegacyRateLimits.Inbound

		case (!defaultOutboundExists && !defaultInboundExists) && imported.LegacyRateLimits == nil:
			if input.RemoteChainAlreadySupported {
				// Idempotent behavior: if we're re-calling this sequence and no rate limits are
				// specified, then we re-use whatever is currently onchain to avoid accidentally
				// overwriting existing onchain config
				rlStateReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetCurrentRateLimiterState, chain, evm_contract.FunctionInput[token_pool.GetCurrentRateLimiterStateArgs]{
					ChainSelector: input.ChainSelector,
					Address:       input.TokenPoolAddress,
					Args: token_pool.GetCurrentRateLimiterStateArgs{
						RemoteChainSelector: input.RemoteChainSelector,
						FastFinality:        false,
					},
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get default rate limiter state for remote chain %d: %w", input.RemoteChainSelector, err)
				}
				outboundRateLimiterConfig = tokens.RateLimiterConfig{
					IsEnabled: rlStateReport.Output.OutboundRateLimiterState.IsEnabled,
					Capacity:  rlStateReport.Output.OutboundRateLimiterState.Capacity,
					Rate:      rlStateReport.Output.OutboundRateLimiterState.Rate,
				}
				inboundRateLimiterConfig = tokens.RateLimiterConfig{
					IsEnabled: rlStateReport.Output.InboundRateLimiterState.IsEnabled,
					Capacity:  rlStateReport.Output.InboundRateLimiterState.Capacity,
					Rate:      rlStateReport.Output.InboundRateLimiterState.Rate,
				}
			} else {
				// If none of the above sources provided default rate limits, we will fall back to
				// a disabled rate limiter by default (idempotent if chain is not yet supported).
				outboundRateLimiterConfig = tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
				inboundRateLimiterConfig = tokens.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
			}

		default:
			return sequences.OnChainOutput{}, fmt.Errorf(
				"default outbound and inbound rate limits must both be specified together in deployment input, provided on MigrationMetadata, or fully omitted from both for remote chain %d",
				input.RemoteChainSelector,
			)
		}

		// Default rate limit bucket should be fully resolved at this point
		rlInputs := []tokens.TPRLRateLimitBucket{
			{
				OutboundRateLimiterConfig: outboundRateLimiterConfig,
				InboundRateLimiterConfig:  inboundRateLimiterConfig,
				FastFinality:              false,
			},
		}

		// Handle fast finality rate limits
		ffOutboundBucket, ffOutboundExists := outbounds.FastFinalityBucket()
		ffInboundBucket, ffInboundExists := inbounds.FastFinalityBucket()
		if ffOutboundExists != ffInboundExists {
			return sequences.OnChainOutput{}, fmt.Errorf(
				"fast-finality rate limits for remote chain %d require both OutboundRateLimits and InboundRateLimits to include a bucket with fastFinality=true, or neither",
				input.RemoteChainSelector,
			)
		}
		if ffOutboundExists && ffInboundExists {
			customOutboundRateLimiterConfig, customInboundRateLimiterConfig := tokens.GenerateTPRLConfigs(
				ffOutboundBucket.RateLimit,
				ffInboundBucket.RateLimit,
				localDecimalsReport.Output,
				input.RemoteChainConfig.RemoteDecimals,
				chain.Family(),
				tvReport.Output.Version,
				tvReport.Output.Type.String(),
			)
			rlInputs = append(rlInputs, tokens.TPRLRateLimitBucket{
				OutboundRateLimiterConfig: customOutboundRateLimiterConfig,
				InboundRateLimiterConfig:  customInboundRateLimiterConfig,
				FastFinality:              true,
			})
		}

		// Set CCVs for the remote chain (idempotent: only apply when on-chain differs from desired)
		if input.AdvancedPoolHooks != (common.Address{}) {
			ccvArg, needCCVUpdate, err := makeCCVUpdates(
				b, chain, input.ChainSelector, input.AdvancedPoolHooks, input.RemoteChainSelector,
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

		// If the chain is already supported
		// 1. Check remote token, remove and re-add remote config if requested remote token is different
		// 2. Check existing rate limiters and update if necessary
		// 3. Check existing remote pools and add requested remote pool if it does not exist
		removes := make([]uint64, 0, 1) // Cap == 1 because we may need to remove the chain if the remote token is different
		if input.RemoteChainAlreadySupported {
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
				// Check and update TPRL buckets in rlInputs (default always; fast-finality only when paired FF floats exist on RemoteChainConfig).
				rateLimitersReport, err := maybeUpdateRateLimiters(
					b,
					chain,
					input.ChainSelector,
					input.TokenPoolAddress,
					input.RemoteChainSelector,
					rlInputs,
				)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to maybe update rate limiters: %w", err)
				}
				if rateLimitersReport != nil {
					writes = append(writes, *rateLimitersReport)
				}

				// Check existing remote pools and add any missing (active pool's remote pools first for upgrade, then requested pool)
				getRemotePoolsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetRemotePools, chain, evm_contract.FunctionInput[uint64]{
					ChainSelector: input.ChainSelector,
					Address:       input.TokenPoolAddress,
					Args:          input.RemoteChainSelector,
				})
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to get remote pools: %w", err)
				}
				existingPools := getRemotePoolsReport.Output
				containsPool := func(pool []byte) bool {
					return slices.ContainsFunc(existingPools, func(addr []byte) bool { return bytes.Equal(addr, pool) })
				}
				// Add active pool's remote pools first to protect inflight messages during cutover.
				for _, activePoolAddr := range imported.LegacyRemotePools {
					padded := common.LeftPadBytes(activePoolAddr, 32)
					if !containsPool(padded) {
						addReport, err := cldf_ops.ExecuteOperation(b, token_pool.AddRemotePool, chain, evm_contract.FunctionInput[token_pool.AddRemotePoolArgs]{
							ChainSelector: input.ChainSelector,
							Address:       input.TokenPoolAddress,
							Args: token_pool.AddRemotePoolArgs{
								RemoteChainSelector: input.RemoteChainSelector,
								RemotePoolAddress:   padded,
							},
						})
						if err != nil {
							return sequences.OnChainOutput{}, fmt.Errorf("failed to add active pool remote pool: %w", err)
						}
						writes = append(writes, addReport.Output)
						existingPools = append(existingPools, padded)
					}
				}
				if !containsPool(common.LeftPadBytes(input.RemoteChainConfig.RemotePool, 32)) {
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

				// Update token transfer fee configuration (chain is already supported)
				tokenTransferFeeWrites, err := applyTokenTransferFeeConfigIfNeeded(b, chain, input, input.RemoteChainSelector)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to apply token transfer fee config updates: %w", err)
				}
				writes = append(writes, tokenTransferFeeWrites...)

				// Return early as no further action is required
				batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
				if err != nil {
					return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
				}

				return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
			}
		}

		// If the chain is not supported, apply the config for the remote chain
		// Build remote pool list: active pool's remote pools first (for upgrade cutover), then the requested pool.
		remotePoolAddresses := make([][]byte, 0, len(imported.LegacyRemotePools)+1)
		for _, p := range imported.LegacyRemotePools {
			remotePoolAddresses = append(remotePoolAddresses, common.LeftPadBytes(p, 32))
		}
		inputPoolPadded := common.LeftPadBytes(input.RemoteChainConfig.RemotePool, 32)
		if !slices.ContainsFunc(remotePoolAddresses, func(b []byte) bool { return bytes.Equal(b, inputPoolPadded) }) {
			remotePoolAddresses = append(remotePoolAddresses, inputPoolPadded)
		}
		applyChainUpdatesReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyChainUpdates, chain, evm_contract.FunctionInput[token_pool.ApplyChainUpdatesArgs]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenPoolAddress,
			Args: token_pool.ApplyChainUpdatesArgs{
				RemoteChainSelectorsToRemove: removes,
				ChainsToAdd: []token_pool.ChainUpdate{
					{
						RemoteChainSelector: input.RemoteChainSelector,
						RemotePoolAddresses: remotePoolAddresses,
						RemoteTokenAddress:  common.LeftPadBytes(input.RemoteChainConfig.RemoteToken, 32),
						OutboundRateLimiterConfig: token_pool.Config{
							IsEnabled: outboundRateLimiterConfig.IsEnabled,
							Capacity:  outboundRateLimiterConfig.Capacity,
							Rate:      outboundRateLimiterConfig.Rate,
						},
						InboundRateLimiterConfig: token_pool.Config{
							IsEnabled: inboundRateLimiterConfig.IsEnabled,
							Capacity:  inboundRateLimiterConfig.Capacity,
							Rate:      inboundRateLimiterConfig.Rate,
						},
					},
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply chain updates: %w", err)
		}
		writes = append(writes, applyChainUpdatesReport.Output)

		// Check and update TPRL buckets in rlInputs (default always; fast-finality only when paired FF floats exist on RemoteChainConfig).
		rateLimitersReport, err := maybeUpdateRateLimiters(
			b,
			chain,
			input.ChainSelector,
			input.TokenPoolAddress,
			input.RemoteChainSelector,
			rlInputs,
		)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to maybe update rate limiters: %w", err)
		}
		if rateLimitersReport != nil {
			writes = append(writes, *rateLimitersReport)
		}

		// Apply token transfer fee config on the v2 pool after the remote chain exists.
		tokenTransferFeeWrites, err := applyTokenTransferFeeConfigIfNeeded(b, chain, input, input.RemoteChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply token transfer fee config updates: %w", err)
		}
		writes = append(writes, tokenTransferFeeWrites...)

		batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	},
)

// maybeUpdateRateLimiters checks and updates each TPRL bucket listed in desiredRateLimiterConfigs (by FastFinality flag).
// Returns nil if no update is needed.
func maybeUpdateRateLimiters(
	b cldf_ops.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	tokenPoolAddress common.Address,
	remoteChainSelector uint64,
	desiredRateLimiterConfigs []tokens.TPRLRateLimitBucket,
) (*evm_contract.WriteOutput, error) {
	args := []token_pool.RateLimitConfigArgs{}
	for _, desiredRL := range desiredRateLimiterConfigs {
		// Check existing rate limiters
		rateLimiterStateReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetCurrentRateLimiterState, chain, evm_contract.FunctionInput[token_pool.GetCurrentRateLimiterStateArgs]{
			ChainSelector: chainSelector,
			Address:       tokenPoolAddress,
			Args: token_pool.GetCurrentRateLimiterStateArgs{
				RemoteChainSelector: remoteChainSelector,
				FastFinality:        desiredRL.FastFinality,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get rate limiter state: %w", err)
		}

		// Update whenever on-chain state differs from the resolved desired config, including
		// IsEnabled=false (explicit limiter off). Stopping traffic without disabling the limiter remains
		// represented as IsEnabled=true with rate and capacity zero.
		onchainRL := rateLimiterStateReport.Output
		if !rateLimiterConfigsEqual(onchainRL.InboundRateLimiterState, desiredRL.InboundRateLimiterConfig) ||
			!rateLimiterConfigsEqual(onchainRL.OutboundRateLimiterState, desiredRL.OutboundRateLimiterConfig) {
			args = append(args, token_pool.RateLimitConfigArgs{
				OutboundRateLimiterConfig: token_pool.Config{
					IsEnabled: desiredRL.OutboundRateLimiterConfig.IsEnabled,
					Capacity:  desiredRL.OutboundRateLimiterConfig.Capacity,
					Rate:      desiredRL.OutboundRateLimiterConfig.Rate,
				},
				InboundRateLimiterConfig: token_pool.Config{
					IsEnabled: desiredRL.InboundRateLimiterConfig.IsEnabled,
					Capacity:  desiredRL.InboundRateLimiterConfig.Capacity,
					Rate:      desiredRL.InboundRateLimiterConfig.Rate,
				},
				RemoteChainSelector: remoteChainSelector,
				FastFinality:        desiredRL.FastFinality,
			})
		}
	}

	if len(args) > 0 {
		setInboundRateLimiterReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetRateLimitConfig, chain, evm_contract.FunctionInput[[]token_pool.RateLimitConfigArgs]{
			ChainSelector: chainSelector,
			Address:       tokenPoolAddress,
			Args:          args,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to set rate limiters config: %w", err)
		}
		return &setInboundRateLimiterReport.Output, nil
	}

	return nil, nil
}

// rateLimiterConfigsEqual returns true if the current rate limiter config on-chain matches the desired config.
func rateLimiterConfigsEqual(current token_pool.TokenBucket, desired tokens.RateLimiterConfig) bool {
	return current.IsEnabled == desired.IsEnabled &&
		current.Capacity.Cmp(desired.Capacity) == 0 &&
		current.Rate.Cmp(desired.Rate) == 0
}

func applyTokenTransferFeeConfigIfNeeded(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput, remoteChainSelector uint64) ([]evm_contract.WriteOutput, error) {
	ttfcArgs, err := makeTokenTransferFeeConfigUpdates(b, chain, input, remoteChainSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to make token transfer fee config updates: %w", err)
	}
	if len(ttfcArgs.DisableTokenTransferFeeConfigs) == 0 && len(ttfcArgs.TokenTransferFeeConfigArgs) == 0 {
		return nil, nil
	}
	report, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyTokenTransferFeeConfigUpdates, chain, evm_contract.FunctionInput[token_pool.ApplyTokenTransferFeeConfigUpdatesArgs]{
		ChainSelector: input.ChainSelector,
		Address:       input.TokenPoolAddress,
		Args:          ttfcArgs,
	})
	if err != nil {
		return nil, err
	}
	return []evm_contract.WriteOutput{report.Output}, nil
}

func makeTokenTransferFeeConfigUpdates(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput, remoteChainSelector uint64) (token_pool.ApplyTokenTransferFeeConfigUpdatesArgs, error) {
	desiredTokenTransferFeeConfig := input.RemoteChainConfig.TokenTransferFeeConfig
	if desiredTokenTransferFeeConfig == nil {
		return token_pool.ApplyTokenTransferFeeConfigUpdatesArgs{}, nil
	}

	report, err := cldf_ops.ExecuteOperation(
		b, token_pool.GetTokenTransferFeeConfig, chain,
		evm_contract.FunctionInput[token_pool.GetTokenTransferFeeConfigArgs]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenPoolAddress,
			Args: token_pool.GetTokenTransferFeeConfigArgs{
				Arg0:              common.Address{},            // unused
				DestChainSelector: remoteChainSelector,         // this IS used
				Arg2:              finality.RawWaitForFinality, // unused
				Arg3:              []byte{},                    // unused
			},
		},
		cldf_ops.WithForceExecute[evm_contract.FunctionInput[token_pool.GetTokenTransferFeeConfigArgs], evm.Chain](),
	)
	if err != nil {
		return token_pool.ApplyTokenTransferFeeConfigUpdatesArgs{}, fmt.Errorf("failed to get token transfer fee config: %w", err)
	}

	defaultConfig := tokens.GetDefaultChainAgnosticTokenTransferFeeConfig(
		input.ChainSelector,
		input.RemoteChainSelector,
	)

	currentConfig := tokens.TokenTransferFeeConfig{
		DefaultFinalityTransferFeeBps: report.Output.FinalityTransferFeeBps,
		CustomFinalityTransferFeeBps:  report.Output.FastFinalityTransferFeeBps,
		DefaultFinalityFeeUSDCents:    report.Output.FinalityFeeUSDCents,
		CustomFinalityFeeUSDCents:     report.Output.FastFinalityFeeUSDCents,
		DestBytesOverhead:             report.Output.DestBytesOverhead,
		DestGasOverhead:               report.Output.DestGasOverhead,
		IsEnabled:                     report.Output.IsEnabled,
	}

	// Resolution strategy:
	// (1) If on-chain config is enabled, merge it with the user's provided config (giving precedence to user's config)
	// (2) Fall back to sensible defaults merged with user's provided config (giving precedence to user's config)
	var resolvedConfig tokens.TokenTransferFeeConfig
	if currentConfig.IsEnabled {
		resolvedConfig = desiredTokenTransferFeeConfig.MergeWith(currentConfig)
	} else {
		resolvedConfig = desiredTokenTransferFeeConfig.MergeWith(defaultConfig)
	}

	if !resolvedConfig.IsEnabled && !currentConfig.IsEnabled {
		return token_pool.ApplyTokenTransferFeeConfigUpdatesArgs{}, nil
	}

	if resolvedConfig == currentConfig {
		return token_pool.ApplyTokenTransferFeeConfigUpdatesArgs{}, nil
	}

	if !resolvedConfig.IsEnabled {
		return token_pool.ApplyTokenTransferFeeConfigUpdatesArgs{
			DisableTokenTransferFeeConfigs: []uint64{
				remoteChainSelector,
			},
		}, nil
	} else {
		return token_pool.ApplyTokenTransferFeeConfigUpdatesArgs{
			TokenTransferFeeConfigArgs: []token_pool.TokenTransferFeeConfigArgs{
				{
					DestChainSelector: remoteChainSelector,
					TokenTransferFeeConfig: token_pool.TokenTransferFeeConfig{
						FastFinalityTransferFeeBps: resolvedConfig.CustomFinalityTransferFeeBps,
						FastFinalityFeeUSDCents:    resolvedConfig.CustomFinalityFeeUSDCents,
						FinalityTransferFeeBps:     resolvedConfig.DefaultFinalityTransferFeeBps,
						FinalityFeeUSDCents:        resolvedConfig.DefaultFinalityFeeUSDCents,
						DestBytesOverhead:          resolvedConfig.DestBytesOverhead,
						DestGasOverhead:            resolvedConfig.DestGasOverhead,
						IsEnabled:                  resolvedConfig.IsEnabled,
					},
				},
			},
		}, nil
	}
}

// makeCCVUpdates returns the CCV config update to apply for the remote chain, or (nil, false, nil) if on-chain
// state already matches desired. Uses GetRequiredCCVs(amount=0) for standard CCVs and GetRequiredCCVs(amount>threshold)
// to derive threshold-only CCVs; if an input list is empty, the on-chain value is used for comparison.
func makeCCVUpdates(b cldf_ops.Bundle, chain evm.Chain, chainSelector uint64, advancedPoolHooksAddr common.Address, remoteChainSelector uint64, inboundCCVs, outboundCCVs, inboundCCVsToAddAboveThreshold, outboundCCVsToAddAboveThreshold []string) (arg *advanced_pool_hooks.CCVConfigArg, needUpdate bool, err error) {
	// MessageDirection: Outbound = 0, Inbound = 1 (IPoolV2.MessageDirection)
	const directionInbound uint8 = 1
	const directionOutbound uint8 = 0

	thresholdReport, err := cldf_ops.ExecuteOperation(b, advanced_pool_hooks.GetThresholdAmount, chain, evm_contract.FunctionInput[struct{}]{
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
