package usdc

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type tokenDataObserver struct {
	lggr              logger.Logger
	destChainSelector cciptypes.ChainSelector
	supportedTokens   map[string]struct{}
	usdcMessageReader reader.USDCMessageReader
	attestationClient AttestationClient
}

//nolint:revive
func NewTokenDataObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	tokens map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig,
	usdcMessageReader reader.USDCMessageReader,
	attestationClient AttestationClient,
) *tokenDataObserver {
	supportedTokens := make(map[string]struct{})
	for chainSelector, tokenConfig := range tokens {
		key := sourceTokenIdentifier(chainSelector, tokenConfig.SourcePoolAddress)
		supportedTokens[key] = struct{}{}
	}

	return &tokenDataObserver{
		lggr:              lggr,
		destChainSelector: destChainSelector,
		supportedTokens:   supportedTokens,
		usdcMessageReader: usdcMessageReader,
		attestationClient: attestationClient,
	}
}

func (u *tokenDataObserver) Observe(
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

func (u *tokenDataObserver) IsTokenSupported(
	sourceChain cciptypes.ChainSelector,
	msgToken cciptypes.RampTokenAmount,
) bool {
	_, ok := u.supportedTokens[sourceTokenIdentifier(sourceChain, msgToken.SourcePoolAddress.String())]
	return ok
}

func (u *tokenDataObserver) pickOnlyUSDCMessages(
	messageObservations exectypes.MessageObservations,
) map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]cciptypes.RampTokenAmount {
	usdcMessages := make(map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]cciptypes.RampTokenAmount)
	for chainSelector, messages := range messageObservations {
		usdcMessages[chainSelector] = make(map[exectypes.MessageTokenID]cciptypes.RampTokenAmount)
		for seqNum, message := range messages {
			for i, tokenAmount := range message.TokenAmounts {
				tokenIdentifier := sourceTokenIdentifier(chainSelector, tokenAmount.SourcePoolAddress.String())
				if _, ok := u.supportedTokens[tokenIdentifier]; !ok {
					continue
				}
				usdcMessages[chainSelector][exectypes.NewMessageTokenID(seqNum, i)] = tokenAmount
			}
		}
	}
	return usdcMessages
}

func (u *tokenDataObserver) fetchUSDCMessageHashes(
	ctx context.Context,
	usdcMessages map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]cciptypes.RampTokenAmount,
) (map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash, error) {
	output := make(map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash)

	for chainSelector, messages := range usdcMessages {
		if len(messages) == 0 {
			continue
		}

		// TODO Sequential reading USDC messages from the source chain
		usdcHashes, err := u.usdcMessageReader.MessageHashes(ctx, chainSelector, u.destChainSelector, messages)
		if err != nil {
			return nil, err
		}
		output[chainSelector] = usdcHashes
	}
	return output, nil
}

func (u *tokenDataObserver) fetchAttestations(
	ctx context.Context,
	usdcMessages map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]reader.MessageHash,
) (map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus, error) {
	attestations, err := u.attestationClient.Attestations(ctx, usdcMessages)
	if err != nil {
		return nil, err
	}
	return attestations, nil
}

func (u *tokenDataObserver) extractTokenData(
	messages exectypes.MessageObservations,
	attestations map[cciptypes.ChainSelector]map[exectypes.MessageTokenID]AttestationStatus,
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
					tokenData[i] = attestationToTokenData(seqNum, i, attestations[chainSelector])
				}
			}

			tokenObservations[chainSelector][seqNum] = exectypes.NewMessageTokenData(tokenData...)
		}
	}
	return tokenObservations, nil
}

func attestationToTokenData(
	seqNr cciptypes.SeqNum,
	tokenIndex int,
	attestations map[exectypes.MessageTokenID]AttestationStatus,
) exectypes.TokenData {
	status, ok := attestations[exectypes.NewMessageTokenID(seqNr, tokenIndex)]
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
