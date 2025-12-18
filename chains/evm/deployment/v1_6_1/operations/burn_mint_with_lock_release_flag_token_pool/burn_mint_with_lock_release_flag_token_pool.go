package burn_mint_with_lock_release_flag_token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_1/burn_mint_with_lock_release_flag_token_pool"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "BurnMintWithLockReleaseFlagTokenPool"
var TypeAndVersion = cldf_deployment.NewTypeAndVersion(ContractType, *Version)

var Version *semver.Version = utils.Version_1_6_1

type ConstructorArgs struct {
	Token              common.Address   // The token managed by this pool
	LocalTokenDecimals uint8            // The token decimals on the local chain
	Allowlist          []common.Address // List of addresses allowed to trigger lockOrBurn
	RmnProxy           common.Address   // The RMN proxy address
	Router             common.Address   // The router address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "burn_mint_with_lock_release_flag_token_pool:deploy",
	Version:          Version,
	Description:      "Deploys the BurnMintWithLockReleaseFlagTokenPool contract",
	ContractMetadata: burn_mint_with_lock_release_flag_token_pool.BurnMintWithLockReleaseFlagTokenPoolMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(burn_mint_with_lock_release_flag_token_pool.BurnMintWithLockReleaseFlagTokenPoolBin),
		},
	},
	Validate: func(args ConstructorArgs) error { return nil },
})
