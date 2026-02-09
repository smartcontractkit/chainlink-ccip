package svm

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/mr-tron/base58"

	"give-me-state-v2/orchestrator"
)

// SVMOrchestrator implements orchestrator.TypedOrchestratorInterface for Solana (SVM) chains:
// cache and dedup per chain, then delegate to the generic engine for getAccountInfo JSON-RPC.
type SVMOrchestrator struct {
	generic   *orchestrator.Generic
	orcIDs    map[uint64]string
	caches    map[uint64]*orchestrator.CacheDedup
	cachesMu  sync.RWMutex
	retryable []string
}

// ChainEndpoints defines the endpoints (RPC URLs + config) for one Solana chain.
type ChainEndpoints struct {
	ChainID   uint64
	Endpoints []orchestrator.EndpointConfig
}

// NewSVMOrchestrator creates a Solana typed orchestrator and registers each chain with the generic engine.
func NewSVMOrchestrator(generic *orchestrator.Generic, chainEndpoints []ChainEndpoints, retryableKeywords []string) (*SVMOrchestrator, error) {
	if generic == nil {
		return nil, fmt.Errorf("generic orchestrator is required")
	}
	s := &SVMOrchestrator{
		generic:   generic,
		orcIDs:    make(map[uint64]string),
		caches:    make(map[uint64]*orchestrator.CacheDedup),
		retryable: retryableKeywords,
	}
	for _, ce := range chainEndpoints {
		id := fmt.Sprintf("svm-%d", ce.ChainID)
		if err := generic.RegisterOrchestrator(id, ce.Endpoints, retryableKeywords); err != nil {
			return nil, fmt.Errorf("register chain %d: %w", ce.ChainID, err)
		}
		s.orcIDs[ce.ChainID] = id
		s.caches[ce.ChainID] = orchestrator.NewCacheDedup()
	}
	return s, nil
}

// Execute implements orchestrator.TypedOrchestratorInterface.
// Target is either 32-byte address (converted to base58) or base58 account address as string bytes.
func (s *SVMOrchestrator) Execute(call orchestrator.Call) orchestrator.CallResult {
	orcID, ok := s.orcIDs[call.ChainID]
	if !ok {
		return orchestrator.CallResult{Error: fmt.Errorf("svm: no endpoints for chain %d", call.ChainID)}
	}
	if len(call.Target) == 0 {
		return orchestrator.CallResult{Error: fmt.Errorf("svm: call target required")}
	}

	s.cachesMu.RLock()
	cache := s.caches[call.ChainID]
	s.cachesMu.RUnlock()
	if cache == nil {
		return orchestrator.CallResult{Error: fmt.Errorf("svm: no cache for chain %d", call.ChainID)}
	}

	key := orchestrator.KeyFromTargetAndData(call.Target, call.Data)
	return cache.GetOrRun(key, func() orchestrator.CallResult {
		return s.doRequest(orcID, call)
	})
}

func (s *SVMOrchestrator) doRequest(orcID string, call orchestrator.Call) orchestrator.CallResult {
	body, err := s.buildGetAccountInfoBody(call)
	if err != nil {
		return orchestrator.CallResult{Error: err}
	}
	req := orchestrator.Request{Body: body, Method: "POST"}
	ctx := context.Background()
	res := s.generic.DoRequest(ctx, orcID, req)
	if res.Error != nil {
		return res
	}
	data, err := s.parseGetAccountInfoResponse(res.Data)
	if err != nil {
		return orchestrator.CallResult{Error: err, Retries: res.Retries}
	}
	return orchestrator.CallResult{Data: data, Retries: res.Retries}
}

func (s *SVMOrchestrator) buildGetAccountInfoBody(call orchestrator.Call) ([]byte, error) {
	var accountAddr string
	if len(call.Target) == 32 {
		accountAddr = base58.Encode(call.Target)
	} else {
		accountAddr = string(call.Target)
	}
	req := jsonRPCRequest{
		JSONRPC: "2.0",
		Method:  "getAccountInfo",
		Params: []any{
			accountAddr,
			map[string]string{"encoding": "base64"},
		},
		ID: 1,
	}
	return json.Marshal(req)
}

func (s *SVMOrchestrator) parseGetAccountInfoResponse(body []byte) ([]byte, error) {
	var rpcResp solanaRPCResponse
	if err := json.Unmarshal(body, &rpcResp); err != nil {
		return nil, fmt.Errorf("decode response: %w", err)
	}
	if rpcResp.Error != nil {
		return nil, fmt.Errorf("rpc error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}
	if rpcResp.Result == nil || rpcResp.Result.Value == nil {
		return nil, fmt.Errorf("account not found")
	}
	v := rpcResp.Result.Value
	result := map[string]any{
		"owner":      v.Owner,
		"lamports":   v.Lamports,
		"executable": v.Executable,
		"rentEpoch":  v.RentEpoch,
	}
	if len(v.Data) > 0 {
		result["data"] = v.Data[0]
		if len(v.Data) > 1 {
			result["encoding"] = v.Data[1]
		}
	}
	return json.Marshal(result)
}

type jsonRPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      int    `json:"id"`
}

type solanaRPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  *struct {
		Value *struct {
			Data       []string `json:"data"`
			Executable bool     `json:"executable"`
			Lamports   uint64   `json:"lamports"`
			Owner      string   `json:"owner"`
			RentEpoch  uint64   `json:"rentEpoch"`
		} `json:"value"`
	} `json:"result"`
	Error *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}
