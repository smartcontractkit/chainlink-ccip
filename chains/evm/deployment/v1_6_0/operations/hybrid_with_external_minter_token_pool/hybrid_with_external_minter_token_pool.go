package hybrid_with_external_minter_token_pool

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

type AddRemotePoolArgs struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
}

type RemoveRemotePoolArgs struct {
	RemoteChainSelector uint64
	RemotePoolAddress   []byte
}

type UpdateGroupsArgs struct {
	GroupUpdates []hybrid_with_external_minter_token_pool.HybridTokenPoolAbstractGroupUpdate
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

var GetGroup = contract.NewRead(contract.ReadParams[uint64, uint8, *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool]{
	Name:         "hybrid_with_external_minter_token_pool:get-group",
	Version:      Version,
	Description:  "Gets the group assigned to a remote chain selector on HybridWithExternalMinterTokenPool",
	ContractType: ContractType,
	NewContract:  hybrid_with_external_minter_token_pool.NewHybridWithExternalMinterTokenPool,
	CallContract: func(c *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool, opts *bind.CallOpts, remoteChainSelector uint64) (uint8, error) {
		return c.GetGroup(opts, remoteChainSelector)
	},
})

var GetRemotePools = contract.NewRead(contract.ReadParams[uint64, [][]byte, *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool]{
	Name:         "hybrid_with_external_minter_token_pool:get-remote-pools",
	Version:      Version,
	Description:  "Gets the registered remote pool addresses for a chain selector on HybridWithExternalMinterTokenPool",
	ContractType: ContractType,
	NewContract:  hybrid_with_external_minter_token_pool.NewHybridWithExternalMinterTokenPool,
	CallContract: func(c *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool, opts *bind.CallOpts, remoteChainSelector uint64) ([][]byte, error) {
		return c.GetRemotePools(opts, remoteChainSelector)
	},
})

var GetLockedTokens = contract.NewRead(contract.ReadParams[struct{}, *big.Int, *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool]{
	Name:         "hybrid_with_external_minter_token_pool:get-locked-tokens",
	Version:      Version,
	Description:  "Gets total locked token accounting from HybridWithExternalMinterTokenPool",
	ContractType: ContractType,
	NewContract:  hybrid_with_external_minter_token_pool.NewHybridWithExternalMinterTokenPool,
	CallContract: func(c *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool, opts *bind.CallOpts, _ struct{}) (*big.Int, error) {
		return c.GetLockedTokens(opts)
	},
})

var AddRemotePool = contract.NewWrite(contract.WriteParams[AddRemotePoolArgs, *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool]{
	Name:            "hybrid_with_external_minter_token_pool:add-remote-pool",
	Version:         Version,
	Description:     "Adds a remote pool for a chain selector on HybridWithExternalMinterTokenPool",
	ContractType:    ContractType,
	ContractABI:     hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPoolABI,
	NewContract:     hybrid_with_external_minter_token_pool.NewHybridWithExternalMinterTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool, AddRemotePoolArgs],
	Validate:        func(AddRemotePoolArgs) error { return nil },
	CallContract: func(c *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool, opts *bind.TransactOpts, args AddRemotePoolArgs) (*types.Transaction, error) {
		return c.AddRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var RemoveRemotePool = contract.NewWrite(contract.WriteParams[RemoveRemotePoolArgs, *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool]{
	Name:            "hybrid_with_external_minter_token_pool:remove-remote-pool",
	Version:         Version,
	Description:     "Removes a remote pool for a chain selector on HybridWithExternalMinterTokenPool",
	ContractType:    ContractType,
	ContractABI:     hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPoolABI,
	NewContract:     hybrid_with_external_minter_token_pool.NewHybridWithExternalMinterTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool, RemoveRemotePoolArgs],
	Validate:        func(RemoveRemotePoolArgs) error { return nil },
	CallContract: func(c *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool, opts *bind.TransactOpts, args RemoveRemotePoolArgs) (*types.Transaction, error) {
		return c.RemoveRemotePool(opts, args.RemoteChainSelector, args.RemotePoolAddress)
	},
})

var UpdateGroups = contract.NewWrite(contract.WriteParams[UpdateGroupsArgs, *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool]{
	Name:            "hybrid_with_external_minter_token_pool:update-groups",
	Version:         Version,
	Description:     "Updates remote chain groups on HybridWithExternalMinterTokenPool",
	ContractType:    ContractType,
	ContractABI:     hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPoolABI,
	NewContract:     hybrid_with_external_minter_token_pool.NewHybridWithExternalMinterTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool, UpdateGroupsArgs],
	Validate:        func(UpdateGroupsArgs) error { return nil },
	CallContract: func(c *hybrid_with_external_minter_token_pool.HybridWithExternalMinterTokenPool, opts *bind.TransactOpts, args UpdateGroupsArgs) (*types.Transaction, error) {
		return c.UpdateGroups(opts, args.GroupUpdates)
	},
})
