package offramp

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

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
