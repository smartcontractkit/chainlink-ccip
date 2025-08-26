package usdc_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	exectypes2 "github.com/smartcontractkit/chainlink-ccip/ocr3/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/execute/tokendata/usdc"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/internal"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/internal/libs/testhelpers"
	reader2 "github.com/smartcontractkit/chainlink-ccip/ocr3/pkg/reader"
)

func TestTokenDataObserver_Observe_USDCAndRegularTokens(t *testing.T) {
	ethereumUSDCPool := internal.RandBytes().String()
	avalancheUSDCPool := internal.RandBytes().String()
	supportedPoolsBySelector := map[cciptypes.ChainSelector]string{
		cciptypes.ChainSelector(1): ethereumUSDCPool,
		cciptypes.ChainSelector(2): avalancheUSDCPool,
	}
	tests := []struct {
		name                string
		messageObservations exectypes2.MessageObservations
		usdcReader          reader2.USDCMessageReader
		attestationClient   tokendata.AttestationClient
		expectedTokenData   exectypes2.TokenDataObservations
	}{
		{
			name:                "no messages",
			messageObservations: exectypes2.MessageObservations{},
			expectedTokenData:   exectypes2.TokenDataObservations{},
			usdcReader:          reader2.FakeUSDCMessageReader{},
			attestationClient:   &tokendata.FakeAttestationClient{},
		},
		{
			name: "no USDC messages",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, internal.RandBytes().String()),
					11: internal.MessageWithTokens(t),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(exectypes2.NotSupportedTokenData()),
					11: exectypes2.NewMessageTokenData(),
				},
			},
			usdcReader:        reader2.FakeUSDCMessageReader{},
			attestationClient: &tokendata.FakeAttestationClient{},
		},
		{
			name: "single USDC message per chain",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, ethereumUSDCPool),
				},
				2: {
					12: internal.MessageWithTokens(t, avalancheUSDCPool),
				},
			},
			usdcReader: reader2.NewFakeUSDCMessageReader(
				map[reader2.MessageTokenID]cciptypes.Bytes{
					reader2.NewMessageTokenID(10, 0): []byte("message10"),
					reader2.NewMessageTokenID(12, 0): []byte("message12"),
				},
			),
			attestationClient: &tokendata.FakeAttestationClient{
				Data: map[string]tokendata.AttestationStatus{
					"message10": {Attestation: []byte{10_1}},
					"message12": {Attestation: []byte{12_1}},
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(newReadyTokenData([]byte{10_1})),
				},
				2: {
					12: exectypes2.NewMessageTokenData(newReadyTokenData([]byte{12_1})),
				},
			},
		},
		{
			name: "USDC messages mixed with regular  within a single msg chain",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, internal.RandBytes().String(), ethereumUSDCPool),
					11: internal.MessageWithTokens(t, internal.RandBytes().String(), ethereumUSDCPool, internal.RandBytes().String()),
					12: internal.MessageWithTokens(t, internal.RandBytes().String(), internal.RandBytes().String(), ethereumUSDCPool),
					13: internal.MessageWithTokens(t),
				},
			},
			usdcReader: reader2.NewFakeUSDCMessageReader(
				map[reader2.MessageTokenID]cciptypes.Bytes{
					reader2.NewMessageTokenID(10, 1): []byte("message10_1"),
					reader2.NewMessageTokenID(11, 1): []byte("message11_1"),
					reader2.NewMessageTokenID(12, 2): []byte("message12_2"),
				},
			),
			attestationClient: &tokendata.FakeAttestationClient{
				Data: map[string]tokendata.AttestationStatus{
					"message10_1": {Attestation: []byte{10_1}},
					"message11_1": {Attestation: []byte{11_1}},
					"message12_2": {Attestation: []byte{12_2}},
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(
						exectypes2.NotSupportedTokenData(),
						newReadyTokenData([]byte{10_1}),
					),
					11: exectypes2.NewMessageTokenData(
						exectypes2.NotSupportedTokenData(),
						newReadyTokenData([]byte{11_1}),
						exectypes2.NotSupportedTokenData(),
					),
					12: exectypes2.NewMessageTokenData(
						exectypes2.NotSupportedTokenData(),
						exectypes2.NotSupportedTokenData(),
						newReadyTokenData([]byte{12_2}),
					),
					13: exectypes2.NewMessageTokenData(),
				},
			},
		},
		{
			name: "multiple USDC transfer in a single message",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, ethereumUSDCPool, ethereumUSDCPool, ethereumUSDCPool),
				},
				2: {
					12: internal.MessageWithTokens(t, avalancheUSDCPool, avalancheUSDCPool),
				},
			},
			usdcReader: reader2.NewFakeUSDCMessageReader(
				map[reader2.MessageTokenID]cciptypes.Bytes{
					reader2.NewMessageTokenID(10, 0): []byte("message10_0"),
					reader2.NewMessageTokenID(10, 1): []byte("message10_1"),
					reader2.NewMessageTokenID(10, 2): []byte("message10_2"),
					reader2.NewMessageTokenID(12, 0): []byte("message12_0"),
					reader2.NewMessageTokenID(12, 1): []byte("message12_1"),
				},
			),
			attestationClient: &tokendata.FakeAttestationClient{
				Data: map[string]tokendata.AttestationStatus{
					"message10_0": {Attestation: []byte{10_0}},
					"message10_1": {Attestation: []byte{10_1}},
					"message10_2": {Attestation: []byte{10_2}},
					"message12_0": {Attestation: []byte{12_0}},
					"message12_1": {Attestation: []byte{12_1}},
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(
						newReadyTokenData([]byte{10_0}),
						newReadyTokenData([]byte{10_1}),
						newReadyTokenData([]byte{10_2}),
					),
				},
				2: {
					12: exectypes2.NewMessageTokenData(
						newReadyTokenData([]byte{12_0}),
						newReadyTokenData([]byte{12_1}),
					),
				},
			},
		},
		{
			name: "not ready messages are populated to the result set",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, ethereumUSDCPool, ethereumUSDCPool, internal.RandBytes().String()),
				},
			},
			usdcReader: reader2.NewFakeUSDCMessageReader(
				map[reader2.MessageTokenID]cciptypes.Bytes{
					reader2.NewMessageTokenID(10, 0): []byte("message10_0"),
					reader2.NewMessageTokenID(10, 1): []byte("message10_1"),
				},
			),
			attestationClient: &tokendata.FakeAttestationClient{
				Data: map[string]tokendata.AttestationStatus{
					"message10_0": {Attestation: []byte{10_0}},
					"message10_1": {Error: tokendata.ErrNotReady},
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(
						newReadyTokenData([]byte{10_0}),
						exectypes2.TokenData{
							Ready:     false,
							Data:      nil,
							Error:     tokendata.ErrNotReady,
							Supported: true,
						},
						exectypes2.NotSupportedTokenData(),
					),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			observer := usdc.InitUSDCTokenDataObserver(
				logger.Test(t),
				1,
				supportedPoolsBySelector,
				testhelpers.USDCEncoder,
				test.usdcReader,
				test.attestationClient,
			)

			tkData, err := observer.Observe(context.Background(), test.messageObservations)
			require.NoError(t, err)

			require.Equal(t, test.expectedTokenData, tkData)
		})
	}
}

func newReadyTokenData(data []byte) exectypes2.TokenData {
	return exectypes2.TokenData{
		Ready:     true,
		Error:     nil,
		Data:      data,
		Supported: true,
	}
}
