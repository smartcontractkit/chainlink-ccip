package usdc_test

import (
	"context"
	"testing"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/usdc"
	"github.com/smartcontractkit/chainlink-ccip/internal"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func TestTokenDataObserver_Observe_USDCAndRegularTokens(t *testing.T) {
	ethereumUSDCPool := internal.RandBytes().String()
	avalancheUSDCPool := internal.RandBytes().String()
	config := pluginconfig.USDCCCTPObserverConfig{
		AttestationAPI:         "https://attestation.api",
		AttestationAPITimeout:  commonconfig.MustNewDuration(1),
		AttestationAPIInterval: commonconfig.MustNewDuration(1),
		Tokens: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
			1: {
				SourcePoolAddress:            ethereumUSDCPool,
				SourceMessageTransmitterAddr: internal.RandBytes().String(),
			},
			2: {
				SourcePoolAddress:            avalancheUSDCPool,
				SourceMessageTransmitterAddr: internal.RandBytes().String(),
			},
		},
	}

	tests := []struct {
		name                string
		messageObservations exectypes.MessageObservations
		usdcReader          reader.USDCMessageReader
		attestationClient   usdc.AttestationClient
		expectedTokenData   exectypes.TokenDataObservations
	}{
		{
			name:                "no messages",
			messageObservations: exectypes.MessageObservations{},
			expectedTokenData:   exectypes.TokenDataObservations{},
			usdcReader:          reader.FakeUSDCMessageReader{},
			attestationClient:   usdc.FakeAttestationClient{},
		},
		{
			name: "no USDC messages",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, internal.RandBytes().String()),
					11: internal.MessageWithTokens(t),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(exectypes.NotSupportedTokenData()),
					11: exectypes.NewMessageTokenData(),
				},
			},
			usdcReader:        reader.FakeUSDCMessageReader{},
			attestationClient: usdc.FakeAttestationClient{},
		},
		{
			name: "single USDC message per chain",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, ethereumUSDCPool),
				},
				2: {
					12: internal.MessageWithTokens(t, avalancheUSDCPool),
				},
			},
			usdcReader: reader.FakeUSDCMessageReader{
				Messages: map[cciptypes.SeqNum]map[int][]byte{
					10: {
						0: []byte("message10"),
					},
					12: {
						0: []byte("message12"),
					},
				},
			},
			attestationClient: usdc.FakeAttestationClient{
				Data: map[string]usdc.AttestationStatus{
					"message10": {Data: [32]byte{10_1}},
					"message12": {Data: [32]byte{12_1}},
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(newReadyTokenData([32]byte{10_1})),
				},
				2: {
					12: exectypes.NewMessageTokenData(newReadyTokenData([32]byte{12_1})),
				},
			},
		},
		{
			name: "USDC messages mixed with regular  within a single msg chain",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, internal.RandBytes().String(), ethereumUSDCPool),
					11: internal.MessageWithTokens(t, internal.RandBytes().String(), ethereumUSDCPool, internal.RandBytes().String()),
					12: internal.MessageWithTokens(t, internal.RandBytes().String(), internal.RandBytes().String(), ethereumUSDCPool),
					13: internal.MessageWithTokens(t),
				},
			},
			usdcReader: reader.FakeUSDCMessageReader{
				Messages: map[cciptypes.SeqNum]map[int][]byte{
					10: {
						1: []byte("message10_1"),
					},
					11: {
						1: []byte("message11_1"),
					},
					12: {
						2: []byte("message12_2"),
					},
				},
			},
			attestationClient: usdc.FakeAttestationClient{
				Data: map[string]usdc.AttestationStatus{
					"message10_1": {Data: [32]byte{10_1}},
					"message11_1": {Data: [32]byte{11_1}},
					"message12_2": {Data: [32]byte{12_2}},
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(
						exectypes.NotSupportedTokenData(),
						newReadyTokenData([32]byte{10_1}),
					),
					11: exectypes.NewMessageTokenData(
						exectypes.NotSupportedTokenData(),
						newReadyTokenData([32]byte{11_1}),
						exectypes.NotSupportedTokenData(),
					),
					12: exectypes.NewMessageTokenData(
						exectypes.NotSupportedTokenData(),
						exectypes.NotSupportedTokenData(),
						newReadyTokenData([32]byte{12_2}),
					),
					13: exectypes.NewMessageTokenData(),
				},
			},
		},
		{
			name: "multiple USDC transfer in a single message",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, ethereumUSDCPool, ethereumUSDCPool, ethereumUSDCPool),
				},
				2: {
					12: internal.MessageWithTokens(t, avalancheUSDCPool, avalancheUSDCPool),
				},
			},
			usdcReader: reader.FakeUSDCMessageReader{
				Messages: map[cciptypes.SeqNum]map[int][]byte{
					10: {
						0: []byte("message10_0"),
						1: []byte("message10_1"),
						2: []byte("message10_2"),
					},
					12: {
						0: []byte("message12_0"),
						1: []byte("message12_1"),
					},
				},
			},
			attestationClient: usdc.FakeAttestationClient{
				Data: map[string]usdc.AttestationStatus{
					"message10_0": {Data: [32]byte{10_0}},
					"message10_1": {Data: [32]byte{10_1}},
					"message10_2": {Data: [32]byte{10_2}},
					"message12_0": {Data: [32]byte{12_0}},
					"message12_1": {Data: [32]byte{12_1}},
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(
						newReadyTokenData([32]byte{10_0}),
						newReadyTokenData([32]byte{10_1}),
						newReadyTokenData([32]byte{10_2}),
					),
				},
				2: {
					12: exectypes.NewMessageTokenData(
						newReadyTokenData([32]byte{12_0}),
						newReadyTokenData([32]byte{12_1}),
					),
				},
			},
		},
		{
			name: "not ready messages are populated to the result set",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, ethereumUSDCPool, ethereumUSDCPool, internal.RandBytes().String()),
				},
			},
			usdcReader: reader.FakeUSDCMessageReader{
				Messages: map[cciptypes.SeqNum]map[int][]byte{
					10: {
						0: []byte("message10_0"),
						1: []byte("message10_1"),
					},
				},
			},
			attestationClient: usdc.FakeAttestationClient{
				Data: map[string]usdc.AttestationStatus{
					"message10_0": {Data: [32]byte{10_0}},
					"message10_1": {Error: usdc.ErrNotReady},
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(
						newReadyTokenData([32]byte{10_0}),
						exectypes.TokenData{
							Ready:     false,
							Data:      nil,
							Error:     usdc.ErrNotReady,
							Supported: true,
						},
						exectypes.NotSupportedTokenData(),
					),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			observer := usdc.NewTokenDataObserver(
				config,
				test.usdcReader,
				test.attestationClient,
			)

			tkData, err := observer.Observe(context.Background(), test.messageObservations)
			require.NoError(t, err)

			require.Equal(t, test.expectedTokenData, tkData)
		})
	}
}

func newReadyTokenData(data [32]byte) exectypes.TokenData {
	return exectypes.TokenData{
		Ready:     true,
		Error:     nil,
		Data:      data[:],
		Supported: true,
	}
}
