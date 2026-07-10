package tokenimpl

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc677"
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type tokenBurnMintERC677 struct{}

func (tokenBurnMintERC677) ContractType() deployment.ContractType {
	return utils.BurnMintToken
}

func (tokenBurnMintERC677) Capabilities() CapabilitySet {
	return CapabilitySet{
		ParticipatesInPoolRoleGrant: true,
		SupportsAdminRole:           false,
		SupportsCCIPAdmin:           false,
		SupportsPreMint:             false,
	}
}

func (tokenBurnMintERC677) RevokeAdminRole(_ cldf_ops.Bundle, _ evm.Chain, _, _ common.Address) ([]contract.WriteOutput, error) {
	return nil, fmt.Errorf("admin role revoke not supported for BurnMintERC677 token type")
}

func (tokenBurnMintERC677) HasAdminRole(_ context.Context, _ evm.Chain, _, _ common.Address) (bool, error) {
	return false, fmt.Errorf("admin role checks not supported for BurnMintERC677 token type")
}

func (tokenBurnMintERC677) GrantAdminRole(_ cldf_ops.Bundle, _ evm.Chain, _, _ common.Address) ([]contract.WriteOutput, error) {
	return nil, fmt.Errorf("admin role grant not supported for BurnMintERC677 token type")
}

func (tokenBurnMintERC677) GrantPoolRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool, proposalExecutor common.Address) ([]contract.WriteOutput, error) {
	return burn_mint_erc677.PrepareGrantMintAndBurnRoles(b, chain, token, pool, proposalExecutor)
}

func (tokenBurnMintERC677) SetCCIPAdmin(_ cldf_ops.Bundle, _ evm.Chain, _, _ common.Address) ([]contract.WriteOutput, error) {
	return nil, fmt.Errorf("CCIP admin not supported for BurnMintERC677 token type via this deployment path")
}

func (tokenBurnMintERC677) Transfer(b cldf_ops.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	return transferTokensERC20(b, chain, token, to, scaledAmount)
}

func (tokenBurnMintERC677) Deploy(b cldf_ops.Bundle, chain evm.Chain, in tokensapi.DeployTokenInput) (datastore.AddressRef, []contract.WriteOutput, error) {
	maxSupply := big.NewInt(0)
	if in.Supply != nil {
		maxSupply = tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*in.Supply), in.Decimals)
	}

	ref, err := evmops.MaybeDeployContract(b, burn_mint_erc677.Deploy, chain,
		contract.DeployInput[burn_mint_erc677.ConstructorArgs]{
			TypeAndVersion: deployment.NewTypeAndVersion(burn_mint_erc677.ContractType, *utils.Version_1_0_0),
			Qualifier:      &in.Symbol,
			Args: burn_mint_erc677.ConstructorArgs{
				Name:      in.Name,
				Symbol:    in.Symbol,
				Decimals:  in.Decimals,
				MaxSupply: maxSupply,
			},
		},
		nil,
	)
	if err != nil {
		return datastore.AddressRef{}, nil, fmt.Errorf("failed to deploy BurnMintERC677 token: %w", err)
	}

	return ref, nil, nil
}
