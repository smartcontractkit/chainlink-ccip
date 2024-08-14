package commitrmnocb

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type Query struct {
	RmnOnRampMaxSeqNums []plugintypes.SeqNumChain
	MerkleRoots         []cciptypes.MerkleRootChain
}

func (q Query) Encode() ([]byte, error) {
	return json.Marshal(q)
}

func DecodeCommitPluginQuery(encodedQuery []byte) (Query, error) {
	q := Query{}
	err := json.Unmarshal(encodedQuery, &q)
	return q, err
}

func NewCommitQuery(rmnOnRampMaxSeqNums []plugintypes.SeqNumChain, merkleRoots []cciptypes.MerkleRootChain) Query {
	return Query{
		RmnOnRampMaxSeqNums: rmnOnRampMaxSeqNums,
		MerkleRoots:         merkleRoots,
	}
}

type Observation struct {
	MerkleRoots        []cciptypes.MerkleRootChain     `json:"merkleRoots"`
	GasPrices          []cciptypes.GasPriceChain       `json:"gasPrices"`
	TokenPrices        []cciptypes.TokenPrice          `json:"tokenPrices"`
	OnRampMaxSeqNums   []plugintypes.SeqNumChain       `json:"onRampMaxSeqNums"`
	OffRampNextSeqNums []plugintypes.SeqNumChain       `json:"offRampNextSeqNums"`
	FChain             map[cciptypes.ChainSelector]int `json:"fChain"`
}

func (obs Observation) Encode() ([]byte, error) {
	encodedObservation, err := json.Marshal(obs)
	if err != nil {
		return nil, fmt.Errorf("failed to encode Observation: %w", err)
	}

	return encodedObservation, nil
}

func DecodeCommitPluginObservation(encodedObservation []byte) (Observation, error) {
	o := Observation{}
	err := json.Unmarshal(encodedObservation, &o)
	return o, err
}

// AggregatedObservation is the aggregation of a list of observations
type AggregatedObservation struct {
	// A map from chain selectors to the list of merkle roots observed for each chain
	MerkleRoots map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain

	// A map from chain selectors to the list of gas prices observed for each chain
	GasPrices map[cciptypes.ChainSelector][]cciptypes.BigInt

	// A map from token IDs to the list of prices observed for each token
	TokenPrices map[types.Account][]cciptypes.BigInt

	// A map from chain selectors to the list of OnRamp max sequence numbers observed for each chain
	OnRampMaxSeqNums map[cciptypes.ChainSelector][]cciptypes.SeqNum

	// A map from chain selectors to the list of OffRamp next sequence numbers observed for each chain
	OffRampNextSeqNums map[cciptypes.ChainSelector][]cciptypes.SeqNum

	// A map from chain selectors to the list of f (failure tolerance) observed for each chain
	FChain map[cciptypes.ChainSelector][]int
}

// aggregateObservations takes a list of observations and produces an AggregatedObservation
func aggregateObservations(aos []types.AttributedObservation) AggregatedObservation {
	aggObs := AggregatedObservation{
		MerkleRoots:        make(map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain),
		GasPrices:          make(map[cciptypes.ChainSelector][]cciptypes.BigInt),
		TokenPrices:        make(map[types.Account][]cciptypes.BigInt),
		OnRampMaxSeqNums:   make(map[cciptypes.ChainSelector][]cciptypes.SeqNum),
		OffRampNextSeqNums: make(map[cciptypes.ChainSelector][]cciptypes.SeqNum),
		FChain:             make(map[cciptypes.ChainSelector][]int),
	}

	for _, ao := range aos {
		obs, err := DecodeCommitPluginObservation(ao.Observation)
		if err != nil {
			// TODO: lggr
			continue
		}

		// MerkleRoots
		for _, merkleRoot := range obs.MerkleRoots {
			aggObs.MerkleRoots[merkleRoot.ChainSel] =
				append(aggObs.MerkleRoots[merkleRoot.ChainSel], merkleRoot)
		}

		// GasPrices
		for _, gasPriceChain := range obs.GasPrices {
			aggObs.GasPrices[gasPriceChain.ChainSel] =
				append(aggObs.GasPrices[gasPriceChain.ChainSel], gasPriceChain.GasPrice)
		}

		// TokenPrices
		for _, tokenPrice := range obs.TokenPrices {
			aggObs.TokenPrices[tokenPrice.TokenID] =
				append(aggObs.TokenPrices[tokenPrice.TokenID], tokenPrice.Price)
		}

		// OnRampMaxSeqNums
		for _, seqNumChain := range obs.OnRampMaxSeqNums {
			aggObs.OnRampMaxSeqNums[seqNumChain.ChainSel] =
				append(aggObs.OnRampMaxSeqNums[seqNumChain.ChainSel], seqNumChain.SeqNum)
		}

		// OffRampNextSeqNums
		for _, seqNumChain := range obs.OffRampNextSeqNums {
			aggObs.OffRampNextSeqNums[seqNumChain.ChainSel] =
				append(aggObs.OffRampNextSeqNums[seqNumChain.ChainSel], seqNumChain.SeqNum)
		}

		// FChain
		for chainSel, f := range obs.FChain {
			aggObs.FChain[chainSel] = append(aggObs.FChain[chainSel], f)
		}
	}

	return aggObs
}

