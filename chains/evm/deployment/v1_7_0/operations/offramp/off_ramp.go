package offramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/offramp"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "OffRamp"

type StaticConfig = offramp.OffRampStaticConfig

type ConstructorArgs struct {
	StaticConfig StaticConfig
}

type SourceChainConfigArgs = offramp.OffRampSourceChainConfigArgs

type SourceChainConfig = offramp.OffRampSourceChainConfig

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "off-ramp:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the OffRamp contract",
	ContractMetadata: offramp.OffRampMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(offramp.OffRampBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplySourceChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]SourceChainConfigArgs, *offramp.OffRamp]{
	Name:            "off-ramp:apply-source-chain-config-updates",
	Version:         semver.MustParse("1.7.0"),
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
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the source chain configuration for a given source chain selector",
	ContractType: ContractType,
	NewContract:  offramp.NewOffRamp,
	CallContract: func(offRamp *offramp.OffRamp, opts *bind.CallOpts, args uint64) (SourceChainConfig, error) {
		return offRamp.GetSourceChainConfig(opts, args)
	},
})
