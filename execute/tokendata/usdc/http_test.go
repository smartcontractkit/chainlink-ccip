package usdc

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
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

func Test_HTTPClient_Get(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name               string
		getTs              func() *httptest.Server
		timeout            time.Duration
		messageHash        [32]byte
		expectedError      error
		expectedResponse   []byte
		expectedStatusCode int
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
			expectedError:      fmt.Errorf("encoding/hex: odd length hex string"),
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

			client := NewHTTPClient(attestationURI.String(), tc.timeout, tc.timeout)
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
		require.LessOrEqual(t, requestCount, 2)
		if requestCount%2 == 1 {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.WriteHeader(http.StatusTooManyRequests)
		}
	}))
	defer ts.Close()

	attestationURI, err := url.ParseRequestURI(ts.URL)
	require.NoError(t, err)

	client := NewHTTPClient(attestationURI.String(), 1*time.Millisecond, longTimeout)
	_, _, err = client.Get(tests.Context(t), [32]byte{1, 2, 3})

	require.EqualError(t, err, ErrUnknownResponse.Error())

	// First rate-limit activates cooldown and other requests should return rate limit immediately
	for i := 0; i < 10; i++ {
		_, _, err = client.Get(tests.Context(t), [32]byte{1, 2, 3})
		require.EqualError(t, err, ErrRateLimit.Error())
	}
}

func mustDecode(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return b
}
