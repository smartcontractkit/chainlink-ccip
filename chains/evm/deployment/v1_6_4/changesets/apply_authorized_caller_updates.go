package changesets

import (
	"github.com/ethereum/go-ethereum/common"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	authorized_caller_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/authorized_caller"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
)

type ApplyAuthorizedCallerUpdatesInput struct {
	ChainInputs []ApplyAuthorizedCallerUpdatesInputPerChain
	MCMS        mcms.Input
}

type ApplyAuthorizedCallerUpdatesInputPerChain struct {
	ChainSelector           uint64
	Address                 common.Address
	AuthorizedCallerUpdates authorized_caller_ops.AuthorizedCallerUpdateArgs
}

// Note: This changeset is tested in the usdc_token_pool_deploy_test.go file since usdc_token_pool implements the
// AuthorizedCallers library. Since there is no gobinding for the AuthorizedCallers library, it is being included here in its own
// file to allow for the use of the ApplyAuthorizedCallerUpdates operation on a variety of different contracts
// that implement the AuthorizedCallers interface (e.g. USDCTokenPool, SiloedUSDCTokenPool, etc.).
func ApplyAuthorizedCallerUpdatesChangeset(mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[ApplyAuthorizedCallerUpdatesInput] {
	return cldf.CreateChangeSet(applyAuthorizedCallerUpdatesApply(mcmsRegistry), applyAuthorizedCallerUpdatesVerify(mcmsRegistry))
}

func applyAuthorizedCallerUpdatesApply(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ApplyAuthorizedCallerUpdatesInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input ApplyAuthorizedCallerUpdatesInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		addressByChain := make(map[uint64]common.Address)
		authorizedCallerUpdatesByChain := make(map[uint64]authorized_caller_ops.AuthorizedCallerUpdateArgs)
		for _, perChainInput := range input.ChainInputs {
			addressByChain[perChainInput.ChainSelector] = perChainInput.Address
			authorizedCallerUpdatesByChain[perChainInput.ChainSelector] = perChainInput.AuthorizedCallerUpdates
		}

		sequenceInput := sequences.ApplyAuthorizedCallerUpdatesSequenceInput{
			Address:                        addressByChain,
			AuthorizedCallerUpdatesByChain: authorizedCallerUpdatesByChain,
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.ApplyAuthorizedCallerUpdatesSequence, e.BlockChains, sequenceInput)
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

func applyAuthorizedCallerUpdatesVerify(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, ApplyAuthorizedCallerUpdatesInput) error {
	return func(e cldf.Environment, input ApplyAuthorizedCallerUpdatesInput) error {
		if err := input.MCMS.Validate(); err != nil {
			return err
		}
		return nil
	}
}
