package usdc

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type TokenDataObserver struct {
	lggr               logger.Logger
	destChainSelector  cciptypes.ChainSelector
	supportedTokens    map[string]struct{}
	usdcMessageReader  reader.USDCMessageReader
	attestationClient  AttestationClient
	attestationEncoder AttestationEncoder
}

func NewTokenDataObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	tokens map[cciptypes.ChainSelector]pluginconfig.USDCCCTPTokenConfig,
	attsetationEncoder AttestationEncoder,
	usdcMessageReader reader.USDCMessageReader,
	attestationClient AttestationClient,
) *TokenDataObserver {
	supportedPoolsBySelector := make(map[string]string)
	supportedTokens := make(map[string]struct{})
	for chainSelector, tokenConfig := range tokens {
		key := sourceTokenIdentifier(chainSelector, tokenConfig.SourcePoolAddress)
		supportedTokens[key] = struct{}{}
		supportedPoolsBySelector[chainSelector.String()] = tokenConfig.SourcePoolAddress
	}
	lggr.Infow("Created USDC Token Data Observer",
		"supportedTokenPools", supportedPoolsBySelector,
	)

	return &TokenDataObserver{
		lggr:               lggr,
		destChainSelector:  destChainSelector,
		supportedTokens:    supportedTokens,
		usdcMessageReader:  usdcMessageReader,
		attestationClient:  attestationClient,
		attestationEncoder: attsetationEncoder,
	}
}

func (u *TokenDataObserver) Observe(
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
	return u.extractTokenData(ctx, messages, attestations)
}

func (u *TokenDataObserver) IsTokenSupported(
	sourceChain cciptypes.ChainSelector,
	msgToken cciptypes.RampTokenAmount,
) bool {
	_, ok := u.supportedTokens[sourceTokenIdentifier(sourceChain, msgToken.SourcePoolAddress.String())]
	return ok
}

func (u *TokenDataObserver) Close() error {
	return nil
}

func (u *TokenDataObserver) pickOnlyUSDCMessages(
	messageObservations exectypes.MessageObservations,
) map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.RampTokenAmount {
	usdcMessages := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.RampTokenAmount)
	for chainSelector, messages := range messageObservations {
		usdcMessages[chainSelector] = make(map[reader.MessageTokenID]cciptypes.RampTokenAmount)
		for seqNum, message := range messages {
			for i, tokenAmount := range message.TokenAmounts {
				tokenIdentifier := sourceTokenIdentifier(chainSelector, tokenAmount.SourcePoolAddress.String())
				_, ok := u.supportedTokens[tokenIdentifier]
				if ok {
					usdcMessages[chainSelector][reader.NewMessageTokenID(seqNum, i)] = tokenAmount
				}

				u.lggr.Debugw(
					"Scanning message's tokens for USDC data",
					"isUSDC", ok,
					"seqNum", seqNum,
					"sourceChainSelector", chainSelector,
					"sourcePoolAddress", tokenAmount.SourcePoolAddress.String(),
					"destTokenAddress", tokenAmount.DestTokenAddress.String(),
				)
			}
		}
	}
	return usdcMessages
}

func (u *TokenDataObserver) fetchUSDCMessageHashes(
	ctx context.Context,
	usdcMessages map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.RampTokenAmount,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes, error) {
	output := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes)

	for chainSelector, messages := range usdcMessages {
		if len(messages) == 0 {
			continue
		}

		// TODO Sequential reading USDC messages from the source chain
		usdcHashes, err := u.usdcMessageReader.MessageHashes(ctx, chainSelector, u.destChainSelector, messages)
		if err != nil {
			u.lggr.Errorw(
				"Failed fetching USDC events from the source chain",
				"sourceChainSelector", chainSelector,
				"destChainSelector", u.destChainSelector,
				"messageTokenIDs", maps.Keys(messages),
				"error", err,
			)
			return nil, err
		}
		output[chainSelector] = usdcHashes
	}
	return output, nil
}

func (u *TokenDataObserver) fetchAttestations(
	ctx context.Context,
	usdcMessages map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus, error) {
	attestations, err := u.attestationClient.Attestations(ctx, usdcMessages)
	if err != nil {
		return nil, err
	}
	return attestations, nil
}

func (u *TokenDataObserver) extractTokenData(
	ctx context.Context,
	messages exectypes.MessageObservations,
	attestations map[cciptypes.ChainSelector]map[reader.MessageTokenID]AttestationStatus,
) (exectypes.TokenDataObservations, error) {
	if attestations == nil {
		// TODO: is this an error?
		u.lggr.Warnw("extractTokenData given a nil attestations map.")
		return exectypes.TokenDataObservations{}, nil
	}

	tokenObservations := make(exectypes.TokenDataObservations)

	for chainSelector, chainMessages := range messages {
		tokenObservations[chainSelector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

		for seqNum, message := range chainMessages {
			tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
			for i, tokenAmount := range message.TokenAmounts {
				if !u.IsTokenSupported(chainSelector, tokenAmount) {
					u.lggr.Debugw(
						"Ignoring unsupported token",
						"seqNum", seqNum,
						"sourceChainSelector", chainSelector,
						"sourcePoolAddress", tokenAmount.SourcePoolAddress.String(),
						"destTokenAddress", tokenAmount.DestTokenAddress.String(),
					)
					tokenData[i] = exectypes.NotSupportedTokenData()
				} else {
					tokenData[i] = u.attestationToTokenData(ctx, seqNum, i, attestations[chainSelector])
				}
			}

			tokenObservations[chainSelector][seqNum] = exectypes.NewMessageTokenData(tokenData...)
		}
	}
	return tokenObservations, nil
}

func (u *TokenDataObserver) attestationToTokenData(
	ctx context.Context,
	seqNr cciptypes.SeqNum,
	tokenIndex int,
	attestations map[reader.MessageTokenID]AttestationStatus,
) exectypes.TokenData {
	if attestations == nil {
		// TODO: is this an error?
		u.lggr.Warnw("attestationToTokenData given a nil attestations map.")
		return exectypes.NewErrorTokenData(fmt.Errorf("missing token status"))
	}
	status, ok := attestations[reader.NewMessageTokenID(seqNr, tokenIndex)]
	if !ok {
		return exectypes.NewErrorTokenData(ErrDataMissing)
	}
	if status.Error != nil {
		return exectypes.NewErrorTokenData(status.Error)
	}
	tokenData, err := u.attestationEncoder(ctx, status.MessageHash, status.Attestation)
	if err != nil {
		return exectypes.NewErrorTokenData(fmt.Errorf("unable to encode attestation: %w", err))
	}
	return exectypes.NewSuccessTokenData(tokenData)
}

func sourceTokenIdentifier(chainSelector cciptypes.ChainSelector, sourcePoolAddress string) string {
	// Lowercase the sourcePoolAddress to make the identifier case-insensitive.
	// It makes the code immune to the differences between checksum and non-checksum addresses.
	return fmt.Sprintf("%d-%s", chainSelector, strings.ToLower(sourcePoolAddress))
}
