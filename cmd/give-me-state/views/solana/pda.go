package solana

import (
	"encoding/binary"

	"github.com/gagliardetto/solana-go"
)

// Re-export solana.PublicKey for convenience
type PublicKey = solana.PublicKey

// PublicKeyFromBase58 decodes a base58 string to a PublicKey
func PublicKeyFromBase58(s string) (PublicKey, error) {
	return solana.PublicKeyFromBase58(s)
}

// Uint64ToLE converts uint64 to little-endian bytes
func Uint64ToLE(n uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, n)
	return b
}

// FindProgramAddress wraps the solana-go SDK's FindProgramAddress
func FindProgramAddress(seeds [][]byte, programID PublicKey) (PublicKey, uint8, error) {
	return solana.FindProgramAddress(seeds, programID)
}

// =====================
// CCIP Router PDAs
// =====================

func FindRouterConfigPDA(routerProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("config")}, routerProgram)
}

func FindDestChainStatePDA(chainSelector uint64, routerProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("dest_chain_state"), Uint64ToLE(chainSelector)}, routerProgram)
}

func FindTokenAdminRegistryPDA(mint PublicKey, routerProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("token_admin_registry"), mint.Bytes()}, routerProgram)
}

// =====================
// Fee Quoter PDAs
// =====================

func FindFeeQuoterConfigPDA(feeQuoterProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("config")}, feeQuoterProgram)
}

func FindFeeQuoterDestChainPDA(chainSelector uint64, feeQuoterProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("dest_chain"), Uint64ToLE(chainSelector)}, feeQuoterProgram)
}

func FindFeeQuoterBillingTokenConfigPDA(mint PublicKey, feeQuoterProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("fee_billing_token_config"), mint.Bytes()}, feeQuoterProgram)
}

// =====================
// OffRamp PDAs
// =====================

func FindOffRampConfigPDA(offrampProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("config")}, offrampProgram)
}

func FindOffRampReferenceAddressesPDA(offrampProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("reference_addresses")}, offrampProgram)
}

func FindOffRampSourceChainPDA(chainSelector uint64, offrampProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("source_chain_state"), Uint64ToLE(chainSelector)}, offrampProgram)
}

// =====================
// RMN Remote PDAs
// =====================

func FindRMNRemoteConfigPDA(rmnRemoteProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("config")}, rmnRemoteProgram)
}

func FindRMNRemoteCursesPDA(rmnRemoteProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("curses")}, rmnRemoteProgram)
}

// =====================
// Token Pool PDAs
// =====================

func FindTokenPoolConfigPDA(mint PublicKey, poolProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("ccip_tokenpool_config"), mint.Bytes()}, poolProgram)
}

func FindTokenPoolChainConfigPDA(chainSelector uint64, mint PublicKey, poolProgram PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("ccip_tokenpool_chainconfig"), Uint64ToLE(chainSelector), mint.Bytes()}, poolProgram)
}

// =====================
// MCMS/Timelock PDAs
// =====================

type PDASeed [32]byte

func FindMCMConfigPDA(mcmProgram PublicKey, msigID PDASeed) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("multisig_config"), msigID[:]}, mcmProgram)
}

func FindTimelockConfigPDA(timelockProgram PublicKey, timelockID PDASeed) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("timelock_config"), timelockID[:]}, timelockProgram)
}

// =====================
// Shared PDAs
// =====================

func FindNameAndVersionPDA(program PublicKey) (PublicKey, uint8, error) {
	return FindProgramAddress([][]byte{[]byte("name_version")}, program)
}
