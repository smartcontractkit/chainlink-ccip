package observer_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	exectypes2 "github.com/smartcontractkit/chainlink-ccip/ocr3/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/execute/tokendata/observer"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/internal"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/pluginconfig"
)

func Test_CompositeTokenDataObserver_EmptyObservers(t *testing.T) {
	mockAddrCodec := internal.NewMockAddressCodecHex(t)
	obs, err := observer.NewConfigBasedCompositeObservers(
		t.Context(),
		logger.Test(t),
		100,
		[]pluginconfig.TokenDataObserverConfig{},
		nil,
		nil,
		mockAddrCodec,
	)
	require.NoError(t, err)

	tests := []struct {
		name                string
		messageObservations exectypes2.MessageObservations
		expectedTokenData   exectypes2.TokenDataObservations
	}{
		{
			name:                "no messages",
			messageObservations: exectypes2.MessageObservations{},
			expectedTokenData:   exectypes2.TokenDataObservations{},
		},
		{
			name: "messages without tokens have empty token data",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t),
					11: internal.MessageWithTokens(t),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(),
					11: exectypes2.NewMessageTokenData(),
				},
			},
		},
		{
			name: "messages with random tokens have empty states for all tokens",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, internal.RandBytes().String()),
					11: internal.MessageWithTokens(t, internal.RandBytes().String()),
				},
				2: {
					20: internal.MessageWithTokens(t, internal.RandBytes().String(), internal.RandBytes().String()),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(exectypes2.NewNoopTokenData()),
					11: exectypes2.NewMessageTokenData(exectypes2.NewNoopTokenData()),
				},
				2: {
					20: exectypes2.NewMessageTokenData(exectypes2.NewNoopTokenData(), exectypes2.NewNoopTokenData()),
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

	composite := observer.NewCompositeObservers(
		logger.Test(t),
		fake(
			"LINK",
			map[cciptypes.ChainSelector]string{
				1: linkEthereumTokenSourcePool,
				2: linkAvalancheTokenSourcePool,
			}),
		fake(
			"USDC",
			map[cciptypes.ChainSelector]string{
				1: usdcEthereumTokenSourcePool,
			}),
	)

	tests := []struct {
		name                string
		messageObservations exectypes2.MessageObservations
		expectedTokenData   exectypes2.TokenDataObservations
	}{
		{
			name:                "no messages",
			messageObservations: exectypes2.MessageObservations{},
			expectedTokenData:   exectypes2.TokenDataObservations{},
		},
		{
			name: "messages without tokens are ignored",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t),
					11: internal.MessageWithTokens(t),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(),
					11: exectypes2.NewMessageTokenData(),
				},
			},
		},
		{
			name: "only not-supported tokens",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, internal.RandBytes().String()),
					11: internal.MessageWithTokens(t, internal.RandBytes().String()),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(exectypes2.NewNoopTokenData()),
					11: exectypes2.NewMessageTokenData(exectypes2.NewNoopTokenData()),
				},
			},
		},
		{
			name: "only mixed not-supported tokens",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, internal.RandBytes().String(), internal.RandBytes().String()),
					11: internal.MessageWithTokens(t, internal.RandBytes().String()),
				},
				2: {
					12: internal.MessageWithTokens(t, internal.RandBytes().String(), internal.RandBytes().String()),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(exectypes2.NewNoopTokenData(), exectypes2.NewNoopTokenData()),
					11: exectypes2.NewMessageTokenData(exectypes2.NewNoopTokenData()),
				},
				2: {
					12: exectypes2.NewMessageTokenData(exectypes2.NewNoopTokenData(), exectypes2.NewNoopTokenData()),
				},
			},
		},
		{
			name: "mixed usdc and link tokens",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, linkEthereumTokenSourcePool, usdcEthereumTokenSourcePool),
					11: internal.MessageWithTokens(t, linkEthereumTokenSourcePool, linkEthereumTokenSourcePool),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(
						exectypes2.NewSuccessTokenData([]byte("LINK_10_0")),
						exectypes2.NewSuccessTokenData([]byte("USDC_10_1")),
					),
					11: exectypes2.NewMessageTokenData(
						exectypes2.NewSuccessTokenData([]byte("LINK_11_0")),
						exectypes2.NewSuccessTokenData([]byte("LINK_11_1")),
					),
				},
			},
		},
		{
			name: "mixed tokens",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, linkEthereumTokenSourcePool, internal.RandBytes().String()),
					11: internal.MessageWithTokens(t, linkEthereumTokenSourcePool, linkEthereumTokenSourcePool),
				},
				2: {
					12: internal.MessageWithTokens(t, linkAvalancheTokenSourcePool, internal.RandBytes().String()),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(
						exectypes2.NewSuccessTokenData([]byte("LINK_10_0")),
						exectypes2.NewNoopTokenData(),
					),
					11: exectypes2.NewMessageTokenData(
						exectypes2.NewSuccessTokenData([]byte("LINK_11_0")),
						exectypes2.NewSuccessTokenData([]byte("LINK_11_1")),
					),
				},
				2: {
					12: exectypes2.NewMessageTokenData(
						exectypes2.NewSuccessTokenData([]byte("LINK_12_0")),
						exectypes2.NewNoopTokenData(),
					),
				},
			},
		},
		{
			name: "not supported tokens for chain selector are ignored",
			messageObservations: exectypes2.MessageObservations{
				3: {
					10: internal.MessageWithTokens(t, linkAvalancheTokenSourcePool, internal.RandBytes().String()),
				},
				5: {
					12: internal.MessageWithTokens(t, usdcEthereumTokenSourcePool),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				3: {
					10: exectypes2.NewMessageTokenData(exectypes2.NewNoopTokenData(), exectypes2.NewNoopTokenData()),
				},
				5: {
					12: exectypes2.NewMessageTokenData(exectypes2.NewNoopTokenData()),
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

func Test_CompositeTokenDataObserver_Failures(t *testing.T) {
	linkEthereumTokenSourcePool := internal.RandBytes().String()
	linkAvalancheTokenSourcePool := internal.RandBytes().String()
	usdcEthereumTokenSourcePool := internal.RandBytes().String()

	tests := []struct {
		name                string
		observers           []observer.TokenDataObserver
		messageObservations exectypes2.MessageObservations
		expectedTokenData   exectypes2.TokenDataObservations
	}{
		{
			name: "single observer returns an error but no tokens in messages",
			observers: []observer.TokenDataObserver{
				faulty(
					"stLINK",
					map[cciptypes.ChainSelector]string{
						1: linkEthereumTokenSourcePool,
					}),
			},
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t),
					11: internal.MessageWithTokens(t),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(),
					11: exectypes2.NewMessageTokenData(),
				},
			},
		},
		{
			name: "faulty observer doesn't affect other tokens",
			observers: []observer.TokenDataObserver{
				faulty(
					"LINK",
					map[cciptypes.ChainSelector]string{
						1: linkEthereumTokenSourcePool,
					}),
			},
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, linkEthereumTokenSourcePool, internal.RandBytes().String()),
					11: internal.MessageWithTokens(t, internal.RandBytes().String(), internal.RandBytes().String()),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(
						exectypes2.TokenData{
							Ready:     false,
							Data:      nil,
							Error:     nil,
							Supported: true,
						},
						exectypes2.NewNoopTokenData(),
					),
					11: exectypes2.NewMessageTokenData(
						exectypes2.NewNoopTokenData(),
						exectypes2.NewNoopTokenData(),
					),
				},
			},
		},
		{
			name: "single observer returns an error for tokens for different chains",
			observers: []observer.TokenDataObserver{
				faulty(
					"LINK",
					map[cciptypes.ChainSelector]string{
						1: linkEthereumTokenSourcePool,
					}),
				fake(
					"LINK",
					map[cciptypes.ChainSelector]string{
						2: linkAvalancheTokenSourcePool,
					}),
			},
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, linkEthereumTokenSourcePool),
				},
				2: {
					20: internal.MessageWithTokens(t, linkAvalancheTokenSourcePool),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					// TokenObserver failed to process that
					10: exectypes2.NewMessageTokenData(exectypes2.TokenData{
						Ready:     false,
						Data:      nil,
						Error:     nil,
						Supported: true,
					}),
				},
				2: {
					20: exectypes2.NewMessageTokenData(exectypes2.NewSuccessTokenData([]byte("LINK_20_0"))),
				},
			},
		},
		{
			name: "multiple observers return an error for tokens for different chains",
			observers: []observer.TokenDataObserver{
				faulty(
					"LINK",
					map[cciptypes.ChainSelector]string{
						1: linkEthereumTokenSourcePool,
					}),
				fake(
					"USDC",
					map[cciptypes.ChainSelector]string{
						1: usdcEthereumTokenSourcePool,
					}),
			},
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, linkEthereumTokenSourcePool),
					11: internal.MessageWithTokens(t, usdcEthereumTokenSourcePool),
					12: internal.MessageWithTokens(t, internal.RandBytes().String(), linkEthereumTokenSourcePool),
					13: internal.MessageWithTokens(t, internal.RandBytes().String(), usdcEthereumTokenSourcePool),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(exectypes2.TokenData{
						Ready:     false,
						Data:      nil,
						Error:     nil,
						Supported: true,
					}),
					11: exectypes2.NewMessageTokenData(
						exectypes2.NewSuccessTokenData([]byte("USDC_11_0")),
					),
					12: exectypes2.NewMessageTokenData(
						exectypes2.NewNoopTokenData(),
						exectypes2.TokenData{
							Ready:     false,
							Data:      nil,
							Error:     nil,
							Supported: true,
						},
					),
					13: exectypes2.NewMessageTokenData(
						exectypes2.NewNoopTokenData(),
						exectypes2.NewSuccessTokenData([]byte("USDC_13_1")),
					),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			composite := observer.NewCompositeObservers(logger.Test(t), test.observers...)

			tkData, err := composite.Observe(context.Background(), test.messageObservations)
			require.NoError(t, err)

			require.Equal(t, test.expectedTokenData, tkData)
		})

	}
}

