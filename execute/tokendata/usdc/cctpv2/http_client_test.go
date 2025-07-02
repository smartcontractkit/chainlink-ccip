package cctpv2

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/http"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/usdc"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
)

const (
	// A real response from the CCTPv2 API for a single message.
	//nolint:lll
	realResponse = `{
	  "messages": [
	    {
	      "attestation": "0xeb8f70de97040b6af857960ed32d0510f663d760b9faa54fa45249a9bdba7c8213136a7b66f608e3acaddddb1dac0cdcffdabfbe320eef1b7063b0549728a7581b8c4d19d02a2f72f0273d744d00d83493945abdaa811b3a9ae8c69ccb5b69cce40c0a4c8a2dafcb38b4b4c31a6b343cc220d3fca91c45d90a572479eecb4d641c1b",
	      "message": "0x000000010000000000000001ab82c6ffdf792b349fcaa92ef033a59cb216b48dfb01201583fb0d1752f41b800000000000000000000000008fe6b999dc680ccfdd5bf7eb0974218be2542daa0000000000000000000000008fe6b999dc680ccfdd5bf7eb0974218be2542daa0000000000000000000000000000000000000000000000000000000000000000000003e8000003e8000000010000000000000000000000001c7d4b196cb0c7b01d743fbc6116a902379c7238000000000000000000000000e6094b4f775f9daa59b4dac2535fc0a48f75fc0c00000000000000000000000000000000000000000000000000000000000f4240000000000000000000000000e6094b4f775f9daa59b4dac2535fc0a48f75fc0c00000000000000000000000000000000000000000000000000000000000001f4000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000028d3946",
	      "eventNonce": "0xab82c6ffdf792b349fcaa92ef033a59cb216b48dfb01201583fb0d1752f41b80",
	      "cctpVersion": 2,
	      "status": "complete",
	      "decodedMessage": {
	        "sourceDomain": "0",
	        "destinationDomain": "1",
	        "nonce": "0xab82c6ffdf792b349fcaa92ef033a59cb216b48dfb01201583fb0d1752f41b80",
	        "sender": "0x8fe6b999dc680ccfdd5bf7eb0974218be2542daa",
	        "recipient": "0x8fe6b999dc680ccfdd5bf7eb0974218be2542daa",
	        "destinationCaller": "0x0000000000000000000000000000000000000000000000000000000000000000",
	        "minFinalityThreshold": "1000",
	        "finalityThresholdExecuted": "1000",
	        "messageBody": "0x000000010000000000000000000000001c7d4b196cb0c7b01d743fbc6116a902379c7238000000000000000000000000e6094b4f775f9daa59b4dac2535fc0a48f75fc0c00000000000000000000000000000000000000000000000000000000000f4240000000000000000000000000e6094b4f775f9daa59b4dac2535fc0a48f75fc0c00000000000000000000000000000000000000000000000000000000000001f4000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000028d3946",
	        "decodedMessageBody": {
	          "burnToken": "0x1c7d4b196cb0c7b01d743fbc6116a902379c7238",
	          "mintRecipient": "0xe6094b4f775f9daa59b4dac2535fc0a48f75fc0c",
	          "amount": "1000000",
	          "messageSender": "0xe6094b4f775f9daa59b4dac2535fc0a48f75fc0c",
	          "maxFee": "500",
	          "feeExecuted": "100",
	          "expirationBlock": "42809670",
	          "hookData": null
	        }
	      }
	    }
	  ]
	}`

	validMultiMessages = `{
	  "messages": [
	    {
	      "message": "0xaaa",
	      "eventNonce": "0x11",
	      "attestation": "0x22",
	      "decodedMessage": {
	        "sourceDomain": "3",
	        "destinationDomain": "5",
	        "nonce": "0x11",
	        "sender": "0x456",
	        "recipient": "CCTPV2abc123",
	        "destinationCaller": "0x000",
	        "messageBody": "0x234",
	        "decodedMessageBody": {
	          "burnToken": "0x345",
	          "mintRecipient": "AbC123dEf456",
	          "amount": "5000",
	          "messageSender": "0x901"
	        }
	      },
	      "cctpVersion": 1,
	      "status": "complete"
	    },
	    {
	      "message": "0xbbbb",
	      "eventNonce": "0x22",
	      "attestation": "0x33",
	      "decodedMessage": {
	        "sourceDomain": "1",
	        "destinationDomain": "4",
	        "nonce": "0x22",
	        "sender": "0x789",
	        "recipient": "0x012",
	        "destinationCaller": "0x111",
	        "minFinalityThreshold": "2000",
	        "finalityThresholdExecuted": "2000",
	        "messageBody": "0x234",
	        "decodedMessageBody": {
	          "burnToken": "0xabc",
	          "mintRecipient": "0xdef",
	          "amount": "9001",
	          "messageSender": "0x333",
	          "maxFee": "5",
	          "feeExecuted": "1",
	          "expirationBlock": "42",
	          "hookData": "0x0"
	        }
	      },
	      "cctpVersion": 2,
	      "status": "complete"
	    }
	  ]
	}`

	emptyList = `{"messages":[]}`
	malformed = `{"messages":[{]`          // bad JSON
	wrongType = `{"messages":"not array"}` // messages is a string, not []object
)

