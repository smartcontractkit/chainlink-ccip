package exectypes

import (
	"errors"
	"fmt"

	"golang.org/x/crypto/sha3"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

func (mtd MessageTokenData) Append(index int, td TokenData) MessageTokenData {
	out := mtd
	if index >= len(out.TokenData) {
		newSize := index + 1
		newTokenData := make([]TokenData, newSize)
		// Copy the contents of the old slice into the new slice
		copy(newTokenData, out.TokenData)
		out.TokenData = newTokenData
	}
	out.TokenData[index] = td
	return out
}

func (mtd MessageTokenData) IsReady() bool {
	for _, td := range mtd.TokenData {
		if !td.IsReady() {
			return false
		}
	}
	return true
}

// SupportedAreReady returns true if all the supported TokenData are ready.
func (mtd MessageTokenData) SupportedAreReady() bool {
	for _, td := range mtd.TokenData {
		if td.Supported && !td.IsReady() {
			return false
		}
	}
	return true
}

// Error returns combined errors from all the TokenData children.
// If message IsReady it must return nil. Keep in mind that errors are not preserved when serializing
// TokenDataObservations, so this method is only useful for internal processing. Observation fetched from
// other nodes will return nil even if it's faulty.
func (mtd MessageTokenData) Error() error {
	err := make([]error, 0)
	for _, td := range mtd.TokenData {
		if td.Error != nil {
			err = append(err, td.Error)
		}
	}
	return errors.Join(err...)
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

// NotReadyToken returns a TokenData object with Ready set to false. It doesn't carry additional information,
// it's used to mark tokens which are not ready because consensus hasn't been reached on them.
// By setting them as `Ready=false`, higher level knows this message has to be ignored
func NotReadyToken() TokenData {
	return TokenData{Ready: false}
}

// TokenDataHash returns a hash of the token data. It's used during the consensus process to identify unique token data.
// It intentionally skips the fields that are not relevant for the consensus process and are
// not serialized when nodes gossiping
func TokenDataHash(td TokenData) [32]byte {
	return sha3.Sum256([]byte(fmt.Sprintf("%v_%v", td.Ready, td.Data)))
}

func (td TokenData) IsReady() bool {
	return td.Ready
}
