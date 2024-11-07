package merkleroot

import (
	"testing"

	"github.com/stretchr/testify/assert"

	rmntypes "github.com/smartcontractkit/chainlink-ccip/commit/merkleroot/rmn/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugintypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func TestNextState(t *testing.T) {
	tests := []struct {
		name        string
		outcomeType OutcomeType
		expected    State
	}{
		{
			name:        "ReportIntervalsSelected -> BuildingReport",
			outcomeType: ReportIntervalsSelected,
			expected:    BuildingReport,
		},
		{
			name:        "ReportGenerated -> WaitingForReportTransmission",
			outcomeType: ReportGenerated,
			expected:    WaitingForReportTransmission,
		},
		{
			name:        "ReportEmpty -> SelectingRangesForReport",
			outcomeType: ReportEmpty,
			expected:    SelectingRangesForReport,
		},
		{
			name:        "ReportInFlight -> WaitingForReportTransmission",
			outcomeType: ReportInFlight,
			expected:    WaitingForReportTransmission,
		},
		{
			name:        "ReportTransmitted -> SelectingRangesForReport",
			outcomeType: ReportTransmitted,
			expected:    SelectingRangesForReport,
		},
		{
			name:        "ReportTransmissionFailed -> SelectingRangesForReport",
			outcomeType: ReportTransmissionFailed,
			expected:    SelectingRangesForReport,
		},
		{
			name:        "Unknown -> SelectingRangesForReport",
			outcomeType: OutcomeType(999), // An invalid outcome type to test the default case
			expected:    SelectingRangesForReport,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			outcome := &Outcome{OutcomeType: tt.outcomeType}
			assert.Equal(t, tt.expected, outcome.NextState())
		})
	}
}

func TestObservation_IsEmpty(t *testing.T) {
	tests := []struct {
		name        string
		observation Observation
		expected    bool
	}{
		{
			name: "Empty Observation",
			observation: Observation{
				MerkleRoots:        []cciptypes.MerkleRootChain{},
				OnRampMaxSeqNums:   []plugintypes.SeqNumChain{},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{},
				RMNRemoteConfig:    rmntypes.RemoteConfig{},
				FChain:             map[cciptypes.ChainSelector]int{},
			},
			expected: true,
		},
		{
			name: "Non-empty MerkleRoots",
			observation: Observation{
				MerkleRoots:        []cciptypes.MerkleRootChain{{}},
				OnRampMaxSeqNums:   []plugintypes.SeqNumChain{},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{},
				RMNRemoteConfig:    rmntypes.RemoteConfig{},
				FChain:             map[cciptypes.ChainSelector]int{},
			},
			expected: false,
		},
		{
			name: "Non-empty OnRampMaxSeqNums",
			observation: Observation{
				MerkleRoots:        []cciptypes.MerkleRootChain{},
				OnRampMaxSeqNums:   []plugintypes.SeqNumChain{{}},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{},
				RMNRemoteConfig:    rmntypes.RemoteConfig{},
				FChain:             map[cciptypes.ChainSelector]int{},
			},
			expected: false,
		},
		{
			name: "Non-empty OffRampNextSeqNums",
			observation: Observation{
				MerkleRoots:        []cciptypes.MerkleRootChain{},
				OnRampMaxSeqNums:   []plugintypes.SeqNumChain{},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{{}},
				RMNRemoteConfig:    rmntypes.RemoteConfig{},
				FChain:             map[cciptypes.ChainSelector]int{},
			},
			expected: false,
		},
		{
			name: "Non-empty RMNRemoteConfig",
			observation: Observation{
				MerkleRoots:        []cciptypes.MerkleRootChain{},
				OnRampMaxSeqNums:   []plugintypes.SeqNumChain{},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{},
				RMNRemoteConfig:    rmntypes.RemoteConfig{ContractAddress: cciptypes.UnknownAddress([]byte("hello world address"))},
				FChain:             map[cciptypes.ChainSelector]int{},
			},
			expected: false,
		},
		{
			name: "Non-empty FChain",
			observation: Observation{
				MerkleRoots:        []cciptypes.MerkleRootChain{},
				OnRampMaxSeqNums:   []plugintypes.SeqNumChain{},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{},
				RMNRemoteConfig:    rmntypes.RemoteConfig{},
				FChain:             map[cciptypes.ChainSelector]int{1: 1},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.observation.IsEmpty())
		})
	}
}

