package ccv_ramp_proxy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/ramp_proxy"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "RampProxy"

var SetRamp = contract.NewWrite(
	"ramp-proxy:set-ramp",
	semver.MustParse("1.7.0"),
	"Set the ramp address on the ramp proxy",
	ContractType,
	ramp_proxy.RampProxyABI,
	ramp_proxy.NewRampProxy,
	contract.OnlyOwner,
	func(common.Address) error { return nil },
	func(rampProxy *ramp_proxy.RampProxy, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return rampProxy.SetRamp(opts, args)
	},
)
