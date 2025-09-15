package ccv_ramp_proxy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/ownable_ccv_ramp_proxy"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-evm/pkg/utils"
)

var ContractType cldf_deployment.ContractType = "CCVRampProxy"

type SetRampArgs = ownable_ccv_ramp_proxy.CCVRampProxySetRampsArgs

type GetRampArgs struct {
	RemoteChainSelector uint64
	Version             [32]byte
}

var V1Ramp = utils.MustHash("CCVRamp_V1")

var SetRamp = contract.NewWrite(
	"ccv-ramp-proxy:set-ramps",
	semver.MustParse("1.7.0"),
	"Sets multiple ramp addresses on the CCVRampProxy, each tied to a remote chain selector and CCVRamp version",
	ContractType,
	ownable_ccv_ramp_proxy.OwnableCCVRampProxyABI,
	ownable_ccv_ramp_proxy.NewOwnableCCVRampProxy,
	contract.OnlyOwner,
	func([]SetRampArgs) error { return nil },
	func(ccvRampProxy *ownable_ccv_ramp_proxy.OwnableCCVRampProxy, opts *bind.TransactOpts, args []SetRampArgs) (*types.Transaction, error) {
		return ccvRampProxy.SetRamps(opts, args)
	},
)

var GetRamp = contract.NewRead(
	"ccv-ramp-proxy:get-ramp",
	semver.MustParse("1.7.0"),
	"Gets the ramp address for a given remote chain selector and CCVRamp version",
	ContractType,
	ownable_ccv_ramp_proxy.NewOwnableCCVRampProxy,
	func(ccvRampProxy *ownable_ccv_ramp_proxy.OwnableCCVRampProxy, opts *bind.CallOpts, args GetRampArgs) (common.Address, error) {
		return ccvRampProxy.GetRamp(opts, args.RemoteChainSelector, args.Version)
	},
)
