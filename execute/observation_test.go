package execute

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-ccip/mocks/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/libocr/offchainreporting2plus/ocr3types"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/internal/cache"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/observer"
	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
	codec_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/ocrtypecodec/v1"
	readerpkg_mock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func Test_Observation_CacheUpdate(t *testing.T) {

	getID := func(id string) cciptypes.Bytes32 {
		var b cciptypes.Bytes32
		copy(b[:], id)
		return b
	}
	msgWithID := func(id string) cciptypes.Message {
		var msg cciptypes.Message
		msg.Header.MessageID = getID(id)
		return msg
	}

	homeChain := reader.NewMockHomeChain(t)
	homeChain.EXPECT().GetFChain().Return(nil, fmt.Errorf("early return"))
	plugin := &Plugin{
		lggr:                 mocks.NullLogger,
		homeChain:            homeChain,
		ocrTypeCodec:         ocrTypeCodec,
		inflightMessageCache: cache.NewInflightMessageCache(10 * time.Minute),
	}

	outcome := exectypes.Outcome{
		Report: cciptypes.ExecutePluginReport{
			ChainReports: []cciptypes.ExecutePluginReportSingleChain{
				{
					SourceChainSelector: 1,
					Messages: []cciptypes.Message{
						msgWithID("1"),
						msgWithID("2"),
						msgWithID("3"),
					},
				},
				{
					SourceChainSelector: 2,
					Messages: []cciptypes.Message{
						msgWithID("1"),
						msgWithID("2"),
						msgWithID("3"),
					},
				},
			},
		},
	}

	// No state, report only generated in Filter state so cache is not updated.
	{
		enc, err := ocrTypeCodec.EncodeOutcome(outcome)
		require.NoError(t, err)

		outCtx := ocr3types.OutcomeContext{PreviousOutcome: enc}
		_, err = plugin.Observation(context.Background(), outCtx, nil)
		require.Error(t, err) // home reader error

		require.False(t, plugin.inflightMessageCache.IsInflight(1, getID("1")))
		require.False(t, plugin.inflightMessageCache.IsInflight(1, getID("2")))
		require.False(t, plugin.inflightMessageCache.IsInflight(1, getID("3")))
		require.False(t, plugin.inflightMessageCache.IsInflight(2, getID("1")))
		require.False(t, plugin.inflightMessageCache.IsInflight(2, getID("2")))
		require.False(t, plugin.inflightMessageCache.IsInflight(2, getID("3")))
	}

	// Filter state, cache is updated.
	{
		outcome.State = exectypes.Filter
		enc, err := ocrTypeCodec.EncodeOutcome(outcome)
		require.NoError(t, err)

		outCtx := ocr3types.OutcomeContext{PreviousOutcome: enc}
		_, err = plugin.Observation(context.Background(), outCtx, nil)
		require.Error(t, err)

		require.True(t, plugin.inflightMessageCache.IsInflight(1, getID("1")))
		require.True(t, plugin.inflightMessageCache.IsInflight(1, getID("2")))
		require.True(t, plugin.inflightMessageCache.IsInflight(1, getID("3")))
		require.True(t, plugin.inflightMessageCache.IsInflight(2, getID("1")))
		require.True(t, plugin.inflightMessageCache.IsInflight(2, getID("2")))
		require.True(t, plugin.inflightMessageCache.IsInflight(2, getID("3")))
	}

}

