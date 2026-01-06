package reader

import (
	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// TODO: This is a temporary solution to avoid circular dependencies. There should be a better way to do this.

type MessageTokenID = ccipocr3.MessageTokenID

func NewMessageTokenID(seqNr ccipocr3.SeqNum, index int) MessageTokenID {
	return ccipocr3.NewMessageTokenID(seqNr, index)
}
