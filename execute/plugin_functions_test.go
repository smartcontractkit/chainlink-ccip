package execute

import (
	"fmt"
	"testing"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/plugintypes"
)

func Test_validateObserverReadingEligibility(t *testing.T) {
	tests := []struct {
		name         string
		observerCfg  mapset.Set[cciptypes.ChainSelector]
		observedMsgs plugintypes.ExecutePluginMessageObservations
		expErr       string
	}{
		{
			name:        "ValidObserverAndMessages",
			observerCfg: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
			observedMsgs: plugintypes.ExecutePluginMessageObservations{
				1: {1: {}, 2: {}},
				2: {},
			},
		},
		{
			name:        "ObserverNotAllowedToReadChain",
			observerCfg: mapset.NewSet(cciptypes.ChainSelector(1)),
			observedMsgs: plugintypes.ExecutePluginMessageObservations{
				2: {1: {}},
			},
			expErr: "observer not allowed to read from chain 2",
		},
		{
			name:         "NoMessagesObserved",
			observerCfg:  mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
			observedMsgs: plugintypes.ExecutePluginMessageObservations{},
		},
		{
			name:        "EmptyMessagesInChain",
			observerCfg: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
			observedMsgs: plugintypes.ExecutePluginMessageObservations{
				1: {},
				2: {1: {}, 2: {}},
			},
		},
		{
			name:        "AllMessagesEmpty",
			observerCfg: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
			observedMsgs: plugintypes.ExecutePluginMessageObservations{
				1: {},
				2: {},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObserverReadingEligibility(tc.observerCfg, tc.observedMsgs)
			if len(tc.expErr) != 0 {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tc.expErr)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_validateObservedSequenceNumbers(t *testing.T) {
	testCases := []struct {
		name         string
		observedData map[cciptypes.ChainSelector][]plugintypes.ExecutePluginCommitDataWithMessages
		expErr       bool
	}{
		{
			name: "ValidData",
			observedData: map[cciptypes.ChainSelector][]plugintypes.ExecutePluginCommitDataWithMessages{
				1: {
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							MerkleRoot:          cciptypes.Bytes32{1},
							SequenceNumberRange: cciptypes.SeqNumRange{1, 10},
							ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3},
						},
					},
				},
				2: {
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							MerkleRoot:          cciptypes.Bytes32{2},
							SequenceNumberRange: cciptypes.SeqNumRange{11, 20},
							ExecutedMessages:    []cciptypes.SeqNum{11, 12, 13},
						},
					},
				},
			},
		},
		{
			name: "DuplicateMerkleRoot",
			observedData: map[cciptypes.ChainSelector][]plugintypes.ExecutePluginCommitDataWithMessages{
				1: {
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							MerkleRoot:          cciptypes.Bytes32{1},
							SequenceNumberRange: cciptypes.SeqNumRange{1, 10},
							ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3},
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							MerkleRoot:          cciptypes.Bytes32{1},
							SequenceNumberRange: cciptypes.SeqNumRange{11, 20},
							ExecutedMessages:    []cciptypes.SeqNum{11, 12, 13},
						},
					},
				},
			},
			expErr: true,
		},
		{
			name: "OverlappingSequenceNumberRange",
			observedData: map[cciptypes.ChainSelector][]plugintypes.ExecutePluginCommitDataWithMessages{
				1: {
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							MerkleRoot:          cciptypes.Bytes32{1},
							SequenceNumberRange: cciptypes.SeqNumRange{1, 10},
							ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3},
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							MerkleRoot:          cciptypes.Bytes32{2},
							SequenceNumberRange: cciptypes.SeqNumRange{5, 15},
							ExecutedMessages:    []cciptypes.SeqNum{6, 7, 8},
						},
					},
				},
			},
			expErr: true,
		},
		{
			name: "ExecutedMessageOutsideObservedRange",
			observedData: map[cciptypes.ChainSelector][]plugintypes.ExecutePluginCommitDataWithMessages{
				1: {
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							MerkleRoot:          cciptypes.Bytes32{1},
							SequenceNumberRange: cciptypes.SeqNumRange{1, 10},
							ExecutedMessages:    []cciptypes.SeqNum{1, 2, 11},
						},
					},
				},
			},
			expErr: true,
		},
		{
			name: "NoCommitData",
			observedData: map[cciptypes.ChainSelector][]plugintypes.ExecutePluginCommitDataWithMessages{
				1: {},
			},
		},
		{
			name:         "EmptyObservedData",
			observedData: map[cciptypes.ChainSelector][]plugintypes.ExecutePluginCommitDataWithMessages{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObservedSequenceNumbers(tc.observedData)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_computeRanges(t *testing.T) {
	type args struct {
		reports []plugintypes.ExecutePluginCommitDataWithMessages
	}

	tests := []struct {
		name string
		args args
		want []cciptypes.SeqNumRange
		err  error
	}{
		{
			name: "empty",
			args: args{reports: []plugintypes.ExecutePluginCommitDataWithMessages{}},
			want: nil,
		},
		{
			name: "overlapping ranges",
			args: args{reports: []plugintypes.ExecutePluginCommitDataWithMessages{
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(15, 25),
					},
				},
			},
			},
			err: errOverlappingRanges,
		},
		{
			name: "simple ranges collapsed",
			args: args{reports: []plugintypes.ExecutePluginCommitDataWithMessages{
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(21, 40),
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(41, 60),
					},
				},
			},
			},
			want: []cciptypes.SeqNumRange{{10, 60}},
		},
		{
			name: "non-contiguous ranges",
			args: args{reports: []plugintypes.ExecutePluginCommitDataWithMessages{
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
				},
			},
			},
			want: []cciptypes.SeqNumRange{{10, 20}, {30, 40}, {50, 60}},
		},
		{
			name: "contiguous and non-contiguous ranges",
			args: args{reports: []plugintypes.ExecutePluginCommitDataWithMessages{
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(21, 40),
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
					},
				},
			},
			},
			want: []cciptypes.SeqNumRange{{10, 40}, {50, 60}},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := computeRanges(tt.args.reports)
			if tt.err != nil {
				assert.ErrorIs(t, err, tt.err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_groupByChainSelector(t *testing.T) {
	type args struct {
		reports []plugintypes.CommitPluginReportWithMeta
	}
	tests := []struct {
		name string
		args args
		want plugintypes.ExecutePluginCommitObservations
	}{
		{
			name: "empty",
			args: args{reports: []plugintypes.CommitPluginReportWithMeta{}},
			want: plugintypes.ExecutePluginCommitObservations{},
		},
		{
			name: "reports",
			args: args{reports: []plugintypes.CommitPluginReportWithMeta{{
				Report: cciptypes.CommitPluginReport{
					MerkleRoots: []cciptypes.MerkleRootChain{
						{SourceChainSelector: 1, Interval: cciptypes.NewSeqNumRange(10, 20), MerkleRoot: cciptypes.Bytes32{1}},
						{SourceChainSelector: 2, Interval: cciptypes.NewSeqNumRange(30, 40), MerkleRoot: cciptypes.Bytes32{2}},
					}}}}},
			want: plugintypes.ExecutePluginCommitObservations{
				1: {
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SourceChain:         1,
							MerkleRoot:          cciptypes.Bytes32{1},
							SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
							ExecutedMessages:    nil,
						},
					},
				},
				2: {
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SourceChain:         2,
							MerkleRoot:          cciptypes.Bytes32{2},
							SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
							ExecutedMessages:    nil,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, tt.want, groupByChainSelector(tt.args.reports), "groupByChainSelector(%v)", tt.args.reports)
		})
	}
}

