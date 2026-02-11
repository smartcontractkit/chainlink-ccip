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
	"give-me-state-v2/orchestrator/jd"
	"give-me-state-v2/orchestrator/svm"
	"give-me-state-v2/views"
	_ "give-me-state-v2/views/aptos"
	_ "give-me-state-v2/views/evm"
	_ "give-me-state-v2/views/solana"
)

// TODO smarter json creation and merging? Like if some value wasn't updated, then dont override it.
// This logic should happen for the raw JSON outputted by the tool, not the formatted one.

func main() {
	addressRefsPath := flag.String("addresses", "example_address_refs.json", "Path to address_refs.json")
	networkConfigPath := flag.String("network", "example.yaml", "Path to network config YAML")
	outputPath := flag.String("output", "", "Output file path (default: stdout)")
	timeout := flag.Duration("timeout", 30*time.Minute, "Overall timeout for all operations")
	workersPerEndpoint := flag.Int("workers", 12, "Worker goroutines per RPC endpoint")
	format := flag.Bool("format", false, "Format output to match state.json structure")
	live := flag.Bool("live", true, "Show live RPC stats and progress during run")
	nops := flag.Bool("nops", false, "Include node operator data from Job Distributor (requires JD_* env vars)")
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

	evmChains := buildEVMChainEndpoints(chainRegistry, *workersPerEndpoint)
	evmOrc, err := evm.NewEVMOrchestrator(generic, evmChains, retryable)
	if err != nil {
		fmt.Printf("Error creating EVM orchestrator: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("  Registered %d EVM chains with generic engine\n", len(evmChains))

	svmChains := buildSVMChainEndpoints(chainRegistry, *workersPerEndpoint)
	var svmOrc *svm.SVMOrchestrator
	if len(svmChains) > 0 {
		svmOrc, err = svm.NewSVMOrchestrator(generic, svmChains, retryable)
		if err != nil {
			fmt.Printf("Error creating SVM orchestrator: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("  Registered %d SVM (Solana) chains with generic engine\n", len(svmChains))
	}

	aptosChains := buildAptosChainEndpoints(chainRegistry, *workersPerEndpoint)
	var aptosOrc *aptos.AptosOrchestrator
	if len(aptosChains) > 0 {
		aptosOrc, err = aptos.NewAptosOrchestrator(generic, aptosChains, retryable)
		if err != nil {
			fmt.Printf("Error creating Aptos orchestrator: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("  Registered %d Aptos chains with generic engine\n", len(aptosChains))
	}

	// Job Distributor orchestrator (optional -- only created when -nops flag is provided).
	// Reads connection details from environment variables:
	//   JD_GRPC_URL              - gRPC endpoint (e.g. "jd.example.com:443")
	//   JD_TLS                   - "true" to enable TLS (default: false)
	//   JD_COGNITO_CLIENT_ID     - Cognito OAuth2 client ID (optional, for auth)
	//   JD_COGNITO_CLIENT_SECRET - Cognito OAuth2 client secret
	//   JD_USERNAME              - Cognito username
	//   JD_PASSWORD              - Cognito password
	//   JD_AWS_REGION            - AWS region for Cognito (e.g. "us-west-2")
	var jdOrc *jd.JDOrchestrator
	if *nops {
		// Load .env file if present (won't override already-exported vars).
		if err := loadDotEnv(".env"); err != nil && !os.IsNotExist(err) {
			fmt.Printf("Warning: failed to load .env file: %v\n", err)
		}
		jdGRPCURL := os.Getenv("JD_GRPC_URL")
		if jdGRPCURL == "" {
			fmt.Println("Error: -nops flag requires JD_GRPC_URL environment variable to be set")
			os.Exit(1)
		}
		jdTLS := strings.EqualFold(os.Getenv("JD_TLS"), "true")
		jdCfg := jd.JDConfig{
			GRPCURL: jdGRPCURL,
			TLS:     jdTLS,
		}
		if cognitoClientID := os.Getenv("JD_COGNITO_CLIENT_ID"); cognitoClientID != "" {
			jdCfg.Auth = &jd.JDAuthConfig{
				CognitoClientID:     cognitoClientID,
				CognitoClientSecret: os.Getenv("JD_COGNITO_CLIENT_SECRET"),
				Username:            os.Getenv("JD_USERNAME"),
				Password:            os.Getenv("JD_PASSWORD"),
				AWSRegion:           os.Getenv("JD_AWS_REGION"),
			}
		}
		jdOrc, err = jd.NewJDOrchestrator(ctx, jdCfg)
		if err != nil {
			fmt.Printf("Warning: JD orchestrator setup failed (continuing without JD data): %v\n", err)
			jdOrc = nil
		} else {
			defer jdOrc.Close()
			fmt.Printf("  Connected to Job Distributor at %s\n", jdGRPCURL)
		}
	}

	if len(evmChains)+len(svmChains)+len(aptosChains) == 0 && jdOrc == nil {
		fmt.Println("No EVM/SVM/Aptos chains in network config and no JD configured; nothing to do.")
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

	// Query JD for node operator data (non-blocking on chain views).
	if jdOrc != nil {
		nodeOps, jdErr := queryJD(ctx, jdOrc)
		if jdErr != nil {
			fmt.Printf("Warning: JD query failed (continuing without JD data): %v\n", jdErr)
		} else {
			output["nodeOperators"] = nodeOps
		}
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

	printFinalStats(totalDuration, len(supported), successCount, skippedCount, errorCount, evmOrc, chainRegistry)

	select {
	case <-ctx.Done():
		fmt.Println("\n⚠ Operation timed out or was cancelled")
		os.Exit(1)
	default:
	}
}

func buildEVMChainEndpoints(registry *ChainRegistry, workersPerEndpoint int) []evm.ChainEndpoints {
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
				URL:     rpc.HTTPURL,
				Workers: workersPerEndpoint,
				Timeout: 30,
			})
		}
		if len(endpoints) > 0 {
			out = append(out, evm.ChainEndpoints{ChainID: selector, Endpoints: endpoints})
		}
	}
	return out
}

func buildSVMChainEndpoints(registry *ChainRegistry, workersPerEndpoint int) []svm.ChainEndpoints {
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
				URL:     rpc.HTTPURL,
				Workers: workersPerEndpoint,
				Timeout: 30,
			})
		}
		if len(endpoints) > 0 {
			out = append(out, svm.ChainEndpoints{ChainID: selector, Endpoints: endpoints})
		}
	}
	return out
}