func faulty(prefix string, supportedTokens map[cciptypes.ChainSelector]string) observer.TokenDataObserver {
	return fakeObserver{
		prefix:          prefix,
		faulty:          true,
		supportedTokens: supportedTokens,
	}
}

func fake(prefix string, supportedTokens map[cciptypes.ChainSelector]string) observer.TokenDataObserver {
	return fakeObserver{
		prefix:          prefix,
		faulty:          false,
		supportedTokens: supportedTokens,
	}
}

type fakeObserver struct {
	faulty          bool
	prefix          string
	supportedTokens map[cciptypes.ChainSelector]string
}

func (f fakeObserver) Observe(
	_ context.Context,
	observations exectypes2.MessageObservations,
) (exectypes2.TokenDataObservations, error) {
	if f.faulty {
		return nil, fmt.Errorf("error")
	}

	tokenObservations := make(exectypes2.TokenDataObservations)
	for chainSelector, messages := range observations {
		tokenObservations[chainSelector] = make(map[cciptypes.SeqNum]exectypes2.MessageTokenData)

		for seq, msg := range messages {
			tokenData := make([]exectypes2.TokenData, len(msg.TokenAmounts))
			for i, token := range msg.TokenAmounts {
				if f.IsTokenSupported(chainSelector, token) {
					payload := fmt.Sprintf("%s_%d_%d", f.prefix, seq, i)
					tokenData[i] = exectypes2.NewSuccessTokenData([]byte(payload))
				} else {
					tokenData[i] = exectypes2.NotSupportedTokenData()
				}
			}
			tokenObservations[chainSelector][seq] = exectypes2.NewMessageTokenData(tokenData...)
		}
	}
	return tokenObservations, nil
}

func (f fakeObserver) IsTokenSupported(sourceChain cciptypes.ChainSelector, msgToken cciptypes.RampTokenAmount) bool {
	tokenAddr, ok := f.supportedTokens[sourceChain]
	return ok && tokenAddr == msgToken.SourcePoolAddress.String()
}

func (f fakeObserver) Close() error {
	return nil
}
