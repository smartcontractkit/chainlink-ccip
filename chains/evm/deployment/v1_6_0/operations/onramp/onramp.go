package onramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "OnRamp"
var Version *semver.Version = semver.MustParse("1.6.0")

type StaticConfig = onramp.OnRampStaticConfig
type DynamicConfig = onramp.OnRampDynamicConfig
type DestChainConfigArgs = onramp.OnRampDestChainConfigArgs

type ConstructorArgs struct {
	StaticConfig        StaticConfig
	DynamicConfig       DynamicConfig
	DestChainConfigArgs []DestChainConfigArgs
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "on-ramp:deploy",
	Version:          Version,
	Description:      "Deploys the OnRamp contract",
	ContractMetadata: onramp.OnRampMetaData,
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(onramp.OnRampBin),
		},
	},
	Validate: func(ConstructorArgs) error { return nil },
})

var OnRampApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]onramp.OnRampDestChainConfigArgs, *onramp.OnRamp]{
	Name:            "onramp:apply-dest-chain-config-updates",
	Version:         Version,
	Description:     "Applies updates to destination chain configs on the OnRamp 1.6.0 contract",
	ContractType:    ContractType,
	ContractABI:     onramp.OnRampABI,
	NewContract:     onramp.NewOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*onramp.OnRamp, []onramp.OnRampDestChainConfigArgs],
	Validate:        func([]onramp.OnRampDestChainConfigArgs) error { return nil },
	CallContract: func(onRamp *onramp.OnRamp, opts *bind.TransactOpts, args []onramp.OnRampDestChainConfigArgs) (*types.Transaction, error) {
		return onRamp.ApplyDestChainConfigUpdates(opts, args)
	},
})
