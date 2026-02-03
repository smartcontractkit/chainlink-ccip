package main

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// CallManager is the centralized orchestrator for all RPC calls.
// It provides caching, request deduplication, retry logic, and per-RPC rate limiting.
type CallManager struct {
	// Thread-safe cache for call results
	cache   map[[32]byte]CallResult
	cacheMu sync.RWMutex

	// In-flight request deduplication
	pending   map[[32]byte][]chan CallResult
	pendingMu sync.Mutex

	// Chain-specific RPC pools (keyed by chain selector)
	rpcPools   map[uint64]*RPCPool
	rpcPoolsMu sync.RWMutex

	// Configuration
	maxRetries     int
	initialBackoff time.Duration

	// Statistics (atomic for thread-safety)
	cacheHits    atomic.Int64
	dedupedCalls atomic.Int64
	totalRetries atomic.Int64
	totalCalls   atomic.Int64
	totalErrors  atomic.Int64
}

// RPCPool manages multiple RPC endpoints for a single chain.
// Handles load balancing, rate limiting, and failover.
type RPCPool struct {
	chainID   uint64
	endpoints []*RPCEndpointManager
	mu        sync.Mutex
	nextIdx   int // Round-robin index
}

// RPCEndpointManager manages a single RPC endpoint with rate limiting.
type RPCEndpointManager struct {
	url        string
	name       string
	executor   ChainExecutor
	semaphore  chan struct{} // Limits concurrent requests
	maxWorkers int

	// Stats for this endpoint
	totalCalls   atomic.Int64
	totalErrors  atomic.Int64
	rateLimits   atomic.Int64
	avgLatencyNs atomic.Int64
}

// NewCallManager creates a new CallManager.
func NewCallManager() *CallManager {
	return &CallManager{
		cache:          make(map[[32]byte]CallResult),
		pending:        make(map[[32]byte][]chan CallResult),
		rpcPools:       make(map[uint64]*RPCPool),
		maxRetries:     20,
		initialBackoff: 300 * time.Millisecond,
	}
}

// NewCallManagerWithConfig creates a CallManager configured from network config.
func NewCallManagerWithConfig(registry *ChainRegistry) *CallManager {
	mgr := NewCallManager()

	// Register all chains from the registry
	for selector, chain := range registry.chains {
		// Create RPC pool based on chain family
		pool := &RPCPool{
			chainID:   selector,
			endpoints: make([]*RPCEndpointManager, 0),
		}

		for _, rpc := range chain.RPCs {
			if rpc.HTTPURL == "" {
				continue
			}
			// Skip non-HTTP URLs (like liteserver://)
			if len(rpc.HTTPURL) < 4 || rpc.HTTPURL[:4] != "http" {
				continue
			}

			var executor ChainExecutor
			switch chain.Family {
			case ChainFamilyEVM:
				executor = NewEVMExecutor(rpc.HTTPURL)
			case ChainFamilySVM:
				executor = NewSolanaExecutor(rpc.HTTPURL)
			default:
				// Skip unsupported chain families for now
				continue
			}

			endpoint := &RPCEndpointManager{
				url:        rpc.HTTPURL,
				name:       rpc.RPCName,
				executor:   executor,
				maxWorkers: 8, // Default max concurrent requests per RPC
				semaphore:  make(chan struct{}, 8),
			}
			pool.endpoints = append(pool.endpoints, endpoint)
		}

		if len(pool.endpoints) > 0 {
			mgr.rpcPools[selector] = pool
		}
		// TODO: Add Aptos, TON executors
	}

	return mgr
}

// SetRPCLimit sets the max concurrent workers for all RPCs of a chain.
func (m *CallManager) SetRPCLimit(chainID uint64, maxWorkers int) {
	m.rpcPoolsMu.RLock()
	pool, ok := m.rpcPools[chainID]
	m.rpcPoolsMu.RUnlock()

	if !ok {
		return
	}

	pool.mu.Lock()
	defer pool.mu.Unlock()

	for _, endpoint := range pool.endpoints {
		// Replace semaphore with new size
		endpoint.maxWorkers = maxWorkers
		endpoint.semaphore = make(chan struct{}, maxWorkers)
	}
}

