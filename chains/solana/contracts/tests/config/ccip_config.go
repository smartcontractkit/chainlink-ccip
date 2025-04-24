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
	CcipRouterProgram          = GetProgramID("ccip_router")
	CcipLogicReceiver          = GetProgramID("test_ccip_receiver")
	CcipTokenReceiver          = solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb")
	CcipInvalidReceiverProgram = GetProgramID("test_ccip_invalid_receiver")
	CcipTokenPoolProgram       = GetProgramID("test_token_pool")
	Token2022Program           = solana.Token2022ProgramID
	FeeQuoterProgram           = GetProgramID("fee_quoter")
	CcipOfframpProgram         = GetProgramID("ccip_offramp")
	RMNRemoteProgram           = GetProgramID("rmn_remote")
	PingPongProgram            = GetProgramID("ping_pong_demo")

	// test values
	OnRampAddress                   = []byte{1, 2, 3}
	EnableExecutionAfter            = int64(1800) // 30min
	MaxOracles                      = 16
	OcrF                      uint8 = 5
	ConfigDigest                    = common.MakeRandom32ByteArray()
	Empty24Byte                     = [24]byte{}
	MaxSignersAndTransmitters       = 16

	// valid evm addresses(over precompiled space) - for remote token addresses
	EVMToken0AddressBytes = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 4, 0} // 1024
	EVMToken1AddressBytes = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	EVMToken2AddressBytes = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 20, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	EVMToken3AddressBytes = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 21, 19, 18, 17, 16, 15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	// chain selectors
	SvmChainSelector uint64 = 15
	EvmChainSelector uint64 = 21
	EvmChainLE              = common.Uint64ToLE(EvmChainSelector)
	// bytes4(keccak256("CCIP EVMExtraArgsV2"));
	EvmChainFamilySelector, _ = hex.DecodeString("2812d52c")
	// bytes4(keccak256("CCIP SVMExtraArgsV1"));
	SvmChainFamilySelector, _ = hex.DecodeString("1e10bdc4")

	// example programs
	CcipBaseSender          = GetProgramID("example_ccip_sender")
	CcipBaseReceiver        = GetProgramID("example_ccip_receiver")
	CcipBasePoolBurnMint    = GetProgramID("burnmint_token_pool")
	CcipBasePoolLockRelease = GetProgramID("lockrelease_token_pool")

	// router/onramp PDAs
	RouterConfigPDA, _, _                    = state.FindConfigPDA(CcipRouterProgram)
	ReceiverTargetAccountPDA, _, _           = solana.FindProgramAddress([][]byte{[]byte("counter")}, CcipLogicReceiver)
	ReceiverExternalExecutionConfigPDA, _, _ = solana.FindProgramAddress([][]byte{[]byte("external_execution_config")}, CcipLogicReceiver)
	BillingSignerPDA, _, _                   = state.FindFeeBillingSignerPDA(CcipRouterProgram)
	SvmDestChainStatePDA, _                  = state.FindDestChainStatePDA(SvmChainSelector, CcipRouterProgram)
	EvmDestChainStatePDA, _                  = state.FindDestChainStatePDA(EvmChainSelector, CcipRouterProgram)
	AllowedOfframpEvmPDA, _                  = state.FindAllowedOfframpPDA(EvmChainSelector, CcipOfframpProgram, CcipRouterProgram)
	AllowedOfframpSvmPDA, _                  = state.FindAllowedOfframpPDA(SvmChainSelector, CcipOfframpProgram, CcipRouterProgram)

	// Offramp PDAs
	OfframpConfigPDA, _, _               = state.FindOfframpConfigPDA(CcipOfframpProgram)
	OfframpReferenceAddressesPDA, _, _   = state.FindOfframpReferenceAddressesPDA(CcipOfframpProgram)
	OfframpEvmSourceChainPDA, _, _       = state.FindOfframpSourceChainPDA(EvmChainSelector, CcipOfframpProgram)
	OfframpSvmSourceChainPDA, _, _       = state.FindOfframpSourceChainPDA(SvmChainSelector, CcipOfframpProgram)
	OfframpBillingSignerPDA, _, _        = state.FindOfframpBillingSignerPDA(CcipOfframpProgram)
	OfframpStatePDA, _, _                = state.FindOfframpStatePDA(CcipOfframpProgram)
	OfframpReceiverExternalExecPDA, _, _ = state.FindExternalExecutionConfigPDA(CcipLogicReceiver, CcipOfframpProgram)

	// fee quoter PDAs
	FqConfigPDA, _, _                     = state.FindFqConfigPDA(FeeQuoterProgram)
	FqEvmDestChainPDA, _, _               = state.FindFqDestChainPDA(EvmChainSelector, FeeQuoterProgram)
	FqSvmDestChainPDA, _, _               = state.FindFqDestChainPDA(SvmChainSelector, FeeQuoterProgram)
	FqAllowedPriceUpdaterOfframpPDA, _, _ = state.FindFqAllowedPriceUpdaterPDA(OfframpBillingSignerPDA, FeeQuoterProgram)

	// RMN Remote PDAs
	RMNRemoteConfigPDA, _, _ = state.FindRMNRemoteConfigPDA(RMNRemoteProgram)
	RMNRemoteCursesPDA, _, _ = state.FindRMNRemoteCursesPDA(RMNRemoteProgram)
)
