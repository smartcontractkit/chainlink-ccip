package asynclib

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func TestWaitForAllNoErrOperations_AllOpsRun(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)
	var mu sync.Mutex
	var calledOps []string

	ops := AsyncNoErrOperationsMap{
		"op1": func(_ context.Context, _ logger.Logger) {
			mu.Lock()
			defer mu.Unlock()
			calledOps = append(calledOps, "op1")
		},
		"op2": func(_ context.Context, _ logger.Logger) {
			mu.Lock()
			defer mu.Unlock()
			calledOps = append(calledOps, "op2")
		},
	}

	WaitForAllNoErrOperations(ctx, 2*time.Second, ops, lggr)

	mu.Lock()
	defer mu.Unlock()
	assert.ElementsMatch(t, []string{"op1", "op2"}, calledOps)
}

func TestWaitForAllNoErrOperations_ContextTimeoutRespected(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)
	start := time.Now()

	ops := AsyncNoErrOperationsMap{
		"slowOp": func(ctx context.Context, _ logger.Logger) {
			select {
			case <-ctx.Done():
			case <-time.After(24 * time.Hour):
			}
		},
	}

	timeout := 100 * time.Millisecond
	WaitForAllNoErrOperations(ctx, timeout, ops, lggr)

	elapsed := time.Since(start)
	assert.LessOrEqual(t, elapsed.Milliseconds(), int64(500), "timeout not respected")
}

func TestWaitForAllNoErrOperations_ContextIsPropagated(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	done := make(chan struct{})

	ops := AsyncNoErrOperationsMap{
		"checkCtx": func(ctx context.Context, _ logger.Logger) {
			_, ok := ctx.Deadline()
			assert.True(t, ok, "context should have a deadline")
			close(done)
		},
	}

	WaitForAllNoErrOperations(ctx, 500*time.Millisecond, ops, lggr)

	select {
	case <-done:
	case <-time.After(24 * time.Hour):
		t.Fatal("operation did not complete")
	}
}

func TestWaitForAllNoErrOperations_NoOps(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	// should not panic or block
	WaitForAllNoErrOperations(ctx, 500*time.Millisecond, AsyncNoErrOperationsMap{}, lggr)
}

func TestAsyncOpsRunner_TimeoutRespected(t *testing.T) {
	// Create a runner with a small pool size
	poolSize := map[string]int{"op1": 1}
	runner, err := NewAsyncOpsRunner(poolSize)
	require.NoError(t, err)

	lggr := logger.Nop()

	// Define a hanging operation that ignores the context
	hangingOp := func(_ context.Context, _ logger.Logger) {
		// Simulate work that takes much longer than the timeout
		// and deliberately ignores ctx.Done()
		time.Sleep(2 * time.Second)
	}

	ops := AsyncNoErrOperationsMap{
		"op1": hangingOp,
	}

	start := time.Now()
	timeout := 100 * time.Millisecond

	// Run should return after roughly 'timeout', not 2 seconds
	runner.Run(context.Background(), timeout, ops, lggr)

	elapsed := time.Since(start)

	// Allow a small margin for scheduler overhead, but it should be much less than the 2s sleep
	assert.Less(t, elapsed, 1*time.Second, "Run took too long, timeout not respected")
	assert.GreaterOrEqual(t, elapsed, timeout, "Run returned too early")
}

func TestAsyncOpsRunner_NormalCompletion(t *testing.T) {
	poolSize := map[string]int{"op1": 1}
	runner, err := NewAsyncOpsRunner(poolSize)
	require.NoError(t, err)

	lggr := logger.Nop()

	// Define a fast operation
	fastOp := func(_ context.Context, _ logger.Logger) {
		time.Sleep(10 * time.Millisecond)
	}

	ops := AsyncNoErrOperationsMap{
		"op1": fastOp,
	}

	timeout := 1 * time.Second
	start := time.Now()

	runner.Run(context.Background(), timeout, ops, lggr)

	elapsed := time.Since(start)
	assert.Less(t, elapsed, timeout, "Run should return immediately after tasks complete, not wait for timeout")
}

func TestAsyncOpsRunner_PoolFullDoesNotBlock(t *testing.T) {
	// Create a runner with a pool size of 1 for "op1"
	poolSize := map[string]int{"op1": 1}
	runner, err := NewAsyncOpsRunner(poolSize)
	require.NoError(t, err)

	lggr := logger.Nop()

	// 1. Start a long-running task that occupies the pool
	// We use a channel to signal when it's running to avoid race conditions with sleep
	startedCh := make(chan struct{})
	blockingOp := func(_ context.Context, _ logger.Logger) {
		close(startedCh)
		time.Sleep(500 * time.Millisecond) // Occupy the worker
	}

	// Launch it in a separate goroutine so we can try to schedule another one immediately
	go func() {
		ops := AsyncNoErrOperationsMap{"op1": blockingOp}
		runner.Run(context.Background(), 1*time.Second, ops, lggr)
	}()

	// Wait for the first task to actually start running and occupy the pool
	select {
	case <-startedCh:
	case <-time.After(1 * time.Second):
		t.Fatal("blocking op never started")
	}

	// 2. Try to run another task on the same pool
	// Since the pool size is 1 and it's occupied, this submission should fail immediately
	// due to WithNonblocking(true).
	// Because of our "continue" fix, Run should just skip this task and return immediately
	// since there are no other tasks to wait for.
	start := time.Now()
	ops := AsyncNoErrOperationsMap{"op1": func(ctx context.Context, l logger.Logger) {}}

	// This should return quickly, NOT block waiting for the first op to finish
	runner.Run(context.Background(), 2*time.Second, ops, lggr)

	elapsed := time.Since(start)
	assert.Less(t, elapsed, 200*time.Millisecond, "Run blocked despite pool being full")
}
