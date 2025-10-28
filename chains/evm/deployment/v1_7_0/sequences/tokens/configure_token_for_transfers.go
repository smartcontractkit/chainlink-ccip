package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	v1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_7_0/operations/token_pool"
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

		// Configure remote chains on the token pool as specified
		// This means adding any remote chains not already added, removing any remote chains that are no longer desired,
		// or modifying rate limiters on any existing remote chains.
		customBlockConfirmationArgs := make([]token_pool.CustomBlockConfirmationRateLimitConfigArg, 0, len(input.RemoteChains))
		for destChainSelector, remoteChainConfig := range input.RemoteChains {
			configureTokenPoolForRemoteChainReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPoolForRemoteChain, chain, ConfigureTokenPoolForRemoteChainInput{
				ChainSelector:       input.ChainSelector,
				TokenPoolAddress:    common.HexToAddress(input.TokenPoolAddress),
				RemoteChainSelector: destChainSelector,
				RemoteChainConfig:   remoteChainConfig,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool with address %s on %s for remote chain with selector %d: %w", input.TokenPoolAddress, chain, destChainSelector, err)
			}
			ops = append(ops, configureTokenPoolForRemoteChainReport.Output.BatchOps...)

			if remoteChainConfig.CustomBlockConfirmationConfig != nil {
				cfg := remoteChainConfig.CustomBlockConfirmationConfig
				customBlockConfirmationArgs = append(customBlockConfirmationArgs, token_pool.CustomBlockConfirmationRateLimitConfigArg{
					RemoteChainSelector:       destChainSelector,
					OutboundRateLimiterConfig: cfg.Outbound,
					InboundRateLimiterConfig:  cfg.Inbound,
				})
			}
		}

		blockConfirmationWrites := make([]evm_contract.WriteOutput, 0, 1)
		if input.CustomBlockConfirmationConfig != nil {
			applyConfigReport, err := cldf_ops.ExecuteOperation(b, token_pool.ApplyCustomBlockConfirmationConfigUpdates, chain, evm_contract.FunctionInput[token_pool.ApplyCustomBlockConfirmationConfigArgs]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(input.TokenPoolAddress),
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
				Address:       common.HexToAddress(input.TokenPoolAddress),
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

		tokenAddress, err := cldf_ops.ExecuteOperation(b, token_pool.GetToken, chain, evm_contract.FunctionInput[any]{
			ChainSelector: input.ChainSelector,
			Address:       common.HexToAddress(input.TokenPoolAddress),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
		}

		// Register the token with the token admin registry
		registerTokenReport, err := cldf_ops.ExecuteSequence(b, v1_5_0.RegisterToken, chain, v1_5_0.RegisterTokenInput{
			ChainSelector:             input.ChainSelector,
			TokenAddress:              tokenAddress.Output,
			TokenPoolAddress:          common.HexToAddress(input.TokenPoolAddress),
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
