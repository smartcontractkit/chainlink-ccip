package changesets

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-ccip/chains/evm/deployment/v1_6_4/sequences"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/changesets"
	"github.com/smartcontractkit/chainlink-ccip/deployment/utils/mcms"
	"github.com/smartcontractkit/chainlink-deployments-framework/datastore"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	cldf_ops "github.com/smartcontractkit/chainlink-deployments-framework/operations"
)

type USDCTokenPoolDeployInputPerChain struct {
	ChainSelector               uint64
	TokenMessenger              common.Address
	CCTPMessageTransmitterProxy common.Address
	Token                       common.Address
	Allowlist                   []common.Address
	RMNProxy                    common.Address
	Router                      common.Address
	SupportedUSDCVersion        uint32
}

type USDCTokenPoolDeployInput struct {
	ChainInputs []USDCTokenPoolDeployInputPerChain
	MCMS        mcms.Input
}

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
			sequenceInput := sequences.USDCTokenPoolDeploySequenceInput{
				ChainSelector:               perChainInput.ChainSelector,
				TokenMessenger:              perChainInput.TokenMessenger,
				CCTPMessageTransmitterProxy: perChainInput.CCTPMessageTransmitterProxy,
				Token:                       perChainInput.Token,
				Allowlist:                   perChainInput.Allowlist,
				RMNProxy:                    perChainInput.RMNProxy,
				Router:                      perChainInput.Router,
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
