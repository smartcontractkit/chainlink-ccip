package adapters

import (
	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func init() {
	v := utils.Version_1_5_0

	curseRegistry := fastcurse.GetCurseRegistry()
	curseRegistry.RegisterNewCurse(fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      v,
		CurseAdapter:        NewCurseAdapter(),
		CurseSubjectAdapter: NewCurseAdapter(),
	})

	feeReg := fees.GetRegistry()
	feeReg.RegisterFeeAdapter(chainsel.FamilyEVM, v, NewFeesAdapter())
}
