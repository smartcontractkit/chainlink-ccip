package tokendata

import (
	"context"
	"errors"
	"fmt"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/usdc"
	"github.com/smartcontractkit/chainlink-ccip/pkg/contractreader"
	"github.com/smartcontractkit/chainlink-ccip/pkg/reader"
	cciptypes "github.com/smartcontractkit/chainlink-ccip/pkg/types/ccipocr3"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type TokenDataObserver interface {
	// Observe takes the message observations and returns token data observations for matching tokens.
	// If TokenDataObserver doesn't support the token it must return exectypes.NotSupportedTokenData.
	// This way higher layer will know how to orchestrate Observers and which observations are relevant.
	//
	// TokenDataObserver must return data for every token that is present in the exectypes.MessageObservations.
	// Meaning that exectypes.TokenDataObservations must reflect the structure of exectypes.MessageObservations
	// (the same number of objects and tokens).
	// * if token is supported and requires offchain processing, TokenDataObserver must initialize the token according
	// 	 to its logic. Depending on the result it must initialize with either exectypes.NewSuccessTokenData or
	//	 exectypes.NewErrorTokenData
	// * if token is not supported, TokenDataObserver must set the token data to exectypes.NotSupportedTokenData
	Observe(
		ctx context.Context,
		observations exectypes.MessageObservations,
	) (exectypes.TokenDataObservations, error)

	// IsTokenSupported returns true if the token is supported by the observer.
	IsTokenSupported(sourceChain cciptypes.ChainSelector, msgToken cciptypes.RampTokenAmount) bool
}

// compositeTokenDataObserver is a TokenDataObserver that combines multiple TokenDataObserver behind the same interface.
// Goal of that is to support multiple token observers supporting different tokens (e.g. CCTP, MyFancyToken etc)
type compositeTokenDataObserver struct {
	lggr      logger.Logger
	observers []TokenDataObserver
}

// NewConfigBasedCompositeObservers creates a compositeTokenDataObserver based on the provided configuration.
// Slice of []pluginconfig.TokenDataObserverConfig must be deduped and validated by the plugin.
// Therefore, we don't re-run any validation and only match configs to the proper TokenDataObserver implementation.
// This constructor that should be used by the plugin.
func NewConfigBasedCompositeObservers(
	ctx context.Context,
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	config []pluginconfig.TokenDataObserverConfig,
	encoder cciptypes.TokenDataEncoder,
	readers map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
) (TokenDataObserver, error) {
	observers := make([]TokenDataObserver, len(config))
	for i, c := range config {
		// TODO consider if we can get rid of this switch stmt by moving the logic to the config
		// e.g. observers[i] := config.CreateTokenDataObserver()
		switch {
		case c.USDCCCTPObserverConfig != nil:
			observer, err := createUSDCTokenObserver(ctx, lggr, destChainSelector, *c.USDCCCTPObserverConfig, encoder, readers)
			if err != nil {
				return nil, fmt.Errorf("create USDC/CCTP token observer: %w", err)
			}
			observers[i] = observer
		default:
			return nil, errors.New("unsupported token data observer")
		}
	}
	return NewCompositeObservers(lggr, observers...), nil
}

func createUSDCTokenObserver(
	ctx context.Context,
	lggr logger.Logger,
	destChainSelector cciptypes.ChainSelector,
	cctpConfig pluginconfig.USDCCCTPObserverConfig,
	encoder cciptypes.TokenDataEncoder,
	readers map[cciptypes.ChainSelector]contractreader.ContractReaderFacade,
) (TokenDataObserver, error) {
	usdcReader, err := reader.NewUSDCMessageReader(
		ctx,
		lggr,
		cctpConfig.Tokens,
		readers,
	)
	if err != nil {
		return nil, err
	}

	client, err := usdc.NewSequentialAttestationClient(lggr, cctpConfig)
	if err != nil {
		return nil, fmt.Errorf("create attestation client: %w", err)
	}

	return usdc.NewTokenDataObserver(
		lggr,
		destChainSelector,
		cctpConfig.Tokens,
		encoder.EncodeUSDC,
		usdcReader,
		client,
	), nil
}

// NewCompositeObservers creates a compositeTokenDataObserver based on the provided observers.
// Created mostly for tests purposes, it allows the user to specify custom observers and skip the part
// in which we match the configuration to the proper TokenDataObserver.
func NewCompositeObservers(lggr logger.Logger, observers ...TokenDataObserver) TokenDataObserver {
	return &compositeTokenDataObserver{lggr: lggr, observers: observers}
}

// Observe start with stubbing exectypes.TokenDataObservations with empty data based on the supported tokens.
// Then it iterates over all observers and merges token data returned from them into the final result.
func (c *compositeTokenDataObserver) Observe(
	ctx context.Context,
	msgObservations exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	tokenDataObservations := c.initTokenDataObservations(msgObservations)

	for _, ob := range c.observers {
		tokenData, err := ob.Observe(ctx, msgObservations)
		if err != nil {
			c.lggr.Error("Error while observing token data", "error", err)
			continue
		}
		tokenDataObservations, err = merge(tokenDataObservations, tokenData)
		if err != nil {
			return nil, err
		}
	}
	return tokenDataObservations, nil
}

func (c *compositeTokenDataObserver) IsTokenSupported(
	chainSelector cciptypes.ChainSelector,
	token cciptypes.RampTokenAmount,
) bool {
	for _, ob := range c.observers {
		if ob.IsTokenSupported(chainSelector, token) {
			return true
		}
	}
	return false
}

// initTokenDataObservations initializes the token data observations with empty token data, it asks the child observers
// whether the token is supported or not. If token is supported and requires additional processing we set its state to
// isReady=false. If token is noop (doesn't require offchain processing) we initialize it with empty
// data and set IsReady=true.
func (c *compositeTokenDataObserver) initTokenDataObservations(
	observations exectypes.MessageObservations,
) exectypes.TokenDataObservations {
	tokenObservation := make(exectypes.TokenDataObservations)
	for chainSelector, observation := range observations {
		tokenObservation[chainSelector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

		for seq, message := range observation {
			tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
			for i, token := range message.TokenAmounts {
				if !c.IsTokenSupported(chainSelector, token) {
					// It means that none of the registers support that,
					// we assume token doesn't require additional processing and skip it
					tokenData[i] = exectypes.NewNoopTokenData()
					continue
				}
				// Token is supported by one of the registered observers, mark it as not ready
				// It would be processed by the registered processor in the next step
				tokenData[i] = exectypes.TokenData{Ready: false, Supported: true}
			}
			tokenObservation[chainSelector][seq] = exectypes.MessageTokenData{TokenData: tokenData}
		}
	}
	return tokenObservation
}

func merge(
	base exectypes.TokenDataObservations,
	from exectypes.TokenDataObservations,
) (exectypes.TokenDataObservations, error) {
	for chainSelector, chainObservations := range from {
		for seq, messageTokenData := range chainObservations {
			if len(messageTokenData.TokenData) != len(base[chainSelector][seq].TokenData) {
				return nil, errors.New("token data length mismatch")
			}

			// Merge only TokenData created by the observer
			for i, newTokenData := range messageTokenData.TokenData {
				if base[chainSelector][seq].TokenData[i].IsReady() {
					// Already processed by another observer, skip or raise a warning
					continue
				}
				if newTokenData.Supported {
					base[chainSelector][seq].TokenData[i] = newTokenData
				}
			}
		}
	}
	return base, nil
}

type NoopTokenDataObserver struct {
	tokenSupported bool
}

func (n *NoopTokenDataObserver) Observe(
	_ context.Context,
	observations exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	tokenObservations := make(exectypes.TokenDataObservations)

	for selector, obs := range observations {
		tokenObservations[selector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)

		for seq, message := range obs {
			tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
			for i := range message.TokenAmounts {
				tokenData[i] = exectypes.NewNoopTokenData()
			}
			tokenObservations[selector][seq] = exectypes.MessageTokenData{TokenData: tokenData}
		}
	}
	return tokenObservations, nil
}

func (n *NoopTokenDataObserver) IsTokenSupported(_ cciptypes.ChainSelector, _ cciptypes.RampTokenAmount) bool {
	return n.tokenSupported
}
