package tokendata_test

import (
	"context"
	"fmt"
	"testing"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func Test_CompositeTokenDataObserver_EmptyObservers(t *testing.T) {
	obs, err := tokendata.NewConfigBasedCompositeObservers([]pluginconfig.TokenDataObserverConfig{})
	require.NoError(t, err)

	tests := []struct {
		name                string
		messageObservations exectypes.MessageObservations
		expectedTokenData   exectypes.TokenDataObservations
	}{
		{
			name:                "no messages",
			messageObservations: exectypes.MessageObservations{},
			expectedTokenData:   exectypes.TokenDataObservations{},
		},
		{
			name: "messages without tokens have empty token data",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t),
					11: internal.MessageWithTokens(t),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(),
					11: exectypes.NewMessageTokenData(),
				},
			},
		},
		{
			name: "messages with random tokens have empty states for all tokens",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, internal.RandBytes().String()),
					11: internal.MessageWithTokens(t, internal.RandBytes().String()),
				},
				2: {
					20: internal.MessageWithTokens(t, internal.RandBytes().String(), internal.RandBytes().String()),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(exectypes.NewNoopTokenData()),
					11: exectypes.NewMessageTokenData(exectypes.NewNoopTokenData()),
				},
				2: {
					20: exectypes.NewMessageTokenData(exectypes.NewNoopTokenData(), exectypes.NewNoopTokenData()),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tkData, err1 := obs.Observe(context.Background(), test.messageObservations)
			require.NoError(t, err1)

			require.Equal(t, test.expectedTokenData, tkData)
		})
	}
}

func Test_CompositeTokenDataObserver_ObserveDifferentTokens(t *testing.T) {
	linkEthereumTokenSourcePool := internal.RandBytes().String()
	linkAvalancheTokenSourcePool := internal.RandBytes().String()
	usdcEthereumTokenSourcePool := internal.RandBytes().String()

	composite := tokendata.NewCompositeObservers(
		fakeObserver{
			prefix: "LINK",
			supportedTokens: map[cciptypes.ChainSelector]string{
				1: linkEthereumTokenSourcePool,
				2: linkAvalancheTokenSourcePool,
			}},
		fakeObserver{
			prefix: "USDC",
			supportedTokens: map[cciptypes.ChainSelector]string{
				1: usdcEthereumTokenSourcePool,
			}},
	)

	tests := []struct {
		name                string
		messageObservations exectypes.MessageObservations
		expectedTokenData   exectypes.TokenDataObservations
	}{
		{
			name:                "no messages",
			messageObservations: exectypes.MessageObservations{},
			expectedTokenData:   exectypes.TokenDataObservations{},
		},
		{
			name: "messages without tokens are ignored",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t),
					11: internal.MessageWithTokens(t),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(),
					11: exectypes.NewMessageTokenData(),
				},
			},
		},
		{
			name: "only not-supported tokens",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, internal.RandBytes().String()),
					11: internal.MessageWithTokens(t, internal.RandBytes().String()),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(exectypes.NewNoopTokenData()),
					11: exectypes.NewMessageTokenData(exectypes.NewNoopTokenData()),
				},
			},
		},
		{
			name: "only mixed not-supported tokens",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, internal.RandBytes().String(), internal.RandBytes().String()),
					11: internal.MessageWithTokens(t, internal.RandBytes().String()),
				},
				2: {
					12: internal.MessageWithTokens(t, internal.RandBytes().String(), internal.RandBytes().String()),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(exectypes.NewNoopTokenData(), exectypes.NewNoopTokenData()),
					11: exectypes.NewMessageTokenData(exectypes.NewNoopTokenData()),
				},
				2: {
					12: exectypes.NewMessageTokenData(exectypes.NewNoopTokenData(), exectypes.NewNoopTokenData()),
				},
			},
		},
		{
			name: "mixed usdc and link tokens",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, linkEthereumTokenSourcePool, usdcEthereumTokenSourcePool),
					11: internal.MessageWithTokens(t, linkEthereumTokenSourcePool, linkEthereumTokenSourcePool),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte("LINK_10_0")),
						exectypes.NewSuccessTokenData([]byte("USDC_10_1")),
					),
					11: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte("LINK_11_0")),
						exectypes.NewSuccessTokenData([]byte("LINK_11_1")),
					),
				},
			},
		},
		{
			name: "mixed tokens",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, linkEthereumTokenSourcePool, internal.RandBytes().String()),
					11: internal.MessageWithTokens(t, linkEthereumTokenSourcePool, linkEthereumTokenSourcePool),
				},
				2: {
					12: internal.MessageWithTokens(t, linkAvalancheTokenSourcePool, internal.RandBytes().String()),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte("LINK_10_0")),
						exectypes.NewNoopTokenData(),
					),
					11: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte("LINK_11_0")),
						exectypes.NewSuccessTokenData([]byte("LINK_11_1")),
					),
				},
				2: {
					12: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte("LINK_12_0")),
						exectypes.NewNoopTokenData(),
					),
				},
			},
		},
		{
			name: "not supported tokens for chain selector are ignored",
			messageObservations: exectypes.MessageObservations{
				3: {
					10: internal.MessageWithTokens(t, linkAvalancheTokenSourcePool, internal.RandBytes().String()),
				},
				5: {
					12: internal.MessageWithTokens(t, usdcEthereumTokenSourcePool),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				3: {
					10: exectypes.NewMessageTokenData(exectypes.NewNoopTokenData(), exectypes.NewNoopTokenData()),
				},
				5: {
					12: exectypes.NewMessageTokenData(exectypes.NewNoopTokenData()),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tkData, err := composite.Observe(context.Background(), test.messageObservations)
			require.NoError(t, err)

			require.Equal(t, test.expectedTokenData, tkData)
		})
	}
}

type fakeObserver struct {
	prefix          string
	supportedTokens map[cciptypes.ChainSelector]string
}

func (f fakeObserver) Observe(
	_ context.Context,
	observations exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	tokenObservations := make(exectypes.TokenDataObservations)
	for chainSelector, messages := range observations {
		tokenObservations[chainSelector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

		for seq, msg := range messages {
			tokenData := make([]exectypes.TokenData, len(msg.TokenAmounts))
			for i, token := range msg.TokenAmounts {
				if f.IsTokenSupported(chainSelector, token) {
					payload := fmt.Sprintf("%s_%d_%d", f.prefix, seq, i)
					tokenData[i] = exectypes.NewSuccessTokenData([]byte(payload))
				} else {
					tokenData[i] = exectypes.NotSupportedTokenData()
				}
			}
			tokenObservations[chainSelector][seq] = exectypes.NewMessageTokenData(tokenData...)
		}
	}
	return tokenObservations, nil
}

func (f fakeObserver) IsTokenSupported(sourceChain cciptypes.ChainSelector, msgToken cciptypes.RampTokenAmount) bool {
	tokenAddr, ok := f.supportedTokens[sourceChain]
	return ok && tokenAddr == msgToken.SourcePoolAddress.String()
}
