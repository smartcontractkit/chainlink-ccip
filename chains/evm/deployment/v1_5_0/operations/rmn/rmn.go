package rmn

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/rmn_contract"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
)

var ContractType cldf_deployment.ContractType = "RMN"

type ConstructorArgs struct {
	RMNConfig rmn_contract.RMNConfig
}

type CurseArgs struct {
	CurseID [16]byte
	Subject []fastcurse.Subject
}

type UncurseArgs struct {
	Requests []rmn_contract.RMNOwnerUnvoteToCurseRequest
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "rmn:deploy",
	Version:          semver.MustParse("1.5.0"),
	Description:      "Deploys the RMN contract",
	ContractMetadata: rmn_contract.RMNContractMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.5.0")).String(): {
			EVM: common.FromHex(rmn_contract.RMNContractBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var Curse = contract.NewWrite(contract.WriteParams[CurseArgs, *rmn_contract.RMNContract]{
	Name:            "rmn:curse",
	Version:         semver.MustParse("1.5.0"),
	Description:     "Applies a curse to an RMN contract",
	ContractType:    ContractType,
	ContractABI:     rmn_contract.RMNContractABI,
	NewContract:     rmn_contract.NewRMNContract,
	IsAllowedCaller: contract.OnlyOwner[*rmn_contract.RMNContract, CurseArgs],
	Validate:        func(CurseArgs) error { return nil },
	CallContract: func(rmn *rmn_contract.RMNContract, opts *bind.TransactOpts, args CurseArgs) (*types.Transaction, error) {
		return rmn.OwnerCurse(opts, args.CurseID, args.Subject)
	},
})

var Uncurse = contract.NewWrite(contract.WriteParams[UncurseArgs, *rmn_contract.RMNContract]{
	Name:            "rmn:uncurse",
	Version:         semver.MustParse("1.5.0"),
	Description:     "Uncurses an existing curse on an RMN contract",
	ContractType:    ContractType,
	ContractABI:     rmn_contract.RMNContractABI,
	NewContract:     rmn_contract.NewRMNContract,
	IsAllowedCaller: contract.OnlyOwner[*rmn_contract.RMNContract, UncurseArgs],
	Validate:        func(UncurseArgs) error { return nil },
	CallContract: func(rmn *rmn_contract.RMNContract, opts *bind.TransactOpts, args UncurseArgs) (*types.Transaction, error) {
		return rmn.OwnerUnvoteToCurse(opts, args.Requests)
	},
})

var IsCursed = contract.NewRead(contract.ReadParams[fastcurse.Subject, bool, *rmn_contract.RMNContract]{
	Name:         "rmn:is-cursed",
	Version:      semver.MustParse("1.5.0"),
	Description:  "Checks if a subject is cursed on an RMN contract",
	ContractType: ContractType,
	NewContract:  rmn_contract.NewRMNContract,
	CallContract: func(rmn *rmn_contract.RMNContract, opts *bind.CallOpts, args fastcurse.Subject) (bool, error) {
		return rmn.IsCursed(opts, args)
	},
})

var GetCurseProgress = contract.NewRead(contract.ReadParams[fastcurse.Subject, rmn_contract.GetCurseProgress, *rmn_contract.RMNContract]{
	Name:         "rmn:get-curse-progress",
	Version:      semver.MustParse("1.5.0"),
	Description:  "Gets the curse progress for a given subject on an RMN contract",
	ContractType: ContractType,
	NewContract:  rmn_contract.NewRMNContract,
	CallContract: func(rmn *rmn_contract.RMNContract, opts *bind.CallOpts, args fastcurse.Subject) (rmn_contract.GetCurseProgress, error) {
		return rmn.GetCurseProgress(opts, args)
	},
})

var GetConfigDetails = contract.NewRead(contract.ReadParams[any, rmn_contract.GetConfigDetails, *rmn_contract.RMNContract]{
	Name:         "rmn:get-config-details",
	Version:      semver.MustParse("1.5.0"),
	Description:  "Gets the configuration details of the RMN contract",
	ContractType: ContractType,
	NewContract:  rmn_contract.NewRMNContract,
	CallContract: func(rmn *rmn_contract.RMNContract, opts *bind.CallOpts, _ any) (rmn_contract.GetConfigDetails, error) {
		return rmn.GetConfigDetails(opts)
	},
})
