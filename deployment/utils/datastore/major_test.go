package datastore

import (
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

func TestFindRefByMajor(t *testing.T) {
	const sel = uint64(16423721717087811551)
	mk := func(version, address, qualifier string) datastore.AddressRef {
		return datastore.AddressRef{ChainSelector: sel, Type: "Router", Version: semver.MustParse(version), Address: address, Qualifier: qualifier}
	}
	build := func(refs ...datastore.AddressRef) datastore.DataStore {
		ds := datastore.NewMemoryDataStore()
		for _, r := range refs {
			require.NoError(t, ds.Addresses().Add(r))
		}
		return ds.Seal()
	}
	query := func(version string) datastore.AddressRef {
		return datastore.AddressRef{ChainSelector: sel, Type: "Router", Version: semver.MustParse(version)}
	}

	t.Run("a program upgraded in place resolves per major", func(t *testing.T) {
		ds := build(mk("1.6.0", "a", ""), mk("2.0.0", "a", ""), mk("2.0.1", "a", "x"))
		got, err := FindRefByMajor(ds, query("1.6.0"))
		require.NoError(t, err)
		require.Equal(t, "1.6.0", got.Version.String())

		got, err = FindRefByMajor(ds, query("2.0.0"))
		require.NoError(t, err)
		require.Equal(t, "2.0.1", got.Version.String(), "the greatest 2.x ref wins")
	})
	t.Run("a 1.x lookup does not find a 2.x ref", func(t *testing.T) {
		_, err := FindRefByMajor(build(mk("2.0.1", "a", "")), query("1.6.0"))
		require.ErrorIs(t, err, ErrNoRefOnMajor)
	})
	t.Run("the qualifier still narrows the lookup", func(t *testing.T) {
		q := query("2.0.0")
		q.Qualifier = "x"
		got, err := FindRefByMajor(build(mk("2.0.0", "a", ""), mk("2.0.1", "b", "x")), q)
		require.NoError(t, err)
		require.Equal(t, "b", got.Address)
	})
	t.Run("different addresses at the highest version are ambiguous", func(t *testing.T) {
		_, err := FindRefByMajor(build(mk("2.0.1", "a", ""), mk("2.0.1", "b", "x")), query("2.0.0"))
		require.ErrorContains(t, err, "ambiguous")
	})
	t.Run("a version is required", func(t *testing.T) {
		_, err := FindRefByMajor(build(), datastore.AddressRef{ChainSelector: sel, Type: "Router"})
		require.ErrorContains(t, err, "major version is required")
	})
	t.Run("formats the found ref", func(t *testing.T) {
		got, err := FindAndFormatRefByMajor(build(mk("2.0.1", "a", "")), query("2.0.0"), sel, func(r datastore.AddressRef) (string, error) { return r.Address, nil })
		require.NoError(t, err)
		require.Equal(t, "a", got)
	})
}
