package adapters

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
)

// TestEffectiveRateLimiter covers the four rows of the proxy/previous-pool effective-limit table:
// a disabled bucket imposes no constraint, so the enabled side wins outright, and when both sides
// are enabled the tighter (smaller) capacity/rate wins.
func TestEffectiveRateLimiter(t *testing.T) {
	t.Parallel()

	disabled := tokensapi.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}
	proxyEnabled := tokensapi.RateLimiterConfig{IsEnabled: true, Capacity: big.NewInt(100), Rate: big.NewInt(10)}
	previousEnabled := tokensapi.RateLimiterConfig{IsEnabled: true, Capacity: big.NewInt(50), Rate: big.NewInt(20)}

	tests := []struct {
		name     string
		proxy    tokensapi.RateLimiterConfig
		previous tokensapi.RateLimiterConfig
		want     tokensapi.RateLimiterConfig
	}{
		{
			name:     "both disabled -> disabled",
			proxy:    disabled,
			previous: disabled,
			want:     tokensapi.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)},
		},
		{
			name:     "only proxy enabled -> proxy wins",
			proxy:    proxyEnabled,
			previous: disabled,
			want:     proxyEnabled,
		},
		{
			name:     "only previous enabled -> previous wins",
			proxy:    disabled,
			previous: previousEnabled,
			want:     previousEnabled,
		},
		{
			name:     "both enabled -> tighter of each field wins",
			proxy:    proxyEnabled,
			previous: previousEnabled,
			want:     tokensapi.RateLimiterConfig{IsEnabled: true, Capacity: big.NewInt(50), Rate: big.NewInt(10)},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := effectiveRateLimiter(tc.proxy, tc.previous)
			require.Equal(t, tc.want.IsEnabled, got.IsEnabled)
			require.Equal(t, 0, tc.want.Capacity.Cmp(got.Capacity))
			require.Equal(t, 0, tc.want.Rate.Cmp(got.Rate))
		})
	}
}

func TestMinBigInt(t *testing.T) {
	t.Parallel()

	require.Equal(t, 0, big.NewInt(3).Cmp(minBigInt(big.NewInt(3), big.NewInt(5))))
	require.Equal(t, 0, big.NewInt(3).Cmp(minBigInt(big.NewInt(5), big.NewInt(3))))
	require.Equal(t, 0, big.NewInt(0).Cmp(minBigInt(nil, big.NewInt(5))))
	require.Equal(t, 0, big.NewInt(0).Cmp(minBigInt(big.NewInt(5), nil)))
	require.Equal(t, 0, big.NewInt(0).Cmp(minBigInt(nil, nil)))
}

// TestEffectiveMigrationRateLimits_NoProxyAddress asserts the short-circuit for callers that pass
// a zero-value proxy address: the reader-observed proxy limits are returned unmodified without any
// on-chain calls.
func TestEffectiveMigrationRateLimits_NoProxyAddress(t *testing.T) {
	t.Parallel()

	proxyOut := tokensapi.RateLimiterConfig{IsEnabled: true, Capacity: big.NewInt(1), Rate: big.NewInt(1)}
	proxyIn := tokensapi.RateLimiterConfig{IsEnabled: false, Capacity: big.NewInt(0), Rate: big.NewInt(0)}

	got, err := EffectiveMigrationRateLimits(
		cldf_ops.Bundle{}, testChain, common.Address{}, 1234, proxyOut, proxyIn, 18, 18,
	)
	require.NoError(t, err)
	require.Equal(t, tokensapi.OnchainRateLimits{Outbound: proxyOut, Inbound: proxyIn}, got)
}
