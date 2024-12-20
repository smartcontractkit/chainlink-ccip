package ccip

import (
	"github.com/gagliardetto/solana-go"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/gobindings/ccip_router"
)

// Events - temporary event struct to decode
// anchor-go does not support events
// https://github.com/fragmetric-labs/solana-anchor-go does but requires upgrade to anchor >= v0.30.0
type EventCCIPMessageSent struct {
	Discriminator            [8]byte
	DestinationChainSelector uint64
	SequenceNumber           uint64
	Message                  ccip_router.Solana2AnyRampMessage
}

type EventCommitReportAccepted struct {
	Discriminator [8]byte
	Report        ccip_router.MerkleRoot
}

type EventTransmitted struct {
	Discriminator  [8]byte
	OcrPluginType  uint8
	ConfigDigest   [32]byte
	SequenceNumber uint64
}

type EventExecutionStateChanged struct {
	Discriminator       [8]byte
	SourceChainSelector uint64
	SequenceNumber      uint64
	MessageID           [32]byte
	MessageHash         [32]byte
	State               ccip_router.MessageExecutionState
}

type EventSkippedAlreadyExecutedMessage struct {
	Discriminator       [8]byte
	SourceChainSelector uint64
	SequenceNumber      uint64
}

type EventConfigSet struct {
	Discriminator [8]byte
	OcrPluginType uint8
	ConfigDigest  [32]byte
	Signers       [][20]uint8
	Transmitters  []solana.PublicKey
	F             uint8
}

type UsdPerTokenUpdated struct {
	Discriminator [8]byte
	Token         solana.PublicKey
	Value         [28]byte
	Timestamp     int64
}

type UsdPerUnitGasUpdated struct {
	Discriminator [8]byte
	DestChain     uint64
	Value         [28]byte
	Timestamp     int64
}
