package verifier_test

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"math/big"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/v2/protocol/common"
	"github.com/smartcontractkit/chainlink-ccip/v2/protocol/common/storageaccess"
	"github.com/smartcontractkit/chainlink-ccip/v2/protocol/verifier"
	"github.com/smartcontractkit/chainlink-ccip/v2/protocol/verifier/mocks"
)

// Test constants
const (
	defaultProcessingChannelSize = 10
	defaultProcessingTimeout     = time.Second
	defaultMaxBatchSize          = 100
	defaultDestChain             = cciptypes.ChainSelector(100)
	sourceChain1                 = cciptypes.ChainSelector(42)
	sourceChain2                 = cciptypes.ChainSelector(84)
	unconfiguredChain            = cciptypes.ChainSelector(999)
)

// testSetup contains common test dependencies
type testSetup struct {
	t       *testing.T
	ctx     context.Context
	cancel  context.CancelFunc
	logger  logger.Logger
	storage *storageaccess.InMemoryOffchainStorage
	signer  verifier.MessageSigner
}

// mockSourceReaderSetup contains a mock source reader and its channel
type mockSourceReaderSetup struct {
	reader  *mocks.MockSourceReader
	channel chan common.VerificationTask
}

// newTestSetup creates common test dependencies
func newTestSetup(t *testing.T) *testSetup {
	ctx, cancel := context.WithCancel(context.Background())
	lggr := logger.Test(t)
	storage := storageaccess.NewInMemoryOffchainStorage(lggr)
	signer := createTestSigner(t)

	return &testSetup{
		t:       t,
		ctx:     ctx,
		cancel:  cancel,
		logger:  lggr,
		storage: storage,
		signer:  signer,
	}
}

// cleanup should be called in defer
func (ts *testSetup) cleanup() {
	ts.cancel()
}

// createTestSigner generates a test ECDSA message signer
func createTestSigner(t *testing.T) verifier.MessageSigner {
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	require.NoError(t, err)
	privateKeyBytes := crypto.FromECDSA(privateKey)
	signer, err := verifier.NewECDSAMessageSigner(privateKeyBytes)
	require.NoError(t, err)
	return signer
}

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

func createTestVerificationTask(messageID [32]byte, seqNum cciptypes.SeqNum, sourceChainSelector, destChainSelector cciptypes.ChainSelector) common.VerificationTask {
	message := createTestMessage(messageID, seqNum, sourceChainSelector, destChainSelector)
	return common.VerificationTask{
		Message:      message,
		ReceiptBlobs: [][]byte{[]byte("receipt blob 1"), []byte("receipt blob 2")},
	}
}

// createCoordinatorConfig creates a coordinator config with the given sources
func createCoordinatorConfig(coordinatorID string, sources map[cciptypes.ChainSelector]string) verifier.CoordinatorConfig {
	sourceConfigs := make(map[cciptypes.ChainSelector]verifier.SourceConfig)
	for chainSelector, address := range sources {
		sourceConfigs[chainSelector] = verifier.SourceConfig{
			VerifierAddress: common.UnknownAddress([]byte(address)),
		}
	}

	return verifier.CoordinatorConfig{
		VerifierID:            coordinatorID,
		SourceConfigs:         sourceConfigs,
		ProcessingChannelSize: defaultProcessingChannelSize,
		ProcessingTimeout:     defaultProcessingTimeout,
		MaxBatchSize:          defaultMaxBatchSize,
	}
}

// setupMockSourceReader creates a mock source reader with expectations
func setupMockSourceReader(t *testing.T, shouldClose bool) *mockSourceReaderSetup {
	mockReader := mocks.NewMockSourceReader(t)
	channel := make(chan common.VerificationTask, 10)

	mockReader.EXPECT().Start(mock.Anything).Return(nil)
	mockReader.EXPECT().VerificationTaskChannel().Return((<-chan common.VerificationTask)(channel))

	if shouldClose {
		mockReader.EXPECT().Stop().Run(func() {
			close(channel)
		}).Return(nil)
	} else {
		mockReader.EXPECT().Stop().Return(nil)
	}

	return &mockSourceReaderSetup{
		reader:  mockReader,
		channel: channel,
	}
}

// createVerificationCoordinator creates a verification coordinator with the given setup
func createVerificationCoordinator(ts *testSetup, config verifier.CoordinatorConfig, sourceReaders map[cciptypes.ChainSelector]verifier.SourceReader) (*verifier.VerificationCoordinator, error) {
	commitVerifier := verifier.NewCommitVerifier(config, ts.signer, ts.logger)

	return verifier.NewVerificationCoordinator(
		verifier.WithConfig(config),
		verifier.WithSourceReaders(sourceReaders),
		verifier.WithVerifier(commitVerifier),
		verifier.WithStorage(ts.storage),
		verifier.WithLogger(ts.logger),
	)
}

