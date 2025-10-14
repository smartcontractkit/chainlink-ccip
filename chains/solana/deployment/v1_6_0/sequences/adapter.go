package operations

import (
	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	ccipapi "github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
)

func init() {
	v, err := semver.NewVersion("1.6.0")
	if err != nil {
		panic(err)
	}
	ccipapi.GetLaneAdapterRegistry().RegisterLaneAdapter(chain_selectors.FamilySolana, v, &SolanaAdapter{})
}

type SolanaAdapter struct{}

func (a *SolanaAdapter) GetOnRampAddress(e *cldf.Environment, chainSelector uint64) ([]byte, error) {
	return []byte{}, nil // Not implemented for Solana
}

func (a *SolanaAdapter) GetOffRampAddress(e *cldf.Environment, chainSelector uint64) ([]byte, error) {
	return []byte{}, nil // Not implemented for Solana
}

func (a *SolanaAdapter) GetFQAddress(e *cldf.Environment, chainSelector uint64) ([]byte, error) {
	return []byte{}, nil // Not implemented for Solana
}

func (a *SolanaAdapter) GetRouterAddress(e *cldf.Environment, chainSelector uint64) ([]byte, error) {
	return []byte{}, nil // Not implemented for Solana
}
