package registry

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/address/common"
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type AddressConstructor func(data []byte) (common.Address, error)
type EncodedAddressConstructor func(data string) (common.EncodedAddress, error)

type constructors struct {
	addressConstructor        AddressConstructor
	encodedAddressConstructor EncodedAddressConstructor
}

var chainConstructors map[ccipocr3.ChainSelector]constructors

func init() {
	chainConstructors = make(map[ccipocr3.ChainSelector]constructors)
}

func RegisterConstructors(chainSel ccipocr3.ChainSelector, addressConstructor AddressConstructor, encodedAddressConstructor EncodedAddressConstructor) {
	if _, exists := chainConstructors[chainSel]; exists {
		panic(fmt.Sprintf("Constructors for %d are already registered.", chainSel))
	}

	chainConstructors[chainSel] = constructors{
		addressConstructor:        addressConstructor,
		encodedAddressConstructor: encodedAddressConstructor,
	}
}

func MakeAddress(data []byte, chainSel ccipocr3.ChainSelector) (common.Address, error) {
	constructors, exists := chainConstructors[chainSel]
	if !exists {
		return nil, fmt.Errorf("no constructors registered for chain %d", chainSel)
	}

	return constructors.addressConstructor(data)
}

func MakeEncodedAddress(data string, chainSel ccipocr3.ChainSelector) (common.EncodedAddress, error) {
	constructors, exists := chainConstructors[chainSel]
	if !exists {
		return nil, fmt.Errorf("no constructors registered for chain %d", chainSel)
	}

	return constructors.encodedAddressConstructor(data)
}
