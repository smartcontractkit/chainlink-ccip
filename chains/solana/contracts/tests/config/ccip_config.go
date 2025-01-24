package config

import (
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

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

	RouterConfigPDA, _, _                    = common.FindConfigPDA(CcipRouterProgram)
	RouterStatePDA, _, _                     = common.FindStatePDA(CcipRouterProgram)
	ExternalExecutionConfigPDA, _, _         = common.FindExternalExecutionConfigPDA(CcipRouterProgram)
	ExternalTokenPoolsSignerPDA, _, _        = common.FindExternalTokenPoolsSignerPDA(CcipRouterProgram)
	ReceiverTargetAccountPDA, _, _           = common.FindCounterPDA(CcipLogicReceiver)
	ReceiverExternalExecutionConfigPDA, _, _ = common.FindExternalExecutionConfigPDA(CcipLogicReceiver)
	BillingSignerPDA, _, _                   = common.FindFeeBillingSignerPDA(CcipRouterProgram)

	SVMChainSelector uint64 = 15
	EvmChainSelector uint64 = 21
	EvmChainLE              = common.Uint64ToLE(EvmChainSelector)

	SVMSourceChainStatePDA, _ = common.GetSourceChainStatePDA(SVMChainSelector, CcipRouterProgram)
	SVMDestChainStatePDA, _   = common.GetDestChainStatePDA(SVMChainSelector, CcipRouterProgram)
	EvmSourceChainStatePDA, _ = common.GetSourceChainStatePDA(EvmChainSelector, CcipRouterProgram)
	EvmDestChainStatePDA, _   = common.GetDestChainStatePDA(EvmChainSelector, CcipRouterProgram)

	OnRampAddress        = []byte{1, 2, 3}
	OnRampAddressPadded  = [64]byte{1, 2, 3}
	EnableExecutionAfter = int64(1800) // 30min

	MaxOracles                      = 16
	OcrF                      uint8 = 5
	ConfigDigest                    = common.MakeRandom32ByteArray()
	Empty24Byte                     = [24]byte{}
	MaxSignersAndTransmitters       = 16
)
