package lombard_verifier

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/lombard_verifier"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"

	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "LombardVerifier"

var ResolverType cldf_deployment.ContractType = "LombardVerifierResolver"

var Version *semver.Version = semver.MustParse("1.7.0")

type DynamicConfig = lombard_verifier.LombardVerifierDynamicConfig

type ConstructorArgs struct {
	Bridge           common.Address
	StorageLocations []string
	DynamicConfig    DynamicConfig
	RMN              common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "lombard-verifier:deploy",
	Version:          Version,
	Description:      "Deploys the LombardVerifier contract",
	ContractMetadata: lombard_verifier.LombardVerifierMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(lombard_verifier.LombardVerifierBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
