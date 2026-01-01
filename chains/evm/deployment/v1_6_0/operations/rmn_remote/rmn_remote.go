package rmn_remote

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_remote"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
)

var ContractType cldf_deployment.ContractType = "RMNRemote"
var Version *semver.Version = semver.MustParse("1.6.0")

type ConstructorArgs struct {
	LocalChainSelector uint64
	LegacyRMN          common.Address
}

type CurseArgs struct {
	Subject []fastcurse.Subject
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "rmn-remote:deploy",
	Version:          Version,
	Description:      "Deploys the RMNRemote contract",
	ContractMetadata: rmn_remote.RMNRemoteMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(rmn_remote.RMNRemoteBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var Curse = contract.NewWrite(contract.WriteParams[CurseArgs, *rmn_remote.RMNRemote]{
	Name:            "rmn-remote:curse",
	Version:         Version,
	Description:     "Applies a curse to an RMNRemote contract",
	ContractType:    ContractType,
	ContractABI:     rmn_remote.RMNRemoteABI,
	NewContract:     rmn_remote.NewRMNRemote,
	IsAllowedCaller: contract.OnlyOwner[*rmn_remote.RMNRemote, CurseArgs],
	Validate: func(rmnRemote *rmn_remote.RMNRemote, backend bind.ContractBackend, opts *bind.CallOpts, args CurseArgs) error {
		subjects, err := rmnRemote.GetCursedSubjects(opts)
		if err != nil {
			return fmt.Errorf("failed to get cursed subjects: %w", err)
		}
		for _, subject := range subjects {
			for _, argSubject := range args.Subject {
				if subject == argSubject {
					return fmt.Errorf("subject %s is already cursed", subject)
				}
			}
		}
		return nil
	},
	IsNoop: func(rmnRemote *rmn_remote.RMNRemote, opts *bind.CallOpts, args CurseArgs) (bool, error) {
		// No-ops not possible for this operation given the validation logic.
		return false, nil
	},
	CallContract: func(rmnRemote *rmn_remote.RMNRemote, opts *bind.TransactOpts, args CurseArgs) (*types.Transaction, error) {
		return rmnRemote.Curse0(opts, args.Subject)
	},
})

var Uncurse = contract.NewWrite(contract.WriteParams[CurseArgs, *rmn_remote.RMNRemote]{
	Name:            "rmn-remote:uncurse",
	Version:         Version,
	Description:     "Uncurses an existing curse on an RMNRemote contract",
	ContractType:    ContractType,
	ContractABI:     rmn_remote.RMNRemoteABI,
	NewContract:     rmn_remote.NewRMNRemote,
	IsAllowedCaller: contract.OnlyOwner[*rmn_remote.RMNRemote, CurseArgs],
	Validate: func(rmnRemote *rmn_remote.RMNRemote, backend bind.ContractBackend, opts *bind.CallOpts, args CurseArgs) error {
		subjects, err := rmnRemote.GetCursedSubjects(opts)
		if err != nil {
			return fmt.Errorf("failed to get cursed subjects: %w", err)
		}
		for _, argSubject := range args.Subject {
			found := false
			for _, subject := range subjects {
				if subject == argSubject {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("subject %s is not cursed", argSubject)
			}
		}
		return nil
	},
	IsNoop: func(rmnRemote *rmn_remote.RMNRemote, opts *bind.CallOpts, args CurseArgs) (bool, error) {
		// No-ops not possible for this operation given the validation logic.
		return false, nil
	},
	CallContract: func(rmnRemote *rmn_remote.RMNRemote, opts *bind.TransactOpts, args CurseArgs) (*types.Transaction, error) {
		return rmnRemote.Uncurse0(opts, args.Subject)
	},
})

var IsCursed = contract.NewRead(contract.ReadParams[fastcurse.Subject, bool, *rmn_remote.RMNRemote]{
	Name:         "rmn-remote:is-cursed",
	Version:      Version,
	Description:  "Checks if a subject is cursed on an RMNRemote contract",
	ContractType: ContractType,
	NewContract:  rmn_remote.NewRMNRemote,
	CallContract: func(rmn *rmn_remote.RMNRemote, opts *bind.CallOpts, args fastcurse.Subject) (bool, error) {
		return rmn.IsCursed(opts, args)
	},
})
