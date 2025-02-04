package execute

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/costlymessages"
	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_getMessagesObservation(t *testing.T) {
	ctx := context.Background()

	// Create mock objects
	ccipReader := readerpkg_mock.NewMockCCIPReader(t)
	msgHasher := mocks.NewMessageHasher()
	tokenDataObserver := tokendata.NoopTokenDataObserver{}
	costlyMessageObserver := costlymessages.NoopObserver{}

	//emptyMsgHash, err := msgHasher.Hash(ctx, cciptypes.Message{})
	//require.NoError(t, err)
	// Set up the plugin with mock objects
	plugin := &Plugin{
		lggr:                  mocks.NullLogger,
		ccipReader:            ccipReader,
		msgHasher:             msgHasher,
		tokenDataObserver:     &tokenDataObserver,
		costlyMessageObserver: &costlyMessageObserver,
		ocrTypeCodec:          jsonOcrTypeCodec,
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
			name: "multiple commit reports with missing messages",
			previousOutcome: exectypes.Outcome{
				CommitReports: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					},
					{
						SourceChain: 1,
						// test sets this up with missing messages
						SequenceNumberRange: cciptypes.NewSeqNumRange(6, 10),
					},
					{
						SourceChain: 1,
						// test sets this up with missing messages
						SequenceNumberRange: cciptypes.NewSeqNumRange(11, 12),
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
						{
							SourceChain:         1,
							SequenceNumberRange: cciptypes.NewSeqNumRange(11, 12),
						},
					},
				},
				Messages: exectypes.MessageObservations{
					1: {
						1:  cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 1}},
						2:  cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 2}},
						3:  cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 3}},
						11: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 11}},
						12: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 12}},
					},
				},
				CostlyMessages: []cciptypes.Bytes32{},
				TokenData: exectypes.TokenDataObservations{
					1: {
						1:  exectypes.NewMessageTokenData(),
						2:  exectypes.NewMessageTokenData(),
						3:  exectypes.NewMessageTokenData(),
						11: exectypes.NewMessageTokenData(),
						12: exectypes.NewMessageTokenData(),
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

			// missing message from 7 to 10
			ccipReader.On("MsgsBetweenSeqNums", ctx, cciptypes.ChainSelector(1),
				cciptypes.NewSeqNumRange(6, 10)).Return([]cciptypes.Message{
				{Header: cciptypes.RampMessageHeader{SequenceNumber: 6}},
			}, nil).Maybe()

			ccipReader.On("MsgsBetweenSeqNums", ctx, cciptypes.ChainSelector(1),
				cciptypes.NewSeqNumRange(11, 12)).Return([]cciptypes.Message{
				{Header: cciptypes.RampMessageHeader{SequenceNumber: 11}},
				{Header: cciptypes.RampMessageHeader{SequenceNumber: 12}},
			}, nil).Maybe()

			observation := exectypes.Observation{}
			observation, err := plugin.getMessagesObservation(ctx, plugin.lggr, tt.previousOutcome, observation)
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

func Test_readAllMessages(t *testing.T) {
	ctx := context.Background()

	// Create mock objects
	ccipReader := readerpkg_mock.NewMockCCIPReader(t)
	lggr := mocks.NullLogger
	timestamp := time.Now()

	tests := []struct {
		name               string
		commitData         []exectypes.CommitData
		expectedMsgs       exectypes.MessageObservations
		expectedReports    exectypes.CommitObservations
		expectedTimestamps map[cciptypes.Bytes32]time.Time
		expectedError      bool
	}{
		{
			name:               "no commit data",
			commitData:         []exectypes.CommitData{},
			expectedMsgs:       exectypes.MessageObservations{},
			expectedReports:    exectypes.CommitObservations{},
			expectedTimestamps: map[cciptypes.Bytes32]time.Time{},
			expectedError:      false,
		},
		{
			name: "valid commit data",
			commitData: []exectypes.CommitData{
				{
					SourceChain:         1,
					SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
					Timestamp:           timestamp,
				},
			},
			expectedMsgs: exectypes.MessageObservations{
				1: {
					1: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 1, MessageID: cciptypes.Bytes32{0x01}}},
					2: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 2, MessageID: cciptypes.Bytes32{0x02}}},
					3: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 3, MessageID: cciptypes.Bytes32{0x03}}},
				},
			},
			expectedReports: exectypes.CommitObservations{
				1: []exectypes.CommitData{
					{
						SourceChain:         1,
						SequenceNumberRange: cciptypes.NewSeqNumRange(1, 3),
						Timestamp:           timestamp,
					},
				},
			},
			expectedTimestamps: map[cciptypes.Bytes32]time.Time{
				{0x01}: timestamp,
				{0x02}: timestamp,
				{0x03}: timestamp,
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up mock expectations
			ccipReader.On("MsgsBetweenSeqNums", ctx, cciptypes.ChainSelector(1),
				cciptypes.NewSeqNumRange(1, 3)).Return([]cciptypes.Message{
				{Header: cciptypes.RampMessageHeader{SequenceNumber: 1, MessageID: cciptypes.Bytes32{0x01}}},
				{Header: cciptypes.RampMessageHeader{SequenceNumber: 2, MessageID: cciptypes.Bytes32{0x02}}},
				{Header: cciptypes.RampMessageHeader{SequenceNumber: 3, MessageID: cciptypes.Bytes32{0x03}}},
			}, nil).Maybe()

			msgs, reports, timestamps := readAllMessages(ctx, lggr, ccipReader, tt.commitData)
			assert.Equal(t, tt.expectedMsgs, msgs)
			assert.Equal(t, tt.expectedReports, reports)
			assert.Equal(t, tt.expectedTimestamps, timestamps)
		})
	}
}
