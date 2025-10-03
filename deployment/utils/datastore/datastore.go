package datastore

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

// FormatFn is a function that formats a datastore.AddressRef into a specific type T.
type FormatFn[T any] = func(ref datastore.AddressRef) (T, error)

// FindAndFormatEachRef queries the datastore for multiple AddressRefs.
// AddressRefs specified in the input slice may have only a subset of fields set (e.g. Type and Version).
// This function enforces that exactly one match is found for each provided AddressRef.
// It then formats each AddressRef into the desired type T using the provided FormatFn.
// Example usage: Find contract addresses via type and version, returning each as a native address type.
func FindAndFormatEachRef[T any](ds datastore.DataStore, refs []datastore.AddressRef, format FormatFn[T]) ([]T, error) {
	formattedRefs := make([]T, 0, len(refs))
	for _, ref := range refs {
		refFromStore, err := findSingleRef(ds, ref)
		if err != nil {
			return nil, err
		}
		formattedRef, err := format(refFromStore)
		if err != nil {
			return nil, fmt.Errorf("failed to format ref %s: %w", SprintRef(refFromStore), err)
		}
		formattedRefs = append(formattedRefs, formattedRef)
	}
	return formattedRefs, nil
}

// SprintRef returns a one-line string representation of a datastore.AddressRef for logging.
func SprintRef(ref datastore.AddressRef) string {
	return fmt.Sprintf("{ChainSelector: %d, Type: %s, Version: %s, Qualifier: %s, Address: %s}", ref.ChainSelector, ref.Type, ref.Version, ref.Qualifier, ref.Address)
}

// findSingleRef queries the datastore for an AddressRef matching a subset of fields provided by AddressRef.
// It enforces that exactly one match is found.
func findSingleRef(ds datastore.DataStore, ref datastore.AddressRef) (datastore.AddressRef, error) {
	filterFns := make([]datastore.FilterFunc[datastore.AddressRefKey, datastore.AddressRef], 0, 5)
	// Filter by largest scope (chain) to smallest scope (address)
	// Address is the smallest scope because there can only be one of each address on a given chain
	if ref.ChainSelector != 0 {
		filterFns = append(filterFns, datastore.AddressRefByChainSelector(ref.ChainSelector))
	}
	if ref.Type != "" {
		filterFns = append(filterFns, datastore.AddressRefByType(ref.Type))
	}
	if ref.Version != nil {
		filterFns = append(filterFns, datastore.AddressRefByVersion(ref.Version))
	}
	if ref.Qualifier != "" {
		filterFns = append(filterFns, datastore.AddressRefByQualifier(ref.Qualifier))
	}
	if ref.Address != "" {
		filterFns = append(filterFns, datastore.AddressRefByAddress(ref.Address))
	}
	refs := ds.Addresses().Filter(filterFns...)
	if len(refs) != 1 {
		return datastore.AddressRef{}, fmt.Errorf("expected to find exactly 1 ref with criteria %s, found %d", SprintRef(ref), len(refs))
	}
	return refs[0], nil
}

// FullRef returns the entire datastore.AddressRef
func FullRef(ref datastore.AddressRef) (datastore.AddressRef, error) {
	return ref, nil
}
