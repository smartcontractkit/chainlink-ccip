package usdc_test

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	sel "github.com/smartcontractkit/chain-selectors"
	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/types"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/usdc"
	"github.com/smartcontractkit/chainlink-ccip/internal"
	readermock "github.com/smartcontractkit/chainlink-ccip/mocks/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	readerpkg "github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type usdcMessage struct {
	// nonce is the nonce of the message, generated by Circle's MessageTransmitter
	nonce uint64
	// sourceDomain is the domain of the source chain, check CCTPDestDomains
	sourceDomain uint32
	// eventPayload is the data from the MessageSent(bytes) event, taken directly from chain explorer
	eventPayload string
	// urlMessageHash is the keccak of the eventPayload, used as msg identifier in Attestation API
	//body, _ := hex.DecodeString(eventPayload)
	//urlMessageHash := utils.Keccak256Fixed(body)
	urlMessageHash string
	// attestationResponse is the response from the Attestation API
	attestationResponse string
	// attestationResponseStatus is the status code of the response from the Attestation API
	attestationResponseStatus int
}

// TODO actual tokenData bytes would be abi encoded, but we can't use abi in the repo so only
// passing attestation as it is
func (u *usdcMessage) tokenData() []byte {
	var result map[string]interface{}

	err := json.Unmarshal([]byte(u.attestationResponse), &result)
	if err != nil {
		panic(err)
	}

	attestation, ok := result["attestation"].(string)
	if !ok {
		panic("attestation not found")
	}

	bytes, err := cciptypes.NewBytesFromString(attestation)
	if err != nil {
		panic(err)
	}
	return []byte(bytes)
}

