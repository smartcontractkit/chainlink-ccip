package adapters

import (
	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	// Pull in EVM token-contract strategies so the per-token-type registry is
	// populated before any consumer of strategy.GetRegistry runs.
	_ "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/tokens/strategy/registrations"
	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	tokensapi "github.com/smartcontractkit/chainlink-ccip/deployment/tokens"
	mcmsreaderapi "github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

func init() {
	v, err := semver.NewVersion("1.0.0")
	if err != nil {
		panic(err)
	}
	deployapi.GetRegistry().RegisterDeployer(chain_selectors.FamilyEVM, v, &EVMDeployer{})
	deployapi.GetTransferOwnershipRegistry().RegisterAdapter(chain_selectors.FamilyEVM, v, &EVMTransferOwnershipAdapter{})
	mcmsreaderapi.GetRegistry().RegisterMCMSReader(chain_selectors.FamilyEVM, &EVMMCMSReader{})
	tokensapi.GetTokenAdapterRegistry().RegisterTokenAdapter(chain_selectors.FamilyEVM, v, &EVMTokenBase{})
}
