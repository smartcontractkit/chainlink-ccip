package storageaccess

import (
	"context"
	"math/big"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/v2/protocol/common"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

func createTestMessage(messageID [32]byte, seqNum cciptypes.SeqNum, sourceChainSelector, destChainSelector cciptypes.ChainSelector) common.Any2AnyVerifierMessage {
	return common.Any2AnyVerifierMessage{
		Header: common.MessageHeader{
			MessageID:           messageID,
			SourceChainSelector: sourceChainSelector,
			DestChainSelector:   destChainSelector,
			SequenceNumber:      seqNum,
		},
		Sender:         common.UnknownAddress([]byte("0x9999")),
		OnRampAddress:  common.UnknownAddress([]byte("0x8888")),
		Data:           []byte("test data"),
		Receiver:       common.UnknownAddress([]byte("0x7777")),
		FeeToken:       common.UnknownAddress([]byte("0x6666")),
		FeeTokenAmount: big.NewInt(1000),
		FeeValueJuels:  big.NewInt(500),
		TokenTransfer: common.TokenTransfer{
			SourceTokenAddress: common.UnknownAddress([]byte("0x5555")),
			DestTokenAddress:   common.UnknownAddress([]byte("0x4444")),
			ExtraData:          []byte("token data"),
			Amount:             big.NewInt(2000),
		},
		VerifierReceipts: []common.Receipt{
			{
				Issuer:            common.UnknownAddress([]byte("0x3333")),
				FeeTokenAmount:    big.NewInt(100),
				DestGasLimit:      50000,
				DestBytesOverhead: 1024,
				ExtraArgs:         []byte("receipt args"),
			},
		},
		ExecutorReceipt: &common.Receipt{
			Issuer:            common.UnknownAddress([]byte("0x2222")),
			FeeTokenAmount:    big.NewInt(200),
			DestGasLimit:      60000,
			DestBytesOverhead: 2048,
			ExtraArgs:         []byte("executor args"),
		},
		TokenReceipt: &common.Receipt{
			Issuer:            common.UnknownAddress([]byte("0x1111")),
			FeeTokenAmount:    big.NewInt(300),
			DestGasLimit:      70000,
			DestBytesOverhead: 4096,
			ExtraArgs:         []byte("token args"),
		},
		ExtraArgs: []byte("extra args"),
	}
}

func TestInMemoryOffchainStorage_StoreCCVData(t *testing.T) {
	lggr := logger.Test(t)
	storage := NewInMemoryOffchainStorage(lggr)

	ctx := context.Background()
	verifierAddress := []byte("0x1234")

	// Create test CCV data
	testData := []common.CCVData{
		{
			MessageID:             [32]byte{1, 2, 3},
			SequenceNumber:        100,
			SourceChainSelector:   1,
			DestChainSelector:     2,
			SourceVerifierAddress: verifierAddress,
			DestVerifierAddress:   []byte("0x4567"),
			CCVData:               []byte("signature1"),
			BlobData:              []byte("blob1"),
			Timestamp:             time.Now().Unix(),
			Message:               createTestMessage([32]byte{1, 2, 3}, 100, 1, 2),
		},
		{
			MessageID:             [32]byte{4, 5, 6},
			SequenceNumber:        101,
			SourceChainSelector:   1,
			DestChainSelector:     2,
			SourceVerifierAddress: verifierAddress,
			DestVerifierAddress:   []byte("0x4567"),
			CCVData:               []byte("signature2"),
			BlobData:              []byte("blob2"),
			Timestamp:             time.Now().Unix() + 1,
			Message:               createTestMessage([32]byte{4, 5, 6}, 101, 1, 2),
		},
	}

	// Store data
	err := storage.StoreCCVData(ctx, testData)
	require.NoError(t, err)

	// Retrieve and verify
	storedData, err := storage.GetAllCCVData(verifierAddress)
	require.NoError(t, err)
	require.Len(t, storedData, 2)

	// Verify data is sorted by timestamp
	require.True(t, storedData[0].Timestamp <= storedData[1].Timestamp)

	// Verify content
	require.Equal(t, testData[0].MessageID, storedData[0].MessageID)
	require.Equal(t, testData[1].MessageID, storedData[1].MessageID)
}

func TestInMemoryOffchainStorage_GetCCVDataByTimestamp(t *testing.T) {
	lggr := logger.Test(t)

	// Use fixed time provider for predictable tests
	baseTime := time.Now().UnixMicro()
	timeProvider := func() int64 { return baseTime }
	storage := NewInMemoryOffchainStorageWithTimeProvider(lggr, timeProvider)

	ctx := context.Background()
	verifierAddress := []byte("0x1234")

	// Create test data with different timestamps by storing at different times
	testData1 := []common.CCVData{
		{
			MessageID:             [32]byte{1},
			SequenceNumber:        100,
			SourceChainSelector:   1,
			DestChainSelector:     2,
			SourceVerifierAddress: verifierAddress,
			DestVerifierAddress:   []byte("0x4567"),
			CCVData:               []byte("sig1"),
			Message:               createTestMessage([32]byte{1}, 100, 1, 2),
		},
	}

	// Store first data at baseTime
	err := storage.StoreCCVData(ctx, testData1)
	require.NoError(t, err)

	// Update time provider for second batch
	timeProvider = func() int64 { return baseTime + 10000000 } // 10 seconds later in microseconds
	storage.timeProvider = timeProvider

	testData2 := []common.CCVData{
		{
			MessageID:             [32]byte{2},
			SequenceNumber:        101,
			SourceChainSelector:   1,
			DestChainSelector:     2,
			SourceVerifierAddress: verifierAddress,
			DestVerifierAddress:   []byte("0x4567"),
			CCVData:               []byte("sig2"),
			Message:               createTestMessage([32]byte{2}, 101, 1, 2),
		},
	}

	err = storage.StoreCCVData(ctx, testData2)
	require.NoError(t, err)

	// Update time provider for third batch
	timeProvider = func() int64 { return baseTime + 20000000 } // 20 seconds later in microseconds
	storage.timeProvider = timeProvider

	testData3 := []common.CCVData{
		{
			MessageID:             [32]byte{3},
			SequenceNumber:        102,
			SourceChainSelector:   1,
			DestChainSelector:     2,
			SourceVerifierAddress: verifierAddress,
			DestVerifierAddress:   []byte("0x4567"),
			CCVData:               []byte("sig3"),
			Message:               createTestMessage([32]byte{3}, 102, 1, 2),
		},
	}

	err = storage.StoreCCVData(ctx, testData3)
	require.NoError(t, err)

	tests := []struct {
		name          string
		startTime     int64
		destChains    []cciptypes.ChainSelector
		sourceChains  []cciptypes.ChainSelector
		limit         int
		offset        int
		expectedCount int
		expectedSeqs  []cciptypes.SeqNum
	}{
		{
			name:          "all data",
			startTime:     baseTime - 1,
			destChains:    []cciptypes.ChainSelector{2},
			sourceChains:  []cciptypes.ChainSelector{1},
			limit:         100,
			offset:        0,
			expectedCount: 3,
			expectedSeqs:  []cciptypes.SeqNum{100, 101, 102},
		},
		{
			name:          "middle range",
			startTime:     baseTime + 5000000, // 5 seconds later
			destChains:    []cciptypes.ChainSelector{2},
			sourceChains:  []cciptypes.ChainSelector{1},
			limit:         100,
			offset:        0,
			expectedCount: 2,
			expectedSeqs:  []cciptypes.SeqNum{101, 102},
		},
		{
			name:          "no data in range",
			startTime:     baseTime + 30000000, // 30 seconds later
			destChains:    []cciptypes.ChainSelector{2},
			sourceChains:  []cciptypes.ChainSelector{1},
			limit:         100,
			offset:        0,
			expectedCount: 0,
			expectedSeqs:  nil,
		},
		{
			name:          "pagination test - first page",
			startTime:     baseTime - 1,
			destChains:    []cciptypes.ChainSelector{2},
			sourceChains:  []cciptypes.ChainSelector{1},
			limit:         2,
			offset:        0,
			expectedCount: 2,
			expectedSeqs:  []cciptypes.SeqNum{100, 101},
		},
		{
			name:          "pagination test - second page",
			startTime:     baseTime - 1,
			destChains:    []cciptypes.ChainSelector{2},
			sourceChains:  []cciptypes.ChainSelector{1},
			limit:         2,
			offset:        2,
			expectedCount: 1,
			expectedSeqs:  []cciptypes.SeqNum{102},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := storage.GetCCVDataByTimestamp(ctx, tt.destChains, tt.startTime, tt.sourceChains, tt.limit, tt.offset)
			require.NoError(t, err)
			require.NotNil(t, response)

			require.Equal(t, tt.expectedCount, response.TotalCount)

			// Verify sequence numbers match expected by collecting from all destination chains
			var actualSeqs []cciptypes.SeqNum
			for _, ccvList := range response.Data {
				for _, ccv := range ccvList {
					actualSeqs = append(actualSeqs, ccv.SequenceNumber)
				}
			}

			// Sort actual sequences for comparison
			sort.Slice(actualSeqs, func(i, j int) bool {
				return actualSeqs[i] < actualSeqs[j]
			})

			require.Equal(t, tt.expectedSeqs, actualSeqs)
		})
	}
}