//nolint:lll
var (
	//https://testnet.snowtrace.io/tx/0xeeb0ad6b26bacd1570a9361724a36e338f4aacf1170dec64399220b7483b7eed/eventlog?chainid=43113
	//https://iris-api-sandbox.circle.com/v1/attestations/0x69fb1b419d648cf6c9512acad303746dc85af3b864af81985c76764aba60bf6b
	m1 = usdcMessage{
		nonce:          306189,
		sourceDomain:   1, // Avalanche Fuji
		eventPayload:   "000000000000000100000006000000000004ac0d000000000000000000000000eb08f243e5d3fcff26a9e38ae5520a669f4019d00000000000000000000000009f3b8679c73c2fef8b59b4f3444d4e156fb70aa5000000000000000000000000c08835adf4884e51ff076066706e407506826d9d000000000000000000000000000000005425890298aed601595a70ab815c96711a31bc650000000000000000000000004f32ae7f112c26b109357785e5c66dc5d747fbce00000000000000000000000000000000000000000000000000000000000000640000000000000000000000007a4d8f8c18762d362e64b411d7490fba112811cd",
		urlMessageHash: "0x69fb1b419d648cf6c9512acad303746dc85af3b864af81985c76764aba60bf6b",
		attestationResponse: `{
			"attestation":"0xee466fbd340596aa56e3e40d249869573e4008d84d795b4f2c3cba8649083d08653d38190d0df7e0ee12ae685df2f806d100a03b3716ab1ff2013c7201f1c2d01c9af959b55a4b52dbd0319eed69ce9ace25259830e0b1bff79faf0c9c5d1b5e6d6304e824d657db38f802bcff3e97d0bd30f2ffc62b62381f52c1668ceaa5a73a1b",
			"status":"complete"
		}`,
		attestationResponseStatus: 200,
	}

	//https://testnet.snowtrace.io/tx/0xa7dd97e149c496c52b747bac999d3753a846ea373e1c84b463ffb056516a981b/eventlog?chainid=43113
	//https://iris-api-sandbox.circle.com/v1/attestations/0x6ebe09cc552207bdc7bc688ff9fc149d2fd1b712a9bf369e04f37beee55e959d
	m2 = usdcMessage{
		nonce:          306190,
		sourceDomain:   1, // Avalanche Fuji
		eventPayload:   "000000000000000100000006000000000004ac0e000000000000000000000000eb08f243e5d3fcff26a9e38ae5520a669f4019d00000000000000000000000009f3b8679c73c2fef8b59b4f3444d4e156fb70aa5000000000000000000000000c08835adf4884e51ff076066706e407506826d9d000000000000000000000000000000005425890298aed601595a70ab815c96711a31bc650000000000000000000000004f32ae7f112c26b109357785e5c66dc5d747fbce000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000007a4d8f8c18762d362e64b411d7490fba112811cd",
		urlMessageHash: "0x6ebe09cc552207bdc7bc688ff9fc149d2fd1b712a9bf369e04f37beee55e959d",
		attestationResponse: `{
			"attestation":"0xda3130f99f9029757d3326c48ffada1b0886a463181b65e59d7cc40a8984059f641454ccc9907772e64aa52e3e735b62947eea015d35a6bb1bd2d4822ebbb5b51b630cd1c600c0c5a8646ae64ae7778aa8ce52bd074e4d0414a30d19aad0b50b2408659b2f51e5960e372f8f1277045464cf83bc96154027f7ed2008ece45267641c",
			"status":"complete"
		}`,
		attestationResponseStatus: 200,
	}

	//https://testnet.snowtrace.io/tx/0x389b95c11e2dbe7bf4b8734c4854316707e50e34023299dd37f0f0ebdef9142f
	//https://iris-api-sandbox.circle.com/v1/attestations/0x4d83caf347edd730ccb39afdefdbd312e9ae22a6ec8992087ab0b71818220964
	m3 = usdcMessage{
		nonce:          306191,
		sourceDomain:   1, // Avalanche Fuji
		eventPayload:   "000000000000000100000006000000000004ac0f000000000000000000000000eb08f243e5d3fcff26a9e38ae5520a669f4019d00000000000000000000000009f3b8679c73c2fef8b59b4f3444d4e156fb70aa5000000000000000000000000c08835adf4884e51ff076066706e407506826d9d000000000000000000000000000000005425890298aed601595a70ab815c96711a31bc650000000000000000000000004f32ae7f112c26b109357785e5c66dc5d747fbce0000000000000000000000000000000000000000000000000000000000030d400000000000000000000000007a4d8f8c18762d362e64b411d7490fba112811cd",
		urlMessageHash: "0x4d83caf347edd730ccb39afdefdbd312e9ae22a6ec8992087ab0b71818220964",
		attestationResponse: `{
			"attestation":"0x068c6043f95632cf22eaa552cc08b9ee8fdf635897978ebccbe171a73838ca277248f4ec75d7f677e3bed73868fb08041ba8b71fc9222e63ae9b479d870c2deb1cb68f12cacc918a6638cabea8c36c9ceb8328e149f0465f986ea240a4a5888be43232a90505cb59a66f68a73d23ace3f91c584f9df87cf9012c5a4ca5a746b6ed1b",
			"status":"complete"
		}`,
		attestationResponseStatus: 200,
	}

	//https://sepolia.etherscan.io/tx/0x63eddd816fbf10872aaf27905a07c37cd5c675c785f55175e9f5529bf94ff7e5
	//https://iris-api-sandbox.circle.com/v1/attestations/0x4055282ce9d64f8fb216c3f6ebd121d4601f0292684ebe4850ad80bd28df7581
	m4 = usdcMessage{
		sourceDomain:   0, // Ethereum Sepolia
		nonce:          262600,
		eventPayload:   "00000000000000000000000600000000000401C80000000000000000000000009F3B8679C73C2FEF8B59B4F3444D4E156FB70AA50000000000000000000000009F3B8679C73C2FEF8B59B4F3444D4E156FB70AA5000000000000000000000000C08835ADF4884E51FF076066706E407506826D9D000000000000000000000000000000001C7D4B196CB0C7B01D743FBC6116A902379C72380000000000000000000000004F32AE7F112C26B109357785E5C66DC5D747FBCE00000000000000000000000000000000000000000000000000000000000027100000000000000000000000003FF675B880AC9F67AC6F4342FFD9E99B80469BAD",
		urlMessageHash: "0x4055282ce9d64f8fb216c3f6ebd121d4601f0292684ebe4850ad80bd28df7581",
		attestationResponse: `{
			"attestation":"0xc9f7eb19bb1828413abc2db13fa941b00d0b52f7519b49c44932562612ebd5956a12e04c034f23e749edc9a68f6aff847d201def0860377efe8d92a64d1fc1af1c11f2e8b12b5142f9c37dd6d0429b8d9dd8ea6a032626819a6252ebc813d4907653c95f8a8a51ab534bc744d5a88499cab0887df73b7cc139a636eb9f05f1996d1c",
			"status":"complete"
		}`,
		attestationResponseStatus: 200,
	}

	//https://sepolia.etherscan.io/tx/0x028a2a08f9b6cd74aa013b5300768585eb2ef10a11e24c25bc456eb2223ad34e
	//https://iris-api-sandbox.circle.com/v1/attestations/0x06b43b556e8ad2eb18aecd6051139641fe9e022b4a1af91a05a726d5005aba59
	m5 = usdcMessage{
		sourceDomain:   0, // Ethereum Sepolia
		nonce:          262601,
		eventPayload:   "00000000000000000000000600000000000401C90000000000000000000000009F3B8679C73C2FEF8B59B4F3444D4E156FB70AA50000000000000000000000009F3B8679C73C2FEF8B59B4F3444D4E156FB70AA5000000000000000000000000C08835ADF4884E51FF076066706E407506826D9D000000000000000000000000000000001C7D4B196CB0C7B01D743FBC6116A902379C72380000000000000000000000004F32AE7F112C26B109357785E5C66DC5D747FBCE00000000000000000000000000000000000000000000000000000000000027100000000000000000000000003FF675B880AC9F67AC6F4342FFD9E99B80469BAD",
		urlMessageHash: "0x06b43b556e8ad2eb18aecd6051139641fe9e022b4a1af91a05a726d5005aba59",
		attestationResponse: `{
			"attestation":"0x365d2c7ebc971426b4de119723f56d8c303bd02ba6d93f74bb45766beb8c09ad626171621bbba543acf5adb5c79af1c3358e85ae3553e357f5aa1e42f22799fb1c8efd51b8e32368fce2bafc608070bd4132eedca53b0b4e29d0c4bb22a171aa8602833cd4398098a075d9e1be636d60f277b506f5f0acee6dff298b3893fcd51c1b",
			"status":"complete"
		}`,
		attestationResponseStatus: 200,
	}

	//https://sepolia.etherscan.io/tx/0xad89c8a5b54a9db693c045918e6714553acd217cc50ce6e73d41043baa324722#eventlog
	//https://iris-api-sandbox.circle.com/v1/attestations/0x038cf8dab6b9ec34741ec65aa347eec8690fca821fd2743a1c78efc6a906d28d
	m6 = usdcMessage{
		sourceDomain:              0, // Ethereum Sepolia
		nonce:                     262602,
		eventPayload:              "00000000000000000000000600000000000401D10000000000000000000000009F3B8679C73C2FEF8B59B4F3444D4E156FB70AA50000000000000000000000009F3B8679C73C2FEF8B59B4F3444D4E156FB70AA5000000000000000000000000C08835ADF4884E51FF076066706E407506826D9D000000000000000000000000000000001C7D4B196CB0C7B01D743FBC6116A902379C72380000000000000000000000004F32AE7F112C26B109357785E5C66DC5D747FBCE00000000000000000000000000000000000000000000000000000000000000640000000000000000000000003FF675B880AC9F67AC6F4342FFD9E99B80469BAD",
		urlMessageHash:            "0x038cf8dab6b9ec34741ec65aa347eec8690fca821fd2743a1c78efc6a906d28d",
		attestationResponse:       `{"attestation":"PENDING","status":"pending_confirmations"}`,
		attestationResponseStatus: 404,
	}

	// this is fake message, but event is properly encoded
	m7 = usdcMessage{
		sourceDomain:              0, // Ethereum Sepolia
		nonce:                     262603,
		eventPayload:              "00000000000000000000000600000000000401D20000000000000000000000009F3B8679C73C2FEF8B59B4F3444D4E156FB70AA50000000000000000000000009F3B8679C73C2FEF8B59B4F3444D4E156FB70AA5000000000000000000000000C08835ADF4884E51FF076066706E407506826D9D000000000000000000000000000000001C7D4B196CB0C7B01D743FBC6116A902379C72380000000000000000000000004F32AE7F112C26B109357785E5C66DC5D747FBCE00000000000000000000000000000000000000000000000000000000000000640000000000000000000000003FF675B880AC9F67AC6F4342FFD9E99B80469BAD",
		urlMessageHash:            "0x27d6ea9a1f55e87575400f87e412d8676d40e9b555e1ab7d020d09b7cfd93083",
		attestationResponse:       `Internal Server Error`,
		attestationResponseStatus: 500,
	}
)

