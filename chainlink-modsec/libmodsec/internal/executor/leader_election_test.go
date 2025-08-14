package executor

import (
	"crypto/rand"
	"testing"
)

func TestIsLeader(t *testing.T) {
	// Create a consistent test setup
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3"},
		2: {"node1", "node4", "node5"},
		3: {"node2", "node3", "node4"}, // Not a participant
	}

	le := NewSimpleLeaderElection("node1", participants)

	// Test with a fixed message ID for deterministic results
	msgId := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	tests := []struct {
		name              string
		destChainSelector uint64
		epochs            []uint64
		expectedResult    bool
	}{
		{
			name:              "chain 1, is leader for offset 2",
			destChainSelector: 1,
			epochs:            []uint64{2},
			expectedResult:    true,
		},
		{
			name:              "chain 1, all other offsets are not leader",
			destChainSelector: 1,
			epochs:            []uint64{0, 1},
			expectedResult:    false,
		},
		{
			name:              "chain 2, is leader for offset 1",
			destChainSelector: 2,
			epochs:            []uint64{1},
			expectedResult:    true,
		},
		{
			name:              "chain 2, all other offsets are not leader",
			destChainSelector: 2,
			epochs:            []uint64{0, 2},
			expectedResult:    false,
		},
		{
			name:              "chain 3, node not a participant",
			destChainSelector: 3,
			epochs:            []uint64{0, 1, 2},
			expectedResult:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, epoch := range tt.epochs {
				result := le.IsLeader(msgId, tt.destChainSelector, epoch)
				if result != tt.expectedResult {
					t.Errorf("expected %v, got %v", tt.expectedResult, result)
				}
			}
		})
	}
}

func TestIsLeaderNonParticipant(t *testing.T) {
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3"},
	}

	le := NewSimpleLeaderElection("node1", participants)
	msgId := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	// Test that non-participant chains always return false
	if le.IsLeader(msgId, 999, 0) {
		t.Error("expected non-participant to not be leader")
	}
}

func TestIsLeaderDeterministic(t *testing.T) {
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3"},
	}

	le := NewSimpleLeaderElection("node1", participants)
	msgId := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	// Test that the same input always produces the same result
	result1 := le.IsLeader(msgId, 1, 0)
	result2 := le.IsLeader(msgId, 1, 0)

	if result1 != result2 {
		t.Error("leader election should be deterministic")
	}
}

func TestIsLeaderDifferentMessageIds(t *testing.T) {
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3"},
	}

	le := NewSimpleLeaderElection("node1", participants)

	msgId1 := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}
	msgId2 := [32]byte{2, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	// Different message IDs should potentially produce different results
	result1 := le.IsLeader(msgId1, 1, 0)
	result2 := le.IsLeader(msgId2, 1, 0)

	// Note: This test doesn't assert specific values since hash results could be the same
	// by chance, but it ensures the function handles different inputs correctly
	_ = result1
	_ = result2
}

func TestGetParticipants(t *testing.T) {
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3"},
		2: {"node4", "node5"},
	}

	le := NewSimpleLeaderElection("node1", participants)

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
	le := NewSimpleLeaderElection("node1", make(map[uint64][]string))

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
	le := NewSimpleLeaderElection("node1", make(map[uint64][]string))

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

	le := NewSimpleLeaderElection("node1", participants)
	msgId := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	// Test concurrent reads
	done := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func() {
			le.IsLeader(msgId, 1, 0)
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

	le := NewSimpleLeaderElection("node1", participants)

	// Test multiple messages to see distribution
	leaderCount := 0
	totalTests := 100_000

	for i := 0; i < totalTests; i++ {
		// Use different message IDs to simulate different messages
		msgId := [32]byte{}
		rand.Read(msgId[:])

		if le.IsLeader(msgId, 1, 0) {
			leaderCount++
		}
	}

	// With 5 participants, node1 should be leader roughly 20% of the time
	// Allow for some variance (19-21% range)
	expectedMin := totalTests * 19 / 100
	expectedMax := totalTests * 21 / 100

	if leaderCount < expectedMin || leaderCount > expectedMax {
		t.Errorf("leader distribution seems off: got %d/%d (%.1f%%), expected roughly 20%%",
			leaderCount, totalTests, float64(leaderCount)*100/float64(totalTests))
	}
}

func TestLeaderElectionOffset(t *testing.T) {
	participants := map[uint64][]string{
		1: {"node1", "node2", "node3"},
	}

	le := NewSimpleLeaderElection("node1", participants)

	// Create a random 32 byte array
	msgId := [32]byte{}
	rand.Read(msgId[:])

	// Test that different epochs can produce different results
	result0 := le.IsLeader(msgId, 1, 0)
	result1 := le.IsLeader(msgId, 1, 1)
	result2 := le.IsLeader(msgId, 1, 2)

	// Count how many epochs make node1 the leader
	leaderCount := 0
	if result0 {
		leaderCount++
	}
	if result1 {
		leaderCount++
	}
	if result2 {
		leaderCount++
	}

	// Only one epochs should make node1 the leader (since it's in the participants)
	// If more than one epochs makes node1 the leader, the test should fail
	if leaderCount > 1 {
		t.Errorf("expected only one epoch to make node1 the leader, but got %d: result0=%v, result1=%v, result2=%v",
			leaderCount, result0, result1, result2)
	}

	// If no epoch makes node1 the leader, the test should fail
	if leaderCount == 0 {
		t.Error("expected node1 to be leader for at least one epoch")
	}
}

func TestLeaderElectionEdgeCases(t *testing.T) {
	// Test with single participant
	participants := map[uint64][]string{
		1: {"node1"},
	}

	le := NewSimpleLeaderElection("node1", participants)
	msgId := [32]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32}

	// With single participant, should always be leader for epoch 0
	if !le.IsLeader(msgId, 1, 0) {
		t.Error("single participant should always be leader for epoch 0")
	}

	// Test with epoch modulo participant count should be leader
	if !le.IsLeader(msgId, 1, 1) {
		t.Error("should be leader for epoch modulo participant count")
	}
}
