package merkleroot

import (
	"sort"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
	"github.com/smartcontractkit/chainlink-ccip/shared"
)

type Query struct {
	RetryRMNSignatures bool
	RMNSignatures      *rmn.ReportSignatures
}

type Observation struct {
	MerkleRoots        []cciptypes.MerkleRootChain     `json:"merkleRoots"`
	OnRampMaxSeqNums   []plugintypes.SeqNumChain       `json:"onRampMaxSeqNums"`
	OffRampNextSeqNums []plugintypes.SeqNumChain       `json:"offRampNextSeqNums"`
	FChain             map[cciptypes.ChainSelector]int `json:"fChain"`
}

func (o Observation) IsEmpty() bool {
	return len(o.MerkleRoots) == 0 &&
		len(o.OnRampMaxSeqNums) == 0 &&
		len(o.OffRampNextSeqNums) == 0 &&
		len(o.FChain) == 0
}

// MerkleAggregatedObservation is the aggregation of a list of observations
type MerkleAggregatedObservation struct {
	// A map from chain selectors to the list of merkle roots observed for each chain
	MerkleRoots map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain

	// A map from chain selectors to the list of OnRamp max sequence numbers observed for each chain
	OnRampMaxSeqNums map[cciptypes.ChainSelector][]cciptypes.SeqNum

	// A map from chain selectors to the list of OffRamp next sequence numbers observed for each chain
	OffRampNextSeqNums map[cciptypes.ChainSelector][]cciptypes.SeqNum

	// A map from chain selectors to the list of f (failure tolerance) observed for each chain
	FChain map[cciptypes.ChainSelector][]int
}

// aggregateObservations takes a list of observations and produces an MerkleAggregatedObservation
func aggregateObservations(aos []shared.AttributedObservation[Observation]) MerkleAggregatedObservation {
	aggObs := MerkleAggregatedObservation{
		MerkleRoots:        make(map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain),
		OnRampMaxSeqNums:   make(map[cciptypes.ChainSelector][]cciptypes.SeqNum),
		OffRampNextSeqNums: make(map[cciptypes.ChainSelector][]cciptypes.SeqNum),
		FChain:             make(map[cciptypes.ChainSelector][]int),
	}

	for _, ao := range aos {
		obs := ao.Observation
		// MerkleRoots
		for _, merkleRoot := range obs.MerkleRoots {
			aggObs.MerkleRoots[merkleRoot.ChainSel] =
				append(aggObs.MerkleRoots[merkleRoot.ChainSel], merkleRoot)
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

	// A map from chain selectors to each chain's consensus OnRamp max sequence number
	OnRampMaxSeqNums map[cciptypes.ChainSelector]cciptypes.SeqNum

	// A map from chain selectors to each chain's consensus OffRamp next sequence number
	OffRampNextSeqNums map[cciptypes.ChainSelector]cciptypes.SeqNum

	// A map from chain selectors to each chain's consensus f (failure tolerance)
	FChain map[cciptypes.ChainSelector]int
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
	ReportTransmissionCheckAttempts uint                        `json:"reportTransmissionCheckAttempts"`
}

// Sort all fields of the given Outcome
func (o *Outcome) Sort() {
	sort.Slice(o.RangesSelectedForReport, func(i, j int) bool {
		return o.RangesSelectedForReport[i].ChainSel < o.RangesSelectedForReport[j].ChainSel
	})
	sort.Slice(o.RootsToReport, func(i, j int) bool {
		return o.RootsToReport[i].ChainSel < o.RootsToReport[j].ChainSel
	})
	sort.Slice(o.OffRampNextSeqNums, func(i, j int) bool {
		return o.OffRampNextSeqNums[i].ChainSel < o.OffRampNextSeqNums[j].ChainSel
	})
}

func (o *Outcome) NextState() State {
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

type State int

const (
	SelectingRangesForReport State = iota + 1
	BuildingReport
	WaitingForReportTransmission
)
