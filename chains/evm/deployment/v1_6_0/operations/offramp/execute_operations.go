package offramp

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_6_0/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// OCRConfigArgs is the element type for SetOCR3Configs (matches Solidity OCR config tuple).
type OCRConfigArgs = gobindings.MultiOCR3BaseOCRConfigArgs

// DynamicConfig is the OffRamp dynamic configuration payload for SetDynamicConfig.
type DynamicConfig = gobindings.OffRampDynamicConfig

// SourceChainConfigArgs is the element type for ApplySourceChainConfigUpdates.
type SourceChainConfigArgs = gobindings.OffRampSourceChainConfigArgs

var ApplySourceChainConfigUpdates = contract.NewWrite(contract.WriteParams[[]gobindings.OffRampSourceChainConfigArgs, *gobindings.OffRamp]{
	Name:            "offramp:apply-source-chain-config-updates",
	Version:         Version,
	Description:     "Calls applySourceChainConfigUpdates on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.OffRampMetaData.ABI,
	NewContract:     gobindings.NewOffRamp,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.OffRamp, []gobindings.OffRampSourceChainConfigArgs],
	Validate:        func([]gobindings.OffRampSourceChainConfigArgs) error { return nil },
	CallContract: func(c *gobindings.OffRamp, opts *bind.TransactOpts, args []gobindings.OffRampSourceChainConfigArgs) (*types.Transaction, error) {
		return c.ApplySourceChainConfigUpdates(opts, args)
	},
})

var SetOCR3Configs = contract.NewWrite(contract.WriteParams[[]gobindings.MultiOCR3BaseOCRConfigArgs, *gobindings.OffRamp]{
	Name:            "offramp:set-ocr3-configs",
	Version:         Version,
	Description:     "Calls setOCR3Configs on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.OffRampMetaData.ABI,
	NewContract:     gobindings.NewOffRamp,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.OffRamp, []gobindings.MultiOCR3BaseOCRConfigArgs],
	Validate:        func([]gobindings.MultiOCR3BaseOCRConfigArgs) error { return nil },
	CallContract: func(c *gobindings.OffRamp, opts *bind.TransactOpts, args []gobindings.MultiOCR3BaseOCRConfigArgs) (*types.Transaction, error) {
		return c.SetOCR3Configs(opts, args)
	},
})

var SetDynamicConfig = contract.NewWrite(contract.WriteParams[gobindings.OffRampDynamicConfig, *gobindings.OffRamp]{
	Name:            "offramp:set-dynamic-config",
	Version:         Version,
	Description:     "Calls setDynamicConfig on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.OffRampMetaData.ABI,
	NewContract:     gobindings.NewOffRamp,
	IsAllowedCaller: contract.OnlyOwner[*gobindings.OffRamp, gobindings.OffRampDynamicConfig],
	Validate:        func(gobindings.OffRampDynamicConfig) error { return nil },
	CallContract: func(c *gobindings.OffRamp, opts *bind.TransactOpts, args gobindings.OffRampDynamicConfig) (*types.Transaction, error) {
		return c.SetDynamicConfig(opts, args)
	},
})

var GetStaticConfig = contract.NewRead(contract.ReadParams[struct{}, gobindings.OffRampStaticConfig, *gobindings.OffRamp]{
	Name:         "offramp:get-static-config",
	Version:      Version,
	Description:  "Calls getStaticConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewOffRamp,
	CallContract: func(c *gobindings.OffRamp, opts *bind.CallOpts, args struct{}) (gobindings.OffRampStaticConfig, error) {
		return c.GetStaticConfig(opts)
	},
})

var GetDynamicConfig = contract.NewRead(contract.ReadParams[struct{}, gobindings.OffRampDynamicConfig, *gobindings.OffRamp]{
	Name:         "offramp:get-dynamic-config",
	Version:      Version,
	Description:  "Calls getDynamicConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewOffRamp,
	CallContract: func(c *gobindings.OffRamp, opts *bind.CallOpts, args struct{}) (gobindings.OffRampDynamicConfig, error) {
		return c.GetDynamicConfig(opts)
	},
})

var GetSourceChainConfig = contract.NewRead(contract.ReadParams[uint64, gobindings.OffRampSourceChainConfig, *gobindings.OffRamp]{
	Name:         "offramp:get-source-chain-config",
	Version:      Version,
	Description:  "Calls getSourceChainConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewOffRamp,
	CallContract: func(c *gobindings.OffRamp, opts *bind.CallOpts, args uint64) (gobindings.OffRampSourceChainConfig, error) {
		return c.GetSourceChainConfig(opts, args)
	},
})
