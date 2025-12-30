package link_token

import (
	"github.com/Masterminds/semver/v3"
	cldf_deployment "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
)

var (
	ContractType cldf_deployment.ContractType = "LINK"
	Version                                   = semver.MustParse("1.5.0")
)