func TestInMemoryOffchainStorage_GetCCVDataByMessageID(t *testing.T) {
	lggr := logger.Test(t)
	storage := NewInMemoryOffchainStorage(lggr)

	ctx := context.Background()
	verifierAddress := []byte("0x1234")

	messageID := [32]byte{1, 2, 3, 4, 5}
	testData := []common.CCVData{
		{
			MessageID:             messageID,
			SequenceNumber:        100,
			SourceChainSelector:   1,
			DestChainSelector:     2,
			SourceVerifierAddress: verifierAddress,
			DestVerifierAddress:   []byte("0x4567"),
			CCVData:               []byte("signature"),
			BlobData:              []byte("blob"),
			Timestamp:             time.Now().Unix(),
			Message:               createTestMessage(messageID, 100, 1, 2),
		},
	}

	// Store data
	err := storage.StoreCCVData(ctx, testData)
	require.NoError(t, err)

	// Test finding existing message
	result, err := storage.GetCCVDataByMessageID(messageID)
	require.NoError(t, err)
	require.Equal(t, cciptypes.Bytes32(messageID), result.MessageID)
	require.Equal(t, cciptypes.SeqNum(100), result.SequenceNumber)

	// Test finding non-existing message
	nonExistentID := [32]byte{9, 9, 9}
	result, err = storage.GetCCVDataByMessageID(nonExistentID)
	require.Error(t, err)
	require.Nil(t, result)
	require.Contains(t, err.Error(), "CCV data not found")
}

