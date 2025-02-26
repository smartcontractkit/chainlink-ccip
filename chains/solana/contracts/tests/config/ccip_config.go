package config

import (
	"encoding/hex"

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
	CcipOfframpProgram         = solana.MustPublicKeyFromBase58("offRPDpDxT5MGFNmMh99QKTZfPWTkqYUrStEriAS1H5")
	RMNRemoteProgram           = solana.MustPublicKeyFromBase58("CPkyVFQmyzmb6HfDE5TQr3NmZAsydcYspBoc3bf6Zo5x")

	// test values
	OnRampAddress                   = []byte{1, 2, 3}
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
	// bytes4(keccak256("CCIP EVMExtraArgsV2"));
	EvmChainFamilySelector, _ = hex.DecodeString("2812d52c")
	// bytes4(keccak256("CCIP SVMExtraArgsV1"));
	SvmChainFamilySelector, _ = hex.DecodeString("1e10bdc4")

	// example programs
	CcipBaseSender          = solana.MustPublicKeyFromBase58("CcipSender111111111111111111111111111111111")
	CcipBaseReceiver        = solana.MustPublicKeyFromBase58("CcipReceiver1111111111111111111111111111111")
	CcipBasePoolBurnMint    = solana.MustPublicKeyFromBase58("TokenPooL11111111111111111111111111BurnMint")
	CcipBasePoolLockRelease = solana.MustPublicKeyFromBase58("TokenPooL11111111111111111111111LockReLease")

	// router/onramp PDAs
	RouterConfigPDA, _, _                    = state.FindConfigPDA(CcipRouterProgram)
	ExternalTokenPoolsSignerPDA, _, _        = state.FindExternalTokenPoolsSignerPDA(CcipRouterProgram)
	ReceiverTargetAccountPDA, _, _           = solana.FindProgramAddress([][]byte{[]byte("counter")}, CcipLogicReceiver)
	ReceiverExternalExecutionConfigPDA, _, _ = state.FindExternalExecutionConfigPDA(CcipLogicReceiver)
	BillingSignerPDA, _, _                   = state.FindFeeBillingSignerPDA(CcipRouterProgram)
	SvmDestChainStatePDA, _                  = state.FindDestChainStatePDA(SvmChainSelector, CcipRouterProgram)
	EvmDestChainStatePDA, _                  = state.FindDestChainStatePDA(EvmChainSelector, CcipRouterProgram)
	AllowedOfframpEvmPDA, _                  = state.FindAllowedOfframpPDA(EvmChainSelector, CcipOfframpProgram, CcipRouterProgram)
	AllowedOfframpSvmPDA, _                  = state.FindAllowedOfframpPDA(SvmChainSelector, CcipOfframpProgram, CcipRouterProgram)

	// Offramp PDAs
	OfframpConfigPDA, _, _                  = state.FindOfframpConfigPDA(CcipOfframpProgram)
	OfframpReferenceAddressesPDA, _, _      = state.FindOfframpReferenceAddressesPDA(CcipOfframpProgram)
	OfframpEvmSourceChainPDA, _, _          = state.FindOfframpSourceChainPDA(EvmChainSelector, CcipOfframpProgram)
	OfframpSvmSourceChainPDA, _, _          = state.FindOfframpSourceChainPDA(SvmChainSelector, CcipOfframpProgram)
	OfframpBillingSignerPDA, _, _           = state.FindOfframpBillingSignerPDA(CcipOfframpProgram)
	OfframpStatePDA, _, _                   = state.FindOfframpStatePDA(CcipOfframpProgram)
	OfframpExternalExecutionConfigPDA, _, _ = state.FindExternalExecutionConfigPDA(CcipOfframpProgram)
	OfframpTokenPoolsSignerPDA, _, _        = state.FindExternalTokenPoolsSignerPDA(CcipOfframpProgram)

	// fee quoter PDAs
	FqConfigPDA, _, _                     = state.FindFqConfigPDA(FeeQuoterProgram)
	FqEvmDestChainPDA, _, _               = state.FindFqDestChainPDA(EvmChainSelector, FeeQuoterProgram)
	FqSvmDestChainPDA, _, _               = state.FindFqDestChainPDA(SvmChainSelector, FeeQuoterProgram)
	FqAllowedPriceUpdaterOfframpPDA, _, _ = state.FindFqAllowedPriceUpdaterPDA(OfframpBillingSignerPDA, FeeQuoterProgram)

	// RMN Remote PDAs
	RMNRemoteConfigPDA, _, _ = state.FindRMNRemoteConfigPDA(RMNRemoteProgram)
	RMNRemoteCursesPDA, _, _ = state.FindRMNRemoteCursesPDA(RMNRemoteProgram)
)
