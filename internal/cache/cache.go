package cache

import (
	"sync"
	"time"
)

// Cache defines the interface for cache operations
type Cache[K comparable, V any] interface {
	// Get retrieves a value from the cache
	// Returns the value and true if found, zero value and false if not found
	Get(key K) (V, bool)

	// Set adds or updates a value in the cache
	// Returns true if a new entry was created, false if an existing entry was updated
	Set(key K, value V) bool

	// Delete removes a value from the cache
	// Returns true if an entry was deleted, false if the key wasn't found
	Delete(key K) bool
}

// EvictionPolicy defines how entries should be evicted from the cache
type EvictionPolicy interface {
	// TODO: async process needed
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

// NewInMemoryCache creates a new cache with the specified eviction policy
// The cache is thread-safe and can be used concurrently
func NewInMemoryCache[K comparable, V any](policy EvictionPolicy) Cache[K, V] {
	return &inMemoryCache[K, V]{
		data:   make(map[K]*cacheEntry),
		policy: policy,
	}
}

// Set adds or updates a value in the cache
func (c *inMemoryCache[K, V]) Set(key K, value V) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	_, exists := c.data[key]
	c.data[key] = &cacheEntry{
		value:     value,
		createdAt: time.Now(),
	}
	return !exists
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

	value, ok := entry.value.(V)
	if !ok {
		var zero V
		return zero, false
	}

	return value, true
}

// Delete removes a value from the cache
func (c *inMemoryCache[K, V]) Delete(key K) bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	_, exists := c.data[key]
	if exists {
		delete(c.data, key)
	}
	return exists
}

// neverExpirePolicy implements EvictionPolicy for entries that should never expire
type neverExpirePolicy struct{}

func NewNeverExpirePolicy() EvictionPolicy {
	return neverExpirePolicy{}
}

func (p neverExpirePolicy) ShouldEvict(_ *cacheEntry) bool {
	return false
}

// timeBasedPolicy implements EvictionPolicy for time-based expiration
type timeBasedPolicy struct {
	ttl time.Duration
}

func NewTimeBasedPolicy(ttl time.Duration) EvictionPolicy {
	return &timeBasedPolicy{ttl: ttl}
}

func (p timeBasedPolicy) ShouldEvict(entry *cacheEntry) bool {
	return time.Since(entry.createdAt) > p.ttl
}

// customPolicy implements EvictionPolicy with a custom eviction function
type customPolicy struct {
	evictFunc func(interface{}) bool
}

func NewCustomPolicy(evictFunc func(interface{}) bool) EvictionPolicy {
	return &customPolicy{evictFunc: evictFunc}
}

func (p customPolicy) ShouldEvict(entry *cacheEntry) bool {
	return p.evictFunc(entry.value)
}
