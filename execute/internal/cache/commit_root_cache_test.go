package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

func TestCommitRootsCache_GetTimestampToQueryFrom(t *testing.T) {
	lggr := mocks.NullLogger
	messageVisibilityInterval := 8 * time.Hour
	rootSnoozeTime := 5 * time.Minute
	now := time.Now()

	// Create helper function for readability
	timeBefore := func(hours float64) time.Time {
		return now.Add(time.Duration(-hours * float64(time.Hour)))
	}

	testCases := []struct {
		name                 string
		latestFinalized      time.Time
		expectedResult       time.Time
		expectedOptimization bool
	}{
		{
			name:                 "no timestamp stored should use message visibility window",
			latestFinalized:      time.Time{}, // zero value
			expectedResult:       timeBefore(8),
			expectedOptimization: false,
		},
		{
			name:                 "older than visibility window should use visibility window",
			latestFinalized:      timeBefore(10),
			expectedResult:       timeBefore(8),
			expectedOptimization: true,
		},
		{
			name:                 "newer than visibility window should use finalized timestamp",
			latestFinalized:      timeBefore(2),
			expectedResult:       timeBefore(2),
			expectedOptimization: true,
		},
		{
			name:                 "edge case: equal to visibility window should use finalized timestamp",
			latestFinalized:      timeBefore(8), // Exactly equals the visibility window
			expectedResult:       timeBefore(8),
			expectedOptimization: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cache := internalNewCommitRootsCache(
				lggr,
				messageVisibilityInterval,
				rootSnoozeTime,
				CleanupInterval,
				EvictionGracePeriod,
			)

			// Set latest finalized timestamp if not zero
			if !tc.latestFinalized.IsZero() {
				cache.UpdateLatestFinalizedTimestamp(tc.latestFinalized)
			}

			// Get result
			result := cache.GetTimestampToQueryFrom()

			// Verify result is as expected
			// Using approximate comparison due to potential nanosecond differences
			timeDiff := result.Sub(tc.expectedResult)
			assert.True(t, timeDiff > -time.Second && timeDiff < time.Second,
				"Expected time close to %v but got %v (diff: %v)",
				tc.expectedResult, result, timeDiff)

			// Verify internal state was updated correctly
			messageVisibilityWindow := now.Add(-messageVisibilityInterval)
			if tc.latestFinalized.Before(messageVisibilityWindow) {
				// If original timestamp was before visibility window, internal state should be updated
				timeDiff = cache.latestFinalizedFullyExecutedRoot.Sub(messageVisibilityWindow)
				assert.True(t, timeDiff > -time.Second && timeDiff < time.Second,
					"Internal state should be updated to visibility window")
			} else if !tc.latestFinalized.IsZero() {
				// Otherwise should remain unchanged if not zero (using approximate comparison)
				timeDiff = cache.latestFinalizedFullyExecutedRoot.Sub(tc.latestFinalized)
				assert.True(t, timeDiff > -time.Second && timeDiff < time.Second,
					"Internal timestamp should equal finalized timestamp")
			}
		})
	}
}