func Test_filterOutFullyExecutedMessages(t *testing.T) {
	type args struct {
		reports          []plugintypes.ExecutePluginCommitDataWithMessages
		executedMessages []cciptypes.SeqNumRange
	}
	tests := []struct {
		name    string
		args    args
		want    []plugintypes.ExecutePluginCommitDataWithMessages
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "empty",
			args: args{
				reports:          nil,
				executedMessages: nil,
			},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name: "empty2",
			args: args{
				reports:          []plugintypes.ExecutePluginCommitDataWithMessages{},
				executedMessages: nil,
			},
			want:    []plugintypes.ExecutePluginCommitDataWithMessages{},
			wantErr: assert.NoError,
		},
		{
			name: "no executed messages",
			args: args{
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
						},
					},
				},
				executedMessages: nil,
			},
			want: []plugintypes.ExecutePluginCommitDataWithMessages{
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)}},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)}},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)}},
			},
			wantErr: assert.NoError,
		},
		{
			name: "executed messages",
			args: args{
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)}},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)}},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)}},
				},
				executedMessages: []cciptypes.SeqNumRange{
					cciptypes.NewSeqNumRange(0, 100),
				},
			},
			want:    nil,
			wantErr: assert.NoError,
		},
		{
			name: "2 partially executed",
			args: args{
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
					},
				},
				executedMessages: []cciptypes.SeqNumRange{
					cciptypes.NewSeqNumRange(15, 35),
				},
			},
			want: []plugintypes.ExecutePluginCommitDataWithMessages{
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
						ExecutedMessages:    []cciptypes.SeqNum{15, 16, 17, 18, 19, 20},
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
						ExecutedMessages:    []cciptypes.SeqNum{30, 31, 32, 33, 34, 35},
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "2 partially executed 1 fully executed",
			args: args{
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
						},
					},
				},
				executedMessages: []cciptypes.SeqNumRange{
					cciptypes.NewSeqNumRange(15, 55),
				},
			},
			want: []plugintypes.ExecutePluginCommitDataWithMessages{
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
						ExecutedMessages:    []cciptypes.SeqNum{15, 16, 17, 18, 19, 20},
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
						ExecutedMessages:    []cciptypes.SeqNum{50, 51, 52, 53, 54, 55},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "first report executed",
			args: args{
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
						},
					},
				},
				executedMessages: []cciptypes.SeqNumRange{
					cciptypes.NewSeqNumRange(10, 20),
				},
			},
			want: []plugintypes.ExecutePluginCommitDataWithMessages{
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "last report executed",
			args: args{
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
						},
					},
				},
				executedMessages: []cciptypes.SeqNumRange{
					cciptypes.NewSeqNumRange(50, 60),
				},
			},
			want: []plugintypes.ExecutePluginCommitDataWithMessages{
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "sort-report",
			args: args{
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
						},
					},
				},
				executedMessages: nil,
			},
			want: []plugintypes.ExecutePluginCommitDataWithMessages{
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
					},
				},
				{
					ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
						SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "sort-executed",
			args: args{
				reports: []plugintypes.ExecutePluginCommitDataWithMessages{
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
						},
					},
					{
						ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{
							SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
						},
					},
				},
				executedMessages: []cciptypes.SeqNumRange{
					cciptypes.NewSeqNumRange(50, 60),
					cciptypes.NewSeqNumRange(10, 20),
					cciptypes.NewSeqNumRange(30, 40),
				},
			},
			want:    nil,
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := filterOutExecutedMessages(tt.args.reports, tt.args.executedMessages)
			if !tt.wantErr(t, err, fmt.Sprintf("filterOutExecutedMessages(%v, %v)", tt.args.reports, tt.args.executedMessages)) {
				return
			}
			assert.Equalf(t, tt.want, got, "filterOutExecutedMessages(%v, %v)", tt.args.reports, tt.args.executedMessages)
		})
	}
}

