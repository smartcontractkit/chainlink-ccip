package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/siloed_usdc_token_pool"

	usdc_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type SiloedUSDCTokenPoolDeploySequenceInput struct {
	ChainSelector uint64
	Token         common.Address
	Allowlist     []common.Address
	RMNProxy      common.Address
	Router        common.Address
	LockBox       common.Address
	MCMSAddress   common.Address
}

var SiloedUSDCTokenPoolDeploySequence = operations.NewSequence(
	"SiloedUSDCTokenPoolDeploySequence",
	semver.MustParse("1.6.4"),
	"Deploys the SiloedUSDCTokenPool contract",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input SiloedUSDCTokenPoolDeploySequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
		}

		report, err := operations.ExecuteOperation(b, siloed_usdc_token_pool.Deploy, chain, contract.DeployInput[siloed_usdc_token_pool.ConstructorArgs]{
			ChainSelector: input.ChainSelector,
			TypeAndVersion: deployment.NewTypeAndVersion(
				siloed_usdc_token_pool.ContractType,
				*siloed_usdc_token_pool.Version,
			),
			Args: siloed_usdc_token_pool.ConstructorArgs{
				Token:              input.Token,
				LocalTokenDecimals: 6, // This pool is for USDC which always has 6 decimals
				Allowlist:          input.Allowlist,
				RMNProxy:           input.RMNProxy,
				Router:             input.Router,
				LockBox:            input.LockBox,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to deploy SiloedUSDCTokenPool on %s: %w", chain, err)
		}

		// Begin transferring ownership of the newly deployed token pool to MCMS
		_, err = operations.ExecuteOperation(b, usdc_token_pool_ops.USDCTokenPoolTransferOwnership, chain, contract.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       common.HexToAddress(report.Output.Address),
			Args:          input.MCMSAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to transfer ownership of the token pool to MCMS on %s: %w", chain, err)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{report.Output},
		}, nil
	},
)
