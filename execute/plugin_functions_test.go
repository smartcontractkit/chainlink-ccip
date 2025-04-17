package execute

import (
	"fmt"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_validateObserverReadingEligibility(t *testing.T) {
	tests := []struct {
		name         string
		observerCfg  mapset.Set[cciptypes.ChainSelector]
		observedMsgs exectypes.MessageObservations
		expErr       string
	}{
		{
			name:        "ValidObserverAndMessages",
			observerCfg: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
			observedMsgs: exectypes.MessageObservations{
				1: {1: {}, 2: {}},
				2: {},
			},
		},
		{
			name:        "ObserverNotAllowedToReadChain",
			observerCfg: mapset.NewSet(cciptypes.ChainSelector(1)),
			observedMsgs: exectypes.MessageObservations{
				2: {1: {}},
			},
			expErr: "observer not allowed to read from chain 2",
		},
		{
			name:         "NoMessagesObserved",
			observerCfg:  mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
			observedMsgs: exectypes.MessageObservations{},
		},
		{
			name:        "EmptyMessagesInChain",
			observerCfg: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
			observedMsgs: exectypes.MessageObservations{
				1: {},
				2: {1: {}, 2: {}},
			},
		},
		{
			name:        "AllMessagesEmpty",
			observerCfg: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
			observedMsgs: exectypes.MessageObservations{
				1: {},
				2: {},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateMsgsReadingEligibility(tc.observerCfg, tc.observedMsgs)
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
		name            string
		observedData    map[cciptypes.ChainSelector][]exectypes.CommitData
		supportedChains mapset.Set[cciptypes.ChainSelector]
		expErr          bool
	}{
		{
			name: "ValidData",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.SeqNumRange{1, 3},
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3},
					},
				},
				2: {
					{
						MerkleRoot:          cciptypes.Bytes32{2},
						SequenceNumberRange: cciptypes.SeqNumRange{11, 15},
						ExecutedMessages:    []cciptypes.SeqNum{11, 12, 13},
					},
				},
			},
			supportedChains: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
		},
		{
			name: "UnsupportedChain",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.SeqNumRange{1, 3},
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3},
					},
				},
				2: {
					{
						MerkleRoot:          cciptypes.Bytes32{2},
						SequenceNumberRange: cciptypes.SeqNumRange{11, 15},
						ExecutedMessages:    []cciptypes.SeqNum{11, 12, 13},
					},
				},
			},
			supportedChains: mapset.NewSet(cciptypes.ChainSelector(1)), // <-- 2 is missing
			expErr:          true,
		},
		{
			name: "DuplicateMerkleRoot",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.SeqNumRange{1, 10},
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3},
					},
					{
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.SeqNumRange{11, 20},
						ExecutedMessages:    []cciptypes.SeqNum{11, 12, 13},
					},
				},
			},
			supportedChains: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
			expErr:          true,
		},
		{
			name: "OverlappingSequenceNumberRange",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.SeqNumRange{1, 10},
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3},
					},
					{
						MerkleRoot:          cciptypes.Bytes32{2},
						SequenceNumberRange: cciptypes.SeqNumRange{5, 15},
						ExecutedMessages:    []cciptypes.SeqNum{6, 7, 8},
					},
				},
			},
			supportedChains: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
			expErr:          true,
		},
		{
			name: "ExecutedMessageOutsideObservedRange",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.SeqNumRange{1, 10},
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 11},
					},
				},
			},
			supportedChains: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
			expErr:          true,
		},
		{
			name: "NoCommitData",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {},
			},
			supportedChains: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
		},
		{
			name:            "EmptyObservedData",
			observedData:    map[cciptypes.ChainSelector][]exectypes.CommitData{},
			supportedChains: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateObservedSequenceNumbers(tc.supportedChains, tc.observedData)
			if tc.expErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_validateMessagesConformToCommitReports(t *testing.T) {
	testCases := []struct {
		name         string
		observedData map[cciptypes.ChainSelector][]exectypes.CommitData
		observedMsgs exectypes.MessageObservations
		expErr       bool
	}{
		{
			name: "NoCommitData",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {},
			},
			expErr: true,
		},
		{
			name:         "EmptyObservedData",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{},
		},
		// Tests with messages
		{
			name: "Gap in Sequence Numbers",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.SeqNumRange{1, 10},
						ExecutedMessages:    []cciptypes.SeqNum{1, 2},
						SourceChain:         1,
					},
				},
			},
			observedMsgs: exectypes.MessageObservations{
				1: emptyMessagesMapForRanges([]cciptypes.SeqNumRange{{1, 2}, {5, 10}}),
			},
			expErr: true,
		},
		{
			name: "valid multiple commit reports for multiple chains",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.SeqNumRange{1, 3},
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3},
					},
				},
				2: {
					{
						MerkleRoot:          cciptypes.Bytes32{2},
						SequenceNumberRange: cciptypes.SeqNumRange{11, 15},
						ExecutedMessages:    []cciptypes.SeqNum{11, 12, 13},
					},
				},
			},
			observedMsgs: exectypes.MessageObservations{
				1: emptyMessagesMapForRange(1, 3),
				2: emptyMessagesMapForRange(11, 15),
			},
		},
		{
			name: "valid multiple commit reports for same chain",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.SeqNumRange{1, 3},
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3},
					},
					{
						MerkleRoot:          cciptypes.Bytes32{2},
						SequenceNumberRange: cciptypes.SeqNumRange{4, 6},
					},
					{
						MerkleRoot:          cciptypes.Bytes32{3},
						SequenceNumberRange: cciptypes.SeqNumRange{8, 10},
					},
				},
			},
			observedMsgs: exectypes.MessageObservations{
				1: emptyMessagesMapForRanges([]cciptypes.SeqNumRange{{1, 3}, {4, 6}, {8, 10}}),
			},
		},
		{
			name: "Extra Sequence Numbers",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.SeqNumRange{1, 3},
					},
				},
			},
			observedMsgs: exectypes.MessageObservations{
				1: emptyMessagesMapForRange(1, 4),
			},
			expErr: true,
		},
		{
			name: "Missing Sequence Numbers",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.SeqNumRange{1, 3},
					},
				},
			},
			observedMsgs: exectypes.MessageObservations{
				1: emptyMessagesMapForRange(1, 2),
			},
			expErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := validateMessagesConformToCommitReports(tc.observedData, tc.observedMsgs)
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
		reports []exectypes.CommitData
	}

	tests := []struct {
		name string
		args args
		want []cciptypes.SeqNumRange
		err  error
	}{
		{
			name: "empty",
			args: args{reports: []exectypes.CommitData{}},
			want: nil,
		},
		{
			name: "overlapping ranges",
			args: args{reports: []exectypes.CommitData{
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(15, 25),
				},
			}},
			err: errOverlappingRanges,
		},
		{
			name: "simple ranges collapsed",
			args: args{reports: []exectypes.CommitData{
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(21, 40),
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(41, 60),
				},
			}},
			want: []cciptypes.SeqNumRange{{10, 60}},
		},
		{
			name: "non-contiguous ranges",
			args: args{reports: []exectypes.CommitData{
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
			}},
			want: []cciptypes.SeqNumRange{{10, 20}, {30, 40}, {50, 60}},
		},
		{
			name: "contiguous and non-contiguous ranges",
			args: args{reports: []exectypes.CommitData{
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(21, 40),
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
				},
			}},
			want: []cciptypes.SeqNumRange{{10, 40}, {50, 60}},
		},
		{
			name: "contiguous and non-contiguous ranges",
			args: args{reports: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 12)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(13, 15)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(16, 20)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(22, 33)},
			}},
			want: []cciptypes.SeqNumRange{{10, 20}, {22, 33}},
		},
		{
			name: "overlap on range bound",
			args: args{reports: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 12)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(13, 15)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(15, 20)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(22, 33)},
			}},
			err: errOverlappingRanges,
		},
	}
	for _, tt := range tests {
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

func Test_groupByChainSelectorWithFilter(t *testing.T) {
	type args struct {
		reports            []cciptypes.CommitPluginReportWithMeta
		cursedSourceChains map[cciptypes.ChainSelector]bool
	}
	tests := []struct {
		name string
		args args
		want exectypes.CommitObservations
	}{
		{
			name: "empty",
			args: args{
				reports:            []cciptypes.CommitPluginReportWithMeta{},
				cursedSourceChains: nil,
			},
			want: exectypes.CommitObservations{},
		},
		{
			name: "reports with no cursed chains",
			args: args{
				reports: []cciptypes.CommitPluginReportWithMeta{{
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{ChainSel: 1, SeqNumsRange: cciptypes.NewSeqNumRange(10, 20), MerkleRoot: cciptypes.Bytes32{1}},
							{ChainSel: 2, SeqNumsRange: cciptypes.NewSeqNumRange(30, 40), MerkleRoot: cciptypes.Bytes32{2}},
						}}}},
				cursedSourceChains: map[cciptypes.ChainSelector]bool{},
			},
			want: exectypes.CommitObservations{
				1: {
					{
						SourceChain:         1,
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					},
				},
				2: {
					{
						SourceChain:         2,
						MerkleRoot:          cciptypes.Bytes32{2},
						SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
					},
				},
			},
		},
		{
			name: "reports with cursed chain 1",
			args: args{
				reports: []cciptypes.CommitPluginReportWithMeta{{
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{ChainSel: 1, SeqNumsRange: cciptypes.NewSeqNumRange(10, 20), MerkleRoot: cciptypes.Bytes32{1}},
							{ChainSel: 2, SeqNumsRange: cciptypes.NewSeqNumRange(30, 40), MerkleRoot: cciptypes.Bytes32{2}},
						}}}},
				cursedSourceChains: map[cciptypes.ChainSelector]bool{1: true},
			},
			want: exectypes.CommitObservations{
				2: {
					{
						SourceChain:         2,
						MerkleRoot:          cciptypes.Bytes32{2},
						SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
					},
				},
			},
		},
		{
			name: "reports with all chains cursed",
			args: args{
				reports: []cciptypes.CommitPluginReportWithMeta{{
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{ChainSel: 1, SeqNumsRange: cciptypes.NewSeqNumRange(10, 20), MerkleRoot: cciptypes.Bytes32{1}},
							{ChainSel: 2, SeqNumsRange: cciptypes.NewSeqNumRange(30, 40), MerkleRoot: cciptypes.Bytes32{2}},
						}}}},
				cursedSourceChains: map[cciptypes.ChainSelector]bool{1: true, 2: true},
			},
			want: exectypes.CommitObservations{},
		},
		{
			name: "reports with blessed and unblessed merkle roots",
			args: args{
				reports: []cciptypes.CommitPluginReportWithMeta{{
					Report: cciptypes.CommitPluginReport{
						BlessedMerkleRoots: []cciptypes.MerkleRootChain{
							{ChainSel: 1, SeqNumsRange: cciptypes.NewSeqNumRange(10, 20), MerkleRoot: cciptypes.Bytes32{1}},
						},
						UnblessedMerkleRoots: []cciptypes.MerkleRootChain{
							{ChainSel: 2, SeqNumsRange: cciptypes.NewSeqNumRange(30, 40), MerkleRoot: cciptypes.Bytes32{2}},
						}}}},
				cursedSourceChains: map[cciptypes.ChainSelector]bool{1: true},
			},
			want: exectypes.CommitObservations{
				2: {
					{
						SourceChain:         2,
						MerkleRoot:          cciptypes.Bytes32{2},
						SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
					},
				},
			},
		},
		{
			name: "multiple reports with some cursed chains",
			args: args{
				reports: []cciptypes.CommitPluginReportWithMeta{
					{
						Report: cciptypes.CommitPluginReport{
							BlessedMerkleRoots: []cciptypes.MerkleRootChain{
								{ChainSel: 1, SeqNumsRange: cciptypes.NewSeqNumRange(10, 20), MerkleRoot: cciptypes.Bytes32{1}},
							},
						},
					},
					{
						Report: cciptypes.CommitPluginReport{
							BlessedMerkleRoots: []cciptypes.MerkleRootChain{
								{ChainSel: 2, SeqNumsRange: cciptypes.NewSeqNumRange(30, 40), MerkleRoot: cciptypes.Bytes32{2}},
								{ChainSel: 3, SeqNumsRange: cciptypes.NewSeqNumRange(50, 60), MerkleRoot: cciptypes.Bytes32{3}},
							},
						},
					},
				},
				cursedSourceChains: map[cciptypes.ChainSelector]bool{2: true},
			},
			want: exectypes.CommitObservations{
				1: {
					{
						SourceChain:         1,
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					},
				},
				3: {
					{
						SourceChain:         3,
						MerkleRoot:          cciptypes.Bytes32{3},
						SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			lggr := logger.Test(t)
			assert.Equalf(t, tt.want, groupByChainSelectorWithFilter(lggr, tt.args.reports, tt.args.cursedSourceChains),
				"groupByChainSelectorWithFilter(%v, %v)", tt.args.reports, tt.args.cursedSourceChains)
		})
	}
}

