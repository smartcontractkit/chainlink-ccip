package adapters

import (
	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/hooks" // registers EVM post-proposal CCIP send hook provider
	v1_0_0_adapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_0_0/adapters"
	adapters1_2_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	onrampops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/operations/onramp"
	evmseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
)

func init() {
	curseRegistry := fastcurse.GetCurseRegistry()
	curseRegistry.RegisterNewCurse(fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("1.6.0"),
		CurseAdapter:        NewCurseAdapter(),
		CurseSubjectAdapter: NewCurseAdapter(),
	})
	laneMigratorRegistry := deploy.GetLaneMigratorRegistry()
	laneMigratorRegistry.RegisterRampUpdater(chainsel.FamilyEVM, semver.MustParse("1.6.0"), &LaneMigrator{})
	laneMigratorRegistry.RegisterRouterUpdater(chainsel.FamilyEVM, semver.MustParse("1.2.0"), &adapters1_2_0.RouterUpdater{})

	feeReg := fees.GetRegistry()
	evmAdapter := evmseq.EVMAdapter{}
	feeReg.RegisterFeeAdapter(chainsel.FamilyEVM, semver.MustParse("1.6.0"), NewFeesAdapter(&evmAdapter))

	feeAggReg := fees.GetFeeAggregatorRegistry()
	feeAggReg.RegisterFeeAggregatorAdapter(chainsel.FamilyEVM, semver.MustParse("1.6.0"), NewFeeAggregatorAdapter(&evmAdapter))

	v1_0_0_adapters.GetEVMFeeContractResolver().RegisterOnRampOps(
		datastore.ContractType(onrampops.ContractType),
		onrampops.Version,
		onRampFeeOpsV160{},
	)
}
