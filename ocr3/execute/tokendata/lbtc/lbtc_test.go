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

	exectypes2 "github.com/smartcontractkit/chainlink-ccip/ocr3/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/execute/tokendata/lbtc"
	"github.com/smartcontractkit/chainlink-ccip/ocr3/internal"
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
		messageObservations exectypes2.MessageObservations
		attestationClient   tokendata.AttestationClient
		expectedTokenData   exectypes2.TokenDataObservations
	}{
		{
			name:                "no messages",
			messageObservations: exectypes2.MessageObservations{},
			expectedTokenData:   exectypes2.TokenDataObservations{},
			attestationClient:   &tokendata.FakeAttestationClient{},
		},
		{
			name: "no LBTC messages",
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
			attestationClient: &tokendata.FakeAttestationClient{},
		},
		{
			name: "single LBTC message per chain with non 32 bytes extra data",
			messageObservations: exectypes2.MessageObservations{
				1: {
					10: MessageWithTokensAndExtraData16(t, ethereumLBTCPool),
				},
				2: {
					12: MessageWithTokensAndExtraData16(t, avalancheLBTCPool),
				},
			},
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(exectypes2.NotSupportedTokenData()),
				},
				2: {
					12: exectypes2.NewMessageTokenData(exectypes2.NotSupportedTokenData()),
				},
			},
			attestationClient: &tokendata.FakeAttestationClient{},
		},
		{
			name: "single LBTC message per chain",
			messageObservations: exectypes2.MessageObservations{
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
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					10: exectypes2.NewMessageTokenData(newReadyTokenData([]byte{10_0})),
				},
				2: {
					12: exectypes2.NewMessageTokenData(newReadyTokenData([]byte{12_0})),
				},
			},
		},
		{
			name: "LBTC messages mixed with regular within a single msg chain",
			messageObservations: exectypes2.MessageObservations{
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
			expectedTokenData: exectypes2.TokenDataObservations{
				1: {
					9: exectypes2.NewMessageTokenData(
						exectypes2.NotSupportedTokenData(),
						exectypes2.NotSupportedTokenData(),
					),
					10: exectypes2.NewMessageTokenData(
						newReadyTokenData([]byte{10_0}),
						exectypes2.NotSupportedTokenData(),
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
			name: "multiple LBTC transfer in a single message",
			messageObservations: exectypes2.MessageObservations{
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
					10: MessageWithTokensAndExtraData32(t, ethereumLBTCPool, ethereumLBTCPool, internal.RandBytes().String()),
				},
			},
			attestationClient: &tokendata.FakeAttestationClient{
				Data: map[string]tokendata.AttestationStatus{
					string(bytes32From(ethereumLBTCPool, 0)): {Attestation: []byte{10_0}},
					string(bytes32From(ethereumLBTCPool, 1)): {Error: tokendata.ErrNotReady},
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

func newReadyTokenData(data []byte) exectypes2.TokenData {
	return exectypes2.TokenData{
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
