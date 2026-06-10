package rpctimeout

import (
	"context"
	"time"
)

const (
	// Default is the ceiling for per-chain RPC calls when no OCR deadline is set.
	Default = 10 * time.Second

	ocrBudgetFraction = 0.8
)

// Duration returns the per-chain RPC timeout as min(Default, 80% of remaining OCR budget).
func Duration(ctx context.Context) time.Duration {
	deadline, ok := ctx.Deadline()
	if !ok {
		return Default
	}
	remaining := time.Until(deadline)
	if remaining <= 0 {
		return 0
	}
	ocrBudget := time.Duration(float64(remaining) * ocrBudgetFraction)
	return min(Default, ocrBudget)
}

// Context returns a child context with a per-chain RPC timeout derived from the parent OCR context.
func Context(ctx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(ctx, Duration(ctx))
}
