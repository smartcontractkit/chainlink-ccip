package adapters

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
)

var _ deployapi.AddressNormalizer = (*EVMAddressNormalizer)(nil)

// EVMAddressNormalizer provides EIP-55 checksummed hex for datastore-aligned EVM lookups.
type EVMAddressNormalizer struct{}

func (EVMAddressNormalizer) NormalizeAddress(addr string) (string, error) {
	if !common.IsHexAddress(addr) {
		return "", fmt.Errorf("address %q is not a valid hex address", addr)
	}
	return common.HexToAddress(addr).Hex(), nil
}

// BytesToString converts raw on-chain address bytes to an EIP-55 checksummed hex string. EVM addresses are
// 20 bytes; values read off a pool (e.g. a remote token/pool) are commonly left-padded to 32 bytes, and
// common.BytesToAddress takes the right-most 20 bytes, which is correct for that left-padded input.
func (EVMAddressNormalizer) BytesToString(address []byte) (string, error) {
	if len(address) == 0 {
		return "", fmt.Errorf("address bytes cannot be empty")
	}
	return common.BytesToAddress(address).Hex(), nil
}

// StringToBytes converts a hex address string to its 20-byte representation.
func (EVMAddressNormalizer) StringToBytes(address string) ([]byte, error) {
	if !common.IsHexAddress(address) {
		return nil, fmt.Errorf("address %q is not a valid hex address", address)
	}
	return common.HexToAddress(address).Bytes(), nil
}
