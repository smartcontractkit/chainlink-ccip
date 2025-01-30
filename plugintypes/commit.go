package plugintypes

import (
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// NOTE: The following type should be moved to internal plugin types after it's not required anymore in chainlink repo.
// Right now it's only used in a chainlink repo test: TestCCIPReader_CommitReportsGTETimestamp

type CommitPluginReportWithMeta struct {
	Report    cciptypes.CommitPluginReport `json:"report"`
	Timestamp time.Time                    `json:"timestamp"`
	BlockNum  uint64                       `json:"blockNum"`
	// DisabledSourceChains is a list of source chains that are disabled for the merkleRoots of this report.
	DisabledSourceChains []cciptypes.ChainSelector `json:"disabledSourceChains"`
}
