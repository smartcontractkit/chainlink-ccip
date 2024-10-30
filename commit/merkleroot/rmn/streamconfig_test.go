package rmn

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

func Test_newStreamConfig(t *testing.T) {
	lggr := logger.Test(t)

	streamName := "myCoolStream"

	cfg := newStreamConfig(lggr, streamName)

	assert.Equal(t, streamName, cfg.StreamName)

	assert.Equal(t, 1, cfg.OutgoingBufferSize)

	assert.Equal(t, 1, cfg.IncomingBufferSize)

	// message length
	assert.Greater(t, cfg.MaxMessageLength, 25*megaByte)
	assert.Less(t, cfg.MaxMessageLength, 27*megaByte)

	// message rate
	assert.Greater(t, cfg.MessagesLimit.Rate, 3.4)
	assert.Less(t, cfg.MessagesLimit.Rate, 3.8)

	// message capacity
	assert.Equal(t, 9, int(cfg.MessagesLimit.Capacity))

	// bytes rate
	assert.Greater(t, cfg.BytesLimit.Rate, float64(30*megaByte))
	assert.Less(t, cfg.BytesLimit.Rate, float64(33*megaByte))

	// bytes capacity
	assert.Greater(t, int(cfg.BytesLimit.Capacity), 77*megaByte)
	assert.Less(t, int(cfg.BytesLimit.Capacity), 80*megaByte)
}

const (
	byt      = 1
	kiloByte = 1024 * byt
	megaByte = 1024 * kiloByte
)
