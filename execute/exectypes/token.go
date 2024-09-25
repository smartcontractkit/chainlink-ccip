package exectypes

import (
	"fmt"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type MessageTokenData struct {
	TokenData []TokenData
}

func NewMessageTokenData(tokenData ...TokenData) MessageTokenData {
	if len(tokenData) == 0 {
		return MessageTokenData{TokenData: []TokenData{}}
	}
	return MessageTokenData{TokenData: tokenData}
}

func (mtd MessageTokenData) IsReady() bool {
	for _, td := range mtd.TokenData {
		if !td.IsReady() {
			return false
		}
	}
	return true
}

func (mtd MessageTokenData) Error() error {
	for _, td := range mtd.TokenData {
		if td.Error != nil {
			return td.Error
		}
	}
	return nil
}

func (mtd MessageTokenData) ToByteSlice() [][]byte {
	out := make([][]byte, len(mtd.TokenData))
	for i, td := range mtd.TokenData {
		out[i] = td.Data
	}
	return out
}

// TokenData is the token data for a single token in a message.
// It contains the token data and a flag indicating if the data is ready.
type TokenData struct {
	Ready bool            `json:"ready"`
	Data  cciptypes.Bytes `json:"data"`
	// Error and Supported are used only for internal processing, we don't want nodes to gossip about the
	// internals they see during processing
	Error     error `json:"-"`
	Supported bool  `json:"-"`
}

// NotSupportedTokenData returns a TokenData object with Supported set to false.
// It should be returned by the Observer for tokens that are not supported.
func NotSupportedTokenData() TokenData {
	return TokenData{
		Ready:     false,
		Error:     nil,
		Data:      nil,
		Supported: false,
	}
}

// NewNoopTokenData returns a TokenData object with Ready set to true and empty data.
// It's used for marking tokens that don't require offchain processing.
func NewNoopTokenData() TokenData {
	return TokenData{
		Ready:     true,
		Error:     nil,
		Data:      []byte{},
		Supported: true,
	}
}

// NewSuccessTokenData returns a TokenData object with Ready set to true and the provided data.
func NewSuccessTokenData(data []byte) TokenData {
	return TokenData{
		Ready:     true,
		Error:     nil,
		Data:      data,
		Supported: true,
	}
}

// NewErrorTokenData returns a TokenData object with Ready set to false and the provided error.
func NewErrorTokenData(err error) TokenData {
	return TokenData{
		Ready:     false,
		Error:     err,
		Data:      nil,
		Supported: true,
	}
}

func (td TokenData) IsReady() bool {
	return td.Ready
}

// MessageTokenID is a unique identifier for a message token data (per chain selector). It's a composite key of
// the message sequence number and the token index within the message. It's used to easier identify token data for
// messages without having to deal with nested maps.
type MessageTokenID struct {
	SeqNr cciptypes.SeqNum
	Index int
}

func NewMessageTokenID(seqNr cciptypes.SeqNum, index int) MessageTokenID {
	return MessageTokenID{SeqNr: seqNr, Index: index}
}

func (mti MessageTokenID) String() string {
	return fmt.Sprintf("%d_%d", mti.SeqNr, mti.Index)
}
