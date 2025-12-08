package changesets

import (
	"fmt"

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
	Qualifier      string
	AllowedCallers []erc20_lock_box.AllowedCallerConfigArgs
}

// This changeset is used to configure the allowed callers for a given token in the ERC20Lockbox contract.
// It is different from the apply_authorized_caller_updates changeset which is used for the USDCTokenPool, SiloedUSDCTokenPool, and USDCTokenPoolCCTPV2 contracts.
func ConfigureAllowedCallersChangeset() cldf.ChangeSetV2[ConfigureAllowedCallersInput] {
	return cldf.CreateChangeSet(configureAllowedCallersApply(), configureAllowedCallersVerify())
}

func configureAllowedCallersApply() func(cldf.Environment, ConfigureAllowedCallersInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input ConfigureAllowedCallersInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Build the Address and AllowedCallersByChain maps from per-chain inputs
		addressesByChain := make(map[uint64]common.Address)
		allowedCallersByChain := make(map[uint64][]erc20_lock_box.AllowedCallerConfigArgs)
		for _, perChainInput := range input.ChainInputs {

			// Find the ERC20Lockbox contract address, using the qualifier provided in the input to differentiate between
			// the IOwnable and future IAccessControl lockboxes.
			erc20LockBoxAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:      datastore.ContractType(erc20_lock_box.ContractType),
				Version:   erc20_lock_box.Version,
				Qualifier: perChainInput.Qualifier,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			addressesByChain[perChainInput.ChainSelector] = erc20LockBoxAddress
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

		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithBatchOps(batchOps).
			Build(input.MCMS)
	}

}

func configureAllowedCallersVerify() func(cldf.Environment, ConfigureAllowedCallersInput) error {
	return func(e cldf.Environment, input ConfigureAllowedCallersInput) error {
		for _, perChainInput := range input.ChainInputs {
			if exists := e.BlockChains.Exists(perChainInput.ChainSelector); !exists {
				return fmt.Errorf("chain with selector %d does not exist", perChainInput.ChainSelector)
			}

			for _, allowedCaller := range perChainInput.AllowedCallers {
				if allowedCaller.Caller == (common.Address{}) {
					return fmt.Errorf("caller address cannot be zero for chain selector %d", perChainInput.ChainSelector)
				}
			}

		}
		return nil
	}
}
