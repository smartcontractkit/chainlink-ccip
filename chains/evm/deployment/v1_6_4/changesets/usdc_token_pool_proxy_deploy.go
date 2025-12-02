package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"

	usdc_token_pool_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool"
	usdc_token_pool_cctp_v2_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_cctp_v2"
	usdc_token_pool_proxy_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_proxy"

	utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

type DeployUSDCTokenPoolProxyInput struct {
	ChainInputs []DeployUSDCTokenPoolProxyPerChainInput
	MCMS        mcms.Input
}

type DeployUSDCTokenPoolProxyPerChainInput struct {
	ChainSelector     uint64
	LegacyPoolAddress common.Address
}

// This changeset is used to deploy the USDCTokenPoolProxy contract on a given chain.
// Note: Since this may be deployed on a chain that already has a USDC Token Pool contract deployed,
// the legacy pool address is the only required address to be provided in the input. This is because on a chain such as
// Ethereum Mainnet, which is yet to be updated to CCTP V2, there will only be a V1 deployment of the USDC Token Pool.
func DeployUSDCTokenPoolProxyChangeset() cldf.ChangeSetV2[DeployUSDCTokenPoolProxyInput] {
	return cldf.CreateChangeSet(deployUSDCTokenPoolProxyApply(), deployUSDCTokenPoolProxyVerify())
}

func deployUSDCTokenPoolProxyApply() func(cldf.Environment, DeployUSDCTokenPoolProxyInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input DeployUSDCTokenPoolProxyInput) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()

		// Execute the sequence for each chain input
		for _, perChainInput := range input.ChainInputs {
			// Get the timelock address from the datastore based on the MCMS qualifier and chain selector
			// Without the qualifier, the datastore will sometimes throw an error when fetching the address due to the
			// datastore containing >1 address with the same type and version.
			timeLockAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:      datastore.ContractType(utils.RBACTimelock),
				Version:   semver.MustParse(RBACTimelockVersion),
				Qualifier: input.MCMS.Qualifier,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get time lock address for chain %d: %w", perChainInput.ChainSelector, err)
			}

			routerAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(RouterContractType),
				Version: semver.MustParse(RouterVersion),
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			tokenAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(USDCTokenContractType),
				Version: semver.MustParse(USDCTokenVersion),
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			// It is not necessary to check for the error here since the contract constructor does not require any of the
			// pool addresses to be set, and can be modified later. This also allows for parallel testing/deployment of the USDCTokenPoolProxy
			// and the USDCTokenPool contracts on various chains.
			cctpV1PoolAddress, _ := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(usdc_token_pool_ops.ContractType),
				Version: usdc_token_pool_ops.Version,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)

			cctpV2PoolAddress, _ := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(usdc_token_pool_cctp_v2_ops.ContractType),
				Version: usdc_token_pool_cctp_v2_ops.Version,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)

			sequenceInput := sequences.DeployUSDCTokenPoolProxySequenceInput{
				ChainSelector: perChainInput.ChainSelector,
				Token:         tokenAddress,
				PoolAddresses: usdc_token_pool_proxy_ops.PoolAddresses{
					LegacyCctpV1Pool: perChainInput.LegacyPoolAddress,
					CctpV1Pool:       cctpV1PoolAddress,
					CctpV2Pool:       cctpV2PoolAddress,
				},
				Router:      routerAddress,
				MCMSAddress: timeLockAddress,
			}

			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.DeployUSDCTokenPoolProxySequence, e.BlockChains, sequenceInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy USDCTokenPoolProxy on chain %d: %w", perChainInput.ChainSelector, err)
			}
			// Add the execution reports to the reports slice
			reports = append(reports, report.ExecutionReports...)
			// Add the addresses to the datastore
			for _, r := range report.Output.Addresses {
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
				}
			}
		}

		// Return the changeset output
		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithDataStore(ds).
			Build(input.MCMS)
	}
}

func deployUSDCTokenPoolProxyVerify() func(cldf.Environment, DeployUSDCTokenPoolProxyInput) error {
	return func(e cldf.Environment, input DeployUSDCTokenPoolProxyInput) error {
		return nil
	}
}
