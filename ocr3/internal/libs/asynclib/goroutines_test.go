package asynclib

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

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
