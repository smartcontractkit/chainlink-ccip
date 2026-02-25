package adapters

import (
	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"

	adapters1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	adapters1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/adapters"
	adapters1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"

	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
)

func init() {
	fqReg := deploy.GetFQAndRampUpdaterRegistry()
	fqReg.RegisterFeeQuoterUpdater(chainsel.FamilyEVM, semver.MustParse("1.7.0"), FeeQuoterUpdater[any]{})
	fqReg.RegisterRampUpdater(chainsel.FamilyEVM, semver.MustParse("1.6.0"), adapters1_6.RampUpdateWithFQ{})
	fqReg.RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.6.0"), &adapters1_6.ConfigImportAdapter{})
	fqReg.RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.5.0"), &adapters1_5.ConfigImportAdapter{})
	fqReg.RegisterConfigImporterVersionResolver(chainsel.FamilyEVM, &adapters1_2.LaneVersionResolver{})

	laneMigratorReg := deploy.GetLaneMigratorRegistry()
	laneMigratorReg.RegisterRampUpdater(chainsel.FamilyEVM, semver.MustParse("1.7.0"), &LaneMigrator{})
	laneMigratorReg.RegisterRouterUpdater(chainsel.FamilyEVM, semver.MustParse("1.2.0"), &adapters1_2.RouterUpdater{})
}
