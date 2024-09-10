package rmn

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SignaturesPlayground(t *testing.T) {
	rmnNode := newRmn()

	t.Run("test observation signature verification", func(t *testing.T) {
		observationSignMessage := []byte("some protobuf encoded observation data")
		signature := rmnNode.signObservationMessage(observationSignMessage)

		pubKeyBytes := *rmnNode.o_PK
		pubKeyHex := hex.EncodeToString(pubKeyBytes)
		t.Logf("verifying signature with rmn node public key: %s", pubKeyHex)

		valid := ed25519.Verify(pubKeyBytes, observationSignMessage, signature)
		assert.True(t, valid)

		signature[0] = signature[0] + 1
		valid = ed25519.Verify(pubKeyBytes, observationSignMessage, signature)
		assert.False(t, valid)
	})

	t.Run("test report signature verification", func(t *testing.T) {
		reportSignMessage := []byte("some protobuf encoded report data")
		r, s := rmnNode.signReportMessage(reportSignMessage)
		h := sha256.Sum256(reportSignMessage)

		pubKeyBytes := elliptic.Marshal(elliptic.P256(), rmnNode.r_PK.X, rmnNode.r_PK.Y)
		pubKeyHex := hex.EncodeToString(pubKeyBytes)
		t.Logf("verifying signature with rmn node public key: %s", pubKeyHex)

		valid := ecdsa.Verify(
			rmnNode.r_PK,
			h[:],
			big.NewInt(0).SetBytes(r),
			big.NewInt(0).SetBytes(s),
		)
		assert.True(t, valid)

		r[0] = r[0] + 1
		valid = ecdsa.Verify(
			rmnNode.r_PK,
			h[:],
			big.NewInt(0).SetBytes(r),
			big.NewInt(0).SetBytes(s),
		)
		assert.False(t, valid)
	})

}

type rmn struct {
	// observations signing
	o_sk *ed25519.PrivateKey
	o_PK *ed25519.PublicKey

	// reports signing
	r_sk *ecdsa.PrivateKey
	r_PK *ecdsa.PublicKey
}

func newRmn() *rmn {
	// for observation verification
	o_pk, o_sk, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate Ed25519 key pair: %v", err)
	}

	// for report signature verification
	r_sk, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	return &rmn{
		o_sk: &o_sk,
		o_PK: &o_pk,
		r_sk: r_sk,
		r_PK: &r_sk.PublicKey,
	}
}

func (n *rmn) signObservationMessage(message []byte) []byte {
	return ed25519.Sign(*n.o_sk, message)
}

func (n *rmn) signReportMessage(message []byte) (r, s []byte) {
	h := sha256.Sum256(message)

	rInt, sInt, err := ecdsa.Sign(rand.Reader, n.r_sk, h[:])
	if err != nil {
		panic(err)
	}

	return rInt.Bytes(), sInt.Bytes()
}
