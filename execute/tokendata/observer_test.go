package tokendata

import (
	"context"
	"crypto/rand"
	"fmt"
	"testing"

	commonconfig "github.com/smartcontractkit/chainlink-common/pkg/config"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func TestTokenDataObserver_Observe_USDCAndRegularTokens(t *testing.T) {
	ethereumUSDCPool := randBytes().String()
	avalancheUSDCPool := randBytes().String()
	config := pluginconfig.TokenDataObserverConfig{
		Type:    "usdc-cctp",
		Version: "1.0",
		USDCCCTPObserverConfig: &pluginconfig.USDCCCTPObserverConfig{
			AttestationAPI:         "https://attestation.api",
			AttestationAPITimeout:  commonconfig.MustNewDuration(1),
			AttestationAPIInterval: commonconfig.MustNewDuration(1),
			Tokens: map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig{
				1: {
					SourcePoolAddress:            ethereumUSDCPool,
					SourceMessageTransmitterAddr: randBytes().String(),
				},
				2: {
					SourcePoolAddress:            avalancheUSDCPool,
					SourceMessageTransmitterAddr: randBytes().String(),
				},
			},
		},
	}

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
			name: "no USDC messages",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: messageWithTokens(t, randBytes().String()),
					11: messageWithTokens(t),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(exectypes.NewNoopTokenData()),
					11: exectypes.NewMessageTokenData(),
				},
			},
		},
		{
			name: "single USDC message per chain",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: messageWithTokens(t, ethereumUSDCPool),
				},
				2: {
					12: messageWithTokens(t, avalancheUSDCPool),
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
		},
		{
			name: "USDC messages mixed with regular  within a single msg chain",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: messageWithTokens(t, randBytes().String(), ethereumUSDCPool),
					11: messageWithTokens(t, randBytes().String(), ethereumUSDCPool, randBytes().String()),
					12: messageWithTokens(t, randBytes().String(), randBytes().String(), ethereumUSDCPool),
					13: messageWithTokens(t),
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
		},
		{
			name: "multiple USDC transfer in a single message",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: messageWithTokens(t, ethereumUSDCPool, ethereumUSDCPool, ethereumUSDCPool),
				},
				2: {
					12: messageWithTokens(t, avalancheUSDCPool, avalancheUSDCPool),
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
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			observer, err := NewCompositeTokenDataObserver([]pluginconfig.TokenDataObserverConfig{config})
			require.NoError(t, err)

			tkData, err := observer.Observe(context.Background(), test.messageObservations)
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
		if _, exists := tokenObservations[chainSelector]; !exists {
			tokenObservations[chainSelector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
		}

		for seq, msg := range messages {
			tokenData := make([]exectypes.TokenData, len(msg.TokenAmounts))
			for i, token := range msg.TokenAmounts {
				if f.IsTokenSupported(chainSelector, token) {
					payload := fmt.Sprintf("%s_%d_%d", f.prefix, seq, i)
					tokenData[i] = exectypes.NewTokenData([]byte(payload))
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

func TestNewCompositeTokenDataObserver_ObserveDifferentTokens(t *testing.T) {
	linkEthereumTokenSourcePool := randBytes().String()
	linkAvalancheTokenSourcePool := randBytes().String()
	usdcEthereumTokenSourcePool := randBytes().String()

	composite := CompositeTokenDataObserver{observers: []TokenDataObserver{
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
	}}

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
			name: "only not-supported tokens",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: messageWithTokens(t, randBytes().String()),
					11: messageWithTokens(t, randBytes().String()),
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
			name: "mixed usdc and link tokens",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: messageWithTokens(t, linkEthereumTokenSourcePool, usdcEthereumTokenSourcePool),
					11: messageWithTokens(t, linkEthereumTokenSourcePool, linkEthereumTokenSourcePool),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(
						exectypes.NewTokenData([]byte("LINK_10_0")),
						exectypes.NewTokenData([]byte("USDC_10_1")),
					),
					11: exectypes.NewMessageTokenData(
						exectypes.NewTokenData([]byte("LINK_11_1")),
						exectypes.NewTokenData([]byte("LINK_11_2")),
					),
				},
			},
		},
		{
			name: "mixed tokens",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: messageWithTokens(t, linkEthereumTokenSourcePool, randBytes().String()),
					11: messageWithTokens(t, linkEthereumTokenSourcePool, linkEthereumTokenSourcePool),
				},
				2: {
					12: messageWithTokens(t, linkAvalancheTokenSourcePool, randBytes().String()),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(
						exectypes.NewTokenData([]byte("LINK_10_0")),
						exectypes.NewNoopTokenData(),
					),
					11: exectypes.NewMessageTokenData(
						exectypes.NewTokenData([]byte("LINK_11_1")),
						exectypes.NewTokenData([]byte("LINK_11_2")),
					),
				},
				2: {
					12: exectypes.NewMessageTokenData(
						exectypes.NewTokenData([]byte("LINK_12_0")),
						exectypes.NewNoopTokenData(),
					),
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

func messageWithTokens(t *testing.T, tokenPoolAddr ...string) cciptypes.Message {
	onRampTokens := make([]cciptypes.RampTokenAmount, len(tokenPoolAddr))
	for i, addr := range tokenPoolAddr {
		b, err := cciptypes.NewBytesFromString(addr)
		require.NoError(t, err)
		onRampTokens[i] = cciptypes.RampTokenAmount{
			SourcePoolAddress: b,
			Amount:            cciptypes.NewBigIntFromInt64(int64(i + 1)),
		}
	}
	return cciptypes.Message{
		TokenAmounts: onRampTokens,
	}
}

func randBytes() cciptypes.Bytes {
	var array [32]byte
	_, err := rand.Read(array[:])
	if err != nil {
		panic(err)
	}
	return array[:]
}
