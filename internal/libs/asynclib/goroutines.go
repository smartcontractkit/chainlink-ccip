package asynclib

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/panjf2000/ants/v2"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type AsyncNoErrOperationsMap map[string]func(ctx context.Context, l logger.Logger)

// WaitForAllNoErrOperations spawns goroutines for each operation in the map and waits for
// all of them to finish. It creates a child context with a timeout to ensure that the operations
// do not run indefinitely.
func WaitForAllNoErrOperations(
	ctx context.Context,
	timeout time.Duration,
	operations AsyncNoErrOperationsMap,
	lggr logger.Logger,
) {
	callCtx, cancel := context.WithTimeout(ctx, timeout)
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

	wg.Wait()
}

// AsyncOpsRunner is a runner that runs operations asynchronously in a pool
// and respects context timeouts. It is safe for concurrent use.
type AsyncOpsRunner struct {
	pools map[string]*ants.Pool // one pool per op
}

// NewAsyncOpsRunner creates a new AsyncOpsRunner. It returns an error if the pools cannot be created.
func NewAsyncOpsRunner(poolSizePerOp map[string]int) (*AsyncOpsRunner, error) {
	r := &AsyncOpsRunner{pools: make(map[string]*ants.Pool)}
	for opName, poolSize := range poolSizePerOp {
		if poolSize <= 0 {
			return nil, fmt.Errorf("pool size must be greater than 0, got %d", poolSize)
		}
		p, err := ants.NewPool(
			poolSize,
			// critical: don't block Submit() calls if not enough workers in the pool
			ants.WithNonblocking(true),
		)
		if err != nil {
			return nil, err
		}

		r.pools[opName] = p
	}

	return r, nil
}

// Run runs the operations asynchronously in the pools.
// It returns when the context is done or all the operations are complete.
func (a *AsyncOpsRunner) Run(
	ctx context.Context,
	timeout time.Duration,
	operations AsyncNoErrOperationsMap,
	lggr logger.Logger,
) {
	callCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	lggr.Debugw("spawning goroutines", "timeout", timeout)

	// We use a buffered channel to avoid blocking the worker goroutines if this function
	// returns early (e.g. due to context cancellation). We also avoid closing the channel
	// to prevent "send on closed channel" panics in the workers.
	doneCh := make(chan struct{}, len(operations))
	// numExpectedTasks may not equal len(operations) if some operations are skipped due to config errors
	// or if the pool is full.
	numExpectedTasks := 0
	for opName, opFunc := range operations {
		p, ok := a.pools[opName]
		if !ok {
			// don't have a pool for this op, probably config error
			lggr.Warnw("pool not found for op, config error?", "op", "opName")
			continue
		}

		err := p.Submit(func() {
			defer func() {
				doneCh <- struct{}{}
			}()
			tStart := time.Now()
			opFunc(callCtx, logger.With(lggr, "opID", opName))
			lggr.Debugw("observing goroutine finished", "opID", opName, "duration", time.Since(tStart))
		})
		if err != nil {
			if errors.Is(err, ants.ErrPoolOverload) {
				lggr.Errorw("couldn't start worker for op, pool is full", "op", opName, "err", err)
			} else {
				lggr.Errorw("couldn't start worker for op, some other error", "op", opName, "err", err)
			}
			continue
		}
		numExpectedTasks++
	}

	if numExpectedTasks == 0 {
		lggr.Infow("no tasks to run, returning early")
		return
	}

	// wait for the context to be canceled or everything to return
	tasksDone := 0
	for {
		select {
		case <-callCtx.Done():
			lggr.Infow("async ops runner ctx done, potentially not all tasks complete", "err", callCtx.Err())
			return
		case _, ok := <-doneCh:
			if !ok {
				// channel closed - shouldn't happen because we don't close it.
				lggr.Errorw("async ops runner doneCh closed, this should not happen! Returning early.")
				return
			}
			tasksDone++
			if tasksDone == numExpectedTasks {
				lggr.Infow("async ops runner done all tasks, returning")
				return
			} else {
				lggr.Infow("async ops runner task done",
					"tasksDone", tasksDone,
					"numExpectedTasks", numExpectedTasks)
			}
		}
	}
}
