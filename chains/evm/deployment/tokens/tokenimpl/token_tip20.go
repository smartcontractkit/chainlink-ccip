package tokenimpl

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/tip20"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type tokenTIP20 struct{}

func (tokenTIP20) ContractType() deployment.ContractType {
	return tip20.ContractType
}

func (tokenTIP20) Capabilities() CapabilitySet {
	return CapabilitySet{
		ParticipatesInPoolRoleGrant: true,
		SupportsAdminRole:           true,
		SupportsCCIPAdmin:           false,
		SupportsPreMint:             false,
	}
}

func (tokenTIP20) RevokeAdminRole(b operations.Bundle, chain evm.Chain, token, user common.Address) ([]contract.WriteOutput, error) {
	report, err := operations.ExecuteOperation(b, tip20.RevokeAdminRole, chain, contract.FunctionInput[common.Address]{
		ChainSelector: chain.Selector,
		Address:       token,
		Args:          user,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to revoke TIP-20 admin role: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func (tokenTIP20) HasAdminRole(ctx context.Context, chain evm.Chain, token, user common.Address) (bool, error) {
	tokenContract, err := tip20.NewTIP20Token(token, chain.Client)
	if err != nil {
		return false, fmt.Errorf("failed to instantiate TIP-20 token contract: %w", err)
	}
	hasRole, err := tokenContract.HasRole(&bind.CallOpts{Context: ctx}, user, tip20.DefaultAdminRole)
	if err != nil {
		return false, fmt.Errorf("failed to check TIP-20 admin role for %s: %w", user.Hex(), err)
	}
	return hasRole, nil
}

func (tokenTIP20) KnownAdminRoleHolders(ctx context.Context, chain evm.Chain, token common.Address) ([]common.Address, error) {
	return currentRoleHoldersFromLogs(ctx, chain, token, tip20.DefaultAdminRole)
}

func (tokenTIP20) GrantAdminRole(b operations.Bundle, chain evm.Chain, token, user common.Address) ([]contract.WriteOutput, error) {
	report, err := operations.ExecuteOperation(b, tip20.GrantAdminRole, chain, contract.FunctionInput[common.Address]{
		ChainSelector: chain.Selector,
		Address:       token,
		Args:          user,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant TIP-20 admin role: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func (tokenTIP20) GrantPoolRoles(b operations.Bundle, chain evm.Chain, token, pool, _ common.Address) ([]contract.WriteOutput, error) {
	report, err := operations.ExecuteOperation(b, tip20.GrantIssuerRole, chain, contract.FunctionInput[common.Address]{
		ChainSelector: chain.Selector,
		Address:       token,
		Args:          pool,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant TIP-20 issuer role: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func (tokenTIP20) SetCCIPAdmin(b operations.Bundle, chain evm.Chain, token, ccipAdmin common.Address) ([]contract.WriteOutput, error) {
	return nil, fmt.Errorf("CCIP admin role not supported for TIP-20 tokens")
}

func (tokenTIP20) Transfer(b operations.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	report, err := operations.ExecuteOperation(b, tip20.Transfer, chain, contract.FunctionInput[tip20.TransferArgs]{
		ChainSelector: chain.Selector,
		Address:       token,
		Args: tip20.TransferArgs{
			Amount:   scaledAmount,
			Receiver: to,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to transfer TIP-20 tokens: %w", err)
	}

	return []contract.WriteOutput{report.Output}, nil
}

func (tokenTIP20) Deploy(b operations.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	tokenRef, writes, err := tip20.DeployTokenViaFactory(b, chain, tip20.FactoryDeployArgs{
		QuoteToken: common.Address{},
		Currency:   in.Currency,
		Salt:       [32]byte{},
		Symbol:     in.Symbol,
		Admin:      chain.DeployerKey.From,
		Name:       in.Name,
	})
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy TIP-20 token via factory: %w", err)
	}

	return tokenRef, writes, nil
}
