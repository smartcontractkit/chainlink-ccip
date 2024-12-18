package cache

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInMemoryCache_Basic(t *testing.T) {
	t.Run("never expire policy", func(t *testing.T) {
		cache := NewInMemoryCache[string, int](NewNeverExpirePolicy())

		// Test Set - new entry
		isNew := cache.Set("key1", 100)
		assert.True(t, isNew, "should be a new entry")

		// Test Set - update existing
		isNew = cache.Set("key1", 200)
		assert.False(t, isNew, "should be an update")

		// Test Get - existing key
		value, exists := cache.Get("key1")
		assert.True(t, exists, "should exist")
		assert.Equal(t, 200, value)

		// Test Get - non-existing key
		value, exists = cache.Get("non-existing")
		assert.False(t, exists, "should not exist")
		assert.Equal(t, 0, value)

		// Test Delete - existing key
		deleted := cache.Delete("key1")
		assert.True(t, deleted, "should be deleted")

		// Test Delete - non-existing key
		deleted = cache.Delete("key1")
		assert.False(t, deleted, "should not be deleted")
	})
}

func TestInMemoryCache_TimeBased(t *testing.T) {
	t.Run("time based policy", func(t *testing.T) {
		cache := NewInMemoryCache[string, int](NewTimeBasedPolicy(100 * time.Millisecond))

		cache.Set("key1", 100)

		// Immediate get should succeed
		value, exists := cache.Get("key1")
		assert.True(t, exists)
		assert.Equal(t, 100, value)

		// Wait for TTL to expire
		time.Sleep(150 * time.Millisecond)

		// Get after expiry should fail
		value, exists = cache.Get("key1")
		assert.False(t, exists)
		assert.Equal(t, 0, value)
	})
}

func TestInMemoryCache_Custom(t *testing.T) {
	t.Run("custom policy", func(t *testing.T) {
		// Custom policy that evicts odd numbers
		cache := NewInMemoryCache[string, int](NewCustomPolicy(func(v interface{}) bool {
			if val, ok := v.(int); ok {
				return val%2 != 0
			}
			return false
		}))

		// Even number should stay
		cache.Set("even", 2)
		value, exists := cache.Get("even")
		assert.True(t, exists)
		assert.Equal(t, 2, value)

		// Odd number should be evicted on get
		cache.Set("odd", 3)
		value, exists = cache.Get("odd")
		assert.False(t, exists)
		assert.Equal(t, 0, value)
	})
}

func TestInMemoryCache_Concurrent(t *testing.T) {
	t.Run("concurrent access", func(t *testing.T) {
		cache := NewInMemoryCache[int, int](NewNeverExpirePolicy())
		var wg sync.WaitGroup
		numGoroutines := 10
		numOperations := 100

		// Concurrent writes
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(routine int) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := routine*numOperations + j
					cache.Set(key, key)
				}
			}(i)
		}

		// Concurrent reads
		for i := 0; i < numGoroutines; i++ {
			wg.Add(1)
			go func(routine int) {
				defer wg.Done()
				for j := 0; j < numOperations; j++ {
					key := routine*numOperations + j
					_, _ = cache.Get(key)
				}
			}(i)
		}

		wg.Wait()
	})
}

func TestInMemoryCache_Types(t *testing.T) {
	t.Run("different types", func(t *testing.T) {
		// String cache
		stringCache := NewInMemoryCache[string, string](NewNeverExpirePolicy())
		stringCache.Set("hello", "world")
		val, exists := stringCache.Get("hello")
		assert.True(t, exists)
		assert.Equal(t, "world", val)

		// Struct cache
		type Person struct {
			Name string
			Age  int
		}
		structCache := NewInMemoryCache[string, Person](NewNeverExpirePolicy())
		structCache.Set("person", Person{Name: "John", Age: 30})
		person, exists := structCache.Get("person")
		assert.True(t, exists)
		assert.Equal(t, "John", person.Name)
		assert.Equal(t, 30, person.Age)
	})
}

func TestInMemoryCache_EdgeCases(t *testing.T) {
	t.Run("edge cases", func(t *testing.T) {
		cache := NewInMemoryCache[string, *string](NewNeverExpirePolicy())

		// Nil value
		var nilStr *string
		cache.Set("nil", nilStr)
		val, exists := cache.Get("nil")
		assert.True(t, exists)
		assert.Nil(t, val)

		// Empty string key
		str := "value"
		cache.Set("", &str)
		val, exists = cache.Get("")
		assert.True(t, exists)
		assert.Equal(t, &str, val)

		// Delete empty key
		deleted := cache.Delete("")
		assert.True(t, deleted)
	})
}
