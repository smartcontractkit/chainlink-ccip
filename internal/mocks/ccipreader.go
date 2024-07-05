package mocks

import (
	"context"
	"time"

	"github.com/stretchr/testify/mock"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

type CCIPReader struct {
	*mock.Mock
}

func NewCCIPReader() *CCIPReader {
	return &CCIPReader{
		Mock: &mock.Mock{},
	}
}

func (r CCIPReader) CommitReportsGTETimestamp(
	ctx context.Context, dest cciptypes.ChainSelector, ts time.Time, limit int,
) ([]plugintypes.CommitPluginReportWithMeta, error) {
	args := r.Called(ctx, dest, ts, limit)
	return args.Get(0).([]plugintypes.CommitPluginReportWithMeta), args.Error(1)
}

func (r CCIPReader) ExecutedMessageRanges(
	ctx context.Context, source, dest cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.SeqNumRange, error) {
	args := r.Called(ctx, source, dest, seqNumRange)
	return args.Get(0).([]cciptypes.SeqNumRange), args.Error(1)
}

func (r CCIPReader) MsgsBetweenSeqNums(
	ctx context.Context, chain cciptypes.ChainSelector, seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.Message, error) {
	args := r.Called(ctx, chain, seqNumRange)
	return args.Get(0).([]cciptypes.Message), args.Error(1)
}

func (r CCIPReader) NextSeqNum(ctx context.Context, chains []cciptypes.ChainSelector) (
	seqNum []cciptypes.SeqNum, err error) {
	args := r.Called(ctx, chains)
	return args.Get(0).([]cciptypes.SeqNum), args.Error(1)
}

func (r CCIPReader) GasPrices(ctx context.Context, chains []cciptypes.ChainSelector) ([]cciptypes.BigInt, error) {
	args := r.Called(ctx, chains)
	return args.Get(0).([]cciptypes.BigInt), args.Error(1)
}

func (r CCIPReader) Sync(ctx context.Context) (bool, error) {
	args := r.Called(ctx)
	return args.Bool(0), args.Error(1)
}

func (r CCIPReader) Close(ctx context.Context) error {
	args := r.Called(ctx)
	return args.Error(0)
}
