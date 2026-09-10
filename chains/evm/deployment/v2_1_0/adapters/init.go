package adapters

import (
	chainsel "github.com/smartcontractkit/chain-selectors"

	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
)

// Registration lives alongside the adapter it registers. It is keyed by RMN's own version rather
// than the surrounding EVM lane suite version: RMN is upgraded independently of the rest of the
// lane contracts, behind the RMNProxy.
func init() {
	curseRegistry := fastcurse.GetCurseRegistry()
	curseRegistry.RegisterNewCurse(fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      rmnops.Version,
		CurseAdapter:        NewCurseAdapter(),
		CurseSubjectAdapter: NewCurseAdapter(),
	})
}
