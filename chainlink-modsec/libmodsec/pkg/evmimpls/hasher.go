package evmimpls

import (
	"context"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

var _ modsectypes.Hasher = (*EVMHasher)(nil)

type EVMHasher struct {
}

func (h *EVMHasher) Hash(ctx context.Context, data []byte) ([]byte, error) {
	return crypto.Keccak256(data), nil
}
