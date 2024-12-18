package config

import (
	"github.com/gagliardetto/solana-go"
)

var (
	TimelockProgram         = solana.MustPublicKeyFromBase58("LoCoNsJFuhTkSQjfdDfn3yuwqhSYoPujmviRHVCzsqn")
	TimelockConfigPDA, _, _ = solana.FindProgramAddress([][]byte{[]byte("timelock_config")}, TimelockProgram)
	TimelockSignerPDA, _, _ = solana.FindProgramAddress([][]byte{[]byte("timelock_signer")}, TimelockProgram)
	TimelockOperationPDA    = func(id [32]byte) solana.PublicKey {
		pda, _, _ := solana.FindProgramAddress([][]byte{
			[]byte("timelock_operation"),
			id[:],
		}, TimelockProgram)
		return pda
	}
	NumAccountsPerRole      = 63 // max 64 accounts per role(access list) * 4 - 1(to keep test accounts fits single funding)
	BatchAddAccessChunkSize = 24

	MinDelay = uint64(1)

	TimelockEmptyOpID       = [32]byte{}
	TimelockOpDoneTimestamp = uint64(1)
)
