package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	usdc_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"
)

type DeployUSDCTokenPoolProxySequenceInput struct {
	ChainSelector uint64
	Token         common.Address
	PoolAddresses usdc_token_pool_proxy.PoolAddresses
	Router        common.Address
	MCMSAddress   common.Address
}

var DeployUSDCTokenPoolProxySequence = operations.NewSequence(
	"DeployUSDCTokenPoolProxySequence",
	semver.MustParse("1.6.4"),
	"Deploys the USDCTokenPoolProxy contract",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input DeployUSDCTokenPoolProxySequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
		}

		report, err := operations.ExecuteOperation(b, usdc_token_pool_proxy.Deploy, chain, contract.DeployInput[usdc_token_pool_proxy.ConstructorArgs]{
			ChainSelector: input.ChainSelector,
			TypeAndVersion: deployment.NewTypeAndVersion(
				usdc_token_pool_proxy.ContractType,
				*usdc_token_pool_proxy.Version,
			),
			// TODO: Review and make dynamic?
			Args: usdc_token_pool_proxy.ConstructorArgs{
				Token:         input.Token,
				PoolAddresses: input.PoolAddresses,
				Router:        input.Router,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy USDCTokenPoolProxy on %s: %w", chain, err)
		}

		// Begin transferring ownership to MCMS. A separate changeset will be used to accept ownership.
		_, err = operations.ExecuteOperation(b, usdc_token_pool_ops.USDCTokenPoolTransferOwnership, chain, contract.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       common.HexToAddress(report.Output.Address),
			Args:          input.MCMSAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership to MCMS on %s: %w", chain, err)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{report.Output},
		}, nil
	},
)
