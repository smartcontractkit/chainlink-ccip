package datastore

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

// FormatFn is a function that formats a datastore.AddressRef into a specific type T.
type FormatFn[T any] = func(ref datastore.AddressRef) (T, error)

// ToEVMAddress formats a datastore.AddressRef into an ethereum common.Address.
func ToEVMAddress(ref datastore.AddressRef) (commonAddress common.Address, err error) {
	if ref.Address == "" {
		return common.Address{}, fmt.Errorf("address is empty in ref: %s", sprintRef(ref))
	}
	if !common.IsHexAddress(ref.Address) {
		return common.Address{}, fmt.Errorf("address is not a valid hex address in ref: %s", sprintRef(ref))
	}
	return common.HexToAddress(ref.Address), nil
}

// ToPaddedEVMAddress formats a datastore.AddressRef into a 32-byte padded ethereum address.
func ToPaddedEVMAddress(ref datastore.AddressRef) (paddedAddress []byte, err error) {
	addr, err := ToEVMAddress(ref)
	if err != nil {
		return nil, err
	}
	return common.LeftPadBytes(addr.Bytes(), 32), nil
}

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
			return nil, fmt.Errorf("failed to format ref %s: %w", sprintRef(refFromStore), err)
		}
		formattedRefs = append(formattedRefs, formattedRef)
	}
	return formattedRefs, nil
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
		return datastore.AddressRef{}, fmt.Errorf("expected to find exactly 1 ref with criteria %s, found %d", sprintRef(ref), len(refs))
	}
	return refs[0], nil
}

func sprintRef(ref datastore.AddressRef) string {
	return fmt.Sprintf("{ChainSelector: %d, Type: %s, Version: %s, Qualifier: %s, Address: %s}", ref.ChainSelector, ref.Type, ref.Version, ref.Qualifier, ref.Address)
}
