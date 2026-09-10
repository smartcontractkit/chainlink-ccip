package adapters

import (
	chainsel "github.com/smartcontractkit/chain-selectors"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/hooks" // registers EVM post-proposal CCIP send hook provider
	adapters1_2_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	evmseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	rmnadapters "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v2_1_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func init() {
	evmAdapter := evmseq.EVMAdapter{}
	v1_2_0 := utils.Version_1_2_0
	v1_6_0 := utils.Version_1_6_0
	v1_6_3 := utils.Version_1_6_3

	// Cursing a 1.6 lane suite goes through the RMN 2.1.0 adapter. That adapter resolves the RMN via
	// RMNProxy.getARM() rather than a pinned datastore ref, so it follows whichever RMN the chain
	// actually runs -- which, after the RMN 2.1.0 rollout, is 2.1.0 even on a 1.6 lane suite.
	curseRegistry := fastcurse.GetCurseRegistry()
	curseRegistry.RegisterNewCurse(fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      v1_6_0,
		CurseAdapter:        rmnadapters.NewCurseAdapter(),
		CurseSubjectAdapter: rmnadapters.NewCurseAdapter(),
	})

	laneMigratorRegistry := deploy.GetLaneMigratorRegistry()
	laneMigratorRegistry.RegisterRouterUpdater(chainsel.FamilyEVM, v1_2_0, &adapters1_2_0.RouterUpdater{})
	laneMigratorRegistry.RegisterRampUpdater(chainsel.FamilyEVM, v1_6_0, &LaneMigrator{})

	// NOTE: the fee quoter method signature for updating token transfer fee configs and
	// dest chain configs is the same between versions v1.6.0 and v1.6.3 so this adapter
	// can be reused for both versions.
	feeReg := fees.GetRegistry()
	feeReg.RegisterFeeAdapter(chainsel.FamilyEVM, v1_6_0, NewFeesAdapter(&evmAdapter))
	feeReg.RegisterFeeAdapter(chainsel.FamilyEVM, v1_6_3, NewFeesAdapter(&evmAdapter))

	feeAggReg := fees.GetFeeAggregatorRegistry()
	feeAggReg.RegisterFeeAggregatorAdapter(chainsel.FamilyEVM, v1_6_0, NewFeeAggregatorAdapter(&evmAdapter))
}
