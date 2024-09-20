package exectypes

import (
	"encoding/json"
	"fmt"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
)

// CommitObservations contain the commit plugin report data organized by the source chain selector.
type CommitObservations map[cciptypes.ChainSelector][]CommitData

// MessageObservations contain messages for commit plugin reports organized by source chain selector
// and sequence number.
type MessageObservations map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message

// NonceObservations contain the latest nonce for senders in the previously observed messages.
// Nonces are organized by source chain selector and the string encoded sender address. The address
// must be encoding according to the destination chain requirements with typeconv.AddressBytesToString.
type NonceObservations map[cciptypes.ChainSelector]map[string]uint64

// TokenDataObservations contain token data for messages organized by source chain selector and sequence number.
// There could be multiple tokens per a single message, so MessageTokenData is a slice of TokenData.
// TokenDataObservations are populated during the Observation phase and depend on previously fetched
// MessageObservations details and the `tokenDataObservers` configured in the ExecuteOffchainConfig.
// Content of the MessageTokenData is determined by the tokendata.TokenDataObserver implementations.
//   - if Message doesn't have any tokens, TokenData slice will be empty.
//   - if Message has tokens, but these tokens don't require any special treatment,
//     TokenData slice will contain empty TokenData objects.
//   - if Message has tokens and these tokens require additional processing defined in ExecuteOffchainConfig,
//     specific tokendata.TokenDataObserver will be used to populate the TokenData slice.
type TokenDataObservations map[cciptypes.ChainSelector]map[cciptypes.SeqNum]MessageTokenData

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
	Ready bool   `json:"ready"`
	Data  []byte `json:"data"`
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

// Observation is the observation of the ExecutePlugin.
// TODO: revisit observation types. The maps used here are more space efficient and easier to work
// with but require more transformations compared to the on-chain representations.
type Observation struct {
	// CommitReports are determined during the first phase of execute.
	// It contains the commit reports we would like to execute in the following round.
	CommitReports CommitObservations `json:"commitReports"`

	// Messages are determined during the second phase of execute.
	// Ideally, it contains all the messages identified by the previous outcome's
	// NextCommits. With the previous outcome, and these messsages, we can build the
	// execute report.
	Messages MessageObservations `json:"messages"`

	// TokenData are determined during the second phase of execute.
	// It contains the token data for the messages identified in the same stage as Messages
	TokenData TokenDataObservations `json:"tokenDataObservations"`

	// Nonces are determined during the third phase of execute.
	// It contains the nonces of senders who are being considered for the final report.
	Nonces NonceObservations `json:"nonces"`

	// Contracts are part of the initial discovery phase which runs to initialize the CCIP Reader.
	Contracts dt.Observation `json:"contracts"`
}

// NewObservation constructs an Observation object.
func NewObservation(
	commitReports CommitObservations,
	messages MessageObservations,
	tokenData TokenDataObservations,
	nonces NonceObservations,
	contracts dt.Observation,
) Observation {
	return Observation{
		CommitReports: commitReports,
		Messages:      messages,
		TokenData:     tokenData,
		Nonces:        nonces,
		Contracts:     contracts,
	}
}

// Encode the Observation into a byte slice.
func (obs Observation) Encode() ([]byte, error) {
	return json.Marshal(obs)
}

// DecodeObservation from a byte slice into an Observation.
func DecodeObservation(b []byte) (Observation, error) {
	if len(b) == 0 {
		return Observation{}, nil
	}
	obs := Observation{}
	err := json.Unmarshal(b, &obs)
	return obs, err
}
