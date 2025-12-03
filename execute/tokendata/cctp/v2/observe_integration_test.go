package v2

import (
	"context"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
)

// Helper function to create a complete CCTPv2Message with a specific unique ID
// The uniqueID is used to make messages distinct from each other
func createCircleMessageWithID(uniqueID byte, status string) CCTPv2Message {
	return CCTPv2Message{
		Message:     "0x" + fmt.Sprintf("%02x", uniqueID) + "1234abcd",
		EventNonce:  fmt.Sprintf("%d", uniqueID),
		Attestation: "0x5678ef90" + fmt.Sprintf("%02x", uniqueID),
		DecodedMessage: CCTPv2DecodedMessage{
			SourceDomain:      fmt.Sprintf("%d", uniqueID%10),
			DestinationDomain: "1",
			Nonce:             fmt.Sprintf("%d", uniqueID),
			Sender:            "0x" + fmt.Sprintf("%064x", int(uniqueID)+1),
			Recipient:         "0x" + fmt.Sprintf("%064x", int(uniqueID)+2),
			DestinationCaller: "0x" + fmt.Sprintf("%064x", int(uniqueID)+3),
			MessageBody:       "0xbead" + fmt.Sprintf("%02x", uniqueID), // Use even-length hex
			DecodedMessageBody: CCTPv2DecodedMessageBody{
				BurnToken:     "0x" + fmt.Sprintf("%064x", int(uniqueID)+4),
				MintRecipient: "0x" + fmt.Sprintf("%064x", int(uniqueID)+5),
				Amount:        fmt.Sprintf("%d", 1000000+int(uniqueID)*10000),
				MessageSender: "0x" + fmt.Sprintf("%064x", int(uniqueID)+6),
				MaxFee:        fmt.Sprintf("%d", uniqueID*100),
			},
			MinFinalityThreshold: "0",
		},
		CCTPVersion: 2,
		Status:      status,
	}
}

// Helper to create a test observer with mock client (uses mock from observer_test.go)
func createObserverForIntegrationTest(
	t *testing.T,
	poolConfig map[cciptypes.ChainSelector]string,
	mockClient *mockCCTPv2HTTPClient,
) *CCTPv2TokenDataObserver {
	// Create observer with mock HTTP client
	observer := newTestCCTPv2Observer(t, mockClient, cciptypes.ChainSelector(1))

	// Set pool configuration directly
	observer.supportedPoolsBySelector = poolConfig

	return observer
}

