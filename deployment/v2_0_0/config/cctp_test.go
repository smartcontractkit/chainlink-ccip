package config

import (
	"testing"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"
)

func TestGetCCTPChainDefaults(t *testing.T) {
	tests := []struct {
		name       string
		chainSel   uint64
		want       CCTPChainDefaults
		wantExists bool
	}{
		{
			name:     "ethereum mainnet",
			chainSel: chain_selectors.ETHEREUM_MAINNET.Selector,
			want: CCTPChainDefaults{
				DomainIdentifier: 0,
				TokenMessengerV1: "0xBd3fa81B58Ba92a82136038B25aDec7066af3155",
				TokenMessengerV2: "0x28b5a0e9C621a5BadaA536219b3a228C8168cf5d",
				USDCToken:        "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
			},
			wantExists: true,
		},
		{
			name:     "avalanche fuji testnet",
			chainSel: chain_selectors.AVALANCHE_TESTNET_FUJI.Selector,
			want: CCTPChainDefaults{
				DomainIdentifier: 1,
				TokenMessengerV1: "0xeb08f243E5d3FCFF26A9E38Ae5520A669f4019d0",
				TokenMessengerV2: "0x8FE6B999Dc680CcFDD5Bf7EB0974218be2542DAA",
				USDCToken:        "0x5425890298aed601595a70AB815c96711a31Bc65",
			},
			wantExists: true,
		},
		{
			name:     "solana mainnet",
			chainSel: chain_selectors.SOLANA_MAINNET.Selector,
			want: CCTPChainDefaults{
				DomainIdentifier: 5,
				TokenMessengerV1: "CCTPiPYPc6AsJuwueEnWgSgucamXDZwBd53dQ11YiKX3",
				USDCToken:        "EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v",
			},
			wantExists: true,
		},
		{
			name:       "unsupported chain",
			chainSel:   1234567890,
			wantExists: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := GetCCTPChainDefaults(tt.chainSel)
			require.Equal(t, tt.wantExists, ok)
			if tt.wantExists {
				require.Equal(t, tt.want, got)
			}
		})
	}
}

func TestCCTPChainDefaultsBySelector_AllHaveUSDCToken(t *testing.T) {
	for chainSel, defaults := range CCTPChainDefaultsBySelector {
		require.NotEmptyf(t, defaults.USDCToken, "chain selector %d is missing a USDC token", chainSel)
	}
}
