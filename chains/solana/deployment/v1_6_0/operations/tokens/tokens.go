package tokens

import (
	"github.com/Masterminds/semver/v3"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "LINK"

type ConstructorArgs struct{}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "link:deploy",
	Version:          semver.MustParse("1.0.0"),
	Description:      "Deploys the LINK token contract",
	ContractMetadata: link_token.LinkTokenMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.0.0")).String(): {
			EVM: common.FromHex(link_token.LinkTokenBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
