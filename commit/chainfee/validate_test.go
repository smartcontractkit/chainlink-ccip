package chainfee

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func Test_validateObservedGasPrices(t *testing.T) {
	testCases := []struct {
		name      string
		gasPrices []cciptypes.GasPriceChain
		expErr    bool
	}{
		{
			name:      "empty is valid",
			gasPrices: []cciptypes.GasPriceChain{},
			expErr:    false,
		},
		{
			name: "all valid",
			gasPrices: []cciptypes.GasPriceChain{
				cciptypes.NewGasPriceChain(big.NewInt(10), 1),
				cciptypes.NewGasPriceChain(big.NewInt(20), 2),
				cciptypes.NewGasPriceChain(big.NewInt(1312), 3),
			},
			expErr: false,
		},
		{
			name: "duplicate gas price",
			gasPrices: []cciptypes.GasPriceChain{
				cciptypes.NewGasPriceChain(big.NewInt(10), 1),
				cciptypes.NewGasPriceChain(big.NewInt(20), 2),
				cciptypes.NewGasPriceChain(big.NewInt(1312), 1), // notice we already have a gas price for chain 1
			},
			expErr: true,
		},
		{
			name: "empty gas price",
			gasPrices: []cciptypes.GasPriceChain{
				cciptypes.NewGasPriceChain(big.NewInt(10), 1),
				cciptypes.NewGasPriceChain(big.NewInt(20), 2),
				cciptypes.NewGasPriceChain(nil, 3), // nil
			},
			expErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObservedGasPrices(tc.gasPrices)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
