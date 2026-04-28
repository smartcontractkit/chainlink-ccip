package registrations

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	bnm_erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
)

func scaledSupplyAndPreMint(in tokensapi.DeployTokenInput) (*big.Int, *big.Int) {
	maxSupply := big.NewInt(0)
	if in.Supply != nil {
		maxSupply = tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*in.Supply), in.Decimals)
	}
	preMint := big.NewInt(0)
	if in.PreMint != nil {
		preMint = tokensapi.ScaleTokenAmount(new(big.Int).SetUint64(*in.PreMint), in.Decimals)
	}
	return maxSupply, preMint
}

// grantBnMMintAndBurnRoles is shared by all BnM-family strategies.
// Historically the BnM, BnM+Drip (v1.0.0), and BnM+Drip (v1.5.0) types
// all dispatch to burn_mint_erc20.GrantMintAndBurnRoles (the v1.0.0 op);
// preserved here verbatim.
func grantBnMMintAndBurnRoles(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20.GrantMintAndBurnRoles, chain, contract.FunctionInput[common.Address]{
		ChainSelector: chainSelector,
		Address:       token,
		Args:          pool,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant mint and burn roles: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func grantBnMDefaultAdminRole(b cldf_ops.Bundle, chain evm.Chain, token, externalAdmin common.Address, chainSelector uint64) ([]contract.WriteOutput, error) {
	tokenContract, err := bnm_erc20_bindings.NewBurnMintERC20(token, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate BurnMintERC20 contract: %w", err)
	}
	role, err := tokenContract.DEFAULTADMINROLE(&bind.CallOpts{Context: b.GetContext()})
	if err != nil {
		return nil, fmt.Errorf("failed to get default admin role constant: %w", err)
	}
	report, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20.GrantAdminRole, chain, contract.FunctionInput[burn_mint_erc20.RoleAssignment]{
		ChainSelector: chainSelector,
		Address:       token,
		Args: burn_mint_erc20.RoleAssignment{
			Role: role,
			To:   externalAdmin,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant default admin role: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}
