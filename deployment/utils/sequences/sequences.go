package sequences

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	cldf_chain "github.com/smartcontractkit/chainlink-deployments-framework/chain"
)

// WithChainSelector is implemented by configs that specify a target chain selector.
type WithChainSelector interface {
	ChainSelector() uint64
}

// OnChainOutput is a standard output type for sequences that deploy contracts on-chain and perform write operations.
type OnChainOutput struct {
	// Addresses are the contract addresses that the sequence deployed.
	Addresses []datastore.AddressRef
	// BatchOps are operations that must be executed via MCMS.
	// Order is important and should be preserved during construction of the proposal.
	// Each batch operation is executed atomically.
	BatchOps []mcms_types.BatchOperation
}

func RunAndMergeSequence[IN any](
	b operations.Bundle,
	chains cldf_chain.BlockChains,
	seq *operations.Sequence[IN, OnChainOutput, cldf_chain.BlockChains],
	input IN,
	agg OnChainOutput,
) (OnChainOutput, error) {
	report, err := operations.ExecuteSequence(b, seq, chains, input)
	if err != nil {
		return agg, fmt.Errorf("failed to execute %s: %w", seq.ID(), err)
	}
	agg.BatchOps = append(agg.BatchOps, report.Output.BatchOps...)
	agg.Addresses = append(agg.Addresses, report.Output.Addresses...)
	return agg, nil
}
