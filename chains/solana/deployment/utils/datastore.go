package utils

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

// ToByteArray formats a datastore.AddressRef into a byte slice.
func ToByteArray(ref datastore.AddressRef) (bytes []byte, err error) {
	if ref.Address == "" {
		return nil, fmt.Errorf("address is empty in ref: %s", datastore_utils.SprintRef(ref))
	}
	addr, err := ToAddress(ref)
	if err != nil {
		return nil, err
	}
	return addr.Bytes(), nil
}

// ToAddress formats a datastore.AddressRef into a solana.PublicKey.
func ToAddress(ref datastore.AddressRef) (commonAddress solana.PublicKey, err error) {
	if ref.Address == "" {
		return solana.PublicKey{}, fmt.Errorf("address is empty in ref: %s", datastore_utils.SprintRef(ref))
	}
	out, err := solana.PublicKeyFromBase58(ref.Address)
	if err != nil {
		return solana.PublicKey{}, fmt.Errorf("address is not a valid base58 address in ref: %s", datastore_utils.SprintRef(ref))
	}
	return out, nil
}
