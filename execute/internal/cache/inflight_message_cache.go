package cache

import (
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
)

// InflightMessageCache keeps track of messages that are currently in flight,
// used to prevent duplicate reports from being sent for the same message.
type InflightMessageCache struct {
	inflight *cache.Cache
}

// NewInflightMessageCache creates a new InflightMessageCache with the given cache expiry.
func NewInflightMessageCache(cacheExpiry time.Duration) *InflightMessageCache {
	return &InflightMessageCache{
		inflight: cache.New(cacheExpiry, CleanupInterval),
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
