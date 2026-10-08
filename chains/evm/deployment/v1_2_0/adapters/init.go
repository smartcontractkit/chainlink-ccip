package adapters

import (
	chainsel "github.com/smartcontractkit/chain-selectors"

	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func init() {
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chainsel.FamilyEVM, utils.Version_1_2_0, NewTokenAdapter())
}
