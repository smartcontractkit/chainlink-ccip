package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

// Helper function to create a test CommitData
func createCommitData(
	timestamp time.Time,
	sourceChain ccipocr3.ChainSelector,
	merkleRoot ccipocr3.Bytes32) exectypes.CommitData {
	return exectypes.CommitData{
		Timestamp:   timestamp,
		SourceChain: sourceChain,
		MerkleRoot:  merkleRoot,
	}
}

// Helper function to create a map of commit reports
func createCommitReports(reports ...exectypes.CommitData) map[ccipocr3.ChainSelector][]exectypes.CommitData {
	result := make(map[ccipocr3.ChainSelector][]exectypes.CommitData)
	for _, report := range reports {
		result[report.SourceChain] = append(result[report.SourceChain], report)
	}
	return result
}

func TestCommitRootsCache_GetTimestampToQueryFrom(t *testing.T) {
	lggr := logger.Nop()
	messageVisibilityInterval := 8 * time.Hour
	rootSnoozeTime := 5 * time.Minute

	now := time.Now()
	messageVisibilityWindow := now.Add(-messageVisibilityInterval)

	tests := []struct {
		name                        string
		earliestUnexecutedRoot      time.Time
		messageVisibilityWindow     time.Time
		expectedQueryTimestamp      time.Time
		expectedOptimizationApplied bool
	}{
		{
			name:                        "No unexecuted root, use visibility window",
			earliestUnexecutedRoot:      time.Time{}, // Zero value
			messageVisibilityWindow:     messageVisibilityWindow,
			expectedQueryTimestamp:      messageVisibilityWindow,
			expectedOptimizationApplied: false,
		},
		{
			name:                        "Unexecuted root before visibility window, use visibility window",
			earliestUnexecutedRoot:      messageVisibilityWindow.Add(-1 * time.Hour),
			messageVisibilityWindow:     messageVisibilityWindow,
			expectedQueryTimestamp:      messageVisibilityWindow,
			expectedOptimizationApplied: false,
		},
		{
			name:                        "Unexecuted root after visibility window, optimize query",
			earliestUnexecutedRoot:      messageVisibilityWindow.Add(1 * time.Hour),
			messageVisibilityWindow:     messageVisibilityWindow,
			expectedQueryTimestamp:      messageVisibilityWindow.Add(1 * time.Hour),
			expectedOptimizationApplied: true,
		},
		{
			name:                        "Unexecuted root at visibility window, use visibility window",
			earliestUnexecutedRoot:      messageVisibilityWindow,
			messageVisibilityWindow:     messageVisibilityWindow,
			expectedQueryTimestamp:      messageVisibilityWindow,
			expectedOptimizationApplied: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := internalNewCommitRootsCache(
				lggr,
				messageVisibilityInterval,
				rootSnoozeTime,
				CleanupInterval,
				EvictionGracePeriod,
			)

			// Set the earliestUnexecutedRoot directly for testing
			cache.earliestUnexecutedRoot = tt.earliestUnexecutedRoot

			// Get the query timestamp
			queryTimestamp := cache.GetTimestampToQueryFrom(tt.messageVisibilityWindow)

			// Verify the result
			if tt.expectedOptimizationApplied {
				assert.True(t, queryTimestamp.After(tt.messageVisibilityWindow),
					"Expected query timestamp to be after visibility window")
			} else {
				assert.Equal(t, tt.expectedQueryTimestamp, queryTimestamp,
					"Query timestamp should match expected value")
			}
		})
	}
}

