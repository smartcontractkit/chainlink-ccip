package offramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
)

var ContractType cldf_deployment.ContractType = "OffRamp"
var Version = semver.MustParse("1.6.0")

type ConstructorArgs struct {
	StaticConfig       offramp.OffRampStaticConfig
	DynamicConfig      offramp.OffRampDynamicConfig
	SourceChainConfigs []offramp.OffRampSourceChainConfigArgs
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "offramp:deploy",
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

var ApplySourceChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]offramp.OffRampSourceChainConfigArgs, *offramp.OffRamp]{
	Name:            "offramp:apply-source-chain-config-updates",
	Version:         Version,
	Description:     "Calls applySourceChainConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     offramp.OffRampABI,
	NewContract:     offramp.NewOffRamp,
	IsAllowedCaller: contract.OnlyOwner[*offramp.OffRamp, []offramp.OffRampSourceChainConfigArgs],
	Validate:        func([]offramp.OffRampSourceChainConfigArgs) error { return nil },
	CallContract: func(offRamp *offramp.OffRamp, opts *bind.TransactOpts, args []offramp.OffRampSourceChainConfigArgs) (*types.Transaction, error) {
		return offRamp.ApplySourceChainConfigUpdates(opts, args)
	},
})

var SetOCR3Configs = contract.NewWrite(contract.WriteParams[[]offramp.MultiOCR3BaseOCRConfigArgs, *offramp.OffRamp]{
	Name:            "offramp:set-o-c-r3-configs",
	Version:         Version,
	Description:     "Calls setOCR3Configs on the contract",
	ContractType:    ContractType,
	ContractABI:     offramp.OffRampABI,
	NewContract:     offramp.NewOffRamp,
	IsAllowedCaller: contract.OnlyOwner[*offramp.OffRamp, []offramp.MultiOCR3BaseOCRConfigArgs],
	Validate:        func([]offramp.MultiOCR3BaseOCRConfigArgs) error { return nil },
	CallContract: func(offRamp *offramp.OffRamp, opts *bind.TransactOpts, args []offramp.MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
		return offRamp.SetOCR3Configs(opts, args)
	},
})