// waitForMessages waits for the specified number of messages to be processed
func waitForMessages(ts *testSetup, count int) {
	for i := 0; i < count; i++ {
		err := ts.storage.WaitForStore(ts.ctx)
		require.NoError(ts.t, err)
	}
}

// sendTasksAsync sends verification tasks asynchronously with a delay
func sendTasksAsync(tasks []common.VerificationTask, channel chan<- common.VerificationTask, counter *atomic.Int32, delay time.Duration) {
	go func() {
		for _, task := range tasks {
			channel <- task
			if counter != nil {
				counter.Add(1)
			}
			time.Sleep(delay)
		}
	}()
}

// verifyStoredTasks is a helper to verify stored data matches expected tasks
func verifyStoredTasks(t *testing.T, storedData []common.CCVData, expectedTasks []common.VerificationTask, expectedChain cciptypes.ChainSelector) {
	expectedIDs := make(map[[32]byte]bool)
	for _, task := range expectedTasks {
		expectedIDs[task.Message.Header.MessageID] = true
	}
	for _, data := range storedData {
		assert.True(t, expectedIDs[data.MessageID], "Unexpected message ID: %x", data.MessageID)
		assert.Equal(t, expectedChain, data.SourceChainSelector)
	}
}

func TestVerifier(t *testing.T) {
	ts := newTestSetup(t)
	defer ts.cleanup()

	config := createCoordinatorConfig("test-custom-mockery-verifier", map[cciptypes.ChainSelector]string{
		sourceChain1: "0x1234",
	})

	// Set up mock source reader
	mockSetup := setupMockSourceReader(t, true)
	sourceReaders := map[cciptypes.ChainSelector]verifier.SourceReader{
		sourceChain1: mockSetup.reader,
	}

	// Create and start verifier
	v, err := createVerificationCoordinator(ts, config, sourceReaders)
	require.NoError(t, err)

	err = v.Start(ts.ctx)
	require.NoError(t, err)

	// Create and send test tasks
	testTasks := []common.VerificationTask{
		createTestVerificationTask([32]byte{1, 2, 3}, 100, sourceChain1, defaultDestChain),
		createTestVerificationTask([32]byte{4, 5, 6}, 200, sourceChain1, defaultDestChain),
	}

	var messagesSent atomic.Int32
	sendTasksAsync(testTasks, mockSetup.channel, &messagesSent, 10*time.Millisecond)

	// Wait for processing and verify results
	waitForMessages(ts, len(testTasks))

	err = v.Stop()
	require.NoError(t, err)

	// Verify stored data
	storedData, err := ts.storage.GetAllCCVData(config.SourceConfigs[sourceChain1].VerifierAddress)
	require.NoError(t, err)
	assert.Len(t, storedData, len(testTasks))
	assert.Equal(t, int(messagesSent.Load()), len(testTasks))

	// Verify message IDs
	expectedIDs := make(map[[32]byte]bool)
	for _, task := range testTasks {
		expectedIDs[task.Message.Header.MessageID] = true
	}
	for _, data := range storedData {
		assert.True(t, expectedIDs[data.MessageID], "Unexpected message ID: %x", data.MessageID)
	}
}

