package tokendata_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-common/pkg/utils/tests"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/internal"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

func Test_CompositeTokenDataObserver_EmptyObservers(t *testing.T) {
	obs, err := tokendata.NewConfigBasedCompositeObservers(
		logger.Test(t),
		100,
		[]pluginconfig.TokenDataObserverConfig{},
		nil,
		nil,
	)
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

func Test_CompositeTokenDataObserver_Failures(t *testing.T) {
	linkEthereumTokenSourcePool := internal.RandBytes().String()
	linkAvalancheTokenSourcePool := internal.RandBytes().String()
	usdcEthereumTokenSourcePool := internal.RandBytes().String()

	tests := []struct {
		name                string
		observers           []tokendata.TokenDataObserver
		messageObservations exectypes.MessageObservations
		expectedTokenData   exectypes.TokenDataObservations
	}{
		{
			name: "single observer returns an error but no tokens in messages",
			observers: []tokendata.TokenDataObserver{
				faulty(
					"stLINK",
					map[cciptypes.ChainSelector]string{
						1: linkEthereumTokenSourcePool,
					}),
			},
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
			name: "faulty observer doesn't affect other tokens",
			observers: []tokendata.TokenDataObserver{
				faulty(
					"LINK",
					map[cciptypes.ChainSelector]string{
						1: linkEthereumTokenSourcePool,
					}),
			},
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, linkEthereumTokenSourcePool, internal.RandBytes().String()),
					11: internal.MessageWithTokens(t, internal.RandBytes().String(), internal.RandBytes().String()),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(
						exectypes.TokenData{
							Ready:     false,
							Data:      nil,
							Error:     nil,
							Supported: true,
						},
						exectypes.NewNoopTokenData(),
					),
					11: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
						exectypes.NewNoopTokenData(),
					),
				},
			},
		},
		{
			name: "single observer returns an error for tokens for different chains",
			observers: []tokendata.TokenDataObserver{
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
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, linkEthereumTokenSourcePool),
				},
				2: {
					20: internal.MessageWithTokens(t, linkAvalancheTokenSourcePool),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					// TokenObserver failed to process that
					10: exectypes.NewMessageTokenData(exectypes.TokenData{
						Ready:     false,
						Data:      nil,
						Error:     nil,
						Supported: true,
					}),
				},
				2: {
					20: exectypes.NewMessageTokenData(exectypes.NewSuccessTokenData([]byte("LINK_20_0"))),
				},
			},
		},
		{
			name: "multiple observers return an error for tokens for different chains",
			observers: []tokendata.TokenDataObserver{
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
			messageObservations: exectypes.MessageObservations{
				1: {
					10: internal.MessageWithTokens(t, linkEthereumTokenSourcePool),
					11: internal.MessageWithTokens(t, usdcEthereumTokenSourcePool),
					12: internal.MessageWithTokens(t, internal.RandBytes().String(), linkEthereumTokenSourcePool),
					13: internal.MessageWithTokens(t, internal.RandBytes().String(), usdcEthereumTokenSourcePool),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(exectypes.TokenData{
						Ready:     false,
						Data:      nil,
						Error:     nil,
						Supported: true,
					}),
					11: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte("USDC_11_0")),
					),
					12: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
						exectypes.TokenData{
							Ready:     false,
							Data:      nil,
							Error:     nil,
							Supported: true,
						},
					),
					13: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
						exectypes.NewSuccessTokenData([]byte("USDC_13_1")),
					),
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			composite := tokendata.NewCompositeObservers(logger.Test(t), test.observers...)

			tkData, err := composite.Observe(context.Background(), test.messageObservations)
			require.NoError(t, err)

			require.Equal(t, test.expectedTokenData, tkData)
		})

	}
}

