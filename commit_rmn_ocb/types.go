package commitrmnocb

import (
	"encoding/json"
	"fmt"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type CommitPluginConfig struct {
	// TODO: doc
	AllSourceChains []cciptypes.ChainSelector

	// DestChain is the ccip destination chain configured for the commit plugin DON.
	DestChain cciptypes.ChainSelector `json:"destChain"`

	// PricedTokens is a list of tokens that we want to submit price updates for.
	PricedTokens []types.Account `json:"pricedTokens"`

	// TokenPricesObserver indicates that the node can observe token prices.
	TokenPricesObserver bool `json:"tokenPricesObserver"`

	// NewMsgScanBatchSize is the number of max new messages to scan, typically set to 256.
	NewMsgScanBatchSize int `json:"newMsgScanBatchSize"`
}

type RmnSig struct {
	sig []byte
}

type SignedMerkleRoot struct {
	ChainSel    cciptypes.ChainSelector `json:"chain"`
	SeqNumRange cciptypes.SeqNumRange   `json:"seqNumRange"`
	MerkleRoot  cciptypes.Bytes32       `json:"merkleRoot"`
	RmnSigs     []RmnSig                `json:"rmnSigs"`
}

type CommitQuery struct {
	RmnOnRampMaxSeqNums []plugintypes.SeqNumChain
	SignedMerkleRoots   []SignedMerkleRoot
}

func (q CommitQuery) Encode() ([]byte, error) {
	return json.Marshal(q)
}

func DecodeCommitPluginQuery(encodedQuery []byte) (CommitQuery, error) {
	q := CommitQuery{}
	err := json.Unmarshal(encodedQuery, &q)
	return q, err
}

func NewCommitQuery(rmnOnRampMaxSeqNums []plugintypes.SeqNumChain, signedMerkleRoots []SignedMerkleRoot) CommitQuery {
	return CommitQuery{
		RmnOnRampMaxSeqNums: rmnOnRampMaxSeqNums,
		SignedMerkleRoots:   signedMerkleRoots,
	}
}

type MerkleRootAndChain struct {
	MerkleRoot cciptypes.Bytes32       `json:"merkleRoot"`
	ChainSel   cciptypes.ChainSelector `json:"chain"`
}

type CommitPluginObservation struct {
	MerkleRoots       []MerkleRootAndChain            `json:"merkleRoots"`
	GasPrices         []cciptypes.GasPriceChain       `json:"gasPrices"`
	TokenPrices       []cciptypes.TokenPrice          `json:"tokenPrices"`
	OnRampMaxSeqNums  []plugintypes.SeqNumChain       `json:"onRampMaxSeqNums"`
	OffRampMaxSeqNums []plugintypes.SeqNumChain       `json:"offRampMaxSeqNums"`
	FChain            map[cciptypes.ChainSelector]int `json:"fChain"`
}

func (obs CommitPluginObservation) Encode() ([]byte, error) {
	encodedObservation, err := json.Marshal(obs)
	if err != nil {
		return nil, fmt.Errorf("failed to encode CommitPluginObservation: %w", err)
	}

	return encodedObservation, nil
}

func DecodeCommitPluginObservation(encodedObservation []byte) (CommitPluginObservation, error) {
	o := CommitPluginObservation{}
	err := json.Unmarshal(encodedObservation, &o)
	return o, err
}

// AggregatedObservation TODO: doc
type AggregatedObservation struct {
	// A map from chain selectors to the list of merkle roots observed for each chain
	MerkleRoots map[cciptypes.ChainSelector][]cciptypes.Bytes32

	// A map from chain selectors to the list of gas prices observed for each chain
	GasPrices map[cciptypes.ChainSelector][]cciptypes.BigInt

	// A map from token IDs to the list of prices observed for each token
	TokenPrices map[types.Account][]cciptypes.BigInt

	// A map from chain selectors to the list of OnRamp max sequence numbers observed for each chain
	OnRampMaxSeqNums map[cciptypes.ChainSelector][]cciptypes.SeqNum

	// A map from chain selectors to the list of OffRamp max sequence numbers observed for each chain
	OffRampMaxSeqNums map[cciptypes.ChainSelector][]cciptypes.SeqNum

	// A map from chain selectors to the list of f (failure tolerance) observed for each chain
	FChain map[cciptypes.ChainSelector][]int
}

// aggregateObservations TODO: doc
func aggregateObservations(aos []types.AttributedObservation) AggregatedObservation {
	aggObs := AggregatedObservation{}

	for _, ao := range aos {
		obs, err := DecodeCommitPluginObservation(ao.Observation)
		if err != nil {
			// TODO: log
			continue
		}

		// MerkleRoots
		for _, rootAndChain := range obs.MerkleRoots {
			AppendToMap(aggObs.MerkleRoots, rootAndChain.ChainSel, rootAndChain.MerkleRoot)
		}

		// GasPrices
		for _, gasPriceChain := range obs.GasPrices {
			AppendToMap(aggObs.GasPrices, gasPriceChain.ChainSel, gasPriceChain.GasPrice)
		}

		// TokenPrices
		for _, tokenPrice := range obs.TokenPrices {
			AppendToMap(aggObs.TokenPrices, tokenPrice.TokenID, tokenPrice.Price)
		}

		// OnRampMaxSeqNums
		for _, seqNumChain := range obs.OnRampMaxSeqNums {
			AppendToMap(aggObs.OnRampMaxSeqNums, seqNumChain.ChainSel, seqNumChain.SeqNum)
		}

		// OffRampMaxSeqNums
		for _, seqNumChain := range obs.OffRampMaxSeqNums {
			AppendToMap(aggObs.OffRampMaxSeqNums, seqNumChain.ChainSel, seqNumChain.SeqNum)
		}

		// FChain
		for chainSel, f := range obs.FChain {
			AppendToMap(aggObs.FChain, chainSel, f)
		}
	}

	return aggObs
}

// AppendToMap TODO: doc
func AppendToMap[K comparable, V any](m map[K][]V, k K, v V) {
	if _, exists := m[k]; exists {
		m[k] = append(m[k], v)
	} else {
		m[k] = []V{v}
	}
}

// ConsensusObservation TODO: doc
type ConsensusObservation struct {
	// A map from chain selectors to each chain's consensus merkle root
	MerkleRoots map[cciptypes.ChainSelector]cciptypes.Bytes32

	// A map from chain selectors to each chain's consensus gas prices
	GasPrices map[cciptypes.ChainSelector]cciptypes.BigInt

	// A map from token IDs to each token's consensus price
	TokenPrices map[types.Account]cciptypes.BigInt

	// A map from chain selectors to each chain's consensus OnRamp max sequence number
	OnRampMaxSeqNums map[cciptypes.ChainSelector]cciptypes.SeqNum

	// A map from chain selectors to each chain's consensus OffRamp max sequence number
	OffRampMaxSeqNums map[cciptypes.ChainSelector]cciptypes.SeqNum

	// A map from chain selectors to each chain's consensus f (failure tolerance)
	FChain map[cciptypes.ChainSelector]int
}

type CommitPluginOutcomeType int

const (
	ReportIntervalsSelected CommitPluginOutcomeType = iota
	ReportGenerated
	ReportEmpty
	ReportNotYetTransmitted
	ReportTransmitted
	ReportNotTransmitted
)

type CommitPluginOutcome struct {
	OutcomeType             CommitPluginOutcomeType
	RangesSelectedForReport []ChainRange
	SignedRootsToReport     []SignedMerkleRoot
	MaxOffRampSeqNums       []plugintypes.SeqNumChain
	TokenPrices             []cciptypes.TokenPrice    `json:"tokenPrices"`
	GasPrices               []cciptypes.GasPriceChain `json:"gasPrices"`
}

func (o *CommitPluginOutcome) Encode() ([]byte, error) {
	encodedOutcome, err := json.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("failed to encode CommitPluginOutcome: %w", err)
	}

	return encodedOutcome, nil
}

