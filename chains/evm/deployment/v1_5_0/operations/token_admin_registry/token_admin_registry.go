package token_admin_registry

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/token_admin_registry"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "TokenAdminRegistry"

type ConstructorArgs struct{}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "token-admin-registry:deploy",
	Version:          semver.MustParse("1.5.0"),
	Description:      "Deploys the TokenAdminRegistry contract",
	ContractType:     ContractType,
	ContractMetadata: token_admin_registry.TokenAdminRegistryMetaData,
	BytecodeByVersion: map[string]contract.Bytecode{
		semver.MustParse("1.5.0").String(): {EVM: common.FromHex(token_admin_registry.TokenAdminRegistryBin)},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