func TestCommitRootsCache_TimestampQueryWithMessageVisibilityVariations(t *testing.T) {
	lggr := mocks.NullLogger
	rootSnoozeTime := 5 * time.Minute
	now := time.Now()

	// Helper functions for readability
	timeBefore := func(hours float64) time.Time {
		return now.Add(time.Duration(-hours * float64(time.Hour)))
	}

	testCases := []struct {
		name                      string
		messageVisibilityInterval time.Duration
		latestFinalized           time.Time
		expectedResult            time.Time
		description               string
	}{
		{
			name:                      "short visibility window with recent finalized timestamp",
			messageVisibilityInterval: 2 * time.Hour,
			latestFinalized:           timeBefore(1),
			expectedResult:            timeBefore(1),
		},
		{
			name:                      "standard visibility window with old finalized timestamp",
			messageVisibilityInterval: 8 * time.Hour,
			latestFinalized:           timeBefore(10),
			expectedResult:            timeBefore(8),
		},
		{
			name:                      "long visibility window with recent finalized timestamp",
			messageVisibilityInterval: 24 * time.Hour,
			latestFinalized:           timeBefore(12),
			expectedResult:            timeBefore(12),
		},
		{
			name:                      "zero visibility window (edge case)",
			messageVisibilityInterval: 0,
			latestFinalized:           timeBefore(5),
			expectedResult:            now,
		},
		{
			name:                      "very short visibility window",
			messageVisibilityInterval: 5 * time.Minute,
			latestFinalized:           timeBefore(1),
			expectedResult:            now.Add(-5 * time.Minute),
		},
		{
			name:                      "edge case: finalized at exactly visibility boundary",
			messageVisibilityInterval: 6 * time.Hour,
			latestFinalized:           timeBefore(6),
			expectedResult:            timeBefore(6),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create cache with the specific visibility interval for this test
			cache := internalNewCommitRootsCache(
				lggr,
				tc.messageVisibilityInterval,
				rootSnoozeTime,
				CleanupInterval,
				EvictionGracePeriod,
			)

			// Set latest finalized timestamp
			if !tc.latestFinalized.IsZero() {
				cache.UpdateLatestFinalizedTimestamp(tc.latestFinalized)
			}

			// Get result
			result := cache.GetTimestampToQueryFrom()

			// Verify result is as expected
			// Using approximate comparison due to potential nanosecond differences
			timeDiff := result.Sub(tc.expectedResult)
			assert.True(t, timeDiff > -time.Second && timeDiff < time.Second,
				"Expected time close to %v but got %v (diff: %v) - %s",
				tc.expectedResult, result, timeDiff, tc.description)

			// Additional verification: check internal state
			messageVisibilityWindow := now.Add(-tc.messageVisibilityInterval)
			if tc.latestFinalized.Before(messageVisibilityWindow) {
				// If finalized is before visibility window, internal state should be updated
				assert.True(t, !cache.latestFinalizedFullyExecutedRoot.Before(messageVisibilityWindow),
					"Internal timestamp should have been updated to at least visibility window")
			} else {
				// Otherwise it should match the original finalized timestamp (approximate comparison)
				timeDiff = cache.latestFinalizedFullyExecutedRoot.Sub(tc.latestFinalized)
				assert.True(t, timeDiff > -time.Second && timeDiff < time.Second,
					"Internal timestamp should remain approximately equal to finalized timestamp")
			}
		})
	}

	// Test the behavior with changing visibility windows
	t.Run("changing visibility windows", func(t *testing.T) {
		// Start with a standard window
		initialVisibility := 8 * time.Hour
		initialFinalized := timeBefore(4) // Within initial window

		cache := internalNewCommitRootsCache(
			lggr,
			initialVisibility,
			rootSnoozeTime,
			CleanupInterval,
			EvictionGracePeriod,
		)

		cache.UpdateLatestFinalizedTimestamp(initialFinalized)
		initialResult := cache.GetTimestampToQueryFrom()

		// Should use finalized timestamp
		timeDiff := initialResult.Sub(initialFinalized)
		assert.True(t, timeDiff > -time.Second && timeDiff < time.Second,
			"Initial result should use finalized timestamp")

		// Now change to a shorter window that would make finalized outside window
		newVisibility := 2 * time.Hour
		cache.messageVisibilityInterval = newVisibility

		// This should now use visibility window instead
		newResult := cache.GetTimestampToQueryFrom()
		expectedWindow := now.Add(-newVisibility)

		timeDiff = newResult.Sub(expectedWindow)
		assert.True(t, timeDiff > -time.Second && timeDiff < time.Second,
			"After shortening visibility window, should use visibility window")

		// Verify internal state was updated
		assert.True(t, !cache.latestFinalizedFullyExecutedRoot.Before(expectedWindow),
			"Internal timestamp should have been updated to new visibility window")
	})
}

