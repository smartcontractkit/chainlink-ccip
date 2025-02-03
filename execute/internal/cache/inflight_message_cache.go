package cache

import (
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

func toID(src ccipocr3.ChainSelector, msgID ccipocr3.Bytes32) string {
	return src.String() + msgID.String()
}

func (c *InflightMessageCache) IsInflight(src ccipocr3.ChainSelector, msgID ccipocr3.Bytes32) bool {
	_, found := c.inflight.Get(toID(src, msgID))
	return found
}

func (c *InflightMessageCache) MarkInflight(src ccipocr3.ChainSelector, msgID ccipocr3.Bytes32) {
	c.inflight.SetDefault(toID(src, msgID), struct{}{})
}

func (c *InflightMessageCache) Delete(src ccipocr3.ChainSelector, msgID ccipocr3.Bytes32) {
	c.inflight.Delete(toID(src, msgID))
}
