package onramp

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"

	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

var (
	ContractType cldf_deployment.ContractType = "EVM2EVMOnRamp"
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

var OnRampGetTokenTransferFeeConfig = contract.NewRead(contract.ReadParams[common.Address, evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfig, *evm_2_evm_onramp.EVM2EVMOnRamp]{
	Name:         "onramp:get-token-transfer-fee-config",
	Version:      Version,
	Description:  "Reads the token transfer fee config from the OnRamp 1.5.0 contract",
	ContractType: ContractType,
	NewContract:  evm_2_evm_onramp.NewEVM2EVMOnRamp,
	CallContract: func(onRamp *evm_2_evm_onramp.EVM2EVMOnRamp, opts *bind.CallOpts, args common.Address) (evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfig, error) {
		return onRamp.GetTokenTransferFeeConfig(opts, args)
	},
})

var OnRampStaticConfig = contract.NewRead(contract.ReadParams[any, evm_2_evm_onramp.EVM2EVMOnRampStaticConfig, *evm_2_evm_onramp.EVM2EVMOnRamp]{
	Name:         "onramp:static-config",
	Version:      Version,
	Description:  "Reads the static config from the OnRamp 1.5.0 contract",
	ContractType: ContractType,
	NewContract:  evm_2_evm_onramp.NewEVM2EVMOnRamp,
	CallContract: func(onRamp *evm_2_evm_onramp.EVM2EVMOnRamp, opts *bind.CallOpts, args any) (evm_2_evm_onramp.EVM2EVMOnRampStaticConfig, error) {
		return onRamp.GetStaticConfig(opts)
	},
})

var OnRampDynamicConfig = contract.NewRead(contract.ReadParams[any, evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig, *evm_2_evm_onramp.EVM2EVMOnRamp]{
	Name:         "onramp:dynamic-config",
	Version:      Version,
	Description:  "Reads the dynamic config from the OnRamp 1.5.0 contract",
	ContractType: ContractType,
	NewContract:  evm_2_evm_onramp.NewEVM2EVMOnRamp,
	CallContract: func(onRamp *evm_2_evm_onramp.EVM2EVMOnRamp, opts *bind.CallOpts, args any) (evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig, error) {
		return onRamp.GetDynamicConfig(opts)
	},
})

var OnRampFeeTokenConfig = contract.NewRead(contract.ReadParams[common.Address, evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfig, *evm_2_evm_onramp.EVM2EVMOnRamp]{
	Name:         "onramp:fee-token-config",
	Version:      Version,
	Description:  "Reads the fee token config for a given token from the OnRamp 1.5.0 contract",
	ContractType: ContractType,
	NewContract:  evm_2_evm_onramp.NewEVM2EVMOnRamp,
	CallContract: func(onRamp *evm_2_evm_onramp.EVM2EVMOnRamp, opts *bind.CallOpts, args common.Address) (evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfig, error) {
		return onRamp.GetFeeTokenConfig(opts, args)
	},
})
