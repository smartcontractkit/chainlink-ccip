package adapters

import (
	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	v1_0_0_adapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/onramp"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
)

func init() {
	curseRegistry := fastcurse.GetCurseRegistry()
	curseRegistry.RegisterNewCurse(fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("1.5.0"),
		CurseAdapter:        NewCurseAdapter(),
		CurseSubjectAdapter: NewCurseAdapter(),
	})

	defaultConcurrency := 10
	feeReg := fees.GetRegistry()
	feeReg.RegisterFeeAdapter(chainsel.FamilyEVM, semver.MustParse("1.5.0"), NewFeesAdapter(&defaultConcurrency))

	v1_0_0_adapters.GetEVMFeeContractResolver().RegisterOnRampOps(
		datastore.ContractType(onrampops.ContractType),
		onrampops.Version,
		onRampFeeOpsV150{},
	)
}
