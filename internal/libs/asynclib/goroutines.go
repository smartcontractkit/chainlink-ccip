package asynclib

import (
	"context"
	"maps"
	"slices"
	"sync"
	"time"

	"github.com/google/uuid"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type AsyncOperation func(ctx context.Context, l logger.Logger) any

// ExecuteAsyncOperations spawns goroutines for each operation in the map and waits for
// all of them to finish or for the timeout to expire.
// It returns a map of results for operations that completed successfully within the timeout.
// If an operation times out, it will not be present in the returned map.
// If timeout is 0, no timeout is applied.
func ExecuteAsyncOperations(
	ctx context.Context,
	timeout time.Duration,
	operations map[string]AsyncOperation,
	lggr logger.Logger,
) map[string]any {
	var (
		callCtx context.Context
		cancel  context.CancelFunc
	)

	if timeout > 0 {
		callCtx, cancel = context.WithTimeout(ctx, timeout)
	} else {
		callCtx, cancel = context.WithCancel(ctx)
	}
	defer cancel()

	lggr.Debugw("spawning goroutines", "timeout", timeout)

	type result struct {
		name string
		val  any
	}

	resultsChan := make(chan result, len(operations))
	pendingOps := make(map[string]struct{}, len(operations))

	for opName, op := range operations {
		pendingOps[opName] = struct{}{}
		go func(name string, operation AsyncOperation) {
			defer func() {
				if r := recover(); r != nil {
					lggr.Errorw("AsyncOperation panicked", "opName", name, "panic", r)
					resultsChan <- result{name: name, val: nil}
				}
			}()

			opID := uuid.New().String()
			l := logger.With(lggr, "opName", name, "opID", opID)
			tStart := time.Now()
			val := operation(callCtx, l)
			l.Debugw("observing goroutine finished", "duration", time.Since(tStart))
			resultsChan <- result{name: name, val: val}
		}(opName, op)
	}

	resultsMap := make(map[string]any)
	completed := 0

	for completed < len(operations) {
		select {
		case res := <-resultsChan:
			delete(pendingOps, res.name)
			if res.val != nil {
				resultsMap[res.name] = res.val
			}
			completed++
		case <-callCtx.Done():
			if timeout > 0 && callCtx.Err() != nil {
				lggr.Warnw(
					"ExecuteAsyncOperations's context is done before all operations completed!!!"+
						"This indicates a potential issue in one of the operations.",
					"timeout", timeout,
					"completed", completed,
					"total", len(operations),
					"ctxErr", callCtx.Err(),
					"allOperations", slices.Collect(maps.Keys(operations)),
					"pendingOperations", slices.Collect(maps.Keys(pendingOps)),
				)
			}
			return resultsMap
		}
	}

	return resultsMap
}

// WrapWithSingleFlight ensures that an operation does not overlap with its previous execution.
// It uses a sync.Map to track currently running operations by name.
// If an operation with the same name is already running, the new call returns immediately with nil.
func WrapWithSingleFlight(
	runningOps *sync.Map,
	name string,
	op AsyncOperation,
) AsyncOperation {
	return func(ctx context.Context, l logger.Logger) any {
		if _, loaded := runningOps.LoadOrStore(name, true); loaded {
			l.Warnw("skipping operation because previous run is still active", "opName", name)
			return nil
		}
		defer runningOps.Delete(name)
		return op(ctx, l)
	}
}
