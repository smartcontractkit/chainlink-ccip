package onramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"
)

var ContractType cldf_deployment.ContractType = "OnRamp"
var Version = semver.MustParse("1.6.0")

type ConstructorArgs struct {
	StaticConfig        onramp.OnRampStaticConfig
	DynamicConfig       onramp.OnRampDynamicConfig
	DestChainConfigArgs []onramp.OnRampDestChainConfigArgs
}

var Deploy = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "onramp:deploy",
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

var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]onramp.OnRampDestChainConfigArgs, *onramp.OnRamp]{
	Name:            "onramp:apply-dest-chain-config-updates",
	Version:         Version,
	Description:     "Calls applyDestChainConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     onramp.OnRampABI,
	NewContract:     onramp.NewOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*onramp.OnRamp, []onramp.OnRampDestChainConfigArgs],
	Validate:        func([]onramp.OnRampDestChainConfigArgs) error { return nil },
	CallContract: func(onRamp *onramp.OnRamp, opts *bind.TransactOpts, args []onramp.OnRampDestChainConfigArgs) (*types.Transaction, error) {
		return onRamp.ApplyDestChainConfigUpdates(opts, args)
	},
})
