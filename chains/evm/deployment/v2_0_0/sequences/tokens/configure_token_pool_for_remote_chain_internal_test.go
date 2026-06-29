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

func Test_normalizeInboundRateLimiterConfig(t *testing.T) {
	tests := []struct {
		name         string
		capacity     *big.Int
		rate         *big.Int
		fromDecimals uint8
		toDecimals   uint8
		wantCapacity *big.Int
		wantRate     *big.Int
	}{
		{
			name:         "18 remote decimals to 6 local decimals divides by 1e12",
			capacity:     new(big.Int).Mul(big.NewInt(50_000), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)),
			rate:         new(big.Int).Mul(big.NewInt(100), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)),
			fromDecimals: 18,
			toDecimals:   6,
			wantCapacity: new(big.Int).Mul(big.NewInt(50_000), new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)),
			wantRate:     new(big.Int).Mul(big.NewInt(100), new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)),
		},
		{
			name:         "6 remote decimals to 18 local decimals multiplies by 1e12",
			capacity:     new(big.Int).Mul(big.NewInt(50_000), new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)),
			rate:         new(big.Int).Mul(big.NewInt(100), new(big.Int).Exp(big.NewInt(10), big.NewInt(6), nil)),
			fromDecimals: 6,
			toDecimals:   18,
			wantCapacity: new(big.Int).Mul(big.NewInt(50_000), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)),
			wantRate:     new(big.Int).Mul(big.NewInt(100), new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)),
		},
		{
			name:         "equal decimals returns unchanged values",
			capacity:     big.NewInt(1000),
			rate:         big.NewInt(10),
			fromDecimals: 18,
			toDecimals:   18,
			wantCapacity: big.NewInt(1000),
			wantRate:     big.NewInt(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tokens.RateLimiterConfig{
				IsEnabled: true,
				Capacity:  tt.capacity,
				Rate:      tt.rate,
			}
			got := normalizeInboundRateLimiterConfig(cfg, tt.fromDecimals, tt.toDecimals)
			require.Equal(t, true, got.IsEnabled)
			require.Equal(t, 0, tt.wantCapacity.Cmp(got.Capacity), "capacity mismatch: want %s got %s", tt.wantCapacity, got.Capacity)
			require.Equal(t, 0, tt.wantRate.Cmp(got.Rate), "rate mismatch: want %s got %s", tt.wantRate, got.Rate)
			// original cfg must not be mutated
			require.Equal(t, 0, tt.capacity.Cmp(cfg.Capacity), "original capacity was mutated")
		})
	}
}
