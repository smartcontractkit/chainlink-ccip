package changesets

import (
	"github.com/ethereum/go-ethereum/common"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	usdc_token_pool_proxy_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"

	semver "github.com/Masterminds/semver/v3"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type UpdateLockOrBurnMechanismInput struct {
	ChainInputs []UpdateLockOrBurnMechanismPerChainInput
	MCMS        mcms.Input
}

type UpdateLockOrBurnMechanismPerChainInput struct {
	ChainSelector uint64
	Mechanisms    usdc_token_pool_proxy_ops.UpdateLockOrBurnMechanismsArgs
}

// This changeset is use to update the mechanism to be used for outgoing USDC messages going through the USDCTokenPoolProxy contract.
// It should only be used for the USDCTokenPoolProxy contract.
func UpdateLockOrBurnMechanismChangeset(mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[UpdateLockOrBurnMechanismInput] {
	return cldf.CreateChangeSet(updateLockOrBurnMechanismApply(mcmsRegistry), updateLockOrBurnMechanismVerify(mcmsRegistry))
}

func updateLockOrBurnMechanismApply(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateLockOrBurnMechanismInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input UpdateLockOrBurnMechanismInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		addressByChain := make(map[uint64]common.Address)
		mechanismsByChain := make(map[uint64]usdc_token_pool_proxy_ops.UpdateLockOrBurnMechanismsArgs)
		for _, perChainInput := range input.ChainInputs {

			// Find the USDCTokenPoolProxy address for the given chain selector
			address, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType(usdc_token_pool_proxy_ops.ContractType),
				Version:       semver.MustParse("1.6.4"),
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			addressByChain[perChainInput.ChainSelector] = address
			mechanismsByChain[perChainInput.ChainSelector] = perChainInput.Mechanisms
		}

		sequenceInput := sequences.UpdateLockOrBurnMechanismSequenceInput{
			Address:    addressByChain,
			Mechanisms: mechanismsByChain,
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.USDCTokenPoolProxyUpdateLockOrBurnMechanismSequence, e.BlockChains, sequenceInput)
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

func updateLockOrBurnMechanismVerify(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateLockOrBurnMechanismInput) error {
	return func(e cldf.Environment, input UpdateLockOrBurnMechanismInput) error {
		if err := input.MCMS.Validate(); err != nil {
			return err
		}
		return nil
	}
}
