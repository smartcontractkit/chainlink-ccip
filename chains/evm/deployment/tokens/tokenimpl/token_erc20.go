package tokenimpl

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type tokenERC20 struct{}

func (tokenERC20) ContractType() deployment.ContractType {
	return erc20.ContractType
}

func (tokenERC20) Capabilities() CapabilitySet {
	return CapabilitySet{
		ParticipatesInPoolRoleGrant: false,
		SupportsAdminRole:           false,
		SupportsCCIPAdmin:           false,
		SupportsPreMint:             false,
	}
}

func (tokenERC20) RevokeAdminRole(_ operations.Bundle, _ evm.Chain, _, _ common.Address) ([]contract.WriteOutput, error) {
	return nil, fmt.Errorf("admin role not supported for plain ERC20 token")
}

func (tokenERC20) HasAdminRole(_ context.Context, _ evm.Chain, _, _ common.Address) (bool, error) {
	return false, fmt.Errorf("admin role not supported for plain ERC20 token")
}

func (tokenERC20) KnownAdminRoleHolders(_ context.Context, _ evm.Chain, _ common.Address) ([]common.Address, error) {
	return nil, fmt.Errorf("admin role not supported for plain ERC20 token")
}

func (tokenERC20) GrantAdminRole(_ operations.Bundle, _ evm.Chain, _, _ common.Address) ([]contract.WriteOutput, error) {
	return nil, fmt.Errorf("admin role granting not supported for plain ERC20 token")
}

func (tokenERC20) GrantPoolRoles(_ operations.Bundle, _ evm.Chain, _, _, _ common.Address) ([]contract.WriteOutput, error) {
	return nil, fmt.Errorf("pool role granting not supported for plain ERC20 token")
}

func (tokenERC20) SetCCIPAdmin(_ operations.Bundle, _ evm.Chain, _, _ common.Address) ([]contract.WriteOutput, error) {
	return nil, fmt.Errorf("CCIP admin role not supported for plain ERC20 token")
}

func (tokenERC20) Transfer(b operations.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	return transferTokensERC20(b, chain, token, to, scaledAmount)
}

func (tokenERC20) Deploy(b operations.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	ref, err := contract.MaybeDeployContract(b, erc20.Deploy, chain,
		contract.DeployInput[erc20.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(erc20.ContractType, *utils.Version_1_0_0),
			ChainSelector:  chain.Selector,
			Qualifier:      &in.Symbol,
			Args: erc20.ConstructorArgs{
				Name:   in.Name,
				Symbol: in.Symbol,
			},
		},
		nil,
	)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy ERC20 token: %w", err)
	}

	return ref, nil, nil
}
