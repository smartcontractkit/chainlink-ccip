package exectypes

import (
	"encoding/json"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
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
// There could be multiple tokens per a single message, so MessageTokensData is a slice of TokenData.
// TokenDataObservations are populated during the Observation phase and depends on previously fetched
// MessageObservations details and the `tokenDataProcessors` configured in the ExecuteOffchainConfig.
// Content of the MessageTokensData is determined by the TokenDataProcessor implementations.
//   - if Message doesn't have any tokens, TokenData slice will be empty.
//   - if Message has tokens, but these tokens don't require any special treatment,
//     TokenData slice will contain empty TokenData objects.
//   - if Message has tokens and these tokens require additional processing defined in ExecuteOffchainConfig,
//     specific TokenDataProcessor will be used to populate the TokenData slice.
type TokenDataObservations map[cciptypes.ChainSelector]map[cciptypes.SeqNum]MessageTokensData

type MessageTokensData struct {
	TokenData []TokenData
}

func (mtd MessageTokensData) IsReady() bool {
	for _, td := range mtd.TokenData {
		if !td.IsReady() {
			return false
		}
	}
	return true
}

func (mtd MessageTokensData) Error() error {
	for _, td := range mtd.TokenData {
		if td.Error != nil {
			return td.Error
		}
	}
	return nil
}

func (mtd MessageTokensData) ToByteSlice() [][]byte {
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
	Error error
}

func NewEmptyTokenData() TokenData {
	return TokenData{
		Ready: false,
		Error: nil,
		Data:  nil,
	}
}

func (td TokenData) IsReady() bool {
	return td.Ready
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
}

// NewObservation constructs a Observation object.
func NewObservation(
	commitReports CommitObservations,
	messages MessageObservations,
	tokenData TokenDataObservations,
	nonces NonceObservations,
) Observation {
	return Observation{
		CommitReports: commitReports,
		Messages:      messages,
		TokenData:     tokenData,
		Nonces:        nonces,
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
