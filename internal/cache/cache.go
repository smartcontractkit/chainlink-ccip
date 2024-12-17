package cache

import (
	"sync"
	"time"
)

// Cache defines the interface for cache operations
type Cache[K comparable, V any] interface {
	// Get retrieves a value from the cache
	Get(key K) (V, bool)
	// Set adds or updates a value in the cache
	Set(key K, value V)
	// Delete removes a value from the cache
	Delete(key K)
}

// EvictionPolicy defines how entries should be evicted from the cache
type EvictionPolicy interface {
	ShouldEvict(entry *cacheEntry) bool
}

// cacheEntry represents a single entry in the cache
type cacheEntry struct {
	value     interface{}
	createdAt time.Time
}

// inMemoryCache is an in-memory implementation of the Cache interface
type inMemoryCache[K comparable, V any] struct {
	data   map[K]*cacheEntry
	policy EvictionPolicy
	mutex  sync.RWMutex
}

// NewCache creates a new cache with the specified eviction policy
func NewCache[K comparable, V any](policy EvictionPolicy) Cache[K, V] {
	return &inMemoryCache[K, V]{
		data:   make(map[K]*cacheEntry),
		policy: policy,
	}
}

// Set adds or updates a value in the cache
func (c *inMemoryCache[K, V]) Set(key K, value V) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.data[key] = &cacheEntry{
		value:     value,
		createdAt: time.Now(),
	}
}

// Get retrieves a value from the cache
func (c *inMemoryCache[K, V]) Get(key K) (V, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	entry, exists := c.data[key]
	if !exists {
		var zero V
		return zero, false
	}

	if c.policy.ShouldEvict(entry) {
		var zero V
		return zero, false
	}

	return entry.value.(V), true
}

// Delete removes a value from the cache
func (c *inMemoryCache[K, V]) Delete(key K) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	delete(c.data, key)
}

// NeverExpirePolicy implements EvictionPolicy for entries that should never expire
type NeverExpirePolicy struct{}

func (p NeverExpirePolicy) ShouldEvict(_ *cacheEntry) bool {
	return false
}

// TimeBasedPolicy implements EvictionPolicy for time-based expiration
type TimeBasedPolicy struct {
	ttl time.Duration
}

func NewTimeBasedPolicy(ttl time.Duration) *TimeBasedPolicy {
	return &TimeBasedPolicy{ttl: ttl}
}

func (p TimeBasedPolicy) ShouldEvict(entry *cacheEntry) bool {
	return time.Since(entry.createdAt) > p.ttl
}

// CustomPolicy implements EvictionPolicy with a custom eviction function
type CustomPolicy struct {
	evictFunc func(interface{}) bool
}

func NewCustomPolicy(evictFunc func(interface{}) bool) *CustomPolicy {
	return &CustomPolicy{evictFunc: evictFunc}
}

func (p CustomPolicy) ShouldEvict(entry *cacheEntry) bool {
	return p.evictFunc(entry.value)
}
