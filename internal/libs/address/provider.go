// package address provides a common interface for address types across different blockchains.
package address

import (
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/address/common"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/address/internal/registry"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	// register evm with registry
	_ "github.com/smartcontractkit/chainlink-ccip/internal/libs/address/internal/evm"
)

// UnknownAddress represents an address type which is not attached to a specific chain.
type UnknownAddress []byte

// UnknownEncodedAddress represents an encoded address type which is not attached to a specific chain.
type UnknownEncodedAddress string

// MakeAddress initializes an unknown address type, it knows how to map chainSel to the correct
func MakeAddress(
	data UnknownAddress,
	chainSel ccipocr3.ChainSelector,
) (common.Address, error) {
	return registry.MakeAddress(data, chainSel)
}

func MakeEncodedAddress(
	data UnknownEncodedAddress,
	chainSel ccipocr3.ChainSelector,
) (common.EncodedAddress, error) {
	return registry.MakeEncodedAddress(string(data), chainSel)
}
