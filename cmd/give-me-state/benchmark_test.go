package main

import (
	"encoding/hex"
	"sync"
	"testing"
	"time"
)

// Test data
var (
	testFeeQuoterAddr, _ = hex.DecodeString("aE30bEA16FBA1Acc076138D6C63d61437b25BdB8")
	testRouterAddr, _    = hex.DecodeString("c0f457e615348708FaAB3B40ECC26Badb32B7b30")
	testOnRampAddr, _    = hex.DecodeString("d025EE7070a3EEc33A95bea409b3108d8199bFC7")

	// Function selectors
	testGetStaticConfig, _ = hex.DecodeString("06285c69")
	testOwner, _           = hex.DecodeString("8da5cb5b")
	testTypeAndVersion, _  = hex.DecodeString("181f5a77")
)

// buildTestCallList creates a list of calls for testing
func buildTestCallList(totalCalls int) []Call {
	templates := []Call{
		{ChainID: 1, Target: testFeeQuoterAddr, Data: testGetStaticConfig},
		{ChainID: 1, Target: testFeeQuoterAddr, Data: testOwner},
		{ChainID: 1, Target: testFeeQuoterAddr, Data: testTypeAndVersion},
		{ChainID: 1, Target: testRouterAddr, Data: testOwner},
		{ChainID: 1, Target: testRouterAddr, Data: testTypeAndVersion},
		{ChainID: 1, Target: testOnRampAddr, Data: testGetStaticConfig},
		{ChainID: 1, Target: testOnRampAddr, Data: testOwner},
	}

	calls := make([]Call, 0, totalCalls)
	for len(calls) < totalCalls {
		idx := len(calls) % len(templates)
		calls = append(calls, templates[idx])
	}
	return calls
}

// countTestUniqueCalls returns the number of unique calls
func countTestUniqueCalls(calls []Call) int {
	seen := make(map[[32]byte]bool)
	for _, c := range calls {
		seen[c.CacheKey()] = true
	}
	return len(seen)
}

// BenchmarkSequentialCalls measures baseline sequential execution
func BenchmarkSequentialCalls(b *testing.B) {
	executor := NewMockExecutor(10 * time.Millisecond)
	calls := buildTestCallList(50)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, call := range calls {
			executor.Execute(call.Target, call.Data)
		}
	}
}

// BenchmarkOrchestratedCalls measures orchestrated parallel execution
func BenchmarkOrchestratedCalls(b *testing.B) {
	executor := NewMockExecutor(10 * time.Millisecond)
	calls := buildTestCallList(50)

	mgr := NewCallManager()
	mgr.RegisterExecutor(1, executor)
	defer mgr.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Clear cache between iterations for fair comparison
		mgr.ClearCache()

		var wg sync.WaitGroup
		for _, call := range calls {
			wg.Add(1)
			go func(c Call) {
				defer wg.Done()
				mgr.Execute(c)
			}(call)
		}
		wg.Wait()
	}
}

// BenchmarkOrchestratedWithCache measures performance with warm cache
func BenchmarkOrchestratedWithCache(b *testing.B) {
	executor := NewMockExecutor(10 * time.Millisecond)
	calls := buildTestCallList(50)

	mgr := NewCallManager()
	mgr.RegisterExecutor(1, executor)
	defer mgr.Close()

	// Warm up cache
	var wg sync.WaitGroup
	for _, call := range calls {
		wg.Add(1)
		go func(c Call) {
			defer wg.Done()
			mgr.Execute(c)
		}(call)
	}
	wg.Wait()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var wg sync.WaitGroup
		for _, call := range calls {
			wg.Add(1)
			go func(c Call) {
				defer wg.Done()
				mgr.Execute(c)
			}(call)
		}
		wg.Wait()
	}
}

