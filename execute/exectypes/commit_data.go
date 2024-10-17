package exectypes

import (
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// CommitData is the data that is committed to the chain.
type CommitData struct {
	// SourceChain of the chain that contains the commit report.
	SourceChain cciptypes.ChainSelector `json:"chainSelector"`
	// Timestamp of the block that contains the commit.
	Timestamp time.Time `json:"timestamp"`
	// BlockNum of the block that contains the commit.
	BlockNum uint64 `json:"blockNum"`
	// MerkleRoot of the messages that are in this commit report.
	MerkleRoot cciptypes.Bytes32 `json:"merkleRoot"`
	// SequenceNumberRange of the messages that are in this commit report.
	SequenceNumberRange cciptypes.SeqNumRange `json:"sequenceNumberRange"`

	// Messages that are part of the commit report.
	Messages []cciptypes.Message `json:"messages"`

	// ExecutedMessages are the messages in this report that have already been executed.
	ExecutedMessages []cciptypes.SeqNum `json:"executedMessages"`

	// CostlyMessages are the message IDs of messages that cost more to execute than was paid to execute them (i.e.
	// source fee < execution cost). These messages will not be executed in the current round, but may be executed in
	// future rounds (e.g. if gas prices decrease or if these messages' fees are boosted high enough).
	CostlyMessages []cciptypes.Bytes32 `json:"costlyMessages"`

	// TokenData for each message.
	MessageTokenData []MessageTokenData `json:"messageTokenData"`
}
