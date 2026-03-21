package adapters

import (
	"github.com/Masterminds/semver/v3"

	chainsel "github.com/smartcontractkit/chain-selectors"

	adapters2_0 "github.com/smartcontractkit/chainlink-ccip/ccv/chains/evm/deployment/v2_0_0/adapters"
	adapters1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	adapters1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/adapters"
	adapters1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"

	"github.com/smartcontractkit/chainlink-ccip/deployment/lanes"
	ccvadapters "github.com/smartcontractkit/chainlink-ccip/deployment/v1_7_0/adapters"
)

func init() {
	v := semver.MustParse("2.0.0")
	fqReg := deploy.GetFQAndRampUpdaterRegistry()
	fqReg.RegisterFeeQuoterUpdater(chainsel.FamilyEVM, semver.MustParse("2.0.0"), adapters2_0.FeeQuoterUpdater[any]{})
	fqReg.RegisterRampUpdater(chainsel.FamilyEVM, semver.MustParse("1.6.0"), adapters1_6.RampUpdateWithFQ{})
	fqReg.RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.6.0"), &adapters1_6.ConfigImportAdapter{})
	fqReg.RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.5.0"), &adapters1_5.ConfigImportAdapter{})
	fqReg.RegisterConfigImporterVersionResolver(chainsel.FamilyEVM, &adapters1_2.LaneVersionResolver{})

	laneMigratorReg := deploy.GetLaneMigratorRegistry()
	laneMigratorReg.RegisterRampUpdater(chainsel.FamilyEVM, semver.MustParse("2.0.0"), &LaneMigrator{})
	laneMigratorReg.RegisterRouterUpdater(chainsel.FamilyEVM, semver.MustParse("1.2.0"), &adapters1_2.RouterUpdater{})

	lanes.GetLaneAdapterRegistry().RegisterLaneAdapter(chainsel.FamilyEVM, v, &ChainFamilyAdapter{})
	ccvadapters.GetCommitteeVerifierContractRegistry().Register(chainsel.FamilyEVM, &EVMCommitteeVerifierContractAdapter{})
	ccvadapters.GetExecutorConfigRegistry().Register(chainsel.FamilyEVM, &EVMExecutorConfigAdapter{})
	ccvadapters.GetVerifierJobConfigRegistry().Register(chainsel.FamilyEVM, &EVMVerifierJobConfigAdapter{})
	ccvadapters.GetDeployChainContractsRegistry().Register(chainsel.FamilyEVM, &EVMDeployChainContractsAdapter{})
	ccvadapters.GetDeployChainContractsRegistry().RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.6.0"), &adapters1_6.ConfigImportAdapter{})
	ccvadapters.GetDeployChainContractsRegistry().RegisterLaneVersionResolver(chainsel.FamilyEVM, &adapters1_2.LaneVersionResolver{})
	ccvadapters.GetIndexerConfigRegistry().Register(chainsel.FamilyEVM, &EVMIndexerConfigAdapter{})
	ccvadapters.GetAggregatorConfigRegistry().Register(chainsel.FamilyEVM, &EVMAggregatorConfigAdapter{})
	ccvadapters.GetTokenVerifierConfigRegistry().Register(chainsel.FamilyEVM, &EVMTokenVerifierConfigAdapter{})
}
