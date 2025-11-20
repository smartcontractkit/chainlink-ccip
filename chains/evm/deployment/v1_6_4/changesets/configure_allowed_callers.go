package changesets

import (
	"github.com/ethereum/go-ethereum/common"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/erc20_lock_box"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
)

type ConfigureAllowedCallersInput struct {
	ChainInputs []ConfigureAllowedCallersPerChainInput
	MCMS        mcms.Input
}

type ConfigureAllowedCallersPerChainInput struct {
	ChainSelector  uint64
	Address        common.Address
	AllowedCallers []erc20_lock_box.AllowedCallerConfigArgs
}

func ConfigureAllowedCallersChangeset(mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[ConfigureAllowedCallersInput] {
	return cldf.CreateChangeSet(configureAllowedCallersApply(mcmsRegistry), configureAllowedCallersVerify(mcmsRegistry))
}

func configureAllowedCallersApply(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ConfigureAllowedCallersInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input ConfigureAllowedCallersInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		// Build the Address and AllowedCallersByChain maps from per-chain inputs
		addressByChain := make(map[uint64]common.Address)
		allowedCallersByChain := make(map[uint64][]erc20_lock_box.AllowedCallerConfigArgs)
		for _, perChainInput := range input.ChainInputs {
			addressByChain[perChainInput.ChainSelector] = perChainInput.Address
			allowedCallersByChain[perChainInput.ChainSelector] = perChainInput.AllowedCallers
		}

		// Execute the sequence with the combined input
		sequenceInput := sequences.ConfigureAllowedCallersSequenceInput{
			Address:                        addressByChain,
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
