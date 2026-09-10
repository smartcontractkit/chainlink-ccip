package onramp

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/chain/evm/operations/contract"
)

// getOnRampOwner reads the owner of the OnRamp. It lives here rather than in the
// generated operations package (operations-gen output must not be edited).
var Owner = contract.NewRead(contract.ReadParams[struct{}, common.Address, *OnRampContract]{
	Name:         "onramp:owner",
	Version:      Version,
	Description:  "Returns the owner of the OnRamp",
	ContractType: ContractType,
	NewContract:  NewOnRampContract,
	CallContract: func(c *OnRampContract, opts *bind.CallOpts, _ struct{}) (common.Address, error) {
		return c.Owner(opts)
	},
})
