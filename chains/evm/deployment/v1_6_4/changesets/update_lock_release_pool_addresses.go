package changesets

import (
	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
	mcms_types "github.com/smartcontractkit/mcms/types"

	usdc_token_pool_proxy_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_proxy"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"

	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	datastore "github.com/smartcontractkit/chainlink-deployments-framework/datastore"
)

type UpdateLockReleasePoolAddressesInput struct {
	ChainInputs []UpdateLockReleasePoolAddressesPerChainInput
	MCMS        mcms.Input
}

type UpdateLockReleasePoolAddressesPerChainInput struct {
	ChainSelector        uint64
	LockReleasePoolAddrs usdc_token_pool_proxy_ops.UpdateLockReleasePoolAddressesArgs
}

// This changeset is used to update the lock release pools for a given token in the USDCTokenPoolProxy contract.
// On a given source chain, there will be different SiloedUSDCTokenPool contracts for each destination chain.
// This changeset is used to update the lock release pool addresses for a given source chain and an associated destination chain.
func UpdateLockReleasePoolAddressesChangeset(mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[UpdateLockReleasePoolAddressesInput] {
	return cldf.CreateChangeSet(updateLockReleasePoolAddressesApply(mcmsRegistry), updateLockReleasePoolAddressesVerify(mcmsRegistry))
}

func updateLockReleasePoolAddressesApply(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateLockReleasePoolAddressesInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input UpdateLockReleasePoolAddressesInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		addressByChain := make(map[uint64]common.Address)
		lockReleasePoolAddressesByChain := make(map[uint64]usdc_token_pool_proxy_ops.UpdateLockReleasePoolAddressesArgs)
		for _, perChainInput := range input.ChainInputs {
			address, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType(usdc_token_pool_proxy_ops.ContractType),
				Version:       semver.MustParse("1.6.4"),
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			addressByChain[perChainInput.ChainSelector] = address
			lockReleasePoolAddressesByChain[perChainInput.ChainSelector] = perChainInput.LockReleasePoolAddrs
		}

		sequenceInput := sequences.UpdateLockReleasePoolAddressesSequenceInput{
			Address:                         addressByChain,
			LockReleasePoolAddressesByChain: lockReleasePoolAddressesByChain,
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.USDCTokenPoolProxyUpdateLockReleasePoolAddressesSequence, e.BlockChains, sequenceInput)
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

func updateLockReleasePoolAddressesVerify(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, UpdateLockReleasePoolAddressesInput) error {
	return func(e cldf.Environment, input UpdateLockReleasePoolAddressesInput) error {
		if err := input.MCMS.Validate(); err != nil {
			return err
		}
		return nil
	}
}
