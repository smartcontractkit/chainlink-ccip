package typconv

import (
	"encoding/hex"
	"fmt"
	"strings"
)

// AddressBytesToString converts the given address bytes to a string
// based upon the given chain selector's chain family.
// TODO: only EVM supported for now, fix this.
func AddressBytesToString(addr []byte, chainSelector uint64) string {
	// TODO: not EIP-55. Fix this?
	return "0x" + hex.EncodeToString(addr)
}

// AddressStringToBytes converts the given address string to bytes
// based upon the given chain selector's chain family.
// TODO: only EVM supported for now, fix this.
func AddressStringToBytes(addr string, chainSelector uint64) ([]byte, error) {
	// lower case in case EIP-55 and trim 0x prefix if there
	addrBytes, err := hex.DecodeString(strings.ToLower(strings.TrimPrefix(addr, "0x")))
	if err != nil {
		return nil, fmt.Errorf("failed to decode EVM address '%s': %w", addr, err)
	}

	return addrBytes, nil
}

// KeepNRightBytes returns the last n bytes of the given byte slice.
// Example: KeepNRightBytes([]byte{0x01, 0x02, 0x03, 0x04}, 2) -> []byte{0x03, 0x04}
func KeepNRightBytes(b []byte, n uint) []byte {
	if n >= uint(len(b)) {
		return b
	}
	return b[uint(len(b))-n:]
}
