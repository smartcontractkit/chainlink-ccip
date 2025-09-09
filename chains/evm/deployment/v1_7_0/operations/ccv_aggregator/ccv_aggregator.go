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

type ConstructorArgs = ccv_aggregator.CCVAggregatorStaticConfig

type SourceChainConfigArgs = ccv_aggregator.CCVAggregatorSourceChainConfigArgs

type SourceChainConfig = ccv_aggregator.CCVAggregatorSourceChainConfig

var Deploy = contract.NewDeploy(
	"ccv-aggregator:deploy",
	semver.MustParse("1.7.0"),
	"Deploys the CCVAggregator contract",
	ContractType,
	ccv_aggregator.CCVAggregatorABI,
	func(ConstructorArgs) error { return nil },
	contract.VMDeployers[ConstructorArgs]{
		DeployEVM: func(opts *bind.TransactOpts, backend bind.ContractBackend, args ConstructorArgs) (common.Address, *types.Transaction, error) {
			address, tx, _, err := ccv_aggregator.DeployCCVAggregator(opts, backend, args)
			return address, tx, err
		},
		// DeployZksyncVM: func(opts *accounts.TransactOpts, client *clients.Client, wallet *accounts.Wallet, backend bind.ContractBackend, args ConstructorArgs) (common.Address, error)
	},
)

var ApplySourceChainConfigUpdates = contract.NewWrite(
	"ccv-aggregator:apply-source-chain-config-updates",
	semver.MustParse("1.7.0"),
	"Applies updates to source chain configurations on the CCVAggregator",
	ContractType,
	ccv_aggregator.CCVAggregatorABI,
	ccv_aggregator.NewCCVAggregator,
	contract.OnlyOwner,
	func([]SourceChainConfigArgs) error { return nil },
	func(ccvAggregator *ccv_aggregator.CCVAggregator, opts *bind.TransactOpts, args []SourceChainConfigArgs) (*types.Transaction, error) {
		return ccvAggregator.ApplySourceChainConfigUpdates(opts, args)
	},
)

var GetSourceChainConfig = contract.NewRead(
	"ccv-aggregator:get-source-chain-config",
	semver.MustParse("1.7.0"),
	"Gets the source chain configuration for a given source chain selector",
	ContractType,
	ccv_aggregator.NewCCVAggregator,
	func(ccvAggregator *ccv_aggregator.CCVAggregator, opts *bind.CallOpts, sourceChainSelector uint64) (SourceChainConfig, error) {
		return ccvAggregator.GetSourceChainConfig(opts, sourceChainSelector)
	},
)
