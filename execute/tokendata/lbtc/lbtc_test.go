package lbtc_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink-common/pkg/hashutil"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/lbtc"
	"github.com/smartcontractkit/chainlink-ccip/internal"
)

func TestTokenDataObserver_Observe_LBTCAndRegularTokens(t *testing.T) {
	ethereumLBTCPool := internal.RandBytes().String()
	avalancheLBTCPool := internal.RandBytes().String()
	supportedPoolsBySelector := map[cciptypes.ChainSelector]string{
		cciptypes.ChainSelector(1): ethereumLBTCPool,
		cciptypes.ChainSelector(2): avalancheLBTCPool,
	}
	messageWithTokensAndExtraDataMixedSize :=
		MessageWithTokensAndExtraData32(t, ethereumLBTCPool, internal.RandBytes().String())
	messageWithTokensAndExtraDataMixedSize.TokenAmounts[0].ExtraData =
		messageWithTokensAndExtraDataMixedSize.TokenAmounts[0].ExtraData[:16]
	tests := []struct {
		name                string
		messageObservations exectypes.MessageObservations
		attestationClient   tokendata.AttestationClient
		expectedTokenData   exectypes.TokenDataObservations
	}{
		{
			name:                "no messages",
			messageObservations: exectypes.MessageObservations{},
			expectedTokenData:   exectypes.TokenDataObservations{},
			attestationClient:   &tokendata.FakeAttestationClient{},
		},
		{
			name: "no LBTC messages",
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
			attestationClient: &tokendata.FakeAttestationClient{},
		},
		{
			name: "single LBTC message per chain with non 32 bytes extra data",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: MessageWithTokensAndExtraData16(t, ethereumLBTCPool),
				},
				2: {
					12: MessageWithTokensAndExtraData16(t, avalancheLBTCPool),
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(exectypes.NotSupportedTokenData()),
				},
				2: {
					12: exectypes.NewMessageTokenData(exectypes.NotSupportedTokenData()),
				},
			},
			attestationClient: &tokendata.FakeAttestationClient{},
		},
		{
			name: "single LBTC message per chain",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: MessageWithTokensAndExtraData32(t, ethereumLBTCPool),
				},
				2: {
					12: MessageWithTokensAndExtraData32(t, avalancheLBTCPool),
				},
			},
			attestationClient: &tokendata.FakeAttestationClient{
				Data: map[string]tokendata.AttestationStatus{
					string(bytes32From(ethereumLBTCPool, 0)):  {Attestation: []byte{10_0}},
					string(bytes32From(avalancheLBTCPool, 0)): {Attestation: []byte{12_0}},
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(newReadyTokenData([]byte{10_0})),
				},
				2: {
					12: exectypes.NewMessageTokenData(newReadyTokenData([]byte{12_0})),
				},
			},
		},
		{
			name: "LBTC messages mixed with regular within a single msg chain",
			messageObservations: exectypes.MessageObservations{
				1: {
					9:  messageWithTokensAndExtraDataMixedSize,
					10: MessageWithTokensAndExtraData32(t, ethereumLBTCPool, internal.RandBytes().String()),
					11: MessageWithTokensAndExtraData32(
						t, internal.RandBytes().String(), ethereumLBTCPool, internal.RandBytes().String(),
					),
					12: MessageWithTokensAndExtraData32(
						t, internal.RandBytes().String(), internal.RandBytes().String(), ethereumLBTCPool,
					),
					13: internal.MessageWithTokens(t),
				},
			},
			attestationClient: &tokendata.FakeAttestationClient{
				Data: map[string]tokendata.AttestationStatus{
					string(bytes32From(ethereumLBTCPool, 0)): {Attestation: []byte{10_0}},
					string(bytes32From(ethereumLBTCPool, 1)): {Attestation: []byte{11_1}},
					string(bytes32From(ethereumLBTCPool, 2)): {Attestation: []byte{12_2}},
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					9: exectypes.NewMessageTokenData(
						exectypes.NotSupportedTokenData(),
						exectypes.NotSupportedTokenData(),
					),
					10: exectypes.NewMessageTokenData(
						newReadyTokenData([]byte{10_0}),
						exectypes.NotSupportedTokenData(),
					),
					11: exectypes.NewMessageTokenData(
						exectypes.NotSupportedTokenData(),
						newReadyTokenData([]byte{11_1}),
						exectypes.NotSupportedTokenData(),
					),
					12: exectypes.NewMessageTokenData(
						exectypes.NotSupportedTokenData(),
						exectypes.NotSupportedTokenData(),
						newReadyTokenData([]byte{12_2}),
					),
					13: exectypes.NewMessageTokenData(),
				},
			},
		},
		{
			name: "multiple LBTC transfer in a single message",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: MessageWithTokensAndExtraData32(t, ethereumLBTCPool, ethereumLBTCPool, ethereumLBTCPool),
				},
				2: {
					12: MessageWithTokensAndExtraData32(t, avalancheLBTCPool, avalancheLBTCPool),
				},
			},
			attestationClient: &tokendata.FakeAttestationClient{
				Data: map[string]tokendata.AttestationStatus{
					string(bytes32From(ethereumLBTCPool, 0)):  {Attestation: []byte{10_0}},
					string(bytes32From(ethereumLBTCPool, 1)):  {Attestation: []byte{10_1}},
					string(bytes32From(ethereumLBTCPool, 2)):  {Attestation: []byte{10_2}},
					string(bytes32From(avalancheLBTCPool, 0)): {Attestation: []byte{12_0}},
					string(bytes32From(avalancheLBTCPool, 1)): {Attestation: []byte{12_1}},
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(
						newReadyTokenData([]byte{10_0}),
						newReadyTokenData([]byte{10_1}),
						newReadyTokenData([]byte{10_2}),
					),
				},
				2: {
					12: exectypes.NewMessageTokenData(
						newReadyTokenData([]byte{12_0}),
						newReadyTokenData([]byte{12_1}),
					),
				},
			},
		},
		{
			name: "not ready messages are populated to the result set",
			messageObservations: exectypes.MessageObservations{
				1: {
					10: MessageWithTokensAndExtraData32(t, ethereumLBTCPool, ethereumLBTCPool, internal.RandBytes().String()),
				},
			},
			attestationClient: &tokendata.FakeAttestationClient{
				Data: map[string]tokendata.AttestationStatus{
					string(bytes32From(ethereumLBTCPool, 0)): {Attestation: []byte{10_0}},
					string(bytes32From(ethereumLBTCPool, 1)): {Error: tokendata.ErrNotReady},
				},
			},
			expectedTokenData: exectypes.TokenDataObservations{
				1: {
					10: exectypes.NewMessageTokenData(
						newReadyTokenData([]byte{10_0}),
						exectypes.TokenData{
							Ready:     false,
							Data:      nil,
							Error:     tokendata.ErrNotReady,
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
			observer := lbtc.InitLBTCTokenDataObserver(
				logger.Test(t),
				1,
				supportedPoolsBySelector,
				test.attestationClient,
			)

			tkData, err := observer.Observe(context.Background(), test.messageObservations)
			require.NoError(t, err)

			assert.Equal(t, test.expectedTokenData, tkData)
		})
	}
}

func newReadyTokenData(data []byte) exectypes.TokenData {
	return exectypes.TokenData{
		Ready:     true,
		Error:     nil,
		Data:      data,
		Supported: true,
	}
}

func MessageWithTokensAndExtraData32(t *testing.T, tokenPoolAddr ...string) cciptypes.Message {
	message := internal.MessageWithTokens(t, tokenPoolAddr...)
	for i := range message.TokenAmounts {
		message.TokenAmounts[i].ExtraData = bytes32From(tokenPoolAddr[i], i)
	}
	return message
}

func MessageWithTokensAndExtraData16(t *testing.T, tokenPoolAddr ...string) cciptypes.Message {
	message := internal.MessageWithTokens(t, tokenPoolAddr...)
	for i := range message.TokenAmounts {
		message.TokenAmounts[i].ExtraData = internal.RandBytes()[:16]
	}
	return message
}

func bytes32From(address string, idx int) []byte {
	hasher := hashutil.NewKeccak()
	hash := hasher.Hash([]byte(fmt.Sprintf("%s%d", address, idx)))
	return hash[:]
}
