package adapters

import (
	chainsel "github.com/smartcontractkit/chain-selectors"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func init() {
	v := utils.Version_1_4_0

	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chainsel.FamilyEVM, v, NewTokenAdapter())
}
