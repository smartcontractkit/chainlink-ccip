package inmem

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"

	"github.com/smartcontractkit/chainlink-ccip/internal/libs/slicelib"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func TestInMemoryCCIPReader_CommitReportsGTETimestamp(t *testing.T) {
	type fields struct {
		Reports []plugintypes2.CommitPluginReportWithMeta
	}
	type args struct {
		ts    time.Time
		limit int
	}
	type expected struct {
		block uint64
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []expected
		wantErr bool
	}{
		{
			name:    "empty",
			want:    nil,
			wantErr: false,
		},
		{
			name: "filtered by time",
			args: args{
				ts:    time.UnixMicro(200000000),
				limit: 1000,
			},
			fields: fields{
				Reports: []plugintypes2.CommitPluginReportWithMeta{
					{
						Timestamp: time.UnixMicro(100000000),
						BlockNum:  1000,
					},
					{
						Timestamp: time.UnixMicro(200000000), // GTE this timestamp
						BlockNum:  1001,
					},
					{
						Timestamp: time.UnixMicro(300000000),
						BlockNum:  1002,
					},
				},
			},
			want:    []expected{{block: 1001}, {block: 1002}},
			wantErr: false,
		},
		{
			name: "filter by limit",
			args: args{
				ts:    time.UnixMicro(100000000),
				limit: 1,
			},
			fields: fields{
				Reports: []plugintypes2.CommitPluginReportWithMeta{
					{
						Timestamp: time.UnixMicro(100000000),
						BlockNum:  1000,
					},
					{
						Timestamp: time.UnixMicro(200000000), // GTE this timestamp
						BlockNum:  1001,
					},
				},
			},
			want:    []expected{{block: 1000}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := InMemoryCCIPReader{
				Reports: tt.fields.Reports}
			got, err := r.CommitReportsGTETimestamp(context.Background(), tt.args.ts, tt.args.limit)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommitReportsGTETimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("CommitReportsGTETimestamp() got = %v, want %v", got, tt.want)
				return
			}
			gotBlocks := slicelib.Map(got, func(report plugintypes2.CommitPluginReportWithMeta) expected {
				return expected{block: report.BlockNum}
			})
			require.ElementsMatchf(t, gotBlocks, tt.want, "CommitReportsGTETimestamp() got = %v, want %v", got, tt.want)
		})
	}
}

func makeMsg(seqNum cciptypes.SeqNum, dest cciptypes.ChainSelector, executed bool) MessagesWithMetadata {
	return MessagesWithMetadata{
		Message: cciptypes.Message{
			Header: cciptypes.RampMessageHeader{
				SequenceNumber: seqNum,
			},
		},
		Destination: dest,
		Executed:    executed,
	}
}

