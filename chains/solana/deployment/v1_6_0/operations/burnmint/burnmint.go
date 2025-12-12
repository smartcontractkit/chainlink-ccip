package burnmint

import (
	"github.com/smartcontractkit/chainlink-ccip/chains/solana/deployment/utils"
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

var ProgramName = "burnmint_token_pool"

var Deploy = operations.NewOperation(
	"burnmint:deploy",
	common_utils.Version_1_6_0,
	"Deploys the BurnMintTokenPool program",
	func(b operations.Bundle, chain cldf_solana.Chain, input []datastore.AddressRef) (datastore.AddressRef, error) {
		return utils.MaybeDeployContract(
			b,
			chain,
			input,
			common_utils.BurnMintTokenPool,
			common_utils.Version_1_6_0,
			"",
			ProgramName)
	},
)
