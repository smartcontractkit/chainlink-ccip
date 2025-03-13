package config

import (
	"github.com/gagliardetto/solana-go"
)

var (
	McmProgram = GetProgramID("mcm")
	// For testing CPIs made by other programs (with actual business logic).
	ExternalCpiStubProgram = GetProgramID("external_program_cpi_stub")
	StubAccountPDA, _, _   = solana.FindProgramAddress([][]byte{[]byte("u8_value")}, ExternalCpiStubProgram)

	// ChainID Configuration
	// --------------------
	// Note: This is an arbitrary value used only for localnet testing
	// Value (0x4808e31713a26612) derived from keccak256("solana:localnet")
	//
	// Note: CCIP chain-selector uses genesis hash of each SVM network
	// (mainnet-beta, devnet, testnet) to determine their chain IDs.
	// See: chain-selector specification
	TestChainID             uint64 = 5190648258797659666
	TestChainIDPaddedBuffer        = [32]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x66, 0xa2, 0x13, 0x17, 0xe3, 0x08, 0x48,
	}

	// [0,0,0,...'t','e','s','t','-','m','c','m',]
	TestMsigID = [32]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x6d, 0x63, 0x6d,
	}
	MaxNumSigners               = 180
	MaxAppendSignerBatchSize    = 45
	MaxAppendSignatureBatchSize = 13

	McmEmptyRoot      = [32]byte{}
	McmEmptyTimestamp = uint32(0)
	// root related configs
	// the following diagram shows the structure of the signers and groups:
	// ref: https://github.com/smartcontractkit/ccip-owner-contracts/blob/56f1a8d2cd4ba5ef2b99d2185ffded53957dd410/src/ManyChainMultiSig.sol#L65
	//                    ┌──────┐ root
	//                 ┌─►│2-of-3│◄───────┐
	//                 │  └──────┘        │
	//                 │        ▲         │
	//                 │ group1 │ group2  │
	//              ┌──┴───┐ ┌──┴───┐ ┌───┴────┐
	//          ┌──►│1-of-2│ │2-of-2│ │signer A│
	//          │   └──────┘ └──────┘ └────────┘
	//          │       ▲      ▲  ▲      group3
	//          │       │      │  │     ┌──────┐
	//          │       │      │  └─────┤1-of-2│◄─┐
	//          │       │      │        └──────┘  │
	//  ┌───────┴┐ ┌────┴───┐ ┌┴───────┐ ▲        │
	//  │signer B│ │signer C│ │signer D│ │        │
	//  └────────┘ └────────┘ └────────┘ │        │
	//                                   │        │
	//                            ┌──────┴─┐ ┌────┴───┐
	//                            │signer E│ │signer F│
	//                            └────────┘ └────────┘
	SignerPrivateKeys = []string{
		"aa4dc5ba14d8921dca4f486a0b5bc573502e33d9093025479fc52f22e8d8a4b7",
		"82dbcde99a61371aaa0aee75b04fa5663832a1041232ed8de868b4f4b186f00a",
		"1902e25f992351a3194e16617f879b67c6f7a8a084fea94b77c5b4f1398ff5a6",
		"f25eee250faea64943df683e48a3c804d4616ad35a4dab5654428916bea7e234",
		"169aa793d6325a77b3454572da45d421c189bef166ab778b5f35023522046916",
		"88dfb2c77c655efb5c9d723b4518ea8a7c5416a9951c7f192e254cb446b8db6c",
	}
	SignerGroups = []byte{0, 1, 1, 2, 3, 3}
	GroupQuorums = []uint8{2, 1, 2, 1}
	GroupParents = []uint8{0, 0, 0, 2}
	ClearRoot    = false
)