func TestInMemoryOffchainStorage_MultipleVerifiers(t *testing.T) {
	lggr := logger.Test(t)
	storage := NewInMemoryOffchainStorage(lggr)

	ctx := context.Background()
	verifier1 := []byte("0x1111")
	verifier2 := []byte("0x2222")

	// Create data for different verifiers
	data1 := []common.CCVData{
		{
			MessageID:             [32]byte{1},
			SequenceNumber:        100,
			SourceVerifierAddress: verifier1,
			CCVData:               []byte("sig1"),
			Message:               createTestMessage([32]byte{1}, 100, 1, 2),
		},
		{
			MessageID:             [32]byte{2},
			SequenceNumber:        101,
			SourceVerifierAddress: verifier1,
			CCVData:               []byte("sig2"),
			Message:               createTestMessage([32]byte{2}, 101, 1, 2),
		},
	}

	data2 := []common.CCVData{
		{
			MessageID:             [32]byte{3},
			SequenceNumber:        200,
			SourceVerifierAddress: verifier2,
			CCVData:               []byte("sig3"),
			Message:               createTestMessage([32]byte{3}, 200, 1, 2),
		},
	}

	// Store data for both verifiers
	err := storage.StoreCCVData(ctx, data1)
	require.NoError(t, err)

	err = storage.StoreCCVData(ctx, data2)
	require.NoError(t, err)

	// Verify verifier1 data
	result1, err := storage.GetAllCCVData(verifier1)
	require.NoError(t, err)
	require.Len(t, result1, 2)

	// Verify verifier2 data
	result2, err := storage.GetAllCCVData(verifier2)
	require.NoError(t, err)
	require.Len(t, result2, 1)

	// Verify data is separate
	require.Equal(t, cciptypes.SeqNum(100), result1[0].SequenceNumber)
	require.Equal(t, cciptypes.SeqNum(101), result1[1].SequenceNumber)
	require.Equal(t, cciptypes.SeqNum(200), result2[0].SequenceNumber)
}

