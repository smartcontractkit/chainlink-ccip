package tokenprice

import (
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_validateObservedTokenPrices(t *testing.T) {
	testCases := []struct {
		name          string
		tokenPrices   []cciptypes.TokenPrice
		tokensToQuery map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo
		expErr        bool
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
			tokensToQuery: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
				"0x1": {},
				"0x2": {},
				"0x3": {},
				"0xa": {},
			},
			expErr: false,
		},
		{
			name: "nil price",
			tokenPrices: []cciptypes.TokenPrice{
				cciptypes.NewTokenPrice("0x1", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x2", big.NewInt(1)),
				cciptypes.NewTokenPrice("0x3", nil), // nil price
				cciptypes.NewTokenPrice("0xa", big.NewInt(1)),
			},
			tokensToQuery: map[cciptypes.UnknownEncodedAddress]pluginconfig.TokenInfo{
				"0x1": {},
				"0x2": {},
				"0x3": {},
				"0xa": {},
			},
			expErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tokenPrices := make(cciptypes.TokenPriceMap)
			for _, tp := range tc.tokenPrices {
				tokenPrices[tp.TokenID] = tp.Price
			}
			err := validateObservedTokenPrices(tokenPrices, tc.tokensToQuery)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})

	}
}
