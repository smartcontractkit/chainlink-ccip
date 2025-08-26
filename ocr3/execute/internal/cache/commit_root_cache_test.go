package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// TestCanExecuteAndMarking tests the basic functionality of CanExecute, MarkAsExecuted, and Snooze methods
func TestCanExecuteAndMarking(t *testing.T) {
	lggr := logger.Nop()
	messageVisibilityInterval := 8 * time.Hour
	rootSnoozeTime := 5 * time.Minute

	selector := ccipocr3.ChainSelector(1)
	root1 := ccipocr3.Bytes32{1}
	root2 := ccipocr3.Bytes32{2}
	root3 := ccipocr3.Bytes32{3}

	t.Run("CanExecute returns true for new roots", func(t *testing.T) {
		cache := NewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
		)

		assert.True(t, cache.CanExecute(selector, root1), "New root should be executable")
		assert.True(t, cache.CanExecute(selector, root2), "New root should be executable")
	})

	t.Run("MarkAsExecuted prevents future execution", func(t *testing.T) {
		cache := NewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
		)

		// Initially all roots should be executable
		assert.True(t, cache.CanExecute(selector, root1), "New root should be executable")
		assert.True(t, cache.CanExecute(selector, root2), "New root should be executable")

		// Mark root1 as executed
		cache.MarkAsExecuted(selector, root1)

		// root1 should no longer be executable, but root2 should still be
		assert.False(t, cache.CanExecute(selector, root1), "Executed root should not be executable")
		assert.True(t, cache.CanExecute(selector, root2), "Unexecuted root should be executable")
	})

	t.Run("Snooze temporarily prevents execution", func(t *testing.T) {
		cache := NewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
		)

		// Initially all roots should be executable
		assert.True(t, cache.CanExecute(selector, root1), "New root should be executable")
		assert.True(t, cache.CanExecute(selector, root2), "New root should be executable")
		assert.True(t, cache.CanExecute(selector, root3), "New root should be executable")

		// Snooze root1
		cache.Snooze(selector, root1)

		// Mark root2 as executed
		cache.MarkAsExecuted(selector, root2)

		// root1 should be snoozed, root2 executed, root3 still executable
		assert.False(t, cache.CanExecute(selector, root1), "Snoozed root should not be executable")
		assert.False(t, cache.CanExecute(selector, root2), "Executed root should not be executable")
		assert.True(t, cache.CanExecute(selector, root3), "Untouched root should be executable")

		// In a real scenario, the snooze would expire after rootSnoozeTime
		// For testing, we can't wait, so we're just verifying it's snoozed initially
	})

	t.Run("Different chains are independent", func(t *testing.T) {
		cache := NewCommitRootsCache(
			lggr,
			messageVisibilityInterval,
			rootSnoozeTime,
		)

		selector1 := ccipocr3.ChainSelector(1)
		selector2 := ccipocr3.ChainSelector(2)

		// Mark root1 as executed on chain 1
		cache.MarkAsExecuted(selector1, root1)

		// root1 should not be executable on chain 1, but should be on chain 2
		assert.False(t, cache.CanExecute(selector1, root1), "Root should not be executable on chain where it was executed")
		assert.True(t, cache.CanExecute(selector2, root1), "Root should be executable on different chain")

		// Snooze root2 on chain 2
		cache.Snooze(selector2, root2)

		// root2 should not be executable on chain 2, but should be on chain 1
		assert.True(t, cache.CanExecute(selector1, root2), "Root should be executable on different chain")
		assert.False(t, cache.CanExecute(selector2, root2), "Root should not be executable on chain where it was snoozed")
	})
}
