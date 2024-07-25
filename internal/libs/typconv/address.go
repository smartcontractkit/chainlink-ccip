package typconv

import (
	"encoding/hex"
)

// HexEncode converts a byte slice to an EVM string representation of an address
func HexEncode(addr []byte) string {
	return "0x" + hex.EncodeToString(addr)
}
