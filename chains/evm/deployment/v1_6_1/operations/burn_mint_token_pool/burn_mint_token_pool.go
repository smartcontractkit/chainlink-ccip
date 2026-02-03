package burn_mint_token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/burn_from_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/burn_mint_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/burn_with_from_mint_token_pool"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var TypeAndVersion = cldf_deployment.NewTypeAndVersion(BurnMintContractType, *Version)

var BurnMintContractType cldf_deployment.ContractType = "BurnMintTokenPool"

var BurnWithFromMintContractType cldf_deployment.ContractType = "BurnWithFromMintTokenPool"

var BurnFromMintContractType cldf_deployment.ContractType = "BurnFromMintTokenPool"

var Version = semver.MustParse("1.6.1")

type ConstructorArgs struct {
	Token              common.Address
	LocalTokenDecimals uint8
	Allowlist          []common.Address
	RMNProxy           common.Address
	Router             common.Address
}

var bytecodeByTypeAndVersion = map[string]contract.Bytecode{
	cldf_deployment.NewTypeAndVersion(BurnMintContractType, *Version).String(): {
		EVM: common.FromHex(burn_mint_token_pool.BurnMintTokenPoolBin),
	},
	cldf_deployment.NewTypeAndVersion(BurnWithFromMintContractType, *Version).String(): {
		EVM: common.FromHex(burn_with_from_mint_token_pool.BurnWithFromMintTokenPoolBin),
	},
	cldf_deployment.NewTypeAndVersion(BurnFromMintContractType, *Version).String(): {
		EVM: common.FromHex(burn_from_mint_token_pool.BurnFromMintTokenPoolBin),
	},
}

func IsSupported(tokenPoolType deployment.ContractType, version *semver.Version) bool {
	if version == nil {
		return false
	}
	_, ok := bytecodeByTypeAndVersion[cldf_deployment.NewTypeAndVersion(tokenPoolType, *version).String()]
	return ok
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:                     "burn-mint-token-pool:deploy",
	Version:                  Version,
	Description:              "Deploys the various BurnMintTokenPool contracts",
	ContractMetadata:         burn_mint_token_pool.BurnMintTokenPoolMetaData,
	BytecodeByTypeAndVersion: bytecodeByTypeAndVersion,
	Validate:                 func(ConstructorArgs) error { return nil },
})
