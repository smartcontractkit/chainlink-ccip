package ccipocr3

import (
	ccipocr3common "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// Deprecated: Use ccipocr3common.CommitPluginReport instead.
type CommitPluginReport = ccipocr3common.CommitPluginReport

// Deprecated: Use ccipocr3common.MerkleRootChain instead.
type MerkleRootChain = ccipocr3common.MerkleRootChain

// Deprecated: Use ccipocr3common.PriceUpdates instead.
type PriceUpdates = ccipocr3common.PriceUpdates

// Deprecated: Use ccipocr3common.CommitReportInfo instead.
type CommitReportInfo = ccipocr3common.CommitReportInfo

// Deprecated: Use ccipocr3common.DecodeCommitReportInfo instead.
func DecodeCommitReportInfo(data []byte) (CommitReportInfo, error) {
	return ccipocr3common.DecodeCommitReportInfo(data)
}
