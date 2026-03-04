package tokens

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	token_pool_v161 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
)

func Test_tokenBucketToRateLimiterConfig(t *testing.T) {
	tests := []struct {
		name string
		b    token_pool_v161.TokenBucket
		want *tokens.RateLimiterConfig
	}{
		{
			name: "enabled with capacity and rate",
			b: token_pool_v161.TokenBucket{
				IsEnabled: true,
				Capacity:  big.NewInt(1000),
				Rate:      big.NewInt(100),
			},
			want: &tokens.RateLimiterConfig{
				IsEnabled: true,
				Capacity:  big.NewInt(1000),
				Rate:      big.NewInt(100),
			},
		},
		{
			name: "disabled",
			b: token_pool_v161.TokenBucket{
				IsEnabled: false,
				Capacity:  big.NewInt(0),
				Rate:      big.NewInt(0),
			},
			want: &tokens.RateLimiterConfig{
				IsEnabled: false,
				Capacity:  big.NewInt(0),
				Rate:      big.NewInt(0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tokenBucketToRateLimiterConfig(tt.b)
			require.NotNil(t, got)
			require.Equal(t, tt.want.IsEnabled, got.IsEnabled)
			require.Equal(t, 0, tt.want.Capacity.Cmp(got.Capacity))
			require.Equal(t, 0, tt.want.Rate.Cmp(got.Rate))
		})
	}
}

