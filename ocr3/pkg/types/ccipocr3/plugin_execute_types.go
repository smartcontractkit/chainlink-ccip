package ccipocr3

import (
	ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Deprecated: Use ccipocr3common.ExecutePluginReport instead.
type ExecutePluginReport = ccipocr3common.ExecutePluginReport

// Deprecated: Use ccipocr3common.ExecutePluginReportSingleChain instead.
type ExecutePluginReportSingleChain = ccipocr3common.ExecutePluginReportSingleChain

// Deprecated: Use ccipocr3common.ExecuteReportInfo instead.
type ExecuteReportInfo = ccipocr3common.ExecuteReportInfo

// Deprecated: Use ccipocr3common.DecodeExecuteReportInfo instead.
func DecodeExecuteReportInfo(data []byte) (ExecuteReportInfo, error) {
	return ccipocr3common.DecodeExecuteReportInfo(data)
}
