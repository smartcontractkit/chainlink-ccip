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

type Subject = [16]byte

var ContractType cldf_deployment.ContractType = "RMNRemote"

type ConstructorArgs struct {
	LocalChainSelector uint64
	LegacyRMN          common.Address
}

type CurseArgs struct {
	Subject Subject
}

var Deploy = contract.NewDeploy(
	"rmn-remote:deploy",
	semver.MustParse("1.6.0"),
	"Deploys the RMNRemote contract",
	ContractType,
	rmn_remote.RMNRemoteABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := rmn_remote.DeployRMNRemote(opts, backend, args.LocalChainSelector, args.LegacyRMN)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var Curse = contract.NewWrite(
	"rmn-remote:curse",
	semver.MustParse("1.6.0"),
	"Applies a curse to an RMNRemote contract",
	ContractType,
	rmn_remote.RMNRemoteABI,
	rmn_remote.NewRMNRemote,
	contract.OnlyOwner,
	func(CurseArgs) error { return nil },
	func(rmnRemote *rmn_remote.RMNRemote, opts *bind.TransactOpts, args CurseArgs) (*types.Transaction, error) {
		return rmnRemote.Curse(opts, args.Subject)
	},
)

var Uncurse = contract.NewWrite(
	"rmn-remote:uncurse",
	semver.MustParse("1.6.0"),
	"Uncurses an existing curse on an RMNRemote contract",
	ContractType,
	rmn_remote.RMNRemoteABI,
	rmn_remote.NewRMNRemote,
	contract.OnlyOwner,
	func(CurseArgs) error { return nil },
	func(rmnRemote *rmn_remote.RMNRemote, opts *bind.TransactOpts, args CurseArgs) (*types.Transaction, error) {
		return rmnRemote.Uncurse(opts, args.Subject)
	},
)
