package rmn

import (
	"math"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func Test_messagesLimit(t *testing.T) {
	testCases := []struct {
		roundInterval time.Duration
		wantRate      float64
		wantCapacity  uint32
	}{
		{roundInterval: 50 * time.Millisecond, wantRate: 15, wantCapacity: 15},
		{roundInterval: 100 * time.Millisecond, wantRate: 15, wantCapacity: 15},
		{roundInterval: time.Second, wantRate: 3.6, wantCapacity: 15},
		{roundInterval: 3 * time.Second, wantRate: 1.2, wantCapacity: 15},
		{roundInterval: 10 * time.Second, wantRate: minMsgLimitRate, wantCapacity: 15},
		{roundInterval: 60 * time.Second, wantRate: minMsgLimitRate, wantCapacity: 15},
		{roundInterval: 120 * time.Second, wantRate: minMsgLimitRate, wantCapacity: 15},
	}

	for _, tc := range testCases {
		t.Run(tc.roundInterval.String(), func(t *testing.T) {
			params := messagesLimit(tc.roundInterval)

			// round rate up to 4 decimal points
			params.Rate = math.Round(params.Rate*1000) / 1000

			assert.Equal(t, tc.wantRate, params.Rate)
			assert.Equal(t, tc.wantCapacity, params.Capacity)
		})
	}
}

func Test_bytesLimit(t *testing.T) {
	testCases := []struct {
		roundInterval time.Duration
		wantRate      float64
		wantCapacity  uint32
	}{
		{roundInterval: 50 * time.Millisecond, wantRate: 269765, wantCapacity: 269765},
		{roundInterval: 100 * time.Millisecond, wantRate: 269765, wantCapacity: 269765},
		{roundInterval: time.Second, wantRate: 64743.6, wantCapacity: 269765},
		{roundInterval: 3 * time.Second, wantRate: 21581.2, wantCapacity: 269765},
		{roundInterval: 10 * time.Second, wantRate: 6474.36, wantCapacity: 269765},
		{roundInterval: 60 * time.Second, wantRate: 1079.06, wantCapacity: 269765},
		{roundInterval: 120 * time.Second, wantRate: 539.53, wantCapacity: 269765},
	}

	for _, tc := range testCases {
		t.Run(tc.roundInterval.String(), func(t *testing.T) {
			params := bytesLimit(tc.roundInterval)

			// round rate up to 4 decimal points
			params.Rate = math.Round(params.Rate*1000) / 1000

			assert.Equal(t, tc.wantRate, params.Rate)
			assert.Equal(t, int(tc.wantCapacity), int(params.Capacity))
		})
	}
}

func Test_newStreamConfig(t *testing.T) {
	const kB = 1024
	lggr := logger.Test(t)
	streamName := "myCoolStream"

	cfg := newStreamConfig(lggr, streamName, time.Second)
	assert.Equal(t, streamName, cfg.StreamName)
	assert.Equal(t, 1, cfg.OutgoingBufferSize)
	assert.Equal(t, 1, cfg.IncomingBufferSize)

	// message length
	assert.Greater(t, cfg.MaxMessageLength, 52*kB)
	assert.Less(t, cfg.MaxMessageLength, 53*kB)

	// messages RL
	assert.Greater(t, cfg.MessagesLimit.Rate, 3.58)
	assert.Less(t, cfg.MessagesLimit.Rate, 3.60)
	assert.Equal(t, uint32(15), cfg.MessagesLimit.Capacity)

	// bytes RL
	assert.Greater(t, cfg.BytesLimit.Rate, 64743.59)
	assert.Less(t, cfg.BytesLimit.Rate, 64743.61)
	assert.Equal(t, uint32(269765), cfg.BytesLimit.Capacity)
}