func TestCommitRootsCache_UpdateEarliestUnexecutedRoot(t *testing.T) {
	lggr := logger.Nop()
	messageVisibilityInterval := 8 * time.Hour
	rootSnoozeTime := 5 * time.Minute

	now := time.Now()

	// Create test data
	selector1 := ccipocr3.ChainSelector(1)
	selector2 := ccipocr3.ChainSelector(2)

	root1 := ccipocr3.Bytes32{1}
	root2 := ccipocr3.Bytes32{2}
	root3 := ccipocr3.Bytes32{3}
	root4 := ccipocr3.Bytes32{4}

	timestamp1 := now.Add(-1 * time.Hour)    // 1 hour ago
	timestamp2 := now.Add(-40 * time.Minute) // 40 minutes ago
	timestamp3 := now.Add(-20 * time.Minute) // 20 minutes ago
	timestamp4 := now.Add(-10 * time.Minute) // 10 minutes ago

	report1 := createCommitData(timestamp1, selector1, root1)
	report2 := createCommitData(timestamp2, selector1, root2)
	report3 := createCommitData(timestamp3, selector2, root3)
	report4 := createCommitData(timestamp4, selector2, root4)

	tests := []struct {
		name                 string
		initialValue         time.Time
		remainingReports     map[ccipocr3.ChainSelector][]exectypes.CommitData
		expectedUpdatedValue time.Time
		expectChange         bool
	}{
		{
			name:                 "No unexecuted reports, zero initial value",
			initialValue:         time.Time{},
			remainingReports:     createCommitReports(),
			expectedUpdatedValue: time.Time{},
			expectChange:         false,
		},
		{
			name:                 "No unexecuted reports, reset non-zero initial value",
			initialValue:         timestamp1,
			remainingReports:     createCommitReports(),
			expectedUpdatedValue: time.Time{},
			expectChange:         true,
		},
		{
			name:                 "Single unexecuted report, update from zero",
			initialValue:         time.Time{},
			remainingReports:     createCommitReports(report1),
			expectedUpdatedValue: timestamp1,
			expectChange:         true,
		},
		{
			name:                 "Multiple reports, same chain, find earliest",
			initialValue:         time.Time{},
			remainingReports:     createCommitReports(report1, report2),
			expectedUpdatedValue: timestamp1, // report1 is earlier
			expectChange:         true,
		},
		{
			name:                 "Multiple reports, different chains, find earliest",
			initialValue:         time.Time{},
			remainingReports:     createCommitReports(report2, report3, report4),
			expectedUpdatedValue: timestamp2, // report2 is the earliest
			expectChange:         true,
		},
		{
			name:                 "New earliest is earlier than current",
			initialValue:         timestamp2,
			remainingReports:     createCommitReports(report1, report3), // report1 is earlier than initialValue
			expectedUpdatedValue: timestamp1,
			expectChange:         true,
		},
		{
			name:                 "New earliest is later than current",
			initialValue:         timestamp1,
			remainingReports:     createCommitReports(report2, report3), // report2 is later than initialValue
			expectedUpdatedValue: timestamp2,
			expectChange:         true,
		},
		{
			name:                 "New earliest same as current",
			initialValue:         timestamp1,
			remainingReports:     createCommitReports(report1, report3), // report1 matches initialValue
			expectedUpdatedValue: timestamp1,
			expectChange:         false, // No change expected
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cache := internalNewCommitRootsCache(
				lggr,
				messageVisibilityInterval,
				rootSnoozeTime,
				CleanupInterval,
				EvictionGracePeriod,
			)

			// Set initial value
			cache.earliestUnexecutedRoot = tt.initialValue

			// Call the method under test
			cache.UpdateEarliestUnexecutedRoot(tt.remainingReports)

			// Verify the result
			if tt.expectChange {
				assert.Equal(t, tt.expectedUpdatedValue, cache.earliestUnexecutedRoot,
					"Earliest unexecuted root should be updated to expected value")
			} else {
				assert.Equal(t, tt.initialValue, cache.earliestUnexecutedRoot,
					"Earliest unexecuted root should remain unchanged")
			}
		})
	}
}

