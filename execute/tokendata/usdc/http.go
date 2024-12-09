package usdc

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"sync"
	"time"

	"golang.org/x/time/rate"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
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

type HTTPStatus int

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
	Get(ctx context.Context, messageHash cciptypes.Bytes32) (cciptypes.Bytes, HTTPStatus, error)
}

// httpClient is a client for the USDC attestation API. It encapsulates all the details specific to the Attestation API:
// - rate limiting
// - cool down period
// - parsing JSON response and handling errors
// Therefore AttestationClient is a higher level abstraction that uses httpClient to fetch attestations and can be more
// oriented around caching/processing the attestation data instead of handling the API specifics.
type httpClient struct {
	lggr       logger.Logger
	apiURL     *url.URL
	apiTimeout time.Duration
	rate       *rate.Limiter
	// coolDownUntil defines whether requests are blocked or not.
	coolDownUntil time.Time
	coolDownMu    *sync.RWMutex
}

var (
	clientInstances = make(map[string]*httpClient)
	mutex           sync.Mutex
)

// GetHTTPClient returns a singleton instance of the httpClient for the given API URL.
// It's critical to reuse existing clients because of the self-rate limiting mechanism. Being rate limited by
// Circle comes with a long cool down period, so we should always self-rate limit before hitting the API rate limit.
// IMPORTANT: In the loop world this might require major rework - e.g. making httpClient a loop plugin to
// enforce the singleton pattern.
func GetHTTPClient(
	lggr logger.Logger,
	api string,
	apiInterval time.Duration,
	apiTimeout time.Duration,
) (HTTPClient, error) {
	mutex.Lock()
	defer mutex.Unlock()

	if client, exists := clientInstances[api]; exists {
		return client, nil
	}

	client, err := newHTTPClient(lggr, api, apiInterval, apiTimeout)
	if err != nil {
		return nil, err
	}

	clientInstances[api] = client.(*httpClient)
	return client, nil
}

func newHTTPClient(
	lggr logger.Logger,
	api string,
	apiInterval time.Duration,
	apiTimeout time.Duration,
) (HTTPClient, error) {
	u, err := url.ParseRequestURI(api)
	if err != nil {
		return nil, err
	}

	return &httpClient{
		lggr:       lggr,
		apiURL:     u,
		apiTimeout: apiTimeout,
		rate:       rate.NewLimiter(rate.Every(apiInterval), 1),
		coolDownMu: &sync.RWMutex{},
	}, nil
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

func (r httpResponse) attestationToBytes() (cciptypes.Bytes, error) {
	attestationBytes, err := cciptypes.NewBytesFromString(r.Attestation)
	if err != nil {
		return nil, fmt.Errorf("failed to decode attestation hex: %w", err)
	}
	return attestationBytes, nil
}

func (h *httpClient) Get(ctx context.Context, messageHash cciptypes.Bytes32) (cciptypes.Bytes, HTTPStatus, error) {
	// Terminate immediately when rate limited
	if coolDown, duration := h.inCoolDownPeriod(); coolDown {
		h.lggr.Errorw(
			"Rate limited by the Attestation API, dropping all requests",
			"coolDownDuration", duration,
		)
		return nil, http.StatusTooManyRequests, ErrRateLimit
	}

	if h.rate != nil {
		// Wait blocks until it the attestation API can be called or the
		// context is Done.
		if waitErr := h.rate.Wait(ctx); waitErr != nil {
			h.lggr.Warnw("Self rate-limited, sending too many requests to the Attestation API")
			return nil, http.StatusTooManyRequests, ErrRateLimit
		}
	}

	// Use a timeout to guard against attestation API hanging, causing observation timeout and
	// failing to make any progress.
	timeoutCtx, cancel := context.WithTimeoutCause(ctx, h.apiTimeout, ErrTimeout)
	defer cancel()

	requestURL := *h.apiURL
	requestURL.Path = path.Join(requestURL.Path, apiVersion, attestationPath, messageHash.String())

	response, httpStatus, err := h.callAPI(timeoutCtx, requestURL)
	h.lggr.Debugw(
		"Response from attestation API",
		"messageHash", messageHash,
		"status", httpStatus,
		"err", err,
	)
	return response, httpStatus, err
}

func (h *httpClient) callAPI(ctx context.Context, url url.URL) (cciptypes.Bytes, HTTPStatus, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url.String(), nil)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	req.Header.Add("accept", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctx.Err() != nil && errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return nil, http.StatusRequestTimeout, ErrTimeout
		} else if errors.Is(err, ErrTimeout) {
			return nil, http.StatusRequestTimeout, ErrTimeout
		}
		// On error, res is nil in most cases, do not read res.StatusCode, return BadRequest
		return nil, http.StatusBadRequest, err
	}

	var status HTTPStatus
	if res != nil {
		defer res.Body.Close()
		status = HTTPStatus(res.StatusCode)
	}
	// Explicitly signal if the API is being rate limited
	if res.StatusCode == http.StatusTooManyRequests {
		h.setCoolDownPeriod(res.Header)
		return nil, status, ErrRateLimit
	}
	if res.StatusCode == http.StatusNotFound {
		return nil, status, ErrNotReady
	}
	if res.StatusCode != http.StatusOK {
		return nil, status, ErrUnknownResponse
	}

	response, err := h.parsePayload(res)
	return response, status, err
}

func (h *httpClient) parsePayload(res *http.Response) (cciptypes.Bytes, error) {
	if res == nil {
		return nil, ErrUnknownResponse
	}

	var response httpResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode json: %w", err)
	}

	if err := response.validate(); err != nil {
		return nil, err
	}

	return response.attestationToBytes()
}

func (h *httpClient) setCoolDownPeriod(headers http.Header) {
	coolDownDuration := defaultCoolDownDuration
	if retryAfterHeader, exists := headers["Retry-After"]; exists && len(retryAfterHeader) > 0 {
		retryAfterSec, errParseInt := strconv.ParseInt(retryAfterHeader[0], 10, 64)
		if errParseInt == nil {
			coolDownDuration = time.Duration(retryAfterSec) * time.Second
		} else {
			parsedTime, err := time.Parse(time.RFC1123, retryAfterHeader[0])
			if err == nil {
				coolDownDuration = time.Until(parsedTime)
			}
		}
	}

	coolDownDuration = min(coolDownDuration, maxCoolDownDuration)
	//Logging on the error level, because we should always self-rate limit before hitting the API rate limit
	h.lggr.Errorw(
		"Rate limited by the Attestation API, setting cool down",
		"coolDownDuration", coolDownDuration,
	)

	h.coolDownMu.Lock()
	defer h.coolDownMu.Unlock()
	h.coolDownUntil = time.Now().Add(coolDownDuration)
}

func (h *httpClient) inCoolDownPeriod() (bool, time.Duration) {
	h.coolDownMu.RLock()
	defer h.coolDownMu.RUnlock()
	return time.Now().Before(h.coolDownUntil), time.Until(h.coolDownUntil)
}
