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

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "ccv-proxy:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the CCVProxy contract",
	ContractMetadata: ccv_proxy.CCVProxyMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(ccv_proxy.CCVProxyBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[SetDynamicConfigArgs, *ccv_proxy.CCVProxy]{
	Name:            "ccv-proxy:set-dynamic-config",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Sets the dynamic configuration on the CCVProxy",
	ContractType:    ContractType,
	ContractABI:     ccv_proxy.CCVProxyABI,
	NewContract:     ccv_proxy.NewCCVProxy,
	IsAllowedCaller: contract.OnlyOwner[*ccv_proxy.CCVProxy, SetDynamicConfigArgs],
	Validate:        func(SetDynamicConfigArgs) error { return nil },
	CallContract: func(ccvProxy *ccv_proxy.CCVProxy, opts *bind.TransactOpts, args SetDynamicConfigArgs) (*types.Transaction, error) {
		return ccvProxy.SetDynamicConfig(opts, args.DynamicConfig)
	},
})

var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]DestChainConfigArgs, *ccv_proxy.CCVProxy]{
	Name:            "ccv-proxy:apply-dest-chain-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to destination chain configuration on the CCVProxy",
	ContractType:    ContractType,
	ContractABI:     ccv_proxy.CCVProxyABI,
	NewContract:     ccv_proxy.NewCCVProxy,
	IsAllowedCaller: contract.OnlyOwner[*ccv_proxy.CCVProxy, []DestChainConfigArgs],
	Validate:        func([]DestChainConfigArgs) error { return nil },
	CallContract: func(ccvProxy *ccv_proxy.CCVProxy, opts *bind.TransactOpts, args []DestChainConfigArgs) (*types.Transaction, error) {
		return ccvProxy.ApplyDestChainConfigUpdates(opts, args)
	},
})

var WithdrawFeeTokens = contract.NewWrite(contract.WriteParams[WithdrawFeeTokensArgs, *ccv_proxy.CCVProxy]{
	Name:            "ccv-proxy:withdraw-fee-tokens",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Withdraws fee tokens from the CCVProxy",
	ContractType:    ContractType,
	ContractABI:     ccv_proxy.CCVProxyABI,
	NewContract:     ccv_proxy.NewCCVProxy,
	IsAllowedCaller: contract.OnlyOwner[*ccv_proxy.CCVProxy, WithdrawFeeTokensArgs],
	Validate:        func(WithdrawFeeTokensArgs) error { return nil },
	CallContract: func(ccvProxy *ccv_proxy.CCVProxy, opts *bind.TransactOpts, args WithdrawFeeTokensArgs) (*types.Transaction, error) {
		return ccvProxy.WithdrawFeeTokens(opts, args.FeeTokens)
	},
})

var GetDestChainConfig = contract.NewRead(contract.ReadParams[uint64, DestChainConfig, *ccv_proxy.CCVProxy]{
	Name:         "ccv-proxy:get-dest-chain-config",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the destination chain configuration for a given destination chain selector",
	ContractType: ContractType,
	NewContract:  ccv_proxy.NewCCVProxy,
	CallContract: func(ccvProxy *ccv_proxy.CCVProxy, opts *bind.CallOpts, args uint64) (DestChainConfig, error) {
		return ccvProxy.GetDestChainConfig(opts, args)
	},
})
