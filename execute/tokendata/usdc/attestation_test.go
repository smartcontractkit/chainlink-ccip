package usdc

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func Test_AttestationClient(t *testing.T) {
	type example struct {
		hash   []byte
		keccak string
	}

	messageA := example{
		hash:   []byte{0xA},
		keccak: "0ef9d8f8804d174666011a394cab7901679a8944d24249fd148a6a36071151f8",
	}

	messageB := example{
		hash:   []byte{0xB},
		keccak: "60811857dd566889ff6255277d82526f2d9b3bbcb96076be22a5860765ac3d06",
	}

	messageC := example{
		hash:   []byte{0xC},
		keccak: "4de0e96b0a8886e42a2c35b57df8a9d58a93b5bff655bc37a30e2ab8e29dc066",
	}

	tt := []struct {
		name     string
		success  []string
		pending  []string
		input    map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash
		expected map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus
	}{
		{
			name:     "empty input",
			input:    map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash{},
			expected: map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus{},
		},
		{
			name:    "single success",
			success: []string{messageA.keccak},
			input: map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash{
				cciptypes.ChainSelector(1): {
					exectypes.NewMessageTokenID(1, 1): messageA.hash,
				},
			},
			expected: map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus{
				cciptypes.ChainSelector(1): {
					exectypes.NewMessageTokenID(1, 1): SuccessAttestationStatus(messageA.hash, mustDecode(messageA.keccak)),
				},
			},
		},
		{
			name:    "single pending",
			pending: []string{messageA.keccak},
			input: map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash{
				cciptypes.ChainSelector(1): {
					exectypes.NewMessageTokenID(1, 1): messageA.hash,
				},
			},
			expected: map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus{
				cciptypes.ChainSelector(1): {
					exectypes.NewMessageTokenID(1, 1): ErrorAttestationStatus(ErrNotReady),
				},
			},
		},
		{
			name:    "multiple success",
			success: []string{messageA.keccak, messageB.keccak, messageC.keccak},
			input: map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash{
				cciptypes.ChainSelector(1): {
					exectypes.NewMessageTokenID(1, 1): messageA.hash,
					exectypes.NewMessageTokenID(1, 2): messageB.hash,
				},
				cciptypes.ChainSelector(2): {
					exectypes.NewMessageTokenID(2, 1): messageC.hash,
				},
			},
			expected: map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus{
				cciptypes.ChainSelector(1): {
					exectypes.NewMessageTokenID(1, 1): SuccessAttestationStatus(messageA.hash, mustDecode(messageA.keccak)),
					exectypes.NewMessageTokenID(1, 2): SuccessAttestationStatus(messageB.hash, mustDecode(messageB.keccak)),
				},
				cciptypes.ChainSelector(2): {
					exectypes.NewMessageTokenID(2, 1): SuccessAttestationStatus(messageC.hash, mustDecode(messageC.keccak)),
				},
			},
		},
		{
			name:    "multiple failures - A, C not ready but B internal error",
			pending: []string{messageA.keccak, messageC.keccak},
			input: map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash{
				cciptypes.ChainSelector(1): {
					exectypes.NewMessageTokenID(1, 1): messageA.hash,
					exectypes.NewMessageTokenID(1, 2): messageB.hash,
				},
				cciptypes.ChainSelector(2): {
					exectypes.NewMessageTokenID(2, 1): messageC.hash,
				},
			},
			expected: map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus{
				cciptypes.ChainSelector(1): {
					exectypes.NewMessageTokenID(1, 1): ErrorAttestationStatus(ErrNotReady),
					exectypes.NewMessageTokenID(1, 2): ErrorAttestationStatus(ErrUnknownResponse),
				},
				cciptypes.ChainSelector(2): {
					exectypes.NewMessageTokenID(2, 1): ErrorAttestationStatus(ErrNotReady),
				},
			},
		},
		{
			name:    "mixed success and failure",
			success: []string{messageA.keccak, messageC.keccak},
			pending: []string{messageB.keccak},
			input: map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash{
				cciptypes.ChainSelector(1): {
					exectypes.NewMessageTokenID(1, 1): messageA.hash,
				},
				cciptypes.ChainSelector(2): {
					exectypes.NewMessageTokenID(2, 1): messageB.hash,
				},
				cciptypes.ChainSelector(3): {
					exectypes.NewMessageTokenID(3, 1): messageC.hash,
				},
			},
			expected: map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus{
				cciptypes.ChainSelector(1): {
					exectypes.NewMessageTokenID(1, 1): SuccessAttestationStatus(messageA.hash, mustDecode(messageA.keccak)),
				},
				cciptypes.ChainSelector(2): {
					exectypes.NewMessageTokenID(2, 1): ErrorAttestationStatus(ErrNotReady),
				},
				cciptypes.ChainSelector(3): {
					exectypes.NewMessageTokenID(3, 1): SuccessAttestationStatus(messageC.hash, mustDecode(messageC.keccak)),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			server := httptest.NewServer(createHandler(t, tc.success, tc.pending))
			defer server.Close()

			client, err := NewSequentialAttestationClient(pluginconfig.USDCCCTPObserverConfig{
				AttestationAPI:         server.URL,
				AttestationAPIInterval: commonconfig.MustNewDuration(1 * time.Millisecond),
				AttestationAPITimeout:  commonconfig.MustNewDuration(1 * time.Minute),
			})
			require.NoError(t, err)
			attestations, err := client.Attestations(tests.Context(t), tc.input)
			require.NoError(t, err)
			require.Equal(t, tc.expected, attestations)
		})
	}
}

func groupedByURI(hashes []string) map[string]string {
	out := make(map[string]string)
	for _, hash := range hashes {
		out["/v1/attestations/0x"+hash] = hash
	}
	return out
}

func writeJSONResponse(t *testing.T, w http.ResponseWriter, status, attestation string) {
	response := fmt.Sprintf(`
	{
			"status": "%s",
			"attestation": "%s"
	}`, status, attestation)
	_, err := w.Write([]byte(response))
	require.NoError(t, err)
}

func createHandler(t *testing.T, success []string, pending []string) http.HandlerFunc {
	successURIs := groupedByURI(success)
	pendingURIs := groupedByURI(pending)

	return func(w http.ResponseWriter, r *http.Request) {
		if hash, ok := successURIs[r.URL.String()]; ok {
			writeJSONResponse(t, w, "complete", hash)
		} else if hash1, ok1 := pendingURIs[r.URL.String()]; ok1 {
			writeJSONResponse(t, w, "pending_confirmations", hash1)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}
