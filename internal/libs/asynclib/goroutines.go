package asynclib

import (
	"context"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type AsyncNoErrOperationsMap map[string]func(ctx context.Context, l logger.Logger)

// WaitForAllNoErrOperations spawns goroutines for each operation in the map and waits for
// all of them to finish. It creates a child context with a timeout to ensure that the operations
// do not run indefinitely.
// If the timeout is reached, the function returns, even if some operations are still running.
// This is to prevent the caller from blocking forever if an operation hangs and does not respect the context.
// If timeout is 0, no timeout is applied (waits indefinitely or until parent context is cancelled).
func WaitForAllNoErrOperations(
	ctx context.Context,
	timeout time.Duration,
	operations AsyncNoErrOperationsMap,
	lggr logger.Logger,
) {
	var callCtx context.Context
	var cancel context.CancelFunc

	if timeout > 0 {
		callCtx, cancel = context.WithTimeout(ctx, timeout)
	} else {
		callCtx, cancel = context.WithCancel(ctx)
	}
	defer cancel()

	lggr.Debugw("spawning goroutines", "timeout", timeout)

	var wg sync.WaitGroup
	wg.Add(len(operations))

	for opName, op := range operations {
		go func(opName string, op func(context.Context, logger.Logger)) {
			defer wg.Done()
			tStart := time.Now()
			op(callCtx, logger.With(lggr, "opID", opName))
			lggr.Debugw("observing goroutine finished", "opID", opName, "duration", time.Since(tStart))
		}(opName, op)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return
	case <-callCtx.Done():
		if timeout > 0 && callCtx.Err() == context.DeadlineExceeded {
			lggr.Warnw("WaitForAllNoErrOperations timed out before all operations completed; this indicates a bug in one of the operations not respecting context cancellation", "timeout", timeout)
		}
		return
	}
}

// WrapWithSingleFlight ensures that an operation does not overlap with its previous execution.
// It uses a sync.Map to track currently running operations by name.
// If an operation with the same name is already running, the new call returns immediately.
func WrapWithSingleFlight(
	runningOps *sync.Map,
	name string,
	op func(context.Context, logger.Logger),
) func(context.Context, logger.Logger) {
	return func(ctx context.Context, l logger.Logger) {
		if _, loaded := runningOps.LoadOrStore(name, true); loaded {
			l.Warnw("skipping operation because previous run is still active", "opID", name)
			return
		}
		defer runningOps.Delete(name)
		op(ctx, l)
	}
}
