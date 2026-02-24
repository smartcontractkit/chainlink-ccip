package adapters

import (
	"github.com/Masterminds/semver/v3"
	chainsel "github.com/smartcontractkit/chain-selectors"

	adapters1_2_0 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
	"github.com/smartcontractkit/chainlink-ccip/deployment/fastcurse"
)

func init() {
	curseRegistry := fastcurse.GetCurseRegistry()
	curseRegistry.RegisterNewCurse(fastcurse.CurseRegistryInput{
		CursingFamily:       chainsel.FamilyEVM,
		CursingVersion:      semver.MustParse("1.6.0"),
		CurseAdapter:        NewCurseAdapter(),
		CurseSubjectAdapter: NewCurseAdapter(),
	})
	laneMigratorRegistry := deploy.GetLaneMigratorRegistry()
	laneMigratorRegistry.RegisterRampUpdater(chainsel.FamilyEVM, semver.MustParse("1.6.0"), &LaneMigrater{})
	laneMigratorRegistry.RegisterRouterUpdater(chainsel.FamilyEVM, semver.MustParse("1.2.0"), &adapters1_2_0.RouterUpdater{})

	rampConfigReg := deploy.GetRampConfigUpdaterRegistry()
	rampConfigReg.RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.6.0"), &ConfigImportAdapter{})
	rampConfigReg.RegisterRampConfigApplier(chainsel.FamilyEVM, semver.MustParse("1.6.0"), RampConfigApplier[any]{})
	rampConfigReg.RegisterImporterVersionResolver(chainsel.FamilyEVM, &adapters1_2_0.LaneVersionResolver{})
}
