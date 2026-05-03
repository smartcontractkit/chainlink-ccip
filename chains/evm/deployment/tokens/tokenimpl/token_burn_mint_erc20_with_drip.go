package tokenimpl

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	drip_v150 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/burn_mint_erc20_with_drip"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type tokenBurnMintERC20WithDrip struct{}

func (tokenBurnMintERC20WithDrip) ContractType() deployment.ContractType {
	return drip_v150.ContractType
}

func (tokenBurnMintERC20WithDrip) Capabilities() CapabilitySet {
	return CapabilitySet{
		ParticipatesInPoolRoleGrant: true,
		SupportsAdminRole:           true,
		SupportsCCIPAdmin:           true,
		SupportsPreMint:             false,
	}
}

func (tokenBurnMintERC20WithDrip) RevokeAdminRole(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address) ([]contract.WriteOutput, error) {
	return revokeBnMDefaultAdminRole(b, chain, token, externalAdmin)
}

func (tokenBurnMintERC20WithDrip) GrantAdminRole(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address) ([]contract.WriteOutput, error) {
	return grantBnMDefaultAdminRole(b, chain, token, externalAdmin)
}

func (tokenBurnMintERC20WithDrip) GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address) ([]contract.WriteOutput, error) {
	return grantBnMMintAndBurnRoles(b, chain, token, pool)
}

func (tokenBurnMintERC20WithDrip) SetCCIPAdmin(b cldf_ops.Bundle, chain evm.Chain, token, ccipAdmin common.Address) ([]contract.WriteOutput, error) {
	return setBnMCCIPAdmin(b, chain, token, ccipAdmin)
}

func (tokenBurnMintERC20WithDrip) Transfer(b cldf_ops.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	// NOTE: BnM ERC20 drip tokens inherit from a standard ERC20 implementation, so we can use the same transfer helper function as the plain ERC20 strategy.
	return transferTokensERC20(b, chain, token, to, scaledAmount)
}

func (tokenBurnMintERC20WithDrip) Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	ref, err := contract.MaybeDeployContract(b, drip_v150.Deploy, chain,
		contract.DeployInput[drip_v150.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(drip_v150.ContractType, *drip_v150.Version),
			ChainSelector:  chain.Selector,
			Qualifier:      &in.Symbol,
			Args: drip_v150.ConstructorArgs{
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
