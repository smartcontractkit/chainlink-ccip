package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
	deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/erc20_lock_box"

	datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type ERC20LockboxDeploySequenceInput struct {
	ChainSelector      uint64
	TokenAdminRegistry common.Address
}

var ERC20LockboxDeploySequence = operations.NewSequence(
	"ERC20LockboxDeploySequence",
	semver.MustParse("1.6.4"),
	"Deploys the ERC20Lockbox contract",
	func(b operations.Bundle, chains cldf_chain.BlockChains, input ERC20LockboxDeploySequenceInput) (sequences.OnChainOutput, error) {
		chain, ok := chains.EVMChains()[input.ChainSelector]
		if !ok {
			return sequences.OnChainOutput{}, fmt.Errorf("chain with selector %d not defined", input.ChainSelector)
		}

		report, err := operations.ExecuteOperation(b, erc20_lock_box.Deploy, chain, contract.DeployInput[erc20_lock_box.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(
				erc20_lock_box.ContractType,
				*erc20_lock_box.Version,
			),
			ChainSelector: input.ChainSelector,
			Args: erc20_lock_box.ConstructorArgs{
				TokenAdminRegistry: input.TokenAdminRegistry,
			},
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to execute ERC20LockboxDeploy on %s: %w", chain, err)
		}

		return sequences.OnChainOutput{
			Addresses: []datastore.AddressRef{report.Output},
		}, nil

	},
)