// TestCacheEffectiveness verifies that caching reduces actual RPC calls
func TestCacheEffectiveness(t *testing.T) {
	executor := NewMockExecutor(1 * time.Millisecond)
	calls := buildTestCallList(50)
	uniqueCalls := countTestUniqueCalls(calls)

	mgr := NewCallManager()
	mgr.RegisterExecutor(1, executor)
	defer mgr.Close()

	// Execute all calls
	var wg sync.WaitGroup
	for _, call := range calls {
		wg.Add(1)
		go func(c Call) {
			defer wg.Done()
			mgr.Execute(c)
		}(call)
	}
	wg.Wait()

	stats := mgr.Stats()

	// Verify cache is working: actual RPC calls should be <= unique calls
	actualRPCCalls := executor.CallCount.Load()
	if actualRPCCalls > int64(uniqueCalls) {
		t.Errorf("Expected at most %d RPC calls, got %d", uniqueCalls, actualRPCCalls)
	}

	t.Logf("Total calls: %d, Unique: %d, Actual RPC: %d, Cache hits: %d, Deduped: %d",
		len(calls), uniqueCalls, actualRPCCalls, stats.CacheHits, stats.DedupedCalls)
}

// TestDeduplication verifies that concurrent duplicate requests are deduplicated
func TestDeduplication(t *testing.T) {
	// Use slow executor to ensure requests overlap
	executor := NewMockExecutor(50 * time.Millisecond)

	// All calls are the same
	call := Call{ChainID: 1, Target: testFeeQuoterAddr, Data: testGetStaticConfig}
	numConcurrent := 10

	mgr := NewCallManager()
	mgr.RegisterExecutor(1, executor)
	defer mgr.Close()

	// Launch all calls simultaneously
	var wg sync.WaitGroup
	for i := 0; i < numConcurrent; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mgr.Execute(call)
		}()
	}
	wg.Wait()

	stats := mgr.Stats()

	// Only 1 actual RPC call should be made (rest should be deduped or cached)
	actualRPCCalls := executor.CallCount.Load()
	if actualRPCCalls > 1 {
		t.Errorf("Expected 1 RPC call, got %d (deduplication not working)", actualRPCCalls)
	}

	t.Logf("Concurrent calls: %d, Actual RPC: %d, Deduped: %d, Cache hits: %d",
		numConcurrent, actualRPCCalls, stats.DedupedCalls, stats.CacheHits)
}

// TestRetryLogic verifies that transient failures are retried
func TestRetryLogic(t *testing.T) {
	// Executor that fails 20% of calls
	executor := NewFlakyMockExecutor(1*time.Millisecond, 0.2)

	mgr := NewCallManager()
	mgr.RegisterExecutor(1, executor)
	defer mgr.Close()

	// Make multiple unique calls
	calls := []Call{
		{ChainID: 1, Target: testFeeQuoterAddr, Data: testGetStaticConfig},
		{ChainID: 1, Target: testRouterAddr, Data: testOwner},
		{ChainID: 1, Target: testOnRampAddr, Data: testTypeAndVersion},
	}

	var wg sync.WaitGroup
	results := make([]CallResult, len(calls))
	for i, call := range calls {
		wg.Add(1)
		go func(idx int, c Call) {
			defer wg.Done()
			results[idx] = mgr.Execute(c)
		}(i, call)
	}
	wg.Wait()

	stats := mgr.Stats()

	// Count successes
	successes := 0
	for _, r := range results {
		if r.Error == nil {
			successes++
		}
	}

	t.Logf("Calls: %d, Successes: %d, Retries: %d", len(calls), successes, stats.TotalRetries)

	// With retry logic, we should have more successes than without
	if successes < len(calls)/2 {
		t.Errorf("Expected most calls to succeed with retry, got %d/%d", successes, len(calls))
	}
}