func TestCommitRootsCache_UpdateLatestFinalizedTimestamp(t *testing.T) {
	lggr := mocks.NullLogger
	messageVisibilityInterval := 8 * time.Hour
	rootSnoozeTime := 5 * time.Minute
	now := time.Now()

	testCases := []struct {
		name               string
		initialTimestamp   time.Time
		updateTimestamp    time.Time
		expectUpdate       bool
		secondUpdate       time.Time
		expectSecondUpdate bool
	}{
		{
			name:               "initial update should succeed",
			initialTimestamp:   time.Time{}, // zero value
			updateTimestamp:    now.Add(-2 * time.Hour),
			expectUpdate:       true,
			secondUpdate:       time.Time{}, // no second update
			expectSecondUpdate: false,
		},
		{
			name:               "newer timestamp should update",
			initialTimestamp:   now.Add(-3 * time.Hour),
			updateTimestamp:    now.Add(-1 * time.Hour),
			expectUpdate:       true,
			secondUpdate:       time.Time{}, // no second update
			expectSecondUpdate: false,
		},
		{
			name:               "older timestamp should not update",
			initialTimestamp:   now.Add(-1 * time.Hour),
			updateTimestamp:    now.Add(-2 * time.Hour),
			expectUpdate:       false,
			secondUpdate:       time.Time{}, // no second update
			expectSecondUpdate: false,
		},
		{
			name:               "same timestamp should not update",
			initialTimestamp:   now.Add(-1 * time.Hour),
			updateTimestamp:    now.Add(-1 * time.Hour),
			expectUpdate:       false,
			secondUpdate:       time.Time{}, // no second update
			expectSecondUpdate: false,
		},
		{
			name:               "multiple updates in sequence",
			initialTimestamp:   now.Add(-3 * time.Hour),
			updateTimestamp:    now.Add(-2 * time.Hour),
			expectUpdate:       true,
			secondUpdate:       now.Add(-1 * time.Hour),
			expectSecondUpdate: true,
		},
		{
			name:               "multiple updates with older second timestamp",
			initialTimestamp:   now.Add(-3 * time.Hour),
			updateTimestamp:    now.Add(-1 * time.Hour),
			expectUpdate:       true,
			secondUpdate:       now.Add(-2 * time.Hour),
			expectSecondUpdate: false,
		},
		{
			name:               "zero timestamp should not update",
			initialTimestamp:   now.Add(-1 * time.Hour),
			updateTimestamp:    time.Time{}, // zero value
			expectUpdate:       false,
			secondUpdate:       time.Time{}, // no second update
			expectSecondUpdate: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cache := internalNewCommitRootsCache(
				lggr,
				messageVisibilityInterval,
				rootSnoozeTime,
				CleanupInterval,
				EvictionGracePeriod,
			)

			// Set initial timestamp if not zero
			if !tc.initialTimestamp.IsZero() {
				cache.UpdateLatestFinalizedTimestamp(tc.initialTimestamp)
				assert.Equal(t, tc.initialTimestamp, cache.latestFinalizedFullyExecutedRoot,
					"Initial timestamp should be set correctly")
			}

			// Perform update
			cache.UpdateLatestFinalizedTimestamp(tc.updateTimestamp)

			// Verify update occurred or not
			if tc.expectUpdate {
				assert.Equal(t, tc.updateTimestamp, cache.latestFinalizedFullyExecutedRoot,
					"Timestamp should be updated to new value")
			} else {
				if tc.initialTimestamp.IsZero() {
					assert.True(t, cache.latestFinalizedFullyExecutedRoot.IsZero(),
						"Timestamp should still be zero")
				} else {
					assert.Equal(t, tc.initialTimestamp, cache.latestFinalizedFullyExecutedRoot,
						"Timestamp should not be updated")
				}
			}

			// Perform second update if provided
			if !tc.secondUpdate.IsZero() {
				expectedBeforeSecondUpdate := tc.initialTimestamp
				if tc.expectUpdate {
					expectedBeforeSecondUpdate = tc.updateTimestamp
				}

				cache.UpdateLatestFinalizedTimestamp(tc.secondUpdate)

				// Verify second update occurred or not
				if tc.expectSecondUpdate {
					assert.Equal(t, tc.secondUpdate, cache.latestFinalizedFullyExecutedRoot,
						"Timestamp should be updated to second value")
				} else {
					assert.Equal(t, expectedBeforeSecondUpdate, cache.latestFinalizedFullyExecutedRoot,
						"Timestamp should not be updated by second update")
				}
			}
		})
	}
}