func TestAggregateObservations(t *testing.T) {
	tests := []struct {
		name         string
		observations []plugincommon.AttributedObservation[Observation]
		expected     aggregatedObservation
	}{
		{
			name:         "Empty Observations",
			observations: []plugincommon.AttributedObservation[Observation]{},
			expected: aggregatedObservation{
				MerkleRoots:        make(map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain),
				OnRampMaxSeqNums:   make(map[cciptypes.ChainSelector][]cciptypes.SeqNum),
				OffRampNextSeqNums: make(map[cciptypes.ChainSelector][]cciptypes.SeqNum),
				RMNRemoteConfigs:   make([]rmntypes.RemoteConfig, 0),
				FChain:             make(map[cciptypes.ChainSelector][]int),
			},
		},
		{
			name: "Single Observation",
			observations: []plugincommon.AttributedObservation[Observation]{
				{
					Observation: Observation{
						MerkleRoots:        []cciptypes.MerkleRootChain{{ChainSel: 1}},
						OnRampMaxSeqNums:   []plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 1}},
						OffRampNextSeqNums: []plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 1}},
						RMNRemoteConfig:    rmntypes.RemoteConfig{ContractAddress: cciptypes.UnknownAddress([]byte("address"))},
						FChain:             map[cciptypes.ChainSelector]int{1: 1},
					},
				},
			},
			expected: aggregatedObservation{
				MerkleRoots:        map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain{1: {{ChainSel: 1}}},
				OnRampMaxSeqNums:   map[cciptypes.ChainSelector][]cciptypes.SeqNum{1: {1}},
				OffRampNextSeqNums: map[cciptypes.ChainSelector][]cciptypes.SeqNum{1: {1}},
				RMNRemoteConfigs:   []rmntypes.RemoteConfig{{ContractAddress: cciptypes.UnknownAddress([]byte("address"))}},
				FChain:             map[cciptypes.ChainSelector][]int{1: {1}},
			},
		},
		{
			name: "Multiple Observations",
			observations: []plugincommon.AttributedObservation[Observation]{
				{
					Observation: Observation{
						MerkleRoots:        []cciptypes.MerkleRootChain{{ChainSel: 1}},
						OnRampMaxSeqNums:   []plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 1}},
						OffRampNextSeqNums: []plugintypes.SeqNumChain{{ChainSel: 1, SeqNum: 1}},
						RMNRemoteConfig:    rmntypes.RemoteConfig{ContractAddress: cciptypes.UnknownAddress([]byte("address1"))},
						FChain:             map[cciptypes.ChainSelector]int{1: 1},
					},
				},
				{
					Observation: Observation{
						MerkleRoots:        []cciptypes.MerkleRootChain{{ChainSel: 2}},
						OnRampMaxSeqNums:   []plugintypes.SeqNumChain{{ChainSel: 2, SeqNum: 2}},
						OffRampNextSeqNums: []plugintypes.SeqNumChain{{ChainSel: 2, SeqNum: 2}},
						RMNRemoteConfig:    rmntypes.RemoteConfig{ContractAddress: cciptypes.UnknownAddress([]byte("address2"))},
						FChain:             map[cciptypes.ChainSelector]int{2: 2},
					},
				},
			},
			expected: aggregatedObservation{
				MerkleRoots:        map[cciptypes.ChainSelector][]cciptypes.MerkleRootChain{1: {{ChainSel: 1}}, 2: {{ChainSel: 2}}},
				OnRampMaxSeqNums:   map[cciptypes.ChainSelector][]cciptypes.SeqNum{1: {1}, 2: {2}},
				OffRampNextSeqNums: map[cciptypes.ChainSelector][]cciptypes.SeqNum{1: {1}, 2: {2}},
				RMNRemoteConfigs: []rmntypes.RemoteConfig{
					{ContractAddress: cciptypes.UnknownAddress([]byte("address1"))},
					{ContractAddress: cciptypes.UnknownAddress([]byte("address2"))},
				},
				FChain: map[cciptypes.ChainSelector][]int{1: {1}, 2: {2}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := aggregateObservations(tt.observations)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestOutcome_Sort(t *testing.T) {
	tests := []struct {
		name     string
		outcome  Outcome
		expected Outcome
	}{
		{
			name: "Sort RangesSelectedForReport",
			outcome: Outcome{
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: 2},
					{ChainSel: 1},
				},
				RootsToReport:      []cciptypes.MerkleRootChain{},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{},
			},
			expected: Outcome{
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: 1},
					{ChainSel: 2},
				},
				RootsToReport:      []cciptypes.MerkleRootChain{},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{},
			},
		},
		{
			name: "Sort RootsToReport",
			outcome: Outcome{
				RangesSelectedForReport: []plugintypes.ChainRange{},
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: 2},
					{ChainSel: 1},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{},
			},
			expected: Outcome{
				RangesSelectedForReport: []plugintypes.ChainRange{},
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: 1},
					{ChainSel: 2},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{},
			},
		},
		{
			name: "Sort OffRampNextSeqNums",
			outcome: Outcome{
				RangesSelectedForReport: []plugintypes.ChainRange{},
				RootsToReport:           []cciptypes.MerkleRootChain{},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: 2},
					{ChainSel: 1},
				},
			},
			expected: Outcome{
				RangesSelectedForReport: []plugintypes.ChainRange{},
				RootsToReport:           []cciptypes.MerkleRootChain{},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: 1},
					{ChainSel: 2},
				},
			},
		},
		{
			name: "Sort all fields",
			outcome: Outcome{
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: 3},
					{ChainSel: 1},
					{ChainSel: 2},
				},
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: 3},
					{ChainSel: 1},
					{ChainSel: 2},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: 3},
					{ChainSel: 1},
					{ChainSel: 2},
				},
			},
			expected: Outcome{
				RangesSelectedForReport: []plugintypes.ChainRange{
					{ChainSel: 1},
					{ChainSel: 2},
					{ChainSel: 3},
				},
				RootsToReport: []cciptypes.MerkleRootChain{
					{ChainSel: 1},
					{ChainSel: 2},
					{ChainSel: 3},
				},
				OffRampNextSeqNums: []plugintypes.SeqNumChain{
					{ChainSel: 1},
					{ChainSel: 2},
					{ChainSel: 3},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.outcome.Sort()
			assert.Equal(t, tt.expected, tt.outcome)
		})
	}
}