func DecodeCommitPluginOutcome(b []byte) (CommitPluginOutcome, error) {
	o := CommitPluginOutcome{}
	err := json.Unmarshal(b, &o)
	return o, err
}

// NextState TODO: doc
func (o *CommitPluginOutcome) NextState() CommitPluginState {
	switch o.OutcomeType {
	case ReportIntervalsSelected:
		return BuildingReport
	case ReportGenerated:
		return WaitingForReportTransmission
	case ReportEmpty:
		return SelectingRangesForReport
	case ReportNotYetTransmitted:
		return WaitingForReportTransmission
	case ReportTransmitted:
		return SelectingRangesForReport
	case ReportNotTransmitted:
		return SelectingRangesForReport
	default:
		return SelectingRangesForReport
	}
}

type CommitPluginState int

const (
	SelectingRangesForReport CommitPluginState = iota
	BuildingReport
	WaitingForReportTransmission
)

// ChainRange TODO: doc
type ChainRange struct {
	ChainSel    cciptypes.ChainSelector `json:"chain"`
	SeqNumRange cciptypes.SeqNumRange   `json:"seqNumRange"`
}

// Rmn TODO: doc
type Rmn interface {
	RequestMaxSeqNums(chains []cciptypes.ChainSelector) ([]plugintypes.SeqNumChain, error)
	RequestSignedIntervals(chainRanges []ChainRange) ([]SignedMerkleRoot, error)
}

type OnChain interface {
	GetOnRampMaxSeqNums() ([]plugintypes.SeqNumChain, error)
	GetOffRampMaxSeqNums() ([]plugintypes.SeqNumChain, error)
	GetMerkleRoots([]ChainRange) ([]MerkleRootAndChain, error)
}
