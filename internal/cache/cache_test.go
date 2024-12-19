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

	t.Run("custom policy eviction", func(t *testing.T) {
		isEven := func(v int) bool {
			return v%2 == 0
		}
		cache := NewCustomCache[int](5*time.Minute, 10*time.Minute, isEven)

		// Even number should be evicted
		cache.Set("even", 2, NoExpiration)
		_, found := cache.Get("even")
		assert.False(t, found)

		// Odd number should remain
		cache.Set("odd", 3, NoExpiration)
		value, found := cache.Get("odd")
		assert.True(t, found)
		assert.Equal(t, 3, value)
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

	t.Run("flush operation", func(t *testing.T) {
		cache := NewCustomCache[int](5*time.Minute, 10*time.Minute, nil)

		cache.Set("one", 1, NoExpiration)
		cache.Set("two", 2, NoExpiration)

		cache.Flush()
		items := cache.Items()
		assert.Len(t, items, 0)
	})

	t.Run("type safety", func(t *testing.T) {
		cache := NewCustomCache[int](5*time.Minute, 10*time.Minute, nil)

		// Set with correct type
		cache.Set("good", 123, NoExpiration)

		// Simulate wrong type in underlying cache
		cache.cache.Set("bad", "not an int", NoExpiration)

		// Good type should work
		value, found := cache.Get("good")
		assert.True(t, found)
		assert.Equal(t, 123, value)

		// Bad type should fail safely
		_, found = cache.Get("bad")
		assert.False(t, found)
	})

	t.Run("concurrent access", func(t *testing.T) {
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

	t.Run("complex types", func(t *testing.T) {
		type ComplexType struct {
			ID   int
			Name string
		}

		cache := NewCustomCache[ComplexType](5*time.Minute, 10*time.Minute, nil)

		value := ComplexType{ID: 1, Name: "test"}
		cache.Set("complex", value, NoExpiration)

		retrieved, found := cache.Get("complex")
		assert.True(t, found)
		assert.Equal(t, value, retrieved)
	})

	t.Run("custom policy with nil value", func(t *testing.T) {
		isNil := func(v *string) bool {
			return v == nil
		}
		cache := NewCustomCache[*string](5*time.Minute, 10*time.Minute, isNil)

		str := "test"
		cache.Set("nonnil", &str, NoExpiration)
		cache.Set("nil", nil, NoExpiration)

		// Non-nil value should remain
		value, found := cache.Get("nonnil")
		assert.True(t, found)
		assert.Equal(t, &str, value)

		// Nil value should be evicted
		_, found = cache.Get("nil")
		assert.False(t, found)
	})
}
