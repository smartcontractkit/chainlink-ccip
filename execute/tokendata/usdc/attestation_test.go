package usdc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal"

	"github.com/smartcontractkit/chainlink-ccip/internal/mocks"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func Test_AttestationClient(t *testing.T) {
	type example struct {
		hash   []byte
		keccak string
	}

	messageA := example{
		hash:   []byte{0xA},
		keccak: "0x0ef9d8f8804d174666011a394cab7901679a8944d24249fd148a6a36071151f8",
	}

	messageB := example{
		hash:   []byte{0xB},
		keccak: "0x60811857dd566889ff6255277d82526f2d9b3bbcb96076be22a5860765ac3d06",
	}

	messageC := example{
		hash:   []byte{0xC},
		keccak: "0x4de0e96b0a8886e42a2c35b57df8a9d58a93b5bff655bc37a30e2ab8e29dc066",
	}

	handler := newMockHandler(t)
	server := httptest.NewServer(handler)
	defer server.Close()

	tt := []struct {
		name     string
		success  []string
		pending  []string
		input    map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes
		expected map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus
	}{
		{
			name:     "empty input",
			input:    map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes{},
			expected: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{},
		},
		{
			name:    "single success",
			success: []string{messageA.keccak},
			input: map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): messageA.hash,
				},
			},
			expected: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): tokendata.SuccessAttestationStatus(
						messageA.hash, internal.MustDecode(messageA.keccak),
					),
				},
			},
		},
		{
			name:    "single pending",
			pending: []string{messageA.keccak},
			input: map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): messageA.hash,
				},
			},
			expected: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): tokendata.ErrorAttestationStatus(tokendata.ErrNotReady),
				},
			},
		},
		{
			name:    "multiple success",
			success: []string{messageA.keccak, messageB.keccak, messageC.keccak},
			input: map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): messageA.hash,
					reader.NewMessageTokenID(1, 2): messageB.hash,
				},
				cciptypes.ChainSelector(2): {
					reader.NewMessageTokenID(2, 1): messageC.hash,
				},
			},
			expected: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): tokendata.SuccessAttestationStatus(
						messageA.hash, internal.MustDecode(messageA.keccak),
					),
					reader.NewMessageTokenID(1, 2): tokendata.SuccessAttestationStatus(
						messageB.hash, internal.MustDecode(messageB.keccak),
					),
				},
				cciptypes.ChainSelector(2): {
					reader.NewMessageTokenID(2, 1): tokendata.SuccessAttestationStatus(
						messageC.hash, internal.MustDecode(messageC.keccak),
					),
				},
			},
		},
		{
			name:    "multiple failures - A, C not ready but B internal error",
			pending: []string{messageA.keccak, messageC.keccak},
			input: map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): messageA.hash,
					reader.NewMessageTokenID(1, 2): messageB.hash,
				},
				cciptypes.ChainSelector(2): {
					reader.NewMessageTokenID(2, 1): messageC.hash,
				},
			},
			expected: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): tokendata.ErrorAttestationStatus(tokendata.ErrNotReady),
					reader.NewMessageTokenID(1, 2): tokendata.ErrorAttestationStatus(tokendata.ErrUnknownResponse),
				},
				cciptypes.ChainSelector(2): {
					reader.NewMessageTokenID(2, 1): tokendata.ErrorAttestationStatus(tokendata.ErrNotReady),
				},
			},
		},
		{
			name:    "mixed success and failure",
			success: []string{messageA.keccak, messageC.keccak},
			pending: []string{messageB.keccak},
			input: map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): messageA.hash,
				},
				cciptypes.ChainSelector(2): {
					reader.NewMessageTokenID(2, 1): messageB.hash,
				},
				cciptypes.ChainSelector(3): {
					reader.NewMessageTokenID(3, 1): messageC.hash,
				},
			},
			expected: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): tokendata.SuccessAttestationStatus(
						messageA.hash, internal.MustDecode(messageA.keccak),
					),
				},
				cciptypes.ChainSelector(2): {
					reader.NewMessageTokenID(2, 1): tokendata.ErrorAttestationStatus(tokendata.ErrNotReady),
				},
				cciptypes.ChainSelector(3): {
					reader.NewMessageTokenID(3, 1): tokendata.SuccessAttestationStatus(
						messageC.hash, internal.MustDecode(messageC.keccak),
					),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(
			tc.name, func(t *testing.T) {
				handler.updateURIs(tc.success, tc.pending)

				client, err := NewSequentialAttestationClient(
					mocks.NullLogger, pluginconfig.USDCCCTPObserverConfig{
						AttestationConfig: pluginconfig.AttestationConfig{
							AttestationAPI:         server.URL,
							AttestationAPIInterval: commonconfig.MustNewDuration(1 * time.Millisecond),
							AttestationAPITimeout:  commonconfig.MustNewDuration(5 * time.Second),
							AttestationAPICooldown: commonconfig.MustNewDuration(5 * time.Minute),
						},
					},
				)
				require.NoError(t, err)
				attestations, err := client.Attestations(tests.Context(t), tc.input)
				require.NoError(t, err)
				require.Equal(t, tc.expected, attestations)
			},
		)
	}
}

func Test_httpResponse(t *testing.T) {
	tt := []struct {
		name                string
		response            httpResponse
		expectedError       error
		expectedAttestation cciptypes.Bytes
	}{
		{
			name: "success",
			response: httpResponse{
				Status:      attestationStatusSuccess,
				Attestation: "0x720502893578a89a8a87982982ef781c18b193",
			},
			expectedAttestation: internal.MustDecode("0x720502893578a89a8a87982982ef781c18b193"),
		},
		{
			name: "pending",
			response: httpResponse{
				Status: attestationStatusPending,
			},
			expectedError: tokendata.ErrNotReady,
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
		t.Run(
			tc.name, func(t *testing.T) {
				err := tc.response.validate()
				if tc.expectedError != nil {
					require.EqualError(t, err, tc.expectedError.Error())
					return
				}

				require.NoError(t, err)
				attestation, err1 := tc.response.attestationToBytes()
				require.NoError(t, err1)
				require.Equal(t, tc.expectedAttestation, attestation)
			},
		)
	}
}

type mockHandler struct {
	t *testing.T

	success map[string]string
	pending map[string]string
	mu      sync.RWMutex
}

func newMockHandler(t *testing.T) *mockHandler {
	return &mockHandler{
		t:  t,
		mu: sync.RWMutex{},
	}
}

func (h *mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	if hash, ok := h.success[r.URL.String()]; ok {
		h.writeJSONResponse(w, "complete", hash)
	} else if hash1, ok1 := h.pending[r.URL.String()]; ok1 {
		h.writeJSONResponse(w, "pending_confirmations", hash1)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *mockHandler) updateURIs(success, pending []string) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.success = h.groupedByURI(success)
	h.pending = h.groupedByURI(pending)
}

func (h *mockHandler) writeJSONResponse(w http.ResponseWriter, status, attestation string) {
	response := fmt.Sprintf(
		`
	{
			"status": "%s",
			"attestation": "%s"
	}`, status, attestation,
	)
	_, err := w.Write([]byte(response))
	require.NoError(h.t, err)
}

func (h *mockHandler) groupedByURI(hashes []string) map[string]string {
	out := make(map[string]string)
	for _, hash := range hashes {
		out["/v1/attestations/"+hash] = hash
	}
	return out
}
