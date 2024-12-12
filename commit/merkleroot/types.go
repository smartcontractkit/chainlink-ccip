package merkleroot

import (
	"sort"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type Query struct {
	RetryRMNSignatures bool
	RMNSignatures      *rmn.ReportSignatures
}

type Observation struct {
	MerkleRoots        []cciptypes.MerkleRootChain     `json:"merkleRoots"`
	OnRampMaxSeqNums   []plugintypes.SeqNumChain       `json:"onRampMaxSeqNums"`
	OffRampNextSeqNums []plugintypes.SeqNumChain       `json:"offRampNextSeqNums"`
	RMNRemoteConfig    rmntypes.RemoteConfig           `json:"rmnRemoteConfig"`
	FChain             map[cciptypes.ChainSelector]int `json:"fChain"`
}

func (o Observation) IsEmpty() bool {
	return len(o.MerkleRoots) == 0 &&
		len(o.OnRampMaxSeqNums) == 0 &&
		len(o.OffRampNextSeqNums) == 0 &&
		o.RMNRemoteConfig.IsEmpty() &&
		len(o.FChain) == 0
}

// aggregatedObservation is the aggregation of a list of merkle root processor observations
type aggregatedObservation struct {
	// A map from chain selectors to the list of merkle roots observed for each chain
	MerkleRoots map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain

	// A map from chain selectors to the list of OnRamp max sequence numbers observed for each chain
	OnRampMaxSeqNums map[cciptypes.ChainSelector][]cciptypes.SeqNum

	// A map from chain selectors to the list of OffRamp next sequence numbers observed for each chain
	OffRampNextSeqNums map[cciptypes.ChainSelector][]cciptypes.SeqNum

	// The RMNRemoteConfig observed
	RMNRemoteConfigs []rmntypes.RemoteConfig

	// A map from chain selectors to the list of f (failure tolerance) observed for each chain
	FChain map[cciptypes.ChainSelector][]int
}

// aggregateObservations takes a list of observations and produces an MerkleAggregatedObservation
func aggregateObservations(aos []plugincommon.AttributedObservation[Observation]) aggregatedObservation {
	aggObs := aggregatedObservation{
		MerkleRoots:        make(map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain),
		OnRampMaxSeqNums:   make(map[cciptypes.ChainSelector][]cciptypes.SeqNum),
		OffRampNextSeqNums: make(map[cciptypes.ChainSelector][]cciptypes.SeqNum),
		RMNRemoteConfigs:   make([]rmntypes.RemoteConfig, 0),
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

		// RMNRemoteConfig
		if !obs.RMNRemoteConfig.IsEmpty() {
			aggObs.RMNRemoteConfigs = append(aggObs.RMNRemoteConfigs, obs.RMNRemoteConfig)
		}

		// FChain
		for chainSel, f := range obs.FChain {
			aggObs.FChain[chainSel] = append(aggObs.FChain[chainSel], f)
		}

	}

	return aggObs
}

// consensusObservation holds the consensus values for all chains across all observations in a round
type consensusObservation struct {
	// A map from chain selectors to each chain's consensus merkle root
	MerkleRoots map[cciptypes.ChainSelector]cciptypes.MerkleRootChain

	// A map from chain selectors to each chain's consensus OnRamp max sequence number
	OnRampMaxSeqNums map[cciptypes.ChainSelector]cciptypes.SeqNum

	// A map from chain selectors to each chain's consensus OffRamp next sequence number
	OffRampNextSeqNums map[cciptypes.ChainSelector]cciptypes.SeqNum

	// The consensus RMNRemoteConfig
	RMNRemoteConfig map[cciptypes.ChainSelector]rmntypes.RemoteConfig

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
	OutcomeType                     OutcomeType                   `json:"outcomeType"`
	RangesSelectedForReport         []plugintypes.ChainRange      `json:"rangesSelectedForReport"`
	RootsToReport                   []cciptypes.MerkleRootChain   `json:"rootsToReport"`
	OffRampNextSeqNums              []plugintypes.SeqNumChain     `json:"offRampNextSeqNums"`
	ReportTransmissionCheckAttempts uint                          `json:"reportTransmissionCheckAttempts"`
	RMNReportSignatures             []cciptypes.RMNECDSASignature `json:"rmnReportSignatures"`
	RMNRemoteCfg                    rmntypes.RemoteConfig         `json:"rmnRemoteCfg"`
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

func (o *Outcome) nextState() processorState {
	switch o.OutcomeType {
	case ReportIntervalsSelected:
		return buildingReport
	case ReportGenerated:
		return waitingForReportTransmission
	case ReportEmpty:
		return selectingRangesForReport
	case ReportInFlight:
		return waitingForReportTransmission
	case ReportTransmitted:
		return selectingRangesForReport
	case ReportTransmissionFailed:
		return selectingRangesForReport
	default:
		return selectingRangesForReport
	}
}

type processorState int

const (
	// selectingRangesForReport is the initial state of the processor while oracles agree on which message
	// sequence number ranges to include in the report.
	selectingRangesForReport processorState = iota + 1

	// buildingReport is the state of the processor while the report is being built after message
	// sequence numbers are agreed.
	buildingReport

	// waitingForReportTransmission is the state of the processor after the report is built and waiting for it to
	// get transmitted onchain.
	waitingForReportTransmission
)
