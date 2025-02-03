package config

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
)

var (
	PrintEvents       = true
	DefaultCommitment = rpc.CommitmentConfirmed

	// program ids
	CcipRouterProgram          = solana.MustPublicKeyFromBase58("C8WSPj3yyus1YN3yNB6YA5zStYtbjQWtpmKadmvyUXq8")
	CcipLogicReceiver          = solana.MustPublicKeyFromBase58("CtEVnHsQzhTNWav8skikiV2oF6Xx7r7uGGa8eCDQtTjH")
	CcipTokenReceiver          = solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb")
	CcipInvalidReceiverProgram = solana.MustPublicKeyFromBase58("9Vjda3WU2gsJgE4VdU6QuDw8rfHLyigfFyWs3XDPNUn8")
	CcipTokenPoolProgram       = solana.MustPublicKeyFromBase58("GRvFSLwR7szpjgNEZbGe4HtxfJYXqySXuuRUAJDpu4WH")
	Token2022Program           = solana.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")
	FeeQuoterProgram           = solana.MustPublicKeyFromBase58("FeeVB9Q77QvyaENRL1i77BjW6cTkaWwNLjNbZg9JHqpw")

	// test values
	OnRampAddress                   = []byte{1, 2, 3}
	OnRampAddressPadded             = [64]byte{1, 2, 3}
	EnableExecutionAfter            = int64(1800) // 30min
	MaxOracles                      = 16
	OcrF                      uint8 = 5
	ConfigDigest                    = common.MakeRandom32ByteArray()
	Empty24Byte                     = [24]byte{}
	MaxSignersAndTransmitters       = 16

	// chain selectors
	SvmChainSelector uint64 = 15
	EvmChainSelector uint64 = 21
	EvmChainLE              = common.Uint64ToLE(EvmChainSelector)

	// router/onramp PDAs
	// example programs
	CcipBaseReceiver = solana.MustPublicKeyFromBase58("CcipReceiver1111111111111111111111111111111")
	CcipBaseSender   = solana.MustPublicKeyFromBase58("CcipSender111111111111111111111111111111111")

	RouterConfigPDA, _, _                    = state.FindConfigPDA(CcipRouterProgram)
	RouterStatePDA, _, _                     = state.FindStatePDA(CcipRouterProgram)
	ExternalExecutionConfigPDA, _, _         = state.FindExternalExecutionConfigPDA(CcipRouterProgram)
	ExternalTokenPoolsSignerPDA, _, _        = state.FindExternalTokenPoolsSignerPDA(CcipRouterProgram)
	ReceiverTargetAccountPDA, _, _           = solana.FindProgramAddress([][]byte{[]byte("counter")}, CcipLogicReceiver)
	ReceiverExternalExecutionConfigPDA, _, _ = state.FindExternalExecutionConfigPDA(CcipLogicReceiver)
	BillingSignerPDA, _, _                   = state.FindFeeBillingSignerPDA(CcipRouterProgram)
	SvmSourceChainStatePDA, _                = state.FindSourceChainStatePDA(SvmChainSelector, CcipRouterProgram)
	SvmDestChainStatePDA, _                  = state.FindDestChainStatePDA(SvmChainSelector, CcipRouterProgram)
	EvmSourceChainStatePDA, _                = state.FindSourceChainStatePDA(EvmChainSelector, CcipRouterProgram)
	EvmDestChainStatePDA, _                  = state.FindDestChainStatePDA(EvmChainSelector, CcipRouterProgram)

	// fee quoter PDAs
	FqConfigPDA, _, _        = state.FindFqConfigPDA(FeeQuoterProgram)
	FqEvmDestChainPDA, _, _  = state.FindFqDestChainPDA(EvmChainSelector, FeeQuoterProgram)
	FqSvmDestChainPDA, _, _  = state.FindFqDestChainPDA(SvmChainSelector, FeeQuoterProgram)
	FqBillingSignerPDA, _, _ = state.FindFqBillingSignerPDA(FeeQuoterProgram)
)