// This test focuses on almost e2e flows for USDC message
// It aims to make it as real as possible by using real payloads from the chain and
// real responses from the Attestation API. As long as the Attestation API is supporting old events
// you should be able to just copy-paste block explorer and attestation requests to the browser
// and see those responses in action
func Test_USDC_CCTP_Flow(t *testing.T) {
	fujiChain := cciptypes.ChainSelector(sel.AVALANCHE_TESTNET_FUJI.Selector)
	fujiPool := internal.RandBytes().String()
	fujiTransmitter := "0xa9fB1b3009DCb79E2fe346c16a604B8Fa8aE0a79"

	sepoliaChain := cciptypes.ChainSelector(sel.ETHEREUM_TESTNET_SEPOLIA.Selector)
	sepoliaPool := internal.RandBytes().String()
	sepoliaTransmitter := "0x7865fAfC2db2093669d92c0F33AeEF291086BEFD"

	baseChain := cciptypes.ChainSelector(sel.ETHEREUM_TESTNET_SEPOLIA_BASE_1.Selector)

	config := map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
		fujiChain: {
			SourcePoolAddress:            fujiPool,
			SourceMessageTransmitterAddr: fujiTransmitter,
		},
		sepoliaChain: {
			SourcePoolAddress:            sepoliaPool,
			SourceMessageTransmitterAddr: sepoliaTransmitter,
		},
	}

	fuji := []usdcMessage{m1, m2, m3}
	sepolia := []usdcMessage{m4, m5, m6, m7}
	all := []usdcMessage{m1, m2, m3, m4, m5, m6, m7}

	// Mock http server to return proper payloads
	server := mockHTTPServerResponse(t, all)
	defer server.Close()

	// Always return all the events
	fujiReader := mockReader(t, fuji)
	sepoliaReader := mockReader(t, sepolia)

	usdcReader, err := readerpkg.NewUSDCMessageReader(
		config,
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			fujiChain:    fujiReader,
			sepoliaChain: sepoliaReader,
		})
	require.NoError(t, err)

	attestation, err := usdc.NewSequentialAttestationClient(
		pluginconfig.USDCCCTPObserverConfig{
			AttestationAPI:         server.URL,
			AttestationAPIInterval: commonconfig.MustNewDuration(1 * time.Microsecond),
			AttestationAPITimeout:  commonconfig.MustNewDuration(1 * time.Second),
		})
	require.NoError(t, err)

	tkReader := usdc.NewTokenDataObserver(
		logger.Test(t),
		baseChain,
		config,
		usdcReader,
		attestation,
	)

	tt := []struct {
		name     string
		messages exectypes.MessageObservations
		want     exectypes.TokenDataObservations
	}{
		{
			name: "single valid message from fuji to base",
			messages: exectypes.MessageObservations{
				fujiChain: {
					1: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(t, m1.nonce, m1.sourceDomain, fujiPool),
						},
					},
				},
			},
			want: exectypes.TokenDataObservations{
				fujiChain: {
					1: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m1.tokenData()),
						},
					},
				},
			},
		},
		{
			name: "multiple valid messages and tokens from fuji to base",
			messages: exectypes.MessageObservations{
				fujiChain: {
					1: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(t, m1.nonce, m1.sourceDomain, fujiPool),
							createToken(t, m2.nonce, m2.sourceDomain, fujiPool),
						},
					},
					2: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(t, m3.nonce, m3.sourceDomain, fujiPool),
						},
					},
				},
			},
			want: exectypes.TokenDataObservations{
				fujiChain: {
					1: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m1.tokenData()),
							exectypes.NewSuccessTokenData(m2.tokenData()),
						},
					},
					2: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m3.tokenData()),
						},
					},
				},
			},
		},
		{
			name: "multiple sepolia tokens within a single message to base",
			messages: exectypes.MessageObservations{
				sepoliaChain: {
					10: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(t, m4.nonce, m4.sourceDomain, sepoliaPool),
							createToken(t, m5.nonce, m5.sourceDomain, sepoliaPool),
						},
					},
				},
			},
			want: exectypes.TokenDataObservations{
				sepoliaChain: {
					10: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m4.tokenData()),
							exectypes.NewSuccessTokenData(m5.tokenData()),
						},
					},
				},
			},
		},
		{
			name: "multiple source chain tokens",
			messages: exectypes.MessageObservations{
				fujiChain: {
					1: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(t, m1.nonce, m1.sourceDomain, fujiPool),
							createToken(t, m2.nonce, m2.sourceDomain, fujiPool),
						},
					},
				},
				sepoliaChain: {
					10: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(t, m4.nonce, m4.sourceDomain, sepoliaPool),
						},
					},
					11: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(t, m5.nonce, m5.sourceDomain, sepoliaPool),
						},
					},
				},
			},
			want: exectypes.TokenDataObservations{
				fujiChain: {
					1: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m1.tokenData()),
							exectypes.NewSuccessTokenData(m2.tokenData()),
						},
					},
				},
				sepoliaChain: {
					10: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m4.tokenData()),
						},
					},
					11: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m5.tokenData()),
						},
					},
				},
			},
		},
		{
			name: "messages with tokens failing to get ready",
			messages: exectypes.MessageObservations{
				fujiChain: {
					1: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							createToken(t, m1.nonce, m1.sourceDomain, fujiPool),
							// not matching pool, marked as not supported
							createToken(t, m2.nonce, m2.sourceDomain, sepoliaPool),
							// not matching domain marked as not ready
							createToken(t, m3.nonce, m5.sourceDomain, fujiPool),
						},
					},
					2: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							// not matching nonce
							createToken(t, 123456789, m3.sourceDomain, fujiPool),
						},
					},
				},
				sepoliaChain: {
					10: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							// message pending
							createToken(t, m6.nonce, m6.sourceDomain, sepoliaPool),
						},
					},
					11: cciptypes.Message{
						TokenAmounts: []cciptypes.RampTokenAmount{
							// message with internal server error
							createToken(t, m7.nonce, m7.sourceDomain, sepoliaPool),
						},
					},
				},
			},
			want: exectypes.TokenDataObservations{
				fujiChain: {
					1: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewSuccessTokenData(m1.tokenData()),
							exectypes.NotSupportedTokenData(),
							exectypes.NewErrorTokenData(usdc.ErrDataMissing),
						},
					},
					2: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewErrorTokenData(usdc.ErrDataMissing),
						},
					},
				},
				sepoliaChain: {
					10: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewErrorTokenData(usdc.ErrDataMissing),
						},
					},
					11: exectypes.MessageTokenData{
						TokenData: []exectypes.TokenData{
							exectypes.NewErrorTokenData(usdc.ErrDataMissing),
						},
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err1 := tkReader.Observe(context.Background(), tc.messages)
			require.NoError(t, err1)
			require.Equal(t, tc.want, got)
		})

	}
}

