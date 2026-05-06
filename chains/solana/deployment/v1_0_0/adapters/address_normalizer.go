package adapters

import (
	"fmt"

	"github.com/gagliardetto/solana-go"

	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
)

var _ deployapi.AddressNormalizer = (*SolanaAddressNormalizer)(nil)

// SolanaAddressNormalizer provides canonical base58 for datastore-aligned SVM lookups.
type SolanaAddressNormalizer struct{}

func (SolanaAddressNormalizer) NormalizeAddress(addr string) (string, error) {
	pubKey, err := solana.PublicKeyFromBase58(addr)
	if err != nil {
		return "", fmt.Errorf("failed to parse address '%s' as base58 Solana public key: %w", addr, err)
	}
	return pubKey.String(), nil
}
