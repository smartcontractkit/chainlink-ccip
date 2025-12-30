package registry_module_owner_custom

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/registry_module_owner_custom"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "RegistryModuleOwnerCustom"
var Version *semver.Version = semver.MustParse("1.6.0")

type ConstructorArgs struct {
	TokenAdminRegistry common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "registry-module-owner-custom:deploy",
	Version:          semver.MustParse("1.6.0"),
	Description:      "Deploys the RegistryModuleOwnerCustom contract",
	ContractMetadata: registry_module_owner_custom.RegistryModuleOwnerCustomMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(registry_module_owner_custom.RegistryModuleOwnerCustomBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
