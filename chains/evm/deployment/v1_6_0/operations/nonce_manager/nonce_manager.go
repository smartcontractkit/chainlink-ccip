package nonce_manager

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/nonce_manager"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "NonceManager"

type ConstructorArgs struct {
	AuthorizedCallers []common.Address
}

type AuthorizedCallerArgs = nonce_manager.AuthorizedCallersAuthorizedCallerArgs

type PreviousRampsArgs = nonce_manager.NonceManagerPreviousRampsArgs

var Deploy = contract.NewDeploy(
	"nonce-manager:deploy",
	semver.MustParse("1.6.0"),
	"Deploys the NonceManager contract",
	ContractType,
	nonce_manager.NonceManagerABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := nonce_manager.DeployNonceManager(opts, backend, args.AuthorizedCallers)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var ApplyAuthorizedCallerUpdates = contract.NewWrite(
	"nonce-manager:apply-authorized-caller-updates",
	semver.MustParse("1.6.0"),
	"Applies updates to the list of authorized callers on the NonceManager",
	ContractType,
	nonce_manager.NonceManagerABI,
	nonce_manager.NewNonceManager,
	contract.OnlyOwner,
	func(AuthorizedCallerArgs) error { return nil },
	func(nonceManager *nonce_manager.NonceManager, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return nonceManager.ApplyAuthorizedCallerUpdates(opts, args)
	},
)

var ApplyPreviousRampUpdates = contract.NewWrite(
	"nonce-manager:apply-previous-ramp-updates",
	semver.MustParse("1.6.0"),
	"Applies updates to the list of previous ramps on the NonceManager",
	ContractType,
	nonce_manager.NonceManagerABI,
	nonce_manager.NewNonceManager,
	contract.OnlyOwner,
	func([]PreviousRampsArgs) error { return nil },
	func(nonceManager *nonce_manager.NonceManager, opts *bind.TransactOpts, args []PreviousRampsArgs) (*types.Transaction, error) {
		return nonceManager.ApplyPreviousRampsUpdates(opts, args)
	},
)
