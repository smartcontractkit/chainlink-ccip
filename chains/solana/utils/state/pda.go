package state

import (
	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

func FindConfigPDA(ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, ccipRouterProgram)
}

func FindStatePDA(ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("state")}, ccipRouterProgram)
}

func FindExternalExecutionConfigPDA(ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("external_execution_config")}, ccipRouterProgram)
}

func FindExternalTokenPoolsSignerPDA(ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("external_token_pools_signer")}, ccipRouterProgram)
}

func FindFeeBillingSignerPDA(ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("fee_billing_signer")}, ccipRouterProgram)
}

func FindTokenAdminRegistryPDA(mint solana.PublicKey, ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("token_admin_registry"), mint.Bytes()}, ccipRouterProgram)
}

func FindSourceChainStatePDA(chainSelector uint64, ccipRouterProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("source_chain_state"), chainSelectorLE}, ccipRouterProgram)
	return p, err
}

func FindDestChainStatePDA(chainSelector uint64, ccipRouterProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("dest_chain_state"), chainSelectorLE}, ccipRouterProgram)
	return p, err
}

func FindCommitReportPDA(chainSelector uint64, root [32]byte, ccipRouterProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("commit_report"), chainSelectorLE, root[:]}, ccipRouterProgram)
	return p, err
}

func FindNoncePDA(chainSelector uint64, user solana.PublicKey, ccipRouterProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("nonce"), chainSelectorLE, user.Bytes()}, ccipRouterProgram)
	return p, err
}

func FindApprovedSender(chainSelector uint64, sourceSender []byte, receiverProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("approved_ccip_sender"), chainSelectorLE, []byte{uint8(len(sourceSender))}, sourceSender}, receiverProgram) //nolint:gosec
	return p, err
}

/////////////////////
// Fee Quoter PDAs //
/////////////////////

func FindFqConfigPDA(feeQuoterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, feeQuoterProgram)
}

func FindFqDestChainPDA(chainSelector uint64, feeQuoterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	return solana.FindProgramAddress([][]byte{[]byte("dest_chain"), chainSelectorLE}, feeQuoterProgram)
}

func FindFqBillingTokenConfigPDA(mint solana.PublicKey, feeQuoterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("fee_billing_token_config"), mint.Bytes()}, feeQuoterProgram)
}

func FindFqBillingSignerPDA(feeQuoterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("fee_billing_signer")}, feeQuoterProgram)
}

func FindFqPerChainPerTokenConfigPDA(chainSelector uint64, mint solana.PublicKey, feeQuoterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	return solana.FindProgramAddress([][]byte{[]byte("per_chain_per_token_config"), chainSelectorLE, mint.Bytes()}, feeQuoterProgram)
}
