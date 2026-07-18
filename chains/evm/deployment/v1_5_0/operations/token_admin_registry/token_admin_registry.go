package token_admin_registry

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations2/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cld_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
)

var (
	ContractType cldf_deployment.ContractType = "TokenAdminRegistry"
	Version      *semver.Version              = semver.MustParse("1.5.0")
)

type ConstructorArgs struct{}

type TokenConfig = token_admin_registry.TokenAdminRegistryTokenConfig

type ProposeAdministratorArgs struct {
	TokenAddress  common.Address
	Administrator common.Address
}

type TransferAdminRoleArgs struct {
	TokenAddress  common.Address
	Administrator common.Address
}

type AcceptAdminRoleArgs struct {
	TokenAddress common.Address
}

type SetPoolArgs struct {
	TokenAddress     common.Address
	TokenPoolAddress common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "token-admin-registry:deploy",
	Version:          Version,
	Description:      "Deploys the TokenAdminRegistry contract",
	ContractMetadata: token_admin_registry.TokenAdminRegistryMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(token_admin_registry.TokenAdminRegistryBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

func NewWriteProposeAdministrator(c *token_admin_registry.TokenAdminRegistry) *cld_ops.Operation[contract.FunctionInput[ProposeAdministratorArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[ProposeAdministratorArgs, *token_admin_registry.TokenAdminRegistry]{
		Name:         "token-admin-registry:propose-administrator",
		Version:      Version,
		Description:  "Proposes an administrator for a token on the TokenAdminRegistry contract",
		ContractType: ContractType,
		ContractABI:  token_admin_registry.TokenAdminRegistryABI,
		Contract:     c,
		IsAllowedCaller: func(c *token_admin_registry.TokenAdminRegistry, opts *bind.CallOpts, caller common.Address, input ProposeAdministratorArgs) (bool, error) {
			owner, err := c.Owner(opts)
			if err != nil {
				return false, fmt.Errorf("failed to get owner: %w", err)
			}
			tokenConfig, err := c.GetTokenConfig(opts, input.TokenAddress)
			if err != nil {
				return false, fmt.Errorf("failed to get token config: %w", err)
			}
			return caller == owner && tokenConfig.Administrator == (common.Address{}), nil
		},
		Validate: func(ProposeAdministratorArgs) error { return nil },
		CallContract: func(tokenAdminRegistry *token_admin_registry.TokenAdminRegistry, opts *bind.TransactOpts, args ProposeAdministratorArgs) (*types.Transaction, error) {
			return tokenAdminRegistry.ProposeAdministrator(opts, args.TokenAddress, args.Administrator)
		},
	})
}

func NewWriteTransferAdminRole(c *token_admin_registry.TokenAdminRegistry) *cld_ops.Operation[contract.FunctionInput[TransferAdminRoleArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[TransferAdminRoleArgs, *token_admin_registry.TokenAdminRegistry]{
		Name:         "token-admin-registry:transfer-admin-role",
		Version:      Version,
		Description:  "Transfers the admin role for a token on the TokenAdminRegistry contract",
		ContractType: ContractType,
		ContractABI:  token_admin_registry.TokenAdminRegistryABI,
		Contract:     c,
		IsAllowedCaller: func(c *token_admin_registry.TokenAdminRegistry, opts *bind.CallOpts, caller common.Address, input TransferAdminRoleArgs) (bool, error) {
			tokenConfig, err := c.GetTokenConfig(opts, input.TokenAddress)
			if err != nil {
				return false, fmt.Errorf("failed to get token config: %w", err)
			}
			return tokenConfig.Administrator == caller, nil
		},
		Validate: func(TransferAdminRoleArgs) error { return nil },
		CallContract: func(tokenAdminRegistry *token_admin_registry.TokenAdminRegistry, opts *bind.TransactOpts, args TransferAdminRoleArgs) (*types.Transaction, error) {
			return tokenAdminRegistry.TransferAdminRole(opts, args.TokenAddress, args.Administrator)
		},
	})
}

func NewWriteAcceptAdminRole(c *token_admin_registry.TokenAdminRegistry) *cld_ops.Operation[contract.FunctionInput[AcceptAdminRoleArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[AcceptAdminRoleArgs, *token_admin_registry.TokenAdminRegistry]{
		Name:         "token-admin-registry:accept-admin-role",
		Version:      Version,
		Description:  "Accepts the admin role for a token on the TokenAdminRegistry contract",
		ContractType: ContractType,
		ContractABI:  token_admin_registry.TokenAdminRegistryABI,
		Contract:     c,
		IsAllowedCaller: func(c *token_admin_registry.TokenAdminRegistry, opts *bind.CallOpts, caller common.Address, args AcceptAdminRoleArgs) (bool, error) {
			tokenConfig, err := c.GetTokenConfig(opts, args.TokenAddress)
			if err != nil {
				return false, fmt.Errorf("failed to get token config: %w", err)
			}
			return tokenConfig.PendingAdministrator == caller, nil
		},
		Validate: func(AcceptAdminRoleArgs) error { return nil },
		CallContract: func(tokenAdminRegistry *token_admin_registry.TokenAdminRegistry, opts *bind.TransactOpts, args AcceptAdminRoleArgs) (*types.Transaction, error) {
			return tokenAdminRegistry.AcceptAdminRole(opts, args.TokenAddress)
		},
	})
}

func NewWriteSetPool(c *token_admin_registry.TokenAdminRegistry) *cld_ops.Operation[contract.FunctionInput[SetPoolArgs], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[SetPoolArgs, *token_admin_registry.TokenAdminRegistry]{
		Name:         "token-admin-registry:set-pool",
		Version:      Version,
		Description:  "Sets the token pool for a token on the TokenAdminRegistry contract",
		ContractType: ContractType,
		ContractABI:  token_admin_registry.TokenAdminRegistryABI,
		Contract:     c,
		IsAllowedCaller: func(c *token_admin_registry.TokenAdminRegistry, opts *bind.CallOpts, caller common.Address, input SetPoolArgs) (bool, error) {
			tokenConfig, err := c.GetTokenConfig(opts, input.TokenAddress)
			if err != nil {
				return false, fmt.Errorf("failed to get token config: %w", err)
			}
			return tokenConfig.Administrator == caller, nil
		},
		Validate: func(SetPoolArgs) error { return nil },
		CallContract: func(tokenAdminRegistry *token_admin_registry.TokenAdminRegistry, opts *bind.TransactOpts, args SetPoolArgs) (*types.Transaction, error) {
			return tokenAdminRegistry.SetPool(opts, args.TokenAddress, args.TokenPoolAddress)
		},
	})
}

func NewWriteAddRegistryModule(c *token_admin_registry.TokenAdminRegistry) *cld_ops.Operation[contract.FunctionInput[common.Address], contract.WriteOutput, cldf_evm.Chain] {
	return contract.NewWrite(contract.WriteParams[common.Address, *token_admin_registry.TokenAdminRegistry]{
		Name:            "token-admin-registry:add-registry-module",
		Version:         Version,
		Description:     "Adds a registry module to the TokenAdminRegistry contract",
		ContractType:    ContractType,
		ContractABI:     token_admin_registry.TokenAdminRegistryABI,
		Contract:        c,
		IsAllowedCaller: contract.OnlyOwner[*token_admin_registry.TokenAdminRegistry, common.Address],
		Validate:        func(common.Address) error { return nil },
		CallContract: func(tokenAdminRegistry *token_admin_registry.TokenAdminRegistry, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
			return tokenAdminRegistry.AddRegistryModule(opts, args)
		},
	})
}

func NewReadIsRegistryModule(c *token_admin_registry.TokenAdminRegistry) *cld_ops.Operation[contract.FunctionInput[common.Address], bool, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[common.Address, bool, *token_admin_registry.TokenAdminRegistry]{
		Name:         "token-admin-registry:is-registry-module",
		Version:      Version,
		Description:  "Checks if an address is a registry module in the TokenAdminRegistry contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(tokenAdminRegistry *token_admin_registry.TokenAdminRegistry, opts *bind.CallOpts, args common.Address) (bool, error) {
			return tokenAdminRegistry.IsRegistryModule(opts, args)
		},
	})
}

