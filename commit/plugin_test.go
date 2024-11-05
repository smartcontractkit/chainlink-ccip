package commit

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_queryPhaseRmnRelatedTimers(t *testing.T) {
	testCases := []struct {
		maxQueryDuration           time.Duration
		expInitialObservationTimer time.Duration
		expInitialReportTimer      time.Duration
	}{
		{
			maxQueryDuration:           1 * time.Second,
			expInitialObservationTimer: 550 * time.Millisecond,
			expInitialReportTimer:      200 * time.Millisecond,
		},
		{
			maxQueryDuration:           3 * time.Second,
			expInitialObservationTimer: 1650 * time.Millisecond,
			expInitialReportTimer:      600 * time.Millisecond,
		},
		{
			maxQueryDuration:           5 * time.Second,
			expInitialObservationTimer: 2750 * time.Millisecond,
			expInitialReportTimer:      1000 * time.Millisecond,
		},
		{
			maxQueryDuration:           15 * time.Second,
			expInitialObservationTimer: 3 * time.Second,
			expInitialReportTimer:      2 * time.Second,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.maxQueryDuration.String(), func(t *testing.T) {
			obsTimer := observationsInitialRequestTimerDuration(tc.maxQueryDuration).Round(time.Millisecond)
			sigTimer := reportsInitialRequestTimerDuration(tc.maxQueryDuration).Round(time.Millisecond)
			assert.Equal(t, tc.expInitialObservationTimer, obsTimer)
			assert.Equal(t, tc.expInitialReportTimer, sigTimer)
		})
	}
}
