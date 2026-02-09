package evm

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"give-me-state-v2/orchestrator"
)

// EVMOrchestrator implements orchestrator.TypedOrchestratorInterface for EVM chains:
// cache and dedup per chain selector, then delegate to the generic engine for eth_call.
type EVMOrchestrator struct {
	generic   *orchestrator.Generic
	orcIDs    map[uint64]string // chainID -> registered orchestrator ID
	caches    map[uint64]*orchestrator.CacheDedup
	cachesMu  sync.RWMutex
	retryable []string
}

// ChainEndpoints defines the endpoints (RPC URLs + config) for one chain.
type ChainEndpoints struct {
	ChainID   uint64
	Endpoints []orchestrator.EndpointConfig
}

// NewEVMOrchestrator creates an EVM typed orchestrator and registers each chain with the generic engine.
// retryableKeywords: error messages containing any of these are retried/failed over (e.g. "rate limit", "429", "timeout").
func NewEVMOrchestrator(generic *orchestrator.Generic, chainEndpoints []ChainEndpoints, retryableKeywords []string) (*EVMOrchestrator, error) {
	if generic == nil {
		return nil, fmt.Errorf("generic orchestrator is required")
	}
	e := &EVMOrchestrator{
		generic:   generic,
		orcIDs:    make(map[uint64]string),
		caches:    make(map[uint64]*orchestrator.CacheDedup),
		retryable: retryableKeywords,
	}
	for _, ce := range chainEndpoints {
		id := fmt.Sprintf("evm-%d", ce.ChainID)
		if err := generic.RegisterOrchestrator(id, ce.Endpoints, retryableKeywords); err != nil {
			return nil, fmt.Errorf("register chain %d: %w", ce.ChainID, err)
		}
		e.orcIDs[ce.ChainID] = id
		e.caches[ce.ChainID] = orchestrator.NewCacheDedup()
	}
	return e, nil
}

// Execute implements orchestrator.TypedOrchestratorInterface.
func (e *EVMOrchestrator) Execute(call orchestrator.Call) orchestrator.CallResult {
	orcID, ok := e.orcIDs[call.ChainID]
	if !ok {
		return orchestrator.CallResult{Error: fmt.Errorf("evm: no endpoints for chain %d", call.ChainID)}
	}
	if len(call.Target) == 0 {
		return orchestrator.CallResult{Error: fmt.Errorf("evm: call target required")}
	}

	e.cachesMu.RLock()
	cache := e.caches[call.ChainID]
	e.cachesMu.RUnlock()
	if cache == nil {
		return orchestrator.CallResult{Error: fmt.Errorf("evm: no cache for chain %d", call.ChainID)}
	}

	key := orchestrator.KeyFromTargetAndData(call.Target, call.Data)
	return cache.GetOrRun(key, func() orchestrator.CallResult {
		return e.doRequest(orcID, call)
	})
}

func (e *EVMOrchestrator) doRequest(orcID string, call orchestrator.Call) orchestrator.CallResult {
	body, err := e.buildEthCallBody(call)
	if err != nil {
		return orchestrator.CallResult{Error: err}
	}
	req := orchestrator.Request{Body: body, Method: "POST"}
	ctx := context.Background()
	res := e.generic.DoRequest(ctx, orcID, req)
	if res.Error != nil {
		return res
	}
	// Parse JSON-RPC response and decode result hex to bytes
	data, err := e.parseEthCallResponse(res.Data)
	if err != nil {
		return orchestrator.CallResult{Error: err, Retries: res.Retries}
	}
	return orchestrator.CallResult{Data: data, Retries: res.Retries}
}

func (e *EVMOrchestrator) buildEthCallBody(call orchestrator.Call) ([]byte, error) {
	callObject := map[string]string{
		"to":   "0x" + hex.EncodeToString(call.Target),
		"data": "0x" + hex.EncodeToString(call.Data),
	}
	req := jsonRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_call",
		Params:  []any{callObject, "latest"},
		ID:      1,
	}
	return json.Marshal(req)
}

func (e *EVMOrchestrator) parseEthCallResponse(body []byte) ([]byte, error) {
	var rpcResp jsonRPCResponse
	if err := json.Unmarshal(body, &rpcResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	if rpcResp.Error != nil {
		return nil, fmt.Errorf("rpc error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}
	result := strings.TrimPrefix(rpcResp.Result, "0x")
	if result == "" {
		return []byte{}, nil
	}
	return hex.DecodeString(result)
}

type jsonRPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      int    `json:"id"`
}

type jsonRPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result,omitempty"`
	Error   *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}
