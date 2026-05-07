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
