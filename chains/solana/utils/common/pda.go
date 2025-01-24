package common

import (
	"github.com/gagliardetto/solana-go"
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

func FindTokenAdminRegistryPDA(mint, ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("token_admin_registry"), mint.Bytes()}, ccipRouterProgram)
}

func FindFeeBillingTokenConfigPDA(mint, ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("fee_billing_token_config"), mint.Bytes()}, ccipRouterProgram)
}

func FindCcipTokenpoolBillingPDA(chainSelector uint64, mint, ccipRouterProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	chainSelectorLE := Uint64ToLE(chainSelector)
	return solana.FindProgramAddress([][]byte{[]byte("ccip_tokenpool_billing"), chainSelectorLE, mint.Bytes()}, ccipRouterProgram)
}

func GetSourceChainStatePDA(chainSelector uint64, ccipRouterProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("source_chain_state"), chainSelectorLE}, ccipRouterProgram)
	return p, err
}

func GetDestChainStatePDA(chainSelector uint64, ccipRouterProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("dest_chain_state"), chainSelectorLE}, ccipRouterProgram)
	return p, err
}

func GetCommitReportPDA(chainSelector uint64, root [32]byte, ccipRouterProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("commit_report"), chainSelectorLE, root[:]}, ccipRouterProgram)
	return p, err
}

func GetNoncePDA(chainSelector uint64, user, ccipRouterProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("nonce"), chainSelectorLE, user.Bytes()}, ccipRouterProgram)
	return p, err
}

// Test Code

func FindCounterPDA(ccipReceiverProgram solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("counter")}, ccipReceiverProgram)
}
