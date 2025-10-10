package offramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "OffRamp"

var OffRampApplySourceChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]offramp.OffRampSourceChainConfigArgs, *offramp.OffRamp]{
	Name:            "offramp:apply-source-chain-config-updates",
	Version:         semver.MustParse("1.6.0"),
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
