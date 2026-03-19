package tokens

import (
	"bytes"
	"fmt"
	"math/big"
	"slices"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	routerops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	onrampops_v150 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	onrampops_v160 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	fqops_v163 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_3/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"

	fqops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"

	v17seq "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/token_pool"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/type_and_version"
	token_pool_v150 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_token_pool_and_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	token_pool_v161 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"

	onrampops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/onramp"
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
	// RegistryAddress is the TokenAdminRegistry address; if set, used to fetch the active pool for upgrade import.
	RegistryAddress common.Address
	// TokenAddress is the token address; if set with RegistryAddress, used to fetch the active pool for upgrade import.
	TokenAddress common.Address
	// RemoteChainAlreadySupported is true when the pool already has this remote chain in its supported list (avoids an on-chain read).
	RemoteChainAlreadySupported bool
}

func (c ConfigureTokenPoolForRemoteChainInput) Validate(chain evm.Chain) error {
	if c.ChainSelector != chain.Selector {
		return fmt.Errorf("chain selector %d does not match chain %s", c.ChainSelector, chain)
	}
	return nil
}

// activePoolImportedConfig holds config imported from an active pool (< 2.0.0) for upgrade cutover.
type activePoolImportedConfig struct {
	DefaultOutbound *tokens.RateLimiterConfig
	DefaultInbound  *tokens.RateLimiterConfig
	RemotePools     [][]byte
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

		imported, err := importConfigFromActivePool(b, chain, input.ChainSelector, input.RegistryAddress, input.TokenAddress, input.RemoteChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, err
		}
		var importedDefaultOutbound, importedDefaultInbound *tokens.RateLimiterConfig
		var activePoolRemotePools [][]byte
		if imported != nil {
			importedDefaultOutbound = imported.DefaultOutbound
			importedDefaultInbound = imported.DefaultInbound
			activePoolRemotePools = imported.RemotePools
		}

		localDecimalsReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetTokenDecimals, chain, evm_contract.FunctionInput[struct{}]{
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
		// If input did not provide default finality rate limits, use imported from active pool when available.
		if importedDefaultOutbound != nil && !input.RemoteChainConfig.DefaultFinalityOutboundRateLimiterConfig.IsEnabled {
			defaultFinalityOutboundRateLimiterConfig = *importedDefaultOutbound
		}
		if importedDefaultInbound != nil && !input.RemoteChainConfig.DefaultFinalityInboundRateLimiterConfig.IsEnabled {
			defaultFinalityInboundRateLimiterConfig = *importedDefaultInbound
		}

		customFinalityOutboundRateLimiterConfig, customFinalityInboundRateLimiterConfig := tokens.GenerateTPRLConfigs(
			input.RemoteChainConfig.CustomFinalityOutboundRateLimiterConfig,
			input.RemoteChainConfig.CustomFinalityInboundRateLimiterConfig,
			localDecimalsReport.Output,
			input.RemoteChainConfig.RemoteDecimals,
			chain.Family(),
			token_pool.Version,
		)

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
				for _, activePoolAddr := range activePoolRemotePools {
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
		remotePoolAddresses := make([][]byte, 0, len(activePoolRemotePools)+1)
		for _, p := range activePoolRemotePools {
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
						RemoteChainSelector:       input.RemoteChainSelector,
						RemotePoolAddresses:       remotePoolAddresses,
						RemoteTokenAddress:        common.LeftPadBytes(input.RemoteChainConfig.RemoteToken, 32),
						OutboundRateLimiterConfig: tokensRateLimiterToConfig(defaultFinalityOutboundRateLimiterConfig),
						InboundRateLimiterConfig:  tokensRateLimiterToConfig(defaultFinalityInboundRateLimiterConfig),
					},
				},
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply chain updates: %w", err)
		}
		writes = append(writes, applyChainUpdatesReport.Output)

		// Update token transfer fee configuration (chain was just added)
		tokenTransferFeeWrites, err := applyTokenTransferFeeConfigIfNeeded(b, chain, input, input.RemoteChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to apply token transfer fee config updates: %w", err)
		}
		writes = append(writes, tokenTransferFeeWrites...)

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

		// Update token transfer fee configuration (after applyChainUpdates so chain exists on pool).
		tokenTransferFeeConfigUpdates, err := makeTokenTransferFeeConfigUpdates(b, chain, input, input.RemoteChainSelector)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to make token transfer fee config updates: %w", err)
		}
		if len(tokenTransferFeeConfigUpdates) > 0 {
			applyTokenTransferFeeConfigUpdatesReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyTokenTransferFeeConfigUpdates, chain, evm_contract.FunctionInput[token_pool.ApplyTokenTransferFeeConfigUpdatesArgs]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenPoolAddress,
				Args: token_pool.ApplyTokenTransferFeeConfigUpdatesArgs{
					TokenTransferFeeConfigArgs: tokenTransferFeeConfigUpdates,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply token transfer fee config updates: %w", err)
			}
			writes = append(writes, applyTokenTransferFeeConfigUpdatesReport.Output)
		}

		batchOp, err := evm_contract.NewBatchOperationFromWrites(writes)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from writes: %w", err)
		}

		return sequences.OnChainOutput{BatchOps: []mcms_types.BatchOperation{batchOp}}, nil
	},
)

