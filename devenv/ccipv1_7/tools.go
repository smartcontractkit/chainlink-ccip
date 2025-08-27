package ccipv1_7

import (
	"time"

	"github.com/rs/zerolog"
)

/*
This code should be generalized and moved to devenv library after we finish CCIPv1.7 environment!
*/

type TimeTracker struct {
	logger    zerolog.Logger
	start     time.Time
	last      time.Time
	intervals []interval
}

type interval struct {
	tag   string
	delta time.Duration
}

// NewTimeTracker is a simple utility function that tracks execution time.
func NewTimeTracker(l zerolog.Logger) *TimeTracker { //nolint:gocritic
	now := time.Now()
	return &TimeTracker{
		start:     now,
		last:      now,
		logger:    l,
		intervals: make([]interval, 0),
	}
}

func (t *TimeTracker) Record(tag string) {
	now := time.Now()
	delta := now.Sub(t.last)
	t.intervals = append(t.intervals, interval{
		tag:   tag,
		delta: delta,
	})
	t.last = now
}

func (t *TimeTracker) Print() {
	total := time.Since(t.start)
	t.logger.Debug().Msg("Time tracking results:")
	for _, i := range t.intervals {
		t.logger.Debug().
			Str("Tag", i.tag).
			Str("Duration", i.delta.String()).
			Send()
	}

	t.logger.Debug().
		Str("Duration", total.String()).
		Msg("Total environment boot up time")
}
