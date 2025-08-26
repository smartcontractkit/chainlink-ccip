package rmn_remote

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/operations/call"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/operations/deployment"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/rmn_remote"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "RMNRemote"

type ConstructorArgs struct {
	LocalChainSelector uint64
	LegacyRMN          common.Address
}

type CurseArgs struct {
	Subject [16]byte
}

var Deploy = deployment.New(
	"rmn-remote:deploy",
	semver.MustParse("1.6.0"),
	"Deploys the RMNRemote contract",
	ContractType,
	deployment.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := rmn_remote.DeployRMNRemote(opts, backend, args.LocalChainSelector, args.LegacyRMN)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var Curse = call.NewWrite(
	"rmn-remote:curse",
	semver.MustParse("1.6.0"),
	"Applies a curse to an RMNRemote contract",
	ContractType,
	rmn_remote.NewRMNRemote,
	func(rmnRemote *rmn_remote.RMNRemote, opts *bind.TransactOpts, args CurseArgs) (*types.Transaction, error) {
		return rmnRemote.Curse(opts, args.Subject)
	},
)

var Uncurse = call.NewWrite(
	"rmn-remote:uncurse",
	semver.MustParse("1.6.0"),
	"Uncurses an existing curse on an RMNRemote contract",
	ContractType,
	rmn_remote.NewRMNRemote,
	func(rmnRemote *rmn_remote.RMNRemote, opts *bind.TransactOpts, args CurseArgs) (*types.Transaction, error) {
		return rmnRemote.Uncurse(opts, args.Subject)
	},
)
