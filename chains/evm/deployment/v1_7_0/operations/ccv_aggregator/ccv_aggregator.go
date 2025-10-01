package ccv_aggregator

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/ccv_aggregator"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "CCVAggregator"

type StaticConfig = ccv_aggregator.CCVAggregatorStaticConfig

type ConstructorArgs struct {
	StaticConfig StaticConfig
}

type SourceChainConfigArgs = ccv_aggregator.CCVAggregatorSourceChainConfigArgs

type SourceChainConfig = ccv_aggregator.CCVAggregatorSourceChainConfig

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "ccv-aggregator:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the CCVAggregator contract",
	ContractType:     ContractType,
	ContractMetadata: ccv_aggregator.CCVAggregatorMetaData,
	BytecodeByVersion: map[string]contract.Bytecode{
		semver.MustParse("1.7.0").String(): {EVM: common.FromHex(ccv_aggregator.CCVAggregatorBin)},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplySourceChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]SourceChainConfigArgs, *ccv_aggregator.CCVAggregator]{
	Name:            "ccv-aggregator:apply-source-chain-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to source chain configurations on the CCVAggregator",
	ContractType:    ContractType,
	ContractABI:     ccv_aggregator.CCVAggregatorABI,
	NewContract:     ccv_aggregator.NewCCVAggregator,
	IsAllowedCaller: contract.OnlyOwner[*ccv_aggregator.CCVAggregator],
	Validate:        func([]SourceChainConfigArgs) error { return nil },
	CallContract: func(ccvAggregator *ccv_aggregator.CCVAggregator, opts *bind.TransactOpts, args []SourceChainConfigArgs) (*types.Transaction, error) {
		return ccvAggregator.ApplySourceChainConfigUpdates(opts, args)
	},
})

var GetSourceChainConfig = contract.NewRead(contract.ReadParams[uint64, SourceChainConfig, *ccv_aggregator.CCVAggregator]{
	Name:         "ccv-aggregator:get-source-chain-config",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the source chain configuration for a given source chain selector",
	ContractType: ContractType,
	NewContract:  ccv_aggregator.NewCCVAggregator,
	CallContract: func(ccvAggregator *ccv_aggregator.CCVAggregator, opts *bind.CallOpts, args uint64) (SourceChainConfig, error) {
		return ccvAggregator.GetSourceChainConfig(opts, args)
	},
})
