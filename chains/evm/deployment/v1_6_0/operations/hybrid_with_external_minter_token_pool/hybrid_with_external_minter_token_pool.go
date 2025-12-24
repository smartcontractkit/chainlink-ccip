package hybrid_with_external_minter_token_pool

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/ccip-contract-examples/chains/evm/gobindings/generated/latest/hybrid_with_external_minter_token_pool"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "HybridWithExternalMinterTokenPool"
var TypeAndVersion = cldf_deployment.NewTypeAndVersion(ContractType, *Version)

var Version = utils.Version_1_6_0

type ConstructorArgs struct {
	Minter             common.Address   // The address of the external minter contract (token governor)
	Token              common.Address   // The token managed by this pool
	LocalTokenDecimals uint8            // The token decimals on the local chain
	Allowlist          []common.Address // List of addresses allowed to trigger lockOrBurn
	RmnProxy           common.Address   // The RMN proxy address
	Router             common.Address   // The router address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "hybrid_with_external_minter_token_pool:deploy",
	Version:          Version,
	Description:      "Deploys the HybridWithExternalMinterTokenPool contract",
	ContractMetadata: hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPoolMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPoolBin),
		},
	},
	Validate: func(args ConstructorArgs) error { return nil },
})
