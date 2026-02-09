package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"give-me-state-v2/orchestrator"
	"give-me-state-v2/orchestrator/aptos"
	"give-me-state-v2/orchestrator/evm"
	"give-me-state-v2/orchestrator/svm"
	"give-me-state-v2/views"
	_ "give-me-state-v2/views/aptos"
	_ "give-me-state-v2/views/evm"
	_ "give-me-state-v2/views/solana"
)

func main() {
	addressRefsPath := flag.String("addresses", "address_refs.json", "Path to address_refs.json")
	networkConfigPath := flag.String("network", "testnet.yaml", "Path to network config YAML")
	outputPath := flag.String("output", "", "Output file path (default: stdout)")
	timeout := flag.Duration("timeout", 30*time.Minute, "Overall timeout for all operations")
	maxConcurrent := flag.Int("max-concurrent", 12, "Max concurrent workers per RPC endpoint (scales from min to this when healthy)")
	minConcurrent := flag.Int("min-concurrent", 1, "Min workers per endpoint when unhealthy (scale down to this, not below)")
	format := flag.Bool("format", false, "Format output to match state.json structure")
	live := flag.Bool("live", true, "Show live RPC stats and progress during run")
	flag.Parse()

	fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║           CCIP State View Generator (v2)                    ║")
	fmt.Println("║   Typed orchestrators + generic engine                        ║")
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
	fmt.Println()

	ctx, cancel := context.WithTimeout(context.Background(), *timeout)
	defer cancel()

	fmt.Println("Loading configuration...")
	addressRefs, err := LoadAddressRefs(*addressRefsPath)
	if err != nil {
		fmt.Printf("Error loading address refs: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  Loaded %d address references\n", len(addressRefs))

	networkConfig, err := LoadNetworkConfig(*networkConfigPath)
	if err != nil {
		fmt.Printf("Error loading network config: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  Loaded %d networks\n", len(networkConfig.Networks))

	chainRegistry := NewChainRegistry(networkConfig)

	// Build generic engine and typed orchestrators (EVM, SVM, Aptos)
	generic := orchestrator.NewGeneric()
	retryable := retryableKeywords()

	evmChains := buildEVMChainEndpoints(chainRegistry, *minConcurrent, *maxConcurrent)
	evmOrc, err := evm.NewEVMOrchestrator(generic, evmChains, retryable)
	if err != nil {
		fmt.Printf("Error creating EVM orchestrator: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  Registered %d EVM chains with generic engine\n", len(evmChains))

	svmChains := buildSVMChainEndpoints(chainRegistry, *minConcurrent, *maxConcurrent)
	var svmOrc *svm.SVMOrchestrator
	if len(svmChains) > 0 {
		svmOrc, err = svm.NewSVMOrchestrator(generic, svmChains, retryable)
		if err != nil {
			fmt.Printf("Error creating SVM orchestrator: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("  Registered %d SVM (Solana) chains with generic engine\n", len(svmChains))
	}

	aptosChains := buildAptosChainEndpoints(chainRegistry, *minConcurrent, *maxConcurrent)
	var aptosOrc *aptos.AptosOrchestrator
	if len(aptosChains) > 0 {
		aptosOrc, err = aptos.NewAptosOrchestrator(generic, aptosChains, retryable)
		if err != nil {
			fmt.Printf("Error creating Aptos orchestrator: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("  Registered %d Aptos chains with generic engine\n", len(aptosChains))
	}

	if len(evmChains)+len(svmChains)+len(aptosChains) == 0 {
		fmt.Println("No EVM/SVM/Aptos chains in network config; nothing to do.")
		os.Exit(0)
	}

	// Categorize refs
	supported := make([]AddressRef, 0)
	unsupported := make([]AddressRef, 0)
	noChain := make([]AddressRef, 0)
	for _, ref := range addressRefs {
		chain := chainRegistry.GetChain(ref.ChainSelector)
		if chain == nil {
			noChain = append(noChain, ref)
			continue
		}
		if chain.Family != ChainFamilyEVM && chain.Family != ChainFamilySVM && chain.Family != ChainFamilyAptos {
			unsupported = append(unsupported, ref)
			continue
		}
		if views.IsSupported(string(chain.Family), ref.Type, ref.Version) {
			supported = append(supported, ref)
		} else {
			unsupported = append(unsupported, ref)
		}
	}

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

	fmt.Println("\n━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Printf("Generating views for %d contracts...\n", len(supported))
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	start := time.Now()
	var viewsDone atomic.Int64
	statsDone := make(chan struct{})
	if *live {
		go printLiveStats(generic, chainRegistry, &viewsDone, int64(len(supported)), statsDone)
	}

	type viewResult struct {
		Ref     AddressRef
		View    map[string]any
		Error   error
		Timing  time.Duration
		Skipped bool
	}
	results := make([]viewResult, len(supported))
	var wg sync.WaitGroup

	for i, ref := range supported {
		wg.Add(1)
		go func(idx int, r AddressRef) {
			defer wg.Done()
			defer viewsDone.Add(1)
			chain := chainRegistry.GetChain(r.ChainSelector)
			viewStart := time.Now()

			var addr []byte
			var err error
			switch chain.Family {
			case ChainFamilySVM:
				addr = []byte(r.Address) // base58 string as bytes
			case ChainFamilyAptos:
				addr = []byte(r.Address) // hex string as bytes (e.g. "0x1")
			default:
				addr, err = views.ParseAddress(r.Address)
				if err != nil {
					results[idx] = viewResult{Ref: r, Error: fmt.Errorf("invalid address: %w", err)}
					return
				}
			}

			var typedOrc orchestrator.TypedOrchestratorInterface
			switch chain.Family {
			case ChainFamilyEVM:
				typedOrc = evmOrc
			case ChainFamilySVM:
				if svmOrc == nil {
					results[idx] = viewResult{Ref: r, Skipped: true}
					return
				}
				typedOrc = svmOrc
			case ChainFamilyAptos:
				if aptosOrc == nil {
					results[idx] = viewResult{Ref: r, Skipped: true}
					return
				}
				typedOrc = aptosOrc
			default:
				results[idx] = viewResult{Ref: r, Skipped: true}
				return
			}

			viewCtx := &views.ViewContext{
				Address:           addr,
				AddressHex:        r.Address,
				ChainSelector:     r.ChainSelector,
				Qualifier:         r.Qualifier,
				TypedOrchestrator: typedOrc,
				ChainFamily:       string(chain.Family),
				AllChainSelectors: allChainSelectors,
			}

			viewFn, ok := views.Get(string(chain.Family), r.Type, r.Version)
			if !ok {
				results[idx] = viewResult{Ref: r, Skipped: true}
				return
			}

			view, err := viewFn(viewCtx)
			viewTiming := time.Since(viewStart)
			if view != nil {
				view["_qualifier"] = r.Qualifier
				view["_type"] = r.Type
				view["_requestedVersion"] = r.Version
			}
			results[idx] = viewResult{Ref: r, View: view, Error: err, Timing: viewTiming}
		}(i, ref)
	}

	wg.Wait()
	if *live {
		close(statsDone)
	}
	totalDuration := time.Since(start)

	successCount := 0
	errorCount := 0
	skippedCount := 0
	chainViews := make(map[uint64][]map[string]any)
	var errors []map[string]any
	var skipped []map[string]any

	for _, r := range results {
		if r.Skipped {
			skippedCount++
			skipped = append(skipped, map[string]any{
				"address": r.Ref.Address, "chainSelector": r.Ref.ChainSelector,
				"type": r.Ref.Type, "version": r.Ref.Version, "reason": "no view implementation for this version",
			})
		} else if r.Error != nil {
			errorCount++
			errors = append(errors, map[string]any{
				"address": r.Ref.Address, "chainSelector": r.Ref.ChainSelector,
				"type": r.Ref.Type, "version": r.Ref.Version, "error": r.Error.Error(),
			})
		} else {
			successCount++
			chainViews[r.Ref.ChainSelector] = append(chainViews[r.Ref.ChainSelector], r.View)
		}
	}

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

	jsonOutput, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling output: %v\n", err)
		os.Exit(1)
	}
	if *format {
		legacyOutput, err := formatOutput(jsonOutput, chainRegistry.GetAllChains())
		if err != nil {
			fmt.Printf("Error transforming to legacy format: %v\n", err)
			os.Exit(1)
		}
		jsonOutput = legacyOutput
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

	fmt.Println("\n╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                      STATISTICS                               ║")
	fmt.Println("╠═══════════════════════════════════════════════════════════════╣")
	fmt.Printf("║  Total Duration:        %-37v ║\n", totalDuration.Round(time.Millisecond))
	fmt.Printf("║  Contracts Processed:   %-37d ║\n", len(supported))
	fmt.Printf("║  Successful Views:      %-37d ║\n", successCount)
	fmt.Printf("║  Skipped (no impl):     %-37d ║\n", skippedCount)
	fmt.Printf("║  Failed Views:          %-37d ║\n", errorCount)
	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")

	select {
	case <-ctx.Done():
		fmt.Println("\n⚠ Operation timed out or was cancelled")
		os.Exit(1)
	default:
	}
}

func buildEVMChainEndpoints(registry *ChainRegistry, minConcurrent, maxConcurrent int) []evm.ChainEndpoints {
	allChains := registry.GetAllChains()
	var out []evm.ChainEndpoints
	for selector, info := range allChains {
		if info.Family != ChainFamilyEVM {
			continue
		}
		var endpoints []orchestrator.EndpointConfig
		for _, rpc := range info.RPCs {
			if rpc.HTTPURL == "" || len(rpc.HTTPURL) < 4 || rpc.HTTPURL[:4] != "http" {
				continue
			}
			endpoints = append(endpoints, orchestrator.EndpointConfig{
				URL:               rpc.HTTPURL,
				MinConcurrent:     minConcurrent,
				MaxConcurrent:     maxConcurrent,
				TargetSuccessRate: 0.95,
				Timeout:           30,
			})
		}
		if len(endpoints) > 0 {
			out = append(out, evm.ChainEndpoints{ChainID: selector, Endpoints: endpoints})
		}
	}
	return out
}

func buildSVMChainEndpoints(registry *ChainRegistry, minConcurrent, maxConcurrent int) []svm.ChainEndpoints {
	allChains := registry.GetAllChains()
	var out []svm.ChainEndpoints
	for selector, info := range allChains {
		if info.Family != ChainFamilySVM {
			continue
		}
		var endpoints []orchestrator.EndpointConfig
		for _, rpc := range info.RPCs {
			if rpc.HTTPURL == "" || len(rpc.HTTPURL) < 4 || rpc.HTTPURL[:4] != "http" {
				continue
			}
			endpoints = append(endpoints, orchestrator.EndpointConfig{
				URL:               rpc.HTTPURL,
				MinConcurrent:     minConcurrent,
				MaxConcurrent:     maxConcurrent,
				TargetSuccessRate: 0.95,
				Timeout:           30,
			})
		}
		if len(endpoints) > 0 {
			out = append(out, svm.ChainEndpoints{ChainID: selector, Endpoints: endpoints})
		}
	}
	return out
}

func buildAptosChainEndpoints(registry *ChainRegistry, minConcurrent, maxConcurrent int) []aptos.ChainEndpoints {
	allChains := registry.GetAllChains()
	var out []aptos.ChainEndpoints
	for selector, info := range allChains {
		if info.Family != ChainFamilyAptos {
			continue
		}
		var endpoints []orchestrator.EndpointConfig
		for _, rpc := range info.RPCs {
			if rpc.HTTPURL == "" || len(rpc.HTTPURL) < 4 || rpc.HTTPURL[:4] != "http" {
				continue
			}
			endpoints = append(endpoints, orchestrator.EndpointConfig{
				URL:               rpc.HTTPURL,
				MinConcurrent:     minConcurrent,
				MaxConcurrent:     maxConcurrent,
				TargetSuccessRate: 0.95,
				Timeout:           30,
			})
		}
		if len(endpoints) > 0 {
			out = append(out, aptos.ChainEndpoints{ChainID: selector, Endpoints: endpoints})
		}
	}
	return out
}

// printLiveStats prints live RPC failure rates, queue depth, and progress every 2s.
func printLiveStats(g *orchestrator.Generic, registry *ChainRegistry, viewsDone *atomic.Int64, totalViews int64, done chan struct{}) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	startTime := time.Now()
	lastLines := 0
	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			doneN := viewsDone.Load()
			elapsed := time.Since(startTime).Round(time.Second)
			stats := g.LiveStats()

			// Clear previous block
			if lastLines > 0 {
				for i := 0; i < lastLines; i++ {
					fmt.Print("\033[A\033[K")
				}
			}
			lines := 0

			// Progress + summary
			fmt.Printf("╔══════════════════════════════════════════════════════════════════════════════════╗\n")
			lines++
			fmt.Printf("║ [%v] Views: %d / %d (%.0f%%)                                                    ║\n",
				elapsed, doneN, totalViews, 100*float64(doneN)/float64(totalViews))
			lines++
			fmt.Printf("╠══════════════════════════════════════════════════════════════════════════════════╣\n")
			lines++
			fmt.Printf("║ %-22s │ %8s │ %10s │ %8s │ %-36s ║\n", "Chain", "Queue", "Workers", "Fail%", "Endpoint")
			lines++
			fmt.Printf("╠══════════════════════════╪══════════╪════════════╪══════════╪══════════════════════╣\n")
			lines++

			// Sort orc IDs by queue depth (busiest first), then by chain name
			type orcRow struct {
				id    string
				label string
				orc   orchestrator.OrcLiveStats
			}
			var rows []orcRow
			for id, orc := range stats {
				label := orcIDToChainLabel(id, registry)
				rows = append(rows, orcRow{id, label, orc})
			}
			sort.Slice(rows, func(i, j int) bool {
				if rows[i].orc.QueueDepth != rows[j].orc.QueueDepth {
					return rows[i].orc.QueueDepth > rows[j].orc.QueueDepth
				}
				return rows[i].label < rows[j].label
			})

			shown := 0
			for _, row := range rows {
				if shown >= 12 {
					break
				}
				queueStr := fmt.Sprintf("%d", row.orc.QueueDepth)
				for ei, ep := range row.orc.Endpoints {
					failPct := (1 - ep.SuccessRate) * 100
					workersStr := fmt.Sprintf("%d/%d", ep.Workers, ep.MaxConcurrent)
					endpointShort := shortenURL(ep.URL, 36)
					if ei == 0 {
						fmt.Printf("║ %-22s │ %8s │ %10s │ %7.1f%% │ %-36s ║\n",
							truncate(row.label, 22), queueStr, workersStr, failPct, endpointShort)
					} else {
						fmt.Printf("║ %-22s │ %8s │ %10s │ %7.1f%% │ %-36s ║\n",
							"", "", "", failPct, endpointShort)
					}
					lines++
					shown++
					if shown >= 12 {
						break
					}
				}
				if len(row.orc.Endpoints) == 0 {
					fmt.Printf("║ %-22s │ %8s │ %10s │ %8s │ %-36s ║\n",
						truncate(row.label, 22), queueStr, "-", "-", "-")
					lines++
					shown++
				}
			}
			fmt.Printf("╚══════════════════════════════════════════════════════════════════════════════════╝\n")
			lines++
			lastLines = lines
		}
	}
}

func orcIDToChainLabel(id string, registry *ChainRegistry) string {
	if strings.HasPrefix(id, "evm-") {
		s := strings.TrimPrefix(id, "evm-")
		sel, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return id
		}
		// Prefer canonical chain name (ChainSelectorToName), not the first RPC's name (e.g. "Alchemy")
		if name, ok := ChainSelectorToName[sel]; ok {
			return name
		}
		chain := registry.GetChain(sel)
		if chain != nil && chain.Name != "" {
			return chain.Name
		}
		return s
	}
	return id
}

func shortenURL(url string, maxLen int) string {
	url = strings.TrimPrefix(url, "https://")
	url = strings.TrimPrefix(url, "http://")
	if len(url) <= maxLen {
		return url
	}
	return url[:maxLen-3] + "..."
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

func retryableKeywords() []string {
	return []string{
		// Rate limit
		"rate limit", "429", "too many requests",
		// Timeout
		"timeout", "deadline exceeded",
		// Connection
		"connection refused", "connection reset", "EOF", "broken pipe",
		// DNS
		"no such host", "lookup",
		// TLS / cert
		"tls:", "x509:", "certificate",
		// Server errors (5xx) and connect failures
		"500", "502", "503", "504", "failed to connect",
		// Transient auth (e.g. temporary 401)
		"http error 401", "http 401", "http 429",
		// RPC overload / bad gateway
		"Unsupported RPC call", "http error 400", "http error 404", "http 404", "http 403",
		// RPC node sync / method not available
		"upstream does not have the requested block", "-32601",
	}
}
