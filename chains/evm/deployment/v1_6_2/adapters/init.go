package adapters

import (
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	adaptersV1_6_1 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_1/adapters"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

func init() {
	version := utils.Version_1_6_2

	// NOTE: v1.6.2 token pool contracts still use the base v1.6.1 token pool bindings
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilyEVM, version, adaptersV1_6_1.NewTokenAdapter())
}
