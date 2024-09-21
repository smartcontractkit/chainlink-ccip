package test

import (
	"math"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/address/common"
	"github.com/smartcontractkit/chainlink-ccip/internal/libs/address/registry"
)

const TestChainSelector = ccipocr3.ChainSelector(math.MaxUint64)

func init() {
	registry.RegisterConstructors(TestChainSelector,
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
	return EncodedAddress(a)
}

func (a Address) Bytes() []byte {
	return a
}

type EncodedAddress string

func (ea EncodedAddress) Decode() (common.Address, error) {
	return Address(ea), nil
}

func (ea EncodedAddress) String() string {
	return string(ea)
}
