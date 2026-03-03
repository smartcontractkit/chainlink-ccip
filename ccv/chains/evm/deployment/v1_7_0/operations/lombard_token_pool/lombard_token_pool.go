package lombard_token_pool

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/lombard_token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "LombardTokenPool"

var Version = semver.MustParse("1.7.0")

type ConstructorArgs struct {
	Token             common.Address
	Verifier          common.Address
	Bridge            common.Address
	Adapter           common.Address
	AdvancedPoolHooks common.Address
	RMNProxy          common.Address
	Router            common.Address
	FallbackDecimals  uint8
}

type SetPathArgs struct {
	RemoteChainSelector uint64
	LChainID            [32]byte
	AllowedCaller       [32]byte
	RemoteAdapter       [32]byte
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

var SetPath = contract.NewWrite(contract.WriteParams[SetPathArgs, *lombard_token_pool.LombardTokenPool]{
	Name:            "lombard-token-pool:set-path",
	Version:         Version,
	Description:     "Sets path configuration on the LombardTokenPool",
	ContractType:    ContractType,
	ContractABI:     lombard_token_pool.LombardTokenPoolABI,
	NewContract:     lombard_token_pool.NewLombardTokenPool,
	IsAllowedCaller: contract.OnlyOwner[*lombard_token_pool.LombardTokenPool, SetPathArgs],
	Validate:        func(SetPathArgs) error { return nil },
	CallContract: func(pool *lombard_token_pool.LombardTokenPool, opts *bind.TransactOpts, args SetPathArgs) (*types.Transaction, error) {
		return pool.SetPath(opts, args.RemoteChainSelector, args.LChainID, args.AllowedCaller, args.RemoteAdapter)
	},
})

var GetPath = contract.NewRead(contract.ReadParams[uint64, lombard_token_pool.LombardTokenPoolPath, *lombard_token_pool.LombardTokenPool]{
	Name:         "lombard-token-pool:get-path",
	Version:      Version,
	Description:  "Gets path configuration for a remote chain on the LombardTokenPool",
	ContractType: ContractType,
	NewContract:  lombard_token_pool.NewLombardTokenPool,
	CallContract: func(pool *lombard_token_pool.LombardTokenPool, opts *bind.CallOpts, remoteChainSelector uint64) (lombard_token_pool.LombardTokenPoolPath, error) {
		return pool.GetPath(opts, remoteChainSelector)
	},
})
