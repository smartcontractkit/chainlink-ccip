package tokens

import (
	"math"
	"math/big"
	"testing"

	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
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

// TestErrInboundRateLimitNotPortable pins the message that the v1.5.0, v1.5.1 and v1.6.x configure
// sequences surface when a lane is retargeted to a different remote token while its inbound rate
// limiter is enabled. Which pools reach it is decided by DoesPoolUseLocalDecimals (see the cases in
// TestDoesPoolUseLocalDecimals): only pools that denominate inbound in remote decimals are affected.
func TestErrInboundRateLimitNotPortable(t *testing.T) {
	err := ErrInboundRateLimitNotPortable(1234, big.NewInt(110_000_000), big.NewInt(11_000_000))
	require.Error(t, err)

	// The operator needs to know which lane, what is currently stored, and what to do about it.
	require.ErrorContains(t, err, "remote chain 1234")
	require.ErrorContains(t, err, "capacity 110000000")
	require.ErrorContains(t, err, "rate 11000000")
	require.ErrorContains(t, err, "cannot be reinterpreted for the new one")
	require.ErrorContains(t, err, "Specify the rate limits explicitly")
}

func TestRateLimiterConfigFloatInput_ValidatePrecision(t *testing.T) {
	tests := []struct {
		name    string
		rl      RateLimiterConfigFloatInput
		wantErr string
	}{
		{name: "over-precision capacity", rl: RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1.123456789, Rate: 1}, wantErr: "capacity 1.123456789 has 9 decimal places but the token has only 8"},
		{name: "capacity truncating to zero", rl: RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 0.000000001, Rate: 0}, wantErr: "capacity 1e-09 has 9 decimal places but the token has only 8"},
		{name: "over-precision rate", rl: RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 10, Rate: 0.000000001}, wantErr: "rate 1e-09 has 9 decimal places but the token has only 8"},
		{name: "within precision", rl: RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 0.29, Rate: 0.00000003}},
		{name: "exactly at precision", rl: RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1.12345678, Rate: 0.00000001}},
		{name: "disabled is skipped", rl: RateLimiterConfigFloatInput{IsEnabled: false, Capacity: 1.123456789, Rate: 0.000000001}},
		{name: "NaN capacity", rl: RateLimiterConfigFloatInput{IsEnabled: true, Capacity: math.NaN(), Rate: 1}, wantErr: "capacity NaN is not a finite number"},
		{name: "+Inf capacity", rl: RateLimiterConfigFloatInput{IsEnabled: true, Capacity: math.Inf(1), Rate: 1}, wantErr: "capacity +Inf is not a finite number"},
		{name: "-Inf rate", rl: RateLimiterConfigFloatInput{IsEnabled: true, Capacity: 1, Rate: math.Inf(-1)}, wantErr: "rate -Inf is not a finite number"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rl.ValidatePrecision(8)
			if tt.wantErr == "" {
				require.NoError(t, err)
				return
			}
			require.EqualError(t, err, tt.wantErr)
		})
	}
}

func TestBuildTPRLRemotesForSetRateLimitsLane_Precision(t *testing.T) {
	const localSel, remoteSel = uint64(1), uint64(2)
	enabled := func(capacity, rate float64) *RateLimiterConfigFloatInput {
		return &RateLimiterConfigFloatInput{IsEnabled: true, Capacity: capacity, Rate: rate}
	}
	poolRef := func(version *semver.Version) datastore.AddressRef {
		return datastore.AddressRef{Type: datastore.ContractType(utils.BurnMintTokenPool), Version: version}
	}

	tests := []struct {
		name           string
		poolRef        datastore.AddressRef
		localDecimals  uint8
		remoteDecimals uint8
		local          RemoteOutbounds
		remote         RemoteOutbounds
		wantErr        string
	}{
		{
			name:          "over-precision outbound capacity rejected",
			poolRef:       poolRef(utils.Version_1_6_1),
			localDecimals: 8, remoteDecimals: 8,
			local:   RemoteOutbounds{RateLimit: enabled(1.123456789, 1)},
			remote:  RemoteOutbounds{RateLimit: enabled(10, 1)},
			wantErr: "outbound rate limit for chain selector 1 and remote selector 2: capacity 1.123456789 has 9 decimal places but the token has only 8",
		},
		{
			name:          "inbound zero-truncating capacity and rate rejected",
			poolRef:       poolRef(utils.Version_1_6_1),
			localDecimals: 8, remoteDecimals: 18,
			local:   RemoteOutbounds{RateLimit: enabled(10, 1)},
			remote:  RemoteOutbounds{RateLimit: enabled(0.000000001, 0.000000001)},
			wantErr: "inbound rate limit for chain selector 1 and remote selector 2: capacity 1e-09 has 9 decimal places but the token has only 8",
		},
		{
			name:          "fast-finality outbound rejected",
			poolRef:       poolRef(utils.Version_1_6_1),
			localDecimals: 6, remoteDecimals: 6,
			local:   RemoteOutbounds{Outbounds: []RateLimitConfig{{RateLimit: *enabled(1, 0.1234567), FastFinality: true}}},
			remote:  RemoteOutbounds{Outbounds: []RateLimitConfig{{RateLimit: *enabled(1, 1), FastFinality: true}}},
			wantErr: "fast-finality outbound rate limit for chain selector 1 and remote selector 2: rate 0.1234567 has 7 decimal places but the token has only 6",
		},
		{
			// Pre-1.6.1 EVM pools scale inbound by remote decimals, so 9 places fit an 18-decimal remote.
			name:          "inbound checked against remote decimals for pre-1.6.1 EVM pool",
			poolRef:       poolRef(utils.Version_1_5_1),
			localDecimals: 8, remoteDecimals: 18,
			local:  RemoteOutbounds{RateLimit: enabled(10, 1)},
			remote: RemoteOutbounds{RateLimit: enabled(1.123456789, 1)},
		},
		{
			name:          "outboundOnly ignores unused remote bucket",
			poolRef:       poolRef(utils.Version_1_6_1),
			localDecimals: 8, remoteDecimals: 8,
			local:  RemoteOutbounds{RateLimit: enabled(10, 1), OutboundOnly: true},
			remote: RemoteOutbounds{RateLimit: enabled(1.123456789, 1)},
		},
		{
			name:          "valid precision passes",
			poolRef:       poolRef(utils.Version_1_6_1),
			localDecimals: 8, remoteDecimals: 8,
			local:  RemoteOutbounds{RateLimit: enabled(0.29, 0.00000003)},
			remote: RemoteOutbounds{RateLimit: enabled(100, 1.5)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _, buckets, err := buildTPRLRemotesForSetRateLimitsLane(
				chain_selectors.FamilyEVM, tt.poolRef,
				localSel, tt.localDecimals, tt.local,
				remoteSel, tt.remoteDecimals, tt.remote,
			)
			if tt.wantErr == "" {
				require.NoError(t, err)
				require.NotEmpty(t, buckets)
				return
			}
			require.EqualError(t, err, tt.wantErr)
			require.Nil(t, buckets)
		})
	}
}
