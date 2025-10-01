package link

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/gethwrappers/shared/generated/initial/link_token"
)

var ContractType cldf_deployment.ContractType = "LINK"

type ConstructorArgs struct{}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "link:deploy",
	Version:          semver.MustParse("1.0.0"),
	Description:      "Deploys the LINK token contract",
	ContractType:     ContractType,
	ContractMetadata: link_token.LinkTokenMetaData,
	BytecodeByVersion: map[string]contract.Bytecode{
		semver.MustParse("1.0.0").String(): {EVM: common.FromHex(link_token.LinkTokenBin)},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
