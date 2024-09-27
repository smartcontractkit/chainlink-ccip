package usdc_test

import (
	"context"
	"encoding/hex"
	"fmt"
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

func Test_USDCFlow(t *testing.T) {
	fujiChain := cciptypes.ChainSelector(sel.AVALANCHE_TESTNET_FUJI.Selector)
	fujiPool := internal.RandBytes().String()
	fujiTransmitter := "0xa9fB1b3009DCb79E2fe346c16a604B8Fa8aE0a79"

	baseChain := cciptypes.ChainSelector(sel.ETHEREUM_TESTNET_SEPOLIA_BASE_1.Selector)

	config := map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
		fujiChain: {
			SourcePoolAddress:            fujiPool,
			SourceMessageTransmitterAddr: fujiTransmitter,
		},
	}

	type usdcMessage struct {
		eventPayload        string
		urlMessageHash      string
		attestationResponse string
	}

	//https://testnet.snowtrace.io/tx/0xeeb0ad6b26bacd1570a9361724a36e338f4aacf1170dec64399220b7483b7eed/eventlog?chainid=43113
	usdcMessage1 := "000000000000000100000006000000000004ac0d000000000000000000000000eb08f243e5d3fcff26a9e38ae5520a669f4019d00000000000000000000000009f3b8679c73c2fef8b59b4f3444d4e156fb70aa5000000000000000000000000c08835adf4884e51ff076066706e407506826d9d000000000000000000000000000000005425890298aed601595a70ab815c96711a31bc650000000000000000000000004f32ae7f112c26b109357785e5c66dc5d747fbce00000000000000000000000000000000000000000000000000000000000000640000000000000000000000007a4d8f8c18762d362e64b411d7490fba112811cd"
	//https://testnet.snowtrace.io/tx/0xa7dd97e149c496c52b747bac999d3753a846ea373e1c84b463ffb056516a981b/eventlog?chainid=43113
	usdcMessage2 := "000000000000000100000006000000000004ac0e000000000000000000000000eb08f243e5d3fcff26a9e38ae5520a669f4019d00000000000000000000000009f3b8679c73c2fef8b59b4f3444d4e156fb70aa5000000000000000000000000c08835adf4884e51ff076066706e407506826d9d000000000000000000000000000000005425890298aed601595a70ab815c96711a31bc650000000000000000000000004f32ae7f112c26b109357785e5c66dc5d747fbce000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000007a4d8f8c18762d362e64b411d7490fba112811cd"
	//https://testnet.snowtrace.io/tx/0x389b95c11e2dbe7bf4b8734c4854316707e50e34023299dd37f0f0ebdef9142f
	usdcMessage3 := "000000000000000100000006000000000004ac0f000000000000000000000000eb08f243e5d3fcff26a9e38ae5520a669f4019d00000000000000000000000009f3b8679c73c2fef8b59b4f3444d4e156fb70aa5000000000000000000000000c08835adf4884e51ff076066706e407506826d9d000000000000000000000000000000005425890298aed601595a70ab815c96711a31bc650000000000000000000000004f32ae7f112c26b109357785e5c66dc5d747fbce0000000000000000000000000000000000000000000000000000000000030d400000000000000000000000007a4d8f8c18762d362e64b411d7490fba112811cd"
	// This one is random
	//usdcMessage4 := "000000000000000100000006000000000004ac0f000000000000000000000000eb08f243e5d3fcff26a9e38ae5520a669f4019d00000000000000000000000009f3b8679c73c2fef8b59b4f3444d4e156fb70aa5000000000000000000000000c08835adf4884e51ff076066706e407506826d9d000000000000000000000000000000005425890298aed601595a70ab815c96711a31bc650000000000000000000000004f32ae7f112c26b109357785e5c66dc5d747fbce0000000000000000000000000000000000000000000000000000000000030d400000000000000000000000007a4d8f8c18762d362e64d411d7490fba112811ce"
	//https://iris-api-sandbox.circle.com/v1/attestations/0x69fb1b419d648cf6c9512acad303746dc85af3b864af81985c76764aba60bf6b
	usdcAttestation1 := `{
		"attestation":"0xee466fbd340596aa56e3e40d249869573e4008d84d795b4f2c3cba8649083d08653d38190d0df7e0ee12ae685df2f806d100a03b3716ab1ff2013c7201f1c2d01c9af959b55a4b52dbd0319eed69ce9ace25259830e0b1bff79faf0c9c5d1b5e6d6304e824d657db38f802bcff3e97d0bd30f2ffc62b62381f52c1668ceaa5a73a1b",
		"status":"complete"
	}`
	//https://iris-api-sandbox.circle.com/v1/attestations/0x6ebe09cc552207bdc7bc688ff9fc149d2fd1b712a9bf369e04f37beee55e959d
	usdcAttestation2 := `{
		"attestation":"0xda3130f99f9029757d3326c48ffada1b0886a463181b65e59d7cc40a8984059f641454ccc9907772e64aa52e3e735b62947eea015d35a6bb1bd2d4822ebbb5b51b630cd1c600c0c5a8646ae64ae7778aa8ce52bd074e4d0414a30d19aad0b50b2408659b2f51e5960e372f8f1277045464cf83bc96154027f7ed2008ece45267641c",
		"status":"complete"
	}`
	//https://iris-api-sandbox.circle.com/v1/attestations/0x4d83caf347edd730ccb39afdefdbd312e9ae22a6ec8992087ab0b71818220964
	usdcAttestation3 := `{
		"attestation":"0x068c6043f95632cf22eaa552cc08b9ee8fdf635897978ebccbe171a73838ca277248f4ec75d7f677e3bed73868fb08041ba8b71fc9222e63ae9b479d870c2deb1cb68f12cacc918a6638cabea8c36c9ceb8328e149f0465f986ea240a4a5888be43232a90505cb59a66f68a73d23ace3f91c584f9df87cf9012c5a4ca5a746b6ed1b",
		"status":"complete"
	}`
	// 0xa1e01488a4127753e5008f9e22aef4070be0ac3f795f28a3216d06f50874d493
	usdcAttestation4 := `{
		"error":"Message hash not found"
	}`

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.Contains(r.RequestURI, "0x69fb1b419d648cf6c9512acad303746dc85af3b864af81985c76764aba60bf6b"):
			_, err := w.Write([]byte(usdcAttestation1))
			require.NoError(t, err)
		case strings.Contains(r.RequestURI, "0x6ebe09cc552207bdc7bc688ff9fc149d2fd1b712a9bf369e04f37beee55e959d"):
			_, err := w.Write([]byte(usdcAttestation2))
			require.NoError(t, err)
		case strings.Contains(r.RequestURI, "0x4d83caf347edd730ccb39afdefdbd312e9ae22a6ec8992087ab0b71818220964"):
			_, err := w.Write([]byte(usdcAttestation3))
			require.NoError(t, err)
		case strings.Contains(r.RequestURI, "0xa1e01488a4127753e5008f9e22aef4070be0ac3f795f28a3216d06f50874d493"):
			_, err := w.Write([]byte(usdcAttestation4))
			require.NoError(t, err)
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}

		if strings.Contains(r.RequestURI, "") {

		}
	}))
	defer ts.Close()

	usdcEvents := []types.Sequence{
		{Data: newUSDCMessageEvent(t, usdcMessage1)},
		{Data: newUSDCMessageEvent(t, usdcMessage2)},
		{Data: newUSDCMessageEvent(t, usdcMessage3)},
	}

	// Always return all the events
	r := readermock.NewMockContractReaderFacade(t)
	r.EXPECT().Bind(mock.Anything, mock.Anything).Return(nil)
	r.EXPECT().QueryKey(
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
		mock.Anything,
	).Return(usdcEvents, nil).Maybe()

	usdcReader, err := readerpkg.NewUSDCMessageReader(
		config,
		map[cciptypes.ChainSelector]contractreader.ContractReaderFacade{
			fujiChain: r,
		})
	require.NoError(t, err)

	attestation, err := usdc.NewSequentialAttestationClient(
		pluginconfig.USDCCCTPObserverConfig{
			AttestationAPI:         ts.URL,
			AttestationAPIInterval: commonconfig.MustNewDuration(10 * time.Microsecond),
			AttestationAPITimeout:  commonconfig.MustNewDuration(1 * time.Second),
		})
	require.NoError(t, err)

	ob := usdc.NewTokenDataObserver(
		logger.Test(t),
		baseChain,
		config,
		usdcReader,
		attestation,
	)

	p, err := cciptypes.NewBytesFromString(fujiPool)
	require.NoError(t, err)

	x, err := ob.Observe(context.Background(), exectypes.MessageObservations{
		fujiChain: {
			1: cciptypes.Message{
				TokenAmounts: []cciptypes.RampTokenAmount{
					{
						SourcePoolAddress: p,
						ExtraData:         readerpkg.NewSourceTokenDataPayload(306189, 1).ToBytes(),
						Amount:            cciptypes.NewBigIntFromInt64(100),
					},
				},
			},
		},
	})
	require.NoError(t, err)

	fmt.Println(x)
}

func newUSDCMessageEvent(t *testing.T, messageBody string) *readerpkg.MessageSentEvent {
	body, err := hex.DecodeString(messageBody)
	require.NoError(t, err)

	return &readerpkg.MessageSentEvent{
		Arg0: body,
	}
}
