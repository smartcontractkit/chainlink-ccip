package datastore

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// ToEVMAddress formats a datastore.AddressRef into an ethereum common.Address.
func ToEVMAddress(ref datastore.AddressRef) (commonAddress common.Address, err error) {
	if ref.Address == "" {
		return common.Address{}, fmt.Errorf("address is empty in ref: %s", datastore_utils.SprintRef(ref))
	}
	if !common.IsHexAddress(ref.Address) {
		return common.Address{}, fmt.Errorf("address is not a valid hex address in ref: %s", datastore_utils.SprintRef(ref))
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

func AddressRefOnChain(e cldf.Environment, selector uint64, ref datastore.AddressRef) (common.Address, error) {
	addressRef, err := datastore_utils.FindAndFormatRef(e.DataStore, ref, selector, datastore_utils.FullRef)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to resolve address ref on chain with selector %d: %w", selector, err)
	}
	if addressRef.Address == "" {
		return common.Address{}, fmt.Errorf("address ref resolved to empty address on chain %d for ref %v", selector, addressRef)
	}
	return common.HexToAddress(addressRef.Address), nil
}
