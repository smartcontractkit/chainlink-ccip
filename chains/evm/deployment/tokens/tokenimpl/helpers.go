package tokenimpl

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	evmops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/erc20"
	bnm_erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
)

func revokeDefaultAdminRoleBurnMintERC20(b cldf_ops.Bundle, chain evm.Chain, token, user common.Address) ([]contract.WriteOutput, error) {
	tokenContract, err := bnm_erc20_bindings.NewBurnMintERC20(token, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate BurnMintERC20 contract: %w", err)
	}
	role, err := tokenContract.DEFAULTADMINROLE(&bind.CallOpts{Context: b.GetContext()})
	if err != nil {
		return nil, fmt.Errorf("failed to get default admin role constant: %w", err)
	}
	report, err := evmops.ExecuteWrite(b, chain, token, bnm_erc20_bindings.NewBurnMintERC20, burn_mint_erc20.NewWriteRevokeAdminRole, burn_mint_erc20.RoleAssignment{
		Role: role,
		To:   user,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to revoke default admin role: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func hasDefaultAdminRoleBurnMintERC20(ctx context.Context, chain evm.Chain, token, user common.Address) (bool, error) {
	tokenContract, err := bnm_erc20_bindings.NewBurnMintERC20(token, chain.Client)
	if err != nil {
		return false, fmt.Errorf("failed to instantiate BurnMintERC20 contract: %w", err)
	}
	role, err := tokenContract.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
	if err != nil {
		return false, fmt.Errorf("failed to get default admin role constant: %w", err)
	}
	hasRole, err := tokenContract.HasRole(&bind.CallOpts{Context: ctx}, role, user)
	if err != nil {
		return false, fmt.Errorf("failed to check default admin role for %s: %w", user.Hex(), err)
	}
	return hasRole, nil
}

func grantDefaultAdminRoleBurnMintERC20(b cldf_ops.Bundle, chain evm.Chain, token, user common.Address) ([]contract.WriteOutput, error) {
	tokenContract, err := bnm_erc20_bindings.NewBurnMintERC20(token, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate BurnMintERC20 contract: %w", err)
	}
	role, err := tokenContract.DEFAULTADMINROLE(&bind.CallOpts{Context: b.GetContext()})
	if err != nil {
		return nil, fmt.Errorf("failed to get default admin role constant: %w", err)
	}
	report, err := evmops.ExecuteWrite(b, chain, token, bnm_erc20_bindings.NewBurnMintERC20, burn_mint_erc20.NewWriteGrantAdminRole, burn_mint_erc20.RoleAssignment{
		Role: role,
		To:   user,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant default admin role: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func grantMintAndBurnRolesBurnMintERC20(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address) ([]contract.WriteOutput, error) {
	report, err := evmops.ExecuteWrite(b, chain, token, bnm_erc20_bindings.NewBurnMintERC20, burn_mint_erc20.NewWriteGrantMintAndBurnRoles, pool)
	if err != nil {
		return nil, fmt.Errorf("failed to grant mint and burn roles: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func setCCIPAdminBurnMintERC20(b cldf_ops.Bundle, chain evm.Chain, token, ccipAdmin common.Address) ([]contract.WriteOutput, error) {
	report, err := evmops.ExecuteWrite(b, chain, token, bnm_erc20_bindings.NewBurnMintERC20, burn_mint_erc20.NewWriteSetCCIPAdmin, ccipAdmin.Hex())
	if err != nil {
		return nil, fmt.Errorf("failed to set CCIP admin: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func transferTokensERC20(b cldf_ops.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	report, err := evmops.ExecuteWrite(b, chain, token, erc20_bindings.NewERC20, erc20.NewWriteTransfer, erc20.TransferArgs{
		Amount:   scaledAmount,
		Receiver: to,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to transfer ERC20 tokens: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}
