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
	CcipRouterProgram          = solana.MustPublicKeyFromBase58("7gtMT88cNmZVmVUzkcC6MKZ4NY27ynwodKBVJFQdq8R2")
	CcipLogicReceiver          = solana.MustPublicKeyFromBase58("EvhgrPhTDt4LcSPS2kfJgH6T6XWZ6wT3X9ncDGLT1vui")
	CcipTokenReceiver          = solana.MustPublicKeyFromBase58("DS2tt4BX7YwCw7yrDNwbAdnYrxjeCPeGJbHmZEYC8RTb")
	CcipInvalidReceiverProgram = solana.MustPublicKeyFromBase58("FmyF3oW69MSAhyPSiZ69C4RKBdCPv5vAFTScisV7Me2j")
	CcipTokenPoolProgram       = solana.MustPublicKeyFromBase58("JuCcZ4smxAYv9QHJ36jshA7pA3FuQ3vQeWLUeAtZduJ")
	Token2022Program           = solana.MustPublicKeyFromBase58("TokenzQdBNbLqP5VEhdkAS6EPFLC1PHnBqCXEpPxuEb")
	FeeQuoterProgram           = solana.MustPublicKeyFromBase58("CJaqnyk1hSrDcg8CM2CyWarkgSUDHueUeKKbD6gNM535")
	CcipOfframpProgram         = solana.MustPublicKeyFromBase58("8Qqj9FquL7u6aR7GuY3gePmtmHiZ9Qrbf52mv83CyDHR")
	RMNRemoteProgram           = solana.MustPublicKeyFromBase58("6VdacTfZjiyVkvfGjUfXEgsFSL9HPskyj3MdLqvnqC94")

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
	CcipBaseSender          = solana.MustPublicKeyFromBase58("4LfBQWYaU6zQZbDyYjX8pbY4qjzrhoumUFYZEZEqMNhJ")
	CcipBaseReceiver        = solana.MustPublicKeyFromBase58("48LGpn6tPn5SjTtK2wL9uUx48JUWZdZBv11sboy2orCc")
	CcipBasePoolBurnMint    = solana.MustPublicKeyFromBase58("41FGToCmdaWa1dgZLKFAjvmx6e6AjVTX7SVRibvsMGVB")
	CcipBasePoolLockRelease = solana.MustPublicKeyFromBase58("8eqh8wppT9c5rw4ERqNCffvU6cNFJWff9WmkcYtmGiqC")

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
