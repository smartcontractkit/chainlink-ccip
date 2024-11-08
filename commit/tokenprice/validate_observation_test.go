package tokenprice

import (
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_validateObservedTokenPrices(t *testing.T) {
	testCases := []struct {
		name        string
		tokenPrices []cciptypes.TokenPrice
		expErr      bool
	}{
		{
			name:        "empty is valid",
			tokenPrices: []cciptypes.TokenPrice{},
			expErr:      false,
		},
		{
			name: "all valid",
			tokenPrices: []cciptypes.TokenPrice{
				cciptypes.NewTokenPrice("0x1", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x2", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x3", big.NewInt(1)),
				cciptypes.NewTokenPrice("0xa", big.NewInt(1)),
			},
			expErr: false,
		},
		{
			name: "dup price",
			tokenPrices: []cciptypes.TokenPrice{
				cciptypes.NewTokenPrice("0x1", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x2", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x1", big.NewInt(1)), // dup
				cciptypes.NewTokenPrice("0xa", big.NewInt(1)),
			},
			expErr: true,
		},
		{
			name: "nil price",
			tokenPrices: []cciptypes.TokenPrice{
				cciptypes.NewTokenPrice("0x1", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x2", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x3", nil), // nil price
				cciptypes.NewTokenPrice("0xa", big.NewInt(1)),
			},
			expErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObservedTokenPrices(tc.tokenPrices)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})

	}
}

func Test_validateFChain(t *testing.T) {
	testCases := []struct {
		name   string
		fChain map[cciptypes.ChainSelector]int
		expErr bool
	}{
		{
			name: "FChain contains negative values",
			fChain: map[cciptypes.ChainSelector]int{
				1: 11,
				2: -4,
			},
			expErr: true,
		},
		{
			name: "FChain valid",
			fChain: map[cciptypes.ChainSelector]int{
				12: 6,
				7:  9,
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateFChain(tc.fChain)

			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}
