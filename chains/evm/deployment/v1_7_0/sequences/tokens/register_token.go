package tokens

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_contract "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/sequences"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

// RegisterTokenInput is the input for the RegisterToken sequence.
type RegisterTokenInput struct {
	// ChainSelector is the chain selector for the chain being configured.
	ChainSelector uint64
	// TokenAddress is the address of the token to register.
	TokenAddress common.Address
	// TokenPoolAddress is the address of the token pool.
	TokenPoolAddress common.Address
	// ExternalAdmin is specified when we want to propose an admin that we don't control.
	// In this case, only proposeAdministrator will be called.
	ExternalAdmin common.Address
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
		tokenConfigReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.GetTokenConfig, chain, evm_contract.FunctionInput[common.Address]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenAdminRegistryAddress,
			Args:          input.TokenAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get token config: %w", err)
		}
		admin := tokenConfigReport.Output.Administrator
		pendingAdmin := tokenConfigReport.Output.PendingAdministrator
		tokenPoolSet := tokenConfigReport.Output.TokenPool

		// Get the owner of the token admin registry.
		ownerReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.Owner, chain, evm_contract.FunctionInput[any]{
			ChainSelector: input.ChainSelector,
			Address:       input.TokenAdminRegistryAddress,
		})
		if err != nil {
			return sequences.OnChainOutput{}, fmt.Errorf("failed to get owner of token admin registry: %w", err)
		}
		registryOwner := ownerReport.Output

		// Propose the admin for the token if no admin exists yet.
		if admin == (common.Address{}) {
			desiredAdmin := input.ExternalAdmin
			if desiredAdmin == (common.Address{}) {
				// If no external admin is specified, we set the desired admin to the registry owner.
				desiredAdmin = registryOwner
			}
			proposeAdminReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.ProposeAdministrator, chain, evm_contract.FunctionInput[token_admin_registry.ProposeAdministratorArgs]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenAdminRegistryAddress,
				Args: token_admin_registry.ProposeAdministratorArgs{
					TokenAddress:  input.TokenAddress,
					Administrator: desiredAdmin,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to propose admin: %w", err)
			}
			writes = append(writes, proposeAdminReport.Output)
			pendingAdmin = desiredAdmin
		}

		// Accept the admin role if the registry owner is the pending admin.
		if pendingAdmin == registryOwner {
			acceptAdminReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.AcceptAdminRole, chain, evm_contract.FunctionInput[token_admin_registry.AcceptAdminRoleArgs]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenAdminRegistryAddress,
				Args: token_admin_registry.AcceptAdminRoleArgs{
					TokenAddress: input.TokenAddress,
				},
			})
			if err != nil {
				return sequences.OnChainOutput{}, fmt.Errorf("failed to accept admin role: %w", err)
			}
			writes = append(writes, acceptAdminReport.Output)
			admin = registryOwner
		}

		// Set the token pool on the token admin registry if the registry owner is the admin and the pool is not set.
		if admin == registryOwner && tokenPoolSet != input.TokenPoolAddress {
			setPoolReport, err := cldf_ops.ExecuteOperation(b, token_admin_registry.SetPool, chain, evm_contract.FunctionInput[token_admin_registry.SetPoolArgs]{
				ChainSelector: input.ChainSelector,
				Address:       input.TokenAdminRegistryAddress,
				Args: token_admin_registry.SetPoolArgs{
					TokenAddress:     input.TokenAddress,
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
