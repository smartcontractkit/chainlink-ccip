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

	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"

	utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
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
func USDCTokenPoolCCTPV2DeployChangeset() deployment.ChangeSetV2[USDCTokenPoolCCTPV2DeployInput] {
	return cldf.CreateChangeSet(usdcTokenPoolCCTPV2DeployApply(), usdcTokenPoolCCTPV2DeployVerify())
}

func usdcTokenPoolCCTPV2DeployApply() func(cldf.Environment, USDCTokenPoolCCTPV2DeployInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input USDCTokenPoolCCTPV2DeployInput) (cldf.ChangesetOutput, error) {
		reports := make([]cldf_ops.Report[any, any], 0)
		ds := datastore.NewMemoryDataStore()

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

			// Get the router address from the datastore based on the chain selector.
			routerAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(RouterContractType),
				Version: semver.MustParse(RouterVersion),
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			// Get the RMN address from the datastore based on the chain selector.
			rmnAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(RMNProxyContractType),
				Version: semver.MustParse(RMNProxyVersion),
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			var TokenAddress common.Address
			// If the token address is not found in the datastore, perhaps because it is a new chain
			// then use the token address from the input
			retrievedTokenAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(USDCTokenContractType),
				Version: semver.MustParse(USDCTokenVersion),
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			// If the error is not nil, then check if the token address was provided in the input, and if so use that,
			// otherwise revert because the token address is required.
			if err != nil {
				if perChainInput.Token == (common.Address{}) {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to get token address for chain %d: %w", perChainInput.ChainSelector, err)
				}
				TokenAddress = perChainInput.Token

				// If the token address is provided in the input, add it to the datastore so that it can be used in the
				// future without having to be provided in the input again.
				err = ds.Addresses().Add(datastore.AddressRef{
					Type:          datastore.ContractType(USDCTokenContractType),
					Version:       semver.MustParse(USDCTokenVersion),
					Address:       perChainInput.Token.Hex(),
					ChainSelector: perChainInput.ChainSelector,
				})
				if err != nil {
					return cldf.ChangesetOutput{}, fmt.Errorf("failed to add USDC token address to datastore for chain %d: %w", perChainInput.ChainSelector, err)
				}

			} else {
				TokenAddress = retrievedTokenAddress
			}

			sequenceInput := sequences.USDCTokenPoolCCTPV2DeploySequenceInput{
				ChainSelector:  perChainInput.ChainSelector,
				TokenMessenger: perChainInput.TokenMessenger,
				Token:          TokenAddress,
				Allowlist:      perChainInput.Allowlist,
				RMNProxy:       rmnAddress,
				Router:         routerAddress,
				MCMSAddress:    timeLockAddress,
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

func usdcTokenPoolCCTPV2DeployVerify() func(cldf.Environment, USDCTokenPoolCCTPV2DeployInput) error {
	return func(e cldf.Environment, input USDCTokenPoolCCTPV2DeployInput) error {
		return nil
	}
}
