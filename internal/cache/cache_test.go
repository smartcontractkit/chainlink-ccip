package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomCache(t *testing.T) {
	t.Run("basic operations without custom policy", func(t *testing.T) {
		cache := NewCustomCache[int](5*time.Minute, 10*time.Minute, nil)

		// Test Set and Get
		cache.Set("test1", 100, NoExpiration)
		value, found := cache.Get("test1")
		assert.True(t, found)
		assert.Equal(t, 100, value)

		// Test non-existent key
		_, found = cache.Get("nonexistent")
		assert.False(t, found)

		// Test Delete
		cache.Delete("test1")
		_, found = cache.Get("test1")
		assert.False(t, found)
	})

	t.Run("custom policy with timestamp", func(t *testing.T) {
		now := time.Now()
		isStale := func(v int, storedAt time.Time) bool {
			return storedAt.Before(now)
		}
		cache := NewCustomCache[int](5*time.Minute, 10*time.Minute, isStale)

		// Value stored now should not be evicted
		cache.Set("fresh", 1, NoExpiration)
		value, found := cache.Get("fresh")
		assert.True(t, found)
		assert.Equal(t, 1, value)

		// Simulate old value by manipulating timestamp
		oldValue := timestampedValue[int]{
			Value:    2,
			StoredAt: now.Add(-1 * time.Hour),
		}
		cache.Cache.Set("stale", oldValue, NoExpiration)

		// Stale value should be evicted
		_, found = cache.Get("stale")
		assert.False(t, found)
	})

	t.Run("time based expiration", func(t *testing.T) {
		cache := NewCustomCache[string](1*time.Second, 1*time.Second, nil)

		cache.Set("key", "value", 100*time.Millisecond)

		// Should exist initially
		value, found := cache.Get("key")
		assert.True(t, found)
		assert.Equal(t, "value", value)

		// Should expire
		time.Sleep(200 * time.Millisecond)
		_, found = cache.Get("key")
		assert.False(t, found)
	})

	t.Run("items retrieval", func(t *testing.T) {
		cache := NewCustomCache[int](5*time.Minute, 10*time.Minute, nil)

		cache.Set("one", 1, NoExpiration)
		cache.Set("two", 2, NoExpiration)

		items := cache.Items()
		assert.Len(t, items, 2)
		assert.Equal(t, 1, items["one"])
		assert.Equal(t, 2, items["two"])
	})

	t.Run("concurrent access with timestamps", func(t *testing.T) {
		cache := NewCustomCache[int](5*time.Minute, 10*time.Minute, nil)

		// Run multiple goroutines accessing the cache
		done := make(chan bool)
		for i := 0; i < 10; i++ {
			go func(val int) {
				cache.Set("key", val, NoExpiration)
				_, _ = cache.Get("key")
				done <- true
			}(i)
		}

		// Wait for all goroutines
		for i := 0; i < 10; i++ {
			<-done
		}

		// Should have a value at the end
		_, found := cache.Get("key")
		assert.True(t, found)
	})

	t.Run("complex types with timestamp eviction", func(t *testing.T) {
		type ComplexType struct {
			ID        int
			Name      string
			Timestamp time.Time
		}

		threshold := time.Now()
		cache := NewCustomCache[ComplexType](
			5*time.Minute,
			10*time.Minute,
			func(v ComplexType, storedAt time.Time) bool {
				return storedAt.Before(threshold)
			},
		)

		value := ComplexType{ID: 1, Name: "test", Timestamp: time.Now()}
		cache.Set("complex", value, NoExpiration)

		// Fresh value should not be evicted
		retrieved, found := cache.Get("complex")
		assert.True(t, found)
		assert.Equal(t, value, retrieved)

		// Simulate old value
		oldValue := timestampedValue[ComplexType]{
			Value:    ComplexType{ID: 2, Name: "old"},
			StoredAt: threshold.Add(-1 * time.Hour),
		}
		cache.Cache.Set("old", oldValue, NoExpiration)

		// Old value should be evicted
		_, found = cache.Get("old")
		assert.False(t, found)
	})

	t.Run("correct timestamp storage", func(t *testing.T) {
		cache := NewCustomCache[string](5*time.Minute, 10*time.Minute, nil)

		before := time.Now()
		cache.Set("key", "value", NoExpiration)
		after := time.Now()

		// Get the raw timestamped value
		raw, found := cache.Cache.Get("key")
		assert.True(t, found)

		wrapped, ok := raw.(timestampedValue[string])
		assert.True(t, ok)

		// StoredAt should be between before and after
		assert.True(t, wrapped.StoredAt.After(before) || wrapped.StoredAt.Equal(before))
		assert.True(t, wrapped.StoredAt.Before(after) || wrapped.StoredAt.Equal(after))
	})
}