func TestCommitRootsCache_MarkExecuteAndSnooze(t *testing.T) {
	lggr := mocks.NullLogger
	messageVisibilityInterval := 8 * time.Hour
	rootSnoozeTime := 5 * time.Minute

	// Create unique chain selectors and merkle roots for testing
	sourceChain1 := ccipocr3.ChainSelector(1)
	sourceChain2 := ccipocr3.ChainSelector(2)

	merkleRoot1 := ccipocr3.Bytes32{0x01}
	merkleRoot2 := ccipocr3.Bytes32{0x02}
	merkleRoot3 := ccipocr3.Bytes32{0x03}

	testCases := []struct {
		name         string
		operations   func(cache *commitRootsCache)
		checkResults func(t *testing.T, cache *commitRootsCache)
	}{
		{
			name: "mark as executed",
			operations: func(cache *commitRootsCache) {
				cache.MarkAsExecuted(sourceChain1, merkleRoot1)
			},
			checkResults: func(t *testing.T, cache *commitRootsCache) {
				// Should be executed
				assert.False(t, cache.CanExecute(sourceChain1, merkleRoot1), "Should be marked as executed")
				// Different roots should be executable
				assert.True(t, cache.CanExecute(sourceChain1, merkleRoot2), "Different root should be executable")
				assert.True(t, cache.CanExecute(sourceChain2, merkleRoot1), "Different chain should be executable")
			},
		},
		{
			name: "snooze root",
			operations: func(cache *commitRootsCache) {
				cache.Snooze(sourceChain1, merkleRoot1)
			},
			checkResults: func(t *testing.T, cache *commitRootsCache) {
				// Should be snoozed
				assert.False(t, cache.CanExecute(sourceChain1, merkleRoot1), "Should be snoozed")
				// Different roots should be executable
				assert.True(t, cache.CanExecute(sourceChain1, merkleRoot2), "Different root should be executable")
				assert.True(t, cache.CanExecute(sourceChain2, merkleRoot1), "Different chain should be executable")
			},
		},
		{
			name: "mark multiple roots",
			operations: func(cache *commitRootsCache) {
				cache.MarkAsExecuted(sourceChain1, merkleRoot1)
				cache.MarkAsExecuted(sourceChain1, merkleRoot2)
				cache.MarkAsExecuted(sourceChain2, merkleRoot1)
			},
			checkResults: func(t *testing.T, cache *commitRootsCache) {
				// All marked roots should be not executable
				assert.False(t, cache.CanExecute(sourceChain1, merkleRoot1), "Root 1 chain 1 should be executed")
				assert.False(t, cache.CanExecute(sourceChain1, merkleRoot2), "Root 2 chain 1 should be executed")
				assert.False(t, cache.CanExecute(sourceChain2, merkleRoot1), "Root 1 chain 2 should be executed")
				// Unmarked root should be executable
				assert.True(t, cache.CanExecute(sourceChain2, merkleRoot2), "Unmarked root should be executable")
			},
		},
		{
			name: "snooze after execute",
			operations: func(cache *commitRootsCache) {
				cache.MarkAsExecuted(sourceChain1, merkleRoot1)
				cache.Snooze(sourceChain1, merkleRoot1)
			},
			checkResults: func(t *testing.T, cache *commitRootsCache) {
				// Should be not executable (both executed and snoozed)
				assert.False(t, cache.CanExecute(sourceChain1, merkleRoot1), "Should be not executable")
				// Check internal state - should be both executed and snoozed
				key := getKey(sourceChain1, merkleRoot1)
				assert.True(t, cache.isExecuted(key), "Should be executed")
				assert.True(t, cache.isSnoozed(key), "Should be snoozed")
			},
		},
		{
			name: "execute after snooze",
			operations: func(cache *commitRootsCache) {
				cache.Snooze(sourceChain1, merkleRoot1)
				cache.MarkAsExecuted(sourceChain1, merkleRoot1)
			},
			checkResults: func(t *testing.T, cache *commitRootsCache) {
				// Should be not executable (both snoozed and executed)
				assert.False(t, cache.CanExecute(sourceChain1, merkleRoot1), "Should be not executable")
				// Check internal state - should be both snoozed and executed
				key := getKey(sourceChain1, merkleRoot1)
				assert.True(t, cache.isExecuted(key), "Should be executed")
				assert.True(t, cache.isSnoozed(key), "Should be snoozed")
			},
		},
		{
			name: "multiple operations different roots",
			operations: func(cache *commitRootsCache) {
				cache.MarkAsExecuted(sourceChain1, merkleRoot1)
				cache.Snooze(sourceChain1, merkleRoot2)
				cache.MarkAsExecuted(sourceChain2, merkleRoot3)
			},
			checkResults: func(t *testing.T, cache *commitRootsCache) {
				// Each root should have correct state
				assert.False(t, cache.CanExecute(sourceChain1, merkleRoot1), "Root 1 should be executed")
				assert.False(t, cache.CanExecute(sourceChain1, merkleRoot2), "Root 2 should be snoozed")
				assert.False(t, cache.CanExecute(sourceChain2, merkleRoot3), "Root 3 should be executed")
				// Unmarked root should be executable
				assert.True(t, cache.CanExecute(sourceChain2, merkleRoot2), "Unmarked root should be executable")
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cache := internalNewCommitRootsCache(
				lggr,
				messageVisibilityInterval,
				rootSnoozeTime,
				CleanupInterval,
				EvictionGracePeriod,
			)

			// Apply operations
			tc.operations(cache)

			// Check results
			tc.checkResults(t, cache)
		})
	}
}

