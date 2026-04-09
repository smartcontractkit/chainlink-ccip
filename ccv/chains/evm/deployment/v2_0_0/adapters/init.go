package adapters

import (
	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
)

func init() {
	feeAggReg := fees.GetFeeAggregatorRegistry()
	feeAggReg.RegisterFeeAggregatorAdapter(chainsel.FamilyEVM, semver.MustParse("2.0.0"), NewFeeAggregatorAdapter())
}