func TestMultiSourceVerifier_TwoSources(t *testing.T) {
	ts := newTestSetup(t)
	defer ts.cleanup()

	config := createCoordinatorConfig("test-multi-source-verifier", map[cciptypes.ChainSelector]string{
		sourceChain1: "0x1234",
		sourceChain2: "0x5678",
	})

	// Set up mock source readers
	mockSetup1 := setupMockSourceReader(t, true)
	mockSetup2 := setupMockSourceReader(t, true)
	sourceReaders := map[cciptypes.ChainSelector]verifier.SourceReader{
		sourceChain1: mockSetup1.reader,
		sourceChain2: mockSetup2.reader,
	}

	// Create and start verifier
	v, err := createVerificationCoordinator(ts, config, sourceReaders)
	require.NoError(t, err)

	err = v.Start(ts.ctx)
	require.NoError(t, err)

	// Create test tasks for both sources
	tasksSource1 := []common.VerificationTask{
		createTestVerificationTask([32]byte{1, 1, 1}, 100, sourceChain1, defaultDestChain),
		createTestVerificationTask([32]byte{1, 2, 3}, 101, sourceChain1, defaultDestChain),
	}
	tasksSource2 := []common.VerificationTask{
		createTestVerificationTask([32]byte{2, 1, 1}, 200, sourceChain2, defaultDestChain),
		createTestVerificationTask([32]byte{2, 2, 3}, 201, sourceChain2, defaultDestChain),
	}

	// Send tasks from both sources
	var messagesSent1, messagesSent2 atomic.Int32
	sendTasksAsync(tasksSource1, mockSetup1.channel, &messagesSent1, 5*time.Millisecond)
	sendTasksAsync(tasksSource2, mockSetup2.channel, &messagesSent2, 7*time.Millisecond)

	// Wait for all messages to be processed
	totalMessages := len(tasksSource1) + len(tasksSource2)
	waitForMessages(ts, totalMessages)

	err = v.Stop()
	require.NoError(t, err)

	// Verify stored data for both sources
	storedDataSource1, err := ts.storage.GetAllCCVData(config.SourceConfigs[sourceChain1].VerifierAddress)
	require.NoError(t, err)
	storedDataSource2, err := ts.storage.GetAllCCVData(config.SourceConfigs[sourceChain2].VerifierAddress)
	require.NoError(t, err)

	assert.Len(t, storedDataSource1, len(tasksSource1))
	assert.Len(t, storedDataSource2, len(tasksSource2))
	assert.Equal(t, int(messagesSent1.Load()), len(tasksSource1))
	assert.Equal(t, int(messagesSent2.Load()), len(tasksSource2))

	// Verify message IDs and chain selectors
	verifyStoredTasks(t, storedDataSource1, tasksSource1, sourceChain1)
	verifyStoredTasks(t, storedDataSource2, tasksSource2, sourceChain2)
}

func TestMultiSourceVerifier_SingleSourceFailure(t *testing.T) {
	ts := newTestSetup(t)
	defer ts.cleanup()

	config := createCoordinatorConfig("test-failure-verifier", map[cciptypes.ChainSelector]string{
		sourceChain1: "0x1234",
		sourceChain2: "0x5678",
	})

	// Set up mock source readers - source 2 will fail by closing its channel immediately
	mockSetup1 := setupMockSourceReader(t, true)
	mockSetup2 := setupMockSourceReader(t, false)
	sourceReaders := map[cciptypes.ChainSelector]verifier.SourceReader{
		sourceChain1: mockSetup1.reader,
		sourceChain2: mockSetup2.reader,
	}

	// Create and start verifier
	v, err := createVerificationCoordinator(ts, config, sourceReaders)
	require.NoError(t, err)

	err = v.Start(ts.ctx)
	require.NoError(t, err)

	// Close source 2 channel immediately to simulate failure
	close(mockSetup2.channel)

	// Send verification tasks only to source 1
	tasksSource1 := []common.VerificationTask{
		createTestVerificationTask([32]byte{1, 1, 1}, 100, sourceChain1, defaultDestChain),
		createTestVerificationTask([32]byte{1, 2, 3}, 101, sourceChain1, defaultDestChain),
	}

	sendTasksAsync(tasksSource1, mockSetup1.channel, nil, 5*time.Millisecond)
	waitForMessages(ts, len(tasksSource1))

	err = v.Stop()
	require.NoError(t, err)

	// Verify only source 1 data was stored
	storedDataSource1, err := ts.storage.GetAllCCVData(config.SourceConfigs[sourceChain1].VerifierAddress)
	require.NoError(t, err)
	storedDataSource2, err := ts.storage.GetAllCCVData(config.SourceConfigs[sourceChain2].VerifierAddress)
	require.NoError(t, err)

	assert.Len(t, storedDataSource1, len(tasksSource1))
	assert.Len(t, storedDataSource2, 0) // No messages from failed source
}

