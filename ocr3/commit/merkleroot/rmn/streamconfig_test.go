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
		{roundInterval: 50 * time.Millisecond, wantRate: 120, wantCapacity: 15},
		{roundInterval: 100 * time.Millisecond, wantRate: 60, wantCapacity: 15},
		{roundInterval: time.Second, wantRate: 6, wantCapacity: 15},
		{roundInterval: 3 * time.Second, wantRate: 2, wantCapacity: 15},
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
		{roundInterval: 50 * time.Millisecond, wantRate: 2158120, wantCapacity: 269765},
		{roundInterval: 100 * time.Millisecond, wantRate: 1079060, wantCapacity: 269765},
		{roundInterval: time.Second, wantRate: 107906, wantCapacity: 269765},
		{roundInterval: 3 * time.Second, wantRate: 35968.667, wantCapacity: 269765},
		{roundInterval: 10 * time.Second, wantRate: 10790.6, wantCapacity: 269765},
		{roundInterval: 60 * time.Second, wantRate: 1798.433, wantCapacity: 269765},
		{roundInterval: 120 * time.Second, wantRate: 899.217, wantCapacity: 269765},
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
	assert.Equal(t, cfg.MessagesLimit.Rate, 6.0)
	assert.Equal(t, uint32(15), cfg.MessagesLimit.Capacity)

	// bytes RL
	assert.Equal(t, cfg.BytesLimit.Rate, 107906.0)
	assert.Equal(t, uint32(269765), cfg.BytesLimit.Capacity)
}
