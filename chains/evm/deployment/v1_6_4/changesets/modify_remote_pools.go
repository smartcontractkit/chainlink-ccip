package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	"github.com/smartcontractkit/chainlink-deployments-framework/operations"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type ModifyRemotePoolsInput struct {
	ChainInputs []ModifyRemotePoolsPerChainInput
	MCMS        mcms.Input
}

type ModifyRemotePoolsPerChainInput struct {
	ChainSelector uint64
	Address       common.Address
	Modification  token_pool.RemotePoolModification
}

func ModifyRemotePoolsChangeset() cldf.ChangeSetV2[ModifyRemotePoolsInput] {
	return cldf.CreateChangeSet(modifyRemotePoolsApply(), modifyRemotePoolsVerify())
}

func modifyRemotePoolsApply() func(cldf.Environment, ModifyRemotePoolsInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input ModifyRemotePoolsInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		addressesByChain := make(map[uint64][]common.Address)
		modificationsByChain := make(map[uint64]map[common.Address]token_pool.RemotePoolModification)

		for _, perChainInput := range input.ChainInputs {
			// For each chain input, add the addresses to the addressesByChain map.
			addressesByChain[perChainInput.ChainSelector] = append(addressesByChain[perChainInput.ChainSelector], perChainInput.Address)

			// Initialize the map for the chain selector if it doesn't exist yet to prevent a nil pointer dereference.
			if _, ok := modificationsByChain[perChainInput.ChainSelector]; !ok {
				modificationsByChain[perChainInput.ChainSelector] = make(map[common.Address]token_pool.RemotePoolModification)
			}

			// Map the provided RemotePoolModification to the given address and chain selector.
			// Note: This assumes that a given changeset input will not have with multiple modifications for a specific
			// address on a specific chain selector. If this occurs, then the last one in the input slice will be used.
			modificationsByChain[perChainInput.ChainSelector][perChainInput.Address] = perChainInput.Modification
		}

		// Execute the sequence with the combined input
		sequenceInput := sequences.ModifyRemotePoolsSequenceInput{
			AddressesByChain:     addressesByChain,
			ModificationsByChain: modificationsByChain,
		}
		report, err := operations.ExecuteSequence(e.OperationsBundle, sequences.ModifyRemotePoolsSequence, e.BlockChains, sequenceInput)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(input.MCMS)
	}
}

func modifyRemotePoolsVerify() func(cldf.Environment, ModifyRemotePoolsInput) error {
	return func(e cldf.Environment, input ModifyRemotePoolsInput) error {
		for _, perChainInput := range input.ChainInputs {
			if exists := e.BlockChains.Exists(perChainInput.ChainSelector); !exists {
				return fmt.Errorf("chain with selector %d does not exist", perChainInput.ChainSelector)
			}

			if perChainInput.Address == (common.Address{}) {
				return fmt.Errorf("address must not be zero for chain selector %d", perChainInput.ChainSelector)
			}

			if perChainInput.Modification.Operation == "" {
				return fmt.Errorf("operation cannot be empty for chain selector %d and address %s", perChainInput.ChainSelector, perChainInput.Address.Hex())
			}
		}
		return nil
	}
}
