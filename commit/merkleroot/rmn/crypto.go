// crypto.go contains functions and types related to cryptographic operations
// within the RMN package, e.g. signature verification.

package rmn

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
)

// ED25519Verifier is an interface for verifying ED25519 signatures.
type ED25519Verifier interface {
	Verify(publicKey ed25519.PublicKey, message, sig []byte) bool
}

type ED25519VerifierImpl struct{}

func NewED25519Verifier() ED25519Verifier {
	return ED25519VerifierImpl{}
}

func (ED25519VerifierImpl) Verify(publicKey ed25519.PublicKey, message, sig []byte) bool {
	return ed25519.Verify(publicKey, message, sig)
}

// verifyObservationSignature verifies the signature of the RMN observation.
//
//	e.g. ed25519.sign(sha256("chainlink ccip 1.6 rmn observation"|sha256(observation)))
func verifyObservationSignature(
	rmnNode RMNNodeInfo,
	signedObs *rmnpb.SignedObservation,
	verifier ED25519Verifier,
) error {
	observationBytes, err := proto.Marshal(signedObs.GetObservation())
	if err != nil {
		return fmt.Errorf("failed to marshal observation: %w", err)
	}

	observationBytesSha256 := sha256.Sum256(observationBytes)
	msg := append([]byte(rmnNode.SignObservationPrefix), observationBytesSha256[:]...)
	msgSha256 := sha256.Sum256(msg)

	isValid := verifier.Verify(*rmnNode.SignObservationsPublicKey, msgSha256[:], signedObs.Signature)
	if !isValid {
		return fmt.Errorf("observation signature does not match node %d public key", rmnNode.ID)
	}

	return nil
}
