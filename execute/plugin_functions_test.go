package execute

import (
	"fmt"
	"testing"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"
	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/libocr/commontypes"
	"github.com/smartcontractkit/libocr/offchainreporting2plus/types"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/plugincommon"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	plugintypes2 "github.com/smartcontractkit/chainlink-ccip/plugintypes"
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
		observedData map[cciptypes.ChainSelector][]exectypes.CommitData
		expErr       bool
	}{
		{
			name: "ValidData",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						MerkleRoot:          cciptypes.Bytes32{1},
						SequenceNumberRange: cciptypes.SeqNumRange{1, 10},
						ExecutedMessages:    []cciptypes.SeqNum{1, 2, 3},
					},
				},
				2: {
					{
						MerkleRoot:          cciptypes.Bytes32{2},
						SequenceNumberRange: cciptypes.SeqNumRange{11, 20},
						ExecutedMessages:    []cciptypes.SeqNum{11, 12, 13},
					},
				},
			},
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
			expErr: true,
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
			expErr: true,
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
			expErr: true,
		},
		{
			name: "NoCommitData",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {},
			},
		},
		{
			name:         "EmptyObservedData",
			observedData: map[cciptypes.ChainSelector][]exectypes.CommitData{},
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

func Test_groupByChainSelector(t *testing.T) {
	type args struct {
		reports []plugintypes2.CommitPluginReportWithMeta
	}
	tests := []struct {
		name string
		args args
		want exectypes.CommitObservations
	}{
		{
			name: "empty",
			args: args{reports: []plugintypes2.CommitPluginReportWithMeta{}},
			want: exectypes.CommitObservations{},
		},
		{
			name: "reports",
			args: args{reports: []plugintypes2.CommitPluginReportWithMeta{{
				Report: cciptypes.CommitPluginReport{
					MerkleRoots: []cciptypes.MerkleRootChain{
						{ChainSel: 1, SeqNumsRange: cciptypes.NewSeqNumRange(10, 20), MerkleRoot: cciptypes.Bytes32{1}},
						{ChainSel: 2, SeqNumsRange: cciptypes.NewSeqNumRange(30, 40), MerkleRoot: cciptypes.Bytes32{2}},
					}}}}},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, tt.want, groupByChainSelector(tt.args.reports), "groupByChainSelector(%v)", tt.args.reports)
		})
	}
}

