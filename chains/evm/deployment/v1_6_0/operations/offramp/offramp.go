package offramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "OffRamp"
var Version *semver.Version = semver.MustParse("1.6.0")

type OffRampParams struct {
	GasForCallExactCheck                    uint16
	PermissionLessExecutionThresholdSeconds uint32
	MessageInterceptor                      common.Address
}

type StaticConfig = offramp.OffRampStaticConfig
type DynamicConfig = offramp.OffRampDynamicConfig
type SourceChainConfigArgs = offramp.OffRampSourceChainConfigArgs

type ConstructorArgs struct {
	StaticConfig       StaticConfig
	DynamicConfig      DynamicConfig
	SourceChainConfigs []SourceChainConfigArgs
}

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

var OffRampApplySourceChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]offramp.OffRampSourceChainConfigArgs, *offramp.OffRamp]{
	Name:            "offramp:apply-source-chain-config-updates",
	Version:         Version,
	Description:     "Applies updates to source chain configs on the OffRamp 1.6.0 contract",
	ContractType:    ContractType,
	ContractABI:     offramp.OffRampABI,
	NewContract:     offramp.NewOffRamp,
	IsAllowedCaller: contract.OnlyOwner[*offramp.OffRamp, []offramp.OffRampSourceChainConfigArgs],
	Validate:        func([]offramp.OffRampSourceChainConfigArgs) error { return nil },
	CallContract: func(offRamp *offramp.OffRamp, opts *bind.TransactOpts, args []offramp.OffRampSourceChainConfigArgs) (*types.Transaction, error) {
		return offRamp.ApplySourceChainConfigUpdates(opts, args)
	},
})

var OffRampSetOcr3 = contract.NewWrite(contract.WriteParams[[]offramp.MultiOCR3BaseOCRConfigArgs, *offramp.OffRamp]{
	Name:            "offramp:set-ocr3",
	Version:         Version,
	Description:     "Sets the OCR3 configuration on the OffRamp 1.6.0 contract",
	ContractType:    ContractType,
	ContractABI:     offramp.OffRampABI,
	NewContract:     offramp.NewOffRamp,
	IsAllowedCaller: contract.OnlyOwner[*offramp.OffRamp, []offramp.MultiOCR3BaseOCRConfigArgs],
	Validate:        func([]offramp.MultiOCR3BaseOCRConfigArgs) error { return nil },
	CallContract: func(offRamp *offramp.OffRamp, opts *bind.TransactOpts, args []offramp.MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
		return offRamp.SetOCR3Configs(opts, args)
	},
})
