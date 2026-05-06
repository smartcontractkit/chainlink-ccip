package tokenimpl

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type tokenBurnMintERC20 struct{}

func (tokenBurnMintERC20) ContractType() deployment.ContractType {
	return burn_mint_erc20.ContractType
}

func (tokenBurnMintERC20) Capabilities() CapabilitySet {
	return CapabilitySet{
		ParticipatesInPoolRoleGrant: true,
		SupportsAdminRole:           true,
		SupportsCCIPAdmin:           true,
		SupportsPreMint:             true,
	}
}

func (tokenBurnMintERC20) RevokeAdminRole(b operations.Bundle, chain evm.Chain, token, user common.Address) ([]contract.WriteOutput, error) {
	return revokeDefaultAdminRoleBurnMintERC20(b, chain, token, user)
}

func (tokenBurnMintERC20) GrantAdminRole(b operations.Bundle, chain evm.Chain, token, externalAdmin common.Address) ([]contract.WriteOutput, error) {
	return grantDefaultAdminRoleBurnMintERC20(b, chain, token, externalAdmin)
}

func (tokenBurnMintERC20) GrantPoolRoles(b operations.Bundle, chain evm.Chain, token, pool, _ common.Address) ([]contract.WriteOutput, error) {
	return grantMintAndBurnRolesBurnMintERC20(b, chain, token, pool)
}

func (tokenBurnMintERC20) SetCCIPAdmin(b operations.Bundle, chain evm.Chain, token, ccipAdmin common.Address) ([]contract.WriteOutput, error) {
	return setCCIPAdminBurnMintERC20(b, chain, token, ccipAdmin)
}

func (tokenBurnMintERC20) Transfer(b operations.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	// NOTE: BnM ERC20 tokens inherit from a standard ERC20 implementation, so we can use the same transfer helper function as the plain ERC20 token.
	return transferTokensERC20(b, chain, token, to, scaledAmount)
}

func (tokenBurnMintERC20) Deploy(b operations.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	maxSupply := big.NewInt(0)
	if in.Supply != nil {
		maxSupply = tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*in.Supply), in.Decimals)
	}

	preMint := big.NewInt(0)
	if in.PreMint != nil {
		preMint = tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*in.PreMint), in.Decimals)
	}

	ref, err := contract.MaybeDeployContract(b, burn_mint_erc20.Deploy, chain,
		contract.DeployInput[burn_mint_erc20.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20.ContractType, *utils.Version_1_0_0),
			ChainSelector:  chain.Selector,
			Qualifier:      &in.Symbol,
			Args: burn_mint_erc20.ConstructorArgs{
				Name:      in.Name,
				Symbol:    in.Symbol,
				Decimals:  in.Decimals,
				MaxSupply: maxSupply,
				PreMint:   preMint,
			},
		},
		nil,
	)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy BurnMintERC20 token: %w", err)
	}

	return ref, nil, nil
}
