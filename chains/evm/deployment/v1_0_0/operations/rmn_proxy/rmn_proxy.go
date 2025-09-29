package rmn_proxy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_0_0/rmn_proxy_contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "RMNProxy"

type ConstructorArgs struct {
	RMN common.Address
}

type SetRMNArgs = struct {
	RMN common.Address
}

var Deploy = contract.NewDeploy(
	"rmn_proxy:deploy",
	semver.MustParse("1.0.0"),
	"Deploys the RMNProxy contract",
	ContractType,
	rmn_proxy_contract.RMNProxyABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := rmn_proxy_contract.DeployRMNProxy(opts, backend, args.RMN)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var SetRMN = contract.NewWrite(
	"rmn_proxy:set-rmn",
	semver.MustParse("1.0.0"),
	"Sets the RMN address on the RMNProxy",
	ContractType,
	rmn_proxy_contract.RMNProxyABI,
	rmn_proxy_contract.NewRMNProxy,
	contract.OnlyOwner,
	func(SetRMNArgs) error { return nil },
	func(rmnProxy *rmn_proxy_contract.RMNProxy, opts *bind.TransactOpts, args SetRMNArgs) (*types.Transaction, error) {
		return rmnProxy.SetARM(opts, args.RMN)
	},
)
