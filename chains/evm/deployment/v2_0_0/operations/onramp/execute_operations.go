package onramp

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/onramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// DestChainConfigArgs is the element type for ApplyDestChainConfigUpdates.
type DestChainConfigArgs = gobindings.OnRampDestChainConfigArgs

// DynamicConfig is the OnRamp dynamic configuration for SetDynamicConfig.
type DynamicConfig = gobindings.OnRampDynamicConfig

// StaticConfig is the OnRamp static configuration for deploy/constructor.
type StaticConfig = gobindings.OnRampStaticConfig

var ApplyDestChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]gobindings.OnRampDestChainConfigArgs, *gobindings.OnRamp]{
	Name:            "onramp:apply-dest-chain-config-updates",
	Version:         Version,
	Description:     "Calls applyDestChainConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.OnRampMetaData.ABI,
	NewContract:     gobindings.NewOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.OnRamp, []gobindings.OnRampDestChainConfigArgs],
	Validate:        func([]gobindings.OnRampDestChainConfigArgs) error { return nil },
	CallContract: func(c *gobindings.OnRamp, opts *bind.TransactOpts, args []gobindings.OnRampDestChainConfigArgs) (*types.Transaction, error) {
		return c.ApplyDestChainConfigUpdates(opts, args)
	},
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[gobindings.OnRampDynamicConfig, *gobindings.OnRamp]{
	Name:            "onramp:set-dynamic-config",
	Version:         Version,
	Description:     "Calls setDynamicConfig on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.OnRampMetaData.ABI,
	NewContract:     gobindings.NewOnRamp,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.OnRamp, gobindings.OnRampDynamicConfig],
	Validate:        func(gobindings.OnRampDynamicConfig) error { return nil },
	CallContract: func(c *gobindings.OnRamp, opts *bind.TransactOpts, args gobindings.OnRampDynamicConfig) (*types.Transaction, error) {
		return c.SetDynamicConfig(opts, args)
	},
})

var GetDynamicConfig = contract.NewRead(contract.ReadParams[struct{}, gobindings.OnRampDynamicConfig, *gobindings.OnRamp]{
	Name:         "onramp:get-dynamic-config",
	Version:      Version,
	Description:  "Calls getDynamicConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewOnRamp,
	CallContract: func(c *gobindings.OnRamp, opts *bind.CallOpts, args struct{}) (gobindings.OnRampDynamicConfig, error) {
		return c.GetDynamicConfig(opts)
	},
})

var GetDestChainConfig = contract.NewRead(contract.ReadParams[uint64, gobindings.OnRampDestChainConfig, *gobindings.OnRamp]{
	Name:         "onramp:get-dest-chain-config",
	Version:      Version,
	Description:  "Calls getDestChainConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewOnRamp,
	CallContract: func(c *gobindings.OnRamp, opts *bind.CallOpts, args uint64) (gobindings.OnRampDestChainConfig, error) {
		return c.GetDestChainConfig(opts, args)
	},
})

var GetStaticConfig = contract.NewRead(contract.ReadParams[struct{}, gobindings.OnRampStaticConfig, *gobindings.OnRamp]{
	Name:         "onramp:get-static-config",
	Version:      Version,
	Description:  "Calls getStaticConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewOnRamp,
	CallContract: func(c *gobindings.OnRamp, opts *bind.CallOpts, args struct{}) (gobindings.OnRampStaticConfig, error) {
		return c.GetStaticConfig(opts)
	},
})
