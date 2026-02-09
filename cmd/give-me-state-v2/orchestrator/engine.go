package orchestrator

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"
)

const (
	defaultHTTPTimeout = 30 * time.Second
	retryBaseDelay     = 300 * time.Millisecond
	retryMaxDelay      = 5 * time.Second
	rateLimitCooldown  = 2 * time.Second
)

// Generic is the shared engine: one queue per registered orchestrator,
// a fixed worker pool, and Laplace-smoothed endpoint selection per request.
type Generic struct {
	mu            sync.RWMutex
	orchestrators map[string]*orcState
	client        *http.Client
}

type orcState struct {
	id                 string
	queue              chan *workItem
	endpoints          []*endpointState
	retryableKeywords  []string
	retriesPerEndpoint int
	workers            int // total fixed worker count
}

type endpointState struct {
	url     string
	timeout time.Duration

	mu            sync.Mutex
	successes     int64
	failures      int64
	cooldownUntil time.Time // set on rate-limit; pickEndpoint skips until this passes
}

type workItem struct {
	req        Request
	resultChan chan CallResult
	tried      map[string]int // endpoint url -> attempt count
	mu         sync.Mutex
}

// NewGeneric creates a new generic orchestrator engine.
func NewGeneric() *Generic {
	return &Generic{
		orchestrators: make(map[string]*orcState),
		client: &http.Client{
			Timeout: defaultHTTPTimeout,
		},
	}
}

// laplaceScore returns (successes+1)/(total+2), biased toward 0.5 for new endpoints.
func laplaceScore(es *endpointState) float64 {
	es.mu.Lock()
	s, f := es.successes, es.failures
	es.mu.Unlock()
	return float64(s+1) / float64(s+f+2)
}

// isCoolingDown returns true if the endpoint is in a rate-limit cooldown period.
func isCoolingDown(es *endpointState) bool {
	es.mu.Lock()
	cd := es.cooldownUntil
	es.mu.Unlock()
	return time.Now().Before(cd)
}

// setCooldown marks an endpoint as rate-limited for the cooldown duration.
func setCooldown(es *endpointState) {
	es.mu.Lock()
	es.cooldownUntil = time.Now().Add(rateLimitCooldown)
	es.mu.Unlock()
}

// isRateLimitError returns true if the error looks like a rate-limit response.
func isRateLimitError(err error) bool {
	msg := err.Error()
	return strings.Contains(msg, "429") ||
		strings.Contains(msg, "rate limit") ||
		strings.Contains(msg, "too many requests")
}

// pickEndpoint returns the healthiest, non-cooled-down endpoint that hasn't been
// exhausted for this item. Returns nil if all endpoints are exhausted.
func pickEndpoint(orc *orcState, item *workItem) *endpointState {
	type scored struct {
		es    *endpointState
		score float64
	}
	candidates := make([]scored, 0, len(orc.endpoints))
	for _, es := range orc.endpoints {
		candidates = append(candidates, scored{es: es, score: laplaceScore(es)})
	}
	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].score > candidates[j].score
	})

	now := time.Now()
	item.mu.Lock()
	defer item.mu.Unlock()

	// First pass: prefer endpoints that are not cooling down
	for _, c := range candidates {
		if item.tried[c.es.url] < orc.retriesPerEndpoint {
			c.es.mu.Lock()
			cooled := now.Before(c.es.cooldownUntil)
			c.es.mu.Unlock()
			if !cooled {
				return c.es
			}
		}
	}
	// Second pass: allow cooled-down endpoints if they still have retries
	// (better than returning nil when all are rate-limited)
	for _, c := range candidates {
		if item.tried[c.es.url] < orc.retriesPerEndpoint {
			return c.es
		}
	}
	return nil // all exhausted
}

// RegisterOrchestrator registers a typed orchestrator with its endpoints and retryable keywords.
func (g *Generic) RegisterOrchestrator(id string, endpoints []EndpointConfig, retryableKeywords []string) error {
	if id == "" {
		return fmt.Errorf("orchestrator id is required")
	}
	if len(endpoints) == 0 {
		return fmt.Errorf("at least one endpoint is required")
	}

	g.mu.Lock()
	defer g.mu.Unlock()
	if _, exists := g.orchestrators[id]; exists {
		return fmt.Errorf("orchestrator %q already registered", id)
	}

	queueCap := 10000
	orc := &orcState{
		id:                 id,
		queue:              make(chan *workItem, queueCap),
		retryableKeywords:  retryableKeywords,
		retriesPerEndpoint: DefaultRetriesPerEndpoint,
	}

	// Sum up MaxConcurrent across all endpoints to get total worker count.
	totalWorkers := 0
	for i := range endpoints {
		cfg := &endpoints[i]
		if cfg.MaxConcurrent < 1 {
			cfg.MaxConcurrent = 1
		}
		totalWorkers += cfg.MaxConcurrent

		to := defaultHTTPTimeout
		if cfg.Timeout > 0 {
			to = time.Duration(cfg.Timeout) * time.Second
		}
		orc.endpoints = append(orc.endpoints, &endpointState{
			url:     cfg.URL,
			timeout: to,
		})
	}
	orc.workers = totalWorkers

	g.orchestrators[id] = orc

	// Start fixed worker pool
	for i := 0; i < totalWorkers; i++ {
		go g.runWorker(orc)
	}

	return nil
}

