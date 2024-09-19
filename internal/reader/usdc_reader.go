package reader

import (
	"context"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type USDCMessageReader interface {
	MessageHashes(ctx context.Context,
		source cciptypes.ChainSelector,
		seqNums []cciptypes.SeqNum,
	) (map[cciptypes.SeqNum]map[int][]byte, error)
}

type FakeUSDCMessageReader struct {
	Messages map[cciptypes.SeqNum]map[int][]byte
}

func (f FakeUSDCMessageReader) MessageHashes(
	_ context.Context,
	_ cciptypes.ChainSelector,
	seqNums []cciptypes.SeqNum,
) (map[cciptypes.SeqNum]map[int][]byte, error) {
	outcome := make(map[cciptypes.SeqNum]map[int][]byte)
	for _, seqNum := range seqNums {
		if messages, ok := f.Messages[seqNum]; ok {
			outcome[seqNum] = messages
		}
	}
	return outcome, nil
}
