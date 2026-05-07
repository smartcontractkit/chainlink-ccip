package adapters

import (
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	feesapi "github.com/smartcontractkit/chainlink-ccip/deployment/fees"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	mcmsreaderapi "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

func init() {
	v := utils.Version_1_0_0

	deployapi.GetRegistry().RegisterDeployer(chain_selectors.FamilyEVM, v, &EVMDeployer{})
	deployapi.GetTransferOwnershipRegistry().RegisterAdapter(chain_selectors.FamilyEVM, v, &EVMTransferOwnershipAdapter{})
	mcmsreaderapi.GetRegistry().RegisterMCMSReader(chain_selectors.FamilyEVM, &EVMMCMSReader{})
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilyEVM, v, &EVMTokenBase{})
	feesapi.GetRegistry().RegisterFeeResolver(chain_selectors.FamilyEVM, &EVMFeeResolver{})
	deployapi.GetAddressNormalizerRegistry().RegisterAddressNormalizer(chain_selectors.FamilyEVM, &EVMAddressNormalizer{})
}
