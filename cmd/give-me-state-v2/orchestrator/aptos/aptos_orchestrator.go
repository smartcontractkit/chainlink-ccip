package aptos

import (
	"context"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"

	"give-me-state-v2/orchestrator"
)

// AptosCallType indicates the type of Aptos API call (must match v1 views).
type AptosCallType byte

const (
	AptosCallResources AptosCallType = 0 // GET /v1/accounts/{addr}/resources
	AptosCallResource  AptosCallType = 1 // GET /v1/accounts/{addr}/resource/{type}
	AptosCallView       AptosCallType = 2 // POST /v1/view
)

// AptosOrchestrator implements orchestrator.TypedOrchestratorInterface for Aptos chains:
// cache and dedup per chain, then delegate to the generic engine (GET or POST with FullURL).
type AptosOrchestrator struct {
	generic    *orchestrator.Generic
	orcIDs     map[uint64]string
	baseURLs   map[uint64]string // chainID -> base URL (no trailing slash)
	caches     map[uint64]*orchestrator.CacheDedup
	cachesMu   sync.RWMutex
	retryable  []string
}

// ChainEndpoints defines the endpoints (base URLs + config) for one Aptos chain.
type ChainEndpoints struct {
	ChainID   uint64
	Endpoints []orchestrator.EndpointConfig
}

// NewAptosOrchestrator creates an Aptos typed orchestrator and registers each chain.
// Base URL per chain is taken from the first endpoint (trimmed).
func NewAptosOrchestrator(generic *orchestrator.Generic, chainEndpoints []ChainEndpoints, retryableKeywords []string) (*AptosOrchestrator, error) {
	if generic == nil {
		return nil, fmt.Errorf("generic orchestrator is required")
	}
	a := &AptosOrchestrator{
		generic:   generic,
		orcIDs:    make(map[uint64]string),
		baseURLs:  make(map[uint64]string),
		caches:    make(map[uint64]*orchestrator.CacheDedup),
		retryable: retryableKeywords,
	}
	for _, ce := range chainEndpoints {
		if len(ce.Endpoints) == 0 {
			return nil, fmt.Errorf("aptos chain %d: no endpoints", ce.ChainID)
		}
		baseURL := strings.TrimSuffix(ce.Endpoints[0].URL, "/")
		id := fmt.Sprintf("aptos-%d", ce.ChainID)
		if err := generic.RegisterOrchestrator(id, ce.Endpoints, retryableKeywords); err != nil {
			return nil, fmt.Errorf("register chain %d: %w", ce.ChainID, err)
		}
		a.orcIDs[ce.ChainID] = id
		a.baseURLs[ce.ChainID] = baseURL
		a.caches[ce.ChainID] = orchestrator.NewCacheDedup()
	}
	return a, nil
}

// Execute implements orchestrator.TypedOrchestratorInterface.
// Target is the account address (hex string as bytes). Data: first byte = AptosCallType, rest = type-specific payload.
func (a *AptosOrchestrator) Execute(call orchestrator.Call) orchestrator.CallResult {
	orcID, ok := a.orcIDs[call.ChainID]
	if !ok {
		return orchestrator.CallResult{Error: fmt.Errorf("aptos: no endpoints for chain %d", call.ChainID)}
	}
	baseURL, ok := a.baseURLs[call.ChainID]
	if !ok {
		return orchestrator.CallResult{Error: fmt.Errorf("aptos: no base URL for chain %d", call.ChainID)}
	}

	a.cachesMu.RLock()
	cache := a.caches[call.ChainID]
	a.cachesMu.RUnlock()
	if cache == nil {
		return orchestrator.CallResult{Error: fmt.Errorf("aptos: no cache for chain %d", call.ChainID)}
	}

	key := orchestrator.KeyFromTargetAndData(call.Target, call.Data)
	return cache.GetOrRun(key, func() orchestrator.CallResult {
		return a.doRequest(orcID, baseURL, call)
	})
}

func (a *AptosOrchestrator) doRequest(orcID, baseURL string, call orchestrator.Call) orchestrator.CallResult {
	accountAddr := a.targetToAddr(call.Target)

	callType := AptosCallResources
	callData := call.Data
	if len(call.Data) > 0 {
		callType = AptosCallType(call.Data[0])
		callData = call.Data[1:]
	}

	var req orchestrator.Request
	switch callType {
	case AptosCallResources:
		req = orchestrator.Request{Method: "GET", FullURL: baseURL + "/v1/accounts/" + accountAddr + "/resources"}
	case AptosCallResource:
		req = orchestrator.Request{Method: "GET", FullURL: baseURL + "/v1/accounts/" + accountAddr + "/resource/" + string(callData)}
	case AptosCallView:
		req = orchestrator.Request{Method: "POST", FullURL: baseURL + "/v1/view", Body: callData}
	default:
		return orchestrator.CallResult{Error: fmt.Errorf("unknown Aptos call type: %d", callType)}
	}

	ctx := context.Background()
	res := a.generic.DoRequest(ctx, orcID, req)
	return res
}

// targetToAddr converts Call.Target to Aptos account address (0x-prefixed hex).
// Accepts either 32-byte address (encoded as hex) or hex string as bytes.
func (a *AptosOrchestrator) targetToAddr(target []byte) string {
	if len(target) == 0 {
		return "0x"
	}
	if len(target) == 32 && (target[0] >= 0x7f || target[0] < 0x20) {
		return "0x" + hex.EncodeToString(target)
	}
	s := string(target)
	if !strings.HasPrefix(s, "0x") {
		return "0x" + s
	}
	return s
}
