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

func FindFeeBillingTokenConfigPDA(mint solana.PublicKey, ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("fee_billing_token_config"), mint.Bytes()}, ccipRouterProgram)
}

func FindCcipTokenpoolBillingPDA(chainSelector uint64, mint solana.PublicKey, ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	chainSelectorLE := common.Uint64ToLE(chainSelector)
	return solana.FindProgramAddress([][]byte{[]byte("ccip_tokenpool_billing"), chainSelectorLE, mint.Bytes()}, ccipRouterProgram)
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