func Test_CompositeTokenDataObserver_ParallelExecution(t *testing.T) {
	linkTokenSourcePool := internal.RandBytes().String()
	usdcTokenSourcePool := internal.RandBytes().String()
	randomTokenSourcePool := internal.RandBytes().String()

	lggr := logger.Test(t)

	t.Cleanup(func() {
		_ = lggr.Sync()
	})

	tt := []struct {
		name               string
		observers          []tokendata.TokenDataObserver
		messageObservation exectypes.MessageObservations
		expectedTokenData  exectypes.TokenDataObservations
		timeout            time.Duration
	}{
		{
			name: "multiple valid observers",
			observers: []tokendata.TokenDataObserver{
				fake("LINK", map[cciptypes.ChainSelector]string{
					1: linkTokenSourcePool,
				}).withDelay(100 * time.Millisecond),
				fake("USDC", map[cciptypes.ChainSelector]string{
					1: usdcTokenSourcePool,
				}).withDelay(200 * time.Millisecond),
				fake("TOKEN", map[cciptypes.ChainSelector]string{
					1: randomTokenSourcePool,
				}).withDelay(300 * time.Millisecond),
			},
			messageObservation: exectypes.MessageObservations{
				1: {
					11: internal.MessageWithTokens(t, usdcTokenSourcePool, randomTokenSourcePool),
					12: internal.MessageWithTokens(t, internal.RandBytes().String()),
					13: internal.MessageWithTokens(t, linkTokenSourcePool),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					11: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte("USDC_11_0")),
						exectypes.NewSuccessTokenData([]byte("TOKEN_11_1")),
					),
					12: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
					),
					13: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte("LINK_13_0")),
					),
				},
			},
		},
		{
			name: "faulty observer doesn't impact observation if there are no matching tokens",
			observers: []tokendata.TokenDataObserver{
				fake("LINK", map[cciptypes.ChainSelector]string{
					1: linkTokenSourcePool,
				}).withDelay(100 * time.Millisecond),
				faulty("USDC", map[cciptypes.ChainSelector]string{
					1: usdcTokenSourcePool,
				}).withDelay(200 * time.Millisecond),
			},
			messageObservation: exectypes.MessageObservations{
				1: {
					11: internal.MessageWithTokens(t, linkTokenSourcePool),
					12: internal.MessageWithTokens(t, internal.RandBytes().String()),
					13: internal.MessageWithTokens(t),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					11: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte("LINK_11_0")),
					),
					12: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
					),
					13: exectypes.NewMessageTokenData(),
				},
			},
		},
		{
			name: "timeouting observer doesn't impact observation for other tokens",
			observers: []tokendata.TokenDataObserver{
				fake("LINK", map[cciptypes.ChainSelector]string{
					1: linkTokenSourcePool,
				}).withDelay(100 * time.Millisecond),
				fake("USDC", map[cciptypes.ChainSelector]string{
					1: usdcTokenSourcePool,
				}).withDelay(200 * time.Millisecond),
				fake("TOKEN", map[cciptypes.ChainSelector]string{
					1: randomTokenSourcePool,
				}).withDelay(1 * time.Minute),
			},
			messageObservation: exectypes.MessageObservations{
				1: {
					11: internal.MessageWithTokens(t, linkTokenSourcePool),
					12: internal.MessageWithTokens(t, internal.RandBytes().String()),
					13: internal.MessageWithTokens(t),
					14: internal.MessageWithTokens(t, usdcTokenSourcePool),
					15: internal.MessageWithTokens(t, randomTokenSourcePool, linkTokenSourcePool),
					16: internal.MessageWithTokens(t, randomTokenSourcePool),
				},
			},
			timeout: 500 * time.Millisecond,
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					11: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte("LINK_11_0")),
					),
					12: exectypes.NewMessageTokenData(
						exectypes.NewNoopTokenData(),
					),
					13: exectypes.NewMessageTokenData(),
					14: exectypes.NewMessageTokenData(
						exectypes.NewSuccessTokenData([]byte("USDC_14_0")),
					),
					// These messages will be filtered out by the plugin
					15: exectypes.NewMessageTokenData(
						exectypes.TokenData{Ready: false, Supported: true},
						exectypes.NewSuccessTokenData([]byte("LINK_15_1")),
					),
					16: exectypes.NewMessageTokenData(
						exectypes.TokenData{Ready: false, Supported: true},
					),
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx := tests.Context(t)
			if tc.timeout > 0 {
				var cancel context.CancelFunc
				ctx, cancel = context.WithTimeout(ctx, tc.timeout)
				defer cancel()
			}

			composite := tokendata.NewCompositeObservers(lggr, tc.observers...)

			tkData, err := composite.Observe(ctx, tc.messageObservation)
			require.NoError(t, err)

			require.Equal(t, tc.expectedTokenData, tkData)
		})
	}

}

func faulty(prefix string, supportedTokens map[cciptypes.ChainSelector]string) *fakeObserver {
	return &fakeObserver{
		prefix:          prefix,
		faulty:          true,
		supportedTokens: supportedTokens,
	}
}

func fake(prefix string, supportedTokens map[cciptypes.ChainSelector]string) *fakeObserver {
	return &fakeObserver{
		prefix:          prefix,
		faulty:          false,
		supportedTokens: supportedTokens,
	}
}

type fakeObserver struct {
	faulty          bool
	prefix          string
	supportedTokens map[cciptypes.ChainSelector]string
	sleep           time.Duration
}

func (f *fakeObserver) withDelay(sleep time.Duration) *fakeObserver {
	f.sleep = sleep
	return f
}

func (f *fakeObserver) Observe(
	_ context.Context,
	observations exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	if f.sleep > 0 {
		time.Sleep(f.sleep)
	}

	if f.faulty {
		return nil, fmt.Errorf("error")
	}

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

func (f *fakeObserver) IsTokenSupported(sourceChain cciptypes.ChainSelector, msgToken cciptypes.RampTokenAmount) bool {
	tokenAddr, ok := f.supportedTokens[sourceChain]
	return ok && tokenAddr == msgToken.SourcePoolAddress.String()
}
