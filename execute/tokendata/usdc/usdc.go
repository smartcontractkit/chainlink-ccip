package usdc

import (
	"context"
	"fmt"
	"strings"

	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pkg/logutil"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

const (
	USDCToken = "USDC"
)

type USDCTokenDataObserver struct {
	lggr                     logger.Logger
	destChainSelector        cciptypes.ChainSelector
	supportedPoolsBySelector map[cciptypes.ChainSelector]string
	attestationEncoder       AttestationEncoder
	usdcMessageReader        reader.USDCMessageReader
	attestationClient        tokendata.AttestationClient
}

func NewUSDCTokenDataObserver(
	ctx context.Context,
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	usdcConfig pluginconfig.USDCCCTPObserverConfig,
	attestationEncoder AttestationEncoder,
	readers map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
	addrCodec cciptypes.AddressCodec,
) (*USDCTokenDataObserver, error) {
	usdcReader, err := reader.NewUSDCMessageReader(
		ctx,
		lggr,
		usdcConfig.Tokens,
		readers,
		addrCodec,
	)
	if err != nil {
		return nil, err
	}
	attestationClient, err := NewSequentialAttestationClient(lggr, usdcConfig)
	if err != nil {
		return nil, fmt.Errorf("create attestation client: %w", err)
	}
	supportedPoolsBySelector := make(map[cciptypes.ChainSelector]string)
	for chainSelector, tokenConfig := range usdcConfig.Tokens {
		supportedPoolsBySelector[chainSelector] = tokenConfig.SourcePoolAddress
	}
	lggr.Infow("Created USDC Token Data Observer",
		"supportedTokenPools", supportedPoolsBySelector,
	)
	return &USDCTokenDataObserver{
		lggr:                     lggr,
		destChainSelector:        destChainSelector,
		supportedPoolsBySelector: supportedPoolsBySelector,
		attestationEncoder:       attestationEncoder,
		usdcMessageReader:        usdcReader,
		attestationClient:        attestationClient,
	}, nil
}

func InitUSDCTokenDataObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	supportedPoolsBySelector map[cciptypes.ChainSelector]string,
	attestationEncoder AttestationEncoder,
	usdcMessageReader reader.USDCMessageReader,
	attestationClient tokendata.AttestationClient,
) *USDCTokenDataObserver {
	return &USDCTokenDataObserver{
		lggr:                     lggr,
		destChainSelector:        destChainSelector,
		supportedPoolsBySelector: supportedPoolsBySelector,
		attestationEncoder:       attestationEncoder,
		usdcMessageReader:        usdcMessageReader,
		attestationClient:        attestationClient,
	}
}

func (u *USDCTokenDataObserver) Observe(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	lggr := logutil.WithContextValues(ctx, u.lggr)

	// 1. Pick only messages that contain USDC tokens
	usdcMessages := u.pickOnlyUSDCMessages(lggr, messages)

	// 2. Fetch USDC messages by token id based on the `MessageSent (bytes message)` event
	usdcMessagesByTokenID, err := u.fetchUSDCEventMessages(ctx, lggr, usdcMessages)
	if err != nil {
		return nil, err
	}

	// 3. Fetch attestations for USDC messages
	attestations, err := u.fetchAttestations(ctx, usdcMessagesByTokenID)
	if err != nil {
		return nil, err
	}

	// 4. Add attestations to the token observations
	return u.extractTokenData(ctx, lggr, messages, attestations)
}

func (u *USDCTokenDataObserver) IsTokenSupported(
	sourceChain cciptypes.ChainSelector,
	msgToken cciptypes.RampTokenAmount,
) bool {
	return strings.EqualFold(u.supportedPoolsBySelector[sourceChain], msgToken.SourcePoolAddress.String())
}

func (u *USDCTokenDataObserver) Close() error {
	return nil
}

func (u *USDCTokenDataObserver) pickOnlyUSDCMessages(
	lggr logger.Logger,
	messageObservations exectypes.MessageObservations,
) map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.RampTokenAmount {
	usdcMessages := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.RampTokenAmount)
	for chainSelector, messages := range messageObservations {
		usdcMessages[chainSelector] = make(map[reader.MessageTokenID]cciptypes.RampTokenAmount)
		for seqNum, message := range messages {
			for i, tokenAmount := range message.TokenAmounts {
				isUSDC := u.IsTokenSupported(chainSelector, tokenAmount)
				if isUSDC {
					usdcMessages[chainSelector][reader.NewMessageTokenID(seqNum, i)] = tokenAmount
				}
				lggr.Debugw(
					"Scanning message's tokens for USDC data",
					"isUSDC", isUSDC,
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

func (u *USDCTokenDataObserver) fetchUSDCEventMessages(
	ctx context.Context,
	lggr logger.Logger,
	usdcMessages map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.RampTokenAmount,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes, error) {
	output := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes)

	for chainSelector, messages := range usdcMessages {
		if len(messages) == 0 {
			continue
		}

		// TODO Sequential reading USDC messages from the source chain
		msgByTokenID, err := u.usdcMessageReader.MessagesByTokenID(ctx, chainSelector, u.destChainSelector, messages)
		if err != nil {
			lggr.Errorw(
				"Failed fetching USDC events from the source chain",
				"sourceChainSelector", chainSelector,
				"destChainSelector", u.destChainSelector,
				"messageTokenIDs", maps.Keys(messages),
				"error", err,
			)
			return nil, err
		}
		output[chainSelector] = msgByTokenID
	}
	return output, nil
}

func (u *USDCTokenDataObserver) fetchAttestations(
	ctx context.Context,
	usdcMessages map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes,
) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus, error) {
	attestations, err := u.attestationClient.Attestations(ctx, usdcMessages)
	if err != nil {
		return nil, err
	}
	return attestations, nil
}

func (u *USDCTokenDataObserver) extractTokenData(
	ctx context.Context,
	lggr logger.Logger,
	messages exectypes.MessageObservations,
	attestations map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus,
) (exectypes.TokenDataObservations, error) {
	tokenObservations := make(exectypes.TokenDataObservations)

	for chainSelector, chainMessages := range messages {
		tokenObservations[chainSelector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

		for seqNum, message := range chainMessages {
			tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
			for i, tokenAmount := range message.TokenAmounts {
				if !u.IsTokenSupported(chainSelector, tokenAmount) {
					lggr.Debugw(
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

func (u *USDCTokenDataObserver) attestationToTokenData(
	ctx context.Context,
	seqNr cciptypes.SeqNum,
	tokenIndex int,
	attestations map[reader.MessageTokenID]tokendata.AttestationStatus,
) exectypes.TokenData {
	status, ok := attestations[reader.NewMessageTokenID(seqNr, tokenIndex)]
	if !ok {
		return exectypes.NewErrorTokenData(tokendata.ErrDataMissing)
	}
	if status.Error != nil {
		return exectypes.NewErrorTokenData(status.Error)
	}
	tokenData, err := u.attestationEncoder(ctx, status.MessageBody, status.Attestation)
	if err != nil {
		return exectypes.NewErrorTokenData(fmt.Errorf("unable to encode attestation: %w", err))
	}
	return exectypes.NewSuccessTokenData(tokenData)
}
