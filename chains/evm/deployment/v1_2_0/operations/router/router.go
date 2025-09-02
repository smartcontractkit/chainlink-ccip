package router

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/call"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/deployment"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_2_0/router"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "Router"

type ConstructorArgs struct {
	WrappedNative common.Address
	RMNProxy      common.Address
}

type OnRamp = router.RouterOnRamp

type OffRamp = router.RouterOffRamp

type ApplyRampsUpdatesArgs struct {
	OnRampUpdates  []OnRamp
	OffRampRemoves []OffRamp
	OffRampAdds    []OffRamp
}

var Deploy = deployment.New(
	"router:deploy",
	semver.MustParse("1.2.0"),
	"Deploys the Router contract",
	ContractType,
	router.RouterABI,
	func(ConstructorArgs) error { return nil },
	deployment.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := router.DeployRouter(opts, backend, args.WrappedNative, args.RMNProxy)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var ApplyRampUpdates = call.NewWrite(
	"router:apply-ramp-updates",
	semver.MustParse("1.2.0"),
	"Applies ramp updates to the Router",
	ContractType,
	router.RouterABI,
	router.NewRouter,
	call.OnlyOwner,
	func(ApplyRampsUpdatesArgs) error { return nil },
	func(router *router.Router, opts *bind.TransactOpts, args ApplyRampsUpdatesArgs) (*types.Transaction, error) {
		return router.ApplyRampUpdates(opts, args.OnRampUpdates, args.OffRampRemoves, args.OffRampAdds)
	},
)

var GetOffRamps = call.NewRead(
	"router:get-off-ramps",
	semver.MustParse("1.2.0"),
	"Gets all off ramps on the router",
	ContractType,
	router.NewRouter,
	func(router *router.Router, opts *bind.CallOpts, args any) ([]OffRamp, error) {
		return router.GetOffRamps(opts)
	},
)

var GetOnRamp = call.NewRead(
	"router:get-on-ramp",
	semver.MustParse("1.2.0"),
	"Gets the on ramp for a given destination chain selector",
	ContractType,
	router.NewRouter,
	func(router *router.Router, opts *bind.CallOpts, destChainSelector uint64) (common.Address, error) {
		return router.GetOnRamp(opts, destChainSelector)
	},
)
