package exectypes

import (
	"context"
	"encoding/json"

	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// CommitObservations contain the commit plugin report data organized by the source chain selector.
type CommitObservations map[cciptypes.ChainSelector][]CommitData

// MessageObservations contain messages for commit plugin reports organized by source chain selector
// and sequence number.
type MessageObservations map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message

type MessageHashes map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Bytes32

type EncodedMsgAndTokenDataSizes map[cciptypes.ChainSelector]map[cciptypes.SeqNum]int

// Flatten nested maps into a slice of messages.
func (mo MessageObservations) Flatten() []cciptypes.Message {
	var results []cciptypes.Message
	for _, msgs := range mo {
		for _, msg := range msgs {
			results = append(results, msg)
		}
	}
	return results
}

func GetHashes(ctx context.Context, mo MessageObservations, hasher cciptypes.MessageHasher) (MessageHashes, error) {
	hashes := make(MessageHashes)
	for chain, msgs := range mo {
		hashes[chain] = make(map[cciptypes.SeqNum]cciptypes.Bytes32)
		for seq, msg := range msgs {
			hash, err := hasher.Hash(ctx, msg)
			if err != nil {
				return nil, err
			}
			hashes[chain][seq] = hash
		}
	}
	return hashes, nil
}

// GetEncodedMsgAndTokenDataSizes calculates the encoded sizes of messages and their token data counterpart.
func GetEncodedMsgAndTokenDataSizes(mo MessageObservations, tds TokenDataObservations) EncodedMsgAndTokenDataSizes {
	sizes := make(EncodedMsgAndTokenDataSizes)
	for chain, msgs := range mo {
		sizes[chain] = make(map[cciptypes.SeqNum]int)
		for seq, msg := range msgs {
			td := tds[chain][seq]
			sizes[chain][seq] = EncodedSize(msg) + EncodedSize(td)
		}
	}
	return sizes
}

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
	Hashes   MessageHashes       `json:"messageHashes"`
	// TokenData are determined during the second phase of execute.
	// It contains the token data for the messages identified in the same stage as Messages
	TokenData TokenDataObservations `json:"tokenDataObservations"`

	// CostlyMessages are determined during the GetMessages state of execute.
	// It contains the message IDs of messages that cost more to execute than their source fees. These messages will not
	// be executed in the current round, but may be executed in future rounds (e.g. if gas prices decrease or if
	// these messages' fees are boosted high enough).
	CostlyMessages []cciptypes.Bytes32 `json:"costlyMessages"`

	// Nonces are determined during the third phase of execute.
	// It contains the nonces of senders who are being considered for the final report.
	Nonces NonceObservations `json:"nonces"`

	// Contracts are part of the initial discovery phase which runs to initialize the CCIP Reader.
	Contracts dt.Observation `json:"contracts"`
}

func (co CommitObservations) Flatten() []CommitData {
	var results []CommitData
	for _, reports := range co {
		results = append(results, reports...)
	}
	return results
}

// NewObservation constructs an Observation object.
func NewObservation(
	commitReports CommitObservations,
	messages MessageObservations,
	costlyMessages []cciptypes.Bytes32,
	tokenData TokenDataObservations,
	nonces NonceObservations,
	contracts dt.Observation,
	hashes MessageHashes,
) Observation {
	return Observation{
		CommitReports:  commitReports,
		Messages:       messages,
		CostlyMessages: costlyMessages,
		TokenData:      tokenData,
		Nonces:         nonces,
		Contracts:      contracts,
		Hashes:         hashes,
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
