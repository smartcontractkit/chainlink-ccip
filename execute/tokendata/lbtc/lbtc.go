package lbtc

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"golang.org/x/exp/maps"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
)

type LBTCTokenObserver struct {
	lggr              logger.Logger
	destChainSelector cciptypes.ChainSelector
	config            pluginconfig.LBTCObserverConfig
	client            *LBTCAttestationClient
}

func (o *LBTCTokenObserver) Observe(
	ctx context.Context,
	observations exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	// 1. Pick LBTC messages
	lbtcMessages := o.pickOnlyUSDCMessages(observations)
	// 2. Request attestations
	attestations, err := o.fetchAttestations(ctx, lbtcMessages)
	if err != nil {
		return nil, err
	}
	// 3. Map to result
	return o.createTokenDataObservations(observations, attestations)
}

func (o *LBTCTokenObserver) createTokenDataObservations(messages exectypes.MessageObservations, attestations map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus) (exectypes.TokenDataObservations, error) {
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

func (o *LBTCTokenObserver) attestationToTokenData(
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

// IsTokenSupported returns true if the token is supported by the observer.
func (o *LBTCTokenObserver) IsTokenSupported(sourceChain cciptypes.ChainSelector, msgToken cciptypes.RampTokenAmount) bool {
	return o.config.SourcePoolAddressByChain[sourceChain] == msgToken.SourcePoolAddress.String()
}

// Close closes the observer and releases any resources.
func (*LBTCTokenObserver) Close() error {
	return nil
}

func (u *LBTCTokenObserver) pickOnlyUSDCMessages(
	messageObservations exectypes.MessageObservations,
) map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.RampTokenAmount {
	lbtcMessages := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.RampTokenAmount)
	for chainSelector, messages := range messageObservations {
		lbtcMessages[chainSelector] = make(map[reader.MessageTokenID]cciptypes.RampTokenAmount)
		for seqNum, message := range messages {
			for i, tokenAmount := range message.TokenAmounts {
				isLBTC := u.config.SourcePoolAddressByChain[chainSelector] == tokenAmount.SourcePoolAddress.String()
				if isLBTC {
					lbtcMessages[chainSelector][reader.NewMessageTokenID(seqNum, i)] = tokenAmount
				}
				u.lggr.Debugw(
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

func (u *LBTCTokenObserver) fetchAttestations(ctx context.Context, messages map[cciptypes.ChainSelector]map[reader.MessageTokenID]cciptypes.RampTokenAmount) (map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus, error) {
	res := make(map[cciptypes.ChainSelector]map[reader.MessageTokenID]tokendata.AttestationStatus)
	for chainSelector, tokenAmounts := range messages {
		res[chainSelector] = make(map[reader.MessageTokenID]tokendata.AttestationStatus)
		batch := make(map[string]reader.MessageTokenID)
		for id, message := range tokenAmounts {
			batch[hexutil.Encode(message.ExtraData)] = id
			if len(batch) == u.config.AttestationAPIBatchSize {
				attestations, err := u.client.fetchBatch(ctx, batch)
				if err != nil {
					return nil, err
				}
				maps.Copy(res[chainSelector], attestations)
				batch = make(map[string]reader.MessageTokenID)
			}
		}
	}
	return res, nil
}

func NewLBTCTokenObserver(
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	config pluginconfig.LBTCObserverConfig,
) (*LBTCTokenObserver, error) {
	client, err := NewLBTCAttestationClient(lggr, config)
	if err != nil {
		return nil, fmt.Errorf("create attestation client: %w", err)
	}

	return &LBTCTokenObserver{
		lggr:              lggr,
		destChainSelector: destChainSelector,
		config:            config,
		client:            client,
	}, nil
}
