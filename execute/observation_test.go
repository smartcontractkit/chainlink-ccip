package execute

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/costlymessages"
	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/optimizers"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_filterCommitReports(t *testing.T) {
	tests := []struct {
		name       string
		msgObs     exectypes.MessageObservations
		commitObs  exectypes.CommitObservations
		want       exectypes.CommitObservations
		wantedMsgs exectypes.MessageObservations
	}{
		{
			name:   "no messages",
			msgObs: exectypes.MessageObservations{},
			commitObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 10),
					},
				},
			},
			want:       exectypes.CommitObservations{},
			wantedMsgs: exectypes.MessageObservations{},
		},
		{
			name: "missing messages",
			msgObs: exectypes.MessageObservations{
				1: {
					1: cciptypes.Message{},
					2: cciptypes.Message{},
				},
			},
			commitObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
				},
			},
			want:       exectypes.CommitObservations{},
			wantedMsgs: exectypes.MessageObservations{},
		},
		{
			name: "valid messages",
			msgObs: exectypes.MessageObservations{
				1: {
					1: cciptypes.Message{},
					2: cciptypes.Message{},
					3: cciptypes.Message{},
				},
			},
			commitObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
				},
			},
			want: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
				},
			},
			wantedMsgs: exectypes.MessageObservations{
				1: {
					1: cciptypes.Message{},
					2: cciptypes.Message{},
					3: cciptypes.Message{},
				},
			},
		},
		{
			name: "mixed valid and invalid messages",
			msgObs: exectypes.MessageObservations{
				1: {
					1: cciptypes.Message{},
					2: cciptypes.Message{},
					3: cciptypes.Message{},
					4: cciptypes.Message{},
				},
			},
			commitObs: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(4, 5),
					},
				},
			},
			want: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
				},
			},
			wantedMsgs: exectypes.MessageObservations{
				1: {
					1: cciptypes.Message{},
					2: cciptypes.Message{},
					3: cciptypes.Message{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := intersectCommitReportsAndMessages(tt.msgObs, tt.commitObs)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_getMessagesObservation(t *testing.T) {
	ctx := context.Background()

	// Create mock objects
	ccipReader := readerpkg_mock.NewMockCCIPReader(t)
	msgHasher := mocks.NewMessageHasher()
	tokenDataObserver := tokendata.NoopTokenDataObserver{}
	costlyMessageObserver := costlymessages.NoopObserver{}
	observationOptimizer := optimizers.NewObservationOptimizer(maxObservationLength)

	//emptyMsgHash, err := msgHasher.Hash(ctx, cciptypes.Message{})
	//require.NoError(t, err)
	// Set up the plugin with mock objects
	plugin := &Plugin{
		lggr:                  mocks.NullLogger,
		ccipReader:            ccipReader,
		msgHasher:             msgHasher,
		tokenDataObserver:     &tokenDataObserver,
		costlyMessageObserver: &costlyMessageObserver,
		observationOptimizer:  observationOptimizer,
	}

	tests := []struct {
		name            string
		previousOutcome exectypes.Outcome
		expectedObs     exectypes.Observation
		expectedError   bool
	}{
		{
			name: "no commit reports",
			previousOutcome: exectypes.Outcome{
				CommitReports: []exectypes.CommitData{},
			},
			expectedObs:   exectypes.Observation{},
			expectedError: false,
		},
		{
			name: "valid commit reports",
			previousOutcome: exectypes.Outcome{
				CommitReports: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
				},
			},
			expectedObs: exectypes.Observation{
				CommitReports: exectypes.CommitObservations{
					1: []exectypes.CommitData{
						{
							SourceChain:         1,
							SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
						},
					},
				},
				Messages: exectypes.MessageObservations{
					1: {
						1: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 1}},
						2: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 2}},
						3: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 3}},
					},
				},
				CostlyMessages: []cciptypes.Bytes32{},
				TokenData: exectypes.TokenDataObservations{
					1: {
						1: exectypes.NewMessageTokenData(),
						2: exectypes.NewMessageTokenData(),
						3: exectypes.NewMessageTokenData(),
					},
				},
			},
			expectedError: false,
		},
		{
			name: "multiple commit reports with one missing messages",
			previousOutcome: exectypes.Outcome{
				CommitReports: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(6, 10),
					},
				},
			},
			expectedObs: exectypes.Observation{
				CommitReports: exectypes.CommitObservations{
					1: []exectypes.CommitData{
						{
							SourceChain:         1,
							SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
						},
					},
				},
				Messages: exectypes.MessageObservations{
					1: {
						1: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 1}},
						2: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 2}},
						3: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 3}},
					},
				},
				CostlyMessages: []cciptypes.Bytes32{},
				TokenData: exectypes.TokenDataObservations{
					1: {
						1: exectypes.NewMessageTokenData(),
						2: exectypes.NewMessageTokenData(),
						3: exectypes.NewMessageTokenData(),
					},
				},
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up mock expectations
			ccipReader.On("MsgsBetweenSeqNums", ctx, cciptypes.ChainSelector(1),
				cciptypes.NewSeqNumRange(1, 3)).Return([]cciptypes.Message{
				{Header: cciptypes.RampMessageHeader{SequenceNumber: 1}},
				{Header: cciptypes.RampMessageHeader{SequenceNumber: 2}},
				{Header: cciptypes.RampMessageHeader{SequenceNumber: 3}},
			}, nil).Maybe()

			// missing message 5 and 6
			ccipReader.On("MsgsBetweenSeqNums", ctx, cciptypes.ChainSelector(1),
				cciptypes.NewSeqNumRange(6, 10)).Return([]cciptypes.Message{
				{Header: cciptypes.RampMessageHeader{SequenceNumber: 6}},
			}, nil).Maybe()

			observation := exectypes.Observation{}
			observation, err := plugin.getMessagesObservation(ctx, tt.previousOutcome, observation)
			if tt.expectedError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				observation.Hashes = nil
				assert.Equal(t, tt.expectedObs, observation)
			}
		})
	}
}
