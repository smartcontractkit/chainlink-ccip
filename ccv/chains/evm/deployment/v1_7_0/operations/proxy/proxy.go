package proxy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "Proxy"

type AcceptOwnershipArgs struct {
	IsProposedOwner bool
}

var SetTarget = contract.NewWrite(contract.WriteParams[common.Address, *proxy.Proxy]{
	Name:            "proxy:set-target",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Set the target address on the proxy",
	ContractType:    ContractType,
	ContractABI:     proxy.ProxyABI,
	NewContract:     proxy.NewProxy,
	IsAllowedCaller: contract.OnlyOwner[*proxy.Proxy, common.Address],
	Validate:        func(common.Address) error { return nil },
	CallContract: func(proxy *proxy.Proxy, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return proxy.SetTarget(opts, args)
	},
})

var AcceptOwnership = contract.NewWrite(contract.WriteParams[AcceptOwnershipArgs, *proxy.Proxy]{
	Name:         "proxy:accept-ownership",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Accept ownership of the proxy",
	ContractType: ContractType,
	ContractABI:  proxy.ProxyABI,
	NewContract:  proxy.NewProxy,
	IsAllowedCaller: func(proxy *proxy.Proxy, opts *bind.CallOpts, caller common.Address, args AcceptOwnershipArgs) (bool, error) {
		return args.IsProposedOwner, nil
	},
	Validate: func(AcceptOwnershipArgs) error { return nil },
	CallContract: func(proxy *proxy.Proxy, opts *bind.TransactOpts, _ AcceptOwnershipArgs) (*types.Transaction, error) {
		return proxy.AcceptOwnership(opts)
	},
})
