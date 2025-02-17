package http

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"

	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	longTimeout = 60 * time.Second
)

var (
	validAttestationResponse = []byte(`
	{
		"status": "complete",
		"attestation": "0x720502893578a89a8a87982982ef781c18b193"
	}`)
)

func Test_NewHTTPClient_New(t *testing.T) {
	tt := []struct {
		api     string
		wantErr bool
	}{
		{"http://localhost:8080", false},
		{"https://iris-api-sandbox.circle.com", false},
		{"not_an_url", true},
		{"", true},
		{"   ", true},
	}

	for _, tc := range tt {
		t.Run(tc.api, func(t *testing.T) {
			client, err := newHTTPClient(logger.Test(t), tc.api, 1*time.Millisecond, longTimeout, maxCoolDownDuration)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, client)
			}
		})

	}
}

func Test_HTTPClient_Get(t *testing.T) {
	tt := []struct {
		name               string
		getTs              func(t *testing.T) *httptest.Server
		timeout            time.Duration
		messageHash        cciptypes.Bytes32
		expectedError      error
		expectedResponse   cciptypes.Bytes
		expectedStatusCode HTTPStatus
	}{
		{
			name: "server error",
			getTs: func(t *testing.T) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				}))
			},
			timeout:            time.Hour, // not relevant to the test
			expectedError:      tokendata.ErrUnknownResponse,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "default timeout",
			getTs: func(t *testing.T) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					ctx := r.Context()
					done := make(chan struct{})
					go func() {
						defer close(done)
						time.Sleep(time.Hour) // Simulate long processing time
					}()
					select {
					case <-ctx.Done(): // Listen for request cancellation
						http.Error(w, "request canceled", http.StatusRequestTimeout)
						return
					case <-done: // Proceed when processing is done
						_, err := w.Write(validAttestationResponse)
						require.NoError(t, err)
					}
				}))
			},
			timeout:            50 * time.Millisecond,
			expectedError:      tokendata.ErrTimeout,
			expectedStatusCode: http.StatusRequestTimeout,
		},
		{
			name: "rate limit",
			getTs: func(t *testing.T) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusTooManyRequests)
				}))
			},
			timeout:            time.Hour,
			expectedStatusCode: http.StatusTooManyRequests,
			expectedError:      tokendata.ErrRateLimit,
		},
		{
			name: "not found",
			getTs: func(t *testing.T) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotFound)
				}))
			},
			messageHash:        [32]byte{1, 2, 3, 4, 5},
			timeout:            time.Hour,
			expectedError:      tokendata.ErrNotReady,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "success",
			getTs: func(t *testing.T) *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.RequestURI == "/0x0102030400000000000000000000000000000000000000000000000000000000" {
						_, err := w.Write(validAttestationResponse)
						require.NoError(t, err)
					} else {
						w.WriteHeader(http.StatusNotFound)
					}
				}))
			},
			messageHash:        [32]byte{1, 2, 3, 4},
			timeout:            time.Hour,
			expectedStatusCode: http.StatusOK,
			expectedResponse:   validAttestationResponse,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ts := tc.getTs(t)
			defer ts.Close()

			attestationURI, err := url.ParseRequestURI(ts.URL)
			require.NoError(t, err)

			client, err := newHTTPClient(logger.Test(t), attestationURI.String(), tc.timeout, tc.timeout, maxCoolDownDuration)
			require.NoError(t, err)
			response, statusCode, err := client.Get(tests.Context(t), tc.messageHash.String())

			require.Equal(t, tc.expectedStatusCode, statusCode)

			if tc.expectedError != nil {
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedResponse, response)
			}
		})
	}
}

func Test_HTTPClient_Cooldown(t *testing.T) {
	var requestCount int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if requestCount%2 == 1 {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusTooManyRequests)
		}
	}))
	defer ts.Close()

	attestationURI, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	client, err := newHTTPClient(logger.Test(t), attestationURI.String(),
		1*time.Millisecond, longTimeout, maxCoolDownDuration)
	require.NoError(t, err)
	_, _, err = client.Get(tests.Context(t), cciptypes.Bytes32{1, 2, 3}.String())
	require.EqualError(t, err, tokendata.ErrUnknownResponse.Error())

	// First rate-limit activates cooldown and other requests should return rate limit immediately
	for i := 0; i < 10; i++ {
		_, _, err = client.Get(tests.Context(t), cciptypes.Bytes32{1, 2, 3}.String())
		require.EqualError(t, err, tokendata.ErrRateLimit.Error())
	}
	require.Equal(t, requestCount, 2)
}

