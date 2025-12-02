package changesets

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/common"
	evm_datastore_utils "github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	datastore_utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils/datastore"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"

	utils "github.com/smartcontractkit/chainlink-ccip/deployment/utils"
)

const (
	// The CCTPMessageTransmitterProxy contract is deployed with version v1_6_2 and is used to send messages to the CCTP V1 Token Pool.
	// It is not necessary to deploy the CCTPMessageTransmitterProxy contract for CCTP V1 as V1 is being deprecated in 2026 and
	// as such no new chains should be added that support CCTP V1.
	CCTPMessageTransmitterProxyContractType = "CCTPMessageTransmitterProxy"
	CCTPMessageTransmitterProxyVersion      = "1.6.2"
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
func USDCTokenPoolDeployChangeset() deployment.ChangeSetV2[USDCTokenPoolDeployInput] {
	return cldf.CreateChangeSet(usdcTokenPoolDeployApply(), usdcTokenPoolDeployVerify())
}

func usdcTokenPoolDeployApply() func(cldf.Environment, USDCTokenPoolDeployInput) (cldf.ChangesetOutput, error) {
	return func(e cldf.Environment, input USDCTokenPoolDeployInput) (cldf.ChangesetOutput, error) {
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

			tokenAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(USDCTokenContractType),
				Version: semver.MustParse(USDCTokenVersion),
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			// Since this deploy a CCTP V1 USDC Token Pool, it is assumed that the CCTPMessageTransmitterProxy is already deployed.
			// If it is not, a separate changeset will be used to deploy it. However, you should never be in a position
			// where you need to deploy the CCTPMessageTransmitterProxy for CCTP V1 as V1 is being deprecated in 2026 and
			// as such no new chains should be added that support CCTP V1.
			cctpMessageTransmitterProxyAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(CCTPMessageTransmitterProxyContractType),
				Version: semver.MustParse(CCTPMessageTransmitterProxyVersion),
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			rmnAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(RMNProxyContractType),
				Version: semver.MustParse(RMNProxyVersion),
			}, perChainInput.ChainSelector, evm_datastore_utils.ToEVMAddress)
			if err != nil {
				return cldf.ChangesetOutput{}, err
			}

			routerAddress, err := datastore_utils.FindAndFormatRef(e.DataStore, datastore.AddressRef{
				Type:    datastore.ContractType(RouterContractType),
				Version: semver.MustParse(RouterVersion),
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
				MCMSAddress:                 timeLockAddress,
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

func usdcTokenPoolDeployVerify() func(cldf.Environment, USDCTokenPoolDeployInput) error {
	return func(e cldf.Environment, input USDCTokenPoolDeployInput) error {
		return nil
	}
}
