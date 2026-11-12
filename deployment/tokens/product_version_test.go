package tokens

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"
)

func TestGetTokenAdapterFallsBackWithinAMajorLine(t *testing.T) {
	reg := newTokenAdapterRegistry()
	reg.RegisterTokenAdapter("solana", semver.MustParse("1.6.0"), nil)
	reg.RegisterTokenAdapter("solana", semver.MustParse("2.0.0"), nil)

	for version, want := range map[string]bool{
		"1.6.0": true,
		"2.0.0": true,
		"2.0.1": true,
		"1.6.1": false,
		"3.0.0": false,
	} {
		_, ok := reg.GetTokenAdapter("solana", semver.MustParse(version))
		require.Equal(t, want, ok, version)
	}
	_, ok := reg.GetTokenAdapter("evm", semver.MustParse("2.0.1"))
	require.False(t, ok, "the fallback never crosses chain families")
}
