package config

import (
	"github.com/gagliardetto/solana-go"
)

var (
	McmProgram = solana.MustPublicKeyFromBase58("6UmMZr5MEqiKWD5jqTJd1WCR5kT8oZuFYBLJFi1o6GQX")

	// For testing CPIs made by other programs (with actual business logic).
	ExternalCpiStubProgram = solana.MustPublicKeyFromBase58("4HeqEoSyfYpeC2goFLj9eHgkxV33mR5G7JYAbRsN14uQ")
	StubAccountPDA, _, _   = solana.FindProgramAddress([][]byte{[]byte("u8_value")}, ExternalCpiStubProgram)

	// todo: update chain id following the latest discussion(genesis hash)
	// Last 8 bytes (uint64) of keccak256("solana:localnet") as big-endian
	// This is 0x4808e31713a26612 --> in little-endian, it is "1266a21317e30848"
	TestChainID             uint64 = 5190648258797659666
	TestChainIDPaddedBuffer        = [32]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x12, 0x66, 0xa2, 0x13, 0x17, 0xe3, 0x08, 0x48,
	}

	// [0,0,0,...'t','e','s','t','-','m','c','m',]
	TestMsigNamePaddedBuffer = [32]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x6d, 0x63, 0x6d,
	}
	MaxNumSigners               = 200
	MaxAppendSignerBatchSize    = 45
	MaxAppendSignatureBatchSize = 13
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