// RegisterRPCPool registers an RPC pool for a chain.
func (m *CallManager) RegisterRPCPool(chainID uint64, pool *RPCPool) {
	m.rpcPoolsMu.Lock()
	defer m.rpcPoolsMu.Unlock()
	m.rpcPools[chainID] = pool
}

// RegisterExecutor registers a single executor for a chain (legacy compatibility).
func (m *CallManager) RegisterExecutor(chainID uint64, executor ChainExecutor) {
	pool := &RPCPool{
		chainID: chainID,
		endpoints: []*RPCEndpointManager{
			{
				url:        "legacy",
				name:       "legacy",
				executor:   executor,
				maxWorkers: 8,
				semaphore:  make(chan struct{}, 8),
			},
		},
	}
	m.RegisterRPCPool(chainID, pool)
}

// Execute processes a call, using cache/dedup where possible.
// This is the main entry point for making calls through the orchestrator.
func (m *CallManager) Execute(call Call) CallResult {
	return m.ExecuteWithContext(context.Background(), call)
}

// ExecuteWithContext processes a call with context for cancellation.
func (m *CallManager) ExecuteWithContext(ctx context.Context, call Call) CallResult {
	m.totalCalls.Add(1)
	key := call.CacheKey()

	// 1. Check cache first (read lock for concurrent reads)
	m.cacheMu.RLock()
	if result, ok := m.cache[key]; ok {
		m.cacheMu.RUnlock()
		m.cacheHits.Add(1)
		return CallResult{
			Data:    result.Data,
			Error:   result.Error,
			Cached:  true,
			Retries: result.Retries,
		}
	}
	m.cacheMu.RUnlock()

	// 2. Check if this call is already in-flight (deduplication)
	m.pendingMu.Lock()
	if waiters, ok := m.pending[key]; ok {
		ch := make(chan CallResult, 1)
		m.pending[key] = append(waiters, ch)
		m.pendingMu.Unlock()
		m.dedupedCalls.Add(1)

		// Wait for result or context cancellation
		select {
		case result := <-ch:
			return result
		case <-ctx.Done():
			return CallResult{Error: ctx.Err()}
		}
	}

	// 3. New request - register as pending
	ch := make(chan CallResult, 1)
	m.pending[key] = []chan CallResult{ch}
	m.pendingMu.Unlock()

	// 4. Execute the call (spawns goroutine on-demand)
	go m.executeAndNotify(ctx, call, key)

	// 5. Wait for result
	select {
	case result := <-ch:
		return result
	case <-ctx.Done():
		return CallResult{Error: ctx.Err()}
	}
}

// executeAndNotify executes a call and notifies all waiters.
func (m *CallManager) executeAndNotify(ctx context.Context, call Call, key [32]byte) {
	result := m.executeWithRetry(ctx, call)

	// Store in cache (even errors are cached to prevent retry storms)
	m.cacheMu.Lock()
	m.cache[key] = result
	m.cacheMu.Unlock()

	// Notify all waiters
	m.pendingMu.Lock()
	waiters := m.pending[key]
	delete(m.pending, key)
	m.pendingMu.Unlock()

	for _, ch := range waiters {
		select {
		case ch <- result:
		default:
			// Channel full or closed, skip
		}
	}
}