// DoRequest enqueues a request for the given orchestrator and blocks until a result is returned.
func (g *Generic) DoRequest(ctx context.Context, orchestratorID string, req Request) CallResult {
	g.mu.RLock()
	orc, ok := g.orchestrators[orchestratorID]
	g.mu.RUnlock()
	if !ok {
		return CallResult{Error: fmt.Errorf("orchestrator %q not registered", orchestratorID)}
	}

	method := req.Method
	if method == "" {
		method = "POST"
	}

	item := &workItem{
		req:        Request{Body: req.Body, Method: method},
		resultChan: make(chan CallResult, 1),
		tried:      make(map[string]int),
	}

	select {
	case orc.queue <- item:
	case <-ctx.Done():
		return CallResult{Error: ctx.Err()}
	}

	select {
	case res := <-item.resultChan:
		return res
	case <-ctx.Done():
		return CallResult{Error: ctx.Err()}
	}
}

func (g *Generic) runWorker(orc *orcState) {
	for {
		item, ok := <-orc.queue
		if !ok {
			return
		}

		// Pick the healthiest endpoint with retries remaining
		es := pickEndpoint(orc, item)
		if es == nil {
			// All endpoints exhausted
			item.resultChan <- CallResult{Error: fmt.Errorf("all endpoints exhausted")}
			continue
		}

		// Execute HTTP request
		result := g.doHTTP(es, item.req)
		g.recordOutcome(es, result.Error == nil)

		if result.Error == nil {
			item.resultChan <- result
			continue
		}

		// If rate-limited, cool down this endpoint so other requests avoid it
		if isRateLimitError(result.Error) {
			setCooldown(es)
		}

		// Track the attempt
		item.mu.Lock()
		item.tried[es.url]++
		tried := item.tried[es.url]
		item.mu.Unlock()

		if !g.isRetryable(orc, result.Error) {
			item.resultChan <- result
			continue
		}

		// Check if all endpoints are now exhausted
		if pickEndpoint(orc, item) == nil {
			item.resultChan <- result
			continue
		}

		// Exponential backoff, then re-enqueue (non-blocking to free worker)
		delay := retryBaseDelay
		for i := 1; i < tried; i++ {
			delay *= 2
			if delay > retryMaxDelay {
				delay = retryMaxDelay
				break
			}
		}
		failResult := result
		go func() {
			time.Sleep(delay)
			select {
			case orc.queue <- item:
			default:
				item.resultChan <- failResult
			}
		}()
	}
}

func (g *Generic) doHTTP(es *endpointState, req Request) CallResult {
	url := es.url
	if req.FullURL != "" {
		url = req.FullURL
	}
	httpReq, err := http.NewRequest(req.Method, url, bytes.NewReader(req.Body))
	if err != nil {
		return CallResult{Error: fmt.Errorf("build request: %w", err)}
	}
	if len(req.Body) > 0 {
		httpReq.Header.Set("Content-Type", "application/json")
	}

	ctx, cancel := context.WithTimeout(context.Background(), es.timeout)
	defer cancel()
	httpReq = httpReq.WithContext(ctx)

	resp, err := g.client.Do(httpReq)
	if err != nil {
		return CallResult{Error: err}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return CallResult{Error: fmt.Errorf("http %d: %s", resp.StatusCode, string(body))}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CallResult{Error: fmt.Errorf("read body: %w", err)}
	}

	// If response is JSON-RPC with an "error" field, treat that as error
	var rpcErr struct {
		Error *struct {
			Message string `json:"message"`
		} `json:"error"`
	}
	_ = json.Unmarshal(body, &rpcErr)
	if rpcErr.Error != nil {
		return CallResult{Error: fmt.Errorf("%s", rpcErr.Error.Message)}
	}

	return CallResult{Data: body}
}

func (g *Generic) recordOutcome(es *endpointState, success bool) {
	es.mu.Lock()
	defer es.mu.Unlock()
	if success {
		es.successes++
	} else {
		es.failures++
	}
}

func (g *Generic) isRetryable(orc *orcState, err error) bool {
	if err == nil {
		return false
	}
	msg := err.Error()
	for _, kw := range orc.retryableKeywords {
		if strings.Contains(msg, kw) {
			return true
		}
	}
	return false
}

// EndpointLiveStats is the live view of one RPC endpoint (for display).
type EndpointLiveStats struct {
	URL           string  // RPC URL (may be truncated by caller)
	SuccessRate   float64 // Laplace-smoothed score (0-1)
	Workers       int     // total workers for this orchestrator (shared, not per-endpoint)
	MaxConcurrent int     // kept for display compatibility
}

// OrcLiveStats is the live view of one registered orchestrator (e.g. one chain).
type OrcLiveStats struct {
	QueueDepth int                 // current items in queue
	Endpoints  []EndpointLiveStats // per-endpoint stats
}

// LiveStats returns a snapshot of all orchestrators for live dashboards.
func (g *Generic) LiveStats() map[string]OrcLiveStats {
	g.mu.RLock()
	defer g.mu.RUnlock()
	out := make(map[string]OrcLiveStats, len(g.orchestrators))
	for id, orc := range g.orchestrators {
		eps := make([]EndpointLiveStats, 0, len(orc.endpoints))
		for _, es := range orc.endpoints {
			score := laplaceScore(es)
			eps = append(eps, EndpointLiveStats{
				URL:           es.url,
				SuccessRate:   score,
				Workers:       orc.workers,
				MaxConcurrent: orc.workers,
			})
		}
		out[id] = OrcLiveStats{
			QueueDepth: len(orc.queue),
			Endpoints:  eps,
		}
	}
	return out
}