func Test_combineReportsAndMessages(t *testing.T) {
	type args struct {
		reports          []exectypes.CommitData
		executedMessages []cciptypes.SeqNum
	}
	tests := []struct {
		name         string
		args         args
		wantPending  []exectypes.CommitData
		wantExecuted []exectypes.CommitData
	}{
		{
			name: "empty",
			args: args{
				reports:          nil,
				executedMessages: nil,
			},
			wantPending: nil,
		},
		{
			name: "empty2",
			args: args{
				reports:          []exectypes.CommitData{},
				executedMessages: nil,
			},
			wantPending: []exectypes.CommitData{},
		},
		{
			name: "no executed messages",
			args: args{
				reports: []exectypes.CommitData{
					{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
				},
				executedMessages: nil,
			},
			wantPending: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
			},
		},
		{
			name: "executed messages",
			args: args{
				reports: []exectypes.CommitData{
					{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
				},
				executedMessages: cciptypes.NewSeqNumRange(0, 100).ToSlice(),
			},
			wantPending: nil,
			wantExecuted: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
			},
		},
		{
			name: "2 partially executed",
			args: args{
				reports: []exectypes.CommitData{
					{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
				},
				executedMessages: cciptypes.NewSeqNumRange(15, 35).ToSlice(),
			},
			wantPending: []exectypes.CommitData{
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					ExecutedMessages:    []cciptypes.SeqNum{15, 16, 17, 18, 19, 20},
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
					ExecutedMessages:    []cciptypes.SeqNum{30, 31, 32, 33, 34, 35},
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
				},
			},
		},
		{
			name: "2 partially executed 1 fully executed",
			args: args{
				reports: []exectypes.CommitData{
					{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
				},
				executedMessages: cciptypes.NewSeqNumRange(15, 55).ToSlice(),
			},
			wantPending: []exectypes.CommitData{
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					ExecutedMessages:    []cciptypes.SeqNum{15, 16, 17, 18, 19, 20},
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
					ExecutedMessages:    []cciptypes.SeqNum{50, 51, 52, 53, 54, 55},
				},
			},
			wantExecuted: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
			},
		},
		{
			name: "first report executed",
			args: args{
				reports: []exectypes.CommitData{
					{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
				},
				executedMessages: cciptypes.NewSeqNumRange(10, 20).ToSlice(),
			},
			wantPending: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
			},
			wantExecuted: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
			},
		},
		{
			name: "last report executed",
			args: args{
				reports: []exectypes.CommitData{
					{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
				},
				executedMessages: cciptypes.NewSeqNumRange(50, 60).ToSlice(),
			},
			wantPending: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
			},
			wantExecuted: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
			},
		},
		{
			name: "sort-executed",
			args: args{
				reports: []exectypes.CommitData{
					{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
				},
				executedMessages: cciptypes.NewSeqNumRange(10, 60).ToSlice(),
			},
			wantPending: nil,
			wantExecuted: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := combineReportsAndMessages(tt.args.reports, tt.args.executedMessages)
			assert.Equal(t, tt.wantPending, got)
			assert.Equal(t, tt.wantExecuted, got2)
		})
	}
}

