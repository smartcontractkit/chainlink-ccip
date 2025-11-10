package registry_module_owner_custom

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/registry_module_owner_custom"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "RegistryModuleOwnerCustom"

var Version *semver.Version = semver.MustParse("1.6.0")

type ConstructorArgs struct {
	TokenAdminRegistry common.Address
}

type RegisterAdminViaOwnerArgs struct {
	Token common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "registry-module-owner-custom:deploy",
	Version:          semver.MustParse("1.6.0"),
	Description:      "Deploys the RegistryModuleOwnerCustom contract",
	ContractMetadata: registry_module_owner_custom.RegistryModuleOwnerCustomMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(registry_module_owner_custom.RegistryModuleOwnerCustomBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var RegisterAdminViaOwner = contract.NewWrite(contract.WriteParams[RegisterAdminViaOwnerArgs, *registry_module_owner_custom.RegistryModuleOwnerCustom]{
	Name:            "registry-module-owner-custom:register-admin-via-owner",
	Version:         Version,
	Description:     "Registers an admin for a token, using the token's owner method to prove ownership",
	ContractType:    ContractType,
	ContractABI:     registry_module_owner_custom.RegistryModuleOwnerCustomABI,
	NewContract:     registry_module_owner_custom.NewRegistryModuleOwnerCustom,
	IsAllowedCaller: contract.AllCallersAllowed[*registry_module_owner_custom.RegistryModuleOwnerCustom, RegisterAdminViaOwnerArgs],
	Validate:        func(RegisterAdminViaOwnerArgs) error { return nil },
	CallContract: func(registryModuleOwnerCustom *registry_module_owner_custom.RegistryModuleOwnerCustom, opts *bind.TransactOpts, args RegisterAdminViaOwnerArgs) (*types.Transaction, error) {
		return registryModuleOwnerCustom.RegisterAdminViaOwner(opts, args.Token)
	},
})
