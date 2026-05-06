package adapters

import (
	chainsel "github.com/smartcontractkit/chain-selectors"

	solseq "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func init() {
	curseRegistry := fastcurse.GetCurseRegistry()
	curseRegistry.RegisterNewCurse(fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilySolana,
		CursingVersion:      utils.Version_1_6_0,
		CurseAdapter:        NewCurseAdapter(),
		CurseSubjectAdapter: NewCurseAdapter(),
	})

	feeReg := fees.GetRegistry()
	solAdapter := solseq.SolanaAdapter{}
	feeReg.RegisterFeeAdapter(chainsel.FamilySolana, utils.Version_1_6_0, NewFeesAdapter(&solAdapter))

	feeAggReg := fees.GetFeeAggregatorRegistry()
	feeAggReg.RegisterFeeAggregatorAdapter(chainsel.FamilySolana, utils.Version_1_6_0, NewFeeAggregatorAdapter(&solAdapter))
}
