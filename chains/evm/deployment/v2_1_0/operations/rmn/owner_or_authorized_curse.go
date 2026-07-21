package rmn

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
)

// isOwnerOrAuthorizedCaller allows direct execution for either RMN owner or authorized callers.
// This is needed for pre-transfer curse migration where deployer is owner but not yet an authorized caller.
func isOwnerOrAuthorizedCaller[ARGS any](
	c *RMNContract,
	opts *bind.CallOpts,
	caller common.Address,
	_ ARGS,
) (bool, error) {
	owner, err := c.Owner(opts)
	if err != nil {
		return false, fmt.Errorf("failed to read RMN owner: %w", err)
	}
	if owner == caller {
		return true, nil
	}
	authorizedCallers, err := c.GetAllAuthorizedCallers(opts)
	if err != nil {
		return false, fmt.Errorf("failed to read RMN authorized callers: %w", err)
	}
	for _, authorized := range authorizedCallers {
		if authorized == caller {
			return true, nil
		}
	}
	return false, nil
}

var CurseByOwnerOrAuthorized = contract.NewWrite(contract.WriteParams[[16]byte, *RMNContract]{
	Name:            "rmn:curse-owner-or-authorized",
	Version:         Version,
	Description:     "Calls curse on the contract, allowing owner or authorized callers",
	ContractType:    ContractType,
	ContractABI:     RMNABI,
	NewContract:     NewRMNContract,
	IsAllowedCaller: isOwnerOrAuthorizedCaller[[16]byte],
	Validate:        func([16]byte) error { return nil },
	CallContract: func(
		c *RMNContract,
		opts *bind.TransactOpts,
		args [16]byte,
	) (*types.Transaction, error) {
		return c.Curse(opts, args)
	},
})

var Curse0ByOwnerOrAuthorized = contract.NewWrite(contract.WriteParams[[][16]byte, *RMNContract]{
	Name:            "rmn:curse0-owner-or-authorized",
	Version:         Version,
	Description:     "Calls curse0 on the contract, allowing owner or authorized callers",
	ContractType:    ContractType,
	ContractABI:     RMNABI,
	NewContract:     NewRMNContract,
	IsAllowedCaller: isOwnerOrAuthorizedCaller[[][16]byte],
	Validate:        func([][16]byte) error { return nil },
	CallContract: func(
		c *RMNContract,
		opts *bind.TransactOpts,
		args [][16]byte,
	) (*types.Transaction, error) {
		return c.Curse0(opts, args)
	},
})
