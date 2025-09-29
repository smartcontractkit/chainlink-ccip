package sequences

import (
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/operations/contract"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

// OnChainOutput is a standard output type for sequences that deploy contracts on-chain and perform write operations.
type OnChainOutput struct {
	// Addresses are the contract addresses that the sequence deployed.
	Addresses []datastore.AddressRef
	// Writes are the write operations that the sequence performed.
	Writes []contract.WriteOutput
}
