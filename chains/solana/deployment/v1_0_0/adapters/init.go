package adapters

import (
	chainsel "github.com/smartcontractkit/chain-selectors"

	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
)

func init() {
	fees.GetRegistry().RegisterFeeResolver(chainsel.FamilySolana, &SolanaFeeResolver{})
	deployapi.GetAddressNormalizerRegistry().RegisterAddressNormalizer(chainsel.FamilySolana, &SolanaAddressNormalizer{})
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdminRegistryReader(chainsel.FamilySolana, &SolanaAdminRegistryReader{})
}
