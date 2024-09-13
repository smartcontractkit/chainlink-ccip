package exectypes

import (
	"encoding/json"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// CommitObservations contain the commit plugin report data organized by the source chain selector.
type CommitObservations map[cciptypes.ChainSelector][]CommitData

// MessageObservations contain messages for commit plugin reports organized by source chain selector
// and sequence number.
type MessageObservations map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message

type FeeTokenPriceObservations map[types.Account]cciptypes.TokenPrice

// NonceObservations contain the latest nonce for senders in the previously observed messages.
// Nonces are organized by source chain selector and the string encoded sender address. The address
// must be encoding according to the destination chain requirements with typeconv.AddressBytesToString.
type NonceObservations map[cciptypes.ChainSelector]map[string]uint64

// Observation is the observation of the ExecutePlugin.
// TODO: revisit observation types. The maps used here are more space efficient and easier to work
// with but require more transformations compared to the on-chain representations.
type Observation struct {
	// CommitReports are determined during the first phase of execute.
	// It contains the commit reports we would like to execute in the following round.
	CommitReports CommitObservations `json:"commitReports"`

	// Messages are determined during the second phase of execute.
	// Ideally, it contains all the messages identified by the previous outcome's
	// NextCommits. With the previous outcome, and these messages, we can build the
	// execute report.
	Messages MessageObservations `json:"messages"`

	// FeeTokenPrices are determined during the second phase of execute.
	// We need to observe the token prices for fee tokens so that we can calculate the fees in USD for each
	// message and compare them to the execution cost for each message in Outcome. If the fees are greater than
	// the execution cost, we will not execute the message.
	FeeTokenPrices FeeTokenPriceObservations `json:"feeTokenPrices"`

	// Nonces are determined during the third phase of execute.
	// It contains the nonces of senders who are being considered for the final report.
	Nonces NonceObservations `json:"nonces"`
}

// NewObservation constructs an Observation object.
func NewObservation(
	commitReports CommitObservations,
	messages MessageObservations,
	nonces NonceObservations,
	feeTokenPrices FeeTokenPriceObservations,
) Observation {
	return Observation{
		CommitReports:  commitReports,
		Messages:       messages,
		Nonces:         nonces,
		FeeTokenPrices: feeTokenPrices,
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