// importConfigFromActivePool fetches the active pool from the TokenAdminRegistry and, if its version is < 2.0.0,
// imports rate limit state and remote pools for the given remote chain. Uses 1.5.0 ops for pools < 1.5.1 (including
// proxy/previousPool handling) and 1.6.1 ops for 1.5.1 <= version < 2.0.0.
// Returns nil when registry or token address is zero, when there is no active pool, or when the active pool is >= 2.0.0.
func importConfigFromActivePool(
	b cldf_ops.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	registryAddress, tokenAddress common.Address,
	remoteChainSelector uint64,
) (*activePoolImportedConfig, error) {
	if registryAddress == (common.Address{}) || tokenAddress == (common.Address{}) {
		return nil, nil
	}
	tokenConfigReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.GetTokenConfig, chain, evm_contract.FunctionInput[common.Address]{
		ChainSelector: chainSelector,
		Address:       registryAddress,
		Args:          tokenAddress,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get token config from registry: %w", err)
	}
	activePool := tokenConfigReport.Output.TokenPool
	if activePool == (common.Address{}) {
		return nil, nil
	}
	typeAndVersionReport, err := cldf_ops.ExecuteOperation(b, type_and_version.GetTypeAndVersion, chain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: chainSelector,
		Address:       activePool,
		Args:          struct{}{},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get active pool type and version: %w", err)
	}
	tav := typeAndVersionReport.Output
	if tav.Version.GreaterThanEqual(semver.MustParse("2.0.0")) {
		// Configuration import from another 2.0.0 pool is not currently supported
		return nil, nil
	}
	if tav.Version.LessThan(semver.MustParse("1.5.1")) {
		return importConfigFromActivePoolV150(b, chain, chainSelector, activePool, remoteChainSelector, tav)
	}
	return importConfigFromActivePoolV161(b, chain, chainSelector, activePool, remoteChainSelector)
}

