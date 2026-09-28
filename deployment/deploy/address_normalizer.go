package deploy

import (
	"bytes"
	"fmt"

	chainsel "github.com/smartcontractkit/chain-selectors"

	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

// TryNormalizeAddressRef normalizes the given AddressRef.Address based on the provided chain selector.
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

// BytesToString converts raw on-chain address bytes to a string using the chain family's
// registered AddressNormalizer for the given chain selector.
func BytesToString(sel uint64, addr []byte) (string, error) {
	family, err := chainsel.GetSelectorFamily(sel)
	if err != nil {
		return "", fmt.Errorf("invalid chain selector %d: %w", sel, err)
	}

	normalizer, ok := GetAddressNormalizerRegistry().GetAddressNormalizer(family)
	if !ok {
		return "", fmt.Errorf("no address normalizer registered for chain family %q of chain selector %d", family, sel)
	}

	return normalizer.BytesToString(addr)
}

// StringToBytes converts an address string to raw on-chain bytes using the chain family's
// registered AddressNormalizer for the given chain selector.
func StringToBytes(sel uint64, addr string) ([]byte, error) {
	family, err := chainsel.GetSelectorFamily(sel)
	if err != nil {
		return nil, fmt.Errorf("invalid chain selector %d: %w", sel, err)
	}

	normalizer, ok := GetAddressNormalizerRegistry().GetAddressNormalizer(family)
	if !ok {
		return nil, fmt.Errorf("no address normalizer registered for chain family %q of chain selector %d", family, sel)
	}

	return normalizer.StringToBytes(addr)
}

// RoundTripAddress converts an address string to on-chain bytes and back using the chain family's
// registered AddressNormalizer (StringToBytes then BytesToString), so addresses from different
// sources (user input, datastore, on-chain reads) compare equal as strings.
func RoundTripAddress(sel uint64, addr string) (string, error) {
	if b, err := StringToBytes(sel, addr); err != nil {
		return "", fmt.Errorf("failed to convert address string to bytes: %w", err)
	} else {
		return BytesToString(sel, b)
	}
}

// IsZeroAddress reports whether addr decodes to all-zero bytes in the address encoding of the
// chain family of sel. The check is on bytes because the zero address's string form is
// family-specific (e.g. EVM 0x000…0, Solana 111…1 in base58) while its bytes are all zero.
func IsZeroAddress(sel uint64, addr string) (bool, error) {
	if b, err := StringToBytes(sel, addr); err != nil {
		return false, fmt.Errorf("failed to convert address string to bytes: %w", err)
	} else {
		return len(bytes.Trim(b, "\x00")) == 0, nil
	}
}