func TestCommitRootsCache_ScenarioTests(t *testing.T) {
	lggr := logger.Nop()
	messageVisibilityInterval := 8 * time.Hour
	rootSnoozeTime := 5 * time.Minute

	now := time.Now()
	messageVisibilityWindow := now.Add(-messageVisibilityInterval)

	// Create test data for the colleague's scenario
	selector := ccipocr3.ChainSelector(1)

	root1 := ccipocr3.Bytes32{1}
	root2 := ccipocr3.Bytes32{2}
	root3 := ccipocr3.Bytes32{3}

	timestamp1 := now.Add(-30 * time.Minute) // 10:30am
	timestamp2 := now.Add(-20 * time.Minute) // 10:40am
	timestamp3 := now.Add(-10 * time.Minute) // 10:50am

	report1 := createCommitData(timestamp1, selector, root1)
	report2 := createCommitData(timestamp2, selector, root2)
	report3 := createCommitData(timestamp3, selector, root3)

	t.Run("Colleague's Scenario: Root1 and Root3 executed, Root2 skipped", func(t *testing.T) {
		cache := internalNewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
			CleanupInterval,
			EvictionGracePeriod,
		)

		// Initial state - all roots unexecuted
		allReports := createCommitReports(report1, report2, report3)
		cache.UpdateEarliestUnexecutedRoot(allReports)

		// Verify initial state
		assert.Equal(t, timestamp1, cache.earliestUnexecutedRoot, "Initial earliest should be Root1")

		// First round - execute Root1 and Root3, but not Root2
		cache.MarkAsExecuted(selector, root1)
		cache.MarkAsExecuted(selector, root3)

		// Update with remaining reports (just Root2)
		remainingReports := createCommitReports(report2)
		cache.UpdateEarliestUnexecutedRoot(remainingReports)

		// Verify state after first round
		assert.Equal(t, timestamp2, cache.earliestUnexecutedRoot, "After execution, earliest should be Root2")

		// Get query timestamp - should use Root2's timestamp since it's after visibility window
		queryTimestamp := cache.GetTimestampToQueryFrom(messageVisibilityWindow)
		assert.Equal(t, timestamp2, queryTimestamp, "Query timestamp should be Root2's timestamp")

		// Verify Root2 is still considered executable
		assert.True(t, cache.CanExecute(selector, root2), "Root2 should be executable")
		assert.False(t, cache.CanExecute(selector, root1), "Root1 should not be executable")
		assert.False(t, cache.CanExecute(selector, root3), "Root3 should not be executable")
	})

	t.Run("Edge Case: All roots executed in single round", func(t *testing.T) {
		cache := internalNewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
			CleanupInterval,
			EvictionGracePeriod,
		)

		// Initial state - all roots unexecuted
		allReports := createCommitReports(report1, report2, report3)
		cache.UpdateEarliestUnexecutedRoot(allReports)

		// Execute all roots
		cache.MarkAsExecuted(selector, root1)
		cache.MarkAsExecuted(selector, root2)
		cache.MarkAsExecuted(selector, root3)

		// Update with empty remaining reports
		cache.UpdateEarliestUnexecutedRoot(createCommitReports())

		// Verify state is reset
		assert.True(t, cache.earliestUnexecutedRoot.IsZero(), "Earliest should be reset when all roots executed")

		// Get query timestamp - should use visibility window
		queryTimestamp := cache.GetTimestampToQueryFrom(messageVisibilityWindow)
		assert.Equal(t, messageVisibilityWindow, queryTimestamp, "Query timestamp should be visibility window")
	})

	t.Run("Edge Case: Roots executed out of order", func(t *testing.T) {
		cache := internalNewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
			CleanupInterval,
			EvictionGracePeriod,
		)

		// Initial state - all roots unexecuted
		allReports := createCommitReports(report1, report2, report3)
		cache.UpdateEarliestUnexecutedRoot(allReports)

		// Execute Root3 (latest) first
		cache.MarkAsExecuted(selector, root3)

		// Update with remaining reports
		remainingReports := createCommitReports(report1, report2)
		cache.UpdateEarliestUnexecutedRoot(remainingReports)

		// Verify earliest is still Root1
		assert.Equal(t, timestamp1, cache.earliestUnexecutedRoot, "Earliest should still be Root1")

		// Execute Root1 (earliest) next
		cache.MarkAsExecuted(selector, root1)

		// Update with remaining reports
		remainingReports = createCommitReports(report2)
		cache.UpdateEarliestUnexecutedRoot(remainingReports)

		// Verify earliest is now Root2
		assert.Equal(t, timestamp2, cache.earliestUnexecutedRoot, "Earliest should now be Root2")
	})

	t.Run("Edge Case: Snoozed roots are not considered executed", func(t *testing.T) {
		cache := internalNewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
			CleanupInterval,
			EvictionGracePeriod,
		)

		// Initial state - all roots unexecuted
		allReports := createCommitReports(report1, report2, report3)
		cache.UpdateEarliestUnexecutedRoot(allReports)

		// Snooze Root2
		cache.Snooze(selector, root2)

		// Verify Root2 is not executable but still tracked
		assert.False(t, cache.CanExecute(selector, root2), "Snoozed root should not be executable")
		assert.Equal(t, timestamp1, cache.earliestUnexecutedRoot, "Earliest should still be Root1")

		// After snooze expires, Root2 should be executable again
		// (Would need time mocking to test this thoroughly)
	})

	t.Run("Edge Case: Multiple chains with different execution patterns", func(t *testing.T) {
		cache := internalNewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
			CleanupInterval,
			EvictionGracePeriod,
		)

		// Create roots on different chains
		selector1 := ccipocr3.ChainSelector(1)
		selector2 := ccipocr3.ChainSelector(2)

		root1Chain1 := ccipocr3.Bytes32{11}
		root2Chain1 := ccipocr3.Bytes32{12}
		root1Chain2 := ccipocr3.Bytes32{21}
		root2Chain2 := ccipocr3.Bytes32{22}

		timestamp1Chain1 := now.Add(-40 * time.Minute)
		timestamp2Chain1 := now.Add(-30 * time.Minute)
		timestamp1Chain2 := now.Add(-20 * time.Minute)
		timestamp2Chain2 := now.Add(-10 * time.Minute)

		report1Chain1 := createCommitData(timestamp1Chain1, selector1, root1Chain1)
		report2Chain1 := createCommitData(timestamp2Chain1, selector1, root2Chain1)
		report1Chain2 := createCommitData(timestamp1Chain2, selector2, root1Chain2)
		report2Chain2 := createCommitData(timestamp2Chain2, selector2, root2Chain2)

		// Initial state - all roots unexecuted
		allReports := createCommitReports(report1Chain1, report2Chain1, report1Chain2, report2Chain2)
		cache.UpdateEarliestUnexecutedRoot(allReports)

		// Verify initial state
		assert.Equal(t, timestamp1Chain1, cache.earliestUnexecutedRoot, "Initial earliest should be from Chain1")

		// Execute all Chain1 roots
		cache.MarkAsExecuted(selector1, root1Chain1)
		cache.MarkAsExecuted(selector1, root2Chain1)

		// Update with remaining reports
		remainingReports := createCommitReports(report1Chain2, report2Chain2)
		cache.UpdateEarliestUnexecutedRoot(remainingReports)

		// Verify earliest is now from Chain2
		assert.Equal(t, timestamp1Chain2, cache.earliestUnexecutedRoot, "Earliest should now be from Chain2")
	})
}

