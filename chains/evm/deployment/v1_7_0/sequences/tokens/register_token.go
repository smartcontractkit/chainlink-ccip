package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// RegisterTokenInput is the input for the RegisterToken sequence.
type RegisterTokenInput struct {
	// ChainSelector is the chain selector for the chain being configured.
	ChainSelector uint64
	// TokenPoolAddress is the address of the token pool.
	TokenPoolAddress common.Address
	// AdminAddress is the address of the desired token admin.
	AdminAddress common.Address
	// OnlyPropose indicates whether to only propose the admin, and not accept the role or set the pool.
	// Used in cases where we want to propose to an administrator that we don't control.
	OnlyPropose bool
	// TokenAdminRegistryAddress is the address of the TokenAdminRegistry contract.
	TokenAdminRegistryAddress common.Address
}

func (c RegisterTokenInput) Validate(chain evm.Chain) error {
	if c.ChainSelector != chain.Selector {
		return fmt.Errorf("chain selector %d does not match chain %s", c.ChainSelector, chain)
	}
	return nil
}

var RegisterToken = cldf_ops.NewSequence(
	"register-token",
	semver.MustParse("1.7.0"),
	"Registers a token with CCIP via the token admin registry",
	func(b operations.Bundle, chain evm.Chain, input RegisterTokenInput) (output sequences.OnChainOutput, err error) {
		if err := input.Validate(chain); err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("invalid input: %w", err)
		}
		writes := make([]contract.WriteOutput, 0)

		// Get the current token config from the token admin registry.
		tokenConfigReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.GetTokenConfig, chain, contract.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenAdminRegistryAddress,
			Args:          input.TokenPoolAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token config: %w", err)
		}

		// Propose the admin for the token if no admin exists yet.
		if tokenConfigReport.Output.Administrator == (common.Address{}) {
			proposeAdminReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.ProposeAdministrator, chain, contract.FunctionInput[token_admin_registry.ProposeAdministratorArgs]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenAdminRegistryAddress,
				Args: token_admin_registry.ProposeAdministratorArgs{
					TokenAddress:  input.TokenPoolAddress,
					Administrator: input.AdminAddress,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to propose admin: %w", err)
			}
			writes = append(writes, proposeAdminReport.Output)
		}

		if !input.OnlyPropose {
			// Accept the admin role for the token.
			acceptAdminReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.AcceptAdminRole, chain, contract.FunctionInput[token_admin_registry.AcceptAdminRoleArgs]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenAdminRegistryAddress,
				Args: token_admin_registry.AcceptAdminRoleArgs{
					TokenAddress: input.TokenPoolAddress,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to accept admin role: %w", err)
			}
			writes = append(writes, acceptAdminReport.Output)

			// Set the token pool on the token admin registry.
			setPoolReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.SetPool, chain, contract.FunctionInput[token_admin_registry.SetPoolArgs]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenAdminRegistryAddress,
				Args: token_admin_registry.SetPoolArgs{
					TokenAddress:     input.TokenPoolAddress,
					TokenPoolAddress: input.TokenPoolAddress,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to set token pool: %w", err)
			}
			writes = append(writes, setPoolReport.Output)
		}

		return sequences.OnChainOutput{Writes: writes}, nil
	},
)
