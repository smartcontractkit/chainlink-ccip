package registrations

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/strategy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20_with_drip"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func init() {
	strategy.GetRegistry().RegisterEVM(burnMintERC20WithDripStrategy{})
}

type burnMintERC20WithDripStrategy struct{}

func (burnMintERC20WithDripStrategy) ContractType() deployment.ContractType {
	return burn_mint_erc20_with_drip.ContractType
}

func (burnMintERC20WithDripStrategy) Capabilities() strategy.Capabilities {
	return strategy.Capabilities{
		SupportsAdminRole:           true,
		SupportsCCIPAdmin:           true,
		SupportsPreMint:             true,
		ParticipatesInPoolRoleGrant: true,
	}
}

func (burnMintERC20WithDripStrategy) Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	qualifier := in.Symbol
	maxSupply, preMint := scaledSupplyAndPreMint(in)
	ref, err := contract.MaybeDeployContract(b, burn_mint_erc20_with_drip.Deploy, chain, contract.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
		TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_with_drip.ContractType, *common_utils.Version_1_0_0),
		ChainSelector:  chain.Selector,
		Args: burn_mint_erc20_with_drip.ConstructorArgs{
			Name:      in.Name,
			Symbol:    in.Symbol,
			Decimals:  in.Decimals,
			MaxSupply: maxSupply,
			PreMint:   preMint,
		},
		Qualifier: &qualifier,
	}, nil)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy BurnMintERC20WithDrip token: %w", err)
	}
	return ref, nil, nil
}

func (burnMintERC20WithDripStrategy) GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	return grantBnMMintAndBurnRoles(b, chain, token, pool, chainSelector)
}

func (burnMintERC20WithDripStrategy) GrantExternalAdmin(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	return grantBnMDefaultAdminRole(b, chain, token, externalAdmin, chainSelector)
}
