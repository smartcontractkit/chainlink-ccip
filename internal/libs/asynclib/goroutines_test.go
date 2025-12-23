package asynclib

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func TestExecuteAsyncOperations_AllOpsRun(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)
	var mu sync.Mutex
	var calledOps []string

	ops := map[string]AsyncOperation{
		"op1": func(_ context.Context, _ logger.Logger) any {
			mu.Lock()
			defer mu.Unlock()
			calledOps = append(calledOps, "op1")
			return "op1"
		},
		"op2": func(_ context.Context, _ logger.Logger) any {
			mu.Lock()
			defer mu.Unlock()
			calledOps = append(calledOps, "op2")
			return "op2"
		},
	}

	results := ExecuteAsyncOperations(ctx, 2*time.Second, ops, lggr)

	mu.Lock()
	defer mu.Unlock()
	assert.ElementsMatch(t, []string{"op1", "op2"}, calledOps)
	assert.Equal(t, 2, len(results))
}

func TestExecuteAsyncOperations_ContextTimeoutRespected(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Nop()
	start := time.Now()

	ops := map[string]AsyncOperation{
		"slowOp": func(ctx context.Context, _ logger.Logger) any {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(24 * time.Hour):
				return "done1"
			}
		},
	}

	timeout := 100 * time.Millisecond
	results := ExecuteAsyncOperations(ctx, timeout, ops, lggr)

	elapsed := time.Since(start)
	assert.LessOrEqual(t, elapsed.Milliseconds(), int64(500), "timeout not respected")
	assert.Empty(t, results)
}

func TestExecuteAsyncOperations_ContextIsPropagated(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	done := make(chan struct{})

	ops := map[string]AsyncOperation{
		"checkCtx": func(ctx context.Context, _ logger.Logger) any {
			_, ok := ctx.Deadline()
			assert.True(t, ok, "context should have a deadline")
			close(done)
			return "done2"
		},
	}

	ExecuteAsyncOperations(ctx, 500*time.Millisecond, ops, lggr)

	select {
	case <-done:
	case <-time.After(24 * time.Hour):
		t.Fatal("operation did not complete")
	}
}

func TestExecuteAsyncOperations_NoOps(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	// should not panic or block
	results := ExecuteAsyncOperations(ctx, 500*time.Millisecond, map[string]AsyncOperation{}, lggr)
	assert.Empty(t, results)
}

func TestExecuteAsyncOperations_HangingOpReturns(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Nop()
	start := time.Now()

	ops := map[string]AsyncOperation{
		"hangingOp": func(ctx context.Context, _ logger.Logger) any {
			// This op ignores ctx and sleeps for a long time
			time.Sleep(24 * time.Hour)
			return "done3"
		},
	}

	timeout := 100 * time.Millisecond
	results := ExecuteAsyncOperations(ctx, timeout, ops, lggr)

	elapsed := time.Since(start)
	assert.LessOrEqual(t, elapsed.Milliseconds(), int64(500), "timeout not respected even with hanging op")
	assert.Empty(t, results)
}

func TestExecuteAsyncOperations_PartialSuccess(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Nop()

	ops := map[string]AsyncOperation{
		"fastOp": func(_ context.Context, _ logger.Logger) any {
			return "success"
		},
		"slowOp": func(ctx context.Context, _ logger.Logger) any {
			// This simulates an operation that takes longer than the timeout
			select {
			case <-ctx.Done():
				return nil // or return "partial" but context cancellation usually means incomplete
			case <-time.After(2 * time.Second):
				return "too_slow"
			}
		},
	}

	// Timeout shorter than slowOp but enough for fastOp
	timeout := 100 * time.Millisecond
	results := ExecuteAsyncOperations(ctx, timeout, ops, lggr)

	assert.Equal(t, 1, len(results), "expected only one successful operation")
	assert.Equal(t, "success", results["fastOp"])
	_, exists := results["slowOp"]
	assert.False(t, exists, "slowOp should not be in results")
}

func TestExecuteAsyncOperations_NilResultIgnored(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	ops := map[string]AsyncOperation{
		"nilOp": func(_ context.Context, _ logger.Logger) any {
			return nil
		},
		"validOp": func(_ context.Context, _ logger.Logger) any {
			return "value"
		},
	}

	results := ExecuteAsyncOperations(ctx, 500*time.Millisecond, ops, lggr)

	assert.Equal(t, 1, len(results))
	assert.Equal(t, "value", results["validOp"])
	_, exists := results["nilOp"]
	assert.False(t, exists, "nilOp result should be ignored")
}

func TestExecuteAsyncOperations_InfiniteWait(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)

	ops := map[string]AsyncOperation{
		"op1": func(_ context.Context, _ logger.Logger) any {
			time.Sleep(200 * time.Millisecond)
			return "done4"
		},
	}

	// 0 timeout means wait indefinitely (or until parent context cancels)
	results := ExecuteAsyncOperations(ctx, 0, ops, lggr)

	assert.Equal(t, 1, len(results))
	assert.Equal(t, "done4", results["op1"])
}

func TestExecuteAsyncOperations_PanicRecovered(t *testing.T) {
	ctx := context.Background()
	lggr := logger.Test(t)
	start := time.Now()

	ops := map[string]AsyncOperation{
		"panicOp": func(_ context.Context, _ logger.Logger) any {
			panic("oops")
		},
		"validOp": func(_ context.Context, _ logger.Logger) any {
			return "value"
		},
	}

	// Long timeout to ensure we don't return due to timeout but due to completion (including panic recovery)
	results := ExecuteAsyncOperations(ctx, 5*time.Second, ops, lggr)

	elapsed := time.Since(start)
	assert.Less(t, elapsed.Milliseconds(), int64(1000), "should return immediately after panic, not wait for timeout")
	assert.Equal(t, 1, len(results))
	assert.Equal(t, "value", results["validOp"])
}