func createToken(t *testing.T, nonce uint64, sourceDomain uint32, pool string) cciptypes.RampTokenAmount {
	bytesPool, err := cciptypes.NewBytesFromString(pool)
	require.NoError(t, err)

	return cciptypes.RampTokenAmount{
		SourcePoolAddress: bytesPool,
		ExtraData:         readerpkg.NewSourceTokenDataPayload(nonce, sourceDomain).ToBytes(),
		Amount:            cciptypes.NewBigIntFromInt64(100),
	}
}

func mockReader(t *testing.T, message []usdcMessage) *readermock.MockContractReaderFacade {
	items := make([]types.Sequence, len(message))
	for i, m := range message {
		items[i] = types.Sequence{Data: newUSDCMessageEvent(t, m.eventPayload)}
	}

	r := readermock.NewMockContractReaderFacade(t)
	r.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil).Maybe()
	r.EXPECT().QueryKey(
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(items, nil).Maybe()
	return r
}

func mockHTTPServerResponse(t *testing.T, messages []usdcMessage) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, m := range messages {
			if strings.Contains(r.RequestURI, m.urlMessageHash) {
				w.WriteHeader(m.attestationResponseStatus)
				_, err := w.Write([]byte(m.attestationResponse))
				require.NoError(t, err)
				return
			}
		}
		w.WriteHeader(http.StatusInternalServerError)
	}))
	return ts
}

func newUSDCMessageEvent(t *testing.T, messageBody string) *readerpkg.MessageSentEvent {
	body, err := hex.DecodeString(messageBody)
	require.NoError(t, err)
	return &readerpkg.MessageSentEvent{
		Arg0: body,
	}
}