var expectedParsedRealResponse = Messages{
	//nolint:lll
	Messages: []Message{
		{
			Attestation: "0xeb8f70de97040b6af857960ed32d0510f663d760b9faa54fa45249a9bdba7c8213136a7b66f608e3acaddddb1dac0cdcffdabfbe320eef1b7063b0549728a7581b8c4d19d02a2f72f0273d744d00d83493945abdaa811b3a9ae8c69ccb5b69cce40c0a4c8a2dafcb38b4b4c31a6b343cc220d3fca91c45d90a572479eecb4d641c1b",
			Message:     "0x000000010000000000000001ab82c6ffdf792b349fcaa92ef033a59cb216b48dfb01201583fb0d1752f41b800000000000000000000000008fe6b999dc680ccfdd5bf7eb0974218be2542daa0000000000000000000000008fe6b999dc680ccfdd5bf7eb0974218be2542daa0000000000000000000000000000000000000000000000000000000000000000000003e8000003e8000000010000000000000000000000001c7d4b196cb0c7b01d743fbc6116a902379c7238000000000000000000000000e6094b4f775f9daa59b4dac2535fc0a48f75fc0c00000000000000000000000000000000000000000000000000000000000f4240000000000000000000000000e6094b4f775f9daa59b4dac2535fc0a48f75fc0c00000000000000000000000000000000000000000000000000000000000001f4000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000028d3946",
			EventNonce:  "0xab82c6ffdf792b349fcaa92ef033a59cb216b48dfb01201583fb0d1752f41b80",
			CCTPVersion: 2,
			Status:      "complete",
			DecodedMessage: DecodedMessage{
				SourceDomain:              "0",
				DestinationDomain:         "1",
				Nonce:                     "0xab82c6ffdf792b349fcaa92ef033a59cb216b48dfb01201583fb0d1752f41b80",
				Sender:                    "0x8fe6b999dc680ccfdd5bf7eb0974218be2542daa",
				Recipient:                 "0x8fe6b999dc680ccfdd5bf7eb0974218be2542daa",
				DestinationCaller:         "0x0000000000000000000000000000000000000000000000000000000000000000",
				MinFinalityThreshold:      "1000",
				FinalityThresholdExecuted: "1000",
				MessageBody:               "0x000000010000000000000000000000001c7d4b196cb0c7b01d743fbc6116a902379c7238000000000000000000000000e6094b4f775f9daa59b4dac2535fc0a48f75fc0c00000000000000000000000000000000000000000000000000000000000f4240000000000000000000000000e6094b4f775f9daa59b4dac2535fc0a48f75fc0c00000000000000000000000000000000000000000000000000000000000001f4000000000000000000000000000000000000000000000000000000000000006400000000000000000000000000000000000000000000000000000000028d3946",
				DecodedMessageBody: DecodedMessageBody{
					BurnToken:       "0x1c7d4b196cb0c7b01d743fbc6116a902379c7238",
					MintRecipient:   "0xe6094b4f775f9daa59b4dac2535fc0a48f75fc0c",
					Amount:          "1000000",
					MessageSender:   "0xe6094b4f775f9daa59b4dac2535fc0a48f75fc0c",
					MaxFee:          "500",
					FeeExecuted:     "100",
					ExpirationBlock: "42809670",
					HookData:        "",
				},
			},
		},
	},
}

