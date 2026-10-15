package lock_release_token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_1/lock_release_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "LockReleaseTokenPool"
var TypeAndVersion = cldf_deployment.NewTypeAndVersion(ContractType, *Version)

var Version *semver.Version = utils.Version_1_5_1

type ConstructorArgs struct {
	Token              common.Address   // The token managed by this pool
	LocalTokenDecimals uint8            // The token decimals on the local chain
	Allowlist          []common.Address // List of addresses allowed to trigger lockOrBurn
	RmnProxy           common.Address   // The RMN proxy address
	AcceptLiquidity    bool             // Whether the pool should accept liquidity
	Router             common.Address   // The router address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "lock_release_token_pool:deploy",
	Version:          Version,
	Description:      "Deploys the LockReleaseTokenPool contract",
	ContractMetadata: lock_release_token_pool.LockReleaseTokenPoolMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(lock_release_token_pool.LockReleaseTokenPoolBin),
		},
	},
	Validate: func(args ConstructorArgs) error { return nil },
})

var GetRebalancer = contract.NewRead(contract.ReadParams[struct{}, common.Address, *lock_release_token_pool.LockReleaseTokenPool]{
	Name:         "lock_release_token_pool:get-rebalancer",
	Version:      Version,
	Description:  "Gets the rebalancer of the LockReleaseTokenPool 1.5.1 contract",
	ContractType: ContractType,
	NewContract:  lock_release_token_pool.NewLockReleaseTokenPool,
	CallContract: func(c *lock_release_token_pool.LockReleaseTokenPool, opts *bind.CallOpts, args struct{}) (common.Address, error) {
		return c.GetRebalancer(opts)
	},
})

var SetRebalancer = contract.NewWrite(contract.WriteParams[common.Address, *lock_release_token_pool.LockReleaseTokenPool]{
	Name:            "lock_release_token_pool:set-rebalancer",
	Version:         Version,
	Description:     "Sets the rebalancer of the LockReleaseTokenPool 1.5.1 contract",
	ContractType:    ContractType,
	ContractABI:     lock_release_token_pool.LockReleaseTokenPoolABI,
	NewContract:     lock_release_token_pool.NewLockReleaseTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*lock_release_token_pool.LockReleaseTokenPool, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(c *lock_release_token_pool.LockReleaseTokenPool, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return c.SetRebalancer(opts, args)
	},
})
