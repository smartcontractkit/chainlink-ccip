package adapters

import (
	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
)

func init() {
	v := semver.MustParse("1.2.0")

	laneMigratorReg := deploy.GetLaneMigratorRegistry()
	laneMigratorReg.RegisterRouterUpdater(chainsel.FamilyEVM, v, &RouterUpdater{})
}
