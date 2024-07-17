package typconv

import (
	"encoding/hex"
)

// AddressBytesToString converts the given address bytes to a string
// based upon the given chain selector's chain family.
// TODO: only EVM supported for now, fix this.
func AddressBytesToString(addr []byte, chainSelector uint64) string {
	// TODO: not EIP-55. Fix this?
	return "0x" + hex.EncodeToString(addr)
}
