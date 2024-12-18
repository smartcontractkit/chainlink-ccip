package config

import (
	"encoding/binary"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/contracts/tests/utils"
)

var (
	PrintEvents       = true
	DefaultCommitment = rpc.CommitmentConfirmed

	CcipRouterProgram          = solana.MustPublicKeyFromBase58("C8WSPj3yyus1YN3yNB6YA5zStYtbjQWtpmKadmvyUXq8")
	CcipReceiverProgram        = solana.MustPublicKeyFromBase58("CtEVnHsQzhTNWav8skikiV2oF6Xx7r7uGGa8eCDQtTjH")
	CcipReceiverAddress        = solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb")
	CcipInvalidReceiverProgram = solana.MustPublicKeyFromBase58("9Vjda3WU2gsJgE4VdU6QuDw8rfHLyigfFyWs3XDPNUn8")
	CcipTokenPoolProgram       = solana.MustPublicKeyFromBase58("GRvFSLwR7szpjgNEZbGe4HtxfJYXqySXuuRUAJDpu4WH")
	Token2022Program           = solana.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")

	RouterConfigPDA, _, _                    = solana.FindProgramAddress([][]byte{[]byte("config")}, CcipRouterProgram)
	ExternalExecutionConfigPDA, _, _         = solana.FindProgramAddress([][]byte{[]byte("external_execution_config")}, CcipRouterProgram)
	ExternalTokenPoolsSignerPDA, _, _        = solana.FindProgramAddress([][]byte{[]byte("external_token_pools_signer")}, CcipRouterProgram)
	ReceiverTargetAccountPDA, _, _           = solana.FindProgramAddress([][]byte{[]byte("counter")}, CcipReceiverProgram)
	ReceiverExternalExecutionConfigPDA, _, _ = solana.FindProgramAddress([][]byte{[]byte("external_execution_config")}, CcipReceiverProgram)
	BillingSignerPDA, _, _                   = solana.FindProgramAddress([][]byte{[]byte("fee_billing_signer")}, CcipRouterProgram)

	BillingTokenConfigPrefix = []byte("fee_billing_token_config")
	DestChainConfigPrefix    = []byte("destination_billing_config")

	SolanaChainSelector uint64 = 15
	EvmChainSelector    uint64 = 21

	SolanaChainStatePDA, _, _ = solana.FindProgramAddress([][]byte{[]byte("chain_state"), binary.LittleEndian.AppendUint64([]byte{}, SolanaChainSelector)}, CcipRouterProgram)
	EvmChainLE                = utils.Uint64ToLE(EvmChainSelector)
	EvmChainStatePDA, _, _    = solana.FindProgramAddress([][]byte{[]byte("chain_state"), binary.LittleEndian.AppendUint64([]byte{}, EvmChainSelector)}, CcipRouterProgram)

	OnRampAddress        = []byte{1, 2, 3}
	EnableExecutionAfter = int64(1800) // 30min

	MaxOracles           = 16
	OcrF           uint8 = 5
	ConfigDigest         = utils.MakeRandom32ByteArray()
	Empty24Byte          = [24]byte{}
	ReportSequence       = uint64(8)
	ReportContext        = [3][32]byte{
		ConfigDigest,
		[32]byte(binary.BigEndian.AppendUint64(Empty24Byte[:], ReportSequence)),
		utils.MakeRandom32ByteArray(),
	}
	MaxSignersAndTransmitters = 16
)
