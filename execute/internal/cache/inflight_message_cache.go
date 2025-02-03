package cache

import (
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

type InflightMessageCache struct {
	lggr            logger.Logger
	flightAllowance time.Duration

	// inflight is a cache of messages that are currently in flight. This cache
	// is used to prevent duplicate reports from being sent for the same message.
	inflight *cache.Cache
}

func NewInflightMessageCache(lggr logger.Logger, flightAllowance time.Duration) *InflightMessageCache {
	return &InflightMessageCache{
		lggr:            lggr,
		flightAllowance: flightAllowance,
		inflight:        cache.New(flightAllowance, CleanupInterval),
	}
}

func (c *InflightMessageCache) IsInflight(msgID ccipocr3.Bytes32) bool {
	_, found := c.inflight.Get(msgID.String())
	return found
}

func (c *InflightMessageCache) MarkInflight(msgID ccipocr3.Bytes32) {
	c.inflight.SetDefault(msgID.String(), struct{}{})
}

func (c *InflightMessageCache) Delete(msgID ccipocr3.Bytes32) {
	c.inflight.Delete(msgID.String())
}
