package off_ramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/latest/off_ramp"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "OffRamp"

type StaticConfig = off_ramp.OffRampStaticConfig

type ConstructorArgs struct {
	StaticConfig StaticConfig
}

type SourceChainConfigArgs = off_ramp.OffRampSourceChainConfigArgs

type SourceChainConfig = off_ramp.OffRampSourceChainConfig

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "off-ramp:deploy",
	Version:          semver.MustParse("1.7.0"),
	Description:      "Deploys the OffRamp contract",
	ContractMetadata: off_ramp.OffRampMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *semver.MustParse("1.7.0")).String(): {
			EVM: common.FromHex(off_ramp.OffRampBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var ApplySourceChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]SourceChainConfigArgs, *off_ramp.OffRamp]{
	Name:            "off-ramp:apply-source-chain-config-updates",
	Version:         semver.MustParse("1.7.0"),
	Description:     "Applies updates to source chain configurations on the OffRamp",
	ContractType:    ContractType,
	ContractABI:     off_ramp.OffRampABI,
	NewContract:     off_ramp.NewOffRamp,
	IsAllowedCaller: contract.OnlyOwner[*off_ramp.OffRamp, []SourceChainConfigArgs],
	Validate:        func([]SourceChainConfigArgs) error { return nil },
	CallContract: func(offRamp *off_ramp.OffRamp, opts *bind.TransactOpts, args []SourceChainConfigArgs) (*types.Transaction, error) {
		return offRamp.ApplySourceChainConfigUpdates(opts, args)
	},
})

var GetSourceChainConfig = contract.NewRead(contract.ReadParams[uint64, SourceChainConfig, *off_ramp.OffRamp]{
	Name:         "off-ramp:get-source-chain-config",
	Version:      semver.MustParse("1.7.0"),
	Description:  "Gets the source chain configuration for a given source chain selector",
	ContractType: ContractType,
	NewContract:  off_ramp.NewOffRamp,
	CallContract: func(offRamp *off_ramp.OffRamp, opts *bind.CallOpts, args uint64) (SourceChainConfig, error) {
		return offRamp.GetSourceChainConfig(opts, args)
	},
})
