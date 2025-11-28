package changesets

import (
	"fmt"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/token_admin_registry"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

type ERC20LockboxDeployInput struct {
	ChainInputs []ERC20LockboxDeployInputPerChain
	MCMS        mcms.Input
}

type ERC20LockboxDeployInputPerChain struct {
	ChainSelector uint64
}

func ERC20LockboxDeployChangeset(mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[ERC20LockboxDeployInput] {
	return cldf.CreateChangeSet(erc20LockboxDeployApply(mcmsRegistry), erc20LockboxDeployVerify(mcmsRegistry))
}

func erc20LockboxDeployApply(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ERC20LockboxDeployInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input ERC20LockboxDeployInput) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()

		for _, perChainInput := range input.ChainInputs {
			tokenAdminRegistryAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType(token_admin_registry.ContractType),
				Version:       token_admin_registry.Version,
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			sequenceInput := sequences.ERC20LockboxDeploySequenceInput{
				ChainSelector:      perChainInput.ChainSelector,
				TokenAdminRegistry: tokenAdminRegistryAddress,
			}
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.ERC20LockboxDeploySequence, e.BlockChains, sequenceInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy ERC20Lockbox on chain %d: %w", perChainInput.ChainSelector, err)
			}
			reports = append(reports, report.ExecutionReports...)
			for _, r := range report.Output.Addresses {
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
				}
			}
		}
		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithDataStore(ds).
			Build(input.MCMS)
	}
}

func erc20LockboxDeployVerify(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ERC20LockboxDeployInput) error {
	return func(e cldf.Environment, input ERC20LockboxDeployInput) error {
		return nil
	}
}
