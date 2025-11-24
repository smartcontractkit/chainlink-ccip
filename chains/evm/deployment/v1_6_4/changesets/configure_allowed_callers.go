package changesets

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

type ConfigureAllowedCallersInput struct {
	ChainInputs []ConfigureAllowedCallersPerChainInput
	MCMS        mcms.Input
}

type ConfigureAllowedCallersPerChainInput struct {
	ChainSelector  uint64
	AllowedCallers []erc20_lock_box.AllowedCallerConfigArgs
}

// This changeset is used to configure the allowed callers for a given token in the ERC20Lockbox contract.
// It is different from the apply_authorized_caller_updates changeset which is used for the USDCTokenPool, SiloedUSDCTokenPool, and USDCTokenPoolCCTPV2 contracts.
func ConfigureAllowedCallersChangeset(mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[ConfigureAllowedCallersInput] {
	return cldf.CreateChangeSet(configureAllowedCallersApply(mcmsRegistry), configureAllowedCallersVerify(mcmsRegistry))
}

func configureAllowedCallersApply(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ConfigureAllowedCallersInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input ConfigureAllowedCallersInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Build the Address and AllowedCallersByChain maps from per-chain inputs
		addressesByChain := make(map[uint64]common.Address)
		allowedCallersByChain := make(map[uint64][]erc20_lock_box.AllowedCallerConfigArgs)
		for _, perChainInput := range input.ChainInputs {

			erc20_lock_box_address, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType(erc20_lock_box.ContractType),
				Version:       erc20_lock_box.Version,
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			addressesByChain[perChainInput.ChainSelector] = erc20_lock_box_address
			allowedCallersByChain[perChainInput.ChainSelector] = perChainInput.AllowedCallers
		}

		// Execute the sequence with the combined input
		sequenceInput := sequences.ConfigureAllowedCallersSequenceInput{
			AddressesByChain:               addressesByChain,
			ConfigureAllowedCallersByChain: allowedCallersByChain,
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.ERC20LockboxConfigureAllowedCallersSequence, e.BlockChains, sequenceInput)
		if err != nil {
			return cldf.ChangesetOutput{}, err
		}

		batchOps = append(batchOps, report.Output.BatchOps...)
		reports = append(reports, report.ExecutionReports...)

		return changesets.NewOutputBuilder(e, mcmsRegistry).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(input.MCMS)
	}

}

func configureAllowedCallersVerify(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ConfigureAllowedCallersInput) error {
	return func(e cldf.Environment, input ConfigureAllowedCallersInput) error {
		if err := input.MCMS.Validate(); err != nil {
			return err
		}
		return nil
	}
}
