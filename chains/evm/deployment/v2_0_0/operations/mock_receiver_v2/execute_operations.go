package mock_receiver_v2

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"

	gobindings "github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v2_0_0/mock_receiver_v2"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// MockReceiverV2ABI is the JSON ABI of the MockReceiverV2 contract.
var MockReceiverV2ABI = gobindings.MockReceiverV2MetaData.ABI

var GetCCVsAndFinalityConfig = contract.NewRead(contract.ReadParams[GetCCVsAndFinalityConfigArgs, GetCCVsAndFinalityConfigResult, *gobindings.MockReceiverV2]{
	Name:         "mock-receiver-v2:get-cc-vs-and-finality-config",
	Version:      Version,
	Description:  "Calls getCCVsAndFinalityConfig on the contract",
	ContractType: ContractType,
	NewContract:  gobindings.NewMockReceiverV2,
	CallContract: func(c *gobindings.MockReceiverV2, opts *bind.CallOpts, args GetCCVsAndFinalityConfigArgs) (GetCCVsAndFinalityConfigResult, error) {
		res, err := c.GetCCVsAndFinalityConfig(opts, args.Arg0, args.Arg1)
		if err != nil {
			return GetCCVsAndFinalityConfigResult{}, err
		}
		return GetCCVsAndFinalityConfigResult{
			RequiredVerifier:      res.RequiredVerifier,
			OptionalVerifiers:     res.OptionalVerifiers,
			Threshold:             res.Threshold,
			AllowedFinalityConfig: res.AllowedFinalityConfig,
		}, nil
	},
})

var SetAllowedFinalityConfig = contract.NewWrite(contract.WriteParams[[4]byte, *gobindings.MockReceiverV2]{
	Name:            "mock-receiver-v2:set-allowed-finality-config",
	Version:         Version,
	Description:     "Calls setAllowedFinalityConfig on the contract",
	ContractType:    ContractType,
	ContractABI:     gobindings.MockReceiverV2MetaData.ABI,
	NewContract:     gobindings.NewMockReceiverV2,
	IsAllowedCaller: contract.AllCallersAllowed[*gobindings.MockReceiverV2, [4]byte],
	Validate:        func([4]byte) error { return nil },
	CallContract: func(c *gobindings.MockReceiverV2, opts *bind.TransactOpts, args [4]byte) (*types.Transaction, error) {
		return c.SetAllowedFinalityConfig(opts, args)
	},
})
