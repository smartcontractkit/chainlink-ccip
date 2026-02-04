package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"call-orchestrator-demo/views"
	// Import views to trigger init() registration
	_ "call-orchestrator-demo/views/aptos"
	_ "call-orchestrator-demo/views/evm"
	_ "call-orchestrator-demo/views/solana"
)

func main() {
	// Command line flags
	addressRefsPath := flag.String("addresses", "address_refs.json", "Path to address_refs.json")
	networkConfigPath := flag.String("network", "testnet.yaml", "Path to network config YAML")
	outputPath := flag.String("output", "", "Output file path (default: stdout)")
	timeout := flag.Duration("timeout", 30*time.Minute, "Overall timeout for all operations")
	maxRPCWorkers := flag.Int("rpc-workers", 8, "Max concurrent requests per RPC endpoint")
	flag.Parse()

	fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║           CCIP State View Generator                           ║")
	fmt.Println("║   Using Call Orchestrator for reliable RPC calls              ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
	fmt.Println()

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	// =====================================================
	// 1. Load Configuration
	// =====================================================
	fmt.Println("Loading configuration...")

	// Load address references
	addressRefs, err := LoadAddressRefs(*addressRefsPath)
	if err != nil {
		fmt.Printf("Error loading address refs: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  Loaded %d address references\n", len(addressRefs))

	// Load network config
	networkConfig, err := LoadNetworkConfig(*networkConfigPath)
	if err != nil {
		fmt.Printf("Error loading network config: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  Loaded %d networks\n", len(networkConfig.Networks))

	// Create chain registry
	chainRegistry := NewChainRegistry(networkConfig)

	// Note: Aptos clients are not created here to avoid connection exhaustion.
	// The Aptos bindings bypass our rate-limited CallManager and can overwhelm
	// the system with unthrottled HTTP connections.

	// =====================================================
	// 2. Create Call Manager
	// =====================================================
	fmt.Println("\nInitializing Call Orchestrator...")

	callManager := NewCallManagerWithConfig(chainRegistry)
	defer callManager.Close()

	// Set RPC worker limits
	for selector := range chainRegistry.chains {
		callManager.SetRPCLimit(selector, *maxRPCWorkers)
	}

	// =====================================================
	// 3. Filter and Categorize Address Refs
	// =====================================================
	fmt.Println("\nAnalyzing contracts...")

	supported := make([]AddressRef, 0)
	unsupported := make([]AddressRef, 0)
	noChain := make([]AddressRef, 0)

	for _, ref := range addressRefs {
		// Check if we have the chain configured
		chain := chainRegistry.GetChain(ref.ChainSelector)
		if chain == nil {
			noChain = append(noChain, ref)
			continue
		}

		// Check if we support this contract type + version
		if views.IsSupported(string(chain.Family), ref.Type, ref.Version) {
			supported = append(supported, ref)
		} else {
			unsupported = append(unsupported, ref)
		}
	}

	// Collect all unique chain selectors from address refs (for remote chain discovery)
	chainSelectorSet := make(map[uint64]struct{})
	for _, ref := range addressRefs {
		chainSelectorSet[ref.ChainSelector] = struct{}{}
	}
	allChainSelectors := make([]uint64, 0, len(chainSelectorSet))
	for sel := range chainSelectorSet {
		allChainSelectors = append(allChainSelectors, sel)
	}

	fmt.Printf("  Supported contracts:   %d\n", len(supported))
	fmt.Printf("  Unsupported contracts: %d\n", len(unsupported))
	fmt.Printf("  No chain config:       %d\n", len(noChain))
	fmt.Printf("  Unique chains:         %d\n", len(allChainSelectors))

	// Print warnings for unsupported
	if len(unsupported) > 0 {
		fmt.Println("\n⚠ Unsupported contract types/versions:")
		typeVersionCount := make(map[string]int)
		for _, ref := range unsupported {
			key := fmt.Sprintf("%s@%s", ref.Type, ref.Version)
			typeVersionCount[key]++
		}
		for key, count := range typeVersionCount {
			fmt.Printf("    - %s (x%d)\n", key, count)
		}
	}

	if len(noChain) > 0 {
		fmt.Println("\n⚠ Contracts on unconfigured chains:")
		chainCount := make(map[uint64]int)
		for _, ref := range noChain {
			chainCount[ref.ChainSelector]++
		}
		for selector, count := range chainCount {
			fmt.Printf("    - Chain %d (x%d contracts)\n", selector, count)
		}
	}

	// =====================================================
	// 4. Generate Views Concurrently
	// =====================================================
	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("Generating views for %d contracts...\n", len(supported))
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	start := time.Now()

	// Start real-time stats printer
	statsDone := make(chan struct{})
	go printLiveStats(callManager, statsDone, len(supported))

	type viewResult struct {
		Ref     AddressRef
		View    map[string]any
		Error   error
		Timing  time.Duration
		Skipped bool
	}

	results := make([]viewResult, len(supported))
	var wg sync.WaitGroup

	// Create a wrapper CallManager that implements the views interface
	viewCallManager := &viewCallManagerAdapter{cm: callManager}

	for i, ref := range supported {
		wg.Add(1)
		go func(idx int, r AddressRef) {
			defer wg.Done()

			chain := chainRegistry.GetChain(r.ChainSelector)
			viewStart := time.Now()

			// Parse address based on chain family
			var addr []byte
			var err error
			if chain.Family == ChainFamilySVM {
				// Solana uses base58 addresses - store as raw string bytes
				// The executor will convert to base58 as needed
				addr = []byte(r.Address)
			} else {
				// EVM and others use hex addresses
				addr, err = views.ParseAddress(r.Address)
				if err != nil {
					results[idx] = viewResult{Ref: r, Error: fmt.Errorf("invalid address: %w", err)}
					return
				}
			}

			// Create view context
			viewCtx := &views.ViewContext{
				Address:           addr,
				AddressHex:        r.Address,
				ChainSelector:     r.ChainSelector,
				Qualifier:         r.Qualifier,
				CallManager:       viewCallManager,
				ChainFamily:       string(chain.Family),
				AllChainSelectors: allChainSelectors,
			}

			// Get the view function - exact version match only
			viewFn, ok := views.Get(string(chain.Family), r.Type, r.Version)
			if !ok {
				// Skip unsupported versions silently
				results[idx] = viewResult{Ref: r, Skipped: true}
				return
			}

			// Execute the view
			view, err := viewFn(viewCtx)
			viewTiming := time.Since(viewStart)

			// Add metadata
			if view != nil {
				view["_qualifier"] = r.Qualifier
				view["_type"] = r.Type
				view["_requestedVersion"] = r.Version
			}

			results[idx] = viewResult{
				Ref:    r,
				View:   view,
				Error:  err,
				Timing: viewTiming,
			}
		}(i, ref)
	}

	wg.Wait()
	close(statsDone) // Stop the stats printer
	totalDuration := time.Since(start)
	fmt.Println() // Clear the stats line

	// =====================================================
	// 5. Collect Results
	// =====================================================
	successCount := 0
	errorCount := 0
	skippedCount := 0

	// Group results by chain
	chainViews := make(map[uint64][]map[string]any)
	var errors []map[string]any
	var skipped []map[string]any

	for _, r := range results {
		if r.Skipped {
			skippedCount++
			skipped = append(skipped, map[string]any{
				"address":       r.Ref.Address,
				"chainSelector": r.Ref.ChainSelector,
				"type":          r.Ref.Type,
				"version":       r.Ref.Version,
				"reason":        "no view implementation for this version",
			})
		} else if r.Error != nil {
			errorCount++
			errors = append(errors, map[string]any{
				"address":       r.Ref.Address,
				"chainSelector": r.Ref.ChainSelector,
				"type":          r.Ref.Type,
				"version":       r.Ref.Version,
				"error":         r.Error.Error(),
			})
		} else {
			successCount++
			chainViews[r.Ref.ChainSelector] = append(chainViews[r.Ref.ChainSelector], r.View)
		}
	}

	// Build final output
	output := map[string]any{
		"_meta": map[string]any{
			"generatedAt":             time.Now().UTC().Format(time.RFC3339),
			"totalContracts":          len(supported),
			"successfulViews":         successCount,
			"failedViews":             errorCount,
			"skippedNoImplementation": skippedCount,
			"unsupportedChainFamily":  len(unsupported),
			"noChainConfig":           len(noChain),
			"totalDuration":           totalDuration.String(),
		},
		"chains": chainViews,
	}

	if len(errors) > 0 {
		output["_errors"] = errors
	}

	if len(skipped) > 0 {
		output["_skipped"] = skipped
	}

	// =====================================================
	// 6. Output Results
	// =====================================================
	jsonOutput, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling output: %v\n", err)
		os.Exit(1)
	}

	if *outputPath != "" {
		if err := os.WriteFile(*outputPath, jsonOutput, 0644); err != nil {
			fmt.Printf("Error writing output file: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("\nOutput written to: %s\n", *outputPath)
	} else {
		fmt.Println("\n" + string(jsonOutput))
	}

	// =====================================================
	// 7. Print Statistics
	// =====================================================
	stats := callManager.Stats()

	fmt.Println("\n╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                      STATISTICS                               ║")
	fmt.Println("╠═══════════════════════════════════════════════════════════════╣")
	fmt.Printf("║  Total Duration:        %-37v ║\n", totalDuration.Round(time.Millisecond))
	fmt.Printf("║  Contracts Processed:   %-37d ║\n", len(supported))
	fmt.Printf("║  Successful Views:      %-37d ║\n", successCount)
	fmt.Printf("║  Skipped (no impl):     %-37d ║\n", skippedCount)
	fmt.Printf("║  Failed Views:          %-37d ║\n", errorCount)
	fmt.Println("╠═══════════════════════════════════════════════════════════════╣")
	fmt.Printf("║  Total RPC Calls:       %-37d ║\n", stats.TotalCalls)
	fmt.Printf("║  Cache Hits:            %-37d ║\n", stats.CacheHits)
	fmt.Printf("║  Deduped Calls:         %-37d ║\n", stats.DedupedCalls)
	fmt.Printf("║  Total Retries:         %-37d ║\n", stats.TotalRetries)
	fmt.Printf("║  RPC Errors:            %-37d ║\n", stats.Errors)
	fmt.Printf("║  Cache Size:            %-37d ║\n", callManager.CacheSize())
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")

	// Check context for cancellation
	select {
	case <-ctx.Done():
		fmt.Println("\n⚠ Operation timed out or was cancelled")
		os.Exit(1)
	default:
	}
}

// viewCallManagerAdapter adapts CallManager to views.CallManagerInterface
type viewCallManagerAdapter struct {
	cm *CallManager
}

func (a *viewCallManagerAdapter) Execute(call views.Call) views.CallResult {
	// Convert views.Call to main.Call
	mainCall := Call{
		ChainID: call.ChainID,
		Target:  call.Target,
		Data:    call.Data,
	}
	result := a.cm.Execute(mainCall)
	return views.CallResult{
		Data:    result.Data,
		Error:   result.Error,
		Cached:  result.Cached,
		Retries: result.Retries,
	}
}

// printLiveStats prints real-time stats during view generation.
func printLiveStats(cm *CallManager, done chan struct{}, totalContracts int) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	startTime := time.Now()
	lastPrint := 0

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			stats := cm.LiveStats()
			elapsed := time.Since(startTime).Round(time.Second)

			// Calculate calls per second
			callsPerSec := float64(0)
			if elapsed.Seconds() > 0 {
				callsPerSec = float64(stats.TotalCalls) / elapsed.Seconds()
			}

			// Clear previous output (move up and clear lines)
			if lastPrint > 0 {
				for i := 0; i < lastPrint; i++ {
					fmt.Print("\033[A\033[K") // Move up and clear line
				}
			}

			// Print header
			lines := 0
			fmt.Printf("╔══════════════════════════════════════════════════════════════════════════════╗\n")
			lines++
			fmt.Printf("║ [%v] Calls: %-8d (%.0f/s) | Cache: %-6d | Retries: %-5d | Err: %-6d ║\n",
				elapsed, stats.TotalCalls, callsPerSec, stats.CacheSize, stats.TotalRetries, stats.Errors)
			lines++
			fmt.Printf("╠══════════════════════════════════════════════════════════════════════════════╝\n")
			lines++
			fmt.Printf("║ %-20s │ %10s │ %10s │ %8s ║\n", "Chain", "InFlight", "Capacity", "Usage")
			lines++
			fmt.Printf("╠══════════════════════╪════════════╪════════════╪══════════╣\n")
			lines++

			// Get per-chain stats
			chainIDs := cm.AllChainIDs()

			// Sort chains by InFlight (descending) to show busiest first
			type chainStat struct {
				id       uint64
				inFlight int
				capacity int
			}
			chainStats := make([]chainStat, 0, len(chainIDs))
			for _, chainID := range chainIDs {
				rpcStats := cm.LiveRPCStats(chainID)
				inFlight := 0
				capacity := 0
				for _, rpc := range rpcStats {
					inFlight += rpc.ActiveCalls
					capacity += rpc.MaxWorkers
				}
				if capacity > 0 { // Only show chains with capacity
					chainStats = append(chainStats, chainStat{chainID, inFlight, capacity})
				}
			}

			// Sort by inFlight descending
			for i := 0; i < len(chainStats); i++ {
				for j := i + 1; j < len(chainStats); j++ {
					if chainStats[j].inFlight > chainStats[i].inFlight {
						chainStats[i], chainStats[j] = chainStats[j], chainStats[i]
					}
				}
			}

			// Show top 10 busiest chains
			shown := 0
			for _, cs := range chainStats {
				if shown >= 10 {
					break
				}
				pct := makeProgressBar(cs.inFlight, cs.capacity, 8)
				fmt.Printf("║ %-20d │ %10d │ %10d │ %s     ║\n", cs.id, cs.inFlight, cs.capacity, pct)
				lines++
				shown++
			}

			// Summary row
			fmt.Printf("╠══════════════════════╪════════════╪════════════╪══════════╣\n")
			lines++
			totalBar := makeProgressBar(stats.TotalInFlight, stats.TotalMaxCap, 8)
			fmt.Printf("║ %-20s │ %10d │ %10d │ %s     ║\n", "TOTAL", stats.TotalInFlight, stats.TotalMaxCap, totalBar)
			lines++
			fmt.Printf("╚═══════════════════════════════════════════════════════════╝\n")
			lines++

			lastPrint = lines
		}
	}
}

// makeProgressBar creates a simple progress bar string.
func makeProgressBar(current, max, width int) string {
	if max == 0 {
		return "   0%   "
	}
	pct := float64(current) / float64(max) * 100
	filled := int(float64(width) * float64(current) / float64(max))
	if filled > width {
		filled = width
	}
	bar := ""
	for i := 0; i < filled; i++ {
		bar += "█"
	}
	for i := filled; i < width; i++ {
		bar += "░"
	}
	return fmt.Sprintf("%3.0f%%", pct)
}
