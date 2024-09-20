// package address provides a common interface for address types across different blockchains.
package address

import (
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/address/common"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/address/internal/registry"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	// register evm with registry
	_ "github.com/smartcontractkit/chainlink-ccip/internal/libs/address/internal/evm"
)

// MakeAddress initializes an unknown address type, it knows how to map chainSel to the correct
func MakeAddress(
	data []byte,
	chainSel ccipocr3.ChainSelector,
) (common.Address, error) {
	return registry.MakeAddress(data, chainSel)
}

// MakeAndEncodeAddress initializes an unknown address type and Encodes it.
func MakeAndEncodeAddress(
	data []byte,
	chainSel ccipocr3.ChainSelector,
) (common.EncodedAddress, error) {
	addr, err := MakeAddress(data, chainSel)
	if err != nil {
		return nil, err
	}

	return addr.Encode(), nil
}

func MakeEncodedAddress(
	data string,
	chainSel ccipocr3.ChainSelector,
) (common.EncodedAddress, error) {
	return registry.MakeEncodedAddress(data, chainSel)
}

func MakeAndDecodeEncodedAddress(
	data string,
	chainSel ccipocr3.ChainSelector,
) (common.Address, error) {
	addr, err := MakeEncodedAddress(data, chainSel)
	if err != nil {
		return nil, err
	}

	return addr.Decode()
}