func buildAptosChainEndpoints(registry *ChainRegistry, workersPerEndpoint int) []aptos.ChainEndpoints {
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
				URL:     rpc.HTTPURL,
				Workers: workersPerEndpoint,
				Timeout: 30,
			})
		}
		if len(endpoints) > 0 {
			out = append(out, aptos.ChainEndpoints{ChainID: selector, Endpoints: endpoints})
		}
	}
	return out
}

// printLiveStats prints live progress bar, queue depth, and aggregate failure rates every 2s.
func printLiveStats(g *orchestrator.Generic, registry *ChainRegistry, viewsDone *atomic.Int64, totalViews int64, done chan struct{}) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	startTime := time.Now()
	lastLines := 0

	const barWidth = 50 // characters for the progress bar fill area

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

			pct := float64(doneN) / float64(totalViews)
			filled := int(pct * barWidth)
			if filled > barWidth {
				filled = barWidth
			}
			bar := strings.Repeat("█", filled) + strings.Repeat("░", barWidth-filled)

			fmt.Println("╔═══════════════════════════════════════════════════════════════╗")
			lines++
			viewsStr := fmt.Sprintf(" [%v] Views: %d / %d", elapsed, doneN, totalViews)
			fmt.Printf("║%s%-*s║\n", viewsStr, 61-len(viewsStr), "")
			lines++
			fmt.Printf("║ [%s] %3.0f%%  %-*s║\n", bar, pct*100,
				61-len(fmt.Sprintf(" [%s] %3.0f%%  ", bar, pct*100)), "")
			lines++
			fmt.Println("╠═══════════════════════════════════════════════════════════════╣")
			lines++
			fmt.Printf("║ %-30s │ %7s │ %9s │ %7s ║\n",
				"Chain", "Queue", "Workers", "Fail%")
			lines++
			fmt.Println("╠════════════════════════════════╪═════════╪═══════════╪═════════╣")
			lines++

			// Aggregate per-chain: one row per orchestrator
			type orcRow struct {
				label      string
				queueDepth int
				workers    int
				failPct    float64
			}
			var rows []orcRow
			for id, orc := range stats {
				label := orcIDToChainLabel(id, registry)
				var totalWorkers int
				var failSum float64
				for _, ep := range orc.Endpoints {
					totalWorkers = ep.Workers // same for all endpoints in an orc
					failSum += (1 - ep.SuccessRate)
				}
				avgFail := 0.0
				if len(orc.Endpoints) > 0 {
					avgFail = failSum / float64(len(orc.Endpoints)) * 100
				}
				rows = append(rows, orcRow{
					label:      label,
					queueDepth: orc.QueueDepth,
					workers:    totalWorkers,
					failPct:    avgFail,
				})
			}
			sort.Slice(rows, func(i, j int) bool {
				// Highest fail% first so struggling chains are visible at the top.
				if rows[i].queueDepth != rows[j].queueDepth {
					return rows[i].queueDepth > rows[j].queueDepth
				}
				if rows[i].failPct != rows[j].failPct {
					return rows[i].failPct > rows[j].failPct
				}
				return rows[i].label < rows[j].label
			})

			shown := 0
			for _, row := range rows {
				if shown >= 15 {
					break
				}
				fmt.Printf("║ %-30s │ %7d │ %9d │ %6.1f%% ║\n",
					truncate(row.label, 30),
					row.queueDepth, row.workers, row.failPct)
				lines++
				shown++
			}
			fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
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
		// Don't fall back to chain.Name -- it's often the RPC provider
		// name (e.g. "Alchemy"), not the chain name. Use selector instead.
		return "chain-" + s
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