func NewReadOwner(c *token_admin_registry.TokenAdminRegistry) *cld_ops.Operation[contract.FunctionInput[struct{}], common.Address, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[struct{}, common.Address, *token_admin_registry.TokenAdminRegistry]{
		Name:         "token-admin-registry:owner",
		Version:      Version,
		Description:  "Gets the owner of the TokenAdminRegistry contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(tokenAdminRegistry *token_admin_registry.TokenAdminRegistry, opts *bind.CallOpts, args struct{}) (common.Address, error) {
			return tokenAdminRegistry.Owner(opts)
		},
	})
}

func NewReadGetTokenConfig(c *token_admin_registry.TokenAdminRegistry) *cld_ops.Operation[contract.FunctionInput[common.Address], TokenConfig, cldf_evm.Chain] {
	return contract.NewRead(contract.ReadParams[common.Address, TokenConfig, *token_admin_registry.TokenAdminRegistry]{
		Name:         "token-admin-registry:get-token-config",
		Version:      Version,
		Description:  "Gets the token configuration for a given token address from the TokenAdminRegistry contract",
		ContractType: ContractType,
		Contract:     c,
		CallContract: func(tokenAdminRegistry *token_admin_registry.TokenAdminRegistry, opts *bind.CallOpts, args common.Address) (TokenConfig, error) {
			return tokenAdminRegistry.GetTokenConfig(opts, args)
		},
	})
}
