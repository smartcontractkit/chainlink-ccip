package sequences

import (
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	mcms_types "github.com/smartcontractkit/mcms/types"
)

// OnChainOutput is a standard output type for sequences that deploy contracts on-chain and perform write operations.
type OnChainOutput struct {
	// Addresses are the contract addresses that the sequence deployed.
	Addresses []datastore.AddressRef
	// BatchOps are operations that must be executed via MCMS.
	// Order is important and should be preserved during construction of the proposal.
	// Each batch operation is executed atomically.
	BatchOps []mcms_types.BatchOperation
}
