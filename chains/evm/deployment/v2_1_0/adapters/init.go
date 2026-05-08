package adapters

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"

	rmnops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/operations/rmn"
	"github.com/smartcontractkit/chainlink-ccip/deployment/authorizedcallers"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
)

func init() {
	curseRegistry := fastcurse.GetCurseRegistry()
	curseRegistry.RegisterNewCurse(fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("2.1.0"),
		CurseAdapter:        NewCurseAdapter(),
		CurseSubjectAdapter: NewCurseAdapter(),
	})

	authCallersRegistry := authorizedcallers.GetAuthorizedCallersRegistry()
	authCallersRegistry.RegisterAdapter(
		chainsel.FamilyEVM,
		rmnops.ContractType,
		rmnops.Version,
		NewEVMAuthorizedCallersAdapter(
			rmnops.ApplyAuthorizedCallerUpdates,
			rmnops.GetAllAuthorizedCallers,
			func(added, removed []common.Address) rmnops.AuthorizedCallerArgs {
				return rmnops.AuthorizedCallerArgs{AddedCallers: added, RemovedCallers: removed}
			},
		),
	)
	// Additional AuthorizedCallers-inheriting contracts (FeeQuoter, AdvancedPoolHooks, …)
	// follow the same pattern once their ops packages are available:
	//
	//   authCallersRegistry.RegisterAdapter(
	//       chainsel.FamilyEVM,
	//       feeQuoterOps.ContractType,
	//       feeQuoterOps.Version,
	//       NewEVMAuthorizedCallersAdapter(
	//           feeQuoterOps.ApplyAuthorizedCallerUpdates,
	//           feeQuoterOps.GetAllAuthorizedCallers,
	//           func(added, removed []common.Address) feeQuoterOps.AuthorizedCallerArgs {
	//               return feeQuoterOps.AuthorizedCallerArgs{AddedCallers: added, RemovedCallers: removed}
	//           },
	//       ),
	//   )
}
