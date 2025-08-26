package call

import "github.com/ethereum/go-ethereum/common"

// Input is the input structure for all calls.
type Input[ARGS any] struct {
	// Address defines the contract to call.
	Address common.Address `json:"address"`
	// ChainSelector is the selector for the chain on which the contract resides.
	// Required to differentiate between operation runs with the same data targeting different chains.
	ChainSelector uint64 `json:"chainSelector"`
	// Args are the parameters passed to the contract call.
	Args ARGS `json:"args"`
}
