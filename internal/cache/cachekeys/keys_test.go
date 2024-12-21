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

func TestFeeQuoterTokenUpdate(t *testing.T) {
	testCases := []struct {
		name        string
		token       ccipocr3.UnknownEncodedAddress
		chain       ccipocr3.ChainSelector
		expectedKey string
	}{
		{
			name:        "basic key generation",
			token:       ccipocr3.UnknownEncodedAddress("0x1234"),
			chain:       1,
			expectedKey: "fee-quoter-update:1:0x1234",
		},
		{
			name:        "empty token address",
			token:       ccipocr3.UnknownEncodedAddress(""),
			chain:       1,
			expectedKey: "fee-quoter-update:1:",
		},
		{
			name:        "zero chain selector",
			token:       ccipocr3.UnknownEncodedAddress("0x1234"),
			chain:       0,
			expectedKey: "fee-quoter-update:0:0x1234",
		},
		{
			name:        "empty token and zero chain",
			token:       ccipocr3.UnknownEncodedAddress(""),
			chain:       0,
			expectedKey: "fee-quoter-update:0:",
		},
		{
			name:        "long token address and large chain id",
			token:       ccipocr3.UnknownEncodedAddress("0x1234567890abcdef1234567890abcdef12345678"),
			chain:       999999999,
			expectedKey: "fee-quoter-update:999999999:0x1234567890abcdef1234567890abcdef12345678",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			key := FeeQuoterTokenUpdate(tc.token, tc.chain)
			assert.Equal(t, tc.expectedKey, key)
		})
	}
}
