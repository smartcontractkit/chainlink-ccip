package contract

import "github.com/ethereum/go-ethereum/common"

// FunctionInput is the input for read/write operations that target a contract at
// a specific on-chain address. Address (and optionally ChainSelector) are set at
// execution time; Args are passed to the contract method.
type FunctionInput[ARGS any] struct {
	Args          ARGS           `json:"args"`
	ChainSelector uint64         `json:"chainSelector,omitempty"`
	Address       common.Address `json:"address,omitempty"`
}
