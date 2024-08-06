package commitrmnocb

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type RmnSig struct {
	sig []byte
}

type CommitQuery struct {
	RmnOnRampMaxSeqNums []plugintypes.SeqNumChain
	MerkleRoots         []cciptypes.MerkleRootChain
}

func (q CommitQuery) Encode() ([]byte, error) {
	return json.Marshal(q)
}

func DecodeCommitPluginQuery(encodedQuery []byte) (CommitQuery, error) {
	q := CommitQuery{}
	err := json.Unmarshal(encodedQuery, &q)
	return q, err
}

func NewCommitQuery(rmnOnRampMaxSeqNums []plugintypes.SeqNumChain, merkleRoots []cciptypes.MerkleRootChain) CommitQuery {
	return CommitQuery{
		RmnOnRampMaxSeqNums: rmnOnRampMaxSeqNums,
		MerkleRoots:         merkleRoots,
	}
}

type CommitPluginObservation struct {
	MerkleRoots        []cciptypes.MerkleRootChain     `json:"merkleRoots"`
	GasPrices          []cciptypes.GasPriceChain       `json:"gasPrices"`
	TokenPrices        []cciptypes.TokenPrice          `json:"tokenPrices"`
	OnRampMaxSeqNums   []plugintypes.SeqNumChain       `json:"onRampMaxSeqNums"`
	OffRampNextSeqNums []plugintypes.SeqNumChain       `json:"offRampNextSeqNums"`
	FChain             map[cciptypes.ChainSelector]int `json:"fChain"`
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
	aggObs := AggregatedObservation{}

	for _, ao := range aos {
		obs, err := DecodeCommitPluginObservation(ao.Observation)
		if err != nil {
			// TODO: log
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

// GasPricesSortedArray returns a sorted list of gas prices
func (co ConsensusObservation) GasPricesSortedArray() []cciptypes.GasPriceChain {
	gasPrices := make([]cciptypes.GasPriceChain, 0, len(co.GasPrices))
	for chain, gasPrice := range co.GasPrices {
		gasPrices = append(gasPrices, cciptypes.NewGasPriceChain(gasPrice.Int, chain))
	}

	sort.Slice(gasPrices, func(i, j int) bool {
		return gasPrices[i].ChainSel < gasPrices[j].ChainSel
	})

	return gasPrices
}

// TokenPricesSortedArray returns a sorted list of token prices
func (co ConsensusObservation) TokenPricesSortedArray() []cciptypes.TokenPrice {
	tokenPrices := make([]cciptypes.TokenPrice, 0, len(co.TokenPrices))
	for tokenID, tokenPrice := range co.TokenPrices {
		tokenPrices = append(tokenPrices, cciptypes.NewTokenPrice(tokenID, tokenPrice.Int))
	}

	sort.Slice(tokenPrices, func(i, j int) bool {
		return tokenPrices[i].TokenID < tokenPrices[j].TokenID
	})

	return tokenPrices
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
	OutcomeType                     CommitPluginOutcomeType
	RangesSelectedForReport         []ChainRange
	RootsToReport                   []cciptypes.MerkleRootChain
	OffRampNextSeqNums              []plugintypes.SeqNumChain
	TokenPrices                     []cciptypes.TokenPrice    `json:"tokenPrices"`
	GasPrices                       []cciptypes.GasPriceChain `json:"gasPrices"`
	ReportTransmissionCheckAttempts uint                      `json:"reportTransmissionCheckAttempts"`
}

// Encode TODO: sort all lists here to ensure deterministic serialization
func (o CommitPluginOutcome) Encode() ([]byte, error) {
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

func (o CommitPluginOutcome) NextState() CommitPluginState {
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

type ChainRange struct {
	ChainSel    cciptypes.ChainSelector `json:"chain"`
	SeqNumRange cciptypes.SeqNumRange   `json:"seqNumRange"`
}

// CommitPluginReport is the report that will be transmitted by the Commit Plugin
type CommitPluginReport struct {
	MerkleRoots []cciptypes.MerkleRootChain
	TokenPrices []cciptypes.TokenPrice    `json:"tokenPrices"`
	GasPrices   []cciptypes.GasPriceChain `json:"gasPrices"`
}

func (r CommitPluginReport) IsEmpty() bool {
	return len(r.MerkleRoots) == 0 && len(r.TokenPrices) == 0 && len(r.GasPrices) == 0
}

func (r CommitPluginReport) Encode() ([]byte, error) {
	encodedReport, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("failed to encode CommitPluginReport: %w", err)
	}

	return encodedReport, nil
}

func DecodeCommitPluginReport(b []byte) (CommitPluginReport, error) {
	r := CommitPluginReport{}
	err := json.Unmarshal(b, &r)
	return r, err
}
