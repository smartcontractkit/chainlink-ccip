package tokens

import (
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func TestDoesPoolUseLocalDecimals(t *testing.T) {
	bnmExternal := utils.BurnMintWithExternalMinterTokenPool.String()
	hybExternal := utils.HybridWithExternalMinterTokenPool.String()
	stdBurnMint := utils.BurnMintTokenPool.String()
	v151 := utils.Version_1_5_1
	v160 := utils.Version_1_6_0
	v161 := utils.Version_1_6_1

	tests := []struct {
		name        string
		chainFamily string
		version     *semver.Version
		poolType    string
		want        bool
	}{
		{name: "nil version defaults local", chainFamily: chain_selectors.FamilyEVM, version: nil, poolType: stdBurnMint, want: true},
		{name: "non-EVM pre-1.6.1 uses local", chainFamily: chain_selectors.FamilySolana, version: v151, poolType: stdBurnMint, want: true},
		{name: "EVM 1.6.1+ uses local", chainFamily: chain_selectors.FamilyEVM, version: v161, poolType: stdBurnMint, want: true},
		{name: "EVM 1.5.1 standard uses remote", chainFamily: chain_selectors.FamilyEVM, version: v151, poolType: stdBurnMint, want: false},
		{name: "EVM 1.6.0 standard uses remote", chainFamily: chain_selectors.FamilyEVM, version: v160, poolType: stdBurnMint, want: false},
		{name: "EVM 1.6.0 burn mint external minter uses local", chainFamily: chain_selectors.FamilyEVM, version: v160, poolType: bnmExternal, want: true},
		{name: "EVM 1.6.0 hybrid external minter uses local", chainFamily: chain_selectors.FamilyEVM, version: v160, poolType: hybExternal, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, DoesPoolUseLocalDecimals(tt.chainFamily, tt.version, tt.poolType))
		})
	}
}

func TestRebaseRateLimiterConfig(t *testing.T) {
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
			cfg := RateLimiterConfig{
				IsEnabled: true,
				Capacity:  tt.capacity,
				Rate:      tt.rate,
			}
			got := RebaseRateLimiterConfig(cfg, tt.fromDecimals, tt.toDecimals)
			require.Equal(t, true, got.IsEnabled)
			require.Equal(t, 0, tt.wantCapacity.Cmp(got.Capacity), "capacity mismatch: want %s got %s", tt.wantCapacity, got.Capacity)
			require.Equal(t, 0, tt.wantRate.Cmp(got.Rate), "rate mismatch: want %s got %s", tt.wantRate, got.Rate)
			require.Equal(t, 0, tt.capacity.Cmp(cfg.Capacity), "original capacity was mutated")
		})
	}
}