func TestCCTPv2TokenDataObserver_Observe(t *testing.T) {
	const (
		testChain1   = cciptypes.ChainSelector(1)
		testChain2   = cciptypes.ChainSelector(2)
		testPoolAddr = "0x1234567890123456789012345678901234567890"
		testTxHash1  = "0xabc123"
		testTxHash2  = "0xdef456"
		testTxHash3  = "0x789abc"
	)

	// Create messages with specific sourceDomains and calculate their deposit hashes
	// Use different unique IDs to ensure distinct messages
	// We use domains 100 and 200 in tests, so adjust the messages accordingly
	msg1Base := createCircleMessageWithID(1, CCTPMessageStatusComplete)
	msg1Base.DecodedMessage.SourceDomain = "100" // Override for test consistency
	msg1 := msg1Base
	hash1, _ := CalculateDepositHash(msg1.DecodedMessage)

	msg2Base := createCircleMessageWithID(2, CCTPMessageStatusComplete)
	msg2Base.DecodedMessage.SourceDomain = "100" // Override for test consistency
	msg2 := msg2Base
	hash2, _ := CalculateDepositHash(msg2.DecodedMessage)

	msg3Base := createCircleMessageWithID(3, CCTPMessageStatusComplete)
	msg3Base.DecodedMessage.SourceDomain = "200" // Different domain for test 22
	msg3 := msg3Base
	hash3, _ := CalculateDepositHash(msg3.DecodedMessage)

	// Convert hashes to hex strings for createCCTPv2Token
	hash1Hex := fmt.Sprintf("%x", hash1)
	hash2Hex := fmt.Sprintf("%x", hash2)
	hash3Hex := fmt.Sprintf("%x", hash3)

	// Construct expected TokenData for successful tokens
	// msg1: Message="0x011234abcd", Attestation="0x5678ef9001"
	msgBytes1, _ := hex.DecodeString("011234abcd")
	attBytes1, _ := hex.DecodeString("5678ef9001")
	expectedData1 := append(msgBytes1, attBytes1...)
	successTokenData1 := exectypes.NewSuccessTokenData(expectedData1)

	// msg2: Message="0x021234abcd", Attestation="0x5678ef9002"
	msgBytes2, _ := hex.DecodeString("021234abcd")
	attBytes2, _ := hex.DecodeString("5678ef9002")
	expectedData2 := append(msgBytes2, attBytes2...)
	successTokenData2 := exectypes.NewSuccessTokenData(expectedData2)

	// msg3: Message="0x031234abcd", Attestation="0x5678ef9003"
	msgBytes3, _ := hex.DecodeString("031234abcd")
	attBytes3, _ := hex.DecodeString("5678ef9003")
	expectedData3 := append(msgBytes3, attBytes3...)
	successTokenData3 := exectypes.NewSuccessTokenData(expectedData3)

	// Common error states
	errorTokenDataMissing := exectypes.NewErrorTokenData(tokendata.ErrDataMissing)
	errorTokenDataNotReady := exectypes.NewErrorTokenData(tokendata.ErrNotReady)
	notSupportedTokenData := exectypes.NotSupportedTokenData()

	// Test 28: FIFO consumption test - different attestations for same message
	// msg1 with attestation 0x1111
	msgBytes1Att1, _ := hex.DecodeString("011234abcd")
	attBytes1Att1, _ := hex.DecodeString("1111")
	expectedData1Att1 := append(msgBytes1Att1, attBytes1Att1...)
	successTokenData1Att1 := exectypes.NewSuccessTokenData(expectedData1Att1)

	// msg1 with attestation 0x2222
	msgBytes1Att2, _ := hex.DecodeString("011234abcd")
	attBytes1Att2, _ := hex.DecodeString("2222")
	expectedData1Att2 := append(msgBytes1Att2, attBytes1Att2...)
	successTokenData1Att2 := exectypes.NewSuccessTokenData(expectedData1Att2)

	// msg1 with attestation 0x3333
	msgBytes1Att3, _ := hex.DecodeString("011234abcd")
	attBytes1Att3, _ := hex.DecodeString("3333")
	expectedData1Att3 := append(msgBytes1Att3, attBytes1Att3...)
	successTokenData1Att3 := exectypes.NewSuccessTokenData(expectedData1Att3)

	tests := []struct {
		name       string
		messages   exectypes.MessageObservations
		poolConfig map[cciptypes.ChainSelector]string
		setupMock  func(*mockCCTPv2HTTPClient)
		validate   func(*testing.T, exectypes.TokenDataObservations, *mockCCTPv2HTTPClient)
	}{
		// ============================================================
		// BASIC HAPPY PATH SCENARIOS
		// ============================================================
		{
			name: "Single chain, single message, single supported token with data found",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				require.Contains(t, result, testChain1)
				require.Contains(t, result[testChain1], cciptypes.SeqNum(10))
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 1)

				// Verify token data matches expected
				assert.Equal(t, successTokenData1, tokenData[0])

				// Verify exactly 1 API call
				assert.Equal(t, 1, m.getCallCount())
			},
		},
		{
			name: "Single chain, single message, multiple supported tokens with unique deposit hashes",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
						createCCTPv2Token(testPoolAddr, 100, hash3Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg2,
						msg3,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 3)

				// Verify each token data matches expected
				assert.Equal(t, successTokenData1, tokenData[0])
				assert.Equal(t, successTokenData2, tokenData[1])
				assert.Equal(t, successTokenData3, tokenData[2])

				// Only 1 API call (deduplication by txHash+sourceDomain)
				assert.Equal(t, 1, m.getCallCount())
			},
		},
		{
			name: "Multiple chains, multiple messages",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
					11: createTestMessage(testTxHash2, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
					}),
				},
				testChain2: {
					20: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 200, hash3Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{
				testChain1: testPoolAddr,
				testChain2: testPoolAddr,
			},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{msg1},
				})
				m.addResponse(testChain1, 100, testTxHash2, CCTPv2Messages{
					Messages: []CCTPv2Message{msg2},
				})
				m.addResponse(testChain2, 200, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{msg3},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// Verify both chains present
				require.Contains(t, result, testChain1)
				require.Contains(t, result, testChain2)
				// Verify sequence numbers preserved
				require.Contains(t, result[testChain1], cciptypes.SeqNum(10))
				require.Contains(t, result[testChain1], cciptypes.SeqNum(11))
				require.Contains(t, result[testChain2], cciptypes.SeqNum(20))

				// Verify token data matches expected
				assert.Equal(t, successTokenData1, result[testChain1][10].TokenData[0])
				assert.Equal(t, successTokenData2, result[testChain1][11].TokenData[0])
				assert.Equal(t, successTokenData3, result[testChain2][20].TokenData[0])

				// 3 API calls (each unique request param)
				assert.Equal(t, 3, m.getCallCount())
			},
		},

		// ============================================================
		// SAME DEPOSIT HASH EDGE CASES
		// ============================================================
		{
			name: "Two token transfers with SAME deposit hash in SAME transaction - sufficient data",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash1Hex), // Same hash!
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg1,
						msg1, // Extra
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 2)

				// Both tokens get msg1 data (FIFO consumption)
				assert.Equal(t, successTokenData1, tokenData[0])
				assert.Equal(t, successTokenData1, tokenData[1])
				// Third message unused (excess)
			},
		},
		{
			name: "Three token transfers with SAME deposit hash in SAME transaction - insufficient data",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg1,
						// Only 2 messages for 3 tokens
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 3)

				// First two tokens get data (FIFO consumption)
				assert.Equal(t, successTokenData1, tokenData[0])
				assert.Equal(t, successTokenData1, tokenData[1])

				// Third token gets ErrDataMissing (insufficient data)
				assert.Equal(t, errorTokenDataMissing, tokenData[2])
			},
		},
		{
			name: "Multiple tokens with SAME deposit hash in SAME transaction - exact match",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg1,
						msg1,
						msg1,
						// Exactly 4 messages for 4 tokens
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 4)

				// All tokens get msg1 data (FIFO consumption, exact match)
				for i := 0; i < 4; i++ {
					assert.Equal(t, successTokenData1, tokenData[i])
				}
			},
		},
		{
			name: "Multiple tokens with SAME deposit hash across DIFFERENT messages in same transaction",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
					11: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg1,
						msg1,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// Msg 10: 2 tokens, Msg 11: 1 token
				msg10Data := result[testChain1][10].TokenData
				msg11Data := result[testChain1][11].TokenData
				require.Len(t, msg10Data, 2)
				require.Len(t, msg11Data, 1)

				// Sequential consumption: msg10 gets first 2, msg11 gets third
				assert.Equal(t, successTokenData1, msg10Data[0])
				assert.Equal(t, successTokenData1, msg10Data[1])
				assert.Equal(t, successTokenData1, msg11Data[0])

				// Only 1 API call (same txHash+domain)
				assert.Equal(t, 1, m.getCallCount())
			},
		},

		// ============================================================
		// DIFFERENT DEPOSIT HASH SCENARIOS
		// ============================================================
		{
			name: "Multiple tokens with different deposit hashes - all found",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
						createCCTPv2Token(testPoolAddr, 100, hash3Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg2,
						msg3,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 3)

				// Each token gets its corresponding message data
				assert.Equal(t, successTokenData1, tokenData[0])
				assert.Equal(t, successTokenData2, tokenData[1])
				assert.Equal(t, successTokenData3, tokenData[2])
			},
		},
		{
			name: "Multiple tokens with different deposit hashes - partial matches",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
						createCCTPv2Token(testPoolAddr, 100, hash3Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg3,
						// hash2 missing
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 3)

				// Token 0 (hash1) - found
				assert.Equal(t, successTokenData1, tokenData[0])
				// Token 1 (hash2) - not found
				assert.Equal(t, errorTokenDataMissing, tokenData[1])
				// Token 2 (hash3) - found
				assert.Equal(t, successTokenData3, tokenData[2])
			},
		},

		// ============================================================
		// MIXED TOKEN SUPPORT SCENARIOS
		// ============================================================
		{
			name: "Mix of supported and unsupported tokens (wrong pool)",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token("0xBAD0000000000000000000000000000000000000", 100, hash2Hex),
						createCCTPv2Token(testPoolAddr, 100, hash3Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg3,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 3)

				// Token 0 - supported and found
				assert.Equal(t, successTokenData1, tokenData[0])
				// Token 1 - not supported (wrong pool)
				assert.Equal(t, notSupportedTokenData, tokenData[1])
				// Token 2 - supported and found
				assert.Equal(t, successTokenData3, tokenData[2])
			},
		},
		{
			name: "Mix of supported and unsupported tokens (invalid ExtraData)",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						{
							SourcePoolAddress: mustDecodeAddress(testPoolAddr),
							ExtraData:         cciptypes.Bytes{0x01, 0x02}, // Invalid - too short
						},
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg2,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 3)

				// Token 0 - supported and found
				assert.Equal(t, successTokenData1, tokenData[0])
				// Token 1 - not supported (invalid ExtraData)
				assert.Equal(t, notSupportedTokenData, tokenData[1])
				// Token 2 - supported and found
				assert.Equal(t, successTokenData2, tokenData[2])
			},
		},
		{
			name: "All tokens unsupported",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token("0x0000000000000000000000000000000000000000", 100, hash1Hex),
						createCCTPv2Token("0x0000000000000000000000000000000000000000", 100, hash2Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				// No responses needed - no API calls should be made
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 2)

				// Both tokens not supported
				assert.Equal(t, notSupportedTokenData, tokenData[0])
				assert.Equal(t, notSupportedTokenData, tokenData[1])

				// No API calls made
				assert.Equal(t, 0, m.getCallCount())
			},
		},
		{
			name: "Supported token but not configured for that chain",
			messages: exectypes.MessageObservations{
				testChain2: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{
				testChain1: testPoolAddr, // Only chain1 configured, not chain2
			},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				// No setup needed
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain2][10].TokenData
				require.Len(t, tokenData, 1)

				assert.Equal(t, notSupportedTokenData, tokenData[0])
				assert.Equal(t, 0, m.getCallCount())
			},
		},
		{
			name: "Empty ExtraData",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						{
							SourcePoolAddress: mustDecodeAddress(testPoolAddr),
							ExtraData:         cciptypes.Bytes{},
						},
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock:  func(m *mockCCTPv2HTTPClient) {},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 1)

				assert.Equal(t, notSupportedTokenData, tokenData[0])
				assert.Equal(t, 0, m.getCallCount())
			},
		},

		// ============================================================
		// ERROR HANDLING FROM DEPENDENCIES
		// ============================================================
		{
			name: "HTTP client returns error",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addError(testChain1, 100, testTxHash1, fmt.Errorf("network timeout"))
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 1)

				// Token is supported but data missing due to HTTP error
				assert.Equal(t, errorTokenDataMissing, tokenData[0])
			},
		},
		{
			name: "HTTP client returns error for one request, success for another",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
					}),
				},
				testChain2: {
					20: createTestMessage(testTxHash2, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 200, hash3Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{
				testChain1: testPoolAddr,
				testChain2: testPoolAddr,
			},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addError(testChain2, 200, testTxHash2, fmt.Errorf("api error"))
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg2,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// Chain1 message succeeds
				assert.Equal(t, successTokenData2, result[testChain1][10].TokenData[0])
				// Chain2 message fails
				assert.Equal(t, errorTokenDataMissing, result[testChain2][20].TokenData[0])
			},
		},
		{
			name: "Message status is not 'complete'",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				// Create a pending version of msg1 (same deposit hash but pending status)
				pendingMsg1 := msg1
				pendingMsg1.Status = "pending"
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						pendingMsg1, // Not complete
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 1)
				assert.Equal(t, errorTokenDataNotReady, tokenData[0])
			},
		},
		{
			name: "Empty message list from HTTP client",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{}, // Empty
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 1)
				assert.Equal(t, errorTokenDataMissing, tokenData[0])
			},
		},
		{
			name: "Message with empty txHash is skipped",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage("", []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock:  func(m *mockCCTPv2HTTPClient) {},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 1)
				// Token is supported but no API call made (no txHash)
				assert.Equal(t, errorTokenDataMissing, tokenData[0])
				assert.Equal(t, 0, m.getCallCount())
			},
		},
		{
			name: "Mix of messages with and without txHash",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage("", []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
					11: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg2,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// Message without txHash gets ErrDataMissing
				assert.Equal(t, errorTokenDataMissing, result[testChain1][10].TokenData[0])
				// Message with txHash succeeds
				assert.Equal(t, successTokenData2, result[testChain1][11].TokenData[0])
				// Only one API call
				assert.Equal(t, 1, m.getCallCount())
			},
		},
		{
			name: "calculateDepositHash returns error (simulated with custom function)",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				// Create message with invalid DecodedMessage that will fail deposit hash calc
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						{
							Message:     "0x1234abcd",
							EventNonce:  "1",
							Attestation: "0x5678ef90",
							DecodedMessage: CCTPv2DecodedMessage{
								SourceDomain:      "invalid",
								DestinationDomain: "1",
							},
							Status: CCTPMessageStatusComplete,
						},
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// Token should be supported but data is missing because calculateDepositHash failed
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 1)
				assert.Equal(t, errorTokenDataMissing, tokenData[0])
			},
		},
		{
			name: "Different source domains in same transaction",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 200, hash2Hex), // Different domain
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				// Two separate API calls - different source domains
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
					},
				})
				m.addResponse(testChain1, 200, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg2,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 2)
				assert.Equal(t, successTokenData1, tokenData[0])
				assert.Equal(t, successTokenData2, tokenData[1])
				// 2 API calls (different source domains)
				assert.Equal(t, 2, m.getCallCount())
				assert.True(t, m.wasCalledWith(testChain1, 100, testTxHash1))
				assert.True(t, m.wasCalledWith(testChain1, 200, testTxHash1))
			},
		},

		// ============================================================
		// EMPTY/NIL INPUT SCENARIOS
		// ============================================================
		{
			name:       "Empty MessageObservations",
			messages:   exectypes.MessageObservations{},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock:  func(m *mockCCTPv2HTTPClient) {},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				assert.Empty(t, result)
				assert.Equal(t, 0, m.getCallCount())
			},
		},
		{
			name: "Message with no tokens",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock:  func(m *mockCCTPv2HTTPClient) {},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				require.Contains(t, result, testChain1)
				require.Contains(t, result[testChain1], cciptypes.SeqNum(10))
				tokenData := result[testChain1][10].TokenData
				assert.Empty(t, tokenData)
				assert.Equal(t, 0, m.getCallCount())
			},
		},
		{
			name: "Chain with no messages",
			messages: exectypes.MessageObservations{
				testChain1: {},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock:  func(m *mockCCTPv2HTTPClient) {},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				require.Contains(t, result, testChain1)
				assert.Empty(t, result[testChain1])
				assert.Equal(t, 0, m.getCallCount())
			},
		},
		{
			name: "Empty pool configuration",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{}, // Empty
			setupMock:  func(m *mockCCTPv2HTTPClient) {},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 1)
				// Token not supported (no pool configured)
				assert.Equal(t, notSupportedTokenData, tokenData[0])
				assert.Equal(t, 0, m.getCallCount())
			},
		},
		{
			name: "Nil TokenAmounts slice",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: {
						Header: cciptypes.RampMessageHeader{
							TxHash: testTxHash1,
						},
						TokenAmounts: nil, // Nil slice
					},
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock:  func(m *mockCCTPv2HTTPClient) {},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				assert.Empty(t, tokenData)
				assert.Equal(t, 0, m.getCallCount())
			},
		},

		// ============================================================
		// DATA MATCHING/CONSUMPTION SCENARIOS
		// ============================================================
		{
			name: "Token data assigned in order (FIFO)",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				// Create three different attestations for the same message (same deposit hash)
				// to test FIFO consumption
				msg1Copy1 := msg1
				msg1Copy1.Attestation = "0x1111"
				msg1Copy2 := msg1
				msg1Copy2.Attestation = "0x2222"
				msg1Copy3 := msg1
				msg1Copy3.Attestation = "0x3333"

				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{msg1Copy1, msg1Copy2, msg1Copy3},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 3)
				// All should be ready, assigned in FIFO order
				assert.Equal(t, successTokenData1Att1, tokenData[0])
				assert.Equal(t, successTokenData1Att2, tokenData[1])
				assert.Equal(t, successTokenData1Att3, tokenData[2])
			},
		},
		{
			name: "Request deduplication by (chain, domain, txHash)",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
					11: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
					}),
					12: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash3Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg2,
						msg3,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// All three messages should succeed
				assert.Equal(t, successTokenData1, result[testChain1][10].TokenData[0])
				assert.Equal(t, successTokenData2, result[testChain1][11].TokenData[0])
				assert.Equal(t, successTokenData3, result[testChain1][12].TokenData[0])
				// Only 1 API call despite 3 messages (same chain, domain, txHash)
				assert.Equal(t, 1, m.getCallCount())
			},
		},
		{
			name: "Data not shared across different request params",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
					11: createTestMessage(testTxHash2, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex), // Same hash, different tx
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				// Only provide data for first txHash
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
					},
				})
				m.addResponse(testChain1, 100, testTxHash2, CCTPv2Messages{
					Messages: []CCTPv2Message{}, // Empty for second tx
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// First succeeds
				assert.Equal(t, successTokenData1, result[testChain1][10].TokenData[0])
				// Second fails (different request params, no data)
				assert.Equal(t, errorTokenDataMissing, result[testChain1][11].TokenData[0])
			},
		},
		{
			name: "Multiple deposits with overlapping hashes - complex consumption",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
						createCCTPv2Token(testPoolAddr, 100, hash1Hex), // Duplicate hash1
						createCCTPv2Token(testPoolAddr, 100, hash3Hex),
						createCCTPv2Token(testPoolAddr, 100, hash2Hex), // Duplicate hash2
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg1, // 2x hash1
						msg2,
						msg2, // 2x hash2
						msg3,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 5)
				// All should be ready - sufficient data for all (FIFO consumption)
				// hash1 (0), hash2 (1), hash1 (2), hash3 (3), hash2 (4)
				assert.Equal(t, successTokenData1, tokenData[0])
				assert.Equal(t, successTokenData2, tokenData[1])
				assert.Equal(t, successTokenData1, tokenData[2])
				assert.Equal(t, successTokenData3, tokenData[3])
				assert.Equal(t, successTokenData2, tokenData[4])
				// Only 1 API call (same chain, domain, txHash)
				assert.Equal(t, 1, m.getCallCount())
			},
		},

		// ============================================================
		// STRUCTURE PRESERVATION
		// ============================================================
		{
			name: "Output structure matches input structure exactly",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
					}),
					20: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash3Hex),
					}),
				},
				testChain2: {
					30: createTestMessage(testTxHash2, []cciptypes.RampTokenAmount{
						createCCTPv2Token("0x0000000000000000000000000000000000000000", 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{
				testChain1: testPoolAddr,
				testChain2: testPoolAddr,
			},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg2,
						msg3,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// Verify structure: 2 chains
				require.Len(t, result, 2)
				require.Contains(t, result, testChain1)
				require.Contains(t, result, testChain2)

				// Chain1: 2 messages
				require.Len(t, result[testChain1], 2)
				require.Contains(t, result[testChain1], cciptypes.SeqNum(10))
				require.Contains(t, result[testChain1], cciptypes.SeqNum(20))

				// Chain1 Msg 10: 2 tokens
				assert.Len(t, result[testChain1][10].TokenData, 2)

				// Chain1 Msg 20: 1 token
				assert.Len(t, result[testChain1][20].TokenData, 1)

				// Chain2: 1 message
				require.Len(t, result[testChain2], 1)
				require.Contains(t, result[testChain2], cciptypes.SeqNum(30))

				// Chain2 Msg 30: 1 token
				assert.Len(t, result[testChain2][30].TokenData, 1)
			},
		},
		{
			name: "Sequence numbers preserved correctly",
			messages: exectypes.MessageObservations{
				testChain1: {
					1: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
					100: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
					}),
					9999: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash3Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg2,
						msg3,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				require.Contains(t, result[testChain1], cciptypes.SeqNum(1))
				require.Contains(t, result[testChain1], cciptypes.SeqNum(100))
				require.Contains(t, result[testChain1], cciptypes.SeqNum(9999))
				assert.Equal(t, successTokenData1, result[testChain1][1].TokenData[0])
				assert.Equal(t, successTokenData2, result[testChain1][100].TokenData[0])
				assert.Equal(t, successTokenData3, result[testChain1][9999].TokenData[0])
			},
		},
		{
			name: "Token count matches input for each message",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
					11: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
						createCCTPv2Token(testPoolAddr, 100, hash3Hex),
						createCCTPv2Token("0x0000000000000000000000000000000000000000", 100, hash1Hex),
					}),
					12: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg2,
						msg3,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// Msg 10: 1 token in = 1 token out
				assert.Len(t, result[testChain1][10].TokenData, 1)
				// Msg 11: 3 tokens in = 3 tokens out
				assert.Len(t, result[testChain1][11].TokenData, 3)
				// Msg 12: 0 tokens in = 0 tokens out
				assert.Len(t, result[testChain1][12].TokenData, 0)
			},
		},

		// ============================================================
		// INTEGRATION/COMPLEX SCENARIOS
		// ============================================================
		{
			name: "Large scale: many chains, messages, tokens",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
					}),
					11: createTestMessage(testTxHash2, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash3Hex),
					}),
				},
				testChain2: {
					20: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 200, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{
				testChain1: testPoolAddr,
				testChain2: testPoolAddr,
			},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg2,
					},
				})
				m.addResponse(testChain1, 100, testTxHash2, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg3,
					},
				})
				m.addResponse(testChain2, 200, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// All tokens should succeed
				assert.Equal(t, successTokenData1, result[testChain1][10].TokenData[0])
				assert.Equal(t, successTokenData2, result[testChain1][10].TokenData[1])
				assert.Equal(t, successTokenData3, result[testChain1][11].TokenData[0])
				assert.Equal(t, successTokenData1, result[testChain2][20].TokenData[0])
				// 3 unique API calls
				assert.Equal(t, 3, m.getCallCount())
			},
		},
		{
			name: "Mixed success and failure across multiple messages",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex), // Will succeed
					}),
					11: createTestMessage(testTxHash2, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash2Hex), // Will fail (API error)
					}),
					12: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token("0x0000000000000000000000000000000000000000", 100, hash3Hex), // Unsupported
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
					},
				})
				m.addError(testChain1, 100, testTxHash2, fmt.Errorf("server error"))
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// Msg 10: success
				assert.Equal(t, successTokenData1, result[testChain1][10].TokenData[0])
				// Msg 11: failure (API error)
				assert.Equal(t, errorTokenDataMissing, result[testChain1][11].TokenData[0])
				// Msg 12: unsupported
				assert.Equal(t, notSupportedTokenData, result[testChain1][12].TokenData[0])
			},
		},
		{
			name: "All possible token states in one observation",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						// Ready + Supported
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						// Not supported (wrong pool)
						createCCTPv2Token("0x0000000000000000000000000000000000000000", 100, hash2Hex),
						// Supported but data missing
						createCCTPv2Token(testPoolAddr, 100, hash3Hex),
						// Not supported (invalid ExtraData)
						{
							SourcePoolAddress: mustDecodeAddress(testPoolAddr),
							ExtraData:         cciptypes.Bytes{0x01},
						},
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						// hash3 not provided
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 4)

				// Token 0: Ready and Supported
				assert.Equal(t, successTokenData1, tokenData[0])

				// Token 1: Not Supported (wrong pool)
				assert.Equal(t, notSupportedTokenData, tokenData[1])

				// Token 2: Supported but data missing
				assert.Equal(t, errorTokenDataMissing, tokenData[2])

				// Token 3: Not Supported (invalid ExtraData)
				assert.Equal(t, notSupportedTokenData, tokenData[3])
			},
		},
		{
			name: "Realistic production scenario",
			messages: exectypes.MessageObservations{
				testChain1: {
					100: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
						createCCTPv2Token(testPoolAddr, 100, hash1Hex), // Duplicate - same tx
					}),
					101: createTestMessage(testTxHash2, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash2Hex),
					}),
				},
				testChain2: {
					200: createTestMessage(testTxHash3, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 200, hash3Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{
				testChain1: testPoolAddr,
				testChain2: testPoolAddr,
			},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
						msg1,
					},
				})
				// Create a pending version of msg2 (same deposit hash but pending status)
				pendingMsg2 := msg2
				pendingMsg2.Status = "pending"
				m.addResponse(testChain1, 100, testTxHash2, CCTPv2Messages{
					Messages: []CCTPv2Message{
						pendingMsg2, // Not ready yet
					},
				})
				m.addResponse(testChain2, 200, testTxHash3, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg3,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// Chain1 Msg 100: both tokens ready (duplicates handled)
				assert.Equal(t, successTokenData1, result[testChain1][100].TokenData[0])
				assert.Equal(t, successTokenData1, result[testChain1][100].TokenData[1])

				// Chain1 Msg 101: not ready (pending status)
				assert.Equal(t, errorTokenDataNotReady, result[testChain1][101].TokenData[0])

				// Chain2 Msg 200: ready
				assert.Equal(t, successTokenData3, result[testChain2][200].TokenData[0])

				// 3 API calls total
				assert.Equal(t, 3, m.getCallCount())
			},
		},

		// ============================================================
		// CONTEXT HANDLING
		// ============================================================
		{
			name: "Context passed through to HTTP client",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				m.addResponse(testChain1, 100, testTxHash1, CCTPv2Messages{
					Messages: []CCTPv2Message{
						msg1,
					},
				})
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// Just verify the call was made successfully
				assert.Equal(t, successTokenData1, result[testChain1][10].TokenData[0])
				assert.Equal(t, 1, m.getCallCount())
			},
		},
		{
			name: "Observe never returns error (errors are in TokenData)",
			messages: exectypes.MessageObservations{
				testChain1: {
					10: createTestMessage(testTxHash1, []cciptypes.RampTokenAmount{
						createCCTPv2Token(testPoolAddr, 100, hash1Hex),
					}),
				},
			},
			poolConfig: map[cciptypes.ChainSelector]string{testChain1: testPoolAddr},
			setupMock: func(m *mockCCTPv2HTTPClient) {
				// Return error from HTTP client
				m.addError(testChain1, 100, testTxHash1, fmt.Errorf("catastrophic failure"))
			},
			validate: func(t *testing.T, result exectypes.TokenDataObservations, m *mockCCTPv2HTTPClient) {
				// Observe() should not return error, but TokenData should contain error
				tokenData := result[testChain1][10].TokenData
				require.Len(t, tokenData, 1)
				assert.Equal(t, errorTokenDataMissing, tokenData[0])
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockClient := newMockCCTPv2HTTPClient()
			tt.setupMock(mockClient)

			observer := createObserverForIntegrationTest(t, tt.poolConfig, mockClient)

			result, err := observer.Observe(context.Background(), tt.messages)
			require.NoError(t, err)

			tt.validate(t, result, mockClient)
		})
	}
}
