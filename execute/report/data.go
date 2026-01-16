package report

import (
	"slices"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
)

// markNewMessagesExecuted compares an execute plugin report with the commit report metadata and marks the new messages
// as executed.
func markNewMessagesExecuted(
	execReport cciptypes.ExecutePluginReportSingleChain, report exectypes.CommitData,
) exectypes.CommitData {
	// Mark new messages executed.
	for i := 0; i < len(execReport.Messages); i++ {
		report.ExecutedMessages =
			append(report.ExecutedMessages, execReport.Messages[i].Header.SequenceNumber)
	}
	slices.Sort(
		report.ExecutedMessages)

	return report
}
