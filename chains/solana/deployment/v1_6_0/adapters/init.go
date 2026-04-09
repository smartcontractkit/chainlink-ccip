package adapters

import (
	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"

	solseq "github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
)

func init() {
	curseRegistry := fastcurse.GetCurseRegistry()
	curseRegistry.RegisterNewCurse(fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilySolana,
		CursingVersion:      semver.MustParse("1.6.0"),
		CurseAdapter:        NewCurseAdapter(),
		CurseSubjectAdapter: NewCurseAdapter(),
	})

	feeReg := fees.GetRegistry()
	solAdapter := solseq.SolanaAdapter{}
	feeReg.RegisterFeeAdapter(chainsel.FamilySolana, semver.MustParse("1.6.0"), NewFeesAdapter(&solAdapter))

	feeAggReg := fees.GetFeeAggregatorRegistry()
	feeAggReg.RegisterFeeAggregatorAdapter(chainsel.FamilySolana, semver.MustParse("1.6.0"), NewFeeAggregatorAdapter(&solAdapter))
}
