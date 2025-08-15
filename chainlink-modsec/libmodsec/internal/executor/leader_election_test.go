package executor

import (
	"crypto/rand"
	"fmt"
	"testing"
	"time"
)

func TestIsLeader(t *testing.T) {
	deltaStage := 10 * time.Second

	// Create a consistent test setup
	participants := map[uint64][]string{
		1: {"node3", "node2", "node1"},
		2: {"node1", "node2", "node3"}, // Same participants as chain 1, different order
		3: {"node1", "node2", "node6"},
		4: {"node2", "node3", "node4"}, // Not a participant
	}

	le := NewModuloLeaderElection("node2", participants, deltaStage)

	// Test with a fixed message ID for deterministic results
	msgId := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	tests := []struct {
		name              string
		destChainSelector uint64
		expectedResult    bool
		expectedDelay     time.Duration
	}{
		{
			name:              "chain 1, is leader",
			destChainSelector: 1,
			expectedResult:    true,
			expectedDelay:     0 * time.Second,
		},
		{
			name:              "chain 2, different ordering of participants does not change leader",
			destChainSelector: 2,
			expectedResult:    true,
			expectedDelay:     0 * time.Second,
		},
		{
			name:              "chain 3, is not leader",
			destChainSelector: 3,
			expectedResult:    false,
			expectedDelay:     10 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, delay := le.IsLeader(msgId, tt.destChainSelector)
			if result != tt.expectedResult {
				t.Errorf("expected result %v, got %v", tt.expectedResult, result)
			}
			if delay != tt.expectedDelay {
				t.Errorf("expected delay %v, got %v", tt.expectedDelay, delay)
			}
		})
	}
}

func TestIsLeaderNonParticipant(t *testing.T) {
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3"},
	}

	le := NewModuloLeaderElection("node1", participants, 1*time.Second)
	msgId := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	// Test that non-participant chains always return false
	result, delay := le.IsLeader(msgId, 999)
	if result {
		t.Error("expected non-participant to not be leader")
	}
	if delay != 0 {
		t.Errorf("expected 0 delay for non-participant, got %v", delay)
	}
}

func TestIsLeaderDeterministic(t *testing.T) {
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3"},
	}

	le := NewModuloLeaderElection("node1", participants, 1*time.Second)
	msgId := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	// Test that the same input always produces the same result
	result1, delay1 := le.IsLeader(msgId, 1)
	result2, delay2 := le.IsLeader(msgId, 1)

	if result1 != result2 {
		t.Error("leader election should be deterministic")
	}
	if delay1 != delay2 {
		t.Error("delay should be deterministic")
	}
}

func TestGetParticipants(t *testing.T) {
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3"},
		2: {"node4", "node5"},
	}

	le := NewModuloLeaderElection("node1", participants, 1*time.Second)

	tests := []struct {
		name                 string
		destChainSelector    uint64
		expectedParticipants []string
	}{
		{
			name:                 "chain 1",
			destChainSelector:    1,
			expectedParticipants: []string{"node1", "node2", "node3"},
		},
		{
			name:                 "chain 2",
			destChainSelector:    2,
			expectedParticipants: []string{"node4", "node5"},
		},
		{
			name:                 "non-existent chain",
			destChainSelector:    999,
			expectedParticipants: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := le.GetParticipants(tt.destChainSelector)

			if len(result) != len(tt.expectedParticipants) {
				t.Errorf("expected %d participants, got %d", len(tt.expectedParticipants), len(result))
				return
			}

			for i, expected := range tt.expectedParticipants {
				if result[i] != expected {
					t.Errorf("expected participant %s at index %d, got %s", expected, i, result[i])
				}
			}
		})
	}
}

func TestSetParticipants(t *testing.T) {
	le := NewModuloLeaderElection("node1", make(map[uint64][]string), 1*time.Second)

	// Test setting participants for a new chain
	newParticipants := []string{"node1", "node2", "node3"}
	le.SetParticipants(1, newParticipants)

	// Verify participants were set
	result := le.GetParticipants(1)
	if len(result) != len(newParticipants) {
		t.Errorf("expected %d participants, got %d", len(newParticipants), len(result))
	}

	// Verify self participations were updated
	if !le.IsParticipant(1) {
		t.Error("expected to be participant after setting participants")
	}

	// Test updating existing participants
	updatedParticipants := []string{"node1", "node4", "node5"}
	le.SetParticipants(1, updatedParticipants)

	result = le.GetParticipants(1)
	if len(result) != len(updatedParticipants) {
		t.Errorf("expected %d participants after update, got %d", len(updatedParticipants), len(result))
	}

	// Verify self participations were updated correctly
	if !le.IsParticipant(1) {
		t.Error("expected to still be participant after updating participants")
	}

	// Test setting participants that don't include self
	le.SetParticipants(2, []string{"node2", "node3"})
	if le.IsParticipant(2) {
		t.Error("expected not to be participant in chain where self is not included")
	}
}

