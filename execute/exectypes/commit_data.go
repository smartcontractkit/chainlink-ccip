package exectypes

import (
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// CommitData is the data that is committed to the chain.
// NOTE: This data structure represents a committed root (within a commit report).
type CommitData struct {
	// SourceChain of the chain that contains the commit report.
	SourceChain cciptypes.ChainSelector `json:"chainSelector"`
	// OnRampAddress used on the source chain.
	OnRampAddress cciptypes.UnknownAddress
	// Timestamp of the block that contains the commit.
	Timestamp time.Time `json:"timestamp"`
	// BlockNum of the block that contains the commit.
	BlockNum uint64 `json:"blockNum"`
	// MerkleRoot of the messages that are in this commit report.
	MerkleRoot cciptypes.Bytes32 `json:"merkleRoot"`
	// SequenceNumberRange of the messages that are in this commit report.
	SequenceNumberRange cciptypes.SeqNumRange `json:"sequenceNumberRange"`
	// ExecutedMessages are the messages in this report that have already been executed.
	ExecutedMessages []cciptypes.SeqNum `json:"executedMessages"`

	/////////////////////////////////////////////////////////////////////
	// Fields below here are not always present.                       //
	// They are present only in outcomes and Filtered pending reports. //
	/////////////////////////////////////////////////////////////////////

	// Messages that are part of the commit report.
	Messages []cciptypes.Message `json:"messages"`
	// Hashes are the hashes of the respective messages in Messages slice.
	// Get populated during GetMessages Outcome phase
	// Length of this slice should equal to the length of Messages slice.
	Hashes []cciptypes.Bytes32 `json:"messageHashes"`

	// TokenData for each message.
	// Length of this slice should equal to the length of Messages slice.
	MessageTokenData []MessageTokenData `json:"messageTokenData"`
}

// LessThan CommitData comparator
// Ordering priority is:
// 1. By timestamp (earlier timestamps first)
// 2. By source chain selector (smaller chain selectors first)
// 3. By sequence number range (smaller starting sequence numbers first)
// This function is used for stable sorting of CommitData objects.
// CHANGE WITH CAUTION, not giving priority to timestamps causes lane starvation
func LessThan(i, j CommitData) bool {
	if !i.Timestamp.Equal(j.Timestamp) {
		return i.Timestamp.Before(j.Timestamp)
	}
	if i.SourceChain != j.SourceChain {
		return i.SourceChain < j.SourceChain
	}
	return i.SequenceNumberRange.Start() < j.SequenceNumberRange.Start()
}

// CopyNoMsgData creates a copy of the CommitData without the messages.Data
func (cd CommitData) CopyNoMsgData() CommitData {
	msgsWitoutData := make([]cciptypes.Message, len(cd.Messages))
	for i, msg := range cd.Messages {
		msgsWitoutData[i] = msg.CopyWithoutData()
	}
	return CommitData{
		SourceChain:         cd.SourceChain,
		OnRampAddress:       cd.OnRampAddress,
		Timestamp:           cd.Timestamp,
		BlockNum:            cd.BlockNum,
		MerkleRoot:          cd.MerkleRoot,
		SequenceNumberRange: cd.SequenceNumberRange,
		ExecutedMessages:    append([]cciptypes.SeqNum{}, cd.ExecutedMessages...),
		Messages:            msgsWitoutData,
		Hashes:              append([]cciptypes.Bytes32{}, cd.Hashes...),
		MessageTokenData:    append([]MessageTokenData{}, cd.MessageTokenData...),
	}
}
