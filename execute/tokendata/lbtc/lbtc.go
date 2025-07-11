package lbtc

import (
	"context"
	"fmt"
	"strings"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type LBTCTokenDataObserver struct {
	lggr                     logger.Logger
	destChainSelector        cciptypes.ChainSelector
	supportedPoolsBySelector map[cciptypes.ChainSelector]string
	client                   tokendata.AttestationClient
}

func NewLBTCTokenDataObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	config pluginconfig.LBTCObserverConfig,
) (*LBTCTokenDataObserver, error) {
	client, err := NewLBTCAttestationClient(lggr, config)
	if err != nil {
		return nil, fmt.Errorf("create attestation client: %w", err)
	}
	return &LBTCTokenDataObserver{
		lggr:                     lggr,
		destChainSelector:        destChainSelector,
		supportedPoolsBySelector: config.SourcePoolAddressByChain,
		client:                   client,
	}, nil
}

func InitLBTCTokenDataObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	supportedPoolsBySelector map[cciptypes.ChainSelector]string,
	client tokendata.AttestationClient,
) *LBTCTokenDataObserver {
	return &LBTCTokenDataObserver{
		lggr:                     lggr,
		destChainSelector:        destChainSelector,
		supportedPoolsBySelector: supportedPoolsBySelector,
		client:                   client,
	}
}

func (o *LBTCTokenDataObserver) Observe(
	ctx context.Context,
	observations exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	// 1. Pick LBTC messages
	lbtcMessages := o.pickOnlyLBTCMessages(observations)
	// 2. Request attestations
	attestations, err := o.client.Attestations(ctx, lbtcMessages)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch attestations: %w", err)
	}
	// 3. Map to result
	return o.createTokenDataObservations(observations, attestations)
}

// IsTokenSupported returns true if the token is supported by the observer.
func (o *LBTCTokenDataObserver) IsTokenSupported(
	sourceChain cciptypes.ChainSelector,
	msgToken cciptypes.RampTokenAmount,
) bool {
	return strings.EqualFold(o.supportedPoolsBySelector[sourceChain], msgToken.SourcePoolAddress.String()) &&
		len(msgToken.ExtraData) == 32
}

// Close closes the observer and releases any resources.
func (o *LBTCTokenDataObserver) Close() error {
	return nil
}

func (o *LBTCTokenDataObserver) pickOnlyLBTCMessages(
	messageObservations exectypes.MessageObservations,
) map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes {
	lbtcMessages := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.Bytes)
	for chainSelector, messages := range messageObservations {
		lbtcMessages[chainSelector] = make(map[reader.MessageTokenID]cciptypes.Bytes)
		for seqNum, message := range messages {
			for i, tokenAmount := range message.TokenAmounts {
				isLBTC := o.IsTokenSupported(chainSelector, tokenAmount)
				if isLBTC {
					lbtcMessages[chainSelector][reader.NewMessageTokenID(seqNum, i)] = tokenAmount.ExtraData
				}
				o.lggr.Debugw(
					"Scanning message's tokens for LBTC data",
					"isLBTC", isLBTC,
					"seqNum", seqNum,
					"sourceChainSelector", chainSelector,
					"sourcePoolAddress", tokenAmount.SourcePoolAddress.String(),
					"destTokenAddress", tokenAmount.DestTokenAddress.String(),
				)
			}
		}
	}
	return lbtcMessages
}

func (o *LBTCTokenDataObserver) createTokenDataObservations(
	messages exectypes.MessageObservations,
	attestations map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus,
) (exectypes.TokenDataObservations, error) {
	tokenObservations := make(exectypes.TokenDataObservations)
	for chainSelector, chainMessages := range messages {
		tokenObservations[chainSelector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
		for seqNum, message := range chainMessages {
			tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
			for i, tokenAmount := range message.TokenAmounts {
				if !o.IsTokenSupported(chainSelector, tokenAmount) {
					o.lggr.Debugw(
						"Ignoring unsupported token",
						"seqNum", seqNum,
						"sourceChainSelector", chainSelector,
						"sourcePoolAddress", tokenAmount.SourcePoolAddress.String(),
						"destTokenAddress", tokenAmount.DestTokenAddress.String(),
					)
					tokenData[i] = exectypes.NotSupportedTokenData()
				} else {
					tokenData[i] = o.attestationToTokenData(seqNum, i, attestations[chainSelector])
				}
			}
			tokenObservations[chainSelector][seqNum] = exectypes.NewMessageTokenData(tokenData...)
		}
	}
	return tokenObservations, nil
}

func (o *LBTCTokenDataObserver) attestationToTokenData(
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
	return exectypes.NewSuccessTokenData(status.Attestation)
}
