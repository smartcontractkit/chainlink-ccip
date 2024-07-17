package typconv

import (
	"encoding/hex"
)

func AddressBytesToString(addr []byte) string {
	return "0x" + hex.EncodeToString(addr)
}
