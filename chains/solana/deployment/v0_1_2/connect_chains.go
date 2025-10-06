package v012

import (
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	ccipapi "github.com/smartcontractkit/chainlink-ccip/deployment/v1_6"
)

// high level API
func (a *SolanaAdapter) ConfigureLaneLegAsSource(e cldf.Environment, cfg ccipapi.UpdateLanesInput) (cldf.ChangesetOutput, error) {
	return cldf.ChangesetOutput{}, nil // Not implemented for Solana
}

func (a *SolanaAdapter) ConfigureLaneLegAsDest(e cldf.Environment, cfg ccipapi.UpdateLanesInput) (cldf.ChangesetOutput, error) {
	return cldf.ChangesetOutput{}, nil // Not implemented for Solana
}

func (a *SolanaAdapter) ConfigureLaneAsSourceAndDest(e cldf.Environment, cfg ccipapi.UpdateLanesInput) (cldf.ChangesetOutput, error) {
	return cldf.ChangesetOutput{}, nil // Not implemented for Solana
}
