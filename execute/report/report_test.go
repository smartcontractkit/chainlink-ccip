package report

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/types"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

// TODO: better than this
type tdr struct{}

func (t tdr) ReadTokenData(
	ctx context.Context, srcChain cciptypes.ChainSelector, num cciptypes.SeqNum) ([][]byte, error,
) {
	return nil, nil
}

type badHasher struct{}

func (bh badHasher) Hash(context.Context, cciptypes.Message) (cciptypes.Bytes32, error) {
	return cciptypes.Bytes32{}, fmt.Errorf("bad hasher")
}

type badTokenDataReader struct{}

func (btdr badTokenDataReader) ReadTokenData(
	_ context.Context, _ cciptypes.ChainSelector, _ cciptypes.SeqNum,
) ([][]byte, error) {
	return nil, fmt.Errorf("bad token data reader")
}

type badCodec struct{}

func (bc badCodec) Encode(ctx context.Context, report cciptypes.ExecutePluginReport) ([]byte, error) {
	return nil, fmt.Errorf("bad codec")
}

func (bc badCodec) Decode(ctx context.Context, bytes []byte) (cciptypes.ExecutePluginReport, error) {
	return cciptypes.ExecutePluginReport{}, fmt.Errorf("bad codec")
}

func Test_buildSingleChainReport_Errors(t *testing.T) {
	lggr := logger.Test(t)

	type args struct {
		report          plugintypes.ExecutePluginCommitDataWithMessages
		hasher          cciptypes.MessageHasher
		tokenDataReader types.TokenDataReader
		codec           cciptypes.ExecutePluginCodec
	}
	tests := []struct {
		name    string
		args    args
		wantErr string
	}{
		{
			name:    "wrong number of messages",
			wantErr: "unexpected number of messages: expected 1, got 2",
			args: args{
				report: plugintypes.ExecutePluginCommitDataWithMessages{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					},
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{}},
						{Header: cciptypes.RampMessageHeader{}},
					},
				},
			},
		},
		{
			name:    "wrong sequence numbers",
			wantErr: "sequence number 102 outside of report range [100 -> 101]",
			args: args{
				report: plugintypes.ExecutePluginCommitDataWithMessages{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(101)),
					},
					Messages: []cciptypes.Message{
						{
							Header: cciptypes.RampMessageHeader{
								SequenceNumber: cciptypes.SeqNum(100),
							},
						},
						{
							Header: cciptypes.RampMessageHeader{
								SequenceNumber: cciptypes.SeqNum(102),
							},
						},
					},
				},
			},
		},
		{
			name:    "source mismatch",
			wantErr: "unexpected source chain: expected 1111, got 2222",
			args: args{
				report: plugintypes.ExecutePluginCommitDataWithMessages{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SourceChain:         1111,
						SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					},
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{
							SourceChainSelector: 2222,
							SequenceNumber:      cciptypes.SeqNum(100),
						}},
					},
				},
				hasher: badHasher{},
			},
		},
		{
			name:    "bad hasher",
			wantErr: "unable to hash message (1234567, 100): bad hasher",
			args: args{
				report: plugintypes.ExecutePluginCommitDataWithMessages{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SourceChain:         1234567,
						SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					},
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{
							SourceChainSelector: 1234567,
							SequenceNumber:      cciptypes.SeqNum(100),
						}},
					},
				},
				hasher: badHasher{},
			},
		},
		{
			name:    "bad token data reader",
			wantErr: "unable to read token data for message 100: bad token data reader",
			args: args{
				report: plugintypes.ExecutePluginCommitDataWithMessages{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SourceChain:         1234567,
						SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					},
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{
							SourceChainSelector: 1234567,
							SequenceNumber:      cciptypes.SeqNum(100),
						}},
					},
				},
				tokenDataReader: badTokenDataReader{},
			},
		},
		{
			name:    "bad codec",
			wantErr: "unable to encode report: bad codec",
			args: args{
				report: plugintypes.ExecutePluginCommitDataWithMessages{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SourceChain:         1234567,
						SequenceNumberRange: cciptypes.NewSeqNumRange(cciptypes.SeqNum(100), cciptypes.SeqNum(100)),
					},
					Messages: []cciptypes.Message{
						{Header: cciptypes.RampMessageHeader{
							SourceChainSelector: 1234567,
							SequenceNumber:      cciptypes.SeqNum(100),
						}},
					},
				},
				codec: badCodec{},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			// Select hasher mock.
			var resolvedHasher cciptypes.MessageHasher
			if tt.args.hasher != nil {
				resolvedHasher = tt.args.hasher
			} else {
				resolvedHasher = mocks.NewMessageHasher()
			}

			// Select token data reader mock.
			var resolvedTokenDataReader types.TokenDataReader
			if tt.args.tokenDataReader != nil {
				resolvedTokenDataReader = tt.args.tokenDataReader
			} else {
				resolvedTokenDataReader = tdr{}
			}

			// Select codec mock.
			var resolvedCodec cciptypes.ExecutePluginCodec
			if tt.args.codec != nil {
				resolvedCodec = tt.args.codec
			} else {
				resolvedCodec = mocks.NewExecutePluginJSONReportCodec()
			}

			ctx := context.Background()
			execReport, size, err := buildSingleChainReport(
				ctx, lggr, resolvedHasher, resolvedTokenDataReader, resolvedCodec, tt.args.report, 0)
			if tt.wantErr != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.wantErr)
				return
			}
			require.NoError(t, err)
			fmt.Println(execReport, size, err)
		})
	}
}