func printFinalStats(totalDuration time.Duration, totalContracts, successCount, skippedCount, errorCount int,
	evmOrc *evm.EVMOrchestrator, registry *ChainRegistry) {

	evmStats := evmOrc.Stats()

	// Compute derived metrics
	logicalCalls := evmStats.CacheHits + evmStats.CacheDeduped + evmStats.CacheMisses
	uniqueCalls := evmStats.CacheMisses // calls that actually needed execution

	// Throughput
	durationSecs := totalDuration.Seconds()
	var logicalPerSec, httpPerSec float64
	if durationSecs > 0 {
		logicalPerSec = float64(logicalCalls) / durationSecs
		httpPerSec = float64(evmStats.TotalSuccesses) / durationSecs
	}

	// RPC reduction: compare logical calls vs successful HTTP calls
	// (retries inflate TotalHTTPCalls, so use Successes for a fair comparison)
	var rpcReductionPct float64
	if logicalCalls > 0 {
		rpcReductionPct = (1 - float64(evmStats.TotalSuccesses)/float64(logicalCalls)) * 100
	}

	// Avg batch size
	var avgBatchSize float64
	if evmStats.Multicall.BatchesSent > 0 {
		avgBatchSize = float64(evmStats.Multicall.CallsBatched) / float64(evmStats.Multicall.BatchesSent)
	}

	// Cache hit rate
	var cacheHitPct float64
	if logicalCalls > 0 {
		cacheHitPct = float64(evmStats.CacheHits+evmStats.CacheDeduped) / float64(logicalCalls) * 100
	}

	w := 61 // inner width between ║ and ║
	line := func(label string, value string) {
		fmt.Printf("║  %-28s %-*s ║\n", label, w-30, value)
	}
	sep := "╠═══════════════════════════════════════════════════════════════╣"

	fmt.Println("\n╔═══════════════════════════════════════════════════════════════╗")
	fmt.Println("║                         STATISTICS                           ║")
	fmt.Println(sep)
	line("Total Duration:", totalDuration.Round(time.Millisecond).String())
	line("Contracts Processed:", fmt.Sprintf("%d", totalContracts))
	line("Successful / Skipped / Failed:", fmt.Sprintf("%d / %d / %d", successCount, skippedCount, errorCount))

	fmt.Println(sep)
	fmt.Println("║                      RPC EFFICIENCY                          ║")
	fmt.Println(sep)
	line("Logical Calls (from views):", fmt.Sprintf("%d", logicalCalls))
	line("  Cache Hits:", fmt.Sprintf("%d (%.1f%%)", evmStats.CacheHits, cacheHitPct))
	line("  Deduped (in-flight):", fmt.Sprintf("%d", evmStats.CacheDeduped))
	line("  Unique (executed):", fmt.Sprintf("%d", uniqueCalls))
	line("Multicall Chains:", fmt.Sprintf("%d", evmStats.Multicall.MulticallChains))
	line("Multicall Batches Sent:", fmt.Sprintf("%d (avg %.1f calls/batch)", evmStats.Multicall.BatchesSent, avgBatchSize))
	line("Single Calls (unbatched):", fmt.Sprintf("%d", evmStats.Multicall.SingleCalls))
	if evmStats.Multicall.FallbackCalls > 0 {
		line("Multicall Fallbacks:", fmt.Sprintf("%d", evmStats.Multicall.FallbackCalls))
	}
	line("Successful HTTP Calls:", fmt.Sprintf("%d", evmStats.TotalSuccesses))
	if evmStats.TotalRetries > 0 {
		line("Retries (rate limits etc):", fmt.Sprintf("%d", evmStats.TotalRetries))
		line("Total HTTP Calls:", fmt.Sprintf("%d (incl. retries)", evmStats.TotalHTTPCalls))
	}
	line("RPC Reduction:", fmt.Sprintf("%.1f%% fewer calls than naive", rpcReductionPct))

	fmt.Println(sep)
	fmt.Println("║                       THROUGHPUT                             ║")
	fmt.Println(sep)
	line("Logical Calls/sec:", fmt.Sprintf("%.0f", logicalPerSec))
	line("Successful HTTP/sec:", fmt.Sprintf("%.0f", httpPerSec))

	// Top chains by RPC pressure (successful HTTP calls)
	type chainRow struct {
		label        string
		logicalCalls int64
		successes    int64
		retries      int64
		hasMulticall bool
	}
	var chainRows []chainRow
	for orcID, cs := range evmStats.PerChain {
		label := orcIDToChainLabel(orcID, registry)
		chainRows = append(chainRows, chainRow{
			label:        label,
			logicalCalls: cs.LogicalCalls,
			successes:    cs.Successes,
			retries:      cs.Retries,
			hasMulticall: cs.HasMulticall,
		})
	}
	sort.Slice(chainRows, func(i, j int) bool {
		return chainRows[i].successes > chainRows[j].successes
	})

	if len(chainRows) > 0 {
		fmt.Println(sep)
		fmt.Println("║              TOP CHAINS (RPC PRESSURE)                       ║")
		fmt.Printf("║  %-18s %7s %6s %7s %6s %3s  ║\n",
			"Chain", "Logical", "OK", "Saved", "Retry", "MC")
		fmt.Println(sep)
		shown := 0
		for _, cr := range chainRows {
			if shown >= 10 || (cr.successes == 0 && cr.retries == 0) {
				break
			}
			savedPct := 0.0
			if cr.logicalCalls > 0 {
				savedPct = (1 - float64(cr.successes)/float64(cr.logicalCalls)) * 100
			}
			mc := " -"
			if cr.hasMulticall {
				mc = " Y"
			}
			fmt.Printf("║  %-18s %7d %6d %6.0f%% %6d %3s  ║\n",
				truncate(cr.label, 18),
				cr.logicalCalls, cr.successes, savedPct, cr.retries, mc)
			shown++
		}
	}

	fmt.Println("╚═══════════════════════════════════════════════════════════════╝")
}

