package adapters

import (
	chainsel "github.com/smartcontractkit/chain-selectors"

	burnmintmttpops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/burn_mint_multi_token_pool"
	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
)

func init() {
	// MTTP registration is keyed by the pool's own version, independently of the lane suite.
	tokens.GetMultiTokenPoolAdapterRegistry().RegisterMultiTokenPoolAdapter(
		chainsel.FamilyEVM,
		burnmintmttpops.Version,
		&MultiTokenPoolAdapter{},
	)

	// RMN registration is keyed by RMN's own version rather than the surrounding EVM lane
	// suite version: RMN is upgraded independently behind the RMNProxy.
	curseRegistry := fastcurse.GetCurseRegistry()
	curseRegistry.RegisterNewCurse(fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      rmnops.Version,
		CurseAdapter:        NewCurseAdapter(),
		CurseSubjectAdapter: NewCurseAdapter(),
	})
}
