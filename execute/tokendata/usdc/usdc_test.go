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
	"github.com/smartcontractkit/chainlink-ccip/mocks/internal_/reader"
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
		expectedTokenData   exectypes.TokenDataObservations
		attestationClient   usdc.AttestationClient
	}{
		{
			name:                "no messages",
			messageObservations: exectypes.MessageObservations{},
			expectedTokenData:   exectypes.TokenDataObservations{},
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
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(exectypes.NewTokenData([]byte{10_1})),
				},
				2: {
					12: exectypes.NewMessageTokenData(exectypes.NewTokenData([]byte{12_1})),
				},
			},
			attestationClient: usdc.FakeAttestationClient{},
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
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
						exectypes.NewTokenData([]byte{10_2}),
					),
					11: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
						exectypes.NewTokenData([]byte{11_2}),
						exectypes.NewNoopTokenData(),
					),
					12: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
						exectypes.NewNoopTokenData(),
						exectypes.NewTokenData([]byte{12_3}),
					),
					13: exectypes.NewMessageTokenData(),
				},
			},
			attestationClient: usdc.FakeAttestationClient{},
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
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(
						exectypes.NewTokenData([]byte{10_1}),
						exectypes.NewTokenData([]byte{10_2}),
						exectypes.NewTokenData([]byte{10_3}),
					),
				},
				2: {
					12: exectypes.NewMessageTokenData(
						exectypes.NewTokenData([]byte{12_1}),
						exectypes.NewTokenData([]byte{12_2}),
					),
				},
			},
			attestationClient: usdc.FakeAttestationClient{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			observer := usdc.NewUSDCCCTP(
				config,
				reader.NewMockUSDCMessageReader(t),
				test.attestationClient,
			)

			tkData, err := observer.Observe(context.Background(), test.messageObservations)
			require.NoError(t, err)

			require.Equal(t, test.expectedTokenData, tkData)
		})
	}
}
