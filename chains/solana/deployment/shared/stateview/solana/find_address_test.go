package solana

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

const (
	offRampA = "4HXmLxqCJvxBYnBLMVhRrJoZDrmCvcgYnF2Bo5Rj8cAq"
	offRampB = "9o9vS5dHHQLaZLv8gHuNu6k6J5HjisF9ravgRZigiDkb"
)

func tv(t string, v string) cldf.TypeAndVersion {
	return cldf.NewTypeAndVersion(cldf.ContractType(t), *semver.MustParse(v))
}

func TestFindSolanaAddress(t *testing.T) {
	tests := []struct {
		name  string
		want  cldf.TypeAndVersion
		have  map[string]cldf.TypeAndVersion
		found string
	}{
		{
			name:  "exact match wins",
			want:  tv("OffRamp", "1.6.0"),
			have:  map[string]cldf.TypeAndVersion{offRampA: tv("OffRamp", "1.6.0")},
			found: offRampA,
		},
		{
			// The case the audit found in every environment: the row is keyed at the
			// generation, the chain is on a later patch.
			name:  "generation fallback finds a later patch",
			want:  tv("OffRamp", "1.6.0"),
			have:  map[string]cldf.TypeAndVersion{offRampA: tv("OffRamp", "1.6.3")},
			found: offRampA,
		},
		{
			name:  "exact match preferred over a generation sibling",
			want:  tv("OffRamp", "1.6.0"),
			have:  map[string]cldf.TypeAndVersion{offRampA: tv("OffRamp", "1.6.3"), offRampB: tv("OffRamp", "1.6.0")},
			found: offRampB,
		},
		{
			// Two rows in the same generation: choosing either would return a
			// plausible address for the wrong program.
			name:  "ambiguous generation match returns nothing",
			want:  tv("OffRamp", "1.6.0"),
			have:  map[string]cldf.TypeAndVersion{offRampA: tv("OffRamp", "1.6.2"), offRampB: tv("OffRamp", "1.6.3")},
			found: "",
		},
		{
			name:  "different type never matches",
			want:  tv("OffRamp", "1.6.0"),
			have:  map[string]cldf.TypeAndVersion{offRampA: tv("Router", "1.6.3")},
			found: "",
		},
		{
			name:  "different minor is a different generation",
			want:  tv("OffRamp", "1.6.0"),
			have:  map[string]cldf.TypeAndVersion{offRampA: tv("OffRamp", "1.5.3")},
			found: "",
		},
		{
			// 0.1.0-dev is a miskeyed generation, not a patch of 1.6.0. These rows
			// need re-keying; the lookup must not paper over them.
			name:  "prerelease is not a patch of the release",
			want:  tv("OffRamp", "1.6.0"),
			have:  map[string]cldf.TypeAndVersion{offRampA: tv("OffRamp", "0.1.0-dev")},
			found: "",
		},
		{
			name:  "prerelease matches its own generation",
			want:  tv("BurnMintTokenPool", "0.1.0-dev"),
			have:  map[string]cldf.TypeAndVersion{offRampA: tv("BurnMintTokenPool", "0.1.2-dev")},
			found: offRampA,
		},
		{
			name:  "no addresses at all",
			want:  tv("OffRamp", "1.6.0"),
			have:  map[string]cldf.TypeAndVersion{},
			found: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindSolanaAddress(tt.want, tt.have)
			if tt.found == "" {
				require.True(t, got.IsZero(), "expected no match, got %s", got)
				return
			}
			require.Equal(t, solana.MustPublicKeyFromBase58(tt.found), got)
		})
	}
}
