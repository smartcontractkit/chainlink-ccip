package changesets

import (
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

type UpdatePoolAddressesInput struct {
	ChainInputs []UpdatePoolAddressesPerChainInput
	MCMS        mcms.Input
}

type UpdatePoolAddressesPerChainInput struct {
	ChainSelector uint64
	PoolAddresses usdc_token_pool_proxy_ops.PoolAddresses
}

func UpdatePoolAddressesChangeset() cldf.ChangeSetV2[UpdatePoolAddressesInput] {
	return cldf.CreateChangeSet(updatePoolAddressesApply(), updatePoolAddressesVerify())
}

func updatePoolAddressesApply() func(cldf.Environment, UpdatePoolAddressesInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input UpdatePoolAddressesInput) (cldf.ChangesetOutput, error) {
		batchOps := make([]mcms_types.BatchOperation, 0)
		reports := make([]cldf_ops.Report[any, any], 0)

		addressByChain := make(map[uint64]common.Address)
		poolAddressesByChain := make(map[uint64]usdc_token_pool_proxy_ops.PoolAddresses)
		for _, perChainInput := range input.ChainInputs {
			address, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(usdc_token_pool_proxy_ops.ContractType),
				Version: usdc_token_pool_proxy_ops.Version,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}
			addressByChain[perChainInput.ChainSelector] = address
			poolAddressesByChain[perChainInput.ChainSelector] = perChainInput.PoolAddresses
		}

		sequenceInput := sequences.UpdatePoolAddressesSequenceInput{
			Address:              addressByChain,
			PoolAddressesByChain: poolAddressesByChain,
		}

		report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.USDCTokenPoolProxyUpdatePoolAddressesSequence, e.BlockChains, sequenceInput)
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

func updatePoolAddressesVerify() func(cldf.Environment, UpdatePoolAddressesInput) error {
	return func(e cldf.Environment, input UpdatePoolAddressesInput) error {
		return nil
	}
}
