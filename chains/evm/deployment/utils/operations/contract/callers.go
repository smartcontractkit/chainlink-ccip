package contract

import (
	"slices"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// AuthorizedCallersContract is implemented by generated bindings for contracts that mix in
// AuthorizedCallers (getAllAuthorizedCallers view).
type AuthorizedCallersContract interface {
	GetAllAuthorizedCallers(opts *bind.CallOpts) ([]common.Address, error)
}

// IsAuthorizedCaller returns whether caller is present in the contract's authorized caller set.
func IsAuthorizedCaller[C AuthorizedCallersContract, ARGS any](
	c C,
	opts *bind.CallOpts,
	caller common.Address,
	_ ARGS,
) (bool, error) {
	callers, err := c.GetAllAuthorizedCallers(opts)
	if err != nil {
		return false, err
	}
	return slices.Contains(callers, caller), nil
}