func TestCommitRootsCache_IntegrationScenario(t *testing.T) {
	lggr := logger.Nop()
	messageVisibilityInterval := 8 * time.Hour
	rootSnoozeTime := 5 * time.Minute

	now := time.Now()

	// Create a timeline with 5 roots
	selector := ccipocr3.ChainSelector(1)

	root1 := ccipocr3.Bytes32{1}
	root2 := ccipocr3.Bytes32{2}
	root3 := ccipocr3.Bytes32{3}
	root4 := ccipocr3.Bytes32{4}
	root5 := ccipocr3.Bytes32{5}

	// Create timestamps that align with our example scenario
	timestamp1 := now.Add(-40 * time.Minute) // 10:30am
	timestamp2 := now.Add(-30 * time.Minute) // 10:40am
	timestamp3 := now.Add(-20 * time.Minute) // 10:50am
	timestamp4 := now.Add(-10 * time.Minute) // 11:00am
	timestamp5 := now.Add(-5 * time.Minute)  // 11:10am

	report1 := createCommitData(timestamp1, selector, root1)
	report2 := createCommitData(timestamp2, selector, root2)
	report3 := createCommitData(timestamp3, selector, root3)
	report4 := createCommitData(timestamp4, selector, root4)
	report5 := createCommitData(timestamp5, selector, root5)

	t.Run("Full Timeline Scenario", func(t *testing.T) {
		cache := internalNewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
			CleanupInterval,
			EvictionGracePeriod,
		)

		// Message visibility window is 8 hours ago
		messageVisibilityWindow := now.Add(-messageVisibilityInterval)

		// Initial state - no unexecuted roots
		queryTimestamp1 := cache.GetTimestampToQueryFrom(messageVisibilityWindow)
		assert.Equal(t, messageVisibilityWindow, queryTimestamp1, "Initial query should use visibility window")

		// First query at 11:00am - Discover all 5 roots
		allReports := createCommitReports(report1, report2, report3, report4, report5)
		cache.UpdateEarliestUnexecutedRoot(allReports)

		// Verify initial tracking
		assert.Equal(t, timestamp1, cache.earliestUnexecutedRoot, "Initial earliest should be Root1")

		// Execute Root1, Root3, Root5
		cache.MarkAsExecuted(selector, root1)
		cache.MarkAsExecuted(selector, root3)
		cache.MarkAsExecuted(selector, root5)

		// Update with remaining unexecuted roots
		remainingReports := createCommitReports(report2, report4)
		cache.UpdateEarliestUnexecutedRoot(remainingReports)

		// Verify tracking after first execution round
		assert.Equal(t, timestamp2, cache.earliestUnexecutedRoot, "After first round, earliest should be Root2")

		// Second query at 11:15am
		queryTimestamp2 := cache.GetTimestampToQueryFrom(messageVisibilityWindow)
		assert.Equal(t, timestamp2, queryTimestamp2, "Second query should use Root2's timestamp")

		// Execute Root4, Root2 remains unexecuted
		cache.MarkAsExecuted(selector, root4)

		// Update with remaining unexecuted roots
		remainingReports = createCommitReports(report2)
		cache.UpdateEarliestUnexecutedRoot(remainingReports)

		// Verify tracking after second execution round
		assert.Equal(t, timestamp2, cache.earliestUnexecutedRoot, "After second round, earliest should still be Root2")

		// Third query at 11:30am
		queryTimestamp3 := cache.GetTimestampToQueryFrom(messageVisibilityWindow)
		assert.Equal(t, timestamp2, queryTimestamp3, "Third query should use Root2's timestamp")

		// Root2 still unexecuted

		// Fourth query at 11:45am - Root2 finally executes
		cache.MarkAsExecuted(selector, root2)

		// Update with empty remaining reports
		cache.UpdateEarliestUnexecutedRoot(createCommitReports())

		// Verify tracking after final execution
		assert.True(t, cache.earliestUnexecutedRoot.IsZero(), "After all roots executed, tracking should be reset")

		// Query after all executions
		queryTimestamp4 := cache.GetTimestampToQueryFrom(messageVisibilityWindow)
		assert.Equal(t, messageVisibilityWindow, queryTimestamp4, "Query after all executions should use visibility window")
	})
}