func TestInMemoryOffchainStorage_Clear(t *testing.T) {
	lggr := logger.Test(t)
	storage := NewInMemoryOffchainStorage(lggr)

	ctx := context.Background()
	verifierAddress := []byte("0x1234")

	// Add some data
	testData := []common.CCVData{
		{
			MessageID:             [32]byte{1, 2, 3},
			SourceVerifierAddress: verifierAddress,
			CCVData:               []byte("signature"),
			Message:               createTestMessage([32]byte{1, 2, 3}, 100, 1, 2),
		},
	}

	err := storage.StoreCCVData(ctx, testData)
	require.NoError(t, err)

	// Verify data exists
	result, err := storage.GetAllCCVData(verifierAddress)
	require.NoError(t, err)
	require.Len(t, result, 1)

	// Clear storage
	storage.Clear()

	// Verify data is gone
	result, err = storage.GetAllCCVData(verifierAddress)
	require.NoError(t, err)
	require.Len(t, result, 0)

	// Verify stats are reset
	stats := storage.GetStats()
	require.Equal(t, 0, stats["totalEntries"])
	require.Equal(t, 0, stats["verifierCount"])
}

func TestInMemoryOffchainStorage_EmptyData(t *testing.T) {
	lggr := logger.Test(t)
	storage := NewInMemoryOffchainStorage(lggr)

	ctx := context.Background()

	// Store empty data should not error
	err := storage.StoreCCVData(ctx, []common.CCVData{})
	require.NoError(t, err)

	// Get data for non-existent verifier
	verifierAddress := []byte("0x1234")
	result, err := storage.GetAllCCVData(verifierAddress)
	require.NoError(t, err)
	require.Len(t, result, 0)

	// Get data by timestamp for non-existent verifier
	timestampResult, err := storage.GetCCVDataByTimestamp(
		ctx,
		[]cciptypes.ChainSelector{2}, // dest chains
		0,                            // start timestamp
		[]cciptypes.ChainSelector{1}, // source chains
		100,                          // limit
		0,                            // offset
	)
	require.NoError(t, err)
	require.Equal(t, 0, timestampResult.TotalCount)
}

func TestInMemoryOffchainStorage_TimestampHandling(t *testing.T) {
	lggr := logger.Test(t)
	storage := NewInMemoryOffchainStorage(lggr)

	ctx := context.Background()
	verifierAddress := []byte("0x1234")

	// Create data - timestamp will be set by storage layer
	testData := []common.CCVData{
		{
			MessageID:             [32]byte{1},
			SourceVerifierAddress: verifierAddress,
			Message:               createTestMessage([32]byte{1}, 100, 1, 2),
		},
	}

	err := storage.StoreCCVData(ctx, testData)
	require.NoError(t, err)

	// Verify data was stored - timestamp is managed internally by storage entries
	result, err := storage.GetAllCCVData(verifierAddress)
	require.NoError(t, err)
	require.Len(t, result, 1)

	// The storage layer manages timestamps internally, so we just verify data was stored
	require.Equal(t, cciptypes.Bytes32([32]byte{1}), result[0].MessageID)
}

