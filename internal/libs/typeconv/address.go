package typconv

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/gagliardetto/solana-go"
	chainsel "github.com/smartcontractkit/chain-selectors"
)

// AddressBytesToString converts the given address bytes to a string
// based upon the given chain selector's chain family.
func AddressBytesToString(addr []byte, chainSelector uint64) string {
	family, err := chainsel.GetSelectorFamily(chainSelector)
	if err != nil {
		panic(err)
	}

	switch family {
	case chainsel.FamilyEVM:
		// TODO: not EIP-55. Fix this?
		return "0x" + hex.EncodeToString(addr)
	case chainsel.FamilySolana:
		return solana.PublicKeyFromBytes(addr).String()
	default:
		panic("unsupported chain family")
	}
}

// AddressStringToBytes converts the given address string to bytes
// based upon the given chain selector's chain family.
func AddressStringToBytes(addr string, chainSelector uint64) ([]byte, error) {
	family, err := chainsel.GetSelectorFamily(chainSelector)
	if err != nil {
		panic(err)
	}

	switch family {
	case chainsel.FamilyEVM:
		// lower case in case EIP-55 and trim 0x prefix if there
		addrBytes, err := hex.DecodeString(strings.ToLower(strings.TrimPrefix(addr, "0x")))
		if err != nil {
			return nil, fmt.Errorf("failed to decode EVM address '%s': %w", addr, err)
		}

		return addrBytes, nil
	case chainsel.FamilySolana:
		pk, err := solana.PublicKeyFromBase58(addr)
		if err != nil {
			return nil, fmt.Errorf("failed to decode SVM address '%s': %w", addr, err)
		}
		return pk.Bytes(), nil
	default:
		panic("unsupported chain family")
	}
}

// KeepNRightBytes returns the last n bytes of the given byte slice.
// Example: KeepNRightBytes([]byte{0x01, 0x02, 0x03, 0x04}, 2) -> []byte{0x03, 0x04}
func KeepNRightBytes(b []byte, n uint) []byte {
	if n >= uint(len(b)) {
		return b
	}
	return b[uint(len(b))-n:]
}
