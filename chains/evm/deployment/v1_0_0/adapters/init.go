package adapters

import (
	"github.com/Masterminds/semver/v3"
	chain_selectors "github.com/smartcontractkit/chain-selectors"

	deployapi "github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
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
}