func TestCommitRootsCache_TTLEviction(t *testing.T) {
	// This test uses a very short TTL to test that roots are actually evicted
	lggr := mocks.NullLogger
	messageVisibilityInterval := 50 * time.Millisecond
	rootSnoozeTime := 25 * time.Millisecond
	cleanupInterval := 10 * time.Millisecond // Faster cleanup for testing
	evictionGracePeriod := 10 * time.Millisecond

	sourceChain := ccipocr3.ChainSelector(1)
	merkleRoot := ccipocr3.Bytes32{0x01}

	cache := internalNewCommitRootsCache(
		lggr,
		messageVisibilityInterval,
		rootSnoozeTime,
		cleanupInterval,
		evictionGracePeriod,
	)

	t.Run("snoozed roots should expire", func(t *testing.T) {
		// Snooze a root
		cache.Snooze(sourceChain, merkleRoot)

		// Should be snoozed initially
		assert.False(t, cache.CanExecute(sourceChain, merkleRoot), "Should be snoozed immediately after")

		// Wait for snooze to expire
		time.Sleep(rootSnoozeTime + cleanupInterval*2)

		// Should be executable now
		assert.True(t, cache.CanExecute(sourceChain, merkleRoot), "Should be executable after snooze expiry")
	})

	t.Run("executed roots should expire after visibility interval + grace period", func(t *testing.T) {
		// Mark as executed
		cache.MarkAsExecuted(sourceChain, merkleRoot)

		// Should be not executable initially
		assert.False(t, cache.CanExecute(sourceChain, merkleRoot), "Should be marked as executed immediately after")

		// Wait for execution record to expire (visibility interval + grace period + cleanup buffer)
		time.Sleep(messageVisibilityInterval + evictionGracePeriod + cleanupInterval*2)

		// Should be executable now
		assert.True(t, cache.CanExecute(sourceChain, merkleRoot), "Should be executable after expiry")
	})
}

func TestCommitRootsCache_Concurrency(t *testing.T) {
	lggr := mocks.NullLogger
	messageVisibilityInterval := 8 * time.Hour
	rootSnoozeTime := 5 * time.Minute

	sourceChain := ccipocr3.ChainSelector(1)
	merkleRoot := ccipocr3.Bytes32{0x01}

	cache := internalNewCommitRootsCache(
		lggr,
		messageVisibilityInterval,
		rootSnoozeTime,
		CleanupInterval,
		EvictionGracePeriod,
	)

	// Number of concurrent operations to perform
	const numConcurrentOps = 100
	done := make(chan struct{}, numConcurrentOps*3) // 3 operations per goroutine

	// Run concurrent operations
	for i := 0; i < numConcurrentOps; i++ {
		go func(id int) {
			// Create unique roots for this goroutine
			customRoot := ccipocr3.Bytes32{byte(id % 256)}

			// Perform operations that use locks
			cache.MarkAsExecuted(sourceChain, customRoot)
			done <- struct{}{}

			cache.Snooze(sourceChain, customRoot)
			done <- struct{}{}

			// Also check the shared root
			cache.CanExecute(sourceChain, merkleRoot)
			done <- struct{}{}
		}(i)
	}

	// Wait for all operations to complete
	for i := 0; i < numConcurrentOps*3; i++ {
		select {
		case <-done:
			// Operation completed successfully
		case <-time.After(2 * time.Second):
			t.Fatal("Timeout waiting for concurrent operations to complete")
		}
	}

	// If we've reached here without deadlocks, panics, or races, the test passes
	// (Note: use -race flag with go test to catch races)
}

func TestNewCommitRootsCache(t *testing.T) {
	lggr := mocks.NullLogger
	messageVisibilityInterval := 8 * time.Hour
	rootSnoozeTime := 5 * time.Minute

	// Test the public constructor
	cache := NewCommitRootsCache(
		lggr,
		messageVisibilityInterval,
		rootSnoozeTime,
	)

	// Verify it's properly constructed
	require.NotNil(t, cache, "Cache should not be nil")

	// Type assertion to access internal fields
	concreteCache, ok := cache.(*commitRootsCache)
	require.True(t, ok, "Cache should be of type *commitRootsCache")

	// Verify internal state
	assert.Equal(t, messageVisibilityInterval, concreteCache.messageVisibilityInterval)
	assert.Equal(t, rootSnoozeTime, concreteCache.rootSnoozeTime)
	assert.NotNil(t, concreteCache.executedRoots, "executedRoots should not be nil")
	assert.NotNil(t, concreteCache.snoozedRoots, "snoozedRoots should not be nil")
	assert.True(t, concreteCache.latestFinalizedFullyExecutedRoot.IsZero(), "latestFinalizedTimestamp should be zero value")
}