func TestInMemoryCCIPReader_ExecutedMessages(t *testing.T) {
	type fields struct {
		MessagesWithExecuted map[cciptypes.ChainSelector][]MessagesWithMetadata
		Dest                 cciptypes.ChainSelector
	}
	type args struct {
		source      cciptypes.ChainSelector
		dest        cciptypes.ChainSelector
		seqNumRange cciptypes.SeqNumRange
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []cciptypes.SeqNum
		wantErr bool
	}{
		{
			name:    "empty",
			want:    nil,
			wantErr: false,
		},
		{
			name: "single executed message",
			args: args{
				source:      1,
				dest:        2,
				seqNumRange: cciptypes.NewSeqNumRange(1, 100),
			},
			fields: fields{
				MessagesWithExecuted: map[cciptypes.ChainSelector][]MessagesWithMetadata{
					1: {
						makeMsg(50, 2, false),
						makeMsg(51, 2, true),
						makeMsg(52, 2, false),
					},
				},
				Dest: 2,
			},
			want:    cciptypes.NewSeqNumRange(51, 51).ToSlice(),
			wantErr: false,
		},
		{
			name: "multiple executed message ranges",
			args: args{
				source:      1,
				dest:        2,
				seqNumRange: cciptypes.NewSeqNumRange(1, 100),
			},
			fields: fields{
				MessagesWithExecuted: map[cciptypes.ChainSelector][]MessagesWithMetadata{
					1: {
						makeMsg(50, 2, true),
						makeMsg(51, 2, false),
						makeMsg(52, 2, true),
					},
				},
				Dest: 2,
			},
			want:    []cciptypes.SeqNum{50, 52},
			wantErr: false,
		},
		{
			name: "multiple messages in range",
			args: args{
				source:      1,
				dest:        2,
				seqNumRange: cciptypes.NewSeqNumRange(1, 100),
			},
			fields: fields{
				MessagesWithExecuted: map[cciptypes.ChainSelector][]MessagesWithMetadata{
					1: {
						makeMsg(50, 2, true),
						makeMsg(51, 2, true),
						makeMsg(52, 2, true),
					},
				},
				Dest: 2,
			},
			want:    cciptypes.NewSeqNumRange(50, 52).ToSlice(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := InMemoryCCIPReader{
				Messages: tt.fields.MessagesWithExecuted,
				Dest:     tt.fields.Dest,
			}
			got, err := r.ExecutedMessages(context.Background(), tt.args.source, tt.args.seqNumRange, primitives.Finalized)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecutedMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecutedMessages() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInMemoryCCIPReader_MsgsBetweenSeqNums(t *testing.T) {
	type fields struct {
		MessagesWithMeta map[cciptypes.ChainSelector][]MessagesWithMetadata
		Dest             cciptypes.ChainSelector
	}
	type args struct {
		chain       cciptypes.ChainSelector
		seqNumRange cciptypes.SeqNumRange
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []cciptypes.Message
		wantErr bool
	}{
		{
			name:    "empty",
			want:    nil,
			wantErr: false,
		},
		{
			name: "single message",
			args: args{
				chain:       1,
				seqNumRange: cciptypes.NewSeqNumRange(1, 100),
			},
			fields: fields{
				Dest: 2,
				MessagesWithMeta: map[cciptypes.ChainSelector][]MessagesWithMetadata{
					1: {
						makeMsg(50, 1, false),
						makeMsg(51, 2, false),
						makeMsg(52, 3, false),
					},
				},
			},
			want: []cciptypes.Message{
				makeMsg(51, 2, false).Message,
			},
		},
		{
			name: "multiple messages",
			args: args{
				chain:       1,
				seqNumRange: cciptypes.NewSeqNumRange(1, 100),
			},
			fields: fields{
				Dest: 2,
				MessagesWithMeta: map[cciptypes.ChainSelector][]MessagesWithMetadata{
					1: {
						makeMsg(50, 2, false),
						makeMsg(51, 2, false),
						makeMsg(52, 2, false),
					},
				},
			},
			want: []cciptypes.Message{
				makeMsg(50, 2, false).Message,
				makeMsg(51, 2, false).Message,
				makeMsg(52, 2, false).Message,
			},
		},
		{
			name: "no messages for chain",
			args: args{
				chain:       1,
				seqNumRange: cciptypes.NewSeqNumRange(1, 100),
			},
			fields: fields{
				Dest: 2,
				MessagesWithMeta: map[cciptypes.ChainSelector][]MessagesWithMetadata{
					10: {
						makeMsg(50, 2, false),
						makeMsg(51, 2, false),
						makeMsg(52, 2, false),
					},
				},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			r := InMemoryCCIPReader{
				Reports:  nil,
				Messages: tt.fields.MessagesWithMeta,
				Dest:     tt.fields.Dest,
			}
			got, err := r.MsgsBetweenSeqNums(context.Background(), tt.args.chain, tt.args.seqNumRange)
			if (err != nil) != tt.wantErr {
				t.Errorf("MsgsBetweenSeqNums() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MsgsBetweenSeqNums() got = %v, want %v", got, tt.want)
			}
		})
	}
}
