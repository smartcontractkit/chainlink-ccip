package committypes

import (
	"sort"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type MerkleRootQuery struct {
	RetryRMNSignatures bool
	RMNSignatures      *rmn.ReportSignatures
}

type MerkleRootObservation struct {
	MerkleRoots        []cciptypes.MerkleRootChain     `json:"merkleRoots"`
	OnRampMaxSeqNums   []plugintypes.SeqNumChain       `json:"onRampMaxSeqNums"`
	OffRampNextSeqNums []plugintypes.SeqNumChain       `json:"offRampNextSeqNums"`
	FChain             map[cciptypes.ChainSelector]int `json:"fChain"`
}

func (o MerkleRootObservation) IsEmpty() bool {
	return len(o.MerkleRoots) == 0 &&
		len(o.OnRampMaxSeqNums) == 0 &&
		len(o.OffRampNextSeqNums) == 0 &&
		len(o.FChain) == 0
}

type MerkleRootOutcomeType int

const (
	ReportIntervalsSelected MerkleRootOutcomeType = iota + 1
	ReportGenerated
	ReportEmpty
	ReportInFlight
	ReportTransmitted
	ReportTransmissionFailed
)

type MerkleRootOutcome struct {
	OutcomeType                     MerkleRootOutcomeType       `json:"outcomeType"`
	RangesSelectedForReport         []plugintypes.ChainRange    `json:"rangesSelectedForReport"`
	RootsToReport                   []cciptypes.MerkleRootChain `json:"rootsToReport"`
	OffRampNextSeqNums              []plugintypes.SeqNumChain   `json:"offRampNextSeqNums"`
	ReportTransmissionCheckAttempts uint                        `json:"reportTransmissionCheckAttempts"`
}

// Sort all fields of the given MerkleRootOutcome
func (o *MerkleRootOutcome) Sort() {
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

type MerkleRootState int

const (
	SelectingRangesForReport MerkleRootState = iota + 1
	BuildingReport
	WaitingForReportTransmission
)

func (o *MerkleRootOutcome) NextState() MerkleRootState {
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
