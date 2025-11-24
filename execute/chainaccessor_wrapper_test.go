package execute

import (
	"context"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	ccipocr3mocks "github.com/smartcontractkit/chainlink-ccip/mocks/chainlink_common/ccipocr3"
)

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
		// Setup mock using auto-generated mock
		mockAccessor := ccipocr3mocks.NewMockChainAccessor(t)
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
		// Setup mock using auto-generated mock
		mockAccessor := ccipocr3mocks.NewMockChainAccessor(t)
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
