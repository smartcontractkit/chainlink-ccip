package execute

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/smartcontractkit/chainlink-ccip/mocks/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"

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

func Test_getMessagesObservation(t *testing.T) {
	ctx := context.Background()
	timestamp := time.Now()

	// Configuration constants
	const (
		batchGasLimit    = 10000
		inflightCacheTTL = 10 * time.Minute
	)

	// Chain selectors
	const (
		src1 = cciptypes.ChainSelector(1)
		src2 = cciptypes.ChainSelector(2)
		dest = cciptypes.ChainSelector(3)
	)

	oneByte := make([]byte, 1)

	// Helper functions
	createCommitData := func(
		srcChain cciptypes.ChainSelector,
		from,
		to cciptypes.SeqNum,
		executedMessages ...cciptypes.SeqNum,
	) exectypes.CommitData {
		return exectypes.CommitData{
			SourceChain:         srcChain,
			SequenceNumberRange: cciptypes.NewSeqNumRange(from, to),
			Timestamp:           timestamp,
			ExecutedMessages:    executedMessages,
		}
	}

	createMessages := func(srcChain cciptypes.ChainSelector,
		destChain cciptypes.ChainSelector,
		fromSeq, toSeq cciptypes.SeqNum) []cciptypes.Message {
		var msgs []cciptypes.Message
		for seq := fromSeq; seq <= toSeq; seq++ {
			msgs = append(msgs, NewMessage(int(seq), int(seq), int(srcChain), int(destChain)))
		}
		return msgs
	}

	createHashesMap := func(
		fromSeq, toSeq cciptypes.SeqNum) map[cciptypes.SeqNum]cciptypes.Bytes32 {
		hashes := make(map[cciptypes.SeqNum]cciptypes.Bytes32)
		for seq := fromSeq; seq <= toSeq; seq++ {
			hashes[seq] = cciptypes.Bytes32{byte(seq)}
		}
		return hashes
	}

	createTokenData := func(fromSeq, toSeq cciptypes.SeqNum) map[cciptypes.SeqNum]exectypes.MessageTokenData {
		tokenData := make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
		for seq := fromSeq; seq <= toSeq; seq++ {
			tokenData[seq] = exectypes.NewMessageTokenData()
		}
		return tokenData
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
					src1: createHashesMap(1, 3),
				},
				TokenData: exectypes.TokenDataObservations{
					src1: createTokenData(1, 3),
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
					src1: createHashesMap(1, 2),
					src2: createHashesMap(5, 6),
				},
				TokenData: exectypes.TokenDataObservations{
					src1: createTokenData(1, 2),
					src2: createTokenData(5, 6),
				},
			},
			expectedError: false,
		},
		{
			name: "multiple commit reports with missing messages",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 3),
				createCommitData(src1, 6, 10),
				createCommitData(src1, 11, 12),
			},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				// Any small size that fits within the max observation size
				codec.EXPECT().EncodeObservation(mock.Anything).Return(oneByte, nil).Maybe()

				// Setup messages for range 1-3
				messages1 := createMessages(src1, dest, 1, 3)
				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 3)).
					Return(messages1, nil)

				// Setup messages for range 6-10 with missing messages (only return message 6)
				messages2 := []cciptypes.Message{
					NewMessage(6, 6, int(src1), int(dest)),
					// Missing messages 7-10
				}
				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(6, 10)).
					Return(messages2, nil)

				// Setup messages for range 11-12
				messages3 := createMessages(src1, dest, 11, 12)
				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(11, 12)).
					Return(messages3, nil)
			},
			expectedObs: exectypes.Observation{
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 3),
						createCommitData(src1, 11, 12),
						// Range 6-10 should be excluded due to missing messages
					},
				},
				Messages: exectypes.MessageObservations{
					src1: {
						1:  NewMessage(1, 1, int(src1), int(dest)),
						2:  NewMessage(2, 2, int(src1), int(dest)),
						3:  NewMessage(3, 3, int(src1), int(dest)),
						11: NewMessage(11, 11, int(src1), int(dest)),
						12: NewMessage(12, 12, int(src1), int(dest)),
						// Messages 6-10 should not be included
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: map[cciptypes.SeqNum]cciptypes.Bytes32{
						1:  {1},
						2:  {2},
						3:  {3},
						11: {11},
						12: {12},
					},
				},
				TokenData: exectypes.TokenDataObservations{
					src1: {
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
			},
			expectedObs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
					src1: {
						// pseudo deleted
						1: NewMessage(1, 1, 0, 0),
						2: NewMessage(2, 2, int(src1), int(dest)),
					},
				},
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 2),
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: createHashesMap(1, 2),
				},
				TokenData: exectypes.TokenDataObservations{
					src1: createTokenData(1, 2),
				},
			},
			expectedError: false,
		},
		{
			name: "executed messages are skipped but hashed",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 3, 1, 2), // 1 and 2 are already executed
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
			},
			expectedObs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
					src1: {
						// 1 and 2 are pseudo deleted because they are already executed
						1: NewMessage(1, 1, 0, 0),
						2: NewMessage(2, 2, 0, 0),
						3: NewMessage(3, 3, int(src1), int(dest)),
					},
				},
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 3, 1, 2), // 1 and 2 are already executed
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: createHashesMap(1, 3),
				},
				TokenData: exectypes.TokenDataObservations{
					src1: createTokenData(1, 3),
				},
			},
			expectedError: false,
		},
		{
			name: "encoding size exceeded mid report",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 3),
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
					Return(make([]byte, lenientMaxObservationLength+1), nil)

				messages := createMessages(src1, dest, 1, 3)

				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 3)).
					Return(messages, nil)
			},
			expectedObs: exectypes.Observation{
				Messages: exectypes.MessageObservations{
					src1: {
						1: NewMessage(1, 1, int(src1), int(dest)),
						// pseudo deleted
						2: NewMessage(2, 2, 0, 0),
						3: NewMessage(3, 3, 0, 0),
					},
				},
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 3),
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: createHashesMap(1, 3),
				},
				TokenData: exectypes.TokenDataObservations{
					src1: createTokenData(1, 3),
				},
			},
			expectedError: false,
		},
		{
			name: "encoding size exceeded mid second report",
			commitData: []exectypes.CommitData{
				createCommitData(src1, 1, 2),
				createCommitData(src2, 1, 4),
			},
			setupMocks: func(ccipReader *readerpkg_mock.MockCCIPReader,
				estimateProvider *ccipocr3.MockEstimateProvider,
				inflightCache *cache.InflightMessageCache,
				codec *codec_mock.MockExecCodec,
			) {
				messages1 := createMessages(src1, dest, 1, 2)
				messages2 := createMessages(src2, dest, 1, 4)

				// Return messages for both chains
				ccipReader.On("MsgsBetweenSeqNums", ctx, src1, cciptypes.NewSeqNumRange(1, 2)).
					Return(messages1, nil)
				ccipReader.On("MsgsBetweenSeqNums", ctx, src2, cciptypes.NewSeqNumRange(1, 4)).
					Return(messages2, nil)

				// First 4 calls to EncodeObservation return small sizes
				codec.EXPECT().EncodeObservation(mock.Anything).
					Return(oneByte, nil).
					Times(4)
				// Fifth call to EncodeObservation returns a large size
				codec.EXPECT().EncodeObservation(mock.Anything).
					Return(make([]byte, lenientMaxObservationLength+1), nil)
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
						// pseudo deleted
						3: NewMessage(3, 3, 0, 0),
						4: NewMessage(4, 4, 0, 0),
					},
				},
				CommitReports: exectypes.CommitObservations{
					src1: []exectypes.CommitData{
						createCommitData(src1, 1, 2),
					},
					src2: []exectypes.CommitData{
						createCommitData(src2, 1, 4),
					},
				},
				Hashes: exectypes.MessageHashes{
					src1: createHashesMap(1, 2),
					src2: createHashesMap(1, 4),
				},
				TokenData: exectypes.TokenDataObservations{
					src1: createTokenData(1, 2),
					src2: createTokenData(1, 4),
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
					src1: createHashesMap(1, 2),
				},
				TokenData: exectypes.TokenDataObservations{
					src1: createTokenData(1, 2),
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
			tokenDataObserver := observer.NoopTokenDataObserver{}

			plugin := &Plugin{
				lggr:                 mocks.NullLogger,
				ccipReader:           ccipReader,
				msgHasher:            msgHasher,
				ocrTypeCodec:         codec,
				estimateProvider:     estimateProvider,
				inflightMessageCache: inflightCache,
				tokenDataObserver:    &tokenDataObserver,
				offchainCfg: pluginconfig.ExecuteOffchainConfig{
					BatchGasLimit: uint64(batchGasLimit),
				},
			}

			tt.setupMocks(ccipReader, estimateProvider, inflightCache, codec)

			prevOutcome := exectypes.Outcome{
				CommitReports: tt.commitData,
			}
			obs, err := plugin.getMessagesObservation(ctx, plugin.lggr, prevOutcome, exectypes.Observation{})

			if tt.expectedError {
				require.Error(t, err)
				if tt.errorSubstring != "" {
					require.Contains(t, err.Error(), tt.errorSubstring)
				}
			} else {
				require.NoError(t, err)
				require.Equal(t, tt.expectedObs, obs)
			}
		})
	}
}
