package executor

import (
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-modsec/libmodsec/pkg/modsectypes"
)

func TestTimedMessageChannel(t *testing.T) {
	// Create a new timed message channel
	tmc := NewTimedMessageChannel(10)
	defer tmc.Close()

	// Create a test message
	msg := modsectypes.Message{
		Header: modsectypes.Header{
			MessageID: [32]byte{1, 2, 3, 4},
		},
	}

	// Send a message with a 100ms tick
	tmc.SendMessage(msg, 100*time.Millisecond)

	// Wait a bit less than the tick to ensure it hasn't been delivered yet
	time.Sleep(50 * time.Millisecond)

	// Check that no message is available yet
	select {
	case <-tmc.Messages():
		t.Fatal("Message should not be available yet")
	default:
		// Expected - no message yet
	}

	// Wait for the tick to complete
	time.Sleep(100 * time.Millisecond)

	// Now the message should be available
	select {
	case receivedMsg := <-tmc.Messages():
		if receivedMsg.Header.MessageID != msg.Header.MessageID {
			t.Fatalf("Expected message ID %v, got %v", msg.Header.MessageID, receivedMsg.Header.MessageID)
		}
	default:
		t.Fatal("Message should be available now")
	}
}

func TestTimedMessageChannelMultipleMessages(t *testing.T) {
	tmc := NewTimedMessageChannel(10)
	defer tmc.Close()

	// Create multiple messages with different ticks
	msg1 := modsectypes.Message{Header: modsectypes.Header{MessageID: [32]byte{1}}}
	msg2 := modsectypes.Message{Header: modsectypes.Header{MessageID: [32]byte{2}}}
	msg3 := modsectypes.Message{Header: modsectypes.Header{MessageID: [32]byte{3}}}

	// Send messages with different ticks (order: 3, 1, 2)
	tmc.SendMessage(msg3, 300*time.Millisecond)
	tmc.SendMessage(msg1, 100*time.Millisecond)
	tmc.SendMessage(msg2, 200*time.Millisecond)

	expectedOrder := []modsectypes.Message{msg1, msg2, msg3}
	for i, expectedMsg := range expectedOrder {
		select {
		case receivedMsg := <-tmc.Messages():
			if receivedMsg.Header.MessageID != expectedMsg.Header.MessageID {
				t.Fatalf("Expected message #%d ID %v, got %v", i+1, expectedMsg.Header.MessageID, receivedMsg.Header.MessageID)
			}
		case <-time.After(500 * time.Millisecond):
			t.Fatalf("Timed out waiting for message #%d", i+1)
		}
	}
}

func TestTimedMessageChannelClose(t *testing.T) {
	tmc := NewTimedMessageChannel(10)

	// Send a message
	msg := modsectypes.Message{Header: modsectypes.Header{MessageID: [32]byte{1}}}
	tmc.SendMessage(msg, 100*time.Millisecond)

	// Close immediately
	tmc.Close()

	// Try to send another message (should not panic)
	tmc.SendMessage(msg, 100*time.Millisecond)

	// Messages channel should be closed
	select {
	case _, ok := <-tmc.Messages():
		if ok {
			t.Fatal("Messages channel should be closed")
		}
	default:
		t.Fatal("Messages channel should be closed")
	}
}
