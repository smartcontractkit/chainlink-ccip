package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	tarops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

////////////////////
// Helper methods //
////////////////////

// AddressRefToBytes converts a hex-encoded address to its EVM byte representation.
func (a *EVMAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	if !common.IsHexAddress(ref.Address) {
		return nil, fmt.Errorf("address %q is not a valid hex address", ref.Address)
	}
	return common.HexToAddress(ref.Address).Bytes(), nil
}

func (a *EVMAdapter) GetTokenAdminRegistryAddress(ds datastore.DataStore, selector uint64) (common.Address, error) {
	filters := datastore.AddressRef{
		Type:          datastore.ContractType(tarops.ContractType),
		ChainSelector: selector,
		Version:       tarops.Version,
	}

	ref, err := datastore_utils.FindAndFormatRef(ds, filters, selector, datastore_utils.FullRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to find token admin registry address on chain %d: %w", selector, err)
	}

	addr, err := a.AddressRefToBytes(ref)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	return common.BytesToAddress(addr), nil
}

func (a *EVMAdapter) FindOneTokenAddress(ds datastore.DataStore, chainSelector uint64, partialRef *datastore.AddressRef) (common.Address, error) {
	filters := datastore.AddressRef{
		ChainSelector: chainSelector,
	}
	if partialRef != nil {
		filters.Address = partialRef.Address
		filters.Qualifier = partialRef.Qualifier
		filters.Type = partialRef.Type
	}

	ref, err := datastore_utils.FindAndFormatRef(ds, filters, chainSelector, datastore_utils.FullRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to find token address for ref %v on chain %d: %w", filters, chainSelector, err)
	}

	addr, err := a.AddressRefToBytes(ref)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	return common.BytesToAddress(addr), nil
}

func (a *EVMAdapter) FindLatestAddressRef(ds datastore.DataStore, ref datastore.AddressRef) (common.Address, error) {
	minVersion := semver.MustParse("1.5.0")
	maxVersion := semver.MustParse("2.0.0")

	filter := []datastore.FilterFunc[datastore.AddressRefKey, datastore.AddressRef]{}
	if ref.ChainSelector != 0 {
		filter = append(filter, datastore.AddressRefByChainSelector(ref.ChainSelector))
	}
	if ref.Qualifier != "" {
		filter = append(filter, datastore.AddressRefByQualifier(ref.Qualifier))
	}
	if ref.Version != nil {
		return common.Address{}, fmt.Errorf("ref version should not be set when finding the latest address ref, got version %s", ref.Version.String())
	}
	if ref.Address != "" {
		return common.Address{}, fmt.Errorf("ref address should not be set when finding the latest address ref, got address %q", ref.Address)
	}
	if ref.Type.String() != "" {
		filter = append(filter, datastore.AddressRefByType(ref.Type))
	}

	refs := ds.Addresses().Filter(filter...)

	var latestRef datastore.AddressRef
	latestVer := minVersion
	doesExist := false
	for _, ref := range refs {
		v := ref.Version
		if v == nil {
			continue
		}

		isInside := v.GreaterThanEqual(minVersion) && v.LessThan(maxVersion)
		isBetter := !doesExist || v.GreaterThan(latestVer)
		if isInside && isBetter {
			doesExist = true
			latestRef = ref
			latestVer = v
		}
	}

	if !doesExist {
		return common.Address{}, fmt.Errorf("no address found for ref (%+v) in version range [%s, %s)", ref, minVersion.String(), maxVersion.String())
	}

	addrBytes, err := a.AddressRefToBytes(latestRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to convert address ref to bytes: %w", err)
	}

	return common.BytesToAddress(addrBytes), nil
}

