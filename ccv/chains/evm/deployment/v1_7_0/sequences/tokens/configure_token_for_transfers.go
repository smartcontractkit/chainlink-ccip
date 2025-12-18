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
	"configure-token",
	semver.MustParse("1.7.0"),
	"Configures a token on an EVM chain for usage with CCIP",
	func(b cldf_ops.Bundle, chains chain.BlockChains, input tokens.ConfigureTokenForTransfersInput) (output sequences.OnChainOutput, err error) {
		ops := make([]mcms_types.BatchOperation, 0)
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not found", input.ChainSelector)
		}

		tokenAddress, err := cldf_ops.ExecuteOperation(b, token_pool.GetToken, chain, evm_contract.FunctionInput[any]{
			ChainSelector: input.ChainSelector,
			Address:       common.HexToAddress(input.TokenPoolAddress),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
		}

		// Configure minimum block confirmation
		configureMinBlockConfirmationReport, err := cldf_ops.ExecuteOperation(b, token_pool.SetMinBlockConfirmation, chain, evm_contract.FunctionInput[uint16]{
			ChainSelector: input.ChainSelector,
			Address:       common.HexToAddress(input.TokenPoolAddress),
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

		// Get the advanced pool hooks address
		advancedPoolHooksAddress, err := cldf_ops.ExecuteOperation(b, token_pool.GetAdvancedPoolHooks, chain, evm_contract.FunctionInput[any]{
			ChainSelector: input.ChainSelector,
			Address:       common.HexToAddress(input.TokenPoolAddress),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get advanced pool hooks address from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
		}

		// Configure token pool for each remote chain
		for remoteChainSelector, remoteChainConfig := range input.RemoteChains {
			configureTokenPoolForRemoteChainReport, err := cldf_ops.ExecuteSequence(b, ConfigureTokenPoolForRemoteChain, chain, ConfigureTokenPoolForRemoteChainInput{
				ChainSelector:       input.ChainSelector,
				TokenPoolAddress:    common.HexToAddress(input.TokenPoolAddress),
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
