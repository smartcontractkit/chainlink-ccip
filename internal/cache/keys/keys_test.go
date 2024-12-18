package cachekeys

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func TestTokenDecimals(t *testing.T) {
	testCases := []struct {
		name        string
		token       ccipocr3.UnknownEncodedAddress
		address     string
		expectedKey string
	}{
		{
			name:        "basic key generation",
			token:       ccipocr3.UnknownEncodedAddress("0x1234"),
			address:     "0xabcd",
			expectedKey: "token-decimals:0x1234:0xabcd",
		},
		{
			name:        "empty token address",
			token:       ccipocr3.UnknownEncodedAddress(""),
			address:     "0xabcd",
			expectedKey: "token-decimals::0xabcd",
		},
		{
			name:        "empty contract address",
			token:       ccipocr3.UnknownEncodedAddress("0x1234"),
			address:     "",
			expectedKey: "token-decimals:0x1234:",
		},
		{
			name:        "both addresses empty",
			token:       ccipocr3.UnknownEncodedAddress(""),
			address:     "",
			expectedKey: "token-decimals::",
		},
		{
			name:        "long addresses",
			token:       ccipocr3.UnknownEncodedAddress("0x1234567890abcdef1234567890abcdef12345678"),
			address:     "0xfedcba0987654321fedcba0987654321fedcba09",
			expectedKey: "token-decimals:0x1234567890abcdef1234567890abcdef12345678:0xfedcba0987654321fedcba0987654321fedcba09",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			key := TokenDecimals(tc.token, tc.address)
			assert.Equal(t, tc.expectedKey, key)
		})
	}
}