func Test_HTTPClient_GetInstance(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(validAttestationResponse)
		require.NoError(t, err)
	}))
	defer ts.Close()

	client1, err := GetHTTPClient(logger.Test(t), ts.URL, 1*time.Hour, longTimeout, maxCoolDownDuration)
	require.NoError(t, err)

	client2, err := GetHTTPClient(logger.Test(t), ts.URL, 1*time.Hour, longTimeout, maxCoolDownDuration)
	require.NoError(t, err)

	client3, err := newHTTPClient(logger.Test(t), ts.URL, 1*time.Hour, longTimeout, maxCoolDownDuration)
	require.NoError(t, err)

	assert.True(t, client1 == client2)

	// This not hang and return immediately
	_, _, err = client1.Get(tests.Context(t), cciptypes.Bytes32{1, 2, 3}.String())
	require.NoError(t, err)

	timeoutCtx, cancel := context.WithTimeoutCause(tests.Context(t), 500*time.Millisecond, tokendata.ErrTimeout)
	defer cancel()
	// This should return immediately with timeout error
	_, _, err = client2.Get(timeoutCtx, cciptypes.Bytes32{1, 2, 3}.String())
	require.Error(t, err)
	require.ErrorIs(t, err, tokendata.ErrRateLimit)

	// This is different instance, should return success immediately
	_, _, err = client3.Get(tests.Context(t), cciptypes.Bytes32{1, 2, 3}.String())
	require.NoError(t, err)
}

func Test_HTTPClient_CoolDownWithRetryHeader(t *testing.T) {
	var requestCount int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if requestCount%2 == 1 {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set("Retry-After", time.Now().Add(50*time.Millisecond).Format(time.RFC1123))
			w.WriteHeader(http.StatusTooManyRequests)
		}
	}))
	defer ts.Close()

	attestationURI, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	client, err := newHTTPClient(
		logger.Test(t), attestationURI.String(), 1*time.Millisecond, time.Hour, maxCoolDownDuration,
	)
	require.NoError(t, err)
	_, _, err = client.Get(tests.Context(t), cciptypes.Bytes32{1, 2, 3}.String())
	require.EqualError(t, err, tokendata.ErrUnknownResponse.Error())

	// Getting rate limited, cooling down for 1 second
	_, _, err = client.Get(tests.Context(t), cciptypes.Bytes32{1, 2, 3}.String())
	require.EqualError(t, err, tokendata.ErrRateLimit.Error())

	require.Eventually(t, func() bool {
		_, _, err = client.Get(tests.Context(t), cciptypes.Bytes32{1, 2, 3}.String())
		return errors.Is(err, tokendata.ErrUnknownResponse)
	}, tests.WaitTimeout(t), 50*time.Millisecond)
	require.Equal(t, requestCount, 3)
}

func Test_HTTPClient_RateLimiting_Parallel(t *testing.T) {
	lggr := mocks.NullLogger

	testCases := []struct {
		name         string
		requests     uint64
		rateConfig   time.Duration
		testDuration time.Duration
		timeout      time.Duration
		err          string
	}{
		{
			name:         "rate limited with default config",
			requests:     5,
			rateConfig:   50 * time.Millisecond,
			testDuration: 4 * 50 * time.Millisecond,
		},
		{
			name:         "rate limited with config",
			requests:     5,
			rateConfig:   50 * time.Millisecond,
			testDuration: 4 * 50 * time.Millisecond,
		},
		{
			name:         "timeout after first request",
			requests:     5,
			rateConfig:   50 * time.Millisecond,
			testDuration: 1 * time.Millisecond,
			timeout:      1 * time.Millisecond,
			err:          "token data API is being rate limited",
		},
		{
			name:         "timeout after second request",
			requests:     5,
			rateConfig:   50 * time.Millisecond,
			testDuration: 50 * time.Millisecond,
			timeout:      75 * time.Millisecond,
			err:          "token data API is being rate limited",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, err := w.Write(validAttestationResponse)
				require.NoError(t, err)
			}))
			defer ts.Close()

			attestationURI, err := url.ParseRequestURI(ts.URL)
			require.NoError(t, err)

			client, err := newHTTPClient(lggr, attestationURI.String(), tc.rateConfig, longTimeout, maxCoolDownDuration)
			require.NoError(t, err)

			ctx := context.Background()
			if tc.timeout > 0 {
				var cf context.CancelFunc
				ctx, cf = context.WithTimeout(ctx, tc.timeout)
				defer cf()
			}

			trigger := make(chan struct{})
			errorChan := make(chan error, tc.requests)
			wg := sync.WaitGroup{}
			for i := 0; i < int(tc.requests); i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()

					<-trigger
					_, _, err := client.Get(ctx, cciptypes.Bytes32{0xA}.String())
					if err != nil {
						errorChan <- err
					}
				}()
			}

			// Start the test
			start := time.Now()
			close(trigger)

			// Wait for requests to complete
			wg.Wait()
			finish := time.Now()
			close(errorChan)

			// Collect errors
			errorFound := false
			for err := range errorChan {
				if tc.err != "" && strings.Contains(err.Error(), tc.err) {
					errorFound = true
				} else if err != nil && !strings.Contains(err.Error(), "token data API timed out") {
					// Anything else is unexpected.
					require.Fail(t, "unexpected error", err)
				}
			}

			if tc.err != "" {
				assert.True(t, errorFound)
			}
			assert.WithinDuration(t, start.Add(tc.testDuration), finish, 100*time.Millisecond)
		})
	}
}
