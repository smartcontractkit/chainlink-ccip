package token_admin_registry

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "TokenAdminRegistry"
var Version *semver.Version = semver.MustParse("1.5.0")

type ConstructorArgs struct{}

type TokenConfig = token_admin_registry.TokenAdminRegistryTokenConfig

type ProposeAdministratorArgs struct {
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
	Version:          semver.MustParse("1.5.0"),
	Description:      "Deploys the TokenAdminRegistry contract",
	ContractMetadata: token_admin_registry.TokenAdminRegistryMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.5.0")).String(): {
			EVM: common.FromHex(token_admin_registry.TokenAdminRegistryBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ProposeAdministrator = contract.NewWrite(contract.WriteParams[ProposeAdministratorArgs, *token_admin_registry.TokenAdminRegistry]{
	Name:         "token-admin-registry:propose-administrator",
	Version:      semver.MustParse("1.5.0"),
	Description:  "Proposes an administrator for a token on the TokenAdminRegistry contract",
	ContractType: ContractType,
	ContractABI:  token_admin_registry.TokenAdminRegistryABI,
	NewContract:  token_admin_registry.NewTokenAdminRegistry,
	IsAllowedCaller: func(contract *token_admin_registry.TokenAdminRegistry, opts *bind.CallOpts, caller common.Address, input ProposeAdministratorArgs) (bool, error) {
		owner, err := contract.Owner(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get owner: %w", err)
		}
		tokenConfig, err := contract.GetTokenConfig(opts, input.TokenAddress)
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

var AcceptAdminRole = contract.NewWrite(contract.WriteParams[AcceptAdminRoleArgs, *token_admin_registry.TokenAdminRegistry]{
	Name:         "token-admin-registry:accept-admin-role",
	Version:      semver.MustParse("1.5.0"),
	Description:  "Accepts the admin role for a token on the TokenAdminRegistry contract",
	ContractType: ContractType,
	ContractABI:  token_admin_registry.TokenAdminRegistryABI,
	NewContract:  token_admin_registry.NewTokenAdminRegistry,
	IsAllowedCaller: func(contract *token_admin_registry.TokenAdminRegistry, opts *bind.CallOpts, caller common.Address, args AcceptAdminRoleArgs) (bool, error) {
		tokenConfig, err := contract.GetTokenConfig(opts, args.TokenAddress)
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

var SetPool = contract.NewWrite(contract.WriteParams[SetPoolArgs, *token_admin_registry.TokenAdminRegistry]{
	Name:         "token-admin-registry:set-pool",
	Version:      semver.MustParse("1.5.0"),
	Description:  "Sets the token pool for a token on the TokenAdminRegistry contract",
	ContractType: ContractType,
	ContractABI:  token_admin_registry.TokenAdminRegistryABI,
	NewContract:  token_admin_registry.NewTokenAdminRegistry,
	IsAllowedCaller: func(contract *token_admin_registry.TokenAdminRegistry, opts *bind.CallOpts, caller common.Address, input SetPoolArgs) (bool, error) {
		tokenConfig, err := contract.GetTokenConfig(opts, input.TokenAddress)
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

var Owner = contract.NewRead(contract.ReadParams[any, common.Address, *token_admin_registry.TokenAdminRegistry]{
	Name:         "token-admin-registry:owner",
	Version:      semver.MustParse("1.5.0"),
	Description:  "Gets the owner of the TokenAdminRegistry contract",
	ContractType: ContractType,
	NewContract:  token_admin_registry.NewTokenAdminRegistry,
	CallContract: func(tokenAdminRegistry *token_admin_registry.TokenAdminRegistry, opts *bind.CallOpts, args any) (common.Address, error) {
		return tokenAdminRegistry.Owner(opts)
	},
})

var GetTokenConfig = contract.NewRead(contract.ReadParams[common.Address, TokenConfig, *token_admin_registry.TokenAdminRegistry]{
	Name:         "token-admin-registry:get-token-config",
	Version:      semver.MustParse("1.5.0"),
	Description:  "Gets the token configuration for a given token address from the TokenAdminRegistry contract",
	ContractType: ContractType,
	NewContract:  token_admin_registry.NewTokenAdminRegistry,
	CallContract: func(tokenAdminRegistry *token_admin_registry.TokenAdminRegistry, opts *bind.CallOpts, args common.Address) (TokenConfig, error) {
		return tokenAdminRegistry.GetTokenConfig(opts, args)
	},
})
