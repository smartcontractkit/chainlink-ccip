package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

/*
Package cache provides a generic caching implementation that wraps the go-cache library
with additional support for custom eviction policies. It allows both time-based expiration
(inherited from go-cache) and custom eviction rules through user-defined policies.

The cache is type-safe through Go generics, thread-safe through mutex locks, and supports
all basic cache operations. Keys are strings, and values can be of any type. Each cached
value stores its insertion timestamp, allowing for time-based validation in custom policies.

Example usage with contract reader:
    type Event struct {
        Timestamp int64
        Data      string
    }

    type ContractReader interface {
        QueryEvents(ctx context.Context, filter QueryFilter) ([]Event, error)
    }

    reader := NewContractReader()

    // Create cache with contract reader in closure
    cache := NewCustomCache[Event](
        5*time.Minute,     // Default expiration
        10*time.Minute,    // Cleanup interval
        func(ev Event, _ time.Time) bool {
            ctx := context.Background()
            filter := QueryFilter{
                FromTimestamp: ev.Timestamp(),
                Confidence:   Finalized,
            }

            // Query for any events after our cache insertion time
            newEvents, err := reader.QueryEvents(ctx, filter)
            if err != nil {
                return false // Keep cache on error
            }

            // Evict if new events exist after our cache time
            return len(newEvents) > 0
        },
    )

    // Cache an event
    ev := Event{Timestamp: time.Now().Unix(), Data: "..."}
    cache.Set("key", ev, NoExpiration)

    // Later: event will be evicted if newer ones exist on chain
    ev, found := cache.Get("key")

The cache ensures data freshness through:
  - Automatic time-based expiration from go-cache
  - Custom eviction policies with access to storage timestamps
  - Thread-safe operations for concurrent access
  - Type safety through Go generics
*/

const (
	NoExpiration = cache.NoExpiration
)

// Cache defines the interface for cache operations
type Cache[V any] interface {
	// Set adds an item to the cache with an expiration time
	Set(key string, value V, expiration time.Duration)
	// Get retrieves an item from the cache
	Get(key string) (V, bool)
	// Delete removes an item from the cache
	Delete(key string)
	// Items returns all items in the cache
	Items() map[string]V
}

// timestampedValue wraps a value with its storage timestamp
type timestampedValue[V any] struct {
	Value    V
	StoredAt time.Time
}

type CustomCache[V any] struct {
	*cache.Cache
	customPolicy func(V, time.Time) bool // Updated to include storage time
	mutex        sync.RWMutex
}

// NewCustomCache creates a new cache with both time-based and custom eviction policies
func NewCustomCache[V any](
	defaultExpiration time.Duration,
	cleanupInterval time.Duration,
	customPolicy func(V, time.Time) bool,
) *CustomCache[V] {
	return &CustomCache[V]{
		Cache:        cache.New(defaultExpiration, cleanupInterval),
		customPolicy: customPolicy,
	}
}

// Set adds an item to the cache with current timestamp
func (c *CustomCache[V]) Set(key string, value V, expiration time.Duration) {
	wrapped := timestampedValue[V]{
		Value:    value,
		StoredAt: time.Now(),
	}
	c.Cache.Set(key, wrapped, expiration)
}

// Get retrieves an item from the cache, checking both time-based and custom policies
func (c *CustomCache[V]) Get(key string) (V, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	var zero V
	value, found := c.Cache.Get(key)
	if !found {
		return zero, false
	}

	// Type assertion for timestamped value
	wrapped, ok := value.(timestampedValue[V])
	if !ok {
		return zero, false
	}

	// Check custom policy with timestamp
	if c.customPolicy != nil && c.customPolicy(wrapped.Value, wrapped.StoredAt) {
		c.Cache.Delete(key)
		return zero, false
	}

	return wrapped.Value, true
}

// Delete removes an item from the cache
func (c *CustomCache[V]) Delete(key string) {
	c.Cache.Delete(key)
}

// Items returns all items in the cache
func (c *CustomCache[V]) Items() map[string]V {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	items := c.Cache.Items()
	result := make(map[string]V)

	for k, v := range items {
		if wrapped, ok := v.Object.(timestampedValue[V]); ok {
			result[k] = wrapped.Value
		}
	}

	return result
}