func Test_decodeAttributedObservations(t *testing.T) {
	mustEncode := func(obs exectypes.Observation) []byte {
		enc, err := ocrTypeCodec.EncodeObservation(obs)
		if err != nil {
			t.Fatal("Unable to encode")
		}
		return enc
	}
	tests := []struct {
		name    string
		args    []types.AttributedObservation
		want    []plugincommon.AttributedObservation[exectypes.Observation]
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "empty",
			args:    nil,
			want:    []plugincommon.AttributedObservation[exectypes.Observation]{},
			wantErr: assert.NoError,
		},
		{
			name: "one observation",
			args: []types.AttributedObservation{
				{
					Observer: commontypes.OracleID(1),
					Observation: mustEncode(exectypes.Observation{
						CommitReports: exectypes.CommitObservations{
							1: {{MerkleRoot: cciptypes.Bytes32{1}}},
						},
					}),
				},
			},
			want: []plugincommon.AttributedObservation[exectypes.Observation]{
				{
					OracleID: commontypes.OracleID(1),
					Observation: exectypes.Observation{
						CommitReports: exectypes.CommitObservations{
							1: {{MerkleRoot: cciptypes.Bytes32{1}}},
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
					Observation: mustEncode(exectypes.Observation{
						CommitReports: exectypes.CommitObservations{
							1: {{MerkleRoot: cciptypes.Bytes32{1}}},
						},
					}),
				},
				{
					Observer: commontypes.OracleID(2),
					Observation: mustEncode(exectypes.Observation{
						CommitReports: exectypes.CommitObservations{
							2: {{MerkleRoot: cciptypes.Bytes32{2}}},
						},
					}),
				},
			},
			want: []plugincommon.AttributedObservation[exectypes.Observation]{
				{
					OracleID: commontypes.OracleID(1),
					Observation: exectypes.Observation{
						CommitReports: exectypes.CommitObservations{
							1: {{MerkleRoot: cciptypes.Bytes32{1}}},
						},
					},
				},
				{
					OracleID: commontypes.OracleID(2),
					Observation: exectypes.Observation{
						CommitReports: exectypes.CommitObservations{
							2: {{MerkleRoot: cciptypes.Bytes32{2}}},
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
			got, err := decodeAttributedObservations(tt.args, ocrTypeCodec)
			if !tt.wantErr(t, err, fmt.Sprintf("decodeAttributedObservations(%v)", tt.args)) {
				return
			}
			assert.Equalf(t, tt.want, got, "decodeAttributedObservations(%v)", tt.args)
		})
	}
}

func Test_computeConsensusObservation(t *testing.T) {
	type args struct {
		observation []exectypes.Observation
		F           int
	}
	dstChain := cciptypes.ChainSelector(1)
	defaultFChain := map[cciptypes.ChainSelector]int{
		dstChain: 1,
	}
	tests := []struct {
		name    string
		args    args
		want    exectypes.Observation
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "empty",
			args: args{
				observation: nil,
			},
			want:    exectypes.Observation{},
			wantErr: assert.Error,
		},
		{
			name: "one consensus observation",
			args: args{
				observation: []exectypes.Observation{
					{
						Nonces: exectypes.NonceObservations{
							dstChain: {
								"0x1": 1,
							},
						},
						FChain: map[cciptypes.ChainSelector]int{
							dstChain: 0,
						},
					},
				},
			},
			want: exectypes.Observation{
				Nonces: exectypes.NonceObservations{
					1: {
						"0x1": 1,
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "consensus when exactly 2f+1",
			args: args{
				observation: []exectypes.Observation{
					{
						Nonces: exectypes.NonceObservations{dstChain: {"0x1": 1}},
						FChain: map[cciptypes.ChainSelector]int{dstChain: 1},
					},
					{
						Nonces: exectypes.NonceObservations{dstChain: {"0x1": 1}},
						FChain: map[cciptypes.ChainSelector]int{dstChain: 1},
					},
					{
						Nonces: exectypes.NonceObservations{dstChain: {"0x1": 1}},
						FChain: map[cciptypes.ChainSelector]int{dstChain: 1},
					},
				},
			},
			want: exectypes.Observation{
				Nonces: exectypes.NonceObservations{
					1: {
						"0x1": 1,
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "no consensus when less than f+1",
			args: args{
				observation: []exectypes.Observation{
					{
						Nonces: exectypes.NonceObservations{dstChain: {"0x1": 1}},
						FChain: map[cciptypes.ChainSelector]int{dstChain: 2},
					},
					{
						Nonces: exectypes.NonceObservations{dstChain: {"0x1": 1}},
						FChain: map[cciptypes.ChainSelector]int{dstChain: 2},
					},
				},
			},
			want:    exectypes.Observation{Nonces: exectypes.NonceObservations{}},
			wantErr: assert.NoError,
		},
		{
			name: "one ignored consensus observation",
			args: args{

				observation: []exectypes.Observation{
					{
						Nonces: exectypes.NonceObservations{
							1: {
								"0x1": 1,
							},
						},
						FChain: map[cciptypes.ChainSelector]int{
							1: 1,
						}},
				},
			},
			want:    exectypes.Observation{Nonces: exectypes.NonceObservations{}},
			wantErr: assert.NoError,
		},
		{
			name: "3 observers required to reach consensus on 4 sender values (with 2f+1)",
			args: args{
				// Across 3 observers, need all 3 observers to agree for consensus (2f+1 = 3)
				observation: []exectypes.Observation{
					{
						Nonces: exectypes.NonceObservations{
							1: {
								"0x1": 1,
								"0x2": 2,
								"0x3": 3,
								"0x4": 4,
							},
						},
						FChain: defaultFChain,
					}, {
						Nonces: exectypes.NonceObservations{
							1: {
								"0x1": 1,
								"0x2": 2,
								"0x3": 3,
								"0x4": 4,
							},
						},
						FChain: defaultFChain,
					}, {
						Nonces: exectypes.NonceObservations{
							1: {
								"0x1": 1,
								"0x2": 2,
								"0x3": 3,
								"0x4": 4,
							},
						},
						FChain: defaultFChain,
					},
				},
			},
			want: exectypes.Observation{
				Nonces: exectypes.NonceObservations{
					1: {
						"0x1": 1,
						"0x2": 2,
						"0x3": 3,
						"0x4": 4,
					},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "3 observers but different nonce values. No consensus.",
			args: args{
				// Across 3 observers
				observation: []exectypes.Observation{
					{
						//
						Nonces: exectypes.NonceObservations{
							1: {
								"0x1": 9,
								"0x2": 9,
								"0x3": 9,
								"0x4": 9,
							},
						},
						FChain: map[cciptypes.ChainSelector]int{
							1: 2,
						},
					}, {
						Nonces: exectypes.NonceObservations{
							1: {
								"0x1": 1,
								"0x4": 4,
							},
						},
						FChain: map[cciptypes.ChainSelector]int{
							1: 2,
						},
					}, {
						Nonces: exectypes.NonceObservations{
							1: {
								"0x2": 2,
								"0x3": 3,
							},
						},
						FChain: map[cciptypes.ChainSelector]int{
							1: 2,
						},
					},
				},
			},
			want:    exectypes.Observation{Nonces: exectypes.NonceObservations{}},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Convert observations to the expected decoded type.
			var ao []plugincommon.AttributedObservation[exectypes.Observation]
			for i, observation := range tt.args.observation {
				ao = append(ao, plugincommon.AttributedObservation[exectypes.Observation]{
					Observation: observation,
					OracleID:    commontypes.OracleID(i),
				})
			}

			lggr := logger.Test(t)
			got, err := computeConsensusObservation(lggr, ao, 1, tt.args.F)
			if !tt.wantErr(t, err, "getConsensusObservation(...)") {
				return
			}
			assert.Equalf(t, tt.want, got, "getConsensusObservation(...)")
		})
	}
}

func Test_computeTokenDataObservationsConsensus(t *testing.T) {
	chainSelector := cciptypes.ChainSelector(1)

	type expected struct {
		ready bool
		data  [][]byte
	}

	tt := []struct {
		name        string
		F           int
		observation []map[cciptypes.SeqNum]exectypes.MessageTokenData
		expected    map[cciptypes.SeqNum]expected
	}{
		{
			name: "messages without token data",
			F:    1,
			observation: []map[cciptypes.SeqNum]exectypes.MessageTokenData{
				{
					1: exectypes.NewMessageTokenData(),
					2: exectypes.NewMessageTokenData(),
					3: exectypes.NewMessageTokenData(),
				},
				{
					1: exectypes.NewMessageTokenData(),
					2: exectypes.NewMessageTokenData(),
					3: exectypes.NewMessageTokenData(),
				},
			},
			expected: map[cciptypes.SeqNum]expected{
				1: {ready: true, data: [][]byte{}},
				2: {ready: true, data: [][]byte{}},
				3: {ready: true, data: [][]byte{}},
			},
		},
		{
			name: "messages with empty token data",
			F:    1,
			observation: []map[cciptypes.SeqNum]exectypes.MessageTokenData{
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
					),
				},
			},
			expected: map[cciptypes.SeqNum]expected{
				1: {ready: true, data: [][]byte{{}}},
			},
		},
		{
			name: "plugins seeing completely different tokens",
			F:    1,
			observation: []map[cciptypes.SeqNum]exectypes.MessageTokenData{
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{11}),
						exectypes.NewNoopTokenData(),
					),
					2: exectypes.NewMessageTokenData(),
					3: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{31}),
					),
					5: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{51}),
						exectypes.NewSuccessTokenData([]byte{52}),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
					),
					2: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
						exectypes.NewNoopTokenData(),
					),
					3: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{31}),
						exectypes.NewSuccessTokenData([]byte{32}),
						exectypes.NewSuccessTokenData([]byte{33}),
					),
					4: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{41}),
						exectypes.NewSuccessTokenData([]byte{42}),
					),
					5: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{51}),
						exectypes.NewSuccessTokenData([]byte{52}),
					),
				},
			},
			expected: map[cciptypes.SeqNum]expected{
				1: {ready: false},
				2: {ready: false},
				3: {ready: false},
				4: {ready: false},
				5: {ready: true, data: [][]byte{{51}, {52}}},
			},
		},
		{
			name: "some tokens are not observed by one of the nodes",
			F:    1,
			observation: []map[cciptypes.SeqNum]exectypes.MessageTokenData{
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{11}),
						exectypes.NewNoopTokenData(),
						exectypes.NewSuccessTokenData([]byte{13}),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{11}),
						exectypes.NewNoopTokenData(),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{11}),
						exectypes.NewSuccessTokenData([]byte{12}),
						exectypes.NewSuccessTokenData([]byte{13}),
					),
				},
			},
			expected: map[cciptypes.SeqNum]expected{
				1: {ready: true, data: [][]byte{{11}, {}, {13}}},
			},
		},
		{
			name: "message not ready - only one token has enough observations",
			F:    2,
			observation: []map[cciptypes.SeqNum]exectypes.MessageTokenData{
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
						exectypes.NewSuccessTokenData([]byte{2}),
						exectypes.NewErrorTokenData(fmt.Errorf("error")),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{1}),
						exectypes.NewSuccessTokenData([]byte{2}),
						exectypes.NewErrorTokenData(fmt.Errorf("error")),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
						exectypes.NewSuccessTokenData([]byte{2}),
						exectypes.NewSuccessTokenData([]byte{3}),
					),
				},
			},
			expected: map[cciptypes.SeqNum]expected{
				1: {ready: false},
			},
		},
		{
			name: "message not ready - only some of the tokens have enough observations",
			F:    1,
			observation: []map[cciptypes.SeqNum]exectypes.MessageTokenData{
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
						exectypes.NewSuccessTokenData([]byte{2}),
						exectypes.NewErrorTokenData(fmt.Errorf("error1")),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{1}),
						exectypes.NewNoopTokenData(),
						exectypes.NewErrorTokenData(fmt.Errorf("error2")),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{1}),
						exectypes.NewSuccessTokenData([]byte{2}),
						exectypes.NewSuccessTokenData([]byte{3}),
					),
				},
			},
			expected: map[cciptypes.SeqNum]expected{
				1: {ready: false},
			},
		},
		{
			name: "message ready - all tokens have enough observations",
			F:    1,
			observation: []map[cciptypes.SeqNum]exectypes.MessageTokenData{
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
						exectypes.NewSuccessTokenData([]byte{2}),
						exectypes.NewErrorTokenData(fmt.Errorf("error")),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{1}),
						exectypes.NewNoopTokenData(),
						exectypes.NewSuccessTokenData([]byte{3}),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{1}),
						exectypes.NewSuccessTokenData([]byte{2}),
						exectypes.NewSuccessTokenData([]byte{3}),
					),
				},
			},
			expected: map[cciptypes.SeqNum]expected{
				1: {ready: true, data: [][]byte{{1}, {2}, {3}}},
			},
		},
		{
			name: "all messages have enough observations",
			F:    1,
			observation: []map[cciptypes.SeqNum]exectypes.MessageTokenData{
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewErrorTokenData(fmt.Errorf("error")),
					),
					2: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{90}),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{1}),
					),
					2: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{2}),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{1}),
					),
					2: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{2}),
					),
				},
			},
			expected: map[cciptypes.SeqNum]expected{
				1: {ready: true, data: [][]byte{{1}}},
				2: {ready: true, data: [][]byte{{2}}},
			},
		},
		{
			name: "consensus is not reached for some of the messages",
			F:    1,
			observation: []map[cciptypes.SeqNum]exectypes.MessageTokenData{
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{1}),
					),
					2: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{2}),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{3}),
					),
					2: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{4}),
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.NewErrorTokenData(fmt.Errorf("error")),
					),
					2: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte{2}),
					),
				},
			},
			expected: map[cciptypes.SeqNum]expected{
				1: {ready: false},
				2: {ready: true, data: [][]byte{{2}}},
			},
		},
		{
			name: "message ready - only ready and data are used for reaching consensus",
			F:    1,
			observation: []map[cciptypes.SeqNum]exectypes.MessageTokenData{
				{
					1: exectypes.NewMessageTokenData(
						exectypes.TokenData{Ready: true, Data: []byte{1}},
					),
					2: exectypes.NewMessageTokenData(
						exectypes.TokenData{Ready: true, Data: []byte{2}, Supported: true},
					),
					3: exectypes.NewMessageTokenData(
						exectypes.TokenData{Ready: true, Data: []byte{3}, Supported: false},
					),
				},
				{
					1: exectypes.NewMessageTokenData(
						exectypes.TokenData{Ready: true, Data: []byte{2}},
					),
					2: exectypes.NewMessageTokenData(
						exectypes.TokenData{Ready: true, Data: []byte{2}, Supported: false},
					),
					3: exectypes.NewMessageTokenData(
						exectypes.TokenData{Ready: true, Data: []byte{3}, Supported: false, Error: fmt.Errorf("error")},
					),
				},
			},
			expected: map[cciptypes.SeqNum]expected{
				1: {ready: false},
				2: {ready: true, data: [][]byte{{2}}},
				3: {ready: true, data: [][]byte{{3}}},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			fChain := make(map[cciptypes.ChainSelector]int)
			fChain[chainSelector] = tc.F

			var ao []plugincommon.AttributedObservation[exectypes.Observation]
			for i, observation := range tc.observation {
				formatted := make(exectypes.TokenDataObservations)
				formatted[chainSelector] = observation

				ao = append(ao, plugincommon.AttributedObservation[exectypes.Observation]{
					Observation: exectypes.Observation{
						TokenData: formatted,
					},
					OracleID: commontypes.OracleID(i),
				})
			}

			obs := computeTokenDataObservationsConsensus(logger.Test(t), ao, fChain)

			for seqNum, exp := range tc.expected {
				mtd, ok := obs[chainSelector][seqNum]
				assert.True(t, ok)

				assert.Equal(t, exp.ready, mtd.IsReady())
				// No need to compare bytes when not ready
				if exp.ready {
					assert.Equal(t, exp.data, obs[chainSelector][seqNum].ToByteSlice())
				}
			}
		})
	}
}

