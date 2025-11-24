package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type USDCTokenPoolDeployInputPerChain struct {
	ChainSelector  uint64
	TokenMessenger common.Address
	Allowlist      []common.Address
}

type USDCTokenPoolDeployInput struct {
	ChainInputs []USDCTokenPoolDeployInputPerChain
	MCMS        mcms.Input
}

// This changeset is used to deploy the USDCTokenPool contract on a given chain.
// Note: Unlike the changset for the CCTP V2 Token Pool, this changeset will NOT deploy the CCTPMessageTransmitterProxy contract.
// That must be performed with a separate changeset in v1_6_2 for the CCTPMessageTransmitterProxy contract.
func USDCTokenPoolDeployChangeset(mcmsRegistry *changesets.MCMSReaderRegistry) deployment.ChangeSetV2[USDCTokenPoolDeployInput] {
	return cldf.CreateChangeSet(usdcTokenPoolDeployApply(mcmsRegistry), usdcTokenPoolDeployVerify(mcmsRegistry))
}

func usdcTokenPoolDeployApply(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, USDCTokenPoolDeployInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input USDCTokenPoolDeployInput) (cldf.ChangesetOutput, error) {
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

			tokenAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType("USDCToken"),
				Version:       semver.MustParse("1.0.0"),
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			// Since this deploy a CCTP V1 USDC Token Pool, it is assumed that the CCTPMessageTransmitterProxy is already deployed.
			// If it is not, a separate changeset will be used to deploy it. However, you should never be in a position
			// where you need to deploy the CCTPMessageTransmitterProxy for CCTP V1 as V1 is being deprecated in 2026 and
			// as such no new chains should be added that support CCTP V1.
			cctpMessageTransmitterProxyAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType("CCTPMessageTransmitterProxy"),
				Version:       semver.MustParse("1.6.2"),
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			rmnAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType("RMN"),
				Version:       semver.MustParse("1.5.0"),
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			routerAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType("Router"),
				Version:       semver.MustParse("1.2.0"),
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			sequenceInput := sequences.USDCTokenPoolDeploySequenceInput{
				ChainSelector:               perChainInput.ChainSelector,
				TokenMessenger:              perChainInput.TokenMessenger,
				CCTPMessageTransmitterProxy: cctpMessageTransmitterProxyAddress,
				Token:                       tokenAddress,
				Allowlist:                   perChainInput.Allowlist,
				RMNProxy:                    rmnAddress,
				Router:                      routerAddress,
				SupportedUSDCVersion:        0, // For CCTP V1 the version will always be zero since it is the first version
				MCMSAddress:                 common.HexToAddress(timelockRef.Address),
			}

			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.USDCTokenPoolDeploySequence, e.BlockChains, sequenceInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy USDCTokenPool on chain %d: %w", perChainInput.ChainSelector, err)
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

func usdcTokenPoolDeployVerify(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, USDCTokenPoolDeployInput) error {
	return func(e cldf.Environment, input USDCTokenPoolDeployInput) error {
		if err := input.MCMS.Validate(); err != nil {
			return err
		}
		return nil
	}
}
