package evm

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/address/common"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/address/internal/registry"
)

func init() {
	registry.RegisterConstructors(1,
		func(data []byte) common.Address {
			return Address(data)
		},
		func(data string) common.EncodedAddress {
			return EncodedAddress(data)
		},
	)
}

type Address []byte

func (sa Address) Encode() (common.EncodedAddress, error) {
	// TODO: not EIP-55. Fix this?
	return EncodedAddress("0x" + hex.EncodeToString(sa)), nil
}

type EncodedAddress string

func (esa EncodedAddress) Decode() (common.Address, error) {
	// lower case in case EIP-55 and trim 0x prefix if there
	addrBytes, err := hex.DecodeString(strings.ToLower(strings.TrimPrefix(string(esa), "0x")))
	if err != nil {
		return nil, fmt.Errorf("failed to decode EVM address '%s': %w", esa, err)
	}

	return Address(addrBytes), nil
}
