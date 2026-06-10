package rpctimeout

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDuration(t *testing.T) {
	t.Parallel()

	t.Run("uses default timeout without parent deadline", func(t *testing.T) {
		t.Parallel()
		assert.Equal(t, Default, Duration(context.Background()))
	})

	t.Run("caps at default timeout when OCR budget is larger", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		assert.Equal(t, Default, Duration(ctx))
	})

	t.Run("caps at 80 percent of remaining OCR budget when tighter", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
		defer cancel()

		got := Duration(ctx)
		want := time.Duration(float64(8*time.Second) * ocrBudgetFraction)
		assert.InDelta(t, float64(want), float64(got), float64(50*time.Millisecond))
	})

	t.Run("returns zero when parent deadline has expired", func(t *testing.T) {
		t.Parallel()
		ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(-time.Second))
		defer cancel()

		assert.Equal(t, time.Duration(0), Duration(ctx))
	})
}
