package tokenimpl

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type tokenBurnMintERC20WithDripV1_5_0 struct{}

func (tokenBurnMintERC20WithDripV1_5_0) ContractType() deployment.ContractType {
	return burn_mint_erc20_with_drip.ContractType
}

func (tokenBurnMintERC20WithDripV1_5_0) Capabilities() CapabilitySet {
	return CapabilitySet{
		ParticipatesInPoolRoleGrant: true,
		SupportsAdminRole:           true,
		SupportsCCIPAdmin:           true,
		SupportsPreMint:             false,
	}
}

func (tokenBurnMintERC20WithDripV1_5_0) RevokeAdminRole(b operations.Bundle, chain evm.Chain, token, externalAdmin common.Address) ([]contract.WriteOutput, error) {
	return revokeDefaultAdminRoleBurnMintERC20(b, chain, token, externalAdmin)
}

func (tokenBurnMintERC20WithDripV1_5_0) GrantAdminRole(b operations.Bundle, chain evm.Chain, token, externalAdmin common.Address) ([]contract.WriteOutput, error) {
	return grantDefaultAdminRoleBurnMintERC20(b, chain, token, externalAdmin)
}

func (tokenBurnMintERC20WithDripV1_5_0) GrantPoolRoles(b operations.Bundle, chain evm.Chain, token, pool, _ common.Address) ([]contract.WriteOutput, error) {
	return grantMintAndBurnRolesBurnMintERC20(b, chain, token, pool)
}

func (tokenBurnMintERC20WithDripV1_5_0) SetCCIPAdmin(b operations.Bundle, chain evm.Chain, token, ccipAdmin common.Address) ([]contract.WriteOutput, error) {
	return setCCIPAdminBurnMintERC20(b, chain, token, ccipAdmin)
}

func (tokenBurnMintERC20WithDripV1_5_0) Transfer(b operations.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	// NOTE: BnM ERC20 drip tokens inherit from a standard ERC20 implementation, so we can use the same transfer helper function as the plain ERC20 token.
	return transferTokensERC20(b, chain, token, to, scaledAmount)
}

func (tokenBurnMintERC20WithDripV1_5_0) Deploy(b operations.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	ref, err := contract.MaybeDeployContract(b, burn_mint_erc20_with_drip.Deploy, chain,
		contract.DeployInput[burn_mint_erc20_with_drip.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc20_with_drip.ContractType, *burn_mint_erc20_with_drip.Version),
			ChainSelector:  chain.Selector,
			Qualifier:      &in.Symbol,
			Args: burn_mint_erc20_with_drip.ConstructorArgs{
				Name:   in.Name,
				Symbol: in.Symbol,
			},
		},
		nil,
	)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy BurnMintERC20WithDrip (v1.5.0) token: %w", err)
	}

	return ref, nil, nil
}
