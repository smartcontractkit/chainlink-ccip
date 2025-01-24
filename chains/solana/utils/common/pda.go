package common

import (
	"encoding/binary"

	"github.com/gagliardetto/solana-go"
)

func FindConfigPDA(programID solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("config")}, programID)
}

func FindStatePDA(programID solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("state")}, programID)
}

func FindExternalExecutionConfigPDA(programID solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("external_execution_config")}, programID)
}

func FindExternalTokenPoolsSignerPDA(programID solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("external_token_pools_signer")}, programID)
}

func FindCounterPDA(programID solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("counter")}, programID)
}

func FindFeeBillingSignerPDA(programID solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("fee_billing_signer")}, programID)
}

func FindTokenAdminRegistryPDA(programID, mint solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("token_admin_registry"), mint.Bytes()}, programID)
}

func FindFeeBillingTokenConfigPDA(programID, mint solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("fee_billing_token_config"), mint.Bytes()}, programID)
}

func FindCcipTokenpoolBillingPDA(programID, mint solana.PublicKey, chainSelector uint64) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("ccip_tokenpool_billing"), binary.LittleEndian.AppendUint64([]byte{}, chainSelector), mint.Bytes()}, programID)
}

func FindCcipTokenpoolChainconfigPDA(programID, mint solana.PublicKey, chainSelector uint64) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress([][]byte{[]byte("ccip_tokenpool_chainconfig"), binary.LittleEndian.AppendUint64([]byte{}, chainSelector), mint.Bytes()}, programID)
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

func GetNoncePDA(chainSelector uint64, user solana.PublicKey, ccipRouterProgram solana.PublicKey) (solana.PublicKey, error) {
	chainSelectorLE := Uint64ToLE(chainSelector)
	p, _, err := solana.FindProgramAddress([][]byte{[]byte("nonce"), chainSelectorLE, user.Bytes()}, ccipRouterProgram)
	return p, err
}
