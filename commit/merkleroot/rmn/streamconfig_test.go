package rmn

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func Test_newStreamConfig(t *testing.T) {
	const kB = 1024
	lggr := logger.Test(t)
	streamName := "myCoolStream"

	cfg := newStreamConfig(lggr, streamName)

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
}