// ConsensusObservation holds the consensus values for all chains across all observations in a round
type ConsensusObservation struct {
	// A map from chain selectors to each chain's consensus merkle root
	MerkleRoots map[cciptypes.ChainSelector]cciptypes.MerkleRootChain

	// A map from chain selectors to each chain's consensus gas prices
	GasPrices map[cciptypes.ChainSelector]cciptypes.BigInt

	// A map from token IDs to each token's consensus price
	TokenPrices map[types.Account]cciptypes.BigInt

	// A map from chain selectors to each chain's consensus OnRamp max sequence number
	OnRampMaxSeqNums map[cciptypes.ChainSelector]cciptypes.SeqNum

	// A map from chain selectors to each chain's consensus OffRamp next sequence number
	OffRampNextSeqNums map[cciptypes.ChainSelector]cciptypes.SeqNum

	// A map from chain selectors to each chain's consensus f (failure tolerance)
	FChain map[cciptypes.ChainSelector]int
}

// GasPricesArray returns a list of gas prices
func (co ConsensusObservation) GasPricesArray() []cciptypes.GasPriceChain {
	gasPrices := make([]cciptypes.GasPriceChain, 0, len(co.GasPrices))
	for chain, gasPrice := range co.GasPrices {
		gasPrices = append(gasPrices, cciptypes.NewGasPriceChain(gasPrice.Int, chain))
	}

	return gasPrices
}

// TokenPricesArray returns a list of token prices
func (co ConsensusObservation) TokenPricesArray() []cciptypes.TokenPrice {
	tokenPrices := make([]cciptypes.TokenPrice, 0, len(co.TokenPrices))
	for tokenID, tokenPrice := range co.TokenPrices {
		tokenPrices = append(tokenPrices, cciptypes.NewTokenPrice(tokenID, tokenPrice.Int))
	}

	return tokenPrices
}

type OutcomeType int

const (
	ReportIntervalsSelected OutcomeType = iota + 1
	ReportGenerated
	ReportEmpty
	ReportInFlight
	ReportTransmitted
	ReportTransmissionFailed
)

type Outcome struct {
	OutcomeType                     OutcomeType                 `json:"outcomeType"`
	RangesSelectedForReport         []plugintypes.ChainRange    `json:"rangesSelectedForReport"`
	RootsToReport                   []cciptypes.MerkleRootChain `json:"rootsToReport"`
	OffRampNextSeqNums              []plugintypes.SeqNumChain   `json:"offRampNextSeqNums"`
	TokenPrices                     []cciptypes.TokenPrice      `json:"tokenPrices"`
	GasPrices                       []cciptypes.GasPriceChain   `json:"gasPrices"`
	ReportTransmissionCheckAttempts uint                        `json:"reportTransmissionCheckAttempts"`
}

// Sort all fields of the given Outcome
func (o Outcome) sort() {
	sort.Slice(o.RangesSelectedForReport, func(i, j int) bool {
		return o.RangesSelectedForReport[i].ChainSel < o.RangesSelectedForReport[j].ChainSel
	})
	sort.Slice(o.RootsToReport, func(i, j int) bool {
		return o.RootsToReport[i].ChainSel < o.RootsToReport[j].ChainSel
	})
	sort.Slice(o.OffRampNextSeqNums, func(i, j int) bool {
		return o.OffRampNextSeqNums[i].ChainSel < o.OffRampNextSeqNums[j].ChainSel
	})
	sort.Slice(o.TokenPrices, func(i, j int) bool {
		return o.TokenPrices[i].TokenID < o.TokenPrices[j].TokenID
	})
	sort.Slice(o.GasPrices, func(i, j int) bool {
		return o.GasPrices[i].ChainSel < o.GasPrices[j].ChainSel
	})
}

// Encode encodes an Outcome deterministically
func (o Outcome) Encode() ([]byte, error) {

	// Sort all lists to ensure deterministic serialization
	o.sort()

	encodedOutcome, err := json.Marshal(o)
	if err != nil {
		return nil, fmt.Errorf("failed to encode Outcome: %w", err)
	}

	return encodedOutcome, nil
}

func DecodeCommitPluginOutcome(b []byte) (Outcome, error) {
	o := Outcome{}
	err := json.Unmarshal(b, &o)
	return o, err
}

func (o Outcome) NextState() CommitPluginState {
	switch o.OutcomeType {
	case ReportIntervalsSelected:
		return BuildingReport
	case ReportGenerated:
		return WaitingForReportTransmission
	case ReportEmpty:
		return SelectingRangesForReport
	case ReportInFlight:
		return WaitingForReportTransmission
	case ReportTransmitted:
		return SelectingRangesForReport
	case ReportTransmissionFailed:
		return SelectingRangesForReport
	default:
		return SelectingRangesForReport
	}
}

type CommitPluginState int

const (
	SelectingRangesForReport CommitPluginState = iota + 1
	BuildingReport
	WaitingForReportTransmission
)
