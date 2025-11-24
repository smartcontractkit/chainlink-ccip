package execute

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/types/query/primitives"
)

// MockChainAccessor is a mock implementation of ChainAccessor for testing
type MockChainAccessor struct {
	mock.Mock
}

func (m *MockChainAccessor) MsgsBetweenSeqNums(
	ctx context.Context,
	chain cciptypes.ChainSelector,
	seqNumRange cciptypes.SeqNumRange,
) ([]cciptypes.Message, error) {
	args := m.Called(ctx, chain, seqNumRange)
	return args.Get(0).([]cciptypes.Message), args.Error(1)
}

func (m *MockChainAccessor) GetContractAddress(contractName string) ([]byte, error) {
	args := m.Called(contractName)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockChainAccessor) GetAllConfigsLegacy(
	ctx context.Context,
	destChainSelector cciptypes.ChainSelector,
	sourceChainSelectors []cciptypes.ChainSelector,
) (cciptypes.ChainConfigSnapshot, map[cciptypes.ChainSelector]cciptypes.SourceChainConfig, error) {
	args := m.Called(ctx, destChainSelector, sourceChainSelectors)
	return args.Get(0).(cciptypes.ChainConfigSnapshot),
		args.Get(1).(map[cciptypes.ChainSelector]cciptypes.SourceChainConfig),
		args.Error(2)
}

func (m *MockChainAccessor) GetChainFeeComponents(ctx context.Context) (cciptypes.ChainFeeComponents, error) {
	args := m.Called(ctx)
	return args.Get(0).(cciptypes.ChainFeeComponents), args.Error(1)
}

func (m *MockChainAccessor) Sync(
	ctx context.Context,
	contractName string,
	contractAddress cciptypes.UnknownAddress,
) error {
	args := m.Called(ctx, contractName, contractAddress)
	return args.Error(0)
}

func (m *MockChainAccessor) CommitReportsGTETimestamp(
	ctx context.Context,
	ts time.Time,
	confidence primitives.ConfidenceLevel,
	limit int,
) ([]cciptypes.CommitPluginReportWithMeta, error) {
	args := m.Called(ctx, ts, confidence, limit)
	return args.Get(0).([]cciptypes.CommitPluginReportWithMeta), args.Error(1)
}

func (m *MockChainAccessor) ExecutedMessages(
	ctx context.Context,
	rangesPerChain map[cciptypes.ChainSelector][]cciptypes.SeqNumRange,
	confidence primitives.ConfidenceLevel,
) (map[cciptypes.ChainSelector][]cciptypes.SeqNum, error) {
	args := m.Called(ctx, rangesPerChain, confidence)
	return args.Get(0).(map[cciptypes.ChainSelector][]cciptypes.SeqNum), args.Error(1)
}

func (m *MockChainAccessor) NextSeqNum(ctx context.Context, sources []cciptypes.ChainSelector) (
	map[cciptypes.ChainSelector]cciptypes.SeqNum, error) {
	args := m.Called(ctx, sources)
	return args.Get(0).(map[cciptypes.ChainSelector]cciptypes.SeqNum), args.Error(1)
}

func (m *MockChainAccessor) Nonces(
	ctx context.Context,
	addresses map[cciptypes.ChainSelector][]cciptypes.UnknownEncodedAddress,
) (map[cciptypes.ChainSelector]map[string]uint64, error) {
	args := m.Called(ctx, addresses)
	return args.Get(0).(map[cciptypes.ChainSelector]map[string]uint64), args.Error(1)
}

func (m *MockChainAccessor) GetChainFeePriceUpdate(
	ctx context.Context,
	selectors []cciptypes.ChainSelector,
) (map[cciptypes.ChainSelector]cciptypes.TimestampedUnixBig, error) {
	args := m.Called(ctx, selectors)
	return args.Get(0).(map[cciptypes.ChainSelector]cciptypes.TimestampedUnixBig), args.Error(1)
}

func (m *MockChainAccessor) GetLatestPriceSeqNr(ctx context.Context) (cciptypes.SeqNum, error) {
	args := m.Called(ctx)
	return args.Get(0).(cciptypes.SeqNum), args.Error(1)
}

func (m *MockChainAccessor) LatestMessageTo(
	ctx context.Context,
	dest cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
	args := m.Called(ctx, dest)
	return args.Get(0).(cciptypes.SeqNum), args.Error(1)
}

func (m *MockChainAccessor) GetExpectedNextSequenceNumber(
	ctx context.Context,
	dest cciptypes.ChainSelector,
) (cciptypes.SeqNum, error) {
	args := m.Called(ctx, dest)
	return args.Get(0).(cciptypes.SeqNum), args.Error(1)
}

