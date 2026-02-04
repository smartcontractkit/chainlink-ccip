package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync/atomic"
	"time"
)

// EVMExecutor implements ChainExecutor for EVM-compatible chains.
// It uses raw eth_call JSON-RPC requests (no bindings needed).
type EVMExecutor struct {
	rpcURL string
	client *http.Client

	// Optional: simulate flaky RPC for testing
	simulateFlaky bool
	flakyRate     float64 // 0.0 to 1.0
	callCount     atomic.Int64
}

// NewEVMExecutor creates a new EVM executor for the given RPC URL.
func NewEVMExecutor(rpcURL string) *EVMExecutor {
	return &EVMExecutor{
		rpcURL: rpcURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// NewFlakyEVMExecutor creates an EVM executor that simulates random failures.
// flakyRate is the probability of failure (0.0 = never fail, 1.0 = always fail).
func NewFlakyEVMExecutor(rpcURL string, flakyRate float64) *EVMExecutor {
	return &EVMExecutor{
		rpcURL:        rpcURL,
		client:        &http.Client{Timeout: 30 * time.Second},
		simulateFlaky: true,
		flakyRate:     flakyRate,
	}
}

// jsonRPCRequest represents a JSON-RPC 2.0 request
type jsonRPCRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      int    `json:"id"`
}

// jsonRPCResponse represents a JSON-RPC 2.0 response
type jsonRPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result,omitempty"`
	Error   *struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

// Execute performs an eth_call to the target contract with the given calldata.
func (e *EVMExecutor) Execute(target, data []byte) ([]byte, error) {
	count := e.callCount.Add(1)

	// Simulate flaky RPC if enabled
	if e.simulateFlaky && shouldFail(count, e.flakyRate) {
		return nil, fmt.Errorf("simulated rate limit error (429 too many requests)")
	}

	// Build the eth_call request
	callObject := map[string]string{
		"to":   "0x" + hex.EncodeToString(target),
		"data": "0x" + hex.EncodeToString(data),
	}

	request := jsonRPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_call",
		Params:  []any{callObject, "latest"},
		ID:      1,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := e.client.Post(e.rpcURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	// Check for HTTP-level errors
	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("http error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	// Parse response
	var rpcResp jsonRPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Check for RPC-level errors
	if rpcResp.Error != nil {
		return nil, fmt.Errorf("rpc error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}

	// Decode the hex result
	result := strings.TrimPrefix(rpcResp.Result, "0x")
	if result == "" {
		return []byte{}, nil
	}

	decoded, err := hex.DecodeString(result)
	if err != nil {
		return nil, fmt.Errorf("failed to decode result hex: %w", err)
	}

	return decoded, nil
}

// shouldFail determines if a call should fail based on the flaky rate.
// Uses a simple deterministic pattern for reproducibility in tests.
func shouldFail(callNum int64, rate float64) bool {
	// Use a deterministic pattern: fail every N calls based on rate
	if rate <= 0 {
		return false
	}
	if rate >= 1 {
		return true
	}
	// e.g., rate=0.2 means fail every 5th call
	interval := int64(1 / rate)
	return callNum%interval == 0
}

// MockExecutor is a mock executor for testing without real RPC calls.
type MockExecutor struct {
	// Responses maps cache keys to predefined responses
	responses map[string][]byte
	// Latency simulates network delay
	latency time.Duration
	// CallCount tracks how many times Execute was called
	CallCount atomic.Int64
}

// NewMockExecutor creates a new mock executor with the given latency.
func NewMockExecutor(latency time.Duration) *MockExecutor {
	return &MockExecutor{
		responses: make(map[string][]byte),
		latency:   latency,
	}
}

// SetResponse sets a predefined response for a given target+data combination.
func (m *MockExecutor) SetResponse(target, data, response []byte) {
	key := hex.EncodeToString(target) + ":" + hex.EncodeToString(data)
	m.responses[key] = response
}

// Execute returns the predefined response or a default response.
func (m *MockExecutor) Execute(target, data []byte) ([]byte, error) {
	m.CallCount.Add(1)

	// Simulate network latency
	if m.latency > 0 {
		time.Sleep(m.latency)
	}

	key := hex.EncodeToString(target) + ":" + hex.EncodeToString(data)
	if resp, ok := m.responses[key]; ok {
		return resp, nil
	}

	// Return a default response (32 bytes of zeros - common for uint256 return)
	return make([]byte, 32), nil
}

// FlakyMockExecutor simulates an unreliable RPC endpoint.
type FlakyMockExecutor struct {
	*MockExecutor
	failureRate float64
	callCount   atomic.Int64
}

// NewFlakyMockExecutor creates a mock executor that fails some percentage of calls.
func NewFlakyMockExecutor(latency time.Duration, failureRate float64) *FlakyMockExecutor {
	return &FlakyMockExecutor{
		MockExecutor: NewMockExecutor(latency),
		failureRate:  failureRate,
	}
}

// Execute returns an error based on the failure rate, otherwise delegates to MockExecutor.
func (f *FlakyMockExecutor) Execute(target, data []byte) ([]byte, error) {
	count := f.callCount.Add(1)

	if shouldFail(count, f.failureRate) {
		return nil, fmt.Errorf("simulated rate limit (429 too many requests)")
	}

	return f.MockExecutor.Execute(target, data)
}

// =====================================================
// Solana Executor
// =====================================================

// SolanaExecutor implements ChainExecutor for Solana chains.
// It uses getAccountInfo to read account data.
type SolanaExecutor struct {
	rpcURL string
	client *http.Client
}

// NewSolanaExecutor creates a new Solana executor for the given RPC URL.
func NewSolanaExecutor(rpcURL string) *SolanaExecutor {
	return &SolanaExecutor{
		rpcURL: rpcURL,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// solanaRPCResponse represents a Solana JSON-RPC response for getAccountInfo
type solanaRPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  *struct {
		Context struct {
			Slot uint64 `json:"slot"`
		} `json:"context"`
		Value *struct {
			Data       []string `json:"data"` // [base64_data, encoding]
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

// Execute performs a getAccountInfo call to read Solana account data.
// For Solana, "target" is the account address (32 bytes or base58 string in data).
// The "data" parameter can contain additional options (unused for now).
func (s *SolanaExecutor) Execute(target, data []byte) ([]byte, error) {
	// Convert target to base58 if it's raw bytes
	// For now, assume target is already base58 encoded as a string in data
	// or we convert the 32-byte address to base58
	var accountAddr string
	if len(target) == 32 {
		accountAddr = base58Encode(target)
	} else {
		accountAddr = string(target)
	}

	request := jsonRPCRequest{
		JSONRPC: "2.0",
		Method:  "getAccountInfo",
		Params: []any{
			accountAddr,
			map[string]string{
				"encoding": "base64",
			},
		},
		ID: 1,
	}

	body, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := s.client.Post(s.rpcURL, "application/json", bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("http error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	var rpcResp solanaRPCResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	if rpcResp.Error != nil {
		return nil, fmt.Errorf("rpc error %d: %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}

	if rpcResp.Result == nil || rpcResp.Result.Value == nil {
		return nil, fmt.Errorf("account not found")
	}

	// Return the account info as JSON (for now)
	// Views can parse this to extract specific fields
	result := map[string]any{
		"owner":      rpcResp.Result.Value.Owner,
		"lamports":   rpcResp.Result.Value.Lamports,
		"executable": rpcResp.Result.Value.Executable,
		"rentEpoch":  rpcResp.Result.Value.RentEpoch,
	}

	// Include raw data if present
	if len(rpcResp.Result.Value.Data) > 0 {
		result["data"] = rpcResp.Result.Value.Data[0] // base64 encoded
		if len(rpcResp.Result.Value.Data) > 1 {
			result["encoding"] = rpcResp.Result.Value.Data[1]
		}
	}

	return json.Marshal(result)
}

// base58Encode encodes bytes to base58 (Solana address format)
// This is a simplified implementation - production should use a proper base58 library
func base58Encode(input []byte) string {
	const alphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
	
	// Count leading zeros
	zeros := 0
	for _, b := range input {
		if b == 0 {
			zeros++
		} else {
			break
		}
	}

	// Allocate enough space
	size := len(input)*138/100 + 1
	buf := make([]byte, size)

	// Process the bytes
	high := size - 1
	for _, b := range input {
		carry := int(b)
		for j := size - 1; j > high || carry != 0; j-- {
			carry += 256 * int(buf[j])
			buf[j] = byte(carry % 58)
			carry /= 58
			if j <= high {
				high = j - 1
			}
		}
	}

	// Skip leading zeros in buf
	j := 0
	for j < size && buf[j] == 0 {
		j++
	}

	// Build result
	result := make([]byte, zeros+size-j)
	for i := 0; i < zeros; i++ {
		result[i] = '1'
	}
	for i := zeros; j < size; i, j = i+1, j+1 {
		result[i] = alphabet[buf[j]]
	}

	return string(result)
}

// =====================================================
// Aptos Executor
// =====================================================

// AptosExecutor implements ChainExecutor for Aptos chains.
// It uses the Aptos REST API to fetch account resources.
type AptosExecutor struct {
	rpcURL string
	client *http.Client
}

// NewAptosExecutor creates a new Aptos executor for the given RPC URL.
func NewAptosExecutor(rpcURL string) *AptosExecutor {
	return &AptosExecutor{
		rpcURL: strings.TrimSuffix(rpcURL, "/"),
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// aptosResourceResponse represents an Aptos resource response
type aptosResourceResponse struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

// AptosCallType indicates the type of Aptos API call
type AptosCallType byte

const (
	// AptosCallResources fetches all resources for an account (GET /v1/accounts/{addr}/resources)
	AptosCallResources AptosCallType = 0
	// AptosCallResource fetches a specific resource (GET /v1/accounts/{addr}/resource/{type})
	AptosCallResource AptosCallType = 1
	// AptosCallView executes a view function (POST /v1/view)
	AptosCallView AptosCallType = 2
)

// AptosViewCall represents a view function call payload
type AptosViewCall struct {
	Function      string   `json:"function"`       // e.g., "0x1::module::function"
	TypeArguments []string `json:"type_arguments"` // Generic type args
	Arguments     []any    `json:"arguments"`      // Function arguments
}

// Execute performs an Aptos REST API call.
// For Aptos, "target" is the account address (hex string as bytes).
// The "data" parameter format depends on the call type:
//   - First byte is AptosCallType
//   - For AptosCallResources (0): no additional data needed
//   - For AptosCallResource (1): rest of data is the resource type string
//   - For AptosCallView (2): rest of data is JSON-encoded AptosViewCall
func (a *AptosExecutor) Execute(target, data []byte) ([]byte, error) {
	// Convert target to address string
	accountAddr := string(target)
	if !strings.HasPrefix(accountAddr, "0x") {
		accountAddr = "0x" + accountAddr
	}

	// Determine call type from first byte of data
	callType := AptosCallResources
	callData := data
	if len(data) > 0 {
		callType = AptosCallType(data[0])
		callData = data[1:]
	}

	switch callType {
	case AptosCallResources:
		return a.fetchResources(accountAddr)
	case AptosCallResource:
		return a.fetchResource(accountAddr, string(callData))
	case AptosCallView:
		return a.executeView(callData)
	default:
		return nil, fmt.Errorf("unknown Aptos call type: %d", callType)
	}
}

// fetchResources gets all resources for an account
func (a *AptosExecutor) fetchResources(accountAddr string) ([]byte, error) {
	url := fmt.Sprintf("%s/v1/accounts/%s/resources", a.rpcURL, accountAddr)
	return a.doGet(url)
}

// fetchResource gets a specific resource for an account
func (a *AptosExecutor) fetchResource(accountAddr, resourceType string) ([]byte, error) {
	url := fmt.Sprintf("%s/v1/accounts/%s/resource/%s", a.rpcURL, accountAddr, resourceType)
	return a.doGet(url)
}

// executeView executes a view function via POST /v1/view
func (a *AptosExecutor) executeView(callData []byte) ([]byte, error) {
	url := fmt.Sprintf("%s/v1/view", a.rpcURL)

	req, err := http.NewRequest("POST", url, bytes.NewReader(callData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("view call failed (status %d): %s", resp.StatusCode, string(bodyBytes))
	}

	return bodyBytes, nil
}

// doGet performs a GET request and returns the response body
func (a *AptosExecutor) doGet(url string) ([]byte, error) {
	resp, err := a.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http request failed: %w", err)
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return nil, fmt.Errorf("account or resource not found")
		}
		return nil, fmt.Errorf("http error %d: %s", resp.StatusCode, string(bodyBytes))
	}

	return bodyBytes, nil
}
