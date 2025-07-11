package ccipocr3

import (
	"math/big"

	ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Deprecated: Use ccipocr3common.UnknownAddress instead.
type UnknownAddress = ccipocr3common.UnknownAddress

// Deprecated: Use ccipocr3common.NewUnknownAddressFromHex instead.
func NewUnknownAddressFromHex(s string) (UnknownAddress, error) {
	return ccipocr3common.NewUnknownAddressFromHex(s)
}

// Deprecated: Use ccipocr3common.UnknownEncodedAddress instead.
type UnknownEncodedAddress = ccipocr3common.UnknownEncodedAddress

// Deprecated: Use ccipocr3common.Bytes instead.
type Bytes = ccipocr3common.Bytes

// Deprecated: Use ccipocr3common.NewBytesFromString instead.
func NewBytesFromString(s string) (Bytes, error) {
	return ccipocr3common.NewBytesFromString(s)
}

// Deprecated: Use ccipocr3common.Bytes32 instead.
type Bytes32 = ccipocr3common.Bytes32

// Deprecated: Use ccipocr3common.NewBytes32FromString instead.
func NewBytes32FromString(s string) (Bytes32, error) {
	return ccipocr3common.NewBytes32FromString(s)
}

// Deprecated: Use ccipocr3common.BigInt instead.
type BigInt = ccipocr3common.BigInt

// Deprecated: Use ccipocr3common.NewBigInt instead.
func NewBigInt(i *big.Int) BigInt {
	return ccipocr3common.NewBigInt(i)
}

// Deprecated: Use ccipocr3common.NewBigIntFromInt64 instead.
func NewBigIntFromInt64(i int64) BigInt {
	return ccipocr3common.NewBigIntFromInt64(i)
}
