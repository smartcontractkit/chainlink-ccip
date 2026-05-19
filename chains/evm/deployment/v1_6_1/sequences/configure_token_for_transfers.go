package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	v1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	tpops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	tpbind "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	ops2contract "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

var ConfigureTokenForTransfers = cldf_ops.NewSequence(
	"configure-token-for-transfers",
	semver.MustParse("1.6.1"),
	"Configures a token on an EVM chain for usage with CCIP",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input tokens.ConfigureTokenForTransfersInput) (output sequences.OnChainOutput, err error) {
		ops := make([]mcms_types.BatchOperation, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		tokenPoolAddress := common.HexToAddress(input.TokenPoolAddress)
		tp, err := tpbind.NewTokenPool(tokenPoolAddress, chain.Client)
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to instantiate token pool contract: %w", err)
		}

		var tokenAddress common.Address
		if input.TokenAddress != "" {
			tokenAddress = common.HexToAddress(input.TokenAddress)
		} else {
			tokenAddrReport, err := cldf_ops.ExecuteOperation(b, tpops.NewReadGetToken(tp), chain, ops2contract.FunctionInput[struct{}]{
				Args: struct{}{},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
			}
			tokenAddress = tokenAddrReport.Output
		}
		registryTokenPoolAddress := tokenPoolAddress
		if input.RegistryTokenPoolAddress != "" {
			registryTokenPoolAddress = common.HexToAddress(input.RegistryTokenPoolAddress)
		}

		// Validate the pool supports the token
		isSupported, err := cldf_ops.ExecuteOperation(b, tpops.NewReadIsSupportedToken(tp), chain, ops2contract.FunctionInput[common.Address]{
			Args: tokenAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check if token %s is supported by token pool %s on %s: %w", tokenAddress, tokenPoolAddress, chain, err)
		}
		if !isSupported.Output {
			return sequences.OnChainOutput{}, fmt.Errorf("token %s is not supported by token pool %s", tokenAddress, tokenPoolAddress)
		}

		// Configure token pool for each remote chain
		for remoteChainSelector, remoteChainConfig := range input.RemoteChains {
			configureTokenPoolForRemoteChainReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPoolForRemoteChain, chain, ConfigureTokenPoolForRemoteChainInput{
				ChainSelector:       input.ChainSelector,
				TokenPoolAddress:    tokenPoolAddress,
				RemoteChainSelector: remoteChainSelector,
				RemoteChainConfig:   remoteChainConfig,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure token pool for remote chain %d: %w", remoteChainSelector, err)
			}
			ops = append(ops, configureTokenPoolForRemoteChainReport.Output.BatchOps...)
		}

		// Register the token with the token admin registry
		registerTokenReport, err := cldf_ops.ExecuteSequence(b, v1_5_0.RegisterToken, chain, v1_5_0.RegisterTokenInput{
			ChainSelector:             input.ChainSelector,
			TokenAddress:              tokenAddress,
			TokenPoolAddress:          registryTokenPoolAddress,
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
