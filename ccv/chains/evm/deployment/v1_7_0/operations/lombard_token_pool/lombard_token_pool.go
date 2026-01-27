package lombard_token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/lombard_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "LombardTokenPool"

var Version = semver.MustParse("1.7.0")

type ConstructorArgs struct {
	Token             common.Address
	LombardVerifier   common.Address
	BridgeV2          common.Address
	AdvancedPoolHooks common.Address
	RMNProxy          common.Address
	Router            common.Address
	FallbackDecimals  uint8
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "lombard-token-pool:deploy",
	Version:          Version,
	Description:      "Deploys the LombardTokenPool contract",
	ContractMetadata: lombard_token_pool.LombardTokenPoolMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(lombard_token_pool.LombardTokenPoolBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})
