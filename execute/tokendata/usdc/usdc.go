package usdc

import (
	"context"
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/execute/exectypes"
	"github.com/smartcontractkit/chainlink-ccip/pluginconfig"
)

type UsdcTokenDataProcessor struct {
	configs []pluginconfig.UsdcCctpTokenDataProcessor
}

func (u *UsdcTokenDataProcessor) ProcessTokenData(ctx context.Context, messages exectypes.MessageObservations) (exectypes.TokenDataObservations, error) {
	usdcMessages, err := u.pickOnlyUSDCMessages(messages)
	if err != nil {
		return nil, err
	}

	attestations, err := u.fetchAttestations(ctx, usdcMessages)
	if err != nil {
		return nil, err
	}

	return u.extractTokenData(attestations)
}

func (u *UsdcTokenDataProcessor) pickOnlyUSDCMessages(_ exectypes.MessageObservations) (interface{}, error) {
	fmt.Println(u.configs)
	panic("implement me")
}

func (u *UsdcTokenDataProcessor) fetchAttestations(_ context.Context, _ interface{}) (interface{}, error) {
	panic("implement me")
}

func (u *UsdcTokenDataProcessor) extractTokenData(_ interface{}) (exectypes.TokenDataObservations, error) {
	panic("implement me")
}