// TestErrorIsolation verifies that errors don't cascade
func TestErrorIsolation(t *testing.T) {
	executor := NewMockExecutor(1 * time.Millisecond)

	mgr := NewCallManager()
	// Don't register executor for chain 999 - will cause errors
	mgr.RegisterExecutor(1, executor)
	defer mgr.Close()

	calls := []Call{
		{ChainID: 1, Target: testFeeQuoterAddr, Data: testGetStaticConfig},   // Will succeed
		{ChainID: 999, Target: testFeeQuoterAddr, Data: testGetStaticConfig}, // Will fail (no executor)
		{ChainID: 1, Target: testRouterAddr, Data: testOwner},                // Will succeed
	}

	var wg sync.WaitGroup
	results := make([]CallResult, len(calls))
	for i, call := range calls {
		wg.Add(1)
		go func(idx int, c Call) {
			defer wg.Done()
			results[idx] = mgr.Execute(c)
		}(i, call)
	}
	wg.Wait()

	// Verify that chain 1 calls succeeded despite chain 999 failing
	if results[0].Error != nil {
		t.Errorf("Call 0 should have succeeded: %v", results[0].Error)
	}
	if results[1].Error == nil {
		t.Error("Call 1 should have failed (no executor for chain 999)")
	}
	if results[2].Error != nil {
		t.Errorf("Call 2 should have succeeded: %v", results[2].Error)
	}

	t.Log("Error isolation verified: failures don't cascade to other calls")
}

// TestCacheKey verifies that cache keys are unique for different calls
func TestCacheKey(t *testing.T) {
	calls := []Call{
		{ChainID: 1, Target: testFeeQuoterAddr, Data: testGetStaticConfig},
		{ChainID: 1, Target: testFeeQuoterAddr, Data: testOwner},           // Different data
		{ChainID: 1, Target: testRouterAddr, Data: testGetStaticConfig},    // Different target
		{ChainID: 2, Target: testFeeQuoterAddr, Data: testGetStaticConfig}, // Different chain
	}

	keys := make(map[[32]byte]int)
	for i, call := range calls {
		key := call.CacheKey()
		if existing, ok := keys[key]; ok {
			t.Errorf("Cache key collision between call %d and %d", existing, i)
		}
		keys[key] = i
	}

	// Same call should produce same key
	key1 := calls[0].CacheKey()
	key2 := calls[0].CacheKey()
	if key1 != key2 {
		t.Error("Same call produced different cache keys")
	}
}

// TestSpeedup measures the actual speedup achieved
func TestSpeedup(t *testing.T) {
	latency := 10 * time.Millisecond
	executor := NewMockExecutor(latency)
	calls := buildTestCallList(50)

	// Sequential timing
	start := time.Now()
	for _, call := range calls {
		executor.Execute(call.Target, call.Data)
	}
	sequential := time.Since(start)

	// Orchestrated timing (fresh executor for fair comparison)
	executor = NewMockExecutor(latency)
	mgr := NewCallManager()
	mgr.RegisterExecutor(1, executor)
	defer mgr.Close()

	start = time.Now()
	var wg sync.WaitGroup
	for _, call := range calls {
		wg.Add(1)
		go func(c Call) {
			defer wg.Done()
			mgr.Execute(c)
		}(call)
	}
	wg.Wait()
	orchestrated := time.Since(start)

	speedup := sequential.Seconds() / orchestrated.Seconds()

	t.Logf("Sequential: %v, Orchestrated: %v, Speedup: %.1fx", sequential, orchestrated, speedup)

	// We should see significant speedup
	if speedup < 5 {
		t.Errorf("Expected at least 5x speedup, got %.1fx", speedup)
	}
}

// TestConfigLoading tests loading configuration files
func TestConfigLoading(t *testing.T) {
	// Test loading address refs if file exists
	refs, err := LoadAddressRefs("address_refs.json")
	if err != nil {
		t.Skipf("Skipping address refs test (file not found): %v", err)
	}
	if len(refs) == 0 {
		t.Error("Expected at least one address ref")
	}
	t.Logf("Loaded %d address refs", len(refs))

	// Test loading network config if file exists
	config, err := LoadNetworkConfig("testnet.yaml")
	if err != nil {
		t.Skipf("Skipping network config test (file not found): %v", err)
	}
	if len(config.Networks) == 0 {
		t.Error("Expected at least one network")
	}
	t.Logf("Loaded %d networks", len(config.Networks))
}
