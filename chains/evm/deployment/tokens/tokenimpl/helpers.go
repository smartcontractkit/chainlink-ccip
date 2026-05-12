package tokenimpl

import (
	"context"
	"fmt"
	"math/big"
	"sort"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	bnm_erc20_bindings "github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/burn_mint_erc20"
)

var (
	roleGrantedTopic = crypto.Keccak256Hash([]byte("RoleGranted(bytes32,address,address)"))
	roleRevokedTopic = crypto.Keccak256Hash([]byte("RoleRevoked(bytes32,address,address)"))
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
	report, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20.RevokeAdminRole, chain, contract.FunctionInput[burn_mint_erc20.RoleAssignment]{
		ChainSelector: chain.Selector,
		Address:       token,
		Args: burn_mint_erc20.RoleAssignment{
			Role: role,
			To:   user,
		},
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

func knownDefaultAdminRoleHoldersBurnMintERC20(ctx context.Context, chain evm.Chain, token common.Address) ([]common.Address, error) {
	tokenContract, err := bnm_erc20_bindings.NewBurnMintERC20(token, chain.Client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate BurnMintERC20 contract: %w", err)
	}
	role, err := tokenContract.DEFAULTADMINROLE(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, fmt.Errorf("failed to get default admin role constant: %w", err)
	}
	return currentRoleHoldersFromLogs(ctx, chain, token, role)
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
	report, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20.GrantAdminRole, chain, contract.FunctionInput[burn_mint_erc20.RoleAssignment]{
		ChainSelector: chain.Selector,
		Address:       token,
		Args: burn_mint_erc20.RoleAssignment{
			Role: role,
			To:   user,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant default admin role: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func currentRoleHoldersFromLogs(ctx context.Context, chain evm.Chain, token common.Address, role [32]byte) ([]common.Address, error) {
	roleTopic := common.BytesToHash(role[:])
	logs, err := chain.Client.FilterLogs(ctx, ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		Addresses: []common.Address{token},
		Topics: [][]common.Hash{
			{roleGrantedTopic, roleRevokedTopic},
			{roleTopic},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to filter role grant and revoke logs: %w", err)
	}

	current := make(map[common.Address]bool)
	for _, log := range logs {
		account, ok := roleLogAccount(log)
		if !ok {
			continue
		}
		switch log.Topics[0] {
		case roleGrantedTopic:
			current[account] = true
		case roleRevokedTopic:
			current[account] = false
		}
	}

	holders := make([]common.Address, 0, len(current))
	for account, hasRole := range current {
		if hasRole {
			holders = append(holders, account)
		}
	}
	sort.Slice(holders, func(i, j int) bool {
		return holders[i].Hex() < holders[j].Hex()
	})
	return holders, nil
}

func roleLogAccount(log ethtypes.Log) (common.Address, bool) {
	if len(log.Topics) < 3 {
		return common.Address{}, false
	}
	account := common.BytesToAddress(log.Topics[2].Bytes())
	return account, account != (common.Address{})
}

func grantMintAndBurnRolesBurnMintERC20(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address) ([]contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20.GrantMintAndBurnRoles, chain, contract.FunctionInput[common.Address]{
		ChainSelector: chain.Selector,
		Address:       token,
		Args:          pool,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to grant mint and burn roles: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func setCCIPAdminBurnMintERC20(b cldf_ops.Bundle, chain evm.Chain, token, ccipAdmin common.Address) ([]contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b, burn_mint_erc20.SetCCIPAdmin, chain, contract.FunctionInput[string]{
		ChainSelector: chain.Selector,
		Address:       token,
		Args:          ccipAdmin.Hex(),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to set CCIP admin: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}

func transferTokensERC20(b cldf_ops.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(b, erc20.Transfer, chain, contract.FunctionInput[erc20.TransferArgs]{
		ChainSelector: chain.Selector,
		Address:       token,
		Args: erc20.TransferArgs{
			Amount:   scaledAmount,
			Receiver: to,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to transfer ERC20 tokens: %w", err)
	}
	return []contract.WriteOutput{report.Output}, nil
}