func TestInMemoryOffchainStorage_ReaderWriterViews(t *testing.T) {
	lggr := logger.Test(t)
	storage := NewInMemoryOffchainStorage(lggr)

	// Create reader and writer views
	reader := CreateReaderOnly(storage)
	writer := CreateWriterOnly(storage)

	ctx := context.Background()
	verifierAddress := []byte("0x1234")

	// Test data
	testData := []common.CCVData{
		{
			MessageID:             [32]byte{1, 2, 3},
			SequenceNumber:        100,
			SourceChainSelector:   1,
			DestChainSelector:     2,
			SourceVerifierAddress: verifierAddress,
			DestVerifierAddress:   []byte("0x4567"),
			CCVData:               []byte("signature1"),
			BlobData:              []byte("blob1"),
			Message:               createTestMessage([32]byte{1, 2, 3}, 100, 1, 2),
		},
	}

	// Store data using writer view
	err := writer.StoreCCVData(ctx, testData)
	require.NoError(t, err)

	// Read data using reader view
	response, err := reader.GetCCVDataByTimestamp(
		ctx,
		[]cciptypes.ChainSelector{2}, // dest chains
		0,                            // start timestamp
		[]cciptypes.ChainSelector{1}, // source chains
		100,                          // limit
		0,                            // offset
	)
	require.NoError(t, err)
	require.NotNil(t, response)
	require.Equal(t, 1, response.TotalCount)

	// Verify data is organized by destination chain
	ccvList, exists := response.Data[2]
	require.True(t, exists)
	require.Len(t, ccvList, 1)
	require.Equal(t, cciptypes.Bytes32([32]byte{1, 2, 3}), ccvList[0].MessageID)
}

func TestInMemoryOffchainStorage_DestinationChainOrganization(t *testing.T) {
	lggr := logger.Test(t)
	storage := NewInMemoryOffchainStorage(lggr)

	ctx := context.Background()
	verifierAddress := []byte("0x1234")

	// Create data for different destination chains
	testData := []common.CCVData{
		{
			MessageID:             [32]byte{1},
			SequenceNumber:        100,
			SourceChainSelector:   1,
			DestChainSelector:     10, // Chain 10
			SourceVerifierAddress: verifierAddress,
			CCVData:               []byte("sig1"),
			Message:               createTestMessage([32]byte{1}, 100, 1, 10),
		},
		{
			MessageID:             [32]byte{2},
			SequenceNumber:        101,
			SourceChainSelector:   1,
			DestChainSelector:     20, // Chain 20
			SourceVerifierAddress: verifierAddress,
			CCVData:               []byte("sig2"),
			Message:               createTestMessage([32]byte{2}, 101, 1, 20),
		},
		{
			MessageID:             [32]byte{3},
			SequenceNumber:        102,
			SourceChainSelector:   1,
			DestChainSelector:     10, // Chain 10 again
			SourceVerifierAddress: verifierAddress,
			CCVData:               []byte("sig3"),
			Message:               createTestMessage([32]byte{3}, 102, 1, 10),
		},
	}

	// Store data
	err := storage.StoreCCVData(ctx, testData)
	require.NoError(t, err)

	// Query for both destination chains
	response, err := storage.GetCCVDataByTimestamp(
		ctx,
		[]cciptypes.ChainSelector{10, 20}, // Both dest chains
		0,                                 // start timestamp
		[]cciptypes.ChainSelector{1},      // source chain
		100,                               // limit
		0,                                 // offset
	)
	require.NoError(t, err)
	require.NotNil(t, response)
	require.Equal(t, 3, response.TotalCount)

	// Verify data is organized by destination chain
	chain10Data, exists := response.Data[10]
	require.True(t, exists)
	require.Len(t, chain10Data, 2) // Two messages for chain 10

	chain20Data, exists := response.Data[20]
	require.True(t, exists)
	require.Len(t, chain20Data, 1) // One message for chain 20

	// Query for only chain 10
	response, err = storage.GetCCVDataByTimestamp(
		ctx,
		[]cciptypes.ChainSelector{10}, // Only chain 10
		0,                             // start timestamp
		[]cciptypes.ChainSelector{1},  // source chain
		100,                           // limit
		0,                             // offset
	)
	require.NoError(t, err)
	require.NotNil(t, response)
	require.Equal(t, 2, response.TotalCount)

	// Should only have chain 10 data
	_, exists = response.Data[10]
	require.True(t, exists)
	_, exists = response.Data[20]
	require.False(t, exists)
}
