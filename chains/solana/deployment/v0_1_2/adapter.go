package v012

import (
	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	mcmstypes "github.com/smartcontractkit/mcms/types"

	ccipapi "github.com/smartcontractkit/chainlink-ccip/deployment/v1_6"
)

func init() {
	v, err := semver.NewVersion("0.1.2")
	if err != nil {
		panic(err)
	}
	ccipapi.RegisterChainAdapter(chain_selectors.FamilySolana, v, &SolanaAdapter{})
}

type SolanaAdapter struct{}

func (a *SolanaAdapter) GetOnRampAddress(e cldf.Environment, chainSelector uint64) ([]byte, error) {
	return []byte{}, nil // Not implemented for Solana
}

func (a *SolanaAdapter) GetOffRampAddress(e cldf.Environment, chainSelector uint64) ([]byte, error) {
	return []byte{}, nil // Not implemented for Solana
}

func (a *SolanaAdapter) GetTimelockAddress(e cldf.Environment, chainSelector uint64) (string, error) {
	return "", nil // Not implemented for Solana
}

func (a *SolanaAdapter) GetMCMSMetadata(e cldf.Environment, chainSelector uint64, action mcmstypes.TimelockAction) (mcmstypes.ChainMetadata, error) {
	return mcmstypes.ChainMetadata{}, nil // Not implemented for Solana
}
