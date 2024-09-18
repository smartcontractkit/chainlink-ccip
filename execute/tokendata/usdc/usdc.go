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
	return NewUSDCCCTP(
		configs,
		nil,
		nil,
	)
}

func NewUSDCCCTP(
	configs pluginconfig.USDCCCTPObserverConfig,
	usdcMessageReader reader.USDCMessageReader,
	attestationClient AttestationClient,
) *USDCCCTPTokenDataObserver {
	supportedTokens := make(map[string]struct{})
	for chainSelector, tokenConfig := range configs.Tokens {
		key := sourceTokenIdentifier(chainSelector, tokenConfig.SourcePoolAddress)
		supportedTokens[key] = struct{}{}
	}

	return &USDCCCTPTokenDataObserver{
		configs:           configs,
		supportedTokens:   supportedTokens,
		usdcMessageReader: usdcMessageReader,
		attestationClient: attestationClient,
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
	return u.extractTokenData(messages, attestations)
}

func (u *USDCCCTPTokenDataObserver) IsTokenSupported(
	sourceChain cciptypes.ChainSelector,
	msgToken cciptypes.RampTokenAmount,
) bool {
	_, ok := u.supportedTokens[sourceTokenIdentifier(sourceChain, msgToken.SourcePoolAddress.String())]
	return ok
}

func (u *USDCCCTPTokenDataObserver) pickOnlyUSDCMessages(
	messageObservations exectypes.MessageObservations,
) map[cciptypes.ChainSelector]map[cciptypes.SeqNum][]int {
	usdcMessages := make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum][]int)
	for chainSelector, messages := range messageObservations {
		usdcMessages[chainSelector] = make(map[cciptypes.SeqNum][]int)
		for seqNum, message := range messages {
			var usdcTokens []int
			for i, tokenAmount := range message.TokenAmounts {
				tokenIdentifier := sourceTokenIdentifier(chainSelector, tokenAmount.SourcePoolAddress.String())
				if _, ok := u.supportedTokens[tokenIdentifier]; ok {
					usdcTokens = append(usdcTokens, i)
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
	usdcMessages map[cciptypes.ChainSelector]map[cciptypes.SeqNum][]int,
) (map[cciptypes.ChainSelector]map[cciptypes.SeqNum]map[int][]byte, error) {
	output := make(map[cciptypes.ChainSelector]map[cciptypes.SeqNum]map[int][]byte)

	for chainSelector, messages := range usdcMessages {
		if len(messages) == 0 {
			continue
		}
		// TODO Sequential reading USDC messages from the source chain
		usdcHashes, err := u.usdcMessageReader.MessageHashes(ctx, chainSelector, maps.Keys(messages))
		if err != nil {
			return nil, err
		}
		output[chainSelector] = usdcHashes
	}
	return output, nil
}

func (u *USDCCCTPTokenDataObserver) fetchAttestations(
	ctx context.Context,
	usdcMessages map[cciptypes.ChainSelector]map[cciptypes.SeqNum]map[int][]byte,
) (map[cciptypes.ChainSelector]map[cciptypes.SeqNum]map[int]AttestationStatus, error) {
	attestations, err := u.attestationClient.Attestations(ctx, usdcMessages)
	if err != nil {
		return nil, err
	}
	return attestations, nil
}

func (u *USDCCCTPTokenDataObserver) extractTokenData(
	messages exectypes.MessageObservations,
	attestations map[cciptypes.ChainSelector]map[cciptypes.SeqNum]map[int]AttestationStatus,
) (exectypes.TokenDataObservations, error) {
	tokenObservations := make(exectypes.TokenDataObservations)

	for chainSelector, chainMessages := range messages {
		tokenObservations[chainSelector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

		for seqNum, message := range chainMessages {
			tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
			for i, tokenAmount := range message.TokenAmounts {
				if !u.IsTokenSupported(chainSelector, tokenAmount) {
					tokenData[i] = exectypes.NotSupportedTokenData()
				} else {
					tokenData[i] = attestationToTokenData(i, attestations[chainSelector][seqNum])
				}
			}

			tokenObservations[chainSelector][seqNum] = exectypes.NewMessageTokenData(tokenData...)
		}
	}
	return tokenObservations, nil
}

func attestationToTokenData(tokenIndex int, statuses map[int]AttestationStatus) exectypes.TokenData {
	status, ok := statuses[tokenIndex]
	if !ok {
		return exectypes.NewErrorTokenData(ErrDataMissing)
	}
	if status.Error != nil {
		return exectypes.NewErrorTokenData(status.Error)
	}
	return exectypes.NewSuccessTokenData(status.Data[:])
}

func sourceTokenIdentifier(chainSelector cciptypes.ChainSelector, sourcePoolAddress string) string {
	return fmt.Sprintf("%d-%s", chainSelector, sourcePoolAddress)
}
