package sequences

import (
	"testing"

	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"
)

func TestGetDefaultTokenFeeUSDCents(t *testing.T) {
	tests := []struct {
		name        string
		sourceChain uint64
		remoteChain uint64
		want        uint16
	}{
		{
			name:        "solana destination with ethereum source uses highest solana default",
			sourceChain: chainsel.ETHEREUM_MAINNET.Selector,
			remoteChain: chainsel.SOLANA_DEVNET.Selector,
			want:        60,
		},
		{
			name:        "solana destination with non-ethereum source uses solana default",
			sourceChain: chainsel.TEST_90000001.Selector,
			remoteChain: chainsel.SOLANA_DEVNET.Selector,
			want:        35,
		},
		{
			name:        "ethereum destination uses ethereum remote default",
			sourceChain: chainsel.TEST_90000001.Selector,
			remoteChain: chainsel.ETHEREUM_MAINNET.Selector,
			want:        150,
		},
		{
			name:        "ethereum source with non-solana and non-ethereum destination uses ethereum source default",
			sourceChain: chainsel.ETHEREUM_MAINNET.Selector,
			remoteChain: chainsel.TEST_90000001.Selector,
			want:        50,
		},
		{
			name:        "non-ethereum source and non-solana destination use baseline default",
			sourceChain: chainsel.TEST_90000001.Selector,
			remoteChain: chainsel.TEST_90000002.Selector,
			want:        25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDefaultTokenFeeUSDCents(tt.sourceChain, tt.remoteChain)
			require.Equal(t, tt.want, got)
		})
	}
}
