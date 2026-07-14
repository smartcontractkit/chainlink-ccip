package adapters

import (
	"errors"
	"fmt"

	"github.com/gagliardetto/solana-go"

	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	mcms_solana "github.com/smartcontractkit/mcms/sdk/solana"
)

var _ deployapi.AddressNormalizer = (*SolanaAddressNormalizer)(nil)

// SolanaAddressNormalizer provides canonical base58 for datastore-aligned SVM lookups.
type SolanaAddressNormalizer struct{}

func (SolanaAddressNormalizer) NormalizeAddress(addr string) (string, error) {
	programID, seed, err := mcms_solana.ParseContractAddress(addr)
	if err == nil {
		return mcms_solana.ContractAddress(programID, seed), nil
	}

	if errors.Is(err, mcms_solana.ErrInvalidContractAddressFormat) {
		if pubKey, err := solana.PublicKeyFromBase58(addr); err != nil {
			return "", fmt.Errorf("failed to parse address '%s' as base58 Solana public key: %w", addr, err)
		} else {
			return pubKey.String(), nil
		}
	}

	return "", fmt.Errorf("failed to normalize address '%s': %w", addr, err)
}

// BytesToString converts raw on-chain address bytes (a 32-byte public key) to its canonical base58 string.
func (SolanaAddressNormalizer) BytesToString(address []byte) (string, error) {
	if len(address) != solana.PublicKeyLength {
		return "", fmt.Errorf("solana address must be %d bytes, got %d", solana.PublicKeyLength, len(address))
	}
	return solana.PublicKeyFromBytes(address).String(), nil
}

// StringToBytes converts a base58 Solana public key string to its 32-byte representation.
func (SolanaAddressNormalizer) StringToBytes(address string) ([]byte, error) {
	pubKey, err := solana.PublicKeyFromBase58(address)
	if err != nil {
		return nil, fmt.Errorf("failed to parse address '%s' as base58 Solana public key: %w", address, err)
	}
	return pubKey.Bytes(), nil
}
