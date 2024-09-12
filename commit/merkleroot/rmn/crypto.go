package rmn

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/protobuf/proto"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/rmnpb"
)

func abiEncode(abiStr string, values ...interface{}) ([]byte, error) {
	inDef := fmt.Sprintf(`[{ "name" : "method", "type": "function", "inputs": %s}]`, abiStr)

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

	internalMerkleRoots := make([]InternalMerkleRoot, len(reportData.LaneUpdates))
	for i, lu := range reportData.LaneUpdates {
		internalMerkleRoots[i] = InternalMerkleRoot{
			SourceChainSelector: lu.SourceChainSelector,
			OnRampAddress:       lu.OnRampAddress,
			MinSeqNr:            lu.MinSeqNr,
			MaxSeqNr:            lu.MaxSeqNr,
			MerkleRoot:          lu.MerkleRoot,
		}
	}

	abiDefinition := `[{"name": "", "type": "bytes32"}]`
	rmnVersionHash := crypto.Keccak256Hash([]byte("RMN_V1_6_ANY2EVM_REPORT"))
	rmnVersionHashAbi, err := abiEncode(abiDefinition, rmnVersionHash)
	if err != nil {
		log.Fatalf("Failed to ABI encode: %v", err)
	}

	encodingUtilsABI, err := abi.JSON(strings.NewReader(EncodingUtilsABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	p := &RMNRemoteReport{
		DestChainId:                 reportData.DestChainEvmID,
		DestChainSelector:           reportData.DestChainSelector,
		RmnRemoteContractAddress:    common.HexToAddress(reportData.RmnRemoteContractAddress),
		OfframpAddress:              common.HexToAddress(reportData.OfframpAddress),
		RmnHomeContractConfigDigest: reportData.RmnHomeContractConfigDigest,
		DestLaneUpdates:             internalMerkleRoots,
	}
	data, err := encodingUtilsABI.Methods["_rmnReport"].Inputs.Pack(rmnVersionHash, p)
	if err != nil {
		log.Fatalf("Failed to ABI encode: %v", err)
	}

	packedDataJSON, err := json.MarshalIndent(p, " ", " ")
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}
	fmt.Println(">>> abi encoded data:", string(packedDataJSON))

	// override with correct abi encoded data
	// data = common.FromHex("9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf53000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000038ac16c53d51914c00000000000000000000000000000000000000000000000049153604bdc88fa50000000000000000000000003d015cec4411357eff4ea5f009a581cc519f75d3000000000000000000000000c5cdb7711a478058023373b8ae9e7421925140f8785936570d1c7422ef30b7da5555ad2f175fa2dd97a2429a2e71d1e07c94e06000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000729d724980edda1000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000007d29db169b9649bc00000000000000000000000000000000000000000000000072581dd2208d09ba48e688aefc20a04fdec6b8ff19df358fd532455659dcf529797cda358e9e520500000000000000000000000000000000000000000000000000000000000000200000000000000000000000006662cb20464f4be557262693bea0409f068397ed")
	fmt.Println(">>>", hex.EncodeToString(rmnVersionHashAbi), hex.EncodeToString(data))

	// got
	// 9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf53000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000038ac16c53d51914c00000000000000000000000000000000000000000000000049153604bdc88fa50000000000000000000000003d015cec4411357eff4ea5f009a581cc519f75d3000000000000000000000000c5cdb7711a478058023373b8ae9e7421925140f8785936570d1c7422ef30b7da5555ad2f175fa2dd97a2429a2e71d1e07c94e06000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000729d724980edda1000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000007d29db169b9649bc00000000000000000000000000000000000000000000000072581dd2208d09ba48e688aefc20a04fdec6b8ff19df358fd532455659dcf529797cda358e9e520500000000000000000000000000000000000000000000000000000000000000200000000000000000000000006662cb20464f4be557262693bea0409f068397ed
	// exp
	// 9651943783dbf81935a60e98f218a9d9b5b28823fb2228bbd91320d632facf53000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000038ac16c53d51914c00000000000000000000000000000000000000000000000049153604bdc88fa50000000000000000000000003d015cec4411357eff4ea5f009a581cc519f75d3000000000000000000000000c5cdb7711a478058023373b8ae9e7421925140f8785936570d1c7422ef30b7da5555ad2f175fa2dd97a2429a2e71d1e07c94e06000000000000000000000000000000000000000000000000000000000000000c000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000729d724980edda1000000000000000000000000000000000000000000000000000000000000000a00000000000000000000000000000000000000000000000007d29db169b9649bc00000000000000000000000000000000000000000000000072581dd2208d09ba48e688aefc20a04fdec6b8ff19df358fd532455659dcf529797cda358e9e520500000000000000000000000000000000000000000000000000000000000000200000000000000000000000006662cb20464f4be557262693bea0409f068397ed

	signedHash := crypto.Keccak256Hash(data)
	// -----------------------------------------------------------------------------------

	for _, sig := range reportSigs {
		recoveredAddress, err := recoverAddressFromSig(
			v,
			sig.R,
			sig.S,
			signedHash[:],
		)
		if err != nil {
			return fmt.Errorf("failed to recover public key from signature: %w", err)
		}

		// Check if the public key is in the list of the provided RMN nodes
		found := false
		for _, node := range rmnNodes {
			if node.SignReportsAddress == recoveredAddress {
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

// Recover public address from ECDSA signature using r, s, v, and the hash of the message
func recoverAddressFromSig(v int, r, s []byte, hash []byte) (common.Address, error) {
	// Ensure r and s are 32 bytes each
	if len(r) != 32 || len(s) != 32 {
		return common.Address{}, errors.New("r and s must be 32 bytes")
	}

	// Ensure v is either 27 or 28 (as used in Ethereum)
	if v != 27 && v != 28 {
		return common.Address{}, errors.New("v must be 27 or 28")
	}

	// Construct the signature by concatenating r, s, and the recovery ID (v - 27 to convert to 0/1)
	sig := append(r, s...)
	sig = append(sig, byte(v-27))

	// Recover the public key bytes from the signature and message hash
	pubKeyBytes, err := crypto.Ecrecover(hash, sig)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to recover public key: %v", err)
	}

	// Convert the recovered public key to an ECDSA public key
	pubKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to unmarshal public key: %v", err)
	} // or SigToPub

	return crypto.PubkeyToAddress(*pubKey), nil
}
