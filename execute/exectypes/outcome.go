package exectypes

import (
	"sort"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

type PluginState string

const (
	// Unknown is the zero value, this allows a "Next" state transition for uninitialized values (i.e. the first round).
	Unknown PluginState = ""

	// Initialized is used to indicate that there was nothing to do except initialize contract addresses.
	Initialized PluginState = "Initialized"

	// GetCommitReports is the first step, it is used to select commit reports from the destination chain.
	GetCommitReports PluginState = "GetCommitReports"

	// GetMessages is the second step, given a set of commit reports it fetches the associated messages.
	GetMessages PluginState = "GetMessages"

	// Filter is the final step, any additional destination data is collected to complete the execution report.
	Filter PluginState = "Filter"
)

// Next returns the next state for the plugin. The Unknown state is used to transition from uninitialized values.
func (p PluginState) Next() PluginState {
	switch p {
	case GetCommitReports:
		return GetMessages

	case GetMessages:
		return Filter

	case Unknown:
		fallthrough
	case Initialized:
		fallthrough
	case Filter:
		return GetCommitReports

	default:
		panic("unexpected execute plugin state")
	}
}

// Outcome is the outcome of the ExecutePlugin.
type Outcome struct {
	// State that the outcome was generated for.
	State PluginState

	// CommitReports are the oldest reports with pending commits. The slice is
	// sorted from oldest to newest. The content of this field changes based on
	// the state:
	// * GetCommitReports: All pending commit reports which were observed.
	// * GetMessages: All pending commit reports with messages.
	// * Filter: Commit reports associated with the final execution report.
	CommitReports []CommitData `json:"commitReports"`

	// Report is built from the oldest pending commit reports.
	Report cciptypes.ExecutePluginReport `json:"report"`
}

// IsEmpty returns true if the outcome has no pending commit reports or chain reports.
func (o Outcome) IsEmpty() bool {
	return len(o.CommitReports) == 0 && len(o.Report.ChainReports) == 0
}

func (o *Outcome) Stats() map[string]int {
	counters := map[string]int{
		messagesLabel:     0,
		tokenDataLabel:    0,
		sourceChainsLabel: 0,
	}

	for _, report := range o.Report.ChainReports {
		counters[sourceChainsLabel]++
		counters[messagesLabel] += len(report.Messages)
		counters[tokenDataLabel] += len(report.OffchainTokenData)
	}
	return counters
}

// ToLogFormat creates a copy of the outcome with the messages data removed.
func (o Outcome) ToLogFormat() Outcome {
	commitReports := make([]CommitData, len(o.CommitReports))
	for i, report := range o.CommitReports {
		commitReports[i] = report.CopyNoMsgData()
	}
	reports := make([]cciptypes.ExecutePluginReportSingleChain, len(o.Report.ChainReports))
	for i, report := range o.Report.ChainReports {
		reports[i] = report.CopyNoMsgData()
	}
	cleanedOutcome := Outcome{
		State:         o.State,
		CommitReports: commitReports,
		Report: cciptypes.ExecutePluginReport{
			ChainReports: reports,
		},
	}
	return cleanedOutcome
}

// NewOutcome creates a new Outcome with the pending commit reports and the chain reports sorted.
func NewOutcome(
	state PluginState,
	selectedCommits []CommitData,
	report cciptypes.ExecutePluginReport,
) Outcome {
	return NewSortedOutcome(state, selectedCommits, report)
}

// NewSortedOutcome ensures canonical ordering of the outcome.
// TODO: handle canonicalization in the encoder.
func NewSortedOutcome(
	state PluginState,
	pendingCommits []CommitData,
	report cciptypes.ExecutePluginReport,
) Outcome {
	pendingCommitsCP := append([]CommitData{}, pendingCommits...)
	reportCP := append([]cciptypes.ExecutePluginReportSingleChain{}, report.ChainReports...)
	sort.Slice(
		pendingCommitsCP,
		func(i, j int) bool {
			return LessThan(pendingCommitsCP[i], pendingCommitsCP[j])
		})
	sort.Slice(
		reportCP,
		func(i, j int) bool {
			return reportCP[i].SourceChainSelector < reportCP[j].SourceChainSelector
		})
	return Outcome{
		State:         state,
		CommitReports: pendingCommitsCP,
		Report:        cciptypes.ExecutePluginReport{ChainReports: reportCP},
	}
}
