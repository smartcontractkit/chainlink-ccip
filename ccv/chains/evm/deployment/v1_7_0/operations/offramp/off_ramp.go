package offramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/gobindings/generated/latest/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "OffRamp"

var Version = semver.MustParse("1.7.0")

type StaticConfig = offramp.OffRampStaticConfig

type ConstructorArgs struct {
	StaticConfig StaticConfig
}

type SourceChainConfigArgs = offramp.OffRampSourceChainConfigArgs

type SourceChainConfig = offramp.OffRampSourceChainConfig

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "off-ramp:deploy",
	Version:          Version,
	Description:      "Deploys the OffRamp contract",
	ContractMetadata: offramp.OffRampMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(offramp.OffRampBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplySourceChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]SourceChainConfigArgs, *offramp.OffRamp]{
	Name:            "off-ramp:apply-source-chain-config-updates",
	Version:         Version,
	Description:     "Applies updates to source chain configurations on the OffRamp",
	ContractType:    ContractType,
	ContractABI:     offramp.OffRampABI,
	NewContract:     offramp.NewOffRamp,
	IsAllowedCaller: contract.OnlyOwner[*offramp.OffRamp, []SourceChainConfigArgs],
	Validate:        func([]SourceChainConfigArgs) error { return nil },
	CallContract: func(offRamp *offramp.OffRamp, opts *bind.TransactOpts, args []SourceChainConfigArgs) (*types.Transaction, error) {
		return offRamp.ApplySourceChainConfigUpdates(opts, args)
	},
})

var GetSourceChainConfig = contract.NewRead(contract.ReadParams[uint64, SourceChainConfig, *offramp.OffRamp]{
	Name:         "off-ramp:get-source-chain-config",
	Version:      Version,
	Description:  "Gets the source chain configuration for a given source chain selector",
	ContractType: ContractType,
	NewContract:  offramp.NewOffRamp,
	CallContract: func(offRamp *offramp.OffRamp, opts *bind.CallOpts, args uint64) (SourceChainConfig, error) {
		return offRamp.GetSourceChainConfig(opts, args)
	},
})

type GetAllSourceChainConfigsResult struct {
	SourceChainSelectors []uint64
	SourceChainConfigs   []SourceChainConfig
}

var GetAllSourceChainConfigs = contract.NewRead(contract.ReadParams[any, GetAllSourceChainConfigsResult, *offramp.OffRamp]{
	Name:         "off-ramp:get-all-source-chain-configs",
	Version:      Version,
	Description:  "Gets all source chain configurations from the OffRamp",
	ContractType: ContractType,
	NewContract:  offramp.NewOffRamp,
	CallContract: func(offRamp *offramp.OffRamp, opts *bind.CallOpts, args any) (GetAllSourceChainConfigsResult, error) {
		sourceChainSelectors, sourceChainConfigs, err := offRamp.GetAllSourceChainConfigs(opts)
		if err != nil {
			return GetAllSourceChainConfigsResult{}, err
		}
		return GetAllSourceChainConfigsResult{
			SourceChainSelectors: sourceChainSelectors,
			SourceChainConfigs:   sourceChainConfigs,
		}, nil
	},
})