func Test_filterOutFullyExecutedMessages(t *testing.T) {
	type args struct {
		reports          []exectypes.CommitData
		executedMessages []cciptypes.SeqNumRange
	}
	tests := []struct {
		name    string
		args    args
		want    []exectypes.CommitData
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
				reports:          []exectypes.CommitData{},
				executedMessages: nil,
			},
			want:    []exectypes.CommitData{},
			wantErr: assert.NoError,
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
			want: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
			},
			wantErr: assert.NoError,
		},
		{
			name: "executed messages",
			args: args{
				reports: []exectypes.CommitData{
					{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
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
				reports: []exectypes.CommitData{
					{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
				},
				executedMessages: []cciptypes.SeqNumRange{
					cciptypes.NewSeqNumRange(15, 35),
				},
			},
			want: []exectypes.CommitData{
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
			wantErr: assert.NoError,
		},
		{
			name: "2 partially executed 1 fully executed",
			args: args{
				reports: []exectypes.CommitData{
					{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
				},
				executedMessages: []cciptypes.SeqNumRange{
					cciptypes.NewSeqNumRange(15, 55),
				},
			},
			want: []exectypes.CommitData{
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					ExecutedMessages:    []cciptypes.SeqNum{15, 16, 17, 18, 19, 20},
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
					ExecutedMessages:    []cciptypes.SeqNum{50, 51, 52, 53, 54, 55},
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "first report executed",
			args: args{
				reports: []exectypes.CommitData{
					{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
				},
				executedMessages: []cciptypes.SeqNumRange{
					cciptypes.NewSeqNumRange(10, 20),
				},
			},
			want: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
			},
			wantErr: assert.NoError,
		},
		{
			name: "last report executed",
			args: args{
				reports: []exectypes.CommitData{
					{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
					{SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60)},
				},
				executedMessages: []cciptypes.SeqNumRange{
					cciptypes.NewSeqNumRange(50, 60),
				},
			},
			want: []exectypes.CommitData{
				{SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20)},
				{SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40)},
			},
			wantErr: assert.NoError,
		},
		{
			name: "sort-report",
			args: args{
				reports: []exectypes.CommitData{
					{
						SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
					},
					{
						SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
					},
					{
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					},
				},
				executedMessages: nil,
			},
			want: []exectypes.CommitData{
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
				},
				{
					SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
				},
			},
			wantErr: assert.NoError,
		},
		{
			name: "sort-executed",
			args: args{
				reports: []exectypes.CommitData{
					{
						SequenceNumberRange: cciptypes.NewSeqNumRange(10, 20),
					},
					{
						SequenceNumberRange: cciptypes.NewSeqNumRange(30, 40),
					},
					{
						SequenceNumberRange: cciptypes.NewSeqNumRange(50, 60),
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
	mustEncode := func(obs exectypes.Observation) []byte {
		enc, err := obs.Encode()
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
							1: {{MerkleRoot: cciptypes.Bytes32{1}, OnRampAddress: cciptypes.UnknownAddress{}}},
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
							1: {{MerkleRoot: cciptypes.Bytes32{1}, OnRampAddress: cciptypes.UnknownAddress{}}},
						},
					},
				},
				{
					OracleID: commontypes.OracleID(2),
					Observation: exectypes.Observation{
						CommitReports: exectypes.CommitObservations{
							2: {{MerkleRoot: cciptypes.Bytes32{2}, OnRampAddress: cciptypes.UnknownAddress{}}},
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

func Test_getConsensusObservation(t *testing.T) {
	type args struct {
		observation []exectypes.Observation
		F           int
		fChain      map[cciptypes.ChainSelector]int
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
				fChain: map[cciptypes.ChainSelector]int{
					1: 1,
				},
				observation: nil,
			},
			want:    exectypes.Observation{},
			wantErr: assert.NoError,
		},
		{
			name: "one consensus observation",
			args: args{
				fChain: map[cciptypes.ChainSelector]int{
					1: 0,
				},
				observation: []exectypes.Observation{
					{
						Nonces: exectypes.NonceObservations{
							1: {
								"0x1": 1,
							},
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
			name: "one ignored consensus observation",
			args: args{
				fChain: map[cciptypes.ChainSelector]int{
					1: 1,
				},
				observation: []exectypes.Observation{
					{
						Nonces: exectypes.NonceObservations{
							1: {
								"0x1": 1,
							},
						},
					},
				},
			},
			want:    exectypes.Observation{},
			wantErr: assert.NoError,
		},
		{
			name: "3 observers required to reach consensus on 4 sender values",
			args: args{
				fChain: map[cciptypes.ChainSelector]int{
					1: 1, // f + 1
				},
				// Across 3 observers
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
					}, {
						Nonces: exectypes.NonceObservations{
							1: {
								"0x1": 1,
								"0x4": 4,
							},
						},
					}, {
						Nonces: exectypes.NonceObservations{
							1: {
								"0x2": 2,
								"0x3": 3,
							},
						},
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
				fChain: map[cciptypes.ChainSelector]int{
					1: 2,
				},
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
					}, {
						Nonces: exectypes.NonceObservations{
							1: {
								"0x1": 1,
								"0x4": 4,
							},
						},
					}, {
						Nonces: exectypes.NonceObservations{
							1: {
								"0x2": 2,
								"0x3": 3,
							},
						},
					},
				},
			},
			want:    exectypes.Observation{},
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
			got, err := getConsensusObservation(lggr, ao, 1, tt.args.F, tt.args.fChain)
			if !tt.wantErr(t, err, "getConsensusObservation(...)") {
				return
			}
			assert.Equalf(t, tt.want, got, "getConsensusObservation(...)")
		})
	}
}

func Test_mergeTokenDataObservation(t *testing.T) {
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

			obs, err := mergeTokenObservations(ao, fChain)
			require.NoError(t, err)

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

func Test_mergeCostlyMessages(t *testing.T) {
	tests := []struct {
		name       string
		aos        []plugincommon.AttributedObservation[exectypes.Observation]
		fChainDest int
		want       []cciptypes.Bytes32
	}{
		{
			name:       "no observations",
			aos:        []plugincommon.AttributedObservation[exectypes.Observation]{},
			fChainDest: 1,
			want:       nil,
		},
		{
			name: "observations below threshold",
			aos: []plugincommon.AttributedObservation[exectypes.Observation]{
				{
					Observation: exectypes.Observation{
						CostlyMessages: []cciptypes.Bytes32{
							{0x01},
						},
					},
				},
				{
					Observation: exectypes.Observation{
						CostlyMessages: []cciptypes.Bytes32{
							{0x01},
						},
					},
				},
			},
			fChainDest: 3,
			want:       nil,
		},
		{
			name: "observations above threshold",
			aos: []plugincommon.AttributedObservation[exectypes.Observation]{
				{
					Observation: exectypes.Observation{
						CostlyMessages: []cciptypes.Bytes32{
							{0x01},
						},
					},
				},
				{
					Observation: exectypes.Observation{
						CostlyMessages: []cciptypes.Bytes32{
							{0x01},
						},
					},
				},
				{
					Observation: exectypes.Observation{
						CostlyMessages: []cciptypes.Bytes32{
							{0x01},
						},
					},
				},
			},
			fChainDest: 2,
			want: []cciptypes.Bytes32{
				{0x01},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeCostlyMessages(tt.aos, tt.fChainDest)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getMessageTimestampMap(t *testing.T) {
	tests := []struct {
		name              string
		commitReportCache map[cciptypes.ChainSelector][]exectypes.CommitData
		obs               exectypes.MessageObservations
		want              map[cciptypes.Bytes32]time.Time
		wantErr           assert.ErrorAssertionFunc
	}{
		{
			name:              "empty",
			commitReportCache: map[cciptypes.ChainSelector][]exectypes.CommitData{},
			obs:               exectypes.MessageObservations{},
			want:              map[cciptypes.Bytes32]time.Time{},
			wantErr:           assert.NoError,
		},
		{
			name:              "missing commit data for a chain",
			commitReportCache: map[cciptypes.ChainSelector][]exectypes.CommitData{},
			obs: exectypes.MessageObservations{
				1: {
					1: {
						Header: cciptypes.RampMessageHeader{
							MessageID: cciptypes.Bytes32{0x01},
						},
					},
				},
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "happy path",
			commitReportCache: map[cciptypes.ChainSelector][]exectypes.CommitData{
				1: {
					{
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
						Timestamp:           time.Unix(1, 0),
					},
					{
						SequenceNumberRange: cciptypes.NewSeqNumRange(15, 25),
						Timestamp:           time.Unix(2, 0),
					},
				},
			},
			obs: exectypes.MessageObservations{
				1: {
					1: {
						Header: cciptypes.RampMessageHeader{
							MessageID: cciptypes.Bytes32{0x01},
						},
					},
					16: {
						Header: cciptypes.RampMessageHeader{
							MessageID: cciptypes.Bytes32{0x04},
						},
					},
				},
			},
			want: map[cciptypes.Bytes32]time.Time{
				{0x01}: time.Unix(1, 0),
				{0x04}: time.Unix(2, 0),
			},
			wantErr: assert.NoError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getMessageTimestampMap(tt.commitReportCache, tt.obs)
			if !tt.wantErr(t, err, "getMessageTimestampMap(...)") {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_truncateLastCommit(t *testing.T) {
	tests := []struct {
		name        string
		chain       cciptypes.ChainSelector
		observation exectypes.Observation
		expected    exectypes.Observation
	}{
		{
			name:  "no commits to truncate",
			chain: 1,
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {},
				},
			},
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {},
				},
			},
		},
		{
			name:  "truncate last commit",
			chain: 1,
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
						{SequenceNumberRange: cciptypes.NewSeqNumRange(11, 20)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1:  {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
						11: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x02}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{
					1: {
						1:  exectypes.NewMessageTokenData(),
						11: exectypes.NewMessageTokenData(),
					},
				},
				CostlyMessages: []cciptypes.Bytes32{{0x02}},
			},
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{
					1: {
						1: exectypes.NewMessageTokenData(),
					},
				},
				CostlyMessages: []cciptypes.Bytes32{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			truncated := truncateLastCommit(tt.observation, tt.chain)
			require.Equal(t, tt.expected, truncated)
		})
	}
}

func Test_truncateChain(t *testing.T) {
	tests := []struct {
		name        string
		chain       cciptypes.ChainSelector
		observation exectypes.Observation
		expected    exectypes.Observation
	}{
		{
			name:  "truncate chain data",
			chain: 1,
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
						{SequenceNumberRange: cciptypes.NewSeqNumRange(11, 20)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{
					1: {
						1: exectypes.NewMessageTokenData(),
					},
				},
				CostlyMessages: []cciptypes.Bytes32{{0x01}},
			},
			expected: exectypes.Observation{
				CommitReports:  map[cciptypes.ChainSelector][]exectypes.CommitData{},
				Messages:       map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{},
				TokenData:      map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{},
				CostlyMessages: []cciptypes.Bytes32{},
			},
		},
		{
			name:  "truncate non existent chain",
			chain: 2, // non existent chain
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{
					1: {
						1: exectypes.NewMessageTokenData(),
					},
				},
				CostlyMessages: []cciptypes.Bytes32{{0x01}},
			},
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
				Messages: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]cciptypes.Message{
					1: {
						1: {Header: cciptypes.RampMessageHeader{MessageID: cciptypes.Bytes32{0x01}}},
					},
				},
				TokenData: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.MessageTokenData{
					1: {
						1: exectypes.NewMessageTokenData(),
					},
				},
				CostlyMessages: []cciptypes.Bytes32{{0x01}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			truncated := truncateChain(tt.observation, tt.chain)
			require.Equal(t, tt.expected, truncated)
		})
	}
}

func Test_truncateObservation(t *testing.T) {
	tests := []struct {
		name        string
		observation exectypes.Observation
		maxSize     int
		expected    exectypes.Observation
		deletedMsgs map[cciptypes.ChainSelector]map[cciptypes.SeqNum]struct{}
		wantErr     bool
	}{
		{
			name: "no truncation needed",
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
			},
			maxSize: 1000,
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
				},
			},
		},
		{
			name: "truncate last commit",
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2)},
						{SequenceNumberRange: cciptypes.NewSeqNumRange(3, 4)},
					},
				},
				Messages: makeMessageObservation(
					map[cciptypes.ChainSelector]cciptypes.SeqNumRange{
						1: {1, 4},
					},
					withData(make([]byte, 100)),
				),
			},
			maxSize: 2600, // this number is calculated by checking encoded sizes for the observation we expect
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2)},
					},
				},
				// Not going to check Messages in the test
			},
			deletedMsgs: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]struct{}{
				1: {
					3: struct{}{},
					4: struct{}{},
				},
			},
		},
		{
			name: "truncate entire chain",
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2)},
					},
					2: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(3, 4)},
					},
				},
				Messages: makeMessageObservation(
					map[cciptypes.ChainSelector]cciptypes.SeqNumRange{
						1: {1, 2},
						2: {3, 4},
					},
					withData(make([]byte, 100)),
				),
			},
			maxSize: 1789,
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					2: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(3, 4)},
					},
				},
			},
		},
		//{
		//	name: "truncate one message from first chain",
		//	observation: exectypes.Observation{
		//		CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
		//			1: {
		//				{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2)},
		//			},
		//			2: {
		//				{SequenceNumberRange: cciptypes.NewSeqNumRange(3, 4)},
		//			},
		//		},
		//		Messages: makeMessageObservation(
		//			map[cciptypes.ChainSelector]cciptypes.SeqNumRange{
		//				1: {1, 2},
		//				2: {3, 4},
		//			},
		//			withData(make([]byte, 100)),
		//		),
		//	},
		//	maxSize: 3159, // chain 1, message 1 will be truncated
		//	expected: exectypes.Observation{
		//		CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
		//			1: {
		//				{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 2)},
		//			},
		//			2: {
		//				{SequenceNumberRange: cciptypes.NewSeqNumRange(3, 4)},
		//			},
		//		},
		//	},
		//	deletedMsgs: map[cciptypes.ChainSelector]map[cciptypes.SeqNum]struct{}{
		//		1: {
		//			1: struct{}{},
		//		},
		//	},
		//},
		{
			name: "truncate all",
			observation: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{
					1: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10)},
					},
					2: {
						{SequenceNumberRange: cciptypes.NewSeqNumRange(11, 20)},
					},
				},
				Messages: makeMessageObservation(
					map[cciptypes.ChainSelector]cciptypes.SeqNumRange{
						1: {1, 10},
						2: {11, 20},
					},
				),
			},
			maxSize: 50, // less than what can fit a single commit report for single chain
			expected: exectypes.Observation{
				CommitReports: map[cciptypes.ChainSelector][]exectypes.CommitData{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.observation.TokenData = makeNoopTokenDataObservations(tt.observation.Messages)
			tt.expected.TokenData = tt.observation.TokenData
			obs, err := truncateObservation(tt.observation, tt.maxSize)
			if tt.wantErr {
				require.Error(t, err)
				return
			}
			require.Equal(t, tt.expected.CommitReports, obs.CommitReports)

			// Check only deleted messages were deleted
			for chain, seqNums := range obs.Messages {
				for seqNum := range seqNums {
					if _, ok := tt.deletedMsgs[chain]; ok {
						if _, ok2 := tt.deletedMsgs[chain][seqNum]; ok2 {
							require.True(t, obs.Messages[chain][seqNum].IsEmpty())
							continue
						}
					}
					require.False(t, obs.Messages[chain][seqNum].IsEmpty())
				}
			}
		})
	}
}
