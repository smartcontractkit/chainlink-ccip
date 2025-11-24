package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
)

type USDCTokenPoolCCTPV2DeployInputPerChain struct {
	ChainSelector  uint64
	TokenMessenger common.Address
	Token          common.Address
	Allowlist      []common.Address
}

type USDCTokenPoolCCTPV2DeployInput struct {
	ChainInputs []USDCTokenPoolCCTPV2DeployInputPerChain
	MCMS        mcms.Input
}

// This changeset is used to deploy the USDCTokenPoolCCTPV2 contract on a given chain.
// Note: In addition to deploying the USDCTokenPoolCCTPV2 contract, this changeset will also deploy the CCTPMessageTransmitterProxy contract,
// configure the allowed callers for the CCTPMessageTransmitterProxy contract, and then begin the ownership transfer to MCMS.
// A separate changeset will be used to accept ownership of the USDCTokenPoolCCTPV2 contract.
func USDCTokenPoolCCTPV2DeployChangeset(mcmsRegistry *changesets.MCMSReaderRegistry) deployment.ChangeSetV2[USDCTokenPoolCCTPV2DeployInput] {
	return cldf.CreateChangeSet(usdcTokenPoolCCTPV2DeployApply(mcmsRegistry), usdcTokenPoolCCTPV2DeployVerify(mcmsRegistry))
}

func usdcTokenPoolCCTPV2DeployApply(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, USDCTokenPoolCCTPV2DeployInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input USDCTokenPoolCCTPV2DeployInput) (cldf.ChangesetOutput, error) {
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

			routerAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType("Router"),
				Version:       semver.MustParse("1.2.0"),
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

			var TokenAddress common.Address
			// If the token address is not found in the datastore, perhaps because it is a new chain
			// then use the token address from the input
			retrievedTokenAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:          datastore.ContractType("USDCToken"),
				Version:       semver.MustParse("1.0.0"),
				ChainSelector: perChainInput.ChainSelector,
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			// If the error is not nil, then check if the token address was provided in the input, and if so use that,
			// otherwise revert because the token address is required.
			if err != nil {
				if perChainInput.Token == (common.Address{}) {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get token address for chain %d: %w", perChainInput.ChainSelector, err)
				}
				TokenAddress = retrievedTokenAddress
			} else if err == nil {
				TokenAddress = retrievedTokenAddress
			}

			sequenceInput := sequences.USDCTokenPoolCCTPV2DeploySequenceInput{
				ChainSelector:  perChainInput.ChainSelector,
				TokenMessenger: perChainInput.TokenMessenger,
				Token:          TokenAddress,
				Allowlist:      perChainInput.Allowlist,
				RMNProxy:       rmnAddress,
				Router:         routerAddress,
				MCMSAddress:    common.HexToAddress(timelockRef.Address),
			}
			report, err := cldf_ops.ExecuteSequence(e.OperationsBundle, sequences.USDCTokenPoolCCTPV2DeploySequence, e.BlockChains, sequenceInput)
			if err != nil {
				return cldf.ChangesetOutput{}, fmt.Errorf("failed to deploy USDCTokenPoolCCTPV2 on chain %d: %w", perChainInput.ChainSelector, err)
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

func usdcTokenPoolCCTPV2DeployVerify(mcmsRegistry *changesets.MCMSReaderRegistry) func(cldf.Environment, USDCTokenPoolCCTPV2DeployInput) error {
	return func(e cldf.Environment, input USDCTokenPoolCCTPV2DeployInput) error {
		if err := input.MCMS.Validate(); err != nil {
			return err
		}
		return nil
	}
}
