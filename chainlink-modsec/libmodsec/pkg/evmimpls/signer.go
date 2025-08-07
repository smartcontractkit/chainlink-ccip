package evmimpls

import (
	"context"
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

var _ modsectypes.Signer = (*EVMSigner)(nil)

type EVMSigner struct {
	privateKey *ecdsa.PrivateKey
}

func (s *EVMSigner) Sign(ctx context.Context, digest []byte) ([]byte, error) {
	if len(digest) != 32 {
		return nil, fmt.Errorf("digest must be 32 bytes long")
	}

	return crypto.Sign(digest, s.privateKey)
}
