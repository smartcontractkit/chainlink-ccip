package orchestrator

import (
	"crypto/sha256"
	"sync"
)

// CacheDedup provides cache and in-flight request deduplication for typed orchestrators.
// Use one per chain (e.g. EVM: one per chain selector) so keys never cross chains.
type CacheDedup struct {
	cacheMu   sync.RWMutex
	cache     map[[32]byte]CallResult
	pendingMu sync.Mutex
	pending   map[[32]byte][]chan CallResult
}

// NewCacheDedup creates a new cache+dedup for a single chain (or scope).
func NewCacheDedup() *CacheDedup {
	return &CacheDedup{
		cache:   make(map[[32]byte]CallResult),
		pending: make(map[[32]byte][]chan CallResult),
	}
}

// GetOrRun returns the cached result for key, or runs fn(), caches and deduplicates in-flight.
// If another goroutine is already running the same key, this blocks until it completes and returns that result.
func (c *CacheDedup) GetOrRun(key [32]byte, fn func() CallResult) CallResult {
	// 1. Check cache
	c.cacheMu.RLock()
	if res, ok := c.cache[key]; ok {
		c.cacheMu.RUnlock()
		res.Cached = true
		return res
	}
	c.cacheMu.RUnlock()

	// 2. Check in-flight (dedup)
	c.pendingMu.Lock()
	if waiters, ok := c.pending[key]; ok {
		ch := make(chan CallResult, 1)
		c.pending[key] = append(waiters, ch)
		c.pendingMu.Unlock()
		return <-ch
	}
	ch := make(chan CallResult, 1)
	c.pending[key] = []chan CallResult{ch}
	c.pendingMu.Unlock()

	// 3. Run and store result
	result := fn()
	c.cacheMu.Lock()
	c.cache[key] = result
	c.cacheMu.Unlock()

	// 4. Notify waiters
	c.pendingMu.Lock()
	waiters := c.pending[key]
	delete(c.pending, key)
	c.pendingMu.Unlock()
	for _, w := range waiters {
		select {
		case w <- result:
		default:
		}
	}
	return result
}

// KeyFromTargetAndData builds a 32-byte cache key from target and data (e.g. for EVM per-chain cache).
func KeyFromTargetAndData(target, data []byte) [32]byte {
	h := sha256.New()
	h.Write(target)
	h.Write(data)
	var key [32]byte
	copy(key[:], h.Sum(nil))
	return key
}