func (m *MockChainAccessor) GetTokenPriceUSD(
	ctx context.Context,
	address cciptypes.UnknownAddress,
) (cciptypes.TimestampedUnixBig, error) {
	args := m.Called(ctx, address)
	return args.Get(0).(cciptypes.TimestampedUnixBig), args.Error(1)
}

func (m *MockChainAccessor) GetFeeQuoterDestChainConfig(
	ctx context.Context,
	dest cciptypes.ChainSelector,
) (cciptypes.FeeQuoterDestChainConfig, error) {
	args := m.Called(ctx, dest)
	return args.Get(0).(cciptypes.FeeQuoterDestChainConfig), args.Error(1)
}

func (m *MockChainAccessor) MessagesByTokenID(
	ctx context.Context,
	source, dest cciptypes.ChainSelector,
	tokens map[cciptypes.MessageTokenID]cciptypes.RampTokenAmount,
) (map[cciptypes.MessageTokenID]cciptypes.Bytes, error) {
	args := m.Called(ctx, source, dest, tokens)
	return args.Get(0).(map[cciptypes.MessageTokenID]cciptypes.Bytes), args.Error(1)
}

func (m *MockChainAccessor) GetFeedPricesUSD(
	ctx context.Context,
	tokens []cciptypes.UnknownEncodedAddress,
	tokenInfo map[cciptypes.UnknownEncodedAddress]cciptypes.TokenInfo,
) (cciptypes.TokenPriceMap, error) {
	args := m.Called(ctx, tokens, tokenInfo)
	return args.Get(0).(cciptypes.TokenPriceMap), args.Error(1)
}

func (m *MockChainAccessor) GetFeeQuoterTokenUpdates(
	ctx context.Context,
	tokensBytes []cciptypes.UnknownAddress,
) (map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedUnixBig, error) {
	args := m.Called(ctx, tokensBytes)
	return args.Get(0).(map[cciptypes.UnknownEncodedAddress]cciptypes.TimestampedUnixBig), args.Error(1)
}

func TestChainAccessorWrapper_PopulateTxHashEnabled(t *testing.T) {
	ctx := context.Background()
	chainSelector := cciptypes.ChainSelector(1)
	seqNumRange := cciptypes.NewSeqNumRange(1, 10)

	// Create test messages with TxHash
	messagesWithTxHash := []cciptypes.Message{
		{
			Header: cciptypes.RampMessageHeader{
				TxHash:         "0xabc123",
				SequenceNumber: 1,
			},
		},
		{
			Header: cciptypes.RampMessageHeader{
				TxHash:         "0xdef456",
				SequenceNumber: 2,
			},
		},
	}

	t.Run("TxHash is preserved when enabled", func(t *testing.T) {
		// Setup mock
		mockAccessor := new(MockChainAccessor)
		mockAccessor.On("MsgsBetweenSeqNums", ctx, chainSelector, seqNumRange).Return(messagesWithTxHash, nil)

		// Create wrapper with flag enabled
		wrapper := NewChainAccessorWrapper(mockAccessor, true)

		// Call MsgsBetweenSeqNums
		result, err := wrapper.MsgsBetweenSeqNums(ctx, chainSelector, seqNumRange)
		require.NoError(t, err)

		// Verify TxHash is preserved
		assert.Len(t, result, 2)
		assert.Equal(t, "0xabc123", result[0].Header.TxHash)
		assert.Equal(t, "0xdef456", result[1].Header.TxHash)
	})

	t.Run("TxHash is cleared when disabled", func(t *testing.T) {
		// Setup mock
		mockAccessor := new(MockChainAccessor)
		mockAccessor.On("MsgsBetweenSeqNums", ctx, chainSelector, seqNumRange).Return(messagesWithTxHash, nil)

		// Create wrapper with flag disabled
		wrapper := NewChainAccessorWrapper(mockAccessor, false)

		// Call MsgsBetweenSeqNums
		result, err := wrapper.MsgsBetweenSeqNums(ctx, chainSelector, seqNumRange)
		require.NoError(t, err)

		// Verify TxHash is cleared
		assert.Len(t, result, 2)
		assert.Equal(t, "", result[0].Header.TxHash)
		assert.Equal(t, "", result[1].Header.TxHash)
		// But other fields are preserved
		assert.Equal(t, cciptypes.SeqNum(1), result[0].Header.SequenceNumber)
		assert.Equal(t, cciptypes.SeqNum(2), result[1].Header.SequenceNumber)
	})
}
