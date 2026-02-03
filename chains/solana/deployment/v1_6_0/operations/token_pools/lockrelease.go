package token_pools

import (
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var LockReleaseProgramName = "lockrelease_token_pool"

var DeployLockRelease = operations.NewOperation(
	"lockrelease:deploy",
	common_utils.Version_1_6_0,
	"Deploys the LockReleaseTokenPool program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []datastore.AddressRef) (datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			common_utils.LockReleaseTokenPool,
			common_utils.Version_1_6_0,
			"",
			LockReleaseProgramName)
	},
)
