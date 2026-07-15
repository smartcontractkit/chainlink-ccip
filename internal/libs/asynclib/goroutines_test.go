package asynclib

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func TestRunner_AllOpsRun(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	ops := AsyncNoErrOperationsMap{
		"op1": func(_ context.Context, _ logger.Logger) any { return 1 },
		"op2": func(_ context.Context, _ logger.Logger) any { return "two" },
	}

	results := NewRunner().WaitForAll(ctx, 2*time.Second, ops, lggr)

	require.Len(t, results, 2)
	assert.Equal(t, 1, results["op1"])
	assert.Equal(t, "two", results["op2"])
}

func TestRunner_NoOps(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	// should not panic or block
	results := NewRunner().WaitForAll(ctx, 500*time.Millisecond, AsyncNoErrOperationsMap{}, lggr)
	assert.Empty(t, results)
}

func TestRunner_ContextIsPropagated(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	ops := AsyncNoErrOperationsMap{
		"checkCtx": func(ctx context.Context, _ logger.Logger) any {
			_, ok := ctx.Deadline()
			assert.True(t, ok, "context should have a deadline")
			return nil
		},
	}

	results := NewRunner().WaitForAll(ctx, 500*time.Millisecond, ops, lggr)
	require.Contains(t, results, "checkCtx")
}

func TestRunner_ContextHonoringOpReturnsOnTimeout(t *testing.T) {
	ctx := context.Background()
	// Nop logger: the op's goroutine finishes right as the timeout fires, so it may log after the
	// test returns, which the zap test logger treats as a fatal error.
	lggr := logger.Nop()
	start := time.Now()

	ops := AsyncNoErrOperationsMap{
		"slowOp": func(ctx context.Context, _ logger.Logger) any {
			select {
			case <-ctx.Done():
			case <-time.After(24 * time.Hour):
			}
			return nil
		},
	}

	results := NewRunner().WaitForAll(ctx, 100*time.Millisecond, ops, lggr)

	assert.LessOrEqual(t, time.Since(start).Milliseconds(), int64(500), "timeout not respected")
	// The op returned (nil) right as the context was cancelled, so it may or may not be recorded;
	// either way the call must not have blocked past the timeout, which the assertion above checks.
	_ = results
}

// TestRunner_CtxIgnoringOpDoesNotBlock is the case that would hang the old helper: an operation that
// ignores context cancellation. WaitForAll must still return within the timeout, and the stuck op
// must be absent from the results (so the caller degrades to a zero value rather than blocking).
func TestRunner_CtxIgnoringOpDoesNotBlock(t *testing.T) {
	ctx := context.Background()
	// Nop logger: the abandoned goroutine is released in cleanup and logs after the test returns.
	lggr := logger.Nop()

	release := make(chan struct{})
	t.Cleanup(func() { close(release) }) // let the abandoned goroutine exit at test end

	ops := AsyncNoErrOperationsMap{
		"fastOp": func(_ context.Context, _ logger.Logger) any { return "done" },
		"stuckOp": func(_ context.Context, _ logger.Logger) any {
			<-release // deliberately ignores ctx
			return "too late"
		},
	}

	start := time.Now()
	results := NewRunner().WaitForAll(ctx, 100*time.Millisecond, ops, lggr)

	assert.LessOrEqual(t, time.Since(start).Milliseconds(), int64(500), "did not return within timeout")
	assert.Equal(t, "done", results["fastOp"], "fast op result should be present")
	assert.NotContains(t, results, "stuckOp", "stuck op must not appear in results")
}

// TestRunner_StuckOpNotRespawned verifies that an operation still running from a previous round is
// not spawned again, so a wedged dependency cannot make goroutines accumulate across rounds.
func TestRunner_StuckOpNotRespawned(t *testing.T) {
	ctx := context.Background()
	// Nop logger: the abandoned goroutine is released in cleanup and logs after the test returns.
	lggr := logger.Nop()

	var calls atomic.Int32
	started := make(chan struct{}, 8)
	release := make(chan struct{})
	t.Cleanup(func() { close(release) })

	ops := AsyncNoErrOperationsMap{
		"stuckOp": func(_ context.Context, _ logger.Logger) any {
			calls.Add(1)
			started <- struct{}{}
			<-release // ignores ctx: stays in flight across rounds
			return nil
		},
	}

	r := NewRunner()

	// Round 1: spawns the op, which starts and then blocks past the timeout.
	round1 := r.WaitForAll(ctx, 50*time.Millisecond, ops, lggr)
	<-started // ensure the op actually began
	assert.Equal(t, int32(1), calls.Load())
	assert.NotContains(t, round1, "stuckOp")

	// Round 2: the op is still in flight, so it must be skipped (not re-spawned).
	round2 := r.WaitForAll(ctx, 50*time.Millisecond, ops, lggr)
	assert.Equal(t, int32(1), calls.Load(), "stuck op should not be spawned a second time")
	assert.NotContains(t, round2, "stuckOp")
}
