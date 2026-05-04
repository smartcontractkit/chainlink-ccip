package adapters

import (
	chainsel "github.com/smartcontractkit/chain-selectors"

	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
)

func init() {
	fees.GetRegistry().RegisterFeeResolver(chainsel.FamilySolana, &SolanaFeeResolver{})
}
