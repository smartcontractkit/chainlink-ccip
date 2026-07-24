package asynclib

import (
	"context"
	"maps"
	"sync"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

// AsyncNoErrOperationsMap maps an operation name to a function that performs the operation and
// returns its result. Operations must not return an error; recoverable failures should be handled
// internally (e.g. by returning an empty/zero result and logging).
type AsyncNoErrOperationsMap map[string]func(ctx context.Context, l logger.Logger) any

// Runner executes a set of named operations concurrently on each call to WaitForAll. It bounds the
// number of goroutines it spawns: it never starts a second goroutine for an operation whose previous
// invocation is still running. A still-running operation is treated as stuck and skipped
// for the current round, so a wedged dependency (i.e. stuck LOOP plugin)
// cannot cause goroutines to accumulate across rounds.
//
// A Runner is safe for concurrent use and is intended to be created once and reused across rounds
// so that its in-flight tracking persists between calls.
type Runner struct {
	mu       sync.Mutex
	inFlight map[string]bool
}

// NewRunner returns a ready-to-use Runner.
func NewRunner() *Runner {
	return &Runner{inFlight: make(map[string]bool)}
}

// WaitForAll runs each operation in its own goroutine under a child context bounded by timeout, and
// returns a map of the results of the operations that COMPLETED within the timeout. Operations that
// do not complete in time (or whose previous invocation is still running) are absent from the
// returned map, so callers should treat a missing key as "no fresh result this round" (its zero
// value) rather than blocking on it.
//
// WaitForAll never blocks past the timeout, even if an operation ignores context cancellation. Such
// an operation's goroutine is abandoned (it keeps running until it returns on its own), but:
//   - it is not re-spawned on subsequent rounds while it remains in flight, so goroutines do not
//     accumulate; and
//   - its late result is discarded rather than served, so consensus never acts on stale data.
//
// This trades a bounded, loudly-logged goroutine leak in the rare stuck case for guaranteed
// liveness of the caller.
func (r *Runner) WaitForAll(
	ctx context.Context,
	timeout time.Duration,
	operations AsyncNoErrOperationsMap,
	lggr logger.Logger,
) map[string]any {
	callCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	lggr.Debugw("spawning operation goroutines", "timeout", timeout)

	// results is scoped to this round: the returned snapshot is cloned from it, and any late write
	// from an abandoned goroutine lands here (never in a future round's map), so stale results are
	// never served.
	var (
		resMu   sync.Mutex
		results = make(map[string]any)
	)
	// Buffered to the number of operations so a late-finishing (abandoned) goroutine can always
	// signal completion without blocking, even after we have stopped waiting.
	doneCh := make(chan struct{}, len(operations))

	spawned := 0
	for opName, op := range operations {
		r.mu.Lock()
		if r.inFlight[opName] {
			r.mu.Unlock()
			// A goroutine from a previous round is still running for this op: don't pile on.
			lggr.Warnw("skipping operation: previous invocation still running (likely stuck)", "opID", opName)
			continue
		}
		r.inFlight[opName] = true
		r.mu.Unlock()

		spawned++
		go func(opName string, op func(context.Context, logger.Logger) any) {
			defer func() { doneCh <- struct{}{} }()
			defer func() {
				r.mu.Lock()
				r.inFlight[opName] = false
				r.mu.Unlock()
			}()

			tStart := time.Now()
			res := op(callCtx, logger.With(lggr, "opID", opName))
			resMu.Lock()
			results[opName] = res
			resMu.Unlock()
			lggr.Debugw("operation goroutine finished", "opID", opName, "duration", time.Since(tStart))
		}(opName, op)
	}

	// Wait for the spawned operations to finish, or the timeout to elapse, whichever comes first.
	timedOut := false
waitLoop:
	for range spawned {
		select {
		case <-doneCh:
		case <-callCtx.Done():
			timedOut = true
			break waitLoop
		}
	}
	if timedOut {
		lggr.Errorw(
			"timed out waiting for async operations; proceeding with completed results only",
			"timeout", timeout,
			"err", context.Cause(callCtx),
		)
	}

	resMu.Lock()
	defer resMu.Unlock()

	// Clone to avoid a data race: a late-finishing (abandoned) goroutine
	// may still write to results after we return it, so the
	// returned map must be a snapshot of the results at this point in time.
	return maps.Clone(results)
}
