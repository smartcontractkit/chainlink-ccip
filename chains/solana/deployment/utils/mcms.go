package utils

import (
	common_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"
	cldf_datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"

	"github.com/smartcontractkit/chainlink-ccip/chains/solana/governance"
)

// GetAllMCMS returns the 1.6.0 MCMS refs of the chain in a fixed order; see governance.GetAllMCMS.
func GetAllMCMS(
	chain cldf_solana.Chain,
	qualifier string,
	existingAddresses []cldf_datastore.AddressRef) []cldf_datastore.AddressRef {
	return governance.GetAllMCMS(common_utils.Version_1_6_0, chain.ChainSelector(), qualifier, existingAddresses)
}
