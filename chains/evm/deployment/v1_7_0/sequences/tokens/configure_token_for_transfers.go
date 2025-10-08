package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
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
		// TODO: Change to ConfigureTokenPoolForRemoteChains (plural)?
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
		}

		tokenAddress, err := cldf_ops.ExecuteOperation(b, token_pool.GetToken, chain, evm_contract.FunctionInput[any]{
			ChainSelector: input.ChainSelector,
			Address:       common.HexToAddress(input.TokenPoolAddress),
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token address from token pool with address %s on %s: %w", input.TokenPoolAddress, chain, err)
		}

		// Register the token with the token admin registry
		registerTokenReport, err := cldf_ops.ExecuteSequence(b, RegisterToken, chain, RegisterTokenInput{
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
