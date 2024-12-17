package contracts

import (
	"math/rand"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/decred/dcrd/dcrec/secp256k1/v4/ecdsa"
	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/config"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils/eth"
)

func SignCommitReport(ctx [3][32]byte, report ccip_router.CommitInput, baseSigners []eth.Signer) (sigs [][65]byte, err error) {
	hash, err := HashCommitReport(ctx, report)
	if err != nil {
		return nil, err
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
		baseSig[0] = baseSig[0] - 27 // key signs 27 or 28, but verification expects 0 or 1 (remove offset)
		sigs = append(sigs, [65]byte(baseSig))
	}
	return sigs, nil
}

func getChainStatePDA(chainSelector uint64) (solana.PublicKey, error) {
	chainSelectorLE := utils.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("chain_state"), chainSelectorLE}, config.CcipRouterProgram)
	return p, err
}

func GetCommitReportPDA(chainSelector uint64, root [32]byte) (solana.PublicKey, error) {
	chainSelectorLE := utils.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("commit_report"), chainSelectorLE, root[:]}, config.CcipRouterProgram)
	return p, err
}

func getNoncePDA(chainSelector uint64, user solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := utils.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("nonce"), chainSelectorLE, user.Bytes()}, config.CcipRouterProgram)
	return p, err
}
