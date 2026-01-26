package rmn_remote

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_remote"
)

var ContractType cldf_deployment.ContractType = "RMNRemote"
var Version *semver.Version = semver.MustParse("1.6.0")

type ConstructorArgs struct {
	LocalChainSelector uint64
	LegacyRMN          common.Address
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "rmn-remote:deploy",
	Version:          semver.MustParse("1.6.0"),
	Description:      "Deploys the RMNRemote contract",
	ContractMetadata: rmn_remote.RMNRemoteMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.6.0")).String(): {
			EVM: common.FromHex(rmn_remote.RMNRemoteBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

type CurseArgs struct {
	Subjects [][16]byte
}

var Curse = contract.NewWrite(contract.WriteParams[CurseArgs, *rmn_remote.RMNRemote]{
	Name:            "rmn-remote:curse",
	Version:         semver.MustParse("1.6.0"),
	Description:     "Calls curse on the contract",
	ContractType:    ContractType,
	ContractABI:     rmn_remote.RMNRemoteABI,
	NewContract:     rmn_remote.NewRMNRemote,
	IsAllowedCaller: contract.OnlyOwner[*rmn_remote.RMNRemote, CurseArgs],
	Validate:        func(CurseArgs) error { return nil },
	CallContract: func(rMNRemote *rmn_remote.RMNRemote, opts *bind.TransactOpts, args CurseArgs) (*types.Transaction, error) {
		return rMNRemote.Curse0(opts, args.Subjects)
	},
})

type UncurseArgs struct {
	Subjects [][16]byte
}

var Uncurse = contract.NewWrite(contract.WriteParams[UncurseArgs, *rmn_remote.RMNRemote]{
	Name:            "rmn-remote:uncurse",
	Version:         semver.MustParse("1.6.0"),
	Description:     "Calls uncurse on the contract",
	ContractType:    ContractType,
	ContractABI:     rmn_remote.RMNRemoteABI,
	NewContract:     rmn_remote.NewRMNRemote,
	IsAllowedCaller: contract.OnlyOwner[*rmn_remote.RMNRemote, UncurseArgs],
	Validate:        func(UncurseArgs) error { return nil },
	CallContract: func(rMNRemote *rmn_remote.RMNRemote, opts *bind.TransactOpts, args UncurseArgs) (*types.Transaction, error) {
		return rMNRemote.Uncurse0(opts, args.Subjects)
	},
})

var IsCursed = contract.NewRead(contract.ReadParams[[16]byte, bool, *rmn_remote.RMNRemote]{
	Name:         "rmn-remote:is-cursed",
	Version:      semver.MustParse("1.6.0"),
	Description:  "Calls isCursed on the contract",
	ContractType: ContractType,
	NewContract:  rmn_remote.NewRMNRemote,
	CallContract: func(rMNRemote *rmn_remote.RMNRemote, opts *bind.CallOpts, args [16]byte) (bool, error) {
		return rMNRemote.IsCursed(opts, args)
	},
})
