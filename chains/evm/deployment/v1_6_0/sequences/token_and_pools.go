package sequences

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"

	evm1_0_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	evm_ds_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

////////////////////
// Helper methods //
////////////////////

// AddressRefToBytes converts a hex-encoded address to its EVM byte representation.
// Delegates to the shared evm datastore utility.
func (a *EVMAdapter) AddressRefToBytes(ref datastore.AddressRef) ([]byte, error) {
	return evm_ds_utils.ToEVMAddressBytes(ref)
}

// GetTokenAdminRegistryAddress looks up the TAR (v1.5.0) address from the datastore.
// Delegates to the shared v1.0.0 adapter helper.
func (a *EVMAdapter) GetTokenAdminRegistryAddress(ds datastore.DataStore, selector uint64) (common.Address, error) {
	return evm1_0_0.GetTokenAdminRegistryAddress(ds, selector, &evm1_0_0.EVMTokenBase{})
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
	latestRef, err := datastore_utils.FindLatestRef(ds, ref, semver.MustParse("1.5.0"), semver.MustParse("2.0.0"))
	if err != nil {
		return common.Address{}, err
	}
	return evm_ds_utils.ToEVMAddress(latestRef)
}