// executeWithRetry attempts to execute a call with exponential backoff retry.
func (m *CallManager) executeWithRetry(ctx context.Context, call Call) CallResult {
	m.rpcPoolsMu.RLock()
	pool, ok := m.rpcPools[call.ChainID]
	m.rpcPoolsMu.RUnlock()

	if !ok {
		m.totalErrors.Add(1)
		return CallResult{
			Error: fmt.Errorf("no RPC pool registered for chain %d", call.ChainID),
		}
	}

	backoff := m.initialBackoff
	var lastErr error

	for attempt := 0; attempt <= m.maxRetries; attempt++ {
		// Check context cancellation
		if ctx.Err() != nil {
			return CallResult{Error: ctx.Err(), Retries: attempt}
		}

		// Get an endpoint with available capacity (smart round-robin)
		endpoint := pool.getAvailableEndpoint(ctx)
		if endpoint == nil {
			m.totalErrors.Add(1)
			return CallResult{Error: fmt.Errorf("no available endpoints for chain %d", call.ChainID)}
		}

		// Slot already acquired by getAvailableEndpoint

		// Execute the call
		start := time.Now()
		result, err := endpoint.executor.Execute(call.Target, call.Data)
		elapsed := time.Since(start)

		// Release semaphore
		<-endpoint.semaphore

		// Update endpoint stats
		endpoint.totalCalls.Add(1)
		endpoint.avgLatencyNs.Store(elapsed.Nanoseconds())

		if err == nil {
			return CallResult{Data: result, Retries: attempt}
		}

		lastErr = err
		endpoint.totalErrors.Add(1)

		// Check if the error is retryable
		if !isRetryable(err) {
			m.totalErrors.Add(1)
			return CallResult{Error: err, Retries: attempt}
		}

		// Track rate limits
		if isRateLimitError(err) {
			endpoint.rateLimits.Add(1)
		}

		// Don't retry on the last attempt
		if attempt < m.maxRetries {
			m.totalRetries.Add(1)

			select {
			case <-time.After(backoff):
			case <-ctx.Done():
				return CallResult{Error: ctx.Err(), Retries: attempt}
			}
			backoff *= 2 // Exponential backoff
		}
	}

	m.totalErrors.Add(1)
	return CallResult{
		Error:   fmt.Errorf("max retries (%d) exceeded: %w", m.maxRetries, lastErr),
		Retries: m.maxRetries,
	}
}

// getEndpoint returns the next endpoint using simple round-robin (no capacity check).
func (p *RPCPool) getEndpoint() *RPCEndpointManager {
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.endpoints) == 0 {
		return nil
	}

	endpoint := p.endpoints[p.nextIdx]
	p.nextIdx = (p.nextIdx + 1) % len(p.endpoints)
	return endpoint
}

// getAvailableEndpoint returns an endpoint with available capacity.
// Prioritizes endpoints by success rate and skips those at max capacity.
// Acquires the semaphore slot before returning.
func (p *RPCPool) getAvailableEndpoint(ctx context.Context) *RPCEndpointManager {
	p.mu.Lock()
	numEndpoints := len(p.endpoints)
	if numEndpoints == 0 {
		p.mu.Unlock()
		return nil
	}
	endpoints := make([]*RPCEndpointManager, numEndpoints)
	copy(endpoints, p.endpoints)
	p.mu.Unlock()

	sort.Slice(endpoints, func(i, j int) bool {
		scoreI := endpointScore(endpoints[i])
		scoreJ := endpointScore(endpoints[j])
		if scoreI == scoreJ {
			return endpoints[i].totalCalls.Load() < endpoints[j].totalCalls.Load()
		}
		return scoreI > scoreJ
	})

	// Try each endpoint in priority order
	for _, endpoint := range endpoints {
		// Try to acquire a slot (non-blocking)
		select {
		case endpoint.semaphore <- struct{}{}:
			// Got a slot!
			return endpoint
		default:
			// This endpoint is at capacity, try next
			continue
		}
	}

	// All endpoints at capacity - wait on the highest-ranked endpoint
	endpoint := endpoints[0]

	select {
	case endpoint.semaphore <- struct{}{}:
		return endpoint
	case <-ctx.Done():
		return nil
	}
}

// endpointScore returns a smoothed success rate score for an endpoint.
// Uses Laplace smoothing to avoid extreme scores for low-volume endpoints.
func endpointScore(endpoint *RPCEndpointManager) float64 {
	calls := endpoint.totalCalls.Load()
	errors := endpoint.totalErrors.Load()
	successes := calls - errors
	return float64(successes+1) / float64(calls+2)
}

