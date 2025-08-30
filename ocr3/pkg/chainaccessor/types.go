package chainaccessor

import (
	"math/big"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// ---------------------------------------------------
// The following types match the structs defined in the EVM contracts are used to decode these
// on-chain events. They should eventually be replaced by chain-reader modifiers and use the
// base cciptypes.CommitReport type.

// SendRequestedEvent represents the contents of the event emitted by the CCIP OnRamp when a
// message is sent.
type SendRequestedEvent struct {
	DestChainSelector cciptypes.ChainSelector
	SequenceNumber    cciptypes.SeqNum
	Message           cciptypes.Message
}

// CommitReportAcceptedEvent represents the contents of the event emitted by the CCIP OffRamp when a
// commit report is accepted.
type CommitReportAcceptedEvent struct {
	BlessedMerkleRoots   []MerkleRoot
	UnblessedMerkleRoots []MerkleRoot
	PriceUpdates         PriceUpdates
}

// ExecutionStateChangedEvent represents the contents of the event emitted by the CCIP OffRamp
type ExecutionStateChangedEvent struct {
	SourceChainSelector cciptypes.ChainSelector
	SequenceNumber      cciptypes.SeqNum
	MessageID           cciptypes.Bytes32
	MessageHash         cciptypes.Bytes32
	State               uint8
	ReturnData          cciptypes.Bytes
	GasUsed             big.Int
}

type MerkleRoot struct {
	SourceChainSelector uint64
	OnRampAddress       cciptypes.UnknownAddress
	MinSeqNr            uint64
	MaxSeqNr            uint64
	MerkleRoot          cciptypes.Bytes32
}

type TokenPriceUpdate struct {
	SourceToken cciptypes.UnknownAddress
	UsdPerToken *big.Int
}

type GasPriceUpdate struct {
	// DestChainSelector is the chain that the gas price is for (some plugin source chain).
	// Not the chain that the gas price is stored on.
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

type PriceUpdates struct {
	TokenPriceUpdates []TokenPriceUpdate
	GasPriceUpdates   []GasPriceUpdate
}

type chainAddressNonce struct {
	chain    cciptypes.ChainSelector
	address  string
	response uint64
}
