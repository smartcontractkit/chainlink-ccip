package orchestrator

// TypedOrchestratorInterface is the interface the views package uses.
// Typed orchestrators (EVM, Solana, etc.) implement cache/dedup and call the generic engine for requests.
type TypedOrchestratorInterface interface {
	Execute(call Call) CallResult
}

// Call is the chain-agnostic request from views (chain, target, calldata).
type Call struct {
	ChainID uint64
	Target  []byte // Contract/program address (e.g. 20 bytes EVM)
	Data    []byte
}

// CallResult is the result returned to views.
type CallResult struct {
	Data    []byte
	Error   error
	Cached  bool
	Retries int
}

// Request is what the generic engine sends over HTTP (body + method).
// Typed orchestrators convert Call into Request and call Generic.DoRequest.
type Request struct {
	Body []byte
	// Method is HTTP method; default "POST"
	Method string
	// FullURL if set is used as the request URL (for chains like Aptos with GET to different paths).
	// If empty, the endpoint's URL is used.
	FullURL string
}

// EndpointConfig is per-endpoint configuration for the generic orchestrator.
type EndpointConfig struct {
	URL     string // RPC or API base URL
	Workers int    // Worker goroutines contributed by this endpoint to the shared pool; 0 = use 1
	Timeout int    // Seconds; 0 = use default
}

// DefaultRetriesPerEndpoint is how many times we retry on one endpoint before failover.
const DefaultRetriesPerEndpoint = 3
