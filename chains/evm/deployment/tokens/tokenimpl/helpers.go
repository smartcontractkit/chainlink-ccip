package tokenimpl

import (
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/burn_mint_erc20"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/operations/erc20"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

func getRetryConfig[T any](b cldf_ops.Bundle, chain evm.Chain, token string) cldf_ops.RetryConfig[contract.FunctionInput[T], evm.Chain] {
	retryBuffer := 2 * time.Second
	retryPolicy := cldf_ops.RetryPolicy{MaxAttempts: 5}
	retryConfig := cldf_ops.RetryConfig[contract.FunctionInput[T], evm.Chain]{
		Enabled: true,
		Policy:  retryPolicy,
		InputHook: func(attempt uint, err error, in contract.FunctionInput[T], _ evm.Chain) contract.FunctionInput[T] {
			wait := time.Duration(attempt+1) * retryBuffer

			b.Logger.Infof(
				"Retrying operation for token %s on chain %d, attempt %d/%d, waiting %s, err: %v",
				token, chain.Selector, attempt+1, retryPolicy.MaxAttempts, wait, err,
			)

			// We avoid time.Sleep since it doesn't respect context cancellation. Instead, we either wait
			// for the entire backoff duration or for the context to be cancelled, whichever comes first.
			select {
			case <-b.GetContext().Done():
			case <-time.After(wait):
			}

			return in
		},
	}

	return retryConfig
}

func revokeDefaultAdminRoleBurnMintERC20(b cldf_ops.Bundle, chain evm.Chain, token, user common.Address) ([]contract.WriteOutput, error) {
	role, err := cldf_ops.ExecuteOperation(
		b, burn_mint_erc20.GetDefaultAdminRole, chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args:          struct{}{},
		},
		cldf_ops.WithRetryConfig(getRetryConfig[struct{}](b, chain, token.Hex())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get default admin role: %w", err)
	}

	report, err := cldf_ops.ExecuteOperation(
		b, burn_mint_erc20.RevokeAdminRole, chain,
		contract.FunctionInput[burn_mint_erc20.RoleAssignment]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args: burn_mint_erc20.RoleAssignment{
				Role: role.Output,
				To:   user,
			},
		},
		cldf_ops.WithRetryConfig(getRetryConfig[burn_mint_erc20.RoleAssignment](b, chain, token.Hex())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to revoke default admin role: %w", err)
	}

	return []contract.WriteOutput{report.Output}, nil
}

func hasDefaultAdminRoleBurnMintERC20(b cldf_ops.Bundle, chain evm.Chain, token, user common.Address) (bool, error) {
	role, err := cldf_ops.ExecuteOperation(
		b, burn_mint_erc20.GetDefaultAdminRole, chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args:          struct{}{},
		},
		cldf_ops.WithRetryConfig(getRetryConfig[struct{}](b, chain, token.Hex())),
	)
	if err != nil {
		return false, fmt.Errorf("failed to get default admin role: %w", err)
	}

	hasRole, err := cldf_ops.ExecuteOperation(
		b, burn_mint_erc20.HasRole, chain,
		contract.FunctionInput[burn_mint_erc20.RoleAssignment]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args: burn_mint_erc20.RoleAssignment{
				Role: role.Output,
				To:   user,
			},
		},
		cldf_ops.WithRetryConfig(getRetryConfig[burn_mint_erc20.RoleAssignment](b, chain, token.Hex())),
	)
	if err != nil {
		return false, fmt.Errorf("failed to check default admin role for %s: %w", user.Hex(), err)
	}

	return hasRole.Output, nil
}

func grantDefaultAdminRoleBurnMintERC20(b cldf_ops.Bundle, chain evm.Chain, token, user common.Address) ([]contract.WriteOutput, error) {
	role, err := cldf_ops.ExecuteOperation(
		b, burn_mint_erc20.GetDefaultAdminRole, chain,
		contract.FunctionInput[struct{}]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args:          struct{}{},
		},
		cldf_ops.WithRetryConfig(getRetryConfig[struct{}](b, chain, token.Hex())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get default admin role: %w", err)
	}

	report, err := cldf_ops.ExecuteOperation(
		b, burn_mint_erc20.GrantAdminRole, chain,
		contract.FunctionInput[burn_mint_erc20.RoleAssignment]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args: burn_mint_erc20.RoleAssignment{
				Role: role.Output,
				To:   user,
			},
		},
		cldf_ops.WithRetryConfig(getRetryConfig[burn_mint_erc20.RoleAssignment](b, chain, token.Hex())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to grant default admin role: %w", err)
	}

	return []contract.WriteOutput{report.Output}, nil
}

func grantMintAndBurnRolesBurnMintERC20(b cldf_ops.Bundle, chain evm.Chain, token, pool common.Address) ([]contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(
		b, burn_mint_erc20.GrantMintAndBurnRoles, chain,
		contract.FunctionInput[common.Address]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args:          pool,
		},
		cldf_ops.WithRetryConfig(getRetryConfig[common.Address](b, chain, token.Hex())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to grant mint and burn roles: %w", err)
	}

	return []contract.WriteOutput{report.Output}, nil
}

func setCCIPAdminBurnMintERC20(b cldf_ops.Bundle, chain evm.Chain, token, ccipAdmin common.Address) ([]contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(
		b, burn_mint_erc20.SetCCIPAdmin, chain,
		contract.FunctionInput[string]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args:          ccipAdmin.Hex(),
		},
		cldf_ops.WithRetryConfig(getRetryConfig[string](b, chain, token.Hex())),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to set CCIP admin: %w", err)
	}

	return []contract.WriteOutput{report.Output}, nil
}

// NOTE: transferTokensERC20 is intentionally NOT retried. Transfer moves value and is not idempotent
// by value - a retry after an unclear outcome *risks transferring twice*. The retry path is reserved
// for idempotent role-assignment writes and reads.
func transferTokensERC20(b cldf_ops.Bundle, chain evm.Chain, token, to common.Address, scaledAmount *big.Int) ([]contract.WriteOutput, error) {
	report, err := cldf_ops.ExecuteOperation(
		b, erc20.Transfer, chain, contract.FunctionInput[erc20.TransferArgs]{
			ChainSelector: chain.Selector,
			Address:       token,
			Args: erc20.TransferArgs{
				Amount:   scaledAmount,
				Receiver: to,
			},
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to transfer ERC20 tokens: %w", err)
	}

	return []contract.WriteOutput{report.Output}, nil
}
