package usdc

import (
	"context"
	"fmt"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/internal/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type USDCCCTPTokenDataObserver struct {
	configs           pluginconfig.USDCCCTPObserverConfig
	supportedTokens   map[string]struct{}
	usdcMessageReader reader.USDCMessageReader
	attestationClient AttestationClient
}

func NewUSDCCCTPTokenDataObserver(configs pluginconfig.USDCCCTPObserverConfig) *USDCCCTPTokenDataObserver {
	supportedTokens := make(map[string]struct{})
	for chainSelector, tokenConfig := range configs.Tokens {
		key := sourceTokenIdentifier(chainSelector, tokenConfig.SourcePoolAddress)
		supportedTokens[key] = struct{}{}
	}

	return &USDCCCTPTokenDataObserver{
		configs:         configs,
		supportedTokens: supportedTokens,
		attestationClient: NewSequentialAttestationClient(
			configs.AttestationAPI,
			configs.AttestationAPITimeout.Duration(),
		),
	}
}

func (u *USDCCCTPTokenDataObserver) Observe(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	// 1. Pick only messages that contain USDC tokens
	usdcMessages := u.pickOnlyUSDCMessages(messages)

	// 2. Fetch USDC messages hashes based on the `MessageSent (bytes message)` event
	usdcMessageHashes, err := u.fetchUSDCMessageHashes(ctx, usdcMessages)
	if err != nil {
		return nil, err
	}

	// 3. Fetch attestations for USDC messages
	attestations, err := u.fetchAttestations(ctx, usdcMessageHashes)
	if err != nil {
		return nil, err
	}

	// 4. Add attestations to the token observations
	return u.extractTokenData(attestations)
}

func (u *USDCCCTPTokenDataObserver) IsTokenSupported(
	sourceChain cciptypes.ChainSelector,
	msgToken cciptypes.RampTokenAmount,
) bool {
	_, ok := u.supportedTokens[sourceTokenIdentifier(sourceChain, msgToken.SourcePoolAddress.String())]
	return ok
}

type usdcTokenData struct {
	Index       int
	MessageHash [32]byte
}

func (u *USDCCCTPTokenDataObserver) pickOnlyUSDCMessages(
	messageObservations exectypes.MessageObservations,
) map[cciptypes.ChainSelector]map[cciptypes.SeqNum][]usdcTokenData {
	usdcMessages := make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum][]usdcTokenData)

	for chainSelector, messages := range messageObservations {
		for seqNum, message := range messages {
			var usdcTokens []usdcTokenData
			for i, tokenAmount := range message.TokenAmounts {
				tokenIdentifier := sourceTokenIdentifier(chainSelector, tokenAmount.SourcePoolAddress.String())
				if _, ok := u.supportedTokens[tokenIdentifier]; ok {
					usdcTokens = append(usdcTokens, usdcTokenData{Index: i})
				}
			}
			if len(usdcTokens) > 0 {
				usdcMessages[chainSelector][seqNum] = usdcTokens
			}
		}
	}
	return usdcMessages
}

func (u *USDCCCTPTokenDataObserver) fetchUSDCMessageHashes(
	ctx context.Context,
	usdcMessages map[cciptypes.ChainSelector]map[cciptypes.SeqNum][]usdcTokenData,
) (map[cciptypes.ChainSelector]map[cciptypes.SeqNum][]usdcTokenData, error) {
	for chainSelector, messages := range usdcMessages {
		// TODO Sequential reading USDC messages from the source chain
		usdcHashes, err := u.usdcMessageReader.MessageHashes(ctx, chainSelector, maps.Keys(messages))
		if err != nil {
			return nil, err
		}

		for seqNum, hashes := range usdcHashes {
			if len(messages[seqNum]) != len(hashes) {
				return nil, fmt.Errorf("message hash count mismatch")
			}

			for i, hash := range hashes {
				messages[seqNum][i].MessageHash = hash
			}
		}
	}
	return usdcMessages, nil
}

func (u *USDCCCTPTokenDataObserver) fetchAttestations(
	ctx context.Context,
	usdcMessages map[cciptypes.ChainSelector]map[cciptypes.SeqNum][]usdcTokenData,
) (map[cciptypes.ChainSelector]map[cciptypes.SeqNum]exectypes.TokenData, error) {
	attestations, err := u.attestationClient.Attestations(ctx, usdcMessages)
	if err != nil {
		return nil, err
	}
	return attestations, nil
}

func (u *USDCCCTPTokenDataObserver) extractTokenData(_ interface{}) (exectypes.TokenDataObservations, error) {
	panic("implement me")
}

func sourceTokenIdentifier(chainSelector cciptypes.ChainSelector, sourcePoolAddress string) string {
	return fmt.Sprintf("%d-%s", chainSelector, sourcePoolAddress)
}
