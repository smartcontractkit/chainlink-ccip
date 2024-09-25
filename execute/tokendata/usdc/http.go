package usdc

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

const (
	apiVersion      = "v1"
	attestationPath = "attestations"

	// defaultCoolDownDurationSec defines the default time to wait after getting rate limited.
	// this value is only used if the 429 response does not contain the Retry-After header
	defaultCoolDownDuration = 5 * time.Minute

	// maxCoolDownDuration defines the maximum duration we can wait till firing the next request
	maxCoolDownDuration = 10 * time.Minute
)

type HTTPClient interface {
	// Get calls the USDC attestation API with the given USDC message hash.
	// The attestation service rate limit is 10 requests per second. If you exceed 10 requests
	// per second, the service blocks all API requests for the next 5 minutes and returns an
	// HTTP 429 response.
	//
	// Documentation:
	//
	//	https://developers.circle.com/stablecoins/reference/getattestation
	//	https://developers.circle.com/stablecoins/docs/transfer-usdc-on-testnet-from-ethereum-to-avalanche
	Get(ctx context.Context, messageHash [32]byte) ([]byte, int, error)
}

// httpClient is a client for the USDC attestation API. It encapsulates all the details specific to the Attestation API:
// - rate limiting
// - cool down period
// - parsing JSON response and handling errors
// Therefore AttestationClient is a higher level abstraction that uses httpClient to fetch attestations and can be more
// oriented around caching/processing the attestation data instead of handling the API specifics.
type httpClient struct {
	apiURL     string
	apiTimeout time.Duration
	rate       *rate.Limiter
	// coolDownUntil defines whether requests are blocked or not.
	coolDownUntil time.Time
	coolDownMu    *sync.RWMutex
}

//nolint:revive
func NewHTTPClient(api string, apiInterval time.Duration, apiTimeout time.Duration) *httpClient {
	return &httpClient{
		apiURL:     api,
		apiTimeout: apiTimeout,
		rate:       rate.NewLimiter(rate.Every(apiInterval), 1),
		coolDownMu: &sync.RWMutex{},
	}
}

type attestationStatus string

const (
	attestationStatusSuccess attestationStatus = "complete"
	attestationStatusPending attestationStatus = "pending_confirmations"
)

type httpResponse struct {
	Status      attestationStatus `json:"status"`
	Attestation string            `json:"attestation"`
	Error       string            `json:"error"`
}

func (r httpResponse) validate() error {
	if r.Error != "" {
		return fmt.Errorf("attestation API error: %s", r.Error)
	}
	if r.Status == "" {
		return fmt.Errorf("invalid attestation response")
	}
	if r.Status != attestationStatusSuccess {
		return ErrNotReady
	}
	return nil
}

func (r httpResponse) attestationToBytes() ([]byte, error) {
	attestationBytes, err := hex.DecodeString(strings.TrimPrefix(r.Attestation, "0x"))
	if err != nil {
		return nil, err
	}
	return attestationBytes, nil
}

func (h *httpClient) Get(ctx context.Context, messageHash [32]byte) ([]byte, int, error) {
	var empty []byte
	// Terminate immediately when rate limited
	if h.inCoolDownPeriod() {
		return empty, http.StatusTooManyRequests, ErrRateLimit
	}

	if h.rate != nil {
		// Wait blocks until it the attestation API can be called or the
		// context is Done.
		if waitErr := h.rate.Wait(ctx); waitErr != nil {
			return empty, http.StatusTooManyRequests, ErrRateLimit
		}
	}

	timeoutCtx, cancel := context.WithTimeoutCause(ctx, h.apiTimeout, ErrTimeout)
	defer cancel()

	url := fmt.Sprintf("%s/%s/%s/0x%x", h.apiURL, apiVersion, attestationPath, messageHash)
	return h.callAPI(timeoutCtx, url)
}

func (h *httpClient) callAPI(ctx context.Context, url string) ([]byte, int, error) {
	var response []byte
	// Use a timeout to guard against attestation API hanging, causing observation timeout and
	// failing to make any progress.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return response, http.StatusBadRequest, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return response, http.StatusRequestTimeout, ErrTimeout
		}
		// On error, res is nil in most cases, do not read res.StatusCode, return BadRequest
		return response, http.StatusBadRequest, err
	}
	defer res.Body.Close()

	// Explicitly signal if the API is being rate limited
	if res.StatusCode == http.StatusTooManyRequests {
		h.setCoolDownPeriod(res.Header)
		return response, res.StatusCode, ErrRateLimit
	}
	if res.StatusCode == http.StatusNotFound {
		return response, res.StatusCode, ErrNotReady
	}
	if res.StatusCode != http.StatusOK {
		return response, res.StatusCode, ErrUnknownResponse
	}

	return h.parsePayload(res)
}

func (h *httpClient) parsePayload(res *http.Response) ([]byte, int, error) {
	var response httpResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, res.StatusCode, fmt.Errorf("failed to decode json: %w", err)
	}

	if err1 := response.validate(); err1 != nil {
		return nil, res.StatusCode, err1
	}

	attestationBytes, err := response.attestationToBytes()
	return attestationBytes, res.StatusCode, err
}

func (h *httpClient) setCoolDownPeriod(headers http.Header) {
	coolDownDuration := defaultCoolDownDuration
	if retryAfterHeader, exists := headers["Retry-After"]; exists && len(retryAfterHeader) > 0 {
		if retryAfterSec, errParseInt := strconv.ParseInt(retryAfterHeader[0], 10, 64); errParseInt == nil {
			coolDownDuration = time.Duration(retryAfterSec) * time.Second
		}
	}

	h.coolDownMu.Lock()
	defer h.coolDownMu.Unlock()
	coolDownDuration = min(coolDownDuration, maxCoolDownDuration)
	h.coolDownUntil = time.Now().Add(coolDownDuration)
}

func (h *httpClient) inCoolDownPeriod() bool {
	h.coolDownMu.RLock()
	defer h.coolDownMu.RUnlock()
	return time.Now().Before(h.coolDownUntil)
}
