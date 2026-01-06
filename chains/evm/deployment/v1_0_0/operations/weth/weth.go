package weth

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/weth9"
)

var ContractType cldf_deployment.ContractType = "WETH9"

var Version *semver.Version = semver.MustParse("1.0.0")

type ConstructorArgs struct{}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "weth:deploy",
	Version:          Version,
	Description:      "Deploys the WETH9 contract",
	ContractMetadata: weth9.WETH9MetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(weth9.WETH9Bin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