func TestSetParticipantsEmpty(t *testing.T) {
	le := NewModuloLeaderElection("node1", make(map[uint64][]string), 1*time.Second)

	// Test setting empty participants
	le.SetParticipants(1, []string{})

	result := le.GetParticipants(1)
	if len(result) != 0 {
		t.Errorf("expected 0 participants, got %d", len(result))
	}

	if le.IsParticipant(1) {
		t.Error("expected not to be participant with empty participants list")
	}
}

func TestConcurrentAccess(t *testing.T) {
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3"},
	}

	le := NewModuloLeaderElection("node1", participants, 1*time.Second)
	msgId := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	// Test concurrent reads
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			le.IsLeader(msgId, 1)
			le.IsParticipant(1)
			le.GetParticipants(1)
			le.SelfParticipations()
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}

	// Test concurrent writes
	done = make(chan bool, 5)
	for i := 0; i < 5; i++ {
		go func() {
			le.SetParticipants(uint64(i+100), []string{"node1", "node2"})
			done <- true
		}()
	}

	// Wait for all goroutines to complete
	for i := 0; i < 5; i++ {
		<-done
	}
}

func TestLeaderElectionDistribution(t *testing.T) {
	// Test that leader election provides reasonable distribution
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3", "node4", "node5"},
	}

	le := NewModuloLeaderElection("node1", participants, 1*time.Second)

	// Test multiple messages to see distribution
	leaderCount := 0
	deltaCounts := make(map[time.Duration]int)
	totalTests := 100_000

	for i := 0; i < totalTests; i++ {
		// Use different message IDs to simulate different messages
		msgId := [32]byte{}
		rand.Read(msgId[:])

		result, delay := le.IsLeader(msgId, 1)
		if result {
			leaderCount++
		}

		deltaCounts[delay]++
	}

	// With 5 participants, node1 should be leader roughly 20% of the time
	// Allow for some variance (19-21% range)
	expectedMin := totalTests * 19 / 100
	expectedMax := totalTests * 21 / 100

	if leaderCount < expectedMin || leaderCount > expectedMax {
		t.Errorf("leader distribution seems off: got %d/%d (%.1f%%), expected roughly 20%%",
			leaderCount, totalTests, float64(leaderCount)*100/float64(totalTests))
	}

	// Check that the delay distribution is roughly even
	for delay, count := range deltaCounts {
		if count < totalTests*19/100 || count > totalTests*21/100 {
			t.Errorf("delay %v count seems off: got %d/%d (%.1f%%), expected roughly 20%%",
				delay, count, totalTests, float64(count)*100/float64(totalTests))
		}
	}
}

func TestLeaderElectionDelay(t *testing.T) {
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3"},
	}

	le := NewModuloLeaderElection("node1", participants, 2*time.Second)

	// Create a random 32 byte array
	msgId := [32]byte{}
	rand.Read(msgId[:])

	// Test that the delay is calculated correctly based on position
	result, delay := le.IsLeader(msgId, 1)

	// If node1 is not the leader (index > 0), it should have a delay
	if !result {
		if delay <= 0 {
			t.Errorf("expected positive delay when not leader, got %v", delay)
		}
		// The delay should be a multiple of deltaStage
		if delay%2*time.Second != 0 {
			t.Errorf("expected delay to be multiple of deltaStage, got %v", delay)
		}
	} else {
		// If node1 is the leader, delay should be 0
		if delay != 0 {
			t.Errorf("expected 0 delay when leader, got %v", delay)
		}
	}
}

func TestLeaderElectionEdgeCases(t *testing.T) {
	// Test with single participant
	participants := map[uint64][]string{
		1: {"node1"},
	}

	le := NewModuloLeaderElection("node1", participants, 1*time.Second)
	msgId := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	// With single participant, should always be leader
	result, delay := le.IsLeader(msgId, 1)
	if !result {
		t.Error("single participant should always be leader")
	}
	if delay != 0 {
		t.Errorf("single participant should have 0 delay, got %v", delay)
	}
}

func TestDeltaStageConfiguration(t *testing.T) {
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3"},
	}

	// Test with different deltaStage values
	testCases := []time.Duration{
		100 * time.Millisecond,
		1 * time.Second,
		5 * time.Second,
	}

	for _, deltaStage := range testCases {
		t.Run(fmt.Sprintf("deltaStage_%v", deltaStage), func(t *testing.T) {
			le := NewModuloLeaderElection("node1", participants, deltaStage)

			msgId := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

			// Test that delays are multiples of deltaStage
			result, delay := le.IsLeader(msgId, 1)
			if !result && delay > 0 {
				if delay%deltaStage != 0 {
					t.Errorf("expected delay %v to be multiple of deltaStage %v", delay, deltaStage)
				}
			}
		})
	}
}
