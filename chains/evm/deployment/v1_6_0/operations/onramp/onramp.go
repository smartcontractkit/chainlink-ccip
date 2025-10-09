package onramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/onramp"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "OnRamp"

var OnRampApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]onramp.OnRampDestChainConfigArgs, *onramp.OnRamp]{
	Name:            "onramp:apply-dest-chain-config-updates",
	Version:         semver.MustParse("1.6.0"),
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
