package evm

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/address/common"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/address/registry"
)

func init() {
	registry.RegisterConstructors(1,
		func(data []byte) (common.Address, error) {
			return Address(data), nil
		},
		func(data string) (common.EncodedAddress, error) {
			return EncodedAddress(data), nil
		},
	)
}

type Address []byte

func (a Address) Encode() common.EncodedAddress {
	// TODO: not EIP-55. Fix this?
	return EncodedAddress("0x" + hex.EncodeToString(a))
}

func (a Address) Bytes() []byte {
	return a
}

type EncodedAddress string

func (ea EncodedAddress) Decode() (common.Address, error) {
	// lower case in case EIP-55 and trim 0x prefix if there
	addrBytes, err := hex.DecodeString(strings.ToLower(strings.TrimPrefix(string(ea), "0x")))
	if err != nil {
		return nil, fmt.Errorf("failed to decode EVM address '%s': %w", ea, err)
	}

	return Address(addrBytes), nil
}

func (ea EncodedAddress) String() string {
	return string(ea)
}
