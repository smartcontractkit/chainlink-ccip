package onramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var (
	ContractType cldf_deployment.ContractType = "OnRamp"
	Version      *semver.Version              = semver.MustParse("1.5.0")
)

type SetTokenTransferFeeConfigInput struct {
	TokenTransferFeeConfigArgs   []evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs
	TokensToUseDefaultFeeConfigs []common.Address
}

var OnRampSetTokenTransferFeeConfig = contract.NewWrite(contract.WriteParams[SetTokenTransferFeeConfigInput, *evm_2_evm_onramp.EVM2EVMOnRamp]{
	Name:            "onramp:set-token-transfer-fee-config",
	Version:         Version,
	Description:     "Sets token transfer fee configs on the OnRamp 1.5.0 contract",
	ContractType:    ContractType,
	ContractABI:     evm_2_evm_onramp.EVM2EVMOnRampABI,
	NewContract:     evm_2_evm_onramp.NewEVM2EVMOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*evm_2_evm_onramp.EVM2EVMOnRamp, SetTokenTransferFeeConfigInput],
	Validate:        func(args SetTokenTransferFeeConfigInput) error { return nil },
	CallContract: func(onRamp *evm_2_evm_onramp.EVM2EVMOnRamp, opts *bind.TransactOpts, args SetTokenTransferFeeConfigInput) (*types.Transaction, error) {
		return onRamp.SetTokenTransferFeeConfig(opts, args.TokenTransferFeeConfigArgs, args.TokensToUseDefaultFeeConfigs)
	},
})
