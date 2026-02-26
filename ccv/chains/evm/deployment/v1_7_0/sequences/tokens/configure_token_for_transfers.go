package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v1_7_0/operations/token_pool"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	v1_5_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

var ConfigureTokenForTransfers = cldf_ops.NewSequence(
	"configure-token-for-transfers",
	semver.MustParse("1.7.0"),
	"Configures a token on an EVM chain for usage with CCIP",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input tokens.ConfigureTokenForTransfersInput) (output sequences.OnChainOutput, err error) {
		ops := make([]mcms_types.BatchOperation, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		var tokenAddress common.Address
		if input.TokenAddress != "" {
			tokenAddress = common.HexToAddress(input.TokenAddress)
		} else {
			tokenAddrReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetToken, chain, evm_contract.FunctionInput[any]{
				ChainSelector: input.ChainSelector,
				Address:       common.HexToAddress(input.TokenPoolAddress),
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
			}
			tokenAddress = tokenAddrReport.Output
		}
		tokenPoolAddress := common.HexToAddress(input.TokenPoolAddress)
		registryTokenPoolAddress := tokenPoolAddress
		if input.RegistryTokenPoolAddress != "" {
			registryTokenPoolAddress = common.HexToAddress(input.RegistryTokenPoolAddress)
		}

		// Validate the pool supports the token
		isSupported, err := cldf_ops.ExecuteOperation(b, token_pool.IsSupportedToken, chain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       tokenPoolAddress,
			Args:          tokenAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to check if token %s is supported by token pool %s on %s: %w", tokenAddress, tokenPoolAddress, chain, err)
		}
		if !isSupported.Output {
			return sequences.OnChainOutput{}, fmt.Errorf("token %s is not supported by token pool %s", tokenAddress, tokenPoolAddress)
		}

		// Configure minimum block confirmation (skip if already set to desired value).
		currentMinBlockReport, err := cldf_ops.ExecuteOperation(b, token_pool.GetMinBlockConfirmation, chain, evm_contract.FunctionInput[any]{
			ChainSelector: input.ChainSelector,
			Address:       tokenPoolAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get min block confirmation from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
		}
		if currentMinBlockReport.Output != input.MinFinalityValue {
			configureMinBlockConfirmationReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetMinBlockConfirmations, chain, evm_contract.FunctionInput[uint16]{
				ChainSelector: input.ChainSelector,
				Address:       tokenPoolAddress,
				Args:          input.MinFinalityValue,
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to configure minimum block confirmation for token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
			}
			configureMinBlockConfirmationOps, err := evm_contract.NewBatchOperationFromWrites([]evm_contract.WriteOutput{configureMinBlockConfirmationReport.Output})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to create batch operation from write outputs: %w", err)
			}
			ops = append(ops, configureMinBlockConfirmationOps)
		}

		// Get the advanced pool hooks address
		advancedPoolHooksAddress, err := cldf_ops.ExecuteOperation(b, token_pool.GetAdvancedPoolHooks, chain, evm_contract.FunctionInput[any]{
			ChainSelector: input.ChainSelector,
			Address:       tokenPoolAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get advanced pool hooks address from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
		}

		// Configure token pool for each remote chain
		for remoteChainSelector, remoteChainConfig := range input.RemoteChains {
			configureTokenPoolForRemoteChainReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPoolForRemoteChain, chain, ConfigureTokenPoolForRemoteChainInput{
				ChainSelector:       input.ChainSelector,
				TokenPoolAddress:    tokenPoolAddress,
				AdvancedPoolHooks:   advancedPoolHooksAddress.Output,
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