var expectedParsedValidMultiMessages = Messages{
	Messages: []Message{
		{
			Message:     "0xaaa",
			EventNonce:  "0x11",
			Attestation: "0x22",
			DecodedMessage: DecodedMessage{
				SourceDomain:      "3",
				DestinationDomain: "5",
				Nonce:             "0x11",
				Sender:            "0x456",
				Recipient:         "CCTPV2abc123",
				DestinationCaller: "0x000",
				MessageBody:       "0x234",
				DecodedMessageBody: DecodedMessageBody{
					BurnToken:     "0x345",
					MintRecipient: "AbC123dEf456",
					Amount:        "5000",
					MessageSender: "0x901",
					// MaxFee, FeeExecuted, ExpirationBlock, HookData omitted (empty)
				},
			},
			CCTPVersion: 1,
			Status:      "complete",
		},
		{
			Message:     "0xbbbb",
			EventNonce:  "0x22",
			Attestation: "0x33",
			DecodedMessage: DecodedMessage{
				SourceDomain:              "1",
				DestinationDomain:         "4",
				Nonce:                     "0x22",
				Sender:                    "0x789",
				Recipient:                 "0x012",
				DestinationCaller:         "0x111",
				MinFinalityThreshold:      "2000",
				FinalityThresholdExecuted: "2000",
				MessageBody:               "0x234",
				DecodedMessageBody: DecodedMessageBody{
					BurnToken:       "0xabc",
					MintRecipient:   "0xdef",
					Amount:          "9001",
					MessageSender:   "0x333",
					MaxFee:          "5",
					FeeExecuted:     "1",
					ExpirationBlock: "42",
					HookData:        "0x0",
				},
			},
			CCTPVersion: 2,
			Status:      "complete",
		},
	},
}

type stubHTTPClient struct {
	body       []byte
	status     int
	err        error
	calledPath string
}

func (s *stubHTTPClient) Get(_ context.Context, path string) (cciptypes.Bytes, http.HTTPStatus, error) {
	s.calledPath = path
	return s.body, http.HTTPStatus(s.status), s.err
}

func (s *stubHTTPClient) Post(
	_ context.Context,
	_ string,
	_ cciptypes.Bytes,
) (cciptypes.Bytes, http.HTTPStatus, error) {
	return nil, http.HTTPStatus(0), fmt.Errorf("not implemented")
}