func Test_allSeqNrsObserved(t *testing.T) {
	tests := []struct {
		name        string
		msgs        []cciptypes.Message
		numberRange cciptypes.SeqNumRange
		want        bool
	}{
		{
			name:        "all sequence numbers observed",
			msgs:        emptyMessagesForRange(1, 3),
			numberRange: cciptypes.NewSeqNumRange(1, 3),
			want:        true,
		},
		{
			name:        "missing sequence number",
			msgs:        []cciptypes.Message{emptyMessagesForRange(1, 1)[0], emptyMessagesForRange(3, 3)[0]},
			numberRange: cciptypes.NewSeqNumRange(1, 3),
			want:        false,
		},
		{
			name:        "extra sequence number",
			msgs:        emptyMessagesForRange(1, 4),
			numberRange: cciptypes.NewSeqNumRange(1, 3),
			want:        false,
		},
		{
			name:        "empty messages",
			msgs:        []cciptypes.Message{},
			numberRange: cciptypes.NewSeqNumRange(1, 3),
			want:        false,
		},
		{
			name:        "empty range",
			msgs:        emptyMessagesForRange(1, 4),
			numberRange: cciptypes.NewSeqNumRange(0, 0),
			want:        false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := msgsConformToSeqRange(tt.msgs, tt.numberRange); got != tt.want {
				t.Errorf("msgsConformToSeqRange() = %v, wantPending %v", got, tt.want)
			}
		})
	}
}

