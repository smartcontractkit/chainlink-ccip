package rmn

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

//nolint:dupl
func Test_newStreamConfig(t *testing.T) {
	t.Skipf("until stream config is fine tuned")

	const kB = 1024
	lggr := logger.Test(t)
	streamName := "myCoolStream"

	t.Run("one second rounds", func(t *testing.T) {
		cfg := newStreamConfig(lggr, streamName, time.Second)
		assert.Equal(t, streamName, cfg.StreamName)
		assert.Equal(t, 1, cfg.OutgoingBufferSize)
		assert.Equal(t, 1, cfg.IncomingBufferSize)
		// message length
		assert.Greater(t, cfg.MaxMessageLength, 52*kB)
		assert.Less(t, cfg.MaxMessageLength, 53*kB)
		// message rate
		assert.Greater(t, cfg.MessagesLimit.Rate, 3.5)
		assert.Less(t, cfg.MessagesLimit.Rate, 3.6)
		// message capacity
		assert.Equal(t, 15, int(cfg.MessagesLimit.Capacity))
		// bytes rate
		assert.Greater(t, cfg.BytesLimit.Rate, float64(63*kB))
		assert.Less(t, cfg.BytesLimit.Rate, float64(64*kB))
		// bytes capacity
		assert.Greater(t, int(cfg.BytesLimit.Capacity), 263*kB)
		assert.Less(t, int(cfg.BytesLimit.Capacity), 265*kB)
	})

	t.Run("500ms rounds", func(t *testing.T) {
		cfg := newStreamConfig(lggr, streamName, 500*time.Millisecond)
		assert.Equal(t, streamName, cfg.StreamName)
		assert.Equal(t, 1, cfg.OutgoingBufferSize)
		assert.Equal(t, 1, cfg.IncomingBufferSize)
		// message length
		assert.Greater(t, cfg.MaxMessageLength, 52*kB)
		assert.Less(t, cfg.MaxMessageLength, 53*kB)
		// message rate
		assert.Greater(t, cfg.MessagesLimit.Rate, 7.1)
		assert.Less(t, cfg.MessagesLimit.Rate, 7.2)
		// message capacity
		assert.Equal(t, 30, int(cfg.MessagesLimit.Capacity))
		// bytes rate
		assert.Greater(t, cfg.BytesLimit.Rate, float64(126*kB))
		assert.Less(t, cfg.BytesLimit.Rate, float64(128*kB))
		// bytes capacity
		assert.Greater(t, int(cfg.BytesLimit.Capacity), 263*kB)
		assert.Less(t, int(cfg.BytesLimit.Capacity), 265*kB)
	})

	t.Run("4s rounds", func(t *testing.T) {
		cfg := newStreamConfig(lggr, streamName, 4*time.Second)
		assert.Equal(t, streamName, cfg.StreamName)
		assert.Equal(t, 1, cfg.OutgoingBufferSize)
		assert.Equal(t, 1, cfg.IncomingBufferSize)
		// message length
		assert.Greater(t, cfg.MaxMessageLength, 52*kB)
		assert.Less(t, cfg.MaxMessageLength, 53*kB)
		// message rate
		assert.Greater(t, cfg.MessagesLimit.Rate, 0.8)
		assert.Less(t, cfg.MessagesLimit.Rate, 0.9)
		// message capacity
		assert.Equal(t, 3, int(cfg.MessagesLimit.Capacity))
		// bytes rate
		assert.Greater(t, cfg.BytesLimit.Rate, float64(15*kB))
		assert.Less(t, cfg.BytesLimit.Rate, float64(17*kB))
		// bytes capacity
		assert.Greater(t, int(cfg.BytesLimit.Capacity), 263*kB)
		assert.Less(t, int(cfg.BytesLimit.Capacity), 265*kB)
	})

}
