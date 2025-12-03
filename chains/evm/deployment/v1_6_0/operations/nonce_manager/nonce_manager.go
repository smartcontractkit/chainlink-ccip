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
var Version *semver.Version = semver.MustParse("1.6.0")

type ConstructorArgs struct {
	AuthorizedCallers []common.Address
}

type AuthorizedCallerArgs = nonce_manager.AuthorizedCallersAuthorizedCallerArgs

type PreviousRampsArgs = nonce_manager.NonceManagerPreviousRampsArgs

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "nonce-manager:deploy",
	Version:          semver.MustParse("1.6.0"),
	Description:      "Deploys the NonceManager contract",
	ContractMetadata: nonce_manager.NonceManagerMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.6.0")).String(): {
			EVM: common.FromHex(nonce_manager.NonceManagerBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[AuthorizedCallerArgs, *nonce_manager.NonceManager]{
	Name:            "nonce-manager:apply-authorized-caller-updates",
	Version:         semver.MustParse("1.6.0"),
	Description:     "Applies updates to the list of authorized callers on the NonceManager",
	ContractType:    ContractType,
	ContractABI:     nonce_manager.NonceManagerABI,
	NewContract:     nonce_manager.NewNonceManager,
	IsAllowedCaller: contract.OnlyOwner[*nonce_manager.NonceManager, AuthorizedCallerArgs],
	Validate:        func(AuthorizedCallerArgs) error { return nil },
	CallContract: func(nonceManager *nonce_manager.NonceManager, opts *bind.TransactOpts, args AuthorizedCallerArgs) (*types.Transaction, error) {
		return nonceManager.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

var ApplyPreviousRampUpdates = contract.NewWrite(contract.WriteParams[[]PreviousRampsArgs, *nonce_manager.NonceManager]{
	Name:            "nonce-manager:apply-previous-ramp-updates",
	Version:         semver.MustParse("1.6.0"),
	Description:     "Applies updates to the list of previous ramps on the NonceManager",
	ContractType:    ContractType,
	ContractABI:     nonce_manager.NonceManagerABI,
	NewContract:     nonce_manager.NewNonceManager,
	IsAllowedCaller: contract.OnlyOwner[*nonce_manager.NonceManager, []PreviousRampsArgs],
	Validate:        func([]PreviousRampsArgs) error { return nil },
	CallContract: func(nonceManager *nonce_manager.NonceManager, opts *bind.TransactOpts, args []PreviousRampsArgs) (*types.Transaction, error) {
		return nonceManager.ApplyPreviousRampsUpdates(opts, args)
	},
})