func TestMultiSourceVerifier_ValidationErrors(t *testing.T) {
	ts := newTestSetup(t)
	defer ts.cleanup()

	tests := []struct {
		name        string
		config      verifier.CoordinatorConfig
		readers     map[cciptypes.ChainSelector]verifier.SourceReader
		expectError string
	}{
		{
			name:        "no source readers",
			config:      createCoordinatorConfig("test-no-sources", map[cciptypes.ChainSelector]string{}),
			readers:     map[cciptypes.ChainSelector]verifier.SourceReader{},
			expectError: "at least one source reader is required",
		},
		{
			name: "mismatched source config and readers",
			config: createCoordinatorConfig("test-mismatch", map[cciptypes.ChainSelector]string{
				sourceChain1: "0x1234",
				sourceChain2: "0x5678",
			}),
			readers: func() map[cciptypes.ChainSelector]verifier.SourceReader {
				// Create a mock that only expects VerificationTaskChannel call
				mockReader := mocks.NewMockSourceReader(t)
				mockCh := make(chan common.VerificationTask)
				mockReader.EXPECT().VerificationTaskChannel().Return((<-chan common.VerificationTask)(mockCh))
				return map[cciptypes.ChainSelector]verifier.SourceReader{
					sourceChain1: mockReader, // Missing reader for sourceChain2
				}
			}(),
			expectError: "source reader not found for chain selector 84",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := createVerificationCoordinator(ts, tt.config, tt.readers)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.expectError)
		})
	}
}

func TestMultiSourceVerifier_HealthCheck(t *testing.T) {
	ts := newTestSetup(t)
	defer ts.cleanup()

	config := createCoordinatorConfig("test-health-check", map[cciptypes.ChainSelector]string{
		sourceChain1: "0x1234",
		sourceChain2: "0x5678",
	})

	// Create mock source readers with health check expectations
	mockSetup1 := setupMockSourceReader(t, false)
	mockSetup2 := setupMockSourceReader(t, false)

	// Set up health check expectations - one healthy, one unhealthy
	mockSetup1.reader.EXPECT().HealthCheck(mock.Anything).Return(nil).Maybe()
	mockSetup2.reader.EXPECT().HealthCheck(mock.Anything).Return(assert.AnError).Maybe()

	sourceReaders := map[cciptypes.ChainSelector]verifier.SourceReader{
		sourceChain1: mockSetup1.reader,
		sourceChain2: mockSetup2.reader,
	}

	v, err := createVerificationCoordinator(ts, config, sourceReaders)
	require.NoError(t, err)

	// Start the verifier
	err = v.Start(ts.ctx)
	require.NoError(t, err)

	// Health check should fail if any source reader is unhealthy
	err = v.HealthCheck(ts.ctx)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "source reader unhealthy for chain")

	// Stop the verifier
	err = v.Stop()
	require.NoError(t, err)
}

func TestVerificationErrorHandling(t *testing.T) {
	ts := newTestSetup(t)
	defer ts.cleanup()

	// Create config with only one source chain configured
	config := createCoordinatorConfig("test-error-handling", map[cciptypes.ChainSelector]string{
		sourceChain1: "0x1234",
		// unconfiguredChain is intentionally not included in the config
	})

	// Set up mock source readers for both chains
	mockSetup1 := setupMockSourceReader(t, true)
	mockSetup2 := setupMockSourceReader(t, true)

	// Create source readers map that includes the unconfigured chain
	// This simulates having a reader for a chain that's not in the coordinator config
	sourceReaders := map[cciptypes.ChainSelector]verifier.SourceReader{
		sourceChain1:      mockSetup1.reader,
		unconfiguredChain: mockSetup2.reader,
	}

	// Create and start verifier - this should succeed even with extra readers
	v, err := createVerificationCoordinator(ts, config, sourceReaders)
	require.NoError(t, err)

	err = v.Start(ts.ctx)
	require.NoError(t, err)

	// Create test verification tasks
	validTask := createTestVerificationTask([32]byte{1, 1, 1}, 100, sourceChain1, defaultDestChain)
	invalidTask := createTestVerificationTask([32]byte{2, 2, 2}, 200, unconfiguredChain, defaultDestChain)

	// Send tasks
	sendTasksAsync([]common.VerificationTask{validTask}, mockSetup1.channel, nil, 10*time.Millisecond)
	sendTasksAsync([]common.VerificationTask{invalidTask}, mockSetup2.channel, nil, 10*time.Millisecond)

	// Wait for valid task to be processed
	waitForMessages(ts, 1)

	// Give some time for error processing
	time.Sleep(50 * time.Millisecond)

	err = v.Stop()
	require.NoError(t, err)

	// Verify results - only the configured source chain should have stored data
	storedData, err := ts.storage.GetAllCCVData(config.SourceConfigs[sourceChain1].VerifierAddress)
	require.NoError(t, err)
	assert.Len(t, storedData, 1)
	assert.Equal(t, validTask.Message.Header.MessageID, storedData[0].MessageID)

	// The unconfigured chain is not in the config, so we can't check its data
	// The test validates that tasks from unconfigured chains don't cause crashes
	// and that configured chains continue to work properly
}