func Test_validateCommitReportsReadingEligibility(t *testing.T) {
	tests := []struct {
		name            string
		supportedChains mapset.Set[cciptypes.ChainSelector]
		observedData    exectypes.CommitObservations
		expErr          string
	}{
		{
			name:            "ValidCommitReports",
			supportedChains: mapset.NewSet(cciptypes.ChainSelector(1), cciptypes.ChainSelector(2)),
			observedData: exectypes.CommitObservations{
				1: {
					{SourceChain: 1},
				},
				2: {
					{SourceChain: 2},
				},
			},
		},
		{
			name:            "UnsupportedChain",
			supportedChains: mapset.NewSet(cciptypes.ChainSelector(1)),
			observedData: exectypes.CommitObservations{
				2: {
					{SourceChain: 2},
				},
			},
			expErr: "observer not allowed to read from chain 2",
		},
		{
			name:            "MismatchedSourceChain",
			supportedChains: mapset.NewSet(cciptypes.ChainSelector(1)),
			observedData: exectypes.CommitObservations{
				1: {
					{SourceChain: 2},
				},
			},
			expErr: "invalid observed data, key=1 but data chain=2",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateCommitReportsReadingEligibility(tc.supportedChains, tc.observedData)
			if len(tc.expErr) != 0 {
				assert.Error(t, err)
				assert.ErrorContains(t, err, tc.expErr)
				return
			}
			assert.NoError(t, err)
		})
	}
}

