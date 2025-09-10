package ccv_ramp_proxy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/ccv_ramp_proxy"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CCVRampProxy"
var Version = semver.MustParse("1.7.0")

type SetRampArgs struct {
	RemoteChainSelector uint64
	Version             [32]byte
	RampAddress         common.Address
}

var Deploy = contract.NewDeploy(
	"ccv-ramp-proxy:deploy",
	Version,
	"Deploys the CCVRampProxy contract",
	ContractType,
	ccv_ramp_proxy.CCVRampProxyABI,
	func(any) error { return nil },
	contract.VMDeployers[any]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args any) (common.Address, *types.Transaction, error) {
			address, tx, _, err := ccv_ramp_proxy.DeployCCVRampProxy(opts, backend)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args any) (common.Address, error)
	},
)

var SetRamp = contract.NewWrite(
	"ccv-ramp-proxy:set-ramp",
	Version,
	"Sets the ramp address for a given remote chain selector and version on the CCVRampProxy",
	ContractType,
	ccv_ramp_proxy.CCVRampProxyABI,
	ccv_ramp_proxy.NewCCVRampProxy,
	contract.OnlyOwner,
	func(SetRampArgs) error { return nil },
	func(ccvRampProxy *ccv_ramp_proxy.CCVRampProxy, opts *bind.TransactOpts, args SetRampArgs) (*types.Transaction, error) {
		return ccvRampProxy.SetRamp(opts, args.RemoteChainSelector, args.Version, args.RampAddress)
	},
)