func TestGetMessages(t *testing.T) {
	type want struct {
		msgs Messages
		err  bool
	}
	tests := []struct {
		name           string
		sourceDomainID uint32
		txHash         string
		stub           stubHTTPClient
		expectedPath   string
		want           want
	}{
		{
			name:           "happy path",
			sourceDomainID: 3,
			txHash:         "0xabc",
			stub: stubHTTPClient{
				body:   []byte(realResponse),
				status: 200,
			},
			expectedPath: "v2/messages/3?transactionHash=0xabc",
			want:         want{msgs: expectedParsedRealResponse, err: false},
		},
		{
			name:           "http client error",
			sourceDomainID: 44,
			txHash:         "0x123",
			stub: stubHTTPClient{
				err: fmt.Errorf("network down"),
			},
			expectedPath: "v2/messages/44?transactionHash=0x123",
			want:         want{err: true},
		},
		{
			name:           "non-200 status",
			sourceDomainID: 3,
			txHash:         "0xabc",
			stub: stubHTTPClient{
				body:   []byte(`{}`),
				status: 502,
			},
			expectedPath: "v2/messages/3?transactionHash=0xabc",
			want:         want{err: true},
		},
		{
			name:           "malformed json -> parse failure",
			sourceDomainID: 3,
			txHash:         "0xabc",
			stub: stubHTTPClient{
				body:   []byte(`{"messages":[{]`),
				status: 200,
			},
			expectedPath: "v2/messages/3?transactionHash=0xabc",
			want:         want{err: true},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cli := &CTTPv2AttestationClient{client: &tc.stub}

			got, err := cli.GetMessages(
				context.Background(),
				tc.sourceDomainID,
				tc.txHash,
			)

			if tc.stub.calledPath != tc.expectedPath {
				t.Fatalf("wrong path used\nwant %s\ngot  %s", tc.expectedPath, tc.stub.calledPath)
			}

			if tc.want.err {
				if err == nil {
					t.Fatalf("expected error, got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tc.want.msgs) {
				t.Fatalf("response mismatch\nwant: %#v\ngot:  %#v", tc.want.msgs, got)
			}
		})
	}
}

func TestParseResponseBody(t *testing.T) {
	tests := []struct {
		name        string
		body        string
		expectErr   bool
		expectCount int
		verify      func(t *testing.T, m Messages)
	}{
		{
			name:        "valid – single message",
			body:        realResponse,
			expectErr:   false,
			expectCount: 1,
			verify: func(t *testing.T, actual Messages) {
				if !reflect.DeepEqual(actual, expectedParsedRealResponse) {
					t.Fatalf("parsed struct mismatch\nwant: %#v\ngot:  %#v", expectedParsedRealResponse, actual)
				}
			},
		},
		{
			name:        "valid – multiple messages",
			body:        validMultiMessages,
			expectErr:   false,
			expectCount: 2,
			verify: func(t *testing.T, actual Messages) {
				if !reflect.DeepEqual(actual, expectedParsedValidMultiMessages) {
					t.Fatalf("parsed struct mismatch\nwant: %#v\ngot:  %#v", expectedParsedValidMultiMessages, actual)
				}
			},
		},
		{
			name:        "valid – empty message list",
			body:        emptyList,
			expectErr:   false,
			expectCount: 0,
		},
		{
			name:      "malformed JSON",
			body:      malformed,
			expectErr: true,
		},
		{
			name:      "wrong type for messages",
			body:      wrongType,
			expectErr: true,
		},
		{
			name:      "empty body",
			body:      "",
			expectErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			msgs, err := parseResponseBody([]byte(tc.body))
			if tc.expectErr {
				if err == nil {
					t.Fatalf("expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got := len(msgs.Messages); got != tc.expectCount {
				t.Fatalf("expected %d messages, got %d", tc.expectCount, got)
			}

			if tc.verify != nil {
				tc.verify(t, msgs)
			}
		})
	}
}

func goodEncoder(_ context.Context, message, attestation cciptypes.Bytes) (cciptypes.Bytes, error) {
	return append(message, attestation...), nil
}
func badEncoder(_ context.Context, _, _ cciptypes.Bytes) (cciptypes.Bytes, error) {
	return nil, fmt.Errorf("encode failed")
}

var baseMsg = Message{
	Message:     "0x01",
	Attestation: "0x02",
	EventNonce:  "0xabc",
	Status:      attestationStatusComplete,
	DecodedMessage: DecodedMessage{
		SourceDomain: "3",
	},
}

func TestMessage_TokenData(t *testing.T) {
	tests := []struct {
		name    string
		msg     Message
		encoder usdc.AttestationEncoder
		want    exectypes.TokenData
		wantErr bool
	}{
		{
			name:    "happy path",
			msg:     baseMsg,
			encoder: goodEncoder,
			want:    exectypes.NewSuccessTokenData([]byte{0x01, 0x02}),
		},
		{
			name:    "status not complete",
			msg:     func() Message { m := baseMsg; m.Status = "processing"; return m }(),
			encoder: goodEncoder,
			wantErr: true,
		},
		{
			name:    "invalid message hex",
			msg:     func() Message { m := baseMsg; m.Message = "0xzz"; return m }(),
			encoder: goodEncoder,
			wantErr: true,
		},
		{
			name:    "invalid attestation hex",
			msg:     func() Message { m := baseMsg; m.Attestation = "0xgh"; return m }(),
			encoder: goodEncoder,
			wantErr: true,
		},
		{
			name:    "encoder failure",
			msg:     baseMsg,
			encoder: badEncoder,
			wantErr: true,
		},
	}

	ctx := context.Background()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.msg.TokenData(ctx, tc.encoder)

			if tc.wantErr {
				if got.Ready || got.Error == nil {
					t.Fatalf("expected error token data, got %+v", got)
				}
				return
			}

			if !got.Ready || got.Error != nil {
				t.Fatalf("unexpected error: %+v", got.Error)
			}
			if !reflect.DeepEqual(got.Data, tc.want.Data) {
				t.Fatalf("data mismatch: want %v, got %v", tc.want.Data, got.Data)
			}
		})
	}
}
