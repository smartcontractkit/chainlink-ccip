package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	v1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/token_pool"
	tp_bindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

var ConfigureTokenForTransfers = cldf_ops.NewSequence(
	"configure-token",
	semver.MustParse("1.7.0"),
	"Configures a token on an EVM chain for usage with CCIP",
	func(b operations.Bundle, chains chain.BlockChains, input tokens.ConfigureTokenForTransfersInput) (output sequences.OnChainOutput, err error) {
		ops := make([]mcms_types.BatchOperation, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}
		tokenPoolAddress := common.HexToAddress(input.TokenPoolAddress)

		// Fetch the local token address once to reuse for token transfer fee lookups and registration.
		tokenAddressReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetToken, chain, evm_contract.FunctionInput[any]{
			ChainSelector: input.ChainSelector,
			Address:       tokenPoolAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
		}
		localTokenAddress := tokenAddressReport.Output

		// Configure remote chains on the token pool as specified
		// This means adding any remote chains not already added, removing any remote chains that are no longer desired,
		// or modifying rate limiters on any existing remote chains.
		customBlockConfirmationArgs := make([]token_pool.CustomBlockConfirmationRateLimitConfigArg, 0, len(input.RemoteChains))
		tokenTransferFeeConfigArgs := make([]token_pool.TokenTransferFeeConfigArg, 0, len(input.RemoteChains))
		destToUseDefaultFeeConfigs := make([]uint64, 0)
		for destChainSelector, remoteChainConfig := range input.RemoteChains {
			configureTokenPoolForRemoteChainReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPoolForRemoteChain, chain, ConfigureTokenPoolForRemoteChainInput{
				ChainSelector:       input.ChainSelector,
				TokenPoolAddress:    tokenPoolAddress,
				RemoteChainSelector: destChainSelector,
				RemoteChainConfig:   remoteChainConfig,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool with address %s on %s for remote chain with selector %d: %w", input.TokenPoolAddress, chain, destChainSelector, err)
			}
			ops = append(ops, configureTokenPoolForRemoteChainReport.Output.BatchOps...)

			currentFeeConfigReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetTokenTransferFeeConfig, chain, evm_contract.FunctionInput[token_pool.GetTokenTransferFeeConfigArgs]{
				ChainSelector: input.ChainSelector,
				Address:       tokenPoolAddress,
				Args: token_pool.GetTokenTransferFeeConfigArgs{
					LocalToken:        localTokenAddress,
					DestChainSelector: destChainSelector,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token transfer fee config for token pool %s on %s for remote chain %d: %w", input.TokenPoolAddress, chain, destChainSelector, err)
			}
			currentFeeConfig := currentFeeConfigReport.Output
			if remoteChainConfig.TokenTransferFeeConfig == nil {
				if !tokenTransferFeeConfigIsZero(currentFeeConfig) {
					destToUseDefaultFeeConfigs = append(destToUseDefaultFeeConfigs, destChainSelector)
					tokenTransferFeeConfigArgs = append(tokenTransferFeeConfigArgs, token_pool.TokenTransferFeeConfigArg{
						DestChainSelector:      destChainSelector,
						TokenTransferFeeConfig: tokens.TokenTransferFeeConfig{},
					})
				}
			} else if !tokenTransferFeeConfigsEqual(currentFeeConfig, *remoteChainConfig.TokenTransferFeeConfig) {
				tokenTransferFeeConfigArgs = append(tokenTransferFeeConfigArgs, token_pool.TokenTransferFeeConfigArg{
					DestChainSelector:      destChainSelector,
					TokenTransferFeeConfig: *remoteChainConfig.TokenTransferFeeConfig,
				})
			}

			if remoteChainConfig.CustomBlockConfirmationConfig != nil {
				cfg := remoteChainConfig.CustomBlockConfirmationConfig
				customBlockConfirmationArgs = append(customBlockConfirmationArgs, token_pool.CustomBlockConfirmationRateLimitConfigArg{
					RemoteChainSelector:       destChainSelector,
					OutboundRateLimiterConfig: cfg.Outbound,
					InboundRateLimiterConfig:  cfg.Inbound,
				})
			}
		}

		tokenTransferFeeWrites := make([]evm_contract.WriteOutput, 0, 1)
		if len(tokenTransferFeeConfigArgs) > 0 || len(destToUseDefaultFeeConfigs) > 0 {
			applyTokenTransferFeesReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyTokenTransferFeeConfigUpdates, chain, evm_contract.FunctionInput[token_pool.ApplyTokenTransferFeeConfigUpdatesArgs]{
				ChainSelector: input.ChainSelector,
				Address:       tokenPoolAddress,
				Args: token_pool.ApplyTokenTransferFeeConfigUpdatesArgs{
					TokenTransferFeeConfigArgs: tokenTransferFeeConfigArgs,
					DestToUseDefaultFeeConfigs: destToUseDefaultFeeConfigs,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply token transfer fee config updates on token pool %s: %w", input.TokenPoolAddress, err)
			}
			tokenTransferFeeWrites = append(tokenTransferFeeWrites, applyTokenTransferFeesReport.Output)
		}
		if len(tokenTransferFeeWrites) > 0 {
			tokenTransferFeeBatch, err := evm_contract.NewBatchOperationFromWrites(tokenTransferFeeWrites)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation for token transfer fee configuration: %w", err)
			}
			ops = append(ops, tokenTransferFeeBatch)
		}

		blockConfirmationWrites := make([]evm_contract.WriteOutput, 0, 1)
		if input.CustomBlockConfirmationConfig != nil {
			applyConfigReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyCustomBlockConfirmationConfigUpdates, chain, evm_contract.FunctionInput[token_pool.ApplyCustomBlockConfirmationConfigArgs]{
				ChainSelector: input.ChainSelector,
				Address:       tokenPoolAddress,
				Args: token_pool.ApplyCustomBlockConfirmationConfigArgs{
					MinBlockConfirmation: input.CustomBlockConfirmationConfig.MinBlockConfirmation,
					RateLimitConfigArgs:  customBlockConfirmationArgs,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to apply custom block confirmation config updates on token pool %s: %w", input.TokenPoolAddress, err)
			}
			blockConfirmationWrites = append(blockConfirmationWrites, applyConfigReport.Output)
		} else if len(customBlockConfirmationArgs) > 0 {
			setConfigReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetCustomBlockConfirmationRateLimitConfig, chain, evm_contract.FunctionInput[[]token_pool.CustomBlockConfirmationRateLimitConfigArg]{
				ChainSelector: input.ChainSelector,
				Address:       tokenPoolAddress,
				Args:          customBlockConfirmationArgs,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set custom block confirmation rate limits on token pool %s: %w", input.TokenPoolAddress, err)
			}
			blockConfirmationWrites = append(blockConfirmationWrites, setConfigReport.Output)
		}
		if len(blockConfirmationWrites) > 0 {
			blockConfirmationBatch, err := evm_contract.NewBatchOperationFromWrites(blockConfirmationWrites)
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation for custom block confirmation configuration: %w", err)
			}
			ops = append(ops, blockConfirmationBatch)
		}

		// Register the token with the token admin registry
		registerTokenReport, err := cldf_ops.ExecuteSequence(b, v1_5_0.RegisterToken, chain, v1_5_0.RegisterTokenInput{
			ChainSelector:             input.ChainSelector,
			TokenAddress:              localTokenAddress,
			TokenPoolAddress:          tokenPoolAddress,
			ExternalAdmin:             common.HexToAddress(input.ExternalAdmin),
			TokenAdminRegistryAddress: common.HexToAddress(input.RegistryAddress),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to register token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
		}
		ops = append(ops, registerTokenReport.Output.BatchOps...)

		return sequences.OnChainOutput{
			BatchOps: ops,
		}, nil
	},
)

func tokenTransferFeeConfigIsZero(config tp_bindings.IPoolV2TokenTransferFeeConfig) bool {
	return config.DestGasOverhead == 0 &&
		config.DestBytesOverhead == 0 &&
		config.DefaultBlockConfirmationFeeUSDCents == 0 &&
		config.CustomBlockConfirmationFeeUSDCents == 0 &&
		config.DefaultBlockConfirmationTransferFeeBps == 0 &&
		config.CustomBlockConfirmationTransferFeeBps == 0
}

func tokenTransferFeeConfigsEqual(current tp_bindings.IPoolV2TokenTransferFeeConfig, desired tokens.TokenTransferFeeConfig) bool {
	return current.DestGasOverhead == desired.DestGasOverhead &&
		current.DestBytesOverhead == desired.DestBytesOverhead &&
		current.DefaultBlockConfirmationFeeUSDCents == desired.DefaultBlockConfirmationFeeUSDCents &&
		current.CustomBlockConfirmationFeeUSDCents == desired.CustomBlockConfirmationFeeUSDCents &&
		current.DefaultBlockConfirmationTransferFeeBps == desired.DefaultBlockConfirmationTransferFeeBps &&
		current.CustomBlockConfirmationTransferFeeBps == desired.CustomBlockConfirmationTransferFeeBps
}
