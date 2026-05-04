package executor

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/executor"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// ExecutorABI is the JSON ABI of the Executor contract.
var ExecutorABI = gobindings.ExecutorMetaData.ABI

// DynamicConfig is the Executor dynamic configuration for SetDynamicConfig and constructor input.
type DynamicConfig = gobindings.ExecutorDynamicConfig

// RemoteChainConfigArgs is the element type returned by GetDestChains and passed to ApplyDestChainUpdates.
type RemoteChainConfigArgs = gobindings.ExecutorRemoteChainConfigArgs

var GetDynamicConfig = contract.NewRead(contract.ReadParams[struct{}, gobindings.ExecutorDynamicConfig, *gobindings.Executor]{
	Name:         "executor:get-dynamic-config",
	Version:      Version,
	Description:  "Calls getDynamicConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewExecutor,
	CallContract: func(c *gobindings.Executor, opts *bind.CallOpts, args struct{}) (gobindings.ExecutorDynamicConfig, error) {
		return c.GetDynamicConfig(opts)
	},
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[gobindings.ExecutorDynamicConfig, *gobindings.Executor]{
	Name:            "executor:set-dynamic-config",
	Version:         Version,
	Description:     "Calls setDynamicConfig on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.ExecutorMetaData.ABI,
	NewContract:     gobindings.NewExecutor,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.Executor, gobindings.ExecutorDynamicConfig],
	Validate:        func(gobindings.ExecutorDynamicConfig) error { return nil },
	CallContract: func(c *gobindings.Executor, opts *bind.TransactOpts, args gobindings.ExecutorDynamicConfig) (*types.Transaction, error) {
		return c.SetDynamicConfig(opts, args)
	},
})

var GetDestChains = contract.NewRead(contract.ReadParams[struct{}, []gobindings.ExecutorRemoteChainConfigArgs, *gobindings.Executor]{
	Name:         "executor:get-dest-chains",
	Version:      Version,
	Description:  "Calls getDestChains on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewExecutor,
	CallContract: func(c *gobindings.Executor, opts *bind.CallOpts, args struct{}) ([]gobindings.ExecutorRemoteChainConfigArgs, error) {
		return c.GetDestChains(opts)
	},
})
