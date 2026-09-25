package adapters

import (
	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func init() {
	v := utils.Version_1_5_0

	feeReg := fees.GetRegistry()
	feeReg.RegisterFeeAdapter(chainsel.FamilyEVM, v, NewFeesAdapter())

	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chainsel.FamilyEVM, v, NewTokenAdapter())
}
