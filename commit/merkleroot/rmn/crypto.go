package rmn

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
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

// VerifyRmnReportSignatures verifies if the provided signatures match all the provided rmn node public keys.
//
//	for each report signature:
//		recover the public key based on the laneUpdates
//		make sure the public key is in the list of RMN nodes
func VerifyRmnReportSignatures(
	laneUpdates []*rmnpb.FixedDestLaneUpdate,
	reportSigs []*rmnpb.EcdsaSignature,
	rmnNodes []RMNNodeInfo,
) error {
	const v = 27

	if reportSigs == nil {
		return fmt.Errorf("no signatures provided")
	}
	if laneUpdates == nil {
		return fmt.Errorf("no lane updates provided")
	}

	// todo: should match rmn signed msg but that's abi encoded
	// so we need to add a chain agnostic interface for computing this hash
	msg, err := json.Marshal(laneUpdates)
	if err != nil {
		return fmt.Errorf("failed to marshal lane updates: %w", err)
	}
	h := sha256.Sum256(msg)

	for _, sig := range reportSigs {
		recoveredPubKey, err := recoverPublicKeyFromSignature(
			v,
			new(big.Int).SetBytes(sig.R),
			new(big.Int).SetBytes(sig.S),
			h[:],
		)
		if err != nil {
			return fmt.Errorf("failed to recover public key from signature: %w", err)
		}

		// Check if the public key is in the list of the provided RMN nodes
		found := false
		for _, node := range rmnNodes {
			if node.SignReportsPublicKey.Equal(recoveredPubKey) {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("public key not found in RMN nodes")
		}
	}

	return nil
}

// Recover public key from ECDSA signature using r, s, v, and the hash of the message
func recoverPublicKeyFromSignature(v int, r, s *big.Int, hash []byte) (*ecdsa.PublicKey, error) {
	recoveryID := v - 27
	if recoveryID != 0 && recoveryID != 1 {
		return nil, fmt.Errorf("invalid v value: %d", v)
	}

	// Combine r and s into a 65-byte signature for recovery
	signature := make([]byte, 65)
	copy(signature[0:32], r.Bytes())  // r (32 bytes)
	copy(signature[32:64], s.Bytes()) // s (32 bytes)
	signature[64] = byte(recoveryID)  // v (recovery id)

	sigPublicKey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return nil, fmt.Errorf("failed to recover public key from signature: %w", err)
	}

	return sigPublicKey, nil
}
