package reader

import (
	"context"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

type USDCMessageHash struct {
	Source      cciptypes.ChainSelector
	SeqNr       cciptypes.SeqNum
	MessageHash []byte
}

type USDCMessageReader interface {
	MessageHashes(ctx context.Context,
		source cciptypes.ChainSelector,
		seqNums []cciptypes.SeqNum,
	) (map[cciptypes.SeqNum][][32]byte, error)
}

type NoopUSDCMessageReader struct{}

func (n *NoopUSDCMessageReader) MessageHashes(
	ctx context.Context,
	source cciptypes.ChainSelector,
	seqNums []cciptypes.SeqNum,
) (map[cciptypes.SeqNum][][32]byte, error) {
	result := map[cciptypes.SeqNum][][32]byte{}
	for _, seqNum := range seqNums {
		result[seqNum] = [][32]byte{}
	}
	return result, nil
}