func Test_computeNoncesConsensus(t *testing.T) {
	lggr := logger.Test(t)

	testCases := []struct {
		name                 string
		allNonceObservations []exectypes.NonceObservations
		fChain               int
		expNonceConsensus    exectypes.NonceObservations
	}{
		{
			name:                 "empty",
			allNonceObservations: []exectypes.NonceObservations{},
			fChain:               1,
			expNonceConsensus:    exectypes.NonceObservations{},
		},
		{
			name: "one observation does not reach threshold",
			allNonceObservations: []exectypes.NonceObservations{
				{1: {"0x1": 100}},
			},
			fChain:            1,
			expNonceConsensus: exectypes.NonceObservations{},
		},
		{
			name: "two observations reach threshold",
			allNonceObservations: []exectypes.NonceObservations{
				{1: {"0x1": 100}},
				{1: {"0x1": 100}},
			},
			fChain: 1,
			expNonceConsensus: exectypes.NonceObservations{
				1: {"0x1": 100},
			},
		},
		{
			name: "two observations reach threshold but different values",
			allNonceObservations: []exectypes.NonceObservations{
				{1: {"0x1": 100}},
				{1: {"0x1": 101}},
			},
			fChain: 1,
			expNonceConsensus: exectypes.NonceObservations{
				1: {"0x1": 101},
			},
		},
		{
			name: "multiple observations with different values unordered",
			allNonceObservations: []exectypes.NonceObservations{
				{1: {"0x1": 108}},
				{1: {"0x1": 100}},
				{1: {"0x1": 102}},
				{1: {"0x1": 103}},
				{1: {"0x1": 109}},
				{1: {"0x1": 104}},
				{1: {"0x1": 105}},
				{1: {"0x1": 106}},
				{1: {"0x1": 107}},
				{1: {"0x1": 101}},
			},
			fChain: 5,
			expNonceConsensus: exectypes.NonceObservations{
				1: {"0x1": 105},
			},
		},
		{
			name: "multiple observations deviating on two values",
			allNonceObservations: []exectypes.NonceObservations{
				{1: {"0x1": 100}},
				{1: {"0x1": 100}},
				{1: {"0x1": 100}},
				{1: {"0x1": 100}},
				{1: {"0x1": 100}},
				{1: {"0x1": 100}},
				{1: {"0x1": 101}},
				{1: {"0x1": 101}},
				{1: {"0x1": 101}},
				{1: {"0x1": 101}},
			},
			fChain: 5,
			expNonceConsensus: exectypes.NonceObservations{
				1: {"0x1": 100},
			},
		},
		{
			name: "multiple observations deviating on two values",
			allNonceObservations: []exectypes.NonceObservations{
				{1: {"0x1": 100}},
				{1: {"0x1": 100}},
				{1: {"0x1": 100}},
				{1: {"0x1": 100}},
				{1: {"0x1": 100}},
				{1: {"0x1": 100}},
				{1: {"0x1": 101}},
				{1: {"0x1": 101}},
				{1: {"0x1": 101}},
				{1: {"0x1": 101}},
			},
			fChain: 5,
			expNonceConsensus: exectypes.NonceObservations{
				1: {"0x1": 100},
			},
		},
		{
			name: "12 oracles, 15 observations, 5 f",
			allNonceObservations: []exectypes.NonceObservations{
				{1: {"0x1": 134}},
				{1: {"0x1": 134}},
				{1: {"0x1": 134}},
				{1: {"0x1": 134}},
				{1: {"0x1": 134}},
				{1: {"0x1": 134}},
				{1: {"0x1": 134}},
				{1: {"0x1": 134}},
				{1: {"0x1": 134}},
				{1: {"0x1": 133}},
				{1: {"0x1": 133}},
				{1: {"0x1": 133}},
				{1: {"0x1": 133}},
				{1: {"0x1": 133}},
				{1: {"0x1": 133}},
			},
			fChain: 5,
			expNonceConsensus: exectypes.NonceObservations{
				1: {"0x1": 133},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			observations := make([]plugincommon.AttributedObservation[exectypes.Observation], len(tc.allNonceObservations))
			for i, obs := range tc.allNonceObservations {
				observations[i] = plugincommon.AttributedObservation[exectypes.Observation]{
					Observation: exectypes.Observation{Nonces: obs},
					OracleID:    commontypes.OracleID(i),
				}
			}
			obs := computeNoncesConsensus(lggr, observations, tc.fChain)
			assert.Equal(t, tc.expNonceConsensus, obs)
		})
	}
}