// isRateLimitError checks if an error is specifically a rate limit error
func isRateLimitError(err error) bool {
	if err == nil {
		return false
	}
	errStr := err.Error()
	return contains(errStr, "rate limit") ||
		contains(errStr, "429") ||
		contains(errStr, "too many requests")
}

// isRetryable determines if an error should trigger a retry.
func isRetryable(err error) bool {
	if err == nil {
		return false
	}

	errStr := err.Error()

	// Rate limit errors
	if isRateLimitError(err) {
		return true
	}

	// Timeout errors
	if contains(errStr, "timeout") ||
		contains(errStr, "deadline exceeded") {
		return true
	}

	// Connection errors
	if contains(errStr, "connection refused") ||
		contains(errStr, "connection reset") ||
		contains(errStr, "EOF") ||
		contains(errStr, "broken pipe") {
		return true
	}

	// DNS errors - try next endpoint
	if contains(errStr, "no such host") ||
		contains(errStr, "lookup") {
		return true
	}

	// TLS/Certificate errors - try next endpoint
	if contains(errStr, "tls:") ||
		contains(errStr, "x509:") ||
		contains(errStr, "certificate") {
		return true
	}

	// Server errors (5xx) and connection failures
	if contains(errStr, "500") ||
		contains(errStr, "502") ||
		contains(errStr, "503") ||
		contains(errStr, "504") ||
		contains(errStr, "failed to connect") {
		return true
	}

	// Auth errors (can be transient)
	if contains(errStr, "http error 401") {
		return true
	}

	// RPC overwhelm errors (these are temporary, not actual "unsupported" methods)
	// HTTP 400 with "Unsupported RPC call" or HTTP 404 are often overwhelmed RPCs
	if contains(errStr, "Unsupported RPC call") ||
		contains(errStr, "http error 400") ||
		contains(errStr, "http error 404") {
		return true
	}

	// RPC node sync issues
	if contains(errStr, "upstream does not have the requested block") ||
		contains(errStr, "-32601") { // method not available (temporary RPC issue)
		return true
	}

	return false
}

