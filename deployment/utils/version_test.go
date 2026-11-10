package utils

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"
)

func TestResolveRegisteredVersion(t *testing.T) {
	v := semver.MustParse
	registered := []*semver.Version{v("1.5.0"), v("1.6.0"), v("2.0.0")}

	got, ok := ResolveRegisteredVersion(registered, v("1.6.0"))
	require.True(t, ok)
	require.Equal(t, "1.6.0", got.String())

	got, ok = ResolveRegisteredVersion(registered, v("2.0.1"))
	require.True(t, ok, "a 2.0.1 ref is served by the 2.0.0 registration")
	require.Equal(t, "2.0.0", got.String())

	_, ok = ResolveRegisteredVersion(registered, v("1.6.1"))
	require.False(t, ok, "versions below 2.0.0 never fall back to a neighbour")

	got, ok = ResolveRegisteredVersion([]*semver.Version{v("2.0.0"), v("2.1.0")}, v("2.0.5"))
	require.True(t, ok)
	require.Equal(t, "2.1.0", got.String(), "the highest registration on the line wins")

	_, ok = ResolveRegisteredVersion(registered, v("3.0.0"))
	require.False(t, ok)
	_, ok = ResolveRegisteredVersion(registered, nil)
	require.False(t, ok)
}
