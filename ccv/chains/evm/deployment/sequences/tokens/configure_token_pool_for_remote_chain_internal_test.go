package tokens

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	token_pool_v161 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/operations/token_pool"
	fqops_v163 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_3/operations/fee_quoter"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/gobindings/generated/v1_5_0/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"

	fqops "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/operations/fee_quoter"
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

func Test_v150OnRampConfigToTokenTransferFeeConfig(t *testing.T) {
	cfg := evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfig{
		MinFeeUSDCents:            50,
		MaxFeeUSDCents:            500,
		DeciBps:                   100,
		DestGasOverhead:           200_000,
		DestBytesOverhead:         32,
		AggregateRateLimitEnabled: false,
		IsEnabled:                 true,
	}
	got := v150OnRampConfigToTokenTransferFeeConfig(cfg)
	require.NotNil(t, got)
	require.Equal(t, uint32(200_000), got.DestGasOverhead)
	require.Equal(t, uint32(32), got.DestBytesOverhead)
	require.Equal(t, uint32(50), got.DefaultFinalityFeeUSDCents)
	require.Equal(t, uint32(0), got.CustomFinalityFeeUSDCents)
	require.Equal(t, uint16(100), got.DefaultFinalityTransferFeeBps)
	require.Equal(t, uint16(0), got.CustomFinalityTransferFeeBps)
	require.True(t, got.IsEnabled)
}

func Test_v163FeeQuoterConfigToTokenTransferFeeConfig(t *testing.T) {
	cfg := fqops_v163.TokenTransferFeeConfig{
		MinFeeUSDCents:    75,
		MaxFeeUSDCents:    750,
		DeciBps:           150,
		DestGasOverhead:   150_000,
		DestBytesOverhead: 64,
		IsEnabled:         true,
	}
	got := v163FeeQuoterConfigToTokenTransferFeeConfig(cfg)
	require.NotNil(t, got)
	require.Equal(t, uint32(150_000), got.DestGasOverhead)
	require.Equal(t, uint32(64), got.DestBytesOverhead)
	require.Equal(t, uint32(75), got.DefaultFinalityFeeUSDCents)
	require.Equal(t, uint32(0), got.CustomFinalityFeeUSDCents)
	require.Equal(t, uint16(150), got.DefaultFinalityTransferFeeBps)
	require.Equal(t, uint16(0), got.CustomFinalityTransferFeeBps)
	require.True(t, got.IsEnabled)
}

func Test_v2FeeQuoterConfigToTokenTransferFeeConfig(t *testing.T) {
	cfg := fqops.TokenTransferFeeConfig{
		FeeUSDCents:       100,
		DestGasOverhead:   180_000,
		DestBytesOverhead: 48,
		IsEnabled:         true,
	}
	got := v2FeeQuoterConfigToTokenTransferFeeConfig(cfg)
	require.NotNil(t, got)
	require.Equal(t, uint32(180_000), got.DestGasOverhead)
	require.Equal(t, uint32(48), got.DestBytesOverhead)
	require.Equal(t, uint32(100), got.DefaultFinalityFeeUSDCents)
	require.Equal(t, uint32(0), got.CustomFinalityFeeUSDCents)
	require.Equal(t, uint16(0), got.DefaultFinalityTransferFeeBps)
	require.Equal(t, uint16(0), got.CustomFinalityTransferFeeBps)
	require.True(t, got.IsEnabled)
}

func Test_mergeTokenTransferFeeConfig(t *testing.T) {
	tests := []struct {
		name     string
		desired  *tokens.TokenTransferFeeConfig
		imported *tokens.TokenTransferFeeConfig
		check    func(t *testing.T, got *tokens.TokenTransferFeeConfig)
	}{
		{
			name:     "nil imported returns desired",
			desired:  &tokens.TokenTransferFeeConfig{DestGasOverhead: 1, IsEnabled: true},
			imported: nil,
			check: func(t *testing.T, got *tokens.TokenTransferFeeConfig) {
				require.Equal(t, uint32(1), got.DestGasOverhead)
				require.True(t, got.IsEnabled)
			},
		},
		{
			name: "all desired zero uses imported",
			desired: &tokens.TokenTransferFeeConfig{
				IsEnabled: true,
			},
			imported: &tokens.TokenTransferFeeConfig{
				DestGasOverhead:               100,
				DestBytesOverhead:             50,
				DefaultFinalityFeeUSDCents:    10,
				CustomFinalityFeeUSDCents:     20,
				DefaultFinalityTransferFeeBps: 5,
				CustomFinalityTransferFeeBps:  6,
				IsEnabled:                     true,
			},
			check: func(t *testing.T, got *tokens.TokenTransferFeeConfig) {
				require.Equal(t, uint32(100), got.DestGasOverhead)
				require.Equal(t, uint32(50), got.DestBytesOverhead)
				require.Equal(t, uint32(10), got.DefaultFinalityFeeUSDCents)
				require.Equal(t, uint32(20), got.CustomFinalityFeeUSDCents)
				require.Equal(t, uint16(5), got.DefaultFinalityTransferFeeBps)
				require.Equal(t, uint16(6), got.CustomFinalityTransferFeeBps)
			},
		},
		{
			name: "non-zero desired overrides imported",
			desired: &tokens.TokenTransferFeeConfig{
				DestGasOverhead:              200,
				DefaultFinalityFeeUSDCents:   25,
				CustomFinalityTransferFeeBps: 10,
				IsEnabled:                    true,
			},
			imported: &tokens.TokenTransferFeeConfig{
				DestGasOverhead:               100,
				DestBytesOverhead:             50,
				DefaultFinalityFeeUSDCents:    10,
				CustomFinalityFeeUSDCents:     20,
				DefaultFinalityTransferFeeBps: 5,
				CustomFinalityTransferFeeBps:  6,
				IsEnabled:                     true,
			},
			check: func(t *testing.T, got *tokens.TokenTransferFeeConfig) {
				require.Equal(t, uint32(200), got.DestGasOverhead, "desired overrides")
				require.Equal(t, uint32(50), got.DestBytesOverhead, "imported when desired zero")
				require.Equal(t, uint32(25), got.DefaultFinalityFeeUSDCents, "desired overrides")
				require.Equal(t, uint32(20), got.CustomFinalityFeeUSDCents, "imported when desired zero")
				require.Equal(t, uint16(5), got.DefaultFinalityTransferFeeBps, "imported when desired zero")
				require.Equal(t, uint16(10), got.CustomFinalityTransferFeeBps, "desired overrides")
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTokenTransferFeeConfig(tt.desired, tt.imported)
			require.NotNil(t, got)
			tt.check(t, got)
		})
	}
}
