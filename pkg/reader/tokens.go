package reader

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// TODO: This is a temporary solution to avoid circular dependencies. There should be a better way to do this.

// MessageTokenID is a unique identifier for a message token data (per chain selector). It's a composite key of
// the message sequence number and the token index within the message. It's used to easier identify token data for
// messages without having to deal with nested maps.
type MessageTokenID struct {
	SeqNr ccipocr3.SeqNum
	Index int
}

func NewMessageTokenID(seqNr ccipocr3.SeqNum, index int) MessageTokenID {
	return MessageTokenID{SeqNr: seqNr, Index: index}
}

func (mti MessageTokenID) String() string {
	return fmt.Sprintf("%d_%d", mti.SeqNr, mti.Index)
}
