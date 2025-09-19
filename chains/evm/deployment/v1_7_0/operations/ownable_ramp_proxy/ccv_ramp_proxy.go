package ccv_ramp_proxy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/ownable_ramp_proxy"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "OwnableRampProxy"

var SetRamp = contract.NewWrite(
	"ownable-ramp-proxy:set-ramp",
	semver.MustParse("1.7.0"),
	"Set the ramp address on the ownable ramp proxy",
	ContractType,
	ownable_ramp_proxy.OwnableRampProxyABI,
	ownable_ramp_proxy.NewOwnableRampProxy,
	contract.OnlyOwner,
	func(common.Address) error { return nil },
	func(rampProxy *ownable_ramp_proxy.OwnableRampProxy, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return rampProxy.SetRamp(opts, args)
	},
)
