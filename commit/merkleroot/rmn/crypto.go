package rmn

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/protobuf/proto"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
)

func abiEncode(abiStr string, values ...interface{}) ([]byte, error) {
	inDef := fmt.Sprintf(`[{ "name" : "method", "type": "function", "%s": %s}]`, "inputs", abiStr)

	ab, err := abi.JSON(strings.NewReader(inDef))
	if err != nil {
		return nil, err
	}

	res, err := ab.Pack("method", values...)
	if err != nil {
		return nil, err
	}

	return res[4:], nil
}

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
	reportData ReportData,
	reportSigs []*rmnpb.EcdsaSignature,
	rmnNodes []RMNNodeInfo,
) error {
	const v = 27

	if reportSigs == nil {
		return fmt.Errorf("no signatures provided")
	}
	if reportData.LaneUpdates == nil {
		return fmt.Errorf("no lane updates provided")
	}

	// -----------------------------------------------------------------------------------

	abiDefinition := `[
		{"name": "", "type": "bytes32"},
		{
			"name": "",
			"type": "tuple",
			"components": [
				{"name": "DestChainEvmID", "type": "uint256"},
				{"name": "DestChainSelector", "type": "uint64"},
				{"name": "RmnRemoteContractAddress", "type": "address"},
				{"name": "OfframpAddress", "type": "address"},
				{"name": "RmnHomeContractConfigDigest", "type": "bytes32"},
				{
					"name": "LaneUpdates",
					"type": "tuple[]",
					"components": [
						{"name": "SourceChainSelector", "type": "uint64"},
						{"name": "MinSeqNr", "type": "uint64"},
						{"name": "MaxSeqNr", "type": "uint64"},
						{"name": "MerkleRoot", "type": "bytes32"},
						{"name": "OnRampAddress", "type": "bytes"}
					]
				}
			]
		}
	]`

	rmnVersionHash := crypto.Keccak256Hash([]byte("RMN_V1_6_ANY2EVM_REPORT"))
	fmt.Println(">>>>>>>>>> rmnVersionHash", rmnVersionHash)

	data, err := abiEncode(abiDefinition, rmnVersionHash, reportData)
	if err != nil {
		log.Fatalf("Failed to ABI encode: %v", err)
	}
	fmt.Println(">>>>>>>>>> abi encoded data", data)

	signedHash := crypto.Keccak256Hash(data)
	fmt.Println(">>>>>>>>>>>> signedHash", signedHash)
	// -----------------------------------------------------------------------------------

	for _, sig := range reportSigs {
		recoveredPubKey, err := recoverPublicKeyFromSignature(
			v,
			new(big.Int).SetBytes(sig.R),
			new(big.Int).SetBytes(sig.S),
			signedHash[:],
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
			return fmt.Errorf("public key not found in RMN nodes, signature verification failed")
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