// queryJD fetches node + chain config data from the Job Distributor and
// returns a JSON-serializable structure for the "nodeOperators" output section.
func queryJD(ctx context.Context, jdOrc *jd.JDOrchestrator) (map[string]any, error) {
	// Step 1: List all nodes.
	nodes, err := jdOrc.ListNodes(ctx)
	if err != nil {
		return nil, fmt.Errorf("listing nodes: %w", err)
	}

	// Step 2: Collect node IDs and fetch chain configs for all nodes.
	nodeIDs := make([]string, 0, len(nodes))
	for _, n := range nodes {
		if id, ok := n["id"].(string); ok && id != "" {
			nodeIDs = append(nodeIDs, id)
		}
	}

	var chainConfigsByNode map[string][]map[string]any
	if len(nodeIDs) > 0 {
		chainConfigsByNode, err = jdOrc.ListNodeChainConfigs(ctx, nodeIDs)
		if err != nil {
			return nil, fmt.Errorf("listing chain configs: %w", err)
		}
	}

	// Step 3: Merge chain configs into each node.
	for _, node := range nodes {
		id, _ := node["id"].(string)
		if configs, ok := chainConfigsByNode[id]; ok {
			node["chainConfigs"] = configs
		} else {
			node["chainConfigs"] = []map[string]any{}
		}
	}

	return map[string]any{
		"nodes": nodes,
	}, nil
}

// loadDotEnv reads a .env file and sets any variables that are not already
// present in the process environment. It handles KEY=VALUE, KEY = VALUE,
// quoted values, and comments / blank lines.
func loadDotEnv(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		// Strip surrounding quotes if present.
		if len(value) >= 2 && ((value[0] == '"' && value[len(value)-1] == '"') ||
			(value[0] == '\'' && value[len(value)-1] == '\'')) {
			value = value[1 : len(value)-1]
		}
		// Only set if not already in environment (explicit exports take precedence).
		if os.Getenv(key) == "" {
			os.Setenv(key, value)
		}
	}
	return nil
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
