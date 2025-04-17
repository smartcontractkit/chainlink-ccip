package merkleroot

import (
	"sort"
	"time"

	"github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	processorLabel    = "merkleroot"
	rootsLabel        = "roots"
	messagesLabel     = "messages"
	rmnSignatureLabel = "rmnSignatures"
)

type Query struct {
	RetryRMNSignatures bool
	RMNSignatures      *rmn.ReportSignatures
}

// ContainsRmnSignatures returns true if the query contains RMN signatures.
func (q Query) ContainsRmnSignatures() bool {
	return q.RMNSignatures != nil && len(q.RMNSignatures.Signatures) > 0
}

type Observation struct {
	MerkleRoots        []cciptypes.MerkleRootChain      `json:"merkleRoots"`
	RMNEnabledChains   map[cciptypes.ChainSelector]bool `json:"rmnEnabledChains"`
	OnRampMaxSeqNums   []plugintypes.SeqNumChain        `json:"onRampMaxSeqNums"`
	OffRampNextSeqNums []plugintypes.SeqNumChain        `json:"offRampNextSeqNums"`
	RMNRemoteConfig    cciptypes.RemoteConfig           `json:"rmnRemoteConfig"`
	FChain             map[cciptypes.ChainSelector]int  `json:"fChain"`
}

func (o Observation) Stats() map[string]int {
	counts := map[string]int{
		rootsLabel:    len(o.MerkleRoots),
		messagesLabel: 0,
	}
	for _, root := range o.MerkleRoots {
		counts[messagesLabel] += root.SeqNumsRange.Length()
	}
	return counts
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

	// RMNEnabledChains is a map of the RMN-enabled source chains.
	RMNEnabledChains map[cciptypes.ChainSelector][]bool

	// A map from chain selectors to the list of OnRamp max sequence numbers observed for each chain
	OnRampMaxSeqNums map[cciptypes.ChainSelector][]cciptypes.SeqNum

	// A map from chain selectors to the list of OffRamp next sequence numbers observed for each chain
	OffRampNextSeqNums map[cciptypes.ChainSelector][]cciptypes.SeqNum

	// The RMNRemoteConfig observed
	RMNRemoteConfigs []cciptypes.RemoteConfig

	// A map from chain selectors to the list of f (failure tolerance) observed for each chain
	FChain map[cciptypes.ChainSelector][]int
}

// aggregateObservations takes a list of observations and produces an MerkleAggregatedObservation
func aggregateObservations(aos []plugincommon.AttributedObservation[Observation]) aggregatedObservation {
	aggObs := aggregatedObservation{
		MerkleRoots:        make(map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain),
		RMNEnabledChains:   make(map[cciptypes.ChainSelector][]bool),
		OnRampMaxSeqNums:   make(map[cciptypes.ChainSelector][]cciptypes.SeqNum),
		OffRampNextSeqNums: make(map[cciptypes.ChainSelector][]cciptypes.SeqNum),
		RMNRemoteConfigs:   make([]cciptypes.RemoteConfig, 0),
		FChain:             make(map[cciptypes.ChainSelector][]int),
	}

	for _, ao := range aos {
		obs := ao.Observation
		// MerkleRoots
		for _, merkleRoot := range obs.MerkleRoots {
			aggObs.MerkleRoots[merkleRoot.ChainSel] =
				append(aggObs.MerkleRoots[merkleRoot.ChainSel], merkleRoot)
		}

		// RMNEnabledChains
		for chainSel, enabled := range obs.RMNEnabledChains {
			aggObs.RMNEnabledChains[chainSel] = append(aggObs.RMNEnabledChains[chainSel], enabled)
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

	// RMNEnabledChains holds the consensus of RMNEnabledChains
	RMNEnabledChains map[cciptypes.ChainSelector]bool

	// A map from chain selectors to each chain's consensus OnRamp max sequence number
	OnRampMaxSeqNums map[cciptypes.ChainSelector]cciptypes.SeqNum

	// A map from chain selectors to each chain's consensus OffRamp next sequence number
	OffRampNextSeqNums map[cciptypes.ChainSelector]cciptypes.SeqNum

	// The consensus RMNRemoteConfig
	RMNRemoteConfig map[cciptypes.ChainSelector]cciptypes.RemoteConfig

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
	OutcomeType                     OutcomeType                      `json:"outcomeType"`
	RangesSelectedForReport         []plugintypes.ChainRange         `json:"rangesSelectedForReport"`
	RootsToReport                   []cciptypes.MerkleRootChain      `json:"rootsToReport"`
	RMNEnabledChains                map[cciptypes.ChainSelector]bool `json:"rmnEnabledChains"`
	OffRampNextSeqNums              []plugintypes.SeqNumChain        `json:"offRampNextSeqNums"`
	ReportTransmissionCheckAttempts uint                             `json:"reportTransmissionCheckAttempts"`
	RMNReportSignatures             []cciptypes.RMNECDSASignature    `json:"rmnReportSignatures"`
	RMNRemoteCfg                    cciptypes.RemoteConfig           `json:"rmnRemoteCfg"`
}

func (o Outcome) Stats() map[string]int {
	counts := map[string]int{
		rootsLabel:        len(o.RootsToReport),
		rmnSignatureLabel: len(o.RMNReportSignatures),
		messagesLabel:     0,
	}
	for _, root := range o.RootsToReport {
		counts[messagesLabel] += root.SeqNumsRange.Length()
	}
	return counts
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

func (p processorState) String() string {
	switch p {
	case selectingRangesForReport:
		return "selectingRangesForReport"
	case buildingReport:
		return "buildingReport"
	case waitingForReportTransmission:
		return "waitingForReportTransmission"
	default:
		return "unknown"
	}
}

// MetricsReporter exposes only relevant methods for reporting merkle roots from metrics.Reporter
type MetricsReporter interface {
	TrackRmnReport(latency float64, success bool)
	TrackProcessorLatency(processor string, method string, latency time.Duration, err error)
	TrackProcessorOutput(processor string, method plugincommon.MethodType, obs plugintypes.Trackable)
}

type NoopMetrics struct{}

func (n NoopMetrics) TrackRmnReport(float64, bool) {}

func (n NoopMetrics) TrackProcessorLatency(string, string, time.Duration, error) {}

func (n NoopMetrics) TrackProcessorOutput(string, plugincommon.MethodType, plugintypes.Trackable) {}
