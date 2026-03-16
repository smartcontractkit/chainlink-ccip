package adapters

import (
	"github.com/Masterminds/semver/v3"

	chainsel "github.com/smartcontractkit/chain-selectors"
	adapters1_2 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/adapters"
	adapters1_5 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/adapters"
	adapters1_6 "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/adapters"
	EVMAdapter "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_0/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/deploy"
)

func init() {
	v, err := semver.NewVersion("2.0.0")
	if err != nil {
		panic(err)
	}
	deploy.GetRegistry().RegisterDeployer(chainsel.FamilyEVM, v, &EVMAdapter.EVMAdapter{})

	fqReg := deploy.GetFQAndRampUpdaterRegistry()
	fqReg.RegisterFeeQuoterUpdater(chainsel.FamilyEVM, semver.MustParse("2.0.0"), FeeQuoterUpdater[any]{})
	fqReg.RegisterRampUpdater(chainsel.FamilyEVM, semver.MustParse("1.6.0"), adapters1_6.RampUpdateWithFQ{})
	fqReg.RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.6.0"), &adapters1_6.ConfigImportAdapter{})
	fqReg.RegisterConfigImporter(chainsel.FamilyEVM, semver.MustParse("1.5.0"), &adapters1_5.ConfigImportAdapter{})
	fqReg.RegisterConfigImporterVersionResolver(chainsel.FamilyEVM, &adapters1_2.LaneVersionResolver{})
}
