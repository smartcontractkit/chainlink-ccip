package testutils

import "sync"

// ConcurrencyGroup is a counting semaphore.
type ConcurrencyGroup chan struct{}

// Registry manages named concurrency groups.
type Registry struct {
	sync.Mutex
	m map[string]ConcurrencyGroup
}

// NewRegistry creates a new, empty registry.
func newRegistry() *Registry {
	return &Registry{m: make(map[string]ConcurrencyGroup)}
}

// Global registry used by the helper Get.
var defaultRegistry = newRegistry()

// Get returns the semaphore for id, creating it with the given
// concurrency level if it does not yet exist.
func (r *Registry) get(id string, concurrency int) ConcurrencyGroup {
	r.Lock()
	defer r.Unlock()
	if s, ok := r.m[id]; ok {
		return s
	}
	s := make(ConcurrencyGroup, concurrency)
	r.m[id] = s
	return s
}

// GetConcurrencyGroup uses the global registry.
func GetConcurrencyGroup(id string, concurrency int) ConcurrencyGroup {
	return defaultRegistry.get(id, concurrency)
}

// Enter acquires a token (blocks until one is available).
func (s ConcurrencyGroup) Enter() { s <- struct{}{} }

// Leave releases a token.
func (s ConcurrencyGroup) Leave() { <-s }
