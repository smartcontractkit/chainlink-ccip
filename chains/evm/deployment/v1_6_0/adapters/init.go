package adapters

import (
	chainsel "github.com/smartcontractkit/chain-selectors"

	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/hooks" // registers EVM post-proposal CCIP send hook provider
	adapters1_2_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	evmseq "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
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

	curseRegistry := fastcurse.GetCurseRegistry()
	curseRegistry.RegisterNewCurse(fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      v1_6_0,
		CurseAdapter:        NewCurseAdapter(),
		CurseSubjectAdapter: NewCurseAdapter(),
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