// importConfigFromActivePoolV150 imports rate limits and the single remote pool for the given remote chain
// using 1.5.0 operations. If activePool type contains "Proxy", fetches previousPool; if previousPool version < 1.4.0
// uses the proxy (activePool) for rate limits, otherwise uses previousPool for rate limits. Remote pool is always
// read from activePool (proxy exposes getRemotePool).
func importConfigFromActivePoolV150(
	b cldf_ops.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	activePool common.Address,
	remoteChainSelector uint64,
	tav type_and_version.TypeAndVersion,
) (*activePoolImportedConfig, error) {
	typeStr := string(tav.Type)
	poolForRateLimits := activePool
	if strings.Contains(typeStr, "Proxy") {
		prevReport, err := cldf_ops.ExecuteOperation(b, token_pool_v150.GetPreviousPool, chain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: chainSelector,
			Address:       activePool,
			Args:          struct{}{},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get previous pool from proxy: %w", err)
		}
		previousPool := prevReport.Output
		if previousPool != (common.Address{}) {
			prevTVReport, err := cldf_ops.ExecuteOperation(b, type_and_version.GetTypeAndVersion, chain, evm_contract.FunctionInput[struct{}]{
				ChainSelector: chainSelector,
				Address:       previousPool,
				Args:          struct{}{},
			})
			if err != nil {
				return nil, fmt.Errorf("failed to get previous pool type and version: %w", err)
			}
			if prevTVReport.Output.Version.LessThan(semver.MustParse("1.4.0")) {
				poolForRateLimits = activePool
			} else {
				poolForRateLimits = previousPool
			}
		}
	}
	cfg, err := fetchRateLimitsAndRemotePoolV150(b, chain, chainSelector, poolForRateLimits, activePool, remoteChainSelector)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// fetchRateLimitsAndRemotePoolV150 fetches rate limiter state and the single remote pool for the given remote chain
// using 1.5.0 operations. poolForRateLimits is the address to read rate limits from; poolForRemotePool is the
// address to read getRemotePool from (e.g. proxy when using previousPool for rate limits).
func fetchRateLimitsAndRemotePoolV150(
	b cldf_ops.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	poolForRateLimits, poolForRemotePool common.Address,
	remoteChainSelector uint64,
) (*activePoolImportedConfig, error) {
	inboundReport, err := cldf_ops.ExecuteOperation(b, token_pool_v150.GetCurrentInboundRateLimiterState, chain, evm_contract.FunctionInput[uint64]{
		ChainSelector: chainSelector,
		Address:       poolForRateLimits,
		Args:          remoteChainSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get inbound rate limiter state: %w", err)
	}
	outboundReport, err := cldf_ops.ExecuteOperation(b, token_pool_v150.GetCurrentOutboundRateLimiterState, chain, evm_contract.FunctionInput[uint64]{
		ChainSelector: chainSelector,
		Address:       poolForRateLimits,
		Args:          remoteChainSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get outbound rate limiter state: %w", err)
	}
	remotePoolReport, err := cldf_ops.ExecuteOperation(b, token_pool_v150.GetRemotePool, chain, evm_contract.FunctionInput[uint64]{
		ChainSelector: chainSelector,
		Address:       poolForRemotePool,
		Args:          remoteChainSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get remote pool: %w", err)
	}
	remotePools := [][]byte{}
	if len(remotePoolReport.Output) > 0 {
		remotePools = [][]byte{remotePoolReport.Output}
	}
	return &activePoolImportedConfig{
		DefaultOutbound: tokenBucketV150ToRateLimiterConfig(outboundReport.Output),
		DefaultInbound:  tokenBucketV150ToRateLimiterConfig(inboundReport.Output),
		RemotePools:     remotePools,
	}, nil
}

// importConfigFromActivePoolV161 imports rate limits and remote pools using 1.6.1 operations (for active pool version >= 1.5.1 and < 2.0.0).
func importConfigFromActivePoolV161(
	b cldf_ops.Bundle,
	chain evm.Chain,
	chainSelector uint64,
	activePool common.Address,
	remoteChainSelector uint64,
) (*activePoolImportedConfig, error) {
	inboundReport, err := cldf_ops.ExecuteOperation(b, token_pool_v161.GetCurrentInboundRateLimiterState, chain, evm_contract.FunctionInput[uint64]{
		ChainSelector: chainSelector,
		Address:       activePool,
		Args:          remoteChainSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get active pool inbound rate limiter state: %w", err)
	}
	outboundReport, err := cldf_ops.ExecuteOperation(b, token_pool_v161.GetCurrentOutboundRateLimiterState, chain, evm_contract.FunctionInput[uint64]{
		ChainSelector: chainSelector,
		Address:       activePool,
		Args:          remoteChainSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get active pool outbound rate limiter state: %w", err)
	}
	remotePoolsReport, err := cldf_ops.ExecuteOperation(b, token_pool_v161.GetRemotePools, chain, evm_contract.FunctionInput[uint64]{
		ChainSelector: chainSelector,
		Address:       activePool,
		Args:          remoteChainSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get active pool remote pools: %w", err)
	}
	return &activePoolImportedConfig{
		DefaultOutbound: tokenBucketToRateLimiterConfig(outboundReport.Output),
		DefaultInbound:  tokenBucketToRateLimiterConfig(inboundReport.Output),
		RemotePools:     remotePoolsReport.Output,
	}, nil
}

// tokenBucketV150ToRateLimiterConfig converts a 1.5.0 (proxy bindings) RateLimiterTokenBucket to tokens.RateLimiterConfig.
func tokenBucketV150ToRateLimiterConfig(b token_pool_v150.TokenBucket) *tokens.RateLimiterConfig {
	return &tokens.RateLimiterConfig{
		IsEnabled: b.IsEnabled,
		Capacity:  new(big.Int).Set(b.Capacity),
		Rate:      new(big.Int).Set(b.Rate),
	}
}

// tokenBucketToRateLimiterConfig converts a 1.6.1 TokenBucket to tokens.RateLimiterConfig.
func tokenBucketToRateLimiterConfig(b token_pool_v161.TokenBucket) *tokens.RateLimiterConfig {
	return &tokens.RateLimiterConfig{
		IsEnabled: b.IsEnabled,
		Capacity:  new(big.Int).Set(b.Capacity),
		Rate:      new(big.Int).Set(b.Rate),
	}
}

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
			RemoteChainSelector:      remoteChainSelector,
			CustomBlockConfirmations: customBlockConfirmation,
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
		setInboundRateLimiterReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetRateLimitConfig, chain, evm_contract.FunctionInput[[]token_pool.RateLimitConfigArgs]{
			ChainSelector: chainSelector,
			Address:       tokenPoolAddress,
			Args: []token_pool.RateLimitConfigArgs{
				{
					RemoteChainSelector:       remoteChainSelector,
					CustomBlockConfirmations:  customBlockConfirmation,
					InboundRateLimiterConfig:  tokensRateLimiterToConfig(desiredInboundRateLimiterConfig),
					OutboundRateLimiterConfig: tokensRateLimiterToConfig(desiredOutboundRateLimiterConfig),
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
func rateLimiterConfigsEqual(current token_pool.TokenBucket, desired tokens.RateLimiterConfig) bool {
	return current.IsEnabled == desired.IsEnabled &&
		current.Capacity.Cmp(desired.Capacity) == 0 &&
		current.Rate.Cmp(desired.Rate) == 0
}

// tokensRateLimiterToConfig converts tokens.RateLimiterConfig to token_pool.Config.
func tokensRateLimiterToConfig(c tokens.RateLimiterConfig) token_pool.Config {
	return token_pool.Config{
		IsEnabled: c.IsEnabled,
		Capacity:  c.Capacity,
		Rate:      c.Rate,
	}
}

func applyTokenTransferFeeConfigIfNeeded(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput, remoteChainSelector uint64) ([]evm_contract.WriteOutput, error) {
	tokenTransferFeeConfigUpdates, err := makeTokenTransferFeeConfigUpdates(b, chain, input, remoteChainSelector)
	if err != nil {
		return nil, fmt.Errorf("failed to make token transfer fee config updates: %w", err)
	}
	if len(tokenTransferFeeConfigUpdates) == 0 {
		return nil, nil
	}
	report, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyTokenTransferFeeConfigUpdates, chain, evm_contract.FunctionInput[token_pool.ApplyTokenTransferFeeConfigUpdatesArgs]{
		ChainSelector: input.ChainSelector,
		Address:       input.TokenPoolAddress,
		Args: token_pool.ApplyTokenTransferFeeConfigUpdatesArgs{
			TokenTransferFeeConfigArgs:     tokenTransferFeeConfigUpdates,
			DisableTokenTransferFeeConfigs: nil,
		},
	})
	if err != nil {
		return nil, err
	}
	return []evm_contract.WriteOutput{report.Output}, nil
}

// v150OnRampConfigToTokenTransferFeeConfig converts OnRamp 1.5.0 token transfer fee config to tokens.TokenTransferFeeConfig.
func v150OnRampConfigToTokenTransferFeeConfig(cfg evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfig) *tokens.TokenTransferFeeConfig {
	return &tokens.TokenTransferFeeConfig{
		DestGasOverhead:               cfg.DestGasOverhead,
		DestBytesOverhead:             cfg.DestBytesOverhead,
		DefaultFinalityFeeUSDCents:    cfg.MinFeeUSDCents,
		CustomFinalityFeeUSDCents:     0,
		DefaultFinalityTransferFeeBps: cfg.DeciBps,
		CustomFinalityTransferFeeBps:  0,
		IsEnabled:                     cfg.IsEnabled,
	}
}

// v163FeeQuoterConfigToTokenTransferFeeConfig converts FeeQuoter 1.6.3 token transfer fee config to tokens.TokenTransferFeeConfig.
func v163FeeQuoterConfigToTokenTransferFeeConfig(cfg fqops_v163.TokenTransferFeeConfig) *tokens.TokenTransferFeeConfig {
	return &tokens.TokenTransferFeeConfig{
		DestGasOverhead:               cfg.DestGasOverhead,
		DestBytesOverhead:             cfg.DestBytesOverhead,
		DefaultFinalityFeeUSDCents:    cfg.MinFeeUSDCents,
		CustomFinalityFeeUSDCents:     0,
		DefaultFinalityTransferFeeBps: cfg.DeciBps,
		CustomFinalityTransferFeeBps:  0,
		IsEnabled:                     cfg.IsEnabled,
	}
}

// v2FeeQuoterConfigToTokenTransferFeeConfig converts FeeQuoter 2.0 (2.0.0 onRamp path) token transfer fee config to tokens.TokenTransferFeeConfig.
func v2FeeQuoterConfigToTokenTransferFeeConfig(cfg fqops.TokenTransferFeeConfig) *tokens.TokenTransferFeeConfig {
	return &tokens.TokenTransferFeeConfig{
		DestGasOverhead:               cfg.DestGasOverhead,
		DestBytesOverhead:             cfg.DestBytesOverhead,
		DefaultFinalityFeeUSDCents:    cfg.FeeUSDCents,
		CustomFinalityFeeUSDCents:     0,
		DefaultFinalityTransferFeeBps: 0,
		CustomFinalityTransferFeeBps:  0,
		IsEnabled:                     cfg.IsEnabled,
	}
}

func importTokenTransferFeeConfigFromActivePool(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput) (*tokens.TokenTransferFeeConfig, error) {
	// get router from token
	dCfgReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetDynamicConfig, chain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: input.ChainSelector,
		Address:       input.TokenPoolAddress,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get token dynamic config for pool %s on chain %s: %w", input.TokenPoolAddress.Hex(), chain.String(), err)
	}
	routerAddr := dCfgReport.Output.Router
	// if router is zero, then the pool is not active and we should not import any config
	if routerAddr == (common.Address{}) {
		return nil, nil
	}
	// get onRamp from router for the destination chain
	onRampOnRouterReport, err := cldf_ops.ExecuteOperation(b, routerops.GetOnRamp, chain, evm_contract.FunctionInput[uint64]{
		ChainSelector: input.ChainSelector,
		Address:       routerAddr,
		Args:          input.RemoteChainSelector,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get onRamp from router %s for "+
			"remote chain selector %d on chain %s: %w", routerAddr.Hex(), input.RemoteChainSelector, chain.String(), err)
	}
	if onRampOnRouterReport.Output == (common.Address{}) {
		// No onRamp configured for this lane yet, nothing to import.
		return nil, nil
	}
	onRampAddr := onRampOnRouterReport.Output
	// check the version of the onRamp contract
	onRampTAVReport, err := cldf_ops.ExecuteOperation(b, type_and_version.GetTypeAndVersion, chain, evm_contract.FunctionInput[struct{}]{
		ChainSelector: input.ChainSelector,
		Address:       onRampAddr,
		Args:          struct{}{},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get type and version of onRamp %s on chain %s: %w", onRampAddr.Hex(), chain.String(), err)
	}
	onRampTAV := onRampTAVReport.Output.Version
	switch onRampTAV.String() {
	case semver.MustParse("1.5.0").String():
		// for onRamp version 1.5.0, import tokenTransferFeeConfig from onRamp 1.5.0
		tokenTransferFeeConfigReport, err := cldf_ops.ExecuteOperation(b, onrampops_v150.OnRampGetTokenTransferFeeConfig, chain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       onRampAddr,
			Args:          input.TokenAddress,
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get token transfer fee config from onRamp %s for token %s on chain %s: %w",
				onRampAddr.Hex(), input.TokenAddress.Hex(), chain.String(), err)
		}
		return v150OnRampConfigToTokenTransferFeeConfig(tokenTransferFeeConfigReport.Output), nil
	case semver.MustParse("1.6.0").String(): // for onRamp versions 1.6.0 , import tokenTransferFeeConfig from FeeQuoter
		// get fee quoter from onRamp
		dCfgOnRamp, err := cldf_ops.ExecuteOperation(b, onrampops_v160.GetDynamicConfig, chain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       onRampAddr,
			Args:          struct{}{},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get dynamic config from onRamp %s on chain %s: %w", onRampAddr.Hex(), chain.String(), err)
		}
		feeQuoterAddr := dCfgOnRamp.Output.FeeQuoter
		if feeQuoterAddr == (common.Address{}) {
			return nil, nil
		}
		// get token transfer fee config from fee quoter
		tokenTransferFeeConfigReport, err := cldf_ops.ExecuteOperation(b, fqops_v163.GetTokenTransferFeeConfig, chain, evm_contract.FunctionInput[fqops_v163.GetTokenTransferFeeConfigArgs]{
			ChainSelector: input.ChainSelector,
			Address:       feeQuoterAddr,
			Args: fqops_v163.GetTokenTransferFeeConfigArgs{
				Token:             input.TokenAddress,
				DestChainSelector: input.RemoteChainSelector,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get token transfer fee config from fee quoter %s for token %s and remote chain selector %d on chain %s: %w",
				feeQuoterAddr.Hex(), input.TokenAddress.Hex(), input.RemoteChainSelector, chain.String(), err)
		}
		return v163FeeQuoterConfigToTokenTransferFeeConfig(tokenTransferFeeConfigReport.Output), nil
	case onrampops.Version.String():
		// get fee quoter from onRamp
		dCfgOnRamp, err := cldf_ops.ExecuteOperation(b, onrampops.GetDynamicConfig, chain, evm_contract.FunctionInput[struct{}]{
			ChainSelector: input.ChainSelector,
			Address:       onRampAddr,
			Args:          struct{}{},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get dynamic config from onRamp %s on chain %s: %w", onRampAddr.Hex(), chain.String(), err)
		}
		feeQuoterAddr := dCfgOnRamp.Output.FeeQuoter
		if feeQuoterAddr == (common.Address{}) {
			return nil, nil
		}
		// get token transfer fee config from fee quoter
		tokenTransferFeeConfigReport, err := cldf_ops.ExecuteOperation(b, fqops.GetTokenTransferFeeConfig, chain, evm_contract.FunctionInput[fqops.GetTokenTransferFeeConfigArgs]{
			ChainSelector: input.ChainSelector,
			Address:       feeQuoterAddr,
			Args: fqops.GetTokenTransferFeeConfigArgs{
				Token:             input.TokenAddress,
				DestChainSelector: input.RemoteChainSelector,
			},
		})
		if err != nil {
			return nil, fmt.Errorf("failed to get token transfer fee config from fee quoter %s for token %s and remote chain selector %d on chain %s: %w",
				feeQuoterAddr.Hex(), input.TokenAddress.Hex(), input.RemoteChainSelector, chain.String(), err)
		}
		return v2FeeQuoterConfigToTokenTransferFeeConfig(tokenTransferFeeConfigReport.Output), nil
	default:
		// Unsupported onRamp version, nothing to import.
		return nil, nil
	}
}

// mergeTokenTransferFeeConfig merges the imported token transfer fee config with the desired config,
// giving precedence to desired config values when they are non-zero (i.e. non-default). If imported config is nil, returns desired config.
func mergeTokenTransferFeeConfig(desired, imported *tokens.TokenTransferFeeConfig) *tokens.TokenTransferFeeConfig {
	if imported == nil {
		return desired
	}
	// merge imported config with desired config, giving precedence to desired config values when they are non-zero (i.e. non-default)
	merged := *imported
	if desired.DestGasOverhead != 0 {
		merged.DestGasOverhead = desired.DestGasOverhead
	}
	if desired.DestBytesOverhead != 0 {
		merged.DestBytesOverhead = desired.DestBytesOverhead
	}
	if desired.DefaultFinalityFeeUSDCents != 0 {
		merged.DefaultFinalityFeeUSDCents = desired.DefaultFinalityFeeUSDCents
	}
	if desired.CustomFinalityFeeUSDCents != 0 {
		merged.CustomFinalityFeeUSDCents = desired.CustomFinalityFeeUSDCents
	}
	if desired.DefaultFinalityTransferFeeBps != 0 {
		merged.DefaultFinalityTransferFeeBps = desired.DefaultFinalityTransferFeeBps
	}
	if desired.CustomFinalityTransferFeeBps != 0 {
		merged.CustomFinalityTransferFeeBps = desired.CustomFinalityTransferFeeBps
	}
	return &merged
}

func makeTokenTransferFeeConfigUpdates(b cldf_ops.Bundle, chain evm.Chain, input ConfigureTokenPoolForRemoteChainInput, remoteChainSelector uint64) ([]token_pool.TokenTransferFeeConfigArgs, error) {
	desiredTokenTransferFeeConfig := input.RemoteChainConfig.TokenTransferFeeConfig
	importedConfig, err := importTokenTransferFeeConfigFromActivePool(b, chain, input)
	if err != nil {
		return nil, fmt.Errorf("failed to import token transfer fee config from active pool: %w", err)
	}
	// merge imported config with desired config, giving precedence to desired config values when they are non-zero (i.e. non-default)
	desiredTokenTransferFeeConfig = *mergeTokenTransferFeeConfig(&desiredTokenTransferFeeConfig, importedConfig)
	if !desiredTokenTransferFeeConfig.IsEnabled {
		return nil, nil
	}
	report, err := cldf_ops.ExecuteOperation(b, token_pool.GetTokenTransferFeeConfig, chain, evm_contract.FunctionInput[token_pool.GetTokenTransferFeeConfigArgs]{
		ChainSelector: input.ChainSelector,
		Address:       input.TokenPoolAddress,
		Args: token_pool.GetTokenTransferFeeConfigArgs{
			Arg0:              common.Address{},
			DestChainSelector: remoteChainSelector,
			Arg2:              0,
			Arg3:              []byte{},
		},
	})
	if err != nil {
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

	updates := make([]token_pool.TokenTransferFeeConfigArgs, 0)

	if desiredTokenTransferFeeConfig.DestGasOverhead != currentTokenTransferFeeConfig.DestGasOverhead ||
		desiredTokenTransferFeeConfig.DestBytesOverhead != currentTokenTransferFeeConfig.DestBytesOverhead ||
		desiredTokenTransferFeeConfig.DefaultFinalityFeeUSDCents != currentTokenTransferFeeConfig.DefaultBlockConfirmationsFeeUSDCents ||
		desiredTokenTransferFeeConfig.CustomFinalityFeeUSDCents != currentTokenTransferFeeConfig.CustomBlockConfirmationsFeeUSDCents ||
		desiredTokenTransferFeeConfig.DefaultFinalityTransferFeeBps != currentTokenTransferFeeConfig.DefaultBlockConfirmationsTransferFeeBps ||
		desiredTokenTransferFeeConfig.CustomFinalityTransferFeeBps != currentTokenTransferFeeConfig.CustomBlockConfirmationsTransferFeeBps {
		updates = append(updates, token_pool.TokenTransferFeeConfigArgs{
			DestChainSelector: remoteChainSelector,
			TokenTransferFeeConfig: token_pool.TokenTransferFeeConfig{
				DestGasOverhead:                         desiredTokenTransferFeeConfig.DestGasOverhead,
				DestBytesOverhead:                       desiredTokenTransferFeeConfig.DestBytesOverhead,
				DefaultBlockConfirmationsFeeUSDCents:    desiredTokenTransferFeeConfig.DefaultFinalityFeeUSDCents,
				CustomBlockConfirmationsFeeUSDCents:     desiredTokenTransferFeeConfig.CustomFinalityFeeUSDCents,
				DefaultBlockConfirmationsTransferFeeBps: desiredTokenTransferFeeConfig.DefaultFinalityTransferFeeBps,
				CustomBlockConfirmationsTransferFeeBps:  desiredTokenTransferFeeConfig.CustomFinalityTransferFeeBps,
				IsEnabled:                               true,
			},
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
