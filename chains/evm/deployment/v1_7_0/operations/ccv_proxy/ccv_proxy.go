package ccv_proxy

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/ccv_proxy"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CCVProxy"

type StaticConfig = ccv_proxy.CCVProxyStaticConfig

type DynamicConfig = ccv_proxy.CCVProxyDynamicConfig

type ConstructorArgs struct {
	StaticConfig  StaticConfig
	DynamicConfig DynamicConfig
}

type SetDynamicConfigArgs struct {
	DynamicConfig DynamicConfig
}

type DestChainConfigArgs = ccv_proxy.CCVProxyDestChainConfigArgs

type WithdrawFeeTokensArgs struct {
	FeeTokens []common.Address
}

type DestChainConfig = ccv_proxy.CCVProxyDestChainConfig

var Deploy = contract.NewDeploy(
	"ccv-proxy:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the CCVProxy contract",
	ContractType,
	ccv_proxy.CCVProxyABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := ccv_proxy.DeployCCVProxy(opts, backend, args.StaticConfig, args.DynamicConfig)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, deployInput ConstructorArgs) (common.Address, error)
	},
)

var SetDestChainConfig = contract.NewWrite(
	"ccv-proxy:set-dynamic-config",
	semver.MustParse("1.7.0"),
	"Sets the dynamic configuration on the CCVProxy",
	ContractType,
	ccv_proxy.CCVProxyABI,
	ccv_proxy.NewCCVProxy,
	contract.OnlyOwner,
	func(SetDynamicConfigArgs) error { return nil },
	func(ccvProxy *ccv_proxy.CCVProxy, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return ccvProxy.SetDynamicConfig(opts, args.DynamicConfig)
	},
)

var ApplyDestChainConfigUpdates = contract.NewWrite(
	"ccv-proxy:apply-dest-chain-config-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to destination chain configuration on the CCVProxy",
	ContractType,
	ccv_proxy.CCVProxyABI,
	ccv_proxy.NewCCVProxy,
	contract.OnlyOwner,
	func([]DestChainConfigArgs) error { return nil },
	func(ccvProxy *ccv_proxy.CCVProxy, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return ccvProxy.ApplyDestChainConfigUpdates(opts, args)
	},
)

var WithdrawFeeTokens = contract.NewWrite(
	"ccv-proxy:withdraw-fee-tokens",
	semver.MustParse("1.7.0"),
	"Withdraws fee tokens from the CCVProxy",
	ContractType,
	ccv_proxy.CCVProxyABI,
	ccv_proxy.NewCCVProxy,
	contract.OnlyOwner,
	func(WithdrawFeeTokensArgs) error { return nil },
	func(ccvProxy *ccv_proxy.CCVProxy, opts *bind.TransactOpts, args WithdrawFeeTokensArgs) (*types.Transaction, error) {
		return ccvProxy.WithdrawFeeTokens(opts, args.FeeTokens)
	},
)

var GetDestChainConfig = contract.NewRead(
	"ccv-proxy:get-dest-chain-config",
	semver.MustParse("1.7.0"),
	"Gets the destination chain configuration for a given destination chain selector",
	ContractType,
	ccv_proxy.NewCCVProxy,
	func(ccvProxy *ccv_proxy.CCVProxy, opts *bind.CallOpts, destChainSelector uint64) (DestChainConfig, error) {
		return ccvProxy.GetDestChainConfig(opts, destChainSelector)
	},
)
