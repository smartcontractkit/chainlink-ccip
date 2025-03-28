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

type message struct {
	hash        string
	body        []byte
	attestation string
}

func Test_AttestationClient(t *testing.T) {
	messageA := message{
		body:        []byte{0xA},
		hash:        "0x0ef9d8f8804d174666011a394cab7901679a8944d24249fd148a6a36071151f8",
		attestation: "0xddeabb261b885a9676022149101626834649faf58012ec5c2d1b016f8225b734",
	}

	messageB := message{
		body:        []byte{0xB},
		hash:        "0x60811857dd566889ff6255277d82526f2d9b3bbcb96076be22a5860765ac3d06",
		attestation: "0x5d9b34ac9e9995bf44daa0f6f31edc6503432141834a09e28aa96786107a50b0",
	}

	messageC := message{
		body:        []byte{0xC},
		hash:        "0x4de0e96b0a8886e42a2c35b57df8a9d58a93b5bff655bc37a30e2ab8e29dc066",
		attestation: "0x29f2961ca84b5de35682ed5f3fb65795698b3a0dc36eb287c830897ec740235d",
	}

	handler := newMockHandler(t)
	server := httptest.NewServer(handler)
	defer server.Close()

	tt := []struct {
		name     string
		success  []message
		pending  []message
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
			success: []message{messageA},
			input: map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): messageA.body,
				},
			},
			expected: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): tokendata.SuccessAttestationStatus(
						internal.MustDecode(messageA.hash), messageA.body, internal.MustDecode(messageA.attestation),
					),
				},
			},
		},
		{
			name:    "single pending",
			pending: []message{messageA},
			input: map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): messageA.body,
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
			success: []message{messageA, messageB, messageC},
			input: map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): messageA.body,
					reader.NewMessageTokenID(1, 2): messageB.body,
				},
				cciptypes.ChainSelector(2): {
					reader.NewMessageTokenID(2, 1): messageC.body,
				},
			},
			expected: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): tokendata.SuccessAttestationStatus(
						internal.MustDecode(messageA.hash), messageA.body, internal.MustDecode(messageA.attestation),
					),
					reader.NewMessageTokenID(1, 2): tokendata.SuccessAttestationStatus(
						internal.MustDecode(messageB.hash), messageB.body, internal.MustDecode(messageB.attestation),
					),
				},
				cciptypes.ChainSelector(2): {
					reader.NewMessageTokenID(2, 1): tokendata.SuccessAttestationStatus(
						internal.MustDecode(messageC.hash), messageC.body, internal.MustDecode(messageC.attestation),
					),
				},
			},
		},
		{
			name:    "multiple failures - A, C not ready but B internal error",
			pending: []message{messageA, messageC},
			input: map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): messageA.body,
					reader.NewMessageTokenID(1, 2): messageB.body,
				},
				cciptypes.ChainSelector(2): {
					reader.NewMessageTokenID(2, 1): messageC.body,
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
			success: []message{messageA, messageC},
			pending: []message{messageB},
			input: map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): messageA.body,
				},
				cciptypes.ChainSelector(2): {
					reader.NewMessageTokenID(2, 1): messageB.body,
				},
				cciptypes.ChainSelector(3): {
					reader.NewMessageTokenID(3, 1): messageC.body,
				},
			},
			expected: map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus{
				cciptypes.ChainSelector(1): {
					reader.NewMessageTokenID(1, 1): tokendata.SuccessAttestationStatus(
						internal.MustDecode(messageA.hash), messageA.body, internal.MustDecode(messageA.attestation),
					),
				},
				cciptypes.ChainSelector(2): {
					reader.NewMessageTokenID(2, 1): tokendata.ErrorAttestationStatus(tokendata.ErrNotReady),
				},
				cciptypes.ChainSelector(3): {
					reader.NewMessageTokenID(3, 1): tokendata.SuccessAttestationStatus(
						internal.MustDecode(messageC.hash), messageC.body, internal.MustDecode(messageC.attestation),
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
						},
						AttestationAPICooldown: commonconfig.MustNewDuration(5 * time.Minute),
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
		response            cciptypes.Bytes
		expectedError       error
		expectedAttestation cciptypes.Bytes
	}{
		{
			name: "malformed response",
			response: []byte(`
				{
					"field": 2137
				}`),
			expectedError: fmt.Errorf("invalid attestation response"),
		},
		{
			name: "success",
			response: []byte(`
				{
					"status": "complete",
					"attestation": "0x720502893578a89a8a87982982ef781c18b193"
				}`),
			expectedAttestation: internal.MustDecode("0x720502893578a89a8a87982982ef781c18b193"),
		},
		{
			name: "complete but malformed attestation",
			response: []byte(`
				{
					"status": "complete",
					"attestation": "0"
				}`),
			expectedError: fmt.Errorf("failed to decode attestation hex: " +
				"Bytes must be of at least length 2 (i.e, '0x' prefix): 0"),
		},
		{
			name: "pending",
			response: []byte(`
				{
					"status": "pending_confirmations"
				}`),
			expectedError: tokendata.ErrNotReady,
		},
		{
			name: "error",
			response: []byte(`
				{
					"error": "some error"
				}`),
			expectedError: fmt.Errorf("attestation API error: some error"),
		},
		{
			name:          "empty",
			response:      []byte(`{}`),
			expectedError: fmt.Errorf("invalid attestation response"),
		},
	}

	for _, tc := range tt {
		t.Run(
			tc.name, func(t *testing.T) {
				attestation, err := attestationFromResponse(tc.response)
				if tc.expectedError != nil {
					require.EqualError(t, err, tc.expectedError.Error())
					return
				}
				require.NoError(t, err)
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

func (h *mockHandler) updateURIs(success, pending []message) {
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

func (h *mockHandler) groupedByURI(msgs []message) map[string]string {
	out := make(map[string]string)
	for _, msg := range msgs {
		out["/v1/attestations/"+msg.hash] = msg.attestation
	}
	return out
}
