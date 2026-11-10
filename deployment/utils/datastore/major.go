package datastore

import (
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

// ErrNoRefOnMajor is returned by FindRefByMajor when no ref exists on the requested major line.
var ErrNoRefOnMajor = errors.New("no ref on major version")

// FindRefByMajor returns the ref matching ref's chain selector, type, qualifier and address whose
// version is on the same major line as ref.Version, choosing the highest version on that line.
// Callers name only the line they work with (1.x or 2.x); the datastore decides which release of
// it is deployed. This matters for programs upgraded in place, which can hold refs for 1.6.0 and
// 2.0.0, and later 2.0.1, at the same address.
//
// Several refs at the highest version must share an address, otherwise the lookup is ambiguous.
func FindRefByMajor(ds datastore.DataStore, ref datastore.AddressRef) (datastore.AddressRef, error) {
	if ref.Version == nil {
		return datastore.AddressRef{}, fmt.Errorf("a major version is required to look up ref %s", SprintRef(ref))
	}
	major := ref.Version.Major()
	criteria := ref
	criteria.Version = nil

	var best datastore.AddressRef
	found := false
	for _, r := range ds.Addresses().Filter(AddressRefToFilters(criteria)...) {
		if r.Version == nil || r.Version.Major() != major {
			continue
		}
		switch {
		case !found || r.Version.GreaterThan(best.Version):
			best, found = r, true
		case r.Version.Equal(best.Version) && r.Address != best.Address:
			return datastore.AddressRef{}, fmt.Errorf("ambiguous ref %s: %s and %s are both at version %s", SprintRef(ref), best.Address, r.Address, r.Version)
		}
	}
	if !found {
		return datastore.AddressRef{}, fmt.Errorf("%w %d: criteria %s", ErrNoRefOnMajor, major, SprintRef(ref))
	}
	return best, nil
}

// FindAndFormatRefByMajor is FindRefByMajor followed by format. It is the counterpart of
// FindAndFormatRef for callers that name only a major version.
func FindAndFormatRefByMajor[T any](ds datastore.DataStore, ref datastore.AddressRef, chainSelector uint64, format FormatFn[T]) (T, error) {
	var empty T
	ref.ChainSelector = chainSelector
	found, err := FindRefByMajor(ds, ref)
	if err != nil {
		return empty, err
	}
	formatted, err := format(found)
	if err != nil {
		return empty, fmt.Errorf("failed to format ref %s: %w", SprintRef(found), err)
	}
	return formatted, nil
}
