package config

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/state"
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

	RouterConfigPDA, _, _                    = state.FindConfigPDA(CcipRouterProgram)
	RouterStatePDA, _, _                     = state.FindStatePDA(CcipRouterProgram)
	ExternalExecutionConfigPDA, _, _         = state.FindExternalExecutionConfigPDA(CcipRouterProgram)
	ExternalTokenPoolsSignerPDA, _, _        = state.FindExternalTokenPoolsSignerPDA(CcipRouterProgram)
	ReceiverTargetAccountPDA, _, _           = solana.FindProgramAddress([][]byte{[]byte("counter")}, CcipLogicReceiver)
	ReceiverExternalExecutionConfigPDA, _, _ = state.FindExternalExecutionConfigPDA(CcipLogicReceiver)
	BillingSignerPDA, _, _                   = state.FindFeeBillingSignerPDA(CcipRouterProgram)

	SVMChainSelector uint64 = 15
	EvmChainSelector uint64 = 21
	EvmChainLE              = common.Uint64ToLE(EvmChainSelector)

	DefaultCcipVersion = ccip_router.CcipVersion{Major: 1, Minor: 6}
	BumpedCcipVersion  = ccip_router.CcipVersion{Major: DefaultCcipVersion.Major, Minor: DefaultCcipVersion.Minor + 1}

	SVMSourceChainStatePDA, _       = state.FindSourceChainStatePDA(DefaultCcipVersion, SVMChainSelector, CcipRouterProgram)
	SVMDestChainStatePDA, _         = state.FindDestChainStatePDA(SVMChainSelector, CcipRouterProgram)
	EvmSourceChainStatePDA, _       = state.FindSourceChainStatePDA(DefaultCcipVersion, EvmChainSelector, CcipRouterProgram)
	EvmBumpedSourceChainStatePDA, _ = state.FindSourceChainStatePDA(BumpedCcipVersion, EvmChainSelector, CcipRouterProgram)
	EvmDestChainStatePDA, _         = state.FindDestChainStatePDA(EvmChainSelector, CcipRouterProgram)

	OnRampAddress        = []byte{1, 2, 3}
	OnRampAddressPadded  = [64]byte{1, 2, 3}
	EnableExecutionAfter = int64(1800) // 30min

	MaxOracles                      = 16
	OcrF                      uint8 = 5
	ConfigDigest                    = common.MakeRandom32ByteArray()
	Empty24Byte                     = [24]byte{}
	MaxSignersAndTransmitters       = 16
)
