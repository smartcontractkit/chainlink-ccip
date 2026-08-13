package datastore

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	df "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFindAndFormatCanonicalRef(t *testing.T) {
	ds := df.NewMemoryDataStore()
	// Add a canonical ref (empty qualifier).
	canonical := df.AddressRef{
		ChainSelector: 1,
		Type:          "OnRamp",
		Version:       semver.MustParse("2.0.0"),
		Qualifier:     "",
		Address:       "0x1111",
	}
	require.NoError(t, ds.Addresses().Add(canonical))
	// Add a legacy ref (non-empty qualifier).
	legacy := df.AddressRef{
		ChainSelector: 1,
		Type:          "OnRamp",
		Version:       semver.MustParse("2.0.0"),
		Qualifier:     "legacy",
		Address:       "0x2222",
	}
	require.NoError(t, ds.Addresses().Add(legacy))

	sealed := ds.Seal()

	// Canonical lookup (empty qualifier) should match only the canonical ref.
	got, err := FindAndFormatCanonicalRef(sealed, df.AddressRef{
		Type:    "OnRamp",
		Version: semver.MustParse("2.0.0"),
	}, 1, FullRef)
	require.NoError(t, err)
	assert.Equal(t, "0x1111", got.Address)
	assert.Equal(t, "", got.Qualifier)

	// Legacy lookup should match only the legacy ref.
	got, err = FindAndFormatCanonicalRef(sealed, df.AddressRef{
		Type:      "OnRamp",
		Version:   semver.MustParse("2.0.0"),
		Qualifier: "legacy",
	}, 1, FullRef)
	require.NoError(t, err)
	assert.Equal(t, "0x2222", got.Address)
	assert.Equal(t, "legacy", got.Qualifier)

	// Regular FindAndFormatRef with no qualifier should fail (finds 2).
	_, err = FindAndFormatRef(sealed, df.AddressRef{
		Type:    "OnRamp",
		Version: semver.MustParse("2.0.0"),
	}, 1, FullRef)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "found 2")
}
