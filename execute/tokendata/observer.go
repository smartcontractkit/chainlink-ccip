package tokendata

import (
	"context"
	"errors"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccipocr3"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/execute/tokendata/usdc"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type TokenDataObserver interface {
	// Observe takes the message observations and returns token data observations for matching tokens.
	// If TokenDataObserver doesn't support the token it should return NoopTokenData.
	Observe(
		ctx context.Context,
		observations exectypes.MessageObservations,
	) (exectypes.TokenDataObservations, error)

	// IsTokenSupported returns true if the token is supported by the observer.
	IsTokenSupported(sourceChain cciptypes.ChainSelector, msgToken cciptypes.RampTokenAmount) bool
}

// CompositeTokenDataObserver is a TokenDataObserver that combines multiple TokenDataObserver behind the same interface.
// Goal of that is to support multiple token observers supporting different tokens (e.g. CCTP, MyFancyToken etc)
type CompositeTokenDataObserver struct {
	observers []TokenDataObserver
}

// nolint early-return
func NewCompositeTokenDataObserver(config []pluginconfig.TokenDataObserverConfig) (*CompositeTokenDataObserver, error) {
	observers := make([]TokenDataObserver, len(config))
	for i, c := range config {
		if c.USDCCCTPObserverConfig != nil {
			observers[i] = usdc.NewUSDCCCTPTokenDataObserver(*c.USDCCCTPObserverConfig)
		} else {
			return nil, errors.New("unsupported token data observer")
		}
	}
	return &CompositeTokenDataObserver{observers: observers}, nil
}

// Observe start with stubbing exectypes.TokenDataObservations with empty data based on the supported tokens.
// Then it iterates over all observers and merges token data returned from them into the final result.
func (t *CompositeTokenDataObserver) Observe(
	ctx context.Context,
	msgObservations exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	tokenDataObservations := t.initTokenDataObservations(msgObservations)

	for _, ob := range t.observers {
		tokenData, err := ob.Observe(ctx, msgObservations)
		if err != nil {
			return nil, err
		}
		tokenDataObservations = merge(tokenDataObservations, tokenData)
	}
	return tokenDataObservations, nil
}

func (t *CompositeTokenDataObserver) IsTokenSupported(
	chainSelector cciptypes.ChainSelector,
	token cciptypes.RampTokenAmount,
) bool {
	for _, ob := range t.observers {
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
func (t *CompositeTokenDataObserver) initTokenDataObservations(
	observations exectypes.MessageObservations,
) exectypes.TokenDataObservations {
	tokenObservation := make(exectypes.TokenDataObservations)
	for chainSelector, observation := range observations {
		// Initialize the nested map if not already present
		if _, exists := tokenObservation[chainSelector]; !exists {
			tokenObservation[chainSelector] = make(map[cciptypes.SeqNum]exectypes.MessageTokenData)
		}

		for seq, message := range observation {
			tokenData := make([]exectypes.TokenData, len(message.TokenAmounts))
			for i, token := range message.TokenAmounts {
				if !t.IsTokenSupported(chainSelector, token) {
					tokenData[i] = exectypes.NewNoopTokenData()
					continue
				}
				tokenData[i] = exectypes.NewEmptyTokenData()
			}
			tokenObservation[chainSelector][seq] = exectypes.MessageTokenData{TokenData: tokenData}
		}
	}
	return tokenObservation
}

func merge(base exectypes.TokenDataObservations, _ exectypes.TokenDataObservations) exectypes.TokenDataObservations {
	return base
}

type NoopTokenDataObserver struct{}

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
				tokenData[i] = exectypes.NewEmptyTokenData()
			}
			tokenObservations[selector][seq] = exectypes.MessageTokenData{TokenData: tokenData}
		}
	}
	return tokenObservations, nil
}

func (n *NoopTokenDataObserver) IsTokenSupported(_ cciptypes.ChainSelector, _ cciptypes.RampTokenAmount) bool {
	return false
}