func Test_computeMessageHashesConsensus(t *testing.T) {
	testCases := []struct {
		name           string
		observations   []map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Bytes32
		expectedHashes exectypes.MessageHashes
		fChain         map[cciptypes.ChainSelector]int
	}{
		{
			name: "empty",
			observations: []map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Bytes32{
				{},
			},
			fChain: map[cciptypes.ChainSelector]int{},
		},
		{
			name: "base scenario reaching f+1 observations while having 2 faulty observations",
			observations: []map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Bytes32{
				{1: {10: cciptypes.Bytes32{111}}},
				{1: {10: cciptypes.Bytes32{111}}},
				{1: {10: cciptypes.Bytes32{222}}}, // <--- different observation is ignored
				{1: {10: cciptypes.Bytes32{222}}}, // <--- different observation is ignored
				{1: {10: cciptypes.Bytes32{111}}},
			},
			expectedHashes: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Bytes32{
				1: {10: cciptypes.Bytes32{111}},
			},
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
		},
		{
			name: "missing one observation and not reaching consensus",
			observations: []map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Bytes32{
				{1: {10: cciptypes.Bytes32{111}}},
				{1: {10: cciptypes.Bytes32{111}}},
			},
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
		},
		{
			name: "one observation has a different hash and we don't reach consensus",
			observations: []map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Bytes32{
				{1: {10: cciptypes.Bytes32{111}}},
				{1: {10: cciptypes.Bytes32{222}}}, // <-- different hash
				{1: {10: cciptypes.Bytes32{111}}},
			},
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
		},
		{
			name: "two different hashes reach consensus threshold f+1, meaning no consensus was formed",
			observations: []map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Bytes32{
				{1: {10: cciptypes.Bytes32{111}}},
				{1: {10: cciptypes.Bytes32{111}}},
				{1: {10: cciptypes.Bytes32{222}}},
				{1: {10: cciptypes.Bytes32{222}}},
				{1: {10: cciptypes.Bytes32{222}}},
				{1: {10: cciptypes.Bytes32{111}}},
			},
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
		},
		{
			name: "three different messages, one reaching consensus, the others not",
			observations: []map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Bytes32{
				{1: {10: cciptypes.Bytes32{111}, 20: cciptypes.Bytes32{222}, 30: cciptypes.Bytes32{103}}},
				{1: {10: cciptypes.Bytes32{111}, 20: cciptypes.Bytes32{222}, 30: cciptypes.Bytes32{103}}},
				{1: {10: cciptypes.Bytes32{111}, 20: cciptypes.Bytes32{222}}},
				{1: {10: cciptypes.Bytes32{111}, 20: cciptypes.Bytes32{221}}},
				{1: {10: cciptypes.Bytes32{111}, 20: cciptypes.Bytes32{221}}},
				{1: {10: cciptypes.Bytes32{111}, 20: cciptypes.Bytes32{221}}},
			},
			expectedHashes: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Bytes32{
				1: {10: cciptypes.Bytes32{111}},
			},
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
		},
		{
			name: "fChain not defined, chain is skipped",
			observations: []map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Bytes32{
				{1: {10: cciptypes.Bytes32{111}}},
				{1: {10: cciptypes.Bytes32{111}}},
				{1: {10: cciptypes.Bytes32{111}}},
			},
			fChain: map[cciptypes.ChainSelector]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			observations := make([]plugincommon.AttributedObservation[exectypes.Observation], len(tc.observations))
			for i, obs := range tc.observations {
				observations[i] = plugincommon.AttributedObservation[exectypes.Observation]{
					Observation: exectypes.Observation{Hashes: obs},
					OracleID:    commontypes.OracleID(i),
				}
			}
			obs := computeMessageHashesConsensus(logger.Test(t), observations, tc.fChain)
			assert.Equal(t, tc.expectedHashes, obs)
		})
	}
}

