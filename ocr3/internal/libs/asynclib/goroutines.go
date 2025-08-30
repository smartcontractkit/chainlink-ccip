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
