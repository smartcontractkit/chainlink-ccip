package advanced_pool_hooks

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/advanced_pool_hooks"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// AdvancedPoolHooksABI is the JSON ABI of the AdvancedPoolHooks contract.
var AdvancedPoolHooksABI = gobindings.AdvancedPoolHooksMetaData.ABI

// Type aliases for sequence call sites.
type (
	CCVConfigArg         = gobindings.AdvancedPoolHooksCCVConfigArg
	AuthorizedCallerArgs = gobindings.AuthorizedCallersAuthorizedCallerArgs
)

var ApplyCCVConfigUpdates = contract.NewWrite(contract.WriteParams[[]gobindings.AdvancedPoolHooksCCVConfigArg, *gobindings.AdvancedPoolHooks]{
	Name:            "advanced-pool-hooks:apply-ccv-config-updates",
	Version:         Version,
	Description:     "Calls applyCCVConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.AdvancedPoolHooksMetaData.ABI,
	NewContract:     gobindings.NewAdvancedPoolHooks,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.AdvancedPoolHooks, []gobindings.AdvancedPoolHooksCCVConfigArg],
	Validate:        func([]gobindings.AdvancedPoolHooksCCVConfigArg) error { return nil },
	CallContract: func(c *gobindings.AdvancedPoolHooks, opts *bind.TransactOpts, args []gobindings.AdvancedPoolHooksCCVConfigArg) (*types.Transaction, error) {
		return c.ApplyCCVConfigUpdates(opts, args)
	},
})

var SetThresholdAmount = contract.NewWrite(contract.WriteParams[*big.Int, *gobindings.AdvancedPoolHooks]{
	Name:            "advanced-pool-hooks:set-threshold-amount",
	Version:         Version,
	Description:     "Calls setThresholdAmount on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.AdvancedPoolHooksMetaData.ABI,
	NewContract:     gobindings.NewAdvancedPoolHooks,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.AdvancedPoolHooks, *big.Int],
	Validate:        func(*big.Int) error { return nil },
	CallContract: func(c *gobindings.AdvancedPoolHooks, opts *bind.TransactOpts, args *big.Int) (*types.Transaction, error) {
		return c.SetThresholdAmount(opts, args)
	},
})

var ApplyAuthorizedCallerUpdates = contract.NewWrite(contract.WriteParams[gobindings.AuthorizedCallersAuthorizedCallerArgs, *gobindings.AdvancedPoolHooks]{
	Name:            "advanced-pool-hooks:apply-authorized-caller-updates",
	Version:         Version,
	Description:     "Calls applyAuthorizedCallerUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.AdvancedPoolHooksMetaData.ABI,
	NewContract:     gobindings.NewAdvancedPoolHooks,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.AdvancedPoolHooks, gobindings.AuthorizedCallersAuthorizedCallerArgs],
	Validate:        func(gobindings.AuthorizedCallersAuthorizedCallerArgs) error { return nil },
	CallContract: func(c *gobindings.AdvancedPoolHooks, opts *bind.TransactOpts, args gobindings.AuthorizedCallersAuthorizedCallerArgs) (*types.Transaction, error) {
		return c.ApplyAuthorizedCallerUpdates(opts, args)
	},
})

var GetThresholdAmount = contract.NewRead(contract.ReadParams[struct{}, *big.Int, *gobindings.AdvancedPoolHooks]{
	Name:         "advanced-pool-hooks:get-threshold-amount",
	Version:      Version,
	Description:  "Calls getThresholdAmount on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewAdvancedPoolHooks,
	CallContract: func(c *gobindings.AdvancedPoolHooks, opts *bind.CallOpts, args struct{}) (*big.Int, error) {
		return c.GetThresholdAmount(opts)
	},
})

var GetAllAuthorizedCallers = contract.NewRead(contract.ReadParams[struct{}, []common.Address, *gobindings.AdvancedPoolHooks]{
	Name:         "advanced-pool-hooks:get-all-authorized-callers",
	Version:      Version,
	Description:  "Calls getAllAuthorizedCallers on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewAdvancedPoolHooks,
	CallContract: func(c *gobindings.AdvancedPoolHooks, opts *bind.CallOpts, args struct{}) ([]common.Address, error) {
		return c.GetAllAuthorizedCallers(opts)
	},
})

var GetRequiredCCVs = contract.NewRead(contract.ReadParams[GetRequiredCCVsArgs, []common.Address, *gobindings.AdvancedPoolHooks]{
	Name:         "advanced-pool-hooks:get-required-cc-vs",
	Version:      Version,
	Description:  "Calls getRequiredCCVs on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewAdvancedPoolHooks,
	CallContract: func(c *gobindings.AdvancedPoolHooks, opts *bind.CallOpts, args GetRequiredCCVsArgs) ([]common.Address, error) {
		return c.GetRequiredCCVs(opts, args.Arg0, args.RemoteChainSelector, args.Amount, args.Arg3, args.Arg4, args.Direction)
	},
})