// TODO: Test Gas Limits
// TODO: Test Size limits
func Test_getMessagesObservation(t *testing.T) {
	ctx := context.Background()

	// Create mock objects
	ccipReader := readerpkg_mock.NewMockCCIPReader(t)
	msgHasher := mocks.NewMessageHasher()
	tokenDataObserver := observer.NoopTokenDataObserver{}
	estimateProvider := ccipocr3.NewMockEstimateProvider(t)
	estimateProvider.EXPECT().CalculateMessageMaxGas(mock.Anything).Return(1)
	estimateProvider.EXPECT().CalculateMerkleTreeGas(mock.Anything).Return(1)
	//ccipReader.EXPECT().ExecutedMessages(mock.Anything, mock.Anything, primitives.Unconfirmed).Return(nil, nil)

	// Set up the plugin with mock objects
	plugin := &Plugin{
		lggr:              mocks.NullLogger,
		ccipReader:        ccipReader,
		msgHasher:         msgHasher,
		tokenDataObserver: &tokenDataObserver,
		ocrTypeCodec:      ocrTypeCodec,
		estimateProvider:  estimateProvider,
		offchainCfg: pluginconfig.ExecuteOffchainConfig{
			BatchGasLimit: 10,
		},
		inflightMessageCache: cache.NewInflightMessageCache(1 * time.Minute),
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
						1: NewMessage(1, 1, 1, 2),
						2: NewMessage(2, 2, 1, 2),
						3: NewMessage(3, 3, 1, 2),
					},
				},
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
						SourceChain:         1,
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
						1:  NewMessage(1, 1, 1, 2),
						2:  NewMessage(2, 2, 1, 2),
						3:  NewMessage(3, 3, 1, 2),
						11: NewMessage(11, 11, 1, 2),
						12: NewMessage(12, 12, 1, 2),
					},
				},
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
				NewMessage(1, 1, 1, 2),
				NewMessage(2, 2, 1, 2),
				NewMessage(3, 3, 1, 2),
			}, nil).Maybe()

			// missing message from 7 to 10
			ccipReader.On("MsgsBetweenSeqNums", ctx, cciptypes.ChainSelector(1),
				cciptypes.NewSeqNumRange(6, 10)).Return([]cciptypes.Message{
				NewMessage(6, 6, 1, 2),
			}, nil).Maybe()

			ccipReader.On("MsgsBetweenSeqNums", ctx, cciptypes.ChainSelector(1),
				cciptypes.NewSeqNumRange(11, 12)).Return([]cciptypes.Message{
				NewMessage(11, 11, 1, 2),
				NewMessage(12, 12, 1, 2),
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
	lggr := mocks.NullLogger
	timestamp := time.Now()
	ccipReader := readerpkg_mock.NewMockCCIPReader(t)
	msgHasher := mocks.NewMessageHasher()
	tokenDataObserver := observer.NoopTokenDataObserver{}
	estimateProvider := ccipocr3.NewMockEstimateProvider(t)
	estimateProvider.EXPECT().CalculateMessageMaxGas(mock.Anything).Return(1)
	estimateProvider.EXPECT().CalculateMerkleTreeGas(mock.Anything).Return(1)
	//ccipReader.EXPECT().ExecutedMessages(mock.Anything, mock.Anything, primitives.Unconfirmed).Return(nil, nil)

	// Set up the plugin with mock objects
	plugin := &Plugin{
		lggr:              lggr,
		ccipReader:        ccipReader,
		msgHasher:         msgHasher,
		tokenDataObserver: &tokenDataObserver,
		ocrTypeCodec:      ocrTypeCodec,
		estimateProvider:  estimateProvider,
		offchainCfg: pluginconfig.ExecuteOffchainConfig{
			BatchGasLimit: 10,
		},
		inflightMessageCache: cache.NewInflightMessageCache(1 * time.Minute),
	}

	tests := []struct {
		name            string
		commitData      []exectypes.CommitData
		expectedMsgs    exectypes.MessageObservations
		expectedReports exectypes.CommitObservations
		expectedHashes  exectypes.MessageHashes
		expectedError   bool
	}{
		{
			name:            "no commit data",
			commitData:      []exectypes.CommitData{},
			expectedMsgs:    nil,
			expectedReports: nil,
			expectedHashes:  nil,
			expectedError:   false,
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
					1: NewMessage(1, 1, 1, 2),
					2: NewMessage(2, 2, 1, 2),
					3: NewMessage(3, 3, 1, 2),
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
			expectedHashes: exectypes.MessageHashes{
				1: {
					1: cciptypes.Bytes32{0x01},
					2: cciptypes.Bytes32{0x02},
					3: cciptypes.Bytes32{0x03},
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
				NewMessage(1, 1, 1, 2),
				NewMessage(2, 2, 1, 2),
				NewMessage(3, 3, 1, 2),
			}, nil).Maybe()
			obs, err := plugin.getObsWithoutTokenData(ctx, lggr, tt.commitData, exectypes.Observation{})
			require.NoError(t, err)
			expectedObs := exectypes.Observation{
				Messages:      tt.expectedMsgs,
				CommitReports: tt.expectedReports,
				Hashes:        tt.expectedHashes,
			}
			require.Equal(t, expectedObs, obs)
		})
	}
}

