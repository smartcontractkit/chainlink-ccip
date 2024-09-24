package plugintypes

import (
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

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

type DonID = uint32
