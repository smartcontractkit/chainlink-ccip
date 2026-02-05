package datastore

import (
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

// FormatFn is a function that formats a datastore.AddressRef into a specific type T.
type FormatFn[T any] = func(ref datastore.AddressRef) (T, error)

// FindAndFormatRef queries the datastore for an AddressRef.
// The inputted AddressRef may have only a subset of fields set (e.g. Type and Version).
// This function enforces that exactly one match is found for the AddressRef.
// It then formats the AddressRef into the desired type T using the provided FormatFn.
// Example usage: Find a contract reference via type and version, returning it as a native address type.
func FindAndFormatRef[T any](ds datastore.DataStore, ref datastore.AddressRef, chainSelector uint64, format FormatFn[T]) (T, error) {
	var empty T
	// We set the chain selector here to ensure we are searching within the correct chain scope.
	// Chain selector is usually not provided in the ref since it is often implied by the context of the greater input.
	ref.ChainSelector = chainSelector
	refFromStore, _, err := findRef(ds, ref)
	if err != nil {
		return empty, err
	}
	formattedRef, err := format(refFromStore)
	if err != nil {
		return empty, fmt.Errorf("failed to format ref %s: %w", SprintRef(refFromStore), err)
	}

	return formattedRef, nil
}

// FindAndFormatFirstRef queries the datastore for multiple AddressRefs in order.
// The inputted AddressRefs may have only a subset of fields set (e.g. Type and Version).
// If none of the provided refs are found, an error is returned.
// If any of the refs return multiple matches, an error is returned.
// If a unique match is found for a ref, it is formatted into the desired type T using the provided FormatFn.
// Example usage: Find the first available contract reference for multiple contract types, returning it as a native address type.
func FindAndFormatFirstRef[T any](ds datastore.DataStore, chainSelector uint64, format FormatFn[T], refs ...datastore.AddressRef) (T, error) {
	var empty T
	if len(refs) == 0 {
		return empty, fmt.Errorf("at least one address ref must be specified")
	}

	for _, ref := range refs {
		ref.ChainSelector = chainSelector
		refFromStore, foundRefs, err := findRef(ds, ref)
		if foundRefs == 0 {
			// No refs found, try the next one
			continue
		}
		if err != nil {
			return empty, err
		}

		// Successfully found a unique ref
		formattedRef, err := format(refFromStore)
		if err != nil {
			return empty, fmt.Errorf("failed to format ref %s: %w", SprintRef(refFromStore), err)
		}
		return formattedRef, nil
	}

	// No refs found
	var refsString strings.Builder
	for _, r := range refs {
		refsString.WriteString(SprintRef(r))
		refsString.WriteString(",")
	}

	return empty, fmt.Errorf("failed to uniquely find any of the provided refs: %v", refsString.String())
}

// SprintRef returns a one-line string representation of a datastore.AddressRef for logging.
func SprintRef(ref datastore.AddressRef) string {
	return fmt.Sprintf("{ChainSelector: %d, Type: %s, Version: %s, Qualifier: %s, Address: %s}", ref.ChainSelector, ref.Type, ref.Version, ref.Qualifier, ref.Address)
}

// findRef queries the datastore for an AddressRef matching a subset of fields provided by AddressRef.
// It enforces that exactly one match is found.
func findRef(ds datastore.DataStore, ref datastore.AddressRef) (datastore.AddressRef, int, error) {
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
		return datastore.AddressRef{}, len(refs), fmt.Errorf("expected to find exactly 1 ref with criteria %s, found %d", SprintRef(ref), len(refs))
	}

	return refs[0], 1, nil
}

// IsAddressRefEmpty checks if an AddressRef is empty.
func IsAddressRefEmpty(ref datastore.AddressRef) bool {
	return ref.Address == "" &&
		ref.Type == "" &&
		ref.Version == nil &&
		ref.Qualifier == "" &&
		ref.ChainSelector == 0 &&
		ref.Labels.Length() == 0
}

// FullRef returns the entire datastore.AddressRef
func FullRef(ref datastore.AddressRef) (datastore.AddressRef, error) {
	return ref, nil
}

func GetAddressRef(
	input []datastore.AddressRef,
	selector uint64,
	contractType cldf.ContractType,
	contractVersion *semver.Version,
	contractQualifier string) datastore.AddressRef {
	for _, ref := range input {
		if ref.ChainSelector == selector &&
			ref.Type == datastore.ContractType(contractType) &&
			ref.Version.Equal(contractVersion) {
			if contractQualifier != "" {
				if ref.Qualifier == contractQualifier {
					return ref
				}
			} else {
				return ref
			}
		}
	}
	return datastore.AddressRef{}
}