func Test_getObsWithoutTokenData(t *testing.T) {
	ctx := context.Background()
	timestamp := time.Now()

	// Configuration constants
	const (
		batchGasLimit        = 10000
		gasPerMsg            = 100
		gasPerMerkleTree     = 200
		highGasPerMsg        = 9000
		highGasPerMerkleTree = 2000
		inflightCacheTTL     = 10 * time.Minute
	)

	// Chain selectors
	const (
		src1 = cciptypes.ChainSelector(1)
		src2 = cciptypes.ChainSelector(2)
		dest = cciptypes.ChainSelector(3)
	)

	oneByte := make([]byte, 1)

	// Helper functions
	createCommitData := func(srcChain cciptypes.ChainSelector, from, to cciptypes.SeqNum) exectypes.CommitData {
		return exectypes.CommitData{
			SourceChain:         srcChain,
			SequenceNumberRange: cciptypes.NewSeqNumRange(from, to),
			Timestamp:           timestamp,
		}
	}

	createMessages := func(srcChain cciptypes.ChainSelector, destChain cciptypes.ChainSelector, fromSeq, toSeq cciptypes.SeqNum) []cciptypes.Message {
		var msgs []cciptypes.Message
		for seq := fromSeq; seq <= toSeq; seq++ {
			msgs = append(msgs, NewMessage(int(seq), int(seq), int(srcChain), int(destChain)))
		}
		return msgs
	}

	createHashesMap := func(srcChain cciptypes.ChainSelector, fromSeq, toSeq cciptypes.SeqNum) map[cciptypes.SeqNum]cciptypes.Bytes32 {
		hashes := make(map[cciptypes.SeqNum]cciptypes.Bytes32)
		for seq := fromSeq; seq <= toSeq; seq++ {
			hashes[seq] = cciptypes.Bytes32{byte(seq)}
		}
		return hashes
	}

	setupMerkleGasExpectation := func(estimateProvider *ccipocr3.MockEstimateProvider, numMsgs int, gas uint64) {
		estimateProvider.EXPECT().CalculateMerkleTreeGas(numMsgs).Return(gas)
	}

	setupMessageGasExpectations := func(estimateProvider *ccipocr3.MockEstimateProvider, msgs []cciptypes.Message, gas uint64) {
		for _, msg := range msgs {
			estimateProvider.EXPECT().CalculateMessageMaxGas(msg).Return(gas)
		}
	}

	tests := []struct {
		name       string
		commitData []exectypes.CommitData
		setupMocks func(ccipReader *readerpkg_mock.MockCCIPReader,
			estimateProvider *ccipocr3.MockEstimateProvider,
			inflightCache *cache.InflightMessageCache,
			codec *codec_mock.MockExecCodec)
		expectedObs    exectypes.Observation
		expectedError  bool
		errorSubstring string
	}{
		{
			name:       "no commit data",
			commitData: []exectypes.CommitData{},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				// No mocks needed for empty commit data
			},
			expectedObs:   exectypes.Observation{},
			expectedError: false,
		},
		{
			name: "single valid commit data",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 3),
			},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				messages := createMessages(src1, dest, 1, 3)

				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 3)).
					Return(messages, nil)
				// Any small size that fits within the max observation size
				codec.EXPECT().EncodeObservation(mock.Anything).Return(oneByte, nil).Maybe()

				setupMessageGasExpectations(estimateProvider, messages, gasPerMsg)
				setupMerkleGasExpectation(estimateProvider, len(messages), gasPerMerkleTree) // 300 for 3 messages
			},
			expectedObs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
					src1: {
						1: NewMessage(1, 1, int(src1), int(dest)),
						2: NewMessage(2, 2, int(src1), int(dest)),
						3: NewMessage(3, 3, int(src1), int(dest)),
					},
				},
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 3),
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: createHashesMap(src1, 1, 3),
				},
			},
			expectedError: false,
		},
		{
			name: "multiple commit data from different source chains",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 2),
				createCommitData(src2, 5, 6),
			},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				messages1 := createMessages(src1, dest, 1, 2)
				messages2 := createMessages(src2, dest, 5, 6)
				// Any small size that fits within the max observation size
				codec.EXPECT().EncodeObservation(mock.Anything).Return(oneByte, nil).Maybe()

				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 2)).
					Return(messages1, nil)
				ccipReader.On("MsgsBetweenSeqNums", ctx, src2, cciptypes.NewSeqNumRange(5, 6)).
					Return(messages2, nil)

				setupMessageGasExpectations(estimateProvider, messages1, gasPerMsg)
				setupMessageGasExpectations(estimateProvider, messages2, gasPerMsg)

				setupMerkleGasExpectation(estimateProvider, len(messages1), gasPerMerkleTree)
				setupMerkleGasExpectation(estimateProvider, len(messages2), gasPerMerkleTree)
			},
			expectedObs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
					src1: {
						1: NewMessage(1, 1, int(src1), int(dest)),
						2: NewMessage(2, 2, int(src1), int(dest)),
					},
					src2: {
						5: NewMessage(5, 5, int(src2), int(dest)),
						6: NewMessage(6, 6, int(src2), int(dest)),
					},
				},
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 2),
					},
					src2: []exectypes.CommitData{
						createCommitData(src2, 5, 6),
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: createHashesMap(src1, 1, 2),
					src2: createHashesMap(src2, 5, 6),
				},
			},
			expectedError: false,
		},
		{
			name: "error reading messages",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 3),
			},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				// Any small size that fits within the max observation size
				codec.EXPECT().EncodeObservation(mock.Anything).Return(oneByte, nil).Maybe()
				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 3)).
					Return(nil, fmt.Errorf("failed to read messages"))
			},
			expectedObs:   exectypes.Observation{},
			expectedError: false, // Should continue processing even with errors
		},
		{
			name: "missing messages in range",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 3),
			},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				// Any small size that fits within the max observation size
				codec.EXPECT().EncodeObservation(mock.Anything).Return(oneByte, nil).Maybe()
				// Return only 2 messages for range 1-3
				incompleteMessages := []cciptypes.Message{
					NewMessage(1, 1, int(src1), int(dest)),
					NewMessage(3, 3, int(src1), int(dest)), // Missing message 2
				}
				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 3)).
					Return(incompleteMessages, nil)
			},
			expectedObs:   exectypes.Observation{},
			expectedError: false, // Should continue processing even with errors
		},
		{
			name: "inflight messages are skipped but hashed",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 2),
			},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				// Any small size that fits within the max observation size
				codec.EXPECT().EncodeObservation(mock.Anything).Return(oneByte, nil).Maybe()
				messages := createMessages(src1, dest, 1, 2)

				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 2)).
					Return(messages, nil)

				// Mark message 1 as inflight
				inflightCache.MarkInflight(src1, messages[0].Header.MessageID)

				// Only calculate gas for non-inflight message
				estimateProvider.EXPECT().CalculateMessageMaxGas(messages[1]).Return(uint64(gasPerMsg))
				setupMerkleGasExpectation(estimateProvider, len(messages), gasPerMerkleTree)
			},
			expectedObs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
					src1: {
						1: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 1}},
						2: NewMessage(2, 2, int(src1), int(dest)),
					},
				},
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 2),
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: createHashesMap(src1, 1, 2),
				},
			},
			expectedError: false,
		},
		{
			name: "gas limit exceeded",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 3),
			},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				// Any small size that fits within the max observation size
				codec.EXPECT().EncodeObservation(mock.Anything).Return(oneByte, nil).Maybe()
				messages := createMessages(src1, dest, 1, 3)

				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 3)).
					Return(messages, nil)

				// First message takes almost all gas
				estimateProvider.EXPECT().CalculateMessageMaxGas(messages[0]).Return(uint64(highGasPerMsg))

				// Merkle tree gas calculation pushes us over the limit
				setupMerkleGasExpectation(estimateProvider, len(messages), highGasPerMerkleTree)
			},
			expectedObs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
					src1: {
						1: NewMessage(1, 1, int(src1), int(dest)),
						2: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 2}},
						3: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 3}},
					},
				},
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 3),
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: createHashesMap(src1, 1, 3),
				},
			},
			expectedError: false,
		},
		{
			name: "gas limit exceeded mid second report",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 2),
				createCommitData(src2, 1, 3),
			},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				messages1 := createMessages(src1, dest, 1, 2)
				messages2 := createMessages(src2, dest, 1, 3)

				// Return messages for both chains
				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 2)).
					Return(messages1, nil)
				ccipReader.On("MsgsBetweenSeqNums", ctx, src2, cciptypes.NewSeqNumRange(1, 3)).
					Return(messages2, nil)

				// Any small size that fits within the max observation size
				codec.EXPECT().EncodeObservation(mock.Anything).Return(oneByte, nil).Maybe()

				// First chain consumes moderate gas
				setupMessageGasExpectations(estimateProvider, messages1, gasPerMsg)
				setupMerkleGasExpectation(estimateProvider, len(messages1), gasPerMerkleTree)

				// First message in second chain consumes moderate gas
				estimateProvider.EXPECT().CalculateMessageMaxGas(messages2[0]).Return(uint64(gasPerMsg))

				// Second message in second chain pushes us over the limit
				estimateProvider.EXPECT().CalculateMessageMaxGas(messages2[1]).Return(uint64(batchGasLimit - 500))

				// Merkle tree gas calculation for second chain pushes us over the limit
				setupMerkleGasExpectation(estimateProvider, len(messages2), 600)
			},
			expectedObs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
					src1: {
						1: NewMessage(1, 1, int(src1), int(dest)),
						2: NewMessage(2, 2, int(src1), int(dest)),
					},
					src2: {
						1: NewMessage(1, 1, int(src2), int(dest)),
						// We still include the second message anyway
						2: NewMessage(2, 2, int(src2), int(dest)),
						3: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 3}},
					},
				},
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 2),
					},
					src2: []exectypes.CommitData{
						createCommitData(src2, 1, 3),
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: createHashesMap(src1, 1, 2),
					src2: createHashesMap(src2, 1, 3),
				},
			},
			expectedError: false,
		},
		{
			name: "encoding size exceeded mid report",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 2),
			},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				// first call to EncodeObservation returns a small size
				codec.EXPECT().EncodeObservation(mock.Anything).
					Return(oneByte, nil).
					Once()
				// second call to EncodeObservation returns a large size, second message shouldn't be included
				codec.EXPECT().EncodeObservation(mock.Anything).
					Return(make([]byte, maxObservationLength+1), nil).
					Once()
				messages := createMessages(src1, dest, 1, 2)

				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 2)).
					Return(messages, nil)

				setupMessageGasExpectations(estimateProvider, messages, gasPerMsg)
				setupMerkleGasExpectation(estimateProvider, len(messages), gasPerMerkleTree)
			},
			expectedObs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
					src1: {
						1: NewMessage(1, 1, int(src1), int(dest)),
						2: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 2}},
					},
				},
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 2),
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: createHashesMap(src1, 1, 2),
				},
			},
			expectedError: false,
		},
		{
			name: "encoding size exceeded mid second report",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 2),
				createCommitData(src2, 1, 3),
			},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				messages1 := createMessages(src1, dest, 1, 2)
				messages2 := createMessages(src2, dest, 1, 3)

				// Return messages for both chains
				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 2)).
					Return(messages1, nil)
				ccipReader.On("MsgsBetweenSeqNums", ctx, src2, cciptypes.NewSeqNumRange(1, 3)).
					Return(messages2, nil)

				// First 4 calls to EncodeObservation return small sizes
				codec.EXPECT().EncodeObservation(mock.Anything).
					Return(oneByte, nil).
					Times(4)
				// Fifth call to EncodeObservation returns a large size
				codec.EXPECT().EncodeObservation(mock.Anything).
					Return(make([]byte, maxObservationLength+1), nil).
					Once()

				estimateProvider.EXPECT().CalculateMessageMaxGas(mock.Anything).Return(uint64(gasPerMsg)).Maybe()
				setupMerkleGasExpectation(estimateProvider, len(messages1), gasPerMerkleTree)
				setupMerkleGasExpectation(estimateProvider, len(messages2), gasPerMerkleTree)
			},
			expectedObs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
					src1: {
						1: NewMessage(1, 1, int(src1), int(dest)),
						2: NewMessage(2, 2, int(src1), int(dest)),
					},
					src2: {
						1: NewMessage(1, 1, int(src2), int(dest)),
						2: NewMessage(2, 2, int(src2), int(dest)),
						3: cciptypes.Message{Header: cciptypes.RampMessageHeader{SequenceNumber: 3}},
					},
				},
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 2),
					},
					src2: []exectypes.CommitData{
						createCommitData(src2, 1, 3),
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: createHashesMap(src1, 1, 2),
					src2: createHashesMap(src2, 1, 3),
				},
			},
			expectedError: false,
		},
		{
			name: "multiple sources with one failing",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 2),
				createCommitData(src2, 1, 2),
			},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				// Any small size that fits within the max observation size
				codec.EXPECT().EncodeObservation(mock.Anything).Return(oneByte, nil).Maybe()
				// Chain 1 works
				messages1 := createMessages(src1, dest, 1, 2)

				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 2)).
					Return(messages1, nil)

				// Chain 2 fails
				ccipReader.On("MsgsBetweenSeqNums", ctx, src2, cciptypes.NewSeqNumRange(1, 2)).
					Return(nil, fmt.Errorf("failed to read"))

				// Process chain 1 messages
				setupMessageGasExpectations(estimateProvider, messages1, gasPerMsg)
				setupMerkleGasExpectation(estimateProvider, len(messages1), gasPerMerkleTree)
			},
			expectedObs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
					src1: {
						1: NewMessage(1, 1, int(src1), int(dest)),
						2: NewMessage(2, 2, int(src1), int(dest)),
					},
				},
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 2),
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: createHashesMap(src1, 1, 2),
				},
			},
			expectedError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ccipReader := readerpkg_mock.NewMockCCIPReader(t)
			msgHasher := mocks.NewMessageHasher()
			estimateProvider := ccipocr3.NewMockEstimateProvider(t)
			inflightCache := cache.NewInflightMessageCache(inflightCacheTTL)
			codec := codec_mock.NewMockExecCodec(t)

			plugin := &Plugin{
				lggr:                 mocks.NullLogger,
				ccipReader:           ccipReader,
				msgHasher:            msgHasher,
				ocrTypeCodec:         codec,
				estimateProvider:     estimateProvider,
				inflightMessageCache: inflightCache,
				offchainCfg: pluginconfig.ExecuteOffchainConfig{
					BatchGasLimit: uint64(batchGasLimit),
				},
			}

			tt.setupMocks(ccipReader, estimateProvider, inflightCache, codec)

			obs, err := plugin.getObsWithoutTokenData(ctx, plugin.lggr, tt.commitData, exectypes.Observation{})

			if tt.expectedError {
				require.Error(t, err)
				if tt.errorSubstring != "" {
					require.Contains(t, err.Error(), tt.errorSubstring)
				}
			} else {
				require.NoError(t, err)
				if len(tt.expectedObs.Messages) > 0 {
					require.Equal(t, tt.expectedObs.Messages, obs.Messages)
				}
				if len(tt.expectedObs.CommitReports) > 0 {
					require.Equal(t, tt.expectedObs.CommitReports, obs.CommitReports)
				}
				if len(tt.expectedObs.Hashes) > 0 {
					require.Equal(t, tt.expectedObs.Hashes, obs.Hashes)
				}
			}
		})
	}
}
