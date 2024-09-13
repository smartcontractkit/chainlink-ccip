package rmn

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"

	"google.golang.org/protobuf/proto"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
)

// verifyObservationSignature verifies the signature of the observation.
//
//	e.g. ed25519.sign(sha256("chainlink ccip 1.6 rmn observation"|sha256(observation)))
func verifyObservationSignature(
	rmnNode RMNNodeInfo,
	signedObs *rmnpb.SignedObservation,
) error {
	observationBytes, err := proto.Marshal(signedObs.GetObservation())
	if err != nil {
		return fmt.Errorf("failed to marshal observation: %w", err)
	}

	observationBytesSha256 := sha256.Sum256(observationBytes)
	msg := append([]byte(rmnNode.SignObservationPrefix), observationBytesSha256[:]...)
	msgSha256 := sha256.Sum256(msg)

	isValid := ed25519.Verify(*rmnNode.SignObservationsPublicKey, msgSha256[:], signedObs.Signature)
	if !isValid {
		return fmt.Errorf("observation signature does not match node %d public key", rmnNode.ID)
	}

	return nil
}
