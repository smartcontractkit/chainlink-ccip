package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/token_pool"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
)

type ApplyChainUpdatesPerChainInput struct {
	ChainSelector uint64
	Address       common.Address
	Updates       token_pool_ops.ApplyChainUpdatesArgs
}

type ApplyChainUpdatesInput struct {
	ChainInputs []ApplyChainUpdatesPerChainInput
	MCMS        mcms.Input
}

func ApplyChainUpdatesChangeset() cldf.ChangeSetV2[ApplyChainUpdatesInput] {
	return cldf.CreateChangeSet(applyChainUpdatesApply(), applyChainUpdatesVerify())
}

// This function invokes the ApplyChainUpdatesSequence to apply chain updates to a sequence of TokenPool contracts on multiple chains.
// It reauires that the inputs be fed directly. In a future version this changeset will be modified to allow for a more
// flexible input that requires fewer parameters to be defined.
func applyChainUpdatesApply() func(cldf.Environment, ApplyChainUpdatesInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input ApplyChainUpdatesInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		addressesByChain := make(map[uint64][]common.Address)
		updatesByChain := make(map[uint64]map[common.Address]token_pool_ops.ApplyChainUpdatesArgs)
		for _, perChainInput := range input.ChainInputs {
			// For each chain input, add the address to the addressesByChain map.
			addressesByChain[perChainInput.ChainSelector] = append(addressesByChain[perChainInput.ChainSelector], perChainInput.Address)

			// If the map for the chain selector doesn't exist yet, initialize it to prevent a nil pointer error.
			if _, ok := updatesByChain[perChainInput.ChainSelector]; !ok {
				updatesByChain[perChainInput.ChainSelector] = make(map[common.Address]token_pool_ops.ApplyChainUpdatesArgs)
			}
			updatesByChain[perChainInput.ChainSelector][perChainInput.Address] = perChainInput.Updates
		}

		sequenceInput := sequences.ApplyChainUpdatesSequenceInput{
			AddressesByChain: addressesByChain,
			UpdatesByChain:   updatesByChain,
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.TokenPoolApplyChainUpdatesSequence, e.BlockChains, sequenceInput)
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

func applyChainUpdatesVerify() func(cldf.Environment, ApplyChainUpdatesInput) error {
	return func(e cldf.Environment, input ApplyChainUpdatesInput) error {
		for _, perChainInput := range input.ChainInputs {
			// Check that the address is not zero
			if perChainInput.Address == (common.Address{}) {
				return fmt.Errorf("address must not be zero for chain selector %d", perChainInput.ChainSelector)
			}

			// Check that the chain to operate on exists
			if exists := e.BlockChains.Exists(perChainInput.ChainSelector); !exists {
				return fmt.Errorf("chain with selector %d does not exist and thus cannot be updated", perChainInput.ChainSelector)
			}

			// Check that the chains to remove exist
			for _, remoteChainSelector := range perChainInput.Updates.RemoteChainSelectorsToRemove {
				if exists := e.BlockChains.Exists(remoteChainSelector); !exists {
					return fmt.Errorf("chain with selector %d does not exist and thus cannot be removed", remoteChainSelector)
				}
			}

			// Check that the chains to add exist
			for _, chainUpdate := range perChainInput.Updates.ChainsToAdd {
				if exists := e.BlockChains.Exists(chainUpdate.RemoteChainSelector); !exists {
					return fmt.Errorf("chain with selector %d does not exist and thus cannot be added", chainUpdate.RemoteChainSelector)
				}
			}

			for _, perChainInput := range input.ChainInputs {
				// Check that at least one chain must be added or removed for each chain selector and address
				if len(perChainInput.Updates.ChainsToAdd) == 0 && len(perChainInput.Updates.RemoteChainSelectorsToRemove) == 0 {
					return fmt.Errorf("at least one chain must be added or removed for chain selector %d and address %s", perChainInput.ChainSelector, perChainInput.Address.Hex())
				}
			}
		}
		return nil
	}
}
