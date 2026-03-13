package onramp

import (
	"fmt"
	"math/big"

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

type ConstructorArgs struct {
	StaticConfig               evm_2_evm_onramp.EVM2EVMOnRampStaticConfig
	DynamicConfig              evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig
	RateLimiterConfig          evm_2_evm_onramp.RateLimiterConfig
	FeeTokenConfigs            []evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs
	TokenTransferFeeConfigArgs []evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs
	NopsAndWeights             []evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight
}

var DeployOnRamp = contract.NewDeploy(contract.DeployParams[ConstructorArgs]{
	Name:             "onramp:deploy",
	Version:          Version,
	Description:      "Deploys the OnRamp 1.5.0 contract",
	ContractMetadata: evm_2_evm_onramp.EVM2EVMOnRampMetaData,
	Validate:         func(args ConstructorArgs) error { return nil },
	BytecodeByTypeAndVersion: map[string]contract.Bytecode{
		cldf_deployment.NewTypeAndVersion(ContractType, *Version).String(): {
			EVM: common.FromHex(evm_2_evm_onramp.EVM2EVMOnRampBin),
		},
	},
})

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

// =============================================================================
// Fee Withdrawal Operations
// =============================================================================

// WithdrawNonLinkFeesInput specifies which non-LINK token to withdraw and where to send it
type WithdrawNonLinkFeesInput struct {
	FeeToken common.Address
	To       common.Address
}

// OnRampWithdrawNonLinkFees withdraws accumulated non-LINK fee tokens to a specified address.
var OnRampWithdrawNonLinkFees = contract.NewWrite(contract.WriteParams[WithdrawNonLinkFeesInput, *evm_2_evm_onramp.EVM2EVMOnRamp]{
	Name:            "onramp:withdraw-non-link-fees",
	Version:         Version,
	Description:     "Withdraws non-LINK fee tokens from the OnRamp 1.5.0 contract to a specified address",
	ContractType:    ContractType,
	ContractABI:     evm_2_evm_onramp.EVM2EVMOnRampABI,
	NewContract:     evm_2_evm_onramp.NewEVM2EVMOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*evm_2_evm_onramp.EVM2EVMOnRamp, WithdrawNonLinkFeesInput],
	Validate:        func(args WithdrawNonLinkFeesInput) error { return nil },
	CallContract: func(onRamp *evm_2_evm_onramp.EVM2EVMOnRamp, opts *bind.TransactOpts, args WithdrawNonLinkFeesInput) (*types.Transaction, error) {
		return onRamp.WithdrawNonLinkFees(opts, args.FeeToken, args.To)
	},
})

// NopAndWeight is an alias for the gethwrapper type to simplify consumer imports
type NopAndWeight = evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight

// SetNopsInput specifies the NOPs and their payment weights
type SetNopsInput struct {
	NopsAndWeights []NopAndWeight
}

// OnRampSetNops sets the NOPs and their payment weights for LINK fee distribution.
var OnRampSetNops = contract.NewWrite(contract.WriteParams[SetNopsInput, *evm_2_evm_onramp.EVM2EVMOnRamp]{
	Name:            "onramp:set-nops",
	Version:         Version,
	Description:     "Sets the NOPs and their payment weights on the OnRamp 1.5.0 contract",
	ContractType:    ContractType,
	ContractABI:     evm_2_evm_onramp.EVM2EVMOnRampABI,
	NewContract:     evm_2_evm_onramp.NewEVM2EVMOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*evm_2_evm_onramp.EVM2EVMOnRamp, SetNopsInput],
	Validate: func(args SetNopsInput) error {
		if len(args.NopsAndWeights) == 0 {
			return fmt.Errorf("NopsAndWeights list cannot be empty")
		}
		return nil
	},
	CallContract: func(onRamp *evm_2_evm_onramp.EVM2EVMOnRamp, opts *bind.TransactOpts, args SetNopsInput) (*types.Transaction, error) {
		return onRamp.SetNops(opts, args.NopsAndWeights)
	},
})

// OnRampPayNops triggers payout of accumulated LINK fees to NOPs based on their configured weights
var OnRampPayNops = contract.NewWrite(contract.WriteParams[any, *evm_2_evm_onramp.EVM2EVMOnRamp]{
	Name:            "onramp:pay-nops",
	Version:         Version,
	Description:     "Pays out accumulated LINK fees to NOPs based on their weights on the OnRamp 1.5.0 contract",
	ContractType:    ContractType,
	ContractABI:     evm_2_evm_onramp.EVM2EVMOnRampABI,
	NewContract:     evm_2_evm_onramp.NewEVM2EVMOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*evm_2_evm_onramp.EVM2EVMOnRamp, any],
	Validate:        func(args any) error { return nil },
	CallContract: func(onRamp *evm_2_evm_onramp.EVM2EVMOnRamp, opts *bind.TransactOpts, args any) (*types.Transaction, error) {
		return onRamp.PayNops(opts)
	},
})

// =============================================================================
// Fee Withdrawal Read Operations (for validation)
// =============================================================================

// GetNopsResult contains the current NOP configuration
type GetNopsResult struct {
	NopsAndWeights []NopAndWeight
	WeightsTotal   *big.Int
}

// OnRampGetNops reads the current NOP configuration from the OnRamp
// Use this to verify NOPs are correctly configured before calling PayNops
var OnRampGetNops = contract.NewRead(contract.ReadParams[any, GetNopsResult, *evm_2_evm_onramp.EVM2EVMOnRamp]{
	Name:         "onramp:get-nops",
	Version:      Version,
	Description:  "Reads the current NOP configuration from the OnRamp 1.5.0 contract",
	ContractType: ContractType,
	NewContract:  evm_2_evm_onramp.NewEVM2EVMOnRamp,
	CallContract: func(onRamp *evm_2_evm_onramp.EVM2EVMOnRamp, opts *bind.CallOpts, args any) (GetNopsResult, error) {
		result, err := onRamp.GetNops(opts)
		if err != nil {
			return GetNopsResult{}, err
		}
		return GetNopsResult{
			NopsAndWeights: result.NopsAndWeights,
			WeightsTotal:   result.WeightsTotal,
		}, nil
	},
})

// OnRampGetNopFeesJuels reads the current accumulated LINK fees pending for NOP payout
var OnRampGetNopFeesJuels = contract.NewRead(contract.ReadParams[any, *big.Int, *evm_2_evm_onramp.EVM2EVMOnRamp]{
	Name:         "onramp:get-nop-fees-juels",
	Version:      Version,
	Description:  "Reads the accumulated LINK fees (in juels) pending for NOP payout",
	ContractType: ContractType,
	NewContract:  evm_2_evm_onramp.NewEVM2EVMOnRamp,
	CallContract: func(onRamp *evm_2_evm_onramp.EVM2EVMOnRamp, opts *bind.CallOpts, args any) (*big.Int, error) {
		return onRamp.GetNopFeesJuels(opts)
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
