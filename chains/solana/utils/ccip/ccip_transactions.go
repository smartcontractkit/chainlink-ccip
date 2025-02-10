package ccip

import (
	"math/rand"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_offramp"
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
