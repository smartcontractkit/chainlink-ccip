package usdc

import (
	"context"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	failedAttestationResponse = []byte(`
	{
		"error": "some error"
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
			client, err := NewHTTPClient(tc.api, 1*time.Millisecond, longTimeout)
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
	t.Parallel()

	tt := []struct {
		name               string
		getTs              func() *httptest.Server
		timeout            time.Duration
		messageHash        [32]byte
		expectedError      error
		expectedResponse   []byte
		expectedStatusCode HTTPStatus
	}{
		{
			name: "server error",
			getTs: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusInternalServerError)
				}))
			},
			timeout:            longTimeout,
			expectedError:      ErrUnknownResponse,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "default timeout",
			getTs: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					time.Sleep(1 * time.Second)
					_, err := w.Write(validAttestationResponse)
					require.NoError(t, err)
				}))
			},
			timeout:            100 * time.Millisecond,
			expectedError:      ErrTimeout,
			expectedStatusCode: http.StatusRequestTimeout,
		},
		{
			name: "200 but attestation response contains error",
			getTs: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					_, err := w.Write(failedAttestationResponse)
					require.NoError(t, err)
				}))
			},
			timeout:            longTimeout,
			expectedStatusCode: http.StatusOK,
			expectedError:      fmt.Errorf("attestation API error: some error"),
		},
		{
			name: "invalid status",
			getTs: func() *httptest.Server {
				attestationResponse := []byte(`
				{
					"status": "complete",
					"attestation": "0"
				}`)

				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					_, err := w.Write(attestationResponse)
					require.NoError(t, err)
				}))
			},
			timeout:            longTimeout,
			expectedStatusCode: http.StatusOK,
			expectedError:      fmt.Errorf("failed to decode attestation hex: encoding/hex: odd length hex string"),
		},
		{
			name: "invalid attestation",
			getTs: func() *httptest.Server {
				attestationResponse := []byte(`
				{
					"status": "",
					"attestation": "0x720502893578a89a8a87982982ef781c18b193"
				}`)

				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					_, err := w.Write(attestationResponse)
					require.NoError(t, err)
				}))
			},
			timeout:            longTimeout,
			expectedStatusCode: http.StatusOK,
			expectedError:      fmt.Errorf("invalid attestation response"),
		},
		{
			name: "malformed response",
			getTs: func() *httptest.Server {
				attestationResponse := []byte(`
				{
					"field": 2137
				}`)

				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					_, err := w.Write(attestationResponse)
					require.NoError(t, err)
				}))
			},
			timeout:            longTimeout,
			expectedStatusCode: http.StatusOK,
			expectedError:      fmt.Errorf("invalid attestation response"),
		},
		{
			name: "rate limit",
			getTs: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusTooManyRequests)
				}))
			},
			timeout:            longTimeout,
			expectedStatusCode: http.StatusTooManyRequests,
			expectedError:      ErrRateLimit,
		},
		{
			name: "not found",
			getTs: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusNotFound)
				}))
			},
			messageHash:        [32]byte{1, 2, 3, 4, 5},
			timeout:            longTimeout,
			expectedError:      ErrNotReady,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name: "success",
			getTs: func() *httptest.Server {
				return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					if r.RequestURI == "/v1/attestations/0x0102030400000000000000000000000000000000000000000000000000000000" {
						_, err := w.Write(validAttestationResponse)
						require.NoError(t, err)
					} else {
						w.WriteHeader(http.StatusNotFound)
					}
				}))
			},
			messageHash:        [32]byte{1, 2, 3, 4},
			timeout:            longTimeout,
			expectedStatusCode: http.StatusOK,
			expectedResponse:   mustDecode("720502893578a89a8a87982982ef781c18b193"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ts := tc.getTs()
			defer ts.Close()

			attestationURI, err := url.ParseRequestURI(ts.URL)
			require.NoError(t, err)

			client, err := NewHTTPClient(attestationURI.String(), tc.timeout, tc.timeout)
			require.NoError(t, err)
			response, statusCode, err := client.Get(tests.Context(t), tc.messageHash)

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

	client, err := NewHTTPClient(attestationURI.String(), 1*time.Millisecond, longTimeout)
	require.NoError(t, err)
	_, _, err = client.Get(tests.Context(t), [32]byte{1, 2, 3})
	require.EqualError(t, err, ErrUnknownResponse.Error())

	// First rate-limit activates cooldown and other requests should return rate limit immediately
	for i := 0; i < 10; i++ {
		_, _, err = client.Get(tests.Context(t), [32]byte{1, 2, 3})
		require.EqualError(t, err, ErrRateLimit.Error())
	}
	require.Equal(t, requestCount, 2)
}

func Test_HTTPClient_CoolDownWithRetryHeader(t *testing.T) {
	var requestCount int
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		if requestCount%2 == 1 {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
		}
	}))
	defer ts.Close()

	attestationURI, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	client, err := NewHTTPClient(attestationURI.String(), 1*time.Millisecond, longTimeout)
	require.NoError(t, err)
	_, _, err = client.Get(tests.Context(t), [32]byte{1, 2, 3})
	require.EqualError(t, err, ErrUnknownResponse.Error())

	// Getting rate limited, cooling down for 1 second
	_, _, err = client.Get(tests.Context(t), [32]byte{1, 2, 3})
	require.EqualError(t, err, ErrRateLimit.Error())

	time.Sleep(2 * time.Second)

	// Next request should go through and reach API
	_, _, err = client.Get(tests.Context(t), [32]byte{1, 2, 3})
	require.EqualError(t, err, ErrUnknownResponse.Error())

	require.Equal(t, requestCount, 3)
}

func Test_HTTPClient_RateLimiting_Parallel(t *testing.T) {
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
			rateConfig:   100 * time.Millisecond,
			testDuration: 4 * 100 * time.Millisecond,
		},
		{
			name:         "rate limited with config",
			requests:     10,
			rateConfig:   50 * time.Millisecond,
			testDuration: 9 * 50 * time.Millisecond,
		},
		{
			name:         "timeout after first request",
			requests:     5,
			rateConfig:   100 * time.Millisecond,
			testDuration: 1 * time.Millisecond,
			timeout:      1 * time.Millisecond,
			err:          "token data API is being rate limited",
		},
		{
			name:         "timeout after second request",
			requests:     5,
			rateConfig:   100 * time.Millisecond,
			testDuration: 100 * time.Millisecond,
			timeout:      150 * time.Millisecond,
			err:          "token data API is being rate limited",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				_, err := w.Write(validAttestationResponse)
				require.NoError(t, err)
			}))
			defer ts.Close()

			attestationURI, err := url.ParseRequestURI(ts.URL)
			require.NoError(t, err)

			client, err := NewHTTPClient(attestationURI.String(), tc.rateConfig, longTimeout)
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
					_, _, err := client.Get(ctx, [32]byte{0xA})
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
			assert.WithinDuration(t, start.Add(tc.testDuration), finish, 50*time.Millisecond)
		})
	}
}

func Test_httpResponse(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name                string
		response            httpResponse
		expectedError       error
		expectedAttestation []byte
	}{
		{
			name: "success",
			response: httpResponse{
				Status:      attestationStatusSuccess,
				Attestation: "0x720502893578a89a8a87982982ef781c18b193",
			},
			expectedAttestation: mustDecode("720502893578a89a8a87982982ef781c18b193"),
		},
		{
			name: "pending",
			response: httpResponse{
				Status: attestationStatusPending,
			},
			expectedError: ErrNotReady,
		},
		{
			name: "error",
			response: httpResponse{
				Error: "some error",
			},
			expectedError: fmt.Errorf("attestation API error: some error"),
		},
		{
			name:          "empty",
			response:      httpResponse{},
			expectedError: fmt.Errorf("invalid attestation response"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.response.validate()
			if tc.expectedError != nil {
				require.EqualError(t, err, tc.expectedError.Error())
				return
			}

			require.NoError(t, err)
			attestation, err1 := tc.response.attestationToBytes()
			require.NoError(t, err1)
			require.Equal(t, tc.expectedAttestation, attestation)
		})
	}
}

func mustDecode(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}
