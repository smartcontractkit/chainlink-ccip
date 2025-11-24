package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"

	usdc_token_pool_proxy_ops "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/usdc_token_pool_proxy"
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
// and we do not want the constructor to revert. Similarly, on a new chain, such as Sonic or Linea, there will not
// be a V1 deployment of the USDC Token Pool, and only the CCTP V2 pool will be required.
// This changeset, as validation will attempt to find the CCTP V1 and CCTP V2 pools, will revert if both are not found.
// If that occurs, then the user must deploy a token pool contract separately and update the datastore before deploying the USDCTokenPoolProxy.
func DeployUSDCTokenPoolProxyChangeset(mcmsRegistry *changesets.MCMSReaderRegistry) cldf.ChangeSetV2[DeployUSDCTokenPoolProxyInput] {
	return cldf.CreateChangeSet(deployUSDCTokenPoolProxyApply(mcmsRegistry), deployUSDCTokenPoolProxyVerify(mcmsRegistry))
}

func deployUSDCTokenPoolProxyApply(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, DeployUSDCTokenPoolProxyInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input DeployUSDCTokenPoolProxyInput) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()

		// Execute the sequence for each chain input
		for _, perChainInput := range input.ChainInputs {
			// Find the chain family for the given chain selector
			chainFamily, err := chain_selectors.GetSelectorFamily(perChainInput.ChainSelector)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get chain family for selector %d: %w", perChainInput.ChainSelector, err)
			}
			reader, ok := mcmsRegistry.GetMCMSReader(chainFamily)
			if !ok {
				return cldf.ChangesetOutput{}, fmt.Errorf("no MCMSReader registered for chain family '%s'", chainFamily)
			}

			timelockRef, err := reader.GetTimelockRef(e, perChainInput.ChainSelector, input.MCMS)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to get timelock ref for chain %d: %w", perChainInput.ChainSelector, err)
			}

			routerAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType("Router"),
				Version:       semver.MustParse("1.2.0"),
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			tokenAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType("USDCToken"),
				Version:       semver.MustParse("1.0.0"),
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			// We actually don't need to check for the error here since if the pool is not deployed, the sequence will fail.
			// and it is possible for a chain that the USDCTokenPool for CCTP V1 is not deployed.
			cctpV1PoolAddress, _ := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType("USDCTokenPool"),
				Version:       semver.MustParse("1.6.4"),
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)

			cctpV2PoolAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType("USDCTokenPoolCCTPV2"),
				Version:       semver.MustParse("1.6.4"),
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)

			// If both of these are empty, then revert because there is neither a CCTP V1 OR a V2 pool deployed to utilize.
			if cctpV1PoolAddress == (common.Address{}) && cctpV2PoolAddress == (common.Address{}) {
				return cldf.ChangesetOutput{}, fmt.Errorf("no USDCTokenPool for CCTP V1 or CCTP V2 found for chain %d", perChainInput.ChainSelector)
			}

			sequenceInput := sequences.DeployUSDCTokenPoolProxySequenceInput{
				ChainSelector: perChainInput.ChainSelector,
				Token:         tokenAddress,
				PoolAddresses: usdc_token_pool_proxy_ops.PoolAddresses{
					LegacyCctpV1Pool: perChainInput.LegacyPoolAddress,
					CctpV1Pool:       cctpV1PoolAddress,
					CctpV2Pool:       cctpV2PoolAddress,
				},
				Router:      routerAddress,
				MCMSAddress: common.HexToAddress(timelockRef.Address),
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

func deployUSDCTokenPoolProxyVerify(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, DeployUSDCTokenPoolProxyInput) error {
	return func(e cldf.Environment, input DeployUSDCTokenPoolProxyInput) error {
		if err := input.MCMS.Validate(); err != nil {
			return err
		}
		return nil
	}
}
