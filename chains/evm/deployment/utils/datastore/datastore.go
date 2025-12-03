package datastore

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

// ToByteArray formats a datastore.AddressRef into a byte slice.
func ToByteArray(ref datastore.AddressRef) (bytes []byte, err error) {
	if ref.Address == "" {
		return nil, fmt.Errorf("address is empty in ref: %s", datastore_utils.SprintRef(ref))
	}
	if !common.IsHexAddress(ref.Address) {
		return nil, fmt.Errorf("address is not a valid hex address in ref: %s", datastore_utils.SprintRef(ref))
	}
	addr, err := ToEVMAddress(ref)
	if err != nil {
		return nil, err
	}
	return addr.Bytes(), nil
}

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
