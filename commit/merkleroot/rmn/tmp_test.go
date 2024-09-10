package rmn

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayground(t *testing.T) {
	rmnNode := newRmn()

	message := []byte("some protobuf encoded data")
	r, s := rmnNode.sign(message)
	h := sha256.Sum256(message)

	pubKeyBytes := elliptic.Marshal(elliptic.P256(), rmnNode.PK.X, rmnNode.PK.Y)
	pubKeyHex := hex.EncodeToString(pubKeyBytes)
	t.Logf("verifying signature with rmn node public key: %s", pubKeyHex)

	valid := ecdsa.Verify(
		rmnNode.PK,
		h[:],
		big.NewInt(0).SetBytes(r),
		big.NewInt(0).SetBytes(s),
	)
	assert.True(t, valid)

	r[0] = r[0] + 1 // change r
	valid = ecdsa.Verify(
		rmnNode.PK,
		h[:],
		big.NewInt(0).SetBytes(r),
		big.NewInt(0).SetBytes(s),
	)
	assert.False(t, valid)
}

type rmn struct {
	sk *ecdsa.PrivateKey
	PK *ecdsa.PublicKey
}

func newRmn() *rmn {
	sk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	return &rmn{
		sk: sk,
		PK: &sk.PublicKey,
	}
}

func (n *rmn) sign(message []byte) (r, s []byte) {
	h := sha256.Sum256(message)

	rInt, sInt, err := ecdsa.Sign(rand.Reader, n.sk, h[:])
	if err != nil {
		panic(err)
	}

	return rInt.Bytes(), sInt.Bytes()
}
