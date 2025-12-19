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
	lggr := logger.Test(t)
	start := time.Now()

	ops := map[string]AsyncOperation{
		"slowOp": func(ctx context.Context, _ logger.Logger) any {
			select {
			case <-ctx.Done():
				return nil
			case <-time.After(24 * time.Hour):
				return "done"
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
			return "done"
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
	lggr := logger.Test(t)
	start := time.Now()

	ops := map[string]AsyncOperation{
		"hangingOp": func(ctx context.Context, _ logger.Logger) any {
			// This op ignores ctx and sleeps for a long time
			time.Sleep(24 * time.Hour)
			return "done"
		},
	}

	timeout := 100 * time.Millisecond
	results := ExecuteAsyncOperations(ctx, timeout, ops, lggr)

	elapsed := time.Since(start)
	assert.LessOrEqual(t, elapsed.Milliseconds(), int64(500), "timeout not respected even with hanging op")
	assert.Empty(t, results)
}
