package siloed_usdc_token_pool

import (
	"github.com/Masterminds/semver/v3"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var ContractType cldf_deployment.ContractType = "SiloedUSDCTokenPool"

var Version = semver.MustParse("1.7.0")
