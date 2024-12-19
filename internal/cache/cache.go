package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	NoExpiration = cache.NoExpiration
)

// CustomCache wraps go-cache with additional eviction policies
type CustomCache[V any] struct {
	cache        *cache.Cache
	customPolicy func(V) bool
	mutex        sync.RWMutex
}

// NewCustomCache creates a new cache with both time-based and custom eviction policies
func NewCustomCache[V any](
	defaultExpiration time.Duration,
	cleanupInterval time.Duration,
	customPolicy func(V) bool,
) *CustomCache[V] {
	return &CustomCache[V]{
		cache:        cache.New(defaultExpiration, cleanupInterval),
		customPolicy: customPolicy,
	}
}

// Set adds an item to the cache
func (c *CustomCache[V]) Set(key string, value V, expiration time.Duration) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache.Set(key, value, expiration)
}

// Get retrieves an item from the cache, checking both time-based and custom policies
func (c *CustomCache[V]) Get(key string) (V, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	var zero V
	value, found := c.cache.Get(key)
	if !found {
		return zero, false
	}

	// Type assertion
	typedValue, ok := value.(V)
	if !ok {
		return zero, false
	}

	// Check custom policy
	if c.customPolicy != nil && c.customPolicy(typedValue) {
		c.cache.Delete(key)
		return zero, false
	}

	return typedValue, true
}

// Delete removes an item from the cache
func (c *CustomCache[V]) Delete(key string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache.Delete(key)
}

// Items returns all items in the cache
func (c *CustomCache[V]) Items() map[string]V {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	items := c.cache.Items()
	result := make(map[string]V)

	for k, v := range items {
		if value, ok := v.Object.(V); ok {
			result[k] = value
		}
	}

	return result
}

// Flush removes all items from the cache
func (c *CustomCache[V]) Flush() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache.Flush()
}
