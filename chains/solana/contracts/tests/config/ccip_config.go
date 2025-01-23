package config

import (
	"encoding/binary"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
)

var (
	PrintEvents       = true
	DefaultCommitment = rpc.CommitmentConfirmed

	CcipRouterProgram          = solana.MustPublicKeyFromBase58("C8WSPj3yyus1YN3yNB6YA5zStYtbjQWtpmKadmvyUXq8")
	CcipLogicReceiver          = solana.MustPublicKeyFromBase58("CtEVnHsQzhTNWav8skikiV2oF6Xx7r7uGGa8eCDQtTjH")
	CcipTokenReceiver          = solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb")
	CcipInvalidReceiverProgram = solana.MustPublicKeyFromBase58("9Vjda3WU2gsJgE4VdU6QuDw8rfHLyigfFyWs3XDPNUn8")
	CcipTokenPoolProgram       = solana.MustPublicKeyFromBase58("GRvFSLwR7szpjgNEZbGe4HtxfJYXqySXuuRUAJDpu4WH")
	Token2022Program           = solana.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")

	RouterConfigPDA, _, _                    = solana.FindProgramAddress([][]byte{[]byte("config")}, CcipRouterProgram)
	RouterStatePDA, _, _                     = solana.FindProgramAddress([][]byte{[]byte("state")}, CcipRouterProgram)
	ExternalExecutionConfigPDA, _, _         = solana.FindProgramAddress([][]byte{[]byte("external_execution_config")}, CcipRouterProgram)
	ExternalTokenPoolsSignerPDA, _, _        = solana.FindProgramAddress([][]byte{[]byte("external_token_pools_signer")}, CcipRouterProgram)
	ReceiverTargetAccountPDA, _, _           = solana.FindProgramAddress([][]byte{[]byte("counter")}, CcipLogicReceiver)
	ReceiverExternalExecutionConfigPDA, _, _ = solana.FindProgramAddress([][]byte{[]byte("external_execution_config")}, CcipLogicReceiver)
	BillingSignerPDA, _, _                   = solana.FindProgramAddress([][]byte{[]byte("fee_billing_signer")}, CcipRouterProgram)

	BillingTokenConfigPrefix = []byte("fee_billing_token_config")
	DestChainConfigPrefix    = []byte("destination_billing_config")

	SVMChainSelector uint64 = 15
	EvmChainSelector uint64 = 21
	EvmChainLE              = common.Uint64ToLE(EvmChainSelector)

	DefaultCcipVersion      = ccip_router.CcipVersion{Major: 1, Minor: 6}
	defaultCcipVersionBytes = []byte{DefaultCcipVersion.Major, DefaultCcipVersion.Minor}

	SVMSourceChainStatePDA, _, _ = solana.FindProgramAddress([][]byte{[]byte("source_chain_state"), defaultCcipVersionBytes, binary.LittleEndian.AppendUint64([]byte{}, SVMChainSelector)}, CcipRouterProgram)
	EvmSourceChainStatePDA, _, _ = solana.FindProgramAddress([][]byte{[]byte("source_chain_state"), defaultCcipVersionBytes, binary.LittleEndian.AppendUint64([]byte{}, EvmChainSelector)}, CcipRouterProgram)
	SVMDestChainStatePDA, _, _   = solana.FindProgramAddress([][]byte{[]byte("dest_chain_state"), binary.LittleEndian.AppendUint64([]byte{}, SVMChainSelector)}, CcipRouterProgram)
	EvmDestChainStatePDA, _, _   = solana.FindProgramAddress([][]byte{[]byte("dest_chain_state"), binary.LittleEndian.AppendUint64([]byte{}, EvmChainSelector)}, CcipRouterProgram)

	OnRampAddress        = []byte{1, 2, 3}
	OnRampAddressPadded  = [64]byte{1, 2, 3}
	EnableExecutionAfter = int64(1800) // 30min

	MaxOracles                      = 16
	OcrF                      uint8 = 5
	ConfigDigest                    = common.MakeRandom32ByteArray()
	Empty24Byte                     = [24]byte{}
	MaxSignersAndTransmitters       = 16
)