func Test_computeCommitObservationsConsensus(t *testing.T) {
	t1 := time.Now()

	baseObservation := map[cciptypes.ChainSelector][]exectypes.CommitData{
		1: {
			{
				SourceChain:         1,
				OnRampAddress:       []byte{5, 6, 7},
				Timestamp:           t1,
				BlockNum:            65535,
				MerkleRoot:          cciptypes.Bytes32{1, 2, 3},
				SequenceNumberRange: cciptypes.NewSeqNumRange(200, 300),
				ExecutedMessages:    []cciptypes.SeqNum{200, 201, 202, 299},
			},
		},
	}

	baseObservationDifferentOnRamp := map[cciptypes.ChainSelector][]exectypes.CommitData{
		1: {
			{
				SourceChain:         1,
				OnRampAddress:       []byte{8, 9, 10}, // <-- onRamp is different
				Timestamp:           t1,
				BlockNum:            65535,
				MerkleRoot:          cciptypes.Bytes32{1, 2, 3},
				SequenceNumberRange: cciptypes.NewSeqNumRange(200, 300),
				ExecutedMessages:    []cciptypes.SeqNum{200, 201, 202, 299},
			},
		},
	}

	testCases := []struct {
		name         string
		observations []exectypes.CommitObservations
		exp          exectypes.CommitObservations
		fChain       map[cciptypes.ChainSelector]int
	}{
		{
			name:         "empty",
			observations: []exectypes.CommitObservations{},
			exp:          exectypes.CommitObservations(nil),
			fChain:       map[cciptypes.ChainSelector]int{},
		},
		{
			name: "base scenario reaching consensus",
			observations: []exectypes.CommitObservations{
				baseObservation,
				baseObservation,
				baseObservation,
			},
			exp: baseObservation,
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
		},
		{
			name: "base scenario not reaching consensus, requires at least f+1=4 observations",
			observations: []exectypes.CommitObservations{
				baseObservation,
				baseObservation,
				baseObservation,
			},
			fChain: map[cciptypes.ChainSelector]int{
				1: 3,
			},
		},
		{
			name: "one oracle did not observe some of the executed messages, not reaching consensus on the executions",
			observations: []exectypes.CommitObservations{
				baseObservation,
				baseObservation,
				{
					1: {
						{
							SourceChain:         1,
							OnRampAddress:       []byte{5, 6, 7},
							Timestamp:           t1,
							BlockNum:            65535,
							MerkleRoot:          cciptypes.Bytes32{1, 2, 3},
							SequenceNumberRange: cciptypes.NewSeqNumRange(200, 300),
							ExecutedMessages:    []cciptypes.SeqNum{200, 201}, // <-- 202 and 299 not observed
						},
					},
				},
			},
			exp: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						SourceChain:         1,
						OnRampAddress:       []byte{5, 6, 7},
						Timestamp:           t1,
						BlockNum:            65535,
						MerkleRoot:          cciptypes.Bytes32{1, 2, 3},
						SequenceNumberRange: cciptypes.NewSeqNumRange(200, 300),
						ExecutedMessages:    []cciptypes.SeqNum{200, 201}, // <-- 202 and 299 not included
					},
				},
			},
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
		},
		{
			name: "same root with different data agreed twice, no consensus is reached",
			observations: []exectypes.CommitObservations{
				baseObservation,
				baseObservationDifferentOnRamp,
				baseObservation,
				baseObservation,
				baseObservationDifferentOnRamp,
				baseObservation,
				baseObservationDifferentOnRamp,
				baseObservationDifferentOnRamp,
			},
			fChain: map[cciptypes.ChainSelector]int{
				1: 2,
			},
		},
		{
			name: "fChain not defined",
			observations: []exectypes.CommitObservations{
				baseObservation,
				baseObservation,
				baseObservation,
			},
			fChain: map[cciptypes.ChainSelector]int{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			observations := make([]plugincommon.AttributedObservation[exectypes.Observation], len(tc.observations))
			for i, obs := range tc.observations {
				observations[i] = plugincommon.AttributedObservation[exectypes.Observation]{
					Observation: exectypes.Observation{CommitReports: obs},
					OracleID:    commontypes.OracleID(i),
				}
			}
			obs := computeCommitObservationsConsensus(logger.Test(t), observations, tc.fChain)
			assert.Equal(t, tc.exp, obs)
		})
	}
}

func Test_computeMessageObservationsConsensus(t *testing.T) {
	msgWithID9 := exectypes.MessageObservations{1: map[cciptypes.SeqNum]cciptypes.Message{
		1000: {Header: cciptypes.RampMessageHeader{SequenceNumber: 1000,
			MessageID: cciptypes.Bytes32{9, 9, 9, 9}, SourceChainSelector: 1, DestChainSelector: 2}}}}
	msgWithID2 := exectypes.MessageObservations{1: map[cciptypes.SeqNum]cciptypes.Message{
		1000: {Header: cciptypes.RampMessageHeader{SequenceNumber: 1000,
			MessageID: cciptypes.Bytes32{2, 2, 2, 2}, SourceChainSelector: 1, DestChainSelector: 2}}}}

	testCases := []struct {
		name         string
		observations []exectypes.MessageObservations
		exp          exectypes.MessageObservations
		fChain       map[cciptypes.ChainSelector]int
	}{
		{
			name:         "empty",
			observations: []exectypes.MessageObservations{},
			exp:          exectypes.MessageObservations(nil),
			fChain:       map[cciptypes.ChainSelector]int{},
		},
		{
			name: "base scenario reaching consensus",
			observations: []exectypes.MessageObservations{
				msgWithID9,
				msgWithID9,
				msgWithID9,
				msgWithID9,
			},
			exp: msgWithID9,
			fChain: map[cciptypes.ChainSelector]int{
				1: 3,
			},
		},
		{
			name: "one message has different id, no consensus",
			observations: []exectypes.MessageObservations{
				msgWithID9,
				msgWithID2,
				msgWithID9,
				msgWithID2,
			},
			exp: exectypes.MessageObservations(nil),
			fChain: map[cciptypes.ChainSelector]int{
				1: 3,
			},
		},
		{
			name: "consensus is reached on two different message ids but same seq nums",
			observations: []exectypes.MessageObservations{
				msgWithID9,
				msgWithID2,
				msgWithID9,
				msgWithID2,
			},
			exp: exectypes.MessageObservations(nil),
			fChain: map[cciptypes.ChainSelector]int{
				1: 1,
			},
		},
		{
			name: "no consensus",
			observations: []exectypes.MessageObservations{
				msgWithID9,
				msgWithID9,
			},
			exp: exectypes.MessageObservations(nil),
			fChain: map[cciptypes.ChainSelector]int{
				1: 3,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			observations := make([]plugincommon.AttributedObservation[exectypes.Observation], len(tc.observations))
			for i, obs := range tc.observations {
				observations[i] = plugincommon.AttributedObservation[exectypes.Observation]{
					Observation: exectypes.Observation{Messages: obs},
					OracleID:    commontypes.OracleID(i),
				}
			}
			obs := computeMessageObservationsConsensus(logger.Test(t), observations, tc.fChain)
			assert.Equal(t, tc.exp, obs)
		})
	}
}
