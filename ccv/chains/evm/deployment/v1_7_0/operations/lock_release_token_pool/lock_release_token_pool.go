package lock_release_token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/siloed_lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/siloed_usdc_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "LockReleaseTokenPool"

var SiloedContractType cldf_deployment.ContractType = "SiloedLockReleaseTokenPool"

var SiloedUSDCTokenPoolContractType cldf_deployment.ContractType = "SiloedUSDCTokenPool"

var Version = semver.MustParse("1.7.0")

var bytecodeByTypeAndVersion = map[string]contract.Bytecode{
	cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
		EVM: common.FromHex(lock_release_token_pool.LockReleaseTokenPoolBin),
	},
	cldf_deployment.NewTypeAndVersion(SiloedContractType, *Version).String(): {
		EVM: common.FromHex(siloed_lock_release_token_pool.SiloedLockReleaseTokenPoolBin),
	},
	cldf_deployment.NewTypeAndVersion(SiloedUSDCTokenPoolContractType, *Version).String(): {
		EVM: common.FromHex(siloed_usdc_token_pool.SiloedUSDCTokenPoolBin),
	},
}

func IsSupported(tokenPoolType deployment.ContractType, version *semver.Version) bool {
	if version == nil {
		return false
	}
	_, ok := bytecodeByTypeAndVersion[cldf_deployment.NewTypeAndVersion(tokenPoolType, *version).String()]
	return ok
}

type ConstructorArgs struct {
	Token              common.Address
	LocalTokenDecimals uint8
	AdvancedPoolHooks  common.Address
	RMNProxy           common.Address
	Router             common.Address
	LockBox            common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:                     "lock-release-token-pool:deploy",
	Version:                  Version,
	Description:              "Deploys the LockReleaseTokenPool contract",
	ContractMetadata:         lock_release_token_pool.LockReleaseTokenPoolMetaData,
	BytecodeByTypeAndVersion: bytecodeByTypeAndVersion,
	Validate:                 func(ConstructorArgs) error { return nil },
})
