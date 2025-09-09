package ccip

import (
	"fmt"
	"math/rand"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/latest/ccip_offramp"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/eth"
)

type Signatures struct {
	Rs    [][32]byte
	Ss    [][32]byte
	RawVs [32]byte
}

func SignCommitReport(ctx [2][32]byte, report ccip_offramp.CommitInput, baseSigners []eth.Signer) (sigs Signatures, err error) {
	hash, err := HashCommitReport(ctx, report)
	if err != nil {
		return Signatures{}, err
	}

	// make copy to avoid race flakiness when randomizing with parallel tests
	signers := make([]eth.Signer, len(baseSigners))
	copy(signers, baseSigners)

	// randomize signers
	rand.Shuffle(len(signers), func(i, j int) {
		signers[i], signers[j] = signers[j], signers[i]
	})

	for i := uint8(0); i < config.OcrF+1; i++ {
		baseSig := ecdsa.SignCompact(secp256k1.PrivKeyFromBytes(signers[i].PrivateKey), hash, false)
		sigs.RawVs[i] = baseSig[0] - 27 // key signs 27 or 28, but verification expects 0 or 1 (remove offset)
		sigs.Rs = append(sigs.Rs, [32]byte(baseSig[1:33]))
		sigs.Ss = append(sigs.Ss, [32]byte(baseSig[33:65]))
	}
	return sigs, nil
}

func AttestCCTP(message []byte, attesters []eth.Signer) ([]byte, error) {
	for i := 0; i < len(attesters)-1; i++ {
		if ethAddressSorter(attesters[i].Address, attesters[i+1].Address) >= 0 {
			return nil, fmt.Errorf("attesters are not sorted: %x >= %x", attesters[i].Address, attesters[i+1].Address)
		}
	}
	messageHash := eth.Keccak256(message)

	// Each signature is 65 bytes (32 bytes for R, 32 bytes for S, 1 byte for V)
	const signatureLength = 65
	attestation := make([]byte, len(attesters)*signatureLength)
	writeOffset := 0

	for _, attester := range attesters {
		if len(attester.PrivateKey) == 0 {
			return nil, fmt.Errorf("attester private key is empty for address %x", attester.Address)
		}
		privKey := secp256k1.PrivKeyFromBytes(attester.PrivateKey)

		// SignCompact returns a 65-byte compact signature: [V (1 byte) || R (32 bytes) || S (32 bytes)]
		// V is recoveryID + 27.
		compactSig := ecdsa.SignCompact(privKey, messageHash, false) // isCompressedKey = false

		// We need to store it as R || S || V in the attestation buffer (the V byte moved to the end)
		v := compactSig[0]
		rs := compactSig[1:]

		copy(attestation[writeOffset:], rs)
		writeOffset += signatureLength - 1 // 64 bytes for R and S
		attestation[writeOffset] = v
		writeOffset++
	}
	return attestation, nil
}

func ethAddressSorter(a, b [20]byte) int {
	for i := 0; i < 20; i++ {
		if a[i] < b[i] {
			return -1
		} else if a[i] > b[i] {
			return 1
		}
	}
	return 0
}
