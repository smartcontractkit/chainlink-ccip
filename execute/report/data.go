package report

import (
	"sort"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"slices"
)

// markNewMessagesExecuted compares an execute plugin report with the commit report metadata and marks the new messages
// as executed.
func markNewMessagesExecuted(
	execReport cciptypes.ExecutePluginReportSingleChain, report exectypes.CommitData,
) exectypes.CommitData {
	// Mark new messages executed.
	for i := range execReport.Messages {
		report.ExecutedMessages =
			append(report.ExecutedMessages, execReport.Messages[i].Header.SequenceNumber)
	}
	slices.Sort(
		report.ExecutedMessages)

	return report
}