func Test_decodeAttributedObservations(t *testing.T) {
	mustEncode := func(obs plugintypes.ExecutePluginObservation) []byte {
		enc, err := obs.Encode()
		if err != nil {
			t.Fatal("Unable to encode")
		}
		return enc
	}
	tests := []struct {
		name    string
		args    []types.AttributedObservation
		want    []decodedAttributedObservation
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			args:    nil,
			want:    []decodedAttributedObservation{},
			wantErr: assert.NoError,
		},
		{
			name: "one observation",
			args: []types.AttributedObservation{
				{
					Observer: commontypes.OracleID(1),
					Observation: mustEncode(plugintypes.ExecutePluginObservation{
						CommitReports: plugintypes.ExecutePluginCommitObservations{
							1: {{ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{MerkleRoot: cciptypes.Bytes32{1}}}},
						},
					}),
				},
			},
			want: []decodedAttributedObservation{
				{
					Observer: commontypes.OracleID(1),
					Observation: plugintypes.ExecutePluginObservation{
						CommitReports: plugintypes.ExecutePluginCommitObservations{
							1: {{ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{MerkleRoot: cciptypes.Bytes32{1}}}},
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "multiple observations",
			args: []types.AttributedObservation{
				{
					Observer: commontypes.OracleID(1),
					Observation: mustEncode(plugintypes.ExecutePluginObservation{
						CommitReports: plugintypes.ExecutePluginCommitObservations{
							1: {{ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{MerkleRoot: cciptypes.Bytes32{1}}}},
						},
					}),
				},
				{
					Observer: commontypes.OracleID(2),
					Observation: mustEncode(plugintypes.ExecutePluginObservation{
						CommitReports: plugintypes.ExecutePluginCommitObservations{
							2: {{ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{MerkleRoot: cciptypes.Bytes32{2}}}},
						},
					}),
				},
			},
			want: []decodedAttributedObservation{
				{
					Observer: commontypes.OracleID(1),
					Observation: plugintypes.ExecutePluginObservation{
						CommitReports: plugintypes.ExecutePluginCommitObservations{
							1: {{ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{MerkleRoot: cciptypes.Bytes32{1}}}},
						},
					},
				},
				{
					Observer: commontypes.OracleID(2),
					Observation: plugintypes.ExecutePluginObservation{
						CommitReports: plugintypes.ExecutePluginCommitObservations{
							2: {{ExecutePluginCommitData: plugintypes.ExecutePluginCommitData{MerkleRoot: cciptypes.Bytes32{2}}}},
						},
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "invalid observation",
			args: []types.AttributedObservation{
				{
					Observer:    commontypes.OracleID(1),
					Observation: []byte("invalid"),
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeAttributedObservations(tt.args)
			if !tt.wantErr(t, err, fmt.Sprintf("decodeAttributedObservations(%v)", tt.args)) {
				return
			}
			assert.Equalf(t, tt.want, got, "decodeAttributedObservations(%v)", tt.args)
		})
	}
}
