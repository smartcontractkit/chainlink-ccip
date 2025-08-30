package exectypes

import (
	"sort"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
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
	// * Filter: Commit reports associated with the final execution report, concatenated together.
	CommitReports []CommitData `json:"commitReports"`

	// Report is built from the oldest pending commit reports.
	// Deprecated: Use Reports field instead.
	Report cciptypes.ExecutePluginReport `json:"report"`

	// Reports are built from the oldest pending commit reports.
	Reports []cciptypes.ExecutePluginReport `json:"reports"`
}

// IsEmpty returns true if the outcome has no pending commit reports or chain reports.
func (o Outcome) IsEmpty() bool {
	return len(o.CommitReports) == 0 && (len(o.Reports) == 0 || len(o.Reports[0].ChainReports) == 0)
}

func (o *Outcome) Stats() map[string]int {
	counters := map[string]int{
		messagesLabel:     0,
		tokenDataLabel:    0,
		sourceChainsLabel: 0,
	}

	for _, execReport := range o.Reports {
		for _, report := range execReport.ChainReports {
			counters[sourceChainsLabel]++
			counters[messagesLabel] += len(report.Messages)
			counters[tokenDataLabel] += len(report.OffchainTokenData)
		}
	}
	return counters
}

// ToLogFormat creates a copy of the outcome with the messages data removed.
func (o Outcome) ToLogFormat() Outcome {
	commitReports := make([]CommitData, len(o.CommitReports))
	for i, report := range o.CommitReports {
		commitReports[i] = report.CopyNoMsgData()
	}
	truncatedReports := make([]cciptypes.ExecutePluginReport, len(o.Reports))
	for i, execReport := range o.Reports {
		truncatedReports[i].ChainReports = make([]cciptypes.ExecutePluginReportSingleChain, len(execReport.ChainReports))
		for j, report := range execReport.ChainReports {
			truncatedReports[i].ChainReports[j] = report.CopyNoMsgData()
		}
	}
	cleanedOutcome := Outcome{
		State:         o.State,
		CommitReports: commitReports,
		Reports:       truncatedReports,
	}
	return cleanedOutcome
}

// NewOutcome creates a new Outcome with the pending commit reports and the chain reports sorted.
// No sorting is needed, they are already in a canonical representation from the builder.
func NewOutcome(
	state PluginState,
	selectedCommits []CommitData,
	reports []cciptypes.ExecutePluginReport,
) Outcome {
	return Outcome{
		State:         state,
		CommitReports: selectedCommits,
		Reports:       reports,
	}
}

func NewOutcomeWithSortedCommitReports(
	state PluginState,
	commitReports []CommitData,
) Outcome {
	sort.Slice(commitReports, func(i, j int) bool {
		return LessThan(commitReports[i], commitReports[j])
	})
	return Outcome{
		State:         state,
		CommitReports: commitReports,
		Reports:       nil,
	}
}
