package plugintypes

import (
	"time"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type CommitPluginReportWithMeta struct {
	Report    cciptypes.CommitPluginReport `json:"report"`
	Timestamp time.Time                    `json:"timestamp"`
	BlockNum  uint64                       `json:"blockNum"`
}

type SeqNumChain struct {
	ChainSel cciptypes.ChainSelector `json:"chainSel"`
	SeqNum   cciptypes.SeqNum        `json:"seqNum"`
}

func NewSeqNumChain(chainSel cciptypes.ChainSelector, seqNum cciptypes.SeqNum) SeqNumChain {
	return SeqNumChain{
		ChainSel: chainSel,
		SeqNum:   seqNum,
	}
}

type ChainRange struct {
	ChainSel    cciptypes.ChainSelector `json:"chain"`
	SeqNumRange cciptypes.SeqNumRange   `json:"seqNumRange"`
}
