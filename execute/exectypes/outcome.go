package exectypes

import (
	"encoding/json"
	"sort"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Outcome is the outcome of the ExecutePlugin.
type Outcome struct {
	// PendingCommitReports are the oldest reports with pending commits. The slice is
	// sorted from oldest to newest.
	PendingCommitReports []CommitData `json:"commitReports"`

	// Report is built from the oldest pending commit reports.
	Report cciptypes.ExecutePluginReport `json:"report"`
}

func (o Outcome) IsEmpty() bool {
	return len(o.PendingCommitReports) == 0 && len(o.Report.ChainReports) == 0
}

func NewOutcome(
	pendingCommits []CommitData,
	report cciptypes.ExecutePluginReport,
) Outcome {
	return newSortedOutcome(pendingCommits, report)
}

// Encode encodes the outcome by first sorting the pending commit reports and the chain reports
// and then JSON marshalling.
// The encoding MUST be deterministic.
func (o Outcome) Encode() ([]byte, error) {
	// We sort again here in case construction is not via the constructor.
	return json.Marshal(newSortedOutcome(o.PendingCommitReports, o.Report))
}

func newSortedOutcome(
	pendingCommits []CommitData,
	report cciptypes.ExecutePluginReport) Outcome {
	pendingCommitsCP := append([]CommitData{}, pendingCommits...)
	reportCP := append([]cciptypes.ExecutePluginReportSingleChain{}, report.ChainReports...)
	sort.Slice(
		pendingCommitsCP,
		func(i, j int) bool {
			return pendingCommitsCP[i].SourceChain < pendingCommitsCP[j].SourceChain
		})
	sort.Slice(
		reportCP,
		func(i, j int) bool {
			return reportCP[i].SourceChainSelector < reportCP[j].SourceChainSelector
		})
	return Outcome{
		PendingCommitReports: pendingCommitsCP,
		Report:               cciptypes.ExecutePluginReport{ChainReports: reportCP},
	}
}

func DecodeOutcome(b []byte) (Outcome, error) {
	o := Outcome{}
	err := json.Unmarshal(b, &o)
	return o, err
}
