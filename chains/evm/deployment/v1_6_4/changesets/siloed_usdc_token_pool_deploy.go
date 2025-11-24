package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	chain_selectors "github.com/smartcontractkit/chain-selectors"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	router "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_2_0/operations/router"
	rmn_proxy "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_5_0/operations/rmn"
	erc20_lock_box "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/operations/erc20_lock_box"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

type SiloedUSDCTokenPoolDeployInputPerChain struct {
	ChainSelector uint64
	Allowlist     []common.Address
}

type SiloedUSDCTokenPoolDeployInput struct {
	ChainInputs []SiloedUSDCTokenPoolDeployInputPerChain
	MCMS        mcms.Input
}

func SiloedUSDCTokenPoolDeployChangeset(mcmsRegistry *changesets.MCMSReaderRegistry) deployment.ChangeSetV2[SiloedUSDCTokenPoolDeployInput] {
	return cldf.CreateChangeSet(siloedUSDCTokenPoolDeployApply(mcmsRegistry), siloedUSDCTokenPoolDeployVerify(mcmsRegistry))
}

func siloedUSDCTokenPoolDeployApply(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, SiloedUSDCTokenPoolDeployInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input SiloedUSDCTokenPoolDeployInput) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()

		for _, perChainInput := range input.ChainInputs {

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

			erc20LockboxAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType(erc20_lock_box.ContractType),
				Version:       erc20_lock_box.Version,
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			routerAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType(router.ContractType),
				Version:       router.Version,
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			rmnProxyAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType(rmn_proxy.ContractType),
				Version:       semver.MustParse("1.5.0"),
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

			sequenceInput := sequences.SiloedUSDCTokenPoolDeploySequenceInput{
				ChainSelector: perChainInput.ChainSelector,
				Token:         tokenAddress,
				Allowlist:     perChainInput.Allowlist,
				RMNProxy:      rmnProxyAddress,
				Router:        routerAddress,
				LockBox:       erc20LockboxAddress,
				MCMSAddress:   common.HexToAddress(timelockRef.Address),
			}
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.SiloedUSDCTokenPoolDeploySequence, e.BlockChains, sequenceInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy SiloedUSDCTokenPool on chain %d: %w", perChainInput.ChainSelector, err)
			}
			for _, r := range report.Output.Addresses {
				if err := ds.Addresses().Add(r); err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add %s %s with address %s on chain with selector %d to datastore: %w", r.Type, r.Version, r.Address, r.ChainSelector, err)
				}
			}
			reports = append(reports, report.ExecutionReports...)
		}
		return changesets.NewOutputBuilder(e, nil).
			WithReports(reports).
			WithDataStore(ds).
			Build(input.MCMS)
	}
}

func siloedUSDCTokenPoolDeployVerify(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, SiloedUSDCTokenPoolDeployInput) error {
	return func(e cldf.Environment, input SiloedUSDCTokenPoolDeployInput) error {
		return nil
	}
}
