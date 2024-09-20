package evm

import (
	"encoding/hex"
	"fmt"
	"strings"
)

type EVMAddress []byte

func (sa EVMAddress) Encode() (EncodedEVMAddress, error) {
	// TODO: not EIP-55. Fix this?
	return EncodedEVMAddress("0x" + hex.EncodeToString(sa)), nil
}

type EncodedEVMAddress string

func (esa EncodedEVMAddress) Decode() (EVMAddress, error) {
	// lower case in case EIP-55 and trim 0x prefix if there
	addrBytes, err := hex.DecodeString(strings.ToLower(strings.TrimPrefix(string(esa), "0x")))
	if err != nil {
		return nil, fmt.Errorf("failed to decode EVM address '%s': %w", esa, err)
	}

	return addrBytes, nil
}
