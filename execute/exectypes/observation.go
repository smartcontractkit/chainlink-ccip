package exectypes

import (
	"context"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/execute/internal"
	dt "github.com/smartcontractkit/chainlink-ccip/internal/plugincommon/discovery/discoverytypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	sourceChainsLabel  = "sourceChains"
	messagesLabel      = "messages"
	tokenDataLabel     = "tokenData"
	commitReportsLabel = "commitReports"
	noncesLabel        = "nonces"
	tokenStateReady    = "tokenReady"
	tokenStateWaiting  = "tokenWaiting"
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

// Count the number of messages in the observation.
func (mo MessageObservations) Count() int {
	count := 0
	for _, msgs := range mo {
		count += len(msgs)
	}
	return count
}

func (mo MessageObservations) Stats() map[string]int {
	messagesCount := 0
	for _, chainMessages := range mo {
		messagesCount += len(chainMessages)
	}

	return map[string]int{
		messagesLabel: messagesCount,
	}
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
			sizes[chain][seq] = internal.EncodedSize(msg) + internal.EncodedSize(td)
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

	// Nonces are determined during the third phase of execute.
	// It contains the nonces of senders who are being considered for the final report.
	Nonces NonceObservations `json:"nonces"`

	// Contracts are part of the initial discovery phase which runs to initialize the CCIP Reader.
	Contracts dt.Observation `json:"contracts"`

	FChain map[cciptypes.ChainSelector]int `json:"fChain"`
}

// ToLogFormat creates a copy of the outcome with the messages.data and discovery obs removed
func (o Observation) ToLogFormat() Observation {
	msgsWithEmptyData := make(MessageObservations)
	for srcChain, msgs := range o.Messages {
		msgsWithEmptyData[srcChain] = make(map[cciptypes.SeqNum]cciptypes.Message)
		for seqNum, msg := range msgs {
			msgsWithEmptyData[srcChain][seqNum] = msg.CopyWithoutData()
		}
	}
	cleanedObs := Observation{
		CommitReports: o.CommitReports,
		Hashes:        o.Hashes,
		TokenData:     o.TokenData,
		Nonces:        o.Nonces,
		FChain:        o.FChain,
		Messages:      msgsWithEmptyData,
		Contracts:     dt.Observation{},
	}

	return cleanedObs
}

func (co CommitObservations) Flatten() []CommitData {
	var results []CommitData
	for _, reports := range co {
		results = append(results, reports...)
	}
	return results
}

func (co CommitObservations) Stats() map[string]int {
	commitReportsCount := 0
	for _, commit := range co {
		commitReportsCount += len(commit)
	}

	return map[string]int{
		commitReportsLabel: commitReportsCount,
	}
}

// NewObservation constructs an Observation object.
func NewObservation(
	commitReports CommitObservations,
	messages MessageObservations,
	tokenData TokenDataObservations,
	nonces NonceObservations,
	contracts dt.Observation,
	hashes MessageHashes,
) Observation {
	return Observation{
		CommitReports: commitReports,
		Messages:      messages,
		TokenData:     tokenData,
		Nonces:        nonces,
		Contracts:     contracts,
		Hashes:        hashes,
	}
}

func (o Observation) Stats() map[string]int {
	stats := map[string]int{}
	mergeStats(&stats, o.Nonces)
	mergeStats(&stats, o.CommitReports)
	mergeStats(&stats, o.Messages)
	mergeStats(&stats, o.TokenData)
	return stats
}

func mergeStats(dest *map[string]int, t plugintypes.Trackable) {
	if t == nil {
		return
	}
	maps.Copy(*dest, t.Stats())
}

func (o NonceObservations) Stats() map[string]int {
	noncesCount := 0
	for _, chainNonces := range o {
		noncesCount += len(chainNonces)
	}

	return map[string]int{
		noncesLabel: noncesCount,
	}
}

func (o TokenDataObservations) Stats() map[string]int {
	tokenCounters := map[string]int{
		tokenStateReady:   0,
		tokenStateWaiting: 0,
	}

	for _, chainTokens := range o {
		for _, tokenData := range chainTokens {
			for _, token := range tokenData.TokenData {
				counterKey := tokenStateWaiting
				if token.IsReady() {
					counterKey = tokenStateReady
				}
				tokenCounters[counterKey]++
			}
		}
	}

	return tokenCounters
}
