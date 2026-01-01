package proxy

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "Proxy"

var Version = semver.MustParse("1.7.0")

type AcceptOwnershipArgs struct {
	IsProposedOwner bool
}

var SetTarget = contract.NewWrite(contract.WriteParams[common.Address, *proxy.Proxy]{
	Name:            "proxy:set-target",
	Version:         Version,
	Description:     "Set the target address on the proxy",
	ContractType:    ContractType,
	ContractABI:     proxy.ProxyABI,
	NewContract:     proxy.NewProxy,
	IsAllowedCaller: contract.OnlyOwner[*proxy.Proxy, common.Address],
	Validate: func(proxy *proxy.Proxy, backend bind.ContractBackend, opts *bind.CallOpts, args common.Address) error {
		if args == (common.Address{}) {
			return errors.New("target cannot be the zero address")
		}
		return nil
	},
	IsNoop: func(proxy *proxy.Proxy, opts *bind.CallOpts, args common.Address) (bool, error) {
		actualTarget, err := proxy.GetTarget(opts)
		if err != nil {
			return false, fmt.Errorf("failed to get target on proxy with address %s: %w", proxy.Address(), err)
		}
		return actualTarget == args, nil
	},
	CallContract: func(proxy *proxy.Proxy, opts *bind.TransactOpts, args common.Address) (*types.Transaction, error) {
		return proxy.SetTarget(opts, args)
	},
})

var AcceptOwnership = contract.NewWrite(contract.WriteParams[AcceptOwnershipArgs, *proxy.Proxy]{
	Name:         "proxy:accept-ownership",
	Version:      Version,
	Description:  "Accept ownership of the proxy",
	ContractType: ContractType,
	ContractABI:  proxy.ProxyABI,
	NewContract:  proxy.NewProxy,
	IsAllowedCaller: func(proxy *proxy.Proxy, opts *bind.CallOpts, caller common.Address, args AcceptOwnershipArgs) (bool, error) {
		return args.IsProposedOwner, nil
	},
	Validate: func(proxy *proxy.Proxy, backend bind.ContractBackend, opts *bind.CallOpts, args AcceptOwnershipArgs) error {
		return nil
	},
	IsNoop: func(proxy *proxy.Proxy, opts *bind.CallOpts, args AcceptOwnershipArgs) (bool, error) {
		return false, nil
	},
	CallContract: func(proxy *proxy.Proxy, opts *bind.TransactOpts, _ AcceptOwnershipArgs) (*types.Transaction, error) {
		return proxy.AcceptOwnership(opts)
	},
})

var GetTarget = contract.NewRead(contract.ReadParams[any, common.Address, *proxy.Proxy]{
	Name:         "proxy:get-target",
	Version:      Version,
	Description:  "Gets the target address on the proxy",
	ContractType: ContractType,
	NewContract:  proxy.NewProxy,
	CallContract: func(proxy *proxy.Proxy, opts *bind.CallOpts, args any) (common.Address, error) {
		return proxy.GetTarget(opts)
	},
})
