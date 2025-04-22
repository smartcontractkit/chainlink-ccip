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
)

func TestInMemoryCCIPReader_CommitReportsGTETimestamp(t *testing.T) {
	type fields struct {
		Reports []cciptypes.CommitPluginReportWithMeta
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
				Reports: []cciptypes.CommitPluginReportWithMeta{
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
				Reports: []cciptypes.CommitPluginReportWithMeta{
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
				UnfinalizedReports: tt.fields.Reports}
			got, err := r.CommitReportsGTETimestamp(
				context.Background(),
				tt.args.ts,
				primitives.Unconfirmed,
				tt.args.limit,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("CommitReportsGTETimestamp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(got) != len(tt.want) {
				t.Errorf("CommitReportsGTETimestamp() got = %v, want %v", got, tt.want)
				return
			}
			gotBlocks := slicelib.Map(got, func(report cciptypes.CommitPluginReportWithMeta) expected {
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
		dest          cciptypes.ChainSelector
		rangesByChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[cciptypes.ChainSelector][]cciptypes.SeqNum
		wantErr bool
	}{
		{
			name:    "empty",
			want:    map[cciptypes.ChainSelector][]cciptypes.SeqNum{},
			wantErr: false,
		},
		{
			name: "single executed message",
			args: args{
				rangesByChain: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
					1: {cciptypes.NewSeqNumRange(1, 100)},
				},
				dest: 2,
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
			want: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: cciptypes.NewSeqNumRange(51, 51).ToSlice(),
			},
			wantErr: false,
		},
		{
			name: "multiple executed message ranges",
			args: args{
				rangesByChain: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
					1: {cciptypes.NewSeqNumRange(1, 100)},
				},
				dest: 2,
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
			want: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {50, 52},
			},
			wantErr: false,
		},
		{
			name: "multiple messages in range",
			args: args{
				rangesByChain: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
					1: {cciptypes.NewSeqNumRange(1, 100)},
				},
				dest: 2,
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
			want: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: cciptypes.NewSeqNumRange(50, 52).ToSlice(),
			},
			wantErr: false,
		},
		{
			name: "disjoint range",
			args: args{
				rangesByChain: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
					1: {
						cciptypes.NewSeqNumRange(1, 50),
						cciptypes.NewSeqNumRange(100, 150),
					},
				},
				dest: 2,
			},
			fields: fields{
				MessagesWithExecuted: map[cciptypes.ChainSelector][]MessagesWithMetadata{
					1: {
						makeMsg(25, 2, true),
						makeMsg(26, 2, true),
						makeMsg(75, 2, true),
						makeMsg(125, 2, true),
						makeMsg(126, 2, false),
						makeMsg(127, 2, true),
					},
				},
				Dest: 2,
			},
			want: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {25, 26, 125, 127},
			},
			wantErr: false,
		},
		{
			name: "overlapping range",
			args: args{
				rangesByChain: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
					1: {
						cciptypes.NewSeqNumRange(1, 50),
						cciptypes.NewSeqNumRange(40, 100),
					},
				},
				dest: 2,
			},
			fields: fields{
				MessagesWithExecuted: map[cciptypes.ChainSelector][]MessagesWithMetadata{
					1: {
						makeMsg(25, 2, true),
						makeMsg(45, 2, true),
						makeMsg(46, 2, true),
						makeMsg(75, 2, true),
					},
				},
				Dest: 2,
			},
			want: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {25, 45, 46, 75},
			},
			wantErr: false,
		},
		{
			name: "multiple chains",
			args: args{
				rangesByChain: map[cciptypes.ChainSelector][]cciptypes.SeqNumRange{
					1: {
						cciptypes.NewSeqNumRange(1, 25),
						cciptypes.NewSeqNumRange(50, 75),
					},
					3: {
						cciptypes.NewSeqNumRange(1, 50),
						cciptypes.NewSeqNumRange(51, 100),
					},
				},
				dest: 2,
			},
			fields: fields{
				MessagesWithExecuted: map[cciptypes.ChainSelector][]MessagesWithMetadata{
					1: {
						makeMsg(25, 2, true),
						makeMsg(45, 2, true),
						makeMsg(46, 2, true),
						makeMsg(75, 2, true),
					},
					3: {
						makeMsg(25, 2, true),
						makeMsg(50, 2, true),
						makeMsg(51, 2, true),
						makeMsg(74, 2, false),
						makeMsg(75, 2, true),
					},
				},
				Dest: 2,
			},
			want: map[cciptypes.ChainSelector][]cciptypes.SeqNum{
				1: {25, 75},
				3: {25, 50, 51, 75},
			},
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
			got, err := r.ExecutedMessages(context.Background(), tt.args.rangesByChain, primitives.Finalized)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExecutedMessages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExecutedMessages() got = %+v, want %+v", got, tt.want)
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
				UnfinalizedReports: nil,
				Messages:           tt.fields.MessagesWithMeta,
				Dest:               tt.fields.Dest,
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
