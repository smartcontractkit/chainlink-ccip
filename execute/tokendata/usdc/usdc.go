package usdc

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type USDCCCTPTokenDataObserver struct {
	configs []pluginconfig.USDCCCTPObserverConfig
}

func (u *USDCCCTPTokenDataObserver) Observe(
	ctx context.Context,
	messages exectypes.MessageObservations,
) (exectypes.TokenDataObservations, error) {
	// Ask LogPoller for USDC messages and matching events
	usdcMessages, err := u.pickOnlyUSDCMessages(messages)
	if err != nil {
		return nil, err
	}

	// Fetch attestations for USDC messages
	attestations, err := u.fetchAttestations(ctx, usdcMessages)
	if err != nil {
		return nil, err
	}

	// Add attestations to the token observations
	return u.extractTokenData(attestations)
}

func (u *USDCCCTPTokenDataObserver) pickOnlyUSDCMessages(_ exectypes.MessageObservations) (interface{}, error) {
	fmt.Println(u.configs)
	panic("implement me")
}

func (u *USDCCCTPTokenDataObserver) fetchAttestations(_ context.Context, _ interface{}) (interface{}, error) {
	panic("implement me")
}

func (u *USDCCCTPTokenDataObserver) extractTokenData(_ interface{}) (exectypes.TokenDataObservations, error) {
	panic("implement me")
}