// contains is a simple substring check
func contains(s, substr string) bool {
	if len(substr) == 0 {
		return true
	}
	if len(s) < len(substr) {
		return false
	}
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// ExecuteAll processes multiple calls concurrently and returns all results.
func (m *CallManager) ExecuteAll(calls []Call) []CallResult {
	return m.ExecuteAllWithContext(context.Background(), calls)
}

// ExecuteAllWithContext processes multiple calls with context.
func (m *CallManager) ExecuteAllWithContext(ctx context.Context, calls []Call) []CallResult {
	results := make([]CallResult, len(calls))
	var wg sync.WaitGroup

	for i, call := range calls {
		wg.Add(1)
		go func(idx int, c Call) {
			defer wg.Done()
			results[idx] = m.ExecuteWithContext(ctx, c)
		}(i, call)
	}

	wg.Wait()
	return results
}

// Stats returns current statistics about the CallManager.
func (m *CallManager) Stats() Stats {
	return Stats{
		TotalCalls:   m.totalCalls.Load(),
		CacheHits:    m.cacheHits.Load(),
		DedupedCalls: m.dedupedCalls.Load(),
		TotalRetries: m.totalRetries.Load(),
		Errors:       m.totalErrors.Load(),
	}
}

// RPCStats returns per-RPC statistics for a chain.
func (m *CallManager) RPCStats(chainID uint64) []RPCEndpointStats {
	m.rpcPoolsMu.RLock()
	pool, ok := m.rpcPools[chainID]
	m.rpcPoolsMu.RUnlock()

	if !ok {
		return nil
	}

	stats := make([]RPCEndpointStats, len(pool.endpoints))
	for i, ep := range pool.endpoints {
		stats[i] = RPCEndpointStats{
			Name:       ep.name,
			URL:        ep.url,
			TotalCalls: ep.totalCalls.Load(),
			Errors:     ep.totalErrors.Load(),
			RateLimits: ep.rateLimits.Load(),
			AvgLatency: time.Duration(ep.avgLatencyNs.Load()),
		}
	}
	return stats
}

// RPCEndpointStats holds stats for a single RPC endpoint.
type RPCEndpointStats struct {
	Name        string
	URL         string
	TotalCalls  int64
	Errors      int64
	RateLimits  int64
	AvgLatency  time.Duration
	ActiveCalls int // Current calls in-flight
	MaxWorkers  int // Max concurrent calls allowed
}

// LiveStats holds real-time statistics about the CallManager.
type LiveStats struct {
	TotalCalls    int64
	CacheHits     int64
	DedupedCalls  int64
	TotalRetries  int64
	Errors        int64
	PendingCalls  int // Calls waiting for results (deduplicated)
	CacheSize     int
	TotalInFlight int // Total calls currently executing across all RPCs
	TotalMaxCap   int // Total max capacity across all RPCs
}

// LiveStats returns real-time statistics including in-flight counts.
func (m *CallManager) LiveStats() LiveStats {
	m.pendingMu.Lock()
	pendingCount := len(m.pending)
	m.pendingMu.Unlock()

	m.cacheMu.RLock()
	cacheSize := len(m.cache)
	m.cacheMu.RUnlock()

	// Count total in-flight across all chains
	totalInFlight := 0
	totalMaxCap := 0
	m.rpcPoolsMu.RLock()
	for _, pool := range m.rpcPools {
		for _, ep := range pool.endpoints {
			totalInFlight += len(ep.semaphore)
			totalMaxCap += ep.maxWorkers
		}
	}
	m.rpcPoolsMu.RUnlock()

	return LiveStats{
		TotalCalls:    m.totalCalls.Load(),
		CacheHits:     m.cacheHits.Load(),
		DedupedCalls:  m.dedupedCalls.Load(),
		TotalRetries:  m.totalRetries.Load(),
		Errors:        m.totalErrors.Load(),
		PendingCalls:  pendingCount,
		CacheSize:     cacheSize,
		TotalInFlight: totalInFlight,
		TotalMaxCap:   totalMaxCap,
	}
}

// LiveRPCStats returns per-RPC statistics including current usage.
func (m *CallManager) LiveRPCStats(chainID uint64) []RPCEndpointStats {
	m.rpcPoolsMu.RLock()
	pool, ok := m.rpcPools[chainID]
	m.rpcPoolsMu.RUnlock()

	if !ok {
		return nil
	}

	stats := make([]RPCEndpointStats, len(pool.endpoints))
	for i, ep := range pool.endpoints {
		stats[i] = RPCEndpointStats{
			Name:        ep.name,
			URL:         ep.url,
			TotalCalls:  ep.totalCalls.Load(),
			Errors:      ep.totalErrors.Load(),
			RateLimits:  ep.rateLimits.Load(),
			AvgLatency:  time.Duration(ep.avgLatencyNs.Load()),
			ActiveCalls: len(ep.semaphore),
			MaxWorkers:  ep.maxWorkers,
		}
	}
	return stats
}

// AllChainIDs returns all registered chain IDs.
func (m *CallManager) AllChainIDs() []uint64 {
	m.rpcPoolsMu.RLock()
	defer m.rpcPoolsMu.RUnlock()

	ids := make([]uint64, 0, len(m.rpcPools))
	for id := range m.rpcPools {
		ids = append(ids, id)
	}
	return ids
}

// ClearCache clears the result cache.
func (m *CallManager) ClearCache() {
	m.cacheMu.Lock()
	defer m.cacheMu.Unlock()
	m.cache = make(map[[32]byte]CallResult)
}

// CacheSize returns the number of entries in the cache.
func (m *CallManager) CacheSize() int {
	m.cacheMu.RLock()
	defer m.cacheMu.RUnlock()
	return len(m.cache)
}

// HasChain returns whether the manager has an RPC pool for the given chain.
func (m *CallManager) HasChain(chainID uint64) bool {
	m.rpcPoolsMu.RLock()
	defer m.rpcPoolsMu.RUnlock()
	_, ok := m.rpcPools[chainID]
	return ok
}

// Close shuts down the CallManager.
func (m *CallManager) Close() {
	// Nothing to clean up with on-demand goroutines
}