func TestCommitRootsCache_AdditionalEdgeCases(t *testing.T) {
	lggr := logger.Nop()
	messageVisibilityInterval := 8 * time.Hour
	rootSnoozeTime := 5 * time.Minute

	now := time.Now()
	messageVisibilityWindow := now.Add(-messageVisibilityInterval)

	selector := ccipocr3.ChainSelector(1)

	t.Run("Edge Case: Unexecuted root at exactly visibility window boundary", func(t *testing.T) {
		cache := internalNewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
			CleanupInterval,
			EvictionGracePeriod,
		)

		// Create a root exactly at the visibility window
		rootAtBoundary := ccipocr3.Bytes32{42}
		reportAtBoundary := createCommitData(messageVisibilityWindow, selector, rootAtBoundary)

		// Update cache
		cache.UpdateEarliestUnexecutedRoot(createCommitReports(reportAtBoundary))

		// Query timestamp should be the same as visibility window
		queryTimestamp := cache.GetTimestampToQueryFrom(messageVisibilityWindow)
		assert.Equal(t, messageVisibilityWindow, queryTimestamp,
			"For a root exactly at visibility window, should use visibility window")
	})

	t.Run("Edge Case: Roots with identical timestamps", func(t *testing.T) {
		cache := internalNewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
			CleanupInterval,
			EvictionGracePeriod,
		)

		// Create two roots with identical timestamps
		sameTimestamp := now.Add(-15 * time.Minute)
		root1 := ccipocr3.Bytes32{1}
		root2 := ccipocr3.Bytes32{2}

		report1 := createCommitData(sameTimestamp, selector, root1)
		report2 := createCommitData(sameTimestamp, selector, root2)

		// Update cache with both reports
		cache.UpdateEarliestUnexecutedRoot(createCommitReports(report1, report2))

		// Verify earliest is set correctly
		assert.Equal(t, sameTimestamp, cache.earliestUnexecutedRoot,
			"With identical timestamps, either can be earliest")

		// Execute one root
		cache.MarkAsExecuted(selector, root1)

		// Update cache with remaining report
		cache.UpdateEarliestUnexecutedRoot(createCommitReports(report2))

		// Verify earliest is still the same timestamp
		assert.Equal(t, sameTimestamp, cache.earliestUnexecutedRoot,
			"After executing one root, earliest should still be same timestamp")
	})

	t.Run("Edge Case: Visibility window moves forward", func(t *testing.T) {
		cache := internalNewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
			CleanupInterval,
			EvictionGracePeriod,
		)

		// Create two roots after the initial visibility window
		initialRoot := ccipocr3.Bytes32{77}
		initialTimestamp := messageVisibilityWindow.Add(15 * time.Minute)
		laterRoot := ccipocr3.Bytes32{88}
		laterTimestamp := messageVisibilityWindow.Add(30 * time.Minute)

		// For debugging
		t.Logf("Message visibility window: %v", messageVisibilityWindow)
		t.Logf("Initial root timestamp: %v", initialTimestamp)
		t.Logf("Later root timestamp: %v", laterTimestamp)

		// Test case 1: Only initial root - should use initial root timestamp
		cache.UpdateEarliestUnexecutedRoot(createCommitReports(
			createCommitData(initialTimestamp, selector, initialRoot)))

		queryTimestamp1 := cache.GetTimestampToQueryFrom(messageVisibilityWindow)
		t.Logf("Query timestamp 1: %v", queryTimestamp1)
		assert.Equal(t, initialTimestamp, queryTimestamp1,
			"With only initial root, should use its timestamp")

		// Test case 2: Add later root - should still use initial root timestamp
		// since it's the earliest unexecuted root
		cache.UpdateEarliestUnexecutedRoot(createCommitReports(
			createCommitData(initialTimestamp, selector, initialRoot),
			createCommitData(laterTimestamp, selector, laterRoot)))

		queryTimestamp2 := cache.GetTimestampToQueryFrom(messageVisibilityWindow)
		t.Logf("Query timestamp 2: %v", queryTimestamp2)
		assert.Equal(t, initialTimestamp, queryTimestamp2,
			"With both roots, should use earliest root timestamp")

		// Test case 3: Move visibility window past initial root
		// should now use laterRoot timestamp
		midVisibilityWindow := initialTimestamp.Add(5 * time.Minute)
		t.Logf("Mid visibility window: %v", midVisibilityWindow)

		// Execute initial root so only later root remains
		cache.MarkAsExecuted(selector, initialRoot)
		cache.UpdateEarliestUnexecutedRoot(createCommitReports(
			createCommitData(laterTimestamp, selector, laterRoot)))

		queryTimestamp3 := cache.GetTimestampToQueryFrom(midVisibilityWindow)
		t.Logf("Query timestamp 3: %v", queryTimestamp3)
		assert.Equal(t, laterTimestamp, queryTimestamp3,
			"With visibility window past initial root, should use later root timestamp")

		// Test case 4: Move visibility window past all roots
		// should use visibility window
		lateVisibilityWindow := laterTimestamp.Add(5 * time.Minute)
		t.Logf("Late visibility window: %v", lateVisibilityWindow)

		queryTimestamp4 := cache.GetTimestampToQueryFrom(lateVisibilityWindow)
		t.Logf("Query timestamp 4: %v", queryTimestamp4)
		assert.Equal(t, lateVisibilityWindow, queryTimestamp4,
			"With visibility window past all roots, should use visibility window")
	})

	t.Run("Edge Case: Change tracking after many execution rounds", func(t *testing.T) {
		cache := internalNewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
			CleanupInterval,
			EvictionGracePeriod,
		)

		// Create 20 roots with gradually increasing timestamps
		allReports := make(map[ccipocr3.ChainSelector][]exectypes.CommitData)
		allReports[selector] = make([]exectypes.CommitData, 0, 20)

		for i := 0; i < 20; i++ {
			root := ccipocr3.Bytes32{byte(i)}
			timestamp := now.Add(time.Duration(-100+i*5) * time.Minute)
			report := createCommitData(timestamp, selector, root)
			allReports[selector] = append(allReports[selector], report)
		}

		// Initial update with all reports
		cache.UpdateEarliestUnexecutedRoot(allReports)

		// Verify earliest is set to the first root
		assert.Equal(t, allReports[selector][0].Timestamp, cache.earliestUnexecutedRoot,
			"Initial earliest should be first root")

		// Execute roots one by one from first to last
		for i := 0; i < 20; i++ {
			cache.MarkAsExecuted(selector, allReports[selector][i].MerkleRoot)

			// Create remaining reports
			remainingReports := make(map[ccipocr3.ChainSelector][]exectypes.CommitData)
			if i < 19 {
				remainingReports[selector] = allReports[selector][i+1:]
				cache.UpdateEarliestUnexecutedRoot(remainingReports)

				// Verify earliest is updated correctly
				assert.Equal(t, allReports[selector][i+1].Timestamp, cache.earliestUnexecutedRoot,
					"After executing %d roots, earliest should be next root", i+1)
			} else {
				// After last execution, update with empty reports
				cache.UpdateEarliestUnexecutedRoot(make(map[ccipocr3.ChainSelector][]exectypes.CommitData))

				// Verify tracking is reset
				assert.True(t, cache.earliestUnexecutedRoot.IsZero(),
					"After executing all roots, tracking should be reset")
			}
		}
	})
}
