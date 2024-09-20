package typconv

import (
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/typeconv/evm"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Address is the specific address family which can encode itself.
type Address interface {
	Encode() (EncodedAddress, error)
}

// EncodedAddress is the specific encoded address family which can decode itself.
type EncodedAddress interface {
	Decode() (Address, error)
}

// UnknownAddress represents an address type which is not attached to a specific chain.
type UnknownAddress []byte

// UnknownEncodedAddress represents an encoded address type which is not attached to a specific chain.
type UnknownEncodedAddress string

// AddressProvider initializes an unknown address type, it knows how to map chainSel to the correct
// address implementation. Once initialized the developer does not need to concern themself with
// address semantics.
type AddressProvider struct {
}

// MakeAddress initializes an unknown address type, it knows how to map chainSel to the correct
func (ap AddressProvider) MakeAddress(data UnknownAddress, chainSel ccipocr3.ChainSelector) (Address, error) {
	// TODO: Move to chain agnostic location.
	switch chainSel {
	default:
		return evm.EVMAddress(data), nil
	}
}

//MakeEncodedAddress(data UnknownEncodedAddress, chainSel ccipocr3.ChainSelector) (EncodedAddress, error)
