package deploy

import (
	"fmt"

	chainsel "github.com/smartcontractkit/chain-selectors"

	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

// TryNormalizeAddressRef attempts to normalize the given AddressRef.Address based on the provided chain selector.
func TryNormalizeAddressRef(sel uint64, ref datastore.AddressRef) (datastore.AddressRef, error) {
	if datastore_utils.IsAddressRefEmpty(ref) {
		return ref, nil
	} else {
		// NOTE: `ref.ChainSelector` is intentionally ignored in favor of `sel`
		ref.ChainSelector = sel
	}

	normalized := ref.Clone()
	if normalized.Address == "" {
		return normalized, nil
	}

	family, err := chainsel.GetSelectorFamily(sel)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("invalid chain selector %d: %w", sel, err)
	}

	normalizer, ok := GetAddressNormalizerRegistry().GetAddressNormalizer(family)
	if !ok {
		return normalized, nil
	}

	normalized.Address, err = normalizer.NormalizeAddress(ref.Address)
	if err != nil {
		return datastore.AddressRef{}, fmt.Errorf("failed to normalize address %s: %w", ref.Address, err)
	}

	return normalized, nil
}
