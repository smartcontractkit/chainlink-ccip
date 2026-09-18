package solutils

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

func TestTypeVersionLabelRoundTrip(t *testing.T) {
	tv := TypeVersion{Name: "ccip-router", Version: *semver.MustParse("1.6.2")}
	label := TypeVersionLabel(tv)
	require.Equal(t, "type_version:ccip-router 1.6.2", label)

	got, ok := ParseTypeVersionLabel([]string{"unrelated", label})
	require.True(t, ok)
	require.Equal(t, "ccip-router 1.6.2", got)
}

func TestTypeVersionLabelPreReleaseRoundTrip(t *testing.T) {
	// The -dev builds are the reason this exists; flattening them is how the datastore came to
	// claim 22 mainnet pools were 1.6.0.
	tv := TypeVersion{Name: "burnmint-token-pool", Version: *semver.MustParse("0.1.0-dev")}
	got, ok := ParseTypeVersionLabel([]string{TypeVersionLabel(tv)})
	require.True(t, ok)
	require.Equal(t, "burnmint-token-pool 0.1.0-dev", got)
}

func TestParseTypeVersionLabelAbsent(t *testing.T) {
	_, ok := ParseTypeVersionLabel([]string{"something-else", "owner:mcms"})
	require.False(t, ok)
}

func TestLabelReplacementKeepsOtherLabels(t *testing.T) {
	// RecordOnChainVersion drops only its own label. Simulate the same filtering it does, so the
	// contract is pinned even though the RPC path itself needs a live cluster.
	tv := deployment.TypeAndVersion{
		Type:    "OffRamp",
		Version: *semver.MustParse("1.6.0"),
		Labels:  deployment.NewLabelSet("owner:mcms", "type_version:ccip-offramp 1.6.2"),
	}

	kept := []string{}
	for _, l := range tv.Labels.List() {
		if _, isTV := ParseTypeVersionLabel([]string{l}); !isTV {
			kept = append(kept, l)
		}
	}
	kept = append(kept, TypeVersionLabel(TypeVersion{Name: "ccip-offramp", Version: *semver.MustParse("1.6.3")}))

	require.ElementsMatch(t, []string{"owner:mcms", "type_version:ccip-offramp 1.6.3"}, kept)
}